//https://www.hackerrank.com/challenges/pangrams

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var s string

	in := bufio.NewReader(os.Stdin)
	s, _ = in.ReadString('\n')

	abc := make(map[byte]bool)

	for i := 'a'; i <= 'z'; i++ {
		abc[byte(i)] = true
	}

	for _, c := range s {
		if 'A' <= c && c <= 'Z' { // capital letter
			delete(abc, byte(c)+'a'-'A')
		} else {
			delete(abc, byte(c))
		}
	}

	if len(abc) == 0 {
		fmt.Println("pangram")
	} else {
		fmt.Println("not pangram")
	}
}
