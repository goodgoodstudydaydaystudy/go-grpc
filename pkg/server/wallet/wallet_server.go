package wallet

import (
	"context"
	"goodgoodstudy.com/go-grpc/protocol/common/status"
	"log"

	pb "goodgoodstudy.com/go-grpc/pkg/pb/wallet"
	protocol "goodgoodstudy.com/go-grpc/pkg/procotol"
	"goodgoodstudy.com/go-grpc/pkg/server/wallet/dao"
)

type server struct {
	db dao.WalletDao
}

func NewWalletServer() (*server, error) {
	db, err := dao.NewWalletDao("mysql")
	return &server{
		db: db,
	}, err
}

// 充值 返回 余额
func (s *server) Recharge(ctx context.Context, req *pb.RechargeReq) (*pb.RechargeResp, error) {
	err :=  s.db.Recharge(req)
	if err != nil {
		log.Println("server Recharge failed: ", err)
		return nil, protocol.NewServerError(status.ErrRechargeFailed)
	}
	return &pb.RechargeResp{
		Balance: 0,
	}, nil
}


// 查询 用户余额
func (s *server) GetUserBalance(ctx context.Context, req *pb.GetUserBalanceReq) (resp *pb.GetUserBalanceResp, err error)  {
	userBalance, err := s.db.GetUserBalance(req.GetUserId())
	if err != nil {
		log.Println("server GetUserBalance failed: ", err)
		return nil, protocol.NewServerError(status.ErrGetUserBalanceFailed)
	}
	return &pb.GetUserBalanceResp{
		Balance: userBalance,
	}, err
}