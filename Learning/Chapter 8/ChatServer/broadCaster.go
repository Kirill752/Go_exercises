package main

type Client struct {
	// Имя клиента
	name string
	// Канал исходящих сообщений
	ch chan string
}

var (
	entering = make(chan Client)
	leaving  = make(chan Client)
	messages = make(chan string)
)

func broadcaster() {
	// Все подключенные клиенты
	clients := make(map[Client]bool)
	for {
		select {
		case msg := <-messages:
			// Широковещательное входящее сообщение
			// во все каналы исходящих сообщений для клиентов
			for cli := range clients {
				cli.ch <- msg
			}
		case cli := <-entering:
			if len(clients) != 0 {
				cli.ch <- "На сервере уже:"
				for prevcli := range clients {
					cli.ch <- prevcli.name
				}
			} else {
				cli.ch <- "Вы единственный человек на сервере"
			}
			clients[cli] = true
		case cli := <-leaving:
			delete(clients, cli)
			close(cli.ch)
		}
	}
}
