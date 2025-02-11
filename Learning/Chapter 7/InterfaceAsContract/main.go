package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", sortHandler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// type str []byte

// func (s str) Len() int           { return len(s) }
// func (s str) Less(i, j int) bool { return s[i] < s[j] }
// func (s str) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

// func main() {
// 	// http.HandleFunc("/", sortHandler)
// 	// log.Fatal(http.ListenAndServe("localhost:8000", nil))
// 	fmt.Println(isPalindrom(str([]byte{'a', 'b', 'a', 'a'})))
// }
