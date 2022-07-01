package numbers

import (
	"math"
	"math/cmplx"
)

// Fast Fourier Transformer
type FFT struct {
}

// Transform returns the FFT of 'vec'
func (fft *FFT) Transform(vec []complex128) []complex128 {
	n := len(vec)

	if n == 1 {
		return []complex128{vec[0]}
	}

	if n%2 != 0 {
		panic("len(x) is not a power of 2")
	}

	// even terms
	even := make([]complex128, n/2)
	for i := 0; i < n/2; i++ {
		even[i] = vec[2*i]
	}
	even1 := fft.Transform(even)

	// odd terms
	odd := even // re-using the allocated memory
	for i := 0; i < n/2; i++ {
		odd[i] = vec[2*i+1]
	}
	odd1 := fft.Transform(odd)

	// combine evens and odds
	ret := make([]complex128, n)
	for k := 0; k < n/2; k++ {
		sin, cos := math.Sincos(-2 * math.Pi * float64(k) / float64(n)) // primitive nth root of unity
		wk := complex(cos, sin)                                         // euler's formula
		ret[k] = even1[k] + wk*odd1[k]
		ret[k+n/2] = even1[k] - wk*odd1[k]
	}
	return ret
}

// Inverse returns the inverse FFT of 'vec'
// The same FFT but exp(-2*math.Pi*i*k/n) and divide by n
func (fft *FFT) Inverse(vec []complex128) []complex128 {
	n := len(vec)

	// conjugate
	ret := make([]complex128, n)
	for i := 0; i < n; i++ {
		ret[i] = cmplx.Conj(vec[i])
	}

	// forward fft
	ret = fft.Transform(ret)

	// take conjugate again
	for i := 0; i < n; i++ {
		ret[i] = cmplx.Conj(ret[i])
	}

	// divide by n
	cN := complex(float64(n), 0)
	for i := 0; i < n; i++ {
		ret[i] = ret[i] / cN
	}

	return ret
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

// https://en.wikipedia.org/wiki/Cooley%E2%80%93Tukey_FFT_algorithm
func ditfft2(x []float64, y []complex128, n, s int) {
	if n == 1 {
		y[0] = complex(x[0], 0)
		return
	}
	ditfft2(x, y, n/2, 2*s)
	ditfft2(x[s:], y[n/2:], n/2, 2*s)
	for k := 0; k < n/2; k++ {
		tf := cmplx.Rect(1, -2*math.Pi*float64(k)/float64(n)) * y[k+n/2]
		y[k], y[k+n/2] = y[k]+tf, y[k]-tf
	}
}
