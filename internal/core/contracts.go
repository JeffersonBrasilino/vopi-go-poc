package core

import (
	"context"
	"database/sql"

	"github.com/vopi-go-poc/internal/core/otel"
)

type UseCase[TInput any, TOutput any] interface {
	Execute(ctx context.Context, input TInput, trace otel.OtelTracer) (output TOutput, err error)
}

type DbConnection interface {
	Instance() *sql.DB
	Disconnect() error
}