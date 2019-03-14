// Code generated by protoc-gen-go. DO NOT EDIT.
// source: core.proto

package core

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_core_084e9cef696b9f8e, []int{0}
}
func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (dst *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(dst, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type ServiceStatus struct {
	Version              string   `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServiceStatus) Reset()         { *m = ServiceStatus{} }
func (m *ServiceStatus) String() string { return proto.CompactTextString(m) }
func (*ServiceStatus) ProtoMessage()    {}
func (*ServiceStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_core_084e9cef696b9f8e, []int{1}
}
func (m *ServiceStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceStatus.Unmarshal(m, b)
}
func (m *ServiceStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceStatus.Marshal(b, m, deterministic)
}
func (dst *ServiceStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceStatus.Merge(dst, src)
}
func (m *ServiceStatus) XXX_Size() int {
	return xxx_messageInfo_ServiceStatus.Size(m)
}
func (m *ServiceStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceStatus.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceStatus proto.InternalMessageInfo

func (m *ServiceStatus) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *ServiceStatus) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type ServiceStatistics struct {
	Metrics              []byte   `protobuf:"bytes,1,opt,name=metrics,proto3" json:"metrics,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServiceStatistics) Reset()         { *m = ServiceStatistics{} }
func (m *ServiceStatistics) String() string { return proto.CompactTextString(m) }
func (*ServiceStatistics) ProtoMessage()    {}
func (*ServiceStatistics) Descriptor() ([]byte, []int) {
	return fileDescriptor_core_084e9cef696b9f8e, []int{2}
}
func (m *ServiceStatistics) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceStatistics.Unmarshal(m, b)
}
func (m *ServiceStatistics) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceStatistics.Marshal(b, m, deterministic)
}
func (dst *ServiceStatistics) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceStatistics.Merge(dst, src)
}
func (m *ServiceStatistics) XXX_Size() int {
	return xxx_messageInfo_ServiceStatistics.Size(m)
}
func (m *ServiceStatistics) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceStatistics.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceStatistics proto.InternalMessageInfo

func (m *ServiceStatistics) GetMetrics() []byte {
	if m != nil {
		return m.Metrics
	}
	return nil
}

