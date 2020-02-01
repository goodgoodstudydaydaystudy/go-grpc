package dao

import (
	"goodgoodstudy.com/go-grpc/protocol/common/status"

	pb "goodgoodstudy.com/go-grpc/pkg/pb/wallet"
	protocol "goodgoodstudy.com/go-grpc/pkg/procotol"
	"goodgoodstudy.com/go-grpc/pkg/server/wallet/dao/mysql"
)

type WalletDao interface {
	// 增
	Recharge(req *pb.RechargeReq) protocol.ServerError

	// 查
   	GetUserBalance(userId uint32) (uint64, protocol.ServerError)
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