// Code generated by protoc-gen-go.
// source: common/common.proto
// DO NOT EDIT!

/*
Package common is a generated protocol buffer package.

It is generated from these files:
	common/common.proto
	common/configuration.proto

It has these top-level messages:
	Header
	ChainHeader
	SignatureHeader
	Payload
	Envelope
	Block
	BlockHeader
	BlockData
	BlockMetadata
	ConfigurationEnvelope
	SignedConfigurationItem
	ConfigurationItem
	ConfigurationSignature
	Policy
	SignaturePolicyEnvelope
	SignaturePolicy
*/
package common

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type HeaderType int32

const (
	HeaderType_MESSAGE                   HeaderType = 0
	HeaderType_CONFIGURATION_TRANSACTION HeaderType = 1
	HeaderType_CONFIGURATION_ITEM        HeaderType = 2
	HeaderType_ENDORSER_TRANSACTION      HeaderType = 3
)

var HeaderType_name = map[int32]string{
	0: "MESSAGE",
	1: "CONFIGURATION_TRANSACTION",
	2: "CONFIGURATION_ITEM",
	3: "ENDORSER_TRANSACTION",
}
var HeaderType_value = map[string]int32{
	"MESSAGE":                   0,
	"CONFIGURATION_TRANSACTION": 1,
	"CONFIGURATION_ITEM":        2,
	"ENDORSER_TRANSACTION":      3,
}

func (x HeaderType) String() string {
	return proto.EnumName(HeaderType_name, int32(x))
}
func (HeaderType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Header struct {
	ChainHeader     *ChainHeader     `protobuf:"bytes,1,opt,name=chainHeader" json:"chainHeader,omitempty"`
	SignatureHeader *SignatureHeader `protobuf:"bytes,2,opt,name=signatureHeader" json:"signatureHeader,omitempty"`
}

func (m *Header) Reset()                    { *m = Header{} }
func (m *Header) String() string            { return proto.CompactTextString(m) }
func (*Header) ProtoMessage()               {}
func (*Header) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Header) GetChainHeader() *ChainHeader {
	if m != nil {
		return m.ChainHeader
	}
	return nil
}

func (m *Header) GetSignatureHeader() *SignatureHeader {
	if m != nil {
		return m.SignatureHeader
	}
	return nil
}

// Header is a generic replay prevention and identity message to include in a signed payload
type ChainHeader struct {
	Type int32 `protobuf:"varint,1,opt,name=type" json:"type,omitempty"`
	// Version indicates message protocol version
	Version int32 `protobuf:"varint,2,opt,name=version" json:"version,omitempty"`
	// Timestamp is the local time when the message was created
	// by the sender
	Timestamp *google_protobuf.Timestamp `protobuf:"bytes,3,opt,name=timestamp" json:"timestamp,omitempty"`
	// Identifier of the chain this message is bound for
	ChainID []byte `protobuf:"bytes,4,opt,name=chainID,proto3" json:"chainID,omitempty"`
}

func (m *ChainHeader) Reset()                    { *m = ChainHeader{} }
func (m *ChainHeader) String() string            { return proto.CompactTextString(m) }
func (*ChainHeader) ProtoMessage()               {}
func (*ChainHeader) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ChainHeader) GetTimestamp() *google_protobuf.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

type SignatureHeader struct {
	// Creator of the message, specified as a certificate chain
	Creator []byte `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	// Arbitrary number that may only be used once. Can be used to detect replay attacks.
	Nonce []byte `protobuf:"bytes,2,opt,name=nonce,proto3" json:"nonce,omitempty"`
	// The epoch in which this header was generated, where epoch is defined based on block height
	Epoch uint64 `protobuf:"varint,3,opt,name=epoch" json:"epoch,omitempty"`
}

func (m *SignatureHeader) Reset()                    { *m = SignatureHeader{} }
func (m *SignatureHeader) String() string            { return proto.CompactTextString(m) }
func (*SignatureHeader) ProtoMessage()               {}
func (*SignatureHeader) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

// Payload is the message contents (and header to allow for signing)
type Payload struct {
	// Header is included to provide identity and prevent replay
	Header *Header `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	// Data, the encoding of which is defined by the type in the header
	Data []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *Payload) Reset()                    { *m = Payload{} }
