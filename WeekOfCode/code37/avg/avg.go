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

func round(x float64) float64 {
	return math.Floor(x*100+.5) / 100.0
}

// Complete the averageOfTopEmployees function below.
func averageOfTopEmployees(rating []int32) {
	var s int32
	k := 0
	for i := 0; i < len(rating); i++ {
		if rating[i] >= 90 {
			s += rating[i]
			k++
		}
	}

	fmt.Printf("%.2f\n", round(float64(s)/float64(k)))

}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	var rating []int32

	for ratingItr := 0; ratingItr < int(n); ratingItr++ {
		ratingItemTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
		checkError(err)
		ratingItem := int32(ratingItemTemp)
		rating = append(rating, ratingItem)
	}

	averageOfTopEmployees(rating)
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
