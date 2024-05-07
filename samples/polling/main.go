// sample code for telegram-bot-go (get updates),
//
// last update: 2024.04.03.

package main

import (
	"fmt"
	"log"
	"time"

	bot "github.com/meinside/telegram-bot-go"
)

const (
	apiToken = "01234567:abcdefghijklmn_ABCDEFGHIJKLMNOPQRST"

	pollingIntervalSeconds = 1
	typingDelaySeconds     = 1

	verbose = true
)

// function for handling updates
func updateHandler(b *bot.Bot, update bot.Update, err error) {
	if err == nil {
		if update.HasMessage() {
			// 'is typing...'
			b.SendChatAction(update.Message.Chat.ID, bot.ChatActionTyping, nil)
			time.Sleep(typingDelaySeconds * time.Second)

			var message string

			if update.Message.HasContact() {
				message = fmt.Sprintf(
					"I received @%s's phone no.: %s",
					*update.Message.From.Username,
					update.Message.Contact.PhoneNumber,
				)
			} else if update.Message.HasLocation() {
				message = fmt.Sprintf(
					"I received @%s's location: (%f, %f)",
					*update.Message.From.Username,
					update.Message.Location.Latitude,
					update.Message.Location.Longitude,
				)
			} else {
				if update.Message.HasText() {
					message = fmt.Sprintf(
						"I received @%s's message: %s",
						*update.Message.From.Username,
						*update.Message.Text,
					)
				} else {
					message = fmt.Sprintf(
						"I received @%s's message",
						*update.Message.From.Username,
					)
				}
			}
			// send message
			if sent := b.SendMessage(
				update.Message.Chat.ID,
				message,
				// option
				bot.OptionsSendMessage{}.
					SetReplyParameters(bot.ReplyParameters{
						MessageID: update.Message.MessageID,
					}).                                        // show original message
					SetReplyMarkup(bot.NewReplyKeyboardMarkup( // show keyboards
						[][]bot.KeyboardButton{
							{
								bot.NewKeyboardButton("Just a button"),
							},
							{
								bot.NewKeyboardButton("Request contact").
									SetRequestContact(true),
								bot.NewKeyboardButton("Request location").
									SetRequestLocation(true),
							},
						},
					)),
			); !sent.Ok {
				log.Printf(
					"*** failed to send message: %s",
					*sent.Description,
				)
			}
		} else if update.HasInlineQuery() {
			// articles for inline query
			article1, _ := bot.NewInlineQueryResultArticle(
				"Star Wars quotes",
				"I am your father.",
				"Darth Vader")
			article2, _ := bot.NewInlineQueryResultArticle(
				"Star Wars quotes",
				"I know.",
				"Han Solo")

			results := []any{
				article1,
				article2,
			}

			// answer inline query
			if sent := b.AnswerInlineQuery(
				update.InlineQuery.ID,
				results,
				nil,
			); !sent.Ok {
				log.Printf(
					"*** failed to answer inline query: %s",
					*sent.Description,
				)
			}
		}
	} else {
		log.Printf(
			"*** error while receiving update (%s)",
			err.Error(),
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

	// get info about this bot
	if me := client.GetMe(); me.Ok {
		log.Printf("Bot information: %s", botName(me.Result))

		// delete webhook (getting updates will not work when wehbook is set up)
		if unhooked := client.DeleteWebhook(true); unhooked.Ok {
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
