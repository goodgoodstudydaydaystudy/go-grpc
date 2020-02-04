package gateway

import (
	"context"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"

	gw "goodgoodstudy.com/go-grpc/pkg/pb/logic/user"
)

type Gateway struct{}

func (g *Gateway) ListenAndServe(addr string, logicAddr string) error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()

	// hack
	httpMux := http.NewServeMux()
	httpMux.Handle("/", mux)
	httpMux.Handle("/v1/swagger/", swaggerServer("."))
	serveSwaggerUI(httpMux)
	// hack end

	opts := []grpc.DialOption{grpc.WithInsecure()}

	err := gw.RegisterUserHandlerFromEndpoint(ctx, mux, logicAddr, opts)
	if err != nil {
		return err
	}

	return http.ListenAndServe(addr, httpMux)
}
