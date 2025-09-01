package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vopi-go-poc/internal/core"
	"github.com/vopi-go-poc/internal/person/usecase/create"
)

func CreatePersonHandler(
	context *gin.Context,
	useCase core.UseCase[*create.CreateInputDto, *create.CreateOutputDto],
) {
	var requestPayload struct {
		Name       string   `json:"name" binding:"required"`
		Document   string   `json:"document" binding:"required"`
		Contacts   []string `json:"contacts" binding:"required"`
		PersonType int      `json:"personType" binding:"required"`
	}
	if err := context.ShouldBind(&requestPayload); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"type": err.Error()})//TODO: normalizar errors.
		return
	}

	/* useCaseInput := &create.CreateInputDto{
		Name:       requestPayload.Name,
		Document:   requestPayload.Document,
		Contacts:   requestPayload.Contacts,
		PersonType: requestPayload.PersonType,
	}

	result, err := useCase.Execute(useCaseInput)
	fmt.Println("result processing", result, err)
	if err != nil {
		context.JSON(http.StatusUnprocessableEntity, err)
		return
	}

	context.JSON(http.StatusCreated, result) */
}
