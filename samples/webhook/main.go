// sample code for telegram-bot-go (receive webhooks),
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
	apiToken     = "01234567:abcdefghijklmn_ABCDEFGHIJKLMNOPQRST"
	webhookHost  = "my.host.com"
	webhookPort  = 8443
	certFilepath = "./cert.pem"
	keyFilepath  = "./cert.key"

	typingDelaySeconds = 1

	verbose = true
)

// function for handling webhook updates
func webhookHandler(b *bot.Bot, webhook bot.Update, err error) {
	if err == nil {
		if webhook.HasMessage() {
			// 'is typing...'
			b.SendChatAction(
				webhook.Message.Chat.ID,
				bot.ChatActionTyping,
				nil,
			)
			time.Sleep(typingDelaySeconds * time.Second)

			var message string

			if webhook.Message.HasContact() {
				message = fmt.Sprintf(
					"I received @%s's phone no.: %s",
					*webhook.Message.From.Username,
					webhook.Message.Contact.PhoneNumber,
				)
			} else if webhook.Message.HasLocation() {
				message = fmt.Sprintf(
					"I received @%s's location: (%f, %f)",
					*webhook.Message.From.Username,
					webhook.Message.Location.Latitude,
					webhook.Message.Location.Longitude,
				)
			} else {
				if webhook.Message.HasText() {
					message = fmt.Sprintf(
						"I received @%s's message: %s",
						*webhook.Message.From.Username,
						*webhook.Message.Text,
					)
				} else {
					message = fmt.Sprintf(
						"I received @%s's message",
						*webhook.Message.From.Username,
					)
				}
			}
			// send message
			if sent := b.SendMessage(
				webhook.Message.Chat.ID,
				message,
				bot.OptionsSendMessage{}.
					SetReplyParameters(bot.NewReplyParameters(webhook.Message.MessageID)). // show original message
					SetReplyMarkup(bot.NewReplyKeyboardMarkup(                             // show keyboards
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
		} else if webhook.HasInlineQuery() {
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
				webhook.InlineQuery.ID,
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
			"*** error while receiving webhook (%s)",
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

		// delete webhook
		if unhooked := client.DeleteWebhook(true); unhooked.Ok {
			// generate certificate and private key for testing
			if err := bot.GenCertAndKey(
				webhookHost,
				certFilepath,
				keyFilepath,
				10*365,
			); err == nil {
				// set webhook
				if hooked := client.SetWebhook(
					webhookHost,
					webhookPort,
					bot.OptionsSetWebhook{}.
						SetCertificate(certFilepath),
				); hooked.Ok {
					// on success, start webhook server
					client.StartWebhookServerAndWait(
						certFilepath,
						keyFilepath,
						webhookHandler,
					)
				} else {
					panic("failed to set webhook")
				}
			} else {
				panic("failed to generate cert/key: " + err.Error())
			}
		} else {
			panic("failed to delete webhook")
		}
	} else {
		panic("failed to get info of the bot")
	}
}
