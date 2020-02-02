package walletdao

import (
	"context"

	"goodgoodstudy.com/go-grpc/protocol/common/status"

	protocol "goodgoodstudy.com/go-grpc/pkg/procotol"
)

type WalletDao interface {
	// 增
	Recharge(ctx context.Context, userId uint32, deltaAdd int64) protocol.ServerError

	// 查
	GetUserBalance(ctx context.Context, userId uint32, forUpdate bool) (int64, protocol.ServerError)
}

func NewWalletDao(dbType string) (WalletDao, error) {
	switch dbType {
	case "mysql":
		return nil, protocol.NewServerError(status.ErrDBTypeNotSupport, "dao暂时不用了先")
	case "mongodb":
		return nil, protocol.NewServerError(status.ErrDBTypeNotSupport)
	}
	return nil, protocol.NewServerError(status.ErrDBTypeNotSupport)
}
