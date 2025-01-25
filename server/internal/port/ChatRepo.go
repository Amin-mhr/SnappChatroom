package port

import (
	"SnappChatroom/internal/domain"
	"context"
)

type ChatRepo interface {
	Create(ctx context.Context, chat domain.Chat) error
}
