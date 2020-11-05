// sample code for telegram-bot-go (get updates),
//
// last update: 2018.11.20.
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

// update handler function
func handleUpdate(b *bot.Bot, update bot.Update, err error) {
	if err == nil {
		if update.HasMessage() {
			// 'is typing...'
			b.SendChatAction(update.Message.Chat.ID, bot.ChatActionTyping)
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
					SetReplyToMessageID(update.Message.MessageID). // show original message
					SetReplyMarkup(bot.ReplyKeyboardMarkup{        // show keyboards
						Keyboard: [][]bot.KeyboardButton{
							[]bot.KeyboardButton{
								bot.KeyboardButton{
									Text: "Just a button",
								},
							},
							[]bot.KeyboardButton{
								bot.KeyboardButton{
									Text:           "Request contact",
									RequestContact: true,
								},
								bot.KeyboardButton{
									Text:            "Request location",
									RequestLocation: true,
								},
							},
						},
					}),
			); !sent.Ok {
				log.Printf(
					"*** failed to send message: %s\n",
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

			results := []interface{}{
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
					"*** failed to answer inline query: %s\n",
					*sent.Description,
				)
			}
		}
	} else {
		log.Printf(
			"*** error while receiving update (%s)\n",
			err.Error(),
		)
	}
}

func main() {
	client := bot.NewClient(apiToken)
	client.Verbose = verbose

	// get info about this bot
	if me := client.GetMe(); me.Ok {
		log.Printf(
			"Bot information: @%s (%s)\n",
			*me.Result.Username,
			me.Result.FirstName,
		)

		// delete webhook (getting updates will not work when wehbook is set up)
		if unhooked := client.DeleteWebhook(true); unhooked.Ok {
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
