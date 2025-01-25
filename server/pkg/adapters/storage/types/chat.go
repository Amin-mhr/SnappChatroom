package types

import (
	"gorm.io/gorm"
)

type Chat struct {
	gorm.Model
	Content    string
	UserID     uint
	ChatRoomID uint
}
