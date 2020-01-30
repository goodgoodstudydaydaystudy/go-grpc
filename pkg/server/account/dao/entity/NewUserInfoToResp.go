package account

import (
	rpb "goodgoodstudy.com/go-grpc/pkg/pb/account"
	"goodgoodstudy.com/go-grpc/pkg/server/account/dao"
)

func NewUserInfoToResp(queryId uint32, db dao.AccountDao) (*rpb.UserInfo, error) {
	userInfo, err := db.GetUserById(queryId)
	return &rpb.UserInfo{
		UserId:               userInfo.UserID,
		Account:              userInfo.Account,
		Nickname:             userInfo.Nickname,
	}, err
}