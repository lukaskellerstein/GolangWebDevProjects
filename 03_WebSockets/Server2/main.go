package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

/*
CALL IT FROM
http://jsbin.com/dohuxetufe/edit?js,console

with this code inside Javascript part

let message = {
  topic: 'someMessageName',
  data: {
    text: 'someMessageText'
  }
}

let ws = new WebSocket('ws://localhost:4000/topic1');

ws.onopen = () => {
  ws.send("HELLO2");
}

ws.onmessage = function(message) {
  console.log(message.data);
};

ws.onerror = function(evt) {
  console.log(evt);
}
*/

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
		// msqType, msg, err := socket.ReadMessage()
		// if err != nil {
		// 	fmt.Println(err)
		// 	return
		// }

		var inMessage Message
		if err := socket.ReadJSON(&inMessage); err != nil {
			fmt.Println(err)
			break
		}

		// TODO
		// now he add 'switch-case' for instruction - channel add

		fmt.Println(inMessage.Name)
		fmt.Println(inMessage.Data)

		// fmt.Println(string(msg))
		// if err = socket.WriteMessage(msqType, msg); err != nil {
		// 	fmt.Println(err)
		// 	return
		// }
	}
}
