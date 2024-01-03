package telegrambot

// https://core.telegram.org/bots/api#available-methods

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
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
func (b *Bot) GetUpdates(options OptionsGetUpdates) (result APIResponse[[]Update]) {
	if options == nil {
		options = map[string]any{}
	}

	return requestGeneric[[]Update](b, "getUpdates", options)
}

// SetWebhook sets various options for receiving incoming updates.
//
// `port` should be one of: 443, 80, 88, or 8443.
//
// https://core.telegram.org/bots/api#setwebhook
func (b *Bot) SetWebhook(host string, port int, options OptionsSetWebhook) (result APIResponse[bool]) {
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
			return APIResponse[bool]{
				Ok:          false,
				Description: &errStr,
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

	return requestGeneric[bool](b, "setWebhook", params)
}

// DeleteWebhook deletes webhook for this bot.
// (Function GetUpdates will not work if webhook is set, so in that case you'll need to delete it)
//
// https://core.telegram.org/bots/api#deletewebhook
func (b *Bot) DeleteWebhook(dropPendingUpdates bool) (result APIResponse[bool]) {
	b.webhookHost = ""
	b.webhookPort = 0
	b.webhookURL = ""

	b.verbose("deleting webhook url")

	return requestGeneric[bool](b, "deleteWebhook", map[string]any{
		"drop_pending_updates": dropPendingUpdates,
	})
}

// GetWebhookInfo gets webhook info for this bot.
//
// https://core.telegram.org/bots/api#getwebhookinfo
func (b *Bot) GetWebhookInfo() (result APIResponse[WebhookInfo]) {
	return requestGeneric[WebhookInfo](b, "getWebhookInfo", map[string]any{})
}

// GetMe gets info of this bot.
//
// https://core.telegram.org/bots/api#getme
func (b *Bot) GetMe() (result APIResponse[User]) {
	return requestGeneric[User](b, "getMe", map[string]any{}) // no params
}

// LogOut logs this bot from cloud Bot API server.
//
// https://core.telegram.org/bots/api#logout
func (b *Bot) LogOut() (result APIResponse[bool]) {
	return requestGeneric[bool](b, "logOut", map[string]any{}) // no params
}

// Close closes this bot from local Bot API server.
//
// https://core.telegram.org/bots/api#close
func (b *Bot) Close() (result APIResponse[bool]) {
	return requestGeneric[bool](b, "close", map[string]any{}) // no params
}

// SendMessage sends a message to the bot.
//
// https://core.telegram.org/bots/api#sendmessage
func (b *Bot) SendMessage(chatID ChatID, text string, options OptionsSendMessage) (result APIResponse[Message]) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["chat_id"] = chatID
	options["text"] = text

	return requestGeneric[Message](b, "sendMessage", options)
}

// ForwardMessage forwards a message.
//
// https://core.telegram.org/bots/api#forwardmessage
func (b *Bot) ForwardMessage(chatID, fromChatID ChatID, messageID int64, options OptionsForwardMessage) (result APIResponse[Message]) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["chat_id"] = chatID
	options["from_chat_id"] = fromChatID
	options["message_id"] = messageID

	return requestGeneric[Message](b, "forwardMessage", options)
}

// ForwardMessages forwards messages.
//
// https://core.telegram.org/bots/api#forwardmessages
func (b *Bot) ForwardMessages(chatID, fromChatID ChatID, messageIDs []int64, options OptionsForwardMessage) (result APIResponse[[]MessageID]) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["chat_id"] = chatID
	options["from_chat_id"] = fromChatID
	options["message_ids"] = messageIDs

	return requestGeneric[[]MessageID](b, "forwardMessages", options)
}

// CopyMessage copies a message.
//
// https://core.telegram.org/bots/api#copymessage
func (b *Bot) CopyMessage(chatID, fromChatID ChatID, messageID int64, options OptionsCopyMessage) (result APIResponse[MessageID]) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["chat_id"] = chatID
	options["from_chat_id"] = fromChatID
	options["message_id"] = messageID

	return requestGeneric[MessageID](b, "copyMessage", options)
}

// CopyMessages copies messages.
//
// https://core.telegram.org/bots/api#copymessages
func (b *Bot) CopyMessages(chatID, fromChatID ChatID, messageIDs []int64, options OptionsCopyMessages) (result APIResponse[[]MessageID]) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["chat_id"] = chatID
	options["from_chat_id"] = fromChatID
	options["message_ids"] = messageIDs

	return requestGeneric[[]MessageID](b, "copyMessages", options)
}

// SendPhoto sends a photo.
//
// https://core.telegram.org/bots/api#sendphoto
func (b *Bot) SendPhoto(chatID ChatID, photo InputFile, options OptionsSendPhoto) (result APIResponse[Message]) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["chat_id"] = chatID
	options["photo"] = photo

	return requestGeneric[Message](b, "sendPhoto", options)
}

// SendAudio sends an audio file. (.mp3 format only, will be played with external players)
//
// https://core.telegram.org/bots/api#sendaudio
func (b *Bot) SendAudio(chatID ChatID, audio InputFile, options OptionsSendAudio) (result APIResponse[Message]) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["chat_id"] = chatID
	options["audio"] = audio

	return requestGeneric[Message](b, "sendAudio", options)
}

// SendDocument sends a general file.
//
// https://core.telegram.org/bots/api#senddocument
func (b *Bot) SendDocument(chatID ChatID, document InputFile, options OptionsSendDocument) (result APIResponse[Message]) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["chat_id"] = chatID
	options["document"] = document

	return requestGeneric[Message](b, "sendDocument", options)
}

// SendSticker sends a sticker.
//
// https://core.telegram.org/bots/api#sendsticker
func (b *Bot) SendSticker(chatID ChatID, sticker InputFile, options OptionsSendSticker) (result APIResponse[Message]) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["chat_id"] = chatID
	options["sticker"] = sticker

	return requestGeneric[Message](b, "sendSticker", options)
}

