// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: iov/configuration/v1beta1/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
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

// QueryConfigRequest is the request type for the Query/Configuration RPC
// method.
type QueryConfigRequest struct {
}

func (m *QueryConfigRequest) Reset()         { *m = QueryConfigRequest{} }
func (m *QueryConfigRequest) String() string { return proto.CompactTextString(m) }
func (*QueryConfigRequest) ProtoMessage()    {}
func (*QueryConfigRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7c58aef036fc829a, []int{0}
}
func (m *QueryConfigRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryConfigRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryConfigRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryConfigRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryConfigRequest.Merge(m, src)
}
func (m *QueryConfigRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryConfigRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryConfigRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryConfigRequest proto.InternalMessageInfo

// QueryConfigResponse is the response type for the Query/Configuration RPC
// method.
type QueryConfigResponse struct {
	// Configuration is the starname configuration.
	Config *Config `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty" yaml:"config"`
}

func (m *QueryConfigResponse) Reset()         { *m = QueryConfigResponse{} }
func (m *QueryConfigResponse) String() string { return proto.CompactTextString(m) }
func (*QueryConfigResponse) ProtoMessage()    {}
func (*QueryConfigResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7c58aef036fc829a, []int{1}
}
func (m *QueryConfigResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryConfigResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryConfigResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryConfigResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryConfigResponse.Merge(m, src)
}
func (m *QueryConfigResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryConfigResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryConfigResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryConfigResponse proto.InternalMessageInfo

// QueryFeesRequest is the request type for the Query/Configuration RPC method.
type QueryFeesRequest struct {
}

func (m *QueryFeesRequest) Reset()         { *m = QueryFeesRequest{} }
func (m *QueryFeesRequest) String() string { return proto.CompactTextString(m) }
func (*QueryFeesRequest) ProtoMessage()    {}
func (*QueryFeesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7c58aef036fc829a, []int{2}
}
func (m *QueryFeesRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryFeesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryFeesRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryFeesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryFeesRequest.Merge(m, src)
}
func (m *QueryFeesRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryFeesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryFeesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryFeesRequest proto.InternalMessageInfo

// QueryFeesResponse is the response type for the Query/Fees RPC method
type QueryFeesResponse struct {
	// Fees is the starname product fee object.
	Fees *Fees `protobuf:"bytes,1,opt,name=fees,proto3" json:"fees,omitempty" yaml:"fees"`
}

func (m *QueryFeesResponse) Reset()         { *m = QueryFeesResponse{} }
func (m *QueryFeesResponse) String() string { return proto.CompactTextString(m) }
func (*QueryFeesResponse) ProtoMessage()    {}
func (*QueryFeesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7c58aef036fc829a, []int{3}
}
func (m *QueryFeesResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryFeesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryFeesResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryFeesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryFeesResponse.Merge(m, src)
}
func (m *QueryFeesResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryFeesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryFeesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryFeesResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*QueryConfigRequest)(nil), "starnamed.x.configuration.v1beta1.QueryConfigRequest")
	proto.RegisterType((*QueryConfigResponse)(nil), "starnamed.x.configuration.v1beta1.QueryConfigResponse")
	proto.RegisterType((*QueryFeesRequest)(nil), "starnamed.x.configuration.v1beta1.QueryFeesRequest")
	proto.RegisterType((*QueryFeesResponse)(nil), "starnamed.x.configuration.v1beta1.QueryFeesResponse")
}

func init() {
	proto.RegisterFile("iov/configuration/v1beta1/query.proto", fileDescriptor_7c58aef036fc829a)
}

var fileDescriptor_7c58aef036fc829a = []byte{
	// 399 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0xcd, 0xcc, 0x2f, 0xd3,
	0x4f, 0xce, 0xcf, 0x4b, 0xcb, 0x4c, 0x2f, 0x2d, 0x4a, 0x2c, 0xc9, 0xcc, 0xcf, 0xd3, 0x2f, 0x33,
	0x4c, 0x4a, 0x2d, 0x49, 0x34, 0xd4, 0x2f, 0x2c, 0x4d, 0x2d, 0xaa, 0xd4, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0x52, 0x2c, 0x2e, 0x49, 0x2c, 0xca, 0x4b, 0xcc, 0x4d, 0x4d, 0xd1, 0xab, 0xd0, 0x43,
	0x51, 0xae, 0x07, 0x55, 0x2e, 0x25, 0x92, 0x9e, 0x9f, 0x9e, 0x0f, 0x56, 0xad, 0x0f, 0x62, 0x41,
	0x34, 0x4a, 0xc9, 0xa4, 0xe7, 0xe7, 0xa7, 0xe7, 0xa4, 0xea, 0x27, 0x16, 0x64, 0xea, 0x27, 0xe6,
	0xe5, 0xe5, 0x97, 0x80, 0x35, 0x15, 0x43, 0x65, 0xf1, 0xd8, 0x5e, 0x52, 0x59, 0x90, 0x0a, 0x55,
	0xa6, 0x24, 0xc2, 0x25, 0x14, 0x08, 0x72, 0x8c, 0x33, 0x58, 0x65, 0x50, 0x6a, 0x61, 0x69, 0x6a,
	0x71, 0x89, 0x52, 0x36, 0x97, 0x30, 0x8a, 0x68, 0x71, 0x41, 0x7e, 0x5e, 0x71, 0xaa, 0x50, 0x08,
	0x17, 0x1b, 0xc4, 0x44, 0x09, 0x46, 0x05, 0x46, 0x0d, 0x6e, 0x23, 0x4d, 0x3d, 0x82, 0x6e, 0xd7,
	0x83, 0x18, 0xe1, 0x24, 0xf8, 0xe9, 0x9e, 0x3c, 0x6f, 0x65, 0x62, 0x6e, 0x8e, 0x95, 0x12, 0x44,
	0x9d, 0x52, 0x10, 0xd4, 0x2c, 0x25, 0x21, 0x2e, 0x01, 0xb0, 0x65, 0x6e, 0xa9, 0xa9, 0xc5, 0x30,
	0x07, 0x24, 0x72, 0x09, 0x22, 0x89, 0x41, 0xad, 0xf7, 0xe1, 0x62, 0x49, 0x4b, 0x4d, 0x2d, 0x86,
	0x5a, 0xae, 0x4e, 0x84, 0xe5, 0x20, 0xed, 0x4e, 0xfc, 0x9f, 0xee, 0xc9, 0x73, 0x43, 0xac, 0x06,
	0x69, 0x57, 0x0a, 0x02, 0x9b, 0x62, 0x74, 0x99, 0x89, 0x8b, 0x15, 0x6c, 0x87, 0xd0, 0x72, 0x46,
	0x2e, 0x36, 0x88, 0x33, 0x85, 0x4c, 0x89, 0x30, 0x14, 0x33, 0xbc, 0xa4, 0xcc, 0x48, 0xd5, 0x06,
	0xf1, 0x91, 0x92, 0x5e, 0xd3, 0xe5, 0x27, 0x93, 0x99, 0x34, 0x84, 0xd4, 0xf4, 0x61, 0xfa, 0xe1,
	0x91, 0x84, 0x1a, 0x75, 0x05, 0x89, 0x45, 0x89, 0xb9, 0xc5, 0x42, 0xf3, 0x19, 0xb9, 0x58, 0x40,
	0x7e, 0x12, 0x32, 0x26, 0xd6, 0x42, 0xa4, 0x40, 0x95, 0x32, 0x21, 0x4d, 0x13, 0xd4, 0x8d, 0x3a,
	0x60, 0x37, 0xaa, 0x09, 0xa9, 0x10, 0x72, 0x23, 0x28, 0x54, 0x9d, 0x42, 0x4e, 0x3c, 0x94, 0x63,
	0x58, 0xf1, 0x48, 0x8e, 0xf1, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92,
	0x63, 0x9c, 0xf0, 0x58, 0x8e, 0xe1, 0xc2, 0x63, 0x39, 0x86, 0x1b, 0x8f, 0xe5, 0x18, 0xa2, 0x8c,
	0xd2, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0x33, 0xf3, 0xcb, 0x74, 0xf3,
	0xf3, 0x52, 0xe1, 0x26, 0xa7, 0xe8, 0x57, 0xa0, 0x19, 0x0a, 0x4e, 0xab, 0x49, 0x6c, 0xe0, 0xc4,
	0x6a, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x6d, 0xd5, 0xe0, 0xdb, 0x53, 0x03, 0x00, 0x00,
}

func (this *QueryConfigRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*QueryConfigRequest)
	if !ok {
		that2, ok := that.(QueryConfigRequest)
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
	return true
}
func (this *QueryConfigResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*QueryConfigResponse)
	if !ok {
		that2, ok := that.(QueryConfigResponse)
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
	if !this.Config.Equal(that1.Config) {
		return false
	}
	return true
}
func (this *QueryFeesRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*QueryFeesRequest)
	if !ok {
		that2, ok := that.(QueryFeesRequest)
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
	return true
}
func (this *QueryFeesResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*QueryFeesResponse)
	if !ok {
		that2, ok := that.(QueryFeesResponse)
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
	if !this.Fees.Equal(that1.Fees) {
		return false
	}
	return true
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
	// Config gets starname configuration.
	Config(ctx context.Context, in *QueryConfigRequest, opts ...grpc.CallOption) (*QueryConfigResponse, error)
	// Fees gets starname product fees.
	Fees(ctx context.Context, in *QueryFeesRequest, opts ...grpc.CallOption) (*QueryFeesResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Config(ctx context.Context, in *QueryConfigRequest, opts ...grpc.CallOption) (*QueryConfigResponse, error) {
	out := new(QueryConfigResponse)
	err := c.cc.Invoke(ctx, "/starnamed.x.configuration.v1beta1.Query/Config", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Fees(ctx context.Context, in *QueryFeesRequest, opts ...grpc.CallOption) (*QueryFeesResponse, error) {
	out := new(QueryFeesResponse)
	err := c.cc.Invoke(ctx, "/starnamed.x.configuration.v1beta1.Query/Fees", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Config gets starname configuration.
	Config(context.Context, *QueryConfigRequest) (*QueryConfigResponse, error)
	// Fees gets starname product fees.
	Fees(context.Context, *QueryFeesRequest) (*QueryFeesResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Config(ctx context.Context, req *QueryConfigRequest) (*QueryConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Config not implemented")
}
func (*UnimplementedQueryServer) Fees(ctx context.Context, req *QueryFeesRequest) (*QueryFeesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Fees not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Config_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Config(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/starnamed.x.configuration.v1beta1.Query/Config",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Config(ctx, req.(*QueryConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Fees_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryFeesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Fees(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/starnamed.x.configuration.v1beta1.Query/Fees",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Fees(ctx, req.(*QueryFeesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "starnamed.x.configuration.v1beta1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Config",
			Handler:    _Query_Config_Handler,
		},
		{
			MethodName: "Fees",
			Handler:    _Query_Fees_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "iov/configuration/v1beta1/query.proto",
}

func (m *QueryConfigRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryConfigRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryConfigRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryConfigResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryConfigResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryConfigResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Config != nil {
		{
			size, err := m.Config.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryFeesRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryFeesRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryFeesRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryFeesResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryFeesResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryFeesResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Fees != nil {
		{
			size, err := m.Fees.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
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
func (m *QueryConfigRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryConfigResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Config != nil {
		l = m.Config.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryFeesRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryFeesResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Fees != nil {
		l = m.Fees.Size()
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
func (m *QueryConfigRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryConfigRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryConfigRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
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
func (m *QueryConfigResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryConfigResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryConfigResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Config", wireType)
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
			if m.Config == nil {
				m.Config = &Config{}
			}
			if err := m.Config.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *QueryFeesRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryFeesRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryFeesRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
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
func (m *QueryFeesResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryFeesResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryFeesResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fees", wireType)
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
			if m.Fees == nil {
				m.Fees = &Fees{}
			}
			if err := m.Fees.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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