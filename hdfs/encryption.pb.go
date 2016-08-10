// Code generated by protoc-gen-go.
// source: gopkg.in/freddierice/go-hproto.v1/hdfs/encryption.proto
// DO NOT EDIT!

package hproto_hdfs

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type CreateEncryptionZoneRequestProto struct {
	Src              *string `protobuf:"bytes,1,req,name=src" json:"src,omitempty"`
	KeyName          *string `protobuf:"bytes,2,opt,name=keyName" json:"keyName,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CreateEncryptionZoneRequestProto) Reset()         { *m = CreateEncryptionZoneRequestProto{} }
func (m *CreateEncryptionZoneRequestProto) String() string { return proto.CompactTextString(m) }
func (*CreateEncryptionZoneRequestProto) ProtoMessage()    {}
func (*CreateEncryptionZoneRequestProto) Descriptor() ([]byte, []int) {
	return fileDescriptor11, []int{0}
}

func (m *CreateEncryptionZoneRequestProto) GetSrc() string {
	if m != nil && m.Src != nil {
		return *m.Src
	}
	return ""
}

func (m *CreateEncryptionZoneRequestProto) GetKeyName() string {
	if m != nil && m.KeyName != nil {
		return *m.KeyName
	}
	return ""
}

type CreateEncryptionZoneResponseProto struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *CreateEncryptionZoneResponseProto) Reset()         { *m = CreateEncryptionZoneResponseProto{} }
func (m *CreateEncryptionZoneResponseProto) String() string { return proto.CompactTextString(m) }
func (*CreateEncryptionZoneResponseProto) ProtoMessage()    {}
func (*CreateEncryptionZoneResponseProto) Descriptor() ([]byte, []int) {
	return fileDescriptor11, []int{1}
}

type ListEncryptionZonesRequestProto struct {
	Id               *int64 `protobuf:"varint,1,req,name=id" json:"id,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *ListEncryptionZonesRequestProto) Reset()         { *m = ListEncryptionZonesRequestProto{} }
func (m *ListEncryptionZonesRequestProto) String() string { return proto.CompactTextString(m) }
func (*ListEncryptionZonesRequestProto) ProtoMessage()    {}
func (*ListEncryptionZonesRequestProto) Descriptor() ([]byte, []int) {
	return fileDescriptor11, []int{2}
}

func (m *ListEncryptionZonesRequestProto) GetId() int64 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

type EncryptionZoneProto struct {
	Id                    *int64                      `protobuf:"varint,1,req,name=id" json:"id,omitempty"`
	Path                  *string                     `protobuf:"bytes,2,req,name=path" json:"path,omitempty"`
	Suite                 *CipherSuiteProto           `protobuf:"varint,3,req,name=suite,enum=hproto.hdfs.CipherSuiteProto" json:"suite,omitempty"`
	CryptoProtocolVersion *CryptoProtocolVersionProto `protobuf:"varint,4,req,name=cryptoProtocolVersion,enum=hproto.hdfs.CryptoProtocolVersionProto" json:"cryptoProtocolVersion,omitempty"`
	KeyName               *string                     `protobuf:"bytes,5,req,name=keyName" json:"keyName,omitempty"`
	XXX_unrecognized      []byte                      `json:"-"`
}

func (m *EncryptionZoneProto) Reset()                    { *m = EncryptionZoneProto{} }
func (m *EncryptionZoneProto) String() string            { return proto.CompactTextString(m) }
func (*EncryptionZoneProto) ProtoMessage()               {}
func (*EncryptionZoneProto) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{3} }

func (m *EncryptionZoneProto) GetId() int64 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *EncryptionZoneProto) GetPath() string {
	if m != nil && m.Path != nil {
		return *m.Path
	}
	return ""
}

func (m *EncryptionZoneProto) GetSuite() CipherSuiteProto {
	if m != nil && m.Suite != nil {
		return *m.Suite
	}
	return CipherSuiteProto_UNKNOWN
}

func (m *EncryptionZoneProto) GetCryptoProtocolVersion() CryptoProtocolVersionProto {
	if m != nil && m.CryptoProtocolVersion != nil {
		return *m.CryptoProtocolVersion
	}
	return CryptoProtocolVersionProto_UNKNOWN_PROTOCOL_VERSION
}

func (m *EncryptionZoneProto) GetKeyName() string {
	if m != nil && m.KeyName != nil {
		return *m.KeyName
	}
	return ""
}

