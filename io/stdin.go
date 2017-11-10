// read an integer from stdin
func readInt(reader *bufio.Reader) (int, error) {
	buf, err := reader.ReadBytes('\n')
	if err != nil && err != io.EOF {
		return 0, err
	}

	l := len(buf) - 1
	n := 0
	for i := 0; i < l; i++ { // exclude '\n'
		n = n*10 + int(buf[i]-'0')
	}
	return n, nil
}

// read integers from stdin,  'a' - output array
func readIntegers(reader *bufio.Reader, a []int) error {
	buf, err := reader.ReadBytes('\n')
	if err != nil && err != io.EOF {
		return err
	}

	l := len(buf) - 1
	p := 0
	i := 0
	for i < l && p < len(a) {
		for i < l && ' ' == buf[i] {
			i++
		}
		d := 0
		for i < len(buf) && ' ' != buf[i] {
			d = d*10 + int(buf[i]-'0')
			i++
		}
		a[p] = d
		p++
	}

	return err
}
