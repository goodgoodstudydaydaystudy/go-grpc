// Code generated by protoc-gen-go. DO NOT EDIT.
// source: wallet.proto

package wallet

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

type RechargeReq struct {
	UserId               uint32   `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Count                uint64   `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RechargeReq) Reset()         { *m = RechargeReq{} }
func (m *RechargeReq) String() string { return proto.CompactTextString(m) }
func (*RechargeReq) ProtoMessage()    {}
func (*RechargeReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_b88fd140af4deb6f, []int{0}
}

func (m *RechargeReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RechargeReq.Unmarshal(m, b)
}
func (m *RechargeReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RechargeReq.Marshal(b, m, deterministic)
}
func (m *RechargeReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RechargeReq.Merge(m, src)
}
func (m *RechargeReq) XXX_Size() int {
	return xxx_messageInfo_RechargeReq.Size(m)
}
func (m *RechargeReq) XXX_DiscardUnknown() {
	xxx_messageInfo_RechargeReq.DiscardUnknown(m)
}

var xxx_messageInfo_RechargeReq proto.InternalMessageInfo

func (m *RechargeReq) GetUserId() uint32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *RechargeReq) GetCount() uint64 {
	if m != nil {
		return m.Count
	}
	return 0
}

type RechargeResp struct {
	Balance              uint64   `protobuf:"varint,1,opt,name=balance,proto3" json:"balance,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RechargeResp) Reset()         { *m = RechargeResp{} }
func (m *RechargeResp) String() string { return proto.CompactTextString(m) }
func (*RechargeResp) ProtoMessage()    {}
func (*RechargeResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_b88fd140af4deb6f, []int{1}
}

func (m *RechargeResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RechargeResp.Unmarshal(m, b)
}
func (m *RechargeResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RechargeResp.Marshal(b, m, deterministic)
}
func (m *RechargeResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RechargeResp.Merge(m, src)
}
func (m *RechargeResp) XXX_Size() int {
	return xxx_messageInfo_RechargeResp.Size(m)
}
func (m *RechargeResp) XXX_DiscardUnknown() {
	xxx_messageInfo_RechargeResp.DiscardUnknown(m)
}

var xxx_messageInfo_RechargeResp proto.InternalMessageInfo

func (m *RechargeResp) GetBalance() uint64 {
	if m != nil {
		return m.Balance
	}
	return 0
}

type GetUserBalanceReq struct {
	UserId               uint32   `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserBalanceReq) Reset()         { *m = GetUserBalanceReq{} }
func (m *GetUserBalanceReq) String() string { return proto.CompactTextString(m) }
func (*GetUserBalanceReq) ProtoMessage()    {}
func (*GetUserBalanceReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_b88fd140af4deb6f, []int{2}
}

