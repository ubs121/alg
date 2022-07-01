//https://codingcompetitions.withgoogle.com/kickstart/round/00000000008cb33e/00000000009e7021#problem

package main

import "fmt"

func main() {
	// read input
	var T int
	fmt.Scanf("%d", &T)

	for t := 0; t < T; t++ {
		var str1, str2 string
		fmt.Scanf("%s", &str1)
		fmt.Scanf("%s", &str2)

		// solve all test cases
		ans := solve(str1, str2)
		if ans >= 0 {
			fmt.Printf("Case #%d: %d\n", (t + 1), ans)
		} else {
			fmt.Printf("Case #%d: IMPOSSIBLE\n", (t + 1))
		}
	}
}

func solve(str1, str2 string) int {
	// impossible if no sufficient letters in 'str2'
	if len(str2) < len(str1) {
		return -1 // impossible
	}

	dels := 0    // minimum deletions
	i, j := 0, 0 // current letters in 'str1', 'str2'
	for i < len(str1) && j < len(str2) {

		// skip if matches
		if str1[i] == str2[j] {
			j++
			i++
		} else {
			// delete non-matching letters
			for j < len(str2) && str1[i] != str2[j] {
				dels++
				j++
			}
		}
	}

	if i < len(str1) && j >= len(str2) {
		return -1 // impossible
	}

	if i >= len(str1) && j < len(str2) {
		dels += len(str2) - j
	}

	return dels
}
