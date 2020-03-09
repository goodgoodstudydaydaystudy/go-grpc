package cache

import (
	"log"
	"time"

	protocol "goodgoodstudy.com/go-grpc/pkg/procotol"
	"goodgoodstudy.com/go-grpc/protocol/common/status"

	"github.com/go-redis/redis"
)

type redisCache struct {
	cli *redis.Client
}

// write id-userInfo
// ready id return userInfo
func NewRedis(client *redis.Client) *redisCache {
	return &redisCache{
		cli: client,
	}
}

func (u *redisCache) GetUserInfo(userId string) (map[string]string, protocol.ServerError) {
	userInfo, err := u.cli.HGetAll(userId).Result()
	if err != nil {
		log.Println("redisCache GetUserInfo failed:", err)
		return nil, protocol.NewServerError(status.ErrPwdCache)
	}
	//log.Println("get cache:", userInfo)
	//log.Println("cache got userInfo:", userInfo)
	return userInfo, nil
}


// account different with nickname
func (u *redisCache) WriteUserInfo(userId string, userInfo map[string]interface{}, expired time.Duration) protocol.ServerError {
	_, err := u.cli.HMSet(userId, userInfo).Result()
	if err != nil {
		log.Println("redisCache WriteUserInfo failed:", err)
		return protocol.NewServerError(status.ErrPwdCache)
	}

	u.cli.Expire(userId, expired)

	return nil
}

