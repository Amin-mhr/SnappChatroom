package storage

import (
	"SnappChatroom/internal/domain"
	"SnappChatroom/pkg/adapters/storage/mapper"
	"context"
	"gorm.io/gorm"
	"log"
)

type chatRepo struct {
	db *gorm.DB
}

func NewChatRepo(db *gorm.DB) *chatRepo {
	return &chatRepo{db: db}
}

func (c *chatRepo) Create(ctx context.Context, chatDomain domain.Chat) error {
	chat := mapper.ChatDomain2Storage(chatDomain)
	err := c.db.Table("chats").WithContext(ctx).Create(chat).Error
	if err != nil {
		log.Println("create chat error:", err)
		return err
	}
	return nil
}
