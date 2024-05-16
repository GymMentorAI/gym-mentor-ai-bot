package main

import (
	"gym_mentor_ai/app"
	"log"
	"net/http"
	"os"
)

func init() {
	readEnvFile()
}

func main() {
	log.Println("Main init")
	app := app.NewApp()
	port := os.Getenv("GYM_MENTOR_WEB_PORT")
	http.ListenAndServe(":"+port, app.Router)
}
