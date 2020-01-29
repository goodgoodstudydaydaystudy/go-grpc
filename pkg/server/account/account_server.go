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
// 返回"已注册"的错误信息
// 返回给用户的注册id
// 把 req 传给 db, 是偷懒的做法, 但是确实是最有效的, 可以这么做
func (s *server) Register(ctx context.Context, req *rpb.RegisterReq) (*rpb.RegisterResp, error) {
	err := s.db.InsertInfo(req)
	if err != nil {
		log.Println("db.insert failed: ", err)
		return &rpb.RegisterResp{}, protocol.NewServerError(-1000)
	}
	userId, _, err := s.db.QueryInfo(req.GetAccount())
	return &rpb.RegisterResp{UserId: userId}, nil
}

// 登录功能
func (s *server) Login(ctx context.Context, req *rpb.LoginReq) (*rpb.LoginResp, error) {
	userId, dbPassword, err := s.db.QueryInfo(req.GetAccount())
	//log.Printf("server userName: %v", userName)
	//log.Println("server login err: ", err) // nil
	queryId := uint32(userId)
	userInfo, err := s.db.GetUserById(queryId)
	isResult := dbPassword != req.GetPassword()
	if err != nil || isResult {
		log.Println("server login failed: ", err)
		return &rpb.LoginResp{}, protocol.NewServerError(-1001)
	}
	return &rpb.LoginResp{UserInfo:userInfo}, nil
}
