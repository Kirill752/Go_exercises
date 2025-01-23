package main

func romanToInt(s string) int {
	romanNums := map[string]int{
		"I": 1, "V": 5, "X": 10,
		"L": 50, "C": 100, "D": 500, "M": 1000,
		"IV": 4, "IX": 9, "XL": 40, "XC": 90,
		"CD": 400, "CM": 900}
	intNum := 0
	i := 0
	for i < len(s)-2 {
		if _, ok := romanNums[s[i:i+2]]; ok {
			intNum += romanNums[s[i:i+2]]
			i += 2
		} else {
			if _, ok := romanNums[string(s[i])]; ok {
				intNum += romanNums[string(s[i])]
			}
			i++
		}
	}
	if _, ok := romanNums[s[i:]]; ok {
		intNum += romanNums[s[i:]]
	} else {
		for ; i < len(s); i++ {
			intNum += romanNums[string(s[i])]
		}
	}
	return intNum
}
