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
// https://core.telegram.org/bots/api#getupdates
func (b *Bot) GetUpdates(options OptionsGetUpdates) (result APIResponseUpdates) {
	if options == nil {
		options = map[string]any{}
	}

	return b.requestResponseUpdates("getUpdates", options)
}

// SetWebhook sets various options for receiving incoming updates.
//
// `port` should be one of: 443, 80, 88, or 8443.
//
// https://core.telegram.org/bots/api#setwebhook
func (b *Bot) SetWebhook(host string, port int, options OptionsSetWebhook) (result APIResponseBool) {
	b.webhookHost = host
	b.webhookPort = port
	b.webhookURL = b.getWebhookURL()

	params := map[string]any{
		"url": b.webhookURL,
	}

	if cert, exists := options["certificate"]; exists {
		var errStr string

		if filepath, ok := cert.(string); ok {
			if file, err := os.Open(filepath); err == nil {
				params["certificate"] = file
			} else {
				errStr = fmt.Sprintf("failed to open certificate: %s", err)
			}
		} else {
			errStr = "given filepath of certificate is not a string"
		}

		if errStr != "" {
			return APIResponseBool{
				APIResponseBase: APIResponseBase{
					Ok:          false,
					Description: &errStr,
				},
			}
		}
	}

	if ipAddress, exists := options["ip_address"]; exists {
		params["ip_address"] = ipAddress
	}

	if maxConnections, exists := options["max_connections"]; exists {
		params["max_connections"] = maxConnections
	}

	if allowedUpdates, exists := options["allowed_updates"]; exists {
		params["allowed_updates"] = allowedUpdates
	}

	if dropPendingUpdates, exists := options["drop_pending_updates"]; exists {
		params["drop_pending_updates"] = dropPendingUpdates
	}

	b.verbose("setting webhook url to: %s", b.webhookURL)

	return b.requestResponseBool("setWebhook", params)
}

// DeleteWebhook deletes webhook for this bot.
// (Function GetUpdates will not work if webhook is set, so in that case you'll need to delete it)
//
// https://core.telegram.org/bots/api#deletewebhook
func (b *Bot) DeleteWebhook(dropPendingUpdates bool) (result APIResponseBool) {
	b.webhookHost = ""
	b.webhookPort = 0
	b.webhookURL = ""

	b.verbose("deleting webhook url")

	return b.requestResponseBool("deleteWebhook", map[string]any{
		"drop_pending_updates": dropPendingUpdates,
	})
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
	return b.requestResponseUser("getMe", map[string]any{}) // no params
}

// LogOut logs this bot from cloud Bot API server.
//
// https://core.telegram.org/bots/api#logout
func (b *Bot) LogOut() (result APIResponseBool) {
	return b.requestResponseBool("logOut", map[string]any{}) // no params
}

// Close closes this bot from local Bot API server.
//
// https://core.telegram.org/bots/api#close
func (b *Bot) Close() (result APIResponseBool) {
	return b.requestResponseBool("close", map[string]any{}) // no params
}

// SendMessage sends a message to the bot.
//
// https://core.telegram.org/bots/api#sendmessage
func (b *Bot) SendMessage(chatID ChatID, text string, options OptionsSendMessage) (result APIResponseMessage) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["chat_id"] = chatID
	options["text"] = text

	return b.requestResponseMessage("sendMessage", options)
}

// ForwardMessage forwards a message.
//
// https://core.telegram.org/bots/api#forwardmessage
func (b *Bot) ForwardMessage(chatID, fromChatID ChatID, messageID int64, options OptionsForwardMessage) (result APIResponseMessage) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["chat_id"] = chatID
	options["from_chat_id"] = fromChatID
	options["message_id"] = messageID

	return b.requestResponseMessage("forwardMessage", options)
}

// CopyMessage copies a message.
//
// https://core.telegram.org/bots/api#copymessage
func (b *Bot) CopyMessage(chatID, fromChatID ChatID, messageID int64, options OptionsCopyMessage) (result APIResponseMessageID) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["chat_id"] = chatID
	options["from_chat_id"] = fromChatID
	options["message_id"] = messageID

	return b.requestResponseMessageID("copyMessage", options)
}

// SendPhoto sends a photo.
//
// https://core.telegram.org/bots/api#sendphoto
func (b *Bot) SendPhoto(chatID ChatID, photo InputFile, options OptionsSendPhoto) (result APIResponseMessage) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["chat_id"] = chatID
	options["photo"] = photo

	return b.requestResponseMessage("sendPhoto", options)
}

// SendAudio sends an audio file. (.mp3 format only, will be played with external players)
//
// https://core.telegram.org/bots/api#sendaudio
func (b *Bot) SendAudio(chatID ChatID, audio InputFile, options OptionsSendAudio) (result APIResponseMessage) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["chat_id"] = chatID
	options["audio"] = audio

	return b.requestResponseMessage("sendAudio", options)
}

// SendDocument sends a general file.
//
// https://core.telegram.org/bots/api#senddocument
func (b *Bot) SendDocument(chatID ChatID, document InputFile, options OptionsSendDocument) (result APIResponseMessage) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["chat_id"] = chatID
	options["document"] = document

	return b.requestResponseMessage("sendDocument", options)
}

