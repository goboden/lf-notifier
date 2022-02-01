package main

import (
	"fmt"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	token, ok := os.LookupEnv("BOT_TOKEN")
	if !ok {
		log.Fatal("Telegram bot token not found.")
		return
	}

	fmt.Println(token)

	botApi, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Started")

	botApi.Debug = true

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := botApi.GetUpdatesChan(u)
	if err != nil {
		log.Fatal(err)
	}

	for update := range updates {
		message := update.Message
		if message == nil {
			continue
		}

		send := tgbotapi.NewMessage(message.Chat.ID, message.Text)
		botApi.Send(send)
	}

}
