// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: schema/v1/what_is.proto

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

type WhatIs struct {
	// DID is the DID of the object
	Did string `protobuf:"bytes,1,opt,name=did,proto3" json:"did,omitempty"`
	// Object_doc is the object document
	Schema *Schema `protobuf:"bytes,2,opt,name=schema,proto3" json:"schema,omitempty"`
	// Creator is the DID of the creator
	Creator string `protobuf:"bytes,3,opt,name=creator,proto3" json:"creator,omitempty"`
	// Timestamp is the time of the last update of the DID Document
	Timestamp int64 `protobuf:"varint,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// IsActive is the status of the DID Document
	IsActive bool `protobuf:"varint,5,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty"`
	// Metadata is a map of key-value pairs that can be used to store additional information about the WhatIs (Schema)
	Metadata map[string]string `protobuf:"bytes,6,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (m *WhatIs) Reset()         { *m = WhatIs{} }
func (m *WhatIs) String() string { return proto.CompactTextString(m) }
func (*WhatIs) ProtoMessage()    {}
func (*WhatIs) Descriptor() ([]byte, []int) {
	return fileDescriptor_3aa27f6c5ab8b0e4, []int{0}
}
func (m *WhatIs) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *WhatIs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_WhatIs.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *WhatIs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WhatIs.Merge(m, src)
}
func (m *WhatIs) XXX_Size() int {
	return m.Size()
}
func (m *WhatIs) XXX_DiscardUnknown() {
	xxx_messageInfo_WhatIs.DiscardUnknown(m)
}

var xxx_messageInfo_WhatIs proto.InternalMessageInfo

func (m *WhatIs) GetDid() string {
	if m != nil {
		return m.Did
	}
	return ""
}

func (m *WhatIs) GetSchema() *Schema {
	if m != nil {
		return m.Schema
	}
	return nil
}

func (m *WhatIs) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *WhatIs) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *WhatIs) GetIsActive() bool {
	if m != nil {
		return m.IsActive
	}
	return false
}

func (m *WhatIs) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func init() {
	proto.RegisterType((*WhatIs)(nil), "sonrio.sonr.schema.WhatIs")
	proto.RegisterMapType((map[string]string)(nil), "sonrio.sonr.schema.WhatIs.MetadataEntry")
}

func init() { proto.RegisterFile("schema/v1/what_is.proto", fileDescriptor_3aa27f6c5ab8b0e4) }

var fileDescriptor_3aa27f6c5ab8b0e4 = []byte{
	// 329 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xc1, 0x4a, 0xfb, 0x40,
	0x10, 0xc6, 0xbb, 0xcd, 0xbf, 0xf9, 0x37, 0x5b, 0x04, 0x59, 0x8a, 0x2e, 0x55, 0x42, 0xf0, 0x20,
	0xb9, 0xb8, 0xc1, 0x7a, 0x11, 0x3d, 0x29, 0x7a, 0xf0, 0x20, 0x48, 0x3c, 0x28, 0x5e, 0xca, 0x36,
	0x5d, 0xcc, 0x52, 0xd3, 0x2d, 0xd9, 0x69, 0x35, 0x6f, 0xe1, 0x53, 0xf8, 0x2c, 0x1e, 0x7b, 0xf4,
	0x28, 0xed, 0x8b, 0x48, 0x76, 0x13, 0x8b, 0xe8, 0x69, 0xf6, 0xdb, 0xf9, 0xcd, 0xf0, 0x7d, 0x0c,
	0xde, 0xd6, 0x49, 0x2a, 0x32, 0x1e, 0xcd, 0x0f, 0xa3, 0xe7, 0x94, 0xc3, 0x40, 0x6a, 0x36, 0xcd,
	0x15, 0x28, 0x42, 0xb4, 0x9a, 0xe4, 0x52, 0xb1, 0xb2, 0x30, 0x0b, 0xf5, 0xb6, 0xd6, 0xb0, 0x7d,
	0x59, 0x76, 0xef, 0xad, 0x89, 0xdd, 0xbb, 0x94, 0xc3, 0x95, 0x26, 0x9b, 0xd8, 0x19, 0xc9, 0x11,
	0x45, 0x01, 0x0a, 0xbd, 0xb8, 0x7c, 0x92, 0x3e, 0x76, 0x2d, 0x4c, 0x9b, 0x01, 0x0a, 0x3b, 0xfd,
	0x1e, 0xfb, 0xbd, 0x99, 0xdd, 0x9a, 0x12, 0x57, 0x24, 0xa1, 0xf8, 0x7f, 0x92, 0x0b, 0x0e, 0x2a,
	0xa7, 0x8e, 0xd9, 0x54, 0x4b, 0xb2, 0x8b, 0x3d, 0x90, 0x99, 0xd0, 0xc0, 0xb3, 0x29, 0xfd, 0x17,
	0xa0, 0xd0, 0x89, 0xd7, 0x1f, 0x64, 0x07, 0x7b, 0x52, 0x0f, 0x78, 0x02, 0x72, 0x2e, 0x68, 0x2b,
	0x40, 0x61, 0x3b, 0x6e, 0x4b, 0x7d, 0x66, 0x34, 0xb9, 0xc0, 0xed, 0x4c, 0x00, 0x1f, 0x71, 0xe0,
	0xd4, 0x0d, 0x9c, 0xb0, 0xd3, 0x0f, 0xff, 0xb2, 0x62, 0x83, 0xb0, 0xeb, 0x0a, 0xbd, 0x9c, 0x40,
	0x5e, 0xc4, 0xdf, 0x93, 0xbd, 0x53, 0xbc, 0xf1, 0xa3, 0x55, 0x26, 0x1e, 0x8b, 0xa2, 0x4e, 0x3c,
	0x16, 0x05, 0xe9, 0xe2, 0xd6, 0x9c, 0x3f, 0xcd, 0x84, 0x09, 0xec, 0xc5, 0x56, 0x9c, 0x34, 0x8f,
	0xd1, 0xf9, 0xfd, 0xfb, 0xd2, 0x47, 0x8b, 0xa5, 0x8f, 0x3e, 0x97, 0x3e, 0x7a, 0x5d, 0xf9, 0x8d,
	0xc5, 0xca, 0x6f, 0x7c, 0xac, 0xfc, 0x06, 0xee, 0xd6, 0x2e, 0xa0, 0x98, 0x0a, 0x5d, 0x79, 0xb9,
	0x41, 0x0f, 0xfb, 0x8f, 0x12, 0xd2, 0xd9, 0x90, 0x25, 0x2a, 0x8b, 0xca, 0xfe, 0x81, 0x54, 0xa6,
	0x46, 0x2f, 0xd5, 0x09, 0x22, 0x33, 0x30, 0x74, 0xcd, 0x25, 0x8e, 0xbe, 0x02, 0x00, 0x00, 0xff,
	0xff, 0xd8, 0xca, 0x82, 0x81, 0xd0, 0x01, 0x00, 0x00,
}

