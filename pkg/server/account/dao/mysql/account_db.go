package mysql

import (
	_ "database/sql/driver"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)


type accountMysql struct {
	conn *sqlx.DB
}

// 创建conn，连接db
// 构造函数一般放在结构体下方紧挨着, 容易找
func NewAccountMysql() (*accountMysql, error) {
	db, err := sqlx.Connect("mysql", "root:8918112lu@/goodStudy")
	if err != nil {
		log.Println("DB open failed: ", err)
	}
	return &accountMysql{
		conn: db,
	}, err
}

// 写入't_member'table
func (c *accountMysql) InsertInfo(account string, password string) error {
	_, err := c.conn.Exec("INSERT t_member SET account=?, password=?", account, password)
	if err != nil {
		log.Println("account insert failed: ", err)
		return err
	}
	return nil
}

// 查询
func (c *accountMysql) QueryInfo(account string) error {
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


// mysql -u root -p8918112lu;
// use goodStudy;
// select * from t_member;

// DROP TABLE t_member;

// 创建库 CREATE DATABASE goodStudy;
// 创建表↓
//CREATE TABLE IF NOT EXISTS t_member(
//	id INT UNSIGNED AUTO_INCREMENT,
//	account VARCHAR(255) NOT NULL UNIQUE,
//	pass VARCHAR(255) NOT NULL,
//	PRIMARY KEY (id)
//	)ENGINE=InnoDB DEFAULT CHARSET=utf8;
