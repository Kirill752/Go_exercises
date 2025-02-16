package main

// Праллельный сервер часов
// Данная программа читает ответы различных серверов часов (см. clock.go) с разных портов и выводит результат в виде таблицы
// go run clockwall.go netcat.go Moscow=localhost:8000 NewYork=localhost:8010 Tokyo=localhost:8020

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"text/tabwriter"
	"time"
)

type server struct {
	name    string
	host    string
	massage string
}

func (s *server) String() string {
	return fmt.Sprintf("%s:%s", s.name, s.host)
}

func parseServers(args []string) ([]*server, error) {
	var servers = make([]*server, 0, len(args))
	for _, arg := range args {
		params := strings.SplitN(arg, "=", 2)
		if len(params) != 2 {
			return nil, fmt.Errorf("wrong Format : %s", arg)
		}
		servers = append(servers, &server{name: params[0], host: params[1]})
	}
	return servers, nil
}

func (s *server) getMassages(conn io.ReadCloser) error {
	defer conn.Close()
	for {
		_, err := fmt.Fscanf(conn, "%s", &s.massage)
		if err != nil {
			return err
		}
	}
}

func (s *server) Connect() (io.ReadCloser, error) {
	conn, err := net.Dial("tcp", s.host)
	if err != nil {
		return nil, err
	}
	fmt.Printf("Connected to %q\n", s.host)
	return conn, nil
}

func printMassges(connections []*server) {
	const format = "%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 20, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Name", "Server", "Message")
	fmt.Fprintf(tw, format, "----", "------", "-------")
	tw.Flush()
	fmt.Print("\033[s")
	for {
		for _, s := range connections {
			fmt.Fprintf(tw, format, s.name, s.host, s.massage)
		}
		fmt.Fprintf(tw, format, "----", "------", "-------")
		fmt.Fprint(tw, "\033[u")
		tw.Flush()
		time.Sleep(1 * time.Second)
	}
}

func main() {
	servers, err := parseServers(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	for _, s := range servers {
		c, err := s.Connect()
		if err != nil {
			log.Fatal(err)
		}
		go func(c io.ReadCloser, s *server) {
			if err := s.getMassages(c); err != nil {
				log.Fatal(err)
			}
		}(c, s)
	}
	printMassges(servers)
}
