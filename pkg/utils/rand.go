package utils

import (
	"math/rand"
	"time"
)

// RandomInt returns, as an int64, a negative pseudo-random number in [min,max]
func RandomInt(min, max int) int64 {
	rand.Seed(time.Now().UnixNano())
	return int64(rand.Intn(max-min+1) + min)
}
