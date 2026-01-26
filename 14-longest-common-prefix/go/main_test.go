package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_longestCommonPrefix(t *testing.T) {
	tests := []struct {
		name     string
		strs     []string
		expected string
	}{
		{
			"there is common prefix",
			[]string{"flower", "flow", "flight"},
			"fl",
		},
		{
			"there is no common prefix",
			[]string{"dog", " racecar", " car"},
			"",
		},
		{
			"all the strings contain the prefix",
			[]string{"cat", "cat", "cat"},
			"cat",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.expected, longestCommonPrefix(tt.strs))
		})
	}
}
