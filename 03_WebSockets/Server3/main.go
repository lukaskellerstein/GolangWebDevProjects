package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/mitchellh/mapstructure"
)

/*
CALL IT FROM
http://jsbin.com/dohuxetufe/edit?js,console

with this code inside Javascript part

let message = {
  name: 'channel add',
  data: {
    name: 'Hardware Support'
  }
}

let subMessage = {
  name: 'channel subscribe'
}

let ws = new WebSocket('ws://localhost:4000');

ws.onopen = () => {
  ws.send(JSON.stringify(message));
  ws.send(JSON.stringify(subMessage));
}

ws.onmessage = function(message) {
  console.log(message);
};

ws.onerror = function(evt) {
  console.log(evt);
};

*/

// Message - asdf
type Message struct {
	// Name - asdf
	Name string `json:"name"`
	// Data - asdf
	Data interface{} `json:"data"`
}

// Channel - asdf
type Channel struct {
	// ID - asdf
	ID string `json:"id" gorethink:"id,omitempty"`
	// Name - asdf
	Name string `json:"name" gorethink:"name"`
}

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

		var inMessage Message
		var outMessage Message

		//Reading messages
		if err := socket.ReadJSON(&inMessage); err != nil {
			fmt.Println(err)
			break
		}

		// TODO
		// now he add 'switch-case' for instruction - channel add

		fmt.Println(inMessage.Name)
		fmt.Println(inMessage.Data)

		switch inMessage.Name {
		case "channel add":
			err := addChannel(inMessage.Data)
			if err != nil {
				outMessage = Message{"error", err}
				if err := socket.WriteJSON(outMessage); err != nil {
					fmt.Println(err)
					break
				}
			}
		case "channel subscribe":
			go subscribeChannel(socket)
		}

		// if err = socket.WriteMessage(msgType, msg); err != nil {
		// 	fmt.Println(err)
		// 	return
		// }
	}
}

func addChannel(data interface{}) error {
	var channel Channel

	err := mapstructure.Decode(data, &channel)
	if err != nil {
		return err
	}
	channel.ID = "1"
	fmt.Println("added channel")
	return nil
}

func subscribeChannel(socket *websocket.Conn) {
	for {
		time.Sleep(time.Second * 1)
		message := Message{"channel add", Channel{"1", "Software Support"}}
		socket.WriteJSON(message)
		fmt.Println("sent new channel")
	}
}
