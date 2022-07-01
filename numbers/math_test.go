package numbers

import (
	crypto_rand "crypto/rand"
	"encoding/binary"
	"fmt"
	"math"
	"math/big"
	"math/rand"
	"strings"
	"testing"
)

func TestRemainder(t *testing.T) {
	const epsilon = 0.00000001
	got := math.Remainder(math.Log(float64(1162261466)), math.Log(3))
	if math.Abs(got-0.0) > epsilon {
		t.Errorf("exp: 0.0, got %f", got)
	}
}

func TestAddFloat(t *testing.T) {
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

// https://leetcode.com/problems/multiply-strings/
func multiplyStrings(num1 string, num2 string) string {
	a := new(big.Int)
	b := new(big.Int)
	a.SetString(num1, 10)
	b.SetString(num2, 10)

	a.Mul(a, b)

	return a.String()
}

func TestMultiplyStrings(t *testing.T) {
	testCases := map[string]string{
		"123*456": "56088",
	}

	for tc, exp := range testCases {
		splits := strings.Split(tc, "*")
		got := multiplyStrings(splits[0], splits[1])
		if got != exp {
			t.Errorf("tc %s: exp %s, got %s", tc, exp, got)
		}
	}
}
