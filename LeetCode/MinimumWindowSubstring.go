package main

func minWindow(s string, t string) string {
	need := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}
	minLen := 100001
	minSub := []byte{}
	L := 0
	count := 0
	R := L
	for ; R < len(s); R++ {
		if need[s[R]] > 0 {
			count++
		}
		need[s[R]]--
		for count == len(t) {
			if (R-L)+1 < minLen {
				minSub = []byte(s[L : R+1])
				minLen = len(minSub)
			}
			need[s[L]]++
			if need[s[L]] > 0 {
				count--
			}
			L++
		}
	}
	return string(minSub)
}
