package main

// Echo выводит аргументы командной строки,
import (
	"fmt"
	"os"
)

func Echo() {
	s, sep := "", ""
	fmt.Println(os.Args[0])
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