// SendSticker sends a sticker.
//
// https://core.telegram.org/bots/api#sendsticker
func (b *Bot) SendSticker(chatID ChatID, sticker InputFile, options OptionsSendSticker) (result APIResponseMessage) {
	if options == nil {
		options = map[string]any{}
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
	params := map[string]any{
		"name": name,
	}

	return b.requestResponseStickerSet("getStickerSet", params)
}

// UploadStickerFile uploads a sticker file.
//
// https://core.telegram.org/bots/api#uploadstickerfile
func (b *Bot) UploadStickerFile(userID int64, sticker InputFile) (result APIResponseFile) {
	// essential params
	params := map[string]any{
		"user_id":     userID,
		"png_sticker": sticker,
	}

	return b.requestResponseFile("uploadStickerFile", params)
}

// CreateNewStickerSet creates a new sticker set.
//
// https://core.telegram.org/bots/api#createnewstickerset
func (b *Bot) CreateNewStickerSet(userID int64, name, title string, emojis string, options OptionsCreateNewStickerSet) (result APIResponseBool) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["user_id"] = userID
	options["name"] = name
	options["title"] = title
	options["emojis"] = emojis

	return b.requestResponseBool("createNewStickerSet", options)
}

// AddStickerToSet adds a sticker to set.
//
// https://core.telegram.org/bots/api#addstickertoset
func (b *Bot) AddStickerToSet(userID int64, name string, emojis string, options OptionsAddStickerToSet) (result APIResponseBool) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["user_id"] = userID
	options["name"] = name
	options["emojis"] = emojis

	return b.requestResponseBool("addStickerToSet", options)
}

// SetStickerPositionInSet sets sticker position in set.
//
// https://core.telegram.org/bots/api#setstickerpositioninset
func (b *Bot) SetStickerPositionInSet(sticker string, position int) (result APIResponseBool) {
	// essential params
	params := map[string]any{
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
	params := map[string]any{
		"sticker": sticker,
	}

	return b.requestResponseBool("deleteStickerFromSet", params)
}

// SetStickerSetThumb sets a thumbnail of a sticker set.
//
// https://core.telegram.org/bots/api#setstickersetthumb
func (b *Bot) SetStickerSetThumb(name string, userID int64, options OptionsSetStickerSetThumb) (result APIResponseBool) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["name"] = name
	options["user_id"] = userID

	return b.requestResponseBool("setStickerSetThumb", options)
}

// SendVideo sends a video file.
//
// https://core.telegram.org/bots/api#sendvideo
func (b *Bot) SendVideo(chatID ChatID, video InputFile, options OptionsSendVideo) (result APIResponseMessage) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["chat_id"] = chatID
	options["video"] = video

	return b.requestResponseMessage("sendVideo", options)
}

// SendAnimation sends an animation.
//
// https://core.telegram.org/bots/api#sendanimation
func (b *Bot) SendAnimation(chatID ChatID, animation InputFile, options OptionsSendAnimation) (result APIResponseMessage) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["chat_id"] = chatID
	options["animation"] = animation

	return b.requestResponseMessage("sendAnimation", options)
}

// SendVoice sends a voice file. (.ogg format only, will be played with Telegram itself))
//
// https://core.telegram.org/bots/api#sendvoice
func (b *Bot) SendVoice(chatID ChatID, voice InputFile, options OptionsSendVoice) (result APIResponseMessage) {
	if options == nil {
		options = map[string]any{}
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
// https://core.telegram.org/bots/api#sendvideonote
func (b *Bot) SendVideoNote(chatID ChatID, videoNote InputFile, options OptionsSendVideoNote) (result APIResponseMessage) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["chat_id"] = chatID
	options["video_note"] = videoNote

	return b.requestResponseMessage("sendVideoNote", options)
}

// SendMediaGroup sends a group of photos or videos as an album.
//
// https://core.telegram.org/bots/api#sendmediagroup
func (b *Bot) SendMediaGroup(chatID ChatID, media []InputMedia, options OptionsSendMediaGroup) (result APIResponseMessages) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["chat_id"] = chatID
	options["media"] = media

	return b.requestResponseMessages("sendMediaGroup", options)
}

// SendLocation sends locations.
//
// https://core.telegram.org/bots/api#sendlocation
func (b *Bot) SendLocation(chatID ChatID, latitude, longitude float32, options OptionsSendLocation) (result APIResponseMessage) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["chat_id"] = chatID
	options["latitude"] = latitude
	options["longitude"] = longitude

	return b.requestResponseMessage("sendLocation", options)
}

// SendVenue sends venues.
//
// https://core.telegram.org/bots/api#sendvenue
func (b *Bot) SendVenue(chatID ChatID, latitude, longitude float32, title, address string, options OptionsSendVenue) (result APIResponseMessage) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["chat_id"] = chatID
	options["latitude"] = latitude
	options["longitude"] = longitude
	options["title"] = title
	options["address"] = address

	return b.requestResponseMessage("sendVenue", options)
}

// SendContact sends contacts.
//
// https://core.telegram.org/bots/api#sendcontact
func (b *Bot) SendContact(chatID ChatID, phoneNumber, firstName string, options OptionsSendContact) (result APIResponseMessage) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["chat_id"] = chatID
	options["phone_number"] = phoneNumber
	options["first_name"] = firstName

	return b.requestResponseMessage("sendContact", options)
}

