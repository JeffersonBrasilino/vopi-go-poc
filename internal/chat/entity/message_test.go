package entity_test

import (
	"testing"

	"github.com/vopi-go-poc/internal/chat/entity"
)

func TestMessage(t *testing.T) {
	t.Run("should create success", func(t *testing.T) {
		contact, _ := entity.NewContact("contact-uuid")
		sender, _ := entity.NewPerson("person-uuid", "John Doe", "123456789", 1, contact)
		msg, err := entity.NewMessage("msg-uuid", sender, "Hello, world!", "sent")
		if err != nil {
			t.Fatalf("esperava sucesso, mas retornou erro: %v", err)
		}
		if msg == nil {
			t.Fatalf("esperava mensagem criada, mas obteve nil")
		}
	})

	t.Run("should create error", func(t *testing.T) {
		_, err := entity.NewMessage("", nil, "", "")
		if err == nil {
			t.Fatalf("esperava erro, mas retornou sucesso")
		}
	})
}

func TestMessageAttributes(t *testing.T) {
	t.Run("Should get attributes", func(t *testing.T) {
		contact, _ := entity.NewContact("contact-uuid")
		sender, _ := entity.NewPerson("person-uuid", "John Doe", "123456789", 1, contact)
		msg, err := entity.NewMessage("msg-uuid", sender, "Hello, world!", "sent")
		if err != nil {
			t.Fatalf("esperava sucesso, mas retornou erro: %v", err)
		}

		if msg.Uuid() != "msg-uuid" {
			t.Errorf("esperava msg-uuid, mas obteve: %v", msg.Uuid())
		}
		if msg.Sender() != sender {
			t.Errorf("esperava sender igual, mas obteve: %v", msg.Sender())
		}
		if msg.Content() != "Hello, world!" {
			t.Errorf("esperava 'Hello, world!', mas obteve: %v", msg.Content())
		}
		if msg.Status() != "sent" {
			t.Errorf("esperava 'sent', mas obteve: %v", msg.Status())
		}
		if msg.CreatedAt().IsZero() {
			t.Errorf("esperava CreatedAt preenchido, mas está zero")
		}
		if msg.UpdatedAt().IsZero() {
			t.Errorf("esperava UpdatedAt preenchido, mas está zero")
		}
	})
}
