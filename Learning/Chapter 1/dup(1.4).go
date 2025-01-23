package main

import (
	"bufio"
	"fmt"
	"os"
)

/* dup1 выводит текст каждой строки, которая появляется в
 * стандартном вводе более одного раза, а также количество
 * её появлений.
 * Пример ввода:
 * Как дела?
 * Как дела?
 * Это новый кадилак!
 * Делать деньги!
 * Делать деньги!
 * Делать деньги!
 * Вот так!
 * Ctr+D
 */
func dup1() {
	counts := make(map[string]int)
	countLines(os.Stdin, counts)
	fmt.Println()
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func dup2() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	fmt.Println("Файлы: ")
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()+", Файл: "+f.Name()]++
	}
}
