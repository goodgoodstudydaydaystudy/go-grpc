package main

import (
	"goodgoodstudy.com/go-grpc/pkg/client"
	"log"
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

	log.Println(resp)
}
