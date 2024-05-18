package app

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"

	"log"
	"net/http"
	"time"
)

func createThread(apiKey string, client *http.Client) (string, error) {
	endpoint := "https://api.openai.com/v1/threads"

	// Crear la solicitud HTTP
	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer([]byte(`{}`)))
	if err != nil {
		return "", fmt.Errorf("error al crear la solicitud HTTP: %v", err)
	}

	// Configurar los encabezados
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiKey))
	req.Header.Set("OpenAI-Beta", "assistants=v2")

	// Hacer la solicitud
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error al hacer la solicitud: %v", err)
	}
	defer resp.Body.Close()

	// Leer la respuesta
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error al leer la respuesta: %v", err)
	}

	log.Println("Body: ", string(body))

	// Verificar el código de estado
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("error en la solicitud: %v, cuerpo: %s", resp.StatusCode, body)
	}

	// Parsear la respuesta JSON
	var response map[string]interface{}
	if err := json.Unmarshal(body, &response); err != nil {
		return "", fmt.Errorf("error al parsear la respuesta JSON: %v", err)
	}

	threadID, ok := response["id"].(string)
	if !ok {
		return "", fmt.Errorf("no se pudo encontrar el ID del thread en la respuesta")
	}

	return threadID, nil
}

func addMessageToThread(apiKey, threadID, role, content string) (string, error) {
	endpoint := fmt.Sprintf("https://api.openai.com/v1/threads/%s/messages", threadID)

	// Datos para la solicitud
	requestBody, err := json.Marshal(map[string]interface{}{
		"role":    role,
		"content": content,
	})
	if err != nil {
		return "", fmt.Errorf("error al crear el cuerpo de la solicitud: %v", err)
	}

	// Crear la solicitud HTTP
	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(requestBody))
	if err != nil {
		return "", fmt.Errorf("error al crear la solicitud HTTP: %v", err)
	}

	// Configurar los encabezados
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiKey))
	req.Header.Set("OpenAI-Beta", "assistants=v2")

	// Hacer la solicitud
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error al hacer la solicitud: %v", err)
	}
	defer resp.Body.Close()

	// Leer la respuesta
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error al leer la respuesta: %v", err)
	}

	// Verificar el código de estado
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("error en la solicitud: %v, cuerpo: %s", resp.StatusCode, body)
	}

	// Parsear la respuesta JSON
	var response map[string]interface{}
	if err := json.Unmarshal(body, &response); err != nil {
		return "", fmt.Errorf("error al parsear la respuesta JSON: %v", err)
	}

	messageID, ok := response["id"].(string)
	if !ok {
		return "", fmt.Errorf("no se pudo encontrar el ID del mensaje en la respuesta")
	}

	return messageID, nil
}

func createRun(apiKey, threadID, assistantID, instructions string) (string, error) {
	endpoint := fmt.Sprintf("https://api.openai.com/v1/threads/%s/runs", threadID)

	requestBody, err := json.Marshal(map[string]interface{}{
		"assistant_id": assistantID,
		"instructions": instructions,
	})
	if err != nil {
		return "", fmt.Errorf("error al crear el cuerpo de la solicitud: %v", err)
	}

	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(requestBody))
	if err != nil {
		return "", fmt.Errorf("error al crear la solicitud HTTP: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiKey))
	req.Header.Set("OpenAI-Beta", "assistants=v2")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error al hacer la solicitud: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error al leer la respuesta: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("error en la solicitud: %v, cuerpo: %s", resp.StatusCode, body)
	}

	var response map[string]interface{}
	if err := json.Unmarshal(body, &response); err != nil {
		return "", fmt.Errorf("error al parsear la respuesta JSON: %v", err)
	}

	runID, ok := response["id"].(string)
	if !ok {
		return "", fmt.Errorf("no se pudo encontrar el ID del run en la respuesta")
	}

	return runID, nil
}

func pollRunStatus(apiKey, threadID, runID string) (string, error) {
	endpoint := fmt.Sprintf("https://api.openai.com/v1/threads/%s/runs/%s", threadID, runID)

	for {
		req, err := http.NewRequest("GET", endpoint, nil)
		if err != nil {
			return "", fmt.Errorf("error al crear la solicitud HTTP: %v", err)
		}

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiKey))
		req.Header.Set("OpenAI-Beta", "assistants=v2")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			return "", fmt.Errorf("error al hacer la solicitud: %v", err)
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return "", fmt.Errorf("error al leer la respuesta: %v", err)
		}

		if resp.StatusCode != http.StatusOK {
			return "", fmt.Errorf("error en la solicitud: %v, cuerpo: %s", resp.StatusCode, body)
		}

		var response map[string]interface{}
		if err := json.Unmarshal(body, &response); err != nil {
			return "", fmt.Errorf("error al parsear la respuesta JSON: %v", err)
		}

		status, ok := response["status"].(string)
		if !ok {
			return "", fmt.Errorf("no se pudo encontrar el estado del run en la respuesta")
		}

		if status == "completed" || status == "failed" {
			return status, nil
		}

		time.Sleep(2 * time.Second)
	}
}

func getLastAssistantMessage(apiKey, threadID string) (string, error) {
	messages, err := getMessagesFromThread(apiKey, threadID)
	if err != nil {
		return "", fmt.Errorf("error al obtener el último mensaje del asistente en el thread: %v", err)
	}

	var lastMessage string
	for i := len(messages) - 1; i >= 0; i-- {
		message := messages[i]
		if role, ok := message["role"].(string); ok && role == "assistant" {
			if content, ok := message["content"].([]interface{}); ok {
				for _, contentItem := range content {
					if contentMap, ok := contentItem.(map[string]interface{}); ok {
						if text, ok := contentMap["text"].(map[string]interface{}); ok {
							if value, ok := text["value"].(string); ok {
								lastMessage = value
								break
							}
						}
					}
				}
			}
		}
		if lastMessage != "" {
			break
		}
	}

	if lastMessage == "" {
		return "", fmt.Errorf("no se encontró ningún mensaje del asistente en el thread")
	}

	return lastMessage, nil
}

func getMessagesFromThread(apiKey, threadID string) ([]map[string]interface{}, error) {
	url := fmt.Sprintf("https://api.openai.com/v1/threads/%s/messages", threadID)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("OpenAI-Beta", "assistants=v2")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error en la respuesta de la API: %v", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var jsonResponse struct {
		Object string                   `json:"object"`
		Data   []map[string]interface{} `json:"data"`
	}

	err = json.Unmarshal(body, &jsonResponse)
	if err != nil {
		return nil, err
	}

	return jsonResponse.Data, nil
}
