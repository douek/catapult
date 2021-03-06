// Code generated by protoc-gen-go. DO NOT EDIT.
// source: node.proto

package node

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Response_RequestStatus int32

const (
	Response_SUCCESSFUL Response_RequestStatus = 0
	Response_FAILED     Response_RequestStatus = 1
)

var Response_RequestStatus_name = map[int32]string{
	0: "SUCCESSFUL",
	1: "FAILED",
}

var Response_RequestStatus_value = map[string]int32{
	"SUCCESSFUL": 0,
	"FAILED":     1,
}

func (x Response_RequestStatus) String() string {
	return proto.EnumName(Response_RequestStatus_name, int32(x))
}

func (Response_RequestStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0c843d59d2d938e7, []int{2, 0}
}

type UUID struct {
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UUID) Reset()         { *m = UUID{} }
func (m *UUID) String() string { return proto.CompactTextString(m) }
func (*UUID) ProtoMessage()    {}
func (*UUID) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c843d59d2d938e7, []int{0}
}

func (m *UUID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UUID.Unmarshal(m, b)
}
func (m *UUID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UUID.Marshal(b, m, deterministic)
}
func (m *UUID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UUID.Merge(m, src)
}
func (m *UUID) XXX_Size() int {
	return xxx_messageInfo_UUID.Size(m)
}
func (m *UUID) XXX_DiscardUnknown() {
	xxx_messageInfo_UUID.DiscardUnknown(m)
}

var xxx_messageInfo_UUID proto.InternalMessageInfo

func (m *UUID) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type VmConfig struct {
	VmID                 *UUID    `protobuf:"bytes,1,opt,name=vmID,proto3" json:"vmID,omitempty"`
	Memory               int64    `protobuf:"varint,2,opt,name=memory,proto3" json:"memory,omitempty"`
	Vcpus                int64    `protobuf:"varint,3,opt,name=vcpus,proto3" json:"vcpus,omitempty"`
	KernelImage          string   `protobuf:"bytes,4,opt,name=kernelImage,proto3" json:"kernelImage,omitempty"`
	RootFileSystem       string   `protobuf:"bytes,5,opt,name=rootFileSystem,proto3" json:"rootFileSystem,omitempty"`
	Address              string   `protobuf:"bytes,6,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VmConfig) Reset()         { *m = VmConfig{} }
func (m *VmConfig) String() string { return proto.CompactTextString(m) }
func (*VmConfig) ProtoMessage()    {}
func (*VmConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c843d59d2d938e7, []int{1}
}

func (m *VmConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VmConfig.Unmarshal(m, b)
}
func (m *VmConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VmConfig.Marshal(b, m, deterministic)
}
func (m *VmConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VmConfig.Merge(m, src)
}
func (m *VmConfig) XXX_Size() int {
	return xxx_messageInfo_VmConfig.Size(m)
}
func (m *VmConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_VmConfig.DiscardUnknown(m)
}

var xxx_messageInfo_VmConfig proto.InternalMessageInfo

func (m *VmConfig) GetVmID() *UUID {
	if m != nil {
		return m.VmID
	}
	return nil
}

func (m *VmConfig) GetMemory() int64 {
	if m != nil {
		return m.Memory
	}
	return 0
}

func (m *VmConfig) GetVcpus() int64 {
	if m != nil {
		return m.Vcpus
	}
	return 0
}

func (m *VmConfig) GetKernelImage() string {
	if m != nil {
		return m.KernelImage
	}
	return ""
}

func (m *VmConfig) GetRootFileSystem() string {
	if m != nil {
		return m.RootFileSystem
	}
	return ""
}

func (m *VmConfig) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type Response struct {
	Status               Response_RequestStatus `protobuf:"varint,1,opt,name=status,proto3,enum=Response_RequestStatus" json:"status,omitempty"`
	Config               *VmConfig              `protobuf:"bytes,2,opt,name=config,proto3" json:"config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c843d59d2d938e7, []int{2}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetStatus() Response_RequestStatus {
	if m != nil {
		return m.Status
	}
	return Response_SUCCESSFUL
}

func (m *Response) GetConfig() *VmConfig {
	if m != nil {
		return m.Config
	}
	return nil
}

