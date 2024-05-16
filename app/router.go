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

	tgPayload.CommandMessage = getCommandAndMessage(tgPayload.Message.Text)

	log.Println("CommandMessage", tgPayload.CommandMessage.Command)

	if tgPayload.CommandMessage.Command == "start" {
		app.start(tgPayload)
	}

	if tgPayload.CommandMessage.Command == "new" {
		app.newRoutine((tgPayload))
	}

	if tgPayload.CommandMessage.Command == "" {
		app.freeText(tgPayload)
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
