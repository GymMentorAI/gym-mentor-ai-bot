package app

import "fmt"

/* Create new gym routine */
func (app *App) newRoutine(tgPayload TelegramPayload) {
	/* TODO:

	Controlar si un usuario puede tener múltiples rutinas

	¿Qué prompt habría que preparar a OpenAI para las indicaciones?

	Enviar el prompt

	Serializar la respuesta

	Guardar en la base de datos

	etc.

	etc.

	*/

	fmt.Printf("TODO: proceso de nueva rutina para: %v\n", tgPayload.Message.From.FirstName)
	app.sendTelegramMessageHTML("rutina creada", fmt.Sprint(tgPayload.Message.Chat.ID))

}