// GetStickerSet gets a sticker set.
//
// https://core.telegram.org/bots/api#getstickerset
func (b *Bot) GetStickerSet(name string) (result APIResponse[StickerSet]) {
	// essential params
	params := map[string]any{
		"name": name,
	}

	return requestGeneric[StickerSet](b, "getStickerSet", params)
}

// GetCustomEmojiStickers gets custom emoji stickers.
//
// https://core.telegram.org/bots/api#getcustomemojistickers
func (b *Bot) GetCustomEmojiStickers(customEmojiIDs []string) (result APIResponse[[]Sticker]) {
	// essential params
	params := map[string]any{
		"custom_emoji_ids": customEmojiIDs,
	}

	return requestGeneric[[]Sticker](b, "getCustomEmojiStickers", params)
}

// UploadStickerFile uploads a sticker file.
//
// https://core.telegram.org/bots/api#uploadstickerfile
func (b *Bot) UploadStickerFile(userID int64, sticker InputFile, stickerFormat StickerFormat) (result APIResponse[File]) {
	// essential params
	params := map[string]any{
		"user_id":        userID,
		"sticker":        sticker,
		"sticker_format": stickerFormat,
	}

	return requestGeneric[File](b, "uploadStickerFile", params)
}

// CreateNewStickerSet creates a new sticker set.
//
// https://core.telegram.org/bots/api#createnewstickerset
func (b *Bot) CreateNewStickerSet(userID int64, name, title string, stickers []InputSticker, stickerFormat StickerFormat, options OptionsCreateNewStickerSet) (result APIResponse[bool]) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["user_id"] = userID
	options["name"] = name
	options["title"] = title
	options["stickers"] = stickers
	options["sticker_format"] = stickerFormat

	return requestGeneric[bool](b, "createNewStickerSet", options)
}

// AddStickerToSet adds a sticker to set.
//
// https://core.telegram.org/bots/api#addstickertoset
func (b *Bot) AddStickerToSet(userID int64, name string, sticker InputSticker, options OptionsAddStickerToSet) (result APIResponse[bool]) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["user_id"] = userID
	options["name"] = name
	options["sticker"] = sticker

	return requestGeneric[bool](b, "addStickerToSet", options)
}

// SetStickerPositionInSet sets sticker position in set.
//
// https://core.telegram.org/bots/api#setstickerpositioninset
func (b *Bot) SetStickerPositionInSet(sticker string, position int) (result APIResponse[bool]) {
	// essential params
	params := map[string]any{
		"sticker":  sticker,
		"position": position,
	}

	return requestGeneric[bool](b, "setStickerPositionInSet", params)
}

// DeleteStickerFromSet deletes a sticker from set.
//
// https://core.telegram.org/bots/api#deletestickerfromset
func (b *Bot) DeleteStickerFromSet(sticker string) (result APIResponse[bool]) {
	// essential params
	params := map[string]any{
		"sticker": sticker,
	}

	return requestGeneric[bool](b, "deleteStickerFromSet", params)
}

// SetStickerSetThumbnail sets a thumbnail of a sticker set.
//
// https://core.telegram.org/bots/api#setstickersetthumbnail
func (b *Bot) SetStickerSetThumbnail(name string, userID int64, options OptionsSetStickerSetThumbnail) (result APIResponse[bool]) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["name"] = name
	options["user_id"] = userID

	return requestGeneric[bool](b, "setStickerSetThumbnail", options)
}

// SetCustomEmojiStickerSetThumbnail sets the custom emoji sticker set's thumbnail.
//
// https://core.telegram.org/bots/api#setcustomemojistickersetthumbnail
func (b *Bot) SetCustomEmojiStickerSetThumbnail(name string, options OptionsSetCustomEmojiStickerSetThumbnail) (result APIResponse[bool]) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["name"] = name

	return requestGeneric[bool](b, "setCustomEmojiStickerSetThumbnail", options)
}

// SetStickerSetTitle sets the title of sticker set.
//
// https://core.telegram.org/bots/api#setstickersettitle
func (b *Bot) SetStickerSetTitle(name, title string) (result APIResponse[bool]) {
	return requestGeneric[bool](b, "setStickerSetTitle", map[string]any{
		"name":  name,
		"title": title,
	})
}

// DeleteStickerSet deletes a sticker set.
//
// https://core.telegram.org/bots/api#deletestickerset
func (b *Bot) DeleteStickerSet(name string) (result APIResponse[bool]) {
	return requestGeneric[bool](b, "deleteStickerSet", map[string]any{
		"name": name,
	})
}

// SetStickerEmojiList sets the emoji list of sticker set.
//
// https://core.telegram.org/bots/api#setstickeremojilist
func (b *Bot) SetStickerEmojiList(sticker string, emojiList []string) (result APIResponse[bool]) {
	return requestGeneric[bool](b, "setStickerEmojiList", map[string]any{
		"sticker":    sticker,
		"emoji_list": emojiList,
	})
}

// SetStickerKeywords sets the keywords of sticker.
//
// https://core.telegram.org/bots/api#setstickerkeywords
func (b *Bot) SetStickerKeywords(sticker string, keywords []string) (result APIResponse[bool]) {
	return requestGeneric[bool](b, "setStickerKeywords", map[string]any{
		"sticker":  sticker,
		"keywords": keywords,
	})
}

// SetStickerMaskPosition sets mask position of sticker.
//
// https://core.telegram.org/bots/api#setstickermaskposition
func (b *Bot) SetStickerMaskPosition(sticker string, options OptionsSetStickerMaskPosition) (result APIResponse[bool]) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["sticker"] = sticker

	return requestGeneric[bool](b, "setStickerMaskPosition", options)
}

