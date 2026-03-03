package main

import (
	"testing"

	h "github.com/andresmunozit/leetcode75-notes/helpers-go"
	"github.com/stretchr/testify/require"
)

func Test_hasPathSum(t *testing.T) {
	tcases := []struct {
		name      string
		root      *h.TreeNode
		targetSum int
		expected  bool
	}{
		{
			name: "case 1",
			root: &h.TreeNode{
				Val: 5,
				Left: &h.TreeNode{
					Val: 4,
					Left: &h.TreeNode{
						Val:   11,
						Left:  &h.TreeNode{Val: 7},
						Right: &h.TreeNode{Val: 2},
					},
				},
				Right: &h.TreeNode{
					Val:  8,
					Left: &h.TreeNode{Val: 13},
					Right: &h.TreeNode{
						Val:   4,
						Right: &h.TreeNode{Val: 1},
					},
				},
			},
			targetSum: 22,
			expected:  true,
		},
		{
			name: "case 2",
			root: &h.TreeNode{
				Val:   1,
				Left:  &h.TreeNode{Val: 2},
				Right: &h.TreeNode{Val: 3},
			},
			targetSum: 5,
			expected:  false,
		},
		{
			name:      "case 3",
			root:      nil,
			targetSum: 0,
			expected:  false,
		},
		{
			name: "case 4",
			root: &h.TreeNode{
				Val:  1,
				Left: &h.TreeNode{Val: 2},
			},
			targetSum: 1,
			expected:  false,
		},
		{
			name:      "case 5",
			root:      &h.TreeNode{Val: 1},
			targetSum: 1,
			expected:  true,
		},
	}

	for _, tc := range tcases {
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(t, tc.expected, hasPathSum(tc.root, tc.targetSum))
		})
	}
}
