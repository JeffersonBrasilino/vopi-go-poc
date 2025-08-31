package database

import (
	"database/sql"

	"github.com/vopi-go-poc/internal/chat/entity"
	"github.com/vopi-go-poc/internal/core"
)

type ChatPostgresRepository struct {
	dbInstance *sql.DB
}

func NewChatPostgresRepository(dbInstance core.DbConnection) *ChatPostgresRepository {
	return &ChatPostgresRepository{
		dbInstance: dbInstance.Instance(),
	}
}

func (r *ChatPostgresRepository) Create(chat *entity.Chat) error {
	// Implementation for creating a chat in PostgreSQL
	return nil
}

func (r *ChatPostgresRepository) Exists(channelId string) (bool, error) {
	// Implementation for checking if a chat exists in PostgreSQL
	return false, nil
}
