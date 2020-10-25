package cmd

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

func GetName(update tgbotapi.Update, scope *tgbotapi.BotAPI) error {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
	msg.ReplyToMessageID = update.Message.MessageID
	scope.Send(msg)
	return nil
}