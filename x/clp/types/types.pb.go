// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sifnode/clp/v1/types.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
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

type Asset struct {
	Symbol string `protobuf:"bytes,1,opt,name=symbol,proto3" json:"symbol,omitempty"`
}

func (m *Asset) Reset()         { *m = Asset{} }
func (m *Asset) String() string { return proto.CompactTextString(m) }
func (*Asset) ProtoMessage()    {}
func (*Asset) Descriptor() ([]byte, []int) {
	return fileDescriptor_a09f92a67752e669, []int{0}
}
func (m *Asset) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Asset) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Asset.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Asset) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Asset.Merge(m, src)
}
func (m *Asset) XXX_Size() int {
	return m.Size()
}
func (m *Asset) XXX_DiscardUnknown() {
	xxx_messageInfo_Asset.DiscardUnknown(m)
}

var xxx_messageInfo_Asset proto.InternalMessageInfo

func (m *Asset) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

type Pool struct {
	ExternalAsset        *Asset                                  `protobuf:"bytes,1,opt,name=external_asset,json=externalAsset,proto3" json:"external_asset,omitempty"`
	NativeAssetBalance   github_com_cosmos_cosmos_sdk_types.Uint `protobuf:"bytes,2,opt,name=native_asset_balance,json=nativeAssetBalance,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Uint" json:"native_asset_balance" yaml:"native_asset_balance"`
	ExternalAssetBalance github_com_cosmos_cosmos_sdk_types.Uint `protobuf:"bytes,3,opt,name=external_asset_balance,json=externalAssetBalance,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Uint" json:"external_asset_balance" yaml:"external_asset_balance"`
	PoolUnits            github_com_cosmos_cosmos_sdk_types.Uint `protobuf:"bytes,4,opt,name=pool_units,json=poolUnits,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Uint" json:"pool_units" yaml:"pool_units"`
}

func (m *Pool) Reset()         { *m = Pool{} }
func (m *Pool) String() string { return proto.CompactTextString(m) }
func (*Pool) ProtoMessage()    {}
func (*Pool) Descriptor() ([]byte, []int) {
	return fileDescriptor_a09f92a67752e669, []int{1}
}
func (m *Pool) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Pool) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Pool.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Pool) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pool.Merge(m, src)
}
func (m *Pool) XXX_Size() int {
	return m.Size()
}
func (m *Pool) XXX_DiscardUnknown() {
	xxx_messageInfo_Pool.DiscardUnknown(m)
}

var xxx_messageInfo_Pool proto.InternalMessageInfo

func (m *Pool) GetExternalAsset() *Asset {
	if m != nil {
		return m.ExternalAsset
	}
	return nil
}

type LiquidityProvider struct {
	Asset                    *Asset                                  `protobuf:"bytes,1,opt,name=asset,proto3" json:"asset,omitempty"`
	LiquidityProviderUnits   github_com_cosmos_cosmos_sdk_types.Uint `protobuf:"bytes,2,opt,name=liquidity_provider_units,json=liquidityProviderUnits,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Uint" json:"liquidity_provider_units" yaml:"liquidity_provider_units"`
	LiquidityProviderAddress string                                  `protobuf:"bytes,3,opt,name=liquidity_provider_address,json=liquidityProviderAddress,proto3" json:"liquidity_provider_address,omitempty"`
}

func (m *LiquidityProvider) Reset()         { *m = LiquidityProvider{} }
func (m *LiquidityProvider) String() string { return proto.CompactTextString(m) }
func (*LiquidityProvider) ProtoMessage()    {}
func (*LiquidityProvider) Descriptor() ([]byte, []int) {
	return fileDescriptor_a09f92a67752e669, []int{2}
}
func (m *LiquidityProvider) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LiquidityProvider) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LiquidityProvider.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LiquidityProvider) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LiquidityProvider.Merge(m, src)
}
func (m *LiquidityProvider) XXX_Size() int {
	return m.Size()
}
func (m *LiquidityProvider) XXX_DiscardUnknown() {
	xxx_messageInfo_LiquidityProvider.DiscardUnknown(m)
}

