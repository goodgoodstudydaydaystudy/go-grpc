/**
 * Author: Orange
 * Date: 20-02-06
 */
package grpcAuth

import (
	"context"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"goodgoodstudy.com/go-grpc/internal"
)

type selfJWTClaims struct {
	jwt.StandardClaims
}

func NewJWTClaims(expireSeconds int64) *selfJWTClaims {
	claims := &selfJWTClaims{}
	now := time.Now().Unix()

	claims.ExpiresAt = now + expireSeconds
	claims.IssuedAt = now
	claims.Issuer = "tt-oversea"
	return claims
}

func (builder *authFuncBuilder) BuildJWT() grpc_auth.AuthFunc {
	return func(ctx context.Context) (context.Context, error) {
		log.Println("AuthenticationInterceptor")

		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return ctx, errUnauthenticated
		}

		if fullMethod, ok := grpc.Method(ctx); ok {
			log.Println(fullMethod, builder.fullMethodNameException, builder.fullMethodNameException[fullMethod])
			if builder.fullMethodNameException[fullMethod] {
				return ctx, nil
			}
		}

		_tokenStr, ok := md["token"]
		if !ok || len(_tokenStr) == 0 {
			log.Println("token empty")
			return ctx, errUnauthenticated
		}

		token, err := jwt.ParseWithClaims(_tokenStr[0], &selfJWTClaims{}, func(token *jwt.Token) (i interface{}, err error) {
			return []byte(internal.SecretKey), nil
		})
		if err != nil {
			return ctx, errUnauthenticated
		}

		if !token.Valid {
			if ve, ok := err.(*jwt.ValidationError); ok {
				if ve.Errors&jwt.ValidationErrorMalformed != 0 {
					log.Println("not even a token")
					// This is not even a token
					return ctx, errUnauthenticated
				} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
					log.Println("token is either expired or not active yet")
					// Token is either expired or not active yet
					return ctx, errTokenExpired
				} else {
					log.Println("can not parse token")
					return ctx, errInternal
				}
			} else {
				log.Println("internal error")
				return ctx, errInternal
			}
		}

		var claims *selfJWTClaims
		if claims, ok = token.Claims.(*selfJWTClaims); !ok {
			log.Println("not selfClaims")
			return ctx, errUnauthenticated
		}

		// token is valid
		// TODO may add info into HEADER
		_ = claims
		//ctx, _ := context.WithTimeout(ctx, time.Duration(time.Second*2))
		return ctx, nil
	}
}
