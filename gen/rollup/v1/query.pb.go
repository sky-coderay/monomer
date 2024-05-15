// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: rollup/v1/query.proto

package rollupv1

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// QueryL1BlockInfoRequest is the request type for the Query/L1BlockInfo RPC
type L1BlockInfoRequest struct {
	// L2 block height; use 0 for latest block height
	Height uint64 `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
}

func (m *L1BlockInfoRequest) Reset()         { *m = L1BlockInfoRequest{} }
func (m *L1BlockInfoRequest) String() string { return proto.CompactTextString(m) }
func (*L1BlockInfoRequest) ProtoMessage()    {}
func (*L1BlockInfoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3e27fbb9d8b6a617, []int{0}
}
func (m *L1BlockInfoRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *L1BlockInfoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_L1BlockInfoRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *L1BlockInfoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L1BlockInfoRequest.Merge(m, src)
}
func (m *L1BlockInfoRequest) XXX_Size() int {
	return m.Size()
}
func (m *L1BlockInfoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_L1BlockInfoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_L1BlockInfoRequest proto.InternalMessageInfo

// L1BlockInfoResponse is the stored L1 block info
type L1BlockInfoResponse struct {
	// Block number
	Number uint64 `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	// Block timestamp
	Time uint64 `protobuf:"varint,2,opt,name=time,proto3" json:"time,omitempty"`
	// Base fee for the block
	BaseFee []byte `protobuf:"bytes,3,opt,name=base_fee,json=baseFee,proto3" json:"base_fee,omitempty"`
	// Hash of the blocK; bytes32
	BlockHash []byte `protobuf:"bytes,4,opt,name=block_hash,json=blockHash,proto3" json:"block_hash,omitempty"`
	// Number of L2 blocks since the start of the epoch
	// Not strictly a piece of L1 information. Represents the number of L2 blocks since the start of the epoch,
	// i.e. when the actual L1 info was first introduced.
	SequenceNumber uint64 `protobuf:"varint,5,opt,name=sequence_number,json=sequenceNumber,proto3" json:"sequence_number,omitempty"`
	// Fields 6,7,8 are SystemConfig
	// Address of the batcher; bytes20
	BatcherAddr []byte `protobuf:"bytes,6,opt,name=batcher_addr,json=batcherAddr,proto3" json:"batcher_addr,omitempty"`
	// Overhead fee for L1; bytes32
	L1FeeOverhead []byte `protobuf:"bytes,7,opt,name=l1_fee_overhead,json=l1FeeOverhead,proto3" json:"l1_fee_overhead,omitempty"`
	// Scalar fee for L1; bytes32
	L1FeeScalar []byte `protobuf:"bytes,8,opt,name=l1_fee_scalar,json=l1FeeScalar,proto3" json:"l1_fee_scalar,omitempty"`
}

func (m *L1BlockInfoResponse) Reset()         { *m = L1BlockInfoResponse{} }
func (m *L1BlockInfoResponse) String() string { return proto.CompactTextString(m) }
func (*L1BlockInfoResponse) ProtoMessage()    {}
func (*L1BlockInfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3e27fbb9d8b6a617, []int{1}
}
func (m *L1BlockInfoResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *L1BlockInfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_L1BlockInfoResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *L1BlockInfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L1BlockInfoResponse.Merge(m, src)
}
func (m *L1BlockInfoResponse) XXX_Size() int {
	return m.Size()
}
func (m *L1BlockInfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_L1BlockInfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_L1BlockInfoResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*L1BlockInfoRequest)(nil), "rollup.v1.L1BlockInfoRequest")
	proto.RegisterType((*L1BlockInfoResponse)(nil), "rollup.v1.L1BlockInfoResponse")
}

func init() { proto.RegisterFile("rollup/v1/query.proto", fileDescriptor_3e27fbb9d8b6a617) }

