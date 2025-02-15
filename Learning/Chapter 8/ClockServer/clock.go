package main

import (
	"io"
	"log"
	"net"
	"time"
)

func main() {
	listnener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listnener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		// Обработка единственного подключения
		go handlecon(conn)
	}
}

func handlecon(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("\r15:04:05"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}
