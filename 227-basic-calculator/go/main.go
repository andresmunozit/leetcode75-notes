package main

import (
	"fmt"
	"strconv"
)

https://app.diagrams.net/?splash=0#G1Bdv3pJldRrI0ZvtGPUduoeV_miTkKvzS#%7B%22pageId%22%3A%22B6OI36_zwcyHBWsVnb9e%22%7D

func calculate(s string) int {
	// Insights
	// If op changes convert the str to num and add it to the stack
	currOp := ""
	currStrNum := ""
	stack := []int{}

	for _, c := range s {
		if isOp(c) {
			currNum, _ := strconv.Atoi(currStrNum)
			stack = append(stack, currNum)
			currStrNum = ""
			currOp = string(c)
			continue
		}
		fmt.Println(currOp)
		currStrNum += string(c)
	}

	currNum, _ := strconv.Atoi(currStrNum)
	stack = append(stack, currNum)

	fmt.Println(stack)

	return 0
}

func isOp(c rune) bool {
	return c == '+' || c == '-' || c == '*' || c == '/'
}

func main() {
	fmt.Println(calculate("3+2*2")) // 7
}

// func sumStack(nums []int) int {
// 	sum := 0
// 	for len(nums) > 0 {

// 		n +=
// 	}
// }
