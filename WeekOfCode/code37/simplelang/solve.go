package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func maximumProgramValue(ins []string) int64 {
	var y int64
	var err error
	x := int64(0) // initial value

	for i := 0; i < len(ins); i++ {
		// parse value
		y, err = strconv.ParseInt(ins[i][4:], 10, 64)
		checkError(err)

		// check instruction
		if strings.HasPrefix(ins[i], "set") { // set
			if y > x {
				x = int64(y)
			}
		} else { // add
			if y > 0 {
				x = x + y
			}
		}
	}
	return x
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	// n
	nTemp, err := strconv.ParseInt(readLine(reader), 10, 32)
	checkError(err)
	n := int(nTemp)

	// read all instructions
	buf, err := ioutil.ReadAll(reader)
	checkError(err)
	ins := strings.SplitN(string(buf), "\n", n)

	// solve
	result := maximumProgramValue(ins)

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