// SendPoll sends a poll.
//
// https://core.telegram.org/bots/api#sendpoll
func (b *Bot) SendPoll(chatID ChatID, question string, pollOptions []string, options OptionsSendPoll) (result APIResponseMessage) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["chat_id"] = chatID
	options["question"] = question
	options["options"] = pollOptions

	return b.requestResponseMessage("sendPoll", options)
}

// StopPoll stops a poll.
//
// https://core.telegram.org/bots/api#stoppoll
func (b *Bot) StopPoll(chatID ChatID, messageID int64, options OptionsStopPoll) (result APIResponsePoll) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["chat_id"] = chatID
	options["message_id"] = messageID

	return b.requestResponsePoll("stopPoll", options)
}

// SendDice sends a random dice.
//
// https://core.telegram.org/bots/api#senddice
func (b *Bot) SendDice(chatID ChatID, options OptionsSendDice) (result APIResponseMessage) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["chat_id"] = chatID

	return b.requestResponseMessage("sendDice", options)
}

// SendChatAction sends chat actions.
//
// https://core.telegram.org/bots/api#sendchataction
func (b *Bot) SendChatAction(chatID ChatID, action ChatAction) (result APIResponseBool) {
	// essential params
	params := map[string]any{
		"chat_id": chatID,
		"action":  action,
	}

	return b.requestResponseBool("sendChatAction", params)
}

// GetUserProfilePhotos gets user profile photos.
//
// https://core.telegram.org/bots/api#getuserprofilephotos
func (b *Bot) GetUserProfilePhotos(userID int64, options OptionsGetUserProfilePhotos) (result APIResponseUserProfilePhotos) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["user_id"] = userID

	return b.requestResponseUserProfilePhotos("getUserProfilePhotos", options)
}

// GetFile gets file info and prepare for download.
//
// https://core.telegram.org/bots/api#getfile
func (b *Bot) GetFile(fileID string) (result APIResponseFile) {
	// essential params
	params := map[string]any{
		"file_id": fileID,
	}

	return b.requestResponseFile("getFile", params)
}

// GetFileURL gets download link from a given File.
func (b *Bot) GetFileURL(file File) string {
	return fmt.Sprintf("%s%s/%s", fileBaseURL, b.token, *file.FilePath)
}

// BanChatMember bans a chat member.
//
// https://core.telegram.org/bots/api#banchatmember
func (b *Bot) BanChatMember(chatID ChatID, userID int64, options OptionsBanChatMember) (result APIResponseBool) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["chat_id"] = chatID
	options["user_id"] = userID

	return b.requestResponseBool("banChatMember", options)
}

// LeaveChat leaves a chat.
//
// https://core.telegram.org/bots/api#leavechat
func (b *Bot) LeaveChat(chatID ChatID) (result APIResponseBool) {
	// essential params
	params := map[string]any{
		"chat_id": chatID,
	}

	return b.requestResponseBool("leaveChat", params)
}

// UnbanChatMember unbans a chat member.
//
// https://core.telegram.org/bots/api#unbanchatmember
func (b *Bot) UnbanChatMember(chatID ChatID, userID int64, onlyIfBanned bool) (result APIResponseBool) {
	// essential params
	params := map[string]any{
		"chat_id":        chatID,
		"user_id":        userID,
		"only_if_banned": onlyIfBanned,
	}

	return b.requestResponseBool("unbanChatMember", params)
}

// RestrictChatMember restricts a chat member.
//
// https://core.telegram.org/bots/api#restrictchatmember
func (b *Bot) RestrictChatMember(chatID ChatID, userID int64, permissions ChatPermissions, options OptionsRestrictChatMember) (result APIResponseBool) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["chat_id"] = chatID
	options["user_id"] = userID
	options["permissions"] = permissions

	return b.requestResponseBool("restrictChatMember", options)
}

// PromoteChatMember promotes a chat member.
//
// https://core.telegram.org/bots/api#promotechatmember
func (b *Bot) PromoteChatMember(chatID ChatID, userID int64, options OptionsPromoteChatMember) (result APIResponseBool) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["chat_id"] = chatID
	options["user_id"] = userID

	return b.requestResponseBool("promoteChatMember", options)
}

// SetChatAdministratorCustomTitle sets chat administrator's custom title.
//
// https://core.telegram.org/bots/api#setchatadministratorcustomtitle
func (b *Bot) SetChatAdministratorCustomTitle(chatID ChatID, userID int64, customTitle string) (result APIResponseBool) {
	return b.requestResponseBool("setChatAdministratorCustomTitle", map[string]any{
		"chat_id":      chatID,
		"user_id":      userID,
		"custom_title": customTitle,
	})
}

// BanChatSenderChat bans a channel chat in a supergroup or a channel.
//
// https://core.telegram.org/bots/api#banchatsenderchat
func (b *Bot) BanChatSenderChat(chatID ChatID, senderChatID int64) (result APIResponseBool) {
	return b.requestResponseBool("banChatSenderChat", map[string]any{
		"chat_id":        chatID,
		"sender_chat_id": senderChatID,
	})
}

// UnbanChatSenderChat unbans a previously banned channel chat in a supergroup or a channel.
//
// https://core.telegram.org/bots/api#unbanchatsenderchat
func (b *Bot) UnbanChatSenderChat(chatID ChatID, senderChatID int64) (result APIResponseBool) {
	return b.requestResponseBool("unbanChatSenderChat", map[string]any{
		"chat_id":        chatID,
		"sender_chat_id": senderChatID,
	})
}

