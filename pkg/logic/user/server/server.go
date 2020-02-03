package server

import (
	"context"
	"log"

	rpb "goodgoodstudy.com/go-grpc/pkg/pb/server/account"
	lpb "goodgoodstudy.com/go-grpc/pkg/pb/server/logic"
	pb "goodgoodstudy.com/go-grpc/pkg/pb/server/wallet"
)

//1. 取得accountServer和walletServer的方法
type accountInterface interface {
	// 注册
	Register(ctx context.Context, req *rpb.RegisterReq) (*rpb.RegisterResp, error)

	// 登录
	Login(ctx context.Context, req *rpb.LoginReq) (resp *rpb.LoginResp, err error)
}

type walletInterface interface {
	// 余额
	GetUserBalance(ctx context.Context, req *pb.GetUserBalanceReq) (resp *pb.GetUserBalanceResp, err error)
}


type UserLogic struct {
	accountInterface
	walletInterface
}

// 2. account的登录服务
func (s *UserLogic) AccountLogin(ctx context.Context, req *lpb.AccountSerReq) (*lpb.AccountSerResp, error) {
	// 2.1 获得LoginReq的实例
	loginReq := &rpb.LoginReq{
		Account:		req.GetAccount(),
		Password:		req.GetPassword(),
	}
	getUserInfo, err := s.Login(ctx, loginReq)
	if err != nil {
		log.Println("logic AccountLogin failed: ", err)
		return nil, err
	}
	// 2.2 不知道怎么处理userInfo
	userInfo := getUserInfo.UserInfo

	// 2.3 使用userId查询余额
	userBalance := &pb.GetUserBalanceReq{
		UserId:		userInfo.UserId,
	}
	getUserBalance, err := s.GetUserBalance(ctx, userBalance)
	if err != nil {
		log.Println("AccountLogin getUserBalance failed: ", err)
		return nil, err
	}

	// 2.4 return
	return &lpb.AccountSerResp{
		UserInfo: 	userInfo,
		Balance:	getUserBalance.Balance},nil
}












































}
