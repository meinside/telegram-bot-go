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
	"strings"
)

// Get updates.
//
// options include: offset, limit, and timeout.
//
// https://core.telegram.org/bots/api#getupdates
func (b *Bot) GetUpdates(options map[string]interface{}) (result ApiResponseUpdates) {
	// optional params
	params := map[string]interface{}{}
	for key, val := range options {
		if val != nil {
			params[key] = val
		}
	}

	return b.requestResponseUpdates("getUpdates", params)
}

// Set webhook url and certificate for receiving incoming updates.
// port should be one of: 443, 80, 88, or 8443.
//
// https://core.telegram.org/bots/api#setwebhook
func (b *Bot) SetWebhook(host string, port int, certFilepath string) (result ApiResponse) {
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

	return b.requestResponse("setWebhook", params)
}

// Delete webhook.
// (Function GetUpdates will not work if webhook is set, so in that case you'll need to delete it)
//
// https://core.telegram.org/bots/api#setwebhook
func (b *Bot) DeleteWebhook() (result ApiResponse) {
	b.webhookHost = ""
	b.webhookPort = 0
	b.webhookUrl = ""

	params := map[string]interface{}{
		"url": "",
	}

	b.verbose("deleting webhook url")

	return b.requestResponse("setWebhook", params)
}

// Get info of this bot.
//
// https://core.telegram.org/bots/api#getme
func (b *Bot) GetMe() (result ApiResponseUser) {
	return b.requestResponseUser("getMe", map[string]interface{}{}) // no params
}

// Send a message.
//
// chatId can be Message.Chat.Id or target channel(eg. @channelusername).
//
// options include: parse_mode, disable_web_page_preview, disable_notification, reply_to_message_id, and reply_markup.
//
// https://core.telegram.org/bots/api#sendmessage
func (b *Bot) SendMessage(chatId interface{}, text *string, options map[string]interface{}) (result ApiResponseMessage) {
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

	return b.requestResponseMessage("sendMessage", params)
}

// Forward a message.
//
// chatId and fromChatId can be Message.Chat.Id or target channel(eg. @channelusername).
//
// options include: disable_notification
//
// https://core.telegram.org/bots/api#forwardmessage
func (b *Bot) ForwardMessage(chatId interface{}, fromChatId interface{}, messageId int) (result ApiResponseMessage) {
	// essential params
	params := map[string]interface{}{
		"chat_id":      chatId,
		"from_chat_id": fromChatId,
		"message_id":   messageId,
	}

	return b.requestResponseMessage("forwardMessage", params)
}

// Send photos.
//
// chatId can be Message.Chat.Id or target channel(eg. @channelusername).
//
// photoFilepath can be local filepath or remote http url.
//
// options include: caption, disable_notification, reply_to_message_id, and reply_markup.
//
// https://core.telegram.org/bots/api#sendphoto
func (b *Bot) SendPhoto(chatId interface{}, photoFilepath *string, options map[string]interface{}) (result ApiResponseMessage) {
	return b.sendFile(chatId, "sendPhoto", "photo", photoFilepath, options)
}

// Send audio files. (.mp3 format only, will be played with external players)
//
// chatId can be Message.Chat.Id or target channel(eg. @channelusername).
//
// audioFilepath can be local filepath or remote http url.
//
// options include: caption, duration, performer, title, disable_notification, reply_to_message_id, and reply_markup.
//
// https://core.telegram.org/bots/api#sendaudio
func (b *Bot) SendAudio(chatId interface{}, audioFilepath *string, options map[string]interface{}) (result ApiResponseMessage) {
	return b.sendFile(chatId, "sendAudio", "audio", audioFilepath, options)
}

// Send general files.
//
// chatId can be Message.Chat.Id or target channel(eg. @channelusername).
//
// documentFilepath can be local filepath or remote http url.
//
// options include: disable_notification, reply_to_message_id, and reply_markup.
//
// https://core.telegram.org/bots/api#senddocument
func (b *Bot) SendDocument(chatId interface{}, documentFilepath *string, options map[string]interface{}) (result ApiResponseMessage) {
	return b.sendFile(chatId, "sendDocument", "document", documentFilepath, options)
}

