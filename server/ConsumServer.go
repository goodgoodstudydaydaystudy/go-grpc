package server

import (
	"context"
	"goodgoodstudy.com/go-grpc/pkg/DB"
	"goodgoodstudy.com/go-grpc/pkg/pb/Pay"
	"math/rand"
	"time"
)

type ControlServer struct {
}

// 订单
func (s *ControlServer) Pay(ctx context.Context, consumeReq *Pay.ConsumeReq) (*Pay.ConsumeResp, error) {

	// 消费成功后，返回订单号
	rand.Seed(time.Now().Unix())
	rnd := rand.Int63n(10)

	// 调用数据库
	err := DB.DBconn{}.InsertOrder(rnd, consumeReq.GetItemNum(), consumeReq.GetUserId(), 11)
	if err != nil {
		return &Pay.ConsumeResp{
			OrderId: rnd,
			Message: "consume failed",
		}, err
	}
	return &Pay.ConsumeResp{OrderId: rnd, Message:"consume success"}, nil // 返回Resp里的字段？
}


