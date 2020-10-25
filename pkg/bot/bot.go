package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/lukemilby/telegramBot/pkg/cmd"
	"log"
)

type Bot struct {
	Bot *tgbotapi.BotAPI
	Registry map[string]cmd.Command
	Scope map[string]bool
}


func NewBot(apiKey string) (*Bot, error) {
	// Create bot
	bot, err := tgbotapi.NewBotAPI(apiKey)
	if err != nil {
		return nil, err
	}
	// Set new registry
	reg := make(map[string]cmd.Command)
	scope := make(map[string]bool)

	return &Bot{
		Bot:      bot,
		Registry: reg,
		Scope: scope,
	}, nil
}

func (b *Bot) AddScopes(scopes ...string) {
	for _, scope := range scopes {
		b.Scope[scope] = true
	}
}

func(a *Bot) Register(commands ...cmd.Command) error {
	for _, c := range commands {
		log.Printf("Loading... %s\n", c.GetName())
		a.Registry[c.GetName()] = c
	}
	return nil
}

func (a *Bot) ListCommands() {
	for k, _ := range a.Registry {
		log.Println("Command: ",k)
	}
}

