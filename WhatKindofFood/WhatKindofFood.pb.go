// Code generated by protoc-gen-go. DO NOT EDIT.
// source: WhatKindofFood.proto

package WhatKindofFood

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type FoodTypeRequest struct {
	Foodtype             string   `protobuf:"bytes,1,opt,name=foodtype,proto3" json:"foodtype,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FoodTypeRequest) Reset()         { *m = FoodTypeRequest{} }
func (m *FoodTypeRequest) String() string { return proto.CompactTextString(m) }
func (*FoodTypeRequest) ProtoMessage()    {}
func (*FoodTypeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_36b9b5f7289b099c, []int{0}
}

func (m *FoodTypeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FoodTypeRequest.Unmarshal(m, b)
}
func (m *FoodTypeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FoodTypeRequest.Marshal(b, m, deterministic)
}
func (m *FoodTypeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FoodTypeRequest.Merge(m, src)
}
func (m *FoodTypeRequest) XXX_Size() int {
	return xxx_messageInfo_FoodTypeRequest.Size(m)
}
func (m *FoodTypeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FoodTypeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FoodTypeRequest proto.InternalMessageInfo

func (m *FoodTypeRequest) GetFoodtype() string {
	if m != nil {
		return m.Foodtype
	}
	return ""
}

type FoodTypeResp struct {
	Foodname             string   `protobuf:"bytes,1,opt,name=foodname,proto3" json:"foodname,omitempty"`
	Foodprice            float32  `protobuf:"fixed32,2,opt,name=foodprice,proto3" json:"foodprice,omitempty"`
	Fooddesc             string   `protobuf:"bytes,3,opt,name=fooddesc,proto3" json:"fooddesc,omitempty"`
	Vendor               string   `protobuf:"bytes,4,opt,name=vendor,proto3" json:"vendor,omitempty"`
	Foodloc              string   `protobuf:"bytes,5,opt,name=foodloc,proto3" json:"foodloc,omitempty"`
	Foodtype             string   `protobuf:"bytes,6,opt,name=foodtype,proto3" json:"foodtype,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FoodTypeResp) Reset()         { *m = FoodTypeResp{} }
func (m *FoodTypeResp) String() string { return proto.CompactTextString(m) }
func (*FoodTypeResp) ProtoMessage()    {}
func (*FoodTypeResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_36b9b5f7289b099c, []int{1}
}

func (m *FoodTypeResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FoodTypeResp.Unmarshal(m, b)
}
func (m *FoodTypeResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FoodTypeResp.Marshal(b, m, deterministic)
}
func (m *FoodTypeResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FoodTypeResp.Merge(m, src)
}
func (m *FoodTypeResp) XXX_Size() int {
	return xxx_messageInfo_FoodTypeResp.Size(m)
}
func (m *FoodTypeResp) XXX_DiscardUnknown() {
	xxx_messageInfo_FoodTypeResp.DiscardUnknown(m)
}

var xxx_messageInfo_FoodTypeResp proto.InternalMessageInfo

func (m *FoodTypeResp) GetFoodname() string {
	if m != nil {
		return m.Foodname
	}
	return ""
}

func (m *FoodTypeResp) GetFoodprice() float32 {
	if m != nil {
		return m.Foodprice
	}
	return 0
}

func (m *FoodTypeResp) GetFooddesc() string {
	if m != nil {
		return m.Fooddesc
	}
	return ""
}

func (m *FoodTypeResp) GetVendor() string {
	if m != nil {
		return m.Vendor
	}
	return ""
}

func (m *FoodTypeResp) GetFoodloc() string {
	if m != nil {
		return m.Foodloc
	}
	return ""
}

func (m *FoodTypeResp) GetFoodtype() string {
	if m != nil {
		return m.Foodtype
	}
	return ""
}

func init() {
	proto.RegisterType((*FoodTypeRequest)(nil), "WhatKindofFood.FoodTypeRequest")
	proto.RegisterType((*FoodTypeResp)(nil), "WhatKindofFood.FoodTypeResp")
}

func init() { proto.RegisterFile("WhatKindofFood.proto", fileDescriptor_36b9b5f7289b099c) }

var fileDescriptor_36b9b5f7289b099c = []byte{
	// 221 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x09, 0xcf, 0x48, 0x2c,
	0xf1, 0xce, 0xcc, 0x4b, 0xc9, 0x4f, 0x73, 0xcb, 0xcf, 0x4f, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0xe2, 0x43, 0x15, 0x55, 0xd2, 0xe5, 0xe2, 0x07, 0xd1, 0x21, 0x95, 0x05, 0xa9, 0x41, 0xa9,
	0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x42, 0x52, 0x5c, 0x1c, 0x69, 0xf9, 0xf9, 0x29, 0x25, 0x95, 0x05,
	0xa9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x70, 0xbe, 0xd2, 0x26, 0x46, 0x2e, 0x1e, 0x84,
	0xfa, 0xe2, 0x02, 0x98, 0xe2, 0xbc, 0xc4, 0x5c, 0x14, 0xc5, 0x20, 0xbe, 0x90, 0x0c, 0x17, 0x27,
	0x88, 0x5d, 0x50, 0x94, 0x99, 0x9c, 0x2a, 0xc1, 0xa4, 0xc0, 0xa8, 0xc1, 0x14, 0x84, 0x10, 0x80,
	0xe9, 0x4c, 0x49, 0x2d, 0x4e, 0x96, 0x60, 0x46, 0xe8, 0x04, 0xf1, 0x85, 0xc4, 0xb8, 0xd8, 0xca,
	0x52, 0xf3, 0x52, 0xf2, 0x8b, 0x24, 0x58, 0xc0, 0x32, 0x50, 0x9e, 0x90, 0x04, 0x17, 0x3b, 0x48,
	0x4d, 0x4e, 0x7e, 0xb2, 0x04, 0x2b, 0x58, 0x02, 0xc6, 0x45, 0x71, 0x34, 0x1b, 0xaa, 0xa3, 0x8d,
	0x62, 0xb9, 0x78, 0x41, 0x6e, 0xf6, 0xcc, 0xf3, 0xad, 0x74, 0x4a, 0xcd, 0xc9, 0xa9, 0x14, 0xf2,
	0xe1, 0x62, 0x77, 0x4f, 0x2d, 0x01, 0x89, 0x09, 0xc9, 0xeb, 0xa1, 0x05, 0x13, 0x5a, 0x68, 0x48,
	0xc9, 0xe0, 0x56, 0x50, 0x5c, 0xa0, 0xc4, 0x60, 0xc0, 0x98, 0xc4, 0x06, 0x0e, 0x59, 0x63, 0x40,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x77, 0x4b, 0x51, 0x01, 0x71, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// FoodInMyBellyClient is the client API for FoodInMyBelly service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FoodInMyBellyClient interface {
	GetFood(ctx context.Context, in *FoodTypeRequest, opts ...grpc.CallOption) (FoodInMyBelly_GetFoodClient, error)
}

type foodInMyBellyClient struct {
	cc *grpc.ClientConn
}

func NewFoodInMyBellyClient(cc *grpc.ClientConn) FoodInMyBellyClient {
	return &foodInMyBellyClient{cc}
}

func (c *foodInMyBellyClient) GetFood(ctx context.Context, in *FoodTypeRequest, opts ...grpc.CallOption) (FoodInMyBelly_GetFoodClient, error) {
	stream, err := c.cc.NewStream(ctx, &_FoodInMyBelly_serviceDesc.Streams[0], "/WhatKindofFood.FoodInMyBelly/GetFood", opts...)
	if err != nil {
		return nil, err
	}
	x := &foodInMyBellyGetFoodClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type FoodInMyBelly_GetFoodClient interface {
	Recv() (*FoodTypeResp, error)
	grpc.ClientStream
}

type foodInMyBellyGetFoodClient struct {
	grpc.ClientStream
}

func (x *foodInMyBellyGetFoodClient) Recv() (*FoodTypeResp, error) {
	m := new(FoodTypeResp)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// FoodInMyBellyServer is the server API for FoodInMyBelly service.
type FoodInMyBellyServer interface {
	GetFood(*FoodTypeRequest, FoodInMyBelly_GetFoodServer) error
}

func RegisterFoodInMyBellyServer(s *grpc.Server, srv FoodInMyBellyServer) {
	s.RegisterService(&_FoodInMyBelly_serviceDesc, srv)
}

func _FoodInMyBelly_GetFood_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(FoodTypeRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FoodInMyBellyServer).GetFood(m, &foodInMyBellyGetFoodServer{stream})
}

type FoodInMyBelly_GetFoodServer interface {
	Send(*FoodTypeResp) error
	grpc.ServerStream
}

type foodInMyBellyGetFoodServer struct {
	grpc.ServerStream
}

func (x *foodInMyBellyGetFoodServer) Send(m *FoodTypeResp) error {
	return x.ServerStream.SendMsg(m)
}

var _FoodInMyBelly_serviceDesc = grpc.ServiceDesc{
	ServiceName: "WhatKindofFood.FoodInMyBelly",
	HandlerType: (*FoodInMyBellyServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetFood",
			Handler:       _FoodInMyBelly_GetFood_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "WhatKindofFood.proto",
}
