// Telegram Bot API helper
//
// referenced: https://core.telegram.org/bots/api
//
// created on : 2015.10.06.
//
// by meinside@gmail.com

package telegrambot

import (
	"bytes"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

const (
	ApiBaseUrl         = "https://api.telegram.org/bot"
	DefaultWebhookPort = 443
	WebhookPath        = "/telegram/bot/webhook"
)

const (
	RedactedString = "<REDACTED>"
)

type Bot struct {
	Token           string
	WebhookProtocol string
	WebhookHost     string
	WebhookPort     int
	WebhookUrl      string
	WebhookHandler  func(webhook Webhook, success bool, err error)
	Verbose         bool
}

// Get new bot API client
func NewClient(token string) *Bot {
	return &Bot{
		Token:       token,
		WebhookPort: DefaultWebhookPort,
	}
}

// Generate hash from token
func (b *Bot) getHashedToken() string {
	return fmt.Sprintf("%x", md5.Sum([]byte(b.Token)))
}

// Get webhook path generated with hash
func (b *Bot) getWebhookPath() string {
	return fmt.Sprintf("%s/%s", WebhookPath, b.getHashedToken())
}

// Get full URL of webhook interface
func (b *Bot) getWebhookUrl() string {
	return fmt.Sprintf("https://%s:%d%s", b.WebhookHost, b.WebhookPort, b.getWebhookPath())
}

// Remove confidential info from given string
func (b *Bot) redact(str string) string {
	tokenRemoved := strings.Replace(str, b.Token, RedactedString, -1)
	redacted := strings.Replace(tokenRemoved, b.getHashedToken(), RedactedString, -1)
	return redacted
}

// Print formatted log message (only when Bot.Verbose == true)
func (b *Bot) verbose(str string, args ...interface{}) {
	if b.Verbose {
		fmt.Printf("> %s\n", b.redact(fmt.Sprintf(str, args...)))
	}
}

// Print formatted error message
func (b *Bot) error(str string, args ...interface{}) {
	fmt.Printf("* %s\n", b.redact(fmt.Sprintf(str, args...)))
}

// Check if given http params contain file or not
func checkIfFileParamExists(params map[string]interface{}) bool {
	for _, value := range params {
		switch value.(type) {
		case *os.File:
			return true
		}
	}

	return false
}

// Convert given interface to string (for HTTP params)
func (b *Bot) paramToString(param interface{}) (result string, success bool) {
	switch param.(type) {
	case int:
		if intValue, ok := param.(int); ok {
			return strconv.Itoa(intValue), ok
		} else {
			b.error("parameter '%+v' could not be cast to int value", param)
		}
	case string:
		if strValue, ok := param.(string); ok {
			return strValue, ok
		} else {
			b.error("parameter '%+v' could not be cast to string value", param)
		}
	default:
		b.error("unexpected type: '%+v' (%T)", param, param)
	}

	return "", true
}

// Send request to API server and return the response (synchronous)
//
// @param method [string] HTTP method
// @param params [map[string]interface{}] request parameters
func (b *Bot) sendRequest(method string, params map[string]interface{}) (resp *http.Response, success bool) {
	client := &http.Client{}
	apiUrl := fmt.Sprintf("%s%s/%s", ApiBaseUrl, b.Token, method)

	b.verbose("sending request to api url: %s, params: %#v", apiUrl, params)

	if checkIfFileParamExists(params) {
		// multipart form data

		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)

		for key, value := range params {
			switch value.(type) {
			case *os.File:
				if fileValue, ok := value.(*os.File); ok {
					if part, err := writer.CreateFormFile(key, fileValue.Name()); err == nil {
						if _, err = io.Copy(part, fileValue); err != nil {
							b.error("could now write to multipart: %s", key)
						}

						defer fileValue.Close()
					} else {
						b.error("could not create form file for parameter '%s' (%v)", key, value)
					}
				} else {
					b.error("parameter '%s' (%v) could not be cast to file", key, value)
				}
			default:
				if strValue, ok := b.paramToString(value); ok {
					writer.WriteField(key, strValue)
				}
			}
		}

		if err := writer.Close(); err != nil {
			b.error("error while closing writer (%s)", err.Error())
		}

		if req, err := http.NewRequest("POST", apiUrl, body); err == nil {
			req.Header.Add("Content-Type", writer.FormDataContentType()) // due to file parameter

			if resp, err := client.Do(req); err == nil {
				return resp, true
			} else {
				b.error("request error: %s", err.Error())
			}
		} else {
			b.error("building request error: %s", err.Error())
		}
	} else {
		// www-form urlencoded

		paramValues := url.Values{}
		for key, value := range params {
			if strValue, ok := b.paramToString(value); ok {
				paramValues[key] = []string{strValue}
			}
		}
		encoded := paramValues.Encode()

		if req, err := http.NewRequest("POST", apiUrl, bytes.NewBufferString(encoded)); err == nil {
			req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
			req.Header.Add("Content-Length", strconv.Itoa(len(encoded)))

			if resp, err := client.Do(req); err == nil {
				return resp, true
			} else {
				b.error("request error: %s", err.Error())
			}
		} else {
			b.error("building request error: %s", err.Error())
		}
	}

	return nil, false
}

