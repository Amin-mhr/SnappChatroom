package postgres

import (
	"SnappChatroom/pkg/adapters/storage/types"
	"fmt"
	"gorm.io/gorm"
	"log"
)

func Migrate(db *gorm.DB) error {

	if err := db.AutoMigrate(
		&types.Chat{},
	); err != nil {
		return fmt.Errorf("failed to migrate database: %w", err)
	}

	log.Println("Database migration completed successfully.")
	return nil
}
