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

type mysqlStoreManager struct {
	mysqlConn *sqlx.DB
}

type redisStoreManager struct {
	redisClient *redis.Client
}


func (st *mysqlStoreManager) Recharge(ctx context.Context, userId uint32, deltaAdd int64) protocol.ServerError {
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

func (st *mysqlStoreManager) GetUserBalance(ctx context.Context, userId uint32) (int64, protocol.ServerError) {
	dao := mysql.NewWalletMysql(st.mysqlConn)
	return dao.GetUserBalance(ctx, userId, false)
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
func (st *redisStoreManager) Recharge(account string, amount int64) protocol.ServerError {
	dao := redis2.NewWalletRedis(st.redisClient)
	err:= dao.Recharge(account, amount, rechargeTime())
	if err != nil {
		return err
	}
	return nil
}


// getTop列表，生成TopN的map返回给server
func (st *redisStoreManager) GetTopUser(n uint) (map[string]uint64, protocol.ServerError) {
	dao := redis2.NewWalletRedis(st.redisClient)
	z, err := dao.GetTopData(n, rechargeTime())
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

func rechargeTime() string {
	return time.Now().Format("2006-01-02-15")
}