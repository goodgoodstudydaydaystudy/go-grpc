package redis

import (
	"github.com/go-redis/redis"
	protocol "goodgoodstudy.com/go-grpc/pkg/procotol"
	"goodgoodstudy.com/go-grpc/protocol/common/status"
	"log"
)

type Client struct {
	cli *redis.Client
}

func NewWalletRedis(cli *redis.Client) *Client {
	return &Client{
		cli: cli,
	}
}

// 写入数据{account, amount, time}
func (c *Client) Recharge(userId uint32, amount int64, time string) protocol.ServerError {
	zAmount := float64(amount)
	z := redis.Z{
		Score:  zAmount,
		Member: userId,
	}
	err := c.cli.ZAdd(time, &z).Err()
	if err != nil {
		log.Println("redis recharge failed:", err)
		return protocol.NewServerError(status.ErrRedisDB)
	}
	return nil
}

// 读取数据{top10 []z}
func (c *Client) GetTopData(n uint, time string) ([]redis.Z, protocol.ServerError) {
	stop := int64(n)
	z, err := c.cli.ZRangeWithScores(time, 0, stop).Result()
	if err == redis.Nil {
		log.Println("no data in redis")
		return nil, protocol.NewServerError(status.ErrGetTopUserFailed)
	}else if err != nil {
		log.Println("redis GetTopN failed:", err)
		return nil, protocol.NewServerError(status.ErrRedisDB)
	}
	return z, nil
}

// 写入order的orderId和deadline
// 用于写入订单的deadline
func (c *Client) RecordOrderDeadline(orderId string, deadline string) protocol.ServerError {
	err := c.cli.SAdd(deadline, orderId).Err()
	//err := c.cli.Do("SAdd", deadline, orderId).Err()
	if err != nil {
		log.Println("redis RecordOrderDeadline failed:", err)
		return protocol.NewServerError(status.ErrRedisDB)
	}
	return nil
}

// 返回过期订单
func (c *Client) GetExpiredOrder(deadline string) ([]string, protocol.ServerError){
	s, err := c.cli.SMembers(deadline).Result()
	if err != nil {
		log.Println("redis GetDeadlineOrder failed:", err)
		return nil, protocol.NewServerError(status.ErrRedisDB)
	}
	return s, nil
}