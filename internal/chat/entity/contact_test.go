package entity_test

import (
	"testing"

	"github.com/vopi-go-poc/internal/chat/entity"
)

func TestContact(t *testing.T) {
	t.Run("should create success", func(t *testing.T) {
		c, err := entity.NewContact("contact-uuid")
		if err != nil {
			t.Fatalf("esperava sucesso, mas retornou erro: %v", err)
		}
		if c == nil {
			t.Fatalf("esperava contato criado, mas obteve nil")
		}
	})

	t.Run("should get contact attribute", func(t *testing.T) {
		c, _ := entity.NewContact("contact-uuid")
		if c.Contact() != "contact-uuid" {
			t.Errorf("esperava contact-uuid, mas obteve: %v", c.Contact())
		}
	})
}
