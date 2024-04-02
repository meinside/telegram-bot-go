//go:build js && wasm

// sample code for telegram-bot-go (get updates),
//
// Wasm version
//
// created on: 2024.01.04.

// NOTE: open related files with GOOS and GOARCH environment variables like:
//    `$ GOOS=js GOARCH=wasm vi __FILENAME__`

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
	pollingIntervalSeconds = 1
	typingDelaySeconds     = 1

	//verbose = false
	verbose = true
)

var _wasmHelper *wh.WasmHelper
var _content js.Value

func init() {
	_wasmHelper = wh.New()
	_wasmHelper.SetVerbose(verbose)

	_content, _ = _wasmHelper.Call("document.getElementById", "content")
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
			appendDiv(
				"message",
				message,
				"margin: 5px;",
				_content,
			)

			// and reply to the message
			if sent := b.SendMessage(
				update.Message.Chat.ID,
				message,
				// options
				bot.OptionsSendMessage{}.
					SetReplyParameters(bot.ReplyParameters{
						MessageID: update.Message.MessageID,
					}).                                                              // reply to the original message
					SetReplyMarkup(replyKeyboardMarkup(true, [][]bot.KeyboardButton{ // show keyboards
						{
							keyboardButton("Send contact", true, false),
							keyboardButton("Send location", false, true),
						},
					})),
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
			err,
		)
	}
}

// generate a reply keyboard markup
func replyKeyboardMarkup(resize bool, keyboards [][]bot.KeyboardButton) bot.ReplyKeyboardMarkup {
	return bot.ReplyKeyboardMarkup{
		ResizeKeyboard: &resize,
		Keyboard:       keyboards,
	}
}

// generate a keyboard button
func keyboardButton(text string, contact, location bool) bot.KeyboardButton {
	return bot.KeyboardButton{
		Text:            text,
		RequestContact:  &contact,
		RequestLocation: &location,
	}
}

func main() {
	// `runBot` will be exposed to js
	_wasmHelper.RegisterFunctions(map[string]wh.WasmFunction{
		"runBot": runBot,
	})

	_wasmHelper.Wait() // busy-wait
}

// this function will be called from `index.html`
func runBot(this js.Value, args []js.Value) any {
	apiToken := args[0].String()

	go func() {
		appendDiv(
			"start",
			"Launching bot...",
			"margin: 10px; color: #000000;",
			_content,
		)

		client := bot.NewClient(apiToken)
		client.Verbose = verbose

		// get info about this bot
		if me := client.GetMe(); me.Ok {
			botID := *me.Result.Username
			botName := me.Result.FirstName

			// set bot info on html,
			appendDiv(
				"info",
				fmt.Sprintf("Connected to bot: <a href='https://telegram.me/%s' target='_blank'>@%s</a> (%s)", botID, botID, botName),
				"margin: 10px; color: #0000FF; font-weight: bold;",
				_content,
			)

			log.Printf(
				"Launched bot: @%s (%s)",
				botID,
				botName,
			)

			// delete webhook (getting updates will not work when wehbook is set up)
			if unhooked := client.DeleteWebhook(true); unhooked.Ok {
				// wait for new updates
				client.StartMonitoringUpdates(
					0,
					pollingIntervalSeconds,
					handleUpdate,
				)
			} else {
				appendDiv(
					"error",
					*unhooked.Description,
					"margin: 10px; color: #FF0000;",
					_content,
				)

				panic("failed to delete webhook")
			}
		} else {
			appendDiv(
				"error",
				*me.Description,
				"margin: 10px; color: #FF0000;",
				_content,
			)

			panic("failed to get info of the bot")
		}
	}()

	return nil
}

// append a child div to a parent
func appendDiv(class, content, style string, parent js.Value) {
	if div, err := _wasmHelper.Call("document.createElement", "div"); err == nil {
		_ = _wasmHelper.SetOn(div, "class", class)
		_ = _wasmHelper.SetOn(div, "style", style)
		_ = _wasmHelper.SetOn(div, "innerHTML", content)
		if _, err := _wasmHelper.CallOn(parent, "appendChild", div); err != nil {
			log.Printf("failed to append div: %s", err)
		}
	} else {
		log.Printf("failed to create div: %s", err)
	}
}
