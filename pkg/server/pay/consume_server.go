package pay

import (
	"context"
	pb "goodgoodstudy.com/go-grpc/pkg/pb/pay"
	"goodgoodstudy.com/go-grpc/pkg/server/account/db"
	"math/rand"
	"time"
)

type controlServer struct {
}

// 订单
func (s *controlServer) Pay(ctx context.Context, consumeReq *pb.ConsumeReq) (*pb.ConsumeResp, error) {
	// 消费成功后，返回订单号
	rand.Seed(time.Now().Unix())
	rnd := rand.Int63n(10)
	return &pb.ConsumeResp{OrderId: rnd, Message:"consume success"}, nil // 返回Resp里的字段？
}


func NewConsumeServer() (*controlServer, error) {
	_, err := db.Conn()
	return &controlServer{}, err
}