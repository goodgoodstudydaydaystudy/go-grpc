package server

import (
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"

	apb "goodgoodstudy.com/go-grpc/pkg/pb/server/account"
)

type CustomClaims struct {
	userInfo  *apb.UserInfo
	jwt.StandardClaims
}

func ParseWithClaims(tokenString string) (*CustomClaims, error) {

	// sample token is expired.  override time so it parses as valid
	// 样本口令已过期，使其有效需重写时间
	at(time.Unix(0, 0), func() {
		token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (i interface{}, err error) {
			return []byte("66666"), nil
		})

		if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
			log.Printf("%v", claims.StandardClaims.ExpiresAt)
			return
		}else {
			log.Println(err)
			return
		}
	})
	return nil, nil
}

// 8 Override time value for tests.  Restore default value after. 超时之后，回复默认值
func at(t time.Time, f func())  {
	jwt.TimeFunc = func() time.Time {
		return t
	}
	f()
	jwt.TimeFunc = time.Now
}