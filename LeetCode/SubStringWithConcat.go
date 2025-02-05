package main

func findSubstring(s string, words []string) (idx []int) {
	setOfWords := make(map[string]int)
	for _, word := range words {
		setOfWords[word]++
	}
	if len(words)*len(words[0]) > len(s) {
		return
	}
	length := len(words[0])
	for i := 0; i < len(s)-length*len(words)+1; i++ {
		seen := make(map[string]int)

		for j := 0; j < len(words); j++ {
			nextIndex := i + j*length
			word := s[nextIndex : nextIndex+length]

			if _, ok := setOfWords[word]; !ok {
				break
			}

			seen[word]++

			seenFrequency := seen[word]
			originFrequency := setOfWords[word]

			if seenFrequency > originFrequency {
				break
			}

			if j+1 == len(words) {
				idx = append(idx, i)
			}
		}
	}
	return
}
