package main

func canFinish(numCourses int, prerequisites [][]int) bool {
	if len(prerequisites) == 0 || numCourses == 0 {
		return true
	}
	courses := make(map[int][]int, numCourses)
	coursesRev := make(map[int][]int, numCourses)
	for _, prerequisite := range prerequisites {
		if prerequisite[0] == prerequisite[1] {
			return false
		}
		courses[prerequisite[1]] = append(courses[prerequisite[1]], prerequisite[0])
		coursesRev[prerequisite[0]] = append(coursesRev[prerequisite[0]], prerequisite[1])
	}
	// Определяем начальный курс
	start := []int{}
	for v := range courses {
		if _, ok := coursesRev[v]; !ok {
			start = append(start, v)
		}
	}
	if len(start) == 0 {
		return false
	}
	seen := map[int]bool{}
	for i := range numCourses {
		_, ok1 := courses[i]
		_, ok2 := coursesRev[i]
		if !ok1 && !ok2 {
			seen[i] = true
		}
	}
	// Обход в ширину со стартового курса
	que := []int{}
	for _, s := range start {
		if !seen[s] {
			que = append(que, s)
			for len(que) > 0 {
				crs := que[0]
				que = que[1:]
				seen[crs] = true
				for _, v := range courses[crs] {
					if !seen[v] {
						canGet := true
						for _, n := range coursesRev[v] {
							if !seen[n] {
								canGet = false
							}
						}
						if canGet {
							que = append(que, v)
						}
					}
				}
			}
		}
	}
	return len(seen) == numCourses
}
