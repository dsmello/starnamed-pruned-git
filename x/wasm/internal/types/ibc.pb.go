// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: x/wasm/internal/types/ibc.proto

package types

import (
	encoding_json "encoding/json"
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

// MsgIBCSend
type MsgIBCSend struct {
	// the channel by which the packet will be sent
	Channel string `protobuf:"bytes,2,opt,name=channel,proto3" json:"channel,omitempty" yaml:"source_channel"`
	// Timeout height relative to the current block height.
	// The timeout is disabled when set to 0.
	TimeoutHeight uint64 `protobuf:"varint,4,opt,name=timeout_height,json=timeoutHeight,proto3" json:"timeout_height,omitempty" yaml:"timeout_height"`
	// Timeout timestamp (in nanoseconds) relative to the current block timestamp.
	// The timeout is disabled when set to 0.
	TimeoutTimestamp uint64 `protobuf:"varint,5,opt,name=timeout_timestamp,json=timeoutTimestamp,proto3" json:"timeout_timestamp,omitempty" yaml:"timeout_timestamp"`
	// data is the payload to transfer
	Data encoding_json.RawMessage `protobuf:"bytes,6,opt,name=data,proto3,casttype=encoding/json.RawMessage" json:"data,omitempty"`
}

func (m *MsgIBCSend) Reset()         { *m = MsgIBCSend{} }
func (m *MsgIBCSend) String() string { return proto.CompactTextString(m) }
func (*MsgIBCSend) ProtoMessage()    {}
func (*MsgIBCSend) Descriptor() ([]byte, []int) {
	return fileDescriptor_9e387a38c39d89d0, []int{0}
}
func (m *MsgIBCSend) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgIBCSend) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgIBCSend.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgIBCSend) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgIBCSend.Merge(m, src)
}
func (m *MsgIBCSend) XXX_Size() int {
	return m.Size()
}
func (m *MsgIBCSend) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgIBCSend.DiscardUnknown(m)
}

var xxx_messageInfo_MsgIBCSend proto.InternalMessageInfo

// MsgIBCCloseChannel port and channel need to be owned by the contract
type MsgIBCCloseChannel struct {
	Channel string `protobuf:"bytes,2,opt,name=channel,proto3" json:"channel,omitempty" yaml:"source_channel"`
}

func (m *MsgIBCCloseChannel) Reset()         { *m = MsgIBCCloseChannel{} }
func (m *MsgIBCCloseChannel) String() string { return proto.CompactTextString(m) }
func (*MsgIBCCloseChannel) ProtoMessage()    {}
func (*MsgIBCCloseChannel) Descriptor() ([]byte, []int) {
	return fileDescriptor_9e387a38c39d89d0, []int{1}
}
func (m *MsgIBCCloseChannel) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgIBCCloseChannel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgIBCCloseChannel.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgIBCCloseChannel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgIBCCloseChannel.Merge(m, src)
}
func (m *MsgIBCCloseChannel) XXX_Size() int {
	return m.Size()
}
func (m *MsgIBCCloseChannel) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgIBCCloseChannel.DiscardUnknown(m)
}

var xxx_messageInfo_MsgIBCCloseChannel proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgIBCSend)(nil), "starnamed.x.wasm.v1beta1.MsgIBCSend")
	proto.RegisterType((*MsgIBCCloseChannel)(nil), "starnamed.x.wasm.v1beta1.MsgIBCCloseChannel")
}

func init() { proto.RegisterFile("x/wasm/internal/types/ibc.proto", fileDescriptor_9e387a38c39d89d0) }

