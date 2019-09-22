// Code generated by protoc-gen-go. DO NOT EDIT.
// source: remcap.proto

package remcappb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Packet struct {
	Data                 []byte   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	InterfaceIndex       int32    `protobuf:"varint,2,opt,name=interface_index,json=interfaceIndex,proto3" json:"interface_index,omitempty"`
	ExtIP                string   `protobuf:"bytes,3,opt,name=extIP,proto3" json:"extIP,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Packet) Reset()         { *m = Packet{} }
func (m *Packet) String() string { return proto.CompactTextString(m) }
func (*Packet) ProtoMessage()    {}
func (*Packet) Descriptor() ([]byte, []int) {
	return fileDescriptor_959640b0d5eff11b, []int{0}
}

func (m *Packet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Packet.Unmarshal(m, b)
}
func (m *Packet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Packet.Marshal(b, m, deterministic)
}
func (m *Packet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Packet.Merge(m, src)
}
func (m *Packet) XXX_Size() int {
	return xxx_messageInfo_Packet.Size(m)
}
func (m *Packet) XXX_DiscardUnknown() {
	xxx_messageInfo_Packet.DiscardUnknown(m)
}

var xxx_messageInfo_Packet proto.InternalMessageInfo

func (m *Packet) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Packet) GetInterfaceIndex() int32 {
	if m != nil {
		return m.InterfaceIndex
	}
	return 0
}

func (m *Packet) GetExtIP() string {
	if m != nil {
		return m.ExtIP
	}
	return ""
}

type Summary struct {
	StartTime            int64    `protobuf:"varint,1,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EndTime              int64    `protobuf:"varint,2,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	PacketsCaptured      int64    `protobuf:"varint,3,opt,name=packets_captured,json=packetsCaptured,proto3" json:"packets_captured,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Summary) Reset()         { *m = Summary{} }
func (m *Summary) String() string { return proto.CompactTextString(m) }
func (*Summary) ProtoMessage()    {}
func (*Summary) Descriptor() ([]byte, []int) {
	return fileDescriptor_959640b0d5eff11b, []int{1}
}

func (m *Summary) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Summary.Unmarshal(m, b)
}
func (m *Summary) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Summary.Marshal(b, m, deterministic)
}
func (m *Summary) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Summary.Merge(m, src)
}
func (m *Summary) XXX_Size() int {
	return xxx_messageInfo_Summary.Size(m)
}
func (m *Summary) XXX_DiscardUnknown() {
	xxx_messageInfo_Summary.DiscardUnknown(m)
}

var xxx_messageInfo_Summary proto.InternalMessageInfo

func (m *Summary) GetStartTime() int64 {
	if m != nil {
		return m.StartTime
	}
	return 0
}

func (m *Summary) GetEndTime() int64 {
	if m != nil {
		return m.EndTime
	}
	return 0
}

func (m *Summary) GetPacketsCaptured() int64 {
	if m != nil {
		return m.PacketsCaptured
	}
	return 0
}

func init() {
	proto.RegisterType((*Packet)(nil), "remcap.Packet")
	proto.RegisterType((*Summary)(nil), "remcap.Summary")
}

func init() { proto.RegisterFile("remcap.proto", fileDescriptor_959640b0d5eff11b) }

var fileDescriptor_959640b0d5eff11b = []byte{
	// 235 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0x90, 0x41, 0x4b, 0xc3, 0x40,
	0x10, 0x85, 0x49, 0x63, 0xd2, 0x76, 0x28, 0x8d, 0x0c, 0x1e, 0xa2, 0x20, 0x84, 0x5e, 0x8c, 0x1e,
	0x7a, 0x50, 0x7f, 0x81, 0x3d, 0xf5, 0x56, 0xb6, 0x9e, 0xf4, 0x10, 0xa6, 0xd9, 0x09, 0x2c, 0xb2,
	0xdb, 0x65, 0x3b, 0x85, 0xfa, 0xef, 0xc5, 0xdd, 0xe8, 0x6d, 0xde, 0xf7, 0x0e, 0xef, 0xbd, 0x81,
	0x45, 0x60, 0xdb, 0x93, 0x5f, 0xfb, 0x70, 0x94, 0x23, 0x96, 0x49, 0xad, 0x3e, 0xa1, 0xdc, 0x51,
	0xff, 0xc5, 0x82, 0x08, 0x57, 0x9a, 0x84, 0xea, 0xac, 0xc9, 0xda, 0x85, 0x8a, 0x37, 0x3e, 0x40,
	0x65, 0x9c, 0x70, 0x18, 0xa8, 0xe7, 0xce, 0x38, 0xcd, 0x97, 0x7a, 0xd2, 0x64, 0x6d, 0xa1, 0x96,
	0xff, 0x78, 0xfb, 0x4b, 0xf1, 0x06, 0x0a, 0xbe, 0xc8, 0x76, 0x57, 0xe7, 0x4d, 0xd6, 0xce, 0x55,
	0x12, 0x2b, 0x07, 0xd3, 0xfd, 0xd9, 0x5a, 0x0a, 0xdf, 0x78, 0x0f, 0x70, 0x12, 0x0a, 0xd2, 0x89,
	0xb1, 0x1c, 0x33, 0x72, 0x35, 0x8f, 0xe4, 0xdd, 0x58, 0xc6, 0x5b, 0x98, 0xb1, 0xd3, 0xc9, 0x9c,
	0x44, 0x73, 0xca, 0x4e, 0x47, 0xeb, 0x11, 0xae, 0x7d, 0x6c, 0x78, 0xea, 0x7a, 0xf2, 0x72, 0x0e,
	0xac, 0x63, 0x4a, 0xae, 0xaa, 0x91, 0x6f, 0x46, 0xfc, 0xfc, 0x0a, 0xa5, 0x62, 0xbb, 0x21, 0x8f,
	0x4f, 0x50, 0xec, 0x9d, 0x19, 0x06, 0x5c, 0xae, 0xc7, 0xd9, 0x69, 0xe5, 0x5d, 0xf5, 0xa7, 0xc7,
	0x62, 0x6d, 0xf6, 0x06, 0x1f, 0xb3, 0xc4, 0xfc, 0xe1, 0x50, 0xc6, 0xef, 0xbc, 0xfc, 0x04, 0x00,
	0x00, 0xff, 0xff, 0xeb, 0xd4, 0xed, 0xd0, 0x2d, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RemCapClient is the client API for RemCap service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RemCapClient interface {
	Sniff(ctx context.Context, opts ...grpc.CallOption) (RemCap_SniffClient, error)
}

type remCapClient struct {
	cc *grpc.ClientConn
}

func NewRemCapClient(cc *grpc.ClientConn) RemCapClient {
	return &remCapClient{cc}
}

func (c *remCapClient) Sniff(ctx context.Context, opts ...grpc.CallOption) (RemCap_SniffClient, error) {
	stream, err := c.cc.NewStream(ctx, &_RemCap_serviceDesc.Streams[0], "/remcap.RemCap/Sniff", opts...)
	if err != nil {
		return nil, err
	}
	x := &remCapSniffClient{stream}
	return x, nil
}

type RemCap_SniffClient interface {
	Send(*Packet) error
	CloseAndRecv() (*Summary, error)
	grpc.ClientStream
}

type remCapSniffClient struct {
	grpc.ClientStream
}

func (x *remCapSniffClient) Send(m *Packet) error {
	return x.ClientStream.SendMsg(m)
}

func (x *remCapSniffClient) CloseAndRecv() (*Summary, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Summary)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// RemCapServer is the server API for RemCap service.
type RemCapServer interface {
	Sniff(RemCap_SniffServer) error
}

func RegisterRemCapServer(s *grpc.Server, srv RemCapServer) {
	s.RegisterService(&_RemCap_serviceDesc, srv)
}

func _RemCap_Sniff_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RemCapServer).Sniff(&remCapSniffServer{stream})
}

type RemCap_SniffServer interface {
	SendAndClose(*Summary) error
	Recv() (*Packet, error)
	grpc.ServerStream
}

type remCapSniffServer struct {
	grpc.ServerStream
}

func (x *remCapSniffServer) SendAndClose(m *Summary) error {
	return x.ServerStream.SendMsg(m)
}

func (x *remCapSniffServer) Recv() (*Packet, error) {
	m := new(Packet)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _RemCap_serviceDesc = grpc.ServiceDesc{
	ServiceName: "remcap.RemCap",
	HandlerType: (*RemCapServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Sniff",
			Handler:       _RemCap_Sniff_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "remcap.proto",
}
