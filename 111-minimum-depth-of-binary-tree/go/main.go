package main

import h "github.com/andresmunozit/leetcode75-notes/helpers-go"

func minDepth(root *h.TreeNode) int {
	// dfs and return the shortest height
	if root == nil {
		return 0
	}

	leftDepth := minDepth(root.Left)
	rightDepth := minDepth(root.Right)
	if root.Left == nil {
		return 1 + rightDepth
	}

	if root.Right == nil {
		return 1 + leftDepth
	}

	return 1 + min(leftDepth, rightDepth)
}
