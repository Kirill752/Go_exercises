package main

// Данная программа выводит время в определнном регионе
// Пример запуска
// TZ=Asia/Tokyo ./clock -port 8020  &
import (
	"flag"
	"fmt"
)

func main() {
	p := flag.String("port", "8000", "string")
	flag.Parse()
	fmt.Println(*p)
	clock(*p)
}
