package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rasyidridha532/bot-telegram-webhook/controllers"
	"log"
	"net/http"
	"time"
)

func main() {
	// get port from env
	port := controllers.DotEnvVar("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	// initiate Gin
	r := gin.Default()

	api := r.Group("/api")

	v1 := api.Group("/v1")
	v1.GET("/profile", controllers.Profile)
	v1.GET("/message", controllers.IncomingWebhook)
	v1.POST("/message", controllers.SendMessage)

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Page Not Found",
		})
	})

	r.GET("/", func(c *gin.Context) {
		t := time.Now()

		c.JSON(http.StatusOK, gin.H{
			"message":       "welcome to my webhook :)",
			"response_time": time.Since(t).String(),
		})
	})

	err := r.Run(":" + port)
	if err != nil {
		log.Fatal(err)
	}
}
