package plugins

import (
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func escapeMarkdownV2(text string) string {
	replacer := strings.NewReplacer(
		`_`, `\_`,
		`*`, `\*`,
		`[`, `\[`,
		`]`, `\]`,
		`(`, `\(`,
		`)`, `\)`,
		`~`, `\~`,
		"`", "\\`",
		`>`, `\>`,
		`#`, `\#`,
		`+`, `\+`,
		`-`, `\-`,
		`=`, `\=`,
		`|`, `\|`,
		`{`, `\{`,
		`}`, `\}`,
		`.`, `\.`,
		`!`, `\!`,
	)
	return replacer.Replace(text)
}

func SendMessage(bot *tgbotapi.BotAPI, chatID int64, text string) {
	escaped := escapeMarkdownV2(text)
	msg := tgbotapi.NewMessage(chatID, escaped)
	msg.ParseMode = "MarkdownV2"
	bot.Send(msg)
}

func DeleteMessage(bot *tgbotapi.BotAPI, chatID int64, messageID int) {
	del := tgbotapi.DeleteMessageConfig{
		ChatID:    chatID,
		MessageID: messageID,
	}
	bot.Request(del)
}
