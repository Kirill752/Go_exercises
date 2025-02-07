package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

// Fetch загружает URL и возвращает имя и длину локального файла.
func fetch(url string) (filename string, n int64, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()
	// Получаем имя файла из последнего компонента аути URL
	local := path.Base(resp.Request.URL.Path)
	if local == "/" || local == "" {
		local = "index.html"
	}
	f, err := os.Create(local)
	if err != nil {
		return "", 0, err
	}
	n, err = io.Copy(f, resp.Body)
	// Закрытие файла; если есть ошибка Сору, возвращаем ее.
	defer func() {
		celoseErr := f.Close()
		if err == nil {
			err = celoseErr
		}
	}()
	return local, n, err
}

func main() {
	for _, u := range os.Args[1:] {
		local, n, err := fetch(u)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch %s: %v\n", u, err)
			continue
		}
		fmt.Fprintf(os.Stdout, "%s => %s (%d bytes).\n", u, local, n)
	}
}
