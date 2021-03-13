// https://www.hackerrank.com/challenges/sherlock-and-valid-string/problem
package main

import "fmt"

func main() {
	var s string
	fmt.Scanf("%s", &s)

	c := make([]int, 26)

	// count letters
	for i := 0; i < len(s); i++ {
		c[s[i]-'a']++
	}

	// All characters in 's' have the same exact frequency (i.e., occur the same number of times).
	// Deleting exactly 1 character from  will result in all its characters having the same frequency

	// check bumpiness
	level := 0
	cutIndex := -1
	i := 0
	for ; i < len(c); i++ {
		// skip 0s
		if c[i] == 0 {
			continue
		}
		//fmt.Printf("%v, %d, %d\n", c, level, i)

		// set level
		if level == 0 {
			level = c[i] // use first value as a level, later we can adjust it
		}

		// continue if same level
		if c[i] == level {
			continue
		}

		// above level
		if c[i] > level {
			if c[i]-1 == level && cutIndex < 0 {
				c[i]--       // decrease/cut
				cutIndex = i // continue
			} else {
				break // already cut or too high
			}
		}

		// below level
		if c[i] < level {
			if c[i] == 1 { // when 1
				if cutIndex < 0 {
					c[i]--       // decrease/cut
					cutIndex = i // continue
				} else {
					break // already cut
				}
			} else { // other
				level = c[i] // adjust level !!!
				i = -1       // re-start
			}
		}

	} // i++ works here

	if i >= len(c) {
		fmt.Printf("YES\n")
	} else {
		fmt.Printf("NO\n")
	}

}