// SetChatPermissions sets permissions of a chat.
//
// https://core.telegram.org/bots/api#setchatpermissions
func (b *Bot) SetChatPermissions(chatID ChatID, permissions ChatPermissions) (result APIResponseBool) {
	return b.requestResponseBool("setChatPermissions", map[string]any{
		"chat_id":     chatID,
		"permissions": permissions,
	})
}

// ExportChatInviteLink exports a chat invite link.
//
// https://core.telegram.org/bots/api#exportchatinvitelink
func (b *Bot) ExportChatInviteLink(chatID ChatID) (result APIResponseString) {
	// essential params
	params := map[string]any{
		"chat_id": chatID,
	}

	return b.requestResponseString("exportChatInviteLink", params)
}

// CreateChatInviteLink creates a chat invite link.
//
// https://core.telegram.org/bots/api#createchatinvitelink
func (b *Bot) CreateChatInviteLink(chatID ChatID, options OptionsCreateChatInviteLink) (result APIResponseChatInviteLink) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["chat_id"] = chatID

	return b.requestResponseChatInviteLink("createChatInviteLink", options)
}

// EditChatInviteLink edits a chat invite link.
//
// https://core.telegram.org/bots/api#editchatinvitelink
func (b *Bot) EditChatInviteLink(chatID ChatID, inviteLink string, options OptionsCreateChatInviteLink) (result APIResponseChatInviteLink) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["chat_id"] = chatID
	options["invite_link"] = inviteLink

	return b.requestResponseChatInviteLink("editChatInviteLink", options)
}

// RevokeChatInviteLink revoks a chat invite link.
//
// https://core.telegram.org/bots/api#revokechatinvitelink
func (b *Bot) RevokeChatInviteLink(chatID ChatID, inviteLink string) (result APIResponseChatInviteLink) {
	return b.requestResponseChatInviteLink("revokeChatInviteLink", map[string]any{
		"chat_id":     chatID,
		"invite_link": inviteLink,
	})
}

// ApproveChatJoinRequest approves chat join request.
//
// https://core.telegram.org/bots/api#approvechatjoinrequest
func (b *Bot) ApproveChatJoinRequest(chatID ChatID, userID int64) (result APIResponseBool) {
	// essential params
	params := map[string]any{
		"chat_id": chatID,
		"user_id": userID,
	}

	return b.requestResponseBool("approveChatJoinRequest", params)
}

// DeclineChatJoinRequest declines chat join request.
//
// https://core.telegram.org/bots/api#declinechatjoinrequest
func (b *Bot) DeclineChatJoinRequest(chatID ChatID, userID int64) (result APIResponseBool) {
	// essential params
	params := map[string]any{
		"chat_id": chatID,
		"user_id": userID,
	}

	return b.requestResponseBool("declineChatJoinRequest", params)
}

// SetChatPhoto sets a chat photo.
//
// https://core.telegram.org/bots/api#setchatphoto
func (b *Bot) SetChatPhoto(chatID ChatID, photo InputFile) (result APIResponseBool) {
	// essential params
	params := map[string]any{
		"chat_id": chatID,
		"photo":   photo,
	}

	return b.requestResponseBool("setChatPhoto", params)
}

// DeleteChatPhoto deletes a chat photo.
//
// https://core.telegram.org/bots/api#deletechatphoto
func (b *Bot) DeleteChatPhoto(chatID ChatID) (result APIResponseBool) {
	// essential params
	params := map[string]any{
		"chat_id": chatID,
	}

	return b.requestResponseBool("deleteChatPhoto", params)
}

// SetChatTitle sets a chat title.
//
// https://core.telegram.org/bots/api#setchattitle
func (b *Bot) SetChatTitle(chatID ChatID, title string) (result APIResponseBool) {
	// essential params
	params := map[string]any{
		"chat_id": chatID,
		"title":   title,
	}

	return b.requestResponseBool("setChatTitle", params)
}

// SetChatDescription sets a chat description.
//
// https://core.telegram.org/bots/api#setchatdescription
func (b *Bot) SetChatDescription(chatID ChatID, description string) (result APIResponseBool) {
	// essential params
	params := map[string]any{
		"chat_id":     chatID,
		"description": description,
	}

	return b.requestResponseBool("setChatDescription", params)
}

// PinChatMessage pins a chat message.
//
// https://core.telegram.org/bots/api#pinchatmessage
func (b *Bot) PinChatMessage(chatID ChatID, messageID int64, options OptionsPinChatMessage) (result APIResponseBool) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["chat_id"] = chatID
	options["message_id"] = messageID

	return b.requestResponseBool("pinChatMessage", options)
}

// UnpinChatMessage unpins a chat message.
//
// https://core.telegram.org/bots/api#unpinchatmessage
func (b *Bot) UnpinChatMessage(chatID ChatID, options OptionsUnpinChatMessage) (result APIResponseBool) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["chat_id"] = chatID

	return b.requestResponseBool("unpinChatMessage", options)
}

// UnpinAllChatMessages unpins all chat messages.
//
// https://core.telegram.org/bots/api#unpinallchatmessages
func (b *Bot) UnpinAllChatMessages(chatID ChatID) (result APIResponseBool) {
	// essential params
	params := map[string]any{
		"chat_id": chatID,
	}

	return b.requestResponseBool("unpinAllChatMessages", params)
}

