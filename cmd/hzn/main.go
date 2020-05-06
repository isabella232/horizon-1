package main

import (
	"bytes"
	"context"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/horizon/pkg/control"
	"github.com/hashicorp/horizon/pkg/hub"
	"github.com/hashicorp/horizon/pkg/pb"
	"github.com/hashicorp/horizon/pkg/tlsmanage"
	"github.com/hashicorp/vault/api"
	"github.com/jinzhu/gorm"
	"github.com/mitchellh/cli"
	"google.golang.org/grpc"
)

var (
	sha1ver   string // sha1 revision used to build the program
	buildTime string // when the executable was built
)

func main() {
	var ver string
	if sha1ver == "" {
		ver = "unknown"
	} else {
		ver = sha1ver[:10] + "-" + buildTime
	}

	c := cli.NewCLI("hzn", ver)
	c.Args = os.Args[1:]
	c.Commands = map[string]cli.CommandFactory{
		"control": controlFactory,
		"hub":     hubFactory,
		"migrate": func() (cli.Command, error) {
			return &migrateRunner{}, nil
		},
	}

	fmt.Printf("hzn: %s\n", ver)

	exitStatus, err := c.Run()
	if err != nil {
		log.Println(err)
	}

	os.Exit(exitStatus)

}

func controlFactory() (cli.Command, error) {
	return &controlServer{}, nil
}

type migrateRunner struct{}

func (m *migrateRunner) Help() string {
	return "run any migrations"
}

func (m *migrateRunner) Synopsis() string {
	return "run any migrations"
}

func (mr *migrateRunner) Run(args []string) int {
	url := os.Getenv("DATABASE_URL")
	if url == "" {
		log.Fatal("no DATABASE_URL provided")
	}

	m, err := migrate.New("file:///migrations", url)
	if err != nil {
		log.Fatal(err)
	}

	err = m.Up()
	if err != nil {
		log.Fatal(err)
	}

	return 0
}

func hubFactory() (cli.Command, error) {
	return &hubRunner{}, nil
}

type controlServer struct{}

func (c *controlServer) Help() string {
	return "Start a control server"
}

func (c *controlServer) Synopsis() string {
	return "Start a control server"
}

func (c *controlServer) Run(args []string) int {
	L := hclog.L()

	vcfg := api.DefaultConfig()

	vc, err := api.NewClient(vcfg)
	if err != nil {
		log.Fatal(err)
	}

	// If we have token AND this is kubernetes, then let's try to get a token
	if vc.Token() == "" {
		f, err := os.Open("/var/run/secrets/kubernetes.io/serviceaccount/token")
		if err == nil {
			L.Info("attempting to login to vault via kubernetes auth")

			data, err := ioutil.ReadAll(f)
			if err != nil {
				log.Fatal(err)
			}

			f.Close()

			sec, err := vc.Logical().Write("auth/kubernetes/login", map[string]interface{}{
				"role": "horizon",
				"jwt":  string(bytes.TrimSpace(data)),
			})
			if err != nil {
				log.Fatal(err)
			}

			if sec == nil {
				log.Fatal("unable to login to get token")
			}

			vc.SetToken(sec.Auth.ClientToken)

			L.Info("retrieved token from vault", "accessor", sec.Auth.Accessor)

			go func() {
				tic := time.NewTicker(time.Hour)
				for {
					<-tic.C
					vc.Auth().Token().RenewSelf(86400)
				}
			}()
		}
	}

	url := os.Getenv("DATABASE_URL")
	if url == "" {
		log.Fatal("no DATABASE_URL provided")
	}

	db, err := gorm.Open("postgres", url)
	if err != nil {
		log.Fatal(err)
	}

	sess := session.New()

	bucket := os.Getenv("S3_BUCKET")
	if bucket == "" {
		log.Fatal("S3_BUCKET not set")
	}

	domain := os.Getenv("HUB_DOMAIN")
	if domain == "" {
		log.Fatal("missing HUB_DOMAIN")
	}

	staging := os.Getenv("LETSENCRYPT_STAGING") != ""

	tlsmgr, err := tlsmanage.NewManager(tlsmanage.ManagerConfig{
		L:           L,
		Domain:      domain,
		VaultClient: vc,
		Staging:     staging,
	})
	if err != nil {
		log.Fatal(err)
	}

	zoneId := os.Getenv("ZONE_ID")
	if zoneId == "" {
		log.Fatal("missing ZONE_ID")
	}

	err = tlsmgr.SetupRoute53(sess, zoneId)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	cert, key, err := tlsmgr.HubMaterial(ctx)
	if err != nil {
		log.Fatal(err)
	}

	regTok := os.Getenv("REGISTER_TOKEN")
	if regTok == "" {
		log.Fatal("missing REGISTER_TOKEN")
	}

	opsTok := os.Getenv("OPS_TOKEN")
	if opsTok == "" {
		log.Fatal("missing OPS_TOKEN")
	}

	dynamoTable := os.Getenv("DYNAMO_TABLE")
	if dynamoTable == "" {
		log.Fatal("missing DYNAMO_TABLE")
	}

	asnDB := os.Getenv("ASN_DB_PATH")

	hubAccess := os.Getenv("HUB_ACCESS_KEY")
	hubSecret := os.Getenv("HUB_SECRET_KEY")

	port := os.Getenv("PORT")

	s, err := control.NewServer(control.ServerConfig{
		DB: db,

		RegisterToken: regTok,
		OpsToken:      opsTok,

		VaultClient: vc,
		VaultPath:   "hzn-k1",
		KeyId:       "k1",

		AwsSession: sess,
		Bucket:     bucket,
		LockTable:  dynamoTable,

		ASNDB: asnDB,

		HubAccessKey: hubAccess,
		HubSecretKey: hubSecret,
	})
	if err != nil {
		log.Fatal(err)
	}

	s.SetHubTLS(cert, key)

	gs := grpc.NewServer()
	pb.RegisterControlServicesServer(gs, s)
	pb.RegisterControlManagementServer(gs, s)
	pb.RegisterFlowTopReporterServer(gs, s)

	tlsCert, err := tlsmgr.Certificate()
	if err != nil {
		log.Fatal(err)
	}

	var lcfg tls.Config
	lcfg.Certificates = []tls.Certificate{tlsCert}

	hs := &http.Server{
		TLSConfig: &lcfg,
		Addr:      ":" + port,
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.ProtoMajor == 2 &&
				strings.HasPrefix(r.Header.Get("Content-Type"), "application/grpc") {
				gs.ServeHTTP(w, r)
			} else {
				s.ServeHTTP(w, r)
			}
		}),
		ErrorLog: L.StandardLogger(&hclog.StandardLoggerOptions{
			InferLevels: true,
		}),
	}

	err = hs.ListenAndServeTLS("", "")
	if err != nil {
		log.Fatal(err)
	}

	return 0
}

