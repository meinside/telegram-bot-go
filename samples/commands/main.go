// sample code for telegram-bot-go (handle commands),
//
// last update: 2023.05.30.

package main

import (
	"fmt"
	"log"
	"time"

	bot "github.com/meinside/telegram-bot-go"
)

const (
	apiToken = "01234567:abcdefghijklmn_ABCDEFGHIJKLMNOPQRST"

	monitorIntervalSeconds = 1
	typingDelaySeconds     = 1

	verbose = true
)

// handle '/start' command
func startCommandHandler(b *bot.Bot, update bot.Update, args string) {
	if update.HasMessage() {
		send(b, update.Message.Chat.ID, update.Message.MessageID, "Starting chat...")
	}
}

// handle '/help' command
func helpCommandHandler(b *bot.Bot, update bot.Update, args string) {
	if update.HasMessage() {
		send(b, update.Message.Chat.ID, update.Message.MessageID, "Help message here.")
	}
}

// handle non-supported commands
func noSuchCommandHandler(b *bot.Bot, update bot.Update, cmd, args string) {
	if update.HasMessage() {
		send(b, update.Message.Chat.ID, update.Message.MessageID, fmt.Sprintf("No such command: %s", cmd))
	}
}

// handle non-command updates
func handleUpdate(b *bot.Bot, update bot.Update, err error) {
	if err == nil {
		if update.HasMessage() {
			// 'is typing...'
			b.SendChatAction(update.Message.Chat.ID, bot.ChatActionTyping, nil)
			time.Sleep(typingDelaySeconds * time.Second)

			// send message
			message := fmt.Sprintf("Received your message: %s", *update.Message.Text)
			send(b, update.Message.Chat.ID, update.Message.MessageID, message)
		}
	} else {
		log.Printf(
			"*** error while receiving update (%s)",
			err.Error(),
		)
	}
}

// send message
func send(b *bot.Bot, chatID, messageID int64, message string) {
	if sent := b.SendMessage(
		chatID,
		message,
		// option
		bot.OptionsSendMessage{}.
			SetReplyToMessageID(messageID), // show original message
	); !sent.Ok {
		log.Printf(
			"*** failed to send message: %s",
			*sent.Description,
		)
	}
}

func main() {
	client := bot.NewClient(apiToken)
	client.Verbose = verbose

	// get info about this bot
	if me := client.GetMe(); me.Ok {
		log.Printf(
			"Bot information: @%s (%s)",
			*me.Result.Username,
			me.Result.FirstName,
		)

		// delete webhook (getting updates will not work when wehbook is set up)
		if unhooked := client.DeleteWebhook(true); unhooked.Ok {
			// add command handlers
			client.AddCommandHandler("/start", startCommandHandler)
			client.AddCommandHandler("/help", helpCommandHandler)
			client.SetNoMatchingCommandHandler(noSuchCommandHandler)

			// wait for new updates
			client.StartMonitoringUpdates(
				0,
				monitorIntervalSeconds,
				handleUpdate,
			)
		} else {
			panic("failed to delete webhook")
		}
	} else {
		panic("failed to get info of the bot")
	}
}
