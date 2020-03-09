package account

import (
	"context"
	rpb "goodgoodstudy.com/go-grpc/pkg/pb/server/account"
	"math/rand"
	"testing"
	"time"
)

func TestServer_CheckPwd(t *testing.T) {
	s, err := NewAccountServer()
	if err != nil {
		t.Fatal(err)
	}

	ctx := context.Background()
	resp, err := s.CheckPwd(ctx, &rpb.CheckPwdReq{
		UserId:               562,
		Password:             "123456",
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Log(resp)
}


func BenchmarkServer_CheckPwd(b *testing.B) {
	srv, err := NewAccountServer()
	if err != nil {
		b.Fatal(err)
	}

	ctx := context.Background()

	b.ResetTimer()
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			user := generateInfo()
			req := &rpb.CheckPwdReq{
				UserId:               user.userId,
				Password:             "123456",
			}
			resp, err := srv.CheckPwd(ctx, req)
			if err != nil {
				b.Fatal(err)
			}
			b.Log(resp)
		}
	})
}

func BenchmarkServer_CheckPwd2(b *testing.B) {
	srv, err := NewAccountServer()
	if err != nil {
		b.Fatal(err)
	}

	ctx := context.Background()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		user := generateInfo()
		req := &rpb.CheckPwdReq{
			UserId:               user.userId,
			Password:             "123456",
		}
		resp, err := srv.CheckPwd(ctx, req)
		if err != nil {
			b.Fatal(err)
		}
		b.Log(resp)
	}
}


type Info struct {
	userId   int32
}

func generateInfo() *Info {
	userIdInt := rand.New(rand.NewSource(time.Now().UnixNano())).Int63n(10000)
	if userIdInt == 0 {
		userIdInt += 1
	}
	return &Info{
		userId:   int32(userIdInt),
	}
}
