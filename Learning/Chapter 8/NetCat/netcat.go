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
	// Копирование ответа от сервера на стандартный вывод
	mustcopy(os.Stdout, conn)
}

func netcat2(adress string) {
	conn, err := net.Dial("tcp", adress)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	// Копирование ответа от сервера на стандартный вывод
	go mustcopy(os.Stdout, conn)
	// Считывание стандартного ввода и отправление его на сервер
	mustcopy(conn, os.Stdin)
}

func netcat3(adress string) {
	conn, err := net.Dial("tcp", adress)
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan struct{})
	TCPconn := conn.(*net.TCPConn)
	go func() {
		io.Copy(os.Stdout, conn) // Игнорируем ошибки
		log.Println(done)
		done <- struct{}{} // Сигнал главной go подпрограмме
	}()
	// Считывание стандартного ввода и отправление его на сервер
	mustcopy(conn, os.Stdin)
	// Закрытие TCP соединения на запись
	TCPconn.CloseWrite()
	<-done
	// Ожидание завершения фоновой go подпрограммы
	// Закрытие TCP соединения на чтение
	TCPconn.CloseRead()
}

func mustcopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
