package main

import (
	"fmt"
	"log"
	//"goodgoodstudy.com/go-grpc/pkg/client"
)

func main() {

	// 初始化客户端
	//consumeClient, err := client.NewConsumeClient()
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//
	//resp, err := consumeClient.Pay(nil) // TODO
	//if err != nil {
	//	log.Println("consume failed,", err)
	//	return
	//}
	//// 怎么把Inp的返回值传给client包哦
	//// 除此之外，这里还要做什么

	//log.Println(resp)

	fmt.Printf("InputData: %v", Inp())		// InputData: map[itemId:1 itemNum:1 userId:1]
}

 func Inp() map[string]int {
	var(
		item int64
		ItemNum int32
		user int64
	)
	fmt.Printf("Please intput item_id:")
	itemId, err := fmt.Scanln(&item)
	if err != nil {
		log.Println(err)
	}

	fmt.Printf("Please intput num:")
	itemNum, err := fmt.Scanln(&ItemNum)
	if err != nil {
		log.Println(err)
	}

	fmt.Printf("Please intput user_id:")
	userId, err := fmt.Scanln(&user)
	if err != nil {
		log.Println(err)
	}
	var InputData = map[string]int{
		"itemId":  itemId,
		"itemNum": itemNum,
		"userId":  userId,
	}
	 return InputData
 }