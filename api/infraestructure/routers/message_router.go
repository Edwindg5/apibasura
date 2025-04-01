package routers

import (
	"apibasura/api/infraestructure/controllers"

	"github.com/gin-gonic/gin"
)

func MessageRouter(
	r *gin.Engine, 
	publishController *controllers.PublishMessageController, 
	wsController *controllers.WebSocketController,
	sensorController *controllers.SensorController,
) {
	v1 := r.Group("/v1/messages")
	{
		v1.POST("/publish", publishController.PublishMessage)
		v1.GET("/ws", wsController.HandleWebSocket)
		v1.POST("/sensor", sensorController.ReceiveSensorData) // Nueva ruta
	}
}