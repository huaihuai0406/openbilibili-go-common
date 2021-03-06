// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: api/http/v1/http.proto

package v1

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type XVerifyReq struct {
	XAnti                string   `protobuf:"bytes,1,opt,name=_anti,json=Anti,proto3" json:"_anti,omitempty" form:"_anti" validate:"required"`
	Uid                  int64    `protobuf:"varint,2,opt,name=uid,proto3" json:"uid,omitempty" form:"uid" validate:"required"`
	ClientIp             string   `protobuf:"bytes,3,opt,name=client_ip,json=clientIp,proto3" json:"client_ip,omitempty" form:"client_ip"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *XVerifyReq) Reset()         { *m = XVerifyReq{} }
func (m *XVerifyReq) String() string { return proto.CompactTextString(m) }
func (*XVerifyReq) ProtoMessage()    {}
func (*XVerifyReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_http_6086dd8d3a71844a, []int{0}
}
func (m *XVerifyReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *XVerifyReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_XVerifyReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *XVerifyReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_XVerifyReq.Merge(dst, src)
}
func (m *XVerifyReq) XXX_Size() int {
	return m.Size()
}
func (m *XVerifyReq) XXX_DiscardUnknown() {
	xxx_messageInfo_XVerifyReq.DiscardUnknown(m)
}

var xxx_messageInfo_XVerifyReq proto.InternalMessageInfo

func (m *XVerifyReq) GetXAnti() string {
	if m != nil {
		return m.XAnti
	}
	return ""
}

func (m *XVerifyReq) GetUid() int64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *XVerifyReq) GetClientIp() string {
	if m != nil {
		return m.ClientIp
	}
	return ""
}

type XVerifyResp struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *XVerifyResp) Reset()         { *m = XVerifyResp{} }
func (m *XVerifyResp) String() string { return proto.CompactTextString(m) }
func (*XVerifyResp) ProtoMessage()    {}
func (*XVerifyResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_http_6086dd8d3a71844a, []int{1}
}
func (m *XVerifyResp) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *XVerifyResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_XVerifyResp.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *XVerifyResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_XVerifyResp.Merge(dst, src)
}
func (m *XVerifyResp) XXX_Size() int {
	return m.Size()
}
func (m *XVerifyResp) XXX_DiscardUnknown() {
	xxx_messageInfo_XVerifyResp.DiscardUnknown(m)
}

var xxx_messageInfo_XVerifyResp proto.InternalMessageInfo

func init() {
	proto.RegisterType((*XVerifyReq)(nil), "live.xcaptcha.v1.XVerifyReq")
	proto.RegisterType((*XVerifyResp)(nil), "live.xcaptcha.v1.XVerifyResp")
}
func (m *XVerifyReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *XVerifyReq) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.XAnti) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintHttp(dAtA, i, uint64(len(m.XAnti)))
		i += copy(dAtA[i:], m.XAnti)
	}
	if m.Uid != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintHttp(dAtA, i, uint64(m.Uid))
	}
	if len(m.ClientIp) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintHttp(dAtA, i, uint64(len(m.ClientIp)))
		i += copy(dAtA[i:], m.ClientIp)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *XVerifyResp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *XVerifyResp) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintHttp(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *XVerifyReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.XAnti)
	if l > 0 {
		n += 1 + l + sovHttp(uint64(l))
	}
	if m.Uid != 0 {
		n += 1 + sovHttp(uint64(m.Uid))
	}
	l = len(m.ClientIp)
	if l > 0 {
		n += 1 + l + sovHttp(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *XVerifyResp) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovHttp(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozHttp(x uint64) (n int) {
	return sovHttp(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *XVerifyReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHttp
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: XVerifyReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: XVerifyReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field XAnti", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHttp
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthHttp
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.XAnti = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Uid", wireType)
			}
			m.Uid = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHttp
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Uid |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientIp", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHttp
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthHttp
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClientIp = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipHttp(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHttp
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *XVerifyResp) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHttp
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: XVerifyResp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: XVerifyResp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipHttp(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHttp
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipHttp(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowHttp
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
					return 0, ErrIntOverflowHttp
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowHttp
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthHttp
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowHttp
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipHttp(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthHttp = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowHttp   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("api/http/v1/http.proto", fileDescriptor_http_6086dd8d3a71844a) }

var fileDescriptor_http_6086dd8d3a71844a = []byte{
	// 287 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4b, 0x2c, 0xc8, 0xd4,
	0xcf, 0x28, 0x29, 0x29, 0xd0, 0x2f, 0x33, 0x04, 0xd3, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42,
	0x02, 0x39, 0x99, 0x65, 0xa9, 0x7a, 0x15, 0xc9, 0x89, 0x05, 0x25, 0xc9, 0x19, 0x89, 0x7a, 0x65,
	0x86, 0x52, 0xba, 0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0xfa, 0xe9, 0xf9,
	0xe9, 0xf9, 0xfa, 0x60, 0x85, 0x49, 0xa5, 0x69, 0x60, 0x1e, 0x98, 0x03, 0x66, 0x41, 0x0c, 0x50,
	0xda, 0xc0, 0xc8, 0xc5, 0x15, 0x11, 0x96, 0x5a, 0x94, 0x99, 0x56, 0x19, 0x94, 0x5a, 0x28, 0x64,
	0xc9, 0xc5, 0x1a, 0x9f, 0x98, 0x57, 0x92, 0x29, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0xe9, 0xa4, 0xf2,
	0xe9, 0x9e, 0xbc, 0x42, 0x5a, 0x7e, 0x51, 0xae, 0x95, 0x12, 0x58, 0x58, 0x49, 0xa1, 0x2c, 0x31,
	0x27, 0x33, 0x25, 0xb1, 0x24, 0xd5, 0x4a, 0xa9, 0x28, 0xb5, 0xb0, 0x34, 0xb3, 0x28, 0x35, 0x45,
	0x29, 0x88, 0xc5, 0x31, 0xaf, 0x24, 0x53, 0xc8, 0x84, 0x8b, 0xb9, 0x34, 0x33, 0x45, 0x82, 0x49,
	0x81, 0x51, 0x83, 0xd9, 0x49, 0xe9, 0xd3, 0x3d, 0x79, 0x39, 0x88, 0xc6, 0xd2, 0xcc, 0x14, 0xec,
	0xda, 0x40, 0xca, 0x85, 0x0c, 0xb9, 0x38, 0x93, 0x73, 0x32, 0x53, 0xf3, 0x4a, 0xe2, 0x33, 0x0b,
	0x24, 0x98, 0xc1, 0x96, 0x8a, 0x7c, 0xba, 0x27, 0x2f, 0x00, 0xd1, 0x0b, 0x97, 0x52, 0x0a, 0xe2,
	0x80, 0xb0, 0x3d, 0x0b, 0x94, 0x78, 0xb9, 0xb8, 0xe1, 0x2e, 0x2e, 0x2e, 0x30, 0x0a, 0xe0, 0x62,
	0x77, 0x86, 0x78, 0x5f, 0xc8, 0x95, 0x8b, 0xad, 0x0c, 0x2c, 0x21, 0x24, 0xa3, 0x87, 0x1e, 0x30,
	0x7a, 0x08, 0x5f, 0x4a, 0xc9, 0xe2, 0x91, 0x2d, 0x2e, 0x70, 0x12, 0x38, 0xf1, 0x48, 0x8e, 0xf1,
	0xc2, 0x23, 0x39, 0xc6, 0x07, 0x8f, 0xe4, 0x18, 0xa3, 0x98, 0xca, 0x0c, 0x93, 0xd8, 0xc0, 0x81,
	0x65, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x50, 0xf8, 0xd1, 0xfb, 0x87, 0x01, 0x00, 0x00,
}