// Set webhook url for receiving incoming updates
//
// https://core.telegram.org/bots/api#setwebhook
//
// @param host [string] webhook server host
// @param port [int] webhook server port (443, 80, 88, or 8443)
// @param certFilepath [string] certification file's path
func (b *Bot) SetWebhookUrl(host string, port int, certFilepath string) (result ApiResult) {
	b.WebhookHost = host
	b.WebhookPort = port
	b.WebhookUrl = b.getWebhookUrl()

	file, err := os.Open(certFilepath)
	if err != nil {
		panic(err.Error())
	}

	params := map[string]interface{}{
		"url":         b.WebhookUrl,
		"certificate": file,
	}

	b.verbose("setting webhook url to: %s", b.WebhookUrl)

	var errStr string

	if resp, success := b.sendRequest("setWebhook", params); success {
		defer resp.Body.Close()

		if body, err := ioutil.ReadAll(resp.Body); err == nil {
			var jsonResponse ApiResult
			if err := json.Unmarshal(body, &jsonResponse); err == nil {
				return jsonResponse
			} else {
				errStr = fmt.Sprintf("json parse error: %s (%s)", err.Error(), string(body))
			}
		} else {
			errStr = fmt.Sprintf("response read error: %s", err.Error())
		}
	}

	b.error(errStr)

	return ApiResult{Ok: false, Description: errStr}
}

// Delete webhook url
//
// https://core.telegram.org/bots/api#setwebhook
//
func (b *Bot) DeleteWebhookUrl() (result ApiResult) {
	b.WebhookHost = ""
	b.WebhookUrl = ""

	params := map[string]interface{}{
		"url": "",
	}

	var errStr string

	if resp, success := b.sendRequest("setWebhook", params); success {
		defer resp.Body.Close()

		if body, err := ioutil.ReadAll(resp.Body); err == nil {
			var jsonResponse ApiResult
			if err := json.Unmarshal(body, &jsonResponse); err == nil {
				return jsonResponse
			} else {
				errStr = fmt.Sprintf("json parse error: %s (%s)", err.Error(), string(body))
			}
		} else {
			errStr = fmt.Sprintf("response read error: %s", err.Error())
		}
	}

	b.error(errStr)

	return ApiResult{Ok: false, Description: errStr}
}

// Webhook request handler
func (b *Bot) handleWebhook(writer http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()

	b.verbose("received webhook request: %+v", req)

	if body, err := ioutil.ReadAll(req.Body); err == nil {
		var webhook Webhook
		if err := json.Unmarshal(body, &webhook); err != nil {
			b.error("error while parsing json (%s)", err.Error())
		} else {
			b.verbose("received webhook body: %s", string(body))

			b.WebhookHandler(webhook, true, nil)
		}
	} else {
		b.error("error while reading webhook request (%s)", err.Error())

		b.WebhookHandler(Webhook{}, false, err)
	}
}

