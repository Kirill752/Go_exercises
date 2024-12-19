package main

func SQRT(x float64) float64 {
	count := 10
	result := x / 2
	for i := 0; i < count; i++ {
		result -= (result*result - x) / (2 * result)
	}
	return result
}
