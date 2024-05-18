package app

import (
	"encoding/json"
	"log"
	"net/http"
)

/* Manage index response for telegram payload*/
func (app *App) index(w http.ResponseWriter, r *http.Request) {
	var tgPayload TelegramPayload

	deserializationError := json.NewDecoder(r.Body).Decode(&tgPayload)
	if deserializationError != nil {
		log.Println("tgPayload deserializationError", deserializationError)
	}
	tgUsername := tgPayload.Message.From.Username
	tgText := tgPayload.Message.Text
	tgCommandMessage := getCommandAndMessage(tgText)

	// Recuperar threads del usuario
	dataThreads := app.getThreadInfo(tgUsername)

	if dataThreads.ThreadId == "" {
		// El usuario no tiene ningún thread previo
		createThreadId, createThreadIdError := createThread(app.OpenAIToken, app.Client)
		if createThreadIdError != nil {
			log.Println(createThreadIdError)
			// Cómo se gestiona un create thread error. Mandar un mensaje de volver a intentar en un rato, por ejemplo
			app.sendTelegramMessageHTML("Algo ha fallado, intentarlo de nuevo en un rato", tgUsername)
			return
		}
		dataThreads.ThreadId = createThreadId
		// Insertar en la base de datos
		insertId, insertIdError := app.insertUsernameAndThreadId(tgUsername, createThreadId)
		if insertIdError != nil {
			log.Println("ERROR:", insertIdError)
		} else {
			dataThreads.MySQLId = insertId
		}
	}

	// dataThreads ya tiene un identificador, bien porque se ha recuperado o porque se ha creado nuevo
	// Tampoco es necesario estar pasando todo el rato el payload de telegram que tiene bastante ruido, se podría simplificar

	var userData UserData
	userData.Username = tgUsername
	userData.CommandMessage = tgCommandMessage
	userData.ThreadInfo = dataThreads

	if userData.CommandMessage.Command == "start" {
		app.start(userData)
	}

	if userData.CommandMessage.Command == "new" {
		app.newRoutine(userData)
	}

	if userData.CommandMessage.Command == "" {
		app.freeText(userData)
	}

	w.Write([]byte("OK"))
}

/* Handle index endpoint */
func (app *App) handleRouter() {
	app.Router.HandleFunc("POST /", app.index)
}

/* Set new router for app */
func (app *App) newRouter() {
	app.Router = http.NewServeMux()
}