// GetChat gets a chat.
//
// https://core.telegram.org/bots/api#getchat
func (b *Bot) GetChat(chatID ChatID) (result APIResponseChat) {
	// essential params
	params := map[string]any{
		"chat_id": chatID,
	}

	return b.requestResponseChat("getChat", params)
}

// GetChatAdministrators gets chat administrators.
//
// https://core.telegram.org/bots/api#getchatadministrators
func (b *Bot) GetChatAdministrators(chatID ChatID) (result APIResponseChatAdministrators) {
	// essential params
	params := map[string]any{
		"chat_id": chatID,
	}

	return b.requestResponseChatAdministrators("getChatAdministrators", params)
}

// GetChatMemberCount gets chat members' count.
//
// https://core.telegram.org/bots/api#getchatmembercount
func (b *Bot) GetChatMemberCount(chatID ChatID) (result APIResponseInt) {
	// essential params
	params := map[string]any{
		"chat_id": chatID,
	}

	return b.requestResponseInt("getChatMemberCount", params)
}

// GetChatMember gets a chat member.
//
// https://core.telegram.org/bots/api#getchatmember
func (b *Bot) GetChatMember(chatID ChatID, userID int64) (result APIResponseChatMember) {
	// essential params
	params := map[string]any{
		"chat_id": chatID,
		"user_id": userID,
	}

	return b.requestResponseChatMember("getChatMember", params)
}

// SetChatStickerSet sets a chat sticker set.
//
// https://core.telegram.org/bots/api#setchatstickerset
func (b *Bot) SetChatStickerSet(chatID ChatID, stickerSetName string) (result APIResponseBool) {
	// essential params
	params := map[string]any{
		"chat_id":          chatID,
		"sticker_set_name": stickerSetName,
	}

	return b.requestResponseBool("setChatStickerSet", params)
}

// DeleteChatStickerSet deletes a chat sticker set.
//
// https://core.telegram.org/bots/api#deletechatstickerset
func (b *Bot) DeleteChatStickerSet(chatID ChatID) (result APIResponseBool) {
	// essential params
	params := map[string]any{
		"chat_id": chatID,
	}

	return b.requestResponseBool("deleteChatStickerSet", params)
}

// AnswerCallbackQuery answers a callback query.
//
// https://core.telegram.org/bots/api#answercallbackquery
func (b *Bot) AnswerCallbackQuery(callbackQueryID string, options OptionsAnswerCallbackQuery) (result APIResponseBool) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["callback_query_id"] = callbackQueryID

	return b.requestResponseBool("answerCallbackQuery", options)
}

// GetMyCommands fetches commands of this bot.
//
// https://core.telegram.org/bots/api#getmycommands
func (b *Bot) GetMyCommands(options OptionsGetMyCommands) (result APIResponseBotCommands) {
	return b.requestResponseBotCommands("getMyCommands", options)
}

// SetMyCommands sets commands of this bot.
//
// https://core.telegram.org/bots/api#setmycommands
func (b *Bot) SetMyCommands(commands []BotCommand, options OptionsSetMyCommands) (result APIResponseBool) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["commands"] = commands

	return b.requestResponseBool("setMyCommands", options)
}

// DeleteMyCommands deletes commands of this bot.
//
// https://core.telegram.org/bots/api#deletemycommands
func (b *Bot) DeleteMyCommands(options OptionsDeleteMyCommands) (result APIResponseBool) {
	return b.requestResponseBool("deleteMyCommands", options)
}

// SetChatMenuButton sets chat menu button.
//
// https://core.telegram.org/bots/api#setchatmenubutton
func (b *Bot) SetChatMenuButton(options OptionsSetChatMenuButton) (result APIResponseBool) {
	return b.requestResponseBool("setChatMenuButton", options)
}

// GetChatMenuButton fetches current chat menu button.
//
// https://core.telegram.org/bots/api#getchatmenubutton
func (b *Bot) GetChatMenuButton(options OptionsGetChatMenuButton) (result APIResponseMenuButton) {
	return b.requestResponseMenuButton("getChatMenuButton", options)
}

// SetMyDefaultAdministratorRights sets my default administrator rights.
//
// https://core.telegram.org/bots/api#setmydefaultadministratorrights
func (b *Bot) SetMyDefaultAdministratorRights(options OptionsSetMyDefaultAdministratorRights) (result APIResponseBool) {
	return b.requestResponseBool("setMyDefaultAdministratorRights", options)
}

// GetMyDefaultAdministratorRights gets my default administrator rights.
//
// https://core.telegram.org/bots/api#getmydefaultadministratorrights
func (b *Bot) GetMyDefaultAdministratorRights(options OptionsGetMyDefaultAdministratorRights) (result APIResponseBool) {
	return b.requestResponseBool("getMyDefaultAdministratorRights", options)
}

// Updating messages
//
// https://core.telegram.org/bots/api#updating-messages

// EditMessageText edits text of a message.
//
// https://core.telegram.org/bots/api#editmessagetext
func (b *Bot) EditMessageText(text string, options OptionsEditMessageText) (result APIResponseMessageOrBool) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["text"] = text

	return b.requestResponseMessageOrBool("editMessageText", options)
}

