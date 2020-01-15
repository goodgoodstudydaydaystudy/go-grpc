package server

import (
	"context"
	rpb "goodgoodstudy.com/go-grpc/pkg/pb/Register"
	"log"
)

// 创建 grpc服务的结构体
type NewconstomerServer struct {
}

// 服务器方法
// 方法名要和rpc接口一致，否则client注册服务器会报错

// 注册
func (s *NewconstomerServer) Registered(ctx context.Context, req *rpb.RegisReq) (*rpb.RegisResp, error) {
	log.Println("register id:", req.GetUeserId())
	return &rpb.RegisResp{Message: "register success"}, nil
}

// 登录
func (s *NewconstomerServer) LogIn(ctx context.Context, req *rpb.LogReq) (*rpb.LogResp, error) {
	log.Println("LogIn account: ", req.GetAccount())
	log.Println("LogIn password: ", req.GetPassword())
	return &rpb.LogResp{Message: "login success"}, nil
}