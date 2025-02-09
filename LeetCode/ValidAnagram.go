package main

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	need := make([]int, 28)
	for i := 0; i < len(s); i++ {
		need[s[i]-'a']++
	}
	for i := 0; i < len(t); i++ {
		need[t[i]-'a']--
		if need[t[i]-'a'] < 0 {
			return false
		}
	}
	return true
}
