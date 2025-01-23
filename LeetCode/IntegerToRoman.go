package main

import (
	"slices"
)

func intToRoman(num int) string {
	res := []byte{}
	for i := 0; num > 0; i++ {
		digit := (num % 10)
		num = num / 10
		switch i {
		case 0:
			switch {
			case digit == 0:
				continue
			case digit < 4:
				for i := 0; i < digit; i++ {
					res = append(res, 'I')
				}
			case digit < 9:
				if digit == 4 {
					res = append(res, 'V', 'I')
				} else {
					for i := 0; i < digit-5; i++ {
						res = append(res, 'I')
					}
					res = append(res, 'V')
				}
			case digit < 10:
				res = append(res, 'X', 'I')
			}
		case 1:
			switch {
			case digit == 0:
				continue
			case digit < 4:
				for i := 0; i < digit; i++ {
					res = append(res, 'X')
				}
			case digit < 9:
				if digit == 4 {
					res = append(res, 'L', 'X')
				} else {
					for i := 0; i < digit-5; i++ {
						res = append(res, 'X')
					}
					res = append(res, 'L')
				}
			case digit < 10:
				res = append(res, 'C', 'X')
			}
		case 2:
			switch {
			case digit == 0:
				continue
			case digit < 4:
				for i := 0; i < digit; i++ {
					res = append(res, 'C')
				}
			case digit < 9:
				if digit == 4 {
					res = append(res, 'D', 'C')
				} else {
					for i := 0; i < digit-5; i++ {
						res = append(res, 'C')
					}
					res = append(res, 'D')
				}
			case digit < 10:
				res = append(res, 'M', 'C')
			}
		case 3:
			switch {
			case digit == 0:
				continue
			case digit < 4:
				for i := 0; i < digit; i++ {
					res = append(res, 'M')
				}
			}
		default:
			return "Couldn`t convert to Roman. num > 3999"
		}
	}
	slices.Reverse(res)
	return string(res)
}
