/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: k8s.io/kubernetes/vendor/k8s.io/apimachinery/pkg/util/intstr/generated.proto

/*
	Package intstr is a generated protocol buffer package.

	It is generated from these files:
		k8s.io/kubernetes/vendor/k8s.io/apimachinery/pkg/util/intstr/generated.proto

	It has these top-level messages:
		IntOrString
*/
package intstr

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

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

func (m *IntOrString) Reset()                    { *m = IntOrString{} }
func (*IntOrString) ProtoMessage()               {}
func (*IntOrString) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{0} }

func init() {
	proto.RegisterType((*IntOrString)(nil), "k8s.io.apimachinery.pkg.util.intstr.IntOrString")
}
func (m *IntOrString) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IntOrString) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0x8
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.Type))
	dAtA[i] = 0x10
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(m.IntVal))
	dAtA[i] = 0x1a
	i++
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.StrVal)))
	i += copy(dAtA[i:], m.StrVal)
	return i, nil
}

func encodeVarintGenerated(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *IntOrString) Size() (n int) {
	var l int
	_ = l
	n += 1 + sovGenerated(uint64(m.Type))
	n += 1 + sovGenerated(uint64(m.IntVal))
	l = len(m.StrVal)
	n += 1 + l + sovGenerated(uint64(l))
	return n
}

func sovGenerated(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozGenerated(x uint64) (n int) {
	return sovGenerated(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *IntOrString) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
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
			return fmt.Errorf("proto: IntOrString: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IntOrString: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= (Type(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IntVal", wireType)
			}
			m.IntVal = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.IntVal |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StrVal", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StrVal = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
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
func skipGenerated(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenerated
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
					return 0, ErrIntOverflowGenerated
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
					return 0, ErrIntOverflowGenerated
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
				return 0, ErrInvalidLengthGenerated
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowGenerated
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
				next, err := skipGenerated(dAtA[start:])
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
	ErrInvalidLengthGenerated = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenerated   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("k8s.io/kubernetes/vendor/k8s.io/apimachinery/pkg/util/intstr/generated.proto", fileDescriptorGenerated)
}

var fileDescriptorGenerated = []byte{
	// 292 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x8f, 0x31, 0x4b, 0x33, 0x31,
	0x1c, 0xc6, 0x93, 0xb7, 0x7d, 0x8b, 0x9e, 0xe0, 0x50, 0x1c, 0x8a, 0x43, 0x7a, 0x28, 0xc8, 0x0d,
	0x9a, 0xac, 0xe2, 0xd8, 0xad, 0x20, 0x08, 0x57, 0x71, 0x70, 0xbb, 0x6b, 0x63, 0x1a, 0xae, 0x4d,
	0x42, 0xee, 0x7f, 0xc2, 0x6d, 0xfd, 0x08, 0xba, 0x39, 0xfa, 0x71, 0x6e, 0xec, 0xd8, 0x41, 0x8a,
	0x17, 0xbf, 0x85, 0x93, 0x5c, 0xee, 0x40, 0xa7, 0xe4, 0x79, 0x9e, 0xdf, 0x2f, 0x90, 0xe0, 0x36,
	0xbb, 0xce, 0xa9, 0xd4, 0x2c, 0x2b, 0x52, 0x6e, 0x15, 0x07, 0x9e, 0xb3, 0x67, 0xae, 0x16, 0xda,
	0xb2, 0x6e, 0x48, 0x8c, 0x5c, 0x27, 0xf3, 0xa5, 0x54, 0xdc, 0x96, 0xcc, 0x64, 0x82, 0x15, 0x20,
	0x57, 0x4c, 0x2a, 0xc8, 0xc1, 0x32, 0xc1, 0x15, 0xb7, 0x09, 0xf0, 0x05, 0x35, 0x56, 0x83, 0x1e,
	0x9e, 0xb7, 0x12, 0xfd, 0x2b, 0x51, 0x93, 0x09, 0xda, 0x48, 0xb4, 0x95, 0x4e, 0xaf, 0x84, 0x84,
	0x65, 0x91, 0xd2, 0xb9, 0x5e, 0x33, 0xa1, 0x85, 0x66, 0xde, 0x4d, 0x8b, 0x27, 0x9f, 0x7c, 0xf0,
	0xb7, 0xf6, 0xcd, 0xb3, 0x57, 0x1c, 0x1c, 0x4d, 0x15, 0xdc, 0xd9, 0x19, 0x58, 0xa9, 0xc4, 0x30,
	0x0a, 0xfa, 0x50, 0x1a, 0x3e, 0xc2, 0x21, 0x8e, 0x7a, 0x93, 0x93, 0x6a, 0x3f, 0x46, 0x6e, 0x3f,
	0xee, 0xdf, 0x97, 0x86, 0x7f, 0x77, 0x67, 0xec, 0x89, 0xe1, 0x45, 0x30, 0x90, 0x0a, 0x1e, 0x92,
	0xd5, 0xe8, 0x5f, 0x88, 0xa3, 0xff, 0x93, 0xe3, 0x8e, 0x1d, 0x4c, 0x7d, 0x1b, 0x77, 0x6b, 0xc3,
	0xe5, 0x60, 0x1b, 0xae, 0x17, 0xe2, 0xe8, 0xf0, 0x97, 0x9b, 0xf9, 0x36, 0xee, 0xd6, 0x9b, 0x83,
	0xb7, 0xf7, 0x31, 0xda, 0x7c, 0x84, 0x68, 0x72, 0x59, 0xd5, 0x04, 0x6d, 0x6b, 0x82, 0x76, 0x35,
	0x41, 0x1b, 0x47, 0x70, 0xe5, 0x08, 0xde, 0x3a, 0x82, 0x77, 0x8e, 0xe0, 0x4f, 0x47, 0xf0, 0xcb,
	0x17, 0x41, 0x8f, 0x83, 0xf6, 0xc3, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x52, 0xa0, 0xb5, 0xc9,
	0x64, 0x01, 0x00, 0x00,
}
