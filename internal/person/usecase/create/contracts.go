package create

import "github.com/vopi-go-poc/internal/person/entity"

type CreateRepository interface {
	Create(person *entity.Person) error
	Exists(document string) (bool, error)
}
