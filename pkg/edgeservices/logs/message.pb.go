// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: message.proto

package logs

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"

	proto "github.com/gogo/protobuf/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Timestamp struct {
	Sec  uint64 `protobuf:"varint,1,opt,name=sec,proto3" json:"sec,omitempty"`
	Nsec uint32 `protobuf:"varint,2,opt,name=nsec,proto3" json:"nsec,omitempty"`
}

func (m *Timestamp) Reset()      { *m = Timestamp{} }
func (*Timestamp) ProtoMessage() {}
func (*Timestamp) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{0}
}
func (m *Timestamp) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Timestamp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Timestamp.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Timestamp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Timestamp.Merge(m, src)
}
func (m *Timestamp) XXX_Size() int {
	return m.Size()
}
func (m *Timestamp) XXX_DiscardUnknown() {
	xxx_messageInfo_Timestamp.DiscardUnknown(m)
}

var xxx_messageInfo_Timestamp proto.InternalMessageInfo

func (m *Timestamp) GetSec() uint64 {
	if m != nil {
		return m.Sec
	}
	return 0
}

func (m *Timestamp) GetNsec() uint32 {
	if m != nil {
		return m.Nsec
	}
	return 0
}

type Attribute struct {
	Key  string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Sval string `protobuf:"bytes,2,opt,name=sval,proto3" json:"sval,omitempty"`
	Ival int64  `protobuf:"varint,3,opt,name=ival,proto3" json:"ival,omitempty"`
}

func (m *Attribute) Reset()      { *m = Attribute{} }
func (*Attribute) ProtoMessage() {}
func (*Attribute) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{1}
}
func (m *Attribute) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Attribute) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Attribute.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Attribute) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Attribute.Merge(m, src)
}
func (m *Attribute) XXX_Size() int {
	return m.Size()
}
func (m *Attribute) XXX_DiscardUnknown() {
	xxx_messageInfo_Attribute.DiscardUnknown(m)
}

var xxx_messageInfo_Attribute proto.InternalMessageInfo

func (m *Attribute) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Attribute) GetSval() string {
	if m != nil {
		return m.Sval
	}
	return ""
}

func (m *Attribute) GetIval() int64 {
	if m != nil {
		return m.Ival
	}
	return 0
}

