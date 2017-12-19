package main

import (
	"log"

	r "github.com/dancannon/gorethink"
	"github.com/gorilla/websocket"
)

//Client - asdf
type Client struct {
	id      string
	send    chan Message
	socket  *websocket.Conn
	session *r.Session
}

// Read - asdf
func (client *Client) Read() {
	var message Message
	for {
		if err := client.socket.ReadJSON(&message); err != nil {
			break
		}
		go func() {
			err := r.Table("messages").
				Insert(message).
				Exec(client.session)
			if err != nil {
				client.send <- Message{"error", err.Error()}
			}
		}()
	}
	client.socket.Close()
}

// Write - asdf
func (client *Client) Write() {
	for msg := range client.send {
		if err := client.socket.WriteJSON(msg); err != nil {
			break
		}
	}
	client.socket.Close()
}

//NewClient - asdf
func NewClient(socket *websocket.Conn, session *r.Session) *Client {
	var user User
	res, err := r.Table("user").Insert(user).RunWrite(session)
	if err != nil {
		log.Println(err.Error())
	}
	var id string
	if len(res.GeneratedKeys) > 0 {
		id = res.GeneratedKeys[0]
	}
	return &Client{
		send:    make(chan Message),
		socket:  socket,
		session: session,
		id:      id,
	}
}
