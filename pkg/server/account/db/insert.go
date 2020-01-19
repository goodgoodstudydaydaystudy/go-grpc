package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	ec "goodgoodstudy.com/go-grpc/pkg/server/account/error_tips"

	"log"
)

// 写入't_member'table
func InsertUserInfo(tableName string,userId int32, account string, password string) (string, error) {
	if tableName != "t_member" {
		log.Println("t able name error")
		return "", nil
	}
	db, err := sql.Open("mysql", "root:8918112lu@/goodStudy")
	if err != nil {
		log.Println("DB.Begin failed: ", err)
		return "", err
	}
	stmt, err := db.Prepare("INSERT t_member SET userid=?, account=?, md5=?")

	if err != nil {
		//errorMessage := ec.ErrorInfo{}.DbError(err).Message
		errorMessage := ec.AccountConflict.Message
		log.Println("tx.Prepare failed: ", err)
		return errorMessage, err
	}
	_, err = stmt.Exec(userId, account, password)
	if err != nil {
		log.Println("Exec failed: ", err)
		return "", err
	}
	return "success", nil
}



// mysql -u root -p8918112lu;
// use goodStudy;
// select * from t_member;

// 创建裤 CREATE DATABASE goodStudy;
// 创建表↓
//CREATE TABLE IF NOT EXISTS t_member(
//	id INT UNSIGNED AUTO_INCREMENT,
//	userId INT NOT NULL,
//	account VARCHAR(255) NOT NULL UNIQUE,
//	md5 VARCHAR(255) NOT NULL,
//	PRIMARY KEY (id)
//	)ENGINE=InnoDB DEFAULT CHARSET=utf8;