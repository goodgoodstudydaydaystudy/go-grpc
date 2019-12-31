package main

import (
	"context"
	"fmt"
	pb "gRPC/pb"
	"google.golang.org/grpc"
	"log"
)

const (
	port = ":50051"
)

func main()  {
	conn, err := grpc.Dial(port, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	log.Printf("Connect succee\n")
	defer conn.Close()

	// 初始化客户端
	client := pb.NewControlClient(conn)

	itemid, num, userid := Inp()

	// 希望这个处理过程从main函数中拆出来。TODO
	ConsumeResp, err := client.Pay(context.Background(), &pb.ConsumeReq{
		ItemId: itemid, ItemNum:num, UserId:userid,
	})
	if err != nil {
		log.Fatalf("Consume not succee %v", err)
	}
	log.Printf("Pay succee %v", ConsumeResp)


}

//// send Server TODO
//func ConsumeInfo(itemid int64, num int64, userid int32) {
//	_ = &pb.ConsumeReq{ItemId:itemid, ItemNum:num, UserId:userid}
//	return
//}

// Pay input
func Inp() (itemid int64, num int64, userid int32) {
	var (
		item int
		itemnum int
		user int
	)
	fmt.Printf("Please input item id")
	itemId, err := fmt.Scanln(&item)
	if err != nil{
		log.Printf("item id is err:", err)
	}

	fmt.Printf("Please input num")
	itemNum, err := fmt.Scanln(&itemnum)
	if err != nil{
		log.Printf("num is err:", err)
	}

	fmt.Printf("Please input user id")
	userId, err := fmt.Scanln(&user)
	if err != nil{
		log.Printf("user_id is err:", err)
	}
	return int64(itemId), int64(itemNum), int32(userId)
}
