package mysql

import (
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

func freedConn(tx *sqlx.Tx) {
	err := tx.Commit()
	if err != nil {
		log.Println("wallet DB error: ", err)
	}
}

// 写入
// 开事务，接defer commit
// getBalance
// FRO UPDATE 上锁
// insert or update
// insert 执行失败后 rollback
func (c *WalletMysql)Recharge(req *pb.RechargeReq) protocol.ServerError {
	tx, err := c.conn.Beginx()
	if err != nil {
		log.Println("Recharge Begin error: ", err)
		return protocol.NewServerError(status.ErrDB)
	}
	defer freedConn(tx)

	getBalance, err :=  c.GetUserBalance(req.GetUserId())
	if err != nil {
		log.Println("Recharge getBalance failed: ", err)
		return protocol.NewServerError(status.ErrDB)
	}

	forUpdate := "SELECT * FROM t_wallet WHERE userId=? FOR UPDATE "
	_, err = c.conn.Exec(forUpdate, req.GetUserId())	// TODO Exec的result有啥用哦？
	if err != nil {
		log.Println("db Recharge forUpdate failed: ", err)
		return protocol.NewServerError(status.ErrDB)
	}

	rechargeMoney := getBalance + req.GetCount()

	now := time.Now()
	nowTime := now.Format("2006-01-02 15:04:05")
	walletInfo := "INSERT INTO t_wallet VALUE(?, ?, ?) ON DUPLICATE KEY UPDATE money=VALUES(money), date=VALUES(date) "

	result, err := c.conn.Exec(walletInfo, req.GetUserId(), rechargeMoney, nowTime)
	if err != nil {
		log.Println("Recharge insert failed: ", err)
		err := tx.Rollback()
		if err != nil {
			log.Println("Recharge rollback failed: ", err)
		}
		return protocol.NewServerError(status.ErrDB)
	}
	log.Println("Exec result: ", result)
	return nil
}

// 查询
func (c *WalletMysql)GetUserBalance(userId uint32) (uint64, protocol.ServerError) {
	row := c.conn.QueryRow("SELECT money FROM t_wallet WHERE userId=?", userId)
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