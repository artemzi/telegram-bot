package bot

import "os"

// Config bot configuration endpoints
type Config struct {
	TelegramToken    string
	WebhookURL       string
	ListenWebhookURL string
	ListenAddr       string // TODO do i need it outside development process?
}

// InitConfig read env variables and return Config
// don't forget trailing slash in the end of WEBHOOK_URL,
// ex: https://XXXXXXXX.ngrok.io/
func InitConfig() Config {
	token := os.Getenv("TELEGRAM_TOKEN")
	return Config{
		TelegramToken:    token,
		WebhookURL:       os.Getenv("WEBHOOK_URL") + token,
		ListenWebhookURL: "/" + token,
		ListenAddr:       ":8080",
	}
}
