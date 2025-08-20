package core

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type postgresConnection struct {
	instance *sql.DB
}

func NewPostgresConnection(
	host string,
	port int,
	user string,
	password string,
	dbname string,
	sslMode string,
) *postgresConnection {
	db, err := sql.Open("postgres",
		fmt.Sprintf(
			"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
			host, port, user, password, dbname, sslMode),
	)

	if err != nil {
		log.Fatalf("Falha ao abrir a conexão: %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Falha ao conectar com o banco de dados: %v", err)
	}

	return &postgresConnection{
		instance: db,
	}
}

func (p *postgresConnection) Instance() *sql.DB {
	return p.instance
}

func (p *postgresConnection) Disconnect() error {
	if p.instance == nil {
		return nil
	}

	err := p.instance.Close()
	if err != nil {
		return err
	}

	p.instance = nil
	return nil
}
