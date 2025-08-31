package createchat

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/vopi-go-poc/internal/chat/entity"
	"github.com/vopi-go-poc/internal/core"
	"github.com/vopi-go-poc/internal/core/otel"
)

type UseCase struct {
	database CreateChatRepository
	trace otel.OtelTracer
}

func NewCreateChat(
	database CreateChatRepository,
) *UseCase {
	return &UseCase{
		database: database,
	}
}

func (r *UseCase) Execute(ctx context.Context, data *CreateChatInput, trace otel.OtelTracer) (any, error) {
	_, span := trace.Start(ctx, "application - Use Case", otel.NewOtelAttr("chat.id", data.ChannelId))
	defer span.End()

	span.AddEvent("start use case", otel.NewOtelAttr("chat.channelId", data.ChannelId))
	chat, err := r.makeChat(data)
	if err != nil {
		span.Error(err,"failed to make chat")
		span.AddEvent("failed to make chat")
		return nil, err
	}

	span.Success("chat created successfully")
	return chat, nil
}

func (c *UseCase) makeParticipants(data *CreateChatInput) ([]*entity.Person, error) {
	var participants []*entity.Person
	var errors []string

	for k, p := range data.Participants {
		contact, err := entity.NewContact(p.Contacts)
		if err != nil {
			errors = append(errors, fmt.Sprintf("participants.%d.contacts: %v", k, err))
			continue
		}

		person, err := entity.NewPerson(
			uuid.NewString(),
			p.Name,
			p.Document,
			1,
			contact,
		)
		if err != nil {
			errors = append(errors, fmt.Sprintf("participants.%d: %v", k, err))
		}
		participants = append(participants, person)
	}
	if len(errors) > 0 {
		return nil, core.NewValidationError(errors)
	}

	return participants, nil
}

func (c *UseCase) makeMessages(data *CreateChatInput, participants []*entity.Person) ([]*entity.Message, error) {
	var messages []*entity.Message
	var errors []string

	for k, m := range data.Messages {

		sender := c.getParticipantByContact(participants, m.Sender)
		if sender == nil {
			errors = append(errors, fmt.Sprintf("messages.%d.sender: %v", k, "participant not found"))
			continue
		}

		message, err := entity.NewMessage(
			uuid.NewString(),
			sender,
			m.Content,
			m.Status,
		)

		if err != nil {
			errors = append(errors, fmt.Sprintf("messages.%d: %v", k, err))
			continue
		}

		messages = append(messages, message)
	}

	if len(errors) > 0 {
		return nil, core.NewValidationError(errors)
	}
	return messages, nil
}

func (c *UseCase) makeChat(data *CreateChatInput) (*entity.Chat, error) {
	participants, err := c.makeParticipants(data)
	if err != nil {
		return nil, err
	}

	messages, err := c.makeMessages(data, participants)
	if err != nil {
		return nil, err
	}

	chat, err := entity.NewChat(
		uuid.NewString(),
		participants,
		messages,
	)
	if err != nil {
		return nil, err
	}
	return chat, nil
}

func (c *UseCase) getParticipantByContact(participants []*entity.Person, contact string) *entity.Person {
	for _, p := range participants {
		if p.Contacts().Contact() == contact {
			return p
		}
	}
	return nil
}
