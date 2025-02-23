package main

import (
	"bufio"
	"fmt"
	"net"
)

func handleConn(conn net.Conn) {
	// Новый канал исходящих сообщений клиента
	ch := make(chan string)
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	ch <- "Вы " + who + "\n"
	messages <- who + " подключился\n"
	entering <- ch
	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- who + ": " + input.Text() + "\n"
	}
	// Примечание: игнорируем потенциальные ошибки input.Err()
	leaving <- ch
	messages <- who + " отключился\n"
	conn.Close()
}

// clientWriter получает широковещательные сообщения
// по исходящему каналу клиента и записывает их в сетевое подключение клиента
func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintf(conn, "%s\n", msg)
	}
}
