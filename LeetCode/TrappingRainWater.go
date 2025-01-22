package main

func trap(height []int) int {
	L, R := 0, 0
	totalVol := 0
	notUsedVol := 0
	for ; R < len(height); R++ {
		if height[R] > height[L] {
			totalVol += (R-L-1)*min(height[R], height[L]) - notUsedVol
			notUsedVol = 0
			L = R
		} else {
			if R != L {
				notUsedVol += height[R]
			}
		}
	}
	L = R - 1
	R = R - 1
	notUsedVol = 0
	for ; L > -1; L-- {
		if height[L] >= height[R] && R != L {
			totalVol += (R-L-1)*min(height[R], height[L]) - notUsedVol
			notUsedVol = 0
			R = L
		} else {
			if R != L {
				notUsedVol += height[L]
			}
		}
	}
	return totalVol
}
