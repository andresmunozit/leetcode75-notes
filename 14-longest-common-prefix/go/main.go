package main

func longestCommonPrefix(strs []string) string {
	str := strs[0]
	res := ""

	for i := 0; i < len(str); i++ {
		for j := 0; j < len(strs); j++ {
			if i >= len(strs[j]) || str[i] != strs[j][i] {
				return res
			}
		}
		res += string(str[i])
	}

	return res
}
