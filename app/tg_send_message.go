package app

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

/* Use telegram bot to send HTML message */
func (app *App) sendTelegramMessageHTML(message, username string) {
	/* TODO: comprobar límite del mensaje que se va a enviar, igual es necesario recortarlo */

	payload := map[string]string{
		"chat_id":    username,
		"parse_mode": "html",
		"text":       message,
	}

	payloadJSON, payloadJSONError := json.Marshal(payload)
	if payloadJSONError != nil {
		log.Println("sendTelegramMessageHTML deserialization json error", payloadJSONError)
		return
	}

	endpoint := fmt.Sprintf("https://api.telegram.org/bot%v/sendMessage", app.TelegramBotToken)

	req, reqError := http.NewRequest("POST", endpoint, bytes.NewReader(payloadJSON))
	if reqError != nil {
		log.Println("sendTelegramMessageHTML reqError", reqError)
		return
	}
	req.Header.Add("Content-Type", "application/json")

	r, rError := app.Client.Do(req)
	if rError != nil {
		log.Println("sendTelegramMessageHTML rError", rError)
		return
	}

	if r.StatusCode != 200 {
		// TODO: gestionar un posible exceso de peticiones por concurrencia de usuarios
		// if res.StatusCode == 429 {

		// }
		log.Println("sendTelegramMessageHTML r.StatusCode error", r.Status)

	}

}
