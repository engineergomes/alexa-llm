package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000" 
	}

	
	router := gin.Default()


	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	router.GET("/test", func(context *gin.Context) {
		log.Println(context)
		context.JSON(200, gin.H{
			"message": "Hello World",
		})
	})


	log.Printf("Servidor rodando na porta %s 🚀", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Erro ao iniciar servidor: %v", err)
	}
}
