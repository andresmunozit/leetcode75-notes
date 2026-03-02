package main

import h "github.com/andresmunozit/leetcode75-notes/helpers-go"

func isBalanced(root *h.TreeNode) bool {
	if root == nil {
		return true
	}

	return dfs(root) != -1
}

func dfs(n *h.TreeNode) int {
	// base cases
	if n == nil {
		return 0
	}

	leftHeight := dfs(n.Left)
	if leftHeight == -1 {
		return -1
	}

	rightHeight := dfs(n.Right)
	if rightHeight == -1 {
		return -1
	}

	if abs(leftHeight-rightHeight) > 1 {
		return -1
	}

	// recursive function
	return 1 + max(dfs(n.Left), dfs(n.Right))
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
