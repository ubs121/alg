package main

import "fmt"

type Student struct {
	id     int
	height int
	price  int
	target bool // madison
}

func (from *Student) costTo(to *Student) (cost int) {
	if to.target {
		return to.id - from.id
	}

	cost = to.id - from.id + to.price

	if from.height < to.height {
		cost += to.height - from.height
	} else {
		cost += from.height - to.height
	}
	return
}

// stringify
func (s *Student) String() string {
	return fmt.Sprintf("{%d, %d, %d}", s.id, s.height, s.price)
}
