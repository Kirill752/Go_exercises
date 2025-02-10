package main

import (
	"flag"
	"fmt"
)

func main() {
	var temp = CelsiusFlag("temp", 20.0, "температура")
	flag.Parse()
	fmt.Println(*temp)
}
