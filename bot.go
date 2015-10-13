// Telegram Bot API helper
//
// referenced: https://core.telegram.org/bots/api
//
// created on : 2015.10.06.
//
// by meinside@gmail.com

package telegrambot

import (
	"crypto/md5"
	"fmt"
	"net/http"
	"strings"
)

const (
	ApiBaseUrl         = "https://api.telegram.org/bot"
	FileBaseUrl        = "https://api.telegram.org/file/bot"
	DefaultWebhookPort = 443

	WebhookPath = "/telegram/bot/webhook"
)

const (
	RedactedString = "<REDACTED>" // confidential info will be displayed as this
)

// Bot
type Bot struct {
	// Telegram bot API's token
	Token string

	// Webhook related stuffs
	WebhookProtocol string
	WebhookHost     string
	WebhookPort     int
	WebhookUrl      string
	WebhookHandler  func(webhook Webhook, success bool, err error)

	// print verbose log messages or not
	Verbose bool
}

// Get a new bot API client with given token string.
func NewClient(token string) *Bot {
	return &Bot{
		Token: token,
	}
}

// Start Webhook server(and wait forever).
//
// Function SetWebhook(host, port, certFilepath) should be called priorly to setup host, port, and certification file.
//
// Certification file(.pem) and a private key is needed.
//
// (https://core.telegram.org/bots/self-signed)
//
// Incoming webhooks will be received through webhookHandler function.
func (b *Bot) StartWebhookServerAndWait(certFilepath string, keyFilepath string, webhookHandler func(webhook Webhook, success bool, err error)) {
	b.verbose("starting webhook server on: %s (port: %d) ...", b.getWebhookPath(), b.WebhookPort)

	// routing
	b.WebhookHandler = webhookHandler
	http.HandleFunc(b.getWebhookPath(), b.handleWebhook)

	// start server
	if err := http.ListenAndServeTLS(fmt.Sprintf(":%d", b.WebhookPort), certFilepath, keyFilepath, nil); err != nil {
		panic(err.Error())
	}
}

// Generate hash from token.
func (b *Bot) getHashedToken() string {
	return fmt.Sprintf("%x", md5.Sum([]byte(b.Token)))
}

// Get webhook path generated with hash.
func (b *Bot) getWebhookPath() string {
	return fmt.Sprintf("%s/%s", WebhookPath, b.getHashedToken())
}

// Get full URL of webhook interface.
func (b *Bot) getWebhookUrl() string {
	return fmt.Sprintf("https://%s:%d%s", b.WebhookHost, b.WebhookPort, b.getWebhookPath())
}

// Remove confidential info from given string.
func (b *Bot) redact(str string) string {
	tokenRemoved := strings.Replace(str, b.Token, RedactedString, -1)
	redacted := strings.Replace(tokenRemoved, b.getHashedToken(), RedactedString, -1)
	return redacted
}

// Print formatted log message. (only when Bot.Verbose == true)
func (b *Bot) verbose(str string, args ...interface{}) {
	if b.Verbose {
		fmt.Printf("> %s\n", b.redact(fmt.Sprintf(str, args...)))
	}
}

// Print formatted error message.
func (b *Bot) error(str string, args ...interface{}) {
	fmt.Printf("* %s\n", b.redact(fmt.Sprintf(str, args...)))
}
