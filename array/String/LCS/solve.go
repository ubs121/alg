// https://www.hackerrank.com/challenges/common-child
package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"io"
	"os"
	"strings"
)

// Complete the commonChild function below.
func commonChild(s1 string, s2 string) int {
	return LCS(s1, s2)
}

// Finding common longest substring
// https://en.wikipedia.org/wiki/Longest_common_subsequence_problem
func LCS(s1 string, s2 string) int {
	m := len(s1) + 1
	n := len(s2) + 1

	C := make([][]int, m)
	C[0] = make([]int, n)

	for i := 1; i < m; i++ {
		C[i] = make([]int, n) // row init

		for j := 1; j < n; j++ {
			if s1[i-1] == s2[j-1] {
				C[i][j] = C[i-1][j-1] + 1
			} else {
				C[i][j] = max(C[i][j-1], C[i-1][j])
			}
		}
	}

	return C[m-1][n-1]
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

// Solve using priority queue
func _solveByPQ(s1 string, s2 string) int {
	var s11, s22 []byte
	s1Cnt := countChars(s1)
	s2Cnt := countChars(s2)

	// cleanup s1
	for i := 0; i < len(s1); i++ {
		if _, exists := s2Cnt[s1[i]]; exists {
			s11 = append(s11, s1[i])
		}
	}

	// cleanup s2
	for i := 0; i < len(s2); i++ {
		if _, exists := s1Cnt[s2[i]]; exists {
			s22 = append(s22, s2[i])
		}
	}
	if len(s11) == 0 { //  || len(ref) == 0
		return 0 // no common characters
	}

	// TODO: take s11=min(s11, s22)

	max := 1
	ref := string(s22)

	pq := PriorityQueue{}
	heap.Init(&pq)
	pq.Push(&Item{value: string(s11)}) // push first string

	for 0 < pq.Len() {
		// pop item
		item := heap.Pop(&pq).(*Item)

		if max >= len(item.value) {
			break // no need to check beyond this point
		}

		//fmt.Printf("%s\n", item.value)

		rightMax := make([]int, len(item.value)) // rightMax[i] = number of chars found in right to left order
		ref2 := ref                              // reference string
		maxSoFar := 0
		for i := len(item.value) - 1; i >= 0; i-- {
			// nearest right position of item.value[i] in 'ref2'
			if hi := strings.LastIndexByte(ref2, item.value[i]); hi >= 0 {
				maxSoFar++       // it includes item.value[i]
				ref2 = ref2[:hi] // cut 'hi'
			}
			rightMax[i] = maxSoFar
		}

		// update max
		if max < rightMax[0] {
			max = rightMax[0]
		}

		ref1 := ref       // reference string
		leftMax := 0      // number of chars found in left to right order
		var lastChar byte // to avoid checking repeated chars

		// push neighbors (valid candidates)
		for i := 0; i < len(item.value); i++ {

			// nearest left position of item.value[i] in 'ref'
			if lo := strings.IndexByte(ref1, item.value[i]); lo >= 0 {
				leftMax++          // it includes item.value[i]
				ref1 = ref1[lo+1:] // cut 'lo'
			}

			// push removing i
			if item.value[i] != lastChar {
				if max < leftMax+rightMax[i] {
					var sb strings.Builder
					sb.WriteString(item.value[:i])
					sb.WriteString(item.value[i+1:])
					pq.Push(&Item{value: sb.String(), priority: leftMax + rightMax[i]})
				}

				lastChar = item.value[i]
			}
		}

	}

	return max
}

func countChars(s string) map[byte]int {
	ret := map[byte]int{}
	for i := 0; i < len(s); i++ {
		ret[s[i]]++
	}
	return ret
}

func main() {
	inFile, err := os.Open("input13.txt")
	checkError(err)
	defer inFile.Close()

	reader := bufio.NewReaderSize(inFile, 1024*1024)

	s1 := readLine(reader)
	s2 := readLine(reader)

	result := commonChild(s1, s2)

	fmt.Printf("%d\n", result)
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
