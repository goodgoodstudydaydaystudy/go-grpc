package wallet

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	pb "goodgoodstudy.com/go-grpc/pkg/pb/server/wallet"
	protocol "goodgoodstudy.com/go-grpc/pkg/procotol"
	redisDb "goodgoodstudy.com/go-grpc/pkg/server/wallet/dao/redis"
	"goodgoodstudy.com/go-grpc/protocol/common/status"

	"github.com/go-redis/redis"
	"github.com/jmoiron/sqlx"
)

type server struct {
	db   *StoreManager
	rdb  *redis.Client
	timeManager
}

type timeManager struct {
}

func NewWalletServer() (*server, error) {
	// create mysql&redis conn
	db, err := sqlx.Connect("mysql", "root:8918112lu@/goodStudy")
	if err != nil {
		log.Println("NewWallet conn failed: ", err)
		return nil, err
	}

	redisConn := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	Store := &StoreManager{
		mysqlConn:   db,
	}

	walletServer := &server{
		db:  Store,
		rdb: redisConn,
	}

	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGINT, syscall.SIGCONT)

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
	//expiredOrderList, err := s.db.GetExpiredOrder(deadline)
	expiredOrderList, err := redisDb.NewWalletRedis(s.rdb).GetExpiredOrder(deadline)
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
	// 写入充值用户的信息（
	err = redisDb.NewWalletRedis(s.rdb).Recharge(req.UserId, req.Amount, s.timeManager.rechargeTime())
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
func (s *server) GetTopTenUser(ctx context.Context, req *pb.GetTopTenUserReq) (*pb.GetTopTenUserResp, error) {
	z, err := redisDb.NewWalletRedis(s.rdb).GetTopUserData(uint(req.Top), s.timeManager.rechargeTime())
	if err != nil {
		log.Println("server GetTopUser failed:", err)
		return nil, err
	}

	topUser := make(map[string]uint64)
	for _, val := range z {
		topUser[val.Member.(string)] = uint64(val.Score)
	}

	var topUserList string
	for key, _ := range topUser {
		log.Println("member:", key)
		topUserList += key + ","
	}

	//TODO topUser的充值金额要附上

	return &pb.GetTopTenUserResp{
		UserList:  topUserList,
	}, nil
}

// 没有付款的订单
func (s *server) WriteNoPaidOrder(ctx context.Context, req *pb.WriteNoPaidOrderReq) (*pb.WriteNoPaidOrderResp, error) {
	err := s.db.WriteNoPaidOrder(ctx, req.UserId, req.OrderId)
	if err != nil {
		log.Println("server OrderNotPay failed:", err)
		return nil, err
	}

	// 查询获取TopTenUserList
	err = redisDb.NewWalletRedis(s.rdb).WriteOrderDeadline(req.OrderId, s.timeManager.deadline())
	if err != nil {
		log.Println("wallet server RecordOrderDeadline failed:", err)
		return nil, err
	}
	return &pb.WriteNoPaidOrderResp{}, nil
}

// 支付订单后，orderId写入库。
func (s *server) Pay(ctx context.Context, req *pb.PayReq) (*pb.PayResp, error) {
	err := s.db.Pay(ctx, req.OrderId)
	if err != nil {
		log.Println("server pay failed:", err)
		return nil, err
	}
	return &pb.PayResp{}, nil
}


func (t *timeManager)rechargeTime() string {
	return time.Now().Format("2006-01-02 15")
}

func (t *timeManager) OrderTime() (orderTime string) {
	return time.Now().Format("2006-01-02 15:04:05")
}

func (t *timeManager) deadline() string {
	now := time.Now()
	afterOneMin, _ := time.ParseDuration("+1m")
	r := now.Add(afterOneMin)
	deadline := r.Format("2006-01-02 15:04:05")
	return deadline
}