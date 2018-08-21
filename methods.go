package telegrambot

// https://core.telegram.org/bots/api#available-methods

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
	"time"
)

const (
	httpTimeoutSeconds = 5
)

// GetUpdates retrieves updates from Telegram bot API.
//
// options include: offset, limit, timeout, and allowed_updates.
//
// https://core.telegram.org/bots/api#getupdates
func (b *Bot) GetUpdates(options map[string]interface{}) (result APIResponseUpdates) {
	// optional params
	params := map[string]interface{}{}
	for key, val := range options {
		if val != nil {
			params[key] = val
		}
	}

	return b.requestResponseUpdates("getUpdates", params)
}

// SetWebhookWithOptions sets webhook url, certificate, and various options for receiving incoming updates.
//
// port should be one of: 443, 80, 88, or 8443.
// default maxConnections = 40
//
// https://core.telegram.org/bots/api#setwebhook
func (b *Bot) SetWebhookWithOptions(host string, port int, certFilepath string, maxConnections int, allowedUpdates []UpdateType) (result APIResponseBool) {
	b.webhookHost = host
	b.webhookPort = port
	b.webhookURL = b.getWebhookURL()

	file, err := os.Open(certFilepath)
	if err != nil {
		panic(err)
	}

	params := map[string]interface{}{
		"url":             b.webhookURL,
		"certificate":     file,
		"max_connections": maxConnections,
		"allowed_updates": allowedUpdates,
	}

	b.verbose("setting webhook url to: %s", b.webhookURL)

	return b.requestResponseBool("setWebhook", params)
}

// SetWebhook sets webhook url and certificate for receiving incoming updates.
func (b *Bot) SetWebhook(host string, port int, certFilepath string) (result APIResponseBool) {
	return b.SetWebhookWithOptions(host, port, certFilepath, 40, []UpdateType{})
}

// DeleteWebhook deletes webhook for this bot.
// (Function GetUpdates will not work if webhook is set, so in that case you'll need to delete it)
//
// https://core.telegram.org/bots/api#deletewebhook
func (b *Bot) DeleteWebhook() (result APIResponseBool) {
	b.webhookHost = ""
	b.webhookPort = 0
	b.webhookURL = ""

	b.verbose("deleting webhook url")

	return b.requestResponseBool("deleteWebhook", map[string]interface{}{})
}

// GetWebhookInfo gets webhook info for this bot.
//
// https://core.telegram.org/bots/api#getwebhookinfo
func (b *Bot) GetWebhookInfo() (result APIResponseWebhookInfo) {
	return b.requestResponseWebhookInfo()
}

// GetMe gets info of this bot.
//
// https://core.telegram.org/bots/api#getme
func (b *Bot) GetMe() (result APIResponseUser) {
	return b.requestResponseUser("getMe", map[string]interface{}{}) // no params
}

// SendMessage sends a message to the bot.
//
// options include: parse_mode, disable_web_page_preview, disable_notification, reply_to_message_id, and reply_markup.
//
// https://core.telegram.org/bots/api#sendmessage
func (b *Bot) SendMessage(chatID ChatID, text string, options map[string]interface{}) (result APIResponseMessage) {
	// essential params
	params := map[string]interface{}{
		"chat_id": chatID,
		"text":    text,
	}
	// optional params
	for key, val := range options {
		if val != nil {
			params[key] = val
		}
	}

	return b.requestResponseMessage("sendMessage", params)
}

// ForwardMessage forwards a message.
//
// options include: disable_notification
//
// https://core.telegram.org/bots/api#forwardmessage
func (b *Bot) ForwardMessage(chatID, fromChatID ChatID, messageID int) (result APIResponseMessage) {
	// essential params
	params := map[string]interface{}{
		"chat_id":      chatID,
		"from_chat_id": fromChatID,
		"message_id":   messageID,
	}

	return b.requestResponseMessage("forwardMessage", params)
}

// SendPhoto sends a photo.
//
// options include: caption, parse_mode, disable_notification, reply_to_message_id, and reply_markup.
//
// https://core.telegram.org/bots/api#sendphoto
func (b *Bot) SendPhoto(chatID ChatID, photo InputFile, options map[string]interface{}) (result APIResponseMessage) {
	if options == nil {
		options = map[string]interface{}{}
	}

	// essential params
	options["chat_id"] = chatID
	options["photo"] = photo

	return b.requestResponseMessage("sendPhoto", options)
}

// SendAudio sends an audio file. (.mp3 format only, will be played with external players)
//
// options include: caption, parse_mode, duration, performer, title, disable_notification, reply_to_message_id, and reply_markup.
//
// https://core.telegram.org/bots/api#sendaudio
func (b *Bot) SendAudio(chatID ChatID, audio InputFile, options map[string]interface{}) (result APIResponseMessage) {
	if options == nil {
		options = map[string]interface{}{}
	}

	// essential params
	options["chat_id"] = chatID
	options["audio"] = audio

	return b.requestResponseMessage("sendAudio", options)
}

// SendDocument sends a general file.
//
// options include: caption, parse_mode, disable_notification, reply_to_message_id, and reply_markup.
//
// https://core.telegram.org/bots/api#senddocument
func (b *Bot) SendDocument(chatID ChatID, document InputFile, options map[string]interface{}) (result APIResponseMessage) {
	if options == nil {
		options = map[string]interface{}{}
	}

	// essential params
	options["chat_id"] = chatID
	options["document"] = document

	return b.requestResponseMessage("sendDocument", options)
}

// SendSticker sends a sticker.
//
// options include: disable_notification, reply_to_message_id, and reply_markup.
//
// https://core.telegram.org/bots/api#sendsticker
func (b *Bot) SendSticker(chatID ChatID, sticker InputFile, options map[string]interface{}) (result APIResponseMessage) {
	if options == nil {
		options = map[string]interface{}{}
	}

	// essential params
	options["chat_id"] = chatID
	options["sticker"] = sticker

	return b.requestResponseMessage("sendSticker", options)
}

// GetStickerSet gets a sticker set.
//
// https://core.telegram.org/bots/api#getstickerset
func (b *Bot) GetStickerSet(name string) (result APIResponseStickerSet) {
	// essential params
	params := map[string]interface{}{
		"name": name,
	}

	return b.requestResponseStickerSet("getStickerSet", params)
}

