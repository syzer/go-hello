package lib

import (
	"github.com/gorilla/websocket"
	"github.com/stretchr/objx"
	"github.com/syzer/go-hello/webServer/lib/trace"
	"log"
	"net/http"
)

const (
	socketBufferSize  = 1024
	messageBufferSize = 256
)

type Room struct {
	// forward is a channel that holds incoming messages
	// that should be forwarded to the other clients.
	forward chan *message

	// join is a channel for clients wishing to join the room.
	join chan *client

	// leave is a channel for clients wishing to leave the room.
	leave chan *client

	// clients holds all current clients in this room.
	clients map[*client]bool

	// tracer will receive trace information of activity
	// in the room.
	tracer trace.Tracer
}

// newRoom makes a new room that is ready to go.
// AKA factory
func newRoom() *Room {
	return &Room{
		forward: make(chan *message),
		join:    make(chan *client),
		leave:   make(chan *client),
		clients: make(map[*client]bool),
		tracer:  trace.Off(),
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
			r.tracer.Trace("New client joined")

		// leaving
		case client := <-r.leave:
			// kick client from map
			delete(r.clients, client)
			// close his channel
			close(client.send)
			r.tracer.Trace("Client left")

		// forward/broadcast message to all clients
		case msg := <-r.forward:
			for client := range r.clients {
				select {
				// send the message
				case client.send <- msg:
					r.tracer.Trace(" -- sent to client")
				// failed to send
				default:
					delete(r.clients, client)
					close(client.send)
					r.tracer.Trace(" -- failed to send, cleaned up client")
				}
			}
		}
	}
}

var upgrader = &websocket.Upgrader{
	ReadBufferSize:  socketBufferSize,
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

	authCookie, err := req.Cookie("auth")
	if err != nil {
		log.Fatal("Failed to get auth cookie:", err)
		return
	}

	client := &client{
		socket:   socket,
		send:     make(chan *message, messageBufferSize),
		room:     r,
		userData: objx.MustFromBase64(authCookie.Value),
	}
	r.join <- client
	defer func() {
		r.leave <- client
	}()
	go client.write()
	client.read()
}