var xxx_messageInfo_LiquidityProvider proto.InternalMessageInfo

func (m *LiquidityProvider) GetAsset() *Asset {
	if m != nil {
		return m.Asset
	}
	return nil
}

func (m *LiquidityProvider) GetLiquidityProviderAddress() string {
	if m != nil {
		return m.LiquidityProviderAddress
	}
	return ""
}

type WhiteList struct {
	ValidatorList []string `protobuf:"bytes,1,rep,name=validator_list,json=validatorList,proto3" json:"validator_list,omitempty"`
}

func (m *WhiteList) Reset()         { *m = WhiteList{} }
func (m *WhiteList) String() string { return proto.CompactTextString(m) }
func (*WhiteList) ProtoMessage()    {}
func (*WhiteList) Descriptor() ([]byte, []int) {
	return fileDescriptor_a09f92a67752e669, []int{3}
}
func (m *WhiteList) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *WhiteList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_WhiteList.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *WhiteList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WhiteList.Merge(m, src)
}
func (m *WhiteList) XXX_Size() int {
	return m.Size()
}
func (m *WhiteList) XXX_DiscardUnknown() {
	xxx_messageInfo_WhiteList.DiscardUnknown(m)
}

var xxx_messageInfo_WhiteList proto.InternalMessageInfo

func (m *WhiteList) GetValidatorList() []string {
	if m != nil {
		return m.ValidatorList
	}
	return nil
}

type LiquidityProviderData struct {
	LiquidityProvider    *LiquidityProvider `protobuf:"bytes,1,opt,name=liquidity_provider,json=liquidityProvider,proto3" json:"liquidity_provider,omitempty"`
	Symbol               string             `protobuf:"bytes,2,opt,name=symbol,proto3" json:"symbol,omitempty"`
	NativeAssetBalance   string             `protobuf:"bytes,3,opt,name=native_asset_balance,json=nativeAssetBalance,proto3" json:"native_asset_balance,omitempty"`
	ExternalAssetBalance string             `protobuf:"bytes,4,opt,name=external_asset_balance,json=externalAssetBalance,proto3" json:"external_asset_balance,omitempty"`
}

func (m *LiquidityProviderData) Reset()         { *m = LiquidityProviderData{} }
func (m *LiquidityProviderData) String() string { return proto.CompactTextString(m) }
func (*LiquidityProviderData) ProtoMessage()    {}
func (*LiquidityProviderData) Descriptor() ([]byte, []int) {
	return fileDescriptor_a09f92a67752e669, []int{4}
}
func (m *LiquidityProviderData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LiquidityProviderData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LiquidityProviderData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LiquidityProviderData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LiquidityProviderData.Merge(m, src)
}
func (m *LiquidityProviderData) XXX_Size() int {
	return m.Size()
}
func (m *LiquidityProviderData) XXX_DiscardUnknown() {
	xxx_messageInfo_LiquidityProviderData.DiscardUnknown(m)
}

var xxx_messageInfo_LiquidityProviderData proto.InternalMessageInfo

func (m *LiquidityProviderData) GetLiquidityProvider() *LiquidityProvider {
	if m != nil {
		return m.LiquidityProvider
	}
	return nil
}

func (m *LiquidityProviderData) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *LiquidityProviderData) GetNativeAssetBalance() string {
	if m != nil {
		return m.NativeAssetBalance
	}
	return ""
}

func (m *LiquidityProviderData) GetExternalAssetBalance() string {
	if m != nil {
		return m.ExternalAssetBalance
	}
	return ""
}

func init() {
	proto.RegisterType((*Asset)(nil), "sifnode.clp.v1.Asset")
	proto.RegisterType((*Pool)(nil), "sifnode.clp.v1.Pool")
	proto.RegisterType((*LiquidityProvider)(nil), "sifnode.clp.v1.LiquidityProvider")
	proto.RegisterType((*WhiteList)(nil), "sifnode.clp.v1.WhiteList")
	proto.RegisterType((*LiquidityProviderData)(nil), "sifnode.clp.v1.LiquidityProviderData")
}

