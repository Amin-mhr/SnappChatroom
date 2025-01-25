package websocket

import (
	"SnappChatroom/app"
	"SnappChatroom/config"
	"fmt"
	"net/http"
)

func Run(appContainer app.App, cfg config.ServerConfig) error {
	chatServiceGetter := chatServiceGetter(appContainer)
	http.HandleFunc("/ws", wsHandler(chatServiceGetter))
	fmt.Println("WebSocket server started on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
		return err
	}
	return nil
}
