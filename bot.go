// Package telegrambot / Telegram Bot API helper
//
// https://core.telegram.org/bots/api
//
// Created on : 2015.10.06, meinside@duck.com
package telegrambot

import (
	"crypto/md5"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"
)

const (
	apiBaseURL  = "https://api.telegram.org/bot"
	fileBaseURL = "https://api.telegram.org/file/bot"

	webhookPath = "/telegram/bot/webhook"
)

const (
	redactedString = "<REDACTED>" // confidential info will be displayed as this
)

// loggers
var (
	_stdout = log.New(os.Stdout, "", log.LstdFlags)
	_stderr = log.New(os.Stderr, "", log.LstdFlags)
)

// Bot struct
type Bot struct {
	token       string // Telegram bot API's token
	tokenHashed string // hashed token

	webhookHost string // webhook hostname
	webhookPort int    // webhook port number
	webhookURL  string // webhook url

	httpClient *http.Client // http client

	quitLoop chan struct{} // quit channel of polling loop

	// manual update handler - must be set
	updateHandler func(b *Bot, update Update, err error)

	// update handlers by content type (if not set, update will be passed to `updateHandler`)
	messageHandler            func(b *Bot, update Update, message Message, edited bool)
	mediaGroupHandler         func(b *Bot, updates []Update, mediaGroupID string)
	channelPostHandler        func(b *Bot, update Update, channelPost Message, edited bool)
	inlineQueryHandler        func(b *Bot, update Update, inlineQuery InlineQuery)
	chosenInlineResultHandler func(b *Bot, update Update, chosenInlineResult ChosenInlineResult)
	callbackQueryHandler      func(b *Bot, update Update, callbackQuery CallbackQuery)
	shippingQueryHandler      func(b *Bot, update Update, shippingQuery ShippingQuery)
	preCheckoutQueryHandler   func(b *Bot, update Update, preCheckoutQuery PreCheckoutQuery)
	pollHandler               func(b *Bot, update Update, poll Poll)
	pollAnswerHandler         func(b *Bot, update Update, pollAnswer PollAnswer)
	chatMemberUpdateHandler   func(b *Bot, update Update, memberUpdated ChatMemberUpdated, isMine bool)
	chatJoinRequestHandler    func(b *Bot, update Update, chatJoinRequest ChatJoinRequest)

	// command handlers (if not set, update will be passed to `updateHandler`)
	commandHandlers          map[string](func(b *Bot, update Update, args string)) // command handler functions
	noMatchingCommandHandler func(b *Bot, update Update, cmd, args string)         // handler function for no matching command

	Verbose bool // print verbose log messages or not
}

// NewClient gets a new bot API client with given token string.
func NewClient(token string) *Bot {
	client := Bot{
		token:       token,
		tokenHashed: fmt.Sprintf("%x", md5.Sum([]byte(token))),

		httpClient: nil,

		quitLoop: make(chan struct{}, 1),
	}

	// FIXME: (wasm) with DialContext, HTTP requests fail with "dial tcp: lookup api.telegram.org: Protocol not available"
	if runtime.GOARCH == "wasm" {
		client.httpClient = &http.Client{
			Timeout: 10 * time.Second,
		}
	} else {
		client.httpClient = &http.Client{
			Transport: &http.Transport{
				DialContext: (&net.Dialer{
					Timeout:   10 * time.Second,
					KeepAlive: 300 * time.Second,
				}).DialContext,
				IdleConnTimeout:       90 * time.Second,
				TLSHandshakeTimeout:   10 * time.Second,
				ResponseHeaderTimeout: 10 * time.Second,
				ExpectContinueTimeout: 1 * time.Second,
			},
		}
	}

	return &client
}

// GenCertAndKey generates a certificate and a private key file with given domain.
// (`OpenSSL` is needed.)
func GenCertAndKey(
	domain string,
	outCertFilepath string,
	outKeyFilepath string,
	expiresInDays int,
) error {
	numBits := 2048
	country := "US"
	state := "New York"
	local := "Brooklyn"
	org := "Example Company"

	if _, err := exec.Command("openssl", "req", "-newkey", fmt.Sprintf("rsa:%d", numBits), "-sha256", "-nodes", "-keyout", outKeyFilepath, "-x509", "-days", strconv.Itoa(expiresInDays), "-out", outCertFilepath, "-subj", fmt.Sprintf("/C=%s/ST=%s/L=%s/O=%s/CN=%s", country, state, local, org, domain)).Output(); err != nil {
		return err
	}

	return nil
}

