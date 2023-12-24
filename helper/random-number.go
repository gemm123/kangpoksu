package helper

import (
	"math/rand"
	"time"
)

func RandomNumberOrder(totalOrder int) int {
	rand.Seed(time.Now().UnixNano())
	randNumber := rand.Intn(90) + 10
	totalOrder = totalOrder / 100
	totalOrder = totalOrder*100 + randNumber

	return totalOrder
}
