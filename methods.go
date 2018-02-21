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
)

// GetUpdates retrieves updates from Telegram bot API.
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

// SetWebhookWithOptions sets webhook url, certificate, and various options for receiving incoming updates.
//
// port should be one of: 443, 80, 88, or 8443.
// default maxConnections = 40
//
// https://core.telegram.org/bots/api#setwebhook
func (b *Bot) SetWebhookWithOptions(host string, port int, certFilepath string, maxConnections int, allowedUpdates []UpdateType) (result ApiResponse) {
	b.webhookHost = host
	b.webhookPort = port
	b.webhookUrl = b.getWebhookUrl()

	file, err := os.Open(certFilepath)
	if err != nil {
		panic(err)
	}

	params := map[string]interface{}{
		"url":             b.webhookUrl,
		"certificate":     file,
		"max_connections": maxConnections,
		"allowed_updates": allowedUpdates,
	}

	b.verbose("setting webhook url to: %s", b.webhookUrl)

	return b.requestResponse("setWebhook", params)
}

// SetWebhook sets webhook url and certificate for receiving incoming updates.
func (b *Bot) SetWebhook(host string, port int, certFilepath string) (result ApiResponse) {
	return b.SetWebhookWithOptions(host, port, certFilepath, 40, []UpdateType{})
}

// DeleteWebhook deletes webhook for this bot.
// (Function GetUpdates will not work if webhook is set, so in that case you'll need to delete it)
//
// https://core.telegram.org/bots/api#deletewebhook
func (b *Bot) DeleteWebhook() (result ApiResponse) {
	b.webhookHost = ""
	b.webhookPort = 0
	b.webhookUrl = ""

	b.verbose("deleting webhook url")

	return b.requestResponse("deleteWebhook", map[string]interface{}{})
}

// GetWebhookInfo gets webhook info for this bot.
//
// https://core.telegram.org/bots/api#getwebhookinfo
func (b *Bot) GetWebhookInfo() (result ApiResponseWebhookInfo) {
	return b.requestResponseWebhookInfo()
}

// GetMe gets info of this bot.
//
// https://core.telegram.org/bots/api#getme
func (b *Bot) GetMe() (result ApiResponseUser) {
	return b.requestResponseUser("getMe", map[string]interface{}{}) // no params
}

