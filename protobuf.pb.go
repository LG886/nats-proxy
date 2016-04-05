// Code generated by protoc-gen-go.
// source: protobuf.proto
// DO NOT EDIT!

/*
Package natsproxy is a generated protocol buffer package.

It is generated from these files:
	protobuf.proto

It has these top-level messages:
	Values
	Request
	Response
*/
package natsproxy

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type Values struct {
	Arr []string `protobuf:"bytes,1,rep,name=arr" json:"arr,omitempty"`
}

func (m *Values) Reset()                    { *m = Values{} }
func (m *Values) String() string            { return proto.CompactTextString(m) }
func (*Values) ProtoMessage()               {}
func (*Values) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Request struct {
	URL        string             `protobuf:"bytes,1,opt,name=URL,json=uRL" json:"URL,omitempty"`
	Method     string             `protobuf:"bytes,2,opt,name=Method,json=method" json:"Method,omitempty"`
	RemoteAddr string             `protobuf:"bytes,3,opt,name=RemoteAddr,json=remoteAddr" json:"RemoteAddr,omitempty"`
	Body       []byte             `protobuf:"bytes,4,opt,name=Body,json=body,proto3" json:"Body,omitempty"`
	Form       map[string]*Values `protobuf:"bytes,5,rep,name=Form,json=form" json:"Form,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Header     map[string]*Values `protobuf:"bytes,6,rep,name=Header,json=header" json:"Header,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type Response struct {
	StatusCode int32              `protobuf:"varint,1,opt,name=StatusCode,json=statusCode" json:"StatusCode,omitempty"`
	Header     map[string]*Values `protobuf:"bytes,2,rep,name=Header,json=header" json:"Header,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Body       []byte             `protobuf:"bytes,3,opt,name=Body,json=body,proto3" json:"Body,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func init() {
	proto.RegisterType((*Values)(nil), "Values")
	proto.RegisterType((*Request)(nil), "Request")
	proto.RegisterType((*Response)(nil), "Response")
}

var fileDescriptor0 = []byte{
	// 311 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x52, 0xcd, 0x4a, 0xc4, 0x30,
	0x18, 0xa4, 0x3f, 0x9b, 0xb5, 0x5f, 0x55, 0x24, 0xa8, 0x84, 0x82, 0x22, 0x3d, 0x88, 0x07, 0xed,
	0x61, 0xbd, 0x88, 0x27, 0xad, 0x28, 0x7b, 0x58, 0x2f, 0x11, 0x3d, 0x78, 0x6b, 0x49, 0x96, 0x05,
	0x6d, 0xb3, 0x26, 0xa9, 0xd8, 0xc7, 0xf2, 0x19, 0x7c, 0x31, 0x93, 0xd4, 0x8d, 0x1e, 0x3c, 0x89,
	0xb7, 0xc9, 0xcc, 0x97, 0x49, 0x66, 0xf8, 0x60, 0x73, 0x29, 0x85, 0x16, 0x75, 0x37, 0x2f, 0x1c,
	0xc8, 0x33, 0x40, 0x0f, 0xd5, 0x73, 0xc7, 0x15, 0xde, 0x82, 0xa8, 0x92, 0x92, 0x04, 0x07, 0xd1,
	0x51, 0x42, 0x2d, 0xcc, 0x3f, 0x42, 0x18, 0x53, 0xfe, 0x62, 0x44, 0x6d, 0xd5, 0x7b, 0x3a, 0x33,
	0x6a, 0x60, 0xd5, 0x8e, 0xce, 0xf0, 0x2e, 0xa0, 0x5b, 0xae, 0x17, 0x82, 0x91, 0xd0, 0x91, 0xa8,
	0x71, 0x27, 0xbc, 0x0f, 0x40, 0x79, 0x23, 0x34, 0xbf, 0x64, 0x4c, 0x92, 0xc8, 0x69, 0x20, 0x3d,
	0x83, 0x31, 0xc4, 0xa5, 0x60, 0x3d, 0x89, 0x8d, 0xb2, 0x4e, 0xe3, 0xda, 0x60, 0x7c, 0x08, 0xf1,
	0x8d, 0x90, 0x0d, 0x19, 0x99, 0xc7, 0xd3, 0x09, 0x2e, 0xbe, 0x5e, 0x2d, 0x2c, 0x79, 0xdd, 0x6a,
	0xd9, 0xd3, 0x78, 0x6e, 0x20, 0x3e, 0x06, 0x34, 0xe5, 0x15, 0xe3, 0x92, 0x20, 0x37, 0xb9, 0xed,
	0x27, 0x07, 0x7a, 0x98, 0x45, 0x0b, 0x77, 0xc8, 0x2e, 0x20, 0xf1, 0x06, 0x36, 0xc0, 0x13, 0xef,
	0x57, 0x01, 0x0c, 0xc4, 0x7b, 0x30, 0x7a, 0xb5, 0xd1, 0xdd, 0xff, 0xd3, 0xc9, 0xb8, 0x18, 0x8a,
	0xa0, 0x03, 0x7b, 0x1e, 0x9e, 0x05, 0x59, 0x09, 0xe9, 0x0f, 0xe3, 0x3f, 0x79, 0xe4, 0xef, 0x01,
	0xac, 0x51, 0xae, 0x96, 0xa2, 0x55, 0xdc, 0x96, 0x73, 0xa7, 0x2b, 0xdd, 0xa9, 0x2b, 0xc1, 0xb8,
	0x33, 0x1a, 0x51, 0x50, 0x9e, 0xc1, 0x27, 0x3e, 0x60, 0xe8, 0x02, 0xee, 0x14, 0xab, 0xab, 0xbf,
	0x25, 0xf4, 0x5d, 0x46, 0xdf, 0x5d, 0xfe, 0xc7, 0x9f, 0xcb, 0x8d, 0x69, 0xf8, 0x98, 0xb4, 0x95,
	0x56, 0x66, 0x49, 0xde, 0xfa, 0x1a, 0xb9, 0x5d, 0x39, 0xfd, 0x0c, 0x00, 0x00, 0xff, 0xff, 0x12,
	0x62, 0xec, 0x27, 0x3d, 0x02, 0x00, 0x00,
}
