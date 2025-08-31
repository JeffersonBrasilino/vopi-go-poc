package person

import (
	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/vopi-go-poc/internal/core"
	"github.com/vopi-go-poc/internal/core/otel"
	"github.com/vopi-go-poc/internal/person/infra/database"
	"github.com/vopi-go-poc/internal/person/infra/http"
	"github.com/vopi-go-poc/internal/person/usecase/create"
)
var tracer = otel.InitTrace("person-module")
type PersonModule struct {
	repo          *database.PersonPostgresRepository
	createUseCase core.UseCase[*create.CreateInputDto, *create.CreateOutputDto]
}

func NewPersonModule(dbInstance core.DbConnection) *PersonModule {
	module := &PersonModule{
		repo: database.NewPersonPostgresRepository(dbInstance),
	}
	module.registerUseCases()

	return module
}

func (m *PersonModule) WithHttp(router *gin.Engine) *PersonModule {
	slog.Info("PersonModule HTTP routes registered.")
	router.Group("/person").
		POST("", func(ctx *gin.Context) { http.CreatePersonHandler(ctx, m.createUseCase) })

	return m
}

func (m *PersonModule) registerUseCases() {
	//m.createUseCase = create.NewCreate(m.repo)
}