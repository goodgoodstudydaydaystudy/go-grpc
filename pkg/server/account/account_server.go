package account

import (
	"context"
	rpb "goodgoodstudy.com/go-grpc/pkg/pb/account"
	protocol "goodgoodstudy.com/go-grpc/pkg/procotol"
	"goodgoodstudy.com/go-grpc/pkg/server/account/dao"
	"log"
)

type server struct {
	db dao.AccountDao
}

// server connDb
func NewAccountServer() (*server, error) {
	db, err := dao.NewAccountDao("mysql")
	return &server{
		db: db,
	}, err
}

// 方法名要和rpc接口一致，否则client注册服务器会报错
// server的注册功能
// TODO 返回"已注册"的错误信息
func (s *server) Register(ctx context.Context, req *rpb.RegisterReq) (*rpb.RegisterResp, error) {
	// 返回给用户的注册id
	err := s.db.InsertInfo(req.GetAccount(), req.GetPassword())
	if err != nil {
		log.Println("db.insert failed: ", err)
		return &rpb.RegisterResp{Message: ""}, protocol.ToServerError(err)
	}
	return &rpb.RegisterResp{Message: "register success"}, err
}

// 登录功能
func (s *server) Login(ctx context.Context, req *rpb.LoginReq) (*rpb.LoginResp, error) {
	return &rpb.LoginResp{}, protocol.NewServerError(-999, "test custom error")
	//return &rpb.LoginResp{Message: "login success"}, nil
}
