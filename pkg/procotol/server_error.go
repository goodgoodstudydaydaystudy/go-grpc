package protocol

import (
	"fmt"

	"goodgoodstudy.com/go-grpc/protocol/common/status"
)

var ErrSystem = NewServerError(status.ErrSys)

// ServerError code/message
type ServerError interface {
	error
	Code() int
	Message() string
}

type serverError struct {
	code       int
	message    string
	underlying error
}

// NewServerError create a new instance of ServerError
// with the specific status and message
func NewServerError(code int, message ...string) ServerError {
	msg := ""
	if len(message) > 0 { //&& len(message[0]) > 0 {
		msg = message[0]
	} else {
		msg = status.MessageFromCode(code)
	}

	return &serverError{code: code, message: msg}
}

// Code returns the status code of ServerError
func (e *serverError) Code() int {
	return e.code
}

// Message returns the message of ServerError
func (e *serverError) Message() string {
	return e.message
}

func (e *serverError) Error() string {
	return fmt.Sprintf("status: %d, message: %s, underlying: %v", e.code, e.message, e.underlying)
}

func ToServerError(err error) ServerError {
	if err == nil {
		return nil
	}

	if pe, ok := err.(ServerError); ok {
		if pe.Code() >= 0 { // >= 0 的code认为是成功, 不返回错误; 但是这里既然有error, 那就返回默认错误; 应该打一行日志
			// log here?
			return ErrSystem
		}
		return pe
	}

	return ErrSystem
}
