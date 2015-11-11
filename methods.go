// https://core.telegram.org/bots/api#available-methods

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

// Set webhook url and certificate for receiving incoming updates.
//
// https://core.telegram.org/bots/api#setwebhook
//
// port should be one of: 443, 80, 88, or 8443.
func (b *Bot) SetWebhook(host string, port int, certFilepath string) (result ApiResult) {
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
	} else {
		errStr = fmt.Sprintf("SetWebhook failed")
	}

	b.error(errStr)

	return ApiResult{Ok: false, Description: &errStr}
}

// Delete webhook.
//
// https://core.telegram.org/bots/api#setwebhook
//
// (Function GetUpdates will not work if webhook is set, so in that case you'll need to delete it)
func (b *Bot) DeleteWebhook() (result ApiResult) {
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
	} else {
		errStr = fmt.Sprintf("DeleteWebhook failed")
	}

	b.error(errStr)

	return ApiResult{Ok: false, Description: &errStr}
}

// Get info of this bot.
//
// https://core.telegram.org/bots/api#getme
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
	} else {
		errStr = fmt.Sprintf("GetMe failed")
	}

	b.error(errStr)

	return ApiResultUser{Ok: false, Description: &errStr}
}

// Send a message.
//
// https://core.telegram.org/bots/api#sendmessage
//
// chatId can be Message.Chat.Id or target channel(eg. @channelusername).
//
// options include parse_mode, disable_web_page_preview, reply_to_message_id, and reply_markup.
func (b *Bot) SendMessage(chatId interface{}, text *string, options map[string]interface{}) (result ApiResultMessage) {
	// essential params
	params := map[string]interface{}{
		"chat_id": chatId,
		"text":    *text,
	}
	// optional params
	for key, val := range options {
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
				errStr = fmt.Sprintf("json parse error: %s (%s)", err.Error(), string(body))
			}
		} else {
			errStr = fmt.Sprintf("response read error: %s", err.Error())
		}
	} else {
		errStr = fmt.Sprintf("SendMessage failed")
	}

	b.error(errStr)

	return ApiResultMessage{Ok: false, Description: &errStr}
}

// Forward a message.
//
// https://core.telegram.org/bots/api#forwardmessage
//
// chatId and fromChatId can be Message.Chat.Id or target channel(eg. @channelusername).
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
				errStr = fmt.Sprintf("json parse error: %s (%s)", err.Error(), string(body))
			}
		} else {
			errStr = fmt.Sprintf("response read error: %s", err.Error())
		}
	} else {
		errStr = fmt.Sprintf("ForwardMessage failed")
	}

	b.error(errStr)

	return ApiResultMessage{Ok: false, Description: &errStr}
}

// Send photos.
//
// https://core.telegram.org/bots/api#sendphoto
//
// chatId can be Message.Chat.Id or target channel(eg. @channelusername).
//
// options include caption, reply_to_message_id, and reply_markup.
func (b *Bot) SendPhoto(chatId interface{}, photoFilepath *string, options map[string]interface{}) (result ApiResultMessage) {
	var errStr string

	if file, err := os.Open(*photoFilepath); err == nil {
		// essential params
		params := map[string]interface{}{
			"chat_id": chatId,
			"photo":   file,
		}
		// optional params
		for key, val := range options {
			if val != nil {
				params[key] = val
			}
		}

		if resp, success := b.sendRequest("sendPhoto", params); success {
			defer resp.Body.Close()

			if body, err := ioutil.ReadAll(resp.Body); err == nil {
				var jsonResponse ApiResultMessage
				if err := json.Unmarshal(body, &jsonResponse); err == nil {
					return jsonResponse
				} else {
					errStr = fmt.Sprintf("json parse error: %s (%s)", err.Error(), string(body))
				}
			} else {
				errStr = fmt.Sprintf("response read error: %s", err.Error())
			}
		} else {
			errStr = fmt.Sprintf("SendPhoto failed")
		}
	} else {
		errStr = err.Error()
	}

	b.error(errStr)

	return ApiResultMessage{Ok: false, Description: &errStr}
}

