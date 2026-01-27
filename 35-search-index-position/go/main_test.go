package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_searchInsert(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		target   int
		expected int
	}{
		{
			name:     "nums length is zero",
			nums:     []int{},
			target:   5,
			expected: 0,
		},
		{
			name:     "target is less than min",
			nums:     []int{1, 3, 5, 6},
			target:   0,
			expected: 0,
		},
		{
			name:     "target is greater than max",
			nums:     []int{1, 3, 5, 6},
			target:   7,
			expected: 4,
		},
		{
			name:     "target is greater than max by two",
			nums:     []int{2, 3, 4, 7, 8, 9},
			target:   11,
			expected: 6,
		},
		{
			name:     "target is present",
			nums:     []int{1, 3, 5, 6},
			target:   5,
			expected: 2,
		},
		{
			name:     "target is between two nums",
			nums:     []int{1, 3, 5, 6},
			target:   2,
			expected: 1,
		},
		{
			name:     "nums length is two",
			nums:     []int{1, 3},
			target:   2,
			expected: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.expected, searchInsert(tt.nums, tt.target))
		})
	}
}
