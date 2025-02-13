package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
	"text/template"
)

type dollars float32
type database map[string]dollars

var mu sync.Mutex

func (d dollars) String() string {
	return fmt.Sprintf("%.2f$", d)
}

func (db database) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/list":
		for item, price := range db {
			fmt.Fprintf(w, "%s: %s\n", item, price)
		}
	case "/price":
		item := r.URL.Query().Get("item")
		price, ok := db[item]
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "Товар %q не существует!\n", item)
			return
		}
		fmt.Fprintf(w, "%s\n", price)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Страница %s не существует!\n", r.URL)
	}
}

func (db database) list(w http.ResponseWriter, r *http.Request) {
	resp := template.Must(template.ParseFiles("./pricesTemp.html"))
	err := resp.Execute(w, db)
	if err != nil {
		fmt.Printf("db.list: %v", err)
		return
	}
}

func (db database) price(w http.ResponseWriter, r *http.Request) {
	item := r.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Товар %q не существует!\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}

func (db database) update(w http.ResponseWriter, r *http.Request) {
	// update?item=shoes&price=25
	item := r.URL.Query().Get("item")
	price := r.URL.Query().Get("price")
	_, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Товар %q не существует!\n", item)
		return
	}
	newPrice, err := (strconv.ParseFloat(price, 32))
	if err != nil {
		fmt.Fprintf(w, "Ошибка преобразования цены! %v\n", err)
	}
	db[item] = dollars(newPrice)
	db.list(w, r)
}

func (db database) add(w http.ResponseWriter, r *http.Request) {
	// add?item=car&price=1000
	item := r.URL.Query().Get("item")
	price := r.URL.Query().Get("price")
	_, ok := db[item]
	if ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Товар %q уже существует!\n", item)
		return
	}
	newPrice, err := (strconv.ParseFloat(price, 32))
	if err != nil {
		fmt.Fprintf(w, "Ошибка преобразования цены! %v\n", err)
	}
	mu.Lock()
	db[item] = dollars(newPrice)
	mu.Unlock()
	db.list(w, r)
}

func (db database) delete(w http.ResponseWriter, r *http.Request) {
	// delete?item=shoes
	item := r.URL.Query().Get("item")
	_, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Товар %q не существует!\n", item)
		return
	}
	mu.Lock()
	delete(db, item)
	mu.Unlock()
	db.list(w, r)
}

func (db database) get(w http.ResponseWriter, r *http.Request) {
	// get?item=shoes
	item := r.URL.Query().Get("item")
	_, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Товар %q не существует!\n", item)
		return
	}
	fmt.Fprintf(w, "%s: %.2f$\n", item, db[item])
}

func main() {
	db := database{"shoes": 50, "socks": 5}
	mux := http.NewServeMux()
	mux.Handle("/list", http.HandlerFunc(db.list))
	mux.HandleFunc("/price", db.price)
	mux.HandleFunc("/update", db.update)
	mux.HandleFunc("/add", db.add)
	mux.HandleFunc("/delete", db.delete)
	mux.HandleFunc("/get", db.get)
	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}
