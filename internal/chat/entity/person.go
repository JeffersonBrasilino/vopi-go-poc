package entity

import (
	"time"

	"github.com/vopi-go-poc/internal/core"
)

type Person struct {
	uuid       string
	name       string
	document   string
	personType int
	contacts   *Contact
	createdAt  time.Time
	updatedAt  time.Time
}

func NewPerson(
	uuid string,
	name string,
	document string,
	personType int,
	contacts *Contact,
) (*Person, error) {
	if err := validatePerson(uuid, name, document, personType, contacts); err != nil {
		return nil, core.NewValidationError(err)
	}

	return &Person{
		uuid:       uuid,
		name:       name,
		document:   document,
		personType: personType,
		contacts:   contacts,
		createdAt:  time.Now(),
		updatedAt:  time.Now(),
	}, nil
}

func validatePerson(
	uuid string, name string, document string, personType int, contacts *Contact) []string {
	var errors []string
	if uuid == "" {
		errors = append(errors, "uuid is required")
	}
	if name == "" {
		errors = append(errors, "name is required")
	}
	if document == "" {
		errors = append(errors, "document is required")
	}
	if personType == 0 {
		errors = append(errors, "personType is required")
	}
	if contacts == nil {
		errors = append(errors, "at least one contact is required")
	}
	return errors
}

func (p *Person) Uuid() string {
	return p.uuid
}

func (p *Person) Name() string {
	return p.name
}

func (p *Person) Document() string {
	return p.document
}

func (p *Person) Contacts() *Contact {
	return p.contacts
}

func (p *Person) PersonType() int {
	return p.personType
}
