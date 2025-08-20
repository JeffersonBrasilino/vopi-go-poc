package chat

import (
	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/vopi-go-poc/internal/chat/infra/database"
	"github.com/vopi-go-poc/internal/chat/infra/http"
	"github.com/vopi-go-poc/internal/chat/usecase/createchat"
	"github.com/vopi-go-poc/internal/core"
)

type ChatModule struct {
	repo *database.ChatPostgresRepository
	createUseCase core.UseCase[*createchat.CreateChatInput, any]
}

func NewChatModule(dbInstance core.DbConnection) *ChatModule {
	slog.Info("ChatModule init.")
	module:= &ChatModule{
		repo: nil,
	}

	module.registerUseCases()
	return module
}

func (m *ChatModule) registerUseCases() {
	m.createUseCase = createchat.NewCreateChat(m.repo)
}

func (m *ChatModule) WithHttp(router *gin.Engine) *ChatModule {
	slog.Info("PersonModule HTTP routes registered.")
	router.Group("/chat").
		POST("", func(ctx *gin.Context) { http.CreateChatHandler(ctx, m.createUseCase) })

	return m
}