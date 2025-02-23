package main

import (
	"flag"
	"os"
	"sync"
)

var verbose = flag.Bool("v", false, "Вывод промежуточных результатов")

func main() {
	// Отмена обхода при обнаруженном вводе
	go func() {
		os.Stdin.Read(make([]byte, 1))
		close(done)
	}()
	// Определяем исходные каталоги
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}
	var wg sync.WaitGroup
	for _, root := range roots {
		wg.Add(1)
		go func() {
			defer wg.Done()
			DiskUsage(root)
		}()
	}
	wg.Wait()
}
