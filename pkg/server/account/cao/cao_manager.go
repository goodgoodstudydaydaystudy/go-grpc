package cao

import (
	"log"
	"time"

	t "goodgoodstudy.com/go-grpc/pkg/foundation/time_manager"
	rpb "goodgoodstudy.com/go-grpc/pkg/pb/server/account"
	protocol "goodgoodstudy.com/go-grpc/pkg/procotol"
	"goodgoodstudy.com/go-grpc/pkg/server/account/cao/cache"
	account "goodgoodstudy.com/go-grpc/pkg/server/account/dao/entity"

	"github.com/go-redis/redis"
)

var tempCache = &account.UserInfo{
	UserID:   -1,
	Account:  "",
	Password: "",
	Nickname: "",
	Gender:   0,
}

// create accountCache struct
type AccountCache struct {
	Cli *redis.Client
}

// check pwdCache and return isValid to server
func (c *AccountCache) GetUserInfoByID(userId int32) (*rpb.UserInfo, bool, protocol.ServerError){
	//log.Println("get req:", userId)
	cao := cache.NewRedis(c.Cli)
	userInfoMap, err := cao.GetUserInfo(ToString(userId))
	//log.Println("get map:", userInfoMap)
	if err != nil {
		log.Println("GetUserInfo failed:", err)
	}
	// have no cache
	//log.Println("userInfoMap:", userInfoMap)
	if len(userInfoMap) == 0 {
		//log.Println("len(userInfoMap) == 0")
		return nil, false, nil
	}

	userInfo, err := UserInfoMapToStruct(userInfoMap)
	if err != nil {
		log.Println("ca GetUserInfoByID failed:", err)
		return nil, false, nil
	}

	return userInfo, true, nil
}

// if invalid write id&pwd into cache and return
func (c *AccountCache) WriteIntoCache (userInfo *account.UserInfo, userId ...int32) protocol.ServerError {
	cao := cache.NewRedis(c.Cli)
	var expiredTime time.Duration
	//log.Println("cm WriteIntoCache userInfo:", userInfo)

	if userInfo == nil {
		userInfo = tempCache
		expiredTime = t.ExpiredTimeSort()
	}

	expiredTime = t.ExpiredTime()
	//log.Println("cm WriteIntoCache userInfo:", userInfo)
	localUserInfo := UserInfoPointToStrut(userInfo)
	//log.Println("cm WriteIntoCache userId:", userId)
	//log.Println("cm WriteIntoCache userId:[0]", userId[0])
	//log.Println("cm WriteIntoCache localUserInfo:", localUserInfo)
	err := cao.WriteUserInfo(ToString(userId[0]), StructToMap(localUserInfo), expiredTime)
	if err != nil {
		log.Println("WriteIntoCache:", err)
		return protocol.ToServerError(err)
	}

	return nil
}

