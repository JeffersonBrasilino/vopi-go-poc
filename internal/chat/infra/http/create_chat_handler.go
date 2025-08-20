package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vopi-go-poc/internal/chat/usecase/createchat"
	"github.com/vopi-go-poc/internal/core"
)

func CreateChatHandler(
	context *gin.Context,
	useCase core.UseCase[*createchat.CreateChatInput, any],
) {
		var requestPayload struct {
			ChannelId    string `json:"channelId" binding:"required"`
			Participants []*struct {
				Contact string `json:"contact" binding:"required"`
				Name    string `json:"name" binding:"required"`
			} `json:"participants" binding:"required"`
			Messages []*struct {
				Content string `json:"content" binding:"required"`
				Status  string `json:"status" binding:"required"`
				Sender  *struct {
					Contact string `json:"contact" binding:"required"`
					Name    string `json:"name" binding:"required"`
				} `json:"sender" binding:"required"`
			} `json:"messages" binding:"required"`
		}
		
		if err := context.ShouldBind(&requestPayload); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		useCaseInput := &createchat.CreateChatInput{
			ChannelId: requestPayload.ChannelId,
		}
		for _, p := range requestPayload.Participants {
			useCaseInput.Participants = append(useCaseInput.Participants, &createchat.Person{
				Contact: p.Contact,
				Name:    p.Name,
			})
		}
		for _, m := range requestPayload.Messages {
			useCaseInput.Messages = append(useCaseInput.Messages, &createchat.Message{
				Content: m.Content,
				Status:  m.Status,
				Sender: &createchat.Person{
					Contact: m.Sender.Contact,
					Name:    m.Sender.Name,
				},
			})
		}

		useCase.Execute(useCaseInput)
}
