package mysql

import (
	rpb "goodgoodstudy.com/go-grpc/pkg/pb/server/account"
	"testing"
)


func TestNewAccountMysql(t *testing.T) {
	db, err := NewAccountMysql()
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("get conn from test:%p", db.conn)
	err = db.AddUsr(&rpb.AddUserReq{
		Account:              "testqqq",
		Password:             "123465",
		Name:                 "qqqq",
		Gender:               0,
	})
	t.Log(err)
}
