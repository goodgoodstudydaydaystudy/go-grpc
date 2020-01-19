package grpc

import (
	"context"
	"strconv"

	grpc1 "google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"goodgoodstudy.com/go-grpc/protocol/common/status"
)

const (
	grpcHeaderCode          = "resp_code"
	grpcHeaderMessage       = "resp_msg"
	grpcHeaderRequestUserID = "req_uid"
)

func SetStatus(ctx context.Context, code int, message ...string) {
	var msg string
	if len(message) > 0 {
		msg = message[0]
	} else {
		msg = status.MessageFromCode(code)
	}

	_ = grpc1.SetHeader(ctx, metadata.Pairs(
		grpcHeaderCode, strconv.Itoa(code),
		grpcHeaderMessage, msg,
	))
}

func GetStatus(md metadata.MD) (code int, message string) {
	if v, ok := md[grpcHeaderCode]; ok && len(v) > 0 {
		code, _ = strconv.Atoi(v[0])
	}

	if v, ok := md[grpcHeaderMessage]; ok && len(v) > 0 {
		message = v[0]
	}
	return
}
