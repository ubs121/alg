// https://codingcompetitions.withgoogle.com/kickstart/round/00000000008cb33e/00000000009e7997#problem

package main

import "fmt"

func main() {
	var T int
	fmt.Scanf("%d", &T)

	for t := 0; t < T; t++ {
		var n int
		fmt.Scanf("%d", &n)

		// solve all test cases
		fmt.Printf("Case #%d: %d\n", (t + 1), solve(n))
	}
}

// TODO: handle big numbers 10^123456
func solve(n int) int {
	m := n           // to save 'n'
	s := 0           // sum of digits
	var digits []int // all digits
	for m > 0 {
		d := m % 10
		digits = append([]int{d}, digits...)
		s += d
		m = m / 10
	}

	newDigit := 9 - (s % 9)
	if newDigit == 9 {
		// already multiple of 9, so we can add 0 or 9, but 0 is preferable
		newDigit = 0
	}

	// find a position for 'newDigit'
	if newDigit == 0 {
		// put right after the leading digit
		digits = append(digits[:1], append([]int{newDigit}, digits[1:]...)...)
	} else {
		i := 0
		for i < len(digits) && digits[i] <= newDigit {
			i++
		}
		if i == len(digits) {
			// put at the end
			digits = append(digits, newDigit)
		} else {
			// put before digits[i]
			digits = append(digits[:i], append([]int{newDigit}, digits[i:]...)...)
		}
	}

	// build back the answer
	n = 0
	for i := 0; i < len(digits); i++ {
		n = n*10 + digits[i]
	}
	return n
}
