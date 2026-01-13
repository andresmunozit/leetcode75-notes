package main

import (
	"fmt"
	"strconv"
	"strings"
)

func addBinary(a, b string) string {
	// Make arrays and reverse
	sliceA := strings.Split(a, "")
	sliceB := strings.Split(b, "")

	revA := reverse(sliceA)
	revB := reverse(sliceB)

	carry := 0
	result := ""

	for i := 0; i < max(len(a), len(b)); i++ {
		var numA int
		var numB int

		// Convert strings into numbers
		// If out of slice bounds let the default 0 value
		if i <= len(a)-1 {
			numA, _ = strconv.Atoi(revA[i])
		}
		if i <= len(b)-1 {
			numB, _ = strconv.Atoi(revB[i])
		}

		// Calculate the sum
		sum := numA + numB + carry

		strBinSum := fmt.Sprintf("%b", sum)
		strBinSumBits := strings.Split(strBinSum, "")

		strBit := ""

		if len(strBinSumBits) > 1 {
			carry = 1
			strBit = strBinSumBits[1]
		} else {
			carry = 0
			strBit = strBinSumBits[0]
		}

		result = strBit + result
	}

	if carry == 1 {
		result = "1" + result
	}

	return result
}

func main() {
	fmt.Println(addBinary("11", "1"))
	fmt.Println(addBinary("1010", "1011"))
	fmt.Println(addBinary(
		"10100000100100110110010000010101111011011001101110111111111101000000101111001110001111100001101",
		"110101001011101110001111100110001010100001101011101010000011011011001011101111001100000011011110011",
	))
}

// 100
// 10101
// 110111101100010011000101110110100000011101000101011001000011011000001100011110011010010011000000000

func reverse(s []string) []string {
	rev := []string{}
	for i := len(s) - 1; i >= 0; i-- {
		rev = append(rev, s[i])
	}
	return rev
}

// 100
// 10101
// 110111101100010011000101110110100000011101000101011001000011011000001100011110011010010011000000000
