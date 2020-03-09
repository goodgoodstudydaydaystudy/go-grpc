package account

import (
	"context"
	"goodgoodstudy.com/go-grpc/pkg/server/account/cao"
	account "goodgoodstudy.com/go-grpc/pkg/server/account/dao/entity"
	"log"

	rpb "goodgoodstudy.com/go-grpc/pkg/pb/server/account"
	protocol "goodgoodstudy.com/go-grpc/pkg/procotol"
	"goodgoodstudy.com/go-grpc/pkg/server/account/dao"
	"goodgoodstudy.com/go-grpc/protocol/common/status"

	"github.com/go-redis/redis"
)

type server struct {
	db dao.AccountDao
	ca cao.AccountCache
}

// server connDb
func NewAccountServer() (*server, error) {
	db, err := dao.NewAccountDao("mysql")

	cli := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	ca := cao.AccountCache{
		Cli: cli,
	}

	return &server{
		db: db,
		ca: ca,
	}, err
}

// 方法名要和rpc接口一致，否则client注册服务器会报错
// server的注册功能
// 返回"已注册"的错误信息
// 返回给用户的注册id
// 把 req 传给 db, 是偷懒的做法, 但是确实是最有效的, 可以这么做
func (s *server) AddUser(ctx context.Context, req *rpb.AddUserReq) (*rpb.AddUserResp, error) {
	resp := &rpb.AddUserResp{}
	err := s.db.AddUsr(req)
	if err != nil {
		log.Println("db.register failed: ", err)
		return resp, protocol.NewServerError(status.ErrAccountExists) // 错误码不要hard code
	}
	userInfo, err := s.db.GetUserByAccount(req.GetAccount())
	if err != nil {
		log.Println("Register: GetUserByAccount failed:", err)
		return resp, err
	}
	return &rpb.AddUserResp{UserId: userInfo.UserID}, nil // pb改成uint32
}

// 登录功能
func (s *server) CheckPwd(ctx context.Context, req *rpb.CheckPwdReq) (resp *rpb.CheckPwdResp, err error) {
	resp = &rpb.CheckPwdResp{}
	// get data from cache
	userInfo, isValid, err := s.ca.GetUserInfoByID(req.UserId)
	if err != nil {
		log.Println("server GetUserInfoByID failed:", err)
		return
	}
	//log.Println("server userInfo:", userInfo)
	//log.Printf("userInfo.Password: %v, req.Password: %v", userInfo.Password, req.Password)

	// Err password
	if isValid == true && userInfo.Password != req.Password {
		err = protocol.NewServerError(status.ErrPasswordError) //  自己把密码错误的错误码加进去status
		log.Printf("Login: user account %v password wrong\n", req.UserId)
		return
	}
	// correct password
	if isValid == true {
		return &rpb.CheckPwdResp{
			Nickname:             userInfo.Nickname,
			Gender:               userInfo.Gender,
		}, nil
	}

	// cache have no data
	// maybe db also have no data, and get temp data.
	var newUserInfo *account.UserInfo
	if isValid == false {
		//log.Println("isValid == false")
		// get data from db
		newUserInfo, err = s.db.GetUserById(req.UserId)
		//log.Println("srv newUserInfo:", newUserInfo)
		//log.Println("newUserInfo:", newUserInfo)
		// data is nil, db have no data
		if newUserInfo == nil {
			// temp data write into cache
			err := s.ca.WriteIntoCache(nil)
			if err != nil {
				log.Println("CheckPwd WriteIntoCache failed:", err)
				return nil, protocol.NewServerError(status.ErrPwdCache)
			}
			return nil, protocol.NewServerError(status.ErrAccountExists)
		}
		if err != nil {
			log.Println("CheckPwd GetUserById failed:", err)
			return nil, protocol.NewServerError(status.ErrDB)
		}
	}

	// get valid data from db and write into cache
	if err := s.ca.WriteIntoCache(newUserInfo, req.UserId); err != nil {
		log.Println("serve WriteIntoCache second failed:", err)
		return nil, protocol.NewServerError(status.ErrPwdCache)
	}
	//log.Println("server userInfo:", userInfo)

	// return
	return &rpb.CheckPwdResp{
		Nickname:             newUserInfo.Nickname,
		Gender:               rpb.Gender(newUserInfo.Gender),
	}, nil
}

// 查询
// 通过 account 查询
func (s *server) GetUserByAccount(ctx context.Context, req *rpb.GetUserByAccountReq) (resp *rpb.GetUserByAccountResp, err error) {
	resp = &rpb.GetUserByAccountResp{}
	user, err := s.db.GetUserByAccount(req.GetAccount())
	if err != nil {
		log.Println("server GetUserByAccount failed: ", err)
	}
	if user == nil {
		err = protocol.NewServerError(status.ErrAccountNotExists)
		return
	}
	resp.UserInfo = &rpb.UserInfo{
		UserId:               user.UserID,
		Account:              user.Account,
		Nickname:             user.Nickname,
		Gender:               rpb.Gender(user.Gender),
	}
	return resp, nil
}

func (s *server) GetUserById(ctx context.Context, req *rpb.GetUserByIdReq) (resp *rpb.GetUserByIdResp, err error) {
	resp = &rpb.GetUserByIdResp{}
	user, err := s.db.GetUserById(req.GetUserId())
	if err != nil {
		log.Println("server GetUserById failed: ", err)
		return nil, err
	}
	if user == nil {
		err = protocol.NewServerError(status.ErrAccountNotExists)
		return nil , err
	}

	resp.UserInfo = &rpb.UserInfo{
		UserId:               user.UserID,
		Account:              user.Account,
		Nickname:             user.Nickname,
		Gender:               rpb.Gender(user.Gender),
	}
	return resp, nil
}
