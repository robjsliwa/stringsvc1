// Code generated by protoc-gen-go.
// source: stringsvc1.proto
// DO NOT EDIT!

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	stringsvc1.proto

It has these top-level messages:
	UppercaseRequest
	UppercaseResponse
	CountRequest
	CountResponse
*/
package pb

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

type UppercaseRequest struct {
	InputString string `protobuf:"bytes,1,opt,name=inputString" json:"inputString,omitempty"`
}

func (m *UppercaseRequest) Reset()                    { *m = UppercaseRequest{} }
func (m *UppercaseRequest) String() string            { return proto.CompactTextString(m) }
func (*UppercaseRequest) ProtoMessage()               {}
func (*UppercaseRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *UppercaseRequest) GetInputString() string {
	if m != nil {
		return m.InputString
	}
	return ""
}

type UppercaseResponse struct {
	UppercasedString string `protobuf:"bytes,1,opt,name=uppercasedString" json:"uppercasedString,omitempty"`
	Err              string `protobuf:"bytes,2,opt,name=err" json:"err,omitempty"`
}

func (m *UppercaseResponse) Reset()                    { *m = UppercaseResponse{} }
func (m *UppercaseResponse) String() string            { return proto.CompactTextString(m) }
func (*UppercaseResponse) ProtoMessage()               {}
func (*UppercaseResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *UppercaseResponse) GetUppercasedString() string {
	if m != nil {
		return m.UppercasedString
	}
	return ""
}

func (m *UppercaseResponse) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

type CountRequest struct {
	InputString string `protobuf:"bytes,1,opt,name=inputString" json:"inputString,omitempty"`
}

func (m *CountRequest) Reset()                    { *m = CountRequest{} }
func (m *CountRequest) String() string            { return proto.CompactTextString(m) }
func (*CountRequest) ProtoMessage()               {}
func (*CountRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *CountRequest) GetInputString() string {
	if m != nil {
		return m.InputString
	}
	return ""
}

type CountResponse struct {
	Length int64  `protobuf:"varint,1,opt,name=length" json:"length,omitempty"`
	Err    string `protobuf:"bytes,2,opt,name=err" json:"err,omitempty"`
}

func (m *CountResponse) Reset()                    { *m = CountResponse{} }
func (m *CountResponse) String() string            { return proto.CompactTextString(m) }
func (*CountResponse) ProtoMessage()               {}
func (*CountResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *CountResponse) GetLength() int64 {
	if m != nil {
		return m.Length
	}
	return 0
}

func (m *CountResponse) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

func init() {
	proto.RegisterType((*UppercaseRequest)(nil), "pb.UppercaseRequest")
	proto.RegisterType((*UppercaseResponse)(nil), "pb.UppercaseResponse")
	proto.RegisterType((*CountRequest)(nil), "pb.CountRequest")
	proto.RegisterType((*CountResponse)(nil), "pb.CountResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for StringService service

type StringServiceClient interface {
	Uppercase(ctx context.Context, in *UppercaseRequest, opts ...grpc.CallOption) (*UppercaseResponse, error)
	Count(ctx context.Context, in *CountRequest, opts ...grpc.CallOption) (*CountResponse, error)
}

type stringServiceClient struct {
	cc *grpc.ClientConn
}

func NewStringServiceClient(cc *grpc.ClientConn) StringServiceClient {
	return &stringServiceClient{cc}
}

func (c *stringServiceClient) Uppercase(ctx context.Context, in *UppercaseRequest, opts ...grpc.CallOption) (*UppercaseResponse, error) {
	out := new(UppercaseResponse)
	err := grpc.Invoke(ctx, "/pb.StringService/Uppercase", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stringServiceClient) Count(ctx context.Context, in *CountRequest, opts ...grpc.CallOption) (*CountResponse, error) {
	out := new(CountResponse)
	err := grpc.Invoke(ctx, "/pb.StringService/Count", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for StringService service

type StringServiceServer interface {
	Uppercase(context.Context, *UppercaseRequest) (*UppercaseResponse, error)
	Count(context.Context, *CountRequest) (*CountResponse, error)
}

func RegisterStringServiceServer(s *grpc.Server, srv StringServiceServer) {
	s.RegisterService(&_StringService_serviceDesc, srv)
}

func _StringService_Uppercase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UppercaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StringServiceServer).Uppercase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.StringService/Uppercase",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StringServiceServer).Uppercase(ctx, req.(*UppercaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StringService_Count_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StringServiceServer).Count(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.StringService/Count",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StringServiceServer).Count(ctx, req.(*CountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _StringService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.StringService",
	HandlerType: (*StringServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Uppercase",
			Handler:    _StringService_Uppercase_Handler,
		},
		{
			MethodName: "Count",
			Handler:    _StringService_Count_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "stringsvc1.proto",
}

func init() { proto.RegisterFile("stringsvc1.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 220 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x28, 0x2e, 0x29, 0xca,
	0xcc, 0x4b, 0x2f, 0x2e, 0x4b, 0x36, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48,
	0x52, 0x32, 0xe1, 0x12, 0x08, 0x2d, 0x28, 0x48, 0x2d, 0x4a, 0x4e, 0x2c, 0x4e, 0x0d, 0x4a, 0x2d,
	0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x52, 0xe0, 0xe2, 0xce, 0xcc, 0x2b, 0x28, 0x2d, 0x09, 0x06, 0x6b,
	0x90, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x42, 0x16, 0x52, 0x0a, 0xe4, 0x12, 0x44, 0xd2, 0x55,
	0x5c, 0x90, 0x9f, 0x57, 0x9c, 0x2a, 0xa4, 0xc5, 0x25, 0x50, 0x0a, 0x13, 0x4c, 0x41, 0xd1, 0x8b,
	0x21, 0x2e, 0x24, 0xc0, 0xc5, 0x9c, 0x5a, 0x54, 0x24, 0xc1, 0x04, 0x96, 0x06, 0x31, 0x95, 0x0c,
	0xb8, 0x78, 0x9c, 0xf3, 0x4b, 0xf3, 0x4a, 0x88, 0x77, 0x84, 0x25, 0x17, 0x2f, 0x54, 0x07, 0xd4,
	0x01, 0x62, 0x5c, 0x6c, 0x39, 0xa9, 0x79, 0xe9, 0x25, 0x19, 0x60, 0xd5, 0xcc, 0x41, 0x50, 0x1e,
	0xa6, 0x65, 0x46, 0xd5, 0x5c, 0xbc, 0x10, 0x43, 0x82, 0x53, 0x8b, 0xca, 0x32, 0x93, 0x53, 0x85,
	0xac, 0xb8, 0x38, 0xe1, 0x1e, 0x12, 0x12, 0xd1, 0x2b, 0x48, 0xd2, 0x43, 0x0f, 0x15, 0x29, 0x51,
	0x34, 0x51, 0x88, 0xa5, 0x4a, 0x0c, 0x42, 0x7a, 0x5c, 0xac, 0x60, 0x77, 0x08, 0x09, 0x80, 0x54,
	0x20, 0x7b, 0x42, 0x4a, 0x10, 0x49, 0x04, 0xa6, 0x3e, 0x89, 0x0d, 0x1c, 0xfa, 0xc6, 0x80, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x12, 0x50, 0xf4, 0xb3, 0x91, 0x01, 0x00, 0x00,
}
