// Code generated by protoc-gen-go.
// source: gopkg.in/freddierice/go-hproto.v1/common/RefreshAuthorizationPolicyProtocol.proto
// DO NOT EDIT!

package hproto_common

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

// *
//  Refresh service acl request.
type RefreshServiceAclRequestProto struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *RefreshServiceAclRequestProto) Reset()                    { *m = RefreshServiceAclRequestProto{} }
func (m *RefreshServiceAclRequestProto) String() string            { return proto.CompactTextString(m) }
func (*RefreshServiceAclRequestProto) ProtoMessage()               {}
func (*RefreshServiceAclRequestProto) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{0} }

// *
// void response
type RefreshServiceAclResponseProto struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *RefreshServiceAclResponseProto) Reset()                    { *m = RefreshServiceAclResponseProto{} }
func (m *RefreshServiceAclResponseProto) String() string            { return proto.CompactTextString(m) }
func (*RefreshServiceAclResponseProto) ProtoMessage()               {}
func (*RefreshServiceAclResponseProto) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{1} }

func init() {
	proto.RegisterType((*RefreshServiceAclRequestProto)(nil), "hproto.common.RefreshServiceAclRequestProto")
	proto.RegisterType((*RefreshServiceAclResponseProto)(nil), "hproto.common.RefreshServiceAclResponseProto")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for RefreshAuthorizationPolicyProtocolService service

type RefreshAuthorizationPolicyProtocolServiceClient interface {
	// *
	// Refresh the service-level authorization policy in-effect.
	RefreshServiceAcl(ctx context.Context, in *RefreshServiceAclRequestProto, opts ...grpc.CallOption) (*RefreshServiceAclResponseProto, error)
}

type refreshAuthorizationPolicyProtocolServiceClient struct {
	cc *grpc.ClientConn
}

func NewRefreshAuthorizationPolicyProtocolServiceClient(cc *grpc.ClientConn) RefreshAuthorizationPolicyProtocolServiceClient {
	return &refreshAuthorizationPolicyProtocolServiceClient{cc}
}

func (c *refreshAuthorizationPolicyProtocolServiceClient) RefreshServiceAcl(ctx context.Context, in *RefreshServiceAclRequestProto, opts ...grpc.CallOption) (*RefreshServiceAclResponseProto, error) {
	out := new(RefreshServiceAclResponseProto)
	err := grpc.Invoke(ctx, "/hproto.common.RefreshAuthorizationPolicyProtocolService/refreshServiceAcl", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for RefreshAuthorizationPolicyProtocolService service

type RefreshAuthorizationPolicyProtocolServiceServer interface {
	// *
	// Refresh the service-level authorization policy in-effect.
	RefreshServiceAcl(context.Context, *RefreshServiceAclRequestProto) (*RefreshServiceAclResponseProto, error)
}

func RegisterRefreshAuthorizationPolicyProtocolServiceServer(s *grpc.Server, srv RefreshAuthorizationPolicyProtocolServiceServer) {
	s.RegisterService(&_RefreshAuthorizationPolicyProtocolService_serviceDesc, srv)
}

func _RefreshAuthorizationPolicyProtocolService_RefreshServiceAcl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefreshServiceAclRequestProto)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RefreshAuthorizationPolicyProtocolServiceServer).RefreshServiceAcl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hproto.common.RefreshAuthorizationPolicyProtocolService/RefreshServiceAcl",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RefreshAuthorizationPolicyProtocolServiceServer).RefreshServiceAcl(ctx, req.(*RefreshServiceAclRequestProto))
	}
	return interceptor(ctx, in, info, handler)
}

var _RefreshAuthorizationPolicyProtocolService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "hproto.common.RefreshAuthorizationPolicyProtocolService",
	HandlerType: (*RefreshAuthorizationPolicyProtocolServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "refreshServiceAcl",
			Handler:    _RefreshAuthorizationPolicyProtocolService_RefreshServiceAcl_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor6,
}

func init() {
	proto.RegisterFile("gopkg.in/freddierice/go-hproto.v1/common/RefreshAuthorizationPolicyProtocol.proto", fileDescriptor6)
}

var fileDescriptor6 = []byte{
	// 219 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x90, 0x41, 0x4e, 0xc4, 0x30,
	0x0c, 0x45, 0x95, 0x6d, 0x24, 0x16, 0x74, 0x89, 0x04, 0x8c, 0x66, 0x05, 0x12, 0xe3, 0x08, 0x6e,
	0x30, 0x73, 0x82, 0xa1, 0x9c, 0x20, 0x4a, 0x4d, 0x12, 0xd1, 0xc6, 0xc1, 0x49, 0x2b, 0x95, 0x13,
	0x70, 0x09, 0x24, 0x8e, 0x4a, 0x49, 0xbb, 0x00, 0x84, 0xa6, 0x2b, 0x6f, 0xde, 0xfb, 0xfe, 0xb6,
	0x7c, 0xb4, 0x14, 0x5f, 0x2c, 0xf8, 0xa0, 0x9e, 0x19, 0x9b, 0xc6, 0x23, 0x7b, 0x83, 0xca, 0xd2,
	0xce, 0x45, 0xa6, 0x4c, 0x30, 0xdc, 0x2b, 0x43, 0x5d, 0x47, 0x41, 0xd5, 0x38, 0x11, 0xc9, 0xed,
	0xfb, 0xec, 0x88, 0xfd, 0x9b, 0xce, 0x9e, 0xc2, 0x91, 0x5a, 0x6f, 0xc6, 0xe3, 0x37, 0x69, 0xa8,
	0x85, 0xa2, 0x54, 0x67, 0x8b, 0x3a, 0x7b, 0xdb, 0x6b, 0x79, 0xb9, 0xa8, 0x4f, 0xc8, 0xc3, 0x94,
	0xbe, 0x37, 0x6d, 0x8d, 0xaf, 0x3d, 0xa6, 0x5c, 0xc4, 0xed, 0x46, 0x5e, 0xfd, 0x03, 0xa4, 0x48,
	0x21, 0x61, 0x21, 0x1e, 0x3e, 0x84, 0xbc, 0x5d, 0x5f, 0xbf, 0xd8, 0x55, 0x94, 0xe7, 0xfc, 0x37,
	0xaf, 0xba, 0x83, 0x5f, 0xad, 0xe0, 0x64, 0xa5, 0x8b, 0xdd, 0x3a, 0xfd, 0xa3, 0xdf, 0xa1, 0x96,
	0x1b, 0x62, 0x0b, 0x3a, 0x6a, 0xe3, 0x10, 0x9c, 0x6e, 0x88, 0x22, 0x24, 0x34, 0x3d, 0xfb, 0x3c,
	0xce, 0x5f, 0x39, 0xdc, 0xac, 0x1f, 0x50, 0x66, 0x7a, 0x17, 0xe2, 0x53, 0x88, 0xaf, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x59, 0xd4, 0x90, 0x9a, 0x99, 0x01, 0x00, 0x00,
}
