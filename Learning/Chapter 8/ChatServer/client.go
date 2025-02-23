package main

import (
	"bufio"
	"fmt"
	"net"
)

func handleConn(conn net.Conn) {
	// Новый канал исходящих сообщений клиента
	var newClient Client
	newClient.ch = make(chan string)
	go clientWriter(conn, newClient.ch)

	who := conn.RemoteAddr().String()
	newClient.name = who
	newClient.ch <- "Вы " + who
	messages <- who + " подключился"
	entering <- newClient
	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- who + ": " + input.Text()
	}
	// Примечание: игнорируем потенциальные ошибки input.Err()
	leaving <- newClient
	messages <- who + " отключился"
	conn.Close()
}

// clientWriter получает широковещательные сообщения
// по исходящему каналу клиента и записывает их в сетевое подключение клиента
func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintf(conn, "%s\n", msg)
	}
}
