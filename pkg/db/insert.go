package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"

	"log"
)

// 写入't_member'table
func InsertUserInfo(t string,i int32, a string, p string) error {
	if t != "t_member" {
		log.Println("table name error")
		return nil
	}
	db, err := sql.Open("mysql", "root:8918112lu@/goodStudy")
	if err != nil {
		log.Println("DB.Begin failed: ", err)
		return err
	}
	stmt, err := db.Prepare("INSERT t_member SET userid=?, account=?, md5=?")
	if err != nil {
		log.Println("tx.Prepare failed: ", err)
		return err
	}
	res, err := stmt.Exec(i, a, p)
	if err != nil {
		log.Println("Exec failed: ", err)
		return err
	}
	log.Println("result: ", res)

	return nil
}



// mysql -u root -p8918112lu;
//CREATE DATABASE goodStudy;
//CREATE TABLE IF NOT EXISTS t_member(
//	id INT UNSIGNED AUTO_INCREMENT,
//	userId INT NOT NULL,
//	account VARCHAR(255) NOT NULL,
//	md5 VARCHAR(255) NOT NULL,
//	PRIMARY KEY (id, userId)
//	)ENGINE=InnoDB DEFAULT CHARSET=utf8;