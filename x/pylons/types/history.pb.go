// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pylons/history.proto

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

type History struct {
	Address    string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Amount     string `protobuf:"bytes,2,opt,name=amount,proto3" json:"amount,omitempty"`
	CookbookID string `protobuf:"bytes,3,opt,name=cookbookID,proto3" json:"cookbookID,omitempty"`
	RecipeID   string `protobuf:"bytes,4,opt,name=recipeID,proto3" json:"recipeID,omitempty"`
	CreatedAt  int64  `protobuf:"varint,5,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	Type       string `protobuf:"bytes,6,opt,name=type,proto3" json:"type,omitempty"`
}

func (m *History) Reset()         { *m = History{} }
func (m *History) String() string { return proto.CompactTextString(m) }
func (*History) ProtoMessage()    {}
func (*History) Descriptor() ([]byte, []int) {
	return fileDescriptor_97d2aad0f73cf68e, []int{0}
}
func (m *History) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *History) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_History.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *History) XXX_Merge(src proto.Message) {
	xxx_messageInfo_History.Merge(m, src)
}
func (m *History) XXX_Size() int {
	return m.Size()
}
func (m *History) XXX_DiscardUnknown() {
	xxx_messageInfo_History.DiscardUnknown(m)
}

var xxx_messageInfo_History proto.InternalMessageInfo

func (m *History) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *History) GetAmount() string {
	if m != nil {
		return m.Amount
	}
	return ""
}

func (m *History) GetCookbookID() string {
	if m != nil {
		return m.CookbookID
	}
	return ""
}

func (m *History) GetRecipeID() string {
	if m != nil {
		return m.RecipeID
	}
	return ""
}

func (m *History) GetCreatedAt() int64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func (m *History) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func init() {
	proto.RegisterType((*History)(nil), "pylonstech.pylons.pylons.History")
}

func init() { proto.RegisterFile("pylons/history.proto", fileDescriptor_97d2aad0f73cf68e) }

var fileDescriptor_97d2aad0f73cf68e = []byte{
	// 237 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x29, 0xa8, 0xcc, 0xc9,
	0xcf, 0x2b, 0xd6, 0xcf, 0xc8, 0x2c, 0x2e, 0xc9, 0x2f, 0xaa, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0x92, 0x80, 0x88, 0x96, 0xa4, 0x26, 0x67, 0xe8, 0x41, 0x98, 0x50, 0x4a, 0x69, 0x25, 0x23,
	0x17, 0xbb, 0x07, 0x44, 0xad, 0x90, 0x04, 0x17, 0x7b, 0x62, 0x4a, 0x4a, 0x51, 0x6a, 0x71, 0xb1,
	0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x8c, 0x2b, 0x24, 0xc6, 0xc5, 0x96, 0x98, 0x9b, 0x5f,
	0x9a, 0x57, 0x22, 0xc1, 0x04, 0x96, 0x80, 0xf2, 0x84, 0xe4, 0xb8, 0xb8, 0x92, 0xf3, 0xf3, 0xb3,
	0x93, 0xf2, 0xf3, 0xb3, 0x3d, 0x5d, 0x24, 0x98, 0xc1, 0x72, 0x48, 0x22, 0x42, 0x52, 0x5c, 0x1c,
	0x45, 0xa9, 0xc9, 0x99, 0x05, 0xa9, 0x9e, 0x2e, 0x12, 0x2c, 0x60, 0x59, 0x38, 0x5f, 0x48, 0x86,
	0x8b, 0x33, 0xb9, 0x28, 0x35, 0xb1, 0x24, 0x35, 0xc5, 0xb1, 0x44, 0x82, 0x55, 0x81, 0x51, 0x83,
	0x39, 0x08, 0x21, 0x20, 0x24, 0xc4, 0xc5, 0x52, 0x52, 0x59, 0x90, 0x2a, 0xc1, 0x06, 0xd6, 0x05,
	0x66, 0x3b, 0xb9, 0x9d, 0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c,
	0x13, 0x1e, 0xcb, 0x31, 0x5c, 0x78, 0x2c, 0xc7, 0x70, 0xe3, 0xb1, 0x1c, 0x43, 0x94, 0x4e, 0x7a,
	0x66, 0x49, 0x46, 0x69, 0x92, 0x5e, 0x72, 0x7e, 0xae, 0x7e, 0x00, 0xd8, 0x63, 0xba, 0x20, 0xbf,
	0xea, 0x43, 0x03, 0xa3, 0x02, 0xc6, 0x00, 0x19, 0x53, 0x9c, 0xc4, 0x06, 0x0e, 0x14, 0x63, 0x40,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x93, 0x0e, 0xd5, 0xb3, 0x2c, 0x01, 0x00, 0x00,
}

func (m *History) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *History) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *History) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Type) > 0 {
		i -= len(m.Type)
		copy(dAtA[i:], m.Type)
		i = encodeVarintHistory(dAtA, i, uint64(len(m.Type)))
		i--
		dAtA[i] = 0x32
	}
	if m.CreatedAt != 0 {
		i = encodeVarintHistory(dAtA, i, uint64(m.CreatedAt))
		i--
		dAtA[i] = 0x28
	}
	if len(m.RecipeID) > 0 {
		i -= len(m.RecipeID)
		copy(dAtA[i:], m.RecipeID)
		i = encodeVarintHistory(dAtA, i, uint64(len(m.RecipeID)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.CookbookID) > 0 {
		i -= len(m.CookbookID)
		copy(dAtA[i:], m.CookbookID)
		i = encodeVarintHistory(dAtA, i, uint64(len(m.CookbookID)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Amount) > 0 {
		i -= len(m.Amount)
		copy(dAtA[i:], m.Amount)
		i = encodeVarintHistory(dAtA, i, uint64(len(m.Amount)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintHistory(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintHistory(dAtA []byte, offset int, v uint64) int {
	offset -= sovHistory(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *History) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovHistory(uint64(l))
	}
	l = len(m.Amount)
	if l > 0 {
		n += 1 + l + sovHistory(uint64(l))
	}
	l = len(m.CookbookID)
	if l > 0 {
		n += 1 + l + sovHistory(uint64(l))
	}
	l = len(m.RecipeID)
	if l > 0 {
		n += 1 + l + sovHistory(uint64(l))
	}
	if m.CreatedAt != 0 {
		n += 1 + sovHistory(uint64(m.CreatedAt))
	}
	l = len(m.Type)
	if l > 0 {
		n += 1 + l + sovHistory(uint64(l))
	}
	return n
}

func sovHistory(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozHistory(x uint64) (n int) {
	return sovHistory(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *History) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHistory
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
			return fmt.Errorf("proto: History: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: History: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHistory
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
				return ErrInvalidLengthHistory
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthHistory
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHistory
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
				return ErrInvalidLengthHistory
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthHistory
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Amount = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CookbookID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHistory
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
				return ErrInvalidLengthHistory
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthHistory
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CookbookID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RecipeID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHistory
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
				return ErrInvalidLengthHistory
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthHistory
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RecipeID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatedAt", wireType)
			}
			m.CreatedAt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHistory
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CreatedAt |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHistory
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
				return ErrInvalidLengthHistory
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthHistory
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Type = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipHistory(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthHistory
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
func skipHistory(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowHistory
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
					return 0, ErrIntOverflowHistory
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
					return 0, ErrIntOverflowHistory
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
				return 0, ErrInvalidLengthHistory
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupHistory
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthHistory
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthHistory        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowHistory          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupHistory = fmt.Errorf("proto: unexpected end of group")
)