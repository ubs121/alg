// https://www.hackerrank.com/challenges/common-child
package main

import (
	"bufio"
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
