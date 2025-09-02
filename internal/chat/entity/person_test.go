package entity_test

import (
	"testing"

	"github.com/vopi-go-poc/internal/chat/entity"
)

func TestPerson(t *testing.T) {
	t.Run("should create success", func(t *testing.T) {
		contact, _ := entity.NewContact("3r43r43r43")
		_, err := entity.NewPerson("uuid-1", "John Doe", "123456789", 1, contact)
		if err != nil {
			t.Fatalf("esperava sucesso, mas retornou erro: %v", err)
		}
	})

	t.Run("should create error", func(t *testing.T) {
		_, err := entity.NewPerson("", "", "", 0, nil)
		if err == nil {
			t.Fatalf("esperava erro, mas retornou sucesso")
		}

		if err.Error() != "uuid is required; name is required; document is required; personType is required; at least one contact is required" {
			t.Fatalf("esperava mensagem de erro específica, mas obteve: %v", err)
		}
	})
}

func TestPersonAttributes(t *testing.T) {
	t.Run("Should get attributes", func(t *testing.T) {
		contact, _ := entity.NewContact("3r43r43r43")
		p, err := entity.NewPerson("uuid-1", "John Doe", "123456789", 1, contact)
		if err != nil {
			t.Fatalf("esperava sucesso, mas retornou erro: %v", err)
		}

		if p.Uuid() != "uuid-1" {
			t.Errorf("esperava uuid-1, mas obteve: %v", p.Uuid())
		}
		if p.Name() != "John Doe" {
			t.Errorf("esperava John Doe, mas obteve: %v", p.Name())
		}
		if p.Document() != "123456789" {
			t.Errorf("esperava 123456789, mas obteve: %v", p.Document())
		}
		if p.PersonType() != 1 {
			t.Errorf("esperava 1, mas obteve: %v", p.PersonType())
		}
		if p.Contacts() == nil {
			t.Errorf("esperava pelo menos um contato, mas obteve: %v", p.Contacts())
		}
	})
}

/* func TestNewPerson_ValidationError(t *testing.T) {
	tests := []struct {
		uuid       string
		name       string
		document   string
		personType int
		contacts   *Contact
		desc       string
	}{
		{"", "John", "123", 1, &Contact{}, "uuid vazio"},
		{"uuid", "", "123", 1, &Contact{}, "name vazio"},
		{"uuid", "John", "", 1, &Contact{}, "document vazio"},
		{"uuid", "John", "123", 0, &Contact{}, "personType zero"},
		{"uuid", "John", "123", 1, nil, "contacts nil"},
	}
	for _, tt := range tests {
		_, err := NewPerson(tt.uuid, tt.name, tt.document, tt.personType, tt.contacts)
		if err == nil {
			t.Errorf("esperava erro de validação para caso: %s", tt.desc)
		}
	}
} */
