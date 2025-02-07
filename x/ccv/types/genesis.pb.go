// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: interchain_security/ccv/v1/genesis.proto

package types

import (
	fmt "fmt"
	types1 "github.com/cosmos/ibc-go/modules/core/04-channel/types"
	types "github.com/cosmos/ibc-go/modules/light-clients/07-tendermint/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/tendermint/tendermint/abci/types"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// ChildGenesisState defines the CCV child chain genesis state
type ChildGenesisState struct {
	Disabled        bool   `protobuf:"varint,1,opt,name=disabled,proto3" json:"disabled,omitempty"`
	ParentChannelId string `protobuf:"bytes,2,opt,name=parent_channel_id,json=parentChannelId,proto3" json:"parent_channel_id,omitempty" yaml:"parent_channel_id"`
	NewChain        bool   `protobuf:"varint,3,opt,name=new_chain,json=newChain,proto3" json:"new_chain,omitempty" yaml:"new_chain"`
	// ParentClientState filled in on new chain, nil on restart.
	ParentClientState *types.ClientState `protobuf:"bytes,4,opt,name=parent_client_state,json=parentClientState,proto3" json:"parent_client_state,omitempty" yaml:"parent_client_state"`
	// ParentConsensusState filled in on new chain, nil on restart.
	ParentConsensusState *types.ConsensusState `protobuf:"bytes,5,opt,name=parent_consensus_state,json=parentConsensusState,proto3" json:"parent_consensus_state,omitempty" yaml:"parent_consensus_state"`
	UnbondingSequences   []UnbondingSequence   `protobuf:"bytes,6,rep,name=unbonding_sequences,json=unbondingSequences,proto3" json:"unbonding_sequences" yaml:"unbonding_sequences"`
}

func (m *ChildGenesisState) Reset()         { *m = ChildGenesisState{} }
func (m *ChildGenesisState) String() string { return proto.CompactTextString(m) }
func (*ChildGenesisState) ProtoMessage()    {}
func (*ChildGenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_5e4d0563aa343ce3, []int{0}
}
func (m *ChildGenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ChildGenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ChildGenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ChildGenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChildGenesisState.Merge(m, src)
}
func (m *ChildGenesisState) XXX_Size() int {
	return m.Size()
}
func (m *ChildGenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_ChildGenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_ChildGenesisState proto.InternalMessageInfo

func (m *ChildGenesisState) GetDisabled() bool {
	if m != nil {
		return m.Disabled
	}
	return false
}

func (m *ChildGenesisState) GetParentChannelId() string {
	if m != nil {
		return m.ParentChannelId
	}
	return ""
}

func (m *ChildGenesisState) GetNewChain() bool {
	if m != nil {
		return m.NewChain
	}
	return false
}

func (m *ChildGenesisState) GetParentClientState() *types.ClientState {
	if m != nil {
		return m.ParentClientState
	}
	return nil
}

func (m *ChildGenesisState) GetParentConsensusState() *types.ConsensusState {
	if m != nil {
		return m.ParentConsensusState
	}
	return nil
}

func (m *ChildGenesisState) GetUnbondingSequences() []UnbondingSequence {
	if m != nil {
		return m.UnbondingSequences
	}
	return nil
}

// UnbondingSequence defines the genesis information for each unbonding packet sequence.
type UnbondingSequence struct {
	Sequence        uint64        `protobuf:"varint,1,opt,name=sequence,proto3" json:"sequence,omitempty"`
	UnbondingTime   uint64        `protobuf:"varint,2,opt,name=unbonding_time,json=unbondingTime,proto3" json:"unbonding_time,omitempty" yaml:"unbonding_time"`
	UnbondingPacket types1.Packet `protobuf:"bytes,3,opt,name=unbonding_packet,json=unbondingPacket,proto3" json:"unbonding_packet" yaml:"unbonding_packet"`
}

func (m *UnbondingSequence) Reset()         { *m = UnbondingSequence{} }
func (m *UnbondingSequence) String() string { return proto.CompactTextString(m) }
func (*UnbondingSequence) ProtoMessage()    {}
func (*UnbondingSequence) Descriptor() ([]byte, []int) {
	return fileDescriptor_5e4d0563aa343ce3, []int{1}
}
func (m *UnbondingSequence) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UnbondingSequence) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UnbondingSequence.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UnbondingSequence) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UnbondingSequence.Merge(m, src)
}
func (m *UnbondingSequence) XXX_Size() int {
	return m.Size()
}
func (m *UnbondingSequence) XXX_DiscardUnknown() {
	xxx_messageInfo_UnbondingSequence.DiscardUnknown(m)
}

