package server

import (
	"context"

	"google.golang.org/grpc"

	"goodgoodstudy.com/go-grpc/pkg/procotol"
	protogrpc "goodgoodstudy.com/go-grpc/pkg/procotol/grpc"
)

/*
 * 所谓拦截器, 就是拦截了原本应该做的事情, 在这里就是handler, 其他的参数都是req是handler原本的参数, 也就是grpc处理的时候的req参数.
 * 在调用原本的逻辑(handler)之前和之后, 做一些自己想做的事, 比如这里就是对error和header进行了一些修改;
 * 执行自己的代码, 执行被拦截的代码(handler), 执行自己的代码;
 * 这东西就叫拦截器
 */

// grpc 是基于http的, 如果要返回数据给调用方, 除了在body返回, 还可以在http header
// 利用context, 很容易做到.
func StatusCodeUnaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	resp, err = handler(ctx, req) // 真正处理请求

	var se protocol.ServerError
	se = protocol.ToServerError(err)
	if err != nil {
		if e, ok := err.(protocol.ServerError); ok && e.Code() < 0 {
			se = e
		} else {
			se = protocol.ErrSystem
		}
	}

	if se != nil {
		protogrpc.SetStatus(ctx, se.Code(), se.Message()) // 写入context
	} else {
		protogrpc.SetStatus(ctx, 0, "OK")
	}
	err = nil // 注意这里hack了一下. grpc的err, 一般是传输过程, 或者其他层面出错了, 还不是服务逻辑的错误
	// 所以服务端返回的error, 不能直接以error的形式给grpc, 不然grpc会认为这服务连不通, 而不是执行业务失败.
	// 实际上, 业务逻辑失败, 是业务定义的. 对于grpc来说, 发送和接收都没有出错, 所以业务错误不等于grpc error.
	// 于是把业务的error code写在了header, 以这种方式返回; 而grpc err置为nil
	return // 这里err还是
}
