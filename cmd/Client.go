package main

import (
	"bufio"
	"context"
	"fmt"
	"goodgoodstudy.com/go-grpc/pkg/client"
	pb "goodgoodstudy.com/go-grpc/pkg/pb"
	"log"
	"os"
	"strconv"
)

func main() {

	// 初始化客户端
	consumeClient, err := client.NewConsumeClient()
	if err != nil {
		log.Println("new client:", err)
	}

	ctx := context.Background()

	resp, err := consumeClient.Pay(ctx, nil) // TODO
	if err != nil {
		log.Println("consume failed,", err)
		return
	}

	log.Println(resp)

	_ = consumeClient.Close()
}

// 读取终端输入
// 这个函数没有意义
func Input() {
	ctx := context.Background()
	consumeClient, err := client.NewConsumeClient()
	if err != nil {
		log.Println("NewConsumeClient", err)
	}

	input := bufio.NewReader(os.Stdin)

	fmt.Printf("Please enter item id:")
	item_id, err := input.ReadString('\n')
	if err != nil {
		log.Println("input item_id", err)
	}
	id, _ := strconv.ParseInt(item_id, 10, 64)
	consumeClient.Pay(ctx, &pb.ConsumeReq{ItemId: id})

	fmt.Printf("Please enter Count:")
	itemnum, err := input.ReadString('\n')
	if err != nil {
		log.Println("input itemCount", err)
	}
	num, _ := strconv.ParseInt(itemnum, 10, 64)
	consumeClient.Pay(ctx, &pb.ConsumeReq{ItemNum: num})

	fmt.Printf("Please enter userId:")
	userid, err := input.ReadString('\n')
	if err != nil {
		log.Println("input userid", err)
	}
	constomerid, _ := strconv.ParseInt(userid, 10, 32)
	consumeClient.Pay(ctx, &pb.ConsumeReq{UserId: int32(constomerid)})

	fmt.Printf("remark or not:")
	descr, err := input.ReadString('\n')
	if descr == "" {
		descr = "nil"
	}
	if err != nil {
		log.Println("description", err)
	}
	consumeClient.Pay(ctx, &pb.ConsumeReq{Description: descr})
}
