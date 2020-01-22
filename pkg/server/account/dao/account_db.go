package dao

import (
	_ "database/sql/driver"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type ConnDb struct {
	conn *sqlx.DB
}

// 创建conn，连接db
// 构造函数一般放在结构体下方紧挨着, 容易找
func NewConnDb() (*ConnDb, error) {
	db, err := sqlx.Open("mysql", "root:8918112lu@/goodStudy")
	if err != nil {
		log.Println("DB open failed: ", err)
	}
	return &ConnDb{
		conn: db,
	}, err
}

// 写入't_member'table
func (c *ConnDb) InsertInfo(userId int32, account string, password string) error {
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
func (c *ConnDb) QueryInfo(account string) error {
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

// TODO
// 创建conn，保存conn
// conn是给server调用
// server需要访问 InsertInfo 传入数据
// 提供访问db内部的接口
// InsertInfo 写入数据库

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
