package main

import (
	"log"
	"telegram-ban-bot/internal/api"
	"telegram-ban-bot/internal/utils"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	config := utils.ReadConfig()
	bot, authErr := tgbotapi.NewBotAPI(config.BotToken)
	if authErr != nil {
		log.Fatalf("Bot cant autorize error: %v", authErr)
	}

	bot.Debug = false

	log.Printf("Authorized as: %v", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)
	for update := range updates {
		if update.Message != nil {
			go api.HandleUpdate(bot, update, config)
		}
	}
}
