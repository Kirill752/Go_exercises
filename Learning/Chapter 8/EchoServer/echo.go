package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

var wg sync.WaitGroup

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func handleConn(c net.Conn) {
	input := bufio.NewScanner(c)
	// Канал передачи сообщений
	text := make(chan string)
	defer func() {
		wg.Wait()
		c.Close()
		fmt.Println("You have been disconnected from the server")
	}()
	// Отдельная go подпрограмма, считывающая сообщение из подключения
	go func() {
		for input.Scan() {
			text <- input.Text()
		}
	}()
	for {
		select {
		// Если получаем сообщение, то обрабатываем его
		case mes := <-text:
			wg.Add(1)
			go func() {
				defer wg.Done()
				echo(c, mes, 1*time.Second)
			}()
		// Если сообщение не приходит в течение 10 секунд, срабатывает этот кейс
		case <-time.After(10 * time.Second):
			return
		}
	}
}

func main() {
	listenner, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listenner.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		// Обработка подключения
		go handleConn(conn)
	}
}
