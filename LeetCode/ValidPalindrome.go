package main

import (
	"strings"
	"unicode"
)

func isPalindrome(s string) bool {
	L, R := 0, len(s)-1
	s = strings.ToLower(s)
	for L <= R {
		for L < len(s) && (!unicode.IsDigit(rune(s[L])) && !unicode.IsLetter(rune(s[L]))) {
			L++
		}
		if L >= len(s) {
			return true
		}
		for R >= 0 && (!unicode.IsDigit(rune(s[R])) && !unicode.IsLetter(rune(s[R]))) {
			R--
		}
		if s[L] != s[R] {
			return false
		}
		L++
		R--
	}
	return true
}
