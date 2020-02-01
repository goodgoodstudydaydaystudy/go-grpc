package account

import (
	"context"
	"goodgoodstudy.com/go-grpc/protocol/common/status"
	"log"

	rpb "goodgoodstudy.com/go-grpc/pkg/pb/account"
	protocol "goodgoodstudy.com/go-grpc/pkg/procotol"
	"goodgoodstudy.com/go-grpc/pkg/server/account/dao"
	account "goodgoodstudy.com/go-grpc/pkg/server/account/dao/entity"
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
	resp := &rpb.RegisterResp{}
	err := s.db.Register(req)
	if err != nil {
		log.Println("db.register failed: ", err)
		return resp, protocol.NewServerError(status.ErrAccountExists) // 错误码不要hard code
	}
	userInfo, err := s.db.GetUserByAccount(req.GetAccount())
	if err != nil {
		log.Println("Register: GetUserByAccount failed:", err)
		return resp, err
	}
	return &rpb.RegisterResp{UserId: userInfo.UserID}, nil // pb改成uint32
}

// 登录功能
func (s *server) Login(ctx context.Context, req *rpb.LoginReq) (resp *rpb.LoginResp, err error) {
	resp = &rpb.LoginResp{} // 这个东西不能返回nil, 所以一开始初始化了, 省的麻烦
	pwd, err := s.db.GetUserPasswordByAccount(req.GetAccount())
	if err != nil {
		return
	}

	if req.GetPassword() != pwd {
		err = protocol.NewServerError(status.ErrPasswordError) //  自己把密码错误的错误码加进去status
		log.Printf("Login: user account %s password wrong\n", req.GetAccount())
		return
	}

	// 嫌麻烦的话可以用GetUserByAccount一次性查出user所有信息, 包括密码, 但是要注意密码不要包含在UserInfo结构体里面, 不能返回给client
	// 就不用查两次

	userInfo, err := s.db.GetUserByAccount(req.GetAccount())
	if err != nil {
		log.Println("GetUserByAccount failed:", err)
		return
	}

	return &rpb.LoginResp{
		UserInfo: account.UserInfoToPb(userInfo),
	}, nil
}

// 查询
// 通过 account 查询
func (s *server) GetUserByAccount(ctx context.Context, req *rpb.QueryByAccount) (resp *rpb.UserInfo, err error) {
	resp = &rpb.UserInfo{}
	user, err := s.db.GetUserByAccount(req.GetAccount())
	if err != nil {
		log.Println("server GetUserByAccount failed: ", err)
	}
	if user == nil {
		err = protocol.NewServerError(status.ErrAccountNotExists)
		return
	}

	return account.UserInfoToPb(user), nil
}


func (s *server) GetUserByUserId(ctx context.Context, req *rpb.QueryById) (resp *rpb.UserInfo, err error) {
	resp = &rpb.UserInfo{}
	user, err := s.db.GetUserById(req.GetUserId())
	if err != nil {
		log.Println("server GetUserById failed: ", err)
		return
	}
	if user == nil {
		err = protocol.NewServerError(status.ErrAccountNotExists)
		return
	}

	return account.UserInfoToPb(user), nil
}