// SendMessage sends a message to the bot.
//
// options include: parse_mode, disable_web_page_preview, disable_notification, reply_to_message_id, and reply_markup.
//
// https://core.telegram.org/bots/api#sendmessage
func (b *Bot) SendMessage(chatId ChatId, text string, options map[string]interface{}) (result ApiResponseMessage) {
	// essential params
	params := map[string]interface{}{
		"chat_id": chatId,
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
func (b *Bot) ForwardMessage(chatId, fromChatId ChatId, messageId int) (result ApiResponseMessage) {
	// essential params
	params := map[string]interface{}{
		"chat_id":      chatId,
		"from_chat_id": fromChatId,
		"message_id":   messageId,
	}

	return b.requestResponseMessage("forwardMessage", params)
}

// SendPhoto sends a photo.
//
// options include: caption, parse_mode, disable_notification, reply_to_message_id, and reply_markup.
//
// https://core.telegram.org/bots/api#sendphoto
func (b *Bot) SendPhoto(chatId ChatId, photo InputFile, options map[string]interface{}) (result ApiResponseMessage) {
	return b.sendObjectMessage(chatId, "sendPhoto", "photo", photo, options)
}

// SendAudio sends an audio file. (.mp3 format only, will be played with external players)
//
// options include: caption, parse_mode, duration, performer, title, disable_notification, reply_to_message_id, and reply_markup.
//
// https://core.telegram.org/bots/api#sendaudio
func (b *Bot) SendAudio(chatId ChatId, audio InputFile, options map[string]interface{}) (result ApiResponseMessage) {
	return b.sendObjectMessage(chatId, "sendAudio", "audio", audio, options)
}

// SendDocument sends a general file.
//
// options include: caption, parse_mode, disable_notification, reply_to_message_id, and reply_markup.
//
// https://core.telegram.org/bots/api#senddocument
func (b *Bot) SendDocument(chatId ChatId, document InputFile, options map[string]interface{}) (result ApiResponseMessage) {
	return b.sendObjectMessage(chatId, "sendDocument", "document", document, options)
}

// SendSticker sends a sticker.
//
// options include: disable_notification, reply_to_message_id, and reply_markup.
//
// https://core.telegram.org/bots/api#sendsticker
func (b *Bot) SendSticker(chatId ChatId, sticker InputFile, options map[string]interface{}) (result ApiResponseMessage) {
	return b.sendObjectMessage(chatId, "sendSticker", "sticker", sticker, options)
}

// GetStickerSet gets a sticker set.
//
// https://core.telegram.org/bots/api#getstickerset
func (b *Bot) GetStickerSet(name string) (result ApiResponseStickerSet) {
	// essential params
	params := map[string]interface{}{
		"name": name,
	}

	return b.requestResponseStickerSet("getStickerSet", params)
}

// UploadStickerFile uploads a sticker file.
//
// https://core.telegram.org/bots/api#uploadstickerfile
func (b *Bot) UploadStickerFile(userId int, sticker InputFile) (result ApiResponseFile) {
	// essential params
	params := map[string]interface{}{
		"user_id": userId,
	}

	return b.sendObjectFile("uploadStickerFile", "png_sticker", sticker, params)
}

// CreateNewStickerSet creates a new sticker set.
//
// options include: contains_masks and mask_position
//
// https://core.telegram.org/bots/api#createnewstickerset
func (b *Bot) CreateNewStickerSet(userId int, name, title string, sticker InputFile, emojis string, options map[string]interface{}) (result ApiResponse) {
	// essential params
	params := map[string]interface{}{
		"user_id": userId,
		"name":    name,
		"title":   title,
		"emojis":  emojis,
	}
	// optional params
	for key, val := range options {
		if val != nil {
			params[key] = val
		}
	}

	return b.sendObject("createNewStickerSet", "png_sticker", sticker, params)
}

// AddStickerToSet adds a sticker to set.
//
// options include: mask_position
//
// https://core.telegram.org/bots/api#addstickertoset
func (b *Bot) AddStickerToSet(userId int, name string, sticker InputFile, emojis string, options map[string]interface{}) (result ApiResponse) {
	// essential params
	params := map[string]interface{}{
		"user_id": userId,
		"name":    name,
		"emojis":  emojis,
	}
	// optional params
	for key, val := range options {
		if val != nil {
			params[key] = val
		}
	}

	return b.sendObject("addStickerToSet", "png_sticker", sticker, params)
}

// SetStickerPositionInSet sets sticker position in set.
//
// https://core.telegram.org/bots/api#setstickerpositioninset
func (b *Bot) SetStickerPositionInSet(sticker string, position int) (result ApiResponse) {
	// essential params
	params := map[string]interface{}{
		"sticker":  sticker,
		"position": position,
	}

	return b.requestResponse("setStickerPositionInSet", params)
}

// DeleteStickerFromSet deletes a sticker from set.
//
// https://core.telegram.org/bots/api#deletestickerfromset
func (b *Bot) DeleteStickerFromSet(sticker string) (result ApiResponse) {
	// essential params
	params := map[string]interface{}{
		"sticker": sticker,
	}

	return b.requestResponse("deleteStickerFromSet", params)
}

// SendVideo sends a video file.
//
// options include: duration, caption, parse_mode, supports_streaming, disable_notification, reply_to_message_id, and reply_markup.
//
// https://core.telegram.org/bots/api#sendvideo
func (b *Bot) SendVideo(chatId ChatId, video InputFile, options map[string]interface{}) (result ApiResponseMessage) {
	return b.sendObjectMessage(chatId, "sendVideo", "video", video, options)
}

// SendVoice sends a voice file. (.ogg format only, will be played with Telegram itself))
//
// options include: caption, parse_mode, duration, disable_notification, reply_to_message_id, and reply_markup.
//
// https://core.telegram.org/bots/api#sendvoice
func (b *Bot) SendVoice(chatId ChatId, voice InputFile, options map[string]interface{}) (result ApiResponseMessage) {
	return b.sendObjectMessage(chatId, "sendVoice", "voice", voice, options)
}

// SendVideoNote sends a video note.
//
// videoNote cannot be a remote http url (not supported yet)
//
// options include: duration, length, disable_notification, reply_to_message_id, and reply_markup.
// (XXX: API returns 'Bad Request: wrong video note length' when length is not given / 2017.05.19.)
//
// https://core.telegram.org/bots/api#sendvideonote
func (b *Bot) SendVideoNote(chatId ChatId, videoNote InputFile, options map[string]interface{}) (result ApiResponseMessage) {
	return b.sendObjectMessage(chatId, "sendVideoNote", "video_note", videoNote, options)
}

// SendMediaGroup sends a group of photos or videos as an album.
//
// options include: disable_notification, and reply_to_message_id
//
// https://core.telegram.org/bots/api#sendmediagroup
func (b *Bot) SendMediaGroup(chatId ChatId, media []InputMedia, options map[string]interface{}) (result ApiResponseMessage) {
	// essential params
	params := map[string]interface{}{
		"chat_id": chatId,
		"media":   media,
	}
	// optional params
	for key, val := range options {
		if val != nil {
			params[key] = val
		}
	}

	return b.requestResponseMessage("sendMediaGroup", params)
}

// SendLocation sends locations.
//
// options include: disable_notification, reply_to_message_id, and reply_markup.
//
// https://core.telegram.org/bots/api#sendlocation
func (b *Bot) SendLocation(chatId ChatId, latitude, longitude float32, options map[string]interface{}) (result ApiResponseMessage) {
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

// SendVenue sends venues.
//
// options include: foursquare_id, disable_notification, reply_to_message_id, and reply_markup.
//
// https://core.telegram.org/bots/api#sendvenue
func (b *Bot) SendVenue(chatId ChatId, latitude, longitude float32, title, address string, options map[string]interface{}) (result ApiResponseMessage) {
	// essential params
	params := map[string]interface{}{
		"chat_id":   chatId,
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
// options include: last_name, disable_notification, reply_to_message_id, and reply_markup.
//
// https://core.telegram.org/bots/api#sendcontact
func (b *Bot) SendContact(chatId ChatId, phoneNumber, firstName string, options map[string]interface{}) (result ApiResponseMessage) {
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

// SendChatAction sends chat actions.
//
// https://core.telegram.org/bots/api#sendchataction
func (b *Bot) SendChatAction(chatId ChatId, action ChatAction) (result ApiResponse) {
	// essential params
	params := map[string]interface{}{
		"chat_id": chatId,
		"action":  action,
	}

	return b.requestResponse("sendChatAction", params)
}

// GetUserProfilePhotos gets user profile photos.
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

// GetFile gets file info and prepare for download.
//
// https://core.telegram.org/bots/api#getfile
func (b *Bot) GetFile(fileId string) (result ApiResponseFile) {
	// essential params
	params := map[string]interface{}{
		"file_id": fileId,
	}

	return b.requestResponseFile("getFile", params)
}

// GetFileUrl gets download link from a given File.
func (b *Bot) GetFileUrl(file File) string {
	return fmt.Sprintf("%s%s/%s", fileBaseUrl, b.token, *file.FilePath)
}

// KickChatMember kicks a chat member
//
// https://core.telegram.org/bots/api#kickchatmember
func (b *Bot) KickChatMember(chatId ChatId, userId int) (result ApiResponse) {
	// essential params
	params := map[string]interface{}{
		"chat_id": chatId,
		"user_id": userId,
	}

	return b.requestResponse("kickChatMember", params)
}

// KickChatMemberUntil kicks a chat member until given date
func (b *Bot) KickChatMemberUntil(chatId ChatId, userId int, untilDate int) (result ApiResponse) {
	// essential params
	params := map[string]interface{}{
		"chat_id":    chatId,
		"user_id":    userId,
		"until_date": untilDate,
	}

	return b.requestResponse("kickChatMember", params)
}

// LeaveChat leaves a chat
//
// https://core.telegram.org/bots/api#leavechat
func (b *Bot) LeaveChat(chatId ChatId) (result ApiResponse) {
	// essential params
	params := map[string]interface{}{
		"chat_id": chatId,
	}

	return b.requestResponse("leaveChat", params)
}

// UnbanChatMember unbans a chat member
//
// https://core.telegram.org/bots/api#unbanchatmember
func (b *Bot) UnbanChatMember(chatId ChatId, userId int) (result ApiResponse) {
	// essential params
	params := map[string]interface{}{
		"chat_id": chatId,
		"user_id": userId,
	}

	return b.requestResponse("unbanChatMember", params)
}

// RestrictChatMember restricts a chat member
//
// options include: until_date, can_send_messages, can_send_media_messages, can_send_other_messages, and can_send_web_page_previews
//
// https://core.telegram.org/bots/api#restrictchatmember
func (b *Bot) RestrictChatMember(chatId ChatId, userId int, options map[string]interface{}) (result ApiResponse) {
	// essential params
	params := map[string]interface{}{
		"chat_id": chatId,
		"user_id": userId,
	}
	// optional params
	for key, val := range options {
		if val != nil {
			params[key] = val
		}
	}

	return b.requestResponse("restrictChatMember", params)
}

// PromoteChatMember promotes a chat member
//
// options include: can_change_info, can_post_messages, can_edit_messages, can_delete_messages, can_invite_users, can_restrict_members, can_pin_messages, and can_promote_members
//
// https://core.telegram.org/bots/api#promotechatmember
func (b *Bot) PromoteChatMember(chatId ChatId, userId int, options map[string]interface{}) (result ApiResponse) {
	// essential params
	params := map[string]interface{}{
		"chat_id": chatId,
		"user_id": userId,
	}
	// optional params
	for key, val := range options {
		if val != nil {
			params[key] = val
		}
	}

	return b.requestResponse("promoteChatMember", params)
}

// ExportChatInviteLink exports a chat invite link
//
// https://core.telegram.org/bots/api#exportchatinvitelink
func (b *Bot) ExportChatInviteLink(chatId ChatId) (result ApiResponseString) {
	// essential params
	params := map[string]interface{}{
		"chat_id": chatId,
	}

	return b.requestResponseString("exportChatInviteLink", params)
}

// SetChatPhoto sets a chat photo
//
// https://core.telegram.org/bots/api#setchatphoto
func (b *Bot) SetChatPhoto(chatId ChatId, photo InputFile) (result ApiResponse) {
	// essential params
	params := map[string]interface{}{
		"chat_id": chatId,
	}

	return b.sendObject("setChatPhoto", "photo", photo, params)
}

// DeleteChatPhoto deletes a chat photo
//
// https://core.telegram.org/bots/api#deletechatphoto
func (b *Bot) DeleteChatPhoto(chatId ChatId) (result ApiResponse) {
	// essential params
	params := map[string]interface{}{
		"chat_id": chatId,
	}

	return b.requestResponse("deleteChatPhoto", params)
}

// SetChatTitle sets a chat title
//
// https://core.telegram.org/bots/api#setchattitle
func (b *Bot) SetChatTitle(chatId ChatId, title string) (result ApiResponse) {
	// essential params
	params := map[string]interface{}{
		"chat_id": chatId,
		"title":   title,
	}

	return b.requestResponse("setChatTitle", params)
}

// SetChatDescription sets a chat description
//
// https://core.telegram.org/bots/api#setchatdescription
func (b *Bot) SetChatDescription(chatId ChatId, description string) (result ApiResponse) {
	// essential params
	params := map[string]interface{}{
		"chat_id":     chatId,
		"description": description,
	}

	return b.requestResponse("setChatDescription", params)
}

// PinChatMessage pins a chat message
//
// options include: disable_notification
//
// https://core.telegram.org/bots/api#pinchatmessage
func (b *Bot) PinChatMessage(chatId ChatId, messageId int, options map[string]interface{}) (result ApiResponse) {
	// essential params
	params := map[string]interface{}{
		"chat_id":    chatId,
		"message_id": messageId,
	}
	// optional params
	for key, val := range options {
		if val != nil {
			params[key] = val
		}
	}

	return b.requestResponse("pinChatMessage", params)
}

// UnpinChatMessage unpins a chat message
//
// https://core.telegram.org/bots/api#unpinchatmessage
func (b *Bot) UnpinChatMessage(chatId ChatId) (result ApiResponse) {
	// essential params
	params := map[string]interface{}{
		"chat_id": chatId,
	}

	return b.requestResponse("unpinChatMessage", params)
}

// GetChat gets a chat
//
// https://core.telegram.org/bots/api#getchat
func (b *Bot) GetChat(chatId ChatId) (result ApiResponseChat) {
	// essential params
	params := map[string]interface{}{
		"chat_id": chatId,
	}

	return b.requestResponseChat("getChat", params)
}

// GetChatAdministrators gets chat administrators
//
// https://core.telegram.org/bots/api#getchatadministrators
func (b *Bot) GetChatAdministrators(chatId ChatId) (result ApiResponseChatAdministrators) {
	// essential params
	params := map[string]interface{}{
		"chat_id": chatId,
	}

	return b.requestResponseChatAdministrators("getChatAdministrators", params)
}

// GetChatMembersCount gets chat members' count
//
// https://core.telegram.org/bots/api#getchatmemberscount
func (b *Bot) GetChatMembersCount(chatId ChatId) (result ApiResponseInt) {
	// essential params
	params := map[string]interface{}{
		"chat_id": chatId,
	}

	return b.requestResponseInt("getChatMembersCount", params)
}

// GetChatMember gets a chat member
//
// https://core.telegram.org/bots/api#getchatmember
func (b *Bot) GetChatMember(chatId ChatId, userId int) (result ApiResponseChatMember) {
	// essential params
	params := map[string]interface{}{
		"chat_id": chatId,
		"user_id": userId,
	}

	return b.requestResponseChatMember("getChatMember", params)
}

// SetChatStickerSet sets a chat sticker set
//
// https://core.telegram.org/bots/api#setchatstickerset
func (b *Bot) SetChatStickerSet(chatId ChatId, stickerSetName string) (result ApiResponse) {
	// essential params
	params := map[string]interface{}{
		"chat_id":          chatId,
		"sticker_set_name": stickerSetName,
	}

	return b.requestResponse("setChatStickerSet", params)
}

// DeleteChatStickerSet deletes a chat sticker set
//
// https://core.telegram.org/bots/api#deletechatstickerset
func (b *Bot) DeleteChatStickerSet(chatId ChatId) (result ApiResponse) {
	// essential params
	params := map[string]interface{}{
		"chat_id": chatId,
	}

	return b.requestResponse("deleteChatStickerSet", params)
}

// AnswerCallbackQuery answers a callback query
//
// options include: text, show_alert, url, and cache_time
//
// https://core.telegram.org/bots/api#answercallbackquery
func (b *Bot) AnswerCallbackQuery(callbackQueryId string, options map[string]interface{}) (result ApiResponse) {
	// essential params
	params := map[string]interface{}{
		"callback_query_id": callbackQueryId,
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

// EditMessageText edits text of a message
//
// required options: chat_id + message_id (when inline_message_id is not given)
//                or inline_message_id (when chat_id & message_id is not given)
//
// other options: parse_mode, disable_web_page_preview, and reply_markup
//
// https://core.telegram.org/bots/api#editmessagetext
func (b *Bot) EditMessageText(text string, options map[string]interface{}) (result ApiResponseMessage) {
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

	return b.requestResponseMessage("editMessageText", params)
}

// EditMessageCaption edits caption of a message
//
// required options: chat_id + message_id (when inline_message_id is not given)
//                or inline_message_id (when chat_id & message_id is not given)
//
// other options: parse_mode, or reply_markup
//
// https://core.telegram.org/bots/api#editmessagecaption
func (b *Bot) EditMessageCaption(caption string, options map[string]interface{}) (result ApiResponseMessage) {
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

	return b.requestResponseMessage("editMessageCaption", params)
}

// EditMessageReplyMarkup edits reply markup of a message
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

// EditMessageLiveLocation edits live location of a message
//
// required options: chat_id + message_id (when inline_message_id is not given)
//                or inline_message_id (when chat_id & message_id is not given)
//
// other options: reply_markup
//
// https://core.telegram.org/bots/api#editmessagelivelocation
func (b *Bot) EditMessageLiveLocation(latitude, longitude float32, options map[string]interface{}) (result ApiResponseMessage) {
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

	return b.requestResponseMessage("editMessageLiveLocation", params)
}

// StopMessageLiveLocation stops live location of a message
//
// required options: chat_id + message_id (when inline_message_id is not given)
//                or inline_message_id (when chat_id & message_id is not given)
//
// other options: reply_markup
//
// https://core.telegram.org/bots/api#stopmessagelivelocation
func (b *Bot) StopMessageLiveLocation(options map[string]interface{}) (result ApiResponseMessage) {
	return b.requestResponseMessage("stopMessageLiveLocation", options)
}

// DeleteMessage deletes a message
//
// https://core.telegram.org/bots/api#deletemessage
func (b *Bot) DeleteMessage(chatId ChatId, messageId int) (result ApiResponse) {
	return b.requestResponse("deleteMessage", map[string]interface{}{
		"chat_id":    chatId,
		"message_id": messageId,
	})
}

// AnswerInlineQuery sends answers to an inline query.
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

// SendInvoice sends an invoice.
//
// options include: provider_data, photo_url, photo_size, photo_width, photo_height, need_name, need_phone_number, need_email, need_shipping_address, is_flexible, disable_notification, reply_to_message_id, and reply_markup
//
// https://core.telegram.org/bots/api#sendinvoice
func (b *Bot) SendInvoice(chatId int64, title, description, payload, providerToken, startParameter, currency string, prices []LabeledPrice, options map[string]interface{}) (result ApiResponseMessage) {
	// essential params
	params := map[string]interface{}{
		"chat_id":         chatId,
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
func (b *Bot) AnswerShippingQuery(shippingQueryId string, ok bool, shippingOptions []ShippingOption, errorMessage *string) (result ApiResponse) {
	// essential params
	params := map[string]interface{}{
		"shipping_query_id": shippingQueryId,
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

	return b.requestResponse("answerShippingQuery", params)
}

// AnswerPreCheckoutQuery answers a pre-checkout query.
//
// https://core.telegram.org/bots/api#answerprecheckoutquery
func (b *Bot) AnswerPreCheckoutQuery(preCheckoutQueryId string, ok bool, errorMessage *string) (result ApiResponse) {
	// essential params
	params := map[string]interface{}{
		"pre_checkout_query_id": preCheckoutQueryId,
		"ok": ok,
	}
	// optional params
	if !ok {
		if errorMessage != nil {
			params["error_message"] = *errorMessage
		}
	}

	return b.requestResponse("answerPreCheckoutQuery", params)
}

// SendGame sends a game.
//
// options include: game_short_name, disable_notification, reply_to_message_id, and reply_markup.
//
// https://core.telegram.org/bots/api#sendgame
func (b *Bot) SendGame(chatId ChatId, gameShortName string, options map[string]interface{}) (result ApiResponseMessage) {
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

// SetGameScore sets score of a game.
//
// required options: chat_id + message_id (when inline_message_id is not given)
//                or inline_message_id (when chat_id & message_id is not given)
//
// other options: force, and disable_edit_message
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

// GetGameHighScores gets high scores of a game.
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
		case *os.File, []byte:
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
		}
		b.error("parameter '%+v' could not be cast to int value", param)
	case int64:
		if intValue, ok := param.(int64); ok {
			return strconv.FormatInt(intValue, 10), ok
		}
		b.error("parameter '%+v' could not be cast to int64 value", param)
	case float32:
		if floatValue, ok := param.(float32); ok {
			return fmt.Sprintf("%.8f", floatValue), ok
		}
		b.error("parameter '%+v' could not be cast to float32 value", param)
	case bool:
		if boolValue, ok := param.(bool); ok {
			return strconv.FormatBool(boolValue), ok
		}
		b.error("parameter '%+v' could not be cast to bool value", param)
	case string:
		if strValue, ok := param.(string); ok {
			return strValue, ok
		}
		b.error("parameter '%+v' could not be cast to string value", param)
	case ChatAction:
		if value, ok := param.(ChatAction); ok {
			return string(value), ok
		}
		b.error("parameter '%+v' could not be cast to string value", param)
	case ParseMode:
		if value, ok := param.(ParseMode); ok {
			return string(value), ok
		}
		b.error("parameter '%+v' could not be cast to string value", param)
	default:
		if json, err := json.Marshal(param); err == nil {
			return string(json), true
		} else {
			b.error("parameter '%+v' could not be encoded as json: %s", param, err)
		}
	}

	return "", false
}

// Send request to API server and return the response as bytes(synchronously).
//
// NOTE: If *os.File is included in the params, it will be closed automatically in this function.
func (b *Bot) request(method string, params map[string]interface{}) (respBytes []byte, err0 error) {
	client := &http.Client{}
	apiUrl := fmt.Sprintf("%s%s/%s", apiBaseUrl, b.token, method)

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
					if part, err := writer.CreateFormFile(key, filename); err == nil {
						if _, err = io.Copy(part, bytes.NewReader(fbytes)); err != nil {
							b.error("could not write to multipart: %s", key)
						}
					} else {
						b.error("could not create form file for parameter '%s' ([]byte)", key)
					}
				} else {
					b.error("parameter '%s' could not be cast to []byte", key)
				}
			default:
				if strValue, ok := b.paramToString(value); ok {
					writer.WriteField(key, strValue)
				}
			}
		}

		if err := writer.Close(); err != nil {
			b.error("error while closing writer (%s)", err)
		}

		if req, err := http.NewRequest("POST", apiUrl, body); err == nil {
			req.Header.Add("Content-Type", writer.FormDataContentType()) // due to file parameter

			if resp, err := client.Do(req); err == nil {
				defer resp.Body.Close()

				// FIXXX: check http status code here
				if bytes, err := ioutil.ReadAll(resp.Body); err == nil {
					return bytes, nil
				} else {
					err0 = fmt.Errorf("response read error: %s", err)

					b.error(err0.Error())
				}
			} else {
				err0 = fmt.Errorf("request error: %s", err)

				b.error(err0.Error())
			}
		} else {
			err0 = fmt.Errorf("building request error: %s", err)

			b.error(err0.Error())
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

				// FIXXX: check http status code here
				if bytes, err := ioutil.ReadAll(resp.Body); err == nil {
					return bytes, nil
				} else {
					err0 = fmt.Errorf("response read error: %s", err)

					b.error(err0.Error())
				}
			} else {
				err0 = fmt.Errorf("request error: %s", err)

				b.error(err0.Error())
			}
		} else {
			err0 = fmt.Errorf("building request error: %s", err)

			b.error(err0.Error())
		}
	}

	return []byte{}, err0
}

// Send request for ApiResponse and fetch its result.
func (b *Bot) requestResponse(method string, params map[string]interface{}) (result ApiResponse) {
	var errStr string

	if bytes, err := b.request(method, params); err == nil {
		var jsonResponse ApiResponse
		if err := json.Unmarshal(bytes, &jsonResponse); err == nil {
			return jsonResponse
		} else {
			errStr = fmt.Sprintf("json parse error: %s (%s)", err, string(bytes))
		}
	} else {
		errStr = fmt.Sprintf("%s failed with error: %s", method, err)
	}

	b.error(errStr)

	return ApiResponse{ApiResponseBase: ApiResponseBase{Ok: false, Description: &errStr}}
}

// Send request for ApiResponseWebhookInfo and fetch its result.
func (b *Bot) requestResponseWebhookInfo() (result ApiResponseWebhookInfo) {
	var errStr string

	if bytes, err := b.request("getWebhookInfo", map[string]interface{}{}); err == nil {
		var jsonResponse ApiResponseWebhookInfo
		if err := json.Unmarshal(bytes, &jsonResponse); err == nil {
			return jsonResponse
		} else {
			errStr = fmt.Sprintf("json parse error: %s (%s)", err, string(bytes))
		}
	} else {
		errStr = fmt.Sprintf("getWebhookInfo failed with error: %s", err)
	}

	b.error(errStr)

	return ApiResponseWebhookInfo{ApiResponseBase: ApiResponseBase{Ok: false, Description: &errStr}}
}

// Send request for ApiResponseUser and fetch its result.
func (b *Bot) requestResponseUser(method string, params map[string]interface{}) (result ApiResponseUser) {
	var errStr string

	if bytes, err := b.request(method, params); err == nil {
		var jsonResponse ApiResponseUser
		if err := json.Unmarshal(bytes, &jsonResponse); err == nil {
			return jsonResponse
		} else {
			errStr = fmt.Sprintf("json parse error: %s (%s)", err, string(bytes))
		}
	} else {
		errStr = fmt.Sprintf("%s failed with error: %s", method, err)
	}

	b.error(errStr)

	return ApiResponseUser{ApiResponseBase: ApiResponseBase{Ok: false, Description: &errStr}}
}

// Send request for ApiResponseMessage and fetch its result.
func (b *Bot) requestResponseMessage(method string, params map[string]interface{}) (result ApiResponseMessage) {
	var errStr string

	if bytes, err := b.request(method, params); err == nil {
		var jsonResponse ApiResponseMessage
		if err := json.Unmarshal(bytes, &jsonResponse); err == nil {
			return jsonResponse
		} else {
			errStr = fmt.Sprintf("json parse error: %s (%s)", err, string(bytes))
		}
	} else {
		errStr = fmt.Sprintf("%s failed with error: %s", method, err)
	}

	b.error(errStr)

	return ApiResponseMessage{ApiResponseBase: ApiResponseBase{Ok: false, Description: &errStr}}
}

// Send request for ApiResponseUserProfilePhotos and fetch its result.
func (b *Bot) requestResponseUserProfilePhotos(method string, params map[string]interface{}) (result ApiResponseUserProfilePhotos) {
	var errStr string

	if bytes, err := b.request(method, params); err == nil {
		var jsonResponse ApiResponseUserProfilePhotos
		if err := json.Unmarshal(bytes, &jsonResponse); err == nil {
			return jsonResponse
		} else {
			errStr = fmt.Sprintf("json parse error: %s (%s)", err, string(bytes))
		}
	} else {
		errStr = fmt.Sprintf("%s failed with error: %s", method, err)
	}

	b.error(errStr)

	return ApiResponseUserProfilePhotos{ApiResponseBase: ApiResponseBase{Ok: false, Description: &errStr}}
}

// Send request for ApiResponseUpdates and fetch its result.
func (b *Bot) requestResponseUpdates(method string, params map[string]interface{}) (result ApiResponseUpdates) {
	var errStr string

	if bytes, err := b.request(method, params); err == nil {
		var jsonResponse ApiResponseUpdates
		if err := json.Unmarshal(bytes, &jsonResponse); err == nil {
			return jsonResponse
		} else {
			errStr = fmt.Sprintf("json parse error: %s (%s)", err, string(bytes))
		}
	} else {
		errStr = fmt.Sprintf("%s failed with error: %s", method, err)
	}

	b.error(errStr)

	return ApiResponseUpdates{ApiResponseBase: ApiResponseBase{Ok: false, Description: &errStr}}
}

// Send request for ApiResponseFile and fetch its result.
func (b *Bot) requestResponseFile(method string, params map[string]interface{}) (result ApiResponseFile) {
	var errStr string

	if bytes, err := b.request(method, params); err == nil {
		var jsonResponse ApiResponseFile
		if err := json.Unmarshal(bytes, &jsonResponse); err == nil {
			return jsonResponse
		} else {
			errStr = fmt.Sprintf("json parse error: %s (%s)", err, string(bytes))
		}
	} else {
		errStr = fmt.Sprintf("%s failed with error: %s", method, err)
	}

	b.error(errStr)

	return ApiResponseFile{ApiResponseBase: ApiResponseBase{Ok: false, Description: &errStr}}
}

// Send request for ApiResponseChat and fetch its result.
func (b *Bot) requestResponseChat(method string, params map[string]interface{}) (result ApiResponseChat) {
	var errStr string

	if bytes, err := b.request(method, params); err == nil {
		var jsonResponse ApiResponseChat
		if err := json.Unmarshal(bytes, &jsonResponse); err == nil {
			return jsonResponse
		} else {
			errStr = fmt.Sprintf("json parse error: %s (%s)", err, string(bytes))
		}
	} else {
		errStr = fmt.Sprintf("%s failed with error: %s", method, err)
	}

	b.error(errStr)

	return ApiResponseChat{ApiResponseBase: ApiResponseBase{Ok: false, Description: &errStr}}
}

// Send request for ApiResponseChatAdministrator and fetch its result.
func (b *Bot) requestResponseChatAdministrators(method string, params map[string]interface{}) (result ApiResponseChatAdministrators) {
	var errStr string

	if bytes, err := b.request(method, params); err == nil {
		var jsonResponse ApiResponseChatAdministrators
		if err := json.Unmarshal(bytes, &jsonResponse); err == nil {
			return jsonResponse
		} else {
			errStr = fmt.Sprintf("json parse error: %s (%s)", err, string(bytes))
		}
	} else {
		errStr = fmt.Sprintf("%s failed with error: %s", method, err)
	}

	b.error(errStr)

	return ApiResponseChatAdministrators{ApiResponseBase: ApiResponseBase{Ok: false, Description: &errStr}}
}

// Send request for ApiResponseChatMember and fetch its result.
func (b *Bot) requestResponseChatMember(method string, params map[string]interface{}) (result ApiResponseChatMember) {
	var errStr string

	if bytes, err := b.request(method, params); err == nil {
		var jsonResponse ApiResponseChatMember
		if err := json.Unmarshal(bytes, &jsonResponse); err == nil {
			return jsonResponse
		} else {
			errStr = fmt.Sprintf("json parse error: %s (%s)", err, string(bytes))
		}
	} else {
		errStr = fmt.Sprintf("%s failed with error: %s", method, err)
	}

	b.error(errStr)

	return ApiResponseChatMember{ApiResponseBase: ApiResponseBase{Ok: false, Description: &errStr}}
}

// Send request for ApiResponseInt and fetch its result.
func (b *Bot) requestResponseInt(method string, params map[string]interface{}) (result ApiResponseInt) {
	var errStr string

	if bytes, err := b.request(method, params); err == nil {
		var jsonResponse ApiResponseInt
		if err := json.Unmarshal(bytes, &jsonResponse); err == nil {
			return jsonResponse
		} else {
			errStr = fmt.Sprintf("json parse error: %s (%s)", err, string(bytes))
		}
	} else {
		errStr = fmt.Sprintf("%s failed with error: %s", method, err)
	}

	b.error(errStr)

	return ApiResponseInt{ApiResponseBase: ApiResponseBase{Ok: false, Description: &errStr}}
}

// Send request for ApiResponseString and fetch its result.
func (b *Bot) requestResponseString(method string, params map[string]interface{}) (result ApiResponseString) {
	var errStr string

	if bytes, err := b.request(method, params); err == nil {
		var jsonResponse ApiResponseString
		if err := json.Unmarshal(bytes, &jsonResponse); err == nil {
			return jsonResponse
		} else {
			errStr = fmt.Sprintf("json parse error: %s (%s)", err, string(bytes))
		}
	} else {
		errStr = fmt.Sprintf("%s failed with error: %s", method, err)
	}

	b.error(errStr)

	return ApiResponseString{ApiResponseBase: ApiResponseBase{Ok: false, Description: &errStr}}
}

// Send request for ApiResponseGameHighScores and fetch its result.
func (b *Bot) requestResponseGameHighScores(method string, params map[string]interface{}) (result ApiResponseGameHighScores) {
	var errStr string

	if bytes, err := b.request(method, params); err == nil {
		var jsonResponse ApiResponseGameHighScores
		if err := json.Unmarshal(bytes, &jsonResponse); err == nil {
			return jsonResponse
		} else {
			errStr = fmt.Sprintf("json parse error: %s (%s)", err, string(bytes))
		}
	} else {
		errStr = fmt.Sprintf("%s failed with error: %s", method, err)
	}

	b.error(errStr)

	return ApiResponseGameHighScores{ApiResponseBase: ApiResponseBase{Ok: false, Description: &errStr}}
}

// Send request for ApiResponseStickerSet and fetch its result.
func (b *Bot) requestResponseStickerSet(method string, params map[string]interface{}) (result ApiResponseStickerSet) {
	var errStr string

	if bytes, err := b.request(method, params); err == nil {
		var jsonResponse ApiResponseStickerSet
		if err := json.Unmarshal(bytes, &jsonResponse); err == nil {
			return jsonResponse
		} else {
			errStr = fmt.Sprintf("json parse error: %s (%s)", err, string(bytes))
		}
	} else {
		errStr = fmt.Sprintf("%s failed with error: %s", method, err)
	}

	b.error(errStr)

	return ApiResponseStickerSet{ApiResponseBase: ApiResponseBase{Ok: false, Description: &errStr}}
}

// Handle Webhook request.
func (b *Bot) handleWebhook(writer http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()

	b.verbose("received webhook request: %+v", req)

	if body, err := ioutil.ReadAll(req.Body); err == nil {
		var webhook Update
		if err := json.Unmarshal(body, &webhook); err != nil {
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

// Send file
func (b *Bot) sendFile(chatId ChatId, apiName, paramKey, filepath string, options map[string]interface{}) (result ApiResponseMessage) {
	// essential params
	params := map[string]interface{}{
		"chat_id": chatId,
	}
	if isHttpUrl(filepath) {
		params[paramKey] = filepath
	} else {
		if file, err := os.Open(filepath); err == nil {
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

// Send bytes
func (b *Bot) sendBytes(chatId ChatId, apiName, paramKey string, bytes []byte, options map[string]interface{}) (result ApiResponseMessage) {
	// essential params
	params := map[string]interface{}{
		"chat_id": chatId,
		paramKey:  bytes,
	}
	// optional params
	for key, val := range options {
		if val != nil {
			params[key] = val
		}
	}

	return b.requestResponseMessage(apiName, params)
}

// Send file id (which is already uploaded to Telegram server)
func (b *Bot) sendFileId(chatId ChatId, apiName, paramKey, fileId string, options map[string]interface{}) (result ApiResponseMessage) {
	// essential params
	params := map[string]interface{}{
		"chat_id": chatId,
		paramKey:  fileId,
	}
	// optional params
	for key, val := range options {
		if val != nil {
			params[key] = val
		}
	}

	return b.requestResponseMessage(apiName, params)
}

// Send object (in []byte, filepath, http url, or file id) and return as ApiResponse
func (b *Bot) sendObject(apiName, paramKey string, obj InputFile, options map[string]interface{}) (result ApiResponse) {
	// params
	params := options

	if obj.Filepath != nil { // filepath
		if file, err := os.Open(*obj.Filepath); err == nil {
			params[paramKey] = file
		} else {
			errStr := err.Error()
			return ApiResponse{
				ApiResponseBase: ApiResponseBase{
					Ok:          false,
					Description: &errStr,
				},
			}
		}
	} else if obj.Url != nil { // http url
		params[paramKey] = *obj.Url
	} else if obj.FileId != nil { // file id
		params[paramKey] = *obj.FileId
	} else if len(obj.Bytes) > 0 { // []byte
		params[paramKey] = obj.Bytes
	} else {
		errorMessage := fmt.Sprintf("sendObject - failed to process parameter '%s': %v", paramKey, obj)
		return ApiResponse{
			ApiResponseBase: ApiResponseBase{
				Ok:          false,
				Description: &errorMessage,
			},
		}
	}

	return b.requestResponse(apiName, params)
}

// Send object (in []byte, filepath, http url, or file id) and return as ApiResponseMessage
func (b *Bot) sendObjectMessage(chatId ChatId, apiName, paramKey string, obj InputFile, options map[string]interface{}) (result ApiResponseMessage) {
	if len(obj.Bytes) > 0 {
		return b.sendBytes(chatId, apiName, paramKey, obj.Bytes, options)
	} else if obj.Filepath != nil {
		return b.sendFile(chatId, apiName, paramKey, *obj.Filepath, options)
	} else if obj.Url != nil {
		return b.sendFile(chatId, apiName, paramKey, *obj.Url, options)
	} else if obj.FileId != nil {
		return b.sendFileId(chatId, apiName, paramKey, *obj.FileId, options)
	}

	errorMessage := fmt.Sprintf("sendObjectMessage - failed to process parameter '%s': %v", paramKey, obj)
	return ApiResponseMessage{
		ApiResponseBase: ApiResponseBase{
			Ok:          false,
			Description: &errorMessage,
		},
	}
}

// Send object (in []byte, filepath, http url, or file id) and return as ApiResponseFile
func (b *Bot) sendObjectFile(apiName, paramKey string, obj InputFile, options map[string]interface{}) (result ApiResponseFile) {
	params := options

	if len(obj.Bytes) > 0 {
		params[paramKey] = obj.Bytes
	} else if obj.Filepath != nil {
		if file, err := os.Open(*obj.Filepath); err == nil {
			params[paramKey] = file
		} else {
			errStr := err.Error()
			return ApiResponseFile{
				ApiResponseBase: ApiResponseBase{
					Ok:          false,
					Description: &errStr,
				},
			}
		}
	} else if obj.Url != nil {
		params[paramKey] = *obj.Url
	} else if obj.FileId != nil {
		params[paramKey] = *obj.FileId
	} else {
		errorMessage := fmt.Sprintf("sendObjectFile - failed to process parameter '%s': %v", paramKey, obj)
		return ApiResponseFile{
			ApiResponseBase: ApiResponseBase{
				Ok:          false,
				Description: &errorMessage,
			},
		}
	}

	return b.requestResponseFile(apiName, params)
}

// check if given path is http url
func isHttpUrl(path string) bool {
	if strings.HasPrefix(path, "http://") || strings.HasPrefix(path, "https://") {
		return true
	}
	return false
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
