package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"time"

	"github.com/fatih/color"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"

	"telebot-go/plugins"
)

func launchBot(botName string) {
	color.Green("🚀 Launching Telegram Bot...")
	time.Sleep(300 * time.Millisecond)

	border := "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
	username := fmt.Sprintf("@%s", botName)
	maxLen := len(border)
	padUser := fmt.Sprintf("%-*s", maxLen-11, username)

	fmt.Printf("\n%s\n\n", color.HiBlueString("┏ %s ┓", border))
	fmt.Printf("🤖 %s\n", color.HiGreenString("Telegram Bot is now online!     "))
	fmt.Printf("🔗 %s\n", color.HiCyanString("Logged in as: %s", padUser))
	fmt.Printf("%s\n\n", color.HiBlueString("┗ %s ┛", border))

	time.Sleep(300 * time.Millisecond)
}

func tikTokRegex(text string) bool {
	re := regexp.MustCompile(`(?i)https?://(www\.)?(vm|vt|m|www)?\.?tiktok\.com/[^\s]+`)
	return re.MatchString(text)
}

func main() {
	_ = godotenv.Load()
	botToken := os.Getenv("BOT_TOKEN")
	APIURL := os.Getenv("AMMARICANO_API")
	if botToken == "" {
		log.Fatal("BOT_TOKEN environment variable is not set")
	}

	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = false
	launchBot(bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		if update.Message.IsCommand() {
			command := update.Message.Command()

			switch command {
			case "start":
				plugins.Start(bot, &update)
			case "gpt":
				plugins.GPT(bot, &update, APIURL)
			case "flux":
				plugins.Flux(bot, &update, APIURL)
			default:
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Perintah tidak dikenal. Coba /start atau /gpt.")
				bot.Send(msg)
			}
		} else if tikTokRegex(update.Message.Text) {
			plugins.TikTok(bot, &update, APIURL)
		} else {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Pesan tidak dikenali. Gunakan /start, /gpt, atau kirim link TikTok.")
			bot.Send(msg)
		}
	}
}
