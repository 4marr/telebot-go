package plugins

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Flux(bot *tgbotapi.BotAPI, update *tgbotapi.Update, apiURL string) {
	chatID := update.Message.Chat.ID
	args := update.Message.CommandArguments()

	if strings.TrimSpace(args) == "" {
		SendMessage(bot, chatID, "❗ Kirim perintah seperti ini:\n`/flux Sebuah kota cyberpunk futuristik di malam hari`")
		return
	}

	loadingMsg := tgbotapi.NewMessage(chatID, escapeMarkdownV2("🎨 *Sedang membuat gambar...*"))
	loadingMsg.ParseMode = "MarkdownV2"
	sentMsg, _ := bot.Send(loadingMsg)

	imageURL := fmt.Sprintf("%s/api/ai/flux?prompt=%s", apiURL, args)
	resp, err := http.Get(imageURL)
	if err != nil || resp.StatusCode != 200 {
		DeleteMessage(bot, chatID, sentMsg.MessageID)
		SendMessage(bot, chatID, fmt.Sprintf("❌ Gagal menghasilkan gambar:\n`%v`", err))
		return
	}
	defer resp.Body.Close()

	imageBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		DeleteMessage(bot, chatID, sentMsg.MessageID)
		SendMessage(bot, chatID, "❌ Gagal membaca data gambar.")
		return
	}

	fileBytes := tgbotapi.FileBytes{
		Name:  "flux_image.jpg",
		Bytes: imageBytes,
	}
	photoMsg := tgbotapi.NewPhoto(chatID, fileBytes)
	photoMsg.Caption = args
	bot.Send(photoMsg)

	DeleteMessage(bot, chatID, sentMsg.MessageID)
}
