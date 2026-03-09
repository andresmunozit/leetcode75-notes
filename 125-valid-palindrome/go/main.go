package main

import (
	"regexp"
	"strings"
)

func isPalindrome(s string) bool {
	// border conditions
	if len(s) < 2 {
		return true
	}

	s = strings.ToLower(s)

	re := regexp.MustCompile("^[a-zA-Z0-9]$")

	l := 0
	r := len(s) - 1

	for l <= r {
		if l == r {
			return true
		}

		if !re.MatchString(string(s[l])) {
			l++
			continue
		}

		if !re.MatchString(string(s[r])) {
			r--
			continue
		}

		if s[l] != s[r] {
			return false
		}

		l++
		r--
	}

	return true
}

func main() {
	isPalindrome(".,")
}
