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

// Rollback
func freedConn(tx *sqlx.Tx) {
	err := tx.Commit()
	if err != nil {
		log.Println("wallet DB error: ", err)
	}
}

// 写入	不知道有没有上锁…
func (c *WalletMysql)Recharge(req *pb.RechargeReq) protocol.ServerError {
	tx, err := c.conn.Beginx()
	if err != nil {
		log.Println("Recharge Begin error: ", err)
		return protocol.NewServerError(status.ErrDB)
	}
	defer freedConn(tx)

	now := time.Now()
	nowTime := now.Format("2006-01-02 15:04:05")
	walletInfo := "INSERT INTO t_wallet VALUE(?, ?, ?) ON DUPLICATE KEY UPDATE money=VALUES(money), date=VALUES(date) "

	//_, err = tx.Exec(walletInfo, req.GetUserId(), req.GetCount(), nowTime)
	result, err := c.conn.Exec(walletInfo, req.GetUserId(), req.GetCount(), nowTime)
	if err != nil {
		log.Println("Recharge insert failed: ", err)
		return protocol.NewServerError(status.ErrDB)
	}
	log.Println("Exec result: ", result)
	return nil
}

// 查询
func (c *WalletMysql)GetUserBalance(userId uint32) (uint64, protocol.ServerError) {
	tx, err := c.conn.Beginx()
	if err != nil {
		log.Println("GetUserBalance Begin error: ", err)
		return 0, protocol.NewServerError(status.ErrDB)
	}
	freedConn(tx) // 这个好像有问题

	row := c.conn.QueryRow("SELECT money FROM t_wallet WHERE userId=? FOR UPDATE ", userId)
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