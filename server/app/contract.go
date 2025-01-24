package app

import (
	"SnappChatroom/config"
	"SnappChatroom/internal/port"
	"context"
)

type App interface {
	Config() config.Config
	ChatService(ctx context.Context) port.Chat
}
