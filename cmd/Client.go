package main

import (
	"context"
	"goodgoodstudy.com/go-grpc/client"
	rpb "goodgoodstudy.com/go-grpc/pkg/pb/Account"
	"goodgoodstudy.com/go-grpc/pkg/pb/Pay"
	"log"
)

//type TestArgs struct {
//	id int64
//	count int64
//	user int64
//	descri string
//}

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

	// 付款接口
	Consume(ctx, consumeClient)

	// 注册接口
	Register(ctx, accountClient)

	// 登录接口
	Login(ctx, accountClient)

	_ = consumeClient.Close()	// 不是执行完main就close了？ 为啥不用defer？
	_ = accountClient.Close()	// 执行完close
}

// 付款接口
func Consume(ctx context.Context, consumeClient *client.ConsumeClient)  {
	consumeResp, err := consumeClient.Pay(ctx, &Pay.ConsumeReq{ItemId: 1, ItemNum: 2, UserId:3, Description:"aaaa"})
	if err != nil {
		log.Println("consume failed,", err)
		return
	}
	log.Println(consumeResp)
}

// 注册接口
func Register(ctx context.Context, accountClient *client.Client)  {
	registerResp, err := accountClient.Register(ctx, &rpb.RegisReq{Account:"777", Password:"666"})
	if err != nil {
		log.Println("regisClient.Regis failed: ", err)
	}
	log.Println(registerResp)
}

// 登录接口
func Login(ctx context.Context, accountClient *client.Client)  {
	loginResp, err := accountClient.Login(ctx, &rpb.LoginReq{Account: "0903", Password: "888"})
	if err != nil {
		log.Println("accountClient.Login failed: ", err)
	}
	log.Println(loginResp)
}

// 读取终端输入
// 这个函数没有意义
//func Input() *TestArgs {
//	var t *TestArgs
//	var input *bufio.Reader
//
//	input = bufio.NewReader(os.Stdin)
//
//	fmt.Printf("Please enter item id:")
//	itemId, err := input.ReadString('\n')
//	if err != nil {
//		log.Println("input item_id", err)
//	}
//	t.id, err = strconv.ParseInt(strings.TrimSpace(itemId), 10, 64)
//	if err != nil {
//		log.Println("itemid strconv.ParseInt failed: ", err)
//	}
//	//consumeClient.Pay(ctx, &pb.ConsumeReq{ItemId: id})
//	log.Println("itemid :", t.id) // input
//	//log.Println("itemid type:", reflect.TypeOf(itemid))		// int64
//
//
//	fmt.Printf("Please enter Count:")
//	num, err := input.ReadString('\n')
//	if err != nil {
//		log.Println("input itemCount", err)
//	}
//	t.count, _ = strconv.ParseInt(num, 10, 64)
//	//consumeClient.Pay(ctx, &pb.ConsumeReq{ItemNum: num})
//
//	fmt.Printf("Please enter userId:")
//	uuid, err := input.ReadString('\n')
//	if err != nil {
//		log.Println("input userid", err)
//	}
//	t.user, _ = strconv.ParseInt(uuid, 10, 32)
//	//consumeClient.Pay(ctx, &pb.ConsumeReq{UserId: int32(constomerid)})
//
//	fmt.Printf("remark or not:")
//	t.descri, err = input.ReadString('\n')
//	if t.descri == "" {
//		t.descri = "nil"
//	}
//	if err != nil {
//		log.Println("description", err)
//	}
//	//consumeClient.Pay(ctx, &pb.ConsumeReq{Description: descr})
//	return TestArgs{
//		id:     ,
//		count:  0,
//		user:   0,
//		descri: "",
//	}
//}