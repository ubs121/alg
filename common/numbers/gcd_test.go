package numbers

import "testing"

//also see https://golang.org/pkg/math/big/#Int.GCD

// greatest common divisor
func GCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// greatest common divisor
func gcd2(u, v uint) uint {
	// simple cases (termination)
	if u == v {
		return u
	}

	if u == 0 {
		return v
	}

	if v == 0 {
		return u
	}

	// look for factors of 2
	if u%2 == 0 { // u is even
		if v%2 == 1 { // v is odd
			return gcd2(u>>1, v)
		}

		// both u and v are even
		return gcd2(u>>1, v>>1) << 1
	}

	if v%2 == 0 { // u is odd, v is even
		return gcd2(u, v>>1)
	}

	// reduce larger argument
	if u > v {
		return gcd2((u-v)>>1, v)
	}

	return gcd2((v-u)>>1, u)
}

func TestGCD(t *testing.T) {
	if r := GCD(10, 6); r != 2 {
		t.Errorf("gcd(10,6): exp 2, got %d", r)
	}

	if r := gcd2(12, 8); r != 4 {
		t.Errorf("gcd(12,8): exp 4, got %d", r)
	}
}
