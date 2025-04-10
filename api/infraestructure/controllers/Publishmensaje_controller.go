// apibasura/api/infraestructure/controllers/PublishMessage_controller.go
package controllers

import (
	"apibasura/api/application"
	"apibasura/api/domain/entities"
	"apibasura/api/infraestructure/adapters"

	"github.com/gin-gonic/gin"
	"net/http"
)

type PublishMessageController struct {
	publishMessageUseCase *application.PublishMessageUseCase
	wsAdapter            *adapters.WebSocketAdapter
}

func NewPublisMessageController(publishUseCase *application.PublishMessageUseCase, wsAdapter *adapters.WebSocketAdapter) *PublishMessageController {
	return &PublishMessageController{
		publishMessageUseCase: publishUseCase,
		wsAdapter:            wsAdapter,
	}
}

func (pmsg *PublishMessageController) PublishMessage(g *gin.Context) {
	var message entities.Message
	if err := g.ShouldBindJSON(&message); err != nil {
		g.JSON(http.StatusNotAcceptable, gin.H{"error": err.Error()})
		return
	}

	messageToPublish, err2 := pmsg.publishMessageUseCase.Execute(message.Text, message.Action)
	if err2 != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err2.Error()})
		return
	}

	// Enviar mensaje a todos los clientes WebSocket conectados
	pmsg.wsAdapter.Broadcast(messageToPublish)

	g.JSON(http.StatusOK, gin.H{"data": messageToPublish})
}