package core

import (
	"github.com/gorilla/websocket"
	"log"
	"time"
)

func ConnectWebSocket() *websocket.Conn {
	webSocketURL := "ws://localhost:8080/ws"
	conn, _, err := websocket.DefaultDialer.Dial(webSocketURL, nil)
	if err != nil {
		log.Fatalf("Failed to connect to WebSocket: %v", err)
	}
	log.Println("Connected to WebSocket")

	return conn
}

type Chat struct {
	ID         uint
	UserID     uint
	ChatRoomID uint
	Content    string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  time.Time
}
