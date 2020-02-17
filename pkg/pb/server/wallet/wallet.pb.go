// Code generated by protoc-gen-go. DO NOT EDIT.
// source: server/wallet/wallet.proto

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
	Amount               int64    `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
	Account              string   `protobuf:"bytes,3,opt,name=account,proto3" json:"account,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RechargeReq) Reset()         { *m = RechargeReq{} }
func (m *RechargeReq) String() string { return proto.CompactTextString(m) }
func (*RechargeReq) ProtoMessage()    {}
func (*RechargeReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_6e585c95aeb2fa82, []int{0}
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

func (m *RechargeReq) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *RechargeReq) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

type RechargeResp struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RechargeResp) Reset()         { *m = RechargeResp{} }
func (m *RechargeResp) String() string { return proto.CompactTextString(m) }
func (*RechargeResp) ProtoMessage()    {}
func (*RechargeResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_6e585c95aeb2fa82, []int{1}
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
	return fileDescriptor_6e585c95aeb2fa82, []int{2}
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
	Balance              int64    `protobuf:"varint,1,opt,name=balance,proto3" json:"balance,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserBalanceResp) Reset()         { *m = GetUserBalanceResp{} }
func (m *GetUserBalanceResp) String() string { return proto.CompactTextString(m) }
func (*GetUserBalanceResp) ProtoMessage()    {}
func (*GetUserBalanceResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_6e585c95aeb2fa82, []int{3}
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

func (m *GetUserBalanceResp) GetBalance() int64 {
	if m != nil {
		return m.Balance
	}
	return 0
}

type GetTopUserReq struct {
	Top                  uint32   `protobuf:"varint,1,opt,name=top,proto3" json:"top,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTopUserReq) Reset()         { *m = GetTopUserReq{} }
func (m *GetTopUserReq) String() string { return proto.CompactTextString(m) }
func (*GetTopUserReq) ProtoMessage()    {}
func (*GetTopUserReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_6e585c95aeb2fa82, []int{4}
}

func (m *GetTopUserReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTopUserReq.Unmarshal(m, b)
}
func (m *GetTopUserReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTopUserReq.Marshal(b, m, deterministic)
}
func (m *GetTopUserReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTopUserReq.Merge(m, src)
}
func (m *GetTopUserReq) XXX_Size() int {
	return xxx_messageInfo_GetTopUserReq.Size(m)
}
func (m *GetTopUserReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTopUserReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetTopUserReq proto.InternalMessageInfo

func (m *GetTopUserReq) GetTop() uint32 {
	if m != nil {
		return m.Top
	}
	return 0
}

type GetTopUserResp struct {
	UserList             string   `protobuf:"bytes,1,opt,name=userList,proto3" json:"userList,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTopUserResp) Reset()         { *m = GetTopUserResp{} }
func (m *GetTopUserResp) String() string { return proto.CompactTextString(m) }
func (*GetTopUserResp) ProtoMessage()    {}
func (*GetTopUserResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_6e585c95aeb2fa82, []int{5}
}

func (m *GetTopUserResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTopUserResp.Unmarshal(m, b)
}
func (m *GetTopUserResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTopUserResp.Marshal(b, m, deterministic)
}
func (m *GetTopUserResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTopUserResp.Merge(m, src)
}
func (m *GetTopUserResp) XXX_Size() int {
	return xxx_messageInfo_GetTopUserResp.Size(m)
}
func (m *GetTopUserResp) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTopUserResp.DiscardUnknown(m)
}

var xxx_messageInfo_GetTopUserResp proto.InternalMessageInfo

func (m *GetTopUserResp) GetUserList() string {
	if m != nil {
		return m.UserList
	}
	return ""
}

type RecordOrderNoPaidReq struct {
	UserId               uint32   `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	OrderId              string   `protobuf:"bytes,2,opt,name=orderId,proto3" json:"orderId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RecordOrderNoPaidReq) Reset()         { *m = RecordOrderNoPaidReq{} }
func (m *RecordOrderNoPaidReq) String() string { return proto.CompactTextString(m) }
func (*RecordOrderNoPaidReq) ProtoMessage()    {}
func (*RecordOrderNoPaidReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_6e585c95aeb2fa82, []int{6}
}

func (m *RecordOrderNoPaidReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RecordOrderNoPaidReq.Unmarshal(m, b)
}
func (m *RecordOrderNoPaidReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RecordOrderNoPaidReq.Marshal(b, m, deterministic)
}
func (m *RecordOrderNoPaidReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RecordOrderNoPaidReq.Merge(m, src)
}
func (m *RecordOrderNoPaidReq) XXX_Size() int {
	return xxx_messageInfo_RecordOrderNoPaidReq.Size(m)
}
func (m *RecordOrderNoPaidReq) XXX_DiscardUnknown() {
	xxx_messageInfo_RecordOrderNoPaidReq.DiscardUnknown(m)
}

var xxx_messageInfo_RecordOrderNoPaidReq proto.InternalMessageInfo

func (m *RecordOrderNoPaidReq) GetUserId() uint32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *RecordOrderNoPaidReq) GetOrderId() string {
	if m != nil {
		return m.OrderId
	}
	return ""
}

type RecordOrderNoPaidResp struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RecordOrderNoPaidResp) Reset()         { *m = RecordOrderNoPaidResp{} }
func (m *RecordOrderNoPaidResp) String() string { return proto.CompactTextString(m) }
func (*RecordOrderNoPaidResp) ProtoMessage()    {}
func (*RecordOrderNoPaidResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_6e585c95aeb2fa82, []int{7}
}

func (m *RecordOrderNoPaidResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RecordOrderNoPaidResp.Unmarshal(m, b)
}
func (m *RecordOrderNoPaidResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RecordOrderNoPaidResp.Marshal(b, m, deterministic)
}
func (m *RecordOrderNoPaidResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RecordOrderNoPaidResp.Merge(m, src)
}
func (m *RecordOrderNoPaidResp) XXX_Size() int {
	return xxx_messageInfo_RecordOrderNoPaidResp.Size(m)
}
func (m *RecordOrderNoPaidResp) XXX_DiscardUnknown() {
	xxx_messageInfo_RecordOrderNoPaidResp.DiscardUnknown(m)
}

var xxx_messageInfo_RecordOrderNoPaidResp proto.InternalMessageInfo

type PayReq struct {
	OrderId              string   `protobuf:"bytes,1,opt,name=orderId,proto3" json:"orderId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PayReq) Reset()         { *m = PayReq{} }
func (m *PayReq) String() string { return proto.CompactTextString(m) }
func (*PayReq) ProtoMessage()    {}
func (*PayReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_6e585c95aeb2fa82, []int{8}
}

func (m *PayReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PayReq.Unmarshal(m, b)
}
func (m *PayReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PayReq.Marshal(b, m, deterministic)
}
func (m *PayReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PayReq.Merge(m, src)
}
func (m *PayReq) XXX_Size() int {
	return xxx_messageInfo_PayReq.Size(m)
}
func (m *PayReq) XXX_DiscardUnknown() {
	xxx_messageInfo_PayReq.DiscardUnknown(m)
}

var xxx_messageInfo_PayReq proto.InternalMessageInfo

func (m *PayReq) GetOrderId() string {
	if m != nil {
		return m.OrderId
	}
	return ""
}

type PayResp struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PayResp) Reset()         { *m = PayResp{} }
func (m *PayResp) String() string { return proto.CompactTextString(m) }
func (*PayResp) ProtoMessage()    {}
func (*PayResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_6e585c95aeb2fa82, []int{9}
}

func (m *PayResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PayResp.Unmarshal(m, b)
}
func (m *PayResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PayResp.Marshal(b, m, deterministic)
}
func (m *PayResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PayResp.Merge(m, src)
}
func (m *PayResp) XXX_Size() int {
	return xxx_messageInfo_PayResp.Size(m)
}
func (m *PayResp) XXX_DiscardUnknown() {
	xxx_messageInfo_PayResp.DiscardUnknown(m)
}

var xxx_messageInfo_PayResp proto.InternalMessageInfo

func init() {
	proto.RegisterType((*RechargeReq)(nil), "wallet.RechargeReq")
	proto.RegisterType((*RechargeResp)(nil), "wallet.RechargeResp")
	proto.RegisterType((*GetUserBalanceReq)(nil), "wallet.GetUserBalanceReq")
	proto.RegisterType((*GetUserBalanceResp)(nil), "wallet.GetUserBalanceResp")
	proto.RegisterType((*GetTopUserReq)(nil), "wallet.GetTopUserReq")
	proto.RegisterType((*GetTopUserResp)(nil), "wallet.GetTopUserResp")
	proto.RegisterType((*RecordOrderNoPaidReq)(nil), "wallet.RecordOrderNoPaidReq")
	proto.RegisterType((*RecordOrderNoPaidResp)(nil), "wallet.RecordOrderNoPaidResp")
	proto.RegisterType((*PayReq)(nil), "wallet.PayReq")
	proto.RegisterType((*PayResp)(nil), "wallet.PayResp")
}

func init() { proto.RegisterFile("server/wallet/wallet.proto", fileDescriptor_6e585c95aeb2fa82) }

var fileDescriptor_6e585c95aeb2fa82 = []byte{
	// 393 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x93, 0xc1, 0x6f, 0xda, 0x30,
	0x14, 0xc6, 0x17, 0x22, 0x25, 0xe4, 0x6d, 0xb0, 0xe1, 0x41, 0xc8, 0xac, 0x4d, 0x62, 0x3e, 0xe5,
	0x80, 0x40, 0xda, 0x0e, 0x3b, 0xed, 0xc2, 0x85, 0x21, 0x4d, 0x2d, 0xb2, 0x5a, 0xb5, 0xea, 0xa5,
	0x32, 0x89, 0x45, 0x51, 0x29, 0x76, 0xed, 0xd0, 0x8a, 0x3f, 0xb1, 0xff, 0x55, 0x65, 0x93, 0x94,
	0x40, 0x28, 0xa7, 0xe4, 0xf3, 0xfb, 0xbd, 0xef, 0x59, 0xdf, 0x93, 0x01, 0x6b, 0xae, 0x9e, 0xb8,
	0x1a, 0x3e, 0xb3, 0xe5, 0x92, 0x67, 0xf9, 0x67, 0x20, 0x95, 0xc8, 0x04, 0xf2, 0xb6, 0x8a, 0x5c,
	0xc3, 0x47, 0xca, 0x93, 0x3b, 0xa6, 0xe6, 0x9c, 0xf2, 0x47, 0xd4, 0x05, 0x7f, 0xad, 0xb9, 0xba,
	0x5d, 0xa4, 0x91, 0xd3, 0x73, 0xe2, 0x06, 0xf5, 0x8c, 0x9c, 0xa4, 0x28, 0x04, 0x8f, 0x3d, 0x88,
	0xf5, 0x2a, 0x8b, 0x6a, 0x3d, 0x27, 0x76, 0x69, 0xae, 0x50, 0x04, 0x3e, 0x4b, 0x12, 0x5b, 0x70,
	0x7b, 0x4e, 0x1c, 0xd0, 0x42, 0x92, 0x26, 0x7c, 0xda, 0x39, 0x6b, 0x49, 0xfa, 0xd0, 0x1a, 0xf3,
	0xec, 0x52, 0x73, 0x35, 0x62, 0x4b, 0xb6, 0x4a, 0x4e, 0xce, 0x23, 0x03, 0x40, 0x87, 0xb4, 0x96,
	0x66, 0xda, 0x6c, 0x2b, 0x2d, 0xee, 0xd2, 0x42, 0x92, 0x9f, 0xd0, 0x18, 0xf3, 0xec, 0x42, 0x48,
	0xd3, 0x62, 0x9c, 0xbf, 0x80, 0x9b, 0x09, 0x99, 0xbb, 0x9a, 0x5f, 0xd2, 0x87, 0x66, 0x19, 0xd1,
	0x12, 0x61, 0xa8, 0x9b, 0x71, 0xff, 0x17, 0x3a, 0xb3, 0x60, 0x40, 0xdf, 0x34, 0xf9, 0x07, 0x6d,
	0xca, 0x13, 0xa1, 0xd2, 0x73, 0x95, 0x72, 0x75, 0x26, 0xa6, 0x6c, 0x91, 0x1a, 0xdf, 0x10, 0xf2,
	0x2b, 0x1e, 0x04, 0x14, 0x81, 0x2f, 0x0c, 0x39, 0x49, 0x6d, 0x42, 0x01, 0x2d, 0x24, 0xe9, 0x42,
	0xe7, 0x88, 0x93, 0x96, 0x84, 0x80, 0x37, 0x65, 0x1b, 0x63, 0x5a, 0x6a, 0x76, 0xf6, 0x9b, 0x03,
	0xf0, 0x2d, 0xa3, 0xe5, 0xaf, 0x97, 0x1a, 0x78, 0x57, 0x76, 0x6b, 0xe8, 0x0f, 0xd4, 0x8b, 0x6c,
	0xd1, 0xd7, 0x41, 0xbe, 0xd8, 0xd2, 0x1e, 0x71, 0xbb, 0x7a, 0xa8, 0x25, 0xf9, 0x80, 0x26, 0x36,
	0x83, 0x52, 0xac, 0xe8, 0x5b, 0x41, 0x56, 0x96, 0x83, 0xf1, 0x7b, 0x25, 0x6b, 0xf5, 0x17, 0x60,
	0x17, 0x27, 0xea, 0x94, 0xd8, 0xdd, 0x16, 0x70, 0x78, 0xec, 0xd8, 0xb6, 0x53, 0x68, 0x55, 0x52,
	0x41, 0xdf, 0x4b, 0xd7, 0xae, 0x44, 0x8f, 0x7f, 0x9c, 0xa8, 0x5a, 0xcf, 0x18, 0xdc, 0x29, 0xdb,
	0xa0, 0x66, 0xc1, 0x6d, 0xd3, 0xc5, 0x9f, 0xf7, 0xb4, 0x21, 0x47, 0xe1, 0x4d, 0x5b, 0xde, 0xcf,
	0x87, 0x72, 0x36, 0xdc, 0x7b, 0x23, 0x33, 0xcf, 0xbe, 0x8e, 0xdf, 0xaf, 0x01, 0x00, 0x00, 0xff,
	0xff, 0x8e, 0x62, 0x5b, 0x0c, 0x3b, 0x03, 0x00, 0x00,
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
	GetTopUser(ctx context.Context, in *GetTopUserReq, opts ...grpc.CallOption) (*GetTopUserResp, error)
	RecordOrderNoPaid(ctx context.Context, in *RecordOrderNoPaidReq, opts ...grpc.CallOption) (*RecordOrderNoPaidResp, error)
	Pay(ctx context.Context, in *PayReq, opts ...grpc.CallOption) (*PayResp, error)
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

func (c *walletClient) GetTopUser(ctx context.Context, in *GetTopUserReq, opts ...grpc.CallOption) (*GetTopUserResp, error) {
	out := new(GetTopUserResp)
	err := c.cc.Invoke(ctx, "/wallet.Wallet/GetTopUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletClient) RecordOrderNoPaid(ctx context.Context, in *RecordOrderNoPaidReq, opts ...grpc.CallOption) (*RecordOrderNoPaidResp, error) {
	out := new(RecordOrderNoPaidResp)
	err := c.cc.Invoke(ctx, "/wallet.Wallet/RecordOrderNoPaid", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletClient) Pay(ctx context.Context, in *PayReq, opts ...grpc.CallOption) (*PayResp, error) {
	out := new(PayResp)
	err := c.cc.Invoke(ctx, "/wallet.Wallet/Pay", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WalletServer is the server API for Wallet service.
type WalletServer interface {
	Recharge(context.Context, *RechargeReq) (*RechargeResp, error)
	GetUserBalance(context.Context, *GetUserBalanceReq) (*GetUserBalanceResp, error)
	GetTopUser(context.Context, *GetTopUserReq) (*GetTopUserResp, error)
	RecordOrderNoPaid(context.Context, *RecordOrderNoPaidReq) (*RecordOrderNoPaidResp, error)
	Pay(context.Context, *PayReq) (*PayResp, error)
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
func (*UnimplementedWalletServer) GetTopUser(ctx context.Context, req *GetTopUserReq) (*GetTopUserResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTopUser not implemented")
}
func (*UnimplementedWalletServer) RecordOrderNoPaid(ctx context.Context, req *RecordOrderNoPaidReq) (*RecordOrderNoPaidResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RecordOrderNoPaid not implemented")
}
func (*UnimplementedWalletServer) Pay(ctx context.Context, req *PayReq) (*PayResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Pay not implemented")
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

func _Wallet_GetTopUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTopUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServer).GetTopUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wallet.Wallet/GetTopUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServer).GetTopUser(ctx, req.(*GetTopUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wallet_RecordOrderNoPaid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RecordOrderNoPaidReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServer).RecordOrderNoPaid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wallet.Wallet/RecordOrderNoPaid",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServer).RecordOrderNoPaid(ctx, req.(*RecordOrderNoPaidReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wallet_Pay_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PayReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WalletServer).Pay(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wallet.Wallet/Pay",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WalletServer).Pay(ctx, req.(*PayReq))
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
		{
			MethodName: "GetTopUser",
			Handler:    _Wallet_GetTopUser_Handler,
		},
		{
			MethodName: "RecordOrderNoPaid",
			Handler:    _Wallet_RecordOrderNoPaid_Handler,
		},
		{
			MethodName: "Pay",
			Handler:    _Wallet_Pay_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "server/wallet/wallet.proto",
}