// UploadStickerFile uploads a sticker file.
//
// https://core.telegram.org/bots/api#uploadstickerfile
func (b *Bot) UploadStickerFile(userID int, sticker InputFile) (result APIResponseFile) {
	// essential params
	params := map[string]interface{}{
		"user_id":     userID,
		"png_sticker": sticker,
	}

	return b.requestResponseFile("uploadStickerFile", params)
}

// CreateNewStickerSet creates a new sticker set.
//
// options include: contains_masks and mask_position
//
// https://core.telegram.org/bots/api#createnewstickerset
func (b *Bot) CreateNewStickerSet(userID int, name, title string, sticker InputFile, emojis string, options map[string]interface{}) (result APIResponseBool) {
	// essential params
	params := map[string]interface{}{
		"user_id":     userID,
		"name":        name,
		"title":       title,
		"emojis":      emojis,
		"png_sticker": sticker,
	}

	// optional params
	for key, val := range options {
		if val != nil {
			params[key] = val
		}
	}

	return b.requestResponseBool("createNewStickerSet", params)
}

// AddStickerToSet adds a sticker to set.
//
// options include: mask_position
//
// https://core.telegram.org/bots/api#addstickertoset
func (b *Bot) AddStickerToSet(userID int, name string, sticker InputFile, emojis string, options map[string]interface{}) (result APIResponseBool) {
	// essential params
	params := map[string]interface{}{
		"user_id":     userID,
		"name":        name,
		"emojis":      emojis,
		"png_sticker": sticker,
	}

	// optional params
	for key, val := range options {
		if val != nil {
			params[key] = val
		}
	}

	return b.requestResponseBool("addStickerToSet", params)
}

// SetStickerPositionInSet sets sticker position in set.
//
// https://core.telegram.org/bots/api#setstickerpositioninset
func (b *Bot) SetStickerPositionInSet(sticker string, position int) (result APIResponseBool) {
	// essential params
	params := map[string]interface{}{
		"sticker":  sticker,
		"position": position,
	}

	return b.requestResponseBool("setStickerPositionInSet", params)
}

// DeleteStickerFromSet deletes a sticker from set.
//
// https://core.telegram.org/bots/api#deletestickerfromset
func (b *Bot) DeleteStickerFromSet(sticker string) (result APIResponseBool) {
	// essential params
	params := map[string]interface{}{
		"sticker": sticker,
	}

	return b.requestResponseBool("deleteStickerFromSet", params)
}

// SendVideo sends a video file.
//
// options include: duration, caption, parse_mode, supports_streaming, disable_notification, reply_to_message_id, and reply_markup.
//
// https://core.telegram.org/bots/api#sendvideo
func (b *Bot) SendVideo(chatID ChatID, video InputFile, options map[string]interface{}) (result APIResponseMessage) {
	if options == nil {
		options = map[string]interface{}{}
	}

	// essential params
	options["chat_id"] = chatID
	options["video"] = video

	return b.requestResponseMessage("sendVideo", options)
}

// SendAnimation sends an animation.
//
// options include: duration, width, height, thumb, caption, parse_mode, disable_notification, reply_to_message_id, and reply_markup.
//
// https://core.telegram.org/bots/api#sendanimation
func (b *Bot) SendAnimation(chatID ChatID, animation InputFile, options map[string]interface{}) (result APIResponseMessage) {
	if options == nil {
		options = map[string]interface{}{}
	}

	// essential params
	options["chat_id"] = chatID
	options["animation"] = animation

	return b.requestResponseMessage("sendAnimation", options)
}

// SendVoice sends a voice file. (.ogg format only, will be played with Telegram itself))
//
// options include: caption, parse_mode, duration, disable_notification, reply_to_message_id, and reply_markup.
//
// https://core.telegram.org/bots/api#sendvoice
func (b *Bot) SendVoice(chatID ChatID, voice InputFile, options map[string]interface{}) (result APIResponseMessage) {
	if options == nil {
		options = map[string]interface{}{}
	}

	// essential params
	options["chat_id"] = chatID
	options["voice"] = voice

	return b.requestResponseMessage("sendVoice", options)
}

// SendVideoNote sends a video note.
//
// videoNote cannot be a remote http url (not supported yet)
//
// options include: duration, length, disable_notification, reply_to_message_id, and reply_markup.
// (XXX: API returns 'Bad Request: wrong video note length' when length is not given / 2017.05.19.)
//
// https://core.telegram.org/bots/api#sendvideonote
func (b *Bot) SendVideoNote(chatID ChatID, videoNote InputFile, options map[string]interface{}) (result APIResponseMessage) {
	if options == nil {
		options = map[string]interface{}{}
	}

	// essential params
	options["chat_id"] = chatID
	options["video_note"] = videoNote

	return b.requestResponseMessage("sendVideoNote", options)
}

// SendMediaGroup sends a group of photos or videos as an album.
//
// options include: disable_notification, and reply_to_message_id
//
// https://core.telegram.org/bots/api#sendmediagroup
func (b *Bot) SendMediaGroup(chatID ChatID, media []InputMedia, options map[string]interface{}) (result APIResponseMessages) {
	// essential params
	params := map[string]interface{}{
		"chat_id": chatID,
		"media":   media,
	}
	// optional params
	for key, val := range options {
		if val != nil {
			params[key] = val
		}
	}

	return b.requestResponseMessages("sendMediaGroup", params)
}

