package lib

import (
	"net/http"
	"log"
	"github.com/gorilla/websocket"
)

const (
	socketBufferSize = 1024
	messageBufferSize = 256
)

type Room struct {
	// forward is a channel that holds incoming messages
	// that should be forwarded to the other clients.
	forward chan []byte

	// join is a channel for clients wishing to join the room.
	join    chan *client

	// leave is a channel for clients wishing to leave the room.
	leave   chan *client

	// clients holds all current clients in this room.
	clients map[*client]bool
}

// newRoom makes a new room that is ready to go.
func newRoom() *Room {
	return &Room{
		forward: make(chan []byte),
		join:    make(chan *client),
		leave:   make(chan *client),
		clients: make(map[*client]bool),
	}
}

// the select will run ony one thing at time,
// so map r.clients is synchronized!!!! amazing, right?
func (r *Room) run() {
	// run forever
	for {
		select {
		// joining
		case client := <-r.join:
		// its stores the reference in a map ex{2132342:true}
			r.clients[client] = true

		// leaving
		case client := <-r.leave:
		// kick client from map
			delete(r.clients, client)
		// close his channel
			close(client.send)

		// forward/broadcast message to all clients
		case msg := <-r.forward:
			for client := range r.clients {
				select {
				// send the message
				case client.send <- msg:
				// failed to send
				default:
					delete(r.clients, client)
					close(client.send)
				}
			}
		}
	}
}

var upgrader = &websocket.Upgrader{
	ReadBufferSize:socketBufferSize,
	WriteBufferSize: socketBufferSize,
}

// method to use Room as handler
func (r *Room) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// upgrade http => socket
	socket, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Fatal("ServeHTTP:", err)
		return
	}
	client := &client{
		socket: socket,
		send:   make(chan []byte, messageBufferSize),
		room:   r,
	}
	r.join <- client
	defer func() {
		r.leave <- client
	}()
	go client.write()
	client.read()
}
