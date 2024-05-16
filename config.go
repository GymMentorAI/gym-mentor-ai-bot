package main

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
	"strings"
)

/* read .env file in the same folder and set keys and values for every line*/
func readEnvFile() {
	//execPath, execPathError := os.Executable()
	//if execPathError != nil {
	//	log.Fatalln("execPathError", execPathError)
	//}

	//folderPath := filepath.Dir(execPath)
	// Determine the current working directory
	workingDir, err := os.Getwd()
	if err != nil {
		log.Fatalln("Error getting working directory:", err)
	}
	envPath := filepath.Join(workingDir, ".env")

	log.Println("Looking for .env file at:", envPath)

	env, envError := os.Open(envPath)
	if envError != nil {
		log.Println("WARNING: initializing without .env file")
		return
	}

	buffer := bufio.NewScanner(env)
	buffer.Split(bufio.ScanLines)

	for buffer.Scan() {
		line := strings.TrimSpace(buffer.Text())
		if line == "" {
			continue
		}
		if strings.HasPrefix(line, "#") {
			continue
		}
		var key, value string

		isKey := true
		for _, v := range line {
			if isKey && v == '=' {
				isKey = false
				continue
			}
			if isKey {
				key += string(v)
				continue
			}
			value += string(v)
		}

		if key != "" && value != "" {
			os.Setenv(key, value)
		}
	}

}
