// divide and conquer
// https://www.hackerrank.com/challenges/array-pairs/problem
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	inFile, err := os.Open("testCase13.txt")
	checkError(err)
	defer inFile.Close()

	rdr := bufio.NewReader(inFile) //os.Stdin

	// read 'n'
	line, _ := rdr.ReadString('\n')
	n, err := strconv.Atoi(strings.TrimRight(line, "\n\r"))
	checkError(err)

	// read array 'a'
	line, _ = rdr.ReadString('\n')                               // scan whole line
	aTemp := strings.Split(strings.TrimRight(line, "\r\n"), " ") // split by space

	arr := make([]int, n)
	for i := 0; i < len(aTemp); i++ {
		arr[i], _ = strconv.Atoi(aTemp[i])
	}

	// solve
	result := solveByMaxSplit(arr)

	fmt.Printf("%d\n", result)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
