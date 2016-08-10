// Code generated by protoc-gen-go.
// source: gopkg.in/freddierice/go-hproto.v1/hdfs/InterDatanodeProtocol.proto
// DO NOT EDIT!

package hproto_hdfs

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
// Block with location information and new generation stamp
// to be used for recovery.
type InitReplicaRecoveryRequestProto struct {
	Block            *RecoveringBlockProto `protobuf:"bytes,1,req,name=block" json:"block,omitempty"`
	XXX_unrecognized []byte                `json:"-"`
}

func (m *InitReplicaRecoveryRequestProto) Reset()                    { *m = InitReplicaRecoveryRequestProto{} }
func (m *InitReplicaRecoveryRequestProto) String() string            { return proto.CompactTextString(m) }
func (*InitReplicaRecoveryRequestProto) ProtoMessage()               {}
func (*InitReplicaRecoveryRequestProto) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func (m *InitReplicaRecoveryRequestProto) GetBlock() *RecoveringBlockProto {
	if m != nil {
		return m.Block
	}
	return nil
}

// *
// Repica recovery information
type InitReplicaRecoveryResponseProto struct {
	ReplicaFound *bool `protobuf:"varint,1,req,name=replicaFound" json:"replicaFound,omitempty"`
	// The following entries are not set if there was no replica found.
	State            *ReplicaStateProto `protobuf:"varint,2,opt,name=state,enum=hproto.hdfs.ReplicaStateProto" json:"state,omitempty"`
	Block            *BlockProto        `protobuf:"bytes,3,opt,name=block" json:"block,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *InitReplicaRecoveryResponseProto) Reset()         { *m = InitReplicaRecoveryResponseProto{} }
func (m *InitReplicaRecoveryResponseProto) String() string { return proto.CompactTextString(m) }
func (*InitReplicaRecoveryResponseProto) ProtoMessage()    {}
func (*InitReplicaRecoveryResponseProto) Descriptor() ([]byte, []int) {
	return fileDescriptor3, []int{1}
}

func (m *InitReplicaRecoveryResponseProto) GetReplicaFound() bool {
	if m != nil && m.ReplicaFound != nil {
		return *m.ReplicaFound
	}
	return false
}

func (m *InitReplicaRecoveryResponseProto) GetState() ReplicaStateProto {
	if m != nil && m.State != nil {
		return *m.State
	}
	return ReplicaStateProto_FINALIZED
}

func (m *InitReplicaRecoveryResponseProto) GetBlock() *BlockProto {
	if m != nil {
		return m.Block
	}
	return nil
}

// *
// Update replica with new generation stamp and length
type UpdateReplicaUnderRecoveryRequestProto struct {
	Block      *ExtendedBlockProto `protobuf:"bytes,1,req,name=block" json:"block,omitempty"`
	RecoveryId *uint64             `protobuf:"varint,2,req,name=recoveryId" json:"recoveryId,omitempty"`
	NewLength  *uint64             `protobuf:"varint,3,req,name=newLength" json:"newLength,omitempty"`
	// New blockId for copy (truncate), default is 0.
	NewBlockId       *uint64 `protobuf:"varint,4,opt,name=newBlockId,def=0" json:"newBlockId,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *UpdateReplicaUnderRecoveryRequestProto) Reset() {
	*m = UpdateReplicaUnderRecoveryRequestProto{}
}
func (m *UpdateReplicaUnderRecoveryRequestProto) String() string { return proto.CompactTextString(m) }
func (*UpdateReplicaUnderRecoveryRequestProto) ProtoMessage()    {}
func (*UpdateReplicaUnderRecoveryRequestProto) Descriptor() ([]byte, []int) {
	return fileDescriptor3, []int{2}
}

const Default_UpdateReplicaUnderRecoveryRequestProto_NewBlockId uint64 = 0

func (m *UpdateReplicaUnderRecoveryRequestProto) GetBlock() *ExtendedBlockProto {
	if m != nil {
		return m.Block
	}
	return nil
}

func (m *UpdateReplicaUnderRecoveryRequestProto) GetRecoveryId() uint64 {
	if m != nil && m.RecoveryId != nil {
		return *m.RecoveryId
	}
	return 0
}

