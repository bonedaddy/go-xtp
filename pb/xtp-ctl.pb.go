// Code generated by protoc-gen-gogo.
// source: xtp-ctl.proto
// DO NOT EDIT!

/*
Package xtp_ctl is a generated protocol buffer package.

It is generated from these files:
	xtp-ctl.proto

It has these top-level messages:
	RPC
	Transport
	Listener
	Dialer
	Conn
	Stream
	ListReq
	ListRes
	CloseReq
	ListenReq
	ListenRes
	AcceptReq
	AcceptRes
	DialerReq
	DialerRes
	DialReq
	DialRes
*/
package xtp_ctl

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type TType int32

const (
	TType_TTypeNil       TType = 0
	TType_TTypeTransport TType = 1
	TType_TTypeListener  TType = 2
	TType_TTypeDialer    TType = 3
	TType_TTypeConn      TType = 4
	TType_TTypeStream    TType = 5
)

var TType_name = map[int32]string{
	0: "TTypeNil",
	1: "TTypeTransport",
	2: "TTypeListener",
	3: "TTypeDialer",
	4: "TTypeConn",
	5: "TTypeStream",
}
var TType_value = map[string]int32{
	"TTypeNil":       0,
	"TTypeTransport": 1,
	"TTypeListener":  2,
	"TTypeDialer":    3,
	"TTypeConn":      4,
	"TTypeStream":    5,
}

func (x TType) Enum() *TType {
	p := new(TType)
	*p = x
	return p
}
func (x TType) String() string {
	return proto.EnumName(TType_name, int32(x))
}
func (x *TType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(TType_value, data, "TType")
	if err != nil {
		return err
	}
	*x = TType(value)
	return nil
}
func (TType) EnumDescriptor() ([]byte, []int) { return fileDescriptorXtpCtl, []int{0} }

type RPC_Type int32

const (
	RPC_Null RPC_Type = 0
	RPC_NoOp RPC_Type = 1
	// List the available resources (Transports, Listeners, Dialers, Conns, Streams)
	RPC_ListReq RPC_Type = 2
	RPC_ListRes RPC_Type = 3
	// close a descriptor. used for Listener, Dialer, Conn, Stream
	RPC_CloseReq RPC_Type = 4
	RPC_CloseRes RPC_Type = 5
	// Open a listener
	RPC_ListenReq RPC_Type = 6
	RPC_ListenRes RPC_Type = 7
	// Listener.AcceptConn() or Conn.AcceptStream()
	RPC_AcceptReq RPC_Type = 8
	RPC_AcceptRes RPC_Type = 9
	// Open a dialer
	RPC_DialerReq RPC_Type = 10
	RPC_DialerRes RPC_Type = 11
	// Dialer.DialConn() or Conn.OpenStream()
	RPC_DialReq RPC_Type = 12
	RPC_DialRes RPC_Type = 13
)

var RPC_Type_name = map[int32]string{
	0:  "Null",
	1:  "NoOp",
	2:  "ListReq",
	3:  "ListRes",
	4:  "CloseReq",
	5:  "CloseRes",
	6:  "ListenReq",
	7:  "ListenRes",
	8:  "AcceptReq",
	9:  "AcceptRes",
	10: "DialerReq",
	11: "DialerRes",
	12: "DialReq",
	13: "DialRes",
}
var RPC_Type_value = map[string]int32{
	"Null":      0,
	"NoOp":      1,
	"ListReq":   2,
	"ListRes":   3,
	"CloseReq":  4,
	"CloseRes":  5,
	"ListenReq": 6,
	"ListenRes": 7,
	"AcceptReq": 8,
	"AcceptRes": 9,
	"DialerReq": 10,
	"DialerRes": 11,
	"DialReq":   12,
	"DialRes":   13,
}

