// Telegram Bot API helper (https://core.telegram.org/bots/api)
//
// Created on : 2015.10.06, meinside@gmail.com
package telegrambot

import (
	"crypto/md5"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
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
	token       string // Telegram bot API's token
	tokenHashed string // hashed token

	webhookHost string // webhook hostname
	webhookPort int    // webhook port number
	webhookUrl  string // webhook url

	updateHandler func(b *Bot, update Update, err error) // update(webhook) handler function

	Verbose bool // print verbose log messages or not
}

// Get a new bot API client with given token string.
func NewClient(token string) *Bot {
	return &Bot{
		token:       token,
		tokenHashed: fmt.Sprintf("%x", md5.Sum([]byte(token))),
	}
}

// Generate certificate and private key file with given domain.
// OpenSSL is needed.
func GenCertAndKey(domain string, outCertFilepath string, outKeyFilepath string, expiresInDays int) error {
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

// Set webhook url and certificate for receiving incoming updates.
// port should be one of: 443, 80, 88, or 8443.
//
// https://core.telegram.org/bots/api#setwebhook
func (b *Bot) SetWebhook(host string, port int, certFilepath string) (result ApiResult) {
	b.webhookHost = host
	b.webhookPort = port
	b.webhookUrl = b.getWebhookUrl()

	file, err := os.Open(certFilepath)
	if err != nil {
		panic(err.Error())
	}

	params := map[string]interface{}{
		"url":         b.webhookUrl,
		"certificate": file,
	}

	b.verbose("setting webhook url to: %s", b.webhookUrl)

	return b.requestResult("setWebhook", params)
}

// Delete webhook.
// (Function GetUpdates will not work if webhook is set, so in that case you'll need to delete it)
//
// https://core.telegram.org/bots/api#setwebhook
func (b *Bot) DeleteWebhook() (result ApiResult) {
	b.webhookHost = ""
	b.webhookPort = 0
	b.webhookUrl = ""

	params := map[string]interface{}{
		"url": "",
	}

	b.verbose("deleting webhook url")

	return b.requestResult("setWebhook", params)
}

// Start Webhook server(and wait forever).
// Function SetWebhook(host, port, certFilepath) should be called priorly to setup host, port, and certification file.
// Certification file(.pem) and a private key is needed.
// Incoming webhooks will be received through webhookHandler function.
//
// https://core.telegram.org/bots/self-signed
func (b *Bot) StartWebhookServerAndWait(certFilepath string, keyFilepath string, webhookHandler func(b *Bot, webhook Update, err error)) {
	b.verbose("starting webhook server on: %s (port: %d) ...", b.getWebhookPath(), b.webhookPort)

	// set update handler
	b.updateHandler = webhookHandler

	// routing
	http.HandleFunc(b.getWebhookPath(), b.handleWebhook)

	// start server
	if err := http.ListenAndServeTLS(fmt.Sprintf(":%d", b.webhookPort), certFilepath, keyFilepath, nil); err != nil {
		panic(err.Error())
	}
}

// Retrieve updates from API server constantly.
// If webhook is registered, it may not work properly. So make sure webhook is deleted, or not registered.
func (b *Bot) StartMonitoringUpdates(updateOffset int, interval int, updateHandler func(b *Bot, update Update, err error)) {
	b.verbose("starting monitoring updates (interval seconds: %d) ...", interval)

	options := map[string]interface{}{
		"offset": updateOffset,
	}

	// set update handler
	b.updateHandler = updateHandler

	var updates ApiResultUpdates
	for {
		if updates = b.GetUpdates(options); updates.Ok {
			for _, update := range updates.Result {
				// update offset (max + 1)
				if options["offset"].(int) <= update.UpdateId {
					options["offset"] = update.UpdateId + 1
				}

				go b.updateHandler(b, update, nil)
			}
		} else {
			go b.updateHandler(b, Update{}, fmt.Errorf("error while retrieving updates - %s", *updates.Description))
		}

		time.Sleep(time.Duration(interval) * time.Second)
	}
}

// Get webhook path generated with hash.
func (b *Bot) getWebhookPath() string {
	return fmt.Sprintf("%s/%s", WebhookPath, b.tokenHashed)
}

// Get full URL of webhook interface.
func (b *Bot) getWebhookUrl() string {
	return fmt.Sprintf("https://%s:%d%s", b.webhookHost, b.webhookPort, b.getWebhookPath())
}

// Remove confidential info from given string.
func (b *Bot) redact(str string) string {
	tokenRemoved := strings.Replace(str, b.token, RedactedString, -1)
	redacted := strings.Replace(tokenRemoved, b.tokenHashed, RedactedString, -1)
	return redacted
}

// Print formatted log message. (only when Bot.Verbose == true)
func (b *Bot) verbose(str string, args ...interface{}) {
	if b.Verbose {
		log.Printf("> %s\n", b.redact(fmt.Sprintf(str, args...)))
	}
}

// Print formatted error message.
func (b *Bot) error(str string, args ...interface{}) {
	log.Printf("* %s\n", b.redact(fmt.Sprintf(str, args...)))
}
