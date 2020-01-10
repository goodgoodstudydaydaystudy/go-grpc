package main

import (
	"bufio"
	"fmt"
	"goodgoodstudy.com/go-grpc/pkg/client"
	"log"
	"os"
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
	client.ConsumeClient{}.Pay()

	fmt.Printf("Please enter Count:")
	itemnum, err := input.ReadString('\n')
	if err != nil {
		log.Println("input itemCount", err)
	}

	fmt.Printf("Please enter userId:")
	userid, err := input.ReadString('\n')
	if err != nil {
		log.Println("input userid", err)
	}

	fmt.Printf("remark or not:")
	descr, err := input.ReadString('\n')
	if descr == "" {
		descr = "nil"
	}
	if err != nil {
		log.Println("description", err)
	}
}