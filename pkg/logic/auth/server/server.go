/**
 * Author: Orange
 * Date: 20-02-04
 */
package server

import (
	"context"
	"log"

	"github.com/dgrijalva/jwt-go"

	"goodgoodstudy.com/go-grpc/client/account"
	"goodgoodstudy.com/go-grpc/internal"
	"goodgoodstudy.com/go-grpc/pkg/foundation/grpc/server/grpcAuth"
	pb "goodgoodstudy.com/go-grpc/pkg/pb/logic/auth"
	accountPB "goodgoodstudy.com/go-grpc/pkg/pb/server/account"
)

const codeExpiresIn = int64(10)

type AuthLogic struct {
	pb.AuthServer
	accountClient *account.Client
}

func NewAuthLogic() (*AuthLogic, error) {
	accountClient, err := account.NewAccountClient()
	if err != nil {
		log.Println("logic server NewAccountClient failed: ", err)
		return nil, err
	}

	return &AuthLogic{
		accountClient: accountClient,
	}, nil
}

func (s *AuthLogic) Login(ctx context.Context, req *pb.LoginReq) (resp *pb.LoginResp, err error) {
	resp = &pb.LoginResp{}

	r, err := s.accountClient.CheckPwd(ctx, req.Account, req.Password)
	if err != nil {
		log.Println("logic.Login check password failed: ", err)
		return
	}
	resp.UserInfo = &pb.UserInfo{
		UserId:   r.UserInfo.UserId,
		Account:  r.UserInfo.Account,
		Nickname: r.UserInfo.Nickname,
		Gender:   pb.Gender(r.UserInfo.Gender),
	}

	claims := grpcAuth.NewJWTClaims(codeExpiresIn)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(internal.SecretKey))
	log.Println("token:", ss)

	resp.Token = ss
	return resp, nil
}

func (s *AuthLogic) Register(ctx context.Context, req *pb.RegisterReq) (resp *pb.RegisterResp, err error) {
	resp = &pb.RegisterResp{}
	r, err := s.accountClient.AddUser(ctx, req.Account, req.Password, req.Nickname, accountPB.Gender(req.Gender))
	if err != nil {
		log.Println("logic Register register failed: ", err)
		return
	}
	resp.UserId = r.UserId
	return resp, nil
}
