package main

func findOrder(numCourses int, prerequisites [][]int) []int {
	if numCourses == 0 {
		return []int{}
	}
	if len(prerequisites) == 0 {
		res := make([]int, numCourses)
		for i := range numCourses {
			res[i] = i
		}
		return res
	}
	courses := make(map[int][]int, numCourses)
	coursesRev := make(map[int][]int, numCourses)
	for _, prerequisite := range prerequisites {
		if prerequisite[0] == prerequisite[1] {
			return []int{}
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
		return []int{}
	}
	seen := map[int]bool{}
	res := []int{}
	for i := range numCourses {
		_, ok1 := courses[i]
		_, ok2 := coursesRev[i]
		if !ok1 && !ok2 {
			seen[i] = true
			res = append(res, i)
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
				res = append(res, crs)
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
	if len(res) == numCourses {
		return res
	}
	return []int{}
}
