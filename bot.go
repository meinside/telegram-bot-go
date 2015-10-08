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
	"sync"
)

const (
	ApiBaseUrl  = "https://api.telegram.org/bot"
	DefaultPort = 443
	WebhookPath = "/telegram/bot/webhook"
)

type Bot struct {
	Token           string
	WebhookProtocol string
	WebhookHost     string
	WebhookPort     int
	WebhookUrl      string
	WebhookHandler  func(success bool, err error, writer http.ResponseWriter, webhook Webhook)
	Verbose         bool
	WaitGroup       sync.WaitGroup
}

type ApiResult struct {
	Ok          bool        `json:"ok"`
	Description string      `json:"description,omitempty"`
	Result      interface{} `json:"result,omitempty"`
}

type Webhook struct {
	UpdateId int         `json:"update_id"`
	Chat     interface{} `json:"chat"`
	Message  Message     `json:"message"`
}

// Get new bot API client
func NewClient(token string) *Bot {
	return &Bot{
		Token:       token,
		WebhookPort: DefaultPort,
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

// Send request to API server and callback its response
func (b *Bot) sendRequest(method string, params map[string]interface{}, callback func(res *http.Response)) {
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
			callback(resp)
		} else {
			b.error("request error: %s", err.Error())
		}
	} else {
		b.error("building request error: %s", err.Error())
	}

	// wait group
	b.WaitGroup.Done()
}

// Set webhook url for receiving incoming updates
//
// https://core.telegram.org/bots/api#setwebhook
//
// @param host [string] webhook server host
// @param port [int] webhook server port (443, 80, 88, or 8443)
// @param certFilepath [string] certification file's path
// @param callback [func] callback function
func (b *Bot) SetWebhookUrl(host string, port int, certFilepath string, callback func(success bool, err error, description *string)) {
	// wait group
	b.WaitGroup.Add(1)

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

	go b.sendRequest("setWebhook", params, func(resp *http.Response) {
		defer resp.Body.Close()

		if body, err := ioutil.ReadAll(resp.Body); err == nil {
			var jsonResponse ApiResult
			if err := json.Unmarshal(body, &jsonResponse); err == nil {
				callback(true, nil, &jsonResponse.Description)
			} else {
				b.error("json parse error: %s (%s)", err.Error(), string(body))

				callback(false, err, nil)
			}
		} else {
			b.error("response read error: %s", err.Error())

			callback(false, err, nil)
		}
	})
}

// Delete webhook url
//
// https://core.telegram.org/bots/api#setwebhook
//
// @param callback [func] callback function
func (b *Bot) DeleteWebhookUrl(callback func(success bool, err error, description *string)) {
	// wait group
	b.WaitGroup.Add(1)

	b.WebhookHost = ""
	b.WebhookUrl = ""

	params := map[string]interface{}{}

	go b.sendRequest("setWebhook", params, func(resp *http.Response) {
		defer resp.Body.Close()
		if body, err := ioutil.ReadAll(resp.Body); err == nil {
			var jsonResponse ApiResult
			if err := json.Unmarshal(body, &jsonResponse); err == nil {
				callback(true, nil, &jsonResponse.Description)
			} else {
				b.error("json parse error: %s (%s)", err.Error(), string(body))

				callback(false, err, nil)
			}
		} else {
			b.error("response read error: %s", err.Error())

			callback(false, err, nil)
		}
	})
}

// Get info of this bot
//
// https://core.telegram.org/bots/api#getme
//
// @param callback [func] callback function
func (b *Bot) GetMe(callback func(success bool, err error, result map[string]interface{})) {
	// wait group
	b.WaitGroup.Add(1)

	params := map[string]interface{}{} // no parameters

	go b.sendRequest("getMe", params, func(resp *http.Response) {
		defer resp.Body.Close()
		if body, err := ioutil.ReadAll(resp.Body); err == nil {
			var jsonResponse ApiResult
			if err := json.Unmarshal(body, &jsonResponse); err == nil {
				if resultMap, ok := jsonResponse.Result.(map[string]interface{}); ok {
					callback(true, nil, resultMap)
				}
			} else {
				b.error("json parse error: %s (%s)", err.Error(), string(body))

				callback(false, err, nil)
			}
		} else {
			b.error("response read error: %s", err.Error())

			callback(false, err, nil)
		}
	})
}

// Wait for this bot until all requests & responses are finished
func (b *Bot) Wait() {
	b.WaitGroup.Wait()
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
			b.WebhookHandler(true, nil, writer, webhook)
		}
	} else {
		b.error("error while reading webhook request (%s)", err.Error())

		b.WebhookHandler(false, err, writer, Webhook{})
	}
}

// Start Webhook server and wait forever
//
// (https://core.telegram.org/bots/self-signed)
//
// @param certFilepath [string] certification file's path (.pem)
// @param keyFilepath [string] private key file's path
// @param webhookHandler [func] webhook handler function
func (b *Bot) StartWebhookServerAndWait(certFilepath string, keyFilepath string, webhookHandler func(success bool, err error, writer http.ResponseWriter, webhook Webhook)) {
	b.verbose("starting webhook server on: %s (port: %d) ...", b.getWebhookPath(), b.WebhookPort)

	// routing
	b.WebhookHandler = webhookHandler
	http.HandleFunc(b.getWebhookPath(), b.handleWebhook)

	// start server
	if err := http.ListenAndServeTLS(fmt.Sprintf(":%d", b.WebhookPort), certFilepath, keyFilepath, nil); err != nil {
		panic(err.Error())
	}
}