// Send stickers.
//
// chatId can be Message.Chat.Id or target channel(eg. @channelusername).
//
// stickerFilepath can be local filepath or remote http url.
//
// options include: disable_notification, reply_to_message_id, and reply_markup.
//
// https://core.telegram.org/bots/api#sendsticker
func (b *Bot) SendSticker(chatId interface{}, stickerFilepath *string, options map[string]interface{}) (result ApiResponseMessage) {
	return b.sendFile(chatId, "sendSticker", "sticker", stickerFilepath, options)
}

// Send video files.
//
// chatId can be Message.Chat.Id or target channel(eg. @channelusername).
//
// videoFilepath can be local filepath or remote http url.
//
// options include: duration, caption, disable_notification, reply_to_message_id, and reply_markup.
//
// https://core.telegram.org/bots/api#sendvideo
func (b *Bot) SendVideo(chatId interface{}, videoFilepath *string, options map[string]interface{}) (result ApiResponseMessage) {
	return b.sendFile(chatId, "sendVideo", "video", videoFilepath, options)
}

// Send voice files. (.ogg format only, will be played with Telegram itself))
//
// chatId can be Message.Chat.Id or target channel(eg. @channelusername).
//
// voiceFilepath can be local filepath or remote http url.
//
// options include: caption, duration, disable_notification, reply_to_message_id, and reply_markup.
//
// https://core.telegram.org/bots/api#sendvoice
func (b *Bot) SendVoice(chatId interface{}, voiceFilepath *string, options map[string]interface{}) (result ApiResponseMessage) {
	return b.sendFile(chatId, "sendVoice", "voice", voiceFilepath, options)
}

// Send locations.
//
// chatId can be Message.Chat.Id or target channel(eg. @channelusername).
//
// options include: display_notification, reply_to_message_id, and reply_markup.
//
// https://core.telegram.org/bots/api#sendlocation
func (b *Bot) SendLocation(chatId interface{}, latitude, longitude float32, options map[string]interface{}) (result ApiResponseMessage) {
	// essential params
	params := map[string]interface{}{
		"chat_id":   chatId,
		"latitude":  latitude,
		"longitude": longitude,
	}
	// optional params
	for key, val := range options {
		if val != nil {
			params[key] = val
		}
	}

	return b.requestResponseMessage("sendLocation", params)
}

// Send venues.
//
// chatId can be Message.Chat.Id or target channel(eg. @channelusername).
//
// options include: foursquare_id, disable_notification, reply_to_message_id, and reply_markup.
//
// https://core.telegram.org/bots/api#sendvenue
func (b *Bot) SendVenue(chatId interface{}, latitude, longitude float32, title, address *string, options map[string]interface{}) (result ApiResponseMessage) {
	// essential params
	params := map[string]interface{}{
		"chat_id":   chatId,
		"latitude":  latitude,
		"longitude": longitude,
		"title":     *title,
		"address":   *address,
	}
	// optional params
	for key, val := range options {
		if val != nil {
			params[key] = val
		}
	}

	return b.requestResponseMessage("sendVenue", params)
}

// Send contacts.
//
// chatId can be Message.Chat.Id or target channel(eg. @channelusername).
//
// options include: last_name, disable_notification, reply_to_message_id, and reply_markup.
//
// https://core.telegram.org/bots/api#sendcontact
func (b *Bot) SendContact(chatId interface{}, phoneNumber, firstName *string, options map[string]interface{}) (result ApiResponseMessage) {
	// essential params
	params := map[string]interface{}{
		"chat_id":      chatId,
		"phone_number": phoneNumber,
		"first_name":   firstName,
	}
	// optional params
	for key, val := range options {
		if val != nil {
			params[key] = val
		}
	}

	return b.requestResponseMessage("sendContact", params)
}

// Send chat action.
//
// chatId can be Message.Chat.Id or target channel(eg. @channelusername).
//
// https://core.telegram.org/bots/api#sendchataction
func (b *Bot) SendChatAction(chatId interface{}, action ChatAction) (result ApiResponse) {
	// essential params
	params := map[string]interface{}{
		"chat_id": chatId,
		"action":  action,
	}

	return b.requestResponse("sendChatAction", params)
}