// EditMessageCaption edits caption of a message.
//
// https://core.telegram.org/bots/api#editmessagecaption
func (b *Bot) EditMessageCaption(caption string, options OptionsEditMessageCaption) (result APIResponseMessageOrBool) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["caption"] = caption

	return b.requestResponseMessageOrBool("editMessageCaption", options)
}

// EditMessageMedia edites a media message.
//
// https://core.telegram.org/bots/api#editmessagemedia
func (b *Bot) EditMessageMedia(media InputMedia, options OptionsEditMessageMedia) (result APIResponseMessageOrBool) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["media"] = media

	return b.requestResponseMessageOrBool("editMessageMedia", options)
}

// EditMessageReplyMarkup edits reply markup of a message.
//
// https://core.telegram.org/bots/api#editmessagereplymarkup
func (b *Bot) EditMessageReplyMarkup(options OptionsEditMessageReplyMarkup) (result APIResponseMessageOrBool) {
	return b.requestResponseMessageOrBool("editMessageReplyMarkup", options)
}

// EditMessageLiveLocation edits live location of a message.
//
// https://core.telegram.org/bots/api#editmessagelivelocation
func (b *Bot) EditMessageLiveLocation(latitude, longitude float32, options OptionsEditMessageLiveLocation) (result APIResponseMessageOrBool) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["latitude"] = latitude
	options["longitude"] = longitude

	return b.requestResponseMessageOrBool("editMessageLiveLocation", options)
}

// StopMessageLiveLocation stops live location of a message.
//
// https://core.telegram.org/bots/api#stopmessagelivelocation
func (b *Bot) StopMessageLiveLocation(options OptionsStopMessageLiveLocation) (result APIResponseMessageOrBool) {
	return b.requestResponseMessageOrBool("stopMessageLiveLocation", options)
}

// DeleteMessage deletes a message.
//
// https://core.telegram.org/bots/api#deletemessage
func (b *Bot) DeleteMessage(chatID ChatID, messageID int64) (result APIResponseBool) {
	return b.requestResponseBool("deleteMessage", map[string]any{
		"chat_id":    chatID,
		"message_id": messageID,
	})
}

// AnswerInlineQuery sends answers to an inline query.
//
// results = array of InlineQueryResultArticle, InlineQueryResultPhoto, InlineQueryResultGif, InlineQueryResultMpeg4Gif, or InlineQueryResultVideo.
//
// https://core.telegram.org/bots/api#answerinlinequery
func (b *Bot) AnswerInlineQuery(inlineQueryID string, results []any, options OptionsAnswerInlineQuery) (result APIResponseBool) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["inline_query_id"] = inlineQueryID
	options["results"] = results

	return b.requestResponseBool("answerInlineQuery", options)
}

// SendInvoice sends an invoice.
//
// https://core.telegram.org/bots/api#sendinvoice
func (b *Bot) SendInvoice(chatID int64, title, description, payload, providerToken, currency string, prices []LabeledPrice, options OptionsSendInvoice) (result APIResponseMessage) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["chat_id"] = chatID
	options["title"] = title
	options["description"] = description
	options["payload"] = payload
	options["provider_token"] = providerToken
	options["currency"] = currency
	options["prices"] = prices

	return b.requestResponseMessage("sendInvoice", options)
}

// CreateInvoiceLink creates a link for an invoice.
//
// https://core.telegram.org/bots/api#createinvoicelink
func (b *Bot) CreateInvoiceLink(title, description, payload, providerToken, currency string, prices []LabeledPrice, options OptionsCreateInvoiceLink) (result APIResponseString) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["title"] = title
	options["description"] = description
	options["payload"] = payload
	options["provider_token"] = providerToken
	options["currency"] = currency
	options["prices"] = prices

	return b.requestResponseString("createInvoiceLink", options)
}

// AnswerShippingQuery answers a shipping query.
//
// if ok is true, shippingOptions should be provided.
// otherwise, errorMessage should be provided.
//
// https://core.telegram.org/bots/api#answershippingquery
func (b *Bot) AnswerShippingQuery(shippingQueryID string, ok bool, shippingOptions []ShippingOption, errorMessage *string) (result APIResponseBool) {
	// essential params
	params := map[string]any{
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
	params := map[string]any{
		"pre_checkout_query_id": preCheckoutQueryID,
		"ok":                    ok,
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
// https://core.telegram.org/bots/api#sendgame
func (b *Bot) SendGame(chatID ChatID, gameShortName string, options OptionsSendGame) (result APIResponseMessage) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["chat_id"] = chatID
	options["game_short_name"] = gameShortName

	return b.requestResponseMessage("sendGame", options)
}

// SetGameScore sets score of a game.
//
// https://core.telegram.org/bots/api#setgamescore
func (b *Bot) SetGameScore(userID int64, score int, options OptionsSetGameScore) (result APIResponseMessageOrBool) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["user_id"] = userID
	options["score"] = score

	return b.requestResponseMessageOrBool("setGameScore", options)
}

// GetGameHighScores gets high scores of a game.
//
// https://core.telegram.org/bots/api#getgamehighscores
func (b *Bot) GetGameHighScores(userID int64, options OptionsGetGameHighScores) (result APIResponseGameHighScores) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["user_id"] = userID

	return b.requestResponseGameHighScores("getGameHighScores", options)
}

