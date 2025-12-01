package main

import (
	"fmt"
	"unicode/utf8"
)

func mergeAlternately(word1 string, word2 string) string {
	l1 := utf8.RuneCountInString(word1)
	l2 := utf8.RuneCountInString(word2)
	m := max(l1, l2)

	// keep track of string rune width
	w1 := 0
	w2 := 0
	s := ""
	for i := range m {
		if i < l1 {
			r, w := utf8.DecodeRuneInString(word1[w1:])
			s = s + string(r)
			w1 += w
		}
		if i < l2 {
			r, w := utf8.DecodeRuneInString(word2[w2:])
			s = s + string(r)
			w2 += w
		}
	}
	return s
}

func main() {
	fmt.Println(mergeAlternately("abc", "pqr")) // apbqcr
	fmt.Println(mergeAlternately("ab", "pqrs")) // apbqrs
	fmt.Println(mergeAlternately("abcd", "pq")) // apbqcd
}
