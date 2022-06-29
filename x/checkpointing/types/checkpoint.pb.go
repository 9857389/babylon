// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: babylon/checkpointing/checkpoint.proto

package types

import (
	bytes "bytes"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/regen-network/cosmos-proto"
	_ "google.golang.org/protobuf/types/known/timestamppb"
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

// CkptStatus is the status of a checkpoint.
type CheckpointStatus int32

const (
	// UNCHECKPOINTED defines a checkpoint that is checkpointed on BTC.
	Uncheckpointed CheckpointStatus = 0
	// UNCONFIRMED defines a validator that is checkpointed on BTC but not confirmed.
	Unconfirmed CheckpointStatus = 1
	// CONFIRMED defines a validator that is confirmed on BTC.
	Confirmed CheckpointStatus = 2
)

var CheckpointStatus_name = map[int32]string{
	0: "CKPT_STATUS_UNCHECKPOINTED",
	1: "CKPT_STATUS_UNCONFIRMED",
	2: "CKPT_STATUS_CONFIRMED",
}

var CheckpointStatus_value = map[string]int32{
	"CKPT_STATUS_UNCHECKPOINTED": 0,
	"CKPT_STATUS_UNCONFIRMED":    1,
	"CKPT_STATUS_CONFIRMED":      2,
}

func (x CheckpointStatus) String() string {
	return proto.EnumName(CheckpointStatus_name, int32(x))
}

func (CheckpointStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_63ff05f0a47b36f7, []int{0}
}

// RawCheckpoint wraps the bls multi sig with meta data
type RawCheckpoint struct {
	// epoch_num defines the epoch number the raw checkpoint is for
	EpochNum uint64 `protobuf:"varint,1,opt,name=epoch_num,json=epochNum,proto3" json:"epoch_num,omitempty"`
	// last_commit_hash defines the 'LastCommitHash' that individual bls sigs are signed on
	LastCommitHash []byte `protobuf:"bytes,2,opt,name=last_commit_hash,json=lastCommitHash,proto3" json:"last_commit_hash,omitempty"`
	// bitmap defines the bitmap that indicates the signers of the bls multi sig
	Bitmap []byte `protobuf:"bytes,3,opt,name=bitmap,proto3" json:"bitmap,omitempty"`
	// bls_multi_sig defines the multi sig that is aggregated from individual bls sigs
	BlsMultiSig []byte `protobuf:"bytes,4,opt,name=bls_multi_sig,json=blsMultiSig,proto3" json:"bls_multi_sig,omitempty"`
}

func (m *RawCheckpoint) Reset()         { *m = RawCheckpoint{} }
func (m *RawCheckpoint) String() string { return proto.CompactTextString(m) }
func (*RawCheckpoint) ProtoMessage()    {}
func (*RawCheckpoint) Descriptor() ([]byte, []int) {
	return fileDescriptor_63ff05f0a47b36f7, []int{0}
}
func (m *RawCheckpoint) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RawCheckpoint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RawCheckpoint.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RawCheckpoint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RawCheckpoint.Merge(m, src)
}
func (m *RawCheckpoint) XXX_Size() int {
	return m.Size()
}
func (m *RawCheckpoint) XXX_DiscardUnknown() {
	xxx_messageInfo_RawCheckpoint.DiscardUnknown(m)
}

var xxx_messageInfo_RawCheckpoint proto.InternalMessageInfo

func (m *RawCheckpoint) GetEpochNum() uint64 {
	if m != nil {
		return m.EpochNum
	}
	return 0
}

func (m *RawCheckpoint) GetLastCommitHash() []byte {
	if m != nil {
		return m.LastCommitHash
	}
	return nil
}

func (m *RawCheckpoint) GetBitmap() []byte {
	if m != nil {
		return m.Bitmap
	}
	return nil
}

func (m *RawCheckpoint) GetBlsMultiSig() []byte {
	if m != nil {
		return m.BlsMultiSig
	}
	return nil
}

// RawCheckpointWithMeta wraps the raw checkpoint with meta data.
type RawCheckpointWithMeta struct {
	Ckpt *RawCheckpoint `protobuf:"bytes,1,opt,name=ckpt,proto3" json:"ckpt,omitempty"`
	// status defines the status of the checkpoint
	Status CheckpointStatus `protobuf:"varint,2,opt,name=status,proto3,enum=babylon.checkpointing.v1.CheckpointStatus" json:"status,omitempty"`
}

