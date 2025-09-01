package createchat_test

import (
	"fmt"
	"testing"

	"github.com/vopi-go-poc/internal/chat/entity"
	"github.com/vopi-go-poc/internal/chat/usecase/createchat"
	"github.com/vopi-go-poc/internal/core/mocks"
)

type mockRepo struct{}

func (m *mockRepo) Create(chat *entity.Chat) error {
	return nil
}
func (m *mockRepo) Exists(channelId string) (bool, error) {
	return false, nil
}

func TestExecute(t *testing.T) {
	t.Parallel()
	t.Run("should create chat successfully", func(t *testing.T) {
		// Arrange
		input := &createchat.CreateChatInput{
			ChannelId:    "123",
			Participants: []*createchat.Person{
				{Name: "Alice", Document: "doc1", Contact: "1111"},
				{Name: "Bob", Document: "doc2", Contact: "1111"},
			},
			Messages:     []*createchat.Message{{Content: "Hello", Sender: "1111", Status: "sent"}},
			BotName:      "TestBot",
		}
		useCase := createchat.NewCreateChat(&mockRepo{})
		res, err := useCase.Execute(t.Context(), input, &mocks.MockTrace{})
		if err != nil {
			t.Error("should create chat successfully")
		}
		if _, ok := res.(*entity.Chat); !ok {
			t.Error("should return a Chat entity")
		}
	})

	t.Run("should create chat error", func(t *testing.T) {
		var cases = []struct {
			caseName string
			input    *createchat.CreateChatInput
			message  string
		}{
			{
				caseName: "empty input",
				message:  "at least one participant is required; at least one message is required",
				input:    &createchat.CreateChatInput{},
			},
			{
				caseName: "invalid sender",
				message:  "messages.0.sender: participant not found",
				input: &createchat.CreateChatInput{
					ChannelId: "123",
					Participants: []*createchat.Person{
						{Name: "Alice", Document: "doc1", Contact: "1111"},
						{Name: "Bob", Document: "doc2", Contact: "1111"},
					},
					Messages: []*createchat.Message{
						{Content: "Hello", Sender: "867867867867867", Status: "sent"},
					},
					BotName: "TestBot",
				},
			},
			{
				caseName: "empty participants",
				message:  "messages.0.sender: participant not found",
				input: &createchat.CreateChatInput{
					ChannelId: "123",
					Participants: nil,
					Messages: []*createchat.Message{
						{Content: "Hello", Sender: "9999", Status: "sent"},
					},
					BotName: "TestBot",
				},
			},
			{
				caseName: "invalid participants",
				message:  "participants.0: name is required; document is required",
				input: &createchat.CreateChatInput{
					ChannelId: "123",
					Participants:  []*createchat.Person{
						{Name: "", Document: "", Contact: ""},
						{Name: "Bob", Document: "doc2", Contact: "1111"},
					},
					Messages: []*createchat.Message{
						{Content: "Hello", Sender: "9999", Status: "sent"},
					},
					BotName: "TestBot",
				},
			},
			{
				caseName: "invalid message",
				message:  "messages.0: content: is required; status: is required",
				input: &createchat.CreateChatInput{
					ChannelId: "123",
					Participants:  []*createchat.Person{
						{Name: "Alice", Document: "doc1", Contact: "1111"},
						{Name: "Bob", Document: "doc2", Contact: "1111"},
					},
					Messages: []*createchat.Message{
						
						{Content: "", Sender: "1111", Status: ""},
					},
					BotName: "TestBot",
				},
			},
		}
		for _, c := range cases {
			t.Run(c.caseName, func(t *testing.T) {
				useCase := createchat.NewCreateChat(&mockRepo{})
				res, err := useCase.Execute(t.Context(), c.input, &mocks.MockTrace{})
				if err == nil {
					t.Error("should return an error")
				}
				fmt.Println("error>>>>>", err.Error())

				if err.Error() != c.message {
					t.Error("should return the correct error message")
				}

				if res != nil {
					t.Error("should return nil")
				}
			})
		}
	})
}