// SendVideo sends a video file.
//
// https://core.telegram.org/bots/api#sendvideo
func (b *Bot) SendVideo(chatID ChatID, video InputFile, options OptionsSendVideo) (result APIResponse[Message]) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["chat_id"] = chatID
	options["video"] = video

	return requestGeneric[Message](b, "sendVideo", options)
}

// SendAnimation sends an animation.
//
// https://core.telegram.org/bots/api#sendanimation
func (b *Bot) SendAnimation(chatID ChatID, animation InputFile, options OptionsSendAnimation) (result APIResponse[Message]) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["chat_id"] = chatID
	options["animation"] = animation

	return requestGeneric[Message](b, "sendAnimation", options)
}

// SendVoice sends a voice file. (.ogg format only, will be played with Telegram itself))
//
// https://core.telegram.org/bots/api#sendvoice
func (b *Bot) SendVoice(chatID ChatID, voice InputFile, options OptionsSendVoice) (result APIResponse[Message]) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["chat_id"] = chatID
	options["voice"] = voice

	return requestGeneric[Message](b, "sendVoice", options)
}

// SendVideoNote sends a video note.
//
// videoNote cannot be a remote http url (not supported yet)
//
// https://core.telegram.org/bots/api#sendvideonote
func (b *Bot) SendVideoNote(chatID ChatID, videoNote InputFile, options OptionsSendVideoNote) (result APIResponse[Message]) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["chat_id"] = chatID
	options["video_note"] = videoNote

	return requestGeneric[Message](b, "sendVideoNote", options)
}

// SendMediaGroup sends a group of photos or videos as an album.
//
// https://core.telegram.org/bots/api#sendmediagroup
func (b *Bot) SendMediaGroup(chatID ChatID, media []InputMedia, options OptionsSendMediaGroup) (result APIResponse[[]Message]) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["chat_id"] = chatID
	options["media"] = media

	return requestGeneric[[]Message](b, "sendMediaGroup", options)
}

// SendLocation sends locations.
//
// https://core.telegram.org/bots/api#sendlocation
func (b *Bot) SendLocation(chatID ChatID, latitude, longitude float32, options OptionsSendLocation) (result APIResponse[Message]) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["chat_id"] = chatID
	options["latitude"] = latitude
	options["longitude"] = longitude

	return requestGeneric[Message](b, "sendLocation", options)
}

// SendVenue sends venues.
//
// https://core.telegram.org/bots/api#sendvenue
func (b *Bot) SendVenue(chatID ChatID, latitude, longitude float32, title, address string, options OptionsSendVenue) (result APIResponse[Message]) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["chat_id"] = chatID
	options["latitude"] = latitude
	options["longitude"] = longitude
	options["title"] = title
	options["address"] = address

	return requestGeneric[Message](b, "sendVenue", options)
}

// SendContact sends contacts.
//
// https://core.telegram.org/bots/api#sendcontact
func (b *Bot) SendContact(chatID ChatID, phoneNumber, firstName string, options OptionsSendContact) (result APIResponse[Message]) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["chat_id"] = chatID
	options["phone_number"] = phoneNumber
	options["first_name"] = firstName

	return requestGeneric[Message](b, "sendContact", options)
}

// SendPoll sends a poll.
//
// https://core.telegram.org/bots/api#sendpoll
func (b *Bot) SendPoll(chatID ChatID, question string, pollOptions []string, options OptionsSendPoll) (result APIResponse[Message]) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["chat_id"] = chatID
	options["question"] = question
	options["options"] = pollOptions

	return requestGeneric[Message](b, "sendPoll", options)
}

// StopPoll stops a poll.
//
// https://core.telegram.org/bots/api#stoppoll
func (b *Bot) StopPoll(chatID ChatID, messageID int64, options OptionsStopPoll) (result APIResponse[Poll]) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["chat_id"] = chatID
	options["message_id"] = messageID

	return requestGeneric[Poll](b, "stopPoll", options)
}

// SendDice sends a random dice.
//
// https://core.telegram.org/bots/api#senddice
func (b *Bot) SendDice(chatID ChatID, options OptionsSendDice) (result APIResponse[Message]) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["chat_id"] = chatID

	return requestGeneric[Message](b, "sendDice", options)
}

// SendChatAction sends chat actions.
//
// https://core.telegram.org/bots/api#sendchataction
func (b *Bot) SendChatAction(chatID ChatID, action ChatAction, options OptionsSendChatAction) (result APIResponse[bool]) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["chat_id"] = chatID
	options["action"] = action

	return requestGeneric[bool](b, "sendChatAction", options)
}

// SetMessageReaction sets message reaction.
//
// https://core.telegram.org/bots/api#setmessagereaction
func (b *Bot) SetMessageReaction(chatID ChatID, messageID int64, options OptionsSetMessageReaction) (result APIResponse[bool]) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["chat_id"] = chatID
	options["message_id"] = messageID

	return requestGeneric[bool](b, "setMessageReaction", options)
}

// GetUserProfilePhotos gets user profile photos.
//
// https://core.telegram.org/bots/api#getuserprofilephotos
func (b *Bot) GetUserProfilePhotos(userID int64, options OptionsGetUserProfilePhotos) (result APIResponse[UserProfilePhotos]) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["user_id"] = userID

	return requestGeneric[UserProfilePhotos](b, "getUserProfilePhotos", options)
}

// GetFile gets file info and prepare for download.
//
// https://core.telegram.org/bots/api#getfile
func (b *Bot) GetFile(fileID string) (result APIResponse[File]) {
	// essential params
	params := map[string]any{
		"file_id": fileID,
	}

	return requestGeneric[File](b, "getFile", params)
}

// GetFileURL gets download link from a given File.
func (b *Bot) GetFileURL(file File) string {
	return fmt.Sprintf("%s%s/%s", fileBaseURL, b.token, *file.FilePath)
}

