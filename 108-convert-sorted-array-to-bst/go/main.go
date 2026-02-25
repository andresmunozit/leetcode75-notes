package main

import (
	h "github.com/andresmunozit/leetcode75-notes/helpers-go"
)

// Input example
// -10, -3, 0, 5, 9

// Output example
//         0
//        / \
//      -3   9
//      /   /
//   -10   5

func sortedArrayToBST(nums []int) *h.TreeNode {
	return helper(&nums, 0, len(nums)-1)
}

func helper(nums *[]int, l, r int) *h.TreeNode {
	// base condition
	if r < l {
		return nil
	}

	m := (l + r) / 2

	return &h.TreeNode{
		Val:   (*nums)[m],
		Left:  helper(nums, l, m-1),
		Right: helper(nums, m+1, r),
	}
}

// func main() {
// 	fmt.Println(sortedArrayToBST([]int{-10, -3, 0, 5, 9}).JSONIndent())
// 	fmt.Println(sortedArrayToBST([]int{1, 3}).JSONIndent())
// }