// Send audio files. (.mp3 format only, will be played with external players)
//
// https://core.telegram.org/bots/api#sendaudio
//
// chatId can be Message.Chat.Id or target channel(eg. @channelusername).
//
// options include duration, performer, title, reply_to_message_id, and reply_markup.
func (b *Bot) SendAudio(chatId interface{}, audioFilepath *string, options map[string]interface{}) (result ApiResultMessage) {
	var errStr string

	if file, err := os.Open(*audioFilepath); err == nil {
		// essential params
		params := map[string]interface{}{
			"chat_id": chatId,
			"audio":   file,
		}
		// optional params
		for key, val := range options {
			if val != nil {
				params[key] = val
			}
		}

		if resp, success := b.sendRequest("sendAudio", params); success {
			defer resp.Body.Close()

			if body, err := ioutil.ReadAll(resp.Body); err == nil {
				var jsonResponse ApiResultMessage
				if err := json.Unmarshal(body, &jsonResponse); err == nil {
					return jsonResponse
				} else {
					errStr = fmt.Sprintf("json parse error: %s (%s)", err.Error(), string(body))
				}
			} else {
				errStr = fmt.Sprintf("response read error: %s", err.Error())
			}
		} else {
			errStr = fmt.Sprintf("SendAudio failed")
		}
	} else {
		errStr = err.Error()
	}

	b.error(errStr)

	return ApiResultMessage{Ok: false, Description: &errStr}
}

// Send general files.
//
// https://core.telegram.org/bots/api#senddocument
//
// chatId can be Message.Chat.Id or target channel(eg. @channelusername).
//
// options include reply_to_message_id, and reply_markup.
func (b *Bot) SendDocument(chatId interface{}, documentFilepath *string, options map[string]interface{}) (result ApiResultMessage) {
	var errStr string

	if file, err := os.Open(*documentFilepath); err == nil {
		// essential params
		params := map[string]interface{}{
			"chat_id":  chatId,
			"document": file,
		}
		// optional params
		for key, val := range options {
			if val != nil {
				params[key] = val
			}
		}

		if resp, success := b.sendRequest("sendDocument", params); success {
			defer resp.Body.Close()

			if body, err := ioutil.ReadAll(resp.Body); err == nil {
				var jsonResponse ApiResultMessage
				if err := json.Unmarshal(body, &jsonResponse); err == nil {
					return jsonResponse
				} else {
					errStr = fmt.Sprintf("json parse error: %s (%s)", err.Error(), string(body))
				}
			} else {
				errStr = fmt.Sprintf("response read error: %s", err.Error())
			}
		} else {
			errStr = fmt.Sprintf("SendDocument failed")
		}
	} else {
		errStr = err.Error()
	}

	b.error(errStr)

	return ApiResultMessage{Ok: false, Description: &errStr}
}

// Send stickers.
//
// https://core.telegram.org/bots/api#sendsticker
//
// chatId can be Message.Chat.Id or target channel(eg. @channelusername).
//
// options include reply_to_message_id, and reply_markup.
func (b *Bot) SendSticker(chatId interface{}, stickerFilepath string, options *map[string]interface{}) (result ApiResultMessage) {
	var errStr string

	if file, err := os.Open(stickerFilepath); err == nil {
		// essential params
		params := map[string]interface{}{
			"chat_id": chatId,
			"sticker": file,
		}
		// optional params
		for key, val := range *options {
			if val != nil {
				params[key] = val
			}
		}

		if resp, success := b.sendRequest("sendSticker", params); success {
			defer resp.Body.Close()

			if body, err := ioutil.ReadAll(resp.Body); err == nil {
				var jsonResponse ApiResultMessage
				if err := json.Unmarshal(body, &jsonResponse); err == nil {
					return jsonResponse
				} else {
					errStr = fmt.Sprintf("json parse error: %s (%s)", err.Error(), string(body))
				}
			} else {
				errStr = fmt.Sprintf("response read error: %s", err.Error())
			}
		} else {
			errStr = fmt.Sprintf("SendSticker failed")
		}
	} else {
		errStr = err.Error()
	}

	b.error(errStr)

	return ApiResultMessage{Ok: false, Description: &errStr}
}

