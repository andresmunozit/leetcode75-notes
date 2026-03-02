package main

import (
	"testing"

	h "github.com/andresmunozit/leetcode75-notes/helpers-go"
	"github.com/stretchr/testify/require"
)

func Test_isBalanced(t *testing.T) {
	tcases := []struct {
		name     string
		root     *h.TreeNode
		expected bool
	}{
		{
			name: "balanced tree 1",
			root: &h.TreeNode{
				Val: 3,
				Left: &h.TreeNode{
					Val: 9,
				},
				Right: &h.TreeNode{
					Val: 20,
					Left: &h.TreeNode{
						Val: 15,
					},
					Right: &h.TreeNode{
						Val: 7,
					},
				},
			},
			expected: true,
		},
		{
			name: "unbalanced tree 1",
			root: &h.TreeNode{
				Val: 1,
				Left: &h.TreeNode{
					Val: 2,
					Left: &h.TreeNode{
						Val: 3,
						Left: &h.TreeNode{
							Val: 4,
						},
						Right: &h.TreeNode{
							Val: 4,
						},
					},
					Right: &h.TreeNode{
						Val: 3,
					},
				},
				Right: &h.TreeNode{
					Val: 2,
				},
			},
			expected: false,
		},
		{
			name:     "root is nil",
			expected: true,
		},
		{
			name: "unbalanced tree 2",
			root: &h.TreeNode{
				Val: 1,
				Left: &h.TreeNode{
					Val: 2,
					Left: &h.TreeNode{
						Val: 3,
						Left: &h.TreeNode{
							Val: 4,
						},
					},
				},
				Right: &h.TreeNode{
					Val: 2,
					Right: &h.TreeNode{
						Val: 3,
						Right: &h.TreeNode{
							Val: 4,
						},
					},
				},
			},
			expected: false,
		},
	}

	for _, tc := range tcases {
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(t, tc.expected, isBalanced(tc.root))
		})
	}
}
