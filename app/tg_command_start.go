package app

import "fmt"

/* send message for /start command */
func (app *App) start(tgPayload TelegramPayload) {
	starMessage := fmt.Sprintf(`¡Hola %v!
Revisa la documentación para conocer los comandos existentes`, tgPayload.Message.From.FirstName)

	app.sendTelegramMessageHTML(starMessage, fmt.Sprint(tgPayload.Message.Chat.ID))
}
