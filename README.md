# Telegram Bot API helper for Golang

This package is for building a simple Telegram Bot with or without webhook interface.

View the [documentation here](https://godoc.org/github.com/meinside/telegram-bot-go).

## Install

```
$ go get github.com/meinside/telegram-bot-go
```

## Generate a self-signed certificate (when using incoming webhook)

Generate self-signed certificate & private key with following script:

```bash
#!/bin/bash
#
# Generate self-signed certificates for Telegram Bot API
# (https://core.telegram.org/bots/self-signed)
# 
# created on : 2015.10.08.
# last update: 2015.10.11.
# 
# by meinside@gmail.com

DOMAIN="my.host.com"	# XXX - edit
EXPIRE_IN=3650	# XXX - edit

NUM_BITS=2048	# 2048 bits
C="US"
ST="New York"
L="Brooklyn"
O="Example Company"

PRIVATE_KEY="cert.key"
CERT_PEM="cert.pem"

openssl req -newkey rsa:$NUM_BITS -sha256 -nodes -keyout $PRIVATE_KEY -x509 -days $EXPIRE_IN -out $CERT_PEM -subj "/C=$C/ST=$ST/L=$L/O=$O/CN=$DOMAIN"

# finished
echo "> finished - certificate: $CERT_PEM / private key: $PRIVATE_KEY"
```

Generated *cert.key* and *cert.pem* file will be used in **telegrambot.StartWebhookServerAndWait()** function.

Also, you can generate certificate and private key using **telegrambot.GenCertAndKey()** function.

## Usage: with incoming webhook

```go
// sample code for telegram-bot-go (receive webhooks), last update: 2016.04.14.
package main

import (
	"fmt"
	"log"
	"time"

	bot "github.com/meinside/telegram-bot-go"
)

const (
	ApiToken     = "01234567:abcdefghijklmn_ABCDEFGHIJKLMNOPQRST"
	WebhookHost  = "my.host.com"
	WebhookPort  = 8443
	CertFilepath = "./cert.pem"
	KeyFilepath  = "./cert.key"

	TypingDelaySeconds = 3

	Verbose = true
)

func main() {
	client := bot.NewClient(ApiToken)
	client.Verbose = Verbose

	// get info about this bot
	if me := client.GetMe(); me.Ok {
		log.Printf("Bot information: @%s (%s)\n", *me.Result.Username, *me.Result.FirstName)

		// delete webhook
		if unhooked := client.DeleteWebhook(); unhooked.Ok {
			// generate certificate and private key for testing
			if err := bot.GenCertAndKey(WebhookHost, CertFilepath, KeyFilepath, 10*365); err == nil {
				// set webhook
				if hooked := client.SetWebhook(WebhookHost, WebhookPort, CertFilepath); hooked.Ok {
					// on success, start webhook server
					client.StartWebhookServerAndWait(CertFilepath, KeyFilepath, func(b *bot.Bot, webhook bot.Update, err error) {
						if err == nil {
							if webhook.HasMessage() {
								// 'is typing...'
								b.SendChatAction(webhook.Message.Chat.Id, bot.ChatActionTyping)
								time.Sleep(TypingDelaySeconds * time.Second)

								var message string

								key1 := "Just a button"
								key2 := "Request contact"
								key3 := "Request location"
								options := map[string]interface{}{
									"reply_to_message_id": webhook.Message.MessageId, // show original message
									"reply_markup": bot.ReplyKeyboardMarkup{ // show keyboards
										Keyboard: [][]bot.KeyboardButton{
											[]bot.KeyboardButton{
												bot.KeyboardButton{
													Text: &key1, // string only
												},
											},
											[]bot.KeyboardButton{
												bot.KeyboardButton{
													Text:           &key2,
													RequestContact: true, // request contact
												},
												bot.KeyboardButton{
													Text:            &key3,
													RequestLocation: true, // request location
												},
											},
										},
									},
								}

								if webhook.Message.HasContact() {
									message = fmt.Sprintf("I received @%s's phone no.: %s", *webhook.Message.From.Username, *webhook.Message.Contact.PhoneNumber)
								} else if webhook.Message.HasLocation() {
									message = fmt.Sprintf("I received @%s's location: (%f, %f)", *webhook.Message.From.Username, webhook.Message.Location.Latitude, webhook.Message.Location.Longitude)
								} else {
									if webhook.Message.HasText() {
										message = fmt.Sprintf("I received @%s's message: %s", *webhook.Message.From.Username, *webhook.Message.Text)
									} else {
										message = fmt.Sprintf("I received @%s's message", *webhook.Message.From.Username)
									}

								}
								// send message
								if sent := b.SendMessage(webhook.Message.Chat.Id, &message, options); !sent.Ok {
									log.Printf("*** failed to send message: %s\n", *sent.Description)
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

								results := []interface{}{
									article1,
									article2,
								}

								// answer inline query
								if sent := b.AnswerInlineQuery(*webhook.InlineQuery.Id, results, nil); !sent.Ok {
									log.Printf("*** failed to answer inline query: %s\n", *sent.Description)
								}
							}
						} else {
							log.Printf("*** error while receiving webhook (%s)\n", err.Error())
						}
					})
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
```

## Usage: without webhook

It would be useful when you're behind a firewall or something.

```go
// sample code for telegram-bot-go (get updates), last update: 2016.04.14.
package main

import (
	"fmt"
	"log"
	"time"

	bot "github.com/meinside/telegram-bot-go"
)

const (
	ApiToken = "01234567:abcdefghijklmn_ABCDEFGHIJKLMNOPQRST"

	MonitorIntervalSeconds = 3
	TypingDelaySeconds     = 3

	Verbose = true
)

func main() {
	client := bot.NewClient(ApiToken)
	client.Verbose = Verbose

	// get info about this bot
	if me := client.GetMe(); me.Ok {
		log.Printf("Bot information: @%s (%s)\n", *me.Result.Username, *me.Result.FirstName)

		// delete webhook (getting updates will not work when wehbook is set up)
		if unhooked := client.DeleteWebhook(); unhooked.Ok {
			// wait for new updates
			client.StartMonitoringUpdates(0, MonitorIntervalSeconds, func(b *bot.Bot, update bot.Update, err error) {
				if err == nil {
					if update.HasMessage() {
						// 'is typing...'
						b.SendChatAction(update.Message.Chat.Id, bot.ChatActionTyping)
						time.Sleep(TypingDelaySeconds * time.Second)

						var message string

						key1 := "Just a button"
						key2 := "Request contact"
						key3 := "Request location"
						options := map[string]interface{}{
							"reply_to_message_id": update.Message.MessageId, // show original message
							"reply_markup": bot.ReplyKeyboardMarkup{ // show keyboards
								Keyboard: [][]bot.KeyboardButton{
									[]bot.KeyboardButton{
										bot.KeyboardButton{
											Text: &key1, // string only
										},
									},
									[]bot.KeyboardButton{
										bot.KeyboardButton{
											Text:           &key2,
											RequestContact: true, // request contact
										},
										bot.KeyboardButton{
											Text:            &key3,
											RequestLocation: true, // request location
										},
									},
								},
							},
						}

						if update.Message.HasContact() {
							message = fmt.Sprintf("I received @%s's phone no.: %s", *update.Message.From.Username, *update.Message.Contact.PhoneNumber)
						} else if update.Message.HasLocation() {
							message = fmt.Sprintf("I received @%s's location: (%f, %f)", *update.Message.From.Username, update.Message.Location.Latitude, update.Message.Location.Longitude)
						} else {
							if update.Message.HasText() {
								message = fmt.Sprintf("I received @%s's message: %s", *update.Message.From.Username, *update.Message.Text)
							} else {
								message = fmt.Sprintf("I received @%s's message", *update.Message.From.Username)
							}

						}
						// send message
						if sent := b.SendMessage(update.Message.Chat.Id, &message, options); !sent.Ok {
							log.Printf("*** failed to send message: %s\n", *sent.Description)
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
						if sent := b.AnswerInlineQuery(*update.InlineQuery.Id, results, nil); !sent.Ok {
							log.Printf("*** failed to answer inline query: %s\n", *sent.Description)
						}
					}
				} else {
					log.Printf("*** error while receiving update (%s)\n", err.Error())
				}
			})
		} else {
			panic("failed to delete webhook")
		}
	} else {
		panic("failed to get info of the bot")
	}
}
```

## License

Copyright (c) 2016 Sungjin Han

MIT License

Permission is hereby granted, free of charge, to any person obtaining
a copy of this software and associated documentation files (the
"Software"), to deal in the Software without restriction, including
without limitation the rights to use, copy, modify, merge, publish,
distribute, sublicense, and/or sell copies of the Software, and to
permit persons to whom the Software is furnished to do so, subject to
the following conditions:

The above copyright notice and this permission notice shall be
included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

