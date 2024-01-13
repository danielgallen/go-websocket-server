package main

import (
	"log"
	"net/http"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
    // Upgrade initial GET request to a WebSocket
    ws, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Fatal(err)
    }
    defer ws.Close()

    // Infinite loop to keep connection open
    for {
        messageType, message, err := ws.ReadMessage()
        if err != nil {
            log.Println("Error reading message:", err)
            break
        }
        log.Printf("Received: %s", message)

        // Echo the message back
        err = ws.WriteMessage(messageType, message)
        if err != nil {
            log.Println("Error writing message:", err)
            break
        }
    }
}

func main() {
    // Set route to handle WebSocket connections
    http.HandleFunc("/echo", handleConnections)

    // Start server on localhost port 8080 and log errors
    log.Println("WebSocket server starting on :8081")
    err := http.ListenAndServe(":8081", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
