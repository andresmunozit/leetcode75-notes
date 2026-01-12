package main

import (
	"fmt"
)

// func addBinary(a, b string) string {
// 	// Make arrays and reverse
// }

func main() {
	// fmt.Println(addBinary("11", "1"))      // 100
	// fmt.Println(addBinary("1010", "1011")) // 10101
	fmt.Println(reverse([]string{"a", "b", "c"}))
}

func reverse(s []string) []string {
	rev := []string{}
	for i := len(s) - 1; i >= 0; i-- {
		rev = append(rev, s[i])
	}
	return rev
}
