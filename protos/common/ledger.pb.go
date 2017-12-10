// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common/ledger.proto

package common

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Contains information about the blockchain ledger such as height, current
// block hash, and previous block hash.
type BlockchainInfo struct {
	Height            uint64 `protobuf:"varint,1,opt,name=height" json:"height,omitempty"`
	CurrentBlockHash  []byte `protobuf:"bytes,2,opt,name=currentBlockHash,proto3" json:"currentBlockHash,omitempty"`
	PreviousBlockHash []byte `protobuf:"bytes,3,opt,name=previousBlockHash,proto3" json:"previousBlockHash,omitempty"`
}

func (m *BlockchainInfo) Reset()                    { *m = BlockchainInfo{} }
func (m *BlockchainInfo) String() string            { return proto.CompactTextString(m) }
func (*BlockchainInfo) ProtoMessage()               {}
func (*BlockchainInfo) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func (m *BlockchainInfo) GetHeight() uint64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *BlockchainInfo) GetCurrentBlockHash() []byte {
	if m != nil {
		return m.CurrentBlockHash
	}
	return nil
}

func (m *BlockchainInfo) GetPreviousBlockHash() []byte {
	if m != nil {
		return m.PreviousBlockHash
	}
	return nil
}

func init() {
	proto.RegisterType((*BlockchainInfo)(nil), "common.BlockchainInfo")
}

func init() { proto.RegisterFile("common/ledger.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 192 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x8e, 0xcd, 0x8a, 0x83, 0x30,
	0x14, 0x46, 0xc9, 0xcc, 0xe0, 0x22, 0x0c, 0xc3, 0x4c, 0x06, 0x8a, 0x4b, 0xe9, 0x4a, 0x8a, 0x98,
	0x85, 0x6f, 0xe0, 0xaa, 0xdd, 0xba, 0xec, 0x2e, 0xc6, 0x98, 0x04, 0x35, 0x57, 0xf2, 0xd3, 0x07,
	0xe8, 0x93, 0x17, 0x93, 0x96, 0x2e, 0x5c, 0x7e, 0xe7, 0xde, 0x03, 0x07, 0xff, 0x73, 0x58, 0x16,
	0x30, 0x74, 0x16, 0x83, 0x14, 0xb6, 0x5e, 0x2d, 0x78, 0x20, 0x59, 0x82, 0xc7, 0x3b, 0xc2, 0x3f,
	0xed, 0x0c, 0x7c, 0xe2, 0x8a, 0x69, 0x73, 0x31, 0x23, 0x90, 0x03, 0xce, 0x94, 0xd0, 0x52, 0xf9,
	0x1c, 0x15, 0xa8, 0xfc, 0xea, 0x9e, 0x8b, 0x9c, 0xf0, 0x2f, 0x0f, 0xd6, 0x0a, 0xe3, 0xa3, 0x70,
	0x66, 0x4e, 0xe5, 0x1f, 0x05, 0x2a, 0xbf, 0xbb, 0x1d, 0x27, 0x15, 0xfe, 0x5b, 0xad, 0xb8, 0x69,
	0x08, 0xee, 0xfd, 0xfc, 0x19, 0x9f, 0xf7, 0x87, 0x96, 0xe3, 0x0a, 0xac, 0xac, 0xb5, 0x99, 0x66,
	0xd6, 0xbb, 0x11, 0x82, 0x19, 0x98, 0xd7, 0x60, 0x36, 0x12, 0xbb, 0x52, 0xb4, 0xab, 0x53, 0xf4,
	0xb5, 0x91, 0xda, 0xab, 0xd0, 0x6f, 0x93, 0xee, 0x24, 0xfa, 0x92, 0x68, 0x92, 0x68, 0x92, 0xfa,
	0x2c, 0xce, 0xe6, 0x11, 0x00, 0x00, 0xff, 0xff, 0x80, 0x88, 0xdf, 0x86, 0x0f, 0x01, 0x00, 0x00,
}
