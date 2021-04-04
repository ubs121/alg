package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
	3*3 square

	a f d
	i e g
	b h c
*/
type square [3][3]int

// calculate minumum distance to a magic square
func solver(start *square) int {
	minCost := 81            // maximum cost, all changed by 9
	e := start.guessCentre() // possible center
	sum := start.sum()       // sum of all elements

	// possible centers
	centers := []int{e, e + 1, e - 1, e + 2, e - 2}

	// e - center
	for _, e := range centers {
		if e < 1 || e > 9 {
			continue
		}

		if minCost < abs(9*e-sum) {
			continue
		}

		m := 3 * e

		// check all possible a, b corners
		for a := max(1, m-e-9); a <= e; a++ {
			for b := max(1, m-a-9); b <= min(9, m-a); b++ {

				dist, err := start.distanceTo(e, a, b)

				if err == nil && dist < minCost {
					minCost = dist
				}
			}
		}

	}

	return minCost
}

// distance to a magic square with center 'e', and corner 'a', 'b'
func (sq *square) distanceTo(e, a, b int) (int, error) {

	m := 3 * e // magic number

	var tmp square

	tmp[0][0] = a
	tmp[1][1] = e
	tmp[2][0] = b
	tmp[0][2] = m - e - b
	tmp[2][2] = m - e - a
	tmp[1][0] = m - a - b
	tmp[0][1] = m - a - tmp[0][2]
	tmp[1][2] = m - tmp[2][2] - tmp[0][2]
	tmp[2][1] = m - b - tmp[2][2]

	dist := 0
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if tmp[i][j] < 1 || tmp[i][j] > 9 {
				return 100, errors.New("wrong input")
			}
			dist += abs(tmp[i][j] - sq[i][j])
		}
	}

	/*
		for i := 0; i < 3; i++ {
			fmt.Printf("%d %d %d\n", tmp[i][0], tmp[i][1], tmp[i][2])
		}
		fmt.Printf("dist=%d\n -------\n", dist)
	*/

	return dist, nil
}

// find possible center
func (sq *square) guessCentre() int {
	var avg []int

	// row, column averages
	for i := 0; i < 3; i++ {
		avg = append(avg, (sq[i][0]+sq[i][1]+sq[i][2])/3)
		avg = append(avg, (sq[0][i]+sq[1][i]+sq[2][i])/3)
	}

	// diagonal 1, 2
	avg = append(avg, (sq[0][0]+sq[1][1]+sq[2][2])/3)
	avg = append(avg, (sq[0][2]+sq[1][1]+sq[2][0])/3)

	s := 0
	for i := 0; i < len(avg); i++ {
		s += avg[i]
	}

	return s / len(avg)
}

// find sum
func (sq *square) sum() int {
	s := 0
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			s += sq[i][j]
		}
	}
	return s
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)
	writer := bufio.NewWriterSize(os.Stdout, 1024*1024)

	data := new(square)
	for i := 0; i < 3; i++ {
		sRowTemp := strings.Split(readLine(reader), " ")

		for j, sRowItem := range sRowTemp {
			sItemTemp, err := strconv.ParseInt(sRowItem, 10, 64)
			checkError(err)
			data[i][j] = int(sItemTemp)
		}

	}

	result := solver(data)
	fmt.Fprintf(writer, "%d\n", int(result))

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

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(n int) int {
	y := n >> 63
	return (n ^ y) - y
}
