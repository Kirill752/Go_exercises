package main

func groupAnagrams(strs []string) (res [][]string) {
	anagrams := make(map[[28]int][]string)
	for i := 0; i < len(strs); i++ {
		var chars [28]int
		for j := 0; j < len(strs[i]); j++ {
			chars[strs[i][j]-'a']++
		}
		anagrams[chars] = append(anagrams[chars], strs[i])
	}
	for _, val := range anagrams {
		res = append(res, val)
	}
	return
}
