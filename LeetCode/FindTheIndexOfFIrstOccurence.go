package main

func strStr(haystack string, needle string) int {
	if len(haystack) < len(needle) {
		return -1
	}
	for i := range haystack {
		if haystack[i] == needle[0] {
			j := i
			for ; j < i+len(needle); j++ {
				if j < len(haystack) && haystack[j] == needle[j-i] {
					continue
				} else {
					break
				}
			}
			if j == i+len(needle) {
				return i
			}
		}
	}
	return -1
}
