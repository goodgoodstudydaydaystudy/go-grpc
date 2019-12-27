// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Field.proto

package Field

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

type ConsumeReq struct {
	ItemId               int64    `protobuf:"varint,1,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
	Description          string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	ItemNum              int64    `protobuf:"varint,3,opt,name=item_num,json=itemNum,proto3" json:"item_num,omitempty"`
	UserId               int32    `protobuf:"varint,4,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConsumeReq) Reset()         { *m = ConsumeReq{} }
func (m *ConsumeReq) String() string { return proto.CompactTextString(m) }
func (*ConsumeReq) ProtoMessage()    {}
func (*ConsumeReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_02e63749c35df299, []int{0}
}

func (m *ConsumeReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConsumeReq.Unmarshal(m, b)
}
func (m *ConsumeReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConsumeReq.Marshal(b, m, deterministic)
}
func (m *ConsumeReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConsumeReq.Merge(m, src)
}
func (m *ConsumeReq) XXX_Size() int {
	return xxx_messageInfo_ConsumeReq.Size(m)
}
func (m *ConsumeReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ConsumeReq.DiscardUnknown(m)
}

var xxx_messageInfo_ConsumeReq proto.InternalMessageInfo

func (m *ConsumeReq) GetItemId() int64 {
	if m != nil {
		return m.ItemId
	}
	return 0
}

func (m *ConsumeReq) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *ConsumeReq) GetItemNum() int64 {
	if m != nil {
		return m.ItemNum
	}
	return 0
}

func (m *ConsumeReq) GetUserId() int32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

type ConsumeResp struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	OrderId              int64    `protobuf:"varint,2,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConsumeResp) Reset()         { *m = ConsumeResp{} }
func (m *ConsumeResp) String() string { return proto.CompactTextString(m) }
func (*ConsumeResp) ProtoMessage()    {}
func (*ConsumeResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_02e63749c35df299, []int{1}
}

func (m *ConsumeResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConsumeResp.Unmarshal(m, b)
}
func (m *ConsumeResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConsumeResp.Marshal(b, m, deterministic)
}
func (m *ConsumeResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConsumeResp.Merge(m, src)
}
func (m *ConsumeResp) XXX_Size() int {
	return xxx_messageInfo_ConsumeResp.Size(m)
}
func (m *ConsumeResp) XXX_DiscardUnknown() {
	xxx_messageInfo_ConsumeResp.DiscardUnknown(m)
}

var xxx_messageInfo_ConsumeResp proto.InternalMessageInfo

func (m *ConsumeResp) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *ConsumeResp) GetOrderId() int64 {
	if m != nil {
		return m.OrderId
	}
	return 0
}

func init() {
	proto.RegisterType((*ConsumeReq)(nil), "Field.ConsumeReq")
	proto.RegisterType((*ConsumeResp)(nil), "Field.ConsumeResp")
}

func init() { proto.RegisterFile("Field.proto", fileDescriptor_02e63749c35df299) }

var fileDescriptor_02e63749c35df299 = []byte{
	// 212 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0x3b, 0x4b, 0xc4, 0x40,
	0x14, 0x85, 0x9d, 0x8d, 0x9b, 0xb8, 0x37, 0x95, 0xb7, 0x71, 0xb4, 0x1a, 0x52, 0xa5, 0x4a, 0xa1,
	0x95, 0xad, 0x0b, 0x42, 0x1a, 0x91, 0xf9, 0x03, 0x12, 0x9d, 0x8b, 0x0c, 0x64, 0x1e, 0xce, 0xa3,
	0x10, 0xff, 0xbc, 0xcc, 0x04, 0x1f, 0x6c, 0xf9, 0x5d, 0x38, 0xdf, 0xe1, 0x5c, 0xe8, 0x1f, 0x35,
	0xad, 0x6a, 0xf2, 0xc1, 0x25, 0x87, 0xfb, 0x0a, 0xc3, 0x17, 0xc0, 0xd1, 0xd9, 0x98, 0x0d, 0x49,
	0xfa, 0xc0, 0x2b, 0xe8, 0x74, 0x22, 0xf3, 0xa2, 0x15, 0x67, 0x82, 0x8d, 0x8d, 0x6c, 0x0b, 0xce,
	0x0a, 0x05, 0xf4, 0x8a, 0xe2, 0x5b, 0xd0, 0x3e, 0x69, 0x67, 0xf9, 0x4e, 0xb0, 0xf1, 0x20, 0xff,
	0x9f, 0xf0, 0x1a, 0x2e, 0x6a, 0xd4, 0x66, 0xc3, 0x9b, 0x9a, 0xad, 0xaa, 0xa7, 0x6c, 0x8a, 0x35,
	0x47, 0x0a, 0xc5, 0x7a, 0x2e, 0xd8, 0xb8, 0x97, 0x6d, 0xc1, 0x59, 0x0d, 0x0f, 0xd0, 0xff, 0x96,
	0x47, 0x8f, 0x1c, 0x3a, 0x43, 0x31, 0x2e, 0xef, 0x54, 0xdb, 0x0f, 0xf2, 0x07, 0x8b, 0xdc, 0x05,
	0xb5, 0x29, 0x76, 0x9b, 0xbc, 0xf2, 0xac, 0x6e, 0xef, 0xa1, 0x3b, 0x3a, 0x9b, 0x82, 0x5b, 0x71,
	0x82, 0xe6, 0x79, 0xf9, 0xc4, 0xcb, 0x69, 0xdb, 0xf9, 0xb7, 0xeb, 0x06, 0x4f, 0x4f, 0xd1, 0x0f,
	0x67, 0xaf, 0x6d, 0xfd, 0xc4, 0xdd, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x87, 0x65, 0xea, 0x9c,
	0x18, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ControlClient is the client API for Control service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ControlClient interface {
	Pay(ctx context.Context, in *ConsumeReq, opts ...grpc.CallOption) (*ConsumeResp, error)
}

type controlClient struct {
	cc *grpc.ClientConn
}

func NewControlClient(cc *grpc.ClientConn) ControlClient {
	return &controlClient{cc}
}

func (c *controlClient) Pay(ctx context.Context, in *ConsumeReq, opts ...grpc.CallOption) (*ConsumeResp, error) {
	out := new(ConsumeResp)
	err := c.cc.Invoke(ctx, "/Field.Control/Pay", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ControlServer is the server API for Control service.
type ControlServer interface {
	Pay(context.Context, *ConsumeReq) (*ConsumeResp, error)
}

// UnimplementedControlServer can be embedded to have forward compatible implementations.
type UnimplementedControlServer struct {
}

func (*UnimplementedControlServer) Pay(ctx context.Context, req *ConsumeReq) (*ConsumeResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Pay not implemented")
}

func RegisterControlServer(s *grpc.Server, srv ControlServer) {
	s.RegisterService(&_Control_serviceDesc, srv)
}

func _Control_Pay_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConsumeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControlServer).Pay(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Field.Control/Pay",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControlServer).Pay(ctx, req.(*ConsumeReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Control_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Field.Control",
	HandlerType: (*ControlServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Pay",
			Handler:    _Control_Pay_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "Field.proto",
}
