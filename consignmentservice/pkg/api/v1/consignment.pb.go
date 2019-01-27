// Code generated by protoc-gen-go. DO NOT EDIT.
// source: consignment.proto

package consignment_proto

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

type ConsignmentRequest struct {
	Api                  string       `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	Consignment          *Consignment `protobuf:"bytes,2,opt,name=consignment,proto3" json:"consignment,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ConsignmentRequest) Reset()         { *m = ConsignmentRequest{} }
func (m *ConsignmentRequest) String() string { return proto.CompactTextString(m) }
func (*ConsignmentRequest) ProtoMessage()    {}
func (*ConsignmentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3804bf87090b51a9, []int{0}
}

func (m *ConsignmentRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConsignmentRequest.Unmarshal(m, b)
}
func (m *ConsignmentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConsignmentRequest.Marshal(b, m, deterministic)
}
func (m *ConsignmentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConsignmentRequest.Merge(m, src)
}
func (m *ConsignmentRequest) XXX_Size() int {
	return xxx_messageInfo_ConsignmentRequest.Size(m)
}
func (m *ConsignmentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ConsignmentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ConsignmentRequest proto.InternalMessageInfo

func (m *ConsignmentRequest) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *ConsignmentRequest) GetConsignment() *Consignment {
	if m != nil {
		return m.Consignment
	}
	return nil
}

type Consignment struct {
	Id                   int64        `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Description          string       `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Weight               int64        `protobuf:"varint,3,opt,name=weight,proto3" json:"weight,omitempty"`
	Containers           []*Container `protobuf:"bytes,4,rep,name=containers,proto3" json:"containers,omitempty"`
	VesselId             string       `protobuf:"bytes,5,opt,name=vessel_id,json=vesselId,proto3" json:"vessel_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Consignment) Reset()         { *m = Consignment{} }
func (m *Consignment) String() string { return proto.CompactTextString(m) }
func (*Consignment) ProtoMessage()    {}
func (*Consignment) Descriptor() ([]byte, []int) {
	return fileDescriptor_3804bf87090b51a9, []int{1}
}

func (m *Consignment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Consignment.Unmarshal(m, b)
}
func (m *Consignment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Consignment.Marshal(b, m, deterministic)
}
func (m *Consignment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Consignment.Merge(m, src)
}
func (m *Consignment) XXX_Size() int {
	return xxx_messageInfo_Consignment.Size(m)
}
func (m *Consignment) XXX_DiscardUnknown() {
	xxx_messageInfo_Consignment.DiscardUnknown(m)
}

var xxx_messageInfo_Consignment proto.InternalMessageInfo

func (m *Consignment) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Consignment) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Consignment) GetWeight() int64 {
	if m != nil {
		return m.Weight
	}
	return 0
}

func (m *Consignment) GetContainers() []*Container {
	if m != nil {
		return m.Containers
	}
	return nil
}

func (m *Consignment) GetVesselId() string {
	if m != nil {
		return m.VesselId
	}
	return ""
}

type Container struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	CustomerId           string   `protobuf:"bytes,2,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	Origin               string   `protobuf:"bytes,3,opt,name=origin,proto3" json:"origin,omitempty"`
	UserId               string   `protobuf:"bytes,4,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Container) Reset()         { *m = Container{} }
func (m *Container) String() string { return proto.CompactTextString(m) }
func (*Container) ProtoMessage()    {}
func (*Container) Descriptor() ([]byte, []int) {
	return fileDescriptor_3804bf87090b51a9, []int{2}
}

func (m *Container) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Container.Unmarshal(m, b)
}
func (m *Container) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Container.Marshal(b, m, deterministic)
}
func (m *Container) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Container.Merge(m, src)
}
func (m *Container) XXX_Size() int {
	return xxx_messageInfo_Container.Size(m)
}
func (m *Container) XXX_DiscardUnknown() {
	xxx_messageInfo_Container.DiscardUnknown(m)
}

var xxx_messageInfo_Container proto.InternalMessageInfo

func (m *Container) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Container) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *Container) GetOrigin() string {
	if m != nil {
		return m.Origin
	}
	return ""
}

func (m *Container) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

type ConsignmentResponse struct {
	Api                  string       `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	Created              bool         `protobuf:"varint,2,opt,name=created,proto3" json:"created,omitempty"`
	Consignment          *Consignment `protobuf:"bytes,3,opt,name=consignment,proto3" json:"consignment,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ConsignmentResponse) Reset()         { *m = ConsignmentResponse{} }
func (m *ConsignmentResponse) String() string { return proto.CompactTextString(m) }
func (*ConsignmentResponse) ProtoMessage()    {}
func (*ConsignmentResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3804bf87090b51a9, []int{3}
}

func (m *ConsignmentResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConsignmentResponse.Unmarshal(m, b)
}
func (m *ConsignmentResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConsignmentResponse.Marshal(b, m, deterministic)
}
func (m *ConsignmentResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConsignmentResponse.Merge(m, src)
}
func (m *ConsignmentResponse) XXX_Size() int {
	return xxx_messageInfo_ConsignmentResponse.Size(m)
}
func (m *ConsignmentResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ConsignmentResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ConsignmentResponse proto.InternalMessageInfo

func (m *ConsignmentResponse) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *ConsignmentResponse) GetCreated() bool {
	if m != nil {
		return m.Created
	}
	return false
}

func (m *ConsignmentResponse) GetConsignment() *Consignment {
	if m != nil {
		return m.Consignment
	}
	return nil
}

func init() {
	proto.RegisterType((*ConsignmentRequest)(nil), "consignment_proto.ConsignmentRequest")
	proto.RegisterType((*Consignment)(nil), "consignment_proto.Consignment")
	proto.RegisterType((*Container)(nil), "consignment_proto.Container")
	proto.RegisterType((*ConsignmentResponse)(nil), "consignment_proto.ConsignmentResponse")
}

func init() { proto.RegisterFile("consignment.proto", fileDescriptor_3804bf87090b51a9) }

var fileDescriptor_3804bf87090b51a9 = []byte{
	// 324 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x35, 0x4d, 0x6d, 0x9b, 0x09, 0x88, 0x5d, 0x41, 0x83, 0x8a, 0x86, 0x80, 0xd2, 0x53, 0x0f,
	0xf5, 0xea, 0x41, 0xe8, 0xa9, 0xd7, 0xf5, 0x03, 0x4a, 0xcd, 0x0e, 0xed, 0x80, 0xdd, 0x8d, 0xbb,
	0x9b, 0x0a, 0x5e, 0xfd, 0x23, 0xbf, 0x50, 0xb2, 0xdb, 0xea, 0xd6, 0x16, 0x8a, 0xb7, 0xbc, 0x37,
	0xf3, 0xe6, 0xbd, 0xf0, 0x16, 0xfa, 0xa5, 0x92, 0x86, 0xe6, 0x72, 0x89, 0xd2, 0x0e, 0x2b, 0xad,
	0xac, 0x62, 0x21, 0x35, 0x75, 0x54, 0xb1, 0x00, 0x36, 0xfe, 0x25, 0x39, 0xbe, 0xd5, 0x68, 0x2c,
	0x3b, 0x85, 0x78, 0x56, 0x51, 0x16, 0xe5, 0xd1, 0x20, 0xe1, 0xcd, 0x27, 0x7b, 0x82, 0x34, 0x10,
	0x67, 0xad, 0x3c, 0x1a, 0xa4, 0xa3, 0x9b, 0xe1, 0xce, 0xc1, 0x61, 0x78, 0x2d, 0x94, 0x14, 0x5f,
	0x11, 0xa4, 0xc1, 0x90, 0x9d, 0x40, 0x8b, 0x84, 0xb3, 0x88, 0x79, 0x8b, 0x04, 0xcb, 0x21, 0x15,
	0x68, 0x4a, 0x4d, 0x95, 0x25, 0x25, 0x9d, 0x43, 0xc2, 0x43, 0x8a, 0x9d, 0x43, 0xe7, 0x1d, 0x69,
	0xbe, 0xb0, 0x59, 0xec, 0x54, 0x6b, 0xc4, 0x1e, 0x01, 0x4a, 0x25, 0xed, 0x8c, 0x24, 0x6a, 0x93,
	0xb5, 0xf3, 0x78, 0x90, 0x8e, 0xae, 0xf7, 0x47, 0xf3, 0x4b, 0x3c, 0xd8, 0x67, 0x57, 0x90, 0xac,
	0xd0, 0x18, 0x7c, 0x9d, 0x92, 0xc8, 0x8e, 0x9d, 0x6b, 0xcf, 0x13, 0x13, 0x51, 0x2c, 0x21, 0xf9,
	0x51, 0xed, 0x24, 0xbe, 0x85, 0xb4, 0xac, 0x8d, 0x55, 0x4b, 0xd4, 0x8d, 0xd6, 0x27, 0x86, 0x0d,
	0x35, 0x11, 0x4d, 0x60, 0xa5, 0x69, 0x4e, 0xd2, 0x05, 0x4e, 0xf8, 0x1a, 0xb1, 0x0b, 0xe8, 0xd6,
	0xc6, 0x8b, 0xda, 0x7e, 0xd0, 0xc0, 0x89, 0x28, 0x3e, 0x23, 0x38, 0xdb, 0xaa, 0xc3, 0x54, 0x4a,
	0x1a, 0xdc, 0xd3, 0x47, 0x06, 0xdd, 0x52, 0xe3, 0xcc, 0xa2, 0xf7, 0xed, 0xf1, 0x0d, 0xfc, 0xdb,
	0x54, 0xfc, 0xef, 0xa6, 0x46, 0x1f, 0x5b, 0x6f, 0xe2, 0x19, 0xf5, 0x8a, 0x4a, 0x64, 0x02, 0xfa,
	0x63, 0x67, 0x11, 0x96, 0x78, 0x77, 0xe0, 0xae, 0x7f, 0x4f, 0x97, 0xf7, 0x87, 0xd6, 0xfc, 0x7f,
	0x16, 0x47, 0x2f, 0x1d, 0x37, 0x7c, 0xf8, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x29, 0x7d, 0xf1, 0xf2,
	0xbe, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ConsignmentServiceClient is the client API for ConsignmentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ConsignmentServiceClient interface {
	CreateConsignment(ctx context.Context, in *ConsignmentRequest, opts ...grpc.CallOption) (*ConsignmentResponse, error)
}

type consignmentServiceClient struct {
	cc *grpc.ClientConn
}

func NewConsignmentServiceClient(cc *grpc.ClientConn) ConsignmentServiceClient {
	return &consignmentServiceClient{cc}
}

func (c *consignmentServiceClient) CreateConsignment(ctx context.Context, in *ConsignmentRequest, opts ...grpc.CallOption) (*ConsignmentResponse, error) {
	out := new(ConsignmentResponse)
	err := c.cc.Invoke(ctx, "/consignment_proto.ConsignmentService/CreateConsignment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConsignmentServiceServer is the server API for ConsignmentService service.
type ConsignmentServiceServer interface {
	CreateConsignment(context.Context, *ConsignmentRequest) (*ConsignmentResponse, error)
}

func RegisterConsignmentServiceServer(s *grpc.Server, srv ConsignmentServiceServer) {
	s.RegisterService(&_ConsignmentService_serviceDesc, srv)
}

func _ConsignmentService_CreateConsignment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConsignmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConsignmentServiceServer).CreateConsignment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/consignment_proto.ConsignmentService/CreateConsignment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConsignmentServiceServer).CreateConsignment(ctx, req.(*ConsignmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ConsignmentService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "consignment_proto.ConsignmentService",
	HandlerType: (*ConsignmentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateConsignment",
			Handler:    _ConsignmentService_CreateConsignment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "consignment.proto",
}