package main

import (
	"testing"
)

func TestSolution(t *testing.T) {
	commands := map[string]string{
		"GLL":    "YES",
		"RG":     "YES",
		"GR":     "YES",
		"RGGRGG": "YES", // will be cycle after 8 repeats
		"GGGGG":  "NO",
	}

	for cmd, expected := range commands {
		ret := "NO"
		if isCycle(cmd) {
			ret = "YES"
		}
		if ret != expected {
			t.Errorf("cmd=%s: expected %s, got %s", cmd, expected, ret)
		}
	}
}