func (m *RawCheckpointWithMeta) Reset()         { *m = RawCheckpointWithMeta{} }
func (m *RawCheckpointWithMeta) String() string { return proto.CompactTextString(m) }
func (*RawCheckpointWithMeta) ProtoMessage()    {}
func (*RawCheckpointWithMeta) Descriptor() ([]byte, []int) {
	return fileDescriptor_63ff05f0a47b36f7, []int{1}
}
func (m *RawCheckpointWithMeta) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RawCheckpointWithMeta) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RawCheckpointWithMeta.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RawCheckpointWithMeta) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RawCheckpointWithMeta.Merge(m, src)
}
func (m *RawCheckpointWithMeta) XXX_Size() int {
	return m.Size()
}
func (m *RawCheckpointWithMeta) XXX_DiscardUnknown() {
	xxx_messageInfo_RawCheckpointWithMeta.DiscardUnknown(m)
}

var xxx_messageInfo_RawCheckpointWithMeta proto.InternalMessageInfo

func (m *RawCheckpointWithMeta) GetCkpt() *RawCheckpoint {
	if m != nil {
		return m.Ckpt
	}
	return nil
}

func (m *RawCheckpointWithMeta) GetStatus() CheckpointStatus {
	if m != nil {
		return m.Status
	}
	return Uncheckpointed
}

// BlsSig wraps the bls sig with meta data.
type BlsSig struct {
	// epoch_num defines the epoch number that the bls sig is signed on
	EpochNum uint64 `protobuf:"varint,1,opt,name=epoch_num,json=epochNum,proto3" json:"epoch_num,omitempty"`
	// last_commit_hash defines the 'LastCommitHash' that the bls sig is signed on
	LastCommitHash []byte `protobuf:"bytes,2,opt,name=last_commit_hash,json=lastCommitHash,proto3" json:"last_commit_hash,omitempty"`
	// bls_sig defines the actual bls sig
	BlsSig []byte `protobuf:"bytes,3,opt,name=bls_sig,json=blsSig,proto3" json:"bls_sig,omitempty"`
	// can't find cosmos_proto.scalar when compiling due to cosmos v0.45.4 does not support scalar
	// string signer_address = 5 [(cosmos_proto.scalar) = "cosmos.AddressString"];
	// the signer_address defines the address of the signer
	SignerAddress string `protobuf:"bytes,5,opt,name=signer_address,json=signerAddress,proto3" json:"signer_address,omitempty"`
}

func (m *BlsSig) Reset()         { *m = BlsSig{} }
func (m *BlsSig) String() string { return proto.CompactTextString(m) }
func (*BlsSig) ProtoMessage()    {}
func (*BlsSig) Descriptor() ([]byte, []int) {
	return fileDescriptor_63ff05f0a47b36f7, []int{2}
}
func (m *BlsSig) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BlsSig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BlsSig.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BlsSig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlsSig.Merge(m, src)
}
func (m *BlsSig) XXX_Size() int {
	return m.Size()
}
func (m *BlsSig) XXX_DiscardUnknown() {
	xxx_messageInfo_BlsSig.DiscardUnknown(m)
}

var xxx_messageInfo_BlsSig proto.InternalMessageInfo

func (m *BlsSig) GetEpochNum() uint64 {
	if m != nil {
		return m.EpochNum
	}
	return 0
}

func (m *BlsSig) GetLastCommitHash() []byte {
	if m != nil {
		return m.LastCommitHash
	}
	return nil
}

func (m *BlsSig) GetBlsSig() []byte {
	if m != nil {
		return m.BlsSig
	}
	return nil
}

func (m *BlsSig) GetSignerAddress() string {
	if m != nil {
		return m.SignerAddress
	}
	return ""
}

func init() {
	proto.RegisterEnum("babylon.checkpointing.v1.CheckpointStatus", CheckpointStatus_name, CheckpointStatus_value)
	proto.RegisterType((*RawCheckpoint)(nil), "babylon.checkpointing.v1.RawCheckpoint")
	proto.RegisterType((*RawCheckpointWithMeta)(nil), "babylon.checkpointing.v1.RawCheckpointWithMeta")
	proto.RegisterType((*BlsSig)(nil), "babylon.checkpointing.v1.BlsSig")
}

func init() {
	proto.RegisterFile("babylon/checkpointing/checkpoint.proto", fileDescriptor_63ff05f0a47b36f7)
}

