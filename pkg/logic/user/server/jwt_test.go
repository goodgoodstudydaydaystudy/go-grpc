package server

import (
	"log"
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type userInfo struct {
	id 			int
	account 	string
	gender 		int
}

type CustomClaimsTest struct {
	*userInfo
	jwt.StandardClaims
}

func TestPwt(t *testing.T)  {
	//token1, _ := newWithClaims_standardClaims()
	//parseWithClaims_customClaimsType(token1)
	user := &userInfo{
		id:      8,
		account: "test01",
		gender:  1,
	}
	token2, _ := newWithClaimsCustomClaims(*user)
	r, _ := parseWithClaims(token2)
	t.Log(r)
	//parse_errorCheck(token2)
}

// 5 new claims
func newWithClaims_standardClaims() (string, error) {
	mySigningKey := []byte("water")

	// create claims
	claims := &jwt.StandardClaims{
		Issuer:		"logicLogin",
		ExpiresAt:  15000,
		Subject: 	"test",
		IssuedAt:   time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	log.Printf("%v %v", ss, err)
	return ss, err
}

// 6 creat a token using a custom claims type
func newWithClaimsCustomClaims(user userInfo) (string, error) {
	mySigningKey := []byte("77777")
	// 和上栗的差别, 这个可以加更多额外的字段耶


	claims := &CustomClaims{
		&user,
		jwt.StandardClaims{
			ExpiresAt: 15000,
			Issuer:    "test",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	log.Printf("%v %v", ss, err)
	return ss, err
}

// 7 parse custom claims type
func parseWithClaims(tokenString string) (*CustomClaims, error) {


	// sample token is expired.  override time so it parses as valid   这句话怎么有点奇怪
	At(time.Unix(0, 0), func() {
		token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (i interface{}, err error) {
			return []byte("77777"), nil
		})

		if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
			log.Printf("%v %v", claims.userInfo, claims.Issuer)
			return
		}else {
			log.Println(err)
			return
		}
	})
	return nil, nil
}

// 8 Override time value for tests.  Restore default value after. 超时之后，回复默认值
func At(t time.Time, f func())  {
	jwt.TimeFunc = func() time.Time {
		return t
	}
	f()
	jwt.TimeFunc = time.Now
}

func parse_errorCheck(tokenString string)  {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (i interface{}, err error) {
		return []byte("water"), nil
	})
	if token.Valid {
		log.Println("ok")
	}else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			log.Println("not even a token")
		}else if ve.Errors&(jwt.ValidationErrorMalformed | jwt.ValidationErrorNotValidYet) != 0 {
			// Token is either expired or not active yet
			log.Println("Token expired or not active")
		}else {
			log.Println("could not handle this token: ", err)
		}
	}else {
		log.Println("could not handle this token: ", err)
	}
}
