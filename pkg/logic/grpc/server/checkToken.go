package server

import (
	"github.com/dgrijalva/jwt-go"
	pb "goodgoodstudy.com/go-grpc/pkg/pb/server/account"
	"log"
)

func checkToken(tokenString string) error {
	// 创建claims

	type customsClaims struct {
		UserInfo *pb.UserInfo `json:"user_info"`
		jwt.StandardClaims
	}
	// 获得token
	token, err := jwt.ParseWithClaims(tokenString, &customsClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return []byte("66666"), nil
	})

	if token.Valid {
		return nil
	}else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0{
			log.Println("That's not even a token")
			return err
		}else if ve.Errors&(jwt.ValidationErrorExpired | jwt.ValidationErrorNotValidYet) != 0 {
			log.Println("Token is either expired or not active yet:", err)
			return err
		}else {
			log.Println("Couldn't handle this token:", err)
			return err
		}
	}else {
		return err
	}
}
