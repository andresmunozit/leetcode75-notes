package main

import (
	"fmt"
	"math"
)

func mySqrt(x int) int {
	// Newton algorithm
	guess := float64(x / 2)
	newGuess := (guess + float64(x)/guess) / 2
	for guess-newGuess > math.Abs(float64(0.1)) {
		guess = newGuess
		newGuess = (guess + float64(x)/guess) / 2
	}
	return int(newGuess)
}

func main() {
	fmt.Println(mySqrt(4))
	fmt.Println(mySqrt(8))
	fmt.Println(mySqrt(49))
	fmt.Println(mySqrt(100))
}

// 2
// 2
// 7
// 10
