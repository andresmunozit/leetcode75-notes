package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_lengthOfLastWord(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		expected int
	}{
		{
			name:     "success case",
			s:        "Hello World",
			expected: 5,
		},
		{
			name:     "string padded with spaces",
			s:        "   fly me   to   the moon  ",
			expected: 4,
		},
		{
			name:     "success case 2",
			s:        "luffy is still joyboy",
			expected: 6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.expected, lengthOfLastWord(tt.s))
		})
	}
}
