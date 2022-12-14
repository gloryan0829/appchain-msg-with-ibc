// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: planet/mail/genesis.proto

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

// GenesisState defines the mail module's genesis state.
type GenesisState struct {
	Params               Params            `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	PortId               string            `protobuf:"bytes,2,opt,name=port_id,json=portId,proto3" json:"port_id,omitempty"`
	MessageList          []Message         `protobuf:"bytes,3,rep,name=messageList,proto3" json:"messageList"`
	MessageCount         uint64            `protobuf:"varint,4,opt,name=messageCount,proto3" json:"messageCount,omitempty"`
	SentMessageList      []SentMessage     `protobuf:"bytes,5,rep,name=sentMessageList,proto3" json:"sentMessageList"`
	SentMessageCount     uint64            `protobuf:"varint,6,opt,name=sentMessageCount,proto3" json:"sentMessageCount,omitempty"`
	TimedoutMessageList  []TimedoutMessage `protobuf:"bytes,7,rep,name=timedoutMessageList,proto3" json:"timedoutMessageList"`
	TimedoutMessageCount uint64            `protobuf:"varint,8,opt,name=timedoutMessageCount,proto3" json:"timedoutMessageCount,omitempty"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_f96363cf64ec8577, []int{0}
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

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetPortId() string {
	if m != nil {
		return m.PortId
	}
	return ""
}

func (m *GenesisState) GetMessageList() []Message {
	if m != nil {
		return m.MessageList
	}
	return nil
}

func (m *GenesisState) GetMessageCount() uint64 {
	if m != nil {
		return m.MessageCount
	}
	return 0
}

func (m *GenesisState) GetSentMessageList() []SentMessage {
	if m != nil {
		return m.SentMessageList
	}
	return nil
}

func (m *GenesisState) GetSentMessageCount() uint64 {
	if m != nil {
		return m.SentMessageCount
	}
	return 0
}

func (m *GenesisState) GetTimedoutMessageList() []TimedoutMessage {
	if m != nil {
		return m.TimedoutMessageList
	}
	return nil
}

func (m *GenesisState) GetTimedoutMessageCount() uint64 {
	if m != nil {
		return m.TimedoutMessageCount
	}
	return 0
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "planet.mail.GenesisState")
}

func init() { proto.RegisterFile("planet/mail/genesis.proto", fileDescriptor_f96363cf64ec8577) }

