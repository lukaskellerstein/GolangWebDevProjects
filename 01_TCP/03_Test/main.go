package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()

	request(conn)
	response(conn)
}

func request(conn net.Conn) {
	i := 0
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		if i == 0 {
			method := strings.Fields(ln)[0]
			uri := strings.Fields(ln)[1]
			fmt.Println("METHOD ", method)
			fmt.Println("URI ", uri)
		}
		if ln == "" {
			break
		}
		i++
	}
}

func response(conn net.Conn) {
	body := "<html><head></head><body>asfd</body></html>"
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}
