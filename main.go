package main

import (
	openai "alexa-ai/clients"
	"io"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// Load .env file, making it accessible using os.Getenv() or os.LookupEnv()
func init() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }
}


func main() {
	
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000" 
	}

	router := gin.Default()

	router.GET("/health", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"status": "ok",
		})
	})

	router.POST("/chat", func(context *gin.Context) {
		body, err := io.ReadAll(context.Request.Body)

		if err != nil {
			context.JSON(500, gin.H{
				"message": "Error reading body",
			})
			return
		}

      	response :=  openai.Chat(string(body))
	  	context.JSON(200, gin.H{
			"message": response,
		})
	})

	err := router.Run(":" + port)
    
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
	
}
