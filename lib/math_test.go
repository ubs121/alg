package alg

import (
	crypto_rand "crypto/rand"
	"encoding/binary"
	"fmt"
	"math"
	"math/rand"
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

func TestRandInt(t *testing.T) {
	min := 1
	max := 100
	var b [8]byte
	_, err := crypto_rand.Read(b[:])
	if err != nil {
		t.Error("cannot seed with crypto rand")
	}
	crypto_rand.Read(b[:])
	rand.Seed(int64(binary.LittleEndian.Uint64(b[:])))

	for i := 0; i < 10; i++ {
		val := rand.Intn(max-min) + min
		fmt.Printf("%d\n", val)
	}
}
