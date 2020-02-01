package mysql

import (
	"goodgoodstudy.com/go-grpc/protocol/common/status"
	"log"
	"time"

	pb "goodgoodstudy.com/go-grpc/pkg/pb/wallet"
	protocol "goodgoodstudy.com/go-grpc/pkg/procotol"

	"github.com/jmoiron/sqlx"
)

type WalletMysql struct {
	db *sqlx.DB
}

func NewWalletMysql() (*WalletMysql, error) {
	db, err := sqlx.Connect("mysql", "root:8918112lu@/goodStudy")
	if err != nil {
		log.Println("NewWallet conn failed: ", err)
		return nil, err
	}
	return &WalletMysql{
		db: db,
	}, nil
}

//
func freedConn(tx *sqlx.Tx) {
	err := tx.Rollback()
	if err != nil {
		log.Println("wallet DB error: ", err)
	}
}

// 写入
func (c *WalletMysql)Recharge(req *pb.RechargeReq) protocol.ServerError {
	now := time.Now()
	nowTime := now.Format("2006-01-02 15:04:05")
	walletInfo := "INSERT INTO t_wallet(userId, money, date) VALUE(?, ?, ?)"

	tx, err := c.db.Beginx()
	if err != nil {
		log.Println("Recharge Begin error: ", err)
		return protocol.NewServerError(status.ErrDB)
	}
	defer freedConn(tx)

	_, err = tx.Exec(walletInfo, req.GetUserId(), req.GetCount(), nowTime)
	if err != nil {
		log.Println("Recharge insert failed: ", err)
		return protocol.NewServerError(status.ErrDB)
	}

	return nil
}

// 查询
func (c *WalletMysql)GetUserBalance(userId uint32) (uint64, protocol.ServerError) {
	tx, err := c.db.Beginx()
	if err != nil {
		log.Println("GetUserBalance Begin error: ", err)
		return 0, protocol.NewServerError(status.ErrDB)
	}
	freedConn(tx)

	row := c.db.QueryRow("SELECT money FROM t_wallet WHERE userId=?", userId)

	var accBalance uint64
	err = row.Scan(&accBalance)
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