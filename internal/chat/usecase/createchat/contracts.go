package createchat

import "github.com/vopi-go-poc/internal/chat/entity"

type CreateChatRepository interface {
	CreateChat(chat *entity.Chat) error
	Exists(channelId string) (bool, error)
}
