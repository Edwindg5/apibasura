// apibasura/api/infraestructure/dependencies/setup_messages.go
package dependencies

import (
	"apibasura/api/application"
	"apibasura/api/infraestructure/adapters"
	"apibasura/api/infraestructure/controllers"
	"apibasura/api/infraestructure/routers"
	"fmt"

	"os"

	"github.com/gin-gonic/gin"
)

func InitMessages(r *gin.Engine) {
	// Configuración RabbitMQ (existente)
	rabbitmqUser := os.Getenv("RABBITMQ_USER")
	rabbitmqPass := os.Getenv("RABBITMQ_PASS")
	rabbitmqHost := os.Getenv("RABBITMQ_HOST")
	rabbitmqPort := os.Getenv("RABBITMQ_PORT")
	connStr := fmt.Sprintf("amqp://%s:%s@%s:%s/", rabbitmqUser, rabbitmqPass, rabbitmqHost, rabbitmqPort)

	// Inicializar WebSocket
	wsAdapter := adapters.NewWebSocketAdapter()
	wsController := controllers.NewWebSocketController(wsAdapter)

	// Inicializar RabbitMQ
	rabbitMQAdapter := adapters.NewRabbitMQAdapter(connStr)
	publishMessageUseCase := application.NewSaveMessage(rabbitMQAdapter)
	publishMessageController := controllers.NewPublisMessageController(publishMessageUseCase, wsAdapter)

	// Inicializar controlador de sensor
	sensorController := controllers.NewSensorController(wsAdapter)

	// Configurar routers
	routers.MessageRouter(r, publishMessageController, wsController, sensorController)
}