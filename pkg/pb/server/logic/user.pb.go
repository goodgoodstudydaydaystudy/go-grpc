// Code generated by protoc-gen-go. DO NOT EDIT.
// source: server/logic/user.proto

package logic

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

// 1. account server
type UserInfo struct {
	UserId               uint32   `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Account              string   `protobuf:"bytes,2,opt,name=account,proto3" json:"account,omitempty"`
	Nickname             string   `protobuf:"bytes,3,opt,name=nickname,proto3" json:"nickname,omitempty"`
	Gender               string   `protobuf:"bytes,4,opt,name=gender,proto3" json:"gender,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserInfo) Reset()         { *m = UserInfo{} }
func (m *UserInfo) String() string { return proto.CompactTextString(m) }
func (*UserInfo) ProtoMessage()    {}
func (*UserInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_bfe0c5a6e48944d2, []int{0}
}

func (m *UserInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserInfo.Unmarshal(m, b)
}
func (m *UserInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserInfo.Marshal(b, m, deterministic)
}
func (m *UserInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserInfo.Merge(m, src)
}
func (m *UserInfo) XXX_Size() int {
	return xxx_messageInfo_UserInfo.Size(m)
}
func (m *UserInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_UserInfo.DiscardUnknown(m)
}

var xxx_messageInfo_UserInfo proto.InternalMessageInfo

func (m *UserInfo) GetUserId() uint32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *UserInfo) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *UserInfo) GetNickname() string {
	if m != nil {
		return m.Nickname
	}
	return ""
}

func (m *UserInfo) GetGender() string {
	if m != nil {
		return m.Gender
	}
	return ""
}

type AccountSerReq struct {
	Account              string   `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Nickname             string   `protobuf:"bytes,3,opt,name=nickname,proto3" json:"nickname,omitempty"`
	Gender               uint32   `protobuf:"varint,4,opt,name=gender,proto3" json:"gender,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccountSerReq) Reset()         { *m = AccountSerReq{} }
func (m *AccountSerReq) String() string { return proto.CompactTextString(m) }
func (*AccountSerReq) ProtoMessage()    {}
func (*AccountSerReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_bfe0c5a6e48944d2, []int{1}
}

func (m *AccountSerReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountSerReq.Unmarshal(m, b)
}
func (m *AccountSerReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountSerReq.Marshal(b, m, deterministic)
}
func (m *AccountSerReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountSerReq.Merge(m, src)
}
func (m *AccountSerReq) XXX_Size() int {
	return xxx_messageInfo_AccountSerReq.Size(m)
}
func (m *AccountSerReq) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountSerReq.DiscardUnknown(m)
}

var xxx_messageInfo_AccountSerReq proto.InternalMessageInfo

func (m *AccountSerReq) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *AccountSerReq) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *AccountSerReq) GetNickname() string {
	if m != nil {
		return m.Nickname
	}
	return ""
}

func (m *AccountSerReq) GetGender() uint32 {
	if m != nil {
		return m.Gender
	}
	return 0
}

