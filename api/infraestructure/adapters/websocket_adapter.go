// apibasura/api/infraestructure/adapters/websocket_adapter.go
package adapters

import (
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket" //a nhooyr.io/websocket
)


type WebSocketAdapter struct {
	clients   map[*websocket.Conn]bool
	clientsMu sync.Mutex
	Upgrader  websocket.Upgrader
}

func NewWebSocketAdapter() *WebSocketAdapter {
	return &WebSocketAdapter{
		clients: make(map[*websocket.Conn]bool),
		Upgrader: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin: func(r *http.Request) bool {
				return true // Permite todas las conexiones
			},
		},
	}
}
func (wsa *WebSocketAdapter) AddClient(conn *websocket.Conn) {
	wsa.clientsMu.Lock()
	defer wsa.clientsMu.Unlock()
	wsa.clients[conn] = true
	log.Println("Nuevo cliente WebSocket conectado")
}

func (wsa *WebSocketAdapter) RemoveClient(conn *websocket.Conn) {
	wsa.clientsMu.Lock()
	defer wsa.clientsMu.Unlock()
	delete(wsa.clients, conn)
	log.Println("Cliente WebSocket desconectado")
}

func (wsa *WebSocketAdapter) Broadcast(message interface{}) {
	wsa.clientsMu.Lock()
	defer wsa.clientsMu.Unlock()

	for client := range wsa.clients {
		err := client.WriteJSON(message)
		if err != nil {
			log.Printf("Error enviando mensaje WebSocket: %v", err)
			client.Close()
			delete(wsa.clients, client)
		}
	}
}