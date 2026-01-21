package main

import (
	"fmt"
)

func romanToInt(s string) int {
	// create a byte int map
	values := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	sum := 0

	// iterate over s chars, subtract if curr < next, add otherwise
	for i := 0; i < len(s); i++ {
		curr := values[s[i]]

		if i < len(s)-1 && curr < values[s[i+1]] {
			sum -= curr
		} else {
			sum += curr
		}
	}

	return sum
}

func main() {
	fmt.Println(romanToInt("III"))
	fmt.Println(romanToInt("LVIII"))
	fmt.Println(romanToInt("MCMXCIV"))
}

// 3
// 58
// 1994