var fileDescriptor_3e27fbb9d8b6a617 = []byte{
	// 509 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x52, 0xcf, 0x6b, 0x13, 0x51,
	0x10, 0xce, 0x6e, 0x63, 0xdb, 0xbc, 0x26, 0x16, 0x9e, 0x56, 0xd6, 0xd0, 0xae, 0x35, 0x88, 0xf6,
	0x94, 0xc7, 0x56, 0x4f, 0xf1, 0x64, 0x0e, 0x55, 0xc1, 0x1f, 0x31, 0x85, 0x50, 0xa4, 0x10, 0xde,
	0xee, 0x4e, 0x76, 0x17, 0x77, 0xdf, 0xa4, 0x6f, 0x5f, 0x02, 0x45, 0xbc, 0xf4, 0xe4, 0x51, 0xf0,
	0xe4, 0x4d, 0x3c, 0x7a, 0xf0, 0xef, 0x10, 0x4f, 0x05, 0x2f, 0x1e, 0x25, 0xf1, 0xe4, 0x5f, 0x21,
	0xfb, 0xde, 0x26, 0x2a, 0xe2, 0xed, 0x9b, 0x6f, 0x66, 0xbe, 0x19, 0xbe, 0x19, 0xb2, 0x25, 0x31,
	0x4d, 0x27, 0x63, 0x36, 0xf5, 0xd8, 0xc9, 0x04, 0xe4, 0x69, 0x7b, 0x2c, 0x51, 0x21, 0xad, 0x19,
	0xba, 0x3d, 0xf5, 0x9a, 0x97, 0x23, 0x8c, 0x50, 0xb3, 0xac, 0x40, 0xa6, 0xa0, 0xb9, 0x1d, 0x21,
	0x46, 0x29, 0x30, 0x3e, 0x4e, 0x18, 0x17, 0x02, 0x15, 0x57, 0x09, 0x8a, 0xdc, 0x64, 0x5b, 0xfb,
	0x84, 0x3e, 0xf2, 0xba, 0x29, 0x06, 0x2f, 0x1e, 0x8a, 0x11, 0xf6, 0xe1, 0x64, 0x02, 0xb9, 0xa2,
	0x57, 0xc8, 0x6a, 0x0c, 0x49, 0x14, 0x2b, 0xc7, 0xda, 0xb5, 0xf6, 0xaa, 0xfd, 0x32, 0xea, 0x54,
	0x5f, 0xbf, 0xbf, 0x56, 0x69, 0xbd, 0xb3, 0xc9, 0xa5, 0xbf, 0x9a, 0xf2, 0x31, 0x8a, 0x1c, 0x8a,
	0x2e, 0x31, 0xc9, 0x7c, 0x90, 0x8b, 0x2e, 0x13, 0x51, 0x4a, 0xaa, 0x2a, 0xc9, 0xc0, 0xb1, 0x35,
	0xab, 0x31, 0xbd, 0x4a, 0xd6, 0x7d, 0x9e, 0xc3, 0x70, 0x04, 0xe0, 0xac, 0xec, 0x5a, 0x7b, 0xf5,
	0xfe, 0x5a, 0x11, 0x1f, 0x00, 0xd0, 0x1d, 0x42, 0xfc, 0x42, 0x7b, 0x18, 0xf3, 0x3c, 0x76, 0xaa,
	0x3a, 0x59, 0xd3, 0xcc, 0x03, 0x9e, 0xc7, 0xf4, 0x16, 0xd9, 0xcc, 0x8b, 0x35, 0x45, 0x00, 0xc3,
	0x72, 0xdc, 0x05, 0x2d, 0x7c, 0x71, 0x41, 0x3f, 0x31, 0x63, 0xaf, 0x93, 0xba, 0xcf, 0x55, 0x10,
	0x83, 0x1c, 0xf2, 0x30, 0x94, 0xce, 0xaa, 0x56, 0xda, 0x28, 0xb9, 0x7b, 0x61, 0x28, 0xe9, 0x4d,
	0xb2, 0x99, 0x7a, 0xc5, 0x0e, 0x43, 0x9c, 0x82, 0x8c, 0x81, 0x87, 0xce, 0x9a, 0xae, 0x6a, 0xa4,
	0xde, 0x01, 0xc0, 0xd3, 0x92, 0xa4, 0x2d, 0xd2, 0x28, 0xeb, 0xf2, 0x80, 0xa7, 0x5c, 0x3a, 0xeb,
	0x46, 0x4b, 0x57, 0x1d, 0x6a, 0xca, 0x78, 0xb3, 0x7f, 0x66, 0x91, 0xfa, 0xb3, 0xe2, 0x3c, 0x87,
	0x20, 0xa7, 0x49, 0x00, 0x54, 0x92, 0x8d, 0x3f, 0xbc, 0xa2, 0x3b, 0xed, 0xe5, 0xbd, 0xda, 0xff,
	0x1a, 0xdf, 0x74, 0xff, 0x97, 0x36, 0x16, 0xb7, 0x6e, 0x9c, 0x7d, 0xfd, 0xf1, 0xd6, 0x76, 0xe9,
	0x36, 0x2b, 0xbf, 0x21, 0xf5, 0xb4, 0x33, 0x89, 0x18, 0x21, 0x7b, 0x69, 0xae, 0xf4, 0xaa, 0xfb,
	0xc9, 0xfa, 0x3c, 0x73, 0xad, 0xf3, 0x99, 0x6b, 0x7d, 0x9f, 0xb9, 0xd6, 0x9b, 0xb9, 0x5b, 0x39,
	0x9f, 0xbb, 0x95, 0x6f, 0x73, 0xb7, 0x42, 0x1a, 0x01, 0x66, 0xbf, 0x67, 0x74, 0x89, 0xde, 0xb5,
	0x57, 0xbc, 0x42, 0xcf, 0x7a, 0x7e, 0x27, 0x4a, 0x54, 0x3c, 0xf1, 0xdb, 0x01, 0x66, 0x6c, 0x8c,
	0xe9, 0x69, 0x06, 0x32, 0xe4, 0xc8, 0x32, 0x14, 0x98, 0x81, 0x64, 0x11, 0x08, 0xb6, 0x7c, 0xc2,
	0xbb, 0x06, 0x4d, 0xbd, 0x0f, 0xf6, 0x4a, 0xff, 0xe8, 0xe8, 0xa3, 0x5d, 0xeb, 0x1b, 0xd5, 0x81,
	0xf7, 0x65, 0x81, 0x8f, 0x07, 0xde, 0xcc, 0xde, 0x5a, 0xe2, 0xe3, 0xfb, 0xbd, 0xee, 0x63, 0x50,
	0x3c, 0xe4, 0x8a, 0xff, 0xb4, 0x89, 0xe1, 0x3b, 0x9d, 0x81, 0xe7, 0xaf, 0xea, 0x67, 0xbc, 0xfd,
	0x2b, 0x00, 0x00, 0xff, 0xff, 0x5f, 0x48, 0x35, 0x84, 0xe4, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryServiceClient is the client API for QueryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryServiceClient interface {
	L1BlockInfo(ctx context.Context, in *L1BlockInfoRequest, opts ...grpc.CallOption) (*L1BlockInfoResponse, error)
}

type queryServiceClient struct {
	cc grpc1.ClientConn
}

func NewQueryServiceClient(cc grpc1.ClientConn) QueryServiceClient {
	return &queryServiceClient{cc}
}

func (c *queryServiceClient) L1BlockInfo(ctx context.Context, in *L1BlockInfoRequest, opts ...grpc.CallOption) (*L1BlockInfoResponse, error) {
	out := new(L1BlockInfoResponse)
	err := c.cc.Invoke(ctx, "/rollup.v1.QueryService/L1BlockInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServiceServer is the server API for QueryService service.
type QueryServiceServer interface {
	L1BlockInfo(context.Context, *L1BlockInfoRequest) (*L1BlockInfoResponse, error)
}

// UnimplementedQueryServiceServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServiceServer struct {
}

func (*UnimplementedQueryServiceServer) L1BlockInfo(ctx context.Context, req *L1BlockInfoRequest) (*L1BlockInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method L1BlockInfo not implemented")
}

func RegisterQueryServiceServer(s grpc1.Server, srv QueryServiceServer) {
	s.RegisterService(&_QueryService_serviceDesc, srv)
}

func _QueryService_L1BlockInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(L1BlockInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServiceServer).L1BlockInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rollup.v1.QueryService/L1BlockInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServiceServer).L1BlockInfo(ctx, req.(*L1BlockInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _QueryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rollup.v1.QueryService",
	HandlerType: (*QueryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "L1BlockInfo",
			Handler:    _QueryService_L1BlockInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rollup/v1/query.proto",
}

func (m *L1BlockInfoRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *L1BlockInfoRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *L1BlockInfoRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Height != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.Height))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *L1BlockInfoResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *L1BlockInfoResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *L1BlockInfoResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.L1FeeScalar) > 0 {
		i -= len(m.L1FeeScalar)
		copy(dAtA[i:], m.L1FeeScalar)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.L1FeeScalar)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.L1FeeOverhead) > 0 {
		i -= len(m.L1FeeOverhead)
		copy(dAtA[i:], m.L1FeeOverhead)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.L1FeeOverhead)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.BatcherAddr) > 0 {
		i -= len(m.BatcherAddr)
		copy(dAtA[i:], m.BatcherAddr)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.BatcherAddr)))
		i--
		dAtA[i] = 0x32
	}
	if m.SequenceNumber != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.SequenceNumber))
		i--
		dAtA[i] = 0x28
	}
	if len(m.BlockHash) > 0 {
		i -= len(m.BlockHash)
		copy(dAtA[i:], m.BlockHash)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.BlockHash)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.BaseFee) > 0 {
		i -= len(m.BaseFee)
		copy(dAtA[i:], m.BaseFee)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.BaseFee)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Time != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.Time))
		i--
		dAtA[i] = 0x10
	}
	if m.Number != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.Number))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *L1BlockInfoRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Height != 0 {
		n += 1 + sovQuery(uint64(m.Height))
	}
	return n
}

