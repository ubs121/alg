// https://www.hackerrank.com/challenges/sherlock-and-valid-string/problem
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func isValid(s string) bool {

	// frequency for each letter, ex: {'a':2, 'b':2, 'c':2, 'd': 1}
	freq := map[byte]int{}
	for i := 0; i < len(s); i++ {
		freq[s[i]]++
	}

	// count of each freq, ex: { freq(2):3, freq(1): 1}
	freqCount := map[int]int{}
	for _, v := range freq {
		freqCount[v]++
	}

	if len(freqCount) == 1 {
		return true // string is already valid
	}

	if len(freqCount) > 2 {
		return false // not fixable, too many un-even items
	}

	if c, exists := freqCount[1]; exists && c == 1 {
		return true // only one odd item exists, so it can be just removed to make it valid
	}

	var kv [4]int
	i := 0
	// k - freq, v - count of 'k'
	for k, v := range freqCount {
		kv[i] = k
		kv[i+1] = v
		i += 2
	}
	// to simplify the decision {x:xc, y:yc}
	x, xc, y, yc := kv[0], kv[1], kv[2], kv[3]

	// case {(y+1):1, y:yc}, if we remove 'y+1' item it becomes {y:yc+1} which is valid string
	if xc == 1 && x-1 == y {
		return true
	}

	// case {x:cx, (xc+1):1}, similar as above
	if yc == 1 && y-1 == x {
		return true
	}

	return false
}

func main() {
	// inFile, err := os.Open("input13.txt")
	// checkError(err)
	// defer inFile.Close()

	line, err := bufio.NewReader(os.Stdin).ReadString('\n')
	checkError(err)

	if isValid(line) {
		fmt.Printf("YES\n")
	} else {
		fmt.Printf("NO\n")
	}
}

func checkError(err error) {
	if err != nil && err != io.EOF {
		panic(err)
	}
}
