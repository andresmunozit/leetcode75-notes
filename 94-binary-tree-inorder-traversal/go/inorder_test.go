package bintree

import (
	"testing"

	helpersgo "github.com/andresmunozit/leetcode75-notes/helpers-go"

	"github.com/stretchr/testify/require"
)

func Test_inorderTraversal(t *testing.T) {
	tests := []struct {
		name     string
		root     *helpersgo.TreeNode
		expected []int
	}{
		{
			name: "example 1",
			root: func() *helpersgo.TreeNode {
				root := helpersgo.TreeNode{
					Val: 1,
					Right: &helpersgo.TreeNode{
						Val: 2,
						Left: &helpersgo.TreeNode{
							Val: 3,
						},
					},
				}

				return &root
			}(),
			expected: []int{1, 3, 2},
		},
		{
			name: "example 2",
			root: func() *helpersgo.TreeNode {
				root := helpersgo.TreeNode{
					Val: 1,
					Left: &helpersgo.TreeNode{
						Val: 2,
						Left: &helpersgo.TreeNode{
							Val: 4,
						},
						Right: &helpersgo.TreeNode{
							Val: 5,
							Left: &helpersgo.TreeNode{
								Val: 6,
							},
							Right: &helpersgo.TreeNode{
								Val: 7,
							},
						},
					},
					Right: &helpersgo.TreeNode{
						Val: 3,
						Right: &helpersgo.TreeNode{
							Val: 8,
							Left: &helpersgo.TreeNode{
								Val: 9,
							},
						},
					},
				}

				return &root
			}(),
			expected: []int{4, 2, 6, 5, 7, 1, 3, 9, 8},
		},
		{
			name:     "root is nil",
			root:     nil,
			expected: []int{},
		},
		{
			name: "root has no leafs",
			root: &helpersgo.TreeNode{
				Val: 1,
			},
			expected: []int{1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.expected, inorderTraverse(tt.root))
		})
	}
}
