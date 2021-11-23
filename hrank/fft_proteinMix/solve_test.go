package main

import (
	"strconv"
	"strings"
	"testing"
)

func TestSolve(t *testing.T) {
	testCases := map[string]string{
		"AAAAD,1":  "AAADD",
		"AAAAD,2":  "AADAD",
		"AAAAD,3":  "ADDDD",
		"AAAAD,4":  "DAAAD",
		"AAAAD,5":  "DAADA",
		"AAAAD,15": "DDDDA",
		"DAA,6":    "ADD",
	}

	for in, exp := range testCases {
		inputs := strings.Split(in, ",")
		k, _ := strconv.Atoi(inputs[1])
		got := pmix(inputs[0], k)
		if got != exp {
			t.Errorf("[%s] expected %s, but got %s", in, exp, got)
		}
	}
}