// BanChatMember bans a chat member.
//
// https://core.telegram.org/bots/api#banchatmember
func (b *Bot) BanChatMember(chatID ChatID, userID int64, options OptionsBanChatMember) (result APIResponse[bool]) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["chat_id"] = chatID
	options["user_id"] = userID

	return requestGeneric[bool](b, "banChatMember", options)
}

// LeaveChat leaves a chat.
//
// https://core.telegram.org/bots/api#leavechat
func (b *Bot) LeaveChat(chatID ChatID) (result APIResponse[bool]) {
	// essential params
	params := map[string]any{
		"chat_id": chatID,
	}

	return requestGeneric[bool](b, "leaveChat", params)
}

// UnbanChatMember unbans a chat member.
//
// https://core.telegram.org/bots/api#unbanchatmember
func (b *Bot) UnbanChatMember(chatID ChatID, userID int64, onlyIfBanned bool) (result APIResponse[bool]) {
	// essential params
	params := map[string]any{
		"chat_id":        chatID,
		"user_id":        userID,
		"only_if_banned": onlyIfBanned,
	}

	return requestGeneric[bool](b, "unbanChatMember", params)
}

// RestrictChatMember restricts a chat member.
//
// https://core.telegram.org/bots/api#restrictchatmember
func (b *Bot) RestrictChatMember(chatID ChatID, userID int64, permissions ChatPermissions, options OptionsRestrictChatMember) (result APIResponse[bool]) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["chat_id"] = chatID
	options["user_id"] = userID
	options["permissions"] = permissions

	return requestGeneric[bool](b, "restrictChatMember", options)
}

// PromoteChatMember promotes a chat member.
//
// https://core.telegram.org/bots/api#promotechatmember
func (b *Bot) PromoteChatMember(chatID ChatID, userID int64, options OptionsPromoteChatMember) (result APIResponse[bool]) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["chat_id"] = chatID
	options["user_id"] = userID

	return requestGeneric[bool](b, "promoteChatMember", options)
}

// SetChatAdministratorCustomTitle sets chat administrator's custom title.
//
// https://core.telegram.org/bots/api#setchatadministratorcustomtitle
func (b *Bot) SetChatAdministratorCustomTitle(chatID ChatID, userID int64, customTitle string) (result APIResponse[bool]) {
	return requestGeneric[bool](b, "setChatAdministratorCustomTitle", map[string]any{
		"chat_id":      chatID,
		"user_id":      userID,
		"custom_title": customTitle,
	})
}

// BanChatSenderChat bans a channel chat in a supergroup or a channel.
//
// https://core.telegram.org/bots/api#banchatsenderchat
func (b *Bot) BanChatSenderChat(chatID ChatID, senderChatID int64) (result APIResponse[bool]) {
	return requestGeneric[bool](b, "banChatSenderChat", map[string]any{
		"chat_id":        chatID,
		"sender_chat_id": senderChatID,
	})
}

// UnbanChatSenderChat unbans a previously banned channel chat in a supergroup or a channel.
//
// https://core.telegram.org/bots/api#unbanchatsenderchat
func (b *Bot) UnbanChatSenderChat(chatID ChatID, senderChatID int64) (result APIResponse[bool]) {
	return requestGeneric[bool](b, "unbanChatSenderChat", map[string]any{
		"chat_id":        chatID,
		"sender_chat_id": senderChatID,
	})
}

// SetChatPermissions sets permissions of a chat.
//
// https://core.telegram.org/bots/api#setchatpermissions
func (b *Bot) SetChatPermissions(chatID ChatID, permissions ChatPermissions, options OptionsSetChatPermissions) (result APIResponse[bool]) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["chat_id"] = chatID
	options["permissions"] = permissions

	return requestGeneric[bool](b, "setChatPermissions", options)
}

// ExportChatInviteLink exports a chat invite link.
//
// https://core.telegram.org/bots/api#exportchatinvitelink
func (b *Bot) ExportChatInviteLink(chatID ChatID) (result APIResponse[string]) {
	// essential params
	params := map[string]any{
		"chat_id": chatID,
	}

	return requestGeneric[string](b, "exportChatInviteLink", params)
}

// CreateChatInviteLink creates a chat invite link.
//
// https://core.telegram.org/bots/api#createchatinvitelink
func (b *Bot) CreateChatInviteLink(chatID ChatID, options OptionsCreateChatInviteLink) (result APIResponse[ChatInviteLink]) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["chat_id"] = chatID

	return requestGeneric[ChatInviteLink](b, "createChatInviteLink", options)
}

// EditChatInviteLink edits a chat invite link.
//
// https://core.telegram.org/bots/api#editchatinvitelink
func (b *Bot) EditChatInviteLink(chatID ChatID, inviteLink string, options OptionsCreateChatInviteLink) (result APIResponse[ChatInviteLink]) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["chat_id"] = chatID
	options["invite_link"] = inviteLink

	return requestGeneric[ChatInviteLink](b, "editChatInviteLink", options)
}

// RevokeChatInviteLink revoks a chat invite link.
//
// https://core.telegram.org/bots/api#revokechatinvitelink
func (b *Bot) RevokeChatInviteLink(chatID ChatID, inviteLink string) (result APIResponse[ChatInviteLink]) {
	return requestGeneric[ChatInviteLink](b, "revokeChatInviteLink", map[string]any{
		"chat_id":     chatID,
		"invite_link": inviteLink,
	})
}

// ApproveChatJoinRequest approves chat join request.
//
// https://core.telegram.org/bots/api#approvechatjoinrequest
func (b *Bot) ApproveChatJoinRequest(chatID ChatID, userID int64) (result APIResponse[bool]) {
	// essential params
	params := map[string]any{
		"chat_id": chatID,
		"user_id": userID,
	}

	return requestGeneric[bool](b, "approveChatJoinRequest", params)
}