var fileDescriptor_9e387a38c39d89d0 = []byte{
	// 345 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x91, 0xbf, 0x4e, 0xeb, 0x30,
	0x18, 0xc5, 0x93, 0xaa, 0xb7, 0x57, 0xd7, 0xba, 0xf7, 0x0a, 0x22, 0x90, 0x02, 0xaa, 0x9c, 0x2a,
	0x53, 0x17, 0x62, 0x4a, 0x37, 0x26, 0x94, 0x2e, 0x74, 0xa8, 0x84, 0x02, 0x13, 0x4b, 0xe5, 0x24,
	0x9f, 0x1c, 0xa3, 0xc4, 0xae, 0x62, 0xf7, 0xdf, 0xc6, 0x23, 0xf0, 0x58, 0x1d, 0x3b, 0x32, 0x55,
	0xd0, 0xbe, 0x41, 0x47, 0x26, 0xd4, 0x34, 0xad, 0xa8, 0xc4, 0xc4, 0x64, 0xfb, 0x7c, 0x3f, 0x1f,
	0xe9, 0x3b, 0x07, 0x39, 0x13, 0x32, 0xa6, 0x2a, 0x23, 0x5c, 0x68, 0xc8, 0x05, 0x4d, 0x89, 0x9e,
	0x0e, 0x40, 0x11, 0x1e, 0x46, 0xde, 0x20, 0x97, 0x5a, 0x5a, 0xb6, 0xd2, 0x34, 0x17, 0x34, 0x83,
	0xd8, 0x9b, 0x78, 0x1b, 0xd4, 0x1b, 0xb5, 0x42, 0xd0, 0xb4, 0x75, 0x7e, 0xc2, 0x24, 0x93, 0x05,
	0x44, 0x36, 0xb7, 0x2d, 0xef, 0x3e, 0x57, 0x10, 0xea, 0x29, 0xd6, 0xf5, 0x3b, 0xf7, 0x20, 0x62,
	0xab, 0x8d, 0x7e, 0x47, 0x09, 0x15, 0x02, 0x52, 0xbb, 0xd2, 0x30, 0x9b, 0x7f, 0xfc, 0xb3, 0xf5,
	0xc2, 0x39, 0x9d, 0xd2, 0x2c, 0xbd, 0x76, 0x95, 0x1c, 0xe6, 0x11, 0xf4, 0xcb, 0xb9, 0x1b, 0xec,
	0x48, 0xeb, 0x06, 0xfd, 0xd7, 0x3c, 0x03, 0x39, 0xd4, 0xfd, 0x04, 0x38, 0x4b, 0xb4, 0x5d, 0x6d,
	0x98, 0xcd, 0xea, 0xd7, 0xbf, 0x87, 0x73, 0x37, 0xf8, 0x57, 0x0a, 0xb7, 0xc5, 0xdb, 0xea, 0xa2,
	0xe3, 0x1d, 0xb1, 0x39, 0x95, 0xa6, 0xd9, 0xc0, 0xfe, 0x55, 0x98, 0xd4, 0xd7, 0x0b, 0xc7, 0x3e,
	0x34, 0xd9, 0x23, 0x6e, 0x70, 0x54, 0x6a, 0x0f, 0x3b, 0xc9, 0xba, 0x44, 0xd5, 0x98, 0x6a, 0x6a,
	0xd7, 0x1a, 0x66, 0xf3, 0xaf, 0x5f, 0xff, 0x58, 0x38, 0x36, 0x88, 0x48, 0xc6, 0x5c, 0x30, 0xf2,
	0xa4, 0xa4, 0xf0, 0x02, 0x3a, 0xee, 0x81, 0x52, 0x94, 0x41, 0x50, 0x90, 0x6e, 0x17, 0x59, 0xdb,
	0x04, 0x3a, 0xa9, 0x54, 0xd0, 0x29, 0x97, 0xfa, 0x49, 0x12, 0xfe, 0xdd, 0xec, 0x1d, 0x1b, 0xb3,
	0x25, 0x36, 0xe7, 0x4b, 0x6c, 0xbe, 0x2d, 0xb1, 0xf9, 0xb2, 0xc2, 0xc6, 0x7c, 0x85, 0x8d, 0xd7,
	0x15, 0x36, 0x1e, 0xaf, 0x18, 0xd7, 0xc9, 0x30, 0xf4, 0x22, 0x99, 0x11, 0x2e, 0x47, 0x17, 0x52,
	0x00, 0xd9, 0xd7, 0x45, 0xbe, 0x6d, 0x36, 0xac, 0x15, 0x35, 0xb5, 0x3f, 0x03, 0x00, 0x00, 0xff,
	0xff, 0xd9, 0x29, 0xad, 0x3a, 0xf9, 0x01, 0x00, 0x00,
}

func (m *MsgIBCSend) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgIBCSend) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgIBCSend) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintIbc(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0x32
	}
	if m.TimeoutTimestamp != 0 {
		i = encodeVarintIbc(dAtA, i, uint64(m.TimeoutTimestamp))
		i--
		dAtA[i] = 0x28
	}
	if m.TimeoutHeight != 0 {
		i = encodeVarintIbc(dAtA, i, uint64(m.TimeoutHeight))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Channel) > 0 {
		i -= len(m.Channel)
		copy(dAtA[i:], m.Channel)
		i = encodeVarintIbc(dAtA, i, uint64(len(m.Channel)))
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}

func (m *MsgIBCCloseChannel) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgIBCCloseChannel) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgIBCCloseChannel) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Channel) > 0 {
		i -= len(m.Channel)
		copy(dAtA[i:], m.Channel)
		i = encodeVarintIbc(dAtA, i, uint64(len(m.Channel)))
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}

func encodeVarintIbc(dAtA []byte, offset int, v uint64) int {
	offset -= sovIbc(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgIBCSend) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Channel)
	if l > 0 {
		n += 1 + l + sovIbc(uint64(l))
	}
	if m.TimeoutHeight != 0 {
		n += 1 + sovIbc(uint64(m.TimeoutHeight))
	}
	if m.TimeoutTimestamp != 0 {
		n += 1 + sovIbc(uint64(m.TimeoutTimestamp))
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovIbc(uint64(l))
	}
	return n
}

func (m *MsgIBCCloseChannel) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Channel)
	if l > 0 {
		n += 1 + l + sovIbc(uint64(l))
	}
	return n
}

func sovIbc(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozIbc(x uint64) (n int) {
	return sovIbc(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgIBCSend) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIbc
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
			return fmt.Errorf("proto: MsgIBCSend: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgIBCSend: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Channel", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIbc
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
				return ErrInvalidLengthIbc
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIbc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Channel = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TimeoutHeight", wireType)
			}
			m.TimeoutHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIbc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TimeoutHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TimeoutTimestamp", wireType)
			}
			m.TimeoutTimestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIbc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TimeoutTimestamp |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIbc
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
				return ErrInvalidLengthIbc
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthIbc
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
			skippy, err := skipIbc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthIbc
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthIbc
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
func (m *MsgIBCCloseChannel) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIbc
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
			return fmt.Errorf("proto: MsgIBCCloseChannel: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgIBCCloseChannel: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Channel", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIbc
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
				return ErrInvalidLengthIbc
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIbc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Channel = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipIbc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthIbc
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthIbc
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
func skipIbc(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowIbc
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
					return 0, ErrIntOverflowIbc
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
					return 0, ErrIntOverflowIbc
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
				return 0, ErrInvalidLengthIbc
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupIbc
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthIbc
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthIbc        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowIbc          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupIbc = fmt.Errorf("proto: unexpected end of group")
)
