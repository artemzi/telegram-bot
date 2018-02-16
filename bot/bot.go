package bot

import (
	"log"
	"net/http"
	"os"

	"gopkg.in/telegram-bot-api.v4"
)

// Run start bot and return (instance, updates object)
func Run() (*tgbotapi.BotAPI, tgbotapi.UpdatesChannel) {
	webhookURL := "https://5fe34562.ngrok.io/"
	bot := getBot(os.Getenv("TELEGRAM_TOKEN"))
	bot.GetWebhookInfo() // TODO

	log.Printf("Authorized on account %s", bot.Self.UserName)
	bot.RemoveWebhook() // TODO
	_, err := bot.SetWebhook(tgbotapi.NewWebhook(webhookURL + bot.Token))
	if err != nil {
		log.Fatal(err)
	}

	updates := bot.ListenForWebhook("/" + bot.Token)
	go http.ListenAndServe(":8080", nil)

	return bot, updates
}

func getBot(token string) *tgbotapi.BotAPI {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatalf("Error in getting new bot: %v\n", err)
	}

	bot.Debug = true
	return bot
}