package websocket

import (
	"SnappChatroom/api/service"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

func wsHandler(svcGetter ServiceGetter[*service.ChatService]) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var upgrader = websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		}

		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println("Upgrade error:", err)
			http.Error(w, "Failed to upgrade connection", http.StatusInternalServerError)
			return
		}
		defer conn.Close()

		log.Println("Client connected")

		// Read messages from the client.
		for {
			messageType, message, err := conn.ReadMessage()
			if err != nil {
				log.Println("Read error:", err)
				break
			}

			log.Printf("Received: %s", message)

			// Echo the message back to the client.
			if err := conn.WriteMessage(messageType, message); err != nil {
				log.Println("Write error:", err)
				break
			}
		}

		log.Println("Client disconnected")
	}
}
