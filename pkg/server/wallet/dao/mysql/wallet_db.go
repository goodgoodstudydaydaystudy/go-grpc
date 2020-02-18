package mysql

import (
	"context"
	"database/sql"
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

func NewWalletMysql(qe queryExec) *WalletMysql {
	return &WalletMysql{
		qe: qe,
	}
}

// 只负责写入表，连"锁"都没有
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
// 只管查询，加入 forUpdate bool 提供"查询锁行"的功能
func (c *WalletMysql) GetUserBalance(ctx context.Context, userId uint32, forUpdate bool) (int64, protocol.ServerError) {
	// 这个操作太骚了
	query := "SELECT money FROM t_wallet WHERE userId=?"
	if forUpdate {
		query += " FOR UPDATE"
	}

	row := c.qe.QueryRowxContext(ctx, query, userId)
	var accBalance int64
	err := row.Scan(&accBalance)
	if err == sql.ErrNoRows {
		return 0, nil
	}
	if err != nil {
		log.Println("wallet GetUserBalance failed: ", err)
		return 0, protocol.NewServerError(status.ErrDB)
	}
	return accBalance, nil
}

func (c *WalletMysql) RecordOrderNoPaid(ctx context.Context, userId uint32, orderId string, orderTime string, deadline string) (err protocol.ServerError) {
	var de error
	orderExec := "INSERT INTO t_order(orderId, userId, orderTime, deadline) VALUE(?, ?, ?, ?)"
	_, de = c.qe.ExecContext(ctx, orderExec, orderId, userId, orderTime, deadline)
	if de != nil {
		log.Println("insert order failed:", de)
		return protocol.NewServerError(status.ErrDB)
	}
	return nil
}

func (c *WalletMysql) Pay(ctx context.Context, orderId string) protocol.ServerError {
	now := time.Now()
	payTime := now.Format("2006-01-02-15 04:05:00")

	payExec := "INSERT INTO t_pay(orderId, payTime) VALUE(?, ?)"
	_, err := c.qe.ExecContext(ctx, payExec, orderId, payTime)
	if err != nil {
		log.Println("insert t_pay failed:", err)
		return protocol.NewServerError(status.ErrDB)
	}

	orderExec := "UPDATE t_order SET status=? WHERE orderId=?"
	_, err = c.qe.ExecContext(ctx, orderExec, 1, orderId)
	if err != nil {
		log.Println("update t_order failed:", err)
		return protocol.NewServerError(status.ErrDB)
	}
	return nil
}

// scan t_order
func (c *WalletMysql) GetNoPaid(ctx context.Context) (NoPaid []string, serverError protocol.ServerError) {
	ScanExec := "SELECT orderId FROM t_order WHERE status=?"
	rows, err := c.qe.QueryContext(ctx, ScanExec, 0)
	if err != nil {
		log.Println("GetNoPaid query failed:", err)
		return nil, protocol.NewServerError(status.ErrDB)
	}
	for rows.Next() {
		var noPaidOrder string
		err := rows.Scan(&noPaidOrder)
		if err != nil {
			log.Println("GetNoPaid failed ")
			return nil, protocol.NewServerError(status.ErrDB)
		}
		NoPaid = append(NoPaid, noPaidOrder)
	}
	return NoPaid, nil
}

// 接收过期订单，修改t_order的status=3(expired)
func (c *WalletMysql) MarkExpiredOrder(ctx context.Context, expiredOder []string) (err protocol.ServerError) {
	// range expiredOder []string to get order
	for _, orderId := range expiredOder{
		// TODO 尝试加上 go 协程
		checkOrderExec := "SELECT status FROM t_order WHERE orderId=?"
		row := c.qe.QueryRowxContext(ctx, checkOrderExec, orderId)
		var statusNum int
		err := row.Scan(&statusNum)
		if err != nil {
			log.Println("MarkExpiredOrder get order failed:", err)
		}
		// check status
		// mark status=3
		if statusNum == 0 {
			orderExec := "UPDATE t_order SET status=? WHERE orderId=?"
			_, err := c.qe.ExecContext(ctx, orderExec, 3, orderId)
			if err != nil {
				log.Println("MarkExpiredOrder exec failed:", err)
				return nil
			}
		}
		}
	return nil
}


// mysql goodStudy -uroot -p8918112lu;
// select * from t_wallet;

//DROP TABLE t_wallet;

//创建表↓
// t_wallet
/*
CREATE TABLE IF NOT EXISTS t_wallet(
	userId INT NOT NULL,
	money INT NOT NULL,
	date DATETIME NOT NULL,
	PRIMARY KEY(userId)
	)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
*/

// t_order
/*
CREATE TABLE IF NOT EXISTS t_order(
	id INT UNSIGNED AUTO_INCREMENT,
	orderId CHAR(255) NOT NULL,
	userId INT NOT NULL,
	orderTime CHAR(255) NOT NULL,
	deadline CHAR(255) NOT NULL,
	status INT NOT NULL DEFAULT 0,
	PRIMARY KEY(id)
	)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
 */

// t_pay
/*
CREATE TABLE IF NOT EXISTS t_pay(
	id INT UNSIGNED AUTO_INCREMENT,
	orderId CHAR(255) NOT NULL,
	payTime CHAR(255) NOT NULL,
	PRIMARY KEY(id)
	)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
 */