package main

import (
	"fmt"
	"net/http"
	"strconv"
	"sync"
)

var mu sync.Mutex
var count int

// Минимальный echo сервер
func url_path(w http.ResponseWriter, r *http.Request) {
	// Запрос представлен структурой типа http.Request,
	// которая представляет собой ряд связанных полей, одним из
	// которых является URL входящего запроса. Полученный запрос
	// передаётся функции обработчику, которая извлекает
	// компонент пути (/hello) из URL запроса и отправляет его
	// в качестве ответа с помощью Fprintf.
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

// Пример использоания:
// func main() {
// 	// fetch()
// 	http.HandleFunc("/", url_path) // Каждый запрос вызывает обработчик
// 	log.Fatal(http.ListenAndServe("localhost:8000", nil))
// }
/*
Компилируем программу: go build -o execute main.go server.go
Запускаем в фоновом режиме: ./execute &
Делаем запрос с помощью программы fetch:
1) ./fetch http://localhost:8000
HTTP status: 200 OK
URL.Path = "/"
2) ./fetch http://localhost:8000/help
HTTP status: 200 OK
URL.Path = "/help"
3) Можно обратиться через браузер
*/

// Счетчик запросов
func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}

// Вывод фигур лиссажу в клиент HTTP
func handler_lissajous(w http.ResponseWriter, r *http.Request) {
	cycles, err := strconv.Atoi(r.URL.Query().Get("cycles"))
	if err != nil || cycles <= 0 {
		http.NotFound(w, r)
		return
	}
	fmt.Println(cycles)
	lissajous(w, float64(cycles))
}
