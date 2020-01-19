package account

import (
	"context"
	rpb "goodgoodstudy.com/go-grpc/pkg/pb/account"
	"goodgoodstudy.com/go-grpc/pkg/server/account/db"

	"log"
	"math/rand"
)

type Server struct {
}

// 服务器方法
// 方法名要和rpc接口一致，否则client注册服务器会报错

// 注册
func (s *Server) Register(ctx context.Context, req *rpb.RegisterReq) (*rpb.RegisterResp, error) {

	//accountStr  := req.GetAccount()
	//passwordStr := req.GetPassword()

	// 注册id
	userId := rand.Int31()
	message, err := db.InsertUserInfo("t_member", userId, req.GetAccount(), req.GetPassword())
	if err != nil {
		log.Println("db.insert failed: ", err)
		return &rpb.RegisterResp{Message:message}, err
	}
	return &rpb.RegisterResp{Message:"register success", UeserId:userId}, nil
}

// 登录
func (s *Server) Login(ctx context.Context, req *rpb.LoginReq) (*rpb.LoginResp, error) {
	log.Println("LogIn account: ", req.GetAccount())
	log.Println("LogIn password: ", req.GetPassword())
	return &rpb.LoginResp{Message: "login success"}, nil
}