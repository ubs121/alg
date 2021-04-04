// Protein mix problem, also can be solved using XOR/CNOT matrix multiplication
// https://www.hackerrank.com/challenges/pmix/proble
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the pmix function below.
 */
func pmix(s string, k int) string {
	n := len(s)

	// IDEA1: bit 0 and 1 can be independent systems

	// prepare initial state
	state := make([]byte, n)
	for i := 0; i < n; i++ {
		state[i] = s[i] - 'A' // A=00, B=01, C=10, D=11
	}

	//statesDone := map[string]int{}
	for j := 0; j < k; j++ {
		c := state[0] // to prevent from overwrite
		for i := 0; i < n-1; i++ {
			state[i] = state[i] ^ state[i+1] // update
		}
		state[n-1] = state[n-1] ^ c

		// break if state is repeated
		// key := string(state)
		// if pos, exists := statesDone[key]; exists {

		// 	pattern := map[int][]byte{} // pos -> state
		// 	pattern[pos] = state

		// 	// collect other states after 'pos'
		// 	for st, t := range statesDone {
		// 		if pos < t {
		// 			pattern[t] = []byte(st)
		// 		}
		// 	}

		// 	rem := (k-1)%len(pattern) + pos
		// 	state = pattern[rem]

		// 	break
		// } else {
		// 	statesDone[key] = j
		// }
	}

	// build final string
	var sb strings.Builder
	for i := 0; i < n; i++ {
		sb.WriteByte('A' + state[i])
	}
	return sb.String()
}

func pmix1(s string, k int) string {
	n := len(s)

	// IDEA1: bit 0 and 1 can be independent systems

	// prepare initial state
	state := make([]byte, n)
	for i := 0; i < n; i += 64 {
		state[i] = s[i] - 'A' // A=00, B=01, C=10, D=11
	}

	for j := 0; j < k; j++ {
		c := state[0] // to prevent from overwrite
		for i := 0; i < n/64; i++ {
			state[i] = state[i] ^ state[i+1] // update
		}
		state[n-1] = state[n-1] ^ c
	}

	// build final string
	var sb strings.Builder
	for i := 0; i < n; i++ {
		sb.WriteByte('A' + state[i])
	}
	return sb.String()
}

func main() {
	inFile, err := os.Open("input13.txt")
	checkError(err)
	defer inFile.Close()

	reader := bufio.NewReaderSize(inFile, 1024*1024)
	nk := strings.Split(readLine(reader), " ")

	nTemp, err := strconv.ParseInt(nk[0], 10, 64)
	checkError(err)
	n := int(nTemp)

	kTemp, err := strconv.ParseInt(nk[1], 10, 64)
	checkError(err)
	k := int(kTemp)

	s := readLine(reader)
	if len(s) != n {
		panic(fmt.Errorf("wrong input size: exp %d, got %d", n, len(s)))
	}

	result := pmix1(s, k)
	fmt.Printf("%s\n", result)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
