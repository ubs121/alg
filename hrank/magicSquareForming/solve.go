// https://www.hackerrank.com/challenges/magic-square-forming/problem
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
IDEA: One possible magic square is 8.
So we can check all its permutations including rotated versions.
*/

var allMagicSquares = [][]int{
	{8, 1, 6, 3, 5, 7, 4, 9, 2},
	{8, 3, 4, 1, 5, 9, 6, 7, 2},
	{6, 7, 2, 1, 5, 9, 8, 3, 4},
	{6, 1, 8, 7, 5, 3, 2, 9, 4},
	{4, 9, 2, 3, 5, 7, 8, 1, 6},
	{4, 3, 8, 9, 5, 1, 2, 7, 6},
	{2, 9, 4, 7, 5, 3, 6, 1, 8},
	{2, 7, 6, 9, 5, 1, 4, 3, 8},
}

// Complete the formingMagicSquare function below.
func formingMagicSquare(s [][]int) int {
	// must be sum(row)=sum(col)=15

	// helper function
	cost := func(v []int) int {
		c := 0
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				k := i*3 + j
				// cost calculation
				if s[i][j] < v[k] {
					c += v[k] - s[i][j]
				} else {
					c += s[i][j] - v[k]
				}
			}
		}
		return c
	}

	minCost := 81 // 9*9 cost
	for _, v := range allMagicSquares {
		costV := cost(v)
		if minCost > costV {
			minCost = costV
		}
	}

	return minCost
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	var s [][]int
	for i := 0; i < 3; i++ {
		sRowTemp := strings.Split(readLine(reader), " ")

		var sRow []int
		for _, sRowItem := range sRowTemp {
			sItemTemp, err := strconv.ParseInt(sRowItem, 10, 64)
			checkError(err)
			sItem := int(sItemTemp)
			sRow = append(sRow, sItem)
		}

		if len(sRow) != 3 {
			panic("Bad input")
		}

		s = append(s, sRow)
	}

	result := formingMagicSquare(s)
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
