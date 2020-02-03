package server

import (
	"context"
	"google.golang.org/genproto/googleapis/datastore/v1"
	"log"

	"goodgoodstudy.com/go-grpc/client/account"
	"goodgoodstudy.com/go-grpc/client/wallet"
	pb "goodgoodstudy.com/go-grpc/pkg/pb/logic/user"
)

type UserLogic struct {
	accountClient *account.Client
	walletClient  *wallet.Client
}

func NewUserLogic() (*UserLogic, error) {
	accountClient, err := account.NewAccountClient()
	if err != nil {
		return nil, err
	}

	return &UserLogic{
		accountClient: accountClient,
	}, nil
}

// 2. account的登录服务
func (s *UserLogic) CheckUserPassword(ctx context.Context, req *pb.CheckUserPwdReq) (resp *pb.CheckUserPwdResp, err error) {
	resp = &pb.CheckUserPwdResp{}

	// 1. check password
	r, err := s.accountClient.Login(ctx, req.Account, req.Password)
	if err != nil {
		log.Println("logic.Login check password failed: ", err)
		return nil, err
	}
	resp.UserInfo = &pb.UserInfo{
		UserId:   r.UserInfo.UserId,
		Account:  r.UserInfo.Account,
		Nickname: r.UserInfo.Nickname,
		Gender:   pb.Gender(r.UserInfo.Gender),
	}

	// 2. get balance
	userBalance, err := s.walletClient.GetUserById(ctx, resp.UserInfo.UserId)
	if err != nil {
		log.Println("logic.Login get balance failed: ", err)
		return nil, err
	}
	resp.Balance = userBalance.Balance
	// 3. return
	return resp, nil
}

// 3.  account的register服务
func (s *UserLogic) Register(ctx context.Context, req *pb.RegisterReq) (resp *pb.RegisterResp, err error) {
	resp = &pb.RegisterResp{}
	// 3.1 提交注册信息
	r, err := s.accountClient.Register(ctx, req.Account, req.Password, req.Nickname, int(req.Gender))
	if err != nil {
		log.Println("logic Register register failed: ", err)
		return
	}
	resp.UserId = r.UserId
	// 3.2 return
	return resp, nil
}

