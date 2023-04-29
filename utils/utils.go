package utils

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateAccountNumber(length int) string {
	rand.Seed(time.Now().UnixNano())
	min := int64(pow(10, length-1))
	max := int64(pow(10, length) - 1)
	number := rand.Int63n(max-min+1) + min
	return fmt.Sprintf("%0*d", length, number)
}

func pow(x, y int) int {
	p := 1
	for i := 0; i < y; i++ {
		p *= x
	}
	return p
}
