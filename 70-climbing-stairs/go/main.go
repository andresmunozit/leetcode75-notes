package main

func climbStairs(n int) int {
	ways := []int{1, 1, 2}

	for i := 3; i <= n; i++ {
		ways = append(ways, ways[i-1]+ways[i-2])
	}

	return ways[n]
}