// Get user profile photos.
//
// options include: offset and limit.
//
// https://core.telegram.org/bots/api#getuserprofilephotos
func (b *Bot) GetUserProfilePhotos(userId int, options map[string]interface{}) (result ApiResponseUserProfilePhotos) {
	// essential params
	params := map[string]interface{}{
		"user_id": userId,
	}
	// optional params
	for key, val := range options {
		if val != nil {
			params[key] = val
		}
	}

	return b.requestResponseUserProfilePhotos("getUserProfilePhotos", params)
}

// Get file info and prepare for download.
//
// https://core.telegram.org/bots/api#getfile
func (b *Bot) GetFile(fileId *string) (result ApiResponseFile) {
	// essential params
	params := map[string]interface{}{
		"file_id": *fileId,
	}

	return b.requestResponseFile("getFile", params)
}

// Get download link from given File.
func (b *Bot) GetFileUrl(file File) string {
	return fmt.Sprintf("%s%s/%s", FileBaseUrl, b.token, *file.FilePath)
}

// Kick chat member
//
// https://core.telegram.org/bots/api#kickchatmember
func (b *Bot) KickChatMember(chatId interface{}, userId int) (result ApiResponse) {
	// essential params
	params := map[string]interface{}{
		"chat_id": chatId,
		"user_id": userId,
	}

	return b.requestResponse("kickChatMember", params)
}

// Leave chat
//
// https://core.telegram.org/bots/api#leavechat
func (b *Bot) LeaveChat(chatId interface{}) (result ApiResponse) {
	// essential params
	params := map[string]interface{}{
		"chat_id": chatId,
	}

	return b.requestResponse("leaveChat", params)
}

// Unban chat member
//
// https://core.telegram.org/bots/api#unbanchatmember
func (b *Bot) UnbanChatMember(chatId interface{}, userId int) (result ApiResponse) {
	// essential params
	params := map[string]interface{}{
		"chat_id": chatId,
		"user_id": userId,
	}

	return b.requestResponse("unbanChatMember", params)
}

// Get chat
//
// https://core.telegram.org/bots/api#getchat
func (b *Bot) GetChat(chatId interface{}) (result ApiResponseChat) {
	// essential params
	params := map[string]interface{}{
		"chat_id": chatId,
	}

	return b.requestResponseChat("getChat", params)
}

// Get chat administrators
//
// https://core.telegram.org/bots/api#getchatadministrators
func (b *Bot) GetChatAdministrators(chatId interface{}) (result ApiResponseChatAdministrators) {
	// essential params
	params := map[string]interface{}{
		"chat_id": chatId,
	}

	return b.requestResponseChatAdministrators("getChatAdministrators", params)
}

// Get chat members count
//
// https://core.telegram.org/bots/api#getchatmemberscount
func (b *Bot) GetChatMembersCount(chatId interface{}) (result ApiResponseInt) {
	// essential params
	params := map[string]interface{}{
		"chat_id": chatId,
	}

	return b.requestResponseInt("getChatMembersCount", params)
}

// Get chat member
//
// https://core.telegram.org/bots/api#getchatmember
func (b *Bot) GetChatMember(chatId interface{}, userId int) (result ApiResponseChatMember) {
	// essential params
	params := map[string]interface{}{
		"chat_id": chatId,
		"user_id": userId,
	}

	return b.requestResponseChatMember("getChatMember", params)
}

// Answer callback query
//
// options include: text, show_alert, and url
//
// https://core.telegram.org/bots/api#answercallbackquery
func (b *Bot) AnswerCallbackQuery(callbackQueryId *string, options map[string]interface{}) (result ApiResponse) {
	// essential params
	params := map[string]interface{}{
		"callback_query_id": *callbackQueryId,
	}
	// optional params
	for key, val := range options {
		if val != nil {
			params[key] = val
		}
	}

	return b.requestResponse("answerCallbackQuery", params)
}

// Updating messages
//
// https://core.telegram.org/bots/api#updating-messages