// Start Webhook server and wait forever
//
// (https://core.telegram.org/bots/self-signed)
//
// @param certFilepath [string] certification file's path (.pem)
// @param keyFilepath [string] private key file's path
// @param webhookHandler [func] webhook handler function
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

// (Request methods)
// https://core.telegram.org/bots/api#available-methods

// Get info of this bot
//
// https://core.telegram.org/bots/api#getme
//
func (b *Bot) GetMe() (result ApiResultUser) {
	params := map[string]interface{}{} // no parameters

	var errStr string

	if resp, success := b.sendRequest("getMe", params); success {
		defer resp.Body.Close()

		if body, err := ioutil.ReadAll(resp.Body); err == nil {
			var jsonResponse ApiResultUser
			if err := json.Unmarshal(body, &jsonResponse); err == nil {
				return jsonResponse
			} else {
				errStr = fmt.Sprintf("json parse error: %s (%s)", err.Error(), string(body))
			}
		} else {
			errStr = fmt.Sprintf("response read error: %s", err.Error())
		}
	}

	b.error(errStr)

	return ApiResultUser{Ok: false, Description: errStr}
}

// Send a message
//
// https://core.telegram.org/bots/api#sendmessage
func (b *Bot) SendMessage(chatId interface{}, text *string, parseMode *string, disableWebPagePreview *bool, replyToMessageId *int, replyMarkup *interface{}) (result ApiResultMessage) {
	params := map[string]interface{}{
		"chat_id": chatId,
		"text":    *text,
	}
	if parseMode != nil {
		params["parse_mode"] = *parseMode
	}
	if disableWebPagePreview != nil {
		params["disable_web_preview"] = *disableWebPagePreview
	}
	if replyToMessageId != nil {
		params["reply_to_message_id"] = *replyToMessageId
	}
	if replyMarkup != nil {
		params["reply_markup"] = &replyMarkup
	}

	var errStr string

	if resp, success := b.sendRequest("sendMessage", params); success {
		defer resp.Body.Close()

		if body, err := ioutil.ReadAll(resp.Body); err == nil {
			var jsonResponse ApiResultMessage
			if err := json.Unmarshal(body, &jsonResponse); err == nil {
				return jsonResponse
			} else {
				b.error("json parse error: %s (%s)", err.Error(), string(body))
			}
		} else {
			errStr = fmt.Sprintf("response read error: %s", err.Error())
		}
	} else {
		errStr = fmt.Sprintf("SendMessage failed")
	}

	b.error(errStr)

	return ApiResultMessage{Ok: false, Description: errStr}
}

// Forward a message
//
// https://core.telegram.org/bots/api#forwardmessage
// TODO

// Send photos
//
// https://core.telegram.org/bots/api#sendphoto
// TODO

// Send audio files (will be played with external players)
//
// https://core.telegram.org/bots/api#sendaudio
// TODO

// Send general files
//
// https://core.telegram.org/bots/api#senddocument
// TODO

// Send stickers
//
// https://core.telegram.org/bots/api#sendsticker
// TODO

// Send video files
//
// https://core.telegram.org/bots/api#sendvideo
// TODO

// Send voice files (.ogg format only, will be played with Telegram itself))
//
// https://core.telegram.org/bots/api#sendvoice
// TODO

// Send locations
//
// https://core.telegram.org/bots/api#sendlocation
// TODO

// Send chat action
//
// https://core.telegram.org/bots/api#sendchataction
// TODO

// Get user profile photos
//
// https://core.telegram.org/bots/api#getuserprofilephotos
// TODO

// Get updates
//
// https://core.telegram.org/bots/api#getupdates
// TODO

// Get file info and prepare for download
//
// https://core.telegram.org/bots/api#getfile
// TODO
