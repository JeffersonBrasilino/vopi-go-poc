package core

import "database/sql"

type UseCase[TInput any, TOutput any] interface {
	Execute(input TInput) (output TOutput, err error)
}

type DbConnection interface {
	Instance() *sql.DB
	Disconnect() error
}