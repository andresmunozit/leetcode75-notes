package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_strStr(t *testing.T) {
	tests := []struct {
		name     string
		needle   string
		haystack string
		expected int
	}{
		{
			name:     "needle found",
			needle:   "sad",
			haystack: "sadbutsad",
			expected: 0,
		},
		{
			name:     "needle not found",
			needle:   "leeto",
			haystack: "leetcode",
			expected: -1,
		},
		{
			name:     "needle equals to haystack",
			needle:   "a",
			haystack: "a",
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.expected, strStr(tt.haystack, tt.needle))
		})
	}
}
