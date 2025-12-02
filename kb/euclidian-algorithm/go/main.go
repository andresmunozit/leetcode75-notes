package main

import "fmt"

func gcd1(n1 int, n2 int) int {
	r := n1 % n2

	if r == 0 {
		return n2
	}

	return gcd1(n2, r)
}

func gcd2(n1 int, n2 int) int {
	for n2 > 0 {
		// This works because if n1 < n2, n1 is returned
		n1, n2 = n2, n1%n2
	}

	return n1
}

func main() {
	fmt.Println(gcd1(24, 12))
	fmt.Println(gcd1(12, 24))
	fmt.Println(gcd1(15, 9))

	fmt.Println(gcd2(24, 12))
	fmt.Println(gcd2(12, 24))
	fmt.Println(gcd2(15, 9))
}
