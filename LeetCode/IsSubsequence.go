package main

func isSubsequence(s string, t string) bool {
	if len(s) > len(t) {
		return false
	}
	if len(s) == 0 {
		return true
	}
	sInd := 0
	for _, v := range t {
		if byte(v) == s[sInd] {
			sInd++
		}
		if sInd == len(s) {
			return true
		}
	}
	return false
}
