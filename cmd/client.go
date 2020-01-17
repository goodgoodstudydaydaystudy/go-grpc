package main

import (
	"context"
	ap "goodgoodstudy.com/go-grpc/client/account"
	pp "goodgoodstudy.com/go-grpc/client/pay"
	rpb "goodgoodstudy.com/go-grpc/pkg/pb/account"
	pb "goodgoodstudy.com/go-grpc/pkg/pb/pay"
	md "goodgoodstudy.com/go-grpc/pkg/utils"
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
	consumeClient, err := pp.NewConsumeClient()
	if err != nil {
		log.Println("new client:", err)
	}

	accountClient, err := ap.NewAccountClient()
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
func Consume(ctx context.Context, consumeClient *pp.ConsumeClient)  {
	consumeResp, err := consumeClient.Pay(ctx, &pb.ConsumeReq{ItemId: 1, ItemNum: 2, UserId:3, Description:"aaaa"})
	if err != nil {
		log.Println("consume failed,", err)
		return
	}
	log.Println(consumeResp)
}

// 注册接口
func Register(ctx context.Context, accountClient *ap.Client) (message string) {
	// 判断账密长度
	//accountLen := strings.Count(accountStr, "")
	//passwordLen := strings.Count(passwordStr, "")
	//
	//switch {
	//case accountLen > 15:
	//	return &rpb.RegisResp{Message:"account too long"}, nil
	//case accountLen < 6 :
	//	return &rpb.RegisResp{Message:"account too short"}, nil
	//case passwordLen > 20:
	//	return &rpb.RegisResp{Message:"password too long"}, nil
	//case passwordLen < 8 :
	//	return &rpb.RegisResp{Message:"password too short"}, nil

	md5Account := md.Encryption("7777777")
	md5Password := md.Encryption("6666666")

	registerResp, err := accountClient.Register(ctx, &rpb.RegisterReq{Account:md5Account, Password:md5Password})
	if err != nil {
		log.Println("accountClient.Register failed: ", err)
	}
	log.Println(registerResp)
	return
}

// 登录接口
func Login(ctx context.Context, accountClient *ap.Client)  {
	loginResp, err := accountClient.Login(ctx, "account77", "password8888")
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