# go-grpc
2020/1/6
```
server/Server.go

func (s *ControlServer) Pay(ctx context.Context, consumeReq *pb.ConsumeReq) (*pb.ConsumeResp, error) {
	return &pb.ConsumeResp{OrderId: consumeReq.GetItemId()}, nil
 ```
 Q: OrderId明明是ConsumeResp结构体里面的方法，怎么会引出consumeReq结构体和里面的字段唉

日记：

1、将用户信息以结构体的形式封装传递，便于对用户信息的增删
2、数据库返回的error内容也应该经过拦截器


- [x] 规范命名
- [x] 简单了解client和server的拦截器
- [x] Account Server
  - [x] 在server启动的时候建立数据库连接, 建立失败则启动失败
  - [x] 在server关闭的时候关闭数据库连接
  - [x] 尝试返回自定义错误码
  - [x] server结构体私有化, 提供New函数
- [x] Account Client
  - [x] client结构体私有化
  - [x] 测试自定义错误码能不能用
- [x] 仿照account修改pay的server和client, 先不要做pay的db

- [ ] 注册增加“性别”，登录resp返回增加“性别”
  - [ ] 简化新增属性的修改工作
