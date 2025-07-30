package plugins

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TikTokResponse struct {
	Success bool `json:"success"`
	Code    int  `json:"code"`
	Result  struct {
		Title     string   `json:"title"`
		Play      string   `json:"play"`
		WmPlay    string   `json:"wmplay"`
		HdPlay    string   `json:"hdplay"`
		Music     string   `json:"music"`
		Images    []string `json:"images"`
		MusicInfo struct {
			Title  string `json:"title"`
			Author string `json:"author"`
		} `json:"music_info"`
	} `json:"result"`
}

func TikTok(bot *tgbotapi.BotAPI, update *tgbotapi.Update, apiURL string) {
	chatID := update.Message.Chat.ID
	text := update.Message.Text

	// Tampilkan pesan loading
	loading := tgbotapi.NewMessage(chatID, "⏳ *Mengambil data dari TikTok...*")
	loading.ParseMode = "MarkdownV2"
	sentLoading, _ := bot.Send(loading)

	// Panggil API TikTok
	url := fmt.Sprintf("%s/api/download/tiktok?url=%s", apiURL, text)
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != 200 {
		DeleteMessage(bot, chatID, sentLoading.MessageID)
		SendMessage(bot, chatID, "❌ Gagal menghubungi server TikTok.")
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	var result TikTokResponse
	if err := json.Unmarshal(body, &result); err != nil || !result.Success {
		DeleteMessage(bot, chatID, sentLoading.MessageID)
		SendMessage(bot, chatID, "❌ Gagal membaca data dari TikTok.")
		return
	}

	// Jika ada gambar (photomode)
	if len(result.Result.Images) > 0 {
		for i, img := range result.Result.Images {
			photo := tgbotapi.NewPhoto(chatID, tgbotapi.FileURL(img))
			photo.Caption = fmt.Sprintf("ini kak gambar ke %d", i+1)

			// Gunakan MarkdownV2 (auto-escape caption jika ingin lebih aman)
			photo.ParseMode = "MarkdownV2"

			// Hanya gambar pertama yang dikasih tombol musik
			if i == 0 && result.Result.Music != "" {
				btn := tgbotapi.NewInlineKeyboardMarkup(
					tgbotapi.NewInlineKeyboardRow(
						tgbotapi.NewInlineKeyboardButtonURL("🎵 Music", result.Result.Music),
					),
				)
				photo.ReplyMarkup = btn
			}

			bot.Send(photo)
		}

		DeleteMessage(bot, chatID, sentLoading.MessageID)
		return
	}

	// Jika post berupa video
	video := tgbotapi.NewVideo(chatID, tgbotapi.FileURL(result.Result.Play))
	video.Caption = "Ini kak videonya"
	video.ParseMode = "MarkdownV2"

	// Inline buttons
	buttons := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonURL("🎞 Tanpa Watermark", result.Result.Play),
			tgbotapi.NewInlineKeyboardButtonURL("📹 Watermark", result.Result.WmPlay),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonURL("🎥 HD", result.Result.HdPlay),
			tgbotapi.NewInlineKeyboardButtonURL("🎵 Music", result.Result.Music),
		),
	)
	video.ReplyMarkup = buttons

	bot.Send(video)

	// Hapus pesan loading
	DeleteMessage(bot, chatID, sentLoading.MessageID)
}
