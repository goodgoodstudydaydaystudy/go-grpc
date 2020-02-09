package server

import (
	"context"
	"fmt"
	protocol "goodgoodstudy.com/go-grpc/pkg/procotol"
	protogrpc "goodgoodstudy.com/go-grpc/pkg/procotol/grpc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func LogicReqUnaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	// 过滤login
	methodString, _ :=grpc.Method(ctx)
	if methodString == "/user.User/Login" {
		resp, err = handler(ctx, req)
		return
	}

	// 获取md
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("missing credentials")
	}

	// 传入tokenString，userInfo
	var se protocol.ServerError
	var tokenString string

	if token, ok := md["resp_token"]; ok {
		tokenString = token[0]
	}

	err = checkToken(tokenString)
	if err != nil {
		if e, ok := err.(protocol.ServerError); ok && e.Code() < 0 {
			se = e
		} else {
			se = protocol.ErrSystem
		}
	}


	if se != nil {
		protogrpc.SetStatus(ctx, se.Code(), se.Message())
	} else {
		protogrpc.SetStatus(ctx, 0, "OK")
	}

	// 处理原来的请求
	resp, err = handler(ctx, req)

	se = protocol.ToServerError(err)
	if err != nil {
		if e, ok := err.(protocol.ServerError); ok && e.Code() < 0 {
			se = e
		} else {
			se = protocol.ErrSystem
		}
	}

	if se != nil {
		protogrpc.SetStatus(ctx, se.Code(), se.Message())
	} else {
		protogrpc.SetStatus(ctx, 0, "OK")
	}

	err = nil
	return
}
