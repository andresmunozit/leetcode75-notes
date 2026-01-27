package main

func searchInsert(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}

	if target < nums[0] {
		return 0
	}

	if target > nums[len(nums)-1] {
		return len(nums)
	}

	var result int

	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			result = i
			break
		} else if i > 0 && nums[i-1] < target && nums[i] > target {
			result = i
			break
		}
	}

	return result
}
