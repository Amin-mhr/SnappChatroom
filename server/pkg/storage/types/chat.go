package types

import (
	"gorm.io/gorm"
)

type chat struct {
	gorm.Model
	Content string
	UserID  uint
	ChatID  uint
}
