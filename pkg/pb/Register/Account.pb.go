// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Account.proto

package Account

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

type RegisReq struct {
	UeserId              int32    `protobuf:"varint,1,opt,name=ueser_id,json=ueserId,proto3" json:"ueser_id,omitempty"`
	Account              string   `protobuf:"bytes,2,opt,name=account,proto3" json:"account,omitempty"`
	Password             string   `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisReq) Reset()         { *m = RegisReq{} }
func (m *RegisReq) String() string { return proto.CompactTextString(m) }
func (*RegisReq) ProtoMessage()    {}
func (*RegisReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_628bbec62b553e4f, []int{0}
}

func (m *RegisReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisReq.Unmarshal(m, b)
}
func (m *RegisReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisReq.Marshal(b, m, deterministic)
}
func (m *RegisReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisReq.Merge(m, src)
}
func (m *RegisReq) XXX_Size() int {
	return xxx_messageInfo_RegisReq.Size(m)
}
func (m *RegisReq) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisReq.DiscardUnknown(m)
}

var xxx_messageInfo_RegisReq proto.InternalMessageInfo

func (m *RegisReq) GetUeserId() int32 {
	if m != nil {
		return m.UeserId
	}
	return 0
}

func (m *RegisReq) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *RegisReq) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type RegisResp struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisResp) Reset()         { *m = RegisResp{} }
func (m *RegisResp) String() string { return proto.CompactTextString(m) }
func (*RegisResp) ProtoMessage()    {}
func (*RegisResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_628bbec62b553e4f, []int{1}
}

func (m *RegisResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisResp.Unmarshal(m, b)
}
func (m *RegisResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisResp.Marshal(b, m, deterministic)
}
func (m *RegisResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisResp.Merge(m, src)
}
func (m *RegisResp) XXX_Size() int {
	return xxx_messageInfo_RegisResp.Size(m)
}
func (m *RegisResp) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisResp.DiscardUnknown(m)
}

var xxx_messageInfo_RegisResp proto.InternalMessageInfo

func (m *RegisResp) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type LogReq struct {
	Account              string   `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogReq) Reset()         { *m = LogReq{} }
func (m *LogReq) String() string { return proto.CompactTextString(m) }
func (*LogReq) ProtoMessage()    {}
func (*LogReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_628bbec62b553e4f, []int{2}
}

func (m *LogReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogReq.Unmarshal(m, b)
}
func (m *LogReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogReq.Marshal(b, m, deterministic)
}
func (m *LogReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogReq.Merge(m, src)
}
func (m *LogReq) XXX_Size() int {
	return xxx_messageInfo_LogReq.Size(m)
}
func (m *LogReq) XXX_DiscardUnknown() {
	xxx_messageInfo_LogReq.DiscardUnknown(m)
}

var xxx_messageInfo_LogReq proto.InternalMessageInfo

func (m *LogReq) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *LogReq) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type LogResp struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogResp) Reset()         { *m = LogResp{} }
func (m *LogResp) String() string { return proto.CompactTextString(m) }
func (*LogResp) ProtoMessage()    {}
func (*LogResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_628bbec62b553e4f, []int{3}
}

func (m *LogResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogResp.Unmarshal(m, b)
}
func (m *LogResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogResp.Marshal(b, m, deterministic)
}
func (m *LogResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogResp.Merge(m, src)
}
func (m *LogResp) XXX_Size() int {
	return xxx_messageInfo_LogResp.Size(m)
}
func (m *LogResp) XXX_DiscardUnknown() {
	xxx_messageInfo_LogResp.DiscardUnknown(m)
}

var xxx_messageInfo_LogResp proto.InternalMessageInfo

func (m *LogResp) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*RegisReq)(nil), "Account.RegisReq")
	proto.RegisterType((*RegisResp)(nil), "Account.RegisResp")
	proto.RegisterType((*LogReq)(nil), "Account.LogReq")
	proto.RegisterType((*LogResp)(nil), "Account.LogResp")
}

func init() { proto.RegisterFile("Account.proto", fileDescriptor_628bbec62b553e4f) }

