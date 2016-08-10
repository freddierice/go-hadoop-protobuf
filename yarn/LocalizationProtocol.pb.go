// Code generated by protoc-gen-go.
// source: gopkg.in/freddierice/go-hproto.v1/yarn/LocalizationProtocol.proto
// DO NOT EDIT!

package hproto_yarn

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

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for LocalizationProtocolService service

type LocalizationProtocolServiceClient interface {
	Heartbeat(ctx context.Context, in *LocalizerStatusProto, opts ...grpc.CallOption) (*LocalizerHeartbeatResponseProto, error)
}

type localizationProtocolServiceClient struct {
	cc *grpc.ClientConn
}

func NewLocalizationProtocolServiceClient(cc *grpc.ClientConn) LocalizationProtocolServiceClient {
	return &localizationProtocolServiceClient{cc}
}

func (c *localizationProtocolServiceClient) Heartbeat(ctx context.Context, in *LocalizerStatusProto, opts ...grpc.CallOption) (*LocalizerHeartbeatResponseProto, error) {
	out := new(LocalizerHeartbeatResponseProto)
	err := grpc.Invoke(ctx, "/hproto.yarn.LocalizationProtocolService/heartbeat", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for LocalizationProtocolService service

type LocalizationProtocolServiceServer interface {
	Heartbeat(context.Context, *LocalizerStatusProto) (*LocalizerHeartbeatResponseProto, error)
}

func RegisterLocalizationProtocolServiceServer(s *grpc.Server, srv LocalizationProtocolServiceServer) {
	s.RegisterService(&_LocalizationProtocolService_serviceDesc, srv)
}

func _LocalizationProtocolService_Heartbeat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LocalizerStatusProto)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LocalizationProtocolServiceServer).Heartbeat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hproto.yarn.LocalizationProtocolService/Heartbeat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LocalizationProtocolServiceServer).Heartbeat(ctx, req.(*LocalizerStatusProto))
	}
	return interceptor(ctx, in, info, handler)
}

var _LocalizationProtocolService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "hproto.yarn.LocalizationProtocolService",
	HandlerType: (*LocalizationProtocolServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "heartbeat",
			Handler:    _LocalizationProtocolService_Heartbeat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor15,
}

func init() {
	proto.RegisterFile("gopkg.in/freddierice/go-hproto.v1/yarn/LocalizationProtocol.proto", fileDescriptor15)
}

var fileDescriptor15 = []byte{
	// 212 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x8e, 0xb1, 0x4e, 0xc4, 0x30,
	0x0c, 0x86, 0x95, 0x91, 0xb0, 0x55, 0x4c, 0x85, 0x85, 0x1d, 0x12, 0xc1, 0xc8, 0x46, 0x27, 0x06,
	0x06, 0x44, 0x57, 0xa4, 0xca, 0xa4, 0x26, 0x8d, 0x28, 0x71, 0xe4, 0x84, 0x4a, 0xc0, 0x0b, 0xf0,
	0x18, 0x3c, 0x2a, 0x4d, 0xda, 0x93, 0xee, 0xa4, 0x0e, 0xb7, 0x44, 0x56, 0xfc, 0x7d, 0xfe, 0x7f,
	0x79, 0x6f, 0x29, 0xbc, 0x5b, 0xe5, 0xbc, 0x7e, 0x63, 0xec, 0x7b, 0x87, 0xec, 0x0c, 0x6a, 0x4b,
	0xd7, 0x43, 0x60, 0x4a, 0xa4, 0xa6, 0x1b, 0xfd, 0x05, 0xec, 0xf5, 0x23, 0x19, 0x18, 0xdd, 0x37,
	0x24, 0x47, 0xfe, 0x29, 0x6f, 0x0c, 0x8d, 0xaa, 0x20, 0xd5, 0xe9, 0x8a, 0x66, 0xae, 0x6e, 0x8f,
	0xbc, 0x97, 0x9f, 0x2e, 0x22, 0x4f, 0xc8, 0x9d, 0xa7, 0x1e, 0x3f, 0xc0, 0x83, 0x9d, 0xe7, 0xfc,
	0x35, 0x1b, 0x5d, 0xa1, 0xe3, 0x92, 0x70, 0xfb, 0x23, 0xcf, 0xb7, 0xf2, 0xdb, 0x85, 0xad, 0x5e,
	0xe4, 0xc9, 0x80, 0xc0, 0xe9, 0x15, 0x21, 0x55, 0x97, 0x6a, 0xaf, 0x8e, 0x5a, 0x35, 0xe4, 0x36,
	0x41, 0xfa, 0x8c, 0xc5, 0xac, 0xaf, 0xb6, 0x91, 0x87, 0xdd, 0x8d, 0x67, 0x8c, 0x81, 0x7c, 0xc4,
	0x42, 0x37, 0x77, 0xf2, 0x82, 0xd8, 0x2a, 0x08, 0x60, 0x06, 0x3c, 0x30, 0xcb, 0xd8, 0x9c, 0x6d,
	0x55, 0xfb, 0x15, 0xe2, 0x4f, 0x88, 0xff, 0x00, 0x00, 0x00, 0xff, 0xff, 0xab, 0x0a, 0xf7, 0xd6,
	0x5e, 0x01, 0x00, 0x00,
}