// Edit text of message
//
// required options: chat_id + message_id (when inline_message_id is not given)
//                or inline_message_id (when chat_id & message_id is not given)
//
// other options: parse_mode, disable_web_page_preview, and reply_markup
//
// https://core.telegram.org/bots/api#editmessagetext
func (b *Bot) EditMessageText(text *string, options map[string]interface{}) (result ApiResponseMessage) {
	// essential params
	params := map[string]interface{}{
		"text": *text,
	}
	// optional params
	for key, val := range options {
		if val != nil {
			params[key] = val
		}
	}

	return b.requestResponseMessage("editMessageText", params)
}

// Edit caption of message
//
// required options: chat_id + message_id (when inline_message_id is not given)
//                or inline_message_id (when chat_id & message_id is not given)
//
// other options: reply_markup
//
// https://core.telegram.org/bots/api#editmessagecaption
func (b *Bot) EditMessageCaption(caption *string, options map[string]interface{}) (result ApiResponseMessage) {
	// essential params
	params := map[string]interface{}{
		"caption": *caption,
	}
	// optional params
	for key, val := range options {
		if val != nil {
			params[key] = val
		}
	}

	return b.requestResponseMessage("editMessageCaption", params)
}

// Edit reply markup of message
//
// required options: chat_id + message_id (when inline_message_id is not given)
//                or inline_message_id (when chat_id & message_id is not given)
//
// other options: reply_markup
//
// https://core.telegram.org/bots/api#editmessagereplymarkup
func (b *Bot) EditMessageReplyMarkup(options map[string]interface{}) (result ApiResponseMessage) {
	return b.requestResponseMessage("editMessageReplyMarkup", options)
}

// Send answers to an inline query.
//
// results = array of InlineQueryResultArticle, InlineQueryResultPhoto, InlineQueryResultGif, InlineQueryResultMpeg4Gif, or InlineQueryResultVideo.
//
// options include: cache_time, is_personal, next_offset, switch_pm_text, and switch_pm_parameter.
//
// https://core.telegram.org/bots/api#answerinlinequery
func (b *Bot) AnswerInlineQuery(inlineQueryId string, results []interface{}, options map[string]interface{}) (result ApiResponse) {
	// essential params
	params := map[string]interface{}{
		"inline_query_id": inlineQueryId,
		"results":         results,
	}
	// optional params
	for key, val := range options {
		if val != nil {
			params[key] = val
		}
	}

	return b.requestResponse("answerInlineQuery", params)
}

// Send a game.
//
// options include: game_short_name, disable_notification, reply_to_message_id, and reply_markup.
//
// https://core.telegram.org/bots/api#sendgame
func (b *Bot) SendGame(chatId interface{}, gameShortName string, options map[string]interface{}) (result ApiResponseMessage) {
	// essential params
	params := map[string]interface{}{
		"chat_id":         chatId,
		"game_short_name": gameShortName,
	}
	// optional params
	for key, val := range options {
		if val != nil {
			params[key] = val
		}
	}

	return b.requestResponseMessage("sendGame", params)
}

// Set score of a game.
//
// required options: chat_id + message_id (when inline_message_id is not given)
//                or inline_message_id (when chat_id & message_id is not given)
//
// other options: edit_message
//
// https://core.telegram.org/bots/api#setgamescore
func (b *Bot) SetGameScore(userId int, score int, options map[string]interface{}) (result ApiResponseMessage) {
	// essential params
	params := map[string]interface{}{
		"user_id": userId,
		"score":   score,
	}
	// optional params
	for key, val := range options {
		if val != nil {
			params[key] = val
		}
	}

	return b.requestResponseMessage("setGameScore", params)
}