var fileDescriptor_f96363cf64ec8577 = []byte{
	// 342 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0x41, 0x4b, 0xc3, 0x30,
	0x14, 0xc7, 0x1b, 0x37, 0x3b, 0x4d, 0x07, 0x4a, 0x36, 0x30, 0x0e, 0x89, 0x65, 0xa7, 0x22, 0xd8,
	0xe1, 0xbc, 0x7a, 0x9a, 0x07, 0x15, 0x1c, 0xc8, 0xb6, 0x93, 0x97, 0x11, 0x59, 0x28, 0x85, 0xb5,
	0x29, 0x4b, 0x06, 0xfa, 0x2d, 0xfc, 0x40, 0x7e, 0x80, 0x1d, 0x77, 0xf4, 0x24, 0xb2, 0x7d, 0x11,
	0x69, 0x12, 0x31, 0xa9, 0xbb, 0xf5, 0xfd, 0xff, 0xff, 0xf7, 0x7e, 0xaf, 0x79, 0xf0, 0xb4, 0x98,
	0xd3, 0x9c, 0xc9, 0x5e, 0x46, 0xd3, 0x79, 0x2f, 0x61, 0x39, 0x13, 0xa9, 0x88, 0x8b, 0x05, 0x97,
	0x1c, 0x05, 0xda, 0x8a, 0x4b, 0xab, 0xd3, 0x4e, 0x78, 0xc2, 0x95, 0xde, 0x2b, 0xbf, 0x74, 0xa4,
	0x83, 0xed, 0xee, 0x82, 0x2e, 0x68, 0x66, 0x9a, 0x3b, 0xce, 0xdc, 0x8c, 0x09, 0x41, 0x13, 0x66,
	0x2c, 0x62, 0x5b, 0x82, 0xe5, 0x72, 0xea, 0xfa, 0x5d, 0xdb, 0x97, 0x69, 0xc6, 0x66, 0x7c, 0x59,
	0xc9, 0x74, 0x3f, 0x6a, 0xb0, 0x79, 0xa7, 0xb7, 0x1d, 0x4b, 0x2a, 0x19, 0xba, 0x82, 0xbe, 0xe6,
	0x63, 0x10, 0x82, 0x28, 0xe8, 0xb7, 0x62, 0x6b, 0xfb, 0xf8, 0x49, 0x59, 0x83, 0xfa, 0xea, 0xeb,
	0xdc, 0x1b, 0x99, 0x20, 0x3a, 0x81, 0x8d, 0x82, 0x2f, 0xe4, 0x34, 0x9d, 0xe1, 0xbd, 0x10, 0x44,
	0x87, 0x23, 0xbf, 0x2c, 0x1f, 0x66, 0xe8, 0x06, 0x06, 0x86, 0xf6, 0x98, 0x0a, 0x89, 0x6b, 0x61,
	0x2d, 0x0a, 0xfa, 0x6d, 0x67, 0xe0, 0x50, 0xfb, 0x66, 0xa2, 0x1d, 0x47, 0x5d, 0xd8, 0x34, 0xe5,
	0x2d, 0x5f, 0xe6, 0x12, 0xd7, 0x43, 0x10, 0xd5, 0x47, 0x8e, 0x86, 0xee, 0xe1, 0x51, 0xf9, 0xe3,
	0x43, 0x8b, 0xb2, 0xaf, 0x28, 0xd8, 0xa1, 0x8c, 0xff, 0x32, 0x86, 0x54, 0x6d, 0x43, 0x17, 0xf0,
	0xd8, 0x92, 0x34, 0xd1, 0x57, 0xc4, 0x7f, 0x3a, 0x9a, 0xc0, 0xd6, 0xef, 0x73, 0xda, 0xe4, 0x86,
	0x22, 0x9f, 0x39, 0xe4, 0x89, 0x9b, 0x33, 0xf4, 0x5d, 0xed, 0xa8, 0x0f, 0xdb, 0x15, 0x59, 0x6f,
	0x71, 0xa0, 0xb6, 0xd8, 0xe9, 0x0d, 0x2e, 0x57, 0x1b, 0x02, 0xd6, 0x1b, 0x02, 0xbe, 0x37, 0x04,
	0xbc, 0x6f, 0x89, 0xb7, 0xde, 0x12, 0xef, 0x73, 0x4b, 0xbc, 0xe7, 0x96, 0x39, 0xfe, 0xab, 0x39,
	0xff, 0x5b, 0xc1, 0xc4, 0x8b, 0xaf, 0x8e, 0x7e, 0xfd, 0x13, 0x00, 0x00, 0xff, 0xff, 0x78, 0xff,
	0xd4, 0x34, 0xad, 0x02, 0x00, 0x00,
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
	if m.TimedoutMessageCount != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.TimedoutMessageCount))
		i--
		dAtA[i] = 0x40
	}
	if len(m.TimedoutMessageList) > 0 {
		for iNdEx := len(m.TimedoutMessageList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.TimedoutMessageList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x3a
		}
	}
	if m.SentMessageCount != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.SentMessageCount))
		i--
		dAtA[i] = 0x30
	}
	if len(m.SentMessageList) > 0 {
		for iNdEx := len(m.SentMessageList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.SentMessageList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if m.MessageCount != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.MessageCount))
		i--
		dAtA[i] = 0x20
	}
	if len(m.MessageList) > 0 {
		for iNdEx := len(m.MessageList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.MessageList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.PortId) > 0 {
		i -= len(m.PortId)
		copy(dAtA[i:], m.PortId)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.PortId)))
		i--
		dAtA[i] = 0x12
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
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
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	l = len(m.PortId)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if len(m.MessageList) > 0 {
		for _, e := range m.MessageList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.MessageCount != 0 {
		n += 1 + sovGenesis(uint64(m.MessageCount))
	}
	if len(m.SentMessageList) > 0 {
		for _, e := range m.SentMessageList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.SentMessageCount != 0 {
		n += 1 + sovGenesis(uint64(m.SentMessageCount))
	}
	if len(m.TimedoutMessageList) > 0 {
		for _, e := range m.TimedoutMessageList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.TimedoutMessageCount != 0 {
		n += 1 + sovGenesis(uint64(m.TimedoutMessageCount))
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PortId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PortId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MessageList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MessageList = append(m.MessageList, Message{})
			if err := m.MessageList[len(m.MessageList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MessageCount", wireType)
			}
			m.MessageCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MessageCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SentMessageList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SentMessageList = append(m.SentMessageList, SentMessage{})
			if err := m.SentMessageList[len(m.SentMessageList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SentMessageCount", wireType)
			}
			m.SentMessageCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SentMessageCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TimedoutMessageList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TimedoutMessageList = append(m.TimedoutMessageList, TimedoutMessage{})
			if err := m.TimedoutMessageList[len(m.TimedoutMessageList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TimedoutMessageCount", wireType)
			}
			m.TimedoutMessageCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TimedoutMessageCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
