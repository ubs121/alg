package main

import (
	"bytes"
	"fmt"
)

func main() {
	var buffer bytes.Buffer

	for i := 0; i < 1; i++ {
		buffer.WriteString("ulaanbaatar")
	}
	s := buffer.String()

	k := 3
	seg := make(map[byte]int)

	// count first segment
	for i := 0; i < k; i++ {
		if c, ex := seg[s[i]]; ex {
			seg[s[i]] = c + 1
		} else {
			seg[s[i]] = 1
		}

	}

	for i := 0; i+k < len(s); i++ {
		if len(seg) == k {
			fmt.Println(s[i : i+k])
		}

		// add last char
		if m, ex := seg[s[i+k]]; ex {
			seg[s[i+k]] = m + 1
		} else {
			seg[s[i+k]] = 1
		}

		// remove first char
		seg[s[i]] = seg[s[i]] - 1
		if seg[s[i]] == 0 {
			delete(seg, s[i]) // when 0
		}
	}
}
