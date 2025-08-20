package database

import (
	"database/sql"

	"github.com/vopi-go-poc/internal/core"
	"github.com/vopi-go-poc/internal/person/entity"
)

type PersonPostgresRepository struct {
	dbInstance *sql.DB
}

func NewPersonPostgresRepository(dbInstance core.DbConnection) *PersonPostgresRepository {
	return &PersonPostgresRepository{
		dbInstance: dbInstance.Instance(),
	}
}

func (r *PersonPostgresRepository) Create(person *entity.Person) error {
	query := `INSERT INTO persons (name, document, contacts, person_type) VALUES ($1, $2, $3, $4)`
	_, err := r.dbInstance.Exec(query, person.Name, person.Document, person.Contacts, person.PersonType)
	return err
}

func (r *PersonPostgresRepository) Exists(document string) (bool, error) {
	// Simulate a check in the database
	return false, nil
}
