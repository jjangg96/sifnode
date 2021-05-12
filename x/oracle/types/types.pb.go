// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sifnode/oracle/v1/types.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

// StatusText is an enum used to represent the status of the prophecy
type StatusText int32

const (
	// Default value
	StatusText_STATUS_TEXT_UNSPECIFIED StatusText = 0
	// Pending status
	StatusText_STATUS_TEXT_PENDING StatusText = 1
	// Success status
	StatusText_STATUS_TEXT_SUCCESS StatusText = 2
	// Failed status
	StatusText_STATUS_TEXT_FAILED StatusText = 3
)

var StatusText_name = map[int32]string{
	0: "STATUS_TEXT_UNSPECIFIED",
	1: "STATUS_TEXT_PENDING",
	2: "STATUS_TEXT_SUCCESS",
	3: "STATUS_TEXT_FAILED",
}

var StatusText_value = map[string]int32{
	"STATUS_TEXT_UNSPECIFIED": 0,
	"STATUS_TEXT_PENDING":     1,
	"STATUS_TEXT_SUCCESS":     2,
	"STATUS_TEXT_FAILED":      3,
}

func (x StatusText) String() string {
	return proto.EnumName(StatusText_name, int32(x))
}

func (StatusText) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_dac1b931484f4203, []int{0}
}

// GenesisState - all clp state that must be provided at genesis
// TODO: Add parameters to Genesis state ,such as minimum liquidity required to
// create a pool
type GenesisState struct {
	AddressWhitelist []string `protobuf:"bytes,1,rep,name=address_whitelist,json=addressWhitelist,proto3" json:"address_whitelist,omitempty"`
	AdminAddress     string   `protobuf:"bytes,2,opt,name=admin_address,json=adminAddress,proto3" json:"admin_address,omitempty"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_dac1b931484f4203, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetAddressWhitelist() []string {
	if m != nil {
		return m.AddressWhitelist
	}
	return nil
}

func (m *GenesisState) GetAdminAddress() string {
	if m != nil {
		return m.AdminAddress
	}
	return ""
}

// Claim contrains an arbitrary claim with arbitrary content made by a given
// validator
type Claim struct {
	Id               string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ValidatorAddress string `protobuf:"bytes,2,opt,name=validator_address,json=validatorAddress,proto3" json:"validator_address,omitempty"`
	Content          string `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
}

func (m *Claim) Reset()         { *m = Claim{} }
func (m *Claim) String() string { return proto.CompactTextString(m) }
func (*Claim) ProtoMessage()    {}
func (*Claim) Descriptor() ([]byte, []int) {
	return fileDescriptor_dac1b931484f4203, []int{1}
}
func (m *Claim) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Claim) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Claim.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Claim) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Claim.Merge(m, src)
}
func (m *Claim) XXX_Size() int {
	return m.Size()
}
func (m *Claim) XXX_DiscardUnknown() {
	xxx_messageInfo_Claim.DiscardUnknown(m)
}

var xxx_messageInfo_Claim proto.InternalMessageInfo

func (m *Claim) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Claim) GetValidatorAddress() string {
	if m != nil {
		return m.ValidatorAddress
	}
	return ""
}

func (m *Claim) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

// DBProphecy is what the prophecy becomes when being saved to the database.
//  Tendermint/Amino does not support maps so we must serialize those variables
//  into bytes.
type DBProphecy struct {
	Id              string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Status          Status `protobuf:"bytes,2,opt,name=status,proto3" json:"status"`
	ClaimValidators []byte `protobuf:"bytes,3,opt,name=claim_validators,json=claimValidators,proto3" json:"claim_validators,omitempty"`
	ValidatorClaims []byte `protobuf:"bytes,4,opt,name=validator_claims,json=validatorClaims,proto3" json:"validator_claims,omitempty"`
}

func (m *DBProphecy) Reset()         { *m = DBProphecy{} }
func (m *DBProphecy) String() string { return proto.CompactTextString(m) }
func (*DBProphecy) ProtoMessage()    {}
func (*DBProphecy) Descriptor() ([]byte, []int) {
	return fileDescriptor_dac1b931484f4203, []int{2}
}
func (m *DBProphecy) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DBProphecy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DBProphecy.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DBProphecy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DBProphecy.Merge(m, src)
}
func (m *DBProphecy) XXX_Size() int {
	return m.Size()
}
func (m *DBProphecy) XXX_DiscardUnknown() {
	xxx_messageInfo_DBProphecy.DiscardUnknown(m)
}

var xxx_messageInfo_DBProphecy proto.InternalMessageInfo

func (m *DBProphecy) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *DBProphecy) GetStatus() Status {
	if m != nil {
		return m.Status
	}
	return Status{}
}

func (m *DBProphecy) GetClaimValidators() []byte {
	if m != nil {
		return m.ClaimValidators
	}
	return nil
}

func (m *DBProphecy) GetValidatorClaims() []byte {
	if m != nil {
		return m.ValidatorClaims
	}
	return nil
}

// Status is a struct that contains the status of a given prophecy
type Status struct {
	Text       StatusText `protobuf:"varint,1,opt,name=text,proto3,enum=sifnode.oracle.v1.StatusText" json:"text,omitempty"`
	FinalClaim string     `protobuf:"bytes,2,opt,name=final_claim,json=finalClaim,proto3" json:"final_claim,omitempty"`
}

func (m *Status) Reset()         { *m = Status{} }
func (m *Status) String() string { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()    {}
func (*Status) Descriptor() ([]byte, []int) {
	return fileDescriptor_dac1b931484f4203, []int{3}
}
func (m *Status) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Status) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Status.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Status) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Status.Merge(m, src)
}
func (m *Status) XXX_Size() int {
	return m.Size()
}
func (m *Status) XXX_DiscardUnknown() {
	xxx_messageInfo_Status.DiscardUnknown(m)
}

var xxx_messageInfo_Status proto.InternalMessageInfo

func (m *Status) GetText() StatusText {
	if m != nil {
		return m.Text
	}
	return StatusText_STATUS_TEXT_UNSPECIFIED
}

func (m *Status) GetFinalClaim() string {
	if m != nil {
		return m.FinalClaim
	}
	return ""
}

func init() {
	proto.RegisterEnum("sifnode.oracle.v1.StatusText", StatusText_name, StatusText_value)
	proto.RegisterType((*GenesisState)(nil), "sifnode.oracle.v1.GenesisState")
	proto.RegisterType((*Claim)(nil), "sifnode.oracle.v1.Claim")
	proto.RegisterType((*DBProphecy)(nil), "sifnode.oracle.v1.DBProphecy")
	proto.RegisterType((*Status)(nil), "sifnode.oracle.v1.Status")
}

func init() { proto.RegisterFile("sifnode/oracle/v1/types.proto", fileDescriptor_dac1b931484f4203) }

