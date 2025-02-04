package main

func lengthOfLongestSubstring(s string) int {
	usedLetters := make(map[rune]int)
	L := 0
	maxLength := 0
	for R, ch := range s {
		if _, ok := usedLetters[ch]; ok {
			L = usedLetters[ch] + 1
			for i, v := range usedLetters {
				if v < L {
					delete(usedLetters, rune(i))
				}
			}
		} else {
			maxLength = max(maxLength, (R-L)+1)
		}
		usedLetters[ch] = R
	}
	return maxLength
}
