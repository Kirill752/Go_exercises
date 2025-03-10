package main

import (
	"fmt"
)

// Фугкция withoutReturn возвращает значение без оператора return
func withoutReturn(num int) (res int) {
	defer func() {
		switch p := recover(); p {
		case "always panic":
			res = num + 5
		default:
			panic(p)
		}
	}()
	panic("always panic")
}

func main() {
	fmt.Println(withoutReturn(10))
}