// Get high scores of a game.
//
// required options: chat_id + message_id (when inline_message_id is not given)
//                or inline_message_id (when chat_id & message_id is not given)
//
// https://core.telegram.org/bots/api#getgamehighscores
func (b *Bot) GetGameHighScores(userId int, options map[string]interface{}) (result ApiResponseGameHighScores) {
	// essential params
	params := map[string]interface{}{
		"user_id": userId,
	}
	// optional params
	for key, val := range options {
		if val != nil {
			params[key] = val
		}
	}

	return b.requestResponseGameHighScores("getGameHighScores", params)
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
	case int64:
		if intValue, ok := param.(int64); ok {
			return strconv.FormatInt(intValue, 10), ok
		} else {
			b.error("parameter '%+v' could not be cast to int64 value", param)
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
	case ChatAction:
		if value, ok := param.(ChatAction); ok {
			return string(value), ok
		} else {
			b.error("parameter '%+v' could not be cast to string value", param)
		}
	case ParseMode:
		if value, ok := param.(ParseMode); ok {
			return string(value), ok
		} else {
			b.error("parameter '%+v' could not be cast to string value", param)
		}
	case []interface{},
		InlineKeyboardMarkup, ReplyKeyboardMarkup, ReplyKeyboardHide, ForceReply,
		InlineQueryResultCachedAudio, InlineQueryResultCachedDocument, InlineQueryResultCachedGif, InlineQueryResultCachedMpeg4Gif, InlineQueryResultGame, InlineQueryResultCachedPhoto, InlineQueryResultCachedSticker, InlineQueryResultCachedVideo, InlineQueryResultCachedVoice,
		InlineQueryResultArticle, InlineQueryResultAudio, InlineQueryResultContact, InlineQueryResultDocument, InlineQueryResultGif, InlineQueryResultLocation, InlineQueryResultMpeg4Gif, InlineQueryResultPhoto, InlineQueryResultVenue, InlineQueryResultVideo, InlineQueryResultVoice:
		if json, err := json.Marshal(param); err == nil {
			return string(json), true
		} else {
			b.error(err.Error())
		}
	default:
		b.error("unexpected type: '%+v' (%T)", param, param)
	}

	return "", false
}

// Send request to API server and return the response as bytes(synchronously).
//
// NOTE: If *os.File is included in the params, it will be closed automatically in this function.
func (b *Bot) request(method string, params map[string]interface{}) (respBytes []byte, success bool) {
	client := &http.Client{}
	apiUrl := fmt.Sprintf("%s%s/%s", ApiBaseUrl, b.token, method)

	b.verbose("sending request to api url: %s, params: %#v", apiUrl, params)

	if checkIfFileParamExists(params) { // multipart form data
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
				defer resp.Body.Close()

				if bytes, err := ioutil.ReadAll(resp.Body); err == nil {
					return bytes, true
				} else {
					b.error("response read error: %s", err.Error())
				}
			} else {
				b.error("request error: %s", err.Error())
			}
		} else {
			b.error("building request error: %s", err.Error())
		}
	} else { // www-form urlencoded
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
				defer resp.Body.Close()

				if bytes, err := ioutil.ReadAll(resp.Body); err == nil {
					return bytes, true
				} else {
					b.error("response read error: %s", err.Error())
				}
			} else {
				b.error("request error: %s", err.Error())
			}
		} else {
			b.error("building request error: %s", err.Error())
		}
	}

	return []byte{}, false
}

// Send request for ApiResponse and fetch its result.
func (b *Bot) requestResponse(method string, params map[string]interface{}) (result ApiResponse) {
	var errStr string

	if bytes, success := b.request(method, params); success {
		var jsonResponse ApiResponse
		if err := json.Unmarshal(bytes, &jsonResponse); err == nil {
			return jsonResponse
		} else {
			errStr = fmt.Sprintf("json parse error: %s (%s)", err.Error(), string(bytes))
		}
	} else {
		errStr = fmt.Sprintf("%s failed", method)
	}

	b.error(errStr)

	return ApiResponse{ApiResponseBase: ApiResponseBase{Ok: false, Description: &errStr}}
}

// Send request for ApiResponseUser and fetch its result.
func (b *Bot) requestResponseUser(method string, params map[string]interface{}) (result ApiResponseUser) {
	var errStr string

	if bytes, success := b.request(method, params); success {
		var jsonResponse ApiResponseUser
		if err := json.Unmarshal(bytes, &jsonResponse); err == nil {
			return jsonResponse
		} else {
			errStr = fmt.Sprintf("json parse error: %s (%s)", err.Error(), string(bytes))
		}
	} else {
		errStr = fmt.Sprintf("%s failed", method)
	}

	b.error(errStr)

	return ApiResponseUser{ApiResponseBase: ApiResponseBase{Ok: false, Description: &errStr}}
}

