package sametree

import helpers "github.com/andresmunozit/leetcode75-notes/helpers-go"

// Recursion
func isSameTree(p *helpers.TreeNode, q *helpers.TreeNode) bool {
	// Base cases
	if p == nil && q == nil {
		return true
	}

	if (p == nil) != (q == nil) {
		return false
	}

	if p.Val != q.Val {
		return false
	}

	// Last nodes
	if p.Left == nil && q.Left == nil && p.Right == nil && q.Right == nil {
		return true
	}

	// p == q so far, so we continue the recursion
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
