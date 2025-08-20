package entity

import (
	"time"

	"github.com/vopi-go-poc/internal/core"
)

type Message struct {
	uuid      string
	sender    *Person
	content   string
	status    string
	createdAt time.Time
	updatedAt time.Time
}

func NewMessage(
	sender *Person,
	content string,
	status string,
) (*Message, error) {
	if err := validateMessage(sender, content, status); err != nil {
		return nil, core.NewValidationError(err)
	}
	return &Message{
		uuid:      "UUID message generated",
		sender:    sender,
		content:   content,
		status:    status,
		createdAt: time.Now(),
		updatedAt: time.Now(),
	}, nil
}

func validateMessage(sender *Person, content string, status string) []string {
	var errors []string
	if content == "" {
		errors = append(errors, "content: is required")
	}

	if sender == nil {
		errors = append(errors, "sender: is required")
	}
	return errors
}

// Getters
func (m *Message) Uuid() string {
	return m.uuid
}

func (m *Message) Sender() *Person {
	return m.sender
}

func (m *Message) Content() string {
	return m.content
}

func (m *Message) Status() string {
	return m.status
}

func (m *Message) CreatedAt() time.Time {
	return m.createdAt
}

func (m *Message) UpdatedAt() time.Time {
	return m.updatedAt
}