// Send request for ApiResponseMessage and fetch its result.
func (b *Bot) requestResponseMessage(method string, params map[string]interface{}) (result ApiResponseMessage) {
	var errStr string

	if bytes, success := b.request(method, params); success {
		var jsonResponse ApiResponseMessage
		if err := json.Unmarshal(bytes, &jsonResponse); err == nil {
			return jsonResponse
		} else {
			errStr = fmt.Sprintf("json parse error: %s (%s)", err.Error(), string(bytes))
		}
	} else {
		errStr = fmt.Sprintf("%s failed", method)
	}

	b.error(errStr)

	return ApiResponseMessage{ApiResponseBase: ApiResponseBase{Ok: false, Description: &errStr}}
}

// Send request for ApiResponseUserProfilePhotos and fetch its result.
func (b *Bot) requestResponseUserProfilePhotos(method string, params map[string]interface{}) (result ApiResponseUserProfilePhotos) {
	var errStr string

	if bytes, success := b.request(method, params); success {
		var jsonResponse ApiResponseUserProfilePhotos
		if err := json.Unmarshal(bytes, &jsonResponse); err == nil {
			return jsonResponse
		} else {
			errStr = fmt.Sprintf("json parse error: %s (%s)", err.Error(), string(bytes))
		}
	} else {
		errStr = fmt.Sprintf("%s failed", method)
	}

	b.error(errStr)

	return ApiResponseUserProfilePhotos{ApiResponseBase: ApiResponseBase{Ok: false, Description: &errStr}}
}

// Send request for ApiResponseUpdates and fetch its result.
func (b *Bot) requestResponseUpdates(method string, params map[string]interface{}) (result ApiResponseUpdates) {
	var errStr string

	if bytes, success := b.request(method, params); success {
		var jsonResponse ApiResponseUpdates
		if err := json.Unmarshal(bytes, &jsonResponse); err == nil {
			return jsonResponse
		} else {
			errStr = fmt.Sprintf("json parse error: %s (%s)", err.Error(), string(bytes))
		}
	} else {
		errStr = fmt.Sprintf("%s failed", method)
	}

	b.error(errStr)

	return ApiResponseUpdates{ApiResponseBase: ApiResponseBase{Ok: false, Description: &errStr}}
}

// Send request for ApiResponseFile and fetch its result.
func (b *Bot) requestResponseFile(method string, params map[string]interface{}) (result ApiResponseFile) {
	var errStr string

	if bytes, success := b.request(method, params); success {
		var jsonResponse ApiResponseFile
		if err := json.Unmarshal(bytes, &jsonResponse); err == nil {
			return jsonResponse
		} else {
			errStr = fmt.Sprintf("json parse error: %s (%s)", err.Error(), string(bytes))
		}
	} else {
		errStr = fmt.Sprintf("%s failed", method)
	}

	b.error(errStr)

	return ApiResponseFile{ApiResponseBase: ApiResponseBase{Ok: false, Description: &errStr}}
}

// Send request for ApiResponseChat and fetch its result.
func (b *Bot) requestResponseChat(method string, params map[string]interface{}) (result ApiResponseChat) {
	var errStr string

	if bytes, success := b.request(method, params); success {
		var jsonResponse ApiResponseChat
		if err := json.Unmarshal(bytes, &jsonResponse); err == nil {
			return jsonResponse
		} else {
			errStr = fmt.Sprintf("json parse error: %s (%s)", err.Error(), string(bytes))
		}
	} else {
		errStr = fmt.Sprintf("%s failed", method)
	}

	b.error(errStr)

	return ApiResponseChat{ApiResponseBase: ApiResponseBase{Ok: false, Description: &errStr}}
}

// Send request for ApiResponseChatAdministrator and fetch its result.
func (b *Bot) requestResponseChatAdministrators(method string, params map[string]interface{}) (result ApiResponseChatAdministrators) {
	var errStr string

	if bytes, success := b.request(method, params); success {
		var jsonResponse ApiResponseChatAdministrators
		if err := json.Unmarshal(bytes, &jsonResponse); err == nil {
			return jsonResponse
		} else {
			errStr = fmt.Sprintf("json parse error: %s (%s)", err.Error(), string(bytes))
		}
	} else {
		errStr = fmt.Sprintf("%s failed", method)
	}

	b.error(errStr)

	return ApiResponseChatAdministrators{ApiResponseBase: ApiResponseBase{Ok: false, Description: &errStr}}
}