// AnswerWebAppQuery answers a web app's query
//
// https://core.telegram.org/bots/api#answerwebappquery
func (b *Bot) AnswerWebAppQuery(webAppQueryID string, res InlineQueryResult) (result APIResponseSentWebAppMessage) {
	options := map[string]any{
		"web_app_query_id": webAppQueryID,
		"result":           res,
	}

	return b.requestResponseSentWebAppMessage("answerWebAppQuery", options)
}

// Check if given http params contain file or not.
func checkIfFileParamExists(params map[string]any) bool {
	for _, value := range params {
		switch val := value.(type) {
		case *os.File, []byte:
			return true
		case InputFile:
			if len(val.Bytes) > 0 || val.Filepath != nil {
				return true
			}
		}
	}

	return false
}

// Convert given interface to string. (for HTTP params)
func (b *Bot) paramToString(param any) (result string, success bool) {
	switch val := param.(type) {
	case int:
		return strconv.Itoa(val), true
	case int64:
		return strconv.FormatInt(val, 10), true
	case float32:
		return fmt.Sprintf("%.8f", val), true
	case bool:
		return strconv.FormatBool(val), true
	case string:
		return val, true
	case ChatAction:
		return string(val), true
	case ParseMode:
		return string(val), true
	case InputFile:
		if val.URL != nil {
			return *val.URL, true
		}
		if val.FileID != nil {
			return *val.FileID, true
		}
		b.error("parameter '%+v' could not be cast to string value", param)
	default: // fallback: encode to JSON string
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
// NOTE: If *os.File is included in the params, it will be closed automatically by this function.
func (b *Bot) request(method string, params map[string]any) (resp []byte, err error) {
	apiURL := fmt.Sprintf("%s%s/%s", apiBaseURL, b.token, method)

	b.verbose("sending request to api url: %s, params: %#v", apiURL, params)

	if checkIfFileParamExists(params) {
		// multipart form data
		resp, err = b.requestMultipartFormData(apiURL, params)
	} else {
		// www-form urlencoded
		resp, err = b.requestURLEncodedFormData(apiURL, params)
	}

	if err == nil {
		return resp, nil
	}

	return []byte{}, fmt.Errorf(b.redact(err.Error()))
}

// request multipart form data
func (b *Bot) requestMultipartFormData(apiURL string, params map[string]any) (resp []byte, err error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	for key, value := range params {
		switch val := value.(type) {
		case *os.File:
			defer val.Close() // XXX - close the file

			var part io.Writer
			part, err = writer.CreateFormFile(key, val.Name())
			if err == nil {
				if _, err = io.Copy(part, val); err != nil {
					b.error("could not write to multipart: %s", key)
				}
			} else {
				b.error("could not create form file for parameter '%s' (%v)", key, value)
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
						if err := writer.WriteField(key, strValue); err != nil {
							b.error("failed to write field with key: %s, value: %s (%s)", key, strValue, err)
						}
					} else {
						b.error("invalid InputFile parameter '%s'", key)
					}
				}
			} else {
				b.error("parameter '%s' could not be cast to InputFile", key)
			}
		default:
			if strValue, ok := b.paramToString(value); ok {
				if err := writer.WriteField(key, strValue); err != nil {
					b.error("failed to write filed with key: %s, value: %s (%s)", key, strValue, err)
				}
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
		resp, err = b.httpClient.Do(req)

		if resp != nil { // XXX - in case of http redirect
			defer resp.Body.Close()
		}

		if err == nil {
			// FIXXX: check http status code here
			var bytes []byte
			bytes, err = ioutil.ReadAll(resp.Body)
			if err == nil {
				return bytes, nil
			}

			err = fmt.Errorf("response read error: %w", err)

			b.error(err.Error())
		} else {
			err = fmt.Errorf("request error: %w", err)

			b.error(err.Error())
		}
	} else {
		err = fmt.Errorf("building request error: %w", err)

		b.error(err.Error())
	}

	return []byte{}, err
}

// request urlencoded form data
func (b *Bot) requestURLEncodedFormData(apiURL string, params map[string]any) (resp []byte, err error) {
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
		resp, err = b.httpClient.Do(req)

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

			err = fmt.Errorf("response read error: %w", err)

			b.error(err.Error())
		} else {
			err = fmt.Errorf("request error: %w", err)

			b.error(err.Error())
		}
	} else {
		err = fmt.Errorf("building request error: %w", err)

		b.error(err.Error())
	}

	return []byte{}, err
}

