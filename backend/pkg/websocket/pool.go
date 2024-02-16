package websocket

import "fmt"

// Pool is a collection of clients
type Pool struct {
	Register   chan *Client     // Register channel will send out a message to all the clients within this pool when a new client connects
	Unregister chan *Client     // Unregister channel will send out a message to all the clients within this pool when a client disconnects
	Clients    map[*Client]bool // Clients map will keep track of all the clients within this pool and whether or not they are active/inactive
	Broadcast  chan Message     // a channel which, when it is passed a message, will loop through all clients in the pool and send the message through the socket connection.
}

// NewPool creates a new pool
func NewPool() *Pool {
	return &Pool{
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),
		Broadcast:  make(chan Message),
	}
}

// Start starts an infinite loop to listen for new clients, disconnecting clients, and messages to broadcast
func (pool *Pool) Start() {
	for {
		select {
		case client := <-pool.Register:
			// When a new client connects, add it to the pool
			pool.Clients[client] = true
			fmt.Println("Size of Connection Pool: ", len(pool.Clients))
			// Notify all clients that a new client has joined
			for client := range pool.Clients {
				fmt.Println(client)
				client.Conn.WriteJSON(Message{Type: 1, Body: "New User Joined..."})
			}
			break
		case client := <-pool.Unregister:
			// When a client disconnects, remove it from the pool
			delete(pool.Clients, client)
			fmt.Println("Size of Connection Pool: ", len(pool.Clients))
			// Notify all clients that a client has left
			for client := range pool.Clients {
				client.Conn.WriteJSON(Message{Type: 1, Body: "User Disconnected..."})
			}
			break
		case message := <-pool.Broadcast:
			// When a message is received, send it to all clients
			fmt.Println("Sending message to all clients in Pool")
			for client := range pool.Clients {
				if err := client.Conn.WriteJSON(message); err != nil {
					fmt.Println(err)
					return
				}
			}
		}
	}
}
