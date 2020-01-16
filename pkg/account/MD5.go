package account

import (
	"crypto/md5"
	"encoding/hex"
)


// MD5加密
func Encryption(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}
