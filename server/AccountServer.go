package account

import (
	"context"
	rpb "goodgoodstudy.com/go-grpc/pkg/pb/Register"
	"log"
	"math/rand"
	"strings"
)

// 创建 grpc服务的结构体
type Server struct {
}

// 服务器方法
// 方法名要和rpc接口一致，否则client注册服务器会报错

// 注册
func (s *Server) Register(ctx context.Context, req *rpb.RegisReq) (*rpb.RegisResp, error) {
	log.Println("register id:", req.GetAccount())
	accountStr  := req.GetAccount()
	passwordStr := req.GetPassword()

	accountLen := strings.Count(accountStr, "")
	passwordLen := strings.Count(passwordStr, "")

	switch {
	case accountLen > 15:
		return &rpb.RegisResp{Message:"account too long"}, nil
	case accountLen < 6 :
		return &rpb.RegisResp{Message:"account too short"}, nil
	case passwordLen > 20:
		return &rpb.RegisResp{Message:"password too long"}, nil
	case passwordLen < 8 :
		return &rpb.RegisResp{Message:"password too short"}, nil
	default:
		userId := rand.Int31()
		return &rpb.RegisResp{Message:"register success", UeserId:userId}, nil
	}
}

// 登录
func (s *Server) Login(ctx context.Context, req *rpb.LoginReq) (*rpb.LoginResp, error) {
	log.Println("LogIn account: ", req.GetAccount())
	log.Println("LogIn password: ", req.GetPassword())
	return &rpb.LoginResp{Message: "login success"}, nil
}