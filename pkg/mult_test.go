package pkg

import (
	"math"
	"math/rand/v2"
	"testing"
)

func TestMult(t *testing.T) {
	n := 10_000
	rtp := 0.75
	seq := GenSequence(n)

	sum0 := 0.0
	sum1 := 0.0
	for i, v := range seq {
		sum0 += seq[i]
		if Mult(rtp) <= v {
			seq[i] = 0
		}
		sum1 += seq[i]
	}

	RTP := sum1 / sum0
	if math.Abs(RTP-rtp) > 0.1 {
		t.Errorf("Client RTP is greater than the specified rtp by: %f", RTP-rtp)
	}
}

func GenSequence(n int) []float64 {
	r := make([]float64, n)
	for i := range n {
		r[i] = 1 + rand.Float64()*9999.0
	}
	return r
}
