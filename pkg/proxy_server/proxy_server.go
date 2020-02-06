package main

import (
	"log"

	"goodgoodstudy.com/go-grpc/pkg/proxy_server/gateway"
)

func main() {
	g := &gateway.Gateway{}
	if err := g.ListenAndServe("127.0.0.1:9091", "localhost:50053"); err != nil {
		log.Fatal(err)
	}
}
