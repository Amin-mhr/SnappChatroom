package app

import (
	"SnappChatroom/config"
	"SnappChatroom/internal/port"
	"context"
	"github.com/nats-io/nats.go"
	"gorm.io/gorm"
)

type App interface {
	Config() config.Config
	ChatService(ctx context.Context) port.Chat
	DB() *gorm.DB
	Nats() *nats.Conn
}