// Send video files.
//
// https://core.telegram.org/bots/api#sendvideo
//
// chatId can be Message.Chat.Id or target channel(eg. @channelusername).
//
// options include duration, caption, reply_to_message_id, and reply_markup.
func (b *Bot) SendVideo(chatId interface{}, videoFilepath *string, options map[string]interface{}) (result ApiResultMessage) {
	var errStr string

	if file, err := os.Open(*videoFilepath); err == nil {
		// essential params
		params := map[string]interface{}{
			"chat_id": chatId,
			"video":   file,
		}
		// optional params
		for key, val := range options {
			if val != nil {
				params[key] = val
			}
		}

		if resp, success := b.sendRequest("sendVideo", params); success {
			defer resp.Body.Close()

			if body, err := ioutil.ReadAll(resp.Body); err == nil {
				var jsonResponse ApiResultMessage
				if err := json.Unmarshal(body, &jsonResponse); err == nil {
					return jsonResponse
				} else {
					errStr = fmt.Sprintf("json parse error: %s (%s)", err.Error(), string(body))
				}
			} else {
				errStr = fmt.Sprintf("response read error: %s", err.Error())
			}
		} else {
			errStr = fmt.Sprintf("SendVideo failed")
		}
	} else {
		errStr = err.Error()
	}

	b.error(errStr)

	return ApiResultMessage{Ok: false, Description: &errStr}
}

// Send voice files. (.ogg format only, will be played with Telegram itself))
//
// https://core.telegram.org/bots/api#sendvoice
//
// chatId can be Message.Chat.Id or target channel(eg. @channelusername).
//
// options include duration, reply_to_message_id, and reply_markup.
func (b *Bot) SendVoice(chatId interface{}, voiceFilepath string, options *map[string]interface{}) (result ApiResultMessage) {
	var errStr string

	if file, err := os.Open(voiceFilepath); err == nil {
		// essential params
		params := map[string]interface{}{
			"chat_id": chatId,
			"voice":   file,
		}
		// optional params
		for key, val := range *options {
			if val != nil {
				params[key] = val
			}
		}

		if resp, success := b.sendRequest("sendVoice", params); success {
			defer resp.Body.Close()

			if body, err := ioutil.ReadAll(resp.Body); err == nil {
				var jsonResponse ApiResultMessage
				if err := json.Unmarshal(body, &jsonResponse); err == nil {
					return jsonResponse
				} else {
					errStr = fmt.Sprintf("json parse error: %s (%s)", err.Error(), string(body))
				}
			} else {
				errStr = fmt.Sprintf("response read error: %s", err.Error())
			}
		} else {
			errStr = fmt.Sprintf("SendVoice failed")
		}
	} else {
		errStr = err.Error()
	}

	b.error(errStr)

	return ApiResultMessage{Ok: false, Description: &errStr}
}

// Send locations.
//
// https://core.telegram.org/bots/api#sendlocation
//
// chatId can be Message.Chat.Id or target channel(eg. @channelusername).
//
// options include reply_to_message_id, and reply_markup.
func (b *Bot) SendLocation(chatId interface{}, latitude float32, longitude float32, options *map[string]interface{}) (result ApiResultMessage) {
	var errStr string

	// essential params
	params := map[string]interface{}{
		"chat_id":   chatId,
		"latitude":  latitude,
		"longitude": longitude,
	}
	// optional params
	for key, val := range *options {
		if val != nil {
			params[key] = val
		}
	}

	if resp, success := b.sendRequest("sendLocation", params); success {
		defer resp.Body.Close()

		if body, err := ioutil.ReadAll(resp.Body); err == nil {
			var jsonResponse ApiResultMessage
			if err := json.Unmarshal(body, &jsonResponse); err == nil {
				return jsonResponse
			} else {
				errStr = fmt.Sprintf("json parse error: %s (%s)", err.Error(), string(body))
			}
		} else {
			errStr = fmt.Sprintf("response read error: %s", err.Error())
		}
	} else {
		errStr = fmt.Sprintf("SendLocation failed")
	}

	b.error(errStr)

	return ApiResultMessage{Ok: false, Description: &errStr}
}

// Send chat action.
//
// https://core.telegram.org/bots/api#sendchataction
//
// chatId can be Message.Chat.Id or target channel(eg. @channelusername).
//
// action can be one of: typing, upload_photo, record_video, upload_video, record_audio, upload_audio, upload_document, or find_location
func (b *Bot) SendChatAction(chatId interface{}, action string) (result ApiResult) {
	var errStr string

	// essential params
	params := map[string]interface{}{
		"chat_id": chatId,
		"action":  action,
	}

	if resp, success := b.sendRequest("sendChatAction", params); success {
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
	} else {
		errStr = fmt.Sprintf("SendChatAction failed")
	}

	b.error(errStr)

	return ApiResult{Ok: false, Result: false, Description: &errStr}
}

