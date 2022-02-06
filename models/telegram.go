package models

type Message struct {
	Text   string `json:"text"`
	ChatID string `json:"chat_id"`
}

type Token struct {
	BotToken string `json:"token"`
}
