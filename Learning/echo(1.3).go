package main

// Echo выводит аргументы командной строки,
import (
	"fmt"
	"os"
	"strings"
	"time"
)

func Echo() {
	secs := time.Since(time.Now()).Seconds()
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
	fmt.Printf("Executing time: %fs\n", secs)
}
func Echo_Join() {
	secs := time.Since(time.Now()).Seconds()
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Printf("Executing time: %fs\n", secs)
}
