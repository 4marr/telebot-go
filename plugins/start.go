package plugins

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Start(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	chatID := update.Message.Chat.ID

	SendMessage(bot, chatID, "hello, "+update.Message.From.FirstName+"! I'm Ammaricano Bot, a bot that can help you with various tasks.")
}