// Send request for APIResponseWebhookInfo and fetch its result.
func (b *Bot) requestResponseWebhookInfo() (result APIResponseWebhookInfo) {
	var errStr string

	if bytes, err := b.request("getWebhookInfo", map[string]any{}); err == nil {
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
func (b *Bot) requestResponseUser(method string, params map[string]any) (result APIResponseUser) {
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
func (b *Bot) requestResponseMessage(method string, params map[string]any) (result APIResponseMessage) {
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
func (b *Bot) requestResponseMessages(method string, params map[string]any) (result APIResponseMessages) {
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

// Send request for APIResponseMessageID and fetch its result.
func (b *Bot) requestResponseMessageID(method string, params map[string]any) (result APIResponseMessageID) {
	var errStr string

	if bytes, err := b.request(method, params); err == nil {
		var jsonResponse APIResponseMessageID
		err = json.Unmarshal(bytes, &jsonResponse)
		if err == nil {
			return jsonResponse
		}

		errStr = fmt.Sprintf("json parse error: %s (%s)", err, string(bytes))
	} else {
		errStr = fmt.Sprintf("%s failed with error: %s", method, err)
	}

	b.error(errStr)

	return APIResponseMessageID{APIResponseBase: APIResponseBase{Ok: false, Description: &errStr}}
}

// Send request for APIResponseUserProfilePhotos and fetch its result.
func (b *Bot) requestResponseUserProfilePhotos(method string, params map[string]any) (result APIResponseUserProfilePhotos) {
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
func (b *Bot) requestResponseUpdates(method string, params map[string]any) (result APIResponseUpdates) {
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
func (b *Bot) requestResponseFile(method string, params map[string]any) (result APIResponseFile) {
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
func (b *Bot) requestResponseChat(method string, params map[string]any) (result APIResponseChat) {
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
func (b *Bot) requestResponseChatAdministrators(method string, params map[string]any) (result APIResponseChatAdministrators) {
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
func (b *Bot) requestResponseChatMember(method string, params map[string]any) (result APIResponseChatMember) {
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
func (b *Bot) requestResponseInt(method string, params map[string]any) (result APIResponseInt) {
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
func (b *Bot) requestResponseBool(method string, params map[string]any) (result APIResponseBool) {
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
func (b *Bot) requestResponseString(method string, params map[string]any) (result APIResponseString) {
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
func (b *Bot) requestResponseGameHighScores(method string, params map[string]any) (result APIResponseGameHighScores) {
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

// Send request for APIResponseSentWebAppMessage and fetch its result.
func (b *Bot) requestResponseSentWebAppMessage(method string, params map[string]any) (result APIResponseSentWebAppMessage) {
	var errStr string

	if bytes, err := b.request(method, params); err == nil {
		var jsonResponse APIResponseSentWebAppMessage
		err = json.Unmarshal(bytes, &jsonResponse)
		if err == nil {
			return jsonResponse
		}

		errStr = fmt.Sprintf("json parse error: %s (%s)", err, string(bytes))
	} else {
		errStr = fmt.Sprintf("%s failed with error: %s", method, err)
	}

	b.error(errStr)

	return APIResponseSentWebAppMessage{APIResponseBase: APIResponseBase{Ok: false, Description: &errStr}}
}

// Send request for APIResponseStickerSet and fetch its result.
func (b *Bot) requestResponseStickerSet(method string, params map[string]any) (result APIResponseStickerSet) {
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
func (b *Bot) requestResponseMessageOrBool(method string, params map[string]any) (result APIResponseMessageOrBool) {
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

// Send request for APIResponsePoll and fetch its result.
func (b *Bot) requestResponsePoll(method string, params map[string]any) (result APIResponsePoll) {
	var errStr string

	if bytes, err := b.request(method, params); err == nil {
		var jsonResponse APIResponsePoll
		err = json.Unmarshal(bytes, &jsonResponse)
		if err == nil {
			return jsonResponse
		}

		errStr = fmt.Sprintf("json parse error: %s (%s)", err, string(bytes))
	} else {
		errStr = fmt.Sprintf("%s failed with error: %s", method, err)
	}

	b.error(errStr)

	return APIResponsePoll{APIResponseBase: APIResponseBase{Ok: false, Description: &errStr}}
}

// Send request for APIResponseBotCommands and fetch its result.
func (b *Bot) requestResponseBotCommands(method string, params map[string]any) (result APIResponseBotCommands) {
	var errStr string

	if bytes, err := b.request(method, params); err == nil {
		var jsonResponse APIResponseBotCommands
		err = json.Unmarshal(bytes, &jsonResponse)
		if err == nil {
			return jsonResponse
		}

		errStr = fmt.Sprintf("json parse error: %s (%s)", err, string(bytes))
	} else {
		errStr = fmt.Sprintf("%s failed with error: %s", method, err)
	}

	b.error(errStr)

	return APIResponseBotCommands{APIResponseBase: APIResponseBase{Ok: false, Description: &errStr}}
}

// Send request for APIResponseChatInviteLink and fetch its result.
func (b *Bot) requestResponseChatInviteLink(method string, params map[string]any) (result APIResponseChatInviteLink) {
	var errStr string

	if bytes, err := b.request(method, params); err == nil {
		var jsonResponse APIResponseChatInviteLink
		err = json.Unmarshal(bytes, &jsonResponse)
		if err == nil {
			return jsonResponse
		}

		errStr = fmt.Sprintf("json parse error: %s (%s)", err, string(bytes))
	} else {
		errStr = fmt.Sprintf("%s failed with error: %s", method, err)
	}

	b.error(errStr)

	return APIResponseChatInviteLink{APIResponseBase: APIResponseBase{Ok: false, Description: &errStr}}
}

// Send request for APIResponseMenuButton and fetch its result.
func (b *Bot) requestResponseMenuButton(method string, params map[string]any) (result APIResponseMenuButton) {
	var errStr string

	if bytes, err := b.request(method, params); err == nil {
		var jsonResponse APIResponseMenuButton
		err = json.Unmarshal(bytes, &jsonResponse)
		if err == nil {
			return jsonResponse
		}

		errStr = fmt.Sprintf("json parse error: %s (%s)", err, string(bytes))
	} else {
		errStr = fmt.Sprintf("%s failed with error: %s", method, err)
	}

	b.error(errStr)

	return APIResponseMenuButton{APIResponseBase: APIResponseBase{Ok: false, Description: &errStr}}
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
