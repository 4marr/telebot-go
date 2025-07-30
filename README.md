<!-- <p align="center">
  <img src="https://raw.githubusercontent.com/go-telegram-bot-api/telegram-bot-api/master/logo.png" alt="Telegram Bot API Logo" width="180"/>
</p> -->

<h1 align="center">🤖 Telebot-Go</h1>

<p align="center">
  🚀 A feature-rich Telegram bot built with Go, designed for high modularity and extensibility with a dynamic plugin-based architecture.
  Ready to be your personal AI assistant!
</p>

<p align="center">
  <a href="https://go.dev/" target="_blank">
    <img src="https://img.shields.io/badge/Go-1.21+-00ADD8?style=for-the-badge&logo=go&logoColor=white" alt="Go Version"/>
  </a>
  <a href="https://github.com/go-telegram-bot-api/telegram-bot-api" target="_blank">
    <img src="https://img.shields.io/badge/Telegram%20Bot%20API-v5-blue?style=for-the-badge&logo=telegram&logoColor=white" alt="Telegram Bot API Version"/>
  </a>
  <a href="https://github.com/4marr/telebot-go/blob/main/LICENSE" target="_blank">
    <img src="https://img.shields.io/badge/License-MIT-yellow?style=for-the-badge&logo=github&logoColor=white" alt="License"/>
  </a>
  <a href="https://github.com/4marr/telebot-go/stargazers" target="_blank">
    <img src="https://img.shields.io/github/stars/4marr/telebot-go?style=for-the-badge&color=FFD700" alt="GitHub Stars"/>
  </a>
  <a href="https://github.com/4marr/telebot-go/forks" target="_blank">
    <img src="https://img.shields.io/github/forks/4marr/telebot-go?style=for-the-badge&color=C0C0C0" alt="GitHub Forks"/>
  </a>
</p>

---

## ✨ Key Features

This bot comes packed with advanced functionalities, thanks to its flexible plugin system:

*   🗣️ **AI Chatbots:** Interact with various leading AI models for intelligent conversations.
*   🎨 **Image Generation:** Create stunning images from text prompts using Flux.
*   ✈️ **Auto Download Tiktok:** Download Tiktok video's or photo's
*   🛡️ **Robust Error Handling:** Strong error handling with informative and easy-to-understand messages.

---

## 🚀 Getting Started

Follow these simple steps to get your bot up and running in no time!

### Prerequisites

*   <a href="https://go.dev/dl/" target="_blank">Go</a> (version 1.21 or higher)
*   <a href="https://git-scm.com/downloads" target="_blank">Git</a>
*   <a href="https://t.me/BotFather" target="_blank">Telegram Bot Token</a> from BotFather

### Installation

1.  Clone this repository to your local machine:

    ```bash
    git clone https://github.com/4marr/telebot-go.git
    cd telebot-go
    ```

2.  Create a `.env` file in the root directory and add your environment variables:

    ```env
    BOT_TOKEN=YOUR_TELEGRAM_BOT_TOKEN_HERE
    AMMARICANO_API=https://api.ammaricano.my.id # Default Ammaricano API endpoint
    ```

    ⚠️ Replace `YOUR_TELEGRAM_BOT_TOKEN_HERE` with your bot token obtained from BotFather.
    The `AMMARICANO_API` is crucial for enabling all AI functionalities.

### Running the Bot

You can run the bot directly or use `air` for a smoother development experience with live reloading.

#### Using `air` (Recommended for Development)

The `start.sh` script is configured to use `air`, which will automatically restart the bot whenever code changes are detected.

1.  Ensure `air` is installed. If not, `start.sh` will install it automatically:

    ```bash
    go install github.com/cosmtrek/air@latest
    ```

2.  Run the bot using the script:

    ```bash
    ./start.sh
    ```

#### Running the Go Application Directly

1.  Build the application:

    ```bash
    go build -o main .
    ```

2.  Run the executable:

    ```bash
    ./main
    ```

---

## 🔌 Modular Plugin System

This bot is designed with a highly modular plugin system. Each feature is a separate plugin, making it incredibly easy to add, remove, or modify functionalities without affecting the bot's core.

### Adding a New Plugin

1.  Create a new `.go` file inside the `plugins/` directory (e.g., `plugins/my_awesome_plugin.go`).

2.  Write your plugin's logic and register it within an `init()` function:

    ```go
    // plugins/my_awesome_plugin.go
    package plugins

    import (
    	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
    	// Import lib if needed, e.g.: "mytelebot/lib"
    )

    func myCommandHandler(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
    	// Your plugin logic goes here
    	// Example: Sending a reply message
    	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Hello from my new plugin!")
    	bot.Send(msg)
    }
    ```

---

## ⚙️ Environment Configuration

The bot uses environment variables for all its configurations.

### `.env` File

Create a `.env` file in the root directory of your project with the following variables:

```env
BOT_TOKEN=YOUR_TELEGRAM_BOT_TOKEN
AMMARICANO_API=https://api.ammaricano.my.id
```

*   **`BOT_TOKEN`**: Your unique Telegram bot token obtained from <a href="https://t.me/BotFather" target="_blank">BotFather</a>. This is a mandatory variable for the bot to connect to Telegram.

---

## 🤝 Contributing

Contributions are highly welcome! If you have suggestions, bug reports, or wish to contribute code, please feel free to open an issue or submit a pull request.

<p align="center">
  <a href="https://github.com/4marr/telebot-go/issues" target="_blank">
    <img src="https://img.shields.io/badge/Open%20an%20Issue-red?style=for-the-badge&logo=github" alt="Open an Issue"/>
  </a>
  <a href="https://github.com/4marr/telebot-go/pulls" target="_blank">
    <img src="https://img.shields.io/badge/Submit%20a%20PR-green?style=for-the-badge&logo=github" alt="Submit a PR"/>
  </a>
</p>

---

## 📄 License

This project is licensed under the MIT License. See the <a href="https://github.com/4marr/telebot-go/blob/main/LICENSE" target="_blank">LICENSE</a> file for more details.

---

## 📞 Contact

If you have any questions or require further assistance, please feel free to open an issue on this repository.

<p align="center">
  Made with ❤️ by <a href="https://github.com/4marr" target="_blank">Ammaricano</a>
</p>
