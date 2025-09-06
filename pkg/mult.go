package pkg

import "math/rand/v2"

func Mult(rtp float64) float64 {
	if rand.Float64() < rtp {
		return 10_000.0
	}
	return 1.0
}
