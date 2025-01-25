package mapper

import (
	"SnappChatroom/internal/domain"
	"SnappChatroom/pkg/adapters/storage/types"
	"gorm.io/gorm"
)

func ChatDomain2Storage(chatDomain domain.Chat) *types.Chat {
	return &types.Chat{
		Model: gorm.Model{
			CreatedAt: chatDomain.CreatedAt,
			UpdatedAt: chatDomain.UpdatedAt,
			DeletedAt: gorm.DeletedAt(ToNullTime(chatDomain.DeletedAt)),
		},
		Content:    chatDomain.Content,
		UserID:     chatDomain.UserID,
		ChatRoomID: chatDomain.ChatRoomID,
	}
}
