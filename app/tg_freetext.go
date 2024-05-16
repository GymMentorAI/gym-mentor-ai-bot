package app

import (
	"fmt"
	"gym_mentor_ai/app/openai"
	"log"
	"os"
)

/* send message for /start command */
func (app *App) freeText(tgPayload TelegramPayload) {
	apiKey := os.Getenv("GYM_MENTOR_OPENAI_TOKEN")
	//model := "gpt4o"
	prompt := tgPayload.Message.Text
	//temperature := 0.7

	// Crear un nuevo thread
	threadID, err := openai.CreateThread(apiKey)
	if err != nil {
		log.Fatalf("Error al crear el thread: %v", err)
	}

	fmt.Printf("Thread creado con ID: %s\n", threadID)

	// Añadir un mensaje al thread
	role := "user"
	messageID, err := openai.AddMessageToThread(apiKey, threadID, role, prompt)
	if err != nil {
		log.Fatalf("Error al añadir mensaje al thread: %v", err)
	}

	fmt.Printf("Mensaje añadido al thread con ID: %s\n", messageID)

	// Crear una ejecución (run)
	assistantID := os.Getenv("GYM_MENTOR_OPENAI_ASSISTANT_ID")
	instructions := "En formato json. Por favor, ayuda al usuario siendo su mentor en el mundo del fitness, es un miembro muy importante y necesita tu ayuda."
	runID, err := openai.CreateRun(apiKey, threadID, assistantID, instructions)
	if err != nil {
		log.Fatalf("Error al crear la ejecución: %v", err)
	}

	fmt.Printf("Ejecución creada con ID: %s\n", runID)

	// Monitorear el estado de la ejecución
	status, err := openai.PollRunStatus(apiKey, threadID, runID)
	if err != nil {
		log.Fatalf("Error al monitorear el estado de la ejecución: %v", err)
	}

	fmt.Printf("Ejecución completada con estado: %s\n", status)

	// Obtener los mensajes del thread
	message, err := openai.GetLastAssistantMessage(apiKey, threadID)
	if err != nil {
		log.Fatalf("Error al obtener el último mensaje del thread: %v", err)
	}

	fmt.Printf("Último mensaje del thread:\n%s\n", message)

	app.sendTelegramMessageHTML(message, fmt.Sprint(tgPayload.Message.Chat.ID))
}
