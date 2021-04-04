package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func whichSection(n int, k int, a []int) int {
	i := 0 // section id
	s := 0 // number of students so far
	for k > s {
		s = s + a[i]
		i++
	}
	return i
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)
	writer := bufio.NewWriterSize(os.Stdout, 1024*1024)

	tTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	t := int(tTemp)

	for tItr := 0; tItr < int(t); tItr++ {
		nkm := strings.Split(readLine(reader), " ")

		nTemp, err := strconv.ParseInt(nkm[0], 10, 64)
		checkError(err)
		n := int(nTemp)

		kTemp, err := strconv.ParseInt(nkm[1], 10, 64)
		checkError(err)
		k := int(kTemp)

		mTemp, err := strconv.ParseInt(nkm[2], 10, 64)
		checkError(err)
		m := int(mTemp)

		aTemp := strings.Split(readLine(reader), " ")

		var a []int

		for i := 0; i < int(m); i++ {
			aItemTemp, err := strconv.ParseInt(aTemp[i], 10, 64)
			checkError(err)
			aItem := int(aItemTemp)
			a = append(a, aItem)
		}

		result := whichSection(n, k, a)

		fmt.Fprintf(writer, "%d\n", result)
	}

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
