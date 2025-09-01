package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vopi-go-poc/internal/chat/usecase/createchat"
	"github.com/vopi-go-poc/internal/core"
	"github.com/vopi-go-poc/internal/core/otel"
)

func CreateChatHandler(
	context *gin.Context,
	useCase core.UseCase[*createchat.CreateChatInput, any],
	tracer otel.OtelTracer,
) {
	ctx, span := tracer.Start(
		context.Request.Context(),
		"Create chat",
		otel.NewOtelAttr("http.method", "POST"),
		otel.NewOtelAttr("http.url", context.Request.URL.String()),
	)
	defer span.End()

	span.AddEvent("start request")

	var requestPayload struct {
		ChannelId    string `json:"channelId" binding:"required"`
		BotName      string `json:"botName" binding:"required"`
		Participants []*struct {
			Name     string `json:"name" binding:"required"`
			Document string `json:"document" binding:"required"`
			Contact  string `json:"contact" binding:"required"`
		} `json:"participants" binding:"required"`
		Messages []*struct {
			Content string `json:"content" binding:"required"`
			Status  string `json:"status" binding:"required"`
			Sender  string `json:"sender" binding:"required"`
		} `json:"messages" binding:"required"`
	}

	span.AddEvent("validate request")
	if err := context.ShouldBind(&requestPayload); err != nil {
		span.Error(err, "deu ruim")
		span.AddEvent("failed to bind request")
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//span.AddEvent("create chat input")
	useCaseInput := &createchat.CreateChatInput{
		ChannelId: requestPayload.ChannelId,
		BotName:   requestPayload.BotName,
	}

	for _, m := range requestPayload.Messages {
		useCaseInput.Messages = append(useCaseInput.Messages, &createchat.Message{
			Content: m.Content,
			Status:  m.Status,
			Sender:  m.Sender,
		})
	}

	for _, p := range requestPayload.Participants {
		useCaseInput.Participants = append(useCaseInput.Participants, &createchat.Person{
			Name:     p.Name,
			Document: p.Document,
			Contact:  p.Contact,
		})
	}

	span.AddEvent("execute use case")
	result, err := useCase.Execute(ctx, useCaseInput, tracer)
	if err != nil {
		span.Error(err, "deu ruim")
		span.AddEvent("failed to execute use case")
		context.JSON(http.StatusUnprocessableEntity, err)
		return
	}

	context.JSON(http.StatusCreated, result)
}
