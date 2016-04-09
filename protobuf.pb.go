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
	URL         string             `protobuf:"bytes,1,opt,name=URL,json=uRL" json:"URL,omitempty"`
	Method      string             `protobuf:"bytes,2,opt,name=Method,json=method" json:"Method,omitempty"`
	RemoteAddr  string             `protobuf:"bytes,3,opt,name=RemoteAddr,json=remoteAddr" json:"RemoteAddr,omitempty"`
	Body        []byte             `protobuf:"bytes,4,opt,name=Body,json=body,proto3" json:"Body,omitempty"`
	Form        map[string]*Values `protobuf:"bytes,5,rep,name=Form,json=form" json:"Form,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Header      map[string]*Values `protobuf:"bytes,6,rep,name=Header,json=header" json:"Header,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	WebSocketID string             `protobuf:"bytes,7,opt,name=WebSocketID,json=webSocketID" json:"WebSocketID,omitempty"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type Response struct {
	StatusCode int32              `protobuf:"varint,1,opt,name=StatusCode,json=statusCode" json:"StatusCode,omitempty"`
	Header     map[string]*Values `protobuf:"bytes,2,rep,name=Header,json=header" json:"Header,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Body       []byte             `protobuf:"bytes,3,opt,name=Body,json=body,proto3" json:"Body,omitempty"`
	DoUpgrade  bool               `protobuf:"varint,4,opt,name=DoUpgrade,json=doUpgrade" json:"DoUpgrade,omitempty"`
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
	// 349 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x92, 0xbf, 0x4e, 0xc3, 0x30,
	0x10, 0xc6, 0x95, 0x3f, 0x75, 0x9b, 0x0b, 0x20, 0x64, 0x01, 0xb2, 0x2a, 0x40, 0x55, 0x07, 0xc4,
	0x00, 0x19, 0xca, 0x82, 0x98, 0xa0, 0x14, 0x54, 0xa4, 0xb2, 0xb8, 0x2a, 0x48, 0x6c, 0x09, 0x76,
	0xa9, 0x54, 0x5a, 0x07, 0xc7, 0x01, 0xf2, 0x94, 0x3c, 0x05, 0xef, 0x81, 0xed, 0x90, 0xd0, 0x81,
	0x09, 0xb1, 0x7d, 0xfe, 0xbe, 0x8b, 0x2f, 0xf7, 0xf3, 0xc1, 0x46, 0x2a, 0x85, 0x12, 0x49, 0x3e,
	0x8d, 0xac, 0xe8, 0xb6, 0x01, 0xdd, 0xc5, 0xcf, 0x39, 0xcf, 0xf0, 0x26, 0x78, 0xb1, 0x94, 0xc4,
	0xe9, 0x78, 0x87, 0x01, 0x35, 0xb2, 0xfb, 0xe9, 0x42, 0x93, 0xf2, 0x17, 0x1d, 0x2a, 0x93, 0x4e,
	0xe8, 0x48, 0xa7, 0x8e, 0x49, 0x73, 0x3a, 0xc2, 0x3b, 0x80, 0x6e, 0xb9, 0x9a, 0x09, 0x46, 0x5c,
	0x6b, 0xa2, 0x85, 0x3d, 0xe1, 0x7d, 0x00, 0xca, 0x17, 0x42, 0xf1, 0x0b, 0xc6, 0x24, 0xf1, 0x6c,
	0x06, 0xb2, 0x76, 0x30, 0x06, 0xbf, 0x2f, 0x58, 0x41, 0x7c, 0x9d, 0xac, 0x51, 0x3f, 0xd1, 0x1a,
	0x1f, 0x80, 0x7f, 0x2d, 0xe4, 0x82, 0x34, 0x74, 0xf3, 0xb0, 0x87, 0xa3, 0xef, 0xae, 0x91, 0x31,
	0xaf, 0x96, 0x4a, 0x16, 0xd4, 0x9f, 0x6a, 0x89, 0x8f, 0x00, 0x0d, 0x79, 0xcc, 0xb8, 0x24, 0xc8,
	0x56, 0x6e, 0xd5, 0x95, 0xa5, 0x5d, 0xd6, 0xa2, 0x99, 0x3d, 0xe0, 0x0e, 0x84, 0xf7, 0x3c, 0x19,
	0x8b, 0xc7, 0x39, 0x57, 0x37, 0x03, 0xd2, 0xb4, 0xbf, 0x12, 0xbe, 0xfd, 0x58, 0xed, 0x73, 0x08,
	0xea, 0x16, 0x66, 0xc4, 0x39, 0x2f, 0xaa, 0x11, 0xb5, 0xc4, 0x7b, 0xd0, 0x78, 0x35, 0x70, 0xec,
	0x84, 0x61, 0xaf, 0x19, 0x95, 0xa8, 0x68, 0xe9, 0x9e, 0xb9, 0xa7, 0x4e, 0xbb, 0x0f, 0xe1, 0x4a,
	0xeb, 0x3f, 0xdd, 0xd1, 0xfd, 0x70, 0xa0, 0x45, 0x79, 0x96, 0x8a, 0x65, 0xc6, 0x0d, 0xbe, 0xb1,
	0x8a, 0x55, 0x9e, 0x5d, 0x0a, 0xc6, 0xed, 0x45, 0x0d, 0x0a, 0x59, 0xed, 0xe0, 0xe3, 0x1a, 0x81,
	0x6b, 0x11, 0x6c, 0x47, 0xd5, 0xa7, 0xbf, 0x32, 0xa8, 0x68, 0x7b, 0x2b, 0xb4, 0x77, 0x21, 0x18,
	0x88, 0x49, 0xfa, 0x24, 0x75, 0x85, 0x7d, 0x86, 0x16, 0x0d, 0x58, 0x65, 0xfc, 0xc7, 0x44, 0xfd,
	0xf5, 0xa1, 0xfb, 0x10, 0x2c, 0x63, 0x95, 0xe9, 0x25, 0x7b, 0x2f, 0x12, 0x64, 0x77, 0xed, 0xe4,
	0x2b, 0x00, 0x00, 0xff, 0xff, 0xea, 0xc6, 0x72, 0xd5, 0x7d, 0x02, 0x00, 0x00,
}
