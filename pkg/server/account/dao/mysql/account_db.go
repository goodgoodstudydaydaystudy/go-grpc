package mysql

import (
	"database/sql"
	_ "database/sql/driver"
	"log"

	rpb "goodgoodstudy.com/go-grpc/pkg/pb/account"
	protocol "goodgoodstudy.com/go-grpc/pkg/procotol"
	account "goodgoodstudy.com/go-grpc/pkg/server/account/dao/entity"
	"goodgoodstudy.com/go-grpc/protocol/common/status"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
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
		return nil, err
	}
	return &accountMysql{
		conn: db,
	}, err
}

// 写入't_member'table
// 改名叫Register, 函数名要有辨识性, InsertInfo, 什么Info?
func (c *accountMysql) Register(req *rpb.RegisterReq) protocol.ServerError {
	accountInfo := "INSERT INTO t_member(account, password, name, gender) VALUE (?, ?, ?, ?)"
	_, err := c.conn.Exec(accountInfo, req.GetAccount(), req.GetPassword(), req.Name, req.GetGender())
	if err != nil {
		log.Println("account insert failed: ", err)
		return protocol.NewServerError(status.ErrDB)
	}
	return nil
}

func (c *accountMysql) GetUserPasswordByAccount(acc string) (string, protocol.ServerError) {
	row := c.conn.QueryRow("SELECT password from t_member where account=?", acc)

	var pwd string
	err := row.Scan(&pwd)	// 通过scan获得row里面的数据
	if err == sql.ErrNoRows {
		return "", protocol.NewServerError(status.ErrAccountNotExists)
	}

	if err != nil {
		return "", protocol.NewServerError(status.ErrDB)
	}

	return pwd, nil
}

// 查询
// 改名叫GetUserByAccount
func (c *accountMysql) GetUserByAccount(acc string) (*account.UserInfo, protocol.ServerError) {
	log.Println("GetUserByAccount acc: ", acc)
	userInfo := &account.UserInfo{}
	err := c.conn.Get(userInfo, "SELECT * FROM t_member WHERE account=?", acc)
	if err == sql.ErrNoRows {
		return nil, protocol.NewServerError(status.ErrAccountNotExists)
	}

	if err != nil {
		log.Println("GetUserByAccount err: ", err)
		return nil, protocol.NewServerError(status.ErrDB)
	}
	return userInfo, nil
}

// 通常来说, 查用户都是用uid查询, 用account查询也可以, 就必须在数据库中给account这个字段添加索引, 不然查找很慢
func (c *accountMysql) GetUserById(userId uint32) (*account.UserInfo, protocol.ServerError) {
	userInfo := &account.UserInfo{}
	err := c.conn.Get(userInfo, "SELECT * from t_member WHERE id=?", userId) // 这里用了反射, 看UserInfo结构体后面的tag; 通常select *都用这个来查询; 除非遇到只查一两个字段的, 就用Scan
	if err == sql.ErrNoRows {	// 通过sql的查询err判定。
		return nil, protocol.NewServerError(status.ErrAccountNotExists)
	}

	if err != nil {
		return nil, protocol.NewServerError(status.ErrDB) // 这里要换成真正的错误, 比如判断是不是用户不存在, 还是数据库连接不上
	}

	return userInfo, nil
}

// mysql goodStudy -uroot -p8918112lu;
// select * from t_member;

// DROP TABLE t_member;

// 创建库 CREATE DATABASE goodStudy;
// 创建表↓
//CREATE TABLE IF NOT EXISTS t_member(
//	id INT UNSIGNED AUTO_INCREMENT,
//	account VARCHAR(255) NOT NULL UNIQUE,
//	password VARCHAR(255) NOT NULL,
//  name CHAR(255) NOT NULL,
//  gender CHAR(10) NOT NULL,
//	PRIMARY KEY (id)
//	)ENGINE=InnoDB DEFAULT CHARSET=utf8;
