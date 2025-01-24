package domain

import "time"

type Chat struct {
	ID         uint
	UserID     uint
	ChatRoomID uint
	Content    string
	Timestamp  time.Time
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