func (m *Payload) String() string            { return proto.CompactTextString(m) }
func (*Payload) ProtoMessage()               {}
func (*Payload) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Payload) GetHeader() *Header {
	if m != nil {
		return m.Header
	}
	return nil
}

// Envelope wraps a Payload with a signature so that the message may be authenticated
type Envelope struct {
	// A marshaled Payload
	Payload []byte `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
	// A signature by the creator specified in the Payload header
	Signature []byte `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (m *Envelope) Reset()                    { *m = Envelope{} }
func (m *Envelope) String() string            { return proto.CompactTextString(m) }
func (*Envelope) ProtoMessage()               {}
func (*Envelope) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

// This is finalized block structure to be shared among the orderer and peer
// Note that the BlockHeader chains to the previous BlockHeader, and the BlockData hash is embedded
// in the BlockHeader.  This makes it natural and obvious that the Data is included in the hash, but
// the Metadata is not.
type Block struct {
	Header   *BlockHeader   `protobuf:"bytes,1,opt,name=Header" json:"Header,omitempty"`
	Data     *BlockData     `protobuf:"bytes,2,opt,name=Data" json:"Data,omitempty"`
	Metadata *BlockMetadata `protobuf:"bytes,3,opt,name=Metadata" json:"Metadata,omitempty"`
}

func (m *Block) Reset()                    { *m = Block{} }
func (m *Block) String() string            { return proto.CompactTextString(m) }
func (*Block) ProtoMessage()               {}
func (*Block) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Block) GetHeader() *BlockHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Block) GetData() *BlockData {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Block) GetMetadata() *BlockMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type BlockHeader struct {
	Number       uint64 `protobuf:"varint,1,opt,name=Number" json:"Number,omitempty"`
	PreviousHash []byte `protobuf:"bytes,2,opt,name=PreviousHash,proto3" json:"PreviousHash,omitempty"`
	DataHash     []byte `protobuf:"bytes,3,opt,name=DataHash,proto3" json:"DataHash,omitempty"`
}

func (m *BlockHeader) Reset()                    { *m = BlockHeader{} }
func (m *BlockHeader) String() string            { return proto.CompactTextString(m) }
func (*BlockHeader) ProtoMessage()               {}
func (*BlockHeader) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

type BlockData struct {
	Data [][]byte `protobuf:"bytes,1,rep,name=Data,proto3" json:"Data,omitempty"`
}

