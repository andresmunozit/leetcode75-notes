package main

func strStr(haystack string, needle string) int {
	// iterate over the haystack to check if it contains needle
	for i := 0; i <= len(haystack)-len(needle); i++ {
		if haystack[i:len(needle)+i] == needle {
			return i
		}
	}
	return -1
}
