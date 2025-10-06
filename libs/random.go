package libs

import (
	"math/rand"
)

func RandRangeI(min, max int) int {
	return rand.Intn(max-min+1) + min
}

func RandRangeF(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}
