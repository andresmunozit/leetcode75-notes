package main

import (
	"fmt"
)

func gcdOfStrings(str1 string, str2 string) string {
	if str1+str2 != str2+str1 {
		return ""
	}

	n1 := len(str1)
	n2 := len(str2)

	// Euclidean algorithm
	for n2 > 0 {
		n1, n2 = n2, n1%n2
	}

	return str1[:n1]
}

func main() {
	fmt.Println(gcdOfStrings("ABCABC", "ABC"))                            // ABC
	fmt.Println(gcdOfStrings("ABABAB", "ABAB"))                           // AB
	fmt.Println(gcdOfStrings("LEET", "CODE"))                             //
	fmt.Println(gcdOfStrings("ABABABAB", "ABAB"))                         // ABAB
	fmt.Println(gcdOfStrings("XDFGYTXDFGYTXDFGYTXDFGYT", "XDFGYTXDFGYT")) // XDFGYTXDFGYT
	fmt.Println(gcdOfStrings("DFGDFGDFGDFG", "DFGDFGDFG"))                // DFG
}
