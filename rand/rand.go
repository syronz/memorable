package rand

import (
	"math/rand"
	"time"
)

func RandomNumber() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(30)
}
