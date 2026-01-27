package main

func plusOne(digits []int) []int {
	// iterate backwards, add 1 and track the carry
	carry := 0

	for i := len(digits) - 1; i >= 0; i-- {
		result := 0

		if i == len(digits)-1 {
			result = digits[i] + 1
		} else {
			result = digits[i] + carry
		}

		if result > 9 {
			carry = 1
			digits[i] = result - 10
		} else {
			digits[i] = result
			return digits
		}

		if i == 0 {
			if carry == 0 {
				return digits
			} else {
				carrySlice := []int{1}
				return append(carrySlice, digits...)
			}
		}
	}

	return []int{}
}
