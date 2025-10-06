package libs

import (
	"math"
)

func Normalize(x, y float64) (float64, float64) {
	length := math.Sqrt(x*x + y*y)
	if length == 0 {
		return 0, 0 // avoid division by zero
	}
	return x / length, y / length
}
