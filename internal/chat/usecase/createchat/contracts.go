package createchat

import "github.com/vopi-go-poc/internal/chat/entity"

type CreateChatRepository interface {
	Create(chat *entity.Chat) error
	Exists(channelId string) (bool, error)
}
