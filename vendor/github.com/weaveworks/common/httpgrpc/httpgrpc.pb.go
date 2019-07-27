// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/weaveworks/common/httpgrpc/httpgrpc.proto

/*
	Package httpgrpc is a generated protocol buffer package.

	It is generated from these files:
		github.com/weaveworks/common/httpgrpc/httpgrpc.proto

	It has these top-level messages:
		HTTPRequest
		HTTPResponse
		Header
*/
package httpgrpc

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import bytes "bytes"

import strings "strings"
import reflect "reflect"

import context "golang.org/x/net/context"
import grpc "google.golang.org/grpc"

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

type HTTPRequest struct {
	Method  string    `protobuf:"bytes,1,opt,name=method,proto3" json:"method,omitempty"`
	Url     string    `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	Headers []*Header `protobuf:"bytes,3,rep,name=headers" json:"headers,omitempty"`
	Body    []byte    `protobuf:"bytes,4,opt,name=body,proto3" json:"body,omitempty"`
}

func (m *HTTPRequest) Reset()                    { *m = HTTPRequest{} }
func (*HTTPRequest) ProtoMessage()               {}
func (*HTTPRequest) Descriptor() ([]byte, []int) { return fileDescriptorHttpgrpc, []int{0} }

func (m *HTTPRequest) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *HTTPRequest) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *HTTPRequest) GetHeaders() []*Header {
	if m != nil {
		return m.Headers
	}
	return nil
}

func (m *HTTPRequest) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

type HTTPResponse struct {
	Code    int32     `protobuf:"varint,1,opt,name=Code,json=code,proto3" json:"Code,omitempty"`
	Headers []*Header `protobuf:"bytes,2,rep,name=headers" json:"headers,omitempty"`
	Body    []byte    `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
}

func (m *HTTPResponse) Reset()                    { *m = HTTPResponse{} }
func (*HTTPResponse) ProtoMessage()               {}
func (*HTTPResponse) Descriptor() ([]byte, []int) { return fileDescriptorHttpgrpc, []int{1} }

func (m *HTTPResponse) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *HTTPResponse) GetHeaders() []*Header {
	if m != nil {
		return m.Headers
	}
	return nil
}

func (m *HTTPResponse) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

type Header struct {
	Key    string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Values []string `protobuf:"bytes,2,rep,name=values" json:"values,omitempty"`
}

func (m *Header) Reset()                    { *m = Header{} }
func (*Header) ProtoMessage()               {}
func (*Header) Descriptor() ([]byte, []int) { return fileDescriptorHttpgrpc, []int{2} }

func (m *Header) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Header) GetValues() []string {
	if m != nil {
		return m.Values
	}
	return nil
}