// DeclineChatJoinRequest declines chat join request.
//
// https://core.telegram.org/bots/api#declinechatjoinrequest
func (b *Bot) DeclineChatJoinRequest(chatID ChatID, userID int64) (result APIResponse[bool]) {
	// essential params
	params := map[string]any{
		"chat_id": chatID,
		"user_id": userID,
	}

	return requestGeneric[bool](b, "declineChatJoinRequest", params)
}

// SetChatPhoto sets a chat photo.
//
// https://core.telegram.org/bots/api#setchatphoto
func (b *Bot) SetChatPhoto(chatID ChatID, photo InputFile) (result APIResponse[bool]) {
	// essential params
	params := map[string]any{
		"chat_id": chatID,
		"photo":   photo,
	}

	return requestGeneric[bool](b, "setChatPhoto", params)
}

// DeleteChatPhoto deletes a chat photo.
//
// https://core.telegram.org/bots/api#deletechatphoto
func (b *Bot) DeleteChatPhoto(chatID ChatID) (result APIResponse[bool]) {
	// essential params
	params := map[string]any{
		"chat_id": chatID,
	}

	return requestGeneric[bool](b, "deleteChatPhoto", params)
}

// SetChatTitle sets a chat title.
//
// https://core.telegram.org/bots/api#setchattitle
func (b *Bot) SetChatTitle(chatID ChatID, title string) (result APIResponse[bool]) {
	// essential params
	params := map[string]any{
		"chat_id": chatID,
		"title":   title,
	}

	return requestGeneric[bool](b, "setChatTitle", params)
}

// SetChatDescription sets a chat description.
//
// https://core.telegram.org/bots/api#setchatdescription
func (b *Bot) SetChatDescription(chatID ChatID, description string) (result APIResponse[bool]) {
	// essential params
	params := map[string]any{
		"chat_id":     chatID,
		"description": description,
	}

	return requestGeneric[bool](b, "setChatDescription", params)
}

// PinChatMessage pins a chat message.
//
// https://core.telegram.org/bots/api#pinchatmessage
func (b *Bot) PinChatMessage(chatID ChatID, messageID int64, options OptionsPinChatMessage) (result APIResponse[bool]) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["chat_id"] = chatID
	options["message_id"] = messageID

	return requestGeneric[bool](b, "pinChatMessage", options)
}

// UnpinChatMessage unpins a chat message.
//
// https://core.telegram.org/bots/api#unpinchatmessage
func (b *Bot) UnpinChatMessage(chatID ChatID, options OptionsUnpinChatMessage) (result APIResponse[bool]) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["chat_id"] = chatID

	return requestGeneric[bool](b, "unpinChatMessage", options)
}

// UnpinAllChatMessages unpins all chat messages.
//
// https://core.telegram.org/bots/api#unpinallchatmessages
func (b *Bot) UnpinAllChatMessages(chatID ChatID) (result APIResponse[bool]) {
	// essential params
	params := map[string]any{
		"chat_id": chatID,
	}

	return requestGeneric[bool](b, "unpinAllChatMessages", params)
}

// GetChat gets a chat.
//
// https://core.telegram.org/bots/api#getchat
func (b *Bot) GetChat(chatID ChatID) (result APIResponse[Chat]) {
	// essential params
	params := map[string]any{
		"chat_id": chatID,
	}

	return requestGeneric[Chat](b, "getChat", params)
}

// GetChatAdministrators gets chat administrators.
//
// https://core.telegram.org/bots/api#getchatadministrators
func (b *Bot) GetChatAdministrators(chatID ChatID) (result APIResponse[[]ChatMember]) {
	// essential params
	params := map[string]any{
		"chat_id": chatID,
	}

	return requestGeneric[[]ChatMember](b, "getChatAdministrators", params)
}

// GetChatMemberCount gets chat members' count.
//
// https://core.telegram.org/bots/api#getchatmembercount
func (b *Bot) GetChatMemberCount(chatID ChatID) (result APIResponse[int]) {
	// essential params
	params := map[string]any{
		"chat_id": chatID,
	}

	return requestGeneric[int](b, "getChatMemberCount", params)
}

// GetChatMember gets a chat member.
//
// https://core.telegram.org/bots/api#getchatmember
func (b *Bot) GetChatMember(chatID ChatID, userID int64) (result APIResponse[ChatMember]) {
	// essential params
	params := map[string]any{
		"chat_id": chatID,
		"user_id": userID,
	}

	return requestGeneric[ChatMember](b, "getChatMember", params)
}

// SetChatStickerSet sets a chat sticker set.
//
// https://core.telegram.org/bots/api#setchatstickerset
func (b *Bot) SetChatStickerSet(chatID ChatID, stickerSetName string) (result APIResponse[bool]) {
	// essential params
	params := map[string]any{
		"chat_id":          chatID,
		"sticker_set_name": stickerSetName,
	}

	return requestGeneric[bool](b, "setChatStickerSet", params)
}

// DeleteChatStickerSet deletes a chat sticker set.
//
// https://core.telegram.org/bots/api#deletechatstickerset
func (b *Bot) DeleteChatStickerSet(chatID ChatID) (result APIResponse[bool]) {
	// essential params
	params := map[string]any{
		"chat_id": chatID,
	}

	return requestGeneric[bool](b, "deleteChatStickerSet", params)
}

// AnswerCallbackQuery answers a callback query.
//
// https://core.telegram.org/bots/api#answercallbackquery
func (b *Bot) AnswerCallbackQuery(callbackQueryID string, options OptionsAnswerCallbackQuery) (result APIResponse[bool]) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["callback_query_id"] = callbackQueryID

	return requestGeneric[bool](b, "answerCallbackQuery", options)
}

// GetMyCommands fetches commands of this bot.
//
// https://core.telegram.org/bots/api#getmycommands
func (b *Bot) GetMyCommands(options OptionsGetMyCommands) (result APIResponse[[]BotCommand]) {
	return requestGeneric[[]BotCommand](b, "getMyCommands", options)
}

