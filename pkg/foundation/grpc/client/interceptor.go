package client

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"

	"goodgoodstudy.com/go-grpc/pkg/procotol"
	protogrpc "goodgoodstudy.com/go-grpc/pkg/procotol/grpc"
	"goodgoodstudy.com/go-grpc/protocol/common/status"
)

func StatusCodeUnaryInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	var header, trailer metadata.MD
	var peer_ peer.Peer // 还用不上
	opts = append(opts, grpc.Header(&header), grpc.Trailer(&trailer), grpc.Peer(&peer_))

	err := invoker(ctx, method, req, reply, cc, opts...) // 发送原本的client请求

	if err != nil {
		// grpc层面出错了
		// grpc返回的err, 一般是传输过程, 或者其他层面出错了, 还不是服务逻辑的错误, 所以服务端返回的error, 不是在这里的
	} else {
		code, msg := protogrpc.GetStatus(header) // 服务端把错误码写在了header, 这里取出来, 才是业务逻辑的error
		if code != status.Success {
			err = protocol.NewServerError(code, msg)
		}
	}
	return err
}
