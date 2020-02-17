package encode

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"strconv"
	"time"
)

// Md5
// 方式：nowTime + randomNum + seed(string)
func GenerateMd5(seed string) string {
	now := time.Now()

	seedTime := now.Format("20060102150405.000")

	randomNum := rand.New(rand.NewSource(time.Now().UnixNano())).Int63n(1000000)
	randomNumStr := strconv.FormatInt(randomNum, 10)

	newMd5 := md5.New()
	data := []byte(seedTime + randomNumStr+ seed)
	newMd5.Write(data)
	b := md5.Sum(data)
	code := hex.EncodeToString(b[:])

	return code
}

