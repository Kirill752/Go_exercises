package main

func isHappy(n int) bool {
	was := make(map[int]bool)
	var calc func(int) bool
	calc = func(n int) bool {
		digits := []int{}
		for n != 0 {
			digit := n % 10
			digits = append(digits, digit)
			n /= 10
		}
		newN := 0
		for _, digit := range digits {
			newN += digit * digit
		}
		if newN == 1 {
			return true
		}
		if was[newN] {
			return false
		}
		was[newN] = true
		return calc(newN)
	}
	res := calc(n)
	return res
}
