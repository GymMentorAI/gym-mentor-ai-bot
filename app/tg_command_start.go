package app

import "fmt"

/* send message for /start command */
func (app *App) start(userData UserData) {
	starMessage := fmt.Sprintf(`¡Hola %v!
Revisa la documentación para conocer los comandos existentes`, userData.Username)

	app.sendTelegramMessageHTML(starMessage, fmt.Sprint(userData.Username), fmt.Sprint(userData.ChatId))
}
