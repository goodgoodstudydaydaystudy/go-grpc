package main

import (
	"bufio"
	"context"
	"fmt"
	"goodgoodstudy.com/go-grpc/pkg/client"
	pb "goodgoodstudy.com/go-grpc/pkg/pb"
	rpb "goodgoodstudy.com/go-grpc/pkg/pb/Register"
	"log"
	"os"
	"strconv"
	"strings"
)

type TestArgs struct {
	id int64
	count int64
	user int64
	descri string

}

func main() {

	// 初始化客户端
	consumeClient, err := client.NewConsumeClient()
	if err != nil {
		log.Println("new client:", err)
	}

	accountClient, err := client.NewAccountClient()
	if err != nil {
		log.Println("new client:", err)
	}

	ctx := context.Background()

	//Input()	//输入信息

	// 付款接口
	consuResp, err := consumeClient.Pay(ctx, &pb.ConsumeReq{ItemId: TestArgs{}.id, ItemNum: 2, UserId:3, Description:"aaaa"}) // TODO
	if err != nil {
		log.Println("consume failed,", err)
		return
	}
	log.Println(consuResp)

	// 注册接口
	regisResp, err := accountClient.Regis(ctx, &rpb.RegisReq{UeserId: 007, Account:"777", Password:"666"})
	if err != nil {
		log.Println("regisClient.Regis failed: ", err)
	}
	log.Println(regisResp)

	// 登录接口
	loginResp, err := accountClient.Login(ctx, &rpb.LogReq{Account: "0903", Password: "888"})
	if err != nil {
		log.Println("accountClient.Login failed: ", err)
	}
	log.Println(loginResp)

	_ = consumeClient.Close()	// 不是执行完main就close了？ 为啥不用defer？
	_ = accountClient.RegClose()	// 执行完close
}

// 读取终端输入
// 这个函数没有意义
func Input() (*TestArgs, error){
	input := bufio.NewReader(os.Stdin)

	fmt.Printf("Please enter item id:")
	itemId, err := input.ReadString('\n')
	if err != nil {
		log.Println("input item_id", err)
	}
	itemid, _:= strconv.ParseInt(strings.TrimSpace(itemId), 10, 64)
	//consumeClient.Pay(ctx, &pb.ConsumeReq{ItemId: id})
	//log.Println("itemid :", itemid)							// input
	//log.Println("itemid type:", reflect.TypeOf(itemid))		// int64


	fmt.Printf("Please enter Count:")
	itemnum, err := input.ReadString('\n')
	if err != nil {
		log.Println("input itemCount", err)
	}
	itemcount, _ := strconv.ParseInt(itemnum, 10, 64)
	//consumeClient.Pay(ctx, &pb.ConsumeReq{ItemNum: num})

	fmt.Printf("Please enter userId:")
	userid, err := input.ReadString('\n')
	if err != nil {
		log.Println("input userid", err)
	}
	userId, _ := strconv.ParseInt(userid, 10, 32)
	//consumeClient.Pay(ctx, &pb.ConsumeReq{UserId: int32(constomerid)})

	fmt.Printf("remark or not:")
	descri, err := input.ReadString('\n')
	if descri == "" {
		descri = "nil"
	}
	if err != nil {
		log.Println("description", err)
	}
	//consumeClient.Pay(ctx, &pb.ConsumeReq{Description: descr})
	return &TestArgs{		// 这个是不是赋值操作咧
		id:     itemid,
		count:  itemcount,
		user:   userId,
		descri: descri,
	}, nil
}