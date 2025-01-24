package port

import (
	"SnappChatroom/internal/domain"
	"context"
)

type Chat interface {
	SendChatToChatRoom(ctx context.Context, chat domain.Chat) error
}
