package alg

import (
	"math"
	"testing"
)

func TestRemainder(t *testing.T) {
	const epsilon = 0.00000001
	got := math.Remainder(math.Log(float64(1162261466)), math.Log(3))
	if math.Abs(got-0.0) > epsilon {
		t.Errorf("exp: 0.0, got %f", got)
	}
}

func TestAdd(t *testing.T) {
	got := float64(.3) + float64(.6) + float64(.1)
	if i := math.Trunc(got) - got; i < 1e-9 && i > -1e-9 {
		t.Errorf("exp: 1.0, got %f, %f", got, i)
	}
}