func init() { proto.RegisterFile("sifnode/clp/v1/types.proto", fileDescriptor_a09f92a67752e669) }

var fileDescriptor_a09f92a67752e669 = []byte{
	// 516 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0x4d, 0x6b, 0xdb, 0x30,
	0x18, 0x8e, 0x93, 0xb6, 0x10, 0x8d, 0x06, 0x22, 0xd2, 0x60, 0x32, 0x66, 0x6f, 0x86, 0xb1, 0xc2,
	0x98, 0xbd, 0x76, 0x3b, 0x8d, 0x5e, 0x12, 0x7a, 0x2c, 0x23, 0x68, 0x94, 0xc1, 0x2e, 0x46, 0xb1,
	0xd5, 0x44, 0x4c, 0xb1, 0x3c, 0x4b, 0x31, 0xcd, 0x6d, 0xa7, 0x9d, 0x36, 0xd8, 0x9f, 0xd9, 0x7f,
	0xe8, 0xb1, 0xc7, 0xb1, 0x43, 0x18, 0xc9, 0x69, 0xd7, 0xfd, 0x82, 0x61, 0xc9, 0xe9, 0x92, 0xd8,
	0x81, 0xe6, 0x64, 0x4b, 0x7a, 0xdf, 0xe7, 0x79, 0xde, 0x4f, 0xd0, 0x11, 0xf4, 0x2a, 0xe2, 0x21,
	0xf1, 0x02, 0x16, 0x7b, 0xe9, 0x89, 0x27, 0xa7, 0x31, 0x11, 0x6e, 0x9c, 0x70, 0xc9, 0x61, 0x23,
	0x7f, 0x73, 0x03, 0x16, 0xbb, 0xe9, 0x49, 0xa7, 0x35, 0xe4, 0x43, 0xae, 0x9e, 0xbc, 0xec, 0x4f,
	0x5b, 0x39, 0x36, 0xd8, 0xef, 0x0a, 0x41, 0x24, 0x6c, 0x83, 0x03, 0x31, 0x1d, 0x0f, 0x38, 0x33,
	0x8d, 0xc7, 0xc6, 0x71, 0x1d, 0xe5, 0x27, 0xe7, 0x47, 0x0d, 0xec, 0xf5, 0x39, 0x67, 0xf0, 0x0c,
	0x34, 0xc8, 0xb5, 0x24, 0x49, 0x84, 0x99, 0x8f, 0x33, 0x17, 0x65, 0xf8, 0xe0, 0xf4, 0xc8, 0x5d,
	0x27, 0x72, 0x15, 0x1e, 0x3a, 0x5c, 0x1a, 0x6b, 0xf8, 0xcf, 0x06, 0x68, 0x45, 0x58, 0xd2, 0x94,
	0x68, 0x67, 0x7f, 0x80, 0x19, 0x8e, 0x02, 0x62, 0x56, 0x33, 0xb6, 0xde, 0xdb, 0x9b, 0x99, 0x5d,
	0xf9, 0x35, 0xb3, 0x9f, 0x0d, 0xa9, 0x1c, 0x4d, 0x06, 0x6e, 0xc0, 0xc7, 0x5e, 0xc0, 0xc5, 0x98,
	0x8b, 0xfc, 0xf3, 0x42, 0x84, 0x1f, 0xf3, 0xf0, 0x2e, 0x69, 0x24, 0xff, 0xce, 0xec, 0x87, 0x53,
	0x3c, 0x66, 0x6f, 0x9c, 0x32, 0x50, 0x07, 0x41, 0x7d, 0xad, 0xb8, 0x7b, 0xfa, 0x12, 0x7e, 0x31,
	0x40, 0x7b, 0x3d, 0x82, 0x3b, 0x11, 0x35, 0x25, 0xa2, 0xbf, 0xbb, 0x88, 0x47, 0x5a, 0x44, 0x39,
	0xac, 0x83, 0x5a, 0x6b, 0x49, 0x58, 0x0a, 0x09, 0x00, 0x88, 0x39, 0x67, 0xfe, 0x24, 0xa2, 0x52,
	0x98, 0x7b, 0x8a, 0xfb, 0x7c, 0x77, 0xee, 0xa6, 0xe6, 0xfe, 0x0f, 0xe5, 0xa0, 0x7a, 0x76, 0xb8,
	0x54, 0xff, 0xdf, 0xaa, 0xa0, 0x79, 0x41, 0x3f, 0x4d, 0x68, 0x48, 0xe5, 0xb4, 0x9f, 0xf0, 0x94,
	0x86, 0x24, 0x81, 0xcf, 0xc1, 0xfe, 0x3d, 0x6a, 0xa7, 0x6d, 0xe0, 0x57, 0x03, 0x98, 0x6c, 0x09,
	0xe1, 0xc7, 0x39, 0x46, 0x2e, 0x5b, 0xd7, 0x0d, 0xed, 0x2e, 0xdb, 0xd6, 0xb2, 0xb7, 0x01, 0x3b,
	0xa8, 0xcd, 0x36, 0x65, 0xab, 0x88, 0xe0, 0x19, 0xe8, 0x94, 0x38, 0xe1, 0x30, 0x4c, 0x88, 0x10,
	0xba, 0x84, 0xc8, 0x2c, 0xf8, 0x76, 0xf5, 0xbb, 0x73, 0x0a, 0xea, 0xef, 0x47, 0x54, 0x92, 0x0b,
	0x2a, 0x24, 0x7c, 0x0a, 0x1a, 0x29, 0x66, 0x34, 0xc4, 0x92, 0x27, 0x3e, 0xa3, 0x22, 0xcb, 0x47,
	0xed, 0xb8, 0x8e, 0x0e, 0xef, 0x6e, 0x33, 0x33, 0xe7, 0x8f, 0x01, 0x8e, 0x0a, 0x39, 0x3c, 0xc7,
	0x12, 0xc3, 0x3e, 0x80, 0x45, 0x2d, 0x79, 0x52, 0x9f, 0x6c, 0x26, 0xb5, 0x00, 0x81, 0x9a, 0x05,
	0x99, 0x2b, 0xf3, 0x57, 0x5d, 0x9d, 0x3f, 0xf8, 0x72, 0xcb, 0xdc, 0xe8, 0x78, 0xcb, 0xfa, 0xfc,
	0xf5, 0xd6, 0x36, 0x57, 0xad, 0x56, 0xde, 0x94, 0xbd, 0xee, 0xcd, 0xdc, 0x32, 0x6e, 0xe7, 0x96,
	0xf1, 0x7b, 0x6e, 0x19, 0xdf, 0x17, 0x56, 0xe5, 0x76, 0x61, 0x55, 0x7e, 0x2e, 0xac, 0xca, 0x87,
	0xd5, 0xda, 0xbe, 0xa3, 0x57, 0xc1, 0x08, 0xd3, 0xc8, 0x5b, 0x2e, 0x9e, 0x6b, 0xb5, 0x7a, 0x54,
	0x81, 0x07, 0x07, 0x6a, 0xa5, 0xbc, 0xfa, 0x17, 0x00, 0x00, 0xff, 0xff, 0xeb, 0xfb, 0xd1, 0xeb,
	0x96, 0x04, 0x00, 0x00,
}

