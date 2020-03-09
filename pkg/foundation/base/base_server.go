package base
// base包： 封装main的"注册拦截器，"signal注册"，"反射"
// 创建人：
// 创建日期： 20200227

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

// 定义服务入口的基础参数
type cmdBaseObj struct {
	GrpcServer  *grpc.Server
	Signal      chan bool
	Listen      net.Listener
}

// NewCmdBase, 封装一个服务入口所需参数实例
// 参数:
// 		sop: 拦截器, netType: 连接类型, address: 端口地址
// 返回:
// 		参数类的指针
func NewCmdBase(sop grpc.ServerOption, netType string, address string) (c *cmdBaseObj) {
	c = &cmdBaseObj{}
	c.GrpcServer = c.registerGRPCServer(sop)
	c.Listen = c.netListen(netType, address)
	c.Signal = c.registerSignal()
	c.registerReflection(c.GrpcServer)
	return
}


func (c *cmdBaseObj)registerSignal() chan bool {
	exit := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(exit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<- exit
		done <- true
	}()
	return done
}


func (c *cmdBaseObj)registerGRPCServer(sop grpc.ServerOption) *grpc.Server {
	s := grpc.NewServer(sop)
	return s
}


func (c *cmdBaseObj)netListen(netType string, address string) net.Listener {
	lis, err := net.Listen(netType, address)
	if err != nil {
		log.Fatalf("wallet server listen failed")
		return nil
	}
	return lis
}

func (c *cmdBaseObj)registerReflection(s *grpc.Server) {
	reflection.Register(s)
	return
}