var fileDescriptor_63ff05f0a47b36f7 = []byte{
	// 517 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x53, 0x41, 0x6f, 0xd3, 0x3e,
	0x1c, 0xad, 0xf7, 0xef, 0xbf, 0x50, 0x97, 0x96, 0xca, 0x62, 0x2c, 0x04, 0x29, 0x8b, 0x2a, 0x01,
	0xd5, 0x84, 0x12, 0x51, 0xc4, 0x05, 0x4e, 0x6b, 0x56, 0xb4, 0x69, 0x6a, 0x3b, 0xa5, 0xad, 0x90,
	0xb8, 0x44, 0x4e, 0x9a, 0x25, 0xd6, 0xe2, 0x38, 0xaa, 0x1d, 0x60, 0xdf, 0x00, 0xf5, 0x02, 0x47,
	0x2e, 0x95, 0x90, 0xe0, 0xc3, 0x70, 0xdc, 0x91, 0x23, 0x6a, 0x2f, 0x7c, 0x0c, 0x14, 0xa7, 0xa5,
	0x6b, 0xa5, 0xdd, 0xb8, 0xe5, 0xbd, 0xdf, 0x7b, 0xf1, 0xf3, 0xb3, 0x0d, 0x1f, 0xbb, 0xd8, 0xbd,
	0x8c, 0x58, 0x6c, 0x7a, 0xa1, 0xef, 0x5d, 0x24, 0x8c, 0xc4, 0x82, 0xc4, 0xc1, 0x35, 0x64, 0x24,
	0x13, 0x26, 0x18, 0x52, 0x96, 0x3a, 0x63, 0x43, 0x67, 0xbc, 0x7b, 0xa6, 0x3e, 0xf0, 0x18, 0xa7,
	0x8c, 0x3b, 0x52, 0x67, 0xe6, 0x20, 0x37, 0xa9, 0xf7, 0x02, 0x16, 0xb0, 0x9c, 0xcf, 0xbe, 0x96,
	0xec, 0x7e, 0xc0, 0x58, 0x10, 0xf9, 0xa6, 0x44, 0x6e, 0x7a, 0x6e, 0x0a, 0x42, 0x7d, 0x2e, 0x30,
	0x4d, 0x72, 0x41, 0xe3, 0x0b, 0x80, 0x55, 0x1b, 0xbf, 0xb7, 0xfe, 0xae, 0x84, 0x1e, 0xc2, 0xb2,
	0x9f, 0x30, 0x2f, 0x74, 0xe2, 0x94, 0x2a, 0x40, 0x07, 0xcd, 0xa2, 0x7d, 0x5b, 0x12, 0xbd, 0x94,
	0xa2, 0x26, 0xac, 0x47, 0x98, 0x0b, 0xc7, 0x63, 0x94, 0x12, 0xe1, 0x84, 0x98, 0x87, 0xca, 0x8e,
	0x0e, 0x9a, 0x77, 0xec, 0x5a, 0xc6, 0x5b, 0x92, 0x3e, 0xc6, 0x3c, 0x44, 0xf7, 0x61, 0xc9, 0x25,
	0x82, 0xe2, 0x44, 0xf9, 0x4f, 0xce, 0x97, 0x08, 0x35, 0x60, 0xd5, 0x8d, 0xb8, 0x43, 0xd3, 0x48,
	0x10, 0x87, 0x93, 0x40, 0x29, 0xca, 0x71, 0xc5, 0x8d, 0x78, 0x37, 0xe3, 0x06, 0x24, 0x78, 0x59,
	0xfc, 0xfd, 0x75, 0x1f, 0x64, 0xd1, 0x76, 0x37, 0xa2, 0xbd, 0x21, 0x22, 0xec, 0xfa, 0x02, 0xa3,
	0x57, 0xb0, 0xe8, 0x5d, 0x24, 0x42, 0xa6, 0xab, 0xb4, 0x9e, 0x18, 0x37, 0xf5, 0x65, 0x6c, 0xd8,
	0x6d, 0x69, 0x42, 0x6d, 0x58, 0xe2, 0x02, 0x8b, 0x94, 0xcb, 0xe0, 0xb5, 0xd6, 0xc1, 0xcd, 0xf6,
	0xb5, 0x77, 0x20, 0x1d, 0xf6, 0xd2, 0xd9, 0xf8, 0x04, 0x60, 0xa9, 0x1d, 0xf1, 0x01, 0x09, 0xfe,
	0x55, 0x5d, 0x7b, 0xf0, 0x56, 0x56, 0x4b, 0x56, 0xc8, 0xaa, 0xaf, 0xfc, 0xff, 0x8f, 0x60, 0x8d,
	0x93, 0x20, 0xf6, 0x27, 0x0e, 0x1e, 0x8f, 0x27, 0x3e, 0xe7, 0xca, 0xff, 0x3a, 0x68, 0x96, 0xed,
	0x6a, 0xce, 0x1e, 0xe6, 0xe4, 0xc1, 0x77, 0x00, 0xeb, 0xdb, 0x71, 0x51, 0x0b, 0xaa, 0xd6, 0xe9,
	0xd9, 0xd0, 0x19, 0x0c, 0x0f, 0x87, 0xa3, 0x81, 0x33, 0xea, 0x59, 0xc7, 0x1d, 0xeb, 0xf4, 0xac,
	0x7f, 0xd2, 0x1b, 0x76, 0x8e, 0xea, 0x05, 0x15, 0x4d, 0x67, 0x7a, 0x6d, 0x14, 0xaf, 0xf7, 0xed,
	0x8f, 0xd1, 0x53, 0xb8, 0xb7, 0xe5, 0xe9, 0xf7, 0x5e, 0x9f, 0xd8, 0xdd, 0xce, 0x51, 0x1d, 0xa8,
	0x77, 0xa7, 0x33, 0xbd, 0x32, 0x8a, 0x3d, 0x16, 0x9f, 0x93, 0x09, 0xf5, 0xc7, 0xa8, 0x09, 0x77,
	0xaf, 0xab, 0xd7, 0xda, 0x1d, 0xb5, 0x3a, 0x9d, 0xe9, 0x65, 0x6b, 0xa5, 0x54, 0x8b, 0x1f, 0xbf,
	0x69, 0x85, 0x76, 0xff, 0xc7, 0x5c, 0x03, 0x57, 0x73, 0x0d, 0xfc, 0x9a, 0x6b, 0xe0, 0xf3, 0x42,
	0x2b, 0x5c, 0x2d, 0xb4, 0xc2, 0xcf, 0x85, 0x56, 0x78, 0xfb, 0x22, 0x20, 0x22, 0x4c, 0x5d, 0xc3,
	0x63, 0xd4, 0x5c, 0x1e, 0x88, 0x17, 0x62, 0x12, 0xaf, 0x80, 0xf9, 0x61, 0xeb, 0xd9, 0x88, 0xcb,
	0xc4, 0xe7, 0x6e, 0x49, 0x5e, 0xe3, 0xe7, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x99, 0x75, 0xc9,
	0xc2, 0x5c, 0x03, 0x00, 0x00,
}

