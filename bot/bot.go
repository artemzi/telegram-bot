package bot

import (
	"log"
	"net/http"

	"gopkg.in/telegram-bot-api.v4"
)

var config = InitConfig()

// Run func start a bot and return (instance, updates object)
func Run() (*tgbotapi.BotAPI, tgbotapi.UpdatesChannel) {
	bot := getBot(config.TelegramToken)
	log.Printf("Authorized on account %s", bot.Self.UserName)

	bot.GetWebhookInfo() // TODO remove debug info

	// bot.RemoveWebhook() // TODO check
	_, err := bot.SetWebhook(tgbotapi.NewWebhook(config.WebhookURL + "/" + bot.Token))
	if err != nil {
		log.Fatal(err)
	}

	updates := bot.ListenForWebhook("/" + bot.Token)
	go http.ListenAndServe(":8080", nil) // TODO

	return bot, updates
}

func getBot(token string) *tgbotapi.BotAPI {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatalf("Error in getting new bot: %v\n", err)
	}

	bot.Debug = config.Debug
	return bot
}