// AddCommandHandler adds a handler function for given command.
func (b *Bot) AddCommandHandler(
	command string,
	handler func(b *Bot, update Update, args string),
) {
	// initialize map
	if b.commandHandlers == nil {
		b.commandHandlers = map[string]func(b *Bot, update Update, args string){}
	}

	// prepend '/'
	if !strings.HasPrefix(command, "/") {
		command = "/" + command
	}

	b.commandHandlers[command] = handler
}

// SetNoMatchingCommandHandler sets a function for handling no-matching commands.
func (b *Bot) SetNoMatchingCommandHandler(
	handler func(b *Bot, update Update, cmd, args string),
) {
	b.noMatchingCommandHandler = handler
}

// SetMessageHandler sets a function for handling messages.
func (b *Bot) SetMessageHandler(
	handler func(b *Bot, update Update, message Message, edited bool),
) {
	b.messageHandler = handler
}

// SetMediaGroupHandler sets a function for handling updates with media group id.
func (b *Bot) SetMediaGroupHandler(
	handler func(b *Bot, updates []Update, mediaGroupID string),
) {
	b.mediaGroupHandler = handler
}

// SetChannelPostHandler sets a function for handling channel posts.
func (b *Bot) SetChannelPostHandler(
	handler func(b *Bot, update Update, channelPost Message, edited bool),
) {
	b.channelPostHandler = handler
}

// SetInlineQueryHandler sets a function for handling inline queries.
func (b *Bot) SetInlineQueryHandler(
	handler func(b *Bot, update Update, inlineQuery InlineQuery),
) {
	b.inlineQueryHandler = handler
}

// SetChosenInlineResultHandler sets a function for handling chosen inline results.
func (b *Bot) SetChosenInlineResultHandler(
	handler func(b *Bot, update Update, chosenInlineResult ChosenInlineResult),
) {
	b.chosenInlineResultHandler = handler
}

// SetCallbackQueryHandler sets a function for handling callback queries.
func (b *Bot) SetCallbackQueryHandler(
	handler func(b *Bot, update Update, callbackQuery CallbackQuery),
) {
	b.callbackQueryHandler = handler
}

// SetShippingQueryHandler sets a function for handling shipping queries.
func (b *Bot) SetShippingQueryHandler(
	handler func(b *Bot, update Update, shippingQuery ShippingQuery),
) {
	b.shippingQueryHandler = handler
}

// SetPreCheckoutQueryHandler sets a function for handling pre-checkout queries.
func (b *Bot) SetPreCheckoutQueryHandler(
	handler func(b *Bot, update Update, preCheckoutQuery PreCheckoutQuery),
) {
	b.preCheckoutQueryHandler = handler
}

// SetPollHandler sets a function for handling polls.
func (b *Bot) SetPollHandler(
	handler func(b *Bot, update Update, poll Poll),
) {
	b.pollHandler = handler
}

// SetPollAnswerHandler sets a function for handling poll answers.
func (b *Bot) SetPollAnswerHandler(
	handler func(b *Bot, update Update, pollAnswer PollAnswer),
) {
	b.pollAnswerHandler = handler
}

// SetChatMemberUpdateHandler sets a function for handling chat member updates.
func (b *Bot) SetChatMemberUpdateHandler(
	handler func(b *Bot, update Update, memberUpdated ChatMemberUpdated, isMine bool),
) {
	b.chatMemberUpdateHandler = handler
}

// SetChatJoinRequestHandler sets a function for handling chat join requests.
func (b *Bot) SetChatJoinRequestHandler(
	handler func(b *Bot, update Update, chatJoinRequest ChatJoinRequest),
) {
	b.chatJoinRequestHandler = handler
}

// StartWebhookServerAndWait starts a webhook server(and waits forever).
// Function SetWebhook(host, port, certFilepath) should be called priorly to setup host, port, and certification file.
// Certification file(.pem) and a private key is needed.
// Incoming webhooks will be received through webhookHandler function.
//
// https://core.telegram.org/bots/self-signed
func (b *Bot) StartWebhookServerAndWait(
	certFilepath string,
	keyFilepath string,
	webhookHandler func(b *Bot, webhook Update, err error),
) {
	b.verbose("starting webhook server on: %s (port: %d) ...", b.getWebhookPath(), b.webhookPort)

	// set update handler
	if webhookHandler == nil {
		b.error("given webhook handler is nil")
		return
	}
	b.updateHandler = webhookHandler

	// routing
	mux := http.NewServeMux()
	mux.HandleFunc(b.getWebhookPath(), b.handleWebhook)

	// TODO: check http header: `X-Telegram-Bot-Api-Secret-Token` if `secret_token` is provided
	// (https://core.telegram.org/bots/api#setwebhook)

	// start server
	server := &http.Server{
		Addr:              fmt.Sprintf(":%d", b.webhookPort),
		Handler:           mux,
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       60 * time.Second,
	}
	if err := server.ListenAndServeTLS(certFilepath, keyFilepath); err != nil {
		panic(err.Error())
	}
}

