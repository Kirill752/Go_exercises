package main

func ladderLength(beginWord string, endWord string, wordList []string) int {
	que := []string{}
	seen := make(map[string]bool)
	que = append(que, beginWord)
	st := 1
	seen[que[0]] = true
	for len(que) > 0 {
		l := len(que)
		for i := range l {
			if que[i] == endWord {
				return st
			}
			for _, g := range wordList {
				if !seen[g] && countDIFF(que[i], g) == 1 {
					que = append(que, g)
					seen[g] = true
				}
			}
		}
		que = que[l:]
		st++
	}
	return 0
}