type hubRunner struct{}

func (h *hubRunner) Help() string {
	return "Start a hub"
}

func (h *hubRunner) Synopsis() string {
	return "Start a hub"
}

func (h *hubRunner) Run(args []string) int {
	L := hclog.L().Named("hub")

	if os.Getenv("DEBUG") != "" {
		L.SetLevel(hclog.Trace)
	}

	token := os.Getenv("TOKEN")
	if token == "" {
		log.Fatal("missing TOKEN")
	}

	addr := os.Getenv("CONTROL_ADDR")
	if addr == "" {
		log.Fatal("missing ADDR")
	}

	port := os.Getenv("PORT")
	if port == "" {
		L.Info("defaulting port to 443")
		port = "443"
	}

	httpPort := os.Getenv("HTTP_PORT")

	ctx := context.Background()

	sid := os.Getenv("STABLE_ID")
	if sid == "" {
		log.Fatal("missing STABLE_ID")
	}

	webNamespace := os.Getenv("WEB_NAMESPACE")
	if webNamespace == "" {
		L.Info("defaulting to namespace for frontend", "namespace", "/waypoint")
		webNamespace = "/waypoint"
	}

	id, err := pb.ParseULID(sid)
	if err != nil {
		log.Fatal(err)
	}

	tmpdir, err := ioutil.TempDir("", "hzn")
	if err != nil {
		log.Fatal(err)
	}

	defer os.RemoveAll(tmpdir)

	client, err := control.NewClient(ctx, control.ClientConfig{
		Id:      id,
		Token:   token,
		Version: "test",
		Addr:    addr,
		WorkDir: tmpdir,
	})

	defer client.Close(ctx)

	var labels *pb.LabelSet

	strLabels := os.Getenv("LOCATION_LABELS")
	if strLabels != "" {
		labels = pb.ParseLabelSet(os.Getenv(strLabels))
	}

	locs, err := client.LearnLocations(labels)
	if err != nil {
		log.Fatal(err)
	}

	err = client.BootstrapConfig(ctx)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		err := client.Run(ctx)
		if err != nil {
			L.Error("error running control client background tasks", "error", err)
		}
	}()

	L.Info("generating token to access accounts for web")
	serviceToken, err := client.RequestServiceToken(ctx, webNamespace)
	if err != nil {
		log.Fatal(err)
	}

	ln, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatal(err)
	}

	defer ln.Close()

	hb, err := hub.NewHub(L, client, serviceToken)
	if err != nil {
		log.Fatal(err)
	}

	for _, loc := range locs {
		L.Info("learned network location", "labels", loc.Labels, "addresses", loc.Addresses)
	}

	if httpPort != "" {
		L.Info("listen on http", "port", httpPort)
		go hb.ListenHTTP(":" + httpPort)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err = hb.Run(ctx, ln)
	if err != nil {
		log.Fatal(err)
	}

	return 0
}