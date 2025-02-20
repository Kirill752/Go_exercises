package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sync"

	"gopl.io/ch5/links"
)

type work struct {
	url   string
	depth int
}

var maxDepth = 1

// tokens представляет собой подсчитывающий семафор, используемый
// для ограничения количества параллельных запросов величиной 20.
var tockens = make(chan struct{}, 20)

func crawl(w work) []work {
	fmt.Printf("Depth: %d, URL: %s\n", w.depth, w.url)
	if w.depth > maxDepth {
		return nil
	}
	tockens <- struct{}{} // Захват маркера
	list, err := links.Extract(w.url)
	<-tockens // Освобождение маркера
	if err != nil {
		log.Print(err)
	}
	works := make([]work, len(list))

	for _, url := range list {
		works = append(works, work{url: url, depth: w.depth + 1})
	}
	return works
}

func GetLinks() {
	// flag.IntVar(&maxDepth, "d", 3, "max depth")
	// flag.Parse()

	// Wait group для ожидания окончания работы горутин
	var wg sync.WaitGroup
	worklist := make(chan []work)
	// Количество отправиок в рабочий список
	var n int
	// Запуск с аргументами командной строки
	// первоначальные аргументы командной строки должны передаваться в рабочий список в отдельной
	// go-подпрограмме, чтобы избежать взаимоблокировки, ситуации “зависания”, в которой
	// основная go-подпрограмма и go-подпрограммы сканирования пытаются отправлять
	// информацию друг другу в то время, когда ни одна из них ее не получает.
	n++
	go func() {
		var works []work
		for _, url := range os.Args[1:] {
			works = append(works, work{url, 1})
		}
		worklist <- works
	}()

	// Параллельное сканирование
	seen := make(map[string]bool)
	for ; n > 0; n-- {
		works := <-worklist
		for _, w := range works {
			if !seen[w.url] {
				seen[w.url] = true
				n++
				// go-подпрограмма сканирования принимает link в каче­cтве
				// явного параметра, чтобы избежать проблемы захвата переменной цикла
				wg.Add(1)
				go func(w work) {
					defer wg.Done()
					worklist <- crawl(w)
				}(w)
			}
		}
	}
	wg.Wait()
}

func getData(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return fmt.Errorf("getting %s: %s", url, resp.Status)
	}
	defer resp.Body.Close()
	out, err := os.OpenFile(url[8:]+".html", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}
	out.Close()
	return nil
}
