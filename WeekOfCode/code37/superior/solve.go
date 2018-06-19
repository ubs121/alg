package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func maximumSuperiorCharacters(freq []int) int64 {
	return 0
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	tTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	t := int(tTemp)

	for tItr := 0; tItr < int(t); tItr++ {
		freqTemp := strings.Split(readLine(reader), " ")

		var freq []int

		for freqItr := 0; freqItr < 26; freqItr++ {
			freqItemTemp, err := strconv.ParseInt(freqTemp[freqItr], 10, 32)
			checkError(err)
			freqItem := int(freqItemTemp)
			freq = append(freq, freqItem)
		}

		result := maximumSuperiorCharacters(freq)

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
