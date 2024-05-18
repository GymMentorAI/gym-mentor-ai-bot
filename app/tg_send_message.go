package app

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

/* Use telegram bot to send HTML message */
func (app *App) sendTelegramMessageHTML(message, username string, chatId string) {
	log.Println("sendTelegramMessageHTML: enviando mensaje a", username)

	// Comprobar límite del mensaje y recortar si es necesario
	if len(message) > 4096 {
		log.Println("sendTelegramMessageHTML: el mensaje supera los 4096 caracteres y será recortado")
		message = message[:4096]
	}

	payload := map[string]string{
		"chat_id":    chatId,
		"parse_mode": "HTML", // Asegurarse de usar mayúsculas
		"text":       message,
	}

	payloadJSON, payloadJSONError := json.Marshal(payload)
	if payloadJSONError != nil {
		log.Println("sendTelegramMessageHTML: error al serializar JSON", payloadJSONError)
		return
	}

	endpoint := fmt.Sprintf("https://api.telegram.org/bot%v/sendMessage", app.TelegramBotToken)

	req, reqError := http.NewRequest("POST", endpoint, bytes.NewReader(payloadJSON))
	if reqError != nil {
		log.Println("sendTelegramMessageHTML: error al crear la solicitud", reqError)
		return
	}
	req.Header.Add("Content-Type", "application/json")

	r, rError := app.Client.Do(req)
	if rError != nil {
		log.Println("sendTelegramMessageHTML: error al enviar la solicitud", rError)
		return
	}
	defer r.Body.Close()

	if r.StatusCode != http.StatusOK {
		// Manejar el código de estado 429 para excesivas solicitudes
		if r.StatusCode == http.StatusTooManyRequests {
			log.Println("sendTelegramMessageHTML: demasiadas solicitudes, espera antes de reintentar")
			// Implementar lógica para manejar espera antes de reintentar, si es necesario
			return
		}

		body, _ := io.ReadAll(r.Body)
		log.Printf("sendTelegramMessageHTML: error al enviar mensaje, código de estado: %d, respuesta: %s", r.StatusCode, string(body))
		return
	}

	log.Println("sendTelegramMessageHTML: mensaje enviado correctamente")
}
