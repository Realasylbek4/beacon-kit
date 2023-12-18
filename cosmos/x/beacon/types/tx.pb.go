// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: polaris/beacon/v1alpha1/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types/msgservice"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	grpc "google.golang.org/grpc"
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

// Status represents the status of a transaction.
type Status int32

const (
	// STATUS_REVERT_UNSPECIFIED indicates that the transaction reverted.
	Status_STATUS_REVERT_UNSPECIFIED Status = 0
	// STATUS_SUCCESS indicates that the transaction completed successfully.
	Status_STATUS_SUCCESS Status = 1
	// STATUS_NOT_INCLUDED indicates that the transaction was not included in the
	// `evm` block, due to an critical error.
	Status_STATUS_NOT_INCLUDED Status = 2
)

var Status_name = map[int32]string{
	0: "STATUS_REVERT_UNSPECIFIED",
	1: "STATUS_SUCCESS",
	2: "STATUS_NOT_INCLUDED",
}

var Status_value = map[string]int32{
	"STATUS_REVERT_UNSPECIFIED": 0,
	"STATUS_SUCCESS":            1,
	"STATUS_NOT_INCLUDED":       2,
}

func (x Status) String() string {
	return proto.EnumName(Status_name, int32(x))
}

func (Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_eb97204e4da3cefb, []int{0}
}

// WrappedPayloadEnvelope encapsulates an Ethereum transaction as an SDK message.
type WrappedPayloadEnvelope struct {
	// data is inner transaction data of the Ethereum transaction.
	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *WrappedPayloadEnvelope) Reset()         { *m = WrappedPayloadEnvelope{} }
func (m *WrappedPayloadEnvelope) String() string { return proto.CompactTextString(m) }
func (*WrappedPayloadEnvelope) ProtoMessage()    {}
func (*WrappedPayloadEnvelope) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb97204e4da3cefb, []int{0}
}
func (m *WrappedPayloadEnvelope) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *WrappedPayloadEnvelope) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_WrappedPayloadEnvelope.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *WrappedPayloadEnvelope) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WrappedPayloadEnvelope.Merge(m, src)
}
func (m *WrappedPayloadEnvelope) XXX_Size() int {
	return m.Size()
}
func (m *WrappedPayloadEnvelope) XXX_DiscardUnknown() {
	xxx_messageInfo_WrappedPayloadEnvelope.DiscardUnknown(m)
}

var xxx_messageInfo_WrappedPayloadEnvelope proto.InternalMessageInfo

func (m *WrappedPayloadEnvelope) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

// WrappedPayloadEnvelopeResponse defines the block response.
type WrappedPayloadEnvelopeResponse struct {
}

func (m *WrappedPayloadEnvelopeResponse) Reset()         { *m = WrappedPayloadEnvelopeResponse{} }
func (m *WrappedPayloadEnvelopeResponse) String() string { return proto.CompactTextString(m) }
func (*WrappedPayloadEnvelopeResponse) ProtoMessage()    {}
func (*WrappedPayloadEnvelopeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb97204e4da3cefb, []int{1}
}
func (m *WrappedPayloadEnvelopeResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *WrappedPayloadEnvelopeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_WrappedPayloadEnvelopeResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *WrappedPayloadEnvelopeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WrappedPayloadEnvelopeResponse.Merge(m, src)
}
func (m *WrappedPayloadEnvelopeResponse) XXX_Size() int {
	return m.Size()
}
func (m *WrappedPayloadEnvelopeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_WrappedPayloadEnvelopeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_WrappedPayloadEnvelopeResponse proto.InternalMessageInfo

// WrappedForkChoice encapsulates an Ethereum transaction as an SDK message.
type WrappedForkChoice struct {
	// data is inner transaction data of the Ethereum transaction.
	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *WrappedForkChoice) Reset()         { *m = WrappedForkChoice{} }
func (m *WrappedForkChoice) String() string { return proto.CompactTextString(m) }
func (*WrappedForkChoice) ProtoMessage()    {}
func (*WrappedForkChoice) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb97204e4da3cefb, []int{2}
}
func (m *WrappedForkChoice) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *WrappedForkChoice) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_WrappedForkChoice.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *WrappedForkChoice) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WrappedForkChoice.Merge(m, src)
}
func (m *WrappedForkChoice) XXX_Size() int {
	return m.Size()
}
func (m *WrappedForkChoice) XXX_DiscardUnknown() {
	xxx_messageInfo_WrappedForkChoice.DiscardUnknown(m)
}

