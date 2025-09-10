package pkg

import (
	"math"
	"math/rand/v2"
)

func Mult(rtp float64) float64 {
	if rand.Float64() > rtp {
		value := math.Pow(rand.Float64(), 8)
		return 1 + value*4999
	} else {
		value := 1 - math.Pow(1-rand.Float64(), 8)
		return 5000 + value*5000
	}
}
