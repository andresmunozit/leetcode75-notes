package main

import "fmt"

func isValid(s string) bool {
	stack := []byte{}

	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '{' || s[i] == '[' {
			stack = append(stack, s[i])
			continue
		}

		if len(stack) < 1 {
			return false
		}

		if s[i] == ')' && stack[len(stack)-1] == '(' ||
			s[i] == '}' && stack[len(stack)-1] == '{' ||
			s[i] == ']' && stack[len(stack)-1] == '[' {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}

	return len(stack) == 0
}

func main() {
	fmt.Println(isValid("()"))
	fmt.Println(isValid("()[]{}"))
	fmt.Println(isValid("(]"))
	fmt.Println(isValid("([])"))
	fmt.Println(isValid("([)]"))
}

// true
// true
// false
// true
// false
