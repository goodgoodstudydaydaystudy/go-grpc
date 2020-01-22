package account

import (
	"context"
	rpb "goodgoodstudy.com/go-grpc/pkg/pb/account"
	"goodgoodstudy.com/go-grpc/pkg/server/account/dao"

	"log"
	"math/rand"
)

type server struct {
	db dao.AccountDao
}

// 服务器方法
// 方法名要和rpc接口一致，否则client注册服务器会报错

// 注册
func (s *server) Register(ctx context.Context, req *rpb.RegisterReq) (*rpb.RegisterResp, error) {
	// 返回给用户的注册id
	userId := rand.Int31()
	err := s.db.InsertInfo(userId, req.GetAccount(), req.GetPassword())
	if err != nil {
		log.Println("db.insert failed: ", err)
		return &rpb.RegisterResp{Message: "register failed"}, err
	}
	return &rpb.RegisterResp{Message: "register success", UeserId: userId}, nil
}

// 登录
func (s *server) Login(ctx context.Context, req *rpb.LoginReq) (*rpb.LoginResp, error) {
	return &rpb.LoginResp{Message: "login success"}, nil
}

// server connDb
func NewAccountServer() (*server, error) {
	db, err := dao.NewAccountDao("mysql")
	return &server{
		db: db,
	}, err
}