// SetMyName changes the bot's name.
//
// https://core.telegram.org/bots/api#setmyname
func (b *Bot) SetMyName(name string, options OptionsSetMyName) (result APIResponse[bool]) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["name"] = name

	return requestGeneric[bool](b, "setMyName", options)
}

// GetMyName fetches the bot's name.
//
// https://core.telegram.org/bots/api#getmyname
func (b *Bot) GetMyName(options OptionsGetMyName) (result APIResponse[BotName]) {
	return requestGeneric[BotName](b, "getMyName", options)
}

// SetMyDescription sets the bot's description.
//
// https://core.telegram.org/bots/api#setmydescription
func (b *Bot) SetMyDescription(options OptionsSetMyDescription) (result APIResponse[bool]) {
	return requestGeneric[bool](b, "setMyDescription", options)
}

// GetMyDescription gets the bot's description.
//
// https://core.telegram.org/bots/api#setmydescription
func (b *Bot) GetMyDescription(options OptionsGetMyDescription) (result APIResponse[BotDescription]) {
	return requestGeneric[BotDescription](b, "getMyDescription", options)
}

// SetMyShortDescription sets the bot's short description.
//
// https://core.telegram.org/bots/api#setmyshortdescription
func (b *Bot) SetMyShortDescription(options OptionsSetMyShortDescription) (result APIResponse[bool]) {
	return requestGeneric[bool](b, "setMyShortDescription", options)
}

// GetMyShortDescription gets the bot's short description.
//
// https://core.telegram.org/bots/api#getmyshortdescription
func (b *Bot) GetMyShortDescription(options OptionsGetMyShortDescription) (result APIResponse[BotShortDescription]) {
	return requestGeneric[BotShortDescription](b, "getMyShortDescription", options)
}

// GetUserChatBoosts gets boosts of a user.
//
// https://core.telegram.org/bots/api#getuserchatboosts
func (b *Bot) GetUserChatBoosts(chatID ChatID, userID int64) (result APIResponse[UserChatBoosts]) {
	// essential params
	options := map[string]any{
		"chat_id": chatID,
		"user_id": userID,
	}

	return requestGeneric[UserChatBoosts](b, "getUserChatBoosts", options)
}

// SetMyCommands sets commands of this bot.
//
// https://core.telegram.org/bots/api#setmycommands
func (b *Bot) SetMyCommands(commands []BotCommand, options OptionsSetMyCommands) (result APIResponse[bool]) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["commands"] = commands

	return requestGeneric[bool](b, "setMyCommands", options)
}

// DeleteMyCommands deletes commands of this bot.
//
// https://core.telegram.org/bots/api#deletemycommands
func (b *Bot) DeleteMyCommands(options OptionsDeleteMyCommands) (result APIResponse[bool]) {
	return requestGeneric[bool](b, "deleteMyCommands", options)
}

// SetChatMenuButton sets chat menu button.
//
// https://core.telegram.org/bots/api#setchatmenubutton
func (b *Bot) SetChatMenuButton(options OptionsSetChatMenuButton) (result APIResponse[bool]) {
	return requestGeneric[bool](b, "setChatMenuButton", options)
}

// GetChatMenuButton fetches current chat menu button.
//
// https://core.telegram.org/bots/api#getchatmenubutton
func (b *Bot) GetChatMenuButton(options OptionsGetChatMenuButton) (result APIResponse[MenuButton]) {
	return requestGeneric[MenuButton](b, "getChatMenuButton", options)
}

// SetMyDefaultAdministratorRights sets my default administrator rights.
//
// https://core.telegram.org/bots/api#setmydefaultadministratorrights
func (b *Bot) SetMyDefaultAdministratorRights(options OptionsSetMyDefaultAdministratorRights) (result APIResponse[bool]) {
	return requestGeneric[bool](b, "setMyDefaultAdministratorRights", options)
}

// GetMyDefaultAdministratorRights gets my default administrator rights.
//
// https://core.telegram.org/bots/api#getmydefaultadministratorrights
func (b *Bot) GetMyDefaultAdministratorRights(options OptionsGetMyDefaultAdministratorRights) (result APIResponse[bool]) {
	return requestGeneric[bool](b, "getMyDefaultAdministratorRights", options)
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

	return b.requestMessageOrBool("editMessageText", options)
}

// EditMessageCaption edits caption of a message.
//
// https://core.telegram.org/bots/api#editmessagecaption
func (b *Bot) EditMessageCaption(options OptionsEditMessageCaption) (result APIResponseMessageOrBool) {
	if options == nil {
		options = map[string]any{}
	}

	return b.requestMessageOrBool("editMessageCaption", options)
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

	return b.requestMessageOrBool("editMessageMedia", options)
}

// EditMessageReplyMarkup edits reply markup of a message.
//
// https://core.telegram.org/bots/api#editmessagereplymarkup
func (b *Bot) EditMessageReplyMarkup(options OptionsEditMessageReplyMarkup) (result APIResponseMessageOrBool) {
	return b.requestMessageOrBool("editMessageReplyMarkup", options)
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

	return b.requestMessageOrBool("editMessageLiveLocation", options)
}

// StopMessageLiveLocation stops live location of a message.
//
// https://core.telegram.org/bots/api#stopmessagelivelocation
func (b *Bot) StopMessageLiveLocation(options OptionsStopMessageLiveLocation) (result APIResponseMessageOrBool) {
	return b.requestMessageOrBool("stopMessageLiveLocation", options)
}

// DeleteMessage deletes a message.
//
// https://core.telegram.org/bots/api#deletemessage
func (b *Bot) DeleteMessage(chatID ChatID, messageID int64) (result APIResponse[bool]) {
	return requestGeneric[bool](b, "deleteMessage", map[string]any{
		"chat_id":    chatID,
		"message_id": messageID,
	})
}

