// https://www.hackerrank.com/challenges/queries-with-fixed-length/problem
package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

/* IDEA:
1. divide by max
2. calculate rank for each max element,
3. sort elements by rank
4. find min among the elements ranked higher than query 'd'

	Example: numbers in each rank
	rank1 [5,2,1,3,4,1,2,3] <-original array
	rank2 [5,2,3,4,4,2,3]
	rank3 [5,3,4,4,4,3]
	rank4 [5,4,4,4,4]
	rank5 [5,4,4,4]
	rank6 [5,4,4]
	rank7 [5,4]
	rank8 [5]
*/

// Complete the solve function below.
func solve(arr []int, queries []int) []int {

	// elem2rank[i] - rank of i'th element in the 'arr'
	elem2rank := make([]int, len(arr))
	for i := 0; i < len(elem2rank); i++ {
		elem2rank[i] = 1 // default rank
	}
	updateRanks(elem2rank, arr)

	// group by rank and update by actual/minimum values
	// must be rankMin[i] > 0
	// all rank[i] > q elements will be also in rank 'i'
	rankMin := make([]int, len(elem2rank)+1)
	for i, r := range elem2rank {
		if rankMin[r] != 0 { // already set some value
			if rankMin[r] > arr[i] {
				rankMin[r] = arr[i] // take min among same rank numbers
			}
		} else {
			rankMin[r] = arr[i] // new value
		}
	}

	// calculate minimums from right to left direction
	i := len(rankMin) - 1
	minSoFar := rankMin[i]
	for i >= 0 {
		if rankMin[i] == 0 { // missing rank
			rankMin[i] = minSoFar
		} else {
			// ?????
			if minSoFar > rankMin[i] {
				minSoFar = rankMin[i]
			}
		}
		i--
	}

	// collect answer for each query
	var ret []int
	for _, q := range queries {
		ret = append(ret, rankMin[q])
	}
	return ret
}

// find max element and update ranks in place, repeat it for segment split by max
// element rank is a segment length that the element to be the maxiumum value in that segment
func updateRanks(rank []int, arr []int) {
	if len(arr) < 2 {
		return
	}

	ixs := IndexMax(arr) // indexes of the maximum element
	i := ixs[len(ixs)/2] // take median
	rank[i] = len(arr)   // in-place update, it should be updating the original table

	// recursively find maximums in the left and right sides
	updateRanks(rank[:i], arr[:i])
	updateRanks(rank[i+1:], arr[i+1:])
}

// Lib: IndexMax returns max value indexes
func IndexMax(a []int) []int {
	var maxes []int
	maxVal := math.MinInt64 // default max

	for i := 0; i < len(a); i++ {
		if maxVal < a[i] {
			maxVal = a[i]    // new max
			maxes = []int{i} // another max found, so reset
		} else if maxVal == a[i] {
			// collect max positions
			maxes = append(maxes, i)
		}
	}

	return maxes
}

func main() {
	inFile, err := os.Open("input7.txt")
	checkError(err)
	defer inFile.Close()

	reader := bufio.NewReaderSize(inFile, 1024*1024)

	nq := strings.Split(readLine(reader), " ")

	nTemp, err := strconv.ParseInt(nq[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	qTemp, err := strconv.ParseInt(nq[1], 10, 64)
	checkError(err)
	q := int32(qTemp)

	arrTemp := strings.Split(readLine(reader), " ")

	var arr []int

	for arrItr := 0; arrItr < int(n); arrItr++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[arrItr], 10, 64)
		checkError(err)
		arrItem := int(arrItemTemp)
		arr = append(arr, arrItem)
	}

	var queries []int

	for queriesItr := 0; queriesItr < int(q); queriesItr++ {
		queriesItemTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
		checkError(err)
		queriesItem := int(queriesItemTemp)
		queries = append(queries, queriesItem)
	}

	result := solve(arr, queries)

	for resultItr, resultItem := range result {
		fmt.Printf("%d", resultItem)

		if resultItr != len(result)-1 {
			fmt.Printf("\n")
		}
	}

	fmt.Printf("\n")
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
