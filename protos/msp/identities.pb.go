// Code generated by protoc-gen-go. DO NOT EDIT.
// source: msp/identities.proto

/*
Package msp is a generated protocol buffer package.

It is generated from these files:
	msp/identities.proto
	msp/msp_config.proto
	msp/msp_principal.proto

It has these top-level messages:
	SerializedIdentity
	SerializedIdemixIdentity
	MSPConfig
	FabricMSPConfig
	FabricCryptoConfig
	IdemixMSPConfig
	IdemixMSPSignerConfig
	SigningIdentityInfo
	KeyInfo
	FabricOUIdentifier
	FabricNodeOUs
	MSPPrincipal
	OrganizationUnit
	MSPRole
	MSPIdentityAnonymity
*/
package msp

import proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// This struct represents an Identity
// (with its MSP identifier) to be used
// to serialize it and deserialize it
type SerializedIdentity struct {
	// The identifier of the associated membership service provider
	Mspid string `protobuf:"bytes,1,opt,name=mspid" json:"mspid,omitempty"`
	// the Identity, serialized according to the rules of its MPS
	IdBytes []byte `protobuf:"bytes,2,opt,name=id_bytes,json=idBytes,proto3" json:"id_bytes,omitempty"`
}

func (m *SerializedIdentity) Reset()                    { *m = SerializedIdentity{} }
func (m *SerializedIdentity) String() string            { return proto.CompactTextString(m) }
func (*SerializedIdentity) ProtoMessage()               {}
func (*SerializedIdentity) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *SerializedIdentity) GetMspid() string {
	if m != nil {
		return m.Mspid
	}
	return ""
}

func (m *SerializedIdentity) GetIdBytes() []byte {
	if m != nil {
		return m.IdBytes
	}
	return nil
}

// This struct represents an Idemix Identity
// to be used to serialize it and deserialize it.
// The IdemixMSP will first serialize an idemix identity to bytes using
// this proto, and then uses these bytes as id_bytes in SerializedIdentity
type SerializedIdemixIdentity struct {
	// NymX is the X-component of the pseudonym elliptic curve point.
	// It is a []byte representation of an amcl.BIG
	// The pseudonym can be seen as a public key of the identity, it is used to verify signatures.
	NymX []byte `protobuf:"bytes,1,opt,name=NymX,proto3" json:"NymX,omitempty"`
	// NymY is the Y-component of the pseudonym elliptic curve point.
	// It is a []byte representation of an amcl.BIG
	// The pseudonym can be seen as a public key of the identity, it is used to verify signatures.
	NymY []byte `protobuf:"bytes,2,opt,name=NymY,proto3" json:"NymY,omitempty"`
	// OU contains the organizational unit of the idemix identity
	OU []byte `protobuf:"bytes,3,opt,name=OU,proto3" json:"OU,omitempty"`
	// Role contains the role of this identity (e.g., ADMIN or MEMBER)
	Role []byte `protobuf:"bytes,4,opt,name=Role,proto3" json:"Role,omitempty"`
	// Proof contains the cryptographic evidence that this identity is valid
	Proof []byte `protobuf:"bytes,5,opt,name=Proof,proto3" json:"Proof,omitempty"`
}

func (m *SerializedIdemixIdentity) Reset()                    { *m = SerializedIdemixIdentity{} }
func (m *SerializedIdemixIdentity) String() string            { return proto.CompactTextString(m) }
func (*SerializedIdemixIdentity) ProtoMessage()               {}
func (*SerializedIdemixIdentity) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *SerializedIdemixIdentity) GetNymX() []byte {
	if m != nil {
		return m.NymX
	}
	return nil
}

func (m *SerializedIdemixIdentity) GetNymY() []byte {
	if m != nil {
		return m.NymY
	}
	return nil
}

func (m *SerializedIdemixIdentity) GetOU() []byte {
	if m != nil {
		return m.OU
	}
	return nil
}

func (m *SerializedIdemixIdentity) GetRole() []byte {
	if m != nil {
		return m.Role
	}
	return nil
}

func (m *SerializedIdemixIdentity) GetProof() []byte {
	if m != nil {
		return m.Proof
	}
	return nil
}

func init() {
	proto.RegisterType((*SerializedIdentity)(nil), "msp.SerializedIdentity")
	proto.RegisterType((*SerializedIdemixIdentity)(nil), "msp.SerializedIdemixIdentity")
}

func init() { proto.RegisterFile("msp/identities.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 236 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x8f, 0x4f, 0x4b, 0x03, 0x31,
	0x10, 0xc5, 0xd9, 0x6d, 0xeb, 0x9f, 0x50, 0x3c, 0x84, 0x1e, 0xe2, 0xad, 0xf6, 0xb4, 0xa7, 0xe4,
	0xe0, 0x37, 0x28, 0x78, 0xf0, 0xa0, 0x95, 0x95, 0x82, 0x7a, 0x91, 0x6e, 0x33, 0xdd, 0x0e, 0xec,
	0x98, 0x90, 0x89, 0xe0, 0x8a, 0x1f, 0x5e, 0x36, 0x59, 0xc4, 0xde, 0xde, 0x7b, 0xfc, 0xf8, 0x31,
	0x23, 0x16, 0xc4, 0xde, 0xa0, 0x85, 0x8f, 0x88, 0x11, 0x81, 0xb5, 0x0f, 0x2e, 0x3a, 0x39, 0x21,
	0xf6, 0xab, 0x3b, 0x21, 0x9f, 0x21, 0xe0, 0xae, 0xc3, 0x6f, 0xb0, 0xf7, 0x19, 0xe9, 0xe5, 0x42,
	0xcc, 0x88, 0x3d, 0x5a, 0x55, 0x2c, 0x8b, 0xea, 0xb2, 0xce, 0x45, 0x5e, 0x8b, 0x0b, 0xb4, 0xef,
	0x4d, 0x1f, 0x81, 0x55, 0xb9, 0x2c, 0xaa, 0x79, 0x7d, 0x8e, 0x76, 0x3d, 0xd4, 0xd5, 0x8f, 0x50,
	0x27, 0x1a, 0xc2, 0xaf, 0x3f, 0x99, 0x14, 0xd3, 0xc7, 0x9e, 0x5e, 0x92, 0x6b, 0x5e, 0xa7, 0x3c,
	0x6e, 0xaf, 0xa3, 0x26, 0x65, 0x79, 0x25, 0xca, 0xcd, 0x56, 0x4d, 0xd2, 0x52, 0x6e, 0xb6, 0x03,
	0x53, 0xbb, 0x0e, 0xd4, 0x34, 0x33, 0x43, 0x1e, 0x0e, 0x7b, 0x0a, 0xce, 0x1d, 0xd4, 0x2c, 0x8d,
	0xb9, 0xac, 0x1f, 0xc4, 0x8d, 0x0b, 0xad, 0x3e, 0xf6, 0x1e, 0x42, 0x07, 0xb6, 0x85, 0xa0, 0x0f,
	0xbb, 0x26, 0xe0, 0x3e, 0x7f, 0xca, 0x9a, 0xd8, 0xbf, 0x55, 0x2d, 0xc6, 0xe3, 0x67, 0xa3, 0xf7,
	0x8e, 0xcc, 0x3f, 0xd2, 0x64, 0xd2, 0x64, 0xd2, 0x10, 0xfb, 0xe6, 0x2c, 0xe5, 0xdb, 0xdf, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x97, 0xeb, 0x66, 0x12, 0x37, 0x01, 0x00, 0x00,
}
