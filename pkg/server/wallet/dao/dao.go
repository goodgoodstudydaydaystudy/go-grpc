package dao

import (
	"context"

	"goodgoodstudy.com/go-grpc/protocol/common/status"

	pb "goodgoodstudy.com/go-grpc/pkg/pb/wallet"
	protocol "goodgoodstudy.com/go-grpc/pkg/procotol"
	"goodgoodstudy.com/go-grpc/pkg/server/wallet/dao/mysql"
)

type WalletDao interface {
	// 增
	Recharge(ctx context.Context, req *pb.RechargeReq) protocol.ServerError

	// 查
	GetUserBalance(ctx context.Context, userId uint32) (uint64, protocol.ServerError)
}

func NewWalletDao(dbType string) (WalletDao, error) {
	switch dbType {
	case "mysql":
		return mysql.NewWalletMysql()
	case "mongodb":
		return nil, protocol.NewServerError(status.ErrDBTypeNotSupport)
	}
	return nil, protocol.NewServerError(status.ErrDBTypeNotSupport)
}
