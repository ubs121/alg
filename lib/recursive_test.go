package alg

import (
	"fmt"
	"testing"
)

// generate all the strings of n bits
func printBinaryDigits(digits []byte, i int) {
	if i < 1 {
		fmt.Printf("%s\n", string(digits))
	} else {
		digits[i-1] = '0'
		printBinaryDigits(digits, i-1)
		digits[i-1] = '1'
		printBinaryDigits(digits, i-1)
	}
}

func TestPrintBinary(t *testing.T) {
	digits := []byte("0000")
	printBinaryDigits(digits, len(digits))
}
