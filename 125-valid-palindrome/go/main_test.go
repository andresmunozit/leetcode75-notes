package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_isPalindrome(t *testing.T) {
	tcases := []struct {
		name     string
		s        string
		expected bool
	}{
		{
			name:     "palindrome with punctuation and spaces",
			s:        "A man, a plan, a canal: Panama",
			expected: true,
		},
		{
			name:     "not a palindrome",
			s:        "race a car",
			expected: false,
		},
		{
			name:     "single space is palindrome",
			s:        " ",
			expected: true,
		},
		{
			name:     "two characters, not a palindrome 1",
			s:        "a.",
			expected: true,
		},
		{
			name:     "two characters, not a palindrome 2",
			s:        ".,",
			expected: true,
		},
	}

	for _, tc := range tcases {
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(t, tc.expected, isPalindrome(tc.s))
		})
	}
}
