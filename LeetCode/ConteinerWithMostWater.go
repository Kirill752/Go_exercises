package main

func maxArea(height []int) int {
	L, R := 0, len(height)-1
	volume := min(height[L], height[R]) * (R - L)
	for L < R {
		curVol := min(height[L], height[R]) * (R - L)
		if curVol > volume {
			volume = curVol
		}
		if height[L] >= height[R] {
			R--
		} else {
			L++
		}
	}
	return volume
}
