package main

func deleteNumber(number string) string {
	if len(number) <= 1 {
		return "0"
	}
	var result string = "0"
	marker := false
	for i := 1; i < len(number); i++ {
		if number[i] > number[i-1] {
			result = number[:i-1] + number[i:]
			marker = true
			break
		}
	}
	if marker == false && len(number) > 1 {
		return number[:len(number)-1]
	}
	return result
}
