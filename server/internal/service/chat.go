package service

import (
	"SnappChatroom/internal/domain"
	"SnappChatroom/internal/port"
	"SnappChatroom/pkg/adapters/nats"
	"context"
	"errors"
	"log"
)

var (
	ErrChatCreation     = errors.New("error creating chat")
	ErrInSendingMessage = errors.New("error in sending message")
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
	err := c.repo.Create(ctx, chat)
	if err != nil {
		log.Println(ErrChatCreation)
		return err
	}
	//working nats
	err = nats.PublishToChannel(chat)
	if err != nil {
		log.Println(ErrInSendingMessage)
		return err
	}
	return nil
}
