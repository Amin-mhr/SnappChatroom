package storage

import "gorm.io/gorm"

type chatRepo struct {
	db *gorm.DB
}
