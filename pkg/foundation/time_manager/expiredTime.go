package time_manager

import (
	"math/rand"
	"time"
)

func ExpiredTime() time.Duration{
	seed := rand.Int63n(10000000)
	return time.Duration(seed)*time.Millisecond
}

func ExpiredTimeSort() time.Duration {
	seed := rand.Int63n(100000)
	return time.Duration(seed)*time.Millisecond
}