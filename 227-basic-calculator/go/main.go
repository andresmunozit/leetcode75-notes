package main

import (
	"fmt"
	"strconv"
)

// https://app.diagrams.net/?splash=0#G1Bdv3pJldRrI0ZvtGPUduoeV_miTkKvzS#%7B%22pageId%22%3A%22B6OI36_zwcyHBWsVnb9e%22%7D

func calculate(s string) int {
	currStrNum := ""
	currOp := "+"
	stack := []int{}

	for i, c := range s {
		// Transform c to string for easier debugging
		strC := string(c)
		if strC == " " {
			if i < len(s)-1 {
				continue
			}
			strC = ""
		}

		if !(strC == "+" || strC == "-" || strC == "*" || strC == "/") {
			currStrNum += strC
			if i < len(s)-1 {
				continue
			}
		}

		switch currOp {
		case "*":
			n, _ := strconv.Atoi(currStrNum)
			stack[len(stack)-1] = stack[len(stack)-1] * n
		case "/":
			n, _ := strconv.Atoi(currStrNum)
			stack[len(stack)-1] = stack[len(stack)-1] / n
		case "+":
			n, _ := strconv.Atoi(currStrNum)
			stack = append(stack, n)
		case "-":
			n, _ := strconv.Atoi(currStrNum)
			n = -n
			stack = append(stack, n)
		}

		currOp = strC
		currStrNum = ""
	}

	r := 0
	for _, n := range stack {
		r += n
	}

	return r
}

func main() {
	fmt.Println(calculate(" 3/2 "))     // 1
	fmt.Println(calculate("3+2*2"))     // 7
	fmt.Println(calculate("42"))        // 42
	fmt.Println(calculate(" 3+5 / 2 ")) // 5
}
