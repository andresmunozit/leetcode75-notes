package main

import "fmt"

func searchRange(nums []int, target int) []int {
	return []int{leftSearch(nums, target), rightSearch(nums, target)}
}

func leftSearch(nums []int, target int) int {
	l := 0
	r := len(nums) - 1
	i := -1
	for l <= r {
		m := (l + r) / 2
		if nums[m] < target {
			l = m + 1
		} else {
			r = m - 1
		}
		if nums[m] == target {
			i = m
		}
	}
	return i
}

func rightSearch(nums []int, target int) int {
	l := 0
	r := len(nums) - 1
	i := -1

	for l <= r {
		m := (l + r + 1) / 2
		if nums[m] > target {
			r = m - 1
		} else {
			l = m + 1
		}
		if nums[m] == target {
			i = m
		}
	}
	return i
}

func main() {
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 8)) // [3. 4]
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 6)) // [-1. -1]
	fmt.Println(searchRange([]int{}, 0))                  // [-1. -1]
}
