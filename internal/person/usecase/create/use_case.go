package create

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/vopi-go-poc/internal/core"
	"github.com/vopi-go-poc/internal/person/entity"
)

type UseCase struct {
	repo CreateRepository
}

func NewCreate(repo CreateRepository) *UseCase {
	return &UseCase{
		repo: repo,
	}
}

func (c *UseCase) Execute(data *CreateInputDto) (*CreateOutputDto, error) {
	
	person, err := c.makePerson(data)
	if err != nil {
		return nil, err
	}

	errPersist := c.repo.Create(person)

	if errPersist != nil {
		return nil, errPersist
	}

	return &CreateOutputDto{
		Uuid: person.Uuid(),
	}, nil
}

func (c *UseCase) makePerson(data *CreateInputDto) (*entity.Person, error) {
	contacts, err := c.makeContacts(data)
	if err != nil {
		return nil, err
	}

	person, err := entity.NewPerson(
		uuid.NewString(),
		data.Name,
		data.Document,
		data.PersonType,
		contacts,
	)
	if err != nil {
		return nil, err
	}

	return person, nil
}

func (c *UseCase) makeContacts(data *CreateInputDto) ([]entity.Contact, error) {
	var contacts []entity.Contact
	var errors []string
	for k, contact := range data.Contacts {
		contactEntity, err := entity.NewContact(contact)
		if err != nil {
			errors = append(errors, fmt.Sprintf("contact.%d: %v", k, err))
			continue
		}
		contacts = append(contacts, *contactEntity)
	}
	if len(errors) > 0 {
		return nil, core.NewValidationError(errors)
	}
	return contacts, nil
}
