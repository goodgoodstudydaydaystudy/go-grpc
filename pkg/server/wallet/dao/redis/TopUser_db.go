package redis

import (
	protocol "goodgoodstudy.com/go-grpc/pkg/procotol"
	"goodgoodstudy.com/go-grpc/protocol/common/status"
	"log"

	"github.com/go-redis/redis"
)


type Client struct {
	cli *redis.Client
}

func NewWalletRedis(cli *redis.Client) *Client {
	return &Client{
		cli:cli,
	}
}

// 写入数据{account, amount, time}
func (c *Client) Recharge(account string, amount int64, time string) protocol.ServerError {
	zAmount := float64(amount)
	z := redis.Z{
		Score:  zAmount,
		Member: account,
	}
	err := c.cli.ZAdd(time, &z).Err()
	if err != nil {
		log.Println("redis recharge failed:", err)
		return protocol.NewServerError(status.ErrDB)
	}
	return nil
}

// 读取数据{top10 []z}
func (c *Client) GetTopData(n uint, time string) ([]redis.Z, protocol.ServerError) {
	stop := int64(n)
	z, err := c.cli.ZRangeWithScores(time, 0, stop).Result()
	log.Println("Get z:", )
	if err == redis.Nil {
		log.Println("no data in redis")
		return nil, protocol.NewServerError(status.ErrGetTopUserFailed)
	}else if err != nil {
		log.Println("redis GetTopN failed:", err)
		return nil, protocol.NewServerError(status.ErrDB)
	}
	return z, nil
}
