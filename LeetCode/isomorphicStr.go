package main

func isIsomorphic(s string, t string) bool {
	replace := make(map[rune]rune)
	used := make(map[rune]bool)
	for i, ch := range s {
		if _, ok := replace[ch]; !ok {
			if !used[rune(t[i])] {
				replace[ch] = rune(t[i])
				used[rune(t[i])] = true
			} else {
				return false
			}
		} else {
			if replace[ch] != rune(t[i]) {
				return false
			}
		}
	}
	return true
}