func (m *BlockData) Reset()                    { *m = BlockData{} }
func (m *BlockData) String() string            { return proto.CompactTextString(m) }
func (*BlockData) ProtoMessage()               {}
func (*BlockData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

type BlockMetadata struct {
	Metadata [][]byte `protobuf:"bytes,1,rep,name=Metadata,proto3" json:"Metadata,omitempty"`
}

func (m *BlockMetadata) Reset()                    { *m = BlockMetadata{} }
func (m *BlockMetadata) String() string            { return proto.CompactTextString(m) }
func (*BlockMetadata) ProtoMessage()               {}
func (*BlockMetadata) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func init() {
	proto.RegisterType((*Header)(nil), "common.Header")
	proto.RegisterType((*ChainHeader)(nil), "common.ChainHeader")
	proto.RegisterType((*SignatureHeader)(nil), "common.SignatureHeader")
	proto.RegisterType((*Payload)(nil), "common.Payload")
	proto.RegisterType((*Envelope)(nil), "common.Envelope")
	proto.RegisterType((*Block)(nil), "common.Block")
	proto.RegisterType((*BlockHeader)(nil), "common.BlockHeader")
	proto.RegisterType((*BlockData)(nil), "common.BlockData")
	proto.RegisterType((*BlockMetadata)(nil), "common.BlockMetadata")
	proto.RegisterEnum("common.HeaderType", HeaderType_name, HeaderType_value)
}

func init() { proto.RegisterFile("common/common.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 542 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x5c, 0x54, 0x5d, 0x8f, 0xd2, 0x40,
	0x14, 0xb5, 0xcb, 0xf7, 0x2d, 0xba, 0x38, 0xbb, 0xae, 0x48, 0x34, 0x4b, 0x26, 0xd1, 0x6c, 0x24,
	0x42, 0x5c, 0x63, 0xe2, 0x2b, 0x1f, 0x75, 0x97, 0x07, 0x60, 0x33, 0x60, 0x4c, 0x7c, 0x31, 0x43,
	0x99, 0xa5, 0x55, 0xe8, 0x34, 0xfd, 0x20, 0xe1, 0xd5, 0x1f, 0xa0, 0x7f, 0xd9, 0xcc, 0x57, 0xa1,
	0x3c, 0x31, 0xe7, 0x9e, 0x73, 0x6f, 0xcf, 0x3d, 0x1d, 0x0a, 0x17, 0x2e, 0xdf, 0x6e, 0x79, 0xd0,
	0x53, 0x3f, 0xdd, 0x30, 0xe2, 0x09, 0x47, 0x65, 0x85, 0x5a, 0xd7, 0x6b, 0xce, 0xd7, 0x1b, 0xd6,
	0x93, 0xd5, 0x65, 0xfa, 0xd8, 0x4b, 0xfc, 0x2d, 0x8b, 0x13, 0xba, 0x0d, 0x95, 0x10, 0xff, 0xb1,
	0xa0, 0x7c, 0xcf, 0xe8, 0x8a, 0x45, 0xe8, 0x33, 0xd8, 0xae, 0x47, 0xfd, 0x40, 0xc1, 0xa6, 0xd5,
	0xb6, 0x6e, 0xec, 0xdb, 0x8b, 0xae, 0x9e, 0x3b, 0x3c, 0x50, 0xe4, 0x58, 0x87, 0xfa, 0x70, 0x1e,
	0xfb, 0xeb, 0x80, 0x26, 0x69, 0xc4, 0x74, 0xeb, 0x99, 0x6c, 0x7d, 0x69, 0x5a, 0xe7, 0x79, 0x9a,
	0x9c, 0xea, 0xf1, 0x3f, 0x0b, 0xec, 0xa3, 0xf9, 0x08, 0x41, 0x31, 0xd9, 0x87, 0x4c, 0x5a, 0x28,
	0x11, 0x79, 0x46, 0x4d, 0xa8, 0xec, 0x58, 0x14, 0xfb, 0x3c, 0x90, 0xe3, 0x4b, 0xc4, 0x40, 0xf4,
	0x05, 0x6a, 0xd9, 0x56, 0xcd, 0x82, 0x7c, 0x74, 0xab, 0xab, 0xf6, 0xee, 0x9a, 0xbd, 0xbb, 0x0b,
	0xa3, 0x20, 0x07, 0xb1, 0x98, 0x29, 0x37, 0x19, 0x8f, 0x9a, 0xc5, 0xb6, 0x75, 0x53, 0x27, 0x06,
	0xe2, 0xef, 0x70, 0x7e, 0xe2, 0x5a, 0x8a, 0x23, 0x46, 0x13, 0xae, 0xa2, 0x11, 0x62, 0x05, 0xd1,
	0x25, 0x94, 0x02, 0x1e, 0xb8, 0x4c, 0x1a, 0xab, 0x13, 0x05, 0x44, 0x95, 0x85, 0xdc, 0xf5, 0xa4,
	0xa5, 0x22, 0x51, 0x00, 0x3b, 0x50, 0x79, 0xa0, 0xfb, 0x0d, 0xa7, 0x2b, 0xf4, 0x0e, 0xca, 0xde,
	0x71, 0xd4, 0xcf, 0x4c, 0x5e, 0x3a, 0x26, 0xcd, 0x8a, 0x34, 0x56, 0x34, 0xa1, 0x7a, 0xba, 0x3c,
	0xe3, 0x01, 0x54, 0x9d, 0x60, 0xc7, 0x36, 0x5c, 0x25, 0x13, 0xaa, 0x91, 0xc6, 0x98, 0x86, 0xe8,
	0x35, 0xd4, 0xb2, 0xa8, 0x75, 0xfb, 0xa1, 0x80, 0xff, 0x5a, 0x50, 0x1a, 0x6c, 0xb8, 0xfb, 0x1b,
	0x75, 0xcc, 0x1d, 0x38, 0x7d, 0xe9, 0x92, 0x36, 0x76, 0x74, 0x0e, 0x6f, 0xa1, 0x38, 0x32, 0x76,
	0xec, 0xdb, 0xe7, 0x39, 0xa9, 0x20, 0x88, 0xa4, 0xd1, 0x47, 0xa8, 0x4e, 0x58, 0x42, 0xa5, 0x73,
	0xf5, 0x52, 0x5e, 0xe4, 0xa4, 0x86, 0x24, 0x99, 0x0c, 0x33, 0xb0, 0x8f, 0x1e, 0x88, 0xae, 0xa0,
	0x3c, 0x4d, 0xb7, 0x4b, 0xed, 0xaa, 0x48, 0x34, 0x42, 0x18, 0xea, 0x0f, 0x11, 0xdb, 0xf9, 0x3c,
	0x8d, 0xef, 0x69, 0xec, 0xe9, 0xc5, 0x72, 0x35, 0xd4, 0x82, 0xaa, 0x70, 0x21, 0xf9, 0x82, 0xe4,
	0x33, 0x8c, 0xaf, 0xa1, 0x96, 0x99, 0x15, 0xe1, 0xca, 0x6d, 0xac, 0x76, 0x41, 0x84, 0x2b, 0xce,
	0xb8, 0x03, 0x4f, 0x73, 0x16, 0xc5, 0xb4, 0x6c, 0x17, 0x25, 0xcc, 0xf0, 0xfb, 0x5f, 0x00, 0xca,
	0xef, 0x42, 0xdc, 0x52, 0x1b, 0x2a, 0x13, 0x67, 0x3e, 0xef, 0xdf, 0x39, 0x8d, 0x27, 0xe8, 0x0d,
	0xbc, 0x1a, 0xce, 0xa6, 0x5f, 0xc7, 0x77, 0xdf, 0x48, 0x7f, 0x31, 0x9e, 0x4d, 0x7f, 0x2e, 0x48,
	0x7f, 0x3a, 0xef, 0x0f, 0xc5, 0xb9, 0x61, 0xa1, 0x2b, 0x40, 0x79, 0x7a, 0xbc, 0x70, 0x26, 0x8d,
	0x33, 0xd4, 0x84, 0x4b, 0x67, 0x3a, 0x9a, 0x91, 0xb9, 0x43, 0x72, 0x1d, 0x85, 0xc1, 0x87, 0x1f,
	0x9d, 0xb5, 0x9f, 0x78, 0xe9, 0x52, 0x24, 0xd9, 0xf3, 0xf6, 0x21, 0x8b, 0x36, 0x6c, 0xb5, 0x66,
	0x51, 0xef, 0x91, 0x2e, 0x23, 0xdf, 0x55, 0x7f, 0xf3, 0x58, 0x7f, 0x0a, 0x96, 0x65, 0x09, 0x3f,
	0xfd, 0x0f, 0x00, 0x00, 0xff, 0xff, 0xa4, 0x57, 0xb1, 0x3c, 0x22, 0x04, 0x00, 0x00,
}
