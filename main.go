// apibasura/main.go
package main

import (
	"apibasura/api/infraestructure/dependencies"
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	
	r := gin.Default()

	// Configuración de CORS más completa
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Middleware adicional para WebSocket
	r.Use(func(c *gin.Context) {
		if c.Request.Header.Get("Upgrade") == "websocket" {
			c.Header("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Credentials", "true")
		}
		c.Next()
	})

	log.Println("Inicializando dependencias...")
	dependencies.InitMessages(r)
	
	log.Println("WebSocket inicializado y listo para aceptar conexiones en /v1/messages/ws")
	log.Println("Servidor HTTP iniciado en el puerto :5000")
	
	if err := r.Run(":5000"); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}