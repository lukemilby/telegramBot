package cmd

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

type Command interface {
	Execute (update tgbotapi.Update, scope *tgbotapi.BotAPI) error
	GetName() string
	GetHelp() string
}


type Commander struct {
	Name string
	Help string
	Scope string
	Cmd func(update tgbotapi.Update, scope *tgbotapi.BotAPI) error
}

func (c *Commander) GetName() string {
	return c.Name
}

func (c *Commander) GetHelp() string {
	return c.Help
}

func (c *Commander) Execute(update tgbotapi.Update, scope *tgbotapi.BotAPI) error {
	return c.Cmd(update, scope)
}

func NewCommand(name, help, scope string, cmd func(update tgbotapi.Update, scope *tgbotapi.BotAPI) error) Command {
	if scope == "" {
		scope = "public"
	}
	return &Commander{
		Name: name,
		Help: help,
		Scope: scope,
		Cmd:  cmd,
	}
}
