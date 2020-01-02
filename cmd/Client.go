package main

import (
	"fmt"
	"log"

	"goodgoodstudy.com/pkg/client"
)

func main() {

	// 初始化客户端
	consumeClient, err := client.NewConsumeClient()
	if err != nil {
		log.Fatalln(err)
	}

	resp, err := consumeClient.Pay(nil) // TODO
	if err != nil {
		log.Println("consume failed,", err)
		return
	}
	// 怎么把Inp的返回值传给client包哦
	// 除此之外，这里还要做什么

	log.Println(resp)
}

 func Inp() (ItemId int64, num int64, UserId int32) {
	var(
		item int
		ItemNum int
		user int
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

	return int64(itemId), int64(itemNum), int32(userId)
 }