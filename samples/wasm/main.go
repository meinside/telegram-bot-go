// sample code for telegram-bot-go (get updates),
//
// WASM version
//
// created on: 2018.11.19.
//
// +build: js,wasm

// $ GOOS=js GOARCH=wasm vi __FILENAME__

package main

import (
	"fmt"
	"log"
	"syscall/js"
	"time"

	bot "github.com/meinside/telegram-bot-go"
	wh "github.com/meinside/wasm-helper-go"
)

const (
	apiToken = "01234567:abcdefghijklmn_ABCDEFGHIJKLMNOPQRST"

	monitorIntervalSeconds = 1
	typingDelaySeconds     = 1

	//verbose = false
	verbose = true
)

var _wasmHelper *wh.WasmHelper
var _content js.Value

func init() {
	_wasmHelper = wh.New()
	_wasmHelper.SetVerbose(verbose)

	_content = _wasmHelper.Call("document.getElementById", "content")
}

// update handler function
func handleUpdate(b *bot.Bot, update bot.Update, err error) {
	if err == nil {
		if update.HasMessage() {
			// 'is typing...'
			b.SendChatAction(update.Message.Chat.ID, bot.ChatActionTyping, nil)

			// sleep for a while,
			time.Sleep(typingDelaySeconds * time.Second)

			var sender string = *update.Message.From.Username
			var message, fileURL string

			if update.Message.HasContact() {
				message = fmt.Sprintf(
					"Received @%s's phone no.: %s",
					sender,
					update.Message.Contact.PhoneNumber,
				)
			} else if update.Message.HasLocation() {
				message = fmt.Sprintf(
					"Received @%s's location: (%f, %f)",
					sender,
					update.Message.Location.Latitude,
					update.Message.Location.Longitude,
				)
			} else if update.Message.HasText() {
				message = fmt.Sprintf(
					"Received @%s's message: %s",
					sender,
					*update.Message.Text,
				)
			} else if update.Message.HasPhoto() {
				photo := update.Message.LargestPhoto()

				if fileRes := b.GetFile(photo.FileID); fileRes.Ok {
					fileURL = b.GetFileURL(*fileRes.Result)
				}

				if fileURL != "" {
					message = fmt.Sprintf(
						"Received @%s's photo (file url: %s)",
						sender,
						fileURL,
					)
				} else {
					message = fmt.Sprintf(
						"Received @%s's photo (file id: %s)",
						sender,
						photo.FileID,
					)
				}
			} else if update.Message.HasAnimation() {
				animation := update.Message.Animation

				if fileRes := b.GetFile(animation.FileID); fileRes.Ok {
					fileURL = b.GetFileURL(*fileRes.Result)
				}

				if fileURL != "" {
					message = fmt.Sprintf(
						"Received @%s's animation (file url: %s)",
						sender,
						fileURL,
					)
				} else {
					message = fmt.Sprintf(
						"Received @%s's animation (file id: %s)",
						sender,
						animation.FileID,
					)
				}
			} else {
				message = fmt.Sprintf(
					"Received @%s's message",
					sender,
				)
			}

			// append to the html,
			element := _wasmHelper.Call("document.createElement", "div")
			_wasmHelper.SetOn(element, "style", "margin: 5px")
			_wasmHelper.SetOn(element, "innerHTML", message)
			_wasmHelper.CallOn(_content, "appendChild", element)

			// and reply to the message
			if sent := b.SendMessage(
				update.Message.Chat.ID,
				message,
				// options
				bot.OptionsSendMessage{}.
					SetReplyToMessageID(update.Message.MessageID). // show original message
					SetReplyMarkup(bot.ReplyKeyboardMarkup{        // show keyboards
						ResizeKeyboard: true, // compact keyboard size
						Keyboard: [][]bot.KeyboardButton{
							[]bot.KeyboardButton{
								bot.KeyboardButton{
									Text:           "Send contact",
									RequestContact: true,
								},
								bot.KeyboardButton{
									Text:            "Send location",
									RequestLocation: true,
								},
							},
						},
					}),
			); !sent.Ok {
				log.Printf(
					"*** failed to send message: %s",
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

func main() {
	client := bot.NewClient(apiToken)
	client.Verbose = verbose

	// get info about this bot
	if me := client.GetMe(); me.Ok {
		botID := *me.Result.Username
		botName := me.Result.FirstName

		// set bot info on html,
		info := _wasmHelper.Call("document.createElement", "info")
		_wasmHelper.SetOn(
			info,
			"style",
			"margin: 10px; color: #0000FF; font-weight: bold;",
		)
		_wasmHelper.SetOn(
			info,
			"innerHTML",
			fmt.Sprintf("Connected to bot: <a href='https://telegram.me/%s' target='_blank'>@%s</a> (%s)", botID, botID, botName),
		)
		_wasmHelper.CallOn(_content, "appendChild", info)

		log.Printf(
			"Running bot: @%s (%s)",
			botID,
			botName,
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
