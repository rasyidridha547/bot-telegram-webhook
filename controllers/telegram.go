package controllers

import (
	"crypto/tls"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

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

	url := baseUrl + token.BotToken + "/getMe"

	resp, err := client.R().Get(url)
	if err != nil {
		log.Println(err)
		return
	}

	body := helper.Decode(resp)
	if body["ok"] == false && body["error_code"] == http.StatusUnauthorized {
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

	token := os.Getenv("BOT_TOKEN")
	chatID := os.Getenv("CHAT_ID")

	url := baseUrl + token + "/sendMessage"

	// convert raw data to string
	payload := string(rawData)

	resp, err := client.R().SetQueryParams(map[string]string{
		"chat_id": chatID,
		"text":    payload,
	}).
		Post(url)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"message": "error",
			"reason":  err.Error(),
		})
		log.Println(err)
		return
	}

	body := helper.Decode(resp)

	if body["ok"] == false || body["error_code"] == http.StatusUnauthorized {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "error",
			"reason":  "wrong bot token",
		})
		log.Println("bot token not valid!")
		return
	}

	log.Println("successfully send data to telegram")

	c.JSON(http.StatusOK, gin.H{
		"message":       "success",
		"response_time": resp.Time().String(),
		"result":        body,
	})
}

func GoogleWebhook(c *gin.Context) {

	// get raw data from requested server
	alertData := helper.GetRequestBody(c.Request.Body)

	token := os.Getenv("BOT_TOKEN")
	chatID := os.Getenv("CHAT_ID")

	// convert date to readable human (not readable alien of course)
	timeAlert := alertData["incident"].(map[string]interface{})["started_at"].(float64)
	formatDate := helper.ConvertUnixToDate(timeAlert)

	// payload message
	payload := "Google Cloud Platform Monitoring Alert \n\n" +
		"Alert Title: " + alertData["incident"].(map[string]interface{})["condition_name"].(string) + "\n" +
		"VM: " + alertData["incident"].(map[string]interface{})["resource_name"].(string) + "\n" +
		"Date: " + formatDate + "\n" +
		"Summary: " + alertData["incident"].(map[string]interface{})["summary"].(string) + "\n\n" +
		"Please kindly check. Thank you for using this bot :)"

	url := baseUrl + token + "/sendMessage"

	resp, err := client.R().
		SetQueryParams(map[string]string{
			"chat_id": chatID,
			"text":    payload,
		}).
		Post(url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "error",
			"reason":  err.Error(),
		})
		log.Println(err)
		return
	}

	body := helper.Decode(resp)

	if body["error_code"] == http.StatusUnauthorized {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "error",
			"reason":  "wrong bot token",
		})
		log.Println("bot token not valid")
		return
	}

	log.Println("successfully send Google Cloud alert to telegram")

	c.Status(http.StatusOK)
}

func DatadogWebhook(c *gin.Context) {
	log.Println("getting data from Datadog Alert")

	// get raw data from requested server
	alertData := helper.GetRequestBody(c.Request.Body)

	token := os.Getenv("BOT_TOKEN")
	chatID := os.Getenv("CHAT_ID")

	// convert date to readable human (not readable alien of course)
	timeAlert := alertData["last_updated"].(string)
	timeInt, err := strconv.Atoi(timeAlert)
	timeFloat := float64(timeInt)
	formatDate := helper.ConvertUnixToDate(timeFloat)

	payload := "Datadog Monitoring Alert \n\n" +
		"Message: " + alertData["body"].(string) + "\n" +
		"Date: " + formatDate + "\n\n" +
		"Please kindly check. Thank you for using this bot :)"

	url := baseUrl + token + "/sendMessage"

	resp, err := client.R().
		SetQueryParams(map[string]string{
			"chat_id": chatID,
			"text":    payload,
		}).
		Post(url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "error",
			"reason":  err.Error(),
		})
		log.Println(err)
		return
	}

	body := helper.Decode(resp)

	if body["error_code"] == http.StatusUnauthorized {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "error",
			"reason":  "wrong bot token",
		})
		log.Println("bot token not valid")
		return
	}

	log.Println("successfully send Datadog alert to telegram")

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
		log.Println("bot token not valid!")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":        "success",
		"message":       "sukses mengirim pesan",
		"response_time": resp.Time().String(),
	})
}