var fileDescriptor_628bbec62b553e4f = []byte{
	// 228 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0x3f, 0x4b, 0x03, 0x41,
	0x10, 0xc5, 0xbd, 0x48, 0x72, 0x77, 0x83, 0xa2, 0x4e, 0xb5, 0x5e, 0x15, 0x56, 0x84, 0x14, 0x92,
	0x42, 0xb1, 0x15, 0x2c, 0x03, 0xc1, 0x62, 0x5b, 0x0b, 0x39, 0x6f, 0x87, 0x25, 0x45, 0x6e, 0xd6,
	0x9d, 0x0d, 0xf9, 0xfa, 0x72, 0x9b, 0x6c, 0xd0, 0x03, 0x2d, 0x7f, 0xf3, 0xe7, 0xbd, 0xc7, 0x83,
	0xcb, 0xd7, 0xae, 0xe3, 0x5d, 0x1f, 0x97, 0x3e, 0x70, 0x64, 0x2c, 0x8f, 0xa8, 0xdf, 0xa1, 0x32,
	0xe4, 0x36, 0x62, 0xe8, 0x0b, 0x6f, 0xa1, 0xda, 0x91, 0x50, 0xf8, 0xd8, 0x58, 0x55, 0xcc, 0x8b,
	0xc5, 0xd4, 0x94, 0x89, 0x57, 0x16, 0x15, 0x94, 0xed, 0xe1, 0x43, 0x4d, 0xe6, 0xc5, 0xa2, 0x36,
	0x19, 0xb1, 0x81, 0xca, 0xb7, 0x22, 0x7b, 0x0e, 0x56, 0x9d, 0xa7, 0xd5, 0x89, 0xf5, 0x3d, 0xd4,
	0x47, 0x71, 0xf1, 0x83, 0xc4, 0x96, 0x44, 0x5a, 0x47, 0x49, 0xbc, 0x36, 0x19, 0xf5, 0x0b, 0xcc,
	0xd6, 0xec, 0x86, 0x04, 0x3f, 0x6c, 0x8a, 0xbf, 0x6d, 0x26, 0x23, 0x9b, 0x3b, 0x28, 0xd3, 0xff,
	0x7f, 0x26, 0x8f, 0x02, 0x17, 0x6f, 0xb4, 0xef, 0xb8, 0x97, 0xc8, 0x5b, 0x0a, 0xf8, 0x0c, 0x90,
	0xb2, 0x45, 0x0a, 0x64, 0xf1, 0x66, 0x99, 0xfb, 0xc9, 0x6d, 0x34, 0x38, 0x1e, 0x89, 0xd7, 0x67,
	0xf8, 0x00, 0xd3, 0x35, 0xbb, 0x55, 0x8f, 0x57, 0xa7, 0xf5, 0x21, 0x7b, 0x73, 0xfd, 0x7b, 0x30,
	0x5c, 0x7f, 0xce, 0x52, 0xdb, 0x4f, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x50, 0x52, 0x77, 0xcd,
	0x7e, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// NewconstomerClient is the client API for Newconstomer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NewconstomerClient interface {
	Registered(ctx context.Context, in *RegisReq, opts ...grpc.CallOption) (*RegisResp, error)
	LogIn(ctx context.Context, in *LogReq, opts ...grpc.CallOption) (*LogResp, error)
}

type newconstomerClient struct {
	cc *grpc.ClientConn
}

func NewNewconstomerClient(cc *grpc.ClientConn) NewconstomerClient {
	return &newconstomerClient{cc}
}

func (c *newconstomerClient) Registered(ctx context.Context, in *RegisReq, opts ...grpc.CallOption) (*RegisResp, error) {
	out := new(RegisResp)
	err := c.cc.Invoke(ctx, "/Account.Newconstomer/Registered", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *newconstomerClient) LogIn(ctx context.Context, in *LogReq, opts ...grpc.CallOption) (*LogResp, error) {
	out := new(LogResp)
	err := c.cc.Invoke(ctx, "/Account.Newconstomer/LogIn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NewconstomerServer is the server API for Newconstomer service.
type NewconstomerServer interface {
	Registered(context.Context, *RegisReq) (*RegisResp, error)
	LogIn(context.Context, *LogReq) (*LogResp, error)
}

// UnimplementedNewconstomerServer can be embedded to have forward compatible implementations.
type UnimplementedNewconstomerServer struct {
}

func (*UnimplementedNewconstomerServer) Registered(ctx context.Context, req *RegisReq) (*RegisResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Registered not implemented")
}
func (*UnimplementedNewconstomerServer) LogIn(ctx context.Context, req *LogReq) (*LogResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LogIn not implemented")
}

func RegisterNewconstomerServer(s *grpc.Server, srv NewconstomerServer) {
	s.RegisterService(&_Newconstomer_serviceDesc, srv)
}

func _Newconstomer_Registered_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NewconstomerServer).Registered(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Account.Newconstomer/Registered",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NewconstomerServer).Registered(ctx, req.(*RegisReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Newconstomer_LogIn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NewconstomerServer).LogIn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Account.Newconstomer/LogIn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NewconstomerServer).LogIn(ctx, req.(*LogReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Newconstomer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Account.Newconstomer",
	HandlerType: (*NewconstomerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Registered",
			Handler:    _Newconstomer_Registered_Handler,
		},
		{
			MethodName: "LogIn",
			Handler:    _Newconstomer_LogIn_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "Account.proto",
}