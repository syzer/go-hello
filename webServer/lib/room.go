package lib

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

func (r *Room) run() {
	for {
		select {
		// joining
		case client := <-r.join:
			r.clients[client] = true
		// leaving
		case client := <-r.leave:

			delete(r.clients, client)
			close(client.send)
		// forward message to all clients
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
