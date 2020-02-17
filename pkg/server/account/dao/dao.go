package dao

import (
	rpb "goodgoodstudy.com/go-grpc/pkg/pb/server/account"
	"goodgoodstudy.com/go-grpc/pkg/procotol"
	account "goodgoodstudy.com/go-grpc/pkg/server/account/dao/entity"
	"goodgoodstudy.com/go-grpc/pkg/server/account/dao/mysql"
	"goodgoodstudy.com/go-grpc/protocol/common/status"
)

// 接口按功能排好序, 方便别人看
type AccountDao interface {
	// 查
	GetUserPasswordByAccount(acc string) (string, protocol.ServerError)
	GetUserByAccount(acc string) (*account.UserInfo, protocol.ServerError) // 参数不要和包名冲突, 已经有叫account的包了, 所以要么参数改名, 要么包重命名一下, 建议参数改名, 叫acc或者accStr都可以
	GetUserById(userId uint32) (*account.UserInfo, protocol.ServerError)

	// 增
	AddUsr(req *rpb.AddUserReq) protocol.ServerError
}

func NewAccountDao(dbType string) (AccountDao, error) {
	switch dbType {
	case "mysql":
		return mysql.NewAccountMysql()
	case "sqlite":
		return nil, protocol.NewServerError(status.ErrDBTypeNotSupport)
	case "mongodb":
		return nil, protocol.NewServerError(status.ErrDBTypeNotSupport)
	}
	return nil, protocol.NewServerError(status.ErrDBTypeNotSupport)
}
