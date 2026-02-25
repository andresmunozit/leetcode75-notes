package main

import (
	"testing"

	"github.com/stretchr/testify/require"

	h "github.com/andresmunozit/leetcode75-notes/helpers-go"
)

func Test_sortedArrayToBST(t *testing.T) {
	tcases := []struct {
		name     string
		nums     []int
		expected *h.TreeNode
	}{
		{
			name: "sorted array 1",
			nums: []int{-10, -3, 0, 5, 9},
			expected: &h.TreeNode{
				Val: 0,
				Left: &h.TreeNode{
					Val: -10,
					Right: &h.TreeNode{
						Val: -3,
					},
				},
				Right: &h.TreeNode{
					Val: 5,
					Right: &h.TreeNode{
						Val: 9,
					},
				},
			},
		},
		{
			name: "sorted array 2",
			nums: []int{1, 3},
			expected: &h.TreeNode{
				Val: 1,
				Right: &h.TreeNode{
					Val: 3,
				},
			},
		},
	}

	for _, tc := range tcases {
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(t, tc.expected, sortedArrayToBST(tc.nums))
		})
	}
}