type VmList struct {
	VmID                 []*UUID  `protobuf:"bytes,1,rep,name=vmID,proto3" json:"vmID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VmList) Reset()         { *m = VmList{} }
func (m *VmList) String() string { return proto.CompactTextString(m) }
func (*VmList) ProtoMessage()    {}
func (*VmList) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c843d59d2d938e7, []int{3}
}

func (m *VmList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VmList.Unmarshal(m, b)
}
func (m *VmList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VmList.Marshal(b, m, deterministic)
}
func (m *VmList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VmList.Merge(m, src)
}
func (m *VmList) XXX_Size() int {
	return xxx_messageInfo_VmList.Size(m)
}
func (m *VmList) XXX_DiscardUnknown() {
	xxx_messageInfo_VmList.DiscardUnknown(m)
}

var xxx_messageInfo_VmList proto.InternalMessageInfo

func (m *VmList) GetVmID() []*UUID {
	if m != nil {
		return m.VmID
	}
	return nil
}

func init() {
	proto.RegisterEnum("Response_RequestStatus", Response_RequestStatus_name, Response_RequestStatus_value)
	proto.RegisterType((*UUID)(nil), "UUID")
	proto.RegisterType((*VmConfig)(nil), "VmConfig")
	proto.RegisterType((*Response)(nil), "Response")
	proto.RegisterType((*VmList)(nil), "VmList")
}

func init() { proto.RegisterFile("node.proto", fileDescriptor_0c843d59d2d938e7) }

var fileDescriptor_0c843d59d2d938e7 = []byte{
	// 365 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x51, 0xdb, 0xaa, 0xda, 0x40,
	0x14, 0x35, 0x35, 0x4e, 0x74, 0x4b, 0x45, 0x86, 0x62, 0x53, 0xeb, 0x83, 0xa6, 0x50, 0x84, 0x96,
	0x11, 0xd2, 0x2f, 0x28, 0x5e, 0x20, 0xa0, 0x7d, 0x48, 0x48, 0xde, 0xa3, 0xd9, 0x06, 0x69, 0x26,
	0x93, 0x66, 0x26, 0x82, 0xdf, 0xd0, 0x0f, 0x3a, 0xbf, 0x77, 0x70, 0x62, 0xce, 0x51, 0x1f, 0xd7,
	0x65, 0x5f, 0x17, 0x40, 0x2e, 0x12, 0x64, 0x45, 0x29, 0x94, 0x18, 0x7f, 0x4d, 0x85, 0x48, 0x33,
	0x5c, 0x68, 0xb4, 0xaf, 0x8e, 0x0b, 0xe4, 0x85, 0xba, 0xd4, 0xa2, 0x33, 0x01, 0x33, 0x0c, 0xbd,
	0x15, 0xfd, 0x04, 0x9d, 0x73, 0x9c, 0x55, 0x68, 0x1b, 0x53, 0x63, 0xde, 0xf3, 0x6b, 0xe0, 0xbc,
	0x18, 0xd0, 0x8d, 0xf8, 0x52, 0xe4, 0xc7, 0x53, 0x4a, 0xbf, 0x80, 0x79, 0xe6, 0xde, 0x4a, 0x3b,
	0xfa, 0x6e, 0x87, 0x5d, 0xeb, 0x7c, 0x4d, 0xd1, 0x11, 0x10, 0x8e, 0x5c, 0x94, 0x17, 0xfb, 0xc3,
	0xd4, 0x98, 0xb7, 0xfd, 0x1b, 0xd2, 0x5d, 0x0f, 0x45, 0x25, 0xed, 0xb6, 0xa6, 0x6b, 0x40, 0xa7,
	0xd0, 0xff, 0x8b, 0x65, 0x8e, 0x99, 0xc7, 0xe3, 0x14, 0x6d, 0x53, 0x4f, 0xbc, 0xa7, 0xe8, 0x77,
	0x18, 0x94, 0x42, 0xa8, 0xcd, 0x29, 0xc3, 0xe0, 0x22, 0x15, 0x72, 0xbb, 0xa3, 0x4d, 0x4f, 0x2c,
	0xb5, 0xc1, 0x8a, 0x93, 0xa4, 0x44, 0x29, 0x6d, 0xa2, 0x0d, 0x0d, 0x74, 0xfe, 0x1b, 0xd0, 0xf5,
	0x51, 0x16, 0x22, 0x97, 0x48, 0x17, 0x40, 0xa4, 0x8a, 0x55, 0x25, 0xf5, 0xee, 0x03, 0xf7, 0x33,
	0x6b, 0x24, 0xe6, 0xe3, 0xbf, 0x0a, 0xa5, 0x0a, 0xb4, 0xec, 0xdf, 0x6c, 0x74, 0x06, 0xe4, 0xa0,
	0x8f, 0xd6, 0xf7, 0xf4, 0xdd, 0x1e, 0x6b, 0xbe, 0xe0, 0xdf, 0x04, 0xe7, 0x07, 0x7c, 0x7c, 0xa8,
	0xa5, 0x03, 0x80, 0x20, 0x5c, 0x2e, 0xd7, 0x41, 0xb0, 0x09, 0xb7, 0xc3, 0x16, 0x05, 0x20, 0x9b,
	0xdf, 0xde, 0x76, 0xbd, 0x1a, 0x1a, 0xce, 0x37, 0x20, 0x11, 0xdf, 0x9e, 0xa4, 0xba, 0x7b, 0x62,
	0xfb, 0xe9, 0x89, 0x6e, 0x05, 0xe6, 0x1f, 0x91, 0x20, 0x9d, 0x81, 0x15, 0xa8, 0xb8, 0x54, 0xd1,
	0x8e, 0xbe, 0xcf, 0x1d, 0xf7, 0xde, 0x76, 0x76, 0x5a, 0x74, 0x02, 0x24, 0x50, 0xa2, 0x88, 0x76,
	0xb4, 0xee, 0xf0, 0xa8, 0xfe, 0x04, 0xeb, 0x3a, 0x2b, 0xda, 0x49, 0x3a, 0x62, 0x75, 0xf8, 0xac,
	0x09, 0x9f, 0xad, 0xaf, 0xe1, 0x8f, 0x2d, 0x56, 0xef, 0xe3, 0xb4, 0xf6, 0x44, 0x4b, 0xbf, 0x5e,
	0x03, 0x00, 0x00, 0xff, 0xff, 0x43, 0xbb, 0xd4, 0x39, 0x33, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// NodeClient is the client API for Node service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NodeClient interface {
	StartVM(ctx context.Context, in *VmConfig, opts ...grpc.CallOption) (*Response, error)
	StopVM(ctx context.Context, in *UUID, opts ...grpc.CallOption) (*Response, error)
	ListVMs(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*VmList, error)
}

type nodeClient struct {
	cc *grpc.ClientConn
}

func NewNodeClient(cc *grpc.ClientConn) NodeClient {
	return &nodeClient{cc}
}

func (c *nodeClient) StartVM(ctx context.Context, in *VmConfig, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/Node/StartVM", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeClient) StopVM(ctx context.Context, in *UUID, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/Node/StopVM", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeClient) ListVMs(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*VmList, error) {
	out := new(VmList)
	err := c.cc.Invoke(ctx, "/Node/ListVMs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NodeServer is the server API for Node service.
type NodeServer interface {
	StartVM(context.Context, *VmConfig) (*Response, error)
	StopVM(context.Context, *UUID) (*Response, error)
	ListVMs(context.Context, *empty.Empty) (*VmList, error)
}

// UnimplementedNodeServer can be embedded to have forward compatible implementations.
type UnimplementedNodeServer struct {
}

func (*UnimplementedNodeServer) StartVM(ctx context.Context, req *VmConfig) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartVM not implemented")
}
func (*UnimplementedNodeServer) StopVM(ctx context.Context, req *UUID) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StopVM not implemented")
}
func (*UnimplementedNodeServer) ListVMs(ctx context.Context, req *empty.Empty) (*VmList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListVMs not implemented")
}

func RegisterNodeServer(s *grpc.Server, srv NodeServer) {
	s.RegisterService(&_Node_serviceDesc, srv)
}

func _Node_StartVM_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VmConfig)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).StartVM(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Node/StartVM",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).StartVM(ctx, req.(*VmConfig))
	}
	return interceptor(ctx, in, info, handler)
}

func _Node_StopVM_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UUID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).StopVM(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Node/StopVM",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).StopVM(ctx, req.(*UUID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Node_ListVMs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).ListVMs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Node/ListVMs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).ListVMs(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _Node_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Node",
	HandlerType: (*NodeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StartVM",
			Handler:    _Node_StartVM_Handler,
		},
		{
			MethodName: "StopVM",
			Handler:    _Node_StopVM_Handler,
		},
		{
			MethodName: "ListVMs",
			Handler:    _Node_ListVMs_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "node.proto",
}
