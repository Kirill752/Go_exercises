package main

func hightLoadSys2(commands string) string {
	stack := []rune{}
	for _, command := range commands {
		if len(stack) == 0 {
			if string(command) != "M" {
				return "NO"
			}
			stack = append(stack, command)
			continue
		}
		switch string(command) {
		case "M":
			switch string(stack[len(stack)-1]) {
			case "R":
				return "NO"
			case "M":
				return "NO"
			case "C":
				stack = stack[:len(stack)-1]
			case "D":
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, command)
		case "R":
			switch string(stack[len(stack)-1]) {
			case "R":
				return "NO"
			case "M":
				stack = stack[:len(stack)-1]
			case "C":
				stack = stack[:len(stack)-1]
			case "D":
				return "NO"
			}
			stack = append(stack, command)
		case "C":
			switch string(stack[len(stack)-1]) {
			case "R":
				stack = stack[:len(stack)-1]
			case "M":
				stack = stack[:len(stack)-1]
			case "C":
				return "NO"
			case "D":
				return "NO"
			}
			stack = append(stack, command)
		case "D":
			switch string(stack[len(stack)-1]) {
			case "R":
				return "NO"
			case "M":
				stack = stack[:len(stack)-1]
			case "C":
				return "NO"
			case "D":
				return "NO"
			}
		}
	}
	if len(stack) > 0 {
		return "NO"
	}
	return "YES"
}
