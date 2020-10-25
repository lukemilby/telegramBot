package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joho/godotenv"
	"github.com/lukemilby/telegramBot/pkg/bot"
	"github.com/lukemilby/telegramBot/pkg/cmd"
	"log"
	"os"
	"strings"
)


//TODO: Functioning Scope
//TODO: HTTP Support
//TODO: Elasticsearch Logging
//TODO: Profiling
//TODO: Dockerfile
//TODO: Makefile

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get settings from env file
	apiKey := os.Getenv("API_KEY")


	b, err := bot.NewBot(apiKey)
	// Scoping the operation of the bot
	b.AddScopes("Dust")
	b.AddScopes("Crungeon")


	// Register commands
	b.Register(
		cmd.NewCommand("GetName", "Gets users name", "Crungeon", cmd.GetName),
		)

	log.Printf("Autorized on account %s", b.Bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := b.Bot.GetUpdatesChan(u)

	// Main loop
	for update := range updates {
		if update.Message == nil {
			continue
		}
		log.Printf("[%s][%s] %s", update.Message.Chat.Title, update.Message.From.String(), update.Message.Text)
		// Using Scope to filter unwanted messages
		if b.InScope(update.Message.Chat.Title) {
			msg := strings.Split(update.Message.Text, " ")
			if cmd, ok := b.Registry[msg[0]]; ok {
				err := cmd.Execute(update, b.Bot)
				if err != nil {
					log.Printf("Error %s", err.Error())
				}
			}
		}
	}

}