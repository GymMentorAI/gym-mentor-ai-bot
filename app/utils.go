package app

import (
	"strings"
	"unicode"
)

/* Get command and message from a text with format => /comand rest text of message*/
func getCommandAndMessage(text string) CommandMessage {
	var commandMessage CommandMessage
	if !strings.HasPrefix(text, "/") {
		commandMessage.Message = text
		return commandMessage
	}
	var command, message string
	isCommand := true
	for i, v := range text {
		if i == 0 {
			continue
		}

		if isCommand && unicode.IsSpace(v) {
			isCommand = false
			continue
		}

		if isCommand {
			command += string(v)
			continue
		}
		message += string(v)

	}
	commandMessage.Command = command
	commandMessage.Message = message
	return commandMessage

}
