# go-grpc
2020/1/6
```
server/Server.go

func (s *ControlServer) Pay(ctx context.Context, consumeReq *pb.ConsumeReq) (*pb.ConsumeResp, error) {
	return &pb.ConsumeResp{OrderId: consumeReq.GetItemId()}, nil
 ```
 Q: OrderId明明是ConsumeResp结构体里面的方法，怎么会引出consumeReq结构体和里面的字段唉
