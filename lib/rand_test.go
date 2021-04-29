package alg

import (
	crypto_rand "crypto/rand"
	"encoding/binary"
	"fmt"
	"math/rand"
	"testing"
)

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
