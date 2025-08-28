package plugins

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type GPTResponse struct {
	Success bool   `json:"success"`
	Code    int    `json:"code"`
	Result  string `json:"result"`
	Creator string `json:"creator"`
}

func FetchGPTResponse(apiURL string, prompt string) (string, error) {
	fullURL := fmt.Sprintf("%s/api/ai/gpt?ask=%s", apiURL, prompt)

	resp, err := http.Get(fullURL)
	if err != nil {
		return "", fmt.Errorf("gagal mengambil respon dari API: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("gagal membaca response body: %w", err)
	}

	var gptRes GPTResponse
	err = json.Unmarshal(body, &gptRes)
	if err != nil {
		return "", fmt.Errorf("gagal decode JSON: %w", err)
	}

	if !gptRes.Success {
		return "", fmt.Errorf("API mengembalikan error: %s", gptRes.Result)
	}

	return gptRes.Result, nil
}

func GPT(bot *tgbotapi.BotAPI, update *tgbotapi.Update, apiURL string) {
	chatID := update.Message.Chat.ID
	args := update.Message.CommandArguments()

	if strings.TrimSpace(args) == "" {
		SendMessage(bot, chatID, "❗ Kirim perintah seperti ini:\n`/gpt Apa itu AI?`")
		return
	}

	loadingMsg := tgbotapi.NewMessage(chatID, escapeMarkdownV2("⏳ *Memproses permintaan...*"))
	loadingMsg.ParseMode = "MarkdownV2"
	sentMsg, _ := bot.Send(loadingMsg)

	response, err := FetchGPTResponse(apiURL, args)

	DeleteMessage(bot, chatID, sentMsg.MessageID)

	if err != nil {
		SendMessage(bot, chatID, fmt.Sprintf("❌ Terjadi kesalahan:\n`%v`", err))
		return
	}

	SendMessage(bot, chatID, response)
}
