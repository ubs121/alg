// Protein mix problem, also can be solved using XOR/CNOT matrix multiplication
// https://www.hackerrank.com/challenges/pmix/proble
package main

import (
	"numbers"
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/* IDEA: Encode: A=00, B=01, C=10, D=11
* 	1. bit 0 and 1 can be independent systems
	2. circular matrix mutiplication using FFT should give the answer
*/

/*
 * Complete the pmix function below.
 */
func pmix(s string, k int) string {
	n := len(s)

	// prepare initial state
	bit0 := make([]complex128, n)
	bit1 := make([]complex128, n)
	for i := 0; i < n; i++ {
		v := (s[i] - 'A')
		bit0[i] = complex(float64(v&1), 0)
		bit1[i] = complex(float64((v&0x10)>>1), 0)
	}

	fft := &numbers.FFT{}
	fft.Transform(bit0)

	// build final string
	var sb strings.Builder
	for i := 0; i < n; i++ {
		v := int(real(bit0[i])) & 1
		v += (int(real(bit0[i])) & 0x10) << 1
		sb.WriteByte('A' + byte(v))
	}
	return sb.String()
}

func main() {
	inFile, err := os.Open("input13.txt")
	checkError(err)
	defer inFile.Close()

	reader := bufio.NewReaderSize(inFile, 1024*1024)
	nk := strings.Split(readLine(reader), " ")

	nTemp, err := strconv.ParseInt(nk[0], 10, 64)
	checkError(err)
	n := int(nTemp)

	kTemp, err := strconv.ParseInt(nk[1], 10, 64)
	checkError(err)
	k := int(kTemp)

	s := readLine(reader)
	if len(s) != n {
		panic(fmt.Errorf("wrong input size: exp %d, got %d", n, len(s)))
	}

	result := pmix(s, k)
	fmt.Printf("%s\n", result)
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
