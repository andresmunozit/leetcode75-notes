package main

import (
	"fmt"
	"strings"
)

func gcdOfStrings(str1 string, str2 string) string {
	if str1+str2 != str2+str1 {
		return ""
	}

	var shortest string
	var longest string

	if len(str1) >= len(str2) {
		longest = str1
		shortest = str2
	} else {
		longest = str2
		shortest = str1
	}

	l1 := len(longest)
	l2 := len(shortest)

	var str string
	var gcd string

	for _, c := range shortest {
		str += string(c)
		lgcd := len(str)
		if l1%lgcd == 0 && l2%lgcd == 0 {
			f1 := l1 / lgcd
			f2 := l2 / lgcd

			if strings.Repeat(str, f1) == longest && strings.Repeat(str, f2) == shortest {
				gcd = str
			}
		}
	}
	return gcd
}

func main() {
	fmt.Println(gcdOfStrings("ABCABC", "ABC"))    //
	fmt.Println(gcdOfStrings("ABABAB", "ABAB"))   // AB
	fmt.Println(gcdOfStrings("LEET", "CODE"))     // ""
	fmt.Println(gcdOfStrings("ABABABAB", "ABAB")) // ABAB
}
