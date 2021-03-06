package main

import (
	"log"

	"github.com/artemzi/summarizer"
	"github.com/artemzi/telegram-bot/bot"
	"github.com/asaskevich/govalidator"

	"gopkg.in/telegram-bot-api.v4"
)

func main() {
	app, updates := bot.Run()

	for update := range updates {
		log.Printf("[%s] %+v\n", update.Message.From, update.Message.Text)

		if update.Message.IsCommand() {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
			switch update.Message.Command() {
			case "help":
				msg.Text = `
					Please paste valid English article URL for bot.
You can use /status command for bot status.
				`
			case "start":
				msg.Text = "Hello I'm ready to help you."
			case "status":
				msg.Text = "I'm ok."
			default:
				msg.Text = "I don't know that command"
			}
			app.Send(msg)
			continue
		}

		if govalidator.IsRequestURL(update.Message.Text) { // if valid URL string
			s := summarizer.CreateFromURL(update.Message.Text)
			summary, err := s.Summarize()
			if err != nil {
				log.Println("Error occurred: ", err.Error())
			}

			summaryInfo, err := s.GetSummaryInfo()
			if err != nil {
				log.Println("Error occurred: ", err.Error())
			}

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, summary)
			msgInfo := tgbotapi.NewMessage(update.Message.Chat.ID, summaryInfo)
			app.Send(msg)
			app.Send(msgInfo)
		} else {
			log.Printf("Wrong url %s", update.Message.Text)
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Please provide valid URL string.")
			msg.ReplyToMessageID = update.Message.MessageID
			app.Send(msg)
		}
	}
}
