package core

import (
	"github.com/nats-io/nats.go"
	"log"
)

func ConnectToNats() *nats.Conn {
	natsURL := "nats://localhost:4222"
	nc, err := nats.Connect(natsURL)
	if err != nil {
		log.Fatalf("Failed to connect to NATS: %v", err)
	}
	log.Println("Connected to NATS")
	return nc
}