func (m *L1BlockInfoResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Number != 0 {
		n += 1 + sovQuery(uint64(m.Number))
	}
	if m.Time != 0 {
		n += 1 + sovQuery(uint64(m.Time))
	}
	l = len(m.BaseFee)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.BlockHash)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	if m.SequenceNumber != 0 {
		n += 1 + sovQuery(uint64(m.SequenceNumber))
	}
	l = len(m.BatcherAddr)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.L1FeeOverhead)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.L1FeeScalar)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *L1BlockInfoRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: L1BlockInfoRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: L1BlockInfoRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Height |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *L1BlockInfoResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: L1BlockInfoResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: L1BlockInfoResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Number", wireType)
			}
			m.Number = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Number |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Time", wireType)
			}
			m.Time = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Time |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BaseFee", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BaseFee = append(m.BaseFee[:0], dAtA[iNdEx:postIndex]...)
			if m.BaseFee == nil {
				m.BaseFee = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BlockHash = append(m.BlockHash[:0], dAtA[iNdEx:postIndex]...)
			if m.BlockHash == nil {
				m.BlockHash = []byte{}
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SequenceNumber", wireType)
			}
			m.SequenceNumber = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SequenceNumber |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BatcherAddr", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BatcherAddr = append(m.BatcherAddr[:0], dAtA[iNdEx:postIndex]...)
			if m.BatcherAddr == nil {
				m.BatcherAddr = []byte{}
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field L1FeeOverhead", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.L1FeeOverhead = append(m.L1FeeOverhead[:0], dAtA[iNdEx:postIndex]...)
			if m.L1FeeOverhead == nil {
				m.L1FeeOverhead = []byte{}
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field L1FeeScalar", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.L1FeeScalar = append(m.L1FeeScalar[:0], dAtA[iNdEx:postIndex]...)
			if m.L1FeeScalar == nil {
				m.L1FeeScalar = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)