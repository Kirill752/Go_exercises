package main

import "fmt"

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	// Индексируем переменные
	vars := make(map[string]int)
	for _, eq := range equations {
		for i := 0; i < len(eq); i++ {
			if _, ok := vars[eq[i]]; !ok {
				vars[eq[i]] = len(vars)
			}
		}
	}
	// fmt.Println(vars)
	// Формируем матрицу смежности
	adjac := make([][]float64, len(vars))
	for i := 0; i < len(vars); i++ {
		adjac[i] = make([]float64, len(vars))
		for j := 0; j < len(adjac[i]); j++ {
			adjac[i][j] = -1
		}
	}
	for i, v := range values {
		eq := equations[i]
		adjac[vars[eq[0]]][vars[eq[1]]] = v
		adjac[vars[eq[1]]][vars[eq[0]]] = 1.0 / v
		adjac[vars[eq[0]]][vars[eq[0]]] = 1.0
		adjac[vars[eq[1]]][vars[eq[1]]] = 1.0
	}
	for _, v := range adjac {
		fmt.Println(v)
	}
	// Функция обхода графа и поиска пути от делмого к делителю
	// i - координаты делимого
	// j - координаты делителя
	var DFS func(i, j int, ans float64, seen map[int]bool) float64
	// Маркер того, что можно решить задачу
	marker := false
	DFS = func(i, j int, ans float64, seen map[int]bool) float64 {
		seen[i] = true
		if adjac[i][j] != -1.0 {
			marker = true
			return adjac[i][j]
		}
		for m := 0; m < len(adjac[0]); m++ {
			if marker {
				return ans
			}
			if !seen[m] && adjac[i][m] != -1.0 {
				b := DFS(m, j, ans, seen)
				if b != -1 {
					ans *= adjac[i][m] * b
				} else {
					ans = 1
				}
			}
		}
		if marker {
			return ans
		}
		return -1
	}
	ans := make([]float64, len(queries))
	for i, qu := range queries {
		_, ok1 := vars[qu[0]]
		_, ok2 := vars[qu[1]]
		seen := make(map[int]bool)
		marker = false
		if ok1 && ok2 {
			ans[i] = DFS(vars[qu[0]], vars[qu[1]], 1, seen)
		} else {
			ans[i] = -1.0
		}
	}
	return ans
}
