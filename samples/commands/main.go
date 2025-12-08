// sample code for telegram-bot-go (handle commands),
//
// last update: 2024.04.03.

package main

import (
	"context"
	"fmt"
	"log"
	"time"

	bot "github.com/meinside/telegram-bot-go"
)

const (
	apiToken = "01234567:abcdefghijklmn_ABCDEFGHIJKLMNOPQRST"

	requestTimeoutSeconds  = 10
	pollingIntervalSeconds = 1
	typingDelaySeconds     = 1

	verbose = true
)

// handle '/start' command
func startCommandHandler(
	b *bot.Bot,
	update bot.Update,
	args string,
) {
	if update.HasMessage() {
		send(
			b,
			update.Message.Chat.ID,
			update.Message.MessageID,
			"Starting chat...",
		)
	}
}

// handle '/help' command
func helpCommandHandler(
	b *bot.Bot,
	update bot.Update,
	args string,
) {
	if update.HasMessage() {
		send(
			b,
			update.Message.Chat.ID,
			update.Message.MessageID,
			"Help message here.",
		)
	}
}

// handle non-supported commands
func noSuchCommandHandler(
	b *bot.Bot,
	update bot.Update,
	cmd, args string,
) {
	if update.HasMessage() {
		send(
			b,
			update.Message.Chat.ID,
			update.Message.MessageID,
			fmt.Sprintf("No such command: %s", cmd),
		)
	}
}

// handle non-command updates
func updateHandler(
	b *bot.Bot,
	update bot.Update,
	err error,
) {
	if err == nil {
		if update.HasMessage() {
			ctx, cancel := context.WithTimeout(context.TODO(), requestTimeoutSeconds*time.Second)
			defer cancel()

			// 'is typing...'
			b.SendChatAction(
				ctx,
				update.Message.Chat.ID,
				bot.ChatActionTyping,
				nil,
			)
			time.Sleep(typingDelaySeconds * time.Second)

			// send a reply,
			message := fmt.Sprintf("Received your message: %s", *update.Message.Text)
			send(
				b,
				update.Message.Chat.ID,
				update.Message.MessageID,
				message,
			)

			// and add a reaction on the received message
			react(
				b,
				update.Message.Chat.ID,
				update.Message.MessageID,
				"👍",
			)
		}
	} else {
		log.Printf(
			"*** error while receiving update (%s)",
			err.Error(),
		)
	}
}

// send a message
func send(
	b *bot.Bot,
	chatID, messageID int64,
	message string,
) {
	ctx, cancel := context.WithTimeout(context.TODO(), requestTimeoutSeconds*time.Second)
	defer cancel()

	if sent := b.SendMessage(
		ctx,
		chatID,
		message,
		bot.OptionsSendMessage{}.
			SetReplyParameters(bot.NewReplyParameters(messageID)), // show original message
	); !sent.Ok {
		log.Printf(
			"*** failed to send a message: %s",
			*sent.Description,
		)
	}
}

// leave a reaction on a message
func react(
	b *bot.Bot,
	chatID,
	messageID int64,
	emoji string,
) {
	ctx, cancel := context.WithTimeout(context.TODO(), requestTimeoutSeconds*time.Second)
	defer cancel()

	if reacted := b.SetMessageReaction(
		ctx,
		chatID,
		messageID,
		bot.NewMessageReactionWithEmoji(emoji),
	); !reacted.Ok {
		log.Printf(
			"*** failed to leave a reaction on a message: %s",
			*reacted.Description,
		)
	}
}

// generate bot's name
func botName(bot *bot.User) string {
	if bot != nil {
		if bot.Username != nil {
			return fmt.Sprintf("@%s (%s)", *bot.Username, bot.FirstName)
		} else {
			return bot.FirstName
		}
	}

	return "Unknown"
}

func main() {
	client := bot.NewClient(apiToken)
	client.Verbose = verbose

	ctx, cancel := context.WithTimeout(context.TODO(), requestTimeoutSeconds*time.Second)
	defer cancel()

	// get info about this bot
	if me := client.GetMe(ctx); me.Ok {
		log.Printf("Bot information: %s", botName(me.Result))

		ctx, cancel := context.WithTimeout(context.TODO(), requestTimeoutSeconds*time.Second)
		defer cancel()

		// delete webhook (getting updates will not work when wehbook is set up)
		if unhooked := client.DeleteWebhook(ctx, true); unhooked.Ok {
			// add command handlers
			client.AddCommandHandler("/start", startCommandHandler)
			client.AddCommandHandler("/help", helpCommandHandler)
			client.SetNoMatchingCommandHandler(noSuchCommandHandler)

			// wait for new updates
			client.StartPollingUpdates(
				0,
				pollingIntervalSeconds,
				updateHandler,
			)
		} else {
			panic("failed to delete webhook")
		}
	} else {
		panic("failed to get info of the bot")
	}
}
