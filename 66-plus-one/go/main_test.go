package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_plusOne(t *testing.T) {
	tests := []struct {
		name     string
		digits   []int
		expected []int
	}{
		{
			name:     "add one, no carry 1",
			digits:   []int{1, 2, 3},
			expected: []int{1, 2, 4},
		},
		{
			name:     "add one, no carry 2",
			digits:   []int{1, 2, 3},
			expected: []int{1, 2, 4},
		},
		{
			name:     "add one with carry, digits length equals to 1",
			digits:   []int{9},
			expected: []int{1, 0},
		},
		{
			name:     "add one with carry, new digit at the end, digits length greater than 1",
			digits:   []int{9, 9, 9},
			expected: []int{1, 0, 0, 0},
		},
		{
			name:     "add one with carry, digits length greater than 1",
			digits:   []int{1, 5, 9},
			expected: []int{1, 6, 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.expected, plusOne(tt.digits))
		})
	}
}
