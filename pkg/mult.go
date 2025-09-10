package pkg

import "math/rand/v2"

func Mult(rtp float64) float64 {
	if rand.Float64() < rtp {
		return 9000 + rand.NormFloat64()*1000
	}
	return 1.0 + rand.NormFloat64()*0.5*1000
}
