package main

func findTargetSumWays(nums []int, target int) int {
	maxSUM := 0
	for i := 0; i < len(nums); i++ {
		maxSUM += nums[i]
	}
	dp := make([][]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = make([]int, 2*maxSUM+1)
	}
	// Инициализируем первую строку
	dp[0][maxSUM-nums[0]], dp[0][maxSUM+nums[0]] = 1, 1
	for i := 1; i < len(dp); i++ {
		for j := 0; j < 2*maxSUM+1; j++ {
			if (j+nums[i] < 2*maxSUM+1) && (j+nums[i] > -1) {
				dp[i][j] += dp[i-1][j+nums[i]]
			}
			if (j-nums[i] < 2*maxSUM+1) && (j-nums[i] > -1) {
				dp[i][j] += dp[i-1][j-nums[i]]
			}
		}
	}
	if (maxSUM+target < 2*maxSUM+1) && (maxSUM+target > -1) {
		if nums[0] == 0 {
			return 2 * dp[len(dp)-1][maxSUM+target]
		}
		return dp[len(dp)-1][maxSUM+target]
	}
	return 0
}
