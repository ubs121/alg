package strings

import (
	"fmt"
	"regexp"
	"testing"
)

// a function that uses regular expressions to determine a password's strength.
func checkPassword(pwd string) string {
	// valid: a UTF-8 encoded string consisting of a minimum of 6 and a maximum of 25 characters. The accepted range of characters is [U+0021, U+007A]
	// strong: at least 1 lowercase letter, 1 uppercase, 1 special, at least len of 10
	// medium: same as strong except it will not contain special characters and its length must be greater or equal to 8 characters
	// weak: All other valid passwords

	// a minimum of 6 and a maximum of 25 characters
	if len(pwd) < 6 || len(pwd) > 25 {
		return "invalid"
	}

	pwdBytes := []byte(pwd)

	// check if password contains an invalid character, The accepted range of characters is [U+0021, U+007A]
	invalid, err := regexp.Match(`[\x{0000}-\x{0020}\x{007B}-\x{FFFF}]`, pwdBytes)
	if invalid || err != nil {
		return "invalid"
	}

	uppercaseExists, _ := regexp.Match("[A-Z]+", pwdBytes)
	lowercaseExists, _ := regexp.Match("[a-z]+", pwdBytes)
	specialExists, _ := regexp.Match("[!\"#$%&'\\(\\)\\*\\+,-./\\:;<=>?@\\[\\\\\\]\\^\\_]", pwdBytes)

	if uppercaseExists && lowercaseExists {
		if specialExists && len(pwd) >= 10 {
			return "strong"
		} else if len(pwd) >= 8 {
			return "medium" // it will not contain special characters !!! this is conflict
		}
	}

	return "weak"
}

func hasSpecial(pwd []byte) bool {
	specialExists, _ := regexp.Match("[!\"#$%&'\\(\\)\\*\\+,-./\\:;<=>?@\\[\\\\\\]\\^\\_]", pwd)
	return specialExists
}

func TestSpecial(t *testing.T) {
	fmt.Printf("%v", hasSpecial([]byte("a")))
}

func TestPwd(t *testing.T) {
	testCases := map[string]string{
		"Nufu&YM21S": "strong",
		"iT*2spX*8":  "medium", // i think this is a wrong test case, because it has a special character
		"gZAGel":     "weak",
		"2N# 9k":     "invalid",
	}
	for tc, exp := range testCases {
		got := checkPassword(tc)
		if got != exp {
			t.Errorf("%s: exp %s, got %s\n", tc, exp, got)
		}
	}
}
