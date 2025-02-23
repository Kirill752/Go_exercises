package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"
)

// Канал отмены
var done = make(chan struct{})

// cancelled опрашивает канал отмены
func cancelled() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}

// walkDir рекусивно обходит дерево файлов с корнем dir
// и отправляет размер каждого в fileSizes
func walkDir(dir string, fileSizes chan<- int64, wg *sync.WaitGroup) {
	defer wg.Done()
	if cancelled() {
		return
	}
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			wg.Add(1)
			subDir := filepath.Join(dir, entry.Name())
			go walkDir(subDir, fileSizes, wg)
		} else {
			inf, err := entry.Info()
			if err != nil {
				fmt.Fprintf(os.Stderr, "walkDir: %v\n", err)
			}
			fileSizes <- inf.Size()
		}
	}
}

// Подсчитывающий семафор для ограничения параллельности
var sema = make(chan struct{}, 20)

// dirents возвращает записи каталога dir
func dirents(dir string) []os.DirEntry {
	select {
	case <-done:
		return nil
	case sema <- struct{}{}: //Захват маркера
	}
	defer func() { <-sema }() //Освобождение марекра
	entries, err := os.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "dirents: %v\n", err)
		return nil
	}
	return entries
}

func printDiskUsage(dir string, nfiles, nbytes int64) {
	fmt.Printf(`"%s": %d файлов %.1f GB`+"\n", dir, nfiles, float64(nbytes/1e9))
}

// DiskUsage параллельно обходит каталог и выводит количество файлов и занимаемую память

func DiskUsage(dir string) {
	var wg sync.WaitGroup
	// Обход дерева файлов
	fileSizes := make(chan int64)
	wg.Add(1)
	go walkDir(dir, fileSizes, &wg)
	// Закрытие канала
	go func() {
		wg.Wait()
		close(fileSizes)
	}()
	// Плдсчет веса каталогов
	var nfiles, nbytes int64
	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(500 * time.Millisecond)
	}
loop:
	for {
		select {
		case <-done:
			// Опустошение канала fileSizes, чтобы позволить
			// завершиться существующим go -подпрограммам
			for range fileSizes {
			}
			return
		case size, ok := <-fileSizes:
			// Если канал закрыт
			if !ok {
				break loop
			}
			nfiles++
			nbytes += size
		case <-tick:
			printDiskUsage(dir, nfiles, nbytes)
		}
	}
	printDiskUsage(dir, nfiles, nbytes)
}
