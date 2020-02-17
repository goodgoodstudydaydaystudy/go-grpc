package wallet

import (
	"context"
	"log"
	"time"

	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql" // mysql driver
	"github.com/jmoiron/sqlx"
	redis2 "goodgoodstudy.com/go-grpc/pkg/server/wallet/dao/redis"

	protocol "goodgoodstudy.com/go-grpc/pkg/procotol"
	"goodgoodstudy.com/go-grpc/pkg/server/wallet/dao/mysql"
	"goodgoodstudy.com/go-grpc/protocol/common/status"
)

type StoreManager struct {
	mysqlConn   *sqlx.DB
	redisClient *redis.Client
	timeManager
}

type timeManager struct {
}

func (st *StoreManager) Recharge(ctx context.Context, userId uint32, deltaAdd int64) protocol.ServerError {
	// 1. 开启事务
	txErr := doTx(ctx, st.mysqlConn, func(tx *sqlx.Tx) error {
		dao := mysql.NewWalletMysql(tx)
		// 2. 加锁, 查余额
		balance, err := dao.GetUserBalance(ctx, userId, true)
		// 3. 判断充值后, 余额是否足够/溢出
		if balance+deltaAdd < 0 {
			return protocol.NewServerError(status.ErrBalanceNotEnough)
		}
		// 4. 充值/消费
		err = dao.Recharge(ctx, userId, balance+deltaAdd)
		// 5. rollback/commit & return
		return err

	})
	return protocol.ToServerError(txErr)

}

func (st *StoreManager) GetUserBalance(ctx context.Context, userId uint32) (int64, protocol.ServerError) {
	dao := mysql.NewWalletMysql(st.mysqlConn)
	return dao.GetUserBalance(ctx, userId, false)
}

func (st *StoreManager) RecordOrderNoPaid(ctx context.Context, userId uint32, orderId string) protocol.ServerError {
	// 开启事务
	db := st.mysqlConn
	tx, err := db.BeginTxx(ctx, nil)
	if err != nil {
		return protocol.ToServerError(err)
	}
	// 调用 db方法
	dao := mysql.NewWalletMysql(tx)
	err = dao.RecordOrderNoPaid(ctx, userId, orderId, st.rechargeTime(), st.deadline())
	if err != nil {
		_ = tx.Rollback()
		return protocol.ToServerError(err)
	} else {
		_ = tx.Commit()
	}
	// 写入redis RecordOrderDeadline
	Dao2 := redis2.NewWalletRedis(st.redisClient)
	re := Dao2.RecordOrderDeadline(orderId, st.deadline())
	if re != nil {
		return protocol.ToServerError(re)
	}
	// return
	return nil
}

func (st *StoreManager) Pay(ctx context.Context, orderId string) protocol.ServerError {
	txErr := doTx(ctx, st.mysqlConn, func(tx *sqlx.Tx) error {
		dao := mysql.NewWalletMysql(tx)
		err := dao.Pay(ctx, orderId)
		// rollback/commit &return
		return err
	})
	return protocol.ToServerError(txErr)
}

func (st *StoreManager) ScanNoPaid(ctx context.Context) ([]string, protocol.ServerError) {
	tx := st.mysqlConn
	dao := mysql.NewWalletMysql(tx)
	NotPays, err := dao.GetNoPaid(ctx)
	if err != nil {
		return nil, err
	}
	return NotPays, nil
}

func (st *StoreManager) MarkExpiredOrder(ctx context.Context, expiredOder []string) protocol.ServerError {
	// tx
	log.Println("st MarkExpiredOrder")
	txErr := doTx(ctx, st.mysqlConn, func(tx *sqlx.Tx) error {
		dao := mysql.NewWalletMysql(tx)
		err := dao.MarkExpiredOrder(ctx, expiredOder)
		return err
	})
	// return rollback/commit
	return protocol.ToServerError(txErr)
}

// private
// 开启事务, 封装好一个函数, 方便别的地方用
// 函数签名包含一个匿名函数，而且可以在函数最后return，mark
func doTx(ctx context.Context, db *sqlx.DB, fn func(tx *sqlx.Tx) error) error {
	tx, err := db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			_ = tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	return fn(tx)
}

// 管理time和写入recharge recording
func (st *StoreManager) Consume(orderId uint32, amount int64) protocol.ServerError {
	dao := redis2.NewWalletRedis(st.redisClient)
	err:= dao.Recharge(orderId, amount, st.rechargeTime())
	if err != nil {
		return err
	}
	return nil
}


// getTop列表，生成TopN的map返回给server
func (st *StoreManager) GetTopUser(n uint) (map[string]uint64, protocol.ServerError) {
	dao := redis2.NewWalletRedis(st.redisClient)
	z, err := dao.GetTopData(n, st.rechargeTime())
	if err != nil {
		log.Println("st GetTopUser failed:", err)
		return nil, err
	}

	topUser := make(map[string]uint64)
	for _, val := range z {
		topUser[val.Member.(string)] = uint64(val.Score)
	}
	return topUser, nil
}

// 通过时间 返回 过期订单
func (st *StoreManager) GetExpiredOrder(deadline string) ([]string, protocol.ServerError) {
	dao := redis2.NewWalletRedis(st.redisClient)
	expiredOrder, err := dao.GetExpiredOrder(deadline)
	if err != nil {
		return nil, protocol.ToServerError(err)
	}
	return expiredOrder, nil
}


func (t *timeManager)rechargeTime() string {
	return time.Now().Format("2006-01-02 15")
}

func (t *timeManager) RecordOrderTime() (orderTime string) {
	return time.Now().Format("2006-01-02 15:04:05")
}

func (t *timeManager) deadline() string {
	now := time.Now()
	afterOneMin, _ := time.ParseDuration("+1m")
	r := now.Add(afterOneMin)
	deadline := r.Format("2006-01-02 15:04:05")
	return deadline
}