func (x RPC_Type) Enum() *RPC_Type {
	p := new(RPC_Type)
	*p = x
	return p
}
func (x RPC_Type) String() string {
	return proto.EnumName(RPC_Type_name, int32(x))
}
func (x *RPC_Type) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(RPC_Type_value, data, "RPC_Type")
	if err != nil {
		return err
	}
	*x = RPC_Type(value)
	return nil
}
func (RPC_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptorXtpCtl, []int{0, 0} }

type RPC struct {
	Rpc              *RPC_Type `protobuf:"varint,1,opt,name=rpc,enum=RPC_Type" json:"rpc,omitempty"`
	Message          []byte    `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
	Error            *string   `protobuf:"bytes,3,opt,name=error" json:"error,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *RPC) Reset()                    { *m = RPC{} }
func (m *RPC) String() string            { return proto.CompactTextString(m) }
func (*RPC) ProtoMessage()               {}
func (*RPC) Descriptor() ([]byte, []int) { return fileDescriptorXtpCtl, []int{0} }

func (m *RPC) GetRpc() RPC_Type {
	if m != nil && m.Rpc != nil {
		return *m.Rpc
	}
	return RPC_Null
}

func (m *RPC) GetMessage() []byte {
	if m != nil {
		return m.Message
	}
	return nil
}

func (m *RPC) GetError() string {
	if m != nil && m.Error != nil {
		return *m.Error
	}
	return ""
}

// The types of things we use.
type Transport struct {
	Id               *int64  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Transport        *string `protobuf:"bytes,2,opt,name=transport" json:"transport,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Transport) Reset()                    { *m = Transport{} }
func (m *Transport) String() string            { return proto.CompactTextString(m) }
func (*Transport) ProtoMessage()               {}
func (*Transport) Descriptor() ([]byte, []int) { return fileDescriptorXtpCtl, []int{1} }

func (m *Transport) GetId() int64 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *Transport) GetTransport() string {
	if m != nil && m.Transport != nil {
		return *m.Transport
	}
	return ""
}

type Listener struct {
	Id               *int64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	TransportId      *int64 `protobuf:"varint,2,opt,name=transportId" json:"transportId,omitempty"`
	Multiaddr        []byte `protobuf:"bytes,3,opt,name=multiaddr" json:"multiaddr,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Listener) Reset()                    { *m = Listener{} }
func (m *Listener) String() string            { return proto.CompactTextString(m) }
func (*Listener) ProtoMessage()               {}
func (*Listener) Descriptor() ([]byte, []int) { return fileDescriptorXtpCtl, []int{2} }

func (m *Listener) GetId() int64 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *Listener) GetTransportId() int64 {
	if m != nil && m.TransportId != nil {
		return *m.TransportId
	}
	return 0
}

func (m *Listener) GetMultiaddr() []byte {
	if m != nil {
		return m.Multiaddr
	}
	return nil
}

type Dialer struct {
	Id               *int64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	TransportId      *int64 `protobuf:"varint,2,opt,name=transportId" json:"transportId,omitempty"`
	Multiaddr        []byte `protobuf:"bytes,3,opt,name=multiaddr" json:"multiaddr,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Dialer) Reset()                    { *m = Dialer{} }
func (m *Dialer) String() string            { return proto.CompactTextString(m) }
func (*Dialer) ProtoMessage()               {}
func (*Dialer) Descriptor() ([]byte, []int) { return fileDescriptorXtpCtl, []int{3} }

func (m *Dialer) GetId() int64 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *Dialer) GetTransportId() int64 {
	if m != nil && m.TransportId != nil {
		return *m.TransportId
	}
	return 0
}

func (m *Dialer) GetMultiaddr() []byte {
	if m != nil {
		return m.Multiaddr
	}
	return nil
}

type Conn struct {
	Id               *int64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	TransportId      *int64 `protobuf:"varint,2,opt,name=transportId" json:"transportId,omitempty"`
	LocalMultiaddr   []byte `protobuf:"bytes,3,opt,name=localMultiaddr" json:"localMultiaddr,omitempty"`
	RemoteMultiaddr  []byte `protobuf:"bytes,4,opt,name=remoteMultiaddr" json:"remoteMultiaddr,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Conn) Reset()                    { *m = Conn{} }
func (m *Conn) String() string            { return proto.CompactTextString(m) }
func (*Conn) ProtoMessage()               {}
func (*Conn) Descriptor() ([]byte, []int) { return fileDescriptorXtpCtl, []int{4} }

func (m *Conn) GetId() int64 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *Conn) GetTransportId() int64 {
	if m != nil && m.TransportId != nil {
		return *m.TransportId
	}
	return 0
}

func (m *Conn) GetLocalMultiaddr() []byte {
	if m != nil {
		return m.LocalMultiaddr
	}
	return nil
}

func (m *Conn) GetRemoteMultiaddr() []byte {
	if m != nil {
		return m.RemoteMultiaddr
	}
	return nil
}

type Stream struct {
	Id               *int64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	ConnId           *int64 `protobuf:"varint,2,opt,name=connId" json:"connId,omitempty"`
	TransportId      *int64 `protobuf:"varint,3,opt,name=transportId" json:"transportId,omitempty"`
	LocalMultiaddr   []byte `protobuf:"bytes,4,opt,name=localMultiaddr" json:"localMultiaddr,omitempty"`
	RemoteMultiaddr  []byte `protobuf:"bytes,5,opt,name=remoteMultiaddr" json:"remoteMultiaddr,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Stream) Reset()                    { *m = Stream{} }
func (m *Stream) String() string            { return proto.CompactTextString(m) }
func (*Stream) ProtoMessage()               {}
func (*Stream) Descriptor() ([]byte, []int) { return fileDescriptorXtpCtl, []int{5} }

func (m *Stream) GetId() int64 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *Stream) GetConnId() int64 {
	if m != nil && m.ConnId != nil {
		return *m.ConnId
	}
	return 0
}

func (m *Stream) GetTransportId() int64 {
	if m != nil && m.TransportId != nil {
		return *m.TransportId
	}
	return 0
}

func (m *Stream) GetLocalMultiaddr() []byte {
	if m != nil {
		return m.LocalMultiaddr
	}
	return nil
}

func (m *Stream) GetRemoteMultiaddr() []byte {
	if m != nil {
		return m.RemoteMultiaddr
	}
	return nil
}

type ListReq struct {
	// include one per type we want.
	// same TType multiple times is idempotent.
	Types            []TType `protobuf:"varint,1,rep,name=types,enum=TType" json:"types,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ListReq) Reset()                    { *m = ListReq{} }
func (m *ListReq) String() string            { return proto.CompactTextString(m) }
func (*ListReq) ProtoMessage()               {}
func (*ListReq) Descriptor() ([]byte, []int) { return fileDescriptorXtpCtl, []int{6} }

func (m *ListReq) GetTypes() []TType {
	if m != nil {
		return m.Types
	}
	return nil
}

type ListRes struct {
	Items            []*ListRes_Item `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *ListRes) Reset()                    { *m = ListRes{} }
func (m *ListRes) String() string            { return proto.CompactTextString(m) }
func (*ListRes) ProtoMessage()               {}
func (*ListRes) Descriptor() ([]byte, []int) { return fileDescriptorXtpCtl, []int{7} }

func (m *ListRes) GetItems() []*ListRes_Item {
	if m != nil {
		return m.Items
	}
	return nil
}

type ListRes_Item struct {
	Id               *int64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Type             *TType `protobuf:"varint,2,opt,name=type,enum=TType" json:"type,omitempty"`
	Value            []byte `protobuf:"bytes,3,opt,name=value" json:"value,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *ListRes_Item) Reset()                    { *m = ListRes_Item{} }
func (m *ListRes_Item) String() string            { return proto.CompactTextString(m) }
func (*ListRes_Item) ProtoMessage()               {}
func (*ListRes_Item) Descriptor() ([]byte, []int) { return fileDescriptorXtpCtl, []int{7, 0} }

func (m *ListRes_Item) GetId() int64 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *ListRes_Item) GetType() TType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return TType_TTypeNil
}

func (m *ListRes_Item) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type CloseReq struct {
	Id               *int64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *CloseReq) Reset()                    { *m = CloseReq{} }
func (m *CloseReq) String() string            { return proto.CompactTextString(m) }
func (*CloseReq) ProtoMessage()               {}
func (*CloseReq) Descriptor() ([]byte, []int) { return fileDescriptorXtpCtl, []int{8} }

func (m *CloseReq) GetId() int64 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

type ListenReq struct {
	ListenerOpts     *Listener `protobuf:"bytes,1,opt,name=listenerOpts" json:"listenerOpts,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *ListenReq) Reset()                    { *m = ListenReq{} }
func (m *ListenReq) String() string            { return proto.CompactTextString(m) }
func (*ListenReq) ProtoMessage()               {}
func (*ListenReq) Descriptor() ([]byte, []int) { return fileDescriptorXtpCtl, []int{9} }

func (m *ListenReq) GetListenerOpts() *Listener {
	if m != nil {
		return m.ListenerOpts
	}
	return nil
}

type ListenRes struct {
	Listener         *Listener `protobuf:"bytes,1,opt,name=listener" json:"listener,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *ListenRes) Reset()                    { *m = ListenRes{} }
func (m *ListenRes) String() string            { return proto.CompactTextString(m) }
func (*ListenRes) ProtoMessage()               {}
func (*ListenRes) Descriptor() ([]byte, []int) { return fileDescriptorXtpCtl, []int{10} }

func (m *ListenRes) GetListener() *Listener {
	if m != nil {
		return m.Listener
	}
	return nil
}

type AcceptReq struct {
	Id               *int64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *AcceptReq) Reset()                    { *m = AcceptReq{} }
func (m *AcceptReq) String() string            { return proto.CompactTextString(m) }
func (*AcceptReq) ProtoMessage()               {}
func (*AcceptReq) Descriptor() ([]byte, []int) { return fileDescriptorXtpCtl, []int{11} }

func (m *AcceptReq) GetId() int64 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

type AcceptRes struct {
	Conn             *Conn   `protobuf:"bytes,1,opt,name=conn" json:"conn,omitempty"`
	Stream           *Stream `protobuf:"bytes,2,opt,name=stream" json:"stream,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *AcceptRes) Reset()                    { *m = AcceptRes{} }
func (m *AcceptRes) String() string            { return proto.CompactTextString(m) }
func (*AcceptRes) ProtoMessage()               {}
func (*AcceptRes) Descriptor() ([]byte, []int) { return fileDescriptorXtpCtl, []int{12} }

func (m *AcceptRes) GetConn() *Conn {
	if m != nil {
		return m.Conn
	}
	return nil
}

func (m *AcceptRes) GetStream() *Stream {
	if m != nil {
		return m.Stream
	}
	return nil
}

type DialerReq struct {
	DialerOpts       *Dialer `protobuf:"bytes,1,opt,name=dialerOpts" json:"dialerOpts,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *DialerReq) Reset()                    { *m = DialerReq{} }
func (m *DialerReq) String() string            { return proto.CompactTextString(m) }
func (*DialerReq) ProtoMessage()               {}
func (*DialerReq) Descriptor() ([]byte, []int) { return fileDescriptorXtpCtl, []int{13} }

func (m *DialerReq) GetDialerOpts() *Dialer {
	if m != nil {
		return m.DialerOpts
	}
	return nil
}

type DialerRes struct {
	Dialer           *Dialer `protobuf:"bytes,1,opt,name=dialer" json:"dialer,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *DialerRes) Reset()                    { *m = DialerRes{} }
func (m *DialerRes) String() string            { return proto.CompactTextString(m) }
func (*DialerRes) ProtoMessage()               {}
func (*DialerRes) Descriptor() ([]byte, []int) { return fileDescriptorXtpCtl, []int{14} }

func (m *DialerRes) GetDialer() *Dialer {
	if m != nil {
		return m.Dialer
	}
	return nil
}

type DialReq struct {
	Id               *int64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	ConnOpts         *Conn  `protobuf:"bytes,2,opt,name=connOpts" json:"connOpts,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *DialReq) Reset()                    { *m = DialReq{} }
func (m *DialReq) String() string            { return proto.CompactTextString(m) }
func (*DialReq) ProtoMessage()               {}
func (*DialReq) Descriptor() ([]byte, []int) { return fileDescriptorXtpCtl, []int{15} }

func (m *DialReq) GetId() int64 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *DialReq) GetConnOpts() *Conn {
	if m != nil {
		return m.ConnOpts
	}
	return nil
}

type DialRes struct {
	Conn             *Conn   `protobuf:"bytes,1,opt,name=conn" json:"conn,omitempty"`
	Stream           *Stream `protobuf:"bytes,2,opt,name=stream" json:"stream,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *DialRes) Reset()                    { *m = DialRes{} }
func (m *DialRes) String() string            { return proto.CompactTextString(m) }
func (*DialRes) ProtoMessage()               {}
func (*DialRes) Descriptor() ([]byte, []int) { return fileDescriptorXtpCtl, []int{16} }

func (m *DialRes) GetConn() *Conn {
	if m != nil {
		return m.Conn
	}
	return nil
}

func (m *DialRes) GetStream() *Stream {
	if m != nil {
		return m.Stream
	}
	return nil
}

func init() {
	proto.RegisterType((*RPC)(nil), "RPC")
	proto.RegisterType((*Transport)(nil), "Transport")
	proto.RegisterType((*Listener)(nil), "Listener")
	proto.RegisterType((*Dialer)(nil), "Dialer")
	proto.RegisterType((*Conn)(nil), "Conn")
	proto.RegisterType((*Stream)(nil), "Stream")
	proto.RegisterType((*ListReq)(nil), "ListReq")
	proto.RegisterType((*ListRes)(nil), "ListRes")
	proto.RegisterType((*ListRes_Item)(nil), "ListRes.Item")
	proto.RegisterType((*CloseReq)(nil), "CloseReq")
	proto.RegisterType((*ListenReq)(nil), "ListenReq")
	proto.RegisterType((*ListenRes)(nil), "ListenRes")
	proto.RegisterType((*AcceptReq)(nil), "AcceptReq")
	proto.RegisterType((*AcceptRes)(nil), "AcceptRes")
	proto.RegisterType((*DialerReq)(nil), "DialerReq")
	proto.RegisterType((*DialerRes)(nil), "DialerRes")
	proto.RegisterType((*DialReq)(nil), "DialReq")
	proto.RegisterType((*DialRes)(nil), "DialRes")
	proto.RegisterEnum("TType", TType_name, TType_value)
	proto.RegisterEnum("RPC_Type", RPC_Type_name, RPC_Type_value)
}

func init() { proto.RegisterFile("xtp-ctl.proto", fileDescriptorXtpCtl) }

var fileDescriptorXtpCtl = []byte{
	// 627 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x54, 0xcd, 0x6e, 0xd4, 0x3c,
	0x14, 0xfd, 0xf2, 0x37, 0x33, 0xb9, 0x99, 0x99, 0xfa, 0xb3, 0x10, 0x1a, 0xda, 0x4a, 0x0c, 0x46,
	0xd0, 0x11, 0xa2, 0x59, 0x44, 0x6c, 0x40, 0x6c, 0xd0, 0x80, 0xa0, 0x12, 0xb4, 0x95, 0xe9, 0x02,
	0xb1, 0x8b, 0x12, 0x0b, 0x45, 0xca, 0x1f, 0xb1, 0x8b, 0xda, 0x57, 0xe0, 0x0d, 0xd8, 0xf3, 0x28,
	0x3c, 0x18, 0xb2, 0x9d, 0x49, 0x4c, 0xc2, 0xa2, 0x02, 0x76, 0x3e, 0xe7, 0x5c, 0x5f, 0x1f, 0xdb,
	0xf7, 0x5e, 0x58, 0x5c, 0x89, 0xfa, 0x38, 0x11, 0x79, 0x58, 0x37, 0x95, 0xa8, 0xc8, 0x37, 0x1b,
	0x1c, 0x7a, 0xbe, 0xc5, 0x07, 0xe0, 0x34, 0x75, 0xb2, 0xb2, 0xd6, 0xd6, 0x66, 0x19, 0xf9, 0x21,
	0x3d, 0xdf, 0x86, 0x17, 0xd7, 0x35, 0xa3, 0x92, 0xc5, 0x2b, 0x98, 0x16, 0x8c, 0xf3, 0xf8, 0x13,
	0x5b, 0xd9, 0x6b, 0x6b, 0x33, 0xa7, 0x3b, 0x88, 0x6f, 0x81, 0xc7, 0x9a, 0xa6, 0x6a, 0x56, 0xce,
	0xda, 0xda, 0xf8, 0x54, 0x03, 0xf2, 0xc3, 0x02, 0x57, 0xee, 0xc6, 0x33, 0x70, 0x4f, 0x2f, 0xf3,
	0x1c, 0xfd, 0xa7, 0x56, 0xd5, 0x59, 0x8d, 0x2c, 0x1c, 0xc0, 0xf4, 0x6d, 0xc6, 0x05, 0x65, 0x9f,
	0x91, 0xdd, 0x03, 0x8e, 0x1c, 0x3c, 0x87, 0xd9, 0x36, 0xaf, 0x38, 0x93, 0x92, 0x6b, 0x20, 0x8e,
	0x3c, 0xbc, 0x00, 0x5f, 0x06, 0xb2, 0x52, 0x8a, 0x13, 0x13, 0x72, 0x34, 0x95, 0xf0, 0x45, 0x92,
	0xb0, 0x5a, 0x65, 0x9d, 0x99, 0x90, 0x23, 0x5f, 0xc2, 0x97, 0x59, 0x9c, 0xb3, 0x46, 0xaa, 0x60,
	0x42, 0x8e, 0x02, 0x69, 0x41, 0x42, 0xa9, 0xcd, 0x7b, 0xc0, 0xd1, 0x82, 0x3c, 0x05, 0xff, 0xa2,
	0x89, 0x4b, 0x5e, 0x57, 0x8d, 0xc0, 0x4b, 0xb0, 0xb3, 0x54, 0xbd, 0x8f, 0x43, 0xed, 0x2c, 0xc5,
	0x87, 0xe0, 0x8b, 0x9d, 0xa8, 0x5e, 0xc5, 0xa7, 0x3d, 0x41, 0x3e, 0xc2, 0x4c, 0xfb, 0x63, 0xcd,
	0x68, 0xe7, 0x1a, 0x82, 0x2e, 0xf0, 0x24, 0x55, 0x7b, 0x1d, 0x6a, 0x52, 0x32, 0x77, 0x71, 0x99,
	0x8b, 0x2c, 0x4e, 0x53, 0xfd, 0xb2, 0x73, 0xda, 0x13, 0xe4, 0x03, 0x4c, 0xb4, 0xff, 0x7f, 0x9e,
	0xf9, 0xab, 0x05, 0xee, 0xb6, 0x2a, 0xcb, 0x3f, 0x48, 0xfc, 0x10, 0x96, 0x79, 0x95, 0xc4, 0xf9,
	0xbb, 0x41, 0xf6, 0x01, 0x8b, 0x37, 0xb0, 0xd7, 0xb0, 0xa2, 0x12, 0xac, 0x0f, 0x74, 0x55, 0xe0,
	0x90, 0x26, 0xdf, 0x2d, 0x98, 0xbc, 0x17, 0x0d, 0x8b, 0x8b, 0x91, 0x9d, 0xdb, 0x30, 0x49, 0xaa,
	0xb2, 0xec, 0x9c, 0xb4, 0x68, 0x68, 0xd3, 0xb9, 0x89, 0x4d, 0xf7, 0xa6, 0x36, 0xbd, 0xdf, 0xdb,
	0x3c, 0xea, 0xca, 0x19, 0x1f, 0x82, 0x27, 0xae, 0x6b, 0xc6, 0x57, 0xd6, 0xda, 0xd9, 0x2c, 0xa3,
	0x49, 0x78, 0xa1, 0x5a, 0x48, 0x93, 0xe4, 0xaa, 0x2b, 0x75, 0x7c, 0x1f, 0xbc, 0x4c, 0xb0, 0x42,
	0x07, 0x06, 0xd1, 0x22, 0x6c, 0x85, 0xf0, 0x44, 0xb0, 0x82, 0x6a, 0x6d, 0xff, 0x0d, 0xb8, 0x12,
	0x8e, 0x2e, 0xbf, 0x0f, 0xae, 0x4c, 0xa8, 0xae, 0xde, 0x1f, 0xa2, 0x38, 0xd9, 0x8e, 0x5f, 0xe2,
	0xfc, 0x92, 0xb5, 0x8f, 0xaf, 0x01, 0xd9, 0xef, 0xfb, 0x6a, 0x98, 0x8d, 0x3c, 0x33, 0xfa, 0x0a,
	0x1f, 0xc3, 0x3c, 0x6f, 0xab, 0xf6, 0xac, 0x16, 0x5c, 0x85, 0x05, 0x91, 0x1f, 0xee, 0x4a, 0x99,
	0xfe, 0x22, 0x93, 0xc8, 0x68, 0x42, 0xfc, 0x00, 0x66, 0x3b, 0x71, 0xbc, 0xaf, 0x93, 0xc8, 0x81,
	0xd1, 0xa9, 0x23, 0x33, 0xaf, 0x8d, 0xbe, 0xc5, 0x77, 0xc0, 0x95, 0xdf, 0xda, 0x26, 0xf3, 0x42,
	0x59, 0x98, 0x54, 0x51, 0xf8, 0x2e, 0x4c, 0xb8, 0xaa, 0x0c, 0xf5, 0x08, 0x41, 0x34, 0x0d, 0x75,
	0xa1, 0xd0, 0x96, 0x26, 0x4f, 0x8c, 0x8e, 0xc7, 0x47, 0x00, 0xa9, 0x02, 0xc6, 0x9d, 0xa6, 0x61,
	0xab, 0x1b, 0x12, 0x79, 0x6c, 0x0c, 0x06, 0x79, 0x86, 0x96, 0x86, 0x3b, 0x5a, 0x9a, 0x3c, 0xef,
	0xe6, 0xc6, 0xe8, 0x8b, 0xee, 0xc1, 0x4c, 0xfa, 0x54, 0xe7, 0xd9, 0xa6, 0xfd, 0x8e, 0x26, 0xaf,
	0xba, 0x41, 0xf3, 0x37, 0x17, 0x7d, 0x54, 0x80, 0xa7, 0xfe, 0x5f, 0x4e, 0x4b, 0xb5, 0x38, 0xcd,
	0xe4, 0xb4, 0xc5, 0xb0, 0x54, 0xa8, 0x1b, 0x5f, 0xc8, 0xc2, 0xff, 0xc3, 0x42, 0x71, 0xbb, 0x4f,
	0x41, 0x36, 0xde, 0x83, 0x40, 0x51, 0xfa, 0x66, 0xc8, 0x91, 0xa3, 0x51, 0x11, 0xd2, 0x02, 0x72,
	0x3b, 0x5d, 0x1f, 0x8a, 0xbc, 0x9f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x15, 0xef, 0xe0, 0x8e, 0x3d,
	0x06, 0x00, 0x00,
}
