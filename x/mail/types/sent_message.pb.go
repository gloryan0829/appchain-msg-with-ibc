// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: planet/mail/sent_message.proto

package types

import (
	fmt "fmt"
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

type SentMessage struct {
	Id      uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	MsgID   string `protobuf:"bytes,2,opt,name=msgID,proto3" json:"msgID,omitempty"`
	Title   string `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Chain   string `protobuf:"bytes,4,opt,name=chain,proto3" json:"chain,omitempty"`
	Creator string `protobuf:"bytes,5,opt,name=creator,proto3" json:"creator,omitempty"`
}

func (m *SentMessage) Reset()         { *m = SentMessage{} }
func (m *SentMessage) String() string { return proto.CompactTextString(m) }
func (*SentMessage) ProtoMessage()    {}
func (*SentMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_9ba1e7f0c4698d68, []int{0}
}
func (m *SentMessage) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SentMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SentMessage.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SentMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SentMessage.Merge(m, src)
}
func (m *SentMessage) XXX_Size() int {
	return m.Size()
}
func (m *SentMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_SentMessage.DiscardUnknown(m)
}

var xxx_messageInfo_SentMessage proto.InternalMessageInfo

func (m *SentMessage) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *SentMessage) GetMsgID() string {
	if m != nil {
		return m.MsgID
	}
	return ""
}

func (m *SentMessage) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *SentMessage) GetChain() string {
	if m != nil {
		return m.Chain
	}
	return ""
}

func (m *SentMessage) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func init() {
	proto.RegisterType((*SentMessage)(nil), "planet.mail.SentMessage")
}

func init() { proto.RegisterFile("planet/mail/sent_message.proto", fileDescriptor_9ba1e7f0c4698d68) }

var fileDescriptor_9ba1e7f0c4698d68 = []byte{
	// 194 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2b, 0xc8, 0x49, 0xcc,
	0x4b, 0x2d, 0xd1, 0xcf, 0x4d, 0xcc, 0xcc, 0xd1, 0x2f, 0x4e, 0xcd, 0x2b, 0x89, 0xcf, 0x4d, 0x2d,
	0x2e, 0x4e, 0x4c, 0x4f, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x86, 0xc8, 0xeb, 0x81,
	0xe4, 0x95, 0x2a, 0xb9, 0xb8, 0x83, 0x53, 0xf3, 0x4a, 0x7c, 0x21, 0x2a, 0x84, 0xf8, 0xb8, 0x98,
	0x32, 0x53, 0x24, 0x18, 0x15, 0x18, 0x35, 0x58, 0x82, 0x98, 0x32, 0x53, 0x84, 0x44, 0xb8, 0x58,
	0x73, 0x8b, 0xd3, 0x3d, 0x5d, 0x24, 0x98, 0x14, 0x18, 0x35, 0x38, 0x83, 0x20, 0x1c, 0x90, 0x68,
	0x49, 0x66, 0x49, 0x4e, 0xaa, 0x04, 0x33, 0x44, 0x14, 0xcc, 0x01, 0x89, 0x26, 0x67, 0x24, 0x66,
	0xe6, 0x49, 0xb0, 0x40, 0x44, 0xc1, 0x1c, 0x21, 0x09, 0x2e, 0xf6, 0xe4, 0xa2, 0xd4, 0xc4, 0x92,
	0xfc, 0x22, 0x09, 0x56, 0xb0, 0x38, 0x8c, 0xeb, 0xa4, 0x7b, 0xe2, 0x91, 0x1c, 0xe3, 0x85, 0x47,
	0x72, 0x8c, 0x0f, 0x1e, 0xc9, 0x31, 0x4e, 0x78, 0x2c, 0xc7, 0x70, 0xe1, 0xb1, 0x1c, 0xc3, 0x8d,
	0xc7, 0x72, 0x0c, 0x51, 0xc2, 0x50, 0x1f, 0x54, 0x40, 0xfc, 0x50, 0x52, 0x59, 0x90, 0x5a, 0x9c,
	0xc4, 0x06, 0x76, 0xbd, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x9c, 0x27, 0x6b, 0x4b, 0xdf, 0x00,
	0x00, 0x00,
}

func (m *SentMessage) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SentMessage) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SentMessage) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintSentMessage(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Chain) > 0 {
		i -= len(m.Chain)
		copy(dAtA[i:], m.Chain)
		i = encodeVarintSentMessage(dAtA, i, uint64(len(m.Chain)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintSentMessage(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.MsgID) > 0 {
		i -= len(m.MsgID)
		copy(dAtA[i:], m.MsgID)
		i = encodeVarintSentMessage(dAtA, i, uint64(len(m.MsgID)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintSentMessage(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintSentMessage(dAtA []byte, offset int, v uint64) int {
	offset -= sovSentMessage(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *SentMessage) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovSentMessage(uint64(m.Id))
	}
	l = len(m.MsgID)
	if l > 0 {
		n += 1 + l + sovSentMessage(uint64(l))
	}
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovSentMessage(uint64(l))
	}
	l = len(m.Chain)
	if l > 0 {
		n += 1 + l + sovSentMessage(uint64(l))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovSentMessage(uint64(l))
	}
	return n
}

func sovSentMessage(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSentMessage(x uint64) (n int) {
	return sovSentMessage(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SentMessage) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSentMessage
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
			return fmt.Errorf("proto: SentMessage: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SentMessage: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSentMessage
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
				return fmt.Errorf("proto: wrong wireType = %d for field MsgID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSentMessage
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
				return ErrInvalidLengthSentMessage
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSentMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MsgID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSentMessage
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
				return ErrInvalidLengthSentMessage
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSentMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Chain", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSentMessage
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
				return ErrInvalidLengthSentMessage
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSentMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Chain = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSentMessage
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
				return ErrInvalidLengthSentMessage
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSentMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSentMessage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSentMessage
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
func skipSentMessage(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSentMessage
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
					return 0, ErrIntOverflowSentMessage
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
					return 0, ErrIntOverflowSentMessage
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
				return 0, ErrInvalidLengthSentMessage
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSentMessage
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSentMessage
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthSentMessage        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSentMessage          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSentMessage = fmt.Errorf("proto: unexpected end of group")
)