func (this *RawCheckpoint) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*RawCheckpoint)
	if !ok {
		that2, ok := that.(RawCheckpoint)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.EpochNum != that1.EpochNum {
		return false
	}
	if !bytes.Equal(this.LastCommitHash, that1.LastCommitHash) {
		return false
	}
	if !bytes.Equal(this.Bitmap, that1.Bitmap) {
		return false
	}
	if !bytes.Equal(this.BlsMultiSig, that1.BlsMultiSig) {
		return false
	}
	return true
}
func (m *RawCheckpoint) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RawCheckpoint) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RawCheckpoint) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.BlsMultiSig) > 0 {
		i -= len(m.BlsMultiSig)
		copy(dAtA[i:], m.BlsMultiSig)
		i = encodeVarintCheckpoint(dAtA, i, uint64(len(m.BlsMultiSig)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Bitmap) > 0 {
		i -= len(m.Bitmap)
		copy(dAtA[i:], m.Bitmap)
		i = encodeVarintCheckpoint(dAtA, i, uint64(len(m.Bitmap)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.LastCommitHash) > 0 {
		i -= len(m.LastCommitHash)
		copy(dAtA[i:], m.LastCommitHash)
		i = encodeVarintCheckpoint(dAtA, i, uint64(len(m.LastCommitHash)))
		i--
		dAtA[i] = 0x12
	}
	if m.EpochNum != 0 {
		i = encodeVarintCheckpoint(dAtA, i, uint64(m.EpochNum))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *RawCheckpointWithMeta) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RawCheckpointWithMeta) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RawCheckpointWithMeta) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Status != 0 {
		i = encodeVarintCheckpoint(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x10
	}
	if m.Ckpt != nil {
		{
			size, err := m.Ckpt.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintCheckpoint(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *BlsSig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BlsSig) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BlsSig) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.SignerAddress) > 0 {
		i -= len(m.SignerAddress)
		copy(dAtA[i:], m.SignerAddress)
		i = encodeVarintCheckpoint(dAtA, i, uint64(len(m.SignerAddress)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.BlsSig) > 0 {
		i -= len(m.BlsSig)
		copy(dAtA[i:], m.BlsSig)
		i = encodeVarintCheckpoint(dAtA, i, uint64(len(m.BlsSig)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.LastCommitHash) > 0 {
		i -= len(m.LastCommitHash)
		copy(dAtA[i:], m.LastCommitHash)
		i = encodeVarintCheckpoint(dAtA, i, uint64(len(m.LastCommitHash)))
		i--
		dAtA[i] = 0x12
	}
	if m.EpochNum != 0 {
		i = encodeVarintCheckpoint(dAtA, i, uint64(m.EpochNum))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintCheckpoint(dAtA []byte, offset int, v uint64) int {
	offset -= sovCheckpoint(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *RawCheckpoint) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.EpochNum != 0 {
		n += 1 + sovCheckpoint(uint64(m.EpochNum))
	}
	l = len(m.LastCommitHash)
	if l > 0 {
		n += 1 + l + sovCheckpoint(uint64(l))
	}
	l = len(m.Bitmap)
	if l > 0 {
		n += 1 + l + sovCheckpoint(uint64(l))
	}
	l = len(m.BlsMultiSig)
	if l > 0 {
		n += 1 + l + sovCheckpoint(uint64(l))
	}
	return n
}

func (m *RawCheckpointWithMeta) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Ckpt != nil {
		l = m.Ckpt.Size()
		n += 1 + l + sovCheckpoint(uint64(l))
	}
	if m.Status != 0 {
		n += 1 + sovCheckpoint(uint64(m.Status))
	}
	return n
}

func (m *BlsSig) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.EpochNum != 0 {
		n += 1 + sovCheckpoint(uint64(m.EpochNum))
	}
	l = len(m.LastCommitHash)
	if l > 0 {
		n += 1 + l + sovCheckpoint(uint64(l))
	}
	l = len(m.BlsSig)
	if l > 0 {
		n += 1 + l + sovCheckpoint(uint64(l))
	}
	l = len(m.SignerAddress)
	if l > 0 {
		n += 1 + l + sovCheckpoint(uint64(l))
	}
	return n
}

func sovCheckpoint(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCheckpoint(x uint64) (n int) {
	return sovCheckpoint(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RawCheckpoint) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCheckpoint
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
			return fmt.Errorf("proto: RawCheckpoint: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RawCheckpoint: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EpochNum", wireType)
			}
			m.EpochNum = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCheckpoint
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EpochNum |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastCommitHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCheckpoint
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthCheckpoint
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthCheckpoint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LastCommitHash = append(m.LastCommitHash[:0], dAtA[iNdEx:postIndex]...)
			if m.LastCommitHash == nil {
				m.LastCommitHash = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bitmap", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCheckpoint
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthCheckpoint
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthCheckpoint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Bitmap = append(m.Bitmap[:0], dAtA[iNdEx:postIndex]...)
			if m.Bitmap == nil {
				m.Bitmap = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlsMultiSig", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCheckpoint
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthCheckpoint
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthCheckpoint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BlsMultiSig = append(m.BlsMultiSig[:0], dAtA[iNdEx:postIndex]...)
			if m.BlsMultiSig == nil {
				m.BlsMultiSig = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCheckpoint(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCheckpoint
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
func (m *RawCheckpointWithMeta) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCheckpoint
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
			return fmt.Errorf("proto: RawCheckpointWithMeta: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RawCheckpointWithMeta: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ckpt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCheckpoint
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
				return ErrInvalidLengthCheckpoint
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCheckpoint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Ckpt == nil {
				m.Ckpt = &RawCheckpoint{}
			}
			if err := m.Ckpt.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCheckpoint
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= CheckpointStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipCheckpoint(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCheckpoint
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
func (m *BlsSig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCheckpoint
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
			return fmt.Errorf("proto: BlsSig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BlsSig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EpochNum", wireType)
			}
			m.EpochNum = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCheckpoint
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EpochNum |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastCommitHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCheckpoint
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthCheckpoint
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthCheckpoint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LastCommitHash = append(m.LastCommitHash[:0], dAtA[iNdEx:postIndex]...)
			if m.LastCommitHash == nil {
				m.LastCommitHash = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlsSig", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCheckpoint
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthCheckpoint
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthCheckpoint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BlsSig = append(m.BlsSig[:0], dAtA[iNdEx:postIndex]...)
			if m.BlsSig == nil {
				m.BlsSig = []byte{}
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SignerAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCheckpoint
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
				return ErrInvalidLengthCheckpoint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCheckpoint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SignerAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCheckpoint(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCheckpoint
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
func skipCheckpoint(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCheckpoint
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
					return 0, ErrIntOverflowCheckpoint
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
					return 0, ErrIntOverflowCheckpoint
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
				return 0, ErrInvalidLengthCheckpoint
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCheckpoint
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCheckpoint
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCheckpoint        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCheckpoint          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCheckpoint = fmt.Errorf("proto: unexpected end of group")
)