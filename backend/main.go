package main

import (
	"fmt"
	"net/http"

	"github.com/TutorialEdge/realtime-chat-go-react/pkg/websocket"
)

// Define our WebSocket endpoint
func serveWs(pool *websocket.Pool, w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Host)

	// Upgrade this connection to a WebSocket
	// connection
	ws, err := websocket.Upgrade(w, r)
	if err != nil {
		fmt.Fprintf(w, "%+V\n", err)
	}

	// Create a new client
	client := &websocket.Client{
		Conn: ws,
		Pool: pool,
	}

	// Register the client
	pool.Register <- client
	client.Read()
}

// setupRoutes sets up the routes for the server
func setupRoutes() {
	// Create a new WebSocket pool
	pool := websocket.NewPool()
	// Start the pool
	go pool.Start()
	// Handle the WebSocket endpoint
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(pool, w, r)
	})
}

func main() {
	fmt.Println("Distributed Chat App v0.01")
	setupRoutes()
	// Start the server on port 8080
	http.ListenAndServe(":8080", nil)
}
