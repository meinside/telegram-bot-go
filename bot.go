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
	"os"
)

const (
	ApiBaseUrl         = "https://api.telegram.org/bot"
	DefaultWebhookPort = 443
	WebhookPath        = "/telegram/bot/webhook"
)

type Bot struct {
	Token           string
	WebhookProtocol string
	WebhookHost     string
	WebhookPort     int
	WebhookUrl      string
	WebhookHandler  func(success bool, err error, webhook Webhook)
	Verbose         bool
}

// Get new bot API client
func NewClient(token string) *Bot {
	return &Bot{
		Token:       token,
		WebhookPort: DefaultWebhookPort,
	}
}

func (b *Bot) getWebhookPath() string {
	return fmt.Sprintf("%s/%x", WebhookPath, md5.Sum([]byte(b.Token)))
}

func (b *Bot) getWebhookUrl() string {
	return fmt.Sprintf("https://%s:%d%s", b.WebhookHost, b.WebhookPort, b.getWebhookPath())
}

// Print formatted log message (only when Bot.Verbose == true)
func (b *Bot) verbose(str string, args ...interface{}) {
	if b.Verbose {
		fmt.Printf("> %s\n", fmt.Sprintf(str, args...))
	}
}

// Print formatted error message
func (b *Bot) error(str string, args ...interface{}) {
	fmt.Printf("* %s\n", fmt.Sprintf(str, args...))
}

// Send request to API server and return the response (synchronous)
//
// @param method [string] HTTP method
// @param params [map[string]interface{}] request parameters
func (b *Bot) sendRequest(method string, params map[string]interface{}) (success bool, resp *http.Response) {
	client := &http.Client{}
	apiUrl := fmt.Sprintf("%s%s/%s", ApiBaseUrl, b.Token, method)

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	fileIncluded := false

	for key, value := range params {
		switch value.(type) {
		case string:
			if strValue, ok := value.(string); ok {
				writer.WriteField(key, strValue)
			} else {
				b.error("parameter '%s' (%v) could not be cast to string value", key, value)
			}
		case *os.File:
			if fileValue, ok := value.(*os.File); ok {
				if part, err := writer.CreateFormFile(key, fileValue.Name()); err == nil {
					if _, err = io.Copy(part, fileValue); err != nil {
						b.error("could now write to multipart: %s", key)
					} else {
						fileIncluded = true
					}

					defer fileValue.Close()
				} else {
					b.error("could not create form file for parameter '%s' (%v)", key, value)
				}
			} else {
				b.error("parameter '%s' (%v) could not be cast to file", key, value)
			}
		default:
			b.error("unexpected type: %+v (%T)", value, value)
		}
	}

	if err := writer.Close(); err != nil {
		b.error("error while closing writer (%s)", err.Error())
	}

	if req, err := http.NewRequest("POST", apiUrl, body); err == nil {
		if fileIncluded {
			req.Header.Add("Content-Type", writer.FormDataContentType())
		}

		if resp, err := client.Do(req); err == nil {
			return true, resp
		} else {
			b.error("request error: %s", err.Error())
		}
	} else {
		b.error("building request error: %s", err.Error())
	}

	return false, nil
}

// Set webhook url for receiving incoming updates
//
// https://core.telegram.org/bots/api#setwebhook
//
// @param host [string] webhook server host
// @param port [int] webhook server port (443, 80, 88, or 8443)
// @param certFilepath [string] certification file's path
func (b *Bot) SetWebhookUrl(host string, port int, certFilepath string) (success bool, description *string) {
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

	if success, resp := b.sendRequest("setWebhook", params); success {
		defer resp.Body.Close()

		if body, err := ioutil.ReadAll(resp.Body); err == nil {
			var jsonResponse ApiResult
			if err := json.Unmarshal(body, &jsonResponse); err == nil {
				return true, &jsonResponse.Description
			} else {
				b.error("json parse error: %s (%s)", err.Error(), string(body))
			}
		} else {
			b.error("response read error: %s", err.Error())
		}
	}

	return false, nil
}

// Delete webhook url
//
// https://core.telegram.org/bots/api#setwebhook
//
func (b *Bot) DeleteWebhookUrl() (success bool, description *string) {
	b.WebhookHost = ""
	b.WebhookUrl = ""

	params := map[string]interface{}{
		"url": "",
	}

	if success, resp := b.sendRequest("setWebhook", params); success {
		defer resp.Body.Close()

		if body, err := ioutil.ReadAll(resp.Body); err == nil {
			var jsonResponse ApiResult
			if err := json.Unmarshal(body, &jsonResponse); err == nil {
				return true, &jsonResponse.Description
			} else {
				b.error("json parse error: %s (%s)", err.Error(), string(body))
			}
		} else {
			b.error("response read error: %s", err.Error())
		}
	}

	return false, nil
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
			b.verbose("received webhook body = %s", string(body))

			b.WebhookHandler(true, nil, webhook)
		}
	} else {
		b.error("error while reading webhook request (%s)", err.Error())

		b.WebhookHandler(false, err, Webhook{})
	}
}

// Start Webhook server and wait forever
//
// (https://core.telegram.org/bots/self-signed)
//
// @param certFilepath [string] certification file's path (.pem)
// @param keyFilepath [string] private key file's path
// @param webhookHandler [func] webhook handler function
func (b *Bot) StartWebhookServerAndWait(certFilepath string, keyFilepath string, webhookHandler func(success bool, err error, webhook Webhook)) {
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
func (b *Bot) GetMe() (success bool, result map[string]interface{}) {
	params := map[string]interface{}{} // no parameters

	if success, resp := b.sendRequest("getMe", params); success {
		defer resp.Body.Close()

		if body, err := ioutil.ReadAll(resp.Body); err == nil {
			var jsonResponse ApiResult
			if err := json.Unmarshal(body, &jsonResponse); err == nil {
				if resultMap, ok := jsonResponse.Result.(map[string]interface{}); ok {
					return true, resultMap
				}
			} else {
				b.error("json parse error: %s (%s)", err.Error(), string(body))
			}
		} else {
			b.error("response read error: %s", err.Error())
		}
	}

	return false, nil
}

// Send a message
//
// https://core.telegram.org/bots/api#sendmessage
// TODO

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
