package bot

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

// DEBUG app state
var DEBUG = true

// Config bot configuration endpoints
type Config struct {
	TelegramToken string
	WebhookURL    string
	Debug         bool
}

type ngrok struct {
	PublicURL string `json:"public_url"`
}

// InitConfig setup config
func InitConfig() Config {
	data := &ngrok{}

	// get dev server https URL (curl http://localhost:4040/api/tunnels/command_line | jq .public_url)
	if DEBUG {
		cli := &http.Client{Timeout: 3 * time.Second}
		r, err := cli.Get("http://localhost:4040/api/tunnels/command_line")
		if err != nil {
			log.Fatalf("Error in getting config json: %v", err)
		}
		defer r.Body.Close()

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatalf("Error in reading response: %v", err)
		}

		json.Unmarshal(body, &data)
	} else {
		// TODO set data.PublicURL for debug == false
	}

	return Config{
		TelegramToken: os.Getenv("TELEGRAM_TOKEN"),
		WebhookURL:    data.PublicURL,
		Debug:         DEBUG,
	}
}
