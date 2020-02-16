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
	db  *mysqlStoreManager
	rdb *redisStoreManager
}

func NewWalletServer() (*server, error) {
	// create mysql conn
	db, err := sqlx.Connect("mysql", "root:8918112lu@/goodStudy")
	if err != nil {
		log.Println("NewWallet conn failed: ", err)
		return nil, err
	}

	redisClient := NewRedisClient()

	mysqlStore := &mysqlStoreManager{mysqlConn: db}
	redisStore := &redisStoreManager{redisClient:redisClient}

	go NewTimer(mysqlStore)	// timer for scanning t_order

	return &server{
		db:  mysqlStore,
		rdb: redisStore,
	}, err
}


func NewRedisClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	return client
}


func NewTimer(db *mysqlStoreManager)  {
	ctx := context.Background()

	for {
		NotPayList, err := db.ScanNotPay(ctx)
		if err != nil {
			log.Println("server GetNotPayTimer failed:", err)
			return
		}
		if len(NotPayList) != 0 {
			log.Println("NotPayList:", NotPayList)
		}else {
			log.Println("NotPayList is empty")
		}
		time.Sleep(time.Duration(5)*time.Second)
	}
}

// 充值 返回 余额
func (s *server) Recharge(ctx context.Context, req *pb.RechargeReq) (*pb.RechargeResp, error) {
	err := s.db.Recharge(ctx, req.UserId, req.Amount)
	if err != nil {
		log.Println("server Recharge failed: ", err)
		return nil, protocol.NewServerError(status.ErrRechargeFailed)
	}
	err = s.rdb.Recharge(req.Account, req.Amount)
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
	r, err := s.rdb.GetTopUser(uint(req.Top))
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
func (s *server) OrderNotPay(ctx context.Context, req *pb.OrderNotPayReq) (*pb.OrderNotPayResp, error) {
	orderId, err := s.db.OrderNotPay(ctx, req.UserId)
	if err != nil {
		log.Println("server OrderNotPay failed:", err)
		return nil, err
	}

	return &pb.OrderNotPayResp{
		OrderId: orderId,
	}, nil
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
