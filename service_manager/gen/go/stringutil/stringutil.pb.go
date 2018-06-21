// Code generated by protoc-gen-go. DO NOT EDIT.
// source: stringutil.proto

package strutil

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

type String struct {
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *String) Reset()         { *m = String{} }
func (m *String) String() string { return proto.CompactTextString(m) }
func (*String) ProtoMessage()    {}
func (*String) Descriptor() ([]byte, []int) {
	return fileDescriptor_stringutil_32ce04a3d4c026bb, []int{0}
}
func (m *String) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_String.Unmarshal(m, b)
}
func (m *String) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_String.Marshal(b, m, deterministic)
}
func (dst *String) XXX_Merge(src proto.Message) {
	xxx_messageInfo_String.Merge(dst, src)
}
func (m *String) XXX_Size() int {
	return xxx_messageInfo_String.Size(m)
}
func (m *String) XXX_DiscardUnknown() {
	xxx_messageInfo_String.DiscardUnknown(m)
}

var xxx_messageInfo_String proto.InternalMessageInfo

func (m *String) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*String)(nil), "strutil.String")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// StringUtilClient is the client API for StringUtil service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StringUtilClient interface {
	Reverse(ctx context.Context, in *String, opts ...grpc.CallOption) (*String, error)
}

type stringUtilClient struct {
	cc *grpc.ClientConn
}

func NewStringUtilClient(cc *grpc.ClientConn) StringUtilClient {
	return &stringUtilClient{cc}
}

func (c *stringUtilClient) Reverse(ctx context.Context, in *String, opts ...grpc.CallOption) (*String, error) {
	out := new(String)
	err := c.cc.Invoke(ctx, "/strutil.StringUtil/Reverse", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StringUtilServer is the server API for StringUtil service.
type StringUtilServer interface {
	Reverse(context.Context, *String) (*String, error)
}

func RegisterStringUtilServer(s *grpc.Server, srv StringUtilServer) {
	s.RegisterService(&_StringUtil_serviceDesc, srv)
}

func _StringUtil_Reverse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(String)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StringUtilServer).Reverse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/strutil.StringUtil/Reverse",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StringUtilServer).Reverse(ctx, req.(*String))
	}
	return interceptor(ctx, in, info, handler)
}

var _StringUtil_serviceDesc = grpc.ServiceDesc{
	ServiceName: "strutil.StringUtil",
	HandlerType: (*StringUtilServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Reverse",
			Handler:    _StringUtil_Reverse_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "stringutil.proto",
}

func init() { proto.RegisterFile("stringutil.proto", fileDescriptor_stringutil_32ce04a3d4c026bb) }

var fileDescriptor_stringutil_32ce04a3d4c026bb = []byte{
	// 112 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x28, 0x2e, 0x29, 0xca,
	0xcc, 0x4b, 0x2f, 0x2d, 0xc9, 0xcc, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2f, 0x2e,
	0x29, 0x02, 0x71, 0x95, 0xe4, 0xb8, 0xd8, 0x82, 0xc1, 0x92, 0x42, 0x22, 0x5c, 0xac, 0x65, 0x89,
	0x39, 0xa5, 0xa9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x10, 0x8e, 0x91, 0x35, 0x17, 0x17,
	0x44, 0x3e, 0xb4, 0x24, 0x33, 0x47, 0x48, 0x97, 0x8b, 0x3d, 0x28, 0xb5, 0x2c, 0xb5, 0xa8, 0x38,
	0x55, 0x88, 0x5f, 0x0f, 0x6a, 0x84, 0x1e, 0x44, 0x5e, 0x0a, 0x5d, 0x40, 0x89, 0x21, 0x89, 0x0d,
	0x6c, 0x99, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xd0, 0x33, 0xaa, 0x00, 0x80, 0x00, 0x00, 0x00,
}
