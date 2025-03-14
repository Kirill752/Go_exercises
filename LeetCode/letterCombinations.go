package main

func decodeT9(digits string, idx int, currentWord string, result *[]string) {
	if idx == len(digits) {
		*result = append(*result, currentWord)
		return
	}
	// Текущие возможные буквы
	letters := getLetters(digits[idx])
	for i := range len(letters) {
		decodeT9(digits, idx+1, currentWord+string(letters[i]), result)
	}
}

func getLetters(digit byte) string {
	switch digit {
	case '2':
		return "abc"
	case '3':
		return "def"
	case '4':
		return "ghi"
	case '5':
		return "jkl"
	case '6':
		return "mno"
	case '7':
		return "pqrs"
	case '8':
		return "tuv"
	case '9':
		return "wxyz"
	default:
		return ""
	}
}

func letterCombinations(digits string) []string {
	res := []string{}
	decodeT9(digits, 0, "", &res)
	if res[0] == "" {
		return []string{}
	}
	return res
}
