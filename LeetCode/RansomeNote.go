package main

func canConstruct(ransomNote string, magazine string) bool {
	need := make([]int, 28)
	for _, v := range magazine {
		need[v-'a']++
	}
	for _, v := range ransomNote {
		need[v-'a']--
		if need[v-'a'] < 0 {
			return false
		}
	}
	return true
}
