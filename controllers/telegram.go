package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"github.com/joho/godotenv"
	"github.com/rasyidridha532/bot-telegram-webhook/helper"
	"github.com/rasyidridha532/bot-telegram-webhook/models"
	"log"
	"net/http"
	"os"
	"time"
)

var client = resty.New()
var baseUrl = "https://api.telegram.org/bot"

func DotEnvVar(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Println("error load .env file")
	}

	return os.Getenv(key)
}

func Profile(c *gin.Context) {
	token := models.Token{}

	// bind payload body
	err := c.BindJSON(&token)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "error",
			"reason":  err,
		})
		return
	}

	//url := baseUrl + DotEnvVar("BOT_TOKEN") + "/getMe"
	url := baseUrl + token.BotToken + "/getMe"
	t := time.Now()

	resp, err := client.R().Get(url)
	if err != nil {
		log.Fatal(err)
	}

	response := helper.Decode(resp)

	c.JSON(http.StatusOK, gin.H{
		"message":       "success",
		"response_time": time.Since(t).String(),
		"result":        response,
	})
}

func IncomingWebhook(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"message": "test Incoming Webhook",
	})
}

func SendMessage(c *gin.Context) {
	token := &models.Token{}
	message := &models.Message{}

	// bind payload body
	err := c.BindJSON(&token)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "error",
			"reason":  err,
		})
		return
	}

	t := time.Now()

	url := baseUrl + token.BotToken + "/sendMessage"

	_, err = client.R().
		SetQueryParams(map[string]string{
			"chat_id": message.ChatID,
			"text":    message.Text,
		}).
		Post(url)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "error",
			"reason":  err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":        "success",
		"message":       "sukses mengirim pesan",
		"response_time": time.Since(t).String(),
	})
}