// DeleteMessages deletes messages.
//
// https://core.telegram.org/bots/api#deletemessages
func (b *Bot) DeleteMessages(chatID ChatID, messageIDs []int64) (result APIResponse[bool]) {
	return requestGeneric[bool](b, "deleteMessages", map[string]any{
		"chat_id":     chatID,
		"message_ids": messageIDs,
	})
}

// AnswerInlineQuery sends answers to an inline query.
//
// results = array of InlineQueryResultArticle, InlineQueryResultPhoto, InlineQueryResultGif, InlineQueryResultMpeg4Gif, or InlineQueryResultVideo.
//
// https://core.telegram.org/bots/api#answerinlinequery
func (b *Bot) AnswerInlineQuery(inlineQueryID string, results []any, options OptionsAnswerInlineQuery) (result APIResponse[bool]) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["inline_query_id"] = inlineQueryID
	options["results"] = results

	return requestGeneric[bool](b, "answerInlineQuery", options)
}

// SendInvoice sends an invoice.
//
// https://core.telegram.org/bots/api#sendinvoice
func (b *Bot) SendInvoice(chatID int64, title, description, payload, providerToken, currency string, prices []LabeledPrice, options OptionsSendInvoice) (result APIResponse[Message]) {
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

	return requestGeneric[Message](b, "sendInvoice", options)
}

// CreateInvoiceLink creates a link for an invoice.
//
// https://core.telegram.org/bots/api#createinvoicelink
func (b *Bot) CreateInvoiceLink(title, description, payload, providerToken, currency string, prices []LabeledPrice, options OptionsCreateInvoiceLink) (result APIResponse[string]) {
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

	return requestGeneric[string](b, "createInvoiceLink", options)
}

// AnswerShippingQuery answers a shipping query.
//
// if ok is true, shippingOptions should be provided.
// otherwise, errorMessage should be provided.
//
// https://core.telegram.org/bots/api#answershippingquery
func (b *Bot) AnswerShippingQuery(shippingQueryID string, ok bool, shippingOptions []ShippingOption, errorMessage *string) (result APIResponse[bool]) {
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

	return requestGeneric[bool](b, "answerShippingQuery", params)
}

// AnswerPreCheckoutQuery answers a pre-checkout query.
//
// https://core.telegram.org/bots/api#answerprecheckoutquery
func (b *Bot) AnswerPreCheckoutQuery(preCheckoutQueryID string, ok bool, errorMessage *string) (result APIResponse[bool]) {
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

	return requestGeneric[bool](b, "answerPreCheckoutQuery", params)
}

// SendGame sends a game.
//
// https://core.telegram.org/bots/api#sendgame
func (b *Bot) SendGame(chatID ChatID, gameShortName string, options OptionsSendGame) (result APIResponse[Message]) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["chat_id"] = chatID
	options["game_short_name"] = gameShortName

	return requestGeneric[Message](b, "sendGame", options)
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

	return b.requestMessageOrBool("setGameScore", options)
}

// GetGameHighScores gets high scores of a game.
//
// https://core.telegram.org/bots/api#getgamehighscores
func (b *Bot) GetGameHighScores(userID int64, options OptionsGetGameHighScores) (result APIResponse[[]GameHighScore]) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["user_id"] = userID

	return requestGeneric[[]GameHighScore](b, "getGameHighScores", options)
}

// AnswerWebAppQuery answers a web app's query
//
// https://core.telegram.org/bots/api#answerwebappquery
func (b *Bot) AnswerWebAppQuery(webAppQueryID string, res InlineQueryResult) (result APIResponse[SentWebAppMessage]) {
	options := map[string]any{
		"web_app_query_id": webAppQueryID,
		"result":           res,
	}

	return requestGeneric[SentWebAppMessage](b, "answerWebAppQuery", options)
}

// CreateForumTopic creates a topic in a forum supergroup chat.
//
// https://core.telegram.org/bots/api#createforumtopic
func (b *Bot) CreateForumTopic(chatID ChatID, name string, options OptionsCreateForumTopic) (result APIResponse[ForumTopic]) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["chat_id"] = chatID
	options["name"] = name

	return requestGeneric[ForumTopic](b, "createForumTopic", options)
}

// EditForumTopic edits a forum topic.
//
// https://core.telegram.org/bots/api#editforumtopic
func (b *Bot) EditForumTopic(chatID ChatID, messageThreadID int64, options OptionsEditForumTopic) (result APIResponse[bool]) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["chat_id"] = chatID
	options["message_thread_id"] = messageThreadID

	return requestGeneric[bool](b, "editForumTopic", options)
}

// CloseForumTopic closes a forum topic.
//
// https://core.telegram.org/bots/api#closeforumtopic
func (b *Bot) CloseForumTopic(chatID ChatID, messageThreadID int64) (result APIResponse[bool]) {
	options := map[string]any{
		"chat_id":           chatID,
		"message_thread_id": messageThreadID,
	}

	return requestGeneric[bool](b, "closeForumTopic", options)
}

// ReopenForumTopic reopens a forum topic.
//
// https://core.telegram.org/bots/api#reopenforumtopic
func (b *Bot) ReopenForumTopic(chatID ChatID, messageThreadID int64) (result APIResponse[bool]) {
	options := map[string]any{
		"chat_id":           chatID,
		"message_thread_id": messageThreadID,
	}

	return requestGeneric[bool](b, "reopenForumTopic", options)
}

// DeleteForumTopic deletes a forum topic.
//
// https://core.telegram.org/bots/api#deleteforumtopic
func (b *Bot) DeleteForumTopic(chatID ChatID, messageThreadID int64) (result APIResponse[bool]) {
	options := map[string]any{
		"chat_id":           chatID,
		"message_thread_id": messageThreadID,
	}

	return requestGeneric[bool](b, "deleteForumTopic", options)
}

