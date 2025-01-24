package service

import (
	"SnappChatroom/internal/domain"
	"SnappChatroom/internal/port"
	"SnappChatroom/internal/service"
	"context"
)

type ChatService struct {
	svc port.Chat
}

func NewChatService(svc port.Chat) *ChatService {
	return &ChatService{svc: svc}
}

var (
	ErrChatCreation = service.ErrChatCreation
)

func (c *ChatService) SendChatToChatRoom(ctx context.Context, chat domain.Chat) error {

}
