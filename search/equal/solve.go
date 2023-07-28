// https://www.hackerrank.com/challenges/equal/problem
package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

/* IDEA:
1. giving all but one is same as taking from that one
2. sort it,
3. find min and subtract starting from highest numbers to reach a min
4. repeat it for 0, 1, 2 remainders
*/

// Complete the equal function below.
func equal(arr []int) int {
	sort.Ints(arr)

	min := math.MaxInt64
	for to := 0; to < 3; to++ {
		rounds := 0
		for i := 0; i < len(arr); i++ {
			diff := arr[i] - arr[0] + to
			if 0 < diff {
				// 5s
				rounds += diff / 5
				rem := diff % 5
				// 2s
				rounds += rem / 2
				rem = rem % 2
				// 1s
				rounds += rem
			}
		}

		if rounds < min {
			min = rounds
		}
	}

	return min
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	tTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	t := int32(tTemp)

	for tItr := 0; tItr < int(t); tItr++ {
		nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
		checkError(err)
		n := int32(nTemp)

		arrTemp := strings.Split(readLine(reader), " ")

		var arr []int

		for i := 0; i < int(n); i++ {
			arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
			checkError(err)
			arrItem := int(arrItemTemp)
			arr = append(arr, arrItem)
		}

		result := equal(arr)

		fmt.Printf("%d\n", result)
	}
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
