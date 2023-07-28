package strings

import (
	"regexp"
	"testing"
)

func TestNumRegex(t *testing.T) {
	validCases := []string{"2", "0089", "-0.1", "+3.14", "4.", "-.9", "2e10", "-90E3", "3e+7", "+6e-1", "53.5e93", "-123.456e789"}
	invalidCases := []string{"abc", "1a", "1e", "e3", "99e2.5", "--6", "-+3", "95a54e53"}

	var validNumber = regexp.MustCompile(`^(\+|-)?(\d+|(\d+\.\d*)|(\d*\.\d+))((e|E)(\+|-)?\d+)?$`)
	//

	for _, tc := range validCases {
		got := validNumber.MatchString(tc)
		if got != true {
			t.Errorf("%s: exp: true, got: %v", tc, got)
		}
	}

	for _, tc := range invalidCases {
		got := validNumber.MatchString(tc)
		if got != false {
			t.Errorf("%s: exp: false, got: %v", tc, got)
		}
	}
}