type AccountSerResp struct {
	UserInfo             *UserInfo `protobuf:"bytes,1,opt,name=user_info,json=userInfo,proto3" json:"user_info,omitempty"`
	Balance              uint64    `protobuf:"varint,2,opt,name=balance,proto3" json:"balance,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *AccountSerResp) Reset()         { *m = AccountSerResp{} }
func (m *AccountSerResp) String() string { return proto.CompactTextString(m) }
func (*AccountSerResp) ProtoMessage()    {}
func (*AccountSerResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_bfe0c5a6e48944d2, []int{2}
}

func (m *AccountSerResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountSerResp.Unmarshal(m, b)
}
func (m *AccountSerResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountSerResp.Marshal(b, m, deterministic)
}
func (m *AccountSerResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountSerResp.Merge(m, src)
}
func (m *AccountSerResp) XXX_Size() int {
	return xxx_messageInfo_AccountSerResp.Size(m)
}
func (m *AccountSerResp) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountSerResp.DiscardUnknown(m)
}

var xxx_messageInfo_AccountSerResp proto.InternalMessageInfo

func (m *AccountSerResp) GetUserInfo() *UserInfo {
	if m != nil {
		return m.UserInfo
	}
	return nil
}

func (m *AccountSerResp) GetBalance() uint64 {
	if m != nil {
		return m.Balance
	}
	return 0
}

// 2. wallet server
type WalletSerReq struct {
	UserId               uint32   `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Count                int64    `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WalletSerReq) Reset()         { *m = WalletSerReq{} }
func (m *WalletSerReq) String() string { return proto.CompactTextString(m) }
func (*WalletSerReq) ProtoMessage()    {}
func (*WalletSerReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_bfe0c5a6e48944d2, []int{3}
}

func (m *WalletSerReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WalletSerReq.Unmarshal(m, b)
}
func (m *WalletSerReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WalletSerReq.Marshal(b, m, deterministic)
}
func (m *WalletSerReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WalletSerReq.Merge(m, src)
}
func (m *WalletSerReq) XXX_Size() int {
	return xxx_messageInfo_WalletSerReq.Size(m)
}
func (m *WalletSerReq) XXX_DiscardUnknown() {
	xxx_messageInfo_WalletSerReq.DiscardUnknown(m)
}

var xxx_messageInfo_WalletSerReq proto.InternalMessageInfo

func (m *WalletSerReq) GetUserId() uint32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *WalletSerReq) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

type WalletSerResp struct {
	Balance              uint64   `protobuf:"varint,1,opt,name=balance,proto3" json:"balance,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WalletSerResp) Reset()         { *m = WalletSerResp{} }
func (m *WalletSerResp) String() string { return proto.CompactTextString(m) }
func (*WalletSerResp) ProtoMessage()    {}
func (*WalletSerResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_bfe0c5a6e48944d2, []int{4}
}

func (m *WalletSerResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WalletSerResp.Unmarshal(m, b)
}
func (m *WalletSerResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WalletSerResp.Marshal(b, m, deterministic)
}
func (m *WalletSerResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WalletSerResp.Merge(m, src)
}
func (m *WalletSerResp) XXX_Size() int {
	return xxx_messageInfo_WalletSerResp.Size(m)
}
func (m *WalletSerResp) XXX_DiscardUnknown() {
	xxx_messageInfo_WalletSerResp.DiscardUnknown(m)
}

var xxx_messageInfo_WalletSerResp proto.InternalMessageInfo

func (m *WalletSerResp) GetBalance() uint64 {
	if m != nil {
		return m.Balance
	}
	return 0
}

func init() {
	proto.RegisterType((*UserInfo)(nil), "user.UserInfo")
	proto.RegisterType((*AccountSerReq)(nil), "user.AccountSerReq")
	proto.RegisterType((*AccountSerResp)(nil), "user.AccountSerResp")
	proto.RegisterType((*WalletSerReq)(nil), "user.WalletSerReq")
	proto.RegisterType((*WalletSerResp)(nil), "user.WalletSerResp")
}

func init() { proto.RegisterFile("server/logic/user.proto", fileDescriptor_bfe0c5a6e48944d2) }

var fileDescriptor_bfe0c5a6e48944d2 = []byte{
	// 315 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xcf, 0x4b, 0xfb, 0x40,
	0x10, 0xc5, 0xbf, 0xf9, 0x36, 0xf6, 0xc7, 0x68, 0x7a, 0xd8, 0x54, 0xbb, 0xe4, 0x54, 0x72, 0xaa,
	0x08, 0x0d, 0x54, 0x10, 0x3c, 0x78, 0xd0, 0x5b, 0xaf, 0x2b, 0x52, 0xf0, 0x22, 0x9b, 0x64, 0x1b,
	0x42, 0xe3, 0xee, 0x76, 0xb7, 0x51, 0xfa, 0xdf, 0x4b, 0x36, 0x4d, 0xba, 0x25, 0x78, 0xf0, 0x96,
	0x37, 0x9b, 0x99, 0xcf, 0x9b, 0xc7, 0xc0, 0x54, 0x33, 0xf5, 0xc5, 0x54, 0x54, 0x88, 0x2c, 0x4f,
	0xa2, 0x52, 0x33, 0xb5, 0x90, 0x4a, 0xec, 0x05, 0x72, 0xab, 0xef, 0x70, 0x07, 0xc3, 0x37, 0xcd,
	0xd4, 0x8a, 0x6f, 0x04, 0x9a, 0xc2, 0xa0, 0xaa, 0x7d, 0xe4, 0x29, 0x76, 0x66, 0xce, 0xdc, 0x23,
	0xfd, 0x4a, 0xae, 0x52, 0x84, 0x61, 0x40, 0x93, 0x44, 0x94, 0x7c, 0x8f, 0xff, 0xcf, 0x9c, 0xf9,
	0x88, 0x34, 0x12, 0x05, 0x30, 0xe4, 0x79, 0xb2, 0xe5, 0xf4, 0x93, 0xe1, 0x9e, 0x79, 0x6a, 0x35,
	0xba, 0x81, 0x7e, 0xc6, 0x78, 0xca, 0x14, 0x76, 0xcd, 0xcb, 0x51, 0x85, 0x07, 0xf0, 0x9e, 0xeb,
	0xf6, 0x57, 0xa6, 0x08, 0xdb, 0xd9, 0xe3, 0x9d, 0xce, 0x78, 0x49, 0xb5, 0xfe, 0x16, 0x2a, 0x3d,
	0x92, 0x5b, 0xfd, 0x07, 0xb4, 0xd7, 0xa2, 0xd7, 0x30, 0xb6, 0xd1, 0x5a, 0xa2, 0x3b, 0x18, 0xd5,
	0x3b, 0xf3, 0x8d, 0x30, 0xf4, 0xcb, 0xe5, 0x78, 0x61, 0x52, 0x6a, 0x62, 0x21, 0xc3, 0xb2, 0x09,
	0x08, 0xc3, 0x20, 0xa6, 0x05, 0xe5, 0x09, 0x33, 0x6e, 0x5c, 0xd2, 0xc8, 0xf0, 0x09, 0xae, 0xd6,
	0xb4, 0x28, 0x58, 0xb3, 0xd2, 0xaf, 0x51, 0x4e, 0xe0, 0xe2, 0x14, 0x64, 0x8f, 0xd4, 0x22, 0xbc,
	0x05, 0xcf, 0x6a, 0xd7, 0xd2, 0x26, 0x39, 0x67, 0xa4, 0xe5, 0x01, 0xdc, 0xca, 0x19, 0x7a, 0x04,
	0x38, 0xad, 0x82, 0xfc, 0xda, 0xf3, 0x59, 0xae, 0xc1, 0xa4, 0x5b, 0xd4, 0x32, 0xfc, 0x87, 0x1e,
	0x60, 0xd4, 0xd2, 0x10, 0xaa, 0x7f, 0xb2, 0xdd, 0x07, 0x7e, 0xa7, 0x56, 0xf5, 0xbd, 0x5c, 0xbf,
	0xfb, 0x72, 0x9b, 0x45, 0x32, 0x8e, 0xec, 0x9b, 0x8a, 0xfb, 0xe6, 0x9e, 0xee, 0x7f, 0x02, 0x00,
	0x00, 0xff, 0xff, 0xba, 0xb7, 0x87, 0xa0, 0x6a, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserClient is the client API for User service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserClient interface {
	AccountSer(ctx context.Context, in *AccountSerReq, opts ...grpc.CallOption) (*AccountSerResp, error)
	WalletSer(ctx context.Context, in *WalletSerReq, opts ...grpc.CallOption) (*WalletSerResp, error)
}

type userClient struct {
	cc *grpc.ClientConn
}

func NewUserClient(cc *grpc.ClientConn) UserClient {
	return &userClient{cc}
}

func (c *userClient) AccountSer(ctx context.Context, in *AccountSerReq, opts ...grpc.CallOption) (*AccountSerResp, error) {
	out := new(AccountSerResp)
	err := c.cc.Invoke(ctx, "/user.User/AccountSer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) WalletSer(ctx context.Context, in *WalletSerReq, opts ...grpc.CallOption) (*WalletSerResp, error) {
	out := new(WalletSerResp)
	err := c.cc.Invoke(ctx, "/user.User/WalletSer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServer is the server API for User service.
type UserServer interface {
	AccountSer(context.Context, *AccountSerReq) (*AccountSerResp, error)
	WalletSer(context.Context, *WalletSerReq) (*WalletSerResp, error)
}

// UnimplementedUserServer can be embedded to have forward compatible implementations.
type UnimplementedUserServer struct {
}

func (*UnimplementedUserServer) AccountSer(ctx context.Context, req *AccountSerReq) (*AccountSerResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AccountSer not implemented")
}
func (*UnimplementedUserServer) WalletSer(ctx context.Context, req *WalletSerReq) (*WalletSerResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WalletSer not implemented")
}

func RegisterUserServer(s *grpc.Server, srv UserServer) {
	s.RegisterService(&_User_serviceDesc, srv)
}

func _User_AccountSer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountSerReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).AccountSer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/AccountSer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).AccountSer(ctx, req.(*AccountSerReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_WalletSer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WalletSerReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).WalletSer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/WalletSer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).WalletSer(ctx, req.(*WalletSerReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _User_serviceDesc = grpc.ServiceDesc{
	ServiceName: "user.User",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AccountSer",
			Handler:    _User_AccountSer_Handler,
		},
		{
			MethodName: "WalletSer",
			Handler:    _User_WalletSer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "server/logic/user.proto",
}