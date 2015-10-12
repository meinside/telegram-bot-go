// https://core.telegram.org/bots/api#available-methods
//

package telegrambot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

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

// Get info of this bot
//
// @return [ApiResultUser]
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
// @param chatId [int,string] chat id
// @param text [*string] message
// @param options [*map[string]interface{}] optional parameters
//        ( = parse_mode, disable_web_page_preview, reply_to_message_id, reply_markup)
//
// @return [ApiResultMessage]
//
// https://core.telegram.org/bots/api#sendmessage
//
func (b *Bot) SendMessage(chatId interface{}, text *string, options *map[string]interface{}) (result ApiResultMessage) {
	// essential params
	params := map[string]interface{}{
		"chat_id": chatId,
		"text":    *text,
	}
	// optional params
	for key, val := range *options {
		if val != nil {
			params[key] = val
		}
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
// @param chatId [int,string] chat id
// @param fromChatId [int,string] original message's chat id
// @param messageId [int] message id
//
// @return [ApiResultMessage]
//
// https://core.telegram.org/bots/api#forwardmessage
//
func (b *Bot) ForwardMessage(chatId interface{}, fromChatId interface{}, messageId int) (result ApiResultMessage) {
	// essential params
	params := map[string]interface{}{
		"chat_id":      chatId,
		"from_chat_id": fromChatId,
		"message_id":   messageId,
	}

	var errStr string

	if resp, success := b.sendRequest("forwardMessage", params); success {
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
		errStr = fmt.Sprintf("ForwardMessage failed")
	}

	b.error(errStr)

	return ApiResultMessage{Ok: false, Description: errStr}
}

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
	case bool:
		if boolValue, ok := param.(bool); ok {
			return strconv.FormatBool(boolValue), ok
		} else {
			b.error("parameter '%+v' could not be cast to bool value", param)
		}
	case ReplyKeyboardMarkup:
		if value, ok := param.(ReplyKeyboardMarkup); ok {
			if json, err := json.Marshal(value); err == nil {
				return string(json), true
			} else {
				b.error(err.Error())
			}
		} else {
			b.error("parameter '%+v' could not be cast to ReplyKeyboardMarkup value", param)
		}
	case ReplyKeyboardHide:
		if value, ok := param.(ReplyKeyboardHide); ok {
			if json, err := json.Marshal(value); err == nil {
				return string(json), true
			} else {
				b.error(err.Error())
			}
		} else {
			b.error("parameter '%+v' could not be cast to ReplyKeyboardHide value", param)
		}
	case ForceReply:
		if value, ok := param.(ForceReply); ok {
			if json, err := json.Marshal(value); err == nil {
				return string(json), true
			} else {
				b.error(err.Error())
			}
		} else {
			b.error("parameter '%+v' could not be cast to ForceReply value", param)
		}
	default:
		b.error("unexpected type: '%+v' (%T)", param, param)
	}

	return "", true
}

// Send request to API server and return the response (synchronous)
//
// @param method [string] HTTP method
// @param params [map[string]interface{}] request parameters (if *os.File is given, it will be closed automatically)
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
				if file, ok := value.(*os.File); ok {
					defer file.Close()

					if part, err := writer.CreateFormFile(key, file.Name()); err == nil {
						if _, err = io.Copy(part, file); err != nil {
							b.error("could now write to multipart: %s", key)
						}
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
