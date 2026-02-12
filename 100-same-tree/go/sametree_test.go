package sametree

import (
	"testing"

	helpers "github.com/andresmunozit/leetcode75-notes/helpers-go"
	"github.com/stretchr/testify/require"
)

func Test_sameTree(t *testing.T) {
	tests := []struct {
		name     string
		p        *helpers.TreeNode
		q        *helpers.TreeNode
		expected bool
	}{
		{
			name: "equal trees",
			p: &helpers.TreeNode{
				Val: 1,
				Left: &helpers.TreeNode{
					Val: 2,
				},
				Right: &helpers.TreeNode{
					Val: 3,
				},
			},
			q: &helpers.TreeNode{
				Val: 1,
				Left: &helpers.TreeNode{
					Val: 2,
				},
				Right: &helpers.TreeNode{
					Val: 3,
				},
			},
			expected: true,
		},
		{
			name:     "both trees are nil",
			p:        nil,
			q:        nil,
			expected: true,
		},
		{
			name: "different trees 1",
			p: &helpers.TreeNode{
				Val: 1,
				Left: &helpers.TreeNode{
					Val: 2,
				},
			},
			q: &helpers.TreeNode{
				Val: 1,
				Right: &helpers.TreeNode{
					Val: 2,
				},
			},
			expected: false,
		},
		{
			name: "different trees 2",
			p: &helpers.TreeNode{
				Val: 1,
				Left: &helpers.TreeNode{
					Val: 2,
				},
				Right: &helpers.TreeNode{
					Val: 1,
				},
			},
			q: &helpers.TreeNode{
				Val: 1,
				Left: &helpers.TreeNode{
					Val: 1,
				},
				Right: &helpers.TreeNode{
					Val: 2,
				},
			},
			expected: false,
		},
		{
			name: "different trees 3",
			p: &helpers.TreeNode{
				Val: 1,
			},
			q: &helpers.TreeNode{
				Val: 1,
				Right: &helpers.TreeNode{
					Val: 2,
				},
			},
			expected: false,
		},
		{
			name: "different trees 4",
			p: &helpers.TreeNode{
				Val: 1,
			},
			q: &helpers.TreeNode{
				Val: 1,
				Right: &helpers.TreeNode{
					Val: 2,
				},
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.expected, isSameTree(tt.p, tt.q))
		})
	}
}
