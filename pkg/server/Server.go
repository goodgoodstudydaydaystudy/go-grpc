package server

import (
	"context"
	"database/sql"
	"fmt"
	pb "goodgoodstudy.com/go-grpc/pkg/pb"
	"log"
)

type ControlServer struct {
}

func (s *ControlServer) Pay(ctx context.Context, consumeReq *pb.ConsumeReq) (*pb.ConsumeResp, error) {
	// send to 'WriteSql' TODO
	return &pb.ConsumeResp{OrderId: consumeReq.GetItemId()}, nil // 返回Resp里的字段？
}

func WriteMySql() {
	// connect mysql
	db, err := sql.Open("mysql", "root:284927463@/order_sql")
	if err != nil {
		log.Printf("failel to connect mysql %v\n", err)
	}
	Insert(db, 0, 0) // TODO
}

// insert 'item_info' tables
func Insert(db *sql.DB, item_id int64, num int32) {
	stmt, err := db.Prepare("insert ITEMINFO(ITEM_ID, NUM) values (?, ?)")
	if err != nil {
		log.Printf("insert failed:", err, "\n")
	}
	// Exec() 需要传入数据 TODO
	res, err := stmt.Exec()
	classID, err := res.LastInsertId()
	if err != nil {
		log.Printf("lasetInsertId failed:", err, "\n")
	}
	fmt.Println(classID)
}

//mysql -uroot -p284927463 order_sql
//SELECT * FROM ITEM_INFO;
//go build Client.go
//INSERT INTO ITEM_INFO (ITEM_ID, PRICE, NUM) VALUES (2, 18, 5);
//CREATE DATABASE order_sql;
//CREATE TABLE ORDER_INFO(id int NOT NULL auto_increment,
//order_id varchar(255) NOT NULL,
//alive int NOT NULL DEFAULT 0,
//item_id int NOT NULL,
//item_amount int NOT NULL,
//item_money int NOT NULL,
//order_time datetime NOT NULL,
//PRIMARY KEY(id));
//
//CREATE TABLE USER_PAY(id int NOT NULL auto_increment,
//order_id int NOT NULL,
//pay_id int NOT NULL,
//pay_time datetime NOT NULL,
//PRIMARY KEY(id));
//
//CREATE TABLE ITEM_INFO(ITEM_ID int NOT NULL auto_increment,
//PRICE int NOT NULL,
//NUM int NOT NULL,
//PRIMARY KEY(ITEM_ID));
//
//CREATE TABLE MEM_INFO(id int NOT NULL auto_increment,
//name varchar(255) NOT NULL,
//keyword varchar(255) NOT NULL,
//call_num varchar(255) NOT NULL,
//gender varchar(255) NOT NULL,
//birthday date,
//accou_bal int NOT NULL DEFAULT 0,
//times_consu int NOT NULL DEFAULT 0,
//money_consu int NOT NULL DEFAULT 0,
//PRIMARY KEY(id));
//
//CREATE TABLE customer_order(id int not null auto_increment,
//name varchar(255) not null,
//done varchar(255) not null,
//money int not null,
//time datetime not null,
//PRIMARY KEY(id));
