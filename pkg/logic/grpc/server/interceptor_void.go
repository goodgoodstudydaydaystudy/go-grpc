package server

import (
	"context"
	"log"
	"strings"

	"goodgoodstudy.com/go-grpc/pkg/logic/user/server"
	pb "goodgoodstudy.com/go-grpc/pkg/pb/logic/user"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

)

// 作废
func LogicReqUnaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	// 从上下文中获取token
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		log.Println("logic interceptor get md failed")
		return
	}

	var claims *CustomClaims

	if t, ok := md["token"]; ok {
		token := strings.Join(t, ",")
		claims, err = ParseWithClaims(token)
		if err != nil {
			log.Println("parseWithClaims failed: ", err)
			return nil, err
		}
	}
	// 用接口查询userInfo 验证jwt的userInfo.id

	getUserReq := &pb.GetUserByIdReq{
		UserId: claims.userInfo.UserId,
	}
	getUserResp, err := (&server.UserLogic{}).GetUserInfo(ctx, getUserReq)
	if err != nil {
		log.Println("userId not exist or other error")
		return nil, err
	}

	userInfo := getUserResp.UserInfo
	if userInfo == nil {
		log.Println("user not exist ")
		return nil, err
	}

	// 通过验证，继续处理请求
	if userInfo.UserId == claims.userInfo.UserId {
		resp, err = handler(ctx, req)
		return
	}
	err = nil
	return
}