type ListEncryptionZonesResponseProto struct {
	Zones            []*EncryptionZoneProto `protobuf:"bytes,1,rep,name=zones" json:"zones,omitempty"`
	HasMore          *bool                  `protobuf:"varint,2,req,name=hasMore" json:"hasMore,omitempty"`
	XXX_unrecognized []byte                 `json:"-"`
}

func (m *ListEncryptionZonesResponseProto) Reset()         { *m = ListEncryptionZonesResponseProto{} }
func (m *ListEncryptionZonesResponseProto) String() string { return proto.CompactTextString(m) }
func (*ListEncryptionZonesResponseProto) ProtoMessage()    {}
func (*ListEncryptionZonesResponseProto) Descriptor() ([]byte, []int) {
	return fileDescriptor11, []int{4}
}

func (m *ListEncryptionZonesResponseProto) GetZones() []*EncryptionZoneProto {
	if m != nil {
		return m.Zones
	}
	return nil
}

func (m *ListEncryptionZonesResponseProto) GetHasMore() bool {
	if m != nil && m.HasMore != nil {
		return *m.HasMore
	}
	return false
}

type GetEZForPathRequestProto struct {
	Src              *string `protobuf:"bytes,1,req,name=src" json:"src,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *GetEZForPathRequestProto) Reset()                    { *m = GetEZForPathRequestProto{} }
func (m *GetEZForPathRequestProto) String() string            { return proto.CompactTextString(m) }
func (*GetEZForPathRequestProto) ProtoMessage()               {}
func (*GetEZForPathRequestProto) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{5} }

func (m *GetEZForPathRequestProto) GetSrc() string {
	if m != nil && m.Src != nil {
		return *m.Src
	}
	return ""
}

type GetEZForPathResponseProto struct {
	Zone             *EncryptionZoneProto `protobuf:"bytes,1,opt,name=zone" json:"zone,omitempty"`
	XXX_unrecognized []byte               `json:"-"`
}

func (m *GetEZForPathResponseProto) Reset()                    { *m = GetEZForPathResponseProto{} }
func (m *GetEZForPathResponseProto) String() string            { return proto.CompactTextString(m) }
func (*GetEZForPathResponseProto) ProtoMessage()               {}
func (*GetEZForPathResponseProto) Descriptor() ([]byte, []int) { return fileDescriptor11, []int{6} }

func (m *GetEZForPathResponseProto) GetZone() *EncryptionZoneProto {
	if m != nil {
		return m.Zone
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateEncryptionZoneRequestProto)(nil), "hproto.hdfs.CreateEncryptionZoneRequestProto")
	proto.RegisterType((*CreateEncryptionZoneResponseProto)(nil), "hproto.hdfs.CreateEncryptionZoneResponseProto")
	proto.RegisterType((*ListEncryptionZonesRequestProto)(nil), "hproto.hdfs.ListEncryptionZonesRequestProto")
	proto.RegisterType((*EncryptionZoneProto)(nil), "hproto.hdfs.EncryptionZoneProto")
	proto.RegisterType((*ListEncryptionZonesResponseProto)(nil), "hproto.hdfs.ListEncryptionZonesResponseProto")
	proto.RegisterType((*GetEZForPathRequestProto)(nil), "hproto.hdfs.GetEZForPathRequestProto")
	proto.RegisterType((*GetEZForPathResponseProto)(nil), "hproto.hdfs.GetEZForPathResponseProto")
}

func init() {
	proto.RegisterFile("gopkg.in/freddierice/go-hproto.v1/hdfs/encryption.proto", fileDescriptor11)
}

var fileDescriptor11 = []byte{
	// 389 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x91, 0x51, 0x6b, 0xdb, 0x30,
	0x14, 0x85, 0x71, 0x9c, 0xb0, 0xed, 0x06, 0xc2, 0xf0, 0x08, 0x78, 0x83, 0x31, 0xcf, 0x63, 0x2c,
	0x0f, 0x9b, 0x4d, 0xb2, 0xd2, 0x3e, 0x37, 0x21, 0xed, 0x4b, 0x1b, 0x52, 0x17, 0xfa, 0x10, 0xe8,
	0x83, 0xb1, 0x6f, 0x6c, 0x91, 0xd6, 0x52, 0x25, 0xa5, 0x90, 0xfe, 0x9a, 0xfe, 0xbb, 0xfe, 0x8d,
	0x4a, 0x72, 0x12, 0xec, 0x62, 0x4a, 0x5e, 0xc2, 0xd5, 0xd5, 0xd1, 0x39, 0x27, 0x9f, 0xe1, 0x24,
	0xa3, 0x6c, 0x95, 0x05, 0xa4, 0x08, 0x97, 0x1c, 0xd3, 0x94, 0x20, 0x27, 0x09, 0x86, 0x19, 0xfd,
	0x97, 0x33, 0x4e, 0x25, 0x0d, 0x1e, 0x87, 0x61, 0x9e, 0x2e, 0x45, 0x88, 0x45, 0xc2, 0x37, 0x4c,
	0x12, 0x5a, 0x04, 0xe6, 0xc2, 0xe9, 0x6e, 0x05, 0xfa, 0xf6, 0xdb, 0xf0, 0x40, 0x17, 0xfd, 0x53,
	0xbe, 0xf7, 0x67, 0xe0, 0x4d, 0x38, 0xc6, 0x12, 0xa7, 0x7b, 0xe7, 0x05, 0x2d, 0x30, 0xc2, 0x87,
	0x35, 0x0a, 0x39, 0x37, 0x19, 0x9f, 0xc1, 0x16, 0x3c, 0x71, 0x2d, 0xaf, 0x35, 0xf8, 0x14, 0xe9,
	0xd1, 0x71, 0xe1, 0xc3, 0x0a, 0x37, 0xb3, 0xf8, 0x1e, 0xdd, 0x96, 0x67, 0xa9, 0xed, 0xee, 0xe8,
	0xff, 0x82, 0x9f, 0xcd, 0x7e, 0x82, 0xd1, 0x42, 0xa0, 0x31, 0xf4, 0x87, 0xf0, 0xe3, 0x82, 0x08,
	0x59, 0x97, 0x88, 0x5a, 0x66, 0x0f, 0x5a, 0x24, 0x35, 0x91, 0x76, 0xa4, 0x26, 0xff, 0xc5, 0x82,
	0x2f, 0x75, 0x7d, 0xa3, 0xce, 0x71, 0xa0, 0xcd, 0x62, 0x99, 0xab, 0x5a, 0xba, 0xac, 0x99, 0x9d,
	0xff, 0xd0, 0x11, 0x6b, 0x22, 0xd1, 0xb5, 0xd5, 0xb2, 0x37, 0xfa, 0x1e, 0x54, 0x98, 0x05, 0x13,
	0xc2, 0x72, 0xe4, 0xd7, 0xfa, 0xde, 0x38, 0x46, 0xa5, 0xd6, 0xb9, 0x85, 0xbe, 0x49, 0xa3, 0x66,
	0x9b, 0xd0, 0xbb, 0x1b, 0xe4, 0x42, 0x45, 0xbb, 0x6d, 0x63, 0xf2, 0xa7, 0x6e, 0xd2, 0xa4, 0x2c,
	0xed, 0x9a, 0x5d, 0xaa, 0x04, 0x3b, 0xa6, 0xea, 0x9e, 0xa0, 0x04, 0xaf, 0x11, 0x4e, 0x05, 0xa0,
	0x73, 0x0c, 0x9d, 0x27, 0xbd, 0x55, 0x7f, 0xdc, 0x1e, 0x74, 0x47, 0x5e, 0xad, 0x4c, 0x03, 0xa6,
	0xa8, 0x94, 0xeb, 0xd4, 0x3c, 0x16, 0x97, 0x94, 0xa3, 0x01, 0xf4, 0x31, 0xda, 0x1d, 0xfd, 0xbf,
	0xe0, 0x9e, 0xa3, 0x9c, 0x2e, 0xce, 0x28, 0x9f, 0x2b, 0x66, 0xef, 0x7f, 0x7f, 0xff, 0x0a, 0xbe,
	0xd6, 0xd5, 0xd5, 0x72, 0x47, 0xd0, 0xd6, 0x69, 0x4a, 0x6f, 0x1d, 0xd4, 0xcd, 0xa8, 0xc7, 0xa7,
	0xf0, 0x9b, 0xf2, 0x2c, 0x88, 0x59, 0x9c, 0xe4, 0x58, 0x7b, 0xc3, 0xb6, 0xe8, 0xca, 0x61, 0xdc,
	0x7f, 0x43, 0xc6, 0x98, 0x88, 0x67, 0xcb, 0x7a, 0x0d, 0x00, 0x00, 0xff, 0xff, 0xea, 0x94, 0x33,
	0x1a, 0x45, 0x03, 0x00, 0x00,
}