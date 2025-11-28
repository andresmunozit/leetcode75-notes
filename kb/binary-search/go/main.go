package main

import (
	"fmt"
)

var arr = []int{2, 4, 5, 9, 12, 21, 22, 45, 67, 82, 100}

func binarySearch(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	// There can be the case in which left == rigth, win which case both will contain the index
	// of target
	for left <= right {
		// Integer division
		m := (right + left) / 2

		if nums[m] == target {
			return m
		} else if nums[m] > target {
			// This excludes the most right element which we already compared
			right = m - 1
		} else {
			// This excludes the most left element which we already compared
			left = m + 1
		}
	}
	return -1
}

func main() {
	fmt.Println(binarySearch(arr, 9))   // 3
	fmt.Println(binarySearch(arr, 2))   // 0
	fmt.Println(binarySearch(arr, 45))  // 7
	fmt.Println(binarySearch(arr, 101)) // -1
	fmt.Println(binarySearch(arr, 1))   // -1
}
