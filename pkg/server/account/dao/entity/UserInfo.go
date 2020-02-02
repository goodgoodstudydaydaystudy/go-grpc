package account

import (
	pb "goodgoodstudy.com/go-grpc/pkg/pb/server/account"
)

type UserInfo struct {
	UserID   uint32 `db:"id"` // 这玩意叫tag, 反射的时候可以获得tag信息, sqlx可以根据这个tag把和列名和字段名对应上
	Account  string `db:"account"`
	Password string `db:"password"`
	Nickname string `db:"name"`
	Gender   uint32 `db:"gender"` // TODO 要建立一个映射, 把uint32和pb中gender的enum类型对应上, 实际上强制转换也是可以的
}

func UserInfoToPb(userInfo *UserInfo) *pb.UserInfo {
	return &pb.UserInfo{
		UserId:   userInfo.UserID,
		Account:  userInfo.Account,
		Nickname: userInfo.Nickname,
		Gender:   pb.Gender(userInfo.Gender),
	}
}

func PbToUserInfo(pbInfo *pb.UserInfo) *UserInfo {
	return &UserInfo{
		UserID:   pbInfo.UserId,
		Account:  pbInfo.Account,
		Nickname: pbInfo.Nickname,
		Gender:   uint32(pbInfo.Gender),
	}
}
