package db

import (
	"database/sql"
	_ "database/sql/driver"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type connDb struct {
	conn *sql.DB
}


// 写入't_member'table
func (c *connDb) InsertInfo(userId int32, account string, password string) error {
	stmt, err := c.conn.Prepare("INSERT t_member SET userid=?, account=?, md5=?")
	if err != nil {
		log.Println("tx.Prepare failed: ", err)
		return err
	}
	_, err = stmt.Exec(userId, account, password)
	if err != nil {
		log.Println("Exec failed: ", err)
		return err
	}
	return nil
}


// 查询
func (c *connDb) QueryInfo(account string) error {
	stmt, err := c.conn.Prepare("SELECT * FROM t_table WHERE account=?;")
	if err != nil {
		log.Println("query prepare failed:", err)
	}

	rows, err := stmt.Query(account)
	if err != nil {
		log.Println("query account failed: ", err)
	}

	for rows.Next() {
		var account string
		err := rows.Scan(&account)
		if err != nil {
			log.Println("query account result error: ", err)
			return err
		}
	}
	return nil
}


func NewConnDb() (*connDb, error) {
	db ,err := sql.Open("mysql", "root:8918112lu@/goodStudy")
	if err != nil {
		log.Println("DB open failed: ", err)
	}
	return &connDb{
		conn: db,
	}, nil
}

func Insert(userId int32, account string, password string) error {
	err := Insert(userId, account, password)
	return err
}


// mysql -u root -p8918112lu;
// use goodStudy;
// select * from t_member;

// 创建库 CREATE DATABASE goodStudy;
// 创建表↓
//CREATE TABLE IF NOT EXISTS t_member(
//	id INT UNSIGNED AUTO_INCREMENT,
//	userId INT NOT NULL,
//	account VARCHAR(255) NOT NULL UNIQUE,
//	md5 VARCHAR(255) NOT NULL,
//	PRIMARY KEY (id)
//	)ENGINE=InnoDB DEFAULT CHARSET=utf8;