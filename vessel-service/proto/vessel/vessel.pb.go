// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/vessel/vessel.proto

package go_micro_src_vessel

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "golang.org/x/net/context"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Vessel struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Capacity             int32    `protobuf:"varint,2,opt,name=capacity" json:"capacity,omitempty"`
	MaxWeight            int32    `protobuf:"varint,3,opt,name=max_weight,json=maxWeight" json:"max_weight,omitempty"`
	Name                 string   `protobuf:"bytes,4,opt,name=name" json:"name,omitempty"`
	Available            bool     `protobuf:"varint,5,opt,name=available" json:"available,omitempty"`
	OwnerId              string   `protobuf:"bytes,6,opt,name=owner_id,json=ownerId" json:"owner_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Vessel) Reset()         { *m = Vessel{} }
func (m *Vessel) String() string { return proto.CompactTextString(m) }
func (*Vessel) ProtoMessage()    {}
func (*Vessel) Descriptor() ([]byte, []int) {
	return fileDescriptor_vessel_03ff8be4871d373c, []int{0}
}
func (m *Vessel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Vessel.Unmarshal(m, b)
}
func (m *Vessel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Vessel.Marshal(b, m, deterministic)
}
func (dst *Vessel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Vessel.Merge(dst, src)
}
func (m *Vessel) XXX_Size() int {
	return xxx_messageInfo_Vessel.Size(m)
}
func (m *Vessel) XXX_DiscardUnknown() {
	xxx_messageInfo_Vessel.DiscardUnknown(m)
}

var xxx_messageInfo_Vessel proto.InternalMessageInfo

func (m *Vessel) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Vessel) GetCapacity() int32 {
	if m != nil {
		return m.Capacity
	}
	return 0
}

func (m *Vessel) GetMaxWeight() int32 {
	if m != nil {
		return m.MaxWeight
	}
	return 0
}

func (m *Vessel) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Vessel) GetAvailable() bool {
	if m != nil {
		return m.Available
	}
	return false
}

func (m *Vessel) GetOwnerId() string {
	if m != nil {
		return m.OwnerId
	}
	return ""
}

type Specification struct {
	Capacity             int32    `protobuf:"varint,1,opt,name=capacity" json:"capacity,omitempty"`
	MaxWeight            int32    `protobuf:"varint,2,opt,name=max_weight,json=maxWeight" json:"max_weight,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Specification) Reset()         { *m = Specification{} }
func (m *Specification) String() string { return proto.CompactTextString(m) }
func (*Specification) ProtoMessage()    {}
func (*Specification) Descriptor() ([]byte, []int) {
	return fileDescriptor_vessel_03ff8be4871d373c, []int{1}
}
func (m *Specification) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Specification.Unmarshal(m, b)
}
func (m *Specification) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Specification.Marshal(b, m, deterministic)
}
func (dst *Specification) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Specification.Merge(dst, src)
}
func (m *Specification) XXX_Size() int {
	return xxx_messageInfo_Specification.Size(m)
}
func (m *Specification) XXX_DiscardUnknown() {
	xxx_messageInfo_Specification.DiscardUnknown(m)
}

var xxx_messageInfo_Specification proto.InternalMessageInfo

func (m *Specification) GetCapacity() int32 {
	if m != nil {
		return m.Capacity
	}
	return 0
}

func (m *Specification) GetMaxWeight() int32 {
	if m != nil {
		return m.MaxWeight
	}
	return 0
}

