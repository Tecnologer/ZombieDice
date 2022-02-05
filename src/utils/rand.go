package utils

import (
	"math/rand"
	"time"
)

func GetRandInt(limit int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(limit)
}

func GetRandIntRange(start, limit int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(limit-start) + start
}
