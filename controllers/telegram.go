package controllers

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"github.com/rasyidridha532/bot-telegram-webhook/helper"
	"github.com/rasyidridha532/bot-telegram-webhook/models"
)

var client = resty.New().SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
var baseUrl = "https://api.telegram.org/bot"

func Profile(c *gin.Context) {
	token := models.BotToken{}

	// bind payload body
	err := c.BindJSON(&token)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "error",
			"reason":  err.Error(),
		})
		return
	}

	//url := baseUrl + DotEnvVar("BOT_TOKEN") + "/getMe"
	url := baseUrl + token.BotToken + "/getMe"

	resp, err := client.R().Get(url)
	if err != nil {
		log.Println(err)
		return
	}

	body := helper.Decode(resp)
	if body["ok"] != true {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "error",
			"reason":  "bot token is invalid!",
		})
		log.Println("bot token not valid!")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":       "success",
		"response_time": resp.Time().String(),
		"result":        body,
	})
}

func IncomingWebhook(c *gin.Context) {
	log.Println("getting data from webhook")
	// get raw data from requested server
	rawData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Println(err.Error())
		return
	}

	token := helper.DotEnvVar("BOT_TOKEN")
	chatID := helper.DotEnvVar("CHAT_ID")
	port := helper.DotEnvVar("PORT")

	// convert raw data to string
	json := string(rawData)

	_, err = client.R().SetBody(
		models.Message{
			BotToken: token,
			Text:     json,
			ChatID:   chatID,
		},
	).Post(fmt.Sprintf("http://localhost:%s/api/v1/message", port))
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"message": "error",
			"reason":  err.Error(),
		})
		log.Println(err)
		return
	}

	log.Println("successfully send data to telegram")

	c.Status(http.StatusOK)
}

func SendMessage(c *gin.Context) {
	message := &models.Message{}

	// bind payload body
	err := c.BindJSON(&message)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "error",
			"reason":  err,
		})
		return
	}

	url := baseUrl + message.BotToken + "/sendMessage"

	resp, err := client.R().
		SetQueryParams(map[string]string{
			"chat_id": message.ChatID,
			"text":    message.Text,
		}).
		Post(url)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "error",
			"reason":  err.Error(),
		})
		return
	}

	body := helper.Decode(resp)

	if body["ok"] == false && body["error_code"] == http.StatusUnauthorized {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "error",
			"reason":  "wrong bot token",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":        "success",
		"message":       "sukses mengirim pesan",
		"response_time": resp.Time().String(),
	})
}
