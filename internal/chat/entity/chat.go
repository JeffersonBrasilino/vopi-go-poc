package entity

import (
	"time"

	"github.com/vopi-go-poc/internal/core"
)

type Chat struct {
	channelId    string
	uuid         string
	participants []*Person
	messages     []*Message
	createdAt    time.Time
	updatedAt    time.Time
}


func NewChat(
	channelId string,
	participants []*Person,
	messages []*Message,
) (*Chat, error) {
	if err := validateChat(channelId, participants, messages); err != nil {
		return nil, core.NewValidationError(err)
	}
	return &Chat{
		channelId:    channelId,
		uuid:         "uuid generated",
		participants: participants,
		messages:     messages,
		createdAt:    time.Now(),
		updatedAt:    time.Now(),
	}, nil
}

func validateChat(
	channelId string,
	participants []*Person,
	messages []*Message,
) []string {
	var errors []string
	if channelId == "" {
		errors = append(errors, "channelId is required")
	}

	if len(participants) == 0 {
		errors = append(errors, "at least one participant is required")
	}

	if len(messages) == 0 {
		errors = append(errors, "at least one message is required")
	}

	return errors
}

// Getters
func (c *Chat) ChannelId() string {
	return c.channelId
}

func (c *Chat) Uuid() string {
	return c.uuid
}

func (c *Chat) Participants() []*Person {
	return c.participants
}

func (c *Chat) Messages() []*Message {
	return c.messages
}

func (c *Chat) CreatedAt() time.Time {
	return c.createdAt
}

func (c *Chat) UpdatedAt() time.Time {
	return c.updatedAt
}
