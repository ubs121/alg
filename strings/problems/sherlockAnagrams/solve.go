// https://www.hackerrank.com/challenges/sherlock-and-anagrams/problem
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the sherlockAndAnagrams function below.
func sherlockAndAnagrams(s string) int {
	// 2<=n<=100
	n := len(s)

	ans := 0          // final answer
	lsum := [26]int{} // left sum
	l := 1            // len index
	for l < n {
		lsum[s[l-1]-'a']++

		// substr frequencies with length 'l'
		strFreq := map[string]int{}
		strFreq[hashSubstr(lsum)]++

		var csum [26]int
		for i := 0; i < 26; i++ {
			csum[i] = lsum[i] // copy
		}

		for i := 1; i < n-l+1; i++ {
			// signature of substr s[i:i+l]
			csum[s[i-1]-'a']--   // char out
			csum[s[i+l-1]-'a']++ // char in

			// count anagrams
			csumHash := hashSubstr(csum)
			if c, exists := strFreq[csumHash]; exists {
				//fmt.Printf("%s - %d\n", s[i:i+l], csum)
				ans += c
			}

			strFreq[csumHash]++
		}
		l++
	}

	return ans
}

func hashSubstr(buf [26]int) string {
	var sb strings.Builder
	for i := 0; i < 26; i++ {
		sb.WriteByte(byte('a' + i))
		sb.WriteString(strconv.Itoa(buf[i]))
	}
	return sb.String()
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	qTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	q := int32(qTemp)

	for qItr := 0; qItr < int(q); qItr++ {
		s := readLine(reader)

		result := sherlockAndAnagrams(s)

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