func (m *GetUserBalanceReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserBalanceReq.Unmarshal(m, b)
}
func (m *GetUserBalanceReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserBalanceReq.Marshal(b, m, deterministic)
}
func (m *GetUserBalanceReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserBalanceReq.Merge(m, src)
}
func (m *GetUserBalanceReq) XXX_Size() int {
	return xxx_messageInfo_GetUserBalanceReq.Size(m)
}
func (m *GetUserBalanceReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserBalanceReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserBalanceReq proto.InternalMessageInfo

func (m *GetUserBalanceReq) GetUserId() uint32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

type GetUserBalanceResp struct {
	Balance              uint64   `protobuf:"varint,1,opt,name=balance,proto3" json:"balance,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserBalanceResp) Reset()         { *m = GetUserBalanceResp{} }
func (m *GetUserBalanceResp) String() string { return proto.CompactTextString(m) }
func (*GetUserBalanceResp) ProtoMessage()    {}
func (*GetUserBalanceResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_b88fd140af4deb6f, []int{3}
}

func (m *GetUserBalanceResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserBalanceResp.Unmarshal(m, b)
}
func (m *GetUserBalanceResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserBalanceResp.Marshal(b, m, deterministic)
}
func (m *GetUserBalanceResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserBalanceResp.Merge(m, src)
}
func (m *GetUserBalanceResp) XXX_Size() int {
	return xxx_messageInfo_GetUserBalanceResp.Size(m)
}
func (m *GetUserBalanceResp) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserBalanceResp.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserBalanceResp proto.InternalMessageInfo

func (m *GetUserBalanceResp) GetBalance() uint64 {
	if m != nil {
		return m.Balance
	}
	return 0
}

func init() {
	proto.RegisterType((*RechargeReq)(nil), "wallet.RechargeReq")
	proto.RegisterType((*RechargeResp)(nil), "wallet.RechargeResp")
	proto.RegisterType((*GetUserBalanceReq)(nil), "wallet.GetUserBalanceReq")
	proto.RegisterType((*GetUserBalanceResp)(nil), "wallet.GetUserBalanceResp")
}

func init() { proto.RegisterFile("wallet.proto", fileDescriptor_b88fd140af4deb6f) }

var fileDescriptor_b88fd140af4deb6f = []byte{
	// 201 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0x4f, 0xcc, 0xc9,
	0x49, 0x2d, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x83, 0xf0, 0x94, 0x6c, 0xb8, 0xb8,
	0x83, 0x52, 0x93, 0x33, 0x12, 0x8b, 0xd2, 0x53, 0x83, 0x52, 0x0b, 0x85, 0xc4, 0xb9, 0xd8, 0x4b,
	0x8b, 0x53, 0x8b, 0xe2, 0x33, 0x53, 0x24, 0x18, 0x15, 0x18, 0x35, 0x78, 0x83, 0xd8, 0x40, 0x5c,
	0xcf, 0x14, 0x21, 0x11, 0x2e, 0xd6, 0xe4, 0xfc, 0xd2, 0xbc, 0x12, 0x09, 0x26, 0x05, 0x46, 0x0d,
	0x96, 0x20, 0x08, 0x47, 0x49, 0x83, 0x8b, 0x07, 0xa1, 0xbb, 0xb8, 0x40, 0x48, 0x82, 0x8b, 0x3d,
	0x29, 0x31, 0x27, 0x31, 0x2f, 0x39, 0x15, 0xac, 0x9d, 0x25, 0x08, 0xc6, 0x55, 0xd2, 0xe1, 0x12,
	0x74, 0x4f, 0x2d, 0x09, 0x2d, 0x4e, 0x2d, 0x72, 0x82, 0x88, 0xe0, 0xb3, 0x4d, 0x49, 0x8f, 0x4b,
	0x08, 0x5d, 0x35, 0x3e, 0xd3, 0x8d, 0x7a, 0x18, 0xb9, 0xd8, 0xc2, 0xc1, 0x1e, 0x12, 0x32, 0xe7,
	0xe2, 0x80, 0x39, 0x49, 0x48, 0x58, 0x0f, 0xea, 0x67, 0x24, 0x2f, 0x4a, 0x89, 0x60, 0x0a, 0x16,
	0x17, 0x28, 0x31, 0x08, 0x79, 0x72, 0xf1, 0xa1, 0xda, 0x29, 0x24, 0x09, 0x53, 0x89, 0xe1, 0x72,
	0x29, 0x29, 0x5c, 0x52, 0x20, 0xa3, 0x92, 0xd8, 0xc0, 0x61, 0x6c, 0x0c, 0x08, 0x00, 0x00, 0xff,
	0xff, 0xaa, 0xfb, 0x7a, 0xb2, 0x73, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// WalletClient is the client API for Wallet service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type WalletClient interface {
	Recharge(ctx context.Context, in *RechargeReq, opts ...grpc.CallOption) (*RechargeResp, error)
	GetUserBalance(ctx context.Context, in *GetUserBalanceReq, opts ...grpc.CallOption) (*GetUserBalanceResp, error)
}

type walletClient struct {
	cc *grpc.ClientConn
}

func NewWalletClient(cc *grpc.ClientConn) WalletClient {
	return &walletClient{cc}
}

func (c *walletClient) Recharge(ctx context.Context, in *RechargeReq, opts ...grpc.CallOption) (*RechargeResp, error) {
	out := new(RechargeResp)
	err := c.cc.Invoke(ctx, "/wallet.Wallet/Recharge", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletClient) GetUserBalance(ctx context.Context, in *GetUserBalanceReq, opts ...grpc.CallOption) (*GetUserBalanceResp, error) {
	out := new(GetUserBalanceResp)
	err := c.cc.Invoke(ctx, "/wallet.Wallet/GetUserBalance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WalletServer is the server API for Wallet service.
type WalletServer interface {
	Recharge(context.Context, *RechargeReq) (*RechargeResp, error)
	GetUserBalance(context.Context, *GetUserBalanceReq) (*GetUserBalanceResp, error)
}

// UnimplementedWalletServer can be embedded to have forward compatible implementations.
type UnimplementedWalletServer struct {
}

func (*UnimplementedWalletServer) Recharge(ctx context.Context, req *RechargeReq) (*RechargeResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Recharge not implemented")
}
func (*UnimplementedWalletServer) GetUserBalance(ctx context.Context, req *GetUserBalanceReq) (*GetUserBalanceResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserBalance not implemented")
}

func RegisterWalletServer(s *grpc.Server, srv WalletServer) {
	s.RegisterService(&_Wallet_serviceDesc, srv)
}

func _Wallet_Recharge_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RechargeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServer).Recharge(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wallet.Wallet/Recharge",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServer).Recharge(ctx, req.(*RechargeReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wallet_GetUserBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserBalanceReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServer).GetUserBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wallet.Wallet/GetUserBalance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServer).GetUserBalance(ctx, req.(*GetUserBalanceReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Wallet_serviceDesc = grpc.ServiceDesc{
	ServiceName: "wallet.Wallet",
	HandlerType: (*WalletServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Recharge",
			Handler:    _Wallet_Recharge_Handler,
		},
		{
			MethodName: "GetUserBalance",
			Handler:    _Wallet_GetUserBalance_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "wallet.proto",
}
