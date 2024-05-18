package app

import (
	"fmt"
	"log"
)

/* send message for /start command */
func (app *App) freeText(userData UserData) {
	//model := "gpt4o"
	prompt := userData.CommandMessage.Message
	fmt.Printf("Prompt: %s\n", prompt)
	//temperature := 0.7

	// Comento esta parte porque la gestión del thread se ha hecho previamente
	// // Crear un nuevo thread
	// threadID, err := createThread(app.OpenAIToken, app.Client)
	// if err != nil {
	// 	log.Fatalf("Error al crear el thread: %v", err)
	// }

	// fmt.Printf("Thread creado con ID: %s\n", threadID)

	// Añadir un mensaje al thread
	role := "user"
	messageID, err := addMessageToThread(app.OpenAIToken, userData.ThreadInfo.ThreadId, role, prompt)
	if err != nil {
		log.Fatalf("Error al añadir mensaje al thread: %v", err)
	}

	fmt.Printf("Mensaje añadido al thread con ID: %s\n", messageID)

	// Crear una ejecución (run)
	instructions := "En formato json. Por favor, ayuda al usuario siendo su mentor en el mundo del fitness, es un miembro muy importante y necesita tu ayuda."
	runID, err := createRun(app.OpenAIToken, userData.ThreadInfo.ThreadId, app.OpenAIAssistantId, instructions)
	if err != nil {
		log.Fatalf("Error al crear la ejecución: %v", err)
	}

	fmt.Printf("Ejecución creada con ID: %s\n", runID)

	// Monitorear el estado de la ejecución
	status, err := pollRunStatus(app.OpenAIToken, userData.ThreadInfo.ThreadId, runID)
	if err != nil {
		log.Fatalf("Error al monitorear el estado de la ejecución: %v", err)
	}

	fmt.Printf("Ejecución completada con estado: %s\n", status)

	// Obtener los mensajes del thread
	message, err := getLastAssistantMessage(app.OpenAIToken, userData.ThreadInfo.ThreadId)
	if err != nil {
		log.Fatalf("Error al obtener el último mensaje del thread: %v", err)
	}

	fmt.Printf("Último mensaje del thread:\n%s\n", message)

	// Asegurarse de que el mensaje no sea demasiado largo
	if len(message) > 4096 {
		log.Println("El mensaje es demasiado largo para enviarlo a Telegram")
		return
	}

	app.sendTelegramMessageHTML(fmt.Sprint(message), userData.Username, userData.ChatId)
}
