package wallet

import (
	"context"

	"github.com/jmoiron/sqlx"

	"goodgoodstudy.com/go-grpc/protocol/common/status"
	"log"

	pb "goodgoodstudy.com/go-grpc/pkg/pb/server/wallet"
	protocol "goodgoodstudy.com/go-grpc/pkg/procotol"
)

type server struct {
	db *storeManager
}

func NewWalletServer() (*server, error) {
	// create mysql conn
	db, err := sqlx.Connect("mysql", "root:8918112lu@/goodStudy")
	if err != nil {
		log.Println("NewWallet conn failed: ", err)
		return nil, err
	}

	store := &storeManager{mysqlConn: db}
	return &server{
		db: store,
	}, err
}

// 充值 返回 余额
func (s *server) Recharge(ctx context.Context, req *pb.RechargeReq) (*pb.RechargeResp, error) {
	err := s.db.Recharge(ctx, req.UserId, req.Count)
	if err != nil {
		log.Println("server Recharge failed: ", err)
		return nil, protocol.NewServerError(status.ErrRechargeFailed)
	}
	return &pb.RechargeResp{
	}, nil
}

// 查询 用户余额
func (s *server) GetUserBalance(ctx context.Context, req *pb.GetUserBalanceReq) (resp *pb.GetUserBalanceResp, err error) {
	userBalance, err := s.db.GetUserBalance(ctx, req.GetUserId())
	if err != nil {
		log.Println("server GetUserBalance failed: ", err)
		return nil, protocol.NewServerError(status.ErrGetUserBalanceFailed)
	}
	return &pb.GetUserBalanceResp{
		Balance: userBalance,
	}, err
}
