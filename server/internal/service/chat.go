package service

import (
	"SnappChatroom/internal/domain"
	"SnappChatroom/internal/port"
	"context"
	"errors"
)

var (
	ErrChatCreation = errors.New("error creating chat")
)

type chatService struct {
	repo port.ChatRepo
}

func NewChatService(chatRepo port.ChatRepo) port.Chat {
	return &chatService{
		repo: chatRepo,
	}
}

func (c *chatService) SendChatToChatRoom(ctx context.Context, chat domain.Chat) error {
	//should implement with using repo for working with database.
	panic("implement me")
}
