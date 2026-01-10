package main

import (
	"fmt"
	"strconv"
)

func addBinary(a, b string) string {
	aNum, err := strconv.ParseInt(a, 2, 32)
	if err != nil {
		panic(err)
	}

	bNum, err := strconv.ParseInt(b, 2, 32)
	if err != nil {
		panic(err)
	}

	return fmt.Sprintf("%b", aNum+bNum)
}

func main() {
	fmt.Println(addBinary("11", "1"))      // 100
	fmt.Println(addBinary("1010", "1011")) // 10101
}
