// apibasura/api/infraestructure/controllers/websocket_controller.go
package controllers

import (
	"apibasura/api/infraestructure/adapters"
	"log"


	"github.com/gin-gonic/gin"
)

type WebSocketController struct {
	wsAdapter *adapters.WebSocketAdapter
}

func NewWebSocketController(wsAdapter *adapters.WebSocketAdapter) *WebSocketController {
	return &WebSocketController{wsAdapter: wsAdapter}
}

func (wsc *WebSocketController) HandleWebSocket(c *gin.Context) {
	conn, err := wsc.wsAdapter.Upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Printf("Error al actualizar a WebSocket: %v", err)
		return
	}

	wsc.wsAdapter.AddClient(conn)
	defer wsc.wsAdapter.RemoveClient(conn)

	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			break
		}
	}
}