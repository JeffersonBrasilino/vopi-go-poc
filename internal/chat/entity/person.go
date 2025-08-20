package entity

import (
	"time"

	"github.com/vopi-go-poc/internal/core"
)

type Person struct {
	uuId    string
	contact string
	name    string
	createdAt time.Time
	updatedAt time.Time
}

func NewPerson(
	contact string,
	name    string,
) (*Person, error) {
	if err := validatePerson(contact, name); err != nil {
		return nil, core.NewValidationError(err)
	}
	return &Person{
		uuId:    "UUID person generated",
		contact: contact,
		name:    name,
		createdAt: time.Now(),
		updatedAt: time.Now(),
	}, nil
}

func validatePerson(contact, name string) []string {
	var errors []string
	if contact == "" {
		errors = append(errors, "contact is required")
	}

	if name == "" {
		errors = append(errors, "name is required")
	}
	
	return errors
}

// Getters
func (p *Person) UuId() string {
	return p.uuId
}

func (p *Person) Contact() string {
	return p.contact
}

func (p *Person) Name() string {
	return p.name
}

func (p *Person) CreatedAt() time.Time {
	return p.createdAt
}

func (p *Person) UpdatedAt() time.Time {
	return p.updatedAt
}
