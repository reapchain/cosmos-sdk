// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmos/authz/v1beta1/query.proto

package authz

import (
	context "context"
	fmt "fmt"
	query "github.com/reapchain/cosmos-sdk/types/query"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
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

// QueryGrantsRequest is the request type for the Query/Grants RPC method.
type QueryGrantsRequest struct {
	Granter string `protobuf:"bytes,1,opt,name=granter,proto3" json:"granter,omitempty"`
	Grantee string `protobuf:"bytes,2,opt,name=grantee,proto3" json:"grantee,omitempty"`
	// Optional, msg_type_url, when set, will query only grants matching given msg type.
	MsgTypeUrl string `protobuf:"bytes,3,opt,name=msg_type_url,json=msgTypeUrl,proto3" json:"msg_type_url,omitempty"`
	// pagination defines an pagination for the request.
	Pagination *query.PageRequest `protobuf:"bytes,4,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryGrantsRequest) Reset()         { *m = QueryGrantsRequest{} }
func (m *QueryGrantsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryGrantsRequest) ProtoMessage()    {}
func (*QueryGrantsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_376d714ffdeb1545, []int{0}
}
func (m *QueryGrantsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGrantsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGrantsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGrantsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGrantsRequest.Merge(m, src)
}
func (m *QueryGrantsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryGrantsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGrantsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGrantsRequest proto.InternalMessageInfo

func (m *QueryGrantsRequest) GetGranter() string {
	if m != nil {
		return m.Granter
	}
	return ""
}

func (m *QueryGrantsRequest) GetGrantee() string {
	if m != nil {
		return m.Grantee
	}
	return ""
}

func (m *QueryGrantsRequest) GetMsgTypeUrl() string {
	if m != nil {
		return m.MsgTypeUrl
	}
	return ""
}

func (m *QueryGrantsRequest) GetPagination() *query.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}

// QueryGrantsResponse is the response type for the Query/Authorizations RPC method.
type QueryGrantsResponse struct {
	// authorizations is a list of grants granted for grantee by granter.
	Grants []*Grant `protobuf:"bytes,1,rep,name=grants,proto3" json:"grants,omitempty"`
	// pagination defines an pagination for the response.
	Pagination *query.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryGrantsResponse) Reset()         { *m = QueryGrantsResponse{} }
func (m *QueryGrantsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryGrantsResponse) ProtoMessage()    {}
func (*QueryGrantsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_376d714ffdeb1545, []int{1}
}
func (m *QueryGrantsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGrantsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGrantsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGrantsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGrantsResponse.Merge(m, src)
}
func (m *QueryGrantsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryGrantsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGrantsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGrantsResponse proto.InternalMessageInfo

func (m *QueryGrantsResponse) GetGrants() []*Grant {
	if m != nil {
		return m.Grants
	}
	return nil
}

func (m *QueryGrantsResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryGrantsRequest)(nil), "cosmos.authz.v1beta1.QueryGrantsRequest")
	proto.RegisterType((*QueryGrantsResponse)(nil), "cosmos.authz.v1beta1.QueryGrantsResponse")
}

func init() { proto.RegisterFile("cosmos/authz/v1beta1/query.proto", fileDescriptor_376d714ffdeb1545) }

var fileDescriptor_376d714ffdeb1545 = []byte{
	// 381 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x41, 0x4f, 0xf2, 0x30,
	0x18, 0xc7, 0x29, 0xbc, 0x2f, 0xc6, 0xe2, 0xa9, 0x7a, 0x58, 0x90, 0x2c, 0x0b, 0x21, 0x3a, 0x4d,
	0xec, 0x02, 0xdc, 0x3d, 0x78, 0x90, 0xab, 0x2e, 0x7a, 0xf1, 0x42, 0x3a, 0x6c, 0xca, 0x22, 0x5b,
	0xc7, 0xda, 0x19, 0xf1, 0xa8, 0x67, 0x13, 0x13, 0xbe, 0x89, 0x9f, 0xc2, 0x23, 0x89, 0x17, 0x8f,
	0x06, 0xfc, 0x20, 0x86, 0xb6, 0x08, 0xc4, 0x25, 0x7a, 0x5a, 0xb6, 0xe7, 0xf7, 0xfc, 0x9f, 0x5f,
	0x9f, 0x15, 0x3a, 0x3d, 0x2e, 0x22, 0x2e, 0x3c, 0x92, 0xc9, 0xfe, 0xbd, 0x77, 0xdb, 0x0c, 0xa8,
	0x24, 0x4d, 0x6f, 0x98, 0xd1, 0x74, 0x84, 0x93, 0x94, 0x4b, 0x8e, 0x76, 0x34, 0x81, 0x15, 0x81,
	0x0d, 0x51, 0xad, 0x31, 0xce, 0xd9, 0x80, 0x7a, 0x24, 0x09, 0x3d, 0x12, 0xc7, 0x5c, 0x12, 0x19,
	0xf2, 0x58, 0xe8, 0x9e, 0xea, 0xa1, 0x49, 0x0d, 0x88, 0xa0, 0x3a, 0xec, 0x3b, 0x3a, 0x21, 0x2c,
	0x8c, 0x15, 0x6c, 0xd8, 0x7c, 0x03, 0x3d, 0x4d, 0x11, 0xf5, 0x17, 0x00, 0xd1, 0xf9, 0x3c, 0xa4,
	0x93, 0x92, 0x58, 0x0a, 0x9f, 0x0e, 0x33, 0x2a, 0x24, 0xb2, 0xe0, 0x06, 0x9b, 0x7f, 0xa0, 0xa9,
	0x05, 0x1c, 0xe0, 0x6e, 0xfa, 0x8b, 0xd7, 0x65, 0x85, 0x5a, 0xc5, 0xd5, 0x0a, 0x45, 0x0e, 0xdc,
	0x8a, 0x04, 0xeb, 0xca, 0x51, 0x42, 0xbb, 0x59, 0x3a, 0xb0, 0x4a, 0xaa, 0x0c, 0x23, 0xc1, 0x2e,
	0x46, 0x09, 0xbd, 0x4c, 0x07, 0xe8, 0x14, 0xc2, 0xa5, 0xa2, 0xf5, 0xcf, 0x01, 0x6e, 0xa5, 0xb5,
	0x87, 0xcd, 0x0e, 0xe6, 0xe7, 0xc1, 0x7a, 0x39, 0x46, 0x14, 0x9f, 0x11, 0x46, 0x8d, 0x91, 0xbf,
	0xd2, 0x59, 0x1f, 0x03, 0xb8, 0xbd, 0x26, 0x2d, 0x12, 0x1e, 0x0b, 0x8a, 0xda, 0xb0, 0xac, 0x64,
	0x84, 0x05, 0x9c, 0x92, 0x5b, 0x69, 0xed, 0xe2, 0xbc, 0xfd, 0x62, 0xd5, 0xe5, 0x1b, 0x14, 0x75,
	0xd6, 0xa4, 0x8a, 0x4a, 0x6a, 0xff, 0x57, 0x29, 0x3d, 0x71, 0xd5, 0xaa, 0xf5, 0x04, 0xe0, 0x7f,
	0x65, 0x85, 0x1e, 0x01, 0x2c, 0x6b, 0x35, 0xe4, 0xe6, 0x2b, 0xfc, 0x5c, 0x79, 0xf5, 0xe0, 0x0f,
	0xa4, 0x9e, 0x5a, 0x6f, 0x3c, 0xbc, 0x7d, 0x8e, 0x8b, 0x36, 0xaa, 0x79, 0xb9, 0xff, 0x57, 0x1f,
	0xec, 0xe4, 0xf8, 0x75, 0x6a, 0x83, 0xc9, 0xd4, 0x06, 0x1f, 0x53, 0x1b, 0x3c, 0xcf, 0xec, 0xc2,
	0x64, 0x66, 0x17, 0xde, 0x67, 0x76, 0xe1, 0xaa, 0xc1, 0x42, 0xd9, 0xcf, 0x02, 0xdc, 0xe3, 0xd1,
	0x22, 0x41, 0x3f, 0x8e, 0xc4, 0xf5, 0x8d, 0x77, 0xa7, 0xe3, 0x82, 0xb2, 0xba, 0x21, 0xed, 0xaf,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xf2, 0x52, 0x9e, 0x65, 0xc7, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// Returns list of `Authorization`, granted to the grantee by the granter.
	Grants(ctx context.Context, in *QueryGrantsRequest, opts ...grpc.CallOption) (*QueryGrantsResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Grants(ctx context.Context, in *QueryGrantsRequest, opts ...grpc.CallOption) (*QueryGrantsResponse, error) {
	out := new(QueryGrantsResponse)
	err := c.cc.Invoke(ctx, "/cosmos.authz.v1beta1.Query/Grants", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Returns list of `Authorization`, granted to the grantee by the granter.
	Grants(context.Context, *QueryGrantsRequest) (*QueryGrantsResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Grants(ctx context.Context, req *QueryGrantsRequest) (*QueryGrantsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Grants not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Grants_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGrantsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Grants(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.authz.v1beta1.Query/Grants",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Grants(ctx, req.(*QueryGrantsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cosmos.authz.v1beta1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Grants",
			Handler:    _Query_Grants_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cosmos/authz/v1beta1/query.proto",
}

func (m *QueryGrantsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGrantsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGrantsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if len(m.MsgTypeUrl) > 0 {
		i -= len(m.MsgTypeUrl)
		copy(dAtA[i:], m.MsgTypeUrl)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.MsgTypeUrl)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Grantee) > 0 {
		i -= len(m.Grantee)
		copy(dAtA[i:], m.Grantee)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Grantee)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Granter) > 0 {
		i -= len(m.Granter)
		copy(dAtA[i:], m.Granter)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Granter)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryGrantsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGrantsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGrantsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Grants) > 0 {
		for iNdEx := len(m.Grants) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Grants[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
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
func (m *QueryGrantsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Granter)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.Grantee)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	l = len(m.MsgTypeUrl)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryGrantsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Grants) > 0 {
		for _, e := range m.Grants {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
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
func (m *QueryGrantsRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryGrantsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGrantsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Granter", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Granter = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Grantee", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Grantee = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MsgTypeUrl", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MsgTypeUrl = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageRequest{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
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
func (m *QueryGrantsResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryGrantsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGrantsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Grants", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Grants = append(m.Grants, &Grant{})
			if err := m.Grants[len(m.Grants)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageResponse{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
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
