// Code generated by protoc-gen-go. DO NOT EDIT.
// source: orderer/configuration.proto

package orderer

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ConsensusType struct {
	Type string `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
}

func (m *ConsensusType) Reset()                    { *m = ConsensusType{} }
func (m *ConsensusType) String() string            { return proto.CompactTextString(m) }
func (*ConsensusType) ProtoMessage()               {}
func (*ConsensusType) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *ConsensusType) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

type BatchSize struct {
	// Simply specified as number of messages for now, in the future
	// we may want to allow this to be specified by size in bytes
	MaxMessageCount uint32 `protobuf:"varint,1,opt,name=max_message_count,json=maxMessageCount" json:"max_message_count,omitempty"`
	// The byte count of the serialized messages in a batch cannot
	// exceed this value.
	AbsoluteMaxBytes uint32 `protobuf:"varint,2,opt,name=absolute_max_bytes,json=absoluteMaxBytes" json:"absolute_max_bytes,omitempty"`
	// The byte count of the serialized messages in a batch should not
	// exceed this value.
	PreferredMaxBytes uint32 `protobuf:"varint,3,opt,name=preferred_max_bytes,json=preferredMaxBytes" json:"preferred_max_bytes,omitempty"`
}

func (m *BatchSize) Reset()                    { *m = BatchSize{} }
func (m *BatchSize) String() string            { return proto.CompactTextString(m) }
func (*BatchSize) ProtoMessage()               {}
func (*BatchSize) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *BatchSize) GetMaxMessageCount() uint32 {
	if m != nil {
		return m.MaxMessageCount
	}
	return 0
}

func (m *BatchSize) GetAbsoluteMaxBytes() uint32 {
	if m != nil {
		return m.AbsoluteMaxBytes
	}
	return 0
}

func (m *BatchSize) GetPreferredMaxBytes() uint32 {
	if m != nil {
		return m.PreferredMaxBytes
	}
	return 0
}

type BatchTimeout struct {
	// Any duration string parseable by ParseDuration():
	// https://golang.org/pkg/time/#ParseDuration
	Timeout string `protobuf:"bytes,1,opt,name=timeout" json:"timeout,omitempty"`
}

func (m *BatchTimeout) Reset()                    { *m = BatchTimeout{} }
func (m *BatchTimeout) String() string            { return proto.CompactTextString(m) }
func (*BatchTimeout) ProtoMessage()               {}
func (*BatchTimeout) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *BatchTimeout) GetTimeout() string {
	if m != nil {
		return m.Timeout
	}
	return ""
}

// Carries a list of bootstrap brokers, i.e. this is not the exclusive set of
// brokers an ordering service
type KafkaBrokers struct {
	// Each broker here should be identified using the (IP|host):port notation,
	// e.g. 127.0.0.1:7050, or localhost:7050 are valid entries
	Brokers []string `protobuf:"bytes,1,rep,name=brokers" json:"brokers,omitempty"`
}

func (m *KafkaBrokers) Reset()                    { *m = KafkaBrokers{} }
func (m *KafkaBrokers) String() string            { return proto.CompactTextString(m) }
func (*KafkaBrokers) ProtoMessage()               {}
func (*KafkaBrokers) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func (m *KafkaBrokers) GetBrokers() []string {
	if m != nil {
		return m.Brokers
	}
	return nil
}

// ChannelRestrictions is the mssage which conveys restrictions on channel creation for an orderer
type ChannelRestrictions struct {
	MaxCount uint64 `protobuf:"varint,1,opt,name=max_count,json=maxCount" json:"max_count,omitempty"`
}

func (m *ChannelRestrictions) Reset()                    { *m = ChannelRestrictions{} }
func (m *ChannelRestrictions) String() string            { return proto.CompactTextString(m) }
func (*ChannelRestrictions) ProtoMessage()               {}
func (*ChannelRestrictions) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

func (m *ChannelRestrictions) GetMaxCount() uint64 {
	if m != nil {
		return m.MaxCount
	}
	return 0
}

func init() {
	proto.RegisterType((*ConsensusType)(nil), "orderer.ConsensusType")
	proto.RegisterType((*BatchSize)(nil), "orderer.BatchSize")
	proto.RegisterType((*BatchTimeout)(nil), "orderer.BatchTimeout")
	proto.RegisterType((*KafkaBrokers)(nil), "orderer.KafkaBrokers")
	proto.RegisterType((*ChannelRestrictions)(nil), "orderer.ChannelRestrictions")
}

func init() { proto.RegisterFile("orderer/configuration.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 318 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x91, 0x4f, 0x4b, 0xc3, 0x40,
	0x10, 0xc5, 0x89, 0x2d, 0xd6, 0x2e, 0x16, 0xed, 0xf6, 0x12, 0xe8, 0xa5, 0x44, 0x0f, 0x45, 0x24,
	0x01, 0xf5, 0x13, 0xa4, 0x47, 0xe9, 0x25, 0xf6, 0x24, 0x42, 0xd9, 0xa4, 0x93, 0x64, 0x49, 0xb3,
	0x13, 0xf6, 0x0f, 0x24, 0x7e, 0x0f, 0xbf, 0xaf, 0xec, 0x26, 0xb1, 0xde, 0xde, 0xcc, 0xfc, 0x1e,
	0xbc, 0x7d, 0x4b, 0xd6, 0x28, 0x4f, 0x20, 0x41, 0x46, 0x19, 0x8a, 0x9c, 0x17, 0x46, 0x32, 0xcd,
	0x51, 0x84, 0x8d, 0x44, 0x8d, 0x74, 0x36, 0x1c, 0x83, 0x07, 0xb2, 0xd8, 0xa1, 0x50, 0x20, 0x94,
	0x51, 0x87, 0xae, 0x01, 0x4a, 0xc9, 0x54, 0x77, 0x0d, 0xf8, 0xde, 0xc6, 0xdb, 0xce, 0x13, 0xa7,
	0x83, 0x1f, 0x8f, 0xcc, 0x63, 0xa6, 0xb3, 0xf2, 0x83, 0x7f, 0x03, 0x7d, 0x22, 0xcb, 0x9a, 0xb5,
	0xc7, 0x1a, 0x94, 0x62, 0x05, 0x1c, 0x33, 0x34, 0x42, 0x3b, 0x7c, 0x91, 0xdc, 0xd5, 0xac, 0xdd,
	0xf7, 0xfb, 0x9d, 0x5d, 0xd3, 0x67, 0x42, 0x59, 0xaa, 0xf0, 0x6c, 0x34, 0x1c, 0xad, 0x29, 0xed,
	0x34, 0x28, 0xff, 0xca, 0xc1, 0xf7, 0xe3, 0x65, 0xcf, 0xda, 0xd8, 0xee, 0x69, 0x48, 0x56, 0x8d,
	0x84, 0x1c, 0xa4, 0x84, 0xd3, 0x3f, 0x7c, 0xe2, 0xf0, 0xe5, 0xdf, 0x69, 0xe4, 0x83, 0x2d, 0xb9,
	0x75, 0xb1, 0x0e, 0xbc, 0x06, 0x34, 0x9a, 0xfa, 0x64, 0xa6, 0x7b, 0x39, 0xc4, 0x1f, 0x47, 0x4b,
	0xbe, 0xb3, 0xbc, 0x62, 0xb1, 0xc4, 0x0a, 0xa4, 0xb2, 0x64, 0xda, 0x4b, 0xdf, 0xdb, 0x4c, 0x2c,
	0x39, 0x8c, 0xc1, 0x0b, 0x59, 0xed, 0x4a, 0x26, 0x04, 0x9c, 0x13, 0x50, 0x5a, 0xf2, 0xcc, 0xb6,
	0xa6, 0xe8, 0x9a, 0xcc, 0x6d, 0xa0, 0xcb, 0x63, 0xa7, 0xc9, 0x4d, 0xcd, 0x5a, 0xf7, 0xca, 0xf8,
	0x8b, 0x3c, 0xa2, 0x2c, 0x42, 0x2e, 0xaa, 0xac, 0x64, 0x5c, 0x5c, 0x84, 0x2b, 0x5b, 0x85, 0x43,
	0xd9, 0x9f, 0x6f, 0x05, 0xd7, 0xa5, 0x49, 0xc3, 0x0c, 0xeb, 0x88, 0x8b, 0xea, 0xcc, 0x52, 0x95,
	0xa3, 0x11, 0x27, 0xf7, 0x33, 0xd1, 0xe8, 0x8a, 0x7a, 0x57, 0x34, 0xb8, 0xd2, 0x6b, 0x37, 0xbf,
	0xfe, 0x06, 0x00, 0x00, 0xff, 0xff, 0xe3, 0x11, 0xe8, 0x51, 0xd1, 0x01, 0x00, 0x00,
}
