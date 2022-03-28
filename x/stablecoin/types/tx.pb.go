// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: stablecoin/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// Msg to mint USDM: User deposits MTRX and collateral and gets USDM in return.
type MsgMintStable struct {
	Creator string     `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Stable  types.Coin `protobuf:"bytes,2,opt,name=stable,proto3" json:"stable"`
}

func (m *MsgMintStable) Reset()         { *m = MsgMintStable{} }
func (m *MsgMintStable) String() string { return proto.CompactTextString(m) }
func (*MsgMintStable) ProtoMessage()    {}
func (*MsgMintStable) Descriptor() ([]byte, []int) {
	return fileDescriptor_6ef74a087750083d, []int{0}
}
func (m *MsgMintStable) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgMintStable) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgMintStable.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgMintStable) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgMintStable.Merge(m, src)
}
func (m *MsgMintStable) XXX_Size() int {
	return m.Size()
}
func (m *MsgMintStable) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgMintStable.DiscardUnknown(m)
}

var xxx_messageInfo_MsgMintStable proto.InternalMessageInfo

func (m *MsgMintStable) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgMintStable) GetStable() types.Coin {
	if m != nil {
		return m.Stable
	}
	return types.Coin{}
}

type MsgMintStableResponse struct {
	Stable types.Coin `protobuf:"bytes,1,opt,name=stable,proto3" json:"stable"`
}

func (m *MsgMintStableResponse) Reset()         { *m = MsgMintStableResponse{} }
func (m *MsgMintStableResponse) String() string { return proto.CompactTextString(m) }
func (*MsgMintStableResponse) ProtoMessage()    {}
func (*MsgMintStableResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6ef74a087750083d, []int{1}
}
func (m *MsgMintStableResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgMintStableResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgMintStableResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgMintStableResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgMintStableResponse.Merge(m, src)
}
func (m *MsgMintStableResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgMintStableResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgMintStableResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgMintStableResponse proto.InternalMessageInfo

func (m *MsgMintStableResponse) GetStable() types.Coin {
	if m != nil {
		return m.Stable
	}
	return types.Coin{}
}

// Msg to burn USDM: User burns USDM and gets MTRX and collateral in return.
type MsgBurn struct {
	Creator string     `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Stable  types.Coin `protobuf:"bytes,2,opt,name=stable,proto3" json:"stable"`
}

func (m *MsgBurn) Reset()         { *m = MsgBurn{} }
func (m *MsgBurn) String() string { return proto.CompactTextString(m) }
func (*MsgBurn) ProtoMessage()    {}
func (*MsgBurn) Descriptor() ([]byte, []int) {
	return fileDescriptor_6ef74a087750083d, []int{2}
}
func (m *MsgBurn) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgBurn) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgBurn.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgBurn) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgBurn.Merge(m, src)
}
func (m *MsgBurn) XXX_Size() int {
	return m.Size()
}
func (m *MsgBurn) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgBurn.DiscardUnknown(m)
}

var xxx_messageInfo_MsgBurn proto.InternalMessageInfo

func (m *MsgBurn) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgBurn) GetStable() types.Coin {
	if m != nil {
		return m.Stable
	}
	return types.Coin{}
}

type MsgBurnResponse struct {
	Collateral types.Coin `protobuf:"bytes,1,opt,name=collateral,proto3" json:"collateral"`
	Gov        types.Coin `protobuf:"bytes,2,opt,name=gov,proto3" json:"gov"`
}