// StartPollingUpdates retrieves updates from API server constantly, synchronously.
//
// `optionalParams` can be:
//   - []AllowedUpdates
//
// NOTE: Make sure webhook is deleted, or not registered before polling.
func (b *Bot) StartPollingUpdates(
	updateOffset int64,
	interval int,
	updateHandler func(b *Bot, update Update, err error),
	optionalParams ...any,
) {
	b.verbose("starting polling updates (interval seconds: %d) ...", interval)

	// https://core.telegram.org/bots/api#getupdates
	options := OptionsGetUpdates{}.
		SetOffset(updateOffset).
		SetLimit(100). // default: 100
		SetTimeout(1)  // default: 0 for testing

	// iterate optional params and apply to options
	for _, param := range optionalParams {
		if allowedUpdates, ok := param.([]AllowedUpdate); ok {
			options = options.SetAllowedUpdates(allowedUpdates)
		}
	}

	// set update handler
	if updateHandler == nil {
		b.error("given update handler is nil")
		return
	}
	b.updateHandler = updateHandler

	var updates APIResponse[[]Update]
loop:
	for {
		select {
		case <-b.quitLoop:
			break loop
		default:
			if updates = b.GetUpdates(options); updates.Ok {
				// update offset (max + 1)
				for _, update := range *updates.Result {
					if options["offset"].(int64) <= update.UpdateID {
						options["offset"] = update.UpdateID + 1
					}
				}

				if b.mediaGroupHandler != nil {
					// group updates by media group id,
					groups := groupUpdatesByMediaGroupId(*updates.Result)

					// handle updates by group id
					for groupID, groupedUpdates := range groups {
						if groupID == "" { // NOTE: no group id
							for _, update := range groupedUpdates {
								// if there is a matching command, handle it as a command,
								if !handleUpdateAsCommand(b, update) {
									// if it was not handled as a command, handle it by type:
									if !handleUpdateByType(b, update) {
										// otherwise, handle it manually
										go b.updateHandler(b, update, nil)
									}
								}
							}
						} else { // with group id
							go b.mediaGroupHandler(b, groupedUpdates, groupID)
						}
					}
				} else {
					// ordinary handling of updates
					for _, update := range *updates.Result {
						// if there is a matching command, handle it as a command,
						if !handleUpdateAsCommand(b, update) {
							// if it was not handled as a command, handle it by type:
							if !handleUpdateByType(b, update) {
								// otherwise, handle it manually
								go b.updateHandler(b, update, nil)
							}
						}
					}
				}
			} else {
				go b.updateHandler(b, Update{}, fmt.Errorf("%s", *updates.Description))
			}

			time.Sleep(time.Duration(interval) * time.Second)
		}
	}

	b.verbose("stopped polling updates")
}

// DEPRECATED: renamed to `StartPollingUpdates`
func (b *Bot) StartMonitoringUpdates(
	updateOffset int64,
	interval int,
	updateHandler func(b *Bot, update Update, err error),
) {
	b.StartPollingUpdates(updateOffset, interval, updateHandler)
}

// group updates by their media group id
func groupUpdatesByMediaGroupId(updates []Update) (groups map[string][]Update) {
	groups = map[string][]Update{}

	var mediaGroupID string
	for _, update := range updates {
		if update.HasMediaGroup() {
			mediaGroupID = *update.MediaGroupID()
		} else {
			mediaGroupID = "" // NOTE: no group id
		}

		if group, exists := groups[mediaGroupID]; exists {
			groups[mediaGroupID] = append(group, update)
		} else {
			groups[mediaGroupID] = []Update{
				update,
			}
		}
	}

	return groups
}

// checks if given update matches any command and handle it (returns true if handled)
func handleUpdateAsCommand(b *Bot, update Update) bool {
	var message Message
	if update.HasMessage() {
		message = *update.Message
	} else if update.HasEditedMessage() {
		message = *update.EditedMessage
	} else {
		// if it doesn't have a user message, do not handle it
		return false
	}

	// if it doesn't have any text, do not handle it
	if !message.HasText() {
		return false
	}

	txt := *message.Text

	// if a messsage doesn't start with '/', it is not a command
	if !strings.HasPrefix(txt, "/") {
		return false
	}

	command := strings.Split(txt, " ")[0]
	params := strings.TrimSpace(strings.TrimPrefix(txt, command))

	for cmd, cmdHandler := range b.commandHandlers {
		if command == cmd {
			go cmdHandler(b, update, params)

			return true
		}
	}

	// if no-matching-command-handler is set, handle with it
	if b.noMatchingCommandHandler != nil {
		go b.noMatchingCommandHandler(b, update, command, params)

		return true
	}

	return false
}

