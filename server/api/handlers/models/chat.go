package models

import "time"

type Chat struct {
	ID         uint
	UserID     uint
	ChatRoomID uint
	Content    string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  time.Time
}

type User struct {
	UserID     uint
	Username   string
	ChatRoomID map[uint]*ChatRoom //key is ChatRoomID
}

type ChatRoom struct {
	ChatRoomID uint
	UserID     map[uint]*User //key is UserID
}
