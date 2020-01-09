package client

import (
	"bufio"
	"fmt"
	pb "goodgoodstudy.com/go-grpc/pkg/pb"
	"google.golang.org/grpc"
	"log"
	"os"
	"strconv"
)

const (
	port = ":50051"
)

type ConsumeClient struct {
}

func NewConsumeClient() (*ConsumeClient, error) {
	conn, err := grpc.Dial(port, grpc.WithInsecure())
	if err != nil {
		log.Printf("failed to listen: %v", err)
	}
	defer conn.Close()
	client := pb.NewControlClient(conn)
	client.Pay()	// TODO 这里的pay()的参数呀

	return &ConsumeClient{}, nil // TODO
}

// 获取input数据
func (c *ConsumeClient) Pay(req *pb.ConsumeReq) (resp *pb.ConsumeResp, err error) {
	// TODO
	InputData := Input()
	Itemid, _ := strconv.ParseInt(InputData["item_id"], 10 , 64)
	ItemCount, _ := strconv.ParseInt(InputData["itemCount"], 10, 32)
	UeserId, _ := strconv.ParseInt(InputData["user_id"], 10, 64)

	return &pb.ConsumeResp{OrderId: resp.GetOrderId(), Message: resp.GetMessage()}, nil
}

// 读取终端输入
func Input() map[string]string {
	var err error
	input := bufio.NewReader(os.Stdin)

	fmt.Printf("Please enter item id:")
	item_id, err := input.ReadString('\n')
	if err != nil {
		log.Println("input item_id", err)
	}

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

	var InputMap = map[string]string{
		"item_id": item_id,
		"itemCount": itemnum,
		"user_id": userid,
		"descri": descr
	}
	return InputMap
}