package main

import (
	"io"
	"log"
	"net"
	"os"
)

// netcat - считывает данные из подключения и выводит в stdout
func netcat(adress string) {
	conn, err := net.Dial("tcp", adress)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	mustcopy(os.Stdout, conn)
}

func mustcopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