type Message struct {
	Timestamp *Timestamp   `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Mesg      string       `protobuf:"bytes,2,opt,name=mesg,proto3" json:"mesg,omitempty"`
	Attrs     []*Attribute `protobuf:"bytes,3,rep,name=attrs,proto3" json:"attrs,omitempty"`
}

func (m *Message) Reset()      { *m = Message{} }
func (*Message) ProtoMessage() {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{2}
}
func (m *Message) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Message.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return m.Size()
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetTimestamp() *Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *Message) GetMesg() string {
	if m != nil {
		return m.Mesg
	}
	return ""
}

func (m *Message) GetAttrs() []*Attribute {
	if m != nil {
		return m.Attrs
	}
	return nil
}

func init() {
	proto.RegisterType((*Timestamp)(nil), "logs.Timestamp")
	proto.RegisterType((*Attribute)(nil), "logs.Attribute")
	proto.RegisterType((*Message)(nil), "logs.Message")
}

func init() { proto.RegisterFile("message.proto", fileDescriptor_33c57e4bae7b9afd) }

var fileDescriptor_33c57e4bae7b9afd = []byte{
	// 254 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x90, 0x4d, 0x4a, 0xc4, 0x30,
	0x14, 0x80, 0xf3, 0x4c, 0x55, 0x12, 0x19, 0x94, 0xac, 0x66, 0xf5, 0x28, 0x05, 0xa1, 0x1b, 0x0b,
	0xfe, 0x5c, 0x40, 0xc1, 0xa5, 0x9b, 0xe0, 0x05, 0x3a, 0x12, 0x4a, 0x71, 0x6a, 0x87, 0xbe, 0x28,
	0xb8, 0xf3, 0x08, 0x1e, 0xc3, 0xa3, 0xb8, 0xec, 0x72, 0x96, 0x36, 0xdd, 0xb8, 0x9c, 0x23, 0x0c,
	0x2f, 0xc3, 0x4c, 0x77, 0x1f, 0x2f, 0xf9, 0x5e, 0x3e, 0xa2, 0x67, 0x8d, 0x23, 0x2a, 0x2b, 0x57,
	0xac, 0xba, 0xd6, 0xb7, 0x26, 0x59, 0xb6, 0x15, 0x65, 0xd7, 0x5a, 0x3d, 0xd7, 0x8d, 0x23, 0x5f,
	0x36, 0x2b, 0x73, 0xa1, 0x25, 0xb9, 0x97, 0x39, 0xa4, 0x90, 0x27, 0x96, 0xd1, 0x18, 0x9d, 0xbc,
	0xf1, 0xe8, 0x28, 0x85, 0x7c, 0x66, 0x23, 0x67, 0x8f, 0x5a, 0xdd, 0x7b, 0xdf, 0xd5, 0x8b, 0x77,
	0xef, 0x58, 0x79, 0x75, 0x9f, 0x51, 0x51, 0x96, 0x91, 0x15, 0xfa, 0x28, 0x97, 0x51, 0x51, 0x36,
	0x32, 0xcf, 0x6a, 0x9e, 0xc9, 0x14, 0x72, 0x69, 0x23, 0x67, 0xa4, 0x4f, 0x9f, 0x76, 0x41, 0xe6,
	0x4a, 0x2b, 0xbf, 0x8f, 0x88, 0xab, 0xce, 0x6e, 0xce, 0x0b, 0xce, 0x2b, 0x0e, 0x6d, 0x76, 0xba,
	0xc1, 0xdb, 0x1a, 0x47, 0xd5, 0xfe, 0x05, 0x66, 0x73, 0xa9, 0x8f, 0x4b, 0xef, 0x3b, 0x9a, 0xcb,
	0x54, 0x4e, 0xfa, 0xa1, 0xd3, 0xee, 0x4e, 0x1f, 0xee, 0xfa, 0x01, 0xc5, 0x7a, 0x40, 0xb1, 0x19,
	0x10, 0xbe, 0x02, 0xc2, 0x4f, 0x40, 0xf8, 0x0d, 0x08, 0x7d, 0x40, 0xf8, 0x0b, 0x08, 0xff, 0x01,
	0xc5, 0x26, 0x20, 0x7c, 0x8f, 0x28, 0xfa, 0x11, 0xc5, 0x7a, 0x44, 0xb1, 0x38, 0x89, 0x3f, 0x76,
	0xbb, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x39, 0x16, 0xe9, 0xec, 0x42, 0x01, 0x00, 0x00,
}

func (this *Timestamp) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Timestamp)
	if !ok {
		that2, ok := that.(Timestamp)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Sec != that1.Sec {
		return false
	}
	if this.Nsec != that1.Nsec {
		return false
	}
	return true
}
func (this *Attribute) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Attribute)
	if !ok {
		that2, ok := that.(Attribute)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Key != that1.Key {
		return false
	}
	if this.Sval != that1.Sval {
		return false
	}
	if this.Ival != that1.Ival {
		return false
	}
	return true
}
func (this *Message) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Message)
	if !ok {
		that2, ok := that.(Message)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Timestamp.Equal(that1.Timestamp) {
		return false
	}
	if this.Mesg != that1.Mesg {
		return false
	}
	if len(this.Attrs) != len(that1.Attrs) {
		return false
	}
	for i := range this.Attrs {
		if !this.Attrs[i].Equal(that1.Attrs[i]) {
			return false
		}
	}
	return true
}
func (this *Timestamp) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&logs.Timestamp{")
	s = append(s, "Sec: "+fmt.Sprintf("%#v", this.Sec)+",\n")
	s = append(s, "Nsec: "+fmt.Sprintf("%#v", this.Nsec)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *Attribute) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&logs.Attribute{")
	s = append(s, "Key: "+fmt.Sprintf("%#v", this.Key)+",\n")
	s = append(s, "Sval: "+fmt.Sprintf("%#v", this.Sval)+",\n")
	s = append(s, "Ival: "+fmt.Sprintf("%#v", this.Ival)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *Message) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&logs.Message{")
	if this.Timestamp != nil {
		s = append(s, "Timestamp: "+fmt.Sprintf("%#v", this.Timestamp)+",\n")
	}
	s = append(s, "Mesg: "+fmt.Sprintf("%#v", this.Mesg)+",\n")
	if this.Attrs != nil {
		s = append(s, "Attrs: "+fmt.Sprintf("%#v", this.Attrs)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringMessage(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *Timestamp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Timestamp) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Timestamp) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Nsec != 0 {
		i = encodeVarintMessage(dAtA, i, uint64(m.Nsec))
		i--
		dAtA[i] = 0x10
	}
	if m.Sec != 0 {
		i = encodeVarintMessage(dAtA, i, uint64(m.Sec))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Attribute) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Attribute) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Attribute) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Ival != 0 {
		i = encodeVarintMessage(dAtA, i, uint64(m.Ival))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Sval) > 0 {
		i -= len(m.Sval)
		copy(dAtA[i:], m.Sval)
		i = encodeVarintMessage(dAtA, i, uint64(len(m.Sval)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Key) > 0 {
		i -= len(m.Key)
		copy(dAtA[i:], m.Key)
		i = encodeVarintMessage(dAtA, i, uint64(len(m.Key)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Message) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Message) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Message) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Attrs) > 0 {
		for iNdEx := len(m.Attrs) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Attrs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintMessage(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Mesg) > 0 {
		i -= len(m.Mesg)
		copy(dAtA[i:], m.Mesg)
		i = encodeVarintMessage(dAtA, i, uint64(len(m.Mesg)))
		i--
		dAtA[i] = 0x12
	}
	if m.Timestamp != nil {
		{
			size, err := m.Timestamp.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintMessage(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintMessage(dAtA []byte, offset int, v uint64) int {
	offset -= sovMessage(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Timestamp) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Sec != 0 {
		n += 1 + sovMessage(uint64(m.Sec))
	}
	if m.Nsec != 0 {
		n += 1 + sovMessage(uint64(m.Nsec))
	}
	return n
}

func (m *Attribute) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovMessage(uint64(l))
	}
	l = len(m.Sval)
	if l > 0 {
		n += 1 + l + sovMessage(uint64(l))
	}
	if m.Ival != 0 {
		n += 1 + sovMessage(uint64(m.Ival))
	}
	return n
}

func (m *Message) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Timestamp != nil {
		l = m.Timestamp.Size()
		n += 1 + l + sovMessage(uint64(l))
	}
	l = len(m.Mesg)
	if l > 0 {
		n += 1 + l + sovMessage(uint64(l))
	}
	if len(m.Attrs) > 0 {
		for _, e := range m.Attrs {
			l = e.Size()
			n += 1 + l + sovMessage(uint64(l))
		}
	}
	return n
}

func sovMessage(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMessage(x uint64) (n int) {
	return sovMessage(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Timestamp) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Timestamp{`,
		`Sec:` + fmt.Sprintf("%v", this.Sec) + `,`,
		`Nsec:` + fmt.Sprintf("%v", this.Nsec) + `,`,
		`}`,
	}, "")
	return s
}
func (this *Attribute) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Attribute{`,
		`Key:` + fmt.Sprintf("%v", this.Key) + `,`,
		`Sval:` + fmt.Sprintf("%v", this.Sval) + `,`,
		`Ival:` + fmt.Sprintf("%v", this.Ival) + `,`,
		`}`,
	}, "")
	return s
}
func (this *Message) String() string {
	if this == nil {
		return "nil"
	}
	repeatedStringForAttrs := "[]*Attribute{"
	for _, f := range this.Attrs {
		repeatedStringForAttrs += strings.Replace(f.String(), "Attribute", "Attribute", 1) + ","
	}
	repeatedStringForAttrs += "}"
	s := strings.Join([]string{`&Message{`,
		`Timestamp:` + strings.Replace(this.Timestamp.String(), "Timestamp", "Timestamp", 1) + `,`,
		`Mesg:` + fmt.Sprintf("%v", this.Mesg) + `,`,
		`Attrs:` + repeatedStringForAttrs + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringMessage(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *Timestamp) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessage
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Timestamp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Timestamp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sec", wireType)
			}
			m.Sec = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Sec |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nsec", wireType)
			}
			m.Nsec = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Nsec |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipMessage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMessage
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMessage
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Attribute) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessage
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Attribute: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Attribute: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sval", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sval = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ival", wireType)
			}
			m.Ival = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Ival |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipMessage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMessage
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMessage
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Message) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessage
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Message: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Message: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Timestamp == nil {
				m.Timestamp = &Timestamp{}
			}
			if err := m.Timestamp.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Mesg", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Mesg = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Attrs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Attrs = append(m.Attrs, &Attribute{})
			if err := m.Attrs[len(m.Attrs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMessage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMessage
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMessage
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipMessage(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMessage
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthMessage
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMessage
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMessage
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMessage        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMessage          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMessage = fmt.Errorf("proto: unexpected end of group")
)
