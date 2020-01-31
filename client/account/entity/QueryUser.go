package account

import pb "goodgoodstudy.com/go-grpc/pkg/pb/account"

type GetUser struct {
	QueryCode uint32
}

func QueryReqToPb(queryReq *GetUser) *pb.QueryUserReq {
	return &pb.QueryUserReq{
		QueryUser:	pb.QueryIndex(queryReq.QueryCode),
	}
}

func PbToQueryReq(pbQuery *pb.QueryUserReq) *GetUser {
	return &GetUser{
		QueryCode:	uint32(pbQuery.QueryUser),
	}
}