var fileDescriptor_dac1b931484f4203 = []byte{
	// 461 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0xc1, 0x6e, 0xd3, 0x40,
	0x14, 0x45, 0xed, 0x24, 0x04, 0xf5, 0x35, 0x14, 0x77, 0x40, 0xd4, 0x80, 0xea, 0x56, 0x61, 0x53,
	0x8a, 0x64, 0x2b, 0x65, 0xc1, 0x3a, 0x89, 0xdd, 0x2a, 0x12, 0x8a, 0x22, 0x8f, 0x03, 0x08, 0x21,
	0xcc, 0xd4, 0x9e, 0x24, 0x23, 0x39, 0x9e, 0xc8, 0x33, 0x0d, 0xe9, 0x5f, 0xf0, 0x1f, 0xfc, 0x48,
	0x97, 0x5d, 0xb2, 0x42, 0x28, 0xf9, 0x11, 0x94, 0xb1, 0xe3, 0x56, 0xad, 0xba, 0xb3, 0xef, 0xb9,
	0x7e, 0xf7, 0xf9, 0xe9, 0xc2, 0xbe, 0x60, 0xa3, 0x94, 0xc7, 0xd4, 0xe1, 0x19, 0x89, 0x12, 0xea,
	0xcc, 0x5b, 0x8e, 0xbc, 0x9c, 0x51, 0x61, 0xcf, 0x32, 0x2e, 0x39, 0xda, 0x2d, 0xb0, 0x9d, 0x63,
	0x7b, 0xde, 0x7a, 0xf5, 0x7c, 0xcc, 0xc7, 0x5c, 0x51, 0x67, 0xfd, 0x94, 0x1b, 0x9b, 0x3f, 0xa0,
	0x71, 0x46, 0x53, 0x2a, 0x98, 0xc0, 0x92, 0x48, 0x8a, 0xde, 0xc1, 0x2e, 0x89, 0xe3, 0x8c, 0x0a,
	0x11, 0xfe, 0x9c, 0x30, 0x49, 0x13, 0x26, 0xa4, 0xa9, 0x1f, 0x56, 0x8f, 0xb6, 0x7c, 0xa3, 0x00,
	0x9f, 0x37, 0x3a, 0x7a, 0x03, 0x4f, 0x48, 0x3c, 0x65, 0x69, 0x58, 0x10, 0xb3, 0x72, 0xa8, 0x1f,
	0x6d, 0xf9, 0x0d, 0x25, 0xb6, 0x73, 0xad, 0xf9, 0x1d, 0x1e, 0x75, 0x13, 0xc2, 0xa6, 0x68, 0x07,
	0x2a, 0x2c, 0x36, 0x75, 0x65, 0xa9, 0xb0, 0x78, 0x1d, 0x35, 0x27, 0x09, 0x8b, 0x89, 0xe4, 0xd9,
	0x9d, 0x09, 0x46, 0x09, 0x8a, 0x29, 0xc8, 0x84, 0xc7, 0x11, 0x4f, 0x25, 0x4d, 0xa5, 0x59, 0x55,
	0x96, 0xcd, 0x6b, 0xf3, 0xb7, 0x0e, 0xe0, 0x76, 0x06, 0x19, 0x9f, 0x4d, 0x68, 0x74, 0x79, 0x2f,
	0xe5, 0x03, 0xd4, 0x85, 0x24, 0xf2, 0x22, 0x1f, 0xbd, 0x7d, 0xf2, 0xd2, 0xbe, 0x77, 0x1a, 0x1b,
	0x2b, 0x43, 0xa7, 0x76, 0xf5, 0xf7, 0x40, 0xf3, 0x0b, 0x3b, 0x7a, 0x0b, 0x46, 0xb4, 0xde, 0x3b,
	0x2c, 0x77, 0x11, 0x2a, 0xba, 0xe1, 0x3f, 0x55, 0xfa, 0xa7, 0x52, 0x5e, 0x5b, 0x6f, 0xfe, 0x44,
	0x41, 0x61, 0xd6, 0x72, 0x6b, 0xa9, 0xab, 0x1b, 0x88, 0xe6, 0x37, 0xa8, 0xe7, 0x69, 0xa8, 0x05,
	0x35, 0x49, 0x17, 0x52, 0xad, 0xba, 0x73, 0xb2, 0xff, 0xe0, 0x5a, 0x01, 0x5d, 0x48, 0x5f, 0x59,
	0xd1, 0x01, 0x6c, 0x8f, 0x58, 0x4a, 0x92, 0x3c, 0xa3, 0xb8, 0x15, 0x28, 0x49, 0x8d, 0x3f, 0x16,
	0x00, 0x37, 0x1f, 0xa1, 0xd7, 0xb0, 0x87, 0x83, 0x76, 0x30, 0xc4, 0x61, 0xe0, 0x7d, 0x09, 0xc2,
	0x61, 0x1f, 0x0f, 0xbc, 0x6e, 0xef, 0xb4, 0xe7, 0xb9, 0x86, 0x86, 0xf6, 0xe0, 0xd9, 0x6d, 0x38,
	0xf0, 0xfa, 0x6e, 0xaf, 0x7f, 0x66, 0xe8, 0x77, 0x01, 0x1e, 0x76, 0xbb, 0x1e, 0xc6, 0x46, 0x05,
	0xbd, 0x00, 0x74, 0x1b, 0x9c, 0xb6, 0x7b, 0x1f, 0x3d, 0xd7, 0xa8, 0x76, 0xdc, 0xab, 0xa5, 0xa5,
	0x5f, 0x2f, 0x2d, 0xfd, 0xdf, 0xd2, 0xd2, 0x7f, 0xad, 0x2c, 0xed, 0x7a, 0x65, 0x69, 0x7f, 0x56,
	0x96, 0xf6, 0xf5, 0x78, 0xcc, 0xe4, 0xe4, 0xe2, 0xdc, 0x8e, 0xf8, 0xd4, 0xc1, 0x6c, 0x14, 0x4d,
	0x08, 0x4b, 0x9d, 0x4d, 0x71, 0x17, 0x9b, 0xea, 0xaa, 0xde, 0x9e, 0xd7, 0x55, 0x1f, 0xdf, 0xff,
	0x0f, 0x00, 0x00, 0xff, 0xff, 0x66, 0xb2, 0x98, 0xa6, 0xd9, 0x02, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.AdminAddress) > 0 {
		i -= len(m.AdminAddress)
		copy(dAtA[i:], m.AdminAddress)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.AdminAddress)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.AddressWhitelist) > 0 {
		for iNdEx := len(m.AddressWhitelist) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.AddressWhitelist[iNdEx])
			copy(dAtA[i:], m.AddressWhitelist[iNdEx])
			i = encodeVarintTypes(dAtA, i, uint64(len(m.AddressWhitelist[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *Claim) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Claim) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Claim) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Content) > 0 {
		i -= len(m.Content)
		copy(dAtA[i:], m.Content)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Content)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.ValidatorAddress) > 0 {
		i -= len(m.ValidatorAddress)
		copy(dAtA[i:], m.ValidatorAddress)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.ValidatorAddress)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *DBProphecy) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DBProphecy) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DBProphecy) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ValidatorClaims) > 0 {
		i -= len(m.ValidatorClaims)
		copy(dAtA[i:], m.ValidatorClaims)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.ValidatorClaims)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.ClaimValidators) > 0 {
		i -= len(m.ClaimValidators)
		copy(dAtA[i:], m.ClaimValidators)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.ClaimValidators)))
		i--
		dAtA[i] = 0x1a
	}
	{
		size, err := m.Status.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Status) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Status) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Status) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.FinalClaim) > 0 {
		i -= len(m.FinalClaim)
		copy(dAtA[i:], m.FinalClaim)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.FinalClaim)))
		i--
		dAtA[i] = 0x12
	}
	if m.Text != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Text))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintTypes(dAtA []byte, offset int, v uint64) int {
	offset -= sovTypes(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.AddressWhitelist) > 0 {
		for _, s := range m.AddressWhitelist {
			l = len(s)
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	l = len(m.AdminAddress)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func (m *Claim) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.ValidatorAddress)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.Content)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func (m *DBProphecy) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = m.Status.Size()
	n += 1 + l + sovTypes(uint64(l))
	l = len(m.ClaimValidators)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.ValidatorClaims)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func (m *Status) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Text != 0 {
		n += 1 + sovTypes(uint64(m.Text))
	}
	l = len(m.FinalClaim)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func sovTypes(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypes(x uint64) (n int) {
	return sovTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AddressWhitelist", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AddressWhitelist = append(m.AddressWhitelist, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AdminAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AdminAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *Claim) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: Claim: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Claim: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidatorAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ValidatorAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Content", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Content = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *DBProphecy) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: DBProphecy: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DBProphecy: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Status.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClaimValidators", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClaimValidators = append(m.ClaimValidators[:0], dAtA[iNdEx:postIndex]...)
			if m.ClaimValidators == nil {
				m.ClaimValidators = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidatorClaims", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ValidatorClaims = append(m.ValidatorClaims[:0], dAtA[iNdEx:postIndex]...)
			if m.ValidatorClaims == nil {
				m.ValidatorClaims = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *Status) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: Status: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Status: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Text", wireType)
			}
			m.Text = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Text |= StatusText(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FinalClaim", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FinalClaim = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
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
func skipTypes(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
				return 0, ErrInvalidLengthTypes
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTypes
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTypes
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTypes        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypes          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTypes = fmt.Errorf("proto: unexpected end of group")
)
