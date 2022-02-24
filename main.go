package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/rasyidridha532/bot-telegram-webhook/controllers"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println(err)
	}
}

func main() {
	// get port from env
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	env := os.Getenv("GO_ENV")
	if env == "release" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	// initiate Gin
	r := gin.Default()

	r.Use(cors.Default())

	// ::port/api
	api := r.Group("/api")

	// ::port/api/v1
	v1 := api.Group("/v1")
	v1.GET("/profile", controllers.Profile)
	v1.POST("/webhook", controllers.IncomingWebhook)
	v1.POST("/google-alert", controllers.GoogleWebhook)
	v1.POST("/datadog-alert", controllers.DatadogWebhook)
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

	err := r.Run(fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatal(err)
	}
}
