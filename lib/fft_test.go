package alg

import (
	"fmt"
	"math"
	"math/cmplx"
	"testing"
)

func TestFFT(t *testing.T) {
	fft := &FFT{}

	in := []complex128{1, 2, 3, 4}
	exp := []complex128{10, -2 + 2i, -2, -2 - 2i}

	got := fft.Transform(in)
	for i := 0; i < len(exp); i++ {
		if math.Round(real(exp[i])) != math.Round(real(got[i])) ||
			math.Round(imag(exp[i])) != math.Round(imag(got[i])) {
			t.Errorf("Sample1: exp %v, got %v", exp[i], got[i])
		}
	}

	in = []complex128{0, 0, 0, 1}
	exp = []complex128{1, 1i, -1, -1i}

	got = fft.Transform(in)
	for i := 0; i < len(exp); i++ {
		if math.Round(real(exp[i])) != math.Round(real(got[i])) ||
			math.Round(imag(exp[i])) != math.Round(imag(got[i])) {
			t.Errorf("Sample2: exp %v, got %v", exp[i], got[i])
		}
	}
}

func TestCircularMatrixMulti(t *testing.T) {
	/* U does a right shift by 1 on a vector
	U= [[1,1,0,0],
		[0,1,1,0],
		[0,0,1,1],
		[1,0,0,1]]

	exp:=fft.Inverse( diag(U)*fft.Transform(vector) )
	*/
	k := 2 // k times
	vector := []complex128{0, 0, 0, 1}
	n := len(vector)

	fft := &FFT{}
	vector1 := fft.Transform(vector) // fft(vector)

	sin, cos := math.Sincos(2 * float64(math.Pi) * float64(k) / float64(n))
	w := complex(cos, sin)

	// diagonal matrix of U, eigen values in the diagonal
	diagU := make([]complex128, n)
	kCmp := complex(float64(k), 0)
	for i := 0; i < n; i++ {
		iCmp := complex(float64(i), 0)
		diagU[i] = cmplx.Pow(1+cmplx.Pow(w, iCmp), kCmp)
	}

	// multiply by diag(U)
	for i := 0; i < n; i++ {
		vector1[i] = diagU[i] * vector1[i]
	}

	vector2 := fft.Inverse(vector1) // ifft(vector1)

	// apply modulo 2
	ret := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Printf("%v", int(real(vector2[i])))
		ret[i] = int(real(vector2[i])) % 2
	}

	// validate result
	exp := []int{0, 1, 0, 1}
	for i := 0; i < n; i++ {
		if exp[i] != ret[i] {
			t.Errorf("Component %d: exp %d, got %d", i, exp[i], ret[i])
		}
	}
}