func init() {
	proto.RegisterType((*Empty)(nil), "core.Empty")
	proto.RegisterType((*ServiceStatus)(nil), "core.ServiceStatus")
	proto.RegisterType((*ServiceStatistics)(nil), "core.ServiceStatistics")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TemporalCoreClient is the client API for TemporalCore service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TemporalCoreClient interface {
	Status(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ServiceStatus, error)
	Statistics(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ServiceStatistics, error)
}

type temporalCoreClient struct {
	cc *grpc.ClientConn
}

func NewTemporalCoreClient(cc *grpc.ClientConn) TemporalCoreClient {
	return &temporalCoreClient{cc}
}

func (c *temporalCoreClient) Status(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ServiceStatus, error) {
	out := new(ServiceStatus)
	err := c.cc.Invoke(ctx, "/core.TemporalCore/Status", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *temporalCoreClient) Statistics(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ServiceStatistics, error) {
	out := new(ServiceStatistics)
	err := c.cc.Invoke(ctx, "/core.TemporalCore/Statistics", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TemporalCoreServer is the server API for TemporalCore service.
type TemporalCoreServer interface {
	Status(context.Context, *Empty) (*ServiceStatus, error)
	Statistics(context.Context, *Empty) (*ServiceStatistics, error)
}

func RegisterTemporalCoreServer(s *grpc.Server, srv TemporalCoreServer) {
	s.RegisterService(&_TemporalCore_serviceDesc, srv)
}

func _TemporalCore_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TemporalCoreServer).Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/core.TemporalCore/Status",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TemporalCoreServer).Status(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _TemporalCore_Statistics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TemporalCoreServer).Statistics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/core.TemporalCore/Statistics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TemporalCoreServer).Statistics(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _TemporalCore_serviceDesc = grpc.ServiceDesc{
	ServiceName: "core.TemporalCore",
	HandlerType: (*TemporalCoreServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Status",
			Handler:    _TemporalCore_Status_Handler,
		},
		{
			MethodName: "Statistics",
			Handler:    _TemporalCore_Statistics_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "core.proto",
}

func init() { proto.RegisterFile("core.proto", fileDescriptor_core_084e9cef696b9f8e) }

var fileDescriptor_core_084e9cef696b9f8e = []byte{
	// 250 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0xd1, 0x4a, 0xc3, 0x30,
	0x18, 0x85, 0xe9, 0xd0, 0x0d, 0x7f, 0x27, 0xb2, 0x94, 0xb1, 0x32, 0xbd, 0x90, 0x5c, 0x79, 0x63,
	0x03, 0xee, 0x0d, 0x2c, 0x5e, 0x0b, 0x9b, 0x2f, 0x10, 0x63, 0x28, 0x81, 0xb6, 0x7f, 0xc9, 0xff,
	0xaf, 0xe0, 0xad, 0xaf, 0xe0, 0x2b, 0xf8, 0x46, 0xbe, 0x82, 0x0f, 0x22, 0x4d, 0x3a, 0x68, 0xf1,
	0x2e, 0x87, 0x73, 0xf2, 0x9d, 0x93, 0x00, 0x18, 0xf4, 0x36, 0x6f, 0x3d, 0x32, 0x8a, 0xb3, 0xfe,
	0xbc, 0xbd, 0x2d, 0x11, 0xcb, 0xca, 0x2a, 0xdd, 0x3a, 0xa5, 0x9b, 0x06, 0x59, 0xb3, 0xc3, 0x86,
	0x62, 0x46, 0x2e, 0xe0, 0xfc, 0xb9, 0x6e, 0xf9, 0x43, 0x16, 0x70, 0x75, 0xb0, 0xbe, 0x73, 0xc6,
	0x1e, 0x58, 0xf3, 0x91, 0x44, 0x06, 0x8b, 0xce, 0x7a, 0x72, 0xd8, 0x64, 0xc9, 0x5d, 0x72, 0x7f,
	0xb1, 0x3f, 0xc9, 0xde, 0xa9, 0x2d, 0x91, 0x2e, 0x6d, 0x36, 0x8b, 0xce, 0x20, 0xe5, 0x03, 0xac,
	0x46, 0x10, 0x47, 0xec, 0x0c, 0xc5, 0x38, 0x7b, 0x67, 0x28, 0x80, 0x96, 0xfb, 0x93, 0x7c, 0xfc,
	0x4e, 0x60, 0xf9, 0x6a, 0xeb, 0x16, 0xbd, 0xae, 0x0a, 0xf4, 0x56, 0x14, 0x30, 0x1f, 0xda, 0x2f,
	0xf3, 0xf0, 0x90, 0xb0, 0x6d, 0x9b, 0x46, 0x31, 0xd9, 0x27, 0x37, 0x9f, 0x3f, 0xbf, 0x5f, 0xb3,
	0x95, 0xb8, 0x56, 0xdd, 0x4e, 0xf5, 0xbe, 0xa2, 0x78, 0xf5, 0x05, 0x60, 0xd4, 0x3e, 0x01, 0x6d,
	0xfe, 0x81, 0x62, 0x4a, 0xde, 0x04, 0xd8, 0x5a, 0xa4, 0x13, 0x58, 0x34, 0x9f, 0xd6, 0x90, 0x9a,
	0x0a, 0x8f, 0xef, 0x39, 0x0f, 0x5b, 0x03, 0xe5, 0x6d, 0x1e, 0x7e, 0x70, 0xf7, 0x17, 0x00, 0x00,
	0xff, 0xff, 0x69, 0x99, 0x60, 0xde, 0x73, 0x01, 0x00, 0x00,
}