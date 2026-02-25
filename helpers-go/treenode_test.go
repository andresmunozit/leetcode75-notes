package helpersgo

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTreeNodeJSON(t *testing.T) {
	tcases := []struct {
		name     string
		node     TreeNode
		expected string
	}{
		{
			name: "single node",
			node: TreeNode{
				Val: 5,
			},
			expected: `{"val":5}`,
		},
		{
			name: "two levels node",
			node: TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 8,
				},
				Right: &TreeNode{
					Val: 5,
				},
			},
			expected: `{"val":3,"left":{"val":8},"right":{"val":5}}`,
		},
		{
			name: "three levels node",
			node: TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 0,
					Left: &TreeNode{
						Val: 12,
					},
					Right: &TreeNode{
						Val: 4,
					},
				},
				Right: &TreeNode{
					Val: 5,
					Right: &TreeNode{
						Val: 43,
					},
				},
			},
			expected: `{"val":2,"left":{"val":0,"left":{"val":12},"right":{"val":4}},"right":{"val":5,"right":{"val":43}}}`,
		},
	}

	for _, tc := range tcases {
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(t, tc.expected, tc.node.JSON())
		})
	}
}

func TestTreeNodeJSONIndent(t *testing.T) {
	tcases := []struct {
		name     string
		node     TreeNode
		expected string
	}{
		{
			name: "single node",
			node: TreeNode{
				Val: 5,
			},
			expected: `{
	"val": 5
}`,
		},
		{
			name: "two levels node",
			node: TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 8,
				},
				Right: &TreeNode{
					Val: 5,
				},
			},
			expected: `{
	"val": 3,
	"left": {
		"val": 8
	},
	"right": {
		"val": 5
	}
}`,
		},
		{
			name: "three levels node",
			node: TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 0,
					Left: &TreeNode{
						Val: 12,
					},
					Right: &TreeNode{
						Val: 4,
					},
				},
				Right: &TreeNode{
					Val: 5,
					Right: &TreeNode{
						Val: 43,
					},
				},
			},
			expected: `{
	"val": 2,
	"left": {
		"val": 0,
		"left": {
			"val": 12
		},
		"right": {
			"val": 4
		}
	},
	"right": {
		"val": 5,
		"right": {
			"val": 43
		}
	}
}`,
		},
	}

	for _, tc := range tcases {
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(t, tc.expected, tc.node.JSONIndent())
		})
	}
}
