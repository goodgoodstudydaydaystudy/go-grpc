package server

import (
	"context"

	"goodgoodstudy.com/go-grpc/client/account"
	pb "goodgoodstudy.com/go-grpc/pkg/pb/logic/user"
)

type UserLogic struct {
	accountClient *account.Client
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
func (s *UserLogic) Login(ctx context.Context, req *pb.LoginReq) (resp *pb.LoginResp, err error) {
	resp = &pb.LoginResp{}

	// 1. check password
	r, err := s.accountClient.Login(ctx, req.Account, req.Password)
	if err != nil {
		return
	}

	// 2. get balance
	//balanceResp, err := s.walletClient.GetUserBalance(ctx, r.UserInfo.UserId)
	//if err != nil {}

	resp.UserInfo = &pb.UserInfo{
		UserId:   r.UserInfo.UserId,
		Account:  r.UserInfo.Account,
		Nickname: r.UserInfo.Nickname,
		Gender:   pb.Gender(r.UserInfo.Gender),
	}
	//resp.Balance = ...

	// 3. return
	return
}
