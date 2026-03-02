package main

import (
	"testing"

	h "github.com/andresmunozit/leetcode75-notes/helpers-go"
	"github.com/stretchr/testify/require"
)

func Test_minDepth(t *testing.T) {
	tcases := []struct {
		name     string
		root     *h.TreeNode
		expected int
	}{
		{
			name: "tree 1",
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
			expected: 2,
		},
		{
			name: "tree 2",
			root: &h.TreeNode{
				Val: 2,
				Right: &h.TreeNode{
					Val: 3,
					Right: &h.TreeNode{
						Val: 4,
						Right: &h.TreeNode{
							Val: 5,
							Right: &h.TreeNode{
								Val: 6,
							},
						},
					},
				},
			},
			expected: 5,
		},
		{
			name: "tree 2",
			root: &h.TreeNode{
				Val: 2,
				Right: &h.TreeNode{
					Val: 3,
					Right: &h.TreeNode{
						Val: 4,
						Right: &h.TreeNode{
							Val: 5,
							Right: &h.TreeNode{
								Val: 6,
							},
						},
					},
				},
			},
			expected: 5,
		},
		{
			name: "tree 3",
			root: &h.TreeNode{
				Val: -9,
				Left: &h.TreeNode{
					Val: -3,
					Right: &h.TreeNode{
						Val: 4,
						Left: &h.TreeNode{
							Val: -6,
						},
					},
				},
				Right: &h.TreeNode{
					Val: 2,
					Left: &h.TreeNode{
						Val: 4,
						Left: &h.TreeNode{
							Val: -5,
						},
					},
					Right: &h.TreeNode{
						Val: 10,
					},
				},
			},
			expected: 3,
		},
	}

	for _, tc := range tcases {
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(t, tc.expected, minDepth(tc.root))
		})
	}
}
