package main

import "fmt"

type direction int

const (
	right direction = iota
	down
	left
	up
)

func spiralOrder(matrix [][]int) []int {
	// stablish the direction
	// increase corresponding index up to the limit
	// when limit reach update limits
	// change the direction
	d := right
	r := []int{}
	m := len(matrix)
	n := len(matrix[0])
	size := m * n

	rightLimit := n - 1
	downLimit := m - 1
	leftLimit := 0
	upLimit := 0
	x := 0
	y := 0

	for len(r) < size {
		switch d {
		case right:
			for x <= rightLimit {
				r = append(r, matrix[y][x])
				x++
			}
			// Reset to the limit
			x = rightLimit
			y++
			upLimit++
			d = down
		case down:
			for y <= downLimit {
				r = append(r, matrix[y][x])
				y++
			}
			y = downLimit
			x--
			rightLimit--
			d = left
		case left:
			for x >= leftLimit {
				r = append(r, matrix[y][x])
				x--
			}
			x = leftLimit
			y--
			downLimit--
			d = up
		case up:
			for y >= upLimit {
				r = append(r, matrix[y][x])
				y--
			}
			y = upLimit
			x++
			leftLimit++
			d = right
		}
	}

	return r
}

func main() {
	fmt.Println(spiralOrder([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))             // [1,2,3,6,9,8,7,4,5]
	fmt.Println(spiralOrder([][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}})) // [1,2,3,4,8,12,11,10,9,5,6,7]
}
