package alg

import (
	"math"
	"math/cmplx"
)

// Fast Fourier Transformer
type FFT struct {
}

// Transform returns the FFT of 'x'
func (fft *FFT) Transform(x []complex128) []complex128 {
	n := len(x)

	if n == 1 {
		return []complex128{x[0]}
	}

	if n%2 != 0 {
		panic("len(x) is not a power of 2")
	}

	// even terms
	even := make([]complex128, n/2)
	for i := 0; i < n/2; i++ {
		even[i] = x[2*i]
	}
	even1 := fft.Transform(even)

	// odd terms
	odd := even // re-using the allocated memory
	for i := 0; i < n/2; i++ {
		odd[i] = x[2*i+1]
	}
	odd1 := fft.Transform(odd)

	// combine
	cmb := make([]complex128, n)
	for k := 0; k < n/2; k++ {
		kth := -2.0 * float64(k) * float64(math.Pi) / float64(n)
		wk := complex(math.Cos(kth), math.Sin(kth))
		cmb[k] = even1[k] + wk*odd1[k]
		cmb[k+n/2] = even1[k] - wk*odd1[k]
	}
	return cmb
}

// Inverse returns the inverse FFT of 'x'
func (fft *FFT) Inverse(x []complex128) []complex128 {
	n := len(x)

	// conjugate
	y := make([]complex128, n)
	for i := 0; i < n; i++ {
		y[i] = cmplx.Conj(x[i])
	}

	// forward fft
	y = fft.Transform(y)

	// take conjugate again
	for i := 0; i < n; i++ {
		y[i] = cmplx.Conj(y[i])
	}

	// divide by n
	cN := complex(float64(n), 0)
	for i := 0; i < n; i++ {
		y[i] = y[i] / cN
	}

	return y
}

// CConvolve calculates the circular convolution
func (fft *FFT) CConvolve(x, y []complex128) []complex128 {
	if len(x) != len(y) {
		panic("dimensions are different")
	}
	n := len(x)
	a := fft.Transform(x)
	b := fft.Transform(y)

	// multiply
	c := make([]complex128, n)
	for i := 0; i < n; i++ {
		c[i] = a[i] * b[i]
	}

	// inverse
	return fft.Inverse(c)
}

// Convolve returns the linear convolution
func (fft *FFT) Convolve(x, y []complex128) []complex128 {
	a := make([]complex128, 2*len(x))
	b := make([]complex128, 2*len(y))
	copy(a, x)
	copy(b, y)
	return fft.CConvolve(a, b)
}
