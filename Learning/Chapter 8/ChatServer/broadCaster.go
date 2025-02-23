package main

// Канал исходящих сообщений
type client chan<- string

type Client struct {
	// Имя клиента
	name string
	// Канал исходящих сообщений
	ch chan<- string
}

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string)
)

func broadcaster() {
	// Все подключенные клиенты
	clients := make(map[client]bool)
	for {
		select {
		case msg := <-messages:
			// Широковещательное входящее сообщение
			// во все каналы исходящих сообщений для клиентов
			for cli := range clients {
				cli <- msg
			}
		case cli := <-entering:
			clients[cli] = true
		case cli := <-leaving:
			delete(clients, cli)
			close(cli)
		}
	}
}
