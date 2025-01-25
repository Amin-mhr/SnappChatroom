package main

import (
	"SnappChatRoomClient/core"
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/websocket"
	"github.com/nats-io/nats.go"
)

func main() {
	var chatRooms map[string]bool = make(map[string]bool)
	natsConn := core.ConnectToNats()
	defer natsConn.Close()

	wsConn := core.ConnectWebSocket()
	defer wsConn.Close()

	reader := bufio.NewReader(os.Stdin)
	go func() {
		for {
			fmt.Print("Enter your message (format: ChatRoomID:Message): ")
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)

			// Ensure input format is valid
			parts := strings.SplitN(input, ":", 2)
			if len(parts) != 2 {
				fmt.Println("Invalid input format. Use ChatRoomID:Message")
				continue
			}

			chatRoomID := parts[0]
			message := parts[1]

			if exist, _ := chatRooms[chatRoomID]; !exist {
				chatRooms[chatRoomID] = true
				go subscribeNats(chatRoomID, natsConn)
			}

			chatRoomIDUint, err := strconv.ParseUint(chatRoomID, 10, 32)
			var req = core.Chat{
				ID:         0,
				UserID:     1,
				ChatRoomID: uint(chatRoomIDUint),
				Content:    message,
				CreatedAt:  time.Now(),
				UpdatedAt:  time.Time{},
				DeletedAt:  time.Time{},
			}
			// Send message to WebSocket
			reqJson, err := json.Marshal(req)
			if err != nil {
				log.Println(err)
			}
			err = wsConn.WriteMessage(websocket.TextMessage, reqJson)
			if err != nil {
				log.Printf("Failed to send message via WebSocket: %v", err)
				break
			}

			//log.Printf("[WebSocket][%s] Sent message: %s", chatRoomID, message)
		}
	}()

	// Step 4: Subscribe to NATS for all ChatRooms
	//var wg sync.WaitGroup
	//wg.Add(1)

	//go func() {
	//	ticker := time.NewTicker(5 * time.Second) // ایجاد یک Ticker برای هر ۶۰ ثانیه
	//	defer ticker.Stop()
	//
	//	for {
	//		select {
	//		case <-ticker.C: // هر بار که ۶۰ ثانیه می‌گذرد
	//			for _, chatRoomID := range chatRooms {
	//				// استفاده از یک متغیر محلی برای جلوگیری از تغییر مقدار
	//				roomID := chatRoomID
	//				subject := fmt.Sprintf("%s", roomID)
	//				_, err := natsConn.Subscribe(subject, func(msg *nats.Msg) {
	//					fmt.Printf("[NATS][%s] Message received: %s\n", roomID, string(msg.Data))
	//				})
	//				if err != nil {
	//					log.Printf("Failed to subscribe to NATS channel %s: %v", subject, err)
	//				} else {
	//					log.Printf("Subscribed to NATS channel: %s", subject)
	//				}
	//			}
	//
	//			// در صورتی که نیاز باشد گوروتین متوقف شود، می‌توان به یک کانال کنترل خاص گوش داد
	//		}
	//	}
	//}()

	// Wait for the NATS listener goroutine to complete (it won't, but this ensures proper cleanup)
	//wg.Wait()
	select {}
}

func subscribeNats(chatRoomID string, natsConn *nats.Conn) {
	subject := fmt.Sprintf("%s", chatRoomID)
	_, err := natsConn.Subscribe(subject, func(msg *nats.Msg) {
		fmt.Printf("[NATS][%s] Message received: %s\n", chatRoomID, string(msg.Data))
	})
	if err != nil {
		log.Printf("Failed to subscribe to NATS channel %s: %v", subject, err)
	} else {
		log.Printf("Subscribed to NATS channel: %s", subject)
	}
}