// Send request for ApiResponseChatMember and fetch its result.
func (b *Bot) requestResponseChatMember(method string, params map[string]interface{}) (result ApiResponseChatMember) {
	var errStr string

	if bytes, success := b.request(method, params); success {
		var jsonResponse ApiResponseChatMember
		if err := json.Unmarshal(bytes, &jsonResponse); err == nil {
			return jsonResponse
		} else {
			errStr = fmt.Sprintf("json parse error: %s (%s)", err.Error(), string(bytes))
		}
	} else {
		errStr = fmt.Sprintf("%s failed", method)
	}

	b.error(errStr)

	return ApiResponseChatMember{ApiResponseBase: ApiResponseBase{Ok: false, Description: &errStr}}
}

// Send request for ApiResponseInt and fetch its result.
func (b *Bot) requestResponseInt(method string, params map[string]interface{}) (result ApiResponseInt) {
	var errStr string

	if bytes, success := b.request(method, params); success {
		var jsonResponse ApiResponseInt
		if err := json.Unmarshal(bytes, &jsonResponse); err == nil {
			return jsonResponse
		} else {
			errStr = fmt.Sprintf("json parse error: %s (%s)", err.Error(), string(bytes))
		}
	} else {
		errStr = fmt.Sprintf("%s failed", method)
	}

	b.error(errStr)

	return ApiResponseInt{ApiResponseBase: ApiResponseBase{Ok: false, Description: &errStr}}
}

// Send request for ApiResponseGameHighScores and fetch its result.
func (b *Bot) requestResponseGameHighScores(method string, params map[string]interface{}) (result ApiResponseGameHighScores) {
	var errStr string

	if bytes, success := b.request(method, params); success {
		var jsonResponse ApiResponseGameHighScores
		if err := json.Unmarshal(bytes, &jsonResponse); err == nil {
			return jsonResponse
		} else {
			errStr = fmt.Sprintf("json parse error: %s (%s)", err.Error(), string(bytes))
		}
	} else {
		errStr = fmt.Sprintf("%s failed", method)
	}

	b.error(errStr)

	return ApiResponseGameHighScores{ApiResponseBase: ApiResponseBase{Ok: false, Description: &errStr}}
}

// Handle Webhook request.
func (b *Bot) handleWebhook(writer http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()

	b.verbose("received webhook request: %+v", req)

	if body, err := ioutil.ReadAll(req.Body); err == nil {
		var webhook Update
		if err := json.Unmarshal(body, &webhook); err != nil {
			b.error("error while parsing json (%s)", err.Error())
		} else {
			b.verbose("received webhook body: %s", string(body))

			b.updateHandler(b, webhook, nil)
		}
	} else {
		b.error("error while reading webhook request (%s)", err.Error())

		b.updateHandler(b, Update{}, err)
	}
}

// Send file
func (b *Bot) sendFile(chatId interface{}, apiName, paramKey string, filepath *string, options map[string]interface{}) (result ApiResponseMessage) {
	// essential params
	params := map[string]interface{}{
		"chat_id": chatId,
	}
	if isHttpUrl(*filepath) {
		params[paramKey] = *filepath
	} else {
		if file, err := os.Open(*filepath); err == nil {
			params[paramKey] = file
		} else {
			errStr := err.Error()

			b.error(errStr)

			return ApiResponseMessage{
				ApiResponseBase: ApiResponseBase{
					Ok:          false,
					Description: &errStr,
				},
			}
		}
	}
	// optional params
	for key, val := range options {
		if val != nil {
			params[key] = val
		}
	}

	return b.requestResponseMessage(apiName, params)
}

// check if given path is http url
func isHttpUrl(path string) bool {
	if strings.HasPrefix(path, "http://") || strings.HasPrefix(path, "https://") {
		return true
	}
	return false
}