var xxx_messageInfo_WrappedForkChoice proto.InternalMessageInfo

func (m *WrappedForkChoice) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

// WrappedForkChoiceResponse defines the block response.
type WrappedForkChoiceResponse struct {
}

func (m *WrappedForkChoiceResponse) Reset()         { *m = WrappedForkChoiceResponse{} }
func (m *WrappedForkChoiceResponse) String() string { return proto.CompactTextString(m) }
func (*WrappedForkChoiceResponse) ProtoMessage()    {}
func (*WrappedForkChoiceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb97204e4da3cefb, []int{3}
}
func (m *WrappedForkChoiceResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *WrappedForkChoiceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_WrappedForkChoiceResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *WrappedForkChoiceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WrappedForkChoiceResponse.Merge(m, src)
}
func (m *WrappedForkChoiceResponse) XXX_Size() int {
	return m.Size()
}
func (m *WrappedForkChoiceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_WrappedForkChoiceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_WrappedForkChoiceResponse proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("polaris.beacon.v1alpha1.Status", Status_name, Status_value)
	proto.RegisterType((*WrappedPayloadEnvelope)(nil), "polaris.beacon.v1alpha1.WrappedPayloadEnvelope")
	proto.RegisterType((*WrappedPayloadEnvelopeResponse)(nil), "polaris.beacon.v1alpha1.WrappedPayloadEnvelopeResponse")
	proto.RegisterType((*WrappedForkChoice)(nil), "polaris.beacon.v1alpha1.WrappedForkChoice")
	proto.RegisterType((*WrappedForkChoiceResponse)(nil), "polaris.beacon.v1alpha1.WrappedForkChoiceResponse")
}

func init() { proto.RegisterFile("polaris/beacon/v1alpha1/tx.proto", fileDescriptor_eb97204e4da3cefb) }