func (m *WhatIs) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *WhatIs) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *WhatIs) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Metadata) > 0 {
		for k := range m.Metadata {
			v := m.Metadata[k]
			baseI := i
			i -= len(v)
			copy(dAtA[i:], v)
			i = encodeVarintWhatIs(dAtA, i, uint64(len(v)))
			i--
			dAtA[i] = 0x12
			i -= len(k)
			copy(dAtA[i:], k)
			i = encodeVarintWhatIs(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = encodeVarintWhatIs(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0x32
		}
	}
	if m.IsActive {
		i--
		if m.IsActive {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x28
	}
	if m.Timestamp != 0 {
		i = encodeVarintWhatIs(dAtA, i, uint64(m.Timestamp))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintWhatIs(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Schema != nil {
		{
			size, err := m.Schema.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintWhatIs(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Did) > 0 {
		i -= len(m.Did)
		copy(dAtA[i:], m.Did)
		i = encodeVarintWhatIs(dAtA, i, uint64(len(m.Did)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintWhatIs(dAtA []byte, offset int, v uint64) int {
	offset -= sovWhatIs(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *WhatIs) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Did)
	if l > 0 {
		n += 1 + l + sovWhatIs(uint64(l))
	}
	if m.Schema != nil {
		l = m.Schema.Size()
		n += 1 + l + sovWhatIs(uint64(l))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovWhatIs(uint64(l))
	}
	if m.Timestamp != 0 {
		n += 1 + sovWhatIs(uint64(m.Timestamp))
	}
	if m.IsActive {
		n += 2
	}
	if len(m.Metadata) > 0 {
		for k, v := range m.Metadata {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovWhatIs(uint64(len(k))) + 1 + len(v) + sovWhatIs(uint64(len(v)))
			n += mapEntrySize + 1 + sovWhatIs(uint64(mapEntrySize))
		}
	}
	return n
}

func sovWhatIs(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozWhatIs(x uint64) (n int) {
	return sovWhatIs(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *WhatIs) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWhatIs
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
			return fmt.Errorf("proto: WhatIs: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: WhatIs: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Did", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWhatIs
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
				return ErrInvalidLengthWhatIs
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthWhatIs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Did = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Schema", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWhatIs
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
				return ErrInvalidLengthWhatIs
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthWhatIs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Schema == nil {
				m.Schema = &Schema{}
			}
			if err := m.Schema.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWhatIs
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
				return ErrInvalidLengthWhatIs
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthWhatIs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			m.Timestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWhatIs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Timestamp |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsActive", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWhatIs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.IsActive = bool(v != 0)
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWhatIs
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
				return ErrInvalidLengthWhatIs
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthWhatIs
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Metadata == nil {
				m.Metadata = make(map[string]string)
			}
			var mapkey string
			var mapvalue string
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowWhatIs
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
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowWhatIs
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthWhatIs
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthWhatIs
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var stringLenmapvalue uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowWhatIs
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapvalue |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapvalue := int(stringLenmapvalue)
					if intStringLenmapvalue < 0 {
						return ErrInvalidLengthWhatIs
					}
					postStringIndexmapvalue := iNdEx + intStringLenmapvalue
					if postStringIndexmapvalue < 0 {
						return ErrInvalidLengthWhatIs
					}
					if postStringIndexmapvalue > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = string(dAtA[iNdEx:postStringIndexmapvalue])
					iNdEx = postStringIndexmapvalue
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipWhatIs(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if (skippy < 0) || (iNdEx+skippy) < 0 {
						return ErrInvalidLengthWhatIs
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Metadata[mapkey] = mapvalue
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipWhatIs(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthWhatIs
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
func skipWhatIs(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowWhatIs
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
					return 0, ErrIntOverflowWhatIs
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
					return 0, ErrIntOverflowWhatIs
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
				return 0, ErrInvalidLengthWhatIs
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupWhatIs
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthWhatIs
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthWhatIs        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowWhatIs          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupWhatIs = fmt.Errorf("proto: unexpected end of group")
)
