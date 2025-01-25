package service

import (
	"SnappChatroom/api/handlers/models"
	"SnappChatroom/internal/domain"
	"SnappChatroom/internal/port"
	"SnappChatroom/internal/service"
	"context"
	"log"
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

func (c *ChatService) SendChatToChatRoom(ctx context.Context, chat models.Chat) error {
	err := c.svc.SendChatToChatRoom(ctx, domain.Chat{
		ID:         chat.UserID,
		UserID:     chat.UserID,
		ChatRoomID: chat.ChatRoomID,
		Content:    chat.Content,
		CreatedAt:  chat.CreatedAt,
		UpdatedAt:  chat.UpdatedAt,
		DeletedAt:  chat.DeletedAt,
	})

	if err != nil {
		log.Println("error in sending to chatroom: ", err)
	}

	return err
}
