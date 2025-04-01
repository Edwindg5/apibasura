// apibasura/api/infraestructure/controllers/sensor_controller.go
package controllers

import (
	"apibasura/api/domain/entities"
	"apibasura/api/infraestructure/adapters"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SensorController struct {
	wsAdapter *adapters.WebSocketAdapter
}

func NewSensorController(wsAdapter *adapters.WebSocketAdapter) *SensorController {
	return &SensorController{
		wsAdapter: wsAdapter,
	}
}

func (sc *SensorController) ReceiveSensorData(c *gin.Context) {
	var sensorData map[string]interface{}
	
	if err := c.ShouldBindJSON(&sensorData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Crear un mensaje estructurado para el WebSocket
	message := entities.Message{
		Text:   "Datos del sensor recibidos",
		Action: "sensor_update",
	}

	// Enviar tanto los datos originales como el mensaje estructurado
	response := gin.H{
		"message": message,
		"sensor_data": sensorData,
	}

	// Enviar a través del WebSocket
	sc.wsAdapter.Broadcast(response)

	c.JSON(http.StatusOK, gin.H{
		"status": "Datos recibidos y enviados por WebSocket",
		"data":   sensorData,
	})
}