type Response struct {
	Vessel               *Vessel   `protobuf:"bytes,1,opt,name=vessel" json:"vessel,omitempty"`
	Vessels              []*Vessel `protobuf:"bytes,2,rep,name=vessels" json:"vessels,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_vessel_03ff8be4871d373c, []int{2}
}
func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (dst *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(dst, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetVessel() *Vessel {
	if m != nil {
		return m.Vessel
	}
	return nil
}

func (m *Response) GetVessels() []*Vessel {
	if m != nil {
		return m.Vessels
	}
	return nil
}

func init() {
	proto.RegisterType((*Vessel)(nil), "go.micro.src.vessel.Vessel")
	proto.RegisterType((*Specification)(nil), "go.micro.src.vessel.Specification")
	proto.RegisterType((*Response)(nil), "go.micro.src.vessel.Response")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for VesselService service

type VesselServiceClient interface {
	FindAvailable(ctx context.Context, in *Specification, opts ...client.CallOption) (*Response, error)
}

type vesselServiceClient struct {
	c           client.Client
	serviceName string
}

func NewVesselServiceClient(serviceName string, c client.Client) VesselServiceClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "go.micro.src.vessel"
	}
	return &vesselServiceClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *vesselServiceClient) FindAvailable(ctx context.Context, in *Specification, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "VesselService.FindAvailable", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for VesselService service

type VesselServiceHandler interface {
	FindAvailable(context.Context, *Specification, *Response) error
}

func RegisterVesselServiceHandler(s server.Server, hdlr VesselServiceHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&VesselService{hdlr}, opts...))
}

type VesselService struct {
	VesselServiceHandler
}

func (h *VesselService) FindAvailable(ctx context.Context, in *Specification, out *Response) error {
	return h.VesselServiceHandler.FindAvailable(ctx, in, out)
}

func init() { proto.RegisterFile("proto/vessel/vessel.proto", fileDescriptor_vessel_03ff8be4871d373c) }

var fileDescriptor_vessel_03ff8be4871d373c = []byte{
	// 289 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x51, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x75, 0xd3, 0x36, 0x4d, 0x47, 0xe2, 0x61, 0xbd, 0x6c, 0xab, 0x85, 0x90, 0x53, 0x4e, 0x2b,
	0xb4, 0xf8, 0x01, 0x5e, 0x04, 0x3d, 0xa6, 0xa2, 0xc7, 0xb2, 0xdd, 0x8c, 0x75, 0x20, 0xc9, 0x86,
	0x6c, 0x48, 0xeb, 0xdf, 0xf8, 0xa9, 0xc2, 0xa6, 0xad, 0x54, 0x82, 0x9e, 0x76, 0xe6, 0xcd, 0x9b,
	0xc7, 0x9b, 0xb7, 0x30, 0xad, 0x6a, 0xd3, 0x98, 0xbb, 0x16, 0xad, 0xc5, 0xfc, 0xf0, 0x48, 0x87,
	0xf1, 0xeb, 0xad, 0x91, 0x05, 0xe9, 0xda, 0x48, 0x5b, 0x6b, 0xd9, 0x8d, 0xe2, 0x2f, 0x06, 0xfe,
	0xab, 0x2b, 0xf9, 0x15, 0x78, 0x94, 0x09, 0x16, 0xb1, 0x64, 0x92, 0x7a, 0x94, 0xf1, 0x19, 0x04,
	0x5a, 0x55, 0x4a, 0x53, 0xf3, 0x29, 0xbc, 0x88, 0x25, 0xa3, 0xf4, 0xd4, 0xf3, 0x39, 0x40, 0xa1,
	0xf6, 0xeb, 0x1d, 0xd2, 0xf6, 0xa3, 0x11, 0x03, 0x37, 0x9d, 0x14, 0x6a, 0xff, 0xe6, 0x00, 0xce,
	0x61, 0x58, 0xaa, 0x02, 0xc5, 0xd0, 0x89, 0xb9, 0x9a, 0xdf, 0xc2, 0x44, 0xb5, 0x8a, 0x72, 0xb5,
	0xc9, 0x51, 0x8c, 0x22, 0x96, 0x04, 0xe9, 0x0f, 0xc0, 0xa7, 0x10, 0x98, 0x5d, 0x89, 0xf5, 0x9a,
	0x32, 0xe1, 0xbb, 0xad, 0xb1, 0xeb, 0x9f, 0xb2, 0xf8, 0x19, 0xc2, 0x55, 0x85, 0x9a, 0xde, 0x49,
	0xab, 0x86, 0x4c, 0x79, 0x66, 0x8c, 0xfd, 0x69, 0xcc, 0xfb, 0x65, 0x2c, 0x6e, 0x21, 0x48, 0xd1,
	0x56, 0xa6, 0xb4, 0xc8, 0x97, 0xe0, 0x77, 0x21, 0x38, 0x91, 0xcb, 0xc5, 0x8d, 0xec, 0x09, 0x48,
	0x76, 0xe1, 0xa4, 0x07, 0x2a, 0xbf, 0x87, 0x71, 0x57, 0x59, 0xe1, 0x45, 0x83, 0xff, 0xb6, 0x8e,
	0xdc, 0x05, 0x42, 0xd8, 0x41, 0x2b, 0xac, 0x5b, 0xd2, 0xc8, 0x5f, 0x20, 0x7c, 0xa4, 0x32, 0x7b,
	0x38, 0x05, 0x10, 0xf7, 0xea, 0x9c, 0x1d, 0x3e, 0x9b, 0xf7, 0x72, 0x8e, 0x07, 0xc5, 0x17, 0x1b,
	0xdf, 0xfd, 0xf4, 0xf2, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x71, 0x48, 0x62, 0x2a, 0x06, 0x02, 0x00,
	0x00,
}
