package dao

import (
	"fmt"

	"goodgoodstudy.com/go-grpc/pkg/procotol"
	"goodgoodstudy.com/go-grpc/pkg/server/account/dao/mysql"
)

type AccountDao interface {
	InsertInfo(userId int32, account string, password string) error
	QueryInfo(account string) error
}

// TODO 返回 interface，但return 的方法去不在接口中。。
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
