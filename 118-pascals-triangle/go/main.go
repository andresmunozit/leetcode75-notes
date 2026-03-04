package main

func generate(numRows int) [][]int {
	pascalTriangle := [][]int{{1}}

	if numRows == 1 {
		return pascalTriangle
	}

	pascalTriangle = append(pascalTriangle, []int{1, 1})

	if numRows == 2 {
		return pascalTriangle
	}

	for i := 3; i <= numRows; i++ {
		row := make([]int, i)

		row[0] = 1
		row[i-1] = 1

		previousRow := pascalTriangle[i-2]

		for j := 1; j < i-1; j++ {
			row[j] = previousRow[j-1] + previousRow[j]
		}

		pascalTriangle = append(pascalTriangle, row)
	}

	return pascalTriangle
}