func (m *UpdateReplicaUnderRecoveryRequestProto) GetNewLength() uint64 {
	if m != nil && m.NewLength != nil {
		return *m.NewLength
	}
	return 0
}

func (m *UpdateReplicaUnderRecoveryRequestProto) GetNewBlockId() uint64 {
	if m != nil && m.NewBlockId != nil {
		return *m.NewBlockId
	}
	return Default_UpdateReplicaUnderRecoveryRequestProto_NewBlockId
}

// *
// Response returns updated block information
type UpdateReplicaUnderRecoveryResponseProto struct {
	StorageUuid      *string `protobuf:"bytes,1,opt,name=storageUuid" json:"storageUuid,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *UpdateReplicaUnderRecoveryResponseProto) Reset() {
	*m = UpdateReplicaUnderRecoveryResponseProto{}
}
func (m *UpdateReplicaUnderRecoveryResponseProto) String() string { return proto.CompactTextString(m) }
func (*UpdateReplicaUnderRecoveryResponseProto) ProtoMessage()    {}
func (*UpdateReplicaUnderRecoveryResponseProto) Descriptor() ([]byte, []int) {
	return fileDescriptor3, []int{3}
}

func (m *UpdateReplicaUnderRecoveryResponseProto) GetStorageUuid() string {
	if m != nil && m.StorageUuid != nil {
		return *m.StorageUuid
	}
	return ""
}

func init() {
	proto.RegisterType((*InitReplicaRecoveryRequestProto)(nil), "hproto.hdfs.InitReplicaRecoveryRequestProto")
	proto.RegisterType((*InitReplicaRecoveryResponseProto)(nil), "hproto.hdfs.InitReplicaRecoveryResponseProto")
	proto.RegisterType((*UpdateReplicaUnderRecoveryRequestProto)(nil), "hproto.hdfs.UpdateReplicaUnderRecoveryRequestProto")
	proto.RegisterType((*UpdateReplicaUnderRecoveryResponseProto)(nil), "hproto.hdfs.UpdateReplicaUnderRecoveryResponseProto")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for InterDatanodeProtocolService service

type InterDatanodeProtocolServiceClient interface {
	// *
	// Initialize recovery of a replica
	InitReplicaRecovery(ctx context.Context, in *InitReplicaRecoveryRequestProto, opts ...grpc.CallOption) (*InitReplicaRecoveryResponseProto, error)
	// *
	// Update a replica with new generation stamp and length
	UpdateReplicaUnderRecovery(ctx context.Context, in *UpdateReplicaUnderRecoveryRequestProto, opts ...grpc.CallOption) (*UpdateReplicaUnderRecoveryResponseProto, error)
}

type interDatanodeProtocolServiceClient struct {
	cc *grpc.ClientConn
}

func NewInterDatanodeProtocolServiceClient(cc *grpc.ClientConn) InterDatanodeProtocolServiceClient {
	return &interDatanodeProtocolServiceClient{cc}
}

func (c *interDatanodeProtocolServiceClient) InitReplicaRecovery(ctx context.Context, in *InitReplicaRecoveryRequestProto, opts ...grpc.CallOption) (*InitReplicaRecoveryResponseProto, error) {
	out := new(InitReplicaRecoveryResponseProto)
	err := grpc.Invoke(ctx, "/hproto.hdfs.InterDatanodeProtocolService/initReplicaRecovery", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *interDatanodeProtocolServiceClient) UpdateReplicaUnderRecovery(ctx context.Context, in *UpdateReplicaUnderRecoveryRequestProto, opts ...grpc.CallOption) (*UpdateReplicaUnderRecoveryResponseProto, error) {
	out := new(UpdateReplicaUnderRecoveryResponseProto)
	err := grpc.Invoke(ctx, "/hproto.hdfs.InterDatanodeProtocolService/updateReplicaUnderRecovery", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for InterDatanodeProtocolService service

type InterDatanodeProtocolServiceServer interface {
	// *
	// Initialize recovery of a replica
	InitReplicaRecovery(context.Context, *InitReplicaRecoveryRequestProto) (*InitReplicaRecoveryResponseProto, error)
	// *
	// Update a replica with new generation stamp and length
	UpdateReplicaUnderRecovery(context.Context, *UpdateReplicaUnderRecoveryRequestProto) (*UpdateReplicaUnderRecoveryResponseProto, error)
}

func RegisterInterDatanodeProtocolServiceServer(s *grpc.Server, srv InterDatanodeProtocolServiceServer) {
	s.RegisterService(&_InterDatanodeProtocolService_serviceDesc, srv)
}

func _InterDatanodeProtocolService_InitReplicaRecovery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InitReplicaRecoveryRequestProto)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InterDatanodeProtocolServiceServer).InitReplicaRecovery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hproto.hdfs.InterDatanodeProtocolService/InitReplicaRecovery",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InterDatanodeProtocolServiceServer).InitReplicaRecovery(ctx, req.(*InitReplicaRecoveryRequestProto))
	}
	return interceptor(ctx, in, info, handler)
}

func _InterDatanodeProtocolService_UpdateReplicaUnderRecovery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateReplicaUnderRecoveryRequestProto)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InterDatanodeProtocolServiceServer).UpdateReplicaUnderRecovery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hproto.hdfs.InterDatanodeProtocolService/UpdateReplicaUnderRecovery",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InterDatanodeProtocolServiceServer).UpdateReplicaUnderRecovery(ctx, req.(*UpdateReplicaUnderRecoveryRequestProto))
	}
	return interceptor(ctx, in, info, handler)
}

var _InterDatanodeProtocolService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "hproto.hdfs.InterDatanodeProtocolService",
	HandlerType: (*InterDatanodeProtocolServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "initReplicaRecovery",
			Handler:    _InterDatanodeProtocolService_InitReplicaRecovery_Handler,
		},
		{
			MethodName: "updateReplicaUnderRecovery",
			Handler:    _InterDatanodeProtocolService_UpdateReplicaUnderRecovery_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor3,
}

func init() {
	proto.RegisterFile("gopkg.in/freddierice/go-hproto.v1/hdfs/InterDatanodeProtocol.proto", fileDescriptor3)
}

var fileDescriptor3 = []byte{
	// 449 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x53, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0x96, 0x4b, 0x27, 0xb1, 0x57, 0xc4, 0xc1, 0x1c, 0x98, 0xc2, 0xb4, 0x75, 0x95, 0x80, 0x1e,
	0x68, 0xca, 0xca, 0xd0, 0x24, 0x8e, 0x15, 0x20, 0x0a, 0x1c, 0x90, 0xa7, 0x5e, 0xb8, 0x85, 0xf8,
	0x2d, 0xb1, 0x56, 0xd9, 0xc1, 0x71, 0x06, 0xfc, 0x02, 0xf8, 0x11, 0x1c, 0xb8, 0xf3, 0x27, 0xf8,
	0x69, 0xd8, 0x4e, 0x10, 0x31, 0x0d, 0x5b, 0x2e, 0xae, 0xfa, 0xde, 0xf7, 0x7d, 0xfe, 0xde, 0xe7,
	0x17, 0x58, 0x66, 0xaa, 0xb8, 0xc8, 0x62, 0x21, 0xe7, 0xe7, 0x1a, 0x39, 0x17, 0xa8, 0x45, 0x8a,
	0xf3, 0x4c, 0xcd, 0xf2, 0x42, 0x2b, 0xa3, 0xe2, 0xcb, 0xe3, 0x79, 0xce, 0xcf, 0xcb, 0xf9, 0x4a,
	0x1a, 0xd4, 0xcf, 0x13, 0x93, 0x48, 0xc5, 0xf1, 0x9d, 0x6b, 0xa5, 0x6a, 0x13, 0x7b, 0x0c, 0x1d,
	0x35, 0x58, 0x07, 0x8c, 0x8e, 0x7b, 0x0a, 0xba, 0xa3, 0xe6, 0x47, 0xa7, 0x3d, 0x29, 0xaf, 0xec,
	0x71, 0x86, 0xfa, 0x12, 0x75, 0x4d, 0x9c, 0xbc, 0x87, 0xc3, 0x95, 0x14, 0x86, 0x61, 0xb1, 0x11,
	0x69, 0xc2, 0x30, 0x55, 0xb6, 0xf7, 0x85, 0xe1, 0xc7, 0x0a, 0x4b, 0xe3, 0x4d, 0xd2, 0x53, 0xd8,
	0xf9, 0xb0, 0x51, 0xe9, 0xc5, 0x1e, 0x19, 0x0f, 0xa6, 0xa3, 0xc5, 0x51, 0xdc, 0xf2, 0x1a, 0x37,
	0x0c, 0x21, 0xb3, 0xa5, 0xc3, 0x78, 0x06, 0xab, 0xf1, 0x93, 0x9f, 0x04, 0xc6, 0x9d, 0xe2, 0x65,
	0xa1, 0x64, 0x59, 0x47, 0x40, 0x27, 0x70, 0x4b, 0xd7, 0xfd, 0x97, 0xaa, 0x92, 0xdc, 0x5f, 0x72,
	0x93, 0x05, 0x35, 0x7a, 0x02, 0x3b, 0xa5, 0x49, 0x0c, 0xee, 0x0d, 0xc6, 0x64, 0x7a, 0x7b, 0x71,
	0xf0, 0x8f, 0x03, 0x8f, 0x3c, 0x73, 0x80, 0xe6, 0x7a, 0x0f, 0xa6, 0xb3, 0x3f, 0xbe, 0x6f, 0x58,
	0xd6, 0x68, 0x71, 0x37, 0x60, 0x6d, 0xbb, 0xfd, 0x45, 0xe0, 0xc1, 0xba, 0xe0, 0x96, 0xd9, 0x28,
	0xae, 0x25, 0x47, 0xdd, 0x99, 0xc8, 0xd3, 0x30, 0x91, 0xc3, 0x40, 0xf9, 0xc5, 0x67, 0x83, 0x96,
	0xc9, 0xb7, 0x6e, 0xa0, 0x07, 0x00, 0xba, 0x91, 0x5b, 0x71, 0x3b, 0xcb, 0x60, 0x3a, 0x64, 0xad,
	0x0a, 0xdd, 0x87, 0x5d, 0x89, 0x9f, 0xde, 0xa2, 0xcc, 0x4c, 0x6e, 0x4d, 0xbb, 0xf6, 0xdf, 0x02,
	0x3d, 0x02, 0xb0, 0x7f, 0xbc, 0xaa, 0x65, 0x0f, 0xed, 0x4c, 0xc3, 0x67, 0xe4, 0x31, 0x6b, 0x15,
	0x27, 0x6f, 0xe0, 0xe1, 0x55, 0x13, 0xb4, 0x63, 0x1f, 0xc3, 0xa8, 0x34, 0x4a, 0x27, 0x19, 0xae,
	0x2b, 0xe1, 0x52, 0x27, 0xd3, 0x5d, 0xd6, 0x2e, 0x2d, 0xbe, 0x0f, 0x60, 0xbf, 0x73, 0x65, 0xdd,
	0xfe, 0xd8, 0xed, 0xa2, 0x1a, 0xee, 0x88, 0xed, 0xd7, 0xa5, 0x8f, 0x82, 0x34, 0xae, 0x59, 0xae,
	0x68, 0x76, 0x3d, 0xba, 0x6d, 0xfb, 0x2b, 0x81, 0xa8, 0xfa, 0xef, 0x88, 0xf4, 0x49, 0xa0, 0xd6,
	0xef, 0x35, 0xa3, 0x93, 0xde, 0xa4, 0x96, 0x93, 0xe5, 0x6b, 0xb8, 0xaf, 0x74, 0x16, 0x27, 0x45,
	0x92, 0xe6, 0x18, 0x28, 0x14, 0xc1, 0xa7, 0xbd, 0xbc, 0xd7, 0x19, 0xa2, 0xff, 0x2d, 0xbf, 0x11,
	0xf2, 0x83, 0x90, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x19, 0xb7, 0x25, 0x41, 0x42, 0x04, 0x00,
	0x00,
}
