package utils

import (
	"math/rand"
	"time"
)

func GetRandInt(limit int) int {
	time.Sleep(1 * time.Millisecond)
	r := rand.New(rand.NewSource(time.Now().Unix()))
	return r.Intn(limit)
}
