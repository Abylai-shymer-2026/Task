package main

import (
	"log"
	"strings"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("6694955555:AAEHozEjxXkHjK_OEwaoJymIFwEIvQgumtI")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	n := 1
	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		if update.Message.IsCommand() {
			switch update.Message.Command() {
			case "start":
				n = 0
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Привет! Как тебя зовут?")
				bot.Send(msg)

			default:
				log.Printf("Unknown command: %s", update.Message.Text)
			}
		} else if update.Message.Text != "" && update.Message.IsCommand() == false && n == 0 {
			// Если не команда и текстовое сообщение не пустое, обрабатываем как имя
			name := strings.TrimSpace(update.Message.Text)
			replyText := "Привет " + name + "!"
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, replyText)
			msg.ReplyToMessageID = update.Message.MessageID
			bot.Send(msg)
			n = 1
		} else {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Сперва напиши /start")
			msg.ReplyToMessageID = update.Message.MessageID
			bot.Send(msg)

		}
	}
}
