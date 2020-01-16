// MD5包: String MD5加密
// 创建人：LU
// 创建时间：XXXXXXXX
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