// UnpinAllForumTopicMessages unpins all forum topic messages.
//
// https://core.telegram.org/bots/api#unpinallforumtopicmessages
func (b *Bot) UnpinAllForumTopicMessages(chatID ChatID, messageThreadID int64) (result APIResponse[bool]) {
	options := map[string]any{
		"chat_id":           chatID,
		"message_thread_id": messageThreadID,
	}

	return requestGeneric[bool](b, "unpinAllForumTopicMessages", options)
}

// EditGeneralForumTopic edites general forum topic.
//
// https://core.telegram.org/bots/api#editgeneralforumtopic
func (b *Bot) EditGeneralForumTopic(chatID ChatID, name string) (result APIResponse[bool]) {
	options := map[string]any{
		"chat_id": chatID,
		"name":    name,
	}

	return requestGeneric[bool](b, "editGeneralForumTopic", options)
}

// CloseGeneralForumTopic closes general forum topic.
//
// https://core.telegram.org/bots/api#closegeneralforumtopic
func (b *Bot) CloseGeneralForumTopic(chatID ChatID) (result APIResponse[bool]) {
	options := map[string]any{
		"chat_id": chatID,
	}

	return requestGeneric[bool](b, "closeGeneralForumTopic", options)
}

// ReopenGeneralForumTopic reopens general forum topic.
//
// https://core.telegram.org/bots/api#reopengeneralforumtopic
func (b *Bot) ReopenGeneralForumTopic(chatID ChatID) (result APIResponse[bool]) {
	options := map[string]any{
		"chat_id": chatID,
	}

	return requestGeneric[bool](b, "reopenGeneralForumTopic", options)
}

// HideGeneralForumTopic hides general forum topic.
//
// https://core.telegram.org/bots/api#hidegeneralforumtopic
func (b *Bot) HideGeneralForumTopic(chatID ChatID) (result APIResponse[bool]) {
	options := map[string]any{
		"chat_id": chatID,
	}

	return requestGeneric[bool](b, "hideGeneralForumTopic", options)
}

// UnhideGeneralForumTopic unhides general forum topic.
//
// https://core.telegram.org/bots/api#unhidegeneralforumtopic
func (b *Bot) UnhideGeneralForumTopic(chatID ChatID) (result APIResponse[bool]) {
	options := map[string]any{
		"chat_id": chatID,
	}

	return requestGeneric[bool](b, "unhideGeneralForumTopic", options)
}

// https://core.telegram.org/bots/api#unpinallgeneralforumtopicmessages
func (b *Bot) UnpinAllGeneralForumTopicMessages(chatID ChatID) (result APIResponse[bool]) {
	options := map[string]any{
		"chat_id": chatID,
	}

	return requestGeneric[bool](b, "unpinAllGeneralForumTopicMessages", options)
}

// GetForumTopicIconStickers fetches forum topic icon stickers.
//
// https://core.telegram.org/bots/api#getforumtopiciconstickers
func (b *Bot) GetForumTopicIconStickers() (result APIResponse[[]Sticker]) {
	return requestGeneric[[]Sticker](b, "getForumTopicIconStickers", nil)
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
			bytes, err = io.ReadAll(resp.Body)
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
			bytes, err = io.ReadAll(resp.Body)
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

// Send request for APIResponseMessageOrBool and fetch its result.
func (b *Bot) requestMessageOrBool(method string, params map[string]any) (result APIResponseMessageOrBool) {
	var errStr string

	if bytes, err := b.request(method, params); err == nil {
		// try APIResponseMessage type,
		var jsonResponseMessage APIResponse[Message]
		err = json.Unmarshal(bytes, &jsonResponseMessage)
		if err == nil {
			return APIResponseMessageOrBool{
				Ok:            true,
				Description:   jsonResponseMessage.Description,
				ResultMessage: jsonResponseMessage.Result,
			}
		}

		// then try APIResponseBool type,
		var jsonResponseBool APIResponse[bool]
		err = json.Unmarshal(bytes, &jsonResponseBool)
		if err == nil {
			return APIResponseMessageOrBool{
				Ok:          true,
				Description: jsonResponseBool.Description,
				ResultBool:  jsonResponseBool.Result,
			}
		}

		errStr = fmt.Sprintf("json parse error: not in Message nor bool type (%s)", string(bytes))
	} else {
		errStr = fmt.Sprintf("%s failed with error: %s", method, err)
	}

	b.error(errStr)

	return APIResponseMessageOrBool{Ok: false, Description: &errStr}
}

// Send request for APIResponse[T] and fetch its result.
func requestGeneric[T any](b *Bot, method string, params map[string]any) (result APIResponse[T]) {
	var errStr string

	if bytes, err := b.request(method, params); err == nil {
		var jsonResponse APIResponse[T]
		err = json.Unmarshal(bytes, &jsonResponse)
		if err == nil {
			return jsonResponse
		}

		errStr = fmt.Sprintf("json parse error: %s (%s)", err, string(bytes))
	} else {
		errStr = fmt.Sprintf("%s failed with error: %s", method, err)
	}

	b.error(errStr)

	return APIResponse[T]{Ok: false, Description: &errStr}
}

// Handle Webhook request.
func (b *Bot) handleWebhook(writer http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()

	b.verbose("received webhook request: %+v", req)

	if body, err := io.ReadAll(req.Body); err == nil {
		var webhook Update
		if err = json.Unmarshal(body, &webhook); err != nil {
			b.error("error while parsing json (%s)", err)
		} else {
			b.verbose("received webhook body: %s", string(body))

			// if there is a matching command, handle it as a command,
			if !handleUpdateAsCommand(b, webhook) {
				// if there is a matching handler by type, handle with it
				if !handleUpdateByType(b, webhook) {
					// otherwise, handle it manually
					go b.updateHandler(b, webhook, nil)
				}
			}
		}
	} else {
		b.error("error while reading webhook request (%s)", err)

		go b.updateHandler(b, Update{}, err)
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
