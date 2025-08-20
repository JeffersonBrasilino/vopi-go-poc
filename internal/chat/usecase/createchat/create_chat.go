package createchat

import (
	"fmt"

	"github.com/vopi-go-poc/internal/chat/entity"
)

type CreateChat struct {
	database CreateChatRepository
}

func NewCreateChat(
	database CreateChatRepository,
) *CreateChat {
	return &CreateChat{
		database: database,
	}
}

func (r *CreateChat) Execute(data *CreateChatInput) (any, error) {
	fmt.Println("CreateChat use case executed")

	person, errP := entity.NewPerson(
		data.Participants[0].Contact,
		data.Participants[0].Name,
	)
	fmt.Println("Created person:", person, errP)
	if errP != nil {
		return nil, fmt.Errorf("failed to create person: %v", errP)
	}
	return nil, nil
}
