package main

func countDIFF(a, b string) int {
	c := max(len(a), len(b)) - min(len(a), len(b))
	for i := 0; i < min(len(a), len(b)); i++ {
		if a[i] != b[i] {
			c++
		}
	}
	return c
}
func minMutation(startGene string, endGene string, bank []string) int {
	type gen struct {
		name string
		st   int
	}
	que := []gen{}
	seen := make(map[string]bool)
	que = append(que, gen{startGene, 0})
	for len(que) > 0 {
		curr := que[0]
		seen[curr.name] = true
		que = que[1:]
		if curr.name == endGene {
			return curr.st
		}
		for _, g := range bank {
			if !seen[g] && countDIFF(curr.name, g) == 1 {
				que = append(que, gen{g, curr.st + 1})
			}
		}
	}
	return -1
}
