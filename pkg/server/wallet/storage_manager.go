package wallet

import (
	"context"

	_ "github.com/go-sql-driver/mysql" // mysql driver
	"github.com/jmoiron/sqlx"

	protocol "goodgoodstudy.com/go-grpc/pkg/procotol"
	"goodgoodstudy.com/go-grpc/pkg/server/wallet/dao/mysql"
	"goodgoodstudy.com/go-grpc/protocol/common/status"
)

type storeManager struct {
	mysqlConn *sqlx.DB
}

// TODO
// 这个写法结构不是很理解
func (st *storeManager) Recharge(ctx context.Context, userId uint32, deltaAdd int64) protocol.ServerError {
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

func (st *storeManager) GetUserBalance(ctx context.Context, userId uint32) (int64, protocol.ServerError) {
	dao := mysql.NewWalletMysql(st.mysqlConn)
	return dao.GetUserBalance(ctx, userId, false)
}

// private
// 开启事务, 封装好一个函数, 方便别的地方用
// TODO
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
