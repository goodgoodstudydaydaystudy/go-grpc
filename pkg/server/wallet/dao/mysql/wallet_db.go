package mysql

import (
	"context"
	_ "database/sql"
	"log"
	"time"

	pb "goodgoodstudy.com/go-grpc/pkg/pb/wallet"
	protocol "goodgoodstudy.com/go-grpc/pkg/procotol"
	"goodgoodstudy.com/go-grpc/protocol/common/status"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type WalletMysql struct {
	conn *sqlx.DB
}

func NewWalletMysql() (*WalletMysql, error) {
	db, err := sqlx.Connect("mysql", "root:8918112lu@/goodStudy")
	if err != nil {
		log.Println("NewWallet conn failed: ", err)
		return nil, err
	}
	return &WalletMysql{
		conn: db,
	}, nil
}


func (c *WalletMysql) Recharge(ctx context.Context, req *pb.RechargeReq) protocol.ServerError {
	// 1. 开启事务
	tx, err := c.conn.BeginTx(ctx, nil)	// TODO Begin和BeginTx有什么区别哦？
	if err != nil {
		log.Println("Recharge BeginTx failed: ", err)
		return protocol.NewServerError(status.ErrDB)
	}

	// 要理解这个匿名函数的写法
	//defer func() {
	//	if err != nil {
	//		log.Println("db Recharge failed -> rollback: ", err)
	//		tx.Rollback()
	//	} else {
	//		log.Println("Recharge success")
	//		tx.Commit()
	//	}
	//}()

	// 2. 加锁, 查余额
	forUpdate := "SELECT * FROM t_wallet WHERE userId=? FOR UPDATE"
	_, err = tx.Exec(forUpdate, req.GetUserId())
	if err != nil {
		log.Println("db Recharge forUpdate failed: ", err)
		return protocol.NewServerError(status.ErrDB)
	}

	getBalance, err := c.GetUserBalance(ctx, req.GetUserId())
	if err != nil {
		log.Println("db getBalance failed: ", err)
		return protocol.NewServerError(status.ErrDB)
	}

	// 3. 判断充值后, 余额是否足够/溢出
	rechargeMoney := (getBalance + req.Count) * 100
	if rechargeMoney < 0 {
		return protocol.NewServerError(status.ErrBalanceNotEnough)
	}

	now := time.Now()
	rechargeTime := now.Format("2006-01-02 15:04:05")

	recharge := "INSERT INTO t_wallet VALUE(?, ?, ?) ON DUPLICATE KEY UPDATE money=VALUES(money), date=VALUES(date)"
	_, err = tx.Exec(recharge, req.GetUserId(), rechargeMoney, rechargeTime)
	if err != nil {
		log.Println("db Recharge failed -> rollback: ", err)
		err := tx.Rollback()
		if err != nil {
			log.Println("Recharge Rollback failed: ", err)
		}
		return protocol.NewServerError(status.ErrDB)
	}

	// 4. rollback/commit
	err = tx.Commit()
	if err != nil {
		log.Println("Recharge Commit failed: ", err)
	}

	// 5. return
	return nil
}

// 查询
func (c *WalletMysql) GetUserBalance(ctx context.Context, userId uint32) (uint64, protocol.ServerError) {
	row := c.conn.QueryRowContext(ctx, "SELECT money FROM t_wallet WHERE userId=?", userId)
	var accBalance uint64
	err := row.Scan(&accBalance)
	if err != nil {
		log.Println("db GetUserBalance failed: ", err)
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
