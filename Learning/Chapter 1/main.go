package main

/*
* Небольшой сервер, который выводит URL путь,
* число запросов и
* строит фигуры Лиссажу по адресу "http://localhost:8000/"

Запуск :
go run main.go  server.go  lissajous\(1.5\).go
*/
import (
	"log"
	"net/http"
)

func main() {
	// Каждый запрос вызывает обработчик
	http.HandleFunc("/", url_path)
	http.HandleFunc("/count", counter)
	http.HandleFunc("/lissajous", handler_lissajous)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
