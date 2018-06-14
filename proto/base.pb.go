// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: base.proto

/*
	Package base is a generated protocol buffer package.

	It is generated from these files:
		base.proto

	It has these top-level messages:
		CompileRequest
		CompileResponse
*/
package base

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import strings "strings"
import reflect "reflect"

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

type CompileRequest struct {
	SourceCode []string `protobuf:"bytes,1,rep,name=SourceCode" json:"SourceCode,omitempty"`
	Language   string   `protobuf:"bytes,2,opt,name=Language,proto3" json:"Language,omitempty"`
}

func (m *CompileRequest) Reset()                    { *m = CompileRequest{} }
func (*CompileRequest) ProtoMessage()               {}
func (*CompileRequest) Descriptor() ([]byte, []int) { return fileDescriptorBase, []int{0} }

func (m *CompileRequest) GetSourceCode() []string {
	if m != nil {
		return m.SourceCode
	}
	return nil
}

func (m *CompileRequest) GetLanguage() string {
	if m != nil {
		return m.Language
	}
	return ""
}

type CompileResponse struct {
	Result string `protobuf:"bytes,2,opt,name=Result,proto3" json:"Result,omitempty"`
}

func (m *CompileResponse) Reset()                    { *m = CompileResponse{} }
func (*CompileResponse) ProtoMessage()               {}
func (*CompileResponse) Descriptor() ([]byte, []int) { return fileDescriptorBase, []int{1} }

func (m *CompileResponse) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

func init() {
	proto.RegisterType((*CompileRequest)(nil), "CompileRequest")
	proto.RegisterType((*CompileResponse)(nil), "CompileResponse")
}
func (this *CompileRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*CompileRequest)
	if !ok {
		that2, ok := that.(CompileRequest)
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
	if len(this.SourceCode) != len(that1.SourceCode) {
		return false
	}
	for i := range this.SourceCode {
		if this.SourceCode[i] != that1.SourceCode[i] {
			return false
		}
	}
	if this.Language != that1.Language {
		return false
	}
	return true
}
func (this *CompileResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*CompileResponse)
	if !ok {
		that2, ok := that.(CompileResponse)
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
	if this.Result != that1.Result {
		return false
	}
	return true
}
func (this *CompileRequest) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&base.CompileRequest{")
	s = append(s, "SourceCode: "+fmt.Sprintf("%#v", this.SourceCode)+",\n")
	s = append(s, "Language: "+fmt.Sprintf("%#v", this.Language)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *CompileResponse) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&base.CompileResponse{")
	s = append(s, "Result: "+fmt.Sprintf("%#v", this.Result)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringBase(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *CompileRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CompileRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.SourceCode) > 0 {
		for _, s := range m.SourceCode {
			dAtA[i] = 0xa
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	if len(m.Language) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintBase(dAtA, i, uint64(len(m.Language)))
		i += copy(dAtA[i:], m.Language)
	}
	return i, nil
}

func (m *CompileResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CompileResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Result) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintBase(dAtA, i, uint64(len(m.Result)))
		i += copy(dAtA[i:], m.Result)
	}
	return i, nil
}

func encodeVarintBase(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *CompileRequest) Size() (n int) {
	var l int
	_ = l
	if len(m.SourceCode) > 0 {
		for _, s := range m.SourceCode {
			l = len(s)
			n += 1 + l + sovBase(uint64(l))
		}
	}
	l = len(m.Language)
	if l > 0 {
		n += 1 + l + sovBase(uint64(l))
	}
	return n
}

func (m *CompileResponse) Size() (n int) {
	var l int
	_ = l
	l = len(m.Result)
	if l > 0 {
		n += 1 + l + sovBase(uint64(l))
	}
	return n
}

func sovBase(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozBase(x uint64) (n int) {
	return sovBase(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *CompileRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&CompileRequest{`,
		`SourceCode:` + fmt.Sprintf("%v", this.SourceCode) + `,`,
		`Language:` + fmt.Sprintf("%v", this.Language) + `,`,
		`}`,
	}, "")
	return s
}
func (this *CompileResponse) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&CompileResponse{`,
		`Result:` + fmt.Sprintf("%v", this.Result) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringBase(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *CompileRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBase
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
			return fmt.Errorf("proto: CompileRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CompileRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SourceCode", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBase
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
				return ErrInvalidLengthBase
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SourceCode = append(m.SourceCode, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Language", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBase
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
				return ErrInvalidLengthBase
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Language = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBase(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthBase
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
func (m *CompileResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBase
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
			return fmt.Errorf("proto: CompileResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CompileResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Result", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBase
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
				return ErrInvalidLengthBase
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Result = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBase(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthBase
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
func skipBase(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBase
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
					return 0, ErrIntOverflowBase
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
					return 0, ErrIntOverflowBase
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
				return 0, ErrInvalidLengthBase
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowBase
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
				next, err := skipBase(dAtA[start:])
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
	ErrInvalidLengthBase = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBase   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("base.proto", fileDescriptorBase) }

var fileDescriptorBase = []byte{
	// 203 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0x4a, 0x2c, 0x4e,
	0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0xf2, 0xe1, 0xe2, 0x73, 0xce, 0xcf, 0x2d, 0xc8, 0xcc,
	0x49, 0x0d, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x92, 0xe3, 0xe2, 0x0a, 0xce, 0x2f, 0x2d,
	0x4a, 0x4e, 0x75, 0xce, 0x4f, 0x49, 0x95, 0x60, 0x54, 0x60, 0xd6, 0xe0, 0x0c, 0x42, 0x12, 0x11,
	0x92, 0xe2, 0xe2, 0xf0, 0x49, 0xcc, 0x4b, 0x2f, 0x4d, 0x4c, 0x4f, 0x95, 0x60, 0x52, 0x60, 0xd4,
	0xe0, 0x0c, 0x82, 0xf3, 0x95, 0x34, 0xb9, 0xf8, 0xe1, 0xa6, 0x15, 0x17, 0xe4, 0xe7, 0x15, 0xa7,
	0x0a, 0x89, 0x71, 0xb1, 0x05, 0xa5, 0x16, 0x97, 0xe6, 0x94, 0x40, 0x15, 0x43, 0x79, 0x46, 0xa6,
	0x5c, 0xec, 0x50, 0xa5, 0x42, 0x5a, 0x5c, 0xcc, 0x41, 0xa5, 0x79, 0x42, 0xfc, 0x7a, 0xa8, 0x2e,
	0x91, 0x12, 0xd0, 0x43, 0x33, 0x4c, 0x89, 0xc1, 0x49, 0xe7, 0xc2, 0x43, 0x39, 0x86, 0x1b, 0x0f,
	0xe5, 0x18, 0x3e, 0x3c, 0x94, 0x63, 0x6c, 0x78, 0x24, 0xc7, 0xb8, 0xe2, 0x91, 0x1c, 0xe3, 0x89,
	0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7, 0xf8, 0xe2, 0x91, 0x1c, 0xc3,
	0x87, 0x47, 0x72, 0x8c, 0x13, 0x1e, 0xcb, 0x31, 0x24, 0xb1, 0x81, 0x3d, 0x69, 0x0c, 0x08, 0x00,
	0x00, 0xff, 0xff, 0x69, 0xf2, 0x31, 0x79, 0xf2, 0x00, 0x00, 0x00,
}
