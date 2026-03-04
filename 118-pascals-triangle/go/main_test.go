package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_generate(t *testing.T) {
	tcases := []struct {
		numRows  int
		expected [][]int
	}{
		{
			numRows:  5,
			expected: [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}},
		},
		{
			numRows:  1,
			expected: [][]int{{1}},
		},
		{
			numRows:  2,
			expected: [][]int{{1}, {1, 1}},
		},
		{
			numRows:  4,
			expected: [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}},
		},
	}

	for _, tc := range tcases {
		t.Run(fmt.Sprintf("%d rows", tc.numRows), func(t *testing.T) {
			require.Equal(t, tc.expected, generate(tc.numRows))
		})
	}
}