// Get user profile photos.
//
// https://core.telegram.org/bots/api#getuserprofilephotos
//
// options include offset and limit.
func (b *Bot) GetUserProfilePhotos(userId int, options *map[string]interface{}) (result ApiResultUserProfilePhotos) {
	var errStr string

	// essential params
	params := map[string]interface{}{
		"user_id": userId,
	}
	// optional params
	for key, val := range *options {
		if val != nil {
			params[key] = val
		}
	}

	if resp, success := b.sendRequest("getUserProfilePhotos", params); success {
		defer resp.Body.Close()

		if body, err := ioutil.ReadAll(resp.Body); err == nil {
			var jsonResponse ApiResultUserProfilePhotos
			if err := json.Unmarshal(body, &jsonResponse); err == nil {
				return jsonResponse
			} else {
				errStr = fmt.Sprintf("json parse error: %s (%s)", err.Error(), string(body))
			}
		} else {
			errStr = fmt.Sprintf("response read error: %s", err.Error())
		}
	} else {
		errStr = fmt.Sprintf("GetUserProfilePhotos failed")
	}

	b.error(errStr)

	return ApiResultUserProfilePhotos{Ok: false, Description: &errStr}
}

// Get updates.
//
// https://core.telegram.org/bots/api#getupdates
//
// options include offset, limit, and timeout.
func (b *Bot) GetUpdates(options *map[string]interface{}) (result ApiResultUpdates) {
	var errStr string

	// optional params
	params := map[string]interface{}{}
	for key, val := range *options {
		if val != nil {
			params[key] = val
		}
	}

	if resp, success := b.sendRequest("getUpdates", params); success {
		defer resp.Body.Close()

		if body, err := ioutil.ReadAll(resp.Body); err == nil {
			var jsonResponse ApiResultUpdates
			if err := json.Unmarshal(body, &jsonResponse); err == nil {
				return jsonResponse
			} else {
				errStr = fmt.Sprintf("json parse error: %s (%s)", err.Error(), string(body))
			}
		} else {
			errStr = fmt.Sprintf("response read error: %s", err.Error())
		}
	} else {
		errStr = fmt.Sprintf("GetUpdates failed")
	}

	b.error(errStr)

	return ApiResultUpdates{Ok: false, Description: &errStr}
}

// Get file info and prepare for download.
//
// https://core.telegram.org/bots/api#getfile
func (b *Bot) GetFile(fileId *string) (result ApiResultFile) {
	var errStr string

	// essential params
	params := map[string]interface{}{
		"file_id": *fileId,
	}

	if resp, success := b.sendRequest("getFile", params); success {
		defer resp.Body.Close()

		if body, err := ioutil.ReadAll(resp.Body); err == nil {
			var jsonResponse ApiResultFile
			if err := json.Unmarshal(body, &jsonResponse); err == nil {
				return jsonResponse
			} else {
				errStr = fmt.Sprintf("json parse error: %s (%s)", err.Error(), string(body))
			}
		} else {
			errStr = fmt.Sprintf("response read error: %s", err.Error())
		}
	} else {
		errStr = fmt.Sprintf("GetFile failed")
	}

	b.error(errStr)

	return ApiResultFile{Ok: false, Description: &errStr}
}

// Get download link from given File.
func (b *Bot) GetFileUrl(file File) string {
	return fmt.Sprintf("%s%s/%s", FileBaseUrl, b.Token, *file.FilePath)
}

// Check if given http params contain file or not.
func checkIfFileParamExists(params map[string]interface{}) bool {
	for _, value := range params {
		switch value.(type) {
		case *os.File:
			return true
		}
	}

	return false
}

// Convert given interface to string. (for HTTP params)
func (b *Bot) paramToString(param interface{}) (result string, success bool) {
	switch param.(type) {
	case int:
		if intValue, ok := param.(int); ok {
			return strconv.Itoa(intValue), ok
		} else {
			b.error("parameter '%+v' could not be cast to int value", param)
		}
	case float32:
		if floatValue, ok := param.(float32); ok {
			return fmt.Sprintf("%.8f", floatValue), ok
		} else {
			b.error("parameter '%+v' could not be cast to float32 value", param)
		}
	case bool:
		if boolValue, ok := param.(bool); ok {
			return strconv.FormatBool(boolValue), ok
		} else {
			b.error("parameter '%+v' could not be cast to bool value", param)
		}
	case string:
		if strValue, ok := param.(string); ok {
			return strValue, ok
		} else {
			b.error("parameter '%+v' could not be cast to string value", param)
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

// Send request to API server and return the response.
// (synchronous)
//
// If *os.File is included in the params, it will be closed automatically.
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

// Handle Webhook request.
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
