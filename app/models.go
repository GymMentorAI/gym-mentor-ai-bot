package app

import (
	"database/sql"
	"net/http"
)

/*

Response for /start

{
  "update_id": 123456789,
  "message": {
    "message_id": 1,
    "from": {
      "id": 12345678,
      "is_bot": false,
      "first_name": "first_name",
      "username": "username",
      "language_code": "es"
    },
    "chat": {
      "id": 1234567,
      "first_name": "first_name",
      "username": "username",
      "type": "private"
    },
    "date": 1715589607,
    "text": "/new hola mundo",
    "entities": [
      {
        "offset": 0,
        "length": 6,
        "type": "bot_command"
      }
    ]
  }
}

*/

type TelegramPayload struct {
	CommandMessage CommandMessage
	UpdateID       int `json:"update_id"`
	Message        struct {
		MessageID int `json:"message_id"`
		From      struct {
			ID           int    `json:"id"`
			IsBot        bool   `json:"is_bot"`
			FirstName    string `json:"first_name"`
			Username     string `json:"username"`
			LanguageCode string `json:"language_code"`
		} `json:"from"`
		Chat struct {
			ID        int    `json:"id"`
			FirstName string `json:"first_name"`
			Username  string `json:"username"`
			Type      string `json:"type"`
		} `json:"chat"`
		Date     int    `json:"date"`
		Text     string `json:"text"`
		Entities []struct {
			Offset int    `json:"offset"`
			Length int    `json:"length"`
			Type   string `json:"type"`
		} `json:"entities"`
	} `json:"message"`
}

type UserData struct {
	Username       string
	ChatId         string
	ThreadInfo     ThreadInfo
	CommandMessage CommandMessage
}

type CommandMessage struct {
	Command string
	Message string
}

type ThreadInfo struct {
	MySQLId  int
	ThreadId string
}

type App struct {
	TelegramBotToken  string
	OpenAIToken       string
	OpenAIAssistantId string
	Client            *http.Client
	Router            *http.ServeMux
	MySQL             *sql.DB
}
