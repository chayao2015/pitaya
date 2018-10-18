// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pitaya.proto

package protos

import proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PitayaClient is the client API for Pitaya service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PitayaClient interface {
	Call(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	PushToUser(ctx context.Context, in *Push, opts ...grpc.CallOption) (*Response, error)
	SessionBindRemote(ctx context.Context, in *BindMsg, opts ...grpc.CallOption) (*Response, error)
	KickUser(ctx context.Context, in *KickMsg, opts ...grpc.CallOption) (*KickAnswer, error)
}

type pitayaClient struct {
	cc *grpc.ClientConn
}

func NewPitayaClient(cc *grpc.ClientConn) PitayaClient {
	return &pitayaClient{cc}
}

func (c *pitayaClient) Call(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/protos.Pitaya/Call", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pitayaClient) PushToUser(ctx context.Context, in *Push, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/protos.Pitaya/PushToUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pitayaClient) SessionBindRemote(ctx context.Context, in *BindMsg, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/protos.Pitaya/SessionBindRemote", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pitayaClient) KickUser(ctx context.Context, in *KickMsg, opts ...grpc.CallOption) (*KickAnswer, error) {
	out := new(KickAnswer)
	err := c.cc.Invoke(ctx, "/protos.Pitaya/KickUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PitayaServer is the server API for Pitaya service.
type PitayaServer interface {
	Call(context.Context, *Request) (*Response, error)
	PushToUser(context.Context, *Push) (*Response, error)
	SessionBindRemote(context.Context, *BindMsg) (*Response, error)
	KickUser(context.Context, *KickMsg) (*KickAnswer, error)
}

func RegisterPitayaServer(s *grpc.Server, srv PitayaServer) {
	s.RegisterService(&_Pitaya_serviceDesc, srv)
}

func _Pitaya_Call_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PitayaServer).Call(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Pitaya/Call",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PitayaServer).Call(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pitaya_PushToUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Push)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PitayaServer).PushToUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Pitaya/PushToUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PitayaServer).PushToUser(ctx, req.(*Push))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pitaya_SessionBindRemote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BindMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PitayaServer).SessionBindRemote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Pitaya/SessionBindRemote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PitayaServer).SessionBindRemote(ctx, req.(*BindMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pitaya_KickUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KickMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PitayaServer).KickUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Pitaya/KickUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PitayaServer).KickUser(ctx, req.(*KickMsg))
	}
	return interceptor(ctx, in, info, handler)
}

var _Pitaya_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protos.Pitaya",
	HandlerType: (*PitayaServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Call",
			Handler:    _Pitaya_Call_Handler,
		},
		{
			MethodName: "PushToUser",
			Handler:    _Pitaya_PushToUser_Handler,
		},
		{
			MethodName: "SessionBindRemote",
			Handler:    _Pitaya_SessionBindRemote_Handler,
		},
		{
			MethodName: "KickUser",
			Handler:    _Pitaya_KickUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pitaya.proto",
}

func init() { proto.RegisterFile("pitaya.proto", fileDescriptor_pitaya_bd05e00ce1d8d9c1) }

var fileDescriptor_pitaya_bd05e00ce1d8d9c1 = []byte{
	// 218 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0xc8, 0x2c, 0x49,
	0xac, 0x4c, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x03, 0x53, 0xc5, 0x52, 0xbc, 0x45,
	0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x10, 0x61, 0x29, 0xbe, 0xa2, 0xd4, 0xe2, 0x82, 0xfc, 0xbc,
	0xe2, 0x54, 0x28, 0x9f, 0xab, 0xa0, 0xb4, 0x38, 0x03, 0xc6, 0x4e, 0xca, 0xcc, 0x4b, 0x81, 0xb1,
	0xb3, 0x33, 0x93, 0xb3, 0x21, 0x6c, 0xa3, 0x4b, 0x8c, 0x5c, 0x6c, 0x01, 0x60, 0xb3, 0x85, 0xb4,
	0xb9, 0x58, 0x9c, 0x13, 0x73, 0x72, 0x84, 0xf8, 0x21, 0x52, 0xc5, 0x7a, 0x41, 0x10, 0xd3, 0xa5,
	0x04, 0x10, 0x02, 0x10, 0xf3, 0x95, 0x18, 0x84, 0xf4, 0xb8, 0xb8, 0x02, 0x4a, 0x8b, 0x33, 0x42,
	0xf2, 0x43, 0x8b, 0x53, 0x8b, 0x84, 0x78, 0x60, 0x2a, 0x40, 0x62, 0x58, 0xd5, 0x5b, 0x70, 0x09,
	0x06, 0xa7, 0x16, 0x17, 0x67, 0xe6, 0xe7, 0x39, 0x65, 0xe6, 0xa5, 0x04, 0xa5, 0xe6, 0xe6, 0x97,
	0xa4, 0x22, 0x6c, 0x02, 0x89, 0xf9, 0x16, 0xa7, 0x63, 0xd5, 0x69, 0xc8, 0xc5, 0xe1, 0x9d, 0x99,
	0x9c, 0x0d, 0xb6, 0x07, 0xae, 0x01, 0x24, 0x02, 0xd2, 0x20, 0x84, 0x2c, 0xe0, 0x98, 0x57, 0x5c,
	0x9e, 0x5a, 0xa4, 0xc4, 0xe0, 0x24, 0x70, 0xe2, 0x91, 0x1c, 0xe3, 0x85, 0x47, 0x72, 0x8c, 0x0f,
	0x1e, 0xc9, 0x31, 0x4e, 0x78, 0x2c, 0xc7, 0x90, 0x04, 0x09, 0x31, 0x63, 0x40, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x8f, 0xf0, 0x87, 0x00, 0x48, 0x01, 0x00, 0x00,
}
