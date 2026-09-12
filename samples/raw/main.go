// sample code for telegram-bot-go (get updates and print them),
//
// last update: 2026.09.12.

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	bot "github.com/meinside/telegram-bot-go"
)

const (
	requestTimeoutSeconds  = 10
	pollingIntervalSeconds = 1
	typingDelaySeconds     = 1
)

// function for handling updates
func updateHandler(b *bot.Bot, update bot.Update, err error) {
	if err == nil {
		// just print the raw data to the stdout
		if bytes, err := json.MarshalIndent(update, "", "  "); err == nil {
			log.Printf("%s", string(bytes))
		} else {
			log.Printf(
				"* failed to marshal data: %s",
				err,
			)
		}
	} else {
		log.Printf(
			"* error while receiving update: %s",
			err,
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
	if apiToken, exists := os.LookupEnv("TELEGRAM_BOT_API_TOKEN"); exists {
		verbose := false
		if v, exists := os.LookupEnv("VERBOSE"); exists {
			verbose = v == "true"
		}

		client := bot.NewClient(apiToken)
		client.Verbose = verbose

		ctx, cancel := context.WithTimeout(context.TODO(), requestTimeoutSeconds*time.Second)
		defer cancel()

		// get info about this bot
		if me, _ := client.GetMe(ctx); me.OK {
			log.Printf("Bot information: %s", botName(me.Result))

			ctx, cancel := context.WithTimeout(context.TODO(), requestTimeoutSeconds*time.Second)
			defer cancel()

			// delete webhook (getting updates will not work when wehbook is set up)
			if unhooked, _ := client.DeleteWebhook(ctx, true); unhooked.OK {
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
	} else {
		panic("environment variable 'TELEGRAM_BOT_API_TOKEN' is not set")
	}
}
