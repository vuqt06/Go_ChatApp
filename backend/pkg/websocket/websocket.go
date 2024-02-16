package websocket

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

// Define our upgrader with Read and Write buffer sizes
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,

	// Check the origin of the connection,
	//	allowing us to make requests from our React development server
	CheckOrigin: func(r *http.Request) bool { return true },
}

// Upgrade upgrades the HTTP server connection to the WebSocket protocol
func Upgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return ws, err
	}

	return ws, nil
}

// Define a reader which will listen for new messages being sent to our WebSocket endpoint
// func Reader(conn *websocket.Conn) {
// 	for {
// 		// Read in a message
// 		messageType, p, err := conn.ReadMessage()
// 		if err != nil {
// 			fmt.Println(err)
// 			return
// 		}

// 		// Print out that message for clarity
// 		fmt.Println(string(p))

// 		if err := conn.WriteMessage(messageType, p); err != nil {
// 			fmt.Println(err)
// 			return
// 		}
// 	}
// }

// Define a writer which will write a message back to the client
// func Writer(conn *websocket.Conn) {
// 	for {
// 		fmt.Println("Sending")
// 		messageType, r, err := conn.NextReader()
// 		if err != nil {
// 			fmt.Println(err)
// 			return
// 		}

// 		w, err := conn.NextWriter(messageType)
// 		if err != nil {
// 			fmt.Println(err)
// 			return
// 		}

// 		if _, err := io.Copy(w, r); err != nil {
// 			fmt.Println(err)
// 			return
// 		}

// 		if err := w.Close(); err != nil {
// 			fmt.Println(err)
// 			return
// 		}
// 	}
// }
