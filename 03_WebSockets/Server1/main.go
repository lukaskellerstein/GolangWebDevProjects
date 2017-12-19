package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":4000", nil)
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func handler(w http.ResponseWriter, r *http.Request) {
	socket, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		msqType, msg, err := socket.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(msg))
		if err = socket.WriteMessage(msqType, msg); err != nil {
			fmt.Println(err)
			return
		}
	}
}
