package main

import (
	"log"
	"net/http"
	"os"

	"gopkg.in/telegram-bot-api.v4"
)

func getBot(token string) *tgbotapi.BotAPI {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatalf("Error in getting new bot: %v\n", err)
	}

	bot.Debug = false
	return bot
}

// Run starts bot
func main() {
	webhookURL := "https://ef1bc65b.ngrok.io/"

	bot := getBot(os.Getenv("TELEGRAM_TOKEN"))
	// TODO: remove it
	if bot.Debug {
		bot.GetWebhookInfo()
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	bot.RemoveWebhook()
	_, err := bot.SetWebhook(tgbotapi.NewWebhook(webhookURL + bot.Token))
	if err != nil {
		log.Fatal(err)
	}

	updates := bot.ListenForWebhook("/" + bot.Token)
	go http.ListenAndServe(":8080", nil)

	for update := range updates {
		log.Printf("[%s] %+v\n", update.Message.From, update.Message.Text)
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		msg.ReplyToMessageID = update.Message.MessageID

		bot.Send(msg)
	}
}
