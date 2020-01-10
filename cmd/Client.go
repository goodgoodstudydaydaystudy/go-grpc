package main

import (
	"bufio"
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

	resp, err := consumeClient.Pay(nil) // TODO
	if err != nil {
		log.Println("consume failed,", err)
		return
	}



	log.Println(resp)
}


// 读取终端输入
func Input() {
	var err error
	input := bufio.NewReader(os.Stdin)

	fmt.Printf("Please enter item id:")
	item_id, err := input.ReadString('\n')
	if err != nil {
		log.Println("input item_id", err)
	}
	id, _ := strconv.ParseInt(item_id, 10, 64)
	client.ConsumeClient{}.Pay(&pb.ConsumeReq{ItemId:id})


	fmt.Printf("Please enter Count:")
	itemnum, err := input.ReadString('\n')
	if err != nil {
		log.Println("input itemCount", err)
	}
	num, _ := strconv.ParseInt(itemnum, 10, 64)
	client.ConsumeClient{}.Pay(&pb.ConsumeReq{ItemNum:num})

	fmt.Printf("Please enter userId:")
	userid, err := input.ReadString('\n')
	if err != nil {
		log.Println("input userid", err)
	}
	constomerid, _ := strconv.ParseInt(userid, 10, 32)
	client.ConsumeClient{}.Pay(&pb.ConsumeReq{UserId:int32(constomerid)})

	fmt.Printf("remark or not:")
	descr, err := input.ReadString('\n')
	if descr == "" {
		descr = "nil"
	}
	if err != nil {
		log.Println("description", err)
	}
	client.ConsumeClient{}.Pay(&pb.ConsumeReq{Description:descr})
}