var xxx_messageInfo_UnbondingSequence proto.InternalMessageInfo

func (m *UnbondingSequence) GetSequence() uint64 {
	if m != nil {
		return m.Sequence
	}
	return 0
}

func (m *UnbondingSequence) GetUnbondingTime() uint64 {
	if m != nil {
		return m.UnbondingTime
	}
	return 0
}

func (m *UnbondingSequence) GetUnbondingPacket() types1.Packet {
	if m != nil {
		return m.UnbondingPacket
	}
	return types1.Packet{}
}

// ParentGenesisState defines the CCV parent chain genesis state
type ParentGenesisState struct {
	ChildStates []ChildState `protobuf:"bytes,1,rep,name=child_states,json=childStates,proto3" json:"child_states" yaml:"child_states"`
}

func (m *ParentGenesisState) Reset()         { *m = ParentGenesisState{} }
func (m *ParentGenesisState) String() string { return proto.CompactTextString(m) }
func (*ParentGenesisState) ProtoMessage()    {}
func (*ParentGenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_5e4d0563aa343ce3, []int{2}
}
func (m *ParentGenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ParentGenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ParentGenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ParentGenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ParentGenesisState.Merge(m, src)
}
func (m *ParentGenesisState) XXX_Size() int {
	return m.Size()
}
func (m *ParentGenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_ParentGenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_ParentGenesisState proto.InternalMessageInfo

func (m *ParentGenesisState) GetChildStates() []ChildState {
	if m != nil {
		return m.ChildStates
	}
	return nil
}

// ChildState defines the state that the parent chain stores for each child chain
type ChildState struct {
	ChainId   string `protobuf:"bytes,1,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty" yaml:"chain_id"`
	ChannelId string `protobuf:"bytes,2,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty" yaml:"channel_id"`
	Status    Status `protobuf:"varint,3,opt,name=status,proto3,enum=interchain_security.ccv.v1.Status" json:"status,omitempty"`
}

func (m *ChildState) Reset()         { *m = ChildState{} }
func (m *ChildState) String() string { return proto.CompactTextString(m) }
func (*ChildState) ProtoMessage()    {}
func (*ChildState) Descriptor() ([]byte, []int) {
	return fileDescriptor_5e4d0563aa343ce3, []int{3}
}
func (m *ChildState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ChildState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ChildState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ChildState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChildState.Merge(m, src)
}
func (m *ChildState) XXX_Size() int {
	return m.Size()
}
func (m *ChildState) XXX_DiscardUnknown() {
	xxx_messageInfo_ChildState.DiscardUnknown(m)
}

var xxx_messageInfo_ChildState proto.InternalMessageInfo

func (m *ChildState) GetChainId() string {
	if m != nil {
		return m.ChainId
	}
	return ""
}

func (m *ChildState) GetChannelId() string {
	if m != nil {
		return m.ChannelId
	}
	return ""
}

func (m *ChildState) GetStatus() Status {
	if m != nil {
		return m.Status
	}
	return UNINITIALIZED
}

func init() {
	proto.RegisterType((*ChildGenesisState)(nil), "interchain_security.ccv.v1.ChildGenesisState")
	proto.RegisterType((*UnbondingSequence)(nil), "interchain_security.ccv.v1.UnbondingSequence")
	proto.RegisterType((*ParentGenesisState)(nil), "interchain_security.ccv.v1.ParentGenesisState")
	proto.RegisterType((*ChildState)(nil), "interchain_security.ccv.v1.ChildState")
}

func init() {
	proto.RegisterFile("interchain_security/ccv/v1/genesis.proto", fileDescriptor_5e4d0563aa343ce3)
}

var fileDescriptor_5e4d0563aa343ce3 = []byte{
	// 690 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x54, 0xc1, 0x6e, 0xd3, 0x4a,
	0x14, 0x8d, 0x5f, 0xf3, 0xfa, 0x92, 0xc9, 0x7b, 0x6d, 0xe3, 0xb4, 0x0f, 0x93, 0x82, 0x9d, 0x8e,
	0x10, 0x8a, 0x84, 0x3a, 0x56, 0x02, 0xab, 0xae, 0x90, 0xb3, 0x80, 0x8a, 0x4d, 0xe5, 0xc2, 0x86,
	0x4d, 0x64, 0x8f, 0x07, 0x67, 0x44, 0x32, 0x0e, 0x99, 0x71, 0x4a, 0x05, 0x2b, 0x36, 0x6c, 0xf9,
	0x12, 0xd6, 0x7c, 0x42, 0x97, 0x5d, 0x22, 0x16, 0x16, 0x6a, 0xff, 0xc0, 0x5f, 0x80, 0x3c, 0x63,
	0x3b, 0x49, 0x5b, 0xca, 0xee, 0xfa, 0xce, 0x39, 0x67, 0xce, 0xbd, 0xbe, 0x77, 0x40, 0x97, 0x32,
	0x41, 0x66, 0x78, 0xe4, 0x51, 0x36, 0xe4, 0x04, 0xc7, 0x33, 0x2a, 0x4e, 0x6d, 0x8c, 0xe7, 0xf6,
	0xbc, 0x67, 0x87, 0x84, 0x11, 0x4e, 0x39, 0x9a, 0xce, 0x22, 0x11, 0xe9, 0xed, 0x1b, 0x90, 0x08,
	0xe3, 0x39, 0x9a, 0xf7, 0xda, 0x36, 0xf5, 0xb1, 0x3d, 0xa6, 0xe1, 0x48, 0xe0, 0x31, 0x25, 0x4c,
	0x70, 0x5b, 0x10, 0x16, 0x90, 0xd9, 0x84, 0x32, 0x91, 0x29, 0x2d, 0xbe, 0x94, 0x58, 0x7b, 0x2f,
	0x23, 0xe0, 0x68, 0x46, 0x6c, 0x3c, 0xf2, 0x18, 0x23, 0xe3, 0x0c, 0x95, 0x87, 0x39, 0x64, 0x3b,
	0x8c, 0xc2, 0x48, 0x86, 0x76, 0x16, 0xe5, 0xd9, 0xdd, 0x25, 0x61, 0xcf, 0xc7, 0xd4, 0x16, 0xa7,
	0x53, 0x92, 0x5b, 0x6c, 0x3f, 0xb8, 0xa5, 0x98, 0xcc, 0xa9, 0x44, 0xc1, 0x6f, 0x55, 0xd0, 0x1c,
	0x8c, 0xe8, 0x38, 0x78, 0xa6, 0xea, 0x3b, 0x16, 0x9e, 0x20, 0x7a, 0x1b, 0xd4, 0x02, 0xca, 0x3d,
	0x7f, 0x4c, 0x02, 0x43, 0xeb, 0x68, 0xdd, 0x9a, 0x5b, 0x7e, 0xeb, 0xcf, 0x41, 0x73, 0xea, 0xcd,
	0x08, 0x13, 0xc3, 0xdc, 0xe2, 0x90, 0x06, 0xc6, 0x5f, 0x1d, 0xad, 0x5b, 0x77, 0xee, 0xa5, 0x89,
	0x65, 0x9c, 0x7a, 0x93, 0xf1, 0x01, 0xbc, 0x06, 0x81, 0xee, 0xa6, 0xca, 0x0d, 0x54, 0xea, 0x30,
	0xd0, 0x7b, 0xa0, 0xce, 0xc8, 0xc9, 0x50, 0x5a, 0x34, 0xd6, 0xb2, 0x6b, 0x9c, 0xed, 0x34, 0xb1,
	0xb6, 0x94, 0x42, 0x79, 0x04, 0xdd, 0x1a, 0x23, 0x27, 0x83, 0x2c, 0xd4, 0x3f, 0x80, 0x56, 0xa1,
	0x2c, 0x7b, 0x3b, 0xe4, 0x99, 0x5f, 0xa3, 0xda, 0xd1, 0xba, 0x8d, 0xfe, 0x23, 0x44, 0x7d, 0x8c,
	0x96, 0x3b, 0x8f, 0x96, 0x7a, 0x3d, 0xef, 0xa1, 0x81, 0xcc, 0xca, 0x12, 0x1d, 0x33, 0x4d, 0xac,
	0xf6, 0xaa, 0xd7, 0x25, 0x45, 0xe8, 0xe6, 0x45, 0x2e, 0x51, 0xf4, 0xcf, 0x1a, 0xf8, 0xbf, 0xc0,
	0x46, 0x8c, 0x13, 0xc6, 0x63, 0x9e, 0x1b, 0xf8, 0x5b, 0x1a, 0x40, 0x7f, 0x34, 0x50, 0xd0, 0x94,
	0x87, 0xbd, 0x34, 0xb1, 0xee, 0xaf, 0x7a, 0x58, 0xd5, 0x85, 0xee, 0x76, 0x6e, 0x63, 0x85, 0xa8,
	0x7f, 0xd2, 0x40, 0x2b, 0x66, 0x7e, 0xc4, 0x02, 0xca, 0xc2, 0x21, 0x27, 0xef, 0x62, 0xc2, 0x30,
	0xe1, 0xc6, 0x7a, 0x67, 0xad, 0xdb, 0xe8, 0xef, 0xa3, 0xdf, 0x4f, 0x27, 0x7a, 0x55, 0xd0, 0x8e,
	0x73, 0x96, 0x03, 0xcf, 0x12, 0xab, 0xb2, 0xe8, 0xc6, 0x0d, 0xba, 0xd0, 0xd5, 0xe3, 0xab, 0x34,
	0x0e, 0x7f, 0x68, 0xa0, 0x79, 0x4d, 0x2d, 0x1b, 0x9d, 0x82, 0x27, 0x47, 0xa7, 0xea, 0x96, 0xdf,
	0xfa, 0x53, 0xb0, 0xb1, 0x50, 0x17, 0x74, 0x42, 0xe4, 0xdc, 0x54, 0x9d, 0xbb, 0x69, 0x62, 0xed,
	0x5c, 0xbd, 0x3d, 0x3b, 0x87, 0xee, 0x7f, 0x65, 0xe2, 0x25, 0x9d, 0x10, 0x3d, 0x04, 0x5b, 0x0b,
	0xc4, 0xd4, 0xc3, 0x6f, 0x89, 0x90, 0x93, 0xd3, 0xe8, 0xef, 0xca, 0xde, 0x67, 0x5b, 0x84, 0x8a,
	0xd5, 0x99, 0xf7, 0xd0, 0x91, 0x84, 0x38, 0x56, 0x5e, 0xe2, 0x9d, 0xab, 0x97, 0x28, 0x09, 0xe8,
	0x6e, 0x96, 0x29, 0xc5, 0x80, 0x1f, 0x81, 0x7e, 0x24, 0x3b, 0xbf, 0xb2, 0x17, 0x6f, 0xc0, 0xbf,
	0x38, 0x5b, 0x16, 0xf5, 0x77, 0xb8, 0xa1, 0xc9, 0x7e, 0x3f, 0xbc, 0xad, 0xdf, 0x72, 0xb9, 0xd4,
	0xef, 0xde, 0xcd, 0x5d, 0xb4, 0x94, 0x8b, 0x65, 0x25, 0xe8, 0x36, 0x70, 0x09, 0xe4, 0xf0, 0xab,
	0x06, 0xc0, 0x82, 0xa8, 0x23, 0x50, 0x53, 0xe2, 0x54, 0xad, 0x63, 0xdd, 0x69, 0xa5, 0x89, 0xb5,
	0x59, 0xc8, 0xa8, 0x13, 0xe8, 0xfe, 0x23, 0xc3, 0xc3, 0x40, 0x7f, 0x02, 0xc0, 0xb5, 0xdd, 0xdc,
	0x49, 0x13, 0xab, 0x59, 0x32, 0xca, 0xa5, 0xac, 0xe3, 0x72, 0x1d, 0x0f, 0xc0, 0x7a, 0x66, 0x26,
	0xe6, 0xb2, 0xa3, 0x1b, 0x7d, 0x78, 0x5b, 0x59, 0xc7, 0x12, 0xe9, 0xe6, 0x0c, 0xe7, 0xc5, 0xd9,
	0x85, 0xa9, 0x9d, 0x5f, 0x98, 0xda, 0xcf, 0x0b, 0x53, 0xfb, 0x72, 0x69, 0x56, 0xce, 0x2f, 0xcd,
	0xca, 0xf7, 0x4b, 0xb3, 0xf2, 0xba, 0x17, 0x52, 0x31, 0x8a, 0x7d, 0x84, 0xa3, 0x89, 0x8d, 0x23,
	0x3e, 0x89, 0xb8, 0xbd, 0x90, 0xdd, 0x2f, 0x1f, 0xa6, 0xf7, 0xf2, 0x69, 0x92, 0xef, 0x97, 0xbf,
	0x2e, 0x9f, 0xa6, 0xc7, 0xbf, 0x02, 0x00, 0x00, 0xff, 0xff, 0xad, 0x22, 0x91, 0xbe, 0x8f, 0x05,
	0x00, 0x00,
}

func (m *ChildGenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ChildGenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ChildGenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.UnbondingSequences) > 0 {
		for iNdEx := len(m.UnbondingSequences) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.UnbondingSequences[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if m.ParentConsensusState != nil {
		{
			size, err := m.ParentConsensusState.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	if m.ParentClientState != nil {
		{
			size, err := m.ParentClientState.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if m.NewChain {
		i--
		if m.NewChain {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	if len(m.ParentChannelId) > 0 {
		i -= len(m.ParentChannelId)
		copy(dAtA[i:], m.ParentChannelId)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.ParentChannelId)))
		i--
		dAtA[i] = 0x12
	}
	if m.Disabled {
		i--
		if m.Disabled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *UnbondingSequence) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UnbondingSequence) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UnbondingSequence) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.UnbondingPacket.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if m.UnbondingTime != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.UnbondingTime))
		i--
		dAtA[i] = 0x10
	}
	if m.Sequence != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.Sequence))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *ParentGenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ParentGenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ParentGenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ChildStates) > 0 {
		for iNdEx := len(m.ChildStates) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ChildStates[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *ChildState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ChildState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ChildState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Status != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x18
	}
	if len(m.ChannelId) > 0 {
		i -= len(m.ChannelId)
		copy(dAtA[i:], m.ChannelId)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.ChannelId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ChainId) > 0 {
		i -= len(m.ChainId)
		copy(dAtA[i:], m.ChainId)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.ChainId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ChildGenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Disabled {
		n += 2
	}
	l = len(m.ParentChannelId)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.NewChain {
		n += 2
	}
	if m.ParentClientState != nil {
		l = m.ParentClientState.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.ParentConsensusState != nil {
		l = m.ParentConsensusState.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	if len(m.UnbondingSequences) > 0 {
		for _, e := range m.UnbondingSequences {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func (m *UnbondingSequence) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Sequence != 0 {
		n += 1 + sovGenesis(uint64(m.Sequence))
	}
	if m.UnbondingTime != 0 {
		n += 1 + sovGenesis(uint64(m.UnbondingTime))
	}
	l = m.UnbondingPacket.Size()
	n += 1 + l + sovGenesis(uint64(l))
	return n
}

func (m *ParentGenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.ChildStates) > 0 {
		for _, e := range m.ChildStates {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func (m *ChildState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ChainId)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = len(m.ChannelId)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.Status != 0 {
		n += 1 + sovGenesis(uint64(m.Status))
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ChildGenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ChildGenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ChildGenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Disabled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Disabled = bool(v != 0)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ParentChannelId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ParentChannelId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NewChain", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.NewChain = bool(v != 0)
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ParentClientState", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ParentClientState == nil {
				m.ParentClientState = &types.ClientState{}
			}
			if err := m.ParentClientState.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ParentConsensusState", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ParentConsensusState == nil {
				m.ParentConsensusState = &types.ConsensusState{}
			}
			if err := m.ParentConsensusState.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UnbondingSequences", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UnbondingSequences = append(m.UnbondingSequences, UnbondingSequence{})
			if err := m.UnbondingSequences[len(m.UnbondingSequences)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *UnbondingSequence) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: UnbondingSequence: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UnbondingSequence: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sequence", wireType)
			}
			m.Sequence = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Sequence |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UnbondingTime", wireType)
			}
			m.UnbondingTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UnbondingTime |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UnbondingPacket", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.UnbondingPacket.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ParentGenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ParentGenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ParentGenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChildStates", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChildStates = append(m.ChildStates, ChildState{})
			if err := m.ChildStates[len(m.ChildStates)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ChildState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ChildState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ChildState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChannelId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChannelId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= Status(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
