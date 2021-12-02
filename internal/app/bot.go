package go_telegram_rss_bot

import (
	"fmt"
	"log"

	tgbot "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var my_rss_bot *tgbot.BotAPI

func SendMessage(msg tgbot.MessageConfig) {
	if _, err := my_rss_bot.Send(msg); err != nil {
		fmt.Println(err)
	}
}

func StartBot() {
	fmt.Println("Hello telegram rss bot~\nMake sure you've already loaded the cfg.json.")

	bot, err := tgbot.NewBotAPI(cfg.TgBotToken)
	if err != nil {
		log.Panic(err)
	}

	my_rss_bot = bot

	log.Printf("Authorized on account [%s]", my_rss_bot.Self.UserName)

	updater := tgbot.NewUpdate(0)
	updater.Timeout = 60

	updates := my_rss_bot.GetUpdatesChan(updater)

	// start fetching the rss feeds
	go StartFetchRss()

	for update := range updates {
		if update.Message == nil {
			// ignore any non-Message Updates
			continue
		}

		if update.Message.Chat.ID != int64(cfg.TgAdminId) {
			// ignore others
			continue
		}

		if !update.Message.IsCommand() {
			// echo the message
			msg := tgbot.NewMessage(update.Message.Chat.ID, update.Message.Text)
			msg.ReplyToMessageID = update.Message.MessageID
			my_rss_bot.Send(msg)
			continue
		}

		msg := tgbot.NewMessage(update.Message.Chat.ID, "")

		// Extract the command from the Message.
		switch update.Message.Command() {
		case "help":
			msg.Text = "I understand /biu and /status."
		case "biu":
			msg.Text = "I died. XP"
		case "status":
			msg.Text = "I'm ok."

		default:
			msg.Text = "I don't know that command"
		}

		if _, err := my_rss_bot.Send(msg); err != nil {
			log.Panic(err)
		}
	}
}
