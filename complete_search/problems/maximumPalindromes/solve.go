// https://www.hackerrank.com/challenges/maximum-palindromes/problem
package main

import (
	"bufio"
	"fmt"
	"io"
	"math/big"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'initialize' function below.
 *
 * The function accepts STRING s as parameter.
 */

type inData struct {
	txt  string
	freq [][26]int // freq[i] - letter frequencies in range txt[0:i]
	fact map[int]uint64
}

var inp *inData

func initialize(s string) {
	// This function is called once before all queries.
	inp = &inData{}
	inp.txt = s

	// count letters
	freq := make([][26]int, len(s)+1)
	cnt := make([]int, 26)
	for i := 0; i < len(s); i++ {
		c := s[i] - 'a'
		cnt[c]++
		copy(freq[i][:], cnt)
	}
	inp.freq = freq

	// calculate all factorials
	inp.fact = make(map[int]uint64)
	inp.fact[0] = 1
	p := uint64(1)
	for i := 1; i <= len(s)/2; i++ {
		p = (p * uint64(i)) % 1000000007
		inp.fact[i] = p
	}

	// TODO: calculate all multiplicative inverses

}

/*
 * Complete the 'answerQuery' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER l
 *  2. INTEGER r
 */

func answerQuery(l int, r int) int {
	// Return the answer for this query modulo 1000000007.

	// get letter frequencies in range [l-1:r]
	freq := inp.freq[r-1] // range [0:r-1]
	if l > 1 {            // freq[0] is all zero
		cntL := inp.freq[l-2]
		for c := 0; c < 26; c++ {
			freq[c] = freq[c] - cntL[c]
		}
	}

	// IDEA: half the string and calculate permutations of multisets
	perms := uint64(1)
	n := 0
	odds := 0
	for _, v := range freq {
		if v == 0 {
			continue
		}

		if v%2 == 0 {
			n += v / 2
			if p, exists := inp.fact[v/2]; exists {
				perms = (perms * p) % 1000000007
			} else {
				panic(fmt.Errorf("%d factorial not found", v/2))
			}
		} else {
			odds++
			n += (v - 1) / 2 // make it even
			if p, exists := inp.fact[(v-1)/2]; exists {
				perms = (perms * p) % 1000000007
			} else {
				panic(fmt.Errorf("%d factorial not found", (v-1)/2))
			}
		}
	}

	// ans:=(n!/perms)%1000000007 = ((n! % 1000000007)*(inverse(perms)%1000000007)) % 1000000007
	m := new(big.Int).SetUint64(1000000007)
	ans := new(big.Int).SetUint64(inp.fact[n])
	y := new(big.Int).SetUint64(perms)
	y.ModInverse(y, m)
	ans.Mul(ans, y)
	ans.Mod(ans, m)

	if odds > 0 {
		o := new(big.Int).SetInt64(int64(odds))
		m = new(big.Int).SetInt64(1000000007)
		ans.Mul(ans, o) // ans*=odd, always can put one char in the middle
		ans.Mod(ans, m) // ans%=1000000007
	}

	return int(ans.Int64())
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	s := readLine(reader)

	initialize(s)

	qTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	q := int32(qTemp)

	for qItr := 0; qItr < int(q); qItr++ {
		firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

		lTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
		checkError(err)
		l := int(lTemp)

		rTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
		checkError(err)
		r := int(rTemp)

		result := answerQuery(l, r)

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
