package mysql

import (
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
	}
	return &accountMysql{
		conn: db,
	}, err
}

// 写入't_member'table
func (c *accountMysql) InsertInfo(req *rpb.RegisterReq) error {
	accountInfo := "INSERT INTO t_member(account, password, name, gender) VALUE (?, ?, ?, ?)"
	_, err := c.conn.Exec(accountInfo, req.GetAccount(), req.GetPassword(), req.Name, req.GetGender())
	if err != nil {
		log.Println("account insert failed: ", err)
		return err
	}
	return nil
}

// 查询
func (c *accountMysql) QueryInfo(account string) (int32, string, string, error) {
	rows := c.conn.QueryRow("SELECT * FROM t_member WHERE account=?", account)
	var (
		id       int
		password string
		name     string
	)
	err := rows.Scan(&id, &name, &password, &account)
	if err != nil {
		log.Println(err)
		return 0, "", "", nil
	}
	outputId := int32(id)
	return outputId, password, name, nil
}

// 通常来说, 查用户都是用uid查询, 用account查询也可以, 就必须在数据库中给account这个字段添加索引, 不然查找很慢
func (c *accountMysql) GetUserById(userId uint32) (*account.UserInfo, error) {
	userInfo := &account.UserInfo{}
	err := c.conn.Get(userInfo, "SELECT * from t_member WHERE id=?", userId) // 这里用了反射, 看UserInfo结构体后面的tag; 通常select *都用这个来查询; 除非遇到只查一两个字段的, 就用Scan
	if err != nil {
		return nil, protocol.NewServerError(status.ErrDB) // TODO 这里要换成真正的错误, 比如判断是不是用户不存在, 还是数据库连接不上
	}

	return userInfo, err
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
