package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_singleNumber(t *testing.T) {
	tcases := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "single number at end",
			nums:     []int{2, 2, 1},
			expected: 1,
		},
		{
			name:     "single number in middle",
			nums:     []int{4, 1, 2, 1, 2},
			expected: 4,
		},
		{
			name:     "single element",
			nums:     []int{1},
			expected: 1,
		},
	}

	for _, tc := range tcases {
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(t, tc.expected, singleNumber(tc.nums))
		})
	}
}
