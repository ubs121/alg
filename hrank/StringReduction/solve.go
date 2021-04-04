// https://www.hackerrank.com/challenges/string-reduction
package main

import "fmt"

var (
	// solution store ( string -> min_length )
	mem map[string]int
)

/**
  Reduced minimum

  s - string to be reduced
  return - minimum possible length

rmin(s):
  1, len(s) == 1
  if s[0]<>s[1] then 1 else 2, len(s) == 2
  min(rmin(s[1:]), rmin(s[0:len(s)-1]),  len(s)=n
*/
func rmin(s []byte) int {
	fmt.Println(string(s))

	// check from the solution store
	if m, ok := mem[string(s)]; ok {
		return m
	}

	if len(s) == 1 {
		return 1
	}

	if len(s) == 2 {
		if s[0] != s[1] {
			return 1
		} else {
			return 2
		}
	}

	// len(s) > 2
	m := len(s) // minimum length

	// m < 2 бол дахин цааш шалгах шаардлагагүй, үүнээс цааш багасах боломжгүй
	for i := 0; i < len(s)-1 && m > 2; i++ {

		if s[i] != s[i+1] {

			// replace letters at position (i, i+1)
			d := byte('?')
			if s[i] == byte('a') && s[i+1] == byte('b') {
				d = byte('c')
			} else if s[i] == byte('a') && s[i+1] == byte('c') {
				d = byte('b')
			} else if s[i] == byte('b') && s[i+1] == byte('a') {
				d = byte('c')
			} else if s[i] == byte('b') && s[i+1] == byte('c') {
				d = byte('a')
			} else if s[i] == byte('c') && s[i+1] == byte('a') {
				d = byte('b')
			} else if s[i] == byte('c') && s[i+1] == byte('b') {
				d = byte('a')
			}

			var t []byte
			t = append(t, s[:i]...)
			t = append(t, d)
			t = append(t, s[i+2:]...)

			// concate strings and create a new copy
			rm := rmin(t)
			if m > rm {
				m = rm
			}
		}
	}

	// store solution
	mem[string(s)] = m

	return m
}

func main() {
	s := []byte("babcbbaabcbcbcbaabbccaacccbbbcaaacabbbbaaaccbcccacbbccaccbbaacaccbabcaaaacaccacbaacc")
	mem = make(map[string]int)
	fmt.Println(rmin(s))

}