func (m *Asset) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Asset) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Asset) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Symbol) > 0 {
		i -= len(m.Symbol)
		copy(dAtA[i:], m.Symbol)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Symbol)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Pool) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Pool) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Pool) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.PoolUnits.Size()
		i -= size
		if _, err := m.PoolUnits.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size := m.ExternalAssetBalance.Size()
		i -= size
		if _, err := m.ExternalAssetBalance.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.NativeAssetBalance.Size()
		i -= size
		if _, err := m.NativeAssetBalance.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.ExternalAsset != nil {
		{
			size, err := m.ExternalAsset.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *LiquidityProvider) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LiquidityProvider) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LiquidityProvider) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.LiquidityProviderAddress) > 0 {
		i -= len(m.LiquidityProviderAddress)
		copy(dAtA[i:], m.LiquidityProviderAddress)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.LiquidityProviderAddress)))
		i--
		dAtA[i] = 0x1a
	}
	{
		size := m.LiquidityProviderUnits.Size()
		i -= size
		if _, err := m.LiquidityProviderUnits.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.Asset != nil {
		{
			size, err := m.Asset.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *WhiteList) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *WhiteList) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *WhiteList) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ValidatorList) > 0 {
		for iNdEx := len(m.ValidatorList) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.ValidatorList[iNdEx])
			copy(dAtA[i:], m.ValidatorList[iNdEx])
			i = encodeVarintTypes(dAtA, i, uint64(len(m.ValidatorList[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *LiquidityProviderData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LiquidityProviderData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LiquidityProviderData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ExternalAssetBalance) > 0 {
		i -= len(m.ExternalAssetBalance)
		copy(dAtA[i:], m.ExternalAssetBalance)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.ExternalAssetBalance)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.NativeAssetBalance) > 0 {
		i -= len(m.NativeAssetBalance)
		copy(dAtA[i:], m.NativeAssetBalance)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.NativeAssetBalance)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Symbol) > 0 {
		i -= len(m.Symbol)
		copy(dAtA[i:], m.Symbol)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Symbol)))
		i--
		dAtA[i] = 0x12
	}
	if m.LiquidityProvider != nil {
		{
			size, err := m.LiquidityProvider.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
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
func (m *Asset) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Symbol)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func (m *Pool) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ExternalAsset != nil {
		l = m.ExternalAsset.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	l = m.NativeAssetBalance.Size()
	n += 1 + l + sovTypes(uint64(l))
	l = m.ExternalAssetBalance.Size()
	n += 1 + l + sovTypes(uint64(l))
	l = m.PoolUnits.Size()
	n += 1 + l + sovTypes(uint64(l))
	return n
}

func (m *LiquidityProvider) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Asset != nil {
		l = m.Asset.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	l = m.LiquidityProviderUnits.Size()
	n += 1 + l + sovTypes(uint64(l))
	l = len(m.LiquidityProviderAddress)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func (m *WhiteList) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.ValidatorList) > 0 {
		for _, s := range m.ValidatorList {
			l = len(s)
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	return n
}

func (m *LiquidityProviderData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.LiquidityProvider != nil {
		l = m.LiquidityProvider.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.Symbol)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.NativeAssetBalance)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.ExternalAssetBalance)
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
func (m *Asset) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: Asset: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Asset: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Symbol", wireType)
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
			m.Symbol = string(dAtA[iNdEx:postIndex])
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
func (m *Pool) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: Pool: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Pool: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExternalAsset", wireType)
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
			if m.ExternalAsset == nil {
				m.ExternalAsset = &Asset{}
			}
			if err := m.ExternalAsset.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NativeAssetBalance", wireType)
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
			if err := m.NativeAssetBalance.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExternalAssetBalance", wireType)
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
			if err := m.ExternalAssetBalance.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolUnits", wireType)
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
			if err := m.PoolUnits.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
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
func (m *LiquidityProvider) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: LiquidityProvider: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LiquidityProvider: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Asset", wireType)
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
			if m.Asset == nil {
				m.Asset = &Asset{}
			}
			if err := m.Asset.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LiquidityProviderUnits", wireType)
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
			if err := m.LiquidityProviderUnits.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LiquidityProviderAddress", wireType)
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
			m.LiquidityProviderAddress = string(dAtA[iNdEx:postIndex])
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
func (m *WhiteList) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: WhiteList: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: WhiteList: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidatorList", wireType)
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
			m.ValidatorList = append(m.ValidatorList, string(dAtA[iNdEx:postIndex]))
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
func (m *LiquidityProviderData) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: LiquidityProviderData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LiquidityProviderData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LiquidityProvider", wireType)
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
			if m.LiquidityProvider == nil {
				m.LiquidityProvider = &LiquidityProvider{}
			}
			if err := m.LiquidityProvider.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Symbol", wireType)
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
			m.Symbol = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NativeAssetBalance", wireType)
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
			m.NativeAssetBalance = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExternalAssetBalance", wireType)
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
			m.ExternalAssetBalance = string(dAtA[iNdEx:postIndex])
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
