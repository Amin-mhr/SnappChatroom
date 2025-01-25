package nats

import (
	"SnappChatroom/internal/domain"
	"fmt"
	"github.com/nats-io/nats.go"
	"log"
	"strconv"
)

func PublishToChannel(chat domain.Chat) error {
	conn, err := ConnectToNats()
	if err != nil {
		return fmt.Errorf("nats connection is corrupted")
		return err
	}

	channel := strconv.Itoa(int(chat.ChatRoomID))
	err = conn.Publish(channel, []byte(chat.Content))
	if err != nil {
		log.Printf("Failed to publish to channel %s: %v", channel, err)
		return err
	}

	log.Printf("Message published to channel %s: %s", channel, chat.Content)
	return nil
}

func ConnectToNats() (*nats.Conn, error) {
	natsURL := "nats://localhost:4222"
	conn, err := nats.Connect(natsURL)
	if err != nil {
		log.Printf("Failed to connect to nats server: %v", err)
		return nil, err
	}
	log.Println("Connected to nats server:", natsURL)
	return conn, nil
}
