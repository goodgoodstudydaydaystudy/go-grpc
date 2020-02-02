package mysql

import (
	"context"
	_ "database/sql"
	"log"
	"time"

	protocol "goodgoodstudy.com/go-grpc/pkg/procotol"
	"goodgoodstudy.com/go-grpc/protocol/common/status"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// queryExecer is generic interface which is used by both *sqlx.DB and *sqlx.Tx
type queryExec interface {
	sqlx.Execer
	sqlx.ExecerContext
	sqlx.Queryer
	sqlx.QueryerContext
	// sqlx.Preparer
	// sqlx.PreparerContext
}

type WalletMysql struct {
	qe queryExec
}

func NewWalletMysql(qe queryExec) (*WalletMysql, error) {
	return &WalletMysql{
		qe: qe,
	}, nil
}

func (c *WalletMysql) Recharge(ctx context.Context, userId uint32, deltaAdd int64) protocol.ServerError {
	now := time.Now()
	rechargeTime := now.Format("2006-01-02 15:04:05")

	rechargeExec := "INSERT INTO t_wallet VALUE(?, ?, ?) ON DUPLICATE KEY UPDATE money=VALUES(money), date=VALUES(date)"
	_, err := c.qe.ExecContext(ctx, rechargeExec, userId, deltaAdd, rechargeTime)
	if err != nil {
		log.Println("wallet Recharge failed: ", err)
		return protocol.NewServerError(status.ErrDB)
	}

	return nil
}

// 查询
func (c *WalletMysql) GetUserBalance(ctx context.Context, userId uint32, forUpdate bool) (uint64, protocol.ServerError) {
	query := "SELECT money FROM t_wallet WHERE userId=?"
	if forUpdate {
		query += " FOR UPDATE"
	}

	row := c.qe.QueryRowxContext(ctx, query, userId)
	var accBalance uint64
	err := row.Scan(&accBalance)
	if err != nil {
		log.Println("wallet GetUserBalance failed: ", err)
		return 0, protocol.NewServerError(status.ErrDB)
	}
	return accBalance, nil
}

// mysql goodStudy -uroot -p8918112lu;
// select * from t_wallet;

//DROP TABLE t_wallet;

//创建表↓
//CREATE TABLE IF NOT EXISTS t_wallet(
//	userId INT NOT NULL,
//	money INT NOT NULL,
//	date DATETIME NOT NULL,
//	PRIMARY KEY(userId)
//	)ENGINE=InnoDB DEFAULT CHARSET=utf8;
