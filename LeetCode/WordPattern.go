package main

import "strings"

func wordPattern(pattern string, s string) bool {
	replace := make(map[rune]string)
	used := make(map[string]bool)
	words := strings.Split(s, " ")
	if len(pattern) != len(words) {
		return false
	}
	for i, ch := range pattern {
		if _, ok := replace[ch]; !ok {
			if !used[words[i]] {
				replace[ch] = words[i]
				used[words[i]] = true
			} else {
				return false
			}
		} else {
			if replace[ch] != words[i] {
				return false
			}
		}
	}
	return true
}
