package wallet

import (
	"context"
	"log"
	"time"

	pb "goodgoodstudy.com/go-grpc/pkg/pb/server/wallet"
	protocol "goodgoodstudy.com/go-grpc/pkg/procotol"
	"goodgoodstudy.com/go-grpc/protocol/common/status"

	"github.com/go-redis/redis"
	"github.com/jmoiron/sqlx"
)

type server struct {
	db *StoreManager
}


func NewWalletServer() (*server, error) {
	// create mysql&redis conn
	db, err := sqlx.Connect("mysql", "root:8918112lu@/goodStudy")
	if err != nil {
		log.Println("NewWallet conn failed: ", err)
		return nil, err
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	Store := &StoreManager{
		mysqlConn:   db,
		redisClient: rdb,
	}

	walletServer := &server{
		db:  Store,
	}

	go walletServer.scanAndQueryOrderNoPaid()
	go walletServer.expiredOrderToMark(walletServer.getExpiredOrder)

	return walletServer, nil
}


// 获取未支付的订单
func (s *server) scanAndQueryOrderNoPaid()  {
	ctx := context.Background()
	for {
		NoPaidList, err := s.db.ScanNoPaid(ctx)
		if err != nil {
			log.Println("server GetNotPayTimer failed:", err)
			return
		}
		if len(NoPaidList) != 0 {
			log.Println("NotPayList:", NoPaidList)
			// NoPaidList 提供给 app/client核实订单
		}else {
			log.Println("NotPayList is empty")
		}
		time.Sleep(time.Duration(5)*time.Second)
	}
}


// go查询redis的expired订单
func (s *server) getExpiredOrder() []string {
	t := time.Now()
	deadline := t.Format("2006-01-02 15:04:05")
	expiredOrderList, err := s.db.GetExpiredOrder(deadline)
	if err != nil {
		log.Println("wallet server GetExpiredOrder failed:", err)
		return nil
	}
	return expiredOrderList
}


// expired订单发送给db MarkExpiredOrder 修改状态
func (s *server) expiredOrderToMark(fn func() []string){
	ctx := context.Background()
	for {
		expiredOrderList := fn()
		if len(expiredOrderList) == 0 {
			log.Println("expiredOrderList is empty")
		} else {
			err := s.db.MarkExpiredOrder(ctx, expiredOrderList)
			if err != nil {
				log.Println("wallet server expiredOrderToMark failed:", err)
				return
			}
		}
		time.Sleep(time.Duration(1)*time.Second)
	}
}


// 充值 返回 余额
func (s *server) Recharge(ctx context.Context, req *pb.RechargeReq) (*pb.RechargeResp, error) {
	err := s.db.Recharge(ctx, req.UserId, req.Amount)
	if err != nil {
		log.Println("server Recharge failed: ", err)
		return nil, protocol.NewServerError(status.ErrRechargeFailed)
	}
	err = s.db.Consume(req.UserId, req.Amount)
	if err != nil {
		log.Println("server Recharge failed: ", err)
		return nil, protocol.NewServerError(status.ErrRechargeFailed)
	}
	return &pb.RechargeResp{
	}, nil
}


// 查询 用户余额
func (s *server) GetUserBalance(ctx context.Context, req *pb.GetUserBalanceReq) (resp *pb.GetUserBalanceResp, err error) {
	userBalance, err := s.db.GetUserBalance(ctx, req.GetUserId())
	if err != nil {
		log.Println("server GetUserBalance failed: ", err)
		return nil, protocol.NewServerError(status.ErrGetUserBalanceFailed)
	}
	return &pb.GetUserBalanceResp{
		Balance: userBalance,
	}, err
}


// 获取 top 用户
func (s *server) GetTopUser(ctx context.Context, req *pb.GetTopUserReq) (*pb.GetTopUserResp, error) {
	r, err := s.db.GetTopUser(uint(req.Top))
	if err != nil {
		log.Println("server GetTopUser failed:", err)
		return nil, err
	}
	var topUserList string
	for key, _ := range r {
		log.Println("member:", key)
		topUserList += key + ","
	}

	//TODO topUser的充值金额要附上

	return &pb.GetTopUserResp{
		UserList:  topUserList,
	}, nil
}

// 没有付款的订单
func (s *server) RecordOrderNoPaid(ctx context.Context, req *pb.RecordOrderNoPaidReq) (*pb.RecordOrderNoPaidResp, error) {
	err := s.db.RecordOrderNoPaid(ctx, req.UserId, req.OrderId)
	if err != nil {
		log.Println("server OrderNotPay failed:", err)
		return nil, err
	}

	return &pb.RecordOrderNoPaidResp{}, nil
}

// pay
func (s *server) Pay(ctx context.Context, req *pb.PayReq) (*pb.PayResp, error) {
	err := s.db.Pay(ctx, req.OrderId)
	if err != nil {
		log.Println("server pay failed:", err)
		return nil, err
	}
	return &pb.PayResp{}, nil
}