func (m *MsgBurnResponse) Reset()         { *m = MsgBurnResponse{} }
func (m *MsgBurnResponse) String() string { return proto.CompactTextString(m) }
func (*MsgBurnResponse) ProtoMessage()    {}
func (*MsgBurnResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6ef74a087750083d, []int{3}
}
func (m *MsgBurnResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgBurnResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgBurnResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgBurnResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgBurnResponse.Merge(m, src)
}
func (m *MsgBurnResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgBurnResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgBurnResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgBurnResponse proto.InternalMessageInfo

func (m *MsgBurnResponse) GetCollateral() types.Coin {
	if m != nil {
		return m.Collateral
	}
	return types.Coin{}
}

func (m *MsgBurnResponse) GetGov() types.Coin {
	if m != nil {
		return m.Gov
	}
	return types.Coin{}
}

func init() {
	proto.RegisterType((*MsgMintStable)(nil), "MatrixDao.stablecoin.v1.MsgMintStable")
	proto.RegisterType((*MsgMintStableResponse)(nil), "MatrixDao.stablecoin.v1.MsgMintStableResponse")
	proto.RegisterType((*MsgBurn)(nil), "MatrixDao.stablecoin.v1.MsgBurn")
	proto.RegisterType((*MsgBurnResponse)(nil), "MatrixDao.stablecoin.v1.MsgBurnResponse")
}

func init() { proto.RegisterFile("stablecoin/tx.proto", fileDescriptor_6ef74a087750083d) }

var fileDescriptor_6ef74a087750083d = []byte{
	// 330 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x52, 0x3f, 0x4f, 0xfb, 0x30,
	0x14, 0x8c, 0x7f, 0xfd, 0xa9, 0x15, 0x0f, 0x21, 0xa4, 0x00, 0xa2, 0x74, 0x30, 0x55, 0x07, 0xd4,
	0xe9, 0x59, 0x29, 0x03, 0x23, 0x52, 0x61, 0x60, 0x89, 0x84, 0xca, 0x86, 0x58, 0x9c, 0x60, 0x99,
	0x88, 0x34, 0xaf, 0x8a, 0xdd, 0xaa, 0x7c, 0x00, 0x76, 0x3e, 0x56, 0xc7, 0x8e, 0x4c, 0x08, 0xb5,
	0x5f, 0x04, 0xd5, 0xfd, 0x43, 0x3a, 0x20, 0x85, 0x81, 0xed, 0x92, 0x77, 0xef, 0xee, 0x6c, 0x1f,
	0x1c, 0x18, 0x2b, 0xa3, 0x54, 0xc5, 0x94, 0x64, 0xc2, 0x8e, 0x71, 0x90, 0x93, 0x25, 0xff, 0x38,
	0x94, 0x36, 0x4f, 0xc6, 0xd7, 0x92, 0xf0, 0x7b, 0x8c, 0xa3, 0xa0, 0xc1, 0x63, 0x32, 0x7d, 0x32,
	0x22, 0x92, 0x46, 0x89, 0x51, 0x10, 0x29, 0x2b, 0x03, 0xe1, 0x86, 0x6e, 0xb1, 0x71, 0xa8, 0x49,
	0x93, 0x83, 0x62, 0x81, 0x96, 0x7f, 0x5b, 0x11, 0xec, 0x85, 0x46, 0x87, 0x49, 0x66, 0xef, 0x9c,
	0x9a, 0x5f, 0x87, 0x5a, 0x9c, 0x2b, 0x69, 0x29, 0xaf, 0xb3, 0x26, 0x6b, 0xef, 0xf4, 0xd6, 0x9f,
	0xfe, 0x05, 0x54, 0x97, 0x8e, 0xf5, 0x7f, 0x4d, 0xd6, 0xde, 0xed, 0x9c, 0xe0, 0xd2, 0x11, 0x17,
	0x8e, 0xb8, 0x72, 0xc4, 0x2b, 0x4a, 0xb2, 0xee, 0xff, 0xc9, 0xc7, 0xa9, 0xd7, 0x5b, 0xd1, 0x5b,
	0xb7, 0x70, 0xb4, 0xe5, 0xd1, 0x53, 0x66, 0x40, 0x99, 0x51, 0x05, 0x45, 0xf6, 0x3b, 0xc5, 0x07,
	0xa8, 0x85, 0x46, 0x77, 0x87, 0x79, 0xf6, 0x17, 0x79, 0x5f, 0x19, 0xec, 0xaf, 0xe4, 0x37, 0x51,
	0x2f, 0x01, 0x62, 0x4a, 0x53, 0x69, 0x55, 0x2e, 0xd3, 0xb2, 0x71, 0x0b, 0x2b, 0x7e, 0x00, 0x15,
	0x4d, 0xa3, 0xb2, 0x51, 0x16, 0xdc, 0xce, 0x33, 0x54, 0x42, 0xa3, 0xfd, 0x47, 0x80, 0xc2, 0xfb,
	0x9c, 0xe1, 0x0f, 0x05, 0xc0, 0xad, 0x3b, 0x6e, 0x60, 0x39, 0xde, 0xfa, 0x80, 0xdd, 0x9b, 0xc9,
	0x8c, 0xb3, 0xe9, 0x8c, 0xb3, 0xcf, 0x19, 0x67, 0x6f, 0x73, 0xee, 0x4d, 0xe7, 0xdc, 0x7b, 0x9f,
	0x73, 0xef, 0x1e, 0x75, 0x62, 0x9f, 0x86, 0x11, 0xc6, 0xd4, 0x17, 0x1b, 0x4d, 0xd1, 0x77, 0x48,
	0x8c, 0x45, 0xb1, 0xa4, 0x2f, 0x03, 0x65, 0xa2, 0xaa, 0x6b, 0xd6, 0xf9, 0x57, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xb6, 0xf5, 0xa3, 0xbe, 0xbf, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	MintStable(ctx context.Context, in *MsgMintStable, opts ...grpc.CallOption) (*MsgMintStableResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) MintStable(ctx context.Context, in *MsgMintStable, opts ...grpc.CallOption) (*MsgMintStableResponse, error) {
	out := new(MsgMintStableResponse)
	err := c.cc.Invoke(ctx, "/MatrixDao.stablecoin.v1.Msg/MintStable", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	MintStable(context.Context, *MsgMintStable) (*MsgMintStableResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) MintStable(ctx context.Context, req *MsgMintStable) (*MsgMintStableResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MintStable not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_MintStable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgMintStable)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).MintStable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/MatrixDao.stablecoin.v1.Msg/MintStable",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).MintStable(ctx, req.(*MsgMintStable))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "MatrixDao.stablecoin.v1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MintStable",
			Handler:    _Msg_MintStable_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "stablecoin/tx.proto",
}

func (m *MsgMintStable) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgMintStable) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgMintStable) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Stable.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgMintStableResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgMintStableResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgMintStableResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Stable.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *MsgBurn) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgBurn) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgBurn) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Stable.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgBurnResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgBurnResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgBurnResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Gov.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size, err := m.Collateral.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgMintStable) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = m.Stable.Size()
	n += 1 + l + sovTx(uint64(l))
	return n
}

func (m *MsgMintStableResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Stable.Size()
	n += 1 + l + sovTx(uint64(l))
	return n
}

func (m *MsgBurn) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = m.Stable.Size()
	n += 1 + l + sovTx(uint64(l))
	return n
}

func (m *MsgBurnResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Collateral.Size()
	n += 1 + l + sovTx(uint64(l))
	l = m.Gov.Size()
	n += 1 + l + sovTx(uint64(l))
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgMintStable) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgMintStable: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgMintStable: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Stable", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Stable.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgMintStableResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgMintStableResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgMintStableResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Stable", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Stable.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgBurn) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgBurn: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgBurn: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Stable", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Stable.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgBurnResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgBurnResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgBurnResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Collateral", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Collateral.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Gov", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Gov.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
