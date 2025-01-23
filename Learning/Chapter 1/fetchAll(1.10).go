package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

/*
Fetchall выполняет параллельную выборку URL и сообщает
о затраченном времени и размере ответа для каждого из них.
*/
func fetchAll(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // отправка в канал ch
		return
	}
	nbytes, err := io.Copy(io.Discard, resp.Body)
	resp.Body.Close() // Исключение утечки ресурсов
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}

/*
func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetchAll(url, ch)
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)
		// Получение из канала ch
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}
*/
/*Пример запуска:
./execute https://golang.org http://gopl.io https://godoc.org
1.11s   33513 https://godoc.org
1.32s   62384 https://golang.org
2.24s    4154 http://gopl.io
2.24s elapsed
*/
