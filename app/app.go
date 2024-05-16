package app

import (
	"log"
	"net/http"
	"net/http/cookiejar"
	"os"
	"time"
)

func newHttpClient() *http.Client {
	jar, jarErr := cookiejar.New(nil)
	if jarErr != nil {
		log.Fatalln("jarErr", jar)
	}

	client := &http.Client{
		Timeout: 7 * time.Second,
		Jar:     jar,
	}

	return client
}

/*
Create new App and manage the endpoint for the webhook.
Rememeber set the webhook with:
curl -F "url=[URL_ENDPOINT]" https://api.telegram.org/bot[]BOT_TOKEN/setWebhook
*/
func NewApp() App {
	log.Println("App init")
	app := App{
		TelegramBotToken: os.Getenv("GYM_MENTOR_TG_BOT_TOKEN"),
		Client:           newHttpClient(),
	}
	app.newRouter()
	app.handleRouter()

	return app
}
