package server

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"log"
	"time"

	"goodgoodstudy.com/go-grpc/client/account"
	"goodgoodstudy.com/go-grpc/client/wallet"
	pb "goodgoodstudy.com/go-grpc/pkg/pb/logic/user"
	apb "goodgoodstudy.com/go-grpc/pkg/pb/server/account"
)

type CustomClaims struct {
	UserInfo *pb.UserInfo `json:"user_info"`
	jwt.StandardClaims
}


type UserLogic struct {
	accountClient *account.Client
	walletClient  *wallet.Client
}

func NewUserLogic() (*UserLogic, error) {
	accountClient, err := account.NewAccountClient()
	if err != nil {
		log.Println("logic server NewAccountClient failed: ", err)
		return nil, err
	}
	walletClient, err := wallet.NewWalletClient()
	if err != nil {
		log.Println("logic server NewWalletClient failed: ", err)
		return nil, err
	}


	return &UserLogic{
		accountClient: accountClient,
		walletClient:  walletClient,
	}, nil
}


// 2. account的登录服务
func (s *UserLogic) Login(ctx context.Context, req *pb.LoginReq) (resp *pb.LoginResp, err error) {
	resp = &pb.LoginResp{}

	// 1. check password
	// 这里不能返回，否则grpc 框架会报错。
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

	// generate token
	token, err := s.generateToken(*resp.UserInfo)
	if err != nil {
		log.Println("new claims failed: ", err)
		return
	}
	resp.Token = token

	// 2. get balance
	userBalance, err := s.walletClient.GetUserById(ctx, resp.UserInfo.UserId)
	if err != nil {
		log.Println("logic.Login get balance failed: ", err)
		return
	}
	resp.Balance = userBalance.Balance

	// 3. return
	return resp, nil
}

// 3.  account的register服务
func (s *UserLogic) Register(ctx context.Context, req *pb.RegisterReq) (resp *pb.RegisterResp, err error) {
	resp = &pb.RegisterResp{}
	// 3.1 提交注册信息
	r, err := s.accountClient.AddUser(ctx, req.Account, req.Password, req.Nickname, apb.Gender(req.Gender))
	if err != nil {
		log.Println("logic Register register failed: ", err)
		return
	}
	resp.UserId = r.UserId
	// 3.2 return
	return resp, nil
}

// 4. 充值服务
func (s *UserLogic) Recharge(ctx context.Context, req *pb.RechargeReq) (resp *pb.RechargeResp, err error) {
	resp = &pb.RechargeResp{}

	// 4.1 发送请求
	_, err = s.walletClient.Recharge(ctx, req.Account, req.UserId, req.Delta)
	if err != nil {
		log.Println("logic serve Recharge failed: ", err)
		return
	}
	return resp, nil
}

func (s *UserLogic) GetTopUser(ctx context.Context, req *pb.GetTopUserReq) (resp *pb.GetTopUserResp, err error){
	resp = &pb.GetTopUserResp{}

	r, err := s.walletClient.GetTopUser(ctx, req.Top)
	if err != nil {
		log.Println("logic server GetTopUser failed:", err)
		return
	}
	resp.UserList = r.UserList
	log.Println("r.UserList:", r.UserList)
	return resp, nil
}

// 5 generate claims
func (s *UserLogic) generateToken(userInfo pb.UserInfo) (ss string, err error) {
	mySigningKey := []byte("66666")

	durTime, _ := time.ParseDuration("10m")
	expireTime := time.Now().Add(durTime).Unix()

	// create claims
	claims := &CustomClaims{
		UserInfo: &userInfo,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime,
			Subject:   "test",
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err = token.SignedString(mySigningKey)
	if err != nil {
		log.Println("server token failed:", err)
		return ss, err
	}
	return ss, nil
}

func (s *UserLogic) WriteNoPaidOrder(ctx context.Context, req *pb.WriteNoPaidOrderReq) (resp *pb.WriteNoPaidOrderResp, err error) {
	resp = &pb.WriteNoPaidOrderResp{}
	_, err = s.walletClient.WriteNoPaidOrder(ctx, req.UserId, req.OrderId)
	if err != nil {
		log.Println("logic server RecordOrderNoPaid failed:", err)
		return 
	}
	return resp, nil
}

func (s *UserLogic) Pay(ctx context.Context, req *pb.PayReq) (resp *pb.PayResp, err error) {
	resp = &pb.PayResp{}

	_, err = s.walletClient.Pay(ctx, req.OrderId)
	if err != nil {
		log.Println("logic server Pay failed:", err)
		return
	}
	return resp, nil
}