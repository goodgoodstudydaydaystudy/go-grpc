package mysql

import (
	"context"
	_ "database/sql"
	"log"

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

func freedConn(tx *sqlx.Tx) {
	err := tx.Commit()
	if err != nil {
		log.Println("wallet DB error: ", err)
	}
}

func (c *WalletMysql) Recharge(ctx context.Context, req *pb.RechargeReq) protocol.ServerError {
	// 1. 开启事务
	tx, err := c.conn.BeginTx(ctx, nil)
	if err != nil {
		// TODO
	}
	defer func() {
		if err != nil {
			log.Println("XXX")
			tx.Rollback()
		} else {
			log.Println("XXX")
			tx.Commit()
		}
	}()

	// 2. 加锁, 查余额
	// TODO

	// 3. 判断充值后, 余额是否足够/溢出
	// TODO

	// 4. rollback/commit
	// TODO

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