func init() {
	proto.RegisterType((*HTTPRequest)(nil), "httpgrpc.HTTPRequest")
	proto.RegisterType((*HTTPResponse)(nil), "httpgrpc.HTTPResponse")
	proto.RegisterType((*Header)(nil), "httpgrpc.Header")
}
func (this *HTTPRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*HTTPRequest)
	if !ok {
		that2, ok := that.(HTTPRequest)
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
	if this.Method != that1.Method {
		return false
	}
	if this.Url != that1.Url {
		return false
	}
	if len(this.Headers) != len(that1.Headers) {
		return false
	}
	for i := range this.Headers {
		if !this.Headers[i].Equal(that1.Headers[i]) {
			return false
		}
	}
	if !bytes.Equal(this.Body, that1.Body) {
		return false
	}
	return true
}
func (this *HTTPResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*HTTPResponse)
	if !ok {
		that2, ok := that.(HTTPResponse)
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
	if this.Code != that1.Code {
		return false
	}
	if len(this.Headers) != len(that1.Headers) {
		return false
	}
	for i := range this.Headers {
		if !this.Headers[i].Equal(that1.Headers[i]) {
			return false
		}
	}
	if !bytes.Equal(this.Body, that1.Body) {
		return false
	}
	return true
}
func (this *Header) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Header)
	if !ok {
		that2, ok := that.(Header)
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
	if this.Key != that1.Key {
		return false
	}
	if len(this.Values) != len(that1.Values) {
		return false
	}
	for i := range this.Values {
		if this.Values[i] != that1.Values[i] {
			return false
		}
	}
	return true
}
func (this *HTTPRequest) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 8)
	s = append(s, "&httpgrpc.HTTPRequest{")
	s = append(s, "Method: "+fmt.Sprintf("%#v", this.Method)+",\n")
	s = append(s, "Url: "+fmt.Sprintf("%#v", this.Url)+",\n")
	if this.Headers != nil {
		s = append(s, "Headers: "+fmt.Sprintf("%#v", this.Headers)+",\n")
	}
	s = append(s, "Body: "+fmt.Sprintf("%#v", this.Body)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *HTTPResponse) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&httpgrpc.HTTPResponse{")
	s = append(s, "Code: "+fmt.Sprintf("%#v", this.Code)+",\n")
	if this.Headers != nil {
		s = append(s, "Headers: "+fmt.Sprintf("%#v", this.Headers)+",\n")
	}
	s = append(s, "Body: "+fmt.Sprintf("%#v", this.Body)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *Header) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&httpgrpc.Header{")
	s = append(s, "Key: "+fmt.Sprintf("%#v", this.Key)+",\n")
	s = append(s, "Values: "+fmt.Sprintf("%#v", this.Values)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringHttpgrpc(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for HTTP service

type HTTPClient interface {
	Handle(ctx context.Context, in *HTTPRequest, opts ...grpc.CallOption) (*HTTPResponse, error)
}

type hTTPClient struct {
	cc *grpc.ClientConn
}

func NewHTTPClient(cc *grpc.ClientConn) HTTPClient {
	return &hTTPClient{cc}
}

func (c *hTTPClient) Handle(ctx context.Context, in *HTTPRequest, opts ...grpc.CallOption) (*HTTPResponse, error) {
	out := new(HTTPResponse)
	err := grpc.Invoke(ctx, "/httpgrpc.HTTP/Handle", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for HTTP service

type HTTPServer interface {
	Handle(context.Context, *HTTPRequest) (*HTTPResponse, error)
}

func RegisterHTTPServer(s *grpc.Server, srv HTTPServer) {
	s.RegisterService(&_HTTP_serviceDesc, srv)
}

func _HTTP_Handle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HTTPRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HTTPServer).Handle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/httpgrpc.HTTP/Handle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HTTPServer).Handle(ctx, req.(*HTTPRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _HTTP_serviceDesc = grpc.ServiceDesc{
	ServiceName: "httpgrpc.HTTP",
	HandlerType: (*HTTPServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Handle",
			Handler:    _HTTP_Handle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/weaveworks/common/httpgrpc/httpgrpc.proto",
}

func (m *HTTPRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HTTPRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Method) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintHttpgrpc(dAtA, i, uint64(len(m.Method)))
		i += copy(dAtA[i:], m.Method)
	}
	if len(m.Url) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintHttpgrpc(dAtA, i, uint64(len(m.Url)))
		i += copy(dAtA[i:], m.Url)
	}
	if len(m.Headers) > 0 {
		for _, msg := range m.Headers {
			dAtA[i] = 0x1a
			i++
			i = encodeVarintHttpgrpc(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.Body) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintHttpgrpc(dAtA, i, uint64(len(m.Body)))
		i += copy(dAtA[i:], m.Body)
	}
	return i, nil
}

func (m *HTTPResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HTTPResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Code != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintHttpgrpc(dAtA, i, uint64(m.Code))
	}
	if len(m.Headers) > 0 {
		for _, msg := range m.Headers {
			dAtA[i] = 0x12
			i++
			i = encodeVarintHttpgrpc(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.Body) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintHttpgrpc(dAtA, i, uint64(len(m.Body)))
		i += copy(dAtA[i:], m.Body)
	}
	return i, nil
}

func (m *Header) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Header) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Key) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintHttpgrpc(dAtA, i, uint64(len(m.Key)))
		i += copy(dAtA[i:], m.Key)
	}
	if len(m.Values) > 0 {
		for _, s := range m.Values {
			dAtA[i] = 0x12
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
	return i, nil
}

func encodeVarintHttpgrpc(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *HTTPRequest) Size() (n int) {
	var l int
	_ = l
	l = len(m.Method)
	if l > 0 {
		n += 1 + l + sovHttpgrpc(uint64(l))
	}
	l = len(m.Url)
	if l > 0 {
		n += 1 + l + sovHttpgrpc(uint64(l))
	}
	if len(m.Headers) > 0 {
		for _, e := range m.Headers {
			l = e.Size()
			n += 1 + l + sovHttpgrpc(uint64(l))
		}
	}
	l = len(m.Body)
	if l > 0 {
		n += 1 + l + sovHttpgrpc(uint64(l))
	}
	return n
}

func (m *HTTPResponse) Size() (n int) {
	var l int
	_ = l
	if m.Code != 0 {
		n += 1 + sovHttpgrpc(uint64(m.Code))
	}
	if len(m.Headers) > 0 {
		for _, e := range m.Headers {
			l = e.Size()
			n += 1 + l + sovHttpgrpc(uint64(l))
		}
	}
	l = len(m.Body)
	if l > 0 {
		n += 1 + l + sovHttpgrpc(uint64(l))
	}
	return n
}

func (m *Header) Size() (n int) {
	var l int
	_ = l
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovHttpgrpc(uint64(l))
	}
	if len(m.Values) > 0 {
		for _, s := range m.Values {
			l = len(s)
			n += 1 + l + sovHttpgrpc(uint64(l))
		}
	}
	return n
}

func sovHttpgrpc(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozHttpgrpc(x uint64) (n int) {
	return sovHttpgrpc(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *HTTPRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&HTTPRequest{`,
		`Method:` + fmt.Sprintf("%v", this.Method) + `,`,
		`Url:` + fmt.Sprintf("%v", this.Url) + `,`,
		`Headers:` + strings.Replace(fmt.Sprintf("%v", this.Headers), "Header", "Header", 1) + `,`,
		`Body:` + fmt.Sprintf("%v", this.Body) + `,`,
		`}`,
	}, "")
	return s
}
func (this *HTTPResponse) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&HTTPResponse{`,
		`Code:` + fmt.Sprintf("%v", this.Code) + `,`,
		`Headers:` + strings.Replace(fmt.Sprintf("%v", this.Headers), "Header", "Header", 1) + `,`,
		`Body:` + fmt.Sprintf("%v", this.Body) + `,`,
		`}`,
	}, "")
	return s
}
func (this *Header) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Header{`,
		`Key:` + fmt.Sprintf("%v", this.Key) + `,`,
		`Values:` + fmt.Sprintf("%v", this.Values) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringHttpgrpc(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *HTTPRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHttpgrpc
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
			return fmt.Errorf("proto: HTTPRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HTTPRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Method", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHttpgrpc
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
				return ErrInvalidLengthHttpgrpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Method = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Url", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHttpgrpc
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
				return ErrInvalidLengthHttpgrpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Url = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Headers", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHttpgrpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthHttpgrpc
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Headers = append(m.Headers, &Header{})
			if err := m.Headers[len(m.Headers)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Body", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHttpgrpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthHttpgrpc
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Body = append(m.Body[:0], dAtA[iNdEx:postIndex]...)
			if m.Body == nil {
				m.Body = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipHttpgrpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHttpgrpc
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
func (m *HTTPResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHttpgrpc
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
			return fmt.Errorf("proto: HTTPResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HTTPResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Code", wireType)
			}
			m.Code = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHttpgrpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Code |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Headers", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHttpgrpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthHttpgrpc
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Headers = append(m.Headers, &Header{})
			if err := m.Headers[len(m.Headers)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Body", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHttpgrpc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthHttpgrpc
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Body = append(m.Body[:0], dAtA[iNdEx:postIndex]...)
			if m.Body == nil {
				m.Body = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipHttpgrpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHttpgrpc
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
func (m *Header) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHttpgrpc
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
			return fmt.Errorf("proto: Header: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Header: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHttpgrpc
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
				return ErrInvalidLengthHttpgrpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Values", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHttpgrpc
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
				return ErrInvalidLengthHttpgrpc
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Values = append(m.Values, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipHttpgrpc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHttpgrpc
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
func skipHttpgrpc(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowHttpgrpc
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
					return 0, ErrIntOverflowHttpgrpc
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
					return 0, ErrIntOverflowHttpgrpc
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
				return 0, ErrInvalidLengthHttpgrpc
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowHttpgrpc
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
				next, err := skipHttpgrpc(dAtA[start:])
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
	ErrInvalidLengthHttpgrpc = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowHttpgrpc   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("github.com/weaveworks/common/httpgrpc/httpgrpc.proto", fileDescriptorHttpgrpc)
}

var fileDescriptorHttpgrpc = []byte{
	// 319 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x49, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x2f, 0x4f, 0x4d, 0x2c, 0x4b, 0x2d, 0xcf, 0x2f, 0xca,
	0x2e, 0xd6, 0x4f, 0xce, 0xcf, 0xcd, 0xcd, 0xcf, 0xd3, 0xcf, 0x28, 0x29, 0x29, 0x48, 0x2f, 0x2a,
	0x48, 0x86, 0x33, 0xf4, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x38, 0x60, 0x7c, 0xa5, 0x72, 0x2e,
	0x6e, 0x8f, 0x90, 0x90, 0x80, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x21, 0x31, 0x2e, 0xb6,
	0xdc, 0xd4, 0x92, 0x8c, 0xfc, 0x14, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x28, 0x4f, 0x48,
	0x80, 0x8b, 0xb9, 0xb4, 0x28, 0x47, 0x82, 0x09, 0x2c, 0x08, 0x62, 0x0a, 0x69, 0x71, 0xb1, 0x67,
	0xa4, 0x26, 0xa6, 0xa4, 0x16, 0x15, 0x4b, 0x30, 0x2b, 0x30, 0x6b, 0x70, 0x1b, 0x09, 0xe8, 0xc1,
	0x2d, 0xf1, 0x00, 0x4b, 0x04, 0xc1, 0x14, 0x08, 0x09, 0x71, 0xb1, 0x24, 0xe5, 0xa7, 0x54, 0x4a,
	0xb0, 0x28, 0x30, 0x6a, 0xf0, 0x04, 0x81, 0xd9, 0x4a, 0x49, 0x5c, 0x3c, 0x10, 0x8b, 0x8b, 0x0b,
	0xf2, 0xf3, 0x8a, 0x53, 0x41, 0x6a, 0x9c, 0xf3, 0x53, 0x52, 0xc1, 0xf6, 0xb2, 0x06, 0xb1, 0x24,
	0xe7, 0xa7, 0xa4, 0x22, 0xdb, 0xc1, 0x44, 0xac, 0x1d, 0xcc, 0x48, 0x76, 0x18, 0x71, 0xb1, 0x41,
	0x94, 0x81, 0xdc, 0x9f, 0x9d, 0x5a, 0x09, 0xf5, 0x14, 0x88, 0x09, 0xf2, 0x69, 0x59, 0x62, 0x4e,
	0x69, 0x2a, 0xc4, 0x68, 0xce, 0x20, 0x28, 0xcf, 0xc8, 0x91, 0x8b, 0x05, 0xe4, 0x2e, 0x21, 0x4b,
	0x2e, 0x36, 0x8f, 0xc4, 0xbc, 0x94, 0x9c, 0x54, 0x21, 0x51, 0x24, 0x4b, 0x11, 0x41, 0x25, 0x25,
	0x86, 0x2e, 0x0c, 0xf1, 0x88, 0x12, 0x83, 0x53, 0xf0, 0x85, 0x87, 0x72, 0x0c, 0x37, 0x1e, 0xca,
	0x31, 0x7c, 0x78, 0x28, 0xc7, 0xd8, 0xf0, 0x48, 0x8e, 0x71, 0xc5, 0x23, 0x39, 0xc6, 0x13, 0x8f,
	0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e, 0xf1, 0xc5, 0x23, 0x39, 0x86, 0x0f,
	0x8f, 0xe4, 0x18, 0x27, 0x3c, 0x96, 0x63, 0x88, 0x52, 0x25, 0x2a, 0x06, 0x93, 0xd8, 0xc0, 0x31,
	0x67, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x49, 0xec, 0x86, 0xd6, 0xf1, 0x01, 0x00, 0x00,
}
