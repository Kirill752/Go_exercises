package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func clock(port string) {
	listnener, err := net.Listen("tcp", fmt.Sprintf("localhost:%s", port))
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
	z, _ := time.Now().Local().Zone()
	for {
		_, err := io.WriteString(c, fmt.Sprintf("\r%s:%s", z, time.Now().Format("15:04:05")))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}
