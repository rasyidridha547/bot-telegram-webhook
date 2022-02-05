package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"github.com/joho/godotenv"
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
		log.Fatal("error load .env file")
	}

	return os.Getenv(key)
}

func decodejson(resp *resty.Response) map[string]interface{} {
	var response map[string]interface{}

	err := json.Unmarshal(resp.Body(), &response)
	if err != nil {
		return nil
	}

	if response == nil {
		fmt.Println("error json unmarshal")
	}

	return response
}

func Profile(c *gin.Context) {
	token := models.Token{}

	err := c.BindJSON(&token)
	if err != nil {
		log.Fatal(err)
	}

	//url := baseUrl + DotEnvVar("BOT_TOKEN") + "/getMe"
	url := baseUrl + token.BotToken + "/getMe"
	t := time.Now()

	resp, err := client.R().EnableTrace().Get(url)
	if err != nil {
		log.Fatal(err)
	}

	response := decodejson(resp)

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

}
