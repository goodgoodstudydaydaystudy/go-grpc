package cao

import (
	"reflect"
	"strconv"

	rpb "goodgoodstudy.com/go-grpc/pkg/pb/server/account"
	protocol "goodgoodstudy.com/go-grpc/pkg/procotol"
	account "goodgoodstudy.com/go-grpc/pkg/server/account/dao/entity"
)

func ToString(ID int32) string {
	return strconv.Itoa(int(ID))
}

type info struct {
	UserID   int32 `json:"id"`
	Account  string `json:"account"`
	Password string `json:"password"`
	Nickname string `json:"name"`
	Gender   uint32 `json:"gender"`
}

// password unexported
func UserInfoPointToStrut(userInfo *account.UserInfo) info {
	localUserInfo := info{
		UserID:   userInfo.UserID,
		Account:  userInfo.Account,
		Password: userInfo.Password,
		Nickname: userInfo.Nickname,
		Gender:   userInfo.Gender,
	}

	return localUserInfo
}

func StructToMap(userInfo info) map[string]interface{} {
	t := reflect.TypeOf(userInfo)
	v := reflect.ValueOf(userInfo)
	//log.Println("userInfo:", userInfo)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).Interface()
	}

	return data
}

func UserInfoMapToStruct(data map[string]string) (*rpb.UserInfo, protocol.ServerError) {
	userId := data["UserId"]
	tmpGender := data["Gender"]

	ID, _ := strconv.Atoi(userId)
	gender, _ := strconv.Atoi(tmpGender)

	userInfo := rpb.UserInfo{
		UserId:    int32(ID),
		Account:   data["Account"],
		Password:  data["Password"],
		Nickname:  data["Nickname"],
		Gender:   rpb.Gender(gender),
	}

	return &userInfo, nil
}