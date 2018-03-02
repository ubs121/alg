package main

// read integer sequence
func readInts(buf []byte, n int) []int {
	l := len(buf)
	if buf[l-1] == '\n' {
		l-- // remove '\n'
	}

	p := 0       // index for 'arr'
	i := 0       // index for 'buf'
	d := 0       // temp number
	neg := false // negative sign

	arr := make([]int, n)

	for i < l && p < n {
		for i < l && ' ' == buf[i] {
			i++
		}

		d = 0
		neg = false
		if buf[i] == '-' {
			neg = true
			i++
		}

		for i < l && '0' <= buf[i] && buf[i] <= '9' {
			d = d*10 + int(buf[i]-'0')
			i++
		}
		if neg {
			d = -1 * d
		}
		arr[p] = d
		p++
	}

	return arr
}