// checks if given update matches any registered handler by type and handle it (returns true if handled)
func handleUpdateByType(b *Bot, update Update) bool {
	// if it was not handled as a command, handle it by type:
	if b.messageHandler != nil && (update.HasMessage() || update.HasEditedMessage()) {
		var message Message
		if update.HasMessage() {
			message = *update.Message
		} else if update.HasEditedMessage() {
			message = *update.EditedMessage
		}

		go b.messageHandler(b, update, message, update.HasEditedMessage())

		return true
	} else if b.channelPostHandler != nil && (update.HasChannelPost() || update.HasEditedChannelPost()) {
		var channelPost Message
		if update.HasChannelPost() {
			channelPost = *update.ChannelPost
		} else if update.HasEditedMessage() {
			channelPost = *update.EditedChannelPost
		}

		go b.channelPostHandler(b, update, channelPost, update.HasEditedChannelPost())

		return true
	} else if b.inlineQueryHandler != nil && update.HasInlineQuery() {
		go b.inlineQueryHandler(b, update, *update.InlineQuery)

		return true
	} else if b.chosenInlineResultHandler != nil && update.HasChosenInlineResult() {
		go b.chosenInlineResultHandler(b, update, *update.ChosenInlineResult)

		return true
	} else if b.callbackQueryHandler != nil && update.HasCallbackQuery() {
		go b.callbackQueryHandler(b, update, *update.CallbackQuery)

		return true
	} else if b.shippingQueryHandler != nil && update.HasShippingQuery() {
		go b.shippingQueryHandler(b, update, *update.ShippingQuery)

		return true
	} else if b.preCheckoutQueryHandler != nil && update.HasPreCheckoutQuery() {
		go b.preCheckoutQueryHandler(b, update, *update.PreCheckoutQuery)

		return true
	} else if b.pollHandler != nil && update.HasPoll() {
		go b.pollHandler(b, update, *update.Poll)

		return true
	} else if b.pollAnswerHandler != nil && update.HasPollAnswer() {
		go b.pollAnswerHandler(b, update, *update.PollAnswer)

		return true
	} else if b.chatMemberUpdateHandler != nil && (update.HasMyChatMember() || update.HasChatMember()) {
		var chatMemberUpdated ChatMemberUpdated
		if update.HasMyChatMember() {
			chatMemberUpdated = *update.MyChatMember
		} else if update.HasChatMember() {
			chatMemberUpdated = *update.ChatMember
		}

		go b.chatMemberUpdateHandler(b, update, chatMemberUpdated, update.HasMyChatMember())

		return true
	} else if b.chatJoinRequestHandler != nil && update.HasChatJoinRequest() {
		go b.chatJoinRequestHandler(b, update, *update.ChatJoinRequest)

		return true
	}

	return false
}

// StopPollingUpdates stops loop of polling updates
func (b *Bot) StopPollingUpdates() {
	b.verbose("stopping polling updates...")

	b.quitLoop <- struct{}{}
}

// DEPRECATED: renamed to `StopPollingUpdates`
func (b *Bot) StopMonitoringUpdates() {
	b.StopPollingUpdates()
}

// Get webhook path generated with hash.
func (b *Bot) getWebhookPath() string {
	return fmt.Sprintf("%s/%s", webhookPath, b.tokenHashed)
}

// Get full URL of webhook interface.
func (b *Bot) getWebhookURL() string {
	return fmt.Sprintf("https://%s:%d%s", b.webhookHost, b.webhookPort, b.getWebhookPath())
}

// Remove confidential info from given string.
func (b *Bot) redact(str string) string {
	tokenRemoved := strings.ReplaceAll(str, b.token, redactedString)
	redacted := strings.ReplaceAll(tokenRemoved, b.tokenHashed, redactedString)
	return redacted
}

// Print formatted log message. (only when Bot.Verbose == true)
func (b *Bot) verbose(str string, args ...any) {
	if b.Verbose {
		_stdout.Printf("%s\n", b.redact(fmt.Sprintf(str, args...)))
	}
}

// Print formatted error message.
func (b *Bot) error(str string, args ...any) {
	_stderr.Printf("%s\n", b.redact(fmt.Sprintf(str, args...)))
}
