package main

// To test: create web socket request in postman and hit this endpoint ws://localhost:3000/ws
import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/gorilla/websocket"
)

// Define an upgrader to upgrade HTTP connections to WebSocket connections
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	// Allow all origins for simplicity; in production, check the origin
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// WebSocket endpoint handler
func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	// Upgrade the HTTP connection to a WebSocket connection
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error upgrading to WebSocket:", err)
		return
	}
	defer conn.Close()

	// Send initial "hello" message
	if err := conn.WriteMessage(websocket.TextMessage, []byte("hello")); err != nil {
		log.Println("Error writing message:", err)
		return
	}

	// Handle WebSocket messages
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error reading message:", err)
			break
		}
		log.Printf("Received: %s", message)

		// Respond based on the client's message
		var response string
		switch string(message) {
		case "bye":
			response = "goodbye"
		default:
			response = "okay"
		}

		// Send the response back to the client
		if err := conn.WriteMessage(websocket.TextMessage, []byte(response)); err != nil {
			log.Println("Error writing message:", err)
			break
		}

		// Close the connection if the client sends "bye"
		if string(message) == "bye" {
			break
		}
	}
}

func main() {
	r := chi.NewRouter()

	// WebSocket endpoint
	r.Get("/ws", wsEndpoint)

	log.Println("WebSocket server started on :3000")
	log.Fatal(http.ListenAndServe(":3000", r))
}