// SendLocation sends locations.
//
// options include: disable_notification, reply_to_message_id, and reply_markup.
//
// https://core.telegram.org/bots/api#sendlocation
func (b *Bot) SendLocation(chatID ChatID, latitude, longitude float32, options map[string]interface{}) (result APIResponseMessage) {
	// essential params
	params := map[string]interface{}{
		"chat_id":   chatID,
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

// SendVenue sends venues.
//
// options include: foursquare_id, foursquare_type, disable_notification, reply_to_message_id, and reply_markup.
//
// https://core.telegram.org/bots/api#sendvenue
func (b *Bot) SendVenue(chatID ChatID, latitude, longitude float32, title, address string, options map[string]interface{}) (result APIResponseMessage) {
	// essential params
	params := map[string]interface{}{
		"chat_id":   chatID,
		"latitude":  latitude,
		"longitude": longitude,
		"title":     title,
		"address":   address,
	}
	// optional params
	for key, val := range options {
		if val != nil {
			params[key] = val
		}
	}

	return b.requestResponseMessage("sendVenue", params)
}

// SendContact sends contacts.
//
// options include: last_name, vcard, disable_notification, reply_to_message_id, and reply_markup.
//
// https://core.telegram.org/bots/api#sendcontact
func (b *Bot) SendContact(chatID ChatID, phoneNumber, firstName string, options map[string]interface{}) (result APIResponseMessage) {
	// essential params
	params := map[string]interface{}{
		"chat_id":      chatID,
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

// SendChatAction sends chat actions.
//
// https://core.telegram.org/bots/api#sendchataction
func (b *Bot) SendChatAction(chatID ChatID, action ChatAction) (result APIResponseBool) {
	// essential params
	params := map[string]interface{}{
		"chat_id": chatID,
		"action":  action,
	}

	return b.requestResponseBool("sendChatAction", params)
}

// GetUserProfilePhotos gets user profile photos.
//
// options include: offset and limit.
//
// https://core.telegram.org/bots/api#getuserprofilephotos
func (b *Bot) GetUserProfilePhotos(userID int, options map[string]interface{}) (result APIResponseUserProfilePhotos) {
	// essential params
	params := map[string]interface{}{
		"user_id": userID,
	}
	// optional params
	for key, val := range options {
		if val != nil {
			params[key] = val
		}
	}

	return b.requestResponseUserProfilePhotos("getUserProfilePhotos", params)
}

// GetFile gets file info and prepare for download.
//
// https://core.telegram.org/bots/api#getfile
func (b *Bot) GetFile(fileID string) (result APIResponseFile) {
	// essential params
	params := map[string]interface{}{
		"file_id": fileID,
	}

	return b.requestResponseFile("getFile", params)
}

// GetFileURL gets download link from a given File.
func (b *Bot) GetFileURL(file File) string {
	return fmt.Sprintf("%s%s/%s", fileBaseURL, b.token, *file.FilePath)
}

// KickChatMember kicks a chat member
//
// https://core.telegram.org/bots/api#kickchatmember
func (b *Bot) KickChatMember(chatID ChatID, userID int) (result APIResponseBool) {
	// essential params
	params := map[string]interface{}{
		"chat_id": chatID,
		"user_id": userID,
	}

	return b.requestResponseBool("kickChatMember", params)
}

// KickChatMemberUntil kicks a chat member until given date
func (b *Bot) KickChatMemberUntil(chatID ChatID, userID int, untilDate int) (result APIResponseBool) {
	// essential params
	params := map[string]interface{}{
		"chat_id":    chatID,
		"user_id":    userID,
		"until_date": untilDate,
	}

	return b.requestResponseBool("kickChatMember", params)
}

// LeaveChat leaves a chat
//
// https://core.telegram.org/bots/api#leavechat
func (b *Bot) LeaveChat(chatID ChatID) (result APIResponseBool) {
	// essential params
	params := map[string]interface{}{
		"chat_id": chatID,
	}

	return b.requestResponseBool("leaveChat", params)
}

// UnbanChatMember unbans a chat member
//
// https://core.telegram.org/bots/api#unbanchatmember
func (b *Bot) UnbanChatMember(chatID ChatID, userID int) (result APIResponseBool) {
	// essential params
	params := map[string]interface{}{
		"chat_id": chatID,
		"user_id": userID,
	}

	return b.requestResponseBool("unbanChatMember", params)
}

// RestrictChatMember restricts a chat member
//
// options include: until_date, can_send_messages, can_send_media_messages, can_send_other_messages, and can_send_web_page_previews
//
// https://core.telegram.org/bots/api#restrictchatmember
func (b *Bot) RestrictChatMember(chatID ChatID, userID int, options map[string]interface{}) (result APIResponseBool) {
	// essential params
	params := map[string]interface{}{
		"chat_id": chatID,
		"user_id": userID,
	}
	// optional params
	for key, val := range options {
		if val != nil {
			params[key] = val
		}
	}

	return b.requestResponseBool("restrictChatMember", params)
}

// PromoteChatMember promotes a chat member
//
// options include: can_change_info, can_post_messages, can_edit_messages, can_delete_messages, can_invite_users, can_restrict_members, can_pin_messages, and can_promote_members
//
// https://core.telegram.org/bots/api#promotechatmember
func (b *Bot) PromoteChatMember(chatID ChatID, userID int, options map[string]interface{}) (result APIResponseBool) {
	// essential params
	params := map[string]interface{}{
		"chat_id": chatID,
		"user_id": userID,
	}
	// optional params
	for key, val := range options {
		if val != nil {
			params[key] = val
		}
	}

	return b.requestResponseBool("promoteChatMember", params)
}

// ExportChatInviteLink exports a chat invite link
//
// https://core.telegram.org/bots/api#exportchatinvitelink
func (b *Bot) ExportChatInviteLink(chatID ChatID) (result APIResponseString) {
	// essential params
	params := map[string]interface{}{
		"chat_id": chatID,
	}

	return b.requestResponseString("exportChatInviteLink", params)
}

// SetChatPhoto sets a chat photo
//
// https://core.telegram.org/bots/api#setchatphoto
func (b *Bot) SetChatPhoto(chatID ChatID, photo InputFile) (result APIResponseBool) {
	// essential params
	params := map[string]interface{}{
		"chat_id": chatID,
		"photo":   photo,
	}

	return b.requestResponseBool("setChatPhoto", params)
}

// DeleteChatPhoto deletes a chat photo
//
// https://core.telegram.org/bots/api#deletechatphoto
func (b *Bot) DeleteChatPhoto(chatID ChatID) (result APIResponseBool) {
	// essential params
	params := map[string]interface{}{
		"chat_id": chatID,
	}

	return b.requestResponseBool("deleteChatPhoto", params)
}

// SetChatTitle sets a chat title
//
// https://core.telegram.org/bots/api#setchattitle
func (b *Bot) SetChatTitle(chatID ChatID, title string) (result APIResponseBool) {
	// essential params
	params := map[string]interface{}{
		"chat_id": chatID,
		"title":   title,
	}

	return b.requestResponseBool("setChatTitle", params)
}

// SetChatDescription sets a chat description
//
// https://core.telegram.org/bots/api#setchatdescription
func (b *Bot) SetChatDescription(chatID ChatID, description string) (result APIResponseBool) {
	// essential params
	params := map[string]interface{}{
		"chat_id":     chatID,
		"description": description,
	}

	return b.requestResponseBool("setChatDescription", params)
}

// PinChatMessage pins a chat message
//
// options include: disable_notification
//
// https://core.telegram.org/bots/api#pinchatmessage
func (b *Bot) PinChatMessage(chatID ChatID, messageID int, options map[string]interface{}) (result APIResponseBool) {
	// essential params
	params := map[string]interface{}{
		"chat_id":    chatID,
		"message_id": messageID,
	}
	// optional params
	for key, val := range options {
		if val != nil {
			params[key] = val
		}
	}

	return b.requestResponseBool("pinChatMessage", params)
}

// UnpinChatMessage unpins a chat message
//
// https://core.telegram.org/bots/api#unpinchatmessage
func (b *Bot) UnpinChatMessage(chatID ChatID) (result APIResponseBool) {
	// essential params
	params := map[string]interface{}{
		"chat_id": chatID,
	}

	return b.requestResponseBool("unpinChatMessage", params)
}

// GetChat gets a chat
//
// https://core.telegram.org/bots/api#getchat
func (b *Bot) GetChat(chatID ChatID) (result APIResponseChat) {
	// essential params
	params := map[string]interface{}{
		"chat_id": chatID,
	}

	return b.requestResponseChat("getChat", params)
}

// GetChatAdministrators gets chat administrators
//
// https://core.telegram.org/bots/api#getchatadministrators
func (b *Bot) GetChatAdministrators(chatID ChatID) (result APIResponseChatAdministrators) {
	// essential params
	params := map[string]interface{}{
		"chat_id": chatID,
	}

	return b.requestResponseChatAdministrators("getChatAdministrators", params)
}

// GetChatMembersCount gets chat members' count
//
// https://core.telegram.org/bots/api#getchatmemberscount
func (b *Bot) GetChatMembersCount(chatID ChatID) (result APIResponseInt) {
	// essential params
	params := map[string]interface{}{
		"chat_id": chatID,
	}

	return b.requestResponseInt("getChatMembersCount", params)
}

// GetChatMember gets a chat member
//
// https://core.telegram.org/bots/api#getchatmember
func (b *Bot) GetChatMember(chatID ChatID, userID int) (result APIResponseChatMember) {
	// essential params
	params := map[string]interface{}{
		"chat_id": chatID,
		"user_id": userID,
	}

	return b.requestResponseChatMember("getChatMember", params)
}

// SetChatStickerSet sets a chat sticker set
//
// https://core.telegram.org/bots/api#setchatstickerset
func (b *Bot) SetChatStickerSet(chatID ChatID, stickerSetName string) (result APIResponseBool) {
	// essential params
	params := map[string]interface{}{
		"chat_id":          chatID,
		"sticker_set_name": stickerSetName,
	}

	return b.requestResponseBool("setChatStickerSet", params)
}

// DeleteChatStickerSet deletes a chat sticker set
//
// https://core.telegram.org/bots/api#deletechatstickerset
func (b *Bot) DeleteChatStickerSet(chatID ChatID) (result APIResponseBool) {
	// essential params
	params := map[string]interface{}{
		"chat_id": chatID,
	}

	return b.requestResponseBool("deleteChatStickerSet", params)
}

// AnswerCallbackQuery answers a callback query
//
// options include: text, show_alert, url, and cache_time
//
// https://core.telegram.org/bots/api#answercallbackquery
func (b *Bot) AnswerCallbackQuery(callbackQueryID string, options map[string]interface{}) (result APIResponseBool) {
	// essential params
	params := map[string]interface{}{
		"callback_query_id": callbackQueryID,
	}
	// optional params
	for key, val := range options {
		if val != nil {
			params[key] = val
		}
	}

	return b.requestResponseBool("answerCallbackQuery", params)
}

// Updating messages
//
// https://core.telegram.org/bots/api#updating-messages

// EditMessageText edits text of a message
//
// required options: chat_id + message_id (when inline_message_id is not given)
//                or inline_message_id (when chat_id & message_id is not given)
//
// other options: parse_mode, disable_web_page_preview, and reply_markup
//
// https://core.telegram.org/bots/api#editmessagetext
func (b *Bot) EditMessageText(text string, options map[string]interface{}) (result APIResponseMessageOrBool) {
	// essential params
	params := map[string]interface{}{
		"text": text,
	}
	// optional params
	for key, val := range options {
		if val != nil {
			params[key] = val
		}
	}

	return b.requestResponseMessageOrBool("editMessageText", params)
}

// EditMessageCaption edits caption of a message
//
// required options: chat_id + message_id (when inline_message_id is not given)
//                or inline_message_id (when chat_id & message_id is not given)
//
// other options: parse_mode, or reply_markup
//
// https://core.telegram.org/bots/api#editmessagecaption
func (b *Bot) EditMessageCaption(caption string, options map[string]interface{}) (result APIResponseMessageOrBool) {
	// essential params
	params := map[string]interface{}{
		"caption": caption,
	}
	// optional params
	for key, val := range options {
		if val != nil {
			params[key] = val
		}
	}

	return b.requestResponseMessageOrBool("editMessageCaption", params)
}

// EditMessageMedia edites a media message
//
// required options: chat_id + message_id (when inline_message_id is not given)
//                or inline_message_id (when chat_id & message_id is not given)
//
// other options: reply_markup
//
// https://core.telegram.org/bots/api#editmessagemedia
func (b *Bot) EditMessageMedia(media InputMedia, options map[string]interface{}) (result APIResponseMessageOrBool) {
	// essential params
	params := map[string]interface{}{
		"media": media,
	}
	// optional params
	for key, val := range options {
		if val != nil {
			params[key] = val
		}
	}

	return b.requestResponseMessageOrBool("editMessageMedia", params)
}

// EditMessageReplyMarkup edits reply markup of a message
//
// required options: chat_id + message_id (when inline_message_id is not given)
//                or inline_message_id (when chat_id & message_id is not given)
//
// other options: reply_markup
//
// https://core.telegram.org/bots/api#editmessagereplymarkup
func (b *Bot) EditMessageReplyMarkup(options map[string]interface{}) (result APIResponseMessageOrBool) {
	return b.requestResponseMessageOrBool("editMessageReplyMarkup", options)
}

// EditMessageLiveLocation edits live location of a message
//
// required options: chat_id + message_id (when inline_message_id is not given)
//                or inline_message_id (when chat_id & message_id is not given)
//
// other options: reply_markup
//
// https://core.telegram.org/bots/api#editmessagelivelocation
func (b *Bot) EditMessageLiveLocation(latitude, longitude float32, options map[string]interface{}) (result APIResponseMessageOrBool) {
	// essential params
	params := map[string]interface{}{
		"latitude":  latitude,
		"longitude": longitude,
	}
	// optional params
	for key, val := range options {
		if val != nil {
			params[key] = val
		}
	}

	return b.requestResponseMessageOrBool("editMessageLiveLocation", params)
}

// StopMessageLiveLocation stops live location of a message
//
// required options: chat_id + message_id (when inline_message_id is not given)
//                or inline_message_id (when chat_id & message_id is not given)
//
// other options: reply_markup
//
// https://core.telegram.org/bots/api#stopmessagelivelocation
func (b *Bot) StopMessageLiveLocation(options map[string]interface{}) (result APIResponseMessageOrBool) {
	return b.requestResponseMessageOrBool("stopMessageLiveLocation", options)
}

// DeleteMessage deletes a message
//
// https://core.telegram.org/bots/api#deletemessage
func (b *Bot) DeleteMessage(chatID ChatID, messageID int) (result APIResponseBool) {
	return b.requestResponseBool("deleteMessage", map[string]interface{}{
		"chat_id":    chatID,
		"message_id": messageID,
	})
}

// AnswerInlineQuery sends answers to an inline query.
//
// results = array of InlineQueryResultArticle, InlineQueryResultPhoto, InlineQueryResultGif, InlineQueryResultMpeg4Gif, or InlineQueryResultVideo.
//
// options include: cache_time, is_personal, next_offset, switch_pm_text, and switch_pm_parameter.
//
// https://core.telegram.org/bots/api#answerinlinequery
func (b *Bot) AnswerInlineQuery(inlineQueryID string, results []interface{}, options map[string]interface{}) (result APIResponseBool) {
	// essential params
	params := map[string]interface{}{
		"inline_query_id": inlineQueryID,
		"results":         results,
	}
	// optional params
	for key, val := range options {
		if val != nil {
			params[key] = val
		}
	}

	return b.requestResponseBool("answerInlineQuery", params)
}

// SendInvoice sends an invoice.
//
// options include: provider_data, photo_url, photo_size, photo_width, photo_height, need_name, need_phone_number, need_email, need_shipping_address, is_flexible, disable_notification, reply_to_message_id, and reply_markup
//
// https://core.telegram.org/bots/api#sendinvoice
func (b *Bot) SendInvoice(chatID int64, title, description, payload, providerToken, startParameter, currency string, prices []LabeledPrice, options map[string]interface{}) (result APIResponseMessage) {
	// essential params
	params := map[string]interface{}{
		"chat_id":         chatID,
		"title":           title,
		"description":     description,
		"payload":         payload,
		"provider_token":  providerToken,
		"start_parameter": startParameter,
		"currency":        currency,
		"prices":          prices,
	}
	// optional params
	for key, val := range options {
		if val != nil {
			params[key] = val
		}
	}

	return b.requestResponseMessage("sendInvoice", params)
}

// AnswerShippingQuery answers a shipping query.
//
// if ok is true, shippingOptions should be provided.
// otherwise, errorMessage should be provided.
//
// https://core.telegram.org/bots/api#answershippingquery
func (b *Bot) AnswerShippingQuery(shippingQueryID string, ok bool, shippingOptions []ShippingOption, errorMessage *string) (result APIResponseBool) {
	// essential params
	params := map[string]interface{}{
		"shipping_query_id": shippingQueryID,
		"ok":                ok,
	}
	// optional params
	if ok {
		if len(shippingOptions) > 0 {
			params["shipping_options"] = shippingOptions
		}
	} else {
		if errorMessage != nil {
			params["error_message"] = *errorMessage
		}
	}

	return b.requestResponseBool("answerShippingQuery", params)
}

// AnswerPreCheckoutQuery answers a pre-checkout query.
//
// https://core.telegram.org/bots/api#answerprecheckoutquery
func (b *Bot) AnswerPreCheckoutQuery(preCheckoutQueryID string, ok bool, errorMessage *string) (result APIResponseBool) {
	// essential params
	params := map[string]interface{}{
		"pre_checkout_query_id": preCheckoutQueryID,
		"ok": ok,
	}
	// optional params
	if !ok {
		if errorMessage != nil {
			params["error_message"] = *errorMessage
		}
	}

	return b.requestResponseBool("answerPreCheckoutQuery", params)
}

// SendGame sends a game.
//
// options include: disable_notification, reply_to_message_id, and reply_markup.
//
// https://core.telegram.org/bots/api#sendgame
func (b *Bot) SendGame(chatID ChatID, gameShortName string, options map[string]interface{}) (result APIResponseMessage) {
	// essential params
	params := map[string]interface{}{
		"chat_id":         chatID,
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

// SetGameScore sets score of a game.
//
// required options: chat_id + message_id (when inline_message_id is not given)
//                or inline_message_id (when chat_id & message_id is not given)
//
// other options: force, and disable_edit_message
//
// https://core.telegram.org/bots/api#setgamescore
func (b *Bot) SetGameScore(userID int, score int, options map[string]interface{}) (result APIResponseMessageOrBool) {
	// essential params
	params := map[string]interface{}{
		"user_id": userID,
		"score":   score,
	}
	// optional params
	for key, val := range options {
		if val != nil {
			params[key] = val
		}
	}

	return b.requestResponseMessageOrBool("setGameScore", params)
}

// GetGameHighScores gets high scores of a game.
//
// required options: chat_id + message_id (when inline_message_id is not given)
//                or inline_message_id (when chat_id & message_id is not given)
//
// https://core.telegram.org/bots/api#getgamehighscores
func (b *Bot) GetGameHighScores(userID int, options map[string]interface{}) (result APIResponseGameHighScores) {
	// essential params
	params := map[string]interface{}{
		"user_id": userID,
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
		case *os.File, []byte:
			return true
		case InputFile:
			if len(value.(InputFile).Bytes) > 0 || value.(InputFile).Filepath != nil {
				return true
			}
		}
	}

	return false
}

// Convert given interface to string. (for HTTP params)
func (b *Bot) paramToString(param interface{}) (result string, success bool) {
	switch param.(type) {
	case int:
		if intValue, ok := param.(int); ok {
			return strconv.Itoa(intValue), true
		}
		b.error("parameter '%+v' could not be cast to int value", param)
	case int64:
		if intValue, ok := param.(int64); ok {
			return strconv.FormatInt(intValue, 10), true
		}
		b.error("parameter '%+v' could not be cast to int64 value", param)
	case float32:
		if floatValue, ok := param.(float32); ok {
			return fmt.Sprintf("%.8f", floatValue), true
		}
		b.error("parameter '%+v' could not be cast to float32 value", param)
	case bool:
		if boolValue, ok := param.(bool); ok {
			return strconv.FormatBool(boolValue), true
		}
		b.error("parameter '%+v' could not be cast to bool value", param)
	case string:
		if strValue, ok := param.(string); ok {
			return strValue, true
		}
		b.error("parameter '%+v' could not be cast to string value", param)
	case ChatAction:
		if value, ok := param.(ChatAction); ok {
			return string(value), true
		}
		b.error("parameter '%+v' could not be cast to string value", param)
	case ParseMode:
		if value, ok := param.(ParseMode); ok {
			return string(value), true
		}
		b.error("parameter '%+v' could not be cast to string value", param)
	case InputFile:
		if value, ok := param.(InputFile); ok {
			if value.URL != nil {
				return *value.URL, true
			}
			if value.FileID != nil {
				return *value.FileID, true
			}
		}
		b.error("parameter '%+v' could not be cast to string value", param)
	default:
		json, err := json.Marshal(param)
		if err == nil {
			return string(json), true
		}
		b.error("parameter '%+v' could not be encoded as json: %s", param, err)
	}

	return "", false
}

// Send request to API server and return the response as bytes(synchronously).
//
// NOTE: If *os.File is included in the params, it will be closed automatically in this function.
func (b *Bot) request(method string, params map[string]interface{}) (respBytes []byte, err error) {
	client := &http.Client{
		Timeout: httpTimeoutSeconds * time.Second,
	}
	apiURL := fmt.Sprintf("%s%s/%s", apiBaseURL, b.token, method)

	b.verbose("sending request to api url: %s, params: %#v", apiURL, params)

	if checkIfFileParamExists(params) { // multipart form data
		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)

		for key, value := range params {
			switch value.(type) {
			case *os.File:
				if file, ok := value.(*os.File); ok {
					defer file.Close()

					var part io.Writer
					part, err = writer.CreateFormFile(key, file.Name())
					if err == nil {
						if _, err = io.Copy(part, file); err != nil {
							b.error("could not write to multipart: %s", key)
						}
					} else {
						b.error("could not create form file for parameter '%s' (%v)", key, value)
					}
				} else {
					b.error("parameter '%s' (%v) could not be cast to file", key, value)
				}
			case []byte:
				if fbytes, ok := value.([]byte); ok {
					filename := fmt.Sprintf("%s.%s", key, getExtension(fbytes))
					var part io.Writer
					part, err = writer.CreateFormFile(key, filename)
					if err == nil {
						if _, err = io.Copy(part, bytes.NewReader(fbytes)); err != nil {
							b.error("could not write to multipart: %s", key)
						}
					} else {
						b.error("could not create form file for parameter '%s' ([]byte)", key)
					}
				} else {
					b.error("parameter '%s' could not be cast to []byte", key)
				}
			case InputFile:
				if inputFile, ok := value.(InputFile); ok {
					if inputFile.Filepath != nil {
						var file *os.File
						if file, err = os.Open(*inputFile.Filepath); err == nil {
							defer file.Close()

							var part io.Writer
							part, err = writer.CreateFormFile(key, file.Name())
							if err == nil {
								if _, err = io.Copy(part, file); err != nil {
									b.error("could not write to multipart: %s", key)
								}
							} else {
								b.error("could not create form file for parameter '%s' (%v)", key, value)
							}
						} else {
							b.error("parameter '%s' (%v) could not be read from file: %s", key, value, err.Error())
						}
					} else if len(inputFile.Bytes) > 0 {
						filename := fmt.Sprintf("%s.%s", key, getExtension(inputFile.Bytes))
						var part io.Writer
						part, err = writer.CreateFormFile(key, filename)
						if err == nil {
							if _, err = io.Copy(part, bytes.NewReader(inputFile.Bytes)); err != nil {
								b.error("could not write InputFile to multipart: %s", key)
							}
						} else {
							b.error("could not create form file for parameter '%s' (InputFile)", key)
						}
					} else {
						if strValue, ok := b.paramToString(value); ok {
							writer.WriteField(key, strValue)
						} else {
							b.error("invalid InputFile parameter '%s'", key)
						}
					}
				} else {
					b.error("parameter '%s' could not be cast to InputFile", key)
				}
			default:
				if strValue, ok := b.paramToString(value); ok {
					writer.WriteField(key, strValue)
				}
			}
		}

		if err = writer.Close(); err != nil {
			b.error("error while closing writer (%s)", err)
		}

		var req *http.Request
		req, err = http.NewRequest("POST", apiURL, body)
		if err == nil {
			req.Header.Add("Content-Type", writer.FormDataContentType()) // due to file parameter
			req.Close = true

			var resp *http.Response
			resp, err = client.Do(req)
			if resp != nil { // XXX - in case of redirect
				defer resp.Body.Close()
			}
			if err == nil {
				// FIXXX: check http status code here
				var bytes []byte
				bytes, err = ioutil.ReadAll(resp.Body)
				if err == nil {
					return bytes, nil
				}

				err = fmt.Errorf("response read error: %s", err)

				b.error(err.Error())
			} else {
				err = fmt.Errorf("request error: %s", err)

				b.error(err.Error())
			}
		} else {
			err = fmt.Errorf("building request error: %s", err)

			b.error(err.Error())
		}
	} else { // www-form urlencoded
		paramValues := url.Values{}
		for key, value := range params {
			if strValue, ok := b.paramToString(value); ok {
				paramValues[key] = []string{strValue}
			}
		}
		encoded := paramValues.Encode()

		var req *http.Request
		req, err = http.NewRequest("POST", apiURL, bytes.NewBufferString(encoded))
		if err == nil {
			req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
			req.Header.Add("Content-Length", strconv.Itoa(len(encoded)))
			req.Close = true

			var resp *http.Response
			resp, err = client.Do(req)
			if resp != nil { // XXX - in case of redirect
				defer resp.Body.Close()
			}
			if err == nil {
				// FIXXX: check http status code here
				var bytes []byte
				bytes, err = ioutil.ReadAll(resp.Body)
				if err == nil {
					return bytes, nil
				}

				err = fmt.Errorf("response read error: %s", err)

				b.error(err.Error())
			} else {
				err = fmt.Errorf("request error: %s", err)

				b.error(err.Error())
			}
		} else {
			err = fmt.Errorf("building request error: %s", err)

			b.error(err.Error())
		}
	}

	return []byte{}, err
}

// Send request for APIResponseWebhookInfo and fetch its result.
func (b *Bot) requestResponseWebhookInfo() (result APIResponseWebhookInfo) {
	var errStr string

	if bytes, err := b.request("getWebhookInfo", map[string]interface{}{}); err == nil {
		var jsonResponse APIResponseWebhookInfo
		err = json.Unmarshal(bytes, &jsonResponse)
		if err == nil {
			return jsonResponse
		}

		errStr = fmt.Sprintf("json parse error: %s (%s)", err, string(bytes))
	} else {
		errStr = fmt.Sprintf("getWebhookInfo failed with error: %s", err)
	}

	b.error(errStr)

	return APIResponseWebhookInfo{APIResponseBase: APIResponseBase{Ok: false, Description: &errStr}}
}

// Send request for APIResponseUser and fetch its result.
func (b *Bot) requestResponseUser(method string, params map[string]interface{}) (result APIResponseUser) {
	var errStr string

	if bytes, err := b.request(method, params); err == nil {
		var jsonResponse APIResponseUser
		err = json.Unmarshal(bytes, &jsonResponse)
		if err == nil {
			return jsonResponse
		}

		errStr = fmt.Sprintf("json parse error: %s (%s)", err, string(bytes))
	} else {
		errStr = fmt.Sprintf("%s failed with error: %s", method, err)
	}

	b.error(errStr)

	return APIResponseUser{APIResponseBase: APIResponseBase{Ok: false, Description: &errStr}}
}

// Send request for APIResponseMessage and fetch its result.
func (b *Bot) requestResponseMessage(method string, params map[string]interface{}) (result APIResponseMessage) {
	var errStr string

	if bytes, err := b.request(method, params); err == nil {
		var jsonResponse APIResponseMessage
		err = json.Unmarshal(bytes, &jsonResponse)
		if err == nil {
			return jsonResponse
		}

		errStr = fmt.Sprintf("json parse error: %s (%s)", err, string(bytes))
	} else {
		errStr = fmt.Sprintf("%s failed with error: %s", method, err)
	}

	b.error(errStr)

	return APIResponseMessage{APIResponseBase: APIResponseBase{Ok: false, Description: &errStr}}
}

// Send request for APIResponseMessages and fetch its result.
func (b *Bot) requestResponseMessages(method string, params map[string]interface{}) (result APIResponseMessages) {
	var errStr string

	if bytes, err := b.request(method, params); err == nil {
		var jsonResponse APIResponseMessages
		err = json.Unmarshal(bytes, &jsonResponse)
		if err == nil {
			return jsonResponse
		}

		errStr = fmt.Sprintf("json parse error: %s (%s)", err, string(bytes))
	} else {
		errStr = fmt.Sprintf("%s failed with error: %s", method, err)
	}

	b.error(errStr)

	return APIResponseMessages{APIResponseBase: APIResponseBase{Ok: false, Description: &errStr}}
}

// Send request for APIResponseUserProfilePhotos and fetch its result.
func (b *Bot) requestResponseUserProfilePhotos(method string, params map[string]interface{}) (result APIResponseUserProfilePhotos) {
	var errStr string

	if bytes, err := b.request(method, params); err == nil {
		var jsonResponse APIResponseUserProfilePhotos
		err = json.Unmarshal(bytes, &jsonResponse)
		if err == nil {
			return jsonResponse
		}

		errStr = fmt.Sprintf("json parse error: %s (%s)", err, string(bytes))
	} else {
		errStr = fmt.Sprintf("%s failed with error: %s", method, err)
	}

	b.error(errStr)

	return APIResponseUserProfilePhotos{APIResponseBase: APIResponseBase{Ok: false, Description: &errStr}}
}

// Send request for APIResponseUpdates and fetch its result.
func (b *Bot) requestResponseUpdates(method string, params map[string]interface{}) (result APIResponseUpdates) {
	var errStr string

	if bytes, err := b.request(method, params); err == nil {
		var jsonResponse APIResponseUpdates
		err = json.Unmarshal(bytes, &jsonResponse)
		if err == nil {
			return jsonResponse
		}

		errStr = fmt.Sprintf("json parse error: %s (%s)", err, string(bytes))
	} else {
		errStr = fmt.Sprintf("%s failed with error: %s", method, err)
	}

	b.error(errStr)

	return APIResponseUpdates{APIResponseBase: APIResponseBase{Ok: false, Description: &errStr}}
}

// Send request for APIResponseFile and fetch its result.
func (b *Bot) requestResponseFile(method string, params map[string]interface{}) (result APIResponseFile) {
	var errStr string

	if bytes, err := b.request(method, params); err == nil {
		var jsonResponse APIResponseFile
		err = json.Unmarshal(bytes, &jsonResponse)
		if err == nil {
			return jsonResponse
		}

		errStr = fmt.Sprintf("json parse error: %s (%s)", err, string(bytes))
	} else {
		errStr = fmt.Sprintf("%s failed with error: %s", method, err)
	}

	b.error(errStr)

	return APIResponseFile{APIResponseBase: APIResponseBase{Ok: false, Description: &errStr}}
}

// Send request for APIResponseChat and fetch its result.
func (b *Bot) requestResponseChat(method string, params map[string]interface{}) (result APIResponseChat) {
	var errStr string

	if bytes, err := b.request(method, params); err == nil {
		var jsonResponse APIResponseChat
		err = json.Unmarshal(bytes, &jsonResponse)
		if err == nil {
			return jsonResponse
		}

		errStr = fmt.Sprintf("json parse error: %s (%s)", err, string(bytes))
	} else {
		errStr = fmt.Sprintf("%s failed with error: %s", method, err)
	}

	b.error(errStr)

	return APIResponseChat{APIResponseBase: APIResponseBase{Ok: false, Description: &errStr}}
}

// Send request for APIResponseChatAdministrator and fetch its result.
func (b *Bot) requestResponseChatAdministrators(method string, params map[string]interface{}) (result APIResponseChatAdministrators) {
	var errStr string

	if bytes, err := b.request(method, params); err == nil {
		var jsonResponse APIResponseChatAdministrators
		err = json.Unmarshal(bytes, &jsonResponse)
		if err == nil {
			return jsonResponse
		}

		errStr = fmt.Sprintf("json parse error: %s (%s)", err, string(bytes))
	} else {
		errStr = fmt.Sprintf("%s failed with error: %s", method, err)
	}

	b.error(errStr)

	return APIResponseChatAdministrators{APIResponseBase: APIResponseBase{Ok: false, Description: &errStr}}
}

// Send request for APIResponseChatMember and fetch its result.
func (b *Bot) requestResponseChatMember(method string, params map[string]interface{}) (result APIResponseChatMember) {
	var errStr string

	if bytes, err := b.request(method, params); err == nil {
		var jsonResponse APIResponseChatMember
		err = json.Unmarshal(bytes, &jsonResponse)
		if err == nil {
			return jsonResponse
		}

		errStr = fmt.Sprintf("json parse error: %s (%s)", err, string(bytes))
	} else {
		errStr = fmt.Sprintf("%s failed with error: %s", method, err)
	}

	b.error(errStr)

	return APIResponseChatMember{APIResponseBase: APIResponseBase{Ok: false, Description: &errStr}}
}

// Send request for APIResponseInt and fetch its result.
func (b *Bot) requestResponseInt(method string, params map[string]interface{}) (result APIResponseInt) {
	var errStr string

	if bytes, err := b.request(method, params); err == nil {
		var jsonResponse APIResponseInt
		err = json.Unmarshal(bytes, &jsonResponse)
		if err == nil {
			return jsonResponse
		}

		errStr = fmt.Sprintf("json parse error: %s (%s)", err, string(bytes))
	} else {
		errStr = fmt.Sprintf("%s failed with error: %s", method, err)
	}

	b.error(errStr)

	return APIResponseInt{APIResponseBase: APIResponseBase{Ok: false, Description: &errStr}}
}

// Send request for APIResponseBool and fetch its result.
func (b *Bot) requestResponseBool(method string, params map[string]interface{}) (result APIResponseBool) {
	var errStr string

	if bytes, err := b.request(method, params); err == nil {
		var jsonResponse APIResponseBool
		err = json.Unmarshal(bytes, &jsonResponse)
		if err == nil {
			return jsonResponse
		}

		errStr = fmt.Sprintf("json parse error: %s (%s)", err, string(bytes))
	} else {
		errStr = fmt.Sprintf("%s failed with error: %s", method, err)
	}

	b.error(errStr)

	return APIResponseBool{APIResponseBase: APIResponseBase{Ok: false, Description: &errStr}}
}

// Send request for APIResponseString and fetch its result.
func (b *Bot) requestResponseString(method string, params map[string]interface{}) (result APIResponseString) {
	var errStr string

	if bytes, err := b.request(method, params); err == nil {
		var jsonResponse APIResponseString
		err = json.Unmarshal(bytes, &jsonResponse)
		if err == nil {
			return jsonResponse
		}

		errStr = fmt.Sprintf("json parse error: %s (%s)", err, string(bytes))
	} else {
		errStr = fmt.Sprintf("%s failed with error: %s", method, err)
	}

	b.error(errStr)

	return APIResponseString{APIResponseBase: APIResponseBase{Ok: false, Description: &errStr}}
}

// Send request for APIResponseGameHighScores and fetch its result.
func (b *Bot) requestResponseGameHighScores(method string, params map[string]interface{}) (result APIResponseGameHighScores) {
	var errStr string

	if bytes, err := b.request(method, params); err == nil {
		var jsonResponse APIResponseGameHighScores
		err = json.Unmarshal(bytes, &jsonResponse)
		if err == nil {
			return jsonResponse
		}

		errStr = fmt.Sprintf("json parse error: %s (%s)", err, string(bytes))
	} else {
		errStr = fmt.Sprintf("%s failed with error: %s", method, err)
	}

	b.error(errStr)

	return APIResponseGameHighScores{APIResponseBase: APIResponseBase{Ok: false, Description: &errStr}}
}

// Send request for APIResponseStickerSet and fetch its result.
func (b *Bot) requestResponseStickerSet(method string, params map[string]interface{}) (result APIResponseStickerSet) {
	var errStr string

	if bytes, err := b.request(method, params); err == nil {
		var jsonResponse APIResponseStickerSet
		err = json.Unmarshal(bytes, &jsonResponse)
		if err == nil {
			return jsonResponse
		}

		errStr = fmt.Sprintf("json parse error: %s (%s)", err, string(bytes))
	} else {
		errStr = fmt.Sprintf("%s failed with error: %s", method, err)
	}

	b.error(errStr)

	return APIResponseStickerSet{APIResponseBase: APIResponseBase{Ok: false, Description: &errStr}}
}

// Send request for APIResponseMessageOrBool and fetch its result.
func (b *Bot) requestResponseMessageOrBool(method string, params map[string]interface{}) (result APIResponseMessageOrBool) {
	var errStr string

	if bytes, err := b.request(method, params); err == nil {
		// try APIResponseMessage type,
		var jsonResponseMessage APIResponseMessage
		err = json.Unmarshal(bytes, &jsonResponseMessage)
		if err == nil {
			return APIResponseMessageOrBool{
				APIResponseBase: APIResponseBase{Ok: true, Description: jsonResponseMessage.Description},
				ResultMessage:   jsonResponseMessage.Result,
			}
		}

		// then try APIResponseBool type,
		var jsonResponseBool APIResponseBool
		err = json.Unmarshal(bytes, &jsonResponseBool)
		if err == nil {
			return APIResponseMessageOrBool{
				APIResponseBase: APIResponseBase{Ok: true, Description: jsonResponseBool.Description},
				ResultBool:      &jsonResponseBool.Result,
			}
		}

		errStr = fmt.Sprintf("json parse error: not in Message nor bool type (%s)", string(bytes))
	} else {
		errStr = fmt.Sprintf("%s failed with error: %s", method, err)
	}

	b.error(errStr)

	return APIResponseMessageOrBool{APIResponseBase: APIResponseBase{Ok: false, Description: &errStr}}
}

// Handle Webhook request.
func (b *Bot) handleWebhook(writer http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()

	b.verbose("received webhook request: %+v", req)

	if body, err := ioutil.ReadAll(req.Body); err == nil {
		var webhook Update
		if err = json.Unmarshal(body, &webhook); err != nil {
			b.error("error while parsing json (%s)", err)
		} else {
			b.verbose("received webhook body: %s", string(body))

			b.updateHandler(b, webhook, nil)
		}
	} else {
		b.error("error while reading webhook request (%s)", err)

		b.updateHandler(b, Update{}, err)
	}
}

// get file extension from bytes array
//
// https://www.w3.org/Protocols/rfc1341/4_Content-Type.html
func getExtension(bytes []byte) string {
	types := strings.Split(http.DetectContentType(bytes), "/") // ex: "image/jpeg"
	if len(types) >= 2 {
		splitted := strings.Split(types[1], ";") // for removing subtype parameter
		if len(splitted) >= 1 {
			return splitted[0] // return subtype only
		}
	}
	return "" // default
}
