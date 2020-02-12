package wallet

import (
	"context"
	"log"

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