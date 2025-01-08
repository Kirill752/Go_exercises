package main

import "fmt"

func canJump(nums []int) bool {
	var n int
	if len(nums) > 2 {
		n = 2
	} else {
		n = len(nums)
	}
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, len(nums))
	}
	// Заполняем первую строку
	for i := 0; i <= nums[0] && i < len(dp[0]); i++ {
		dp[0][i] = 1
	}
	for i := 1; i < len(nums); i++ {
		// Заполняем строку на основе предыдущей строки
		for j := 0; j < len(dp[0]); j++ {
			if dp[(i-1)%2][j] == 0 {
				break
			}
			dp[i%2][j] = dp[(i-1)%2][j]
		}
		if dp[(i-1)%2][i] != 0 {
			for j := i; j <= i+nums[i] && j < len(dp[0]); j++ {
				dp[i%2][j] = 1
			}
		}
		if dp[i%2][len(dp[0])-1] == 1 {
			return true
		}
	}
	if len(dp) == 0 {
		return true
	}
	if dp[len(dp)-1][len(dp[0])-1] == 1 {
		return true
	} else {
		for i := 0; i < len(dp); i++ {
			if dp[i][len(dp[0])-1] == 1 {
				return true
			}
		}
	}
	return false
}

// не уверен, что она полностью верная, так как может оказаться так, что единственная точка
// из которой можно достигнуть цели не достижима сама по себе
func canJump_v2(nums []int) bool {
	goal := len(nums) - 1
	for i := len(nums) - 2; i > -1; i-- {
		if i+nums[i] >= goal {
			goal = i
		}
	}
	if goal == 0 {
		return true
	}
	return false
}

// Выводит минимальное количество прыжков, необходимое для достижения конца
func jump(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	grath := map[int][]int{}
	for i, v := range nums {
		for j := i + 1; j < len(nums) && j <= i+v; j++ {
			grath[i] = append(grath[i], j)
		}
	}
	visited := map[int]bool{}
	queue := []int{}
	queue = append(queue, 0)
	count := map[int]int{}
	count[0] = 0
	for len(queue) > 0 {
		fmt.Println(queue)
		front := queue[0]
		copy(queue[0:], queue[0+1:])
		queue = queue[:len(queue)-1]
		for i, v := range nums[front : front+nums[front]] {
			if i < len(nums) {
				if _, ok := visited[v]; ok == false {
					visited[v] = true
					count[v] = count[front] + 1
					if v == len(nums)-1 {
						return count[v]
					}
					queue = append(queue, v)
				}
			} else {
				break
			}
		}
	}
	return count[0]
}

func jump_v2(nums []int) int {
	res := 0
	l, r := 0, 0
	for r < len(nums) {
		farthest := 0
		for i, v := range nums[l : r+1] {
			farthest = max(farthest, i+v)
		}
		l = r + 1
		r = farthest
		res++
	}
	return res
}
