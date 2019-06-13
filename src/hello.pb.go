// Code generated by protoc-gen-go. DO NOT EDIT.
// source: src/hello.proto

package demo

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type FncRequest struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FncRequest) Reset()         { *m = FncRequest{} }
func (m *FncRequest) String() string { return proto.CompactTextString(m) }
func (*FncRequest) ProtoMessage()    {}
func (*FncRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_82cf92f603bcd0cf, []int{0}
}

func (m *FncRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FncRequest.Unmarshal(m, b)
}
func (m *FncRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FncRequest.Marshal(b, m, deterministic)
}
func (m *FncRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FncRequest.Merge(m, src)
}
func (m *FncRequest) XXX_Size() int {
	return xxx_messageInfo_FncRequest.Size(m)
}
func (m *FncRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FncRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FncRequest proto.InternalMessageInfo

func (m *FncRequest) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type FncResponse struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	Code                 int32    `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FncResponse) Reset()         { *m = FncResponse{} }
func (m *FncResponse) String() string { return proto.CompactTextString(m) }
func (*FncResponse) ProtoMessage()    {}
func (*FncResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_82cf92f603bcd0cf, []int{1}
}

func (m *FncResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FncResponse.Unmarshal(m, b)
}
func (m *FncResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FncResponse.Marshal(b, m, deterministic)
}
func (m *FncResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FncResponse.Merge(m, src)
}
func (m *FncResponse) XXX_Size() int {
	return xxx_messageInfo_FncResponse.Size(m)
}
func (m *FncResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FncResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FncResponse proto.InternalMessageInfo

func (m *FncResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *FncResponse) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func init() {
	proto.RegisterType((*FncRequest)(nil), "demo.FncRequest")
	proto.RegisterType((*FncResponse)(nil), "demo.FncResponse")
}

func init() { proto.RegisterFile("src/hello.proto", fileDescriptor_82cf92f603bcd0cf) }

var fileDescriptor_82cf92f603bcd0cf = []byte{
	// 151 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2f, 0x2e, 0x4a, 0xd6,
	0xcf, 0x48, 0xcd, 0xc9, 0xc9, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x49, 0x49, 0xcd,
	0xcd, 0x57, 0x92, 0xe3, 0xe2, 0x72, 0xcb, 0x4b, 0x0e, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11,
	0x12, 0xe0, 0x62, 0xce, 0x2d, 0x4e, 0x97, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0x31, 0x95,
	0x8c, 0xb9, 0xb8, 0xc1, 0xf2, 0xc5, 0x05, 0xf9, 0x79, 0xc5, 0xa9, 0x98, 0x0a, 0x84, 0x84, 0xb8,
	0x58, 0x92, 0xf3, 0x53, 0x52, 0x25, 0x98, 0x14, 0x18, 0x35, 0x58, 0x83, 0xc0, 0x6c, 0x23, 0x6b,
	0x2e, 0x6e, 0x97, 0xd4, 0xdc, 0xfc, 0xe0, 0xd4, 0xa2, 0xb2, 0xcc, 0xe4, 0x54, 0x21, 0x1d, 0x2e,
	0x66, 0xb7, 0xbc, 0x64, 0x21, 0x01, 0x3d, 0x90, 0x8d, 0x7a, 0x08, 0xeb, 0xa4, 0x04, 0x91, 0x44,
	0x20, 0x16, 0x28, 0x31, 0x24, 0xb1, 0x81, 0x9d, 0x67, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xb4,
	0x79, 0x60, 0xd9, 0xb1, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DemoServiceClient is the client API for DemoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DemoServiceClient interface {
	Fnc(ctx context.Context, in *FncRequest, opts ...grpc.CallOption) (*FncResponse, error)
}

type demoServiceClient struct {
	cc *grpc.ClientConn
}

func NewDemoServiceClient(cc *grpc.ClientConn) DemoServiceClient {
	return &demoServiceClient{cc}
}

func (c *demoServiceClient) Fnc(ctx context.Context, in *FncRequest, opts ...grpc.CallOption) (*FncResponse, error) {
	out := new(FncResponse)
	err := c.cc.Invoke(ctx, "/demo.DemoService/Fnc", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DemoServiceServer is the server API for DemoService service.
type DemoServiceServer interface {
	Fnc(context.Context, *FncRequest) (*FncResponse, error)
}

// UnimplementedDemoServiceServer can be embedded to have forward compatible implementations.
type UnimplementedDemoServiceServer struct {
}

func (*UnimplementedDemoServiceServer) Fnc(ctx context.Context, req *FncRequest) (*FncResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Fnc not implemented")
}

func RegisterDemoServiceServer(s *grpc.Server, srv DemoServiceServer) {
	s.RegisterService(&_DemoService_serviceDesc, srv)
}

func _DemoService_Fnc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FncRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DemoServiceServer).Fnc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/demo.DemoService/Fnc",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DemoServiceServer).Fnc(ctx, req.(*FncRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DemoService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "demo.DemoService",
	HandlerType: (*DemoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Fnc",
			Handler:    _DemoService_Fnc_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "src/hello.proto",
}
