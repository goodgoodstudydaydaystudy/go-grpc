package pay

import (
	"context"
	pb "goodgoodstudy.com/go-grpc/pkg/pb/pay"
	protocol "goodgoodstudy.com/go-grpc/pkg/procotol"
	"math/rand"
	"time"
)

type server struct {
}

func NewConsumeServer() (*server, error) {
	return &server{}, nil
}

// 订单
func (s *server) Pay(ctx context.Context, consumeReq *pb.ConsumeReq) (*pb.ConsumeResp, error) {
	// 消费成功后，返回订单号
	rand.Seed(time.Now().Unix())
	rnd := rand.Int63n(10)
	return &pb.ConsumeResp{OrderId: rnd}, protocol.NewServerError(-2000)
}
