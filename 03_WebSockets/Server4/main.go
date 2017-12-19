package main

import (
	"fmt"
	"log"
	"net/http"

	r "github.com/dancannon/gorethink"
	"github.com/gorilla/websocket"
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

// User - asdf
type User struct {
	// ID - asdf
	ID string `gorethink:"id,omitempty"`
	// Name - asdf
	Name string `gorethink:"name"`
}

// Message - asdf
type Message struct {
	// Name - asdf
	Name string `json:"name"`
	// Data - asdf
	Data interface{} `json:"data"`
}

// // Channel - asdf
// type Channel struct {
// 	// ID - asdf
// 	ID string `json:"id" gorethink:"id,omitempty"`
// 	// Name - asdf
// 	Name string `json:"name" gorethink:"name"`
// }

var session *r.Session

func main() {
	session2, err := r.Connect(r.ConnectOpts{
		Address:  "localhost:28015",
		Database: "HubDatabase",
	})

	if err != nil {
		log.Panic(err.Error())
	}

	session = session2

	http.HandleFunc("/", handler)
	http.ListenAndServe(":4000", nil)
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func handler(w http.ResponseWriter, req *http.Request) {
	socket, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err.Error())
		return
	}
	client := NewClient(socket, session)
	go client.Write()
	go client.Read()
	subscribeSenzorData(client)

	// socket, err := upgrader.Upgrade(w, r, nil)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// for {

	// 	// var inMessage Message
	// 	// var outMessage Message

	// 	// //Reading messages
	// 	// if err := socket.ReadJSON(&inMessage); err != nil {
	// 	// 	fmt.Println(err)
	// 	// 	break
	// 	// }

	// 	// // TODO
	// 	// // now he add 'switch-case' for instruction - channel add

	// 	// fmt.Println(inMessage.Name)
	// 	// fmt.Println(inMessage.Data)

	// 	// switch inMessage.Name {
	// 	// case "channel add":
	// 	// 	err := addChannel(inMessage.Data)
	// 	// 	if err != nil {
	// 	// 		outMessage = Message{"error", err}
	// 	// 		if err := socket.WriteJSON(outMessage); err != nil {
	// 	// 			fmt.Println(err)
	// 	// 			break
	// 	// 		}
	// 	// 	}
	// 	// case "channel subscribe":
	// 	// 	go subscribeChannel(socket)
	// 	// }

	// }
}

// func addChannel(data interface{}) error {
// 	var channel Channel

// 	err := mapstructure.Decode(data, &channel)
// 	if err != nil {
// 		return err
// 	}
// 	channel.ID = "1"
// 	fmt.Println("added channel")
// 	return nil
// }

// func subscribeChannel(socket *websocket.Conn) {
// 	for {
// 		time.Sleep(time.Second * 1)
// 		message := Message{"channel add", Channel{"1", "Software Support"}}
// 		socket.WriteJSON(message)
// 		fmt.Println("sent new channel")
// 	}
// }

func subscribeSenzorData(client *Client) {
	go func() {
		// stop := client.NewStopChannel(UserStop)
		cursor, err := r.Table("SenzorData").
			Changes(r.ChangesOpts{IncludeInitial: true}).
			Run(client.session)

		if err != nil {
			client.send <- Message{"error", err.Error()}
			return
		}
		// changeFeedHelper(cursor, "user", client.send, stop)

		change := make(chan r.ChangeResponse)
		cursor.Listen(change)
		for {
			eventName := ""
			var data interface{}
			select {
			case val := <-change:
				if val.NewValue != nil && val.OldValue == nil {
					eventName = "message add"
					data = val.NewValue
				} else if val.NewValue == nil && val.OldValue != nil {
					eventName = "message remove"
					data = val.OldValue
				} else if val.NewValue != nil && val.OldValue != nil {
					eventName = "message edit"
					data = val.NewValue
				}
				client.send <- Message{eventName, data}
			}
		}
	}()
}
