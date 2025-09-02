package entity_test

import (
	"testing"
	"github.com/vopi-go-poc/internal/chat/entity"
)

func TestChat(t *testing.T) {
	t.Run("should create success", func(t *testing.T) {
		contact, _ := entity.NewContact("contact-uuid")
		person, _ := entity.NewPerson("person-uuid", "John Doe", "123456789", 1, contact)
		msg, _ := entity.NewMessage("msg-uuid", person, "Hello", "sent")
		participants := []*entity.Person{person}
		messages := []*entity.Message{msg}
		chat, err := entity.NewChat("channel-1", participants, messages)
		if err != nil {
			t.Fatalf("esperava sucesso, mas retornou erro: %v", err)
		}
		if chat == nil {
			t.Fatalf("esperava chat criado, mas obteve nil")
		}
	})

	t.Run("should create error", func(t *testing.T) {
		_, err := entity.NewChat("", nil, nil)
		if err == nil {
			t.Fatalf("esperava erro, mas retornou sucesso")
		}
	})
}

func TestChatAttributes(t *testing.T) {
	t.Run("Should get attributes", func(t *testing.T) {
		contact, _ := entity.NewContact("contact-uuid")
		person, _ := entity.NewPerson("person-uuid", "John Doe", "123456789", 1, contact)
		msg, _ := entity.NewMessage("msg-uuid", person, "Hello", "sent")
		participants := []*entity.Person{person}
		messages := []*entity.Message{msg}
		chat, err := entity.NewChat("channel-1", participants, messages)
		if err != nil {
			t.Fatalf("esperava sucesso, mas retornou erro: %v", err)
		}
		if chat.ChannelId() != "channel-1" {
			t.Errorf("esperava channel-1, mas obteve: %v", chat.ChannelId())
		}
		if chat.Uuid() == "" {
			t.Errorf("esperava uuid preenchido, mas está vazio")
		}
		if len(chat.Participants()) != 1 {
			t.Errorf("esperava 1 participante, mas obteve: %v", len(chat.Participants()))
		}
		if len(chat.Messages()) != 1 {
			t.Errorf("esperava 1 mensagem, mas obteve: %v", len(chat.Messages()))
		}
		if chat.CreatedAt().IsZero() {
			t.Errorf("esperava CreatedAt preenchido, mas está zero")
		}
		if chat.UpdatedAt().IsZero() {
			t.Errorf("esperava UpdatedAt preenchido, mas está zero")
		}
	})
}
