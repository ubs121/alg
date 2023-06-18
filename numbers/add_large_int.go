// You are given two arbitrarily large numbers,
// stored one digit at a time in a slice.
// The first must be added to the second,
// and the second must be reversed before addition.
//
// The goal is to calculate the sum of those two sets of values.
//
// IMPORTANT NOTE:
// - The input can be any lengths (i.e: it can be 20+ digits long).
// - num1 and num2 can be different lengths.
//
// Sample Inputs:
// num1 = 123456
// num2 = 123456
//
// Sample Output:
// Result: 777777
//
// We would also like to see a demonstration of appropriate unit tests
// for this functionality.

package numbers

import (
	"bytes"
)

// NOTE: It assumes both numbers are non negative
// NOTE: "math/big" package is good enough in most cases
func Add(num1 []int, num2 []int) string {
	if len(num2) == 0 {
		return intArray2str(num1)
	}

	if len(num1) == 0 {
		return intArray2str(num2)
	}

	// reverse the first number before addition,
	// because reversing the first is easier for indexing and addition etc
	num1rev := Reverse(num1)

	// check the lengths of arrays
	var shortArr, longArr []int
	if len(num1rev) < len(num2) {
		longArr = num2
		shortArr = num1rev
	} else {
		longArr = num1rev
		shortArr = num2
	}

	var sumRev []int // sum of two numbers, digits are in reversed order
	carry := 0       // carry
	i := 0
	for i < len(shortArr) {
		s := shortArr[i] + longArr[i] + carry
		sumRev = append(sumRev, s%10)
		carry = s / 10
		i++
	}

	// add the carry to the longest
	for i < len(longArr) {
		s := longArr[i] + carry
		sumRev = append(sumRev, s%10)
		carry = s / 10
		i++
	}

	// add the carry to the sum
	if carry > 0 {
		sumRev = append(sumRev, carry)
	}

	// reverse back
	s := Reverse(sumRev)
	return intArray2str(s)
}

// converts from a string formatted number to an intger array
func str2intArray(strNum string) []int {
	out := make([]int, len(strNum))
	for i := 0; i < len(strNum); i++ {
		out[i] = int(strNum[i] - '0')
	}
	return out
}

// converts from an integer array to a string formatted number
func intArray2str(num []int) string {
	var out bytes.Buffer
	for i := 0; i < len(num); i++ {
		out.WriteByte(byte(num[i] + '0'))
	}
	return out.String()
}

// Reverse array of integers
func Reverse(num []int) []int {
	n := len(num)
	revOut := make([]int, n) // reversed output
	copy(revOut, num)        // copy original numbers

	for i := 0; i < n/2; i++ {
		j := n - i - 1                              // an index to be swapped
		revOut[i], revOut[j] = revOut[j], revOut[i] // swap
	}

	return revOut
}

// Reverses a large array in batches concurrently
// NOTE: it can be used in case reversing really large numbers
func ReverseInBatches(num []int, batch int) []int {
	n := len(num)
	var aggChan []chan []int // output channels

	off := 0 // batch offset
	for off+batch < n {
		chOut := make(chan []int)
		go func(in []int, ch chan []int) {
			chOut <- Reverse(in)
		}(num[off:off+batch], chOut)

		aggChan = append(aggChan, chOut)
		off += batch
	}

	// reverse the last remaining part
	if off < n {
		chLast := make(chan []int)
		go func(in []int) {
			chLast <- Reverse(in)
		}(num[off:n])

		aggChan = append(aggChan, chLast)
	}

	// join the results
	var output []int
	for i := len(aggChan) - 1; i >= 0; i-- {
		revArr := <-aggChan[i]
		close(aggChan[i])
		output = append(output, revArr...)
	}

	return output
}
