package random

import (
	"math/rand"
	"time"
)

// Generate generate a random number with time as seed
func Generate(max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max)
}

// GenerateMinMax is used for create random number inside a range
func GenerateMinMax(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}
