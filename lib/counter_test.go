package alg

// a generic counter type
type counter struct {
	digits []digit
}

func (cntr *counter) next() {
	// generate next variant
	i := 0
	for i < len(cntr.digits) && cntr.digits[i].isLast() {
		cntr.digits[i].reset()
		i++
	}

	if i == len(cntr.digits) {
		return // reached the end
	}

	// plus one
	cntr.digits[i].inc()
}

// digit placeholder
type digit struct {
	all  []string // all possible values
	curr int      // current value
}

func (d *digit) isLast() bool {
	return d.curr == len(d.all)-1
}

func (d *digit) inc() {
	d.curr++
}

func (d *digit) reset() {
	d.curr = 0
}

func (d *digit) current() string {
	return d.all[d.curr]
}
