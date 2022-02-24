# bot-telegram-webhook

This is my web app in Golang for getting webhook from Google Cloud and Datadog for alerting purposes and send it with bot telegram to chat/group

## Feature
- Web Framework: Gin
- Env Config: [godotenv](github.com/joho/godotenv)

## ENV Configuration

Enter your configuration in .env.example

| ENV         | Description                               |
|-------------|:------------------------------------------|
| GO_ENV      | app environment (development, production) |
| BOT_TOKEN   | bot telegram token                        |
| CHAT_ID     | chat id for sending the message           |
| PORT        | app port                                  |

## Endpoint
- /api/v1/profile
    - `GET` get overview about bot
- /api/v1/webhook
    - `GET` get raw webhook from anywhere
- /api/v1/google-alert
    - `POST` get alert data from GCP (Google Cloud Platform) and send it with telegram bot
- /api/v1/datadog-alert
    - `POST` get alert data from Datadog and send it with telegram bot
- /api/v1/message
    - `POST` send a message with telegram bot