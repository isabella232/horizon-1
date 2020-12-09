package control

import (
	context "context"
	"crypto/tls"
	"crypto/x509"
	"sync"

	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-multierror"
	"github.com/hashicorp/horizon/pkg/grpc/lz4"
	grpctoken "github.com/hashicorp/horizon/pkg/grpc/token"
	"github.com/hashicorp/horizon/pkg/pb"
	"github.com/hashicorp/horizon/pkg/utils"
	"google.golang.org/grpc"
	gcreds "google.golang.org/grpc/credentials"
)

// HubCatalog is a simple interface to decouple the gather and management of hub addresses from
// the code that broadcasts to them. This is implemented by ConsulMonitor and used
// primarily in production.
type HubCatalog interface {
	Targets() []string
}

// Broadcaster is a simple fan out value. The commands sent to it via funciton calls are
// fanned out to all targets in the given HubCatalog.
type Broadcaster struct {
	L       hclog.Logger
	catalog HubCatalog
	conn    func(addr string) (pb.HubServicesClient, error)
}

// NewBroadcaster creates a new Broadcaster value. The targets to broadcast to come from
// catalog. conn is how we actually open a connection to the target. This conn decoupling
// makes this code much easier to test. In production, conn is usually GRPCDial.Dial.
func NewBroadcaster(
	L hclog.Logger,
	catalog HubCatalog,
	conn func(addr string) (pb.HubServicesClient, error),
) (*Broadcaster, error) {
	br := &Broadcaster{
		L:       L,
		catalog: catalog,
		conn:    conn,
	}

	return br, nil
}

// AdvertiseServices gets a list of targets from the catalog and calls AddService
// on the clients generated from the connect function (which defaults to dialing a grpc
// connection to the target)
func (b *Broadcaster) AdvertiseServices(ctx context.Context, as *pb.AccountServices) error {
	var topError error

	targets := b.catalog.Targets()

	b.L.Info("hub broadcasting beginning", "targets", len(targets))

	for _, tgt := range targets {
		b.L.Info("broadcasting hub update", "target", tgt)
		client, err := b.conn(tgt)
		if err != nil {
			topError = multierror.Append(topError, err)
			continue
		}

		_, err = client.AddServices(ctx, as)
		if err != nil {
			topError = multierror.Append(topError, err)
		}
	}

	return topError
}

// GRPCDial provides connection pooling grpc connections to hubs. It is used to
// avoid spinning up new TCP connections to hubs on every advertise operation.
type GRPCDial struct {
	token string
	cert  []byte

	mu        sync.RWMutex
	grpcConns map[string]*grpc.ClientConn

	tlscfg tls.Config
}

// NewGRPCDial creates a new GRPCDial value. The given token is the authentication
// token that will be included with all calls to the hubs, to identify them as valid.
// cert is a TLS certification that, if set, will be used as the only cert in the TLS
// RootCAs. This further restricts the code to calling valid hubs by making sure that
// the code is only talking to hubs that are using the certs managed by control.
func NewGRPCDial(token string, cert []byte) (*GRPCDial, error) {
	g := &GRPCDial{
		token:     token,
		cert:      cert,
		grpcConns: make(map[string]*grpc.ClientConn),
	}

	if g.cert != nil {
		parsedHubCert, err := utils.ParseCertificate(cert)
		if err != nil {
			return nil, err
		}

		g.tlscfg.RootCAs = x509.NewCertPool()
		g.tlscfg.RootCAs.AddCert(parsedHubCert)
	}

	return g, nil
}

// Dial gets a gRPC client for target. It either generates a new gRPC connection
// to the given target, used as a host:port combo. Or it returns a existing
// connection.
func (g *GRPCDial) Dial(target string) (pb.HubServicesClient, error) {
	g.mu.RLock()
	cc, ok := g.grpcConns[target]
	g.mu.RUnlock()

	if ok {
		return pb.NewHubServicesClient(cc), nil
	}

	g.mu.Lock()
	defer g.mu.Unlock()

	// There is a race here so we have to check again.
	cc, ok = g.grpcConns[target]
	if ok {
		return pb.NewHubServicesClient(cc), nil
	}

	opts := []grpc.DialOption{
		grpc.WithDefaultCallOptions(grpc.UseCompressor(lz4.Name)),
	}

	if g.token != "" {
		opts = append(opts, grpc.WithPerRPCCredentials(grpctoken.Token(g.token)))
	}

	creds := gcreds.NewTLS(&g.tlscfg)

	opts = append(opts, grpc.WithTransportCredentials(creds))

	cc, err := grpc.Dial(target, opts...)
	if err != nil {
		return nil, err
	}

	g.grpcConns[target] = cc

	return pb.NewHubServicesClient(cc), nil
}