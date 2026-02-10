package bintree

import (
	helpersgo "github.com/andresmunozit/leetcode75-notes/helpers-go"
)

func inorderTraverse(root *helpersgo.TreeNode) []int {
	stack := make([]*helpersgo.TreeNode, 0)
	result := make([]int, 0)

	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, root.Val)
		root = root.Right
	}

	return result
}
