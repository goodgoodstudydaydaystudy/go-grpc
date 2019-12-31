package main

import (
	"log"

	"goodgoodstudy.com/go-grpc/pkg/client"
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
