package dao

import (
	"fmt"
	rpb "goodgoodstudy.com/go-grpc/pkg/pb/account"
	account "goodgoodstudy.com/go-grpc/pkg/server/account/dao/entity"

	"goodgoodstudy.com/go-grpc/pkg/procotol"
	"goodgoodstudy.com/go-grpc/pkg/server/account/dao/mysql"
)

type AccountDao interface {
	InsertInfo(req *rpb.RegisterReq) error
	QueryInfo(account string) (int32, string, error)
	GetUserById(userId uint32) (*account.UserInfo, error)
}

func NewAccountDao(dbType string) (AccountDao, error) {
	switch dbType {
	case "mysql":
		return mysql.NewAccountMysql()
	case "sqlite":
		return nil, protocol.NewServerError(-2, "sqlite is not supported yet.")
	case "mongodb":
		return nil, protocol.NewServerError(-2, "mongodb is not supported yet.")
	}
	return nil, protocol.NewServerError(-2, fmt.Sprintf("wrong dbType %s", dbType))
}
