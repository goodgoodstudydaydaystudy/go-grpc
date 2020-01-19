# go-grpc
2020/1/6
```
server/Server.go

func (s *ControlServer) Pay(ctx context.Context, consumeReq *pb.ConsumeReq) (*pb.ConsumeResp, error) {
	return &pb.ConsumeResp{OrderId: consumeReq.GetItemId()}, nil
 ```
 Q: OrderId明明是ConsumeResp结构体里面的方法，怎么会引出consumeReq结构体和里面的字段唉

2020/1/7
```
Client的方法，不知道对不对。
```

- [x] 规范命名
- [x] 简单了解client和server的拦截器
- [ ] Account Server
  - [x] 在server启动的时候建立数据库连接, 建立失败则启动失败
  - [x] 在server关闭的时候关闭数据库连接
  - [ ] 尝试返回自定义错误码
  - [ ] server结构体私有化, 提供New函数
- [ ] Account Client
  - [ ] client结构体私有化
  - [ ] 测试自定义错误码能不能用
- [ ] 仿照account修改pay的server和client, 先不要做pay的db
