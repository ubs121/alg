//https://www.hackerrank.com/challenges/array-pairs/problem
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the solve function below.
func solve(arr []int) int {
	n := len(arr)
	max := 1

	/* count elements */
	cntr := make([]int, 500001)
	for i := 0; i < n; i++ {
		cntr[arr[i]]++ // increase cntr

		if max < arr[i] {
			max = arr[i] // keep max
		}
	}
	for i := 1; i < n; i++ {
		cntr[i] += cntr[i-1] // accumulate
	}

	fmt.Printf("%v\n", cntr[:5])

	/* count pairs */
	pairs := 0
	cntr2 := make([]int, 500001)
	for i := 0; i < n; i++ {
		k := max / arr[i]
		pairs += cntr[k] - cntr2[arr[i]]
		cntr2[arr[i]]++
	}

	return pairs
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)
	writer := bufio.NewWriterSize(os.Stdout, 1024*1024)

	n, err := strconv.ParseInt(readLine(reader), 10, 32)
	checkError(err)

	arrTemp := strings.Split(readLine(reader), " ")

	var arr []int

	for i := 0; i < int(n); i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 32)
		checkError(err)
		arrItem := int(arrItemTemp)
		arr = append(arr, arrItem)
	}

	result := solve(arr)

	fmt.Fprintf(writer, "%d\n", result)

	writer.Flush()
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
