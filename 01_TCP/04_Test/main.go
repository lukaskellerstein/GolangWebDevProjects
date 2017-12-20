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

	mux(conn)
}

func mux(conn net.Conn) {
	i := 0
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		line := scanner.Text()
		if i == 0 {
			method := strings.Fields(line)[0]
			uri := strings.Fields(line)[1]
			fmt.Println("METHOD ", method)
			fmt.Println("URI ", uri)

			if method == "GET" && uri == "/" {
				response1(conn)
			}
			if method == "GET" && uri == "/about" {
				response2(conn)
			}

		}
		if line == "" {
			break
		}
		i++
	}
}

func response1(conn net.Conn) {
	body := "<html><head></head><body>TEXT11</body></html>"

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/hmtl\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

func response2(conn net.Conn) {
	body := "<html><head></head><body>TEXT22</body></html>"

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/hmtl\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}
