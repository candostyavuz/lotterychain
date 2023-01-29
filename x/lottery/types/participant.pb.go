// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lotterychain/lottery/participant.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
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

type Participant struct {
	Id      uint64     `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Address string     `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Bet     types.Coin `protobuf:"bytes,3,opt,name=bet,proto3" json:"bet"`
	TxData  string     `protobuf:"bytes,4,opt,name=txData,proto3" json:"txData,omitempty"`
}

func (m *Participant) Reset()         { *m = Participant{} }
func (m *Participant) String() string { return proto.CompactTextString(m) }
func (*Participant) ProtoMessage()    {}
func (*Participant) Descriptor() ([]byte, []int) {
	return fileDescriptor_001780f7f33de392, []int{0}
}
func (m *Participant) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Participant) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Participant.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Participant) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Participant.Merge(m, src)
}
func (m *Participant) XXX_Size() int {
	return m.Size()
}
func (m *Participant) XXX_DiscardUnknown() {
	xxx_messageInfo_Participant.DiscardUnknown(m)
}

var xxx_messageInfo_Participant proto.InternalMessageInfo

func (m *Participant) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Participant) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Participant) GetBet() types.Coin {
	if m != nil {
		return m.Bet
	}
	return types.Coin{}
}

func (m *Participant) GetTxData() string {
	if m != nil {
		return m.TxData
	}
	return ""
}

func init() {
	proto.RegisterType((*Participant)(nil), "lotterychain.lottery.Participant")
}

func init() {
	proto.RegisterFile("lotterychain/lottery/participant.proto", fileDescriptor_001780f7f33de392)
}

var fileDescriptor_001780f7f33de392 = []byte{
	// 250 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0xcb, 0xc9, 0x2f, 0x29,
	0x49, 0x2d, 0xaa, 0x4c, 0xce, 0x48, 0xcc, 0xcc, 0xd3, 0x87, 0x72, 0xf4, 0x0b, 0x12, 0x8b, 0x4a,
	0x32, 0x93, 0x33, 0x0b, 0x12, 0xf3, 0x4a, 0xf4, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x44, 0x90,
	0xd5, 0xe9, 0x41, 0x39, 0x52, 0x22, 0xe9, 0xf9, 0xe9, 0xf9, 0x60, 0x05, 0xfa, 0x20, 0x16, 0x44,
	0xad, 0x94, 0x5c, 0x72, 0x7e, 0x71, 0x6e, 0x7e, 0xb1, 0x7e, 0x52, 0x62, 0x71, 0xaa, 0x7e, 0x99,
	0x61, 0x52, 0x6a, 0x49, 0xa2, 0xa1, 0x7e, 0x72, 0x7e, 0x66, 0x1e, 0x44, 0x5e, 0xa9, 0x89, 0x91,
	0x8b, 0x3b, 0x00, 0x61, 0x83, 0x10, 0x1f, 0x17, 0x53, 0x66, 0x8a, 0x04, 0xa3, 0x02, 0xa3, 0x06,
	0x4b, 0x10, 0x53, 0x66, 0x8a, 0x90, 0x04, 0x17, 0x7b, 0x62, 0x4a, 0x4a, 0x51, 0x6a, 0x71, 0xb1,
	0x04, 0x93, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x8c, 0x2b, 0x64, 0xc8, 0xc5, 0x9c, 0x94, 0x5a, 0x22,
	0xc1, 0xac, 0xc0, 0xa8, 0xc1, 0x6d, 0x24, 0xa9, 0x07, 0xb1, 0x47, 0x0f, 0x64, 0x8f, 0x1e, 0xd4,
	0x1e, 0x3d, 0xe7, 0xfc, 0xcc, 0x3c, 0x27, 0x96, 0x13, 0xf7, 0xe4, 0x19, 0x82, 0x40, 0x6a, 0x85,
	0xc4, 0xb8, 0xd8, 0x4a, 0x2a, 0x5c, 0x12, 0x4b, 0x12, 0x25, 0x58, 0xc0, 0x66, 0x41, 0x79, 0x4e,
	0x66, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0xe3, 0x84, 0xc7,
	0x72, 0x0c, 0x17, 0x1e, 0xcb, 0x31, 0xdc, 0x78, 0x2c, 0xc7, 0x10, 0x25, 0x83, 0x12, 0x24, 0x15,
	0xf0, 0x40, 0x29, 0xa9, 0x2c, 0x48, 0x2d, 0x4e, 0x62, 0x03, 0xfb, 0xc1, 0x18, 0x10, 0x00, 0x00,
	0xff, 0xff, 0x2a, 0x89, 0xa6, 0xe4, 0x39, 0x01, 0x00, 0x00,
}

func (m *Participant) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Participant) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Participant) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.TxData) > 0 {
		i -= len(m.TxData)
		copy(dAtA[i:], m.TxData)
		i = encodeVarintParticipant(dAtA, i, uint64(len(m.TxData)))
		i--
		dAtA[i] = 0x22
	}
	{
		size, err := m.Bet.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintParticipant(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintParticipant(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintParticipant(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintParticipant(dAtA []byte, offset int, v uint64) int {
	offset -= sovParticipant(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Participant) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovParticipant(uint64(m.Id))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovParticipant(uint64(l))
	}
	l = m.Bet.Size()
	n += 1 + l + sovParticipant(uint64(l))
	l = len(m.TxData)
	if l > 0 {
		n += 1 + l + sovParticipant(uint64(l))
	}
	return n
}

func sovParticipant(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParticipant(x uint64) (n int) {
	return sovParticipant(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Participant) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParticipant
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
			return fmt.Errorf("proto: Participant: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Participant: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParticipant
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParticipant
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
				return ErrInvalidLengthParticipant
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParticipant
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bet", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParticipant
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
				return ErrInvalidLengthParticipant
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParticipant
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Bet.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxData", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParticipant
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
				return ErrInvalidLengthParticipant
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParticipant
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TxData = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParticipant(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParticipant
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
func skipParticipant(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParticipant
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
					return 0, ErrIntOverflowParticipant
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
					return 0, ErrIntOverflowParticipant
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
				return 0, ErrInvalidLengthParticipant
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParticipant
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParticipant
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParticipant        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParticipant          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParticipant = fmt.Errorf("proto: unexpected end of group")
)