var fileDescriptor_eb97204e4da3cefb = []byte{
	// 331 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xc1, 0x4a, 0x02, 0x41,
	0x18, 0xc7, 0x77, 0xa2, 0x3c, 0x0c, 0x11, 0x36, 0x42, 0xa6, 0xd1, 0x20, 0x5e, 0x0a, 0x89, 0x1d,
	0xcc, 0x27, 0xa8, 0x75, 0x04, 0x21, 0x4d, 0x76, 0x76, 0x0b, 0xba, 0xc8, 0xec, 0xee, 0xa0, 0x4b,
	0xab, 0x33, 0xec, 0x8c, 0x8b, 0xde, 0x7a, 0x84, 0x1e, 0xa5, 0xc7, 0xe8, 0xe8, 0xb1, 0x63, 0xe8,
	0xa1, 0xd7, 0x88, 0x74, 0x3d, 0xe5, 0xe9, 0xfb, 0xf8, 0xbe, 0xdf, 0xff, 0x7f, 0xf8, 0xff, 0x61,
	0x4d, 0xc9, 0x84, 0xa7, 0xb1, 0x26, 0x81, 0xe0, 0xa1, 0x9c, 0x92, 0xac, 0xc9, 0x13, 0x35, 0xe6,
	0x4d, 0x62, 0xe6, 0xb6, 0x4a, 0xa5, 0x91, 0xa8, 0x9c, 0x13, 0xf6, 0x96, 0xb0, 0x77, 0x44, 0xb5,
	0x1c, 0x4a, 0x3d, 0x91, 0x9a, 0x4c, 0xf4, 0x88, 0x64, 0xcd, 0xbf, 0xb1, 0x55, 0xd4, 0x6f, 0xe0,
	0xd9, 0x73, 0xca, 0x95, 0x12, 0xd1, 0x80, 0x2f, 0x12, 0xc9, 0x23, 0x3a, 0xcd, 0x44, 0x22, 0x95,
	0x40, 0x08, 0x1e, 0x46, 0xdc, 0xf0, 0x73, 0x50, 0x03, 0xd7, 0xc7, 0xee, 0x66, 0xaf, 0xd7, 0x20,
	0xde, 0x4f, 0xbb, 0x42, 0x2b, 0x39, 0xd5, 0xa2, 0x7e, 0x05, 0x4f, 0x73, 0xa2, 0x23, 0xd3, 0x57,
	0x67, 0x2c, 0xe3, 0x70, 0xbf, 0xd5, 0x05, 0xac, 0xfc, 0x03, 0x77, 0x2e, 0x0d, 0x0f, 0x16, 0x98,
	0xe1, 0x66, 0xa6, 0xd1, 0x25, 0xac, 0x30, 0xef, 0xce, 0xf3, 0xd9, 0xd0, 0xa5, 0x4f, 0xd4, 0xf5,
	0x86, 0x7e, 0x9f, 0x0d, 0xa8, 0xd3, 0xed, 0x74, 0x69, 0xbb, 0x68, 0x21, 0x04, 0x4f, 0xf2, 0x37,
	0xf3, 0x1d, 0x87, 0x32, 0x56, 0x04, 0xa8, 0x0c, 0x4b, 0xf9, 0xad, 0xff, 0xe8, 0x0d, 0xbb, 0x7d,
	0xe7, 0xc1, 0x6f, 0xd3, 0x76, 0xf1, 0xe0, 0xb6, 0x04, 0x61, 0x4f, 0x8f, 0x98, 0x48, 0xb3, 0x38,
	0x14, 0xd5, 0xa3, 0xb7, 0x9f, 0x8f, 0x06, 0xb8, 0xef, 0x7d, 0xae, 0x30, 0x58, 0xae, 0x30, 0xf8,
	0x5e, 0x61, 0xf0, 0xbe, 0xc6, 0xd6, 0x72, 0x8d, 0xad, 0xaf, 0x35, 0xb6, 0x5e, 0x5a, 0xa3, 0xd8,
	0x8c, 0x67, 0x81, 0x1d, 0xca, 0x09, 0x89, 0x8d, 0x8e, 0x44, 0x16, 0x08, 0x9e, 0x92, 0x20, 0x2f,
	0x21, 0x4f, 0x74, 0xbe, 0x6b, 0xc3, 0x2c, 0x94, 0xd0, 0x41, 0x61, 0x13, 0x6b, 0xeb, 0x37, 0x00,
	0x00, 0xff, 0xff, 0x2c, 0x3b, 0xc2, 0x4e, 0xac, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgServiceClient is the client API for MsgService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgServiceClient interface {
}

type msgServiceClient struct {
	cc grpc1.ClientConn
}

func NewMsgServiceClient(cc grpc1.ClientConn) MsgServiceClient {
	return &msgServiceClient{cc}
}

// MsgServiceServer is the server API for MsgService service.
type MsgServiceServer interface {
}

// UnimplementedMsgServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServiceServer struct {
}

func RegisterMsgServiceServer(s grpc1.Server, srv MsgServiceServer) {
	s.RegisterService(&_MsgService_serviceDesc, srv)
}

var _MsgService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "polaris.beacon.v1alpha1.MsgService",
	HandlerType: (*MsgServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "polaris/beacon/v1alpha1/tx.proto",
}

func (m *WrappedPayloadEnvelope) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *WrappedPayloadEnvelope) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *WrappedPayloadEnvelope) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *WrappedPayloadEnvelopeResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *WrappedPayloadEnvelopeResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *WrappedPayloadEnvelopeResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *WrappedForkChoice) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *WrappedForkChoice) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *WrappedForkChoice) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *WrappedForkChoiceResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *WrappedForkChoiceResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *WrappedForkChoiceResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
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
func (m *WrappedPayloadEnvelope) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *WrappedPayloadEnvelopeResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *WrappedForkChoice) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *WrappedForkChoiceResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *WrappedPayloadEnvelope) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: WrappedPayloadEnvelope: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: WrappedPayloadEnvelope: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
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
func (m *WrappedPayloadEnvelopeResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: WrappedPayloadEnvelopeResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: WrappedPayloadEnvelopeResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
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
func (m *WrappedForkChoice) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: WrappedForkChoice: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: WrappedForkChoice: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
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
func (m *WrappedForkChoiceResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: WrappedForkChoiceResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: WrappedForkChoiceResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
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