package models

type Message struct {
	BotToken string `json:"bot_token" binding:"required"`
	Text     string `json:"text"`
	ChatID   string `json:"chat_id" binding:"required"`
}

type BotToken struct {
	BotToken string `json:"bot_token" binding:"required"`
}
