package telegrambot

// https://core.telegram.org/bots/api#available-methods

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"mime/multipart"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// GetUpdates retrieves updates from Telegram bot API.
//
// https://core.telegram.org/bots/api#getupdates
func (b *Bot) GetUpdates(
	ctx context.Context,
	options OptionsGetUpdates,
) (result APIResponse[[]Update], err error) {
	if options == nil {
		options = map[string]any{}
	}

	return requestGeneric[[]Update](ctx, b, "getUpdates", options)
}

// SetWebhook sets various options for receiving incoming updates.
//
// `port` should be one of: 443, 80, 88, or 8443.
//
// https://core.telegram.org/bots/api#setwebhook
func (b *Bot) SetWebhook(
	ctx context.Context,
	host string,
	port int,
	options OptionsSetWebhook,
) (result APIResponse[bool], err error) {
	b.webhookHost = host
	b.webhookPort = port
	b.webhookURL = b.getWebhookURL()

	params := map[string]any{
		"url": b.webhookURL,
	}

	if cert, exists := options["certificate"]; exists {
		if filepath, ok := cert.(string); ok {
			var file *os.File
			if file, err = os.Open(filepath); err == nil {
				params["certificate"] = file
			} else {
				err = fmt.Errorf("failed to open certificate: %w", err)
			}
		} else {
			err = fmt.Errorf("given filepath of certificate is not a string")
		}

		if err != nil {
			errStr := err.Error()

			return APIResponse[bool]{
				OK:          false,
				Description: &errStr,
			}, err
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

	return requestGeneric[bool](ctx, b, "setWebhook", params)
}

// DeleteWebhook deletes webhook for this bot.
// (Function GetUpdates will not work if webhook is set, so in that case you'll need to delete it)
//
// https://core.telegram.org/bots/api#deletewebhook
func (b *Bot) DeleteWebhook(
	ctx context.Context,
	dropPendingUpdates bool,
) (result APIResponse[bool], err error) {
	b.webhookHost = ""
	b.webhookPort = 0
	b.webhookURL = ""

	b.verbose("deleting webhook url")

	return requestGeneric[bool](ctx, b, "deleteWebhook", map[string]any{
		"drop_pending_updates": dropPendingUpdates,
	})
}

// GetWebhookInfo gets webhook info for this bot.
//
// https://core.telegram.org/bots/api#getwebhookinfo
func (b *Bot) GetWebhookInfo(
	ctx context.Context,
) (result APIResponse[WebhookInfo], err error) {
	return requestGeneric[WebhookInfo](ctx, b, "getWebhookInfo", map[string]any{})
}

// GetMe gets info of this bot.
//
// https://core.telegram.org/bots/api#getme
func (b *Bot) GetMe(
	ctx context.Context,
) (result APIResponse[User], err error) {
	return requestGeneric[User](ctx, b, "getMe", map[string]any{}) // no params
}

// LogOut logs this bot from cloud Bot API server.
//
// https://core.telegram.org/bots/api#logout
func (b *Bot) LogOut(
	ctx context.Context,
) (result APIResponse[bool], err error) {
	return requestGeneric[bool](ctx, b, "logOut", map[string]any{}) // no params
}

// Close closes this bot from local Bot API server.
//
// https://core.telegram.org/bots/api#close
func (b *Bot) Close(
	ctx context.Context,
) (result APIResponse[bool], err error) {
	return requestGeneric[bool](ctx, b, "close", map[string]any{}) // no params
}

// SendMessage sends a message to the bot.
//
// https://core.telegram.org/bots/api#sendmessage
func (b *Bot) SendMessage(
	ctx context.Context,
	chatID ChatID,
	text string,
	options OptionsSendMessage,
) (result APIResponse[Message], err error) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["chat_id"] = chatID
	options["text"] = text

	return requestGeneric[Message](ctx, b, "sendMessage", options)
}

// ForwardMessage forwards a message.
//
// https://core.telegram.org/bots/api#forwardmessage
func (b *Bot) ForwardMessage(
	ctx context.Context,
	chatID, fromChatID ChatID,
	messageID int64,
	options OptionsForwardMessage,
) (result APIResponse[Message], err error) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["chat_id"] = chatID
	options["from_chat_id"] = fromChatID
	options["message_id"] = messageID

	return requestGeneric[Message](ctx, b, "forwardMessage", options)
}

// ForwardMessages forwards messages.
//
// https://core.telegram.org/bots/api#forwardmessages
func (b *Bot) ForwardMessages(
	ctx context.Context,
	chatID, fromChatID ChatID,
	messageIDs []int64,
	options OptionsForwardMessages,
) (result APIResponse[[]MessageID], err error) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["chat_id"] = chatID
	options["from_chat_id"] = fromChatID
	options["message_ids"] = messageIDs

	return requestGeneric[[]MessageID](ctx, b, "forwardMessages", options)
}

// CopyMessage copies a message.
//
// https://core.telegram.org/bots/api#copymessage
func (b *Bot) CopyMessage(
	ctx context.Context,
	chatID, fromChatID ChatID,
	messageID int64,
	options OptionsCopyMessage,
) (result APIResponse[MessageID], err error) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["chat_id"] = chatID
	options["from_chat_id"] = fromChatID
	options["message_id"] = messageID

	return requestGeneric[MessageID](ctx, b, "copyMessage", options)
}

// CopyMessages copies messages.
//
// https://core.telegram.org/bots/api#copymessages
func (b *Bot) CopyMessages(
	ctx context.Context,
	chatID, fromChatID ChatID,
	messageIDs []int64,
	options OptionsCopyMessages,
) (result APIResponse[[]MessageID], err error) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["chat_id"] = chatID
	options["from_chat_id"] = fromChatID
	options["message_ids"] = messageIDs

	return requestGeneric[[]MessageID](ctx, b, "copyMessages", options)
}

// SendPhoto sends a photo.
//
// https://core.telegram.org/bots/api#sendphoto
func (b *Bot) SendPhoto(
	ctx context.Context,
	chatID ChatID,
	photo InputFile,
	options OptionsSendPhoto,
) (result APIResponse[Message], err error) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["chat_id"] = chatID
	options["photo"] = photo

	return requestGeneric[Message](ctx, b, "sendPhoto", options)
}

// SendAudio sends an audio file. (.mp3 or .m4a format, will be played with external players)
//
// https://core.telegram.org/bots/api#sendaudio
func (b *Bot) SendAudio(
	ctx context.Context,
	chatID ChatID,
	audio InputFile,
	options OptionsSendAudio,
) (result APIResponse[Message], err error) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["chat_id"] = chatID
	options["audio"] = audio

	return requestGeneric[Message](ctx, b, "sendAudio", options)
}

// SendDocument sends a general file.
//
// https://core.telegram.org/bots/api#senddocument
func (b *Bot) SendDocument(
	ctx context.Context,
	chatID ChatID,
	document InputFile,
	options OptionsSendDocument,
) (result APIResponse[Message], err error) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["chat_id"] = chatID
	options["document"] = document

	return requestGeneric[Message](ctx, b, "sendDocument", options)
}

// SendSticker sends a sticker.
//
// https://core.telegram.org/bots/api#sendsticker
func (b *Bot) SendSticker(
	ctx context.Context,
	chatID ChatID,
	sticker InputFile,
	options OptionsSendSticker,
) (result APIResponse[Message], err error) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["chat_id"] = chatID
	options["sticker"] = sticker

	return requestGeneric[Message](ctx, b, "sendSticker", options)
}

// GetStickerSet gets a sticker set.
//
// https://core.telegram.org/bots/api#getstickerset
func (b *Bot) GetStickerSet(
	ctx context.Context,
	name string,
) (result APIResponse[StickerSet], err error) {
	// essential params
	options := map[string]any{
		"name": name,
	}

	return requestGeneric[StickerSet](ctx, b, "getStickerSet", options)
}

// GetCustomEmojiStickers gets custom emoji stickers.
//
// https://core.telegram.org/bots/api#getcustomemojistickers
func (b *Bot) GetCustomEmojiStickers(
	ctx context.Context,
	customEmojiIDs []string,
) (result APIResponse[[]Sticker], err error) {
	// essential options
	options := map[string]any{
		"custom_emoji_ids": customEmojiIDs,
	}

	return requestGeneric[[]Sticker](ctx, b, "getCustomEmojiStickers", options)
}

// UploadStickerFile uploads a sticker file.
//
// https://core.telegram.org/bots/api#uploadstickerfile
func (b *Bot) UploadStickerFile(
	ctx context.Context,
	userID int64,
	sticker InputFile,
	stickerFormat StickerFormat,
) (result APIResponse[File], err error) {
	// essential options
	options := map[string]any{
		"user_id":        userID,
		"sticker":        sticker,
		"sticker_format": stickerFormat,
	}

	return requestGeneric[File](ctx, b, "uploadStickerFile", options)
}

// CreateNewStickerSet creates a new sticker set.
//
// https://core.telegram.org/bots/api#createnewstickerset
func (b *Bot) CreateNewStickerSet(
	ctx context.Context,
	userID int64,
	name, title string,
	stickers []InputSticker,
	options OptionsCreateNewStickerSet,
) (result APIResponse[bool], err error) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["user_id"] = userID
	options["name"] = name
	options["title"] = title
	options["stickers"] = stickers

	return requestGeneric[bool](ctx, b, "createNewStickerSet", options)
}

// AddStickerToSet adds a sticker to set.
//
// https://core.telegram.org/bots/api#addstickertoset
func (b *Bot) AddStickerToSet(
	ctx context.Context,
	userID int64,
	name string,
	sticker InputSticker,
	options OptionsAddStickerToSet,
) (result APIResponse[bool], err error) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["user_id"] = userID
	options["name"] = name
	options["sticker"] = sticker

	return requestGeneric[bool](ctx, b, "addStickerToSet", options)
}

// SetStickerPositionInSet sets sticker position in set.
//
// https://core.telegram.org/bots/api#setstickerpositioninset
func (b *Bot) SetStickerPositionInSet(
	ctx context.Context,
	sticker string,
	position int,
) (result APIResponse[bool], err error) {
	// essential options
	options := map[string]any{
		"sticker":  sticker,
		"position": position,
	}

	return requestGeneric[bool](ctx, b, "setStickerPositionInSet", options)
}

// DeleteStickerFromSet deletes a sticker from set.
//
// https://core.telegram.org/bots/api#deletestickerfromset
func (b *Bot) DeleteStickerFromSet(
	ctx context.Context,
	sticker string,
) (result APIResponse[bool], err error) {
	// essential options
	options := map[string]any{
		"sticker": sticker,
	}

	return requestGeneric[bool](ctx, b, "deleteStickerFromSet", options)
}

// SetStickerSetThumbnail sets a thumbnail of a sticker set.
//
// https://core.telegram.org/bots/api#setstickersetthumbnail
func (b *Bot) SetStickerSetThumbnail(
	ctx context.Context,
	name string,
	userID int64,
	format StickerFormat,
	options OptionsSetStickerSetThumbnail,
) (result APIResponse[bool], err error) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["name"] = name
	options["user_id"] = userID
	options["format"] = format

	return requestGeneric[bool](ctx, b, "setStickerSetThumbnail", options)
}

// SetCustomEmojiStickerSetThumbnail sets the custom emoji sticker set's thumbnail.
//
// https://core.telegram.org/bots/api#setcustomemojistickersetthumbnail
func (b *Bot) SetCustomEmojiStickerSetThumbnail(
	ctx context.Context,
	name string,
	options OptionsSetCustomEmojiStickerSetThumbnail,
) (result APIResponse[bool], err error) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["name"] = name

	return requestGeneric[bool](ctx, b, "setCustomEmojiStickerSetThumbnail", options)
}

// SetStickerSetTitle sets the title of sticker set.
//
// https://core.telegram.org/bots/api#setstickersettitle
func (b *Bot) SetStickerSetTitle(
	ctx context.Context,
	name, title string,
) (result APIResponse[bool], err error) {
	return requestGeneric[bool](ctx, b, "setStickerSetTitle", map[string]any{
		"name":  name,
		"title": title,
	})
}

// DeleteStickerSet deletes a sticker set.
//
// https://core.telegram.org/bots/api#deletestickerset
func (b *Bot) DeleteStickerSet(
	ctx context.Context,
	name string,
) (result APIResponse[bool], err error) {
	return requestGeneric[bool](ctx, b, "deleteStickerSet", map[string]any{
		"name": name,
	})
}

// ReplaceStickerInSet replaces an existing sticker in a sticker set with a new one.
//
// https://core.telegram.org/bots/api#replacestickerinset
func (b *Bot) ReplaceStickerInSet(
	ctx context.Context,
	userID, name, oldSticker string,
	sticker InputSticker,
) (result APIResponse[bool], err error) {
	return requestGeneric[bool](ctx, b, "replaceStickerInSet", map[string]any{
		"user_id":     userID,
		"name":        name,
		"old_sticker": oldSticker,
		"sticker":     sticker,
	})
}

// SetStickerEmojiList sets the emoji list of sticker set.
//
// https://core.telegram.org/bots/api#setstickeremojilist
func (b *Bot) SetStickerEmojiList(
	ctx context.Context,
	sticker string,
	emojiList []string,
) (result APIResponse[bool], err error) {
	return requestGeneric[bool](ctx, b, "setStickerEmojiList", map[string]any{
		"sticker":    sticker,
		"emoji_list": emojiList,
	})
}

// SetStickerKeywords sets the keywords of sticker.
//
// https://core.telegram.org/bots/api#setstickerkeywords
func (b *Bot) SetStickerKeywords(
	ctx context.Context,
	sticker string,
	keywords []string,
) (result APIResponse[bool], err error) {
	return requestGeneric[bool](ctx, b, "setStickerKeywords", map[string]any{
		"sticker":  sticker,
		"keywords": keywords,
	})
}

// SetStickerMaskPosition sets mask position of sticker.
//
// https://core.telegram.org/bots/api#setstickermaskposition
func (b *Bot) SetStickerMaskPosition(
	ctx context.Context,
	sticker string,
	options OptionsSetStickerMaskPosition,
) (result APIResponse[bool], err error) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["sticker"] = sticker

	return requestGeneric[bool](ctx, b, "setStickerMaskPosition", options)
}

// GetAvailableGifts returns the list of gifts that can be sent by the bot to users.
//
// https://core.telegram.org/bots/api#getavailablegifts
func (b *Bot) GetAvailableGifts(
	ctx context.Context,
) (result APIResponse[Gifts], err error) {
	return requestGeneric[Gifts](ctx, b, "getAvailableGifts", nil)
}

// SendGift sends a gift to the given user.
//
// https://core.telegram.org/bots/api#sendgift
func (b *Bot) SendGift(
	ctx context.Context,
	giftID string,
	options OptionsSendGift,
) (result APIResponse[bool], err error) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["gift_id"] = giftID

	return requestGeneric[bool](ctx, b, "sendGift", options)
}

// GiftPremiumSubscription gifts a Telegram Premium subscription to the given user.
//
// https://core.telegram.org/bots/api#giftpremiumsubscription
func (b *Bot) GiftPremiumSubscription(
	ctx context.Context,
	userID int64,
	monthCount, starCount int,
	options OptionsGiftPremiumSubscription,
) (result APIResponse[bool], err error) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["user_id"] = userID
	options["month_count"] = monthCount
	options["star_count"] = starCount

	return requestGeneric[bool](ctx, b, "giftPremiumSubscription", options)
}

// VerifyUser verifies a user.
//
// https://core.telegram.org/bots/api#verifyuser
func (b *Bot) VerifyUser(
	ctx context.Context,
	userID int64,
	options OptionsVerifyUser,
) (result APIResponse[bool], err error) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["user_id"] = userID

	return requestGeneric[bool](ctx, b, "verifyUser", options)
}

// VerifyChat verifies a chat.
//
// https://core.telegram.org/bots/api#verifychat
func (b *Bot) VerifyChat(
	ctx context.Context,
	chatID ChatID,
	options OptionsVerifyChat,
) (result APIResponse[bool], err error) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["chat_id"] = chatID

	return requestGeneric[bool](ctx, b, "verifyChat", options)
}

// RemoveUserVerification removes a user's verification.
//
// https://core.telegram.org/bots/api#removeuserverification
func (b *Bot) RemoveUserVerification(
	ctx context.Context,
	userID int64,
) (result APIResponse[bool], err error) {
	return requestGeneric[bool](ctx, b, "removeUserVerification", map[string]any{
		"user_id": userID,
	})
}

// RemoveChatVerification removes a chat's verification.
//
// https://core.telegram.org/bots/api#removechatverification
func (b *Bot) RemoveChatVerification(
	ctx context.Context,
	chatID ChatID,
) (result APIResponse[bool], err error) {
	return requestGeneric[bool](ctx, b, "removeChatVerification", map[string]any{
		"chat_id": chatID,
	})
}

// ReadBusinessMessage marks an incoming message as read on behalf of a business account.
//
// https://core.telegram.org/bots/api#readbusinessmessage
func (b *Bot) ReadBusinessMessage(
	ctx context.Context,
	businessConnectionID string,
	chatID, messageID int64,
) (result APIResponse[bool], err error) {
	return requestGeneric[bool](ctx, b, "readBusinessMessage", map[string]any{
		"business_connection_id": businessConnectionID,
		"chat_id":                chatID,
		"message_id":             messageID,
	})
}

// DeleteBusinessMessages deletes messages on behalf of a business account.
//
// https://core.telegram.org/bots/api#deletebusinessmessages
func (b *Bot) DeleteBusinessMessages(
	ctx context.Context,
	businessConnectionID string,
	messageIDs []int64,
) (result APIResponse[bool], err error) {
	return requestGeneric[bool](ctx, b, "deleteBusinessMessages", map[string]any{
		"business_connection_id": businessConnectionID,
		"message_ids":            messageIDs,
	})
}

// SetBusinessAccountName changes the first and last name of a managed business account.
//
// https://core.telegram.org/bots/api#setbusinessaccountname
func (b *Bot) SetBusinessAccountName(
	ctx context.Context,
	businessConnectionID string,
	firstName string,
	options OptionsSetBusinessAccountName,
) (result APIResponse[bool], err error) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["business_connection_id"] = businessConnectionID
	options["first_name"] = firstName

	return requestGeneric[bool](ctx, b, "setBusinessAccountName", options)
}

// SetBusinessAccountUsername changes the username of a managed business account.
//
// https://core.telegram.org/bots/api#setbusinessaccountusername
func (b *Bot) SetBusinessAccountUsername(
	ctx context.Context,
	businessConnectionID string,
	options OptionsSetBusinessAccountUsername,
) (result APIResponse[bool], err error) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["business_connection_id"] = businessConnectionID

	return requestGeneric[bool](ctx, b, "setBusinessAccountUsername", options)
}

// SetBusinessAccountBio changes the bio of a managed business account.
//
// https://core.telegram.org/bots/api#setbusinessaccountbio
func (b *Bot) SetBusinessAccountBio(
	ctx context.Context,
	businessConnectionID string,
	options OptionsSetBusinessAccountBio,
) (result APIResponse[bool], err error) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["business_connection_id"] = businessConnectionID

	return requestGeneric[bool](ctx, b, "setBusinessAccountBio", options)
}

// SetBusinessAccountProfilePhoto changes the profile photo of a managed business account.
//
// https://core.telegram.org/bots/api#setbusinessaccountprofilephoto
func (b *Bot) SetBusinessAccountProfilePhoto(
	ctx context.Context,
	businessConnectionID string,
	photo InputProfilePhoto,
	options OptionsSetBusinessAccountProfilePhoto,
) (result APIResponse[bool], err error) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["business_connection_id"] = businessConnectionID
	options["photo"] = photo

	return requestGeneric[bool](ctx, b, "setBusinessAccountProfilePhoto", options)
}

// RemoveBusinessAccountProfilePhoto removes the current profile photo of a managed business account.
//
// https://core.telegram.org/bots/api#removebusinessaccountprofilephoto
func (b *Bot) RemoveBusinessAccountProfilePhoto(
	ctx context.Context,
	businessConnectionID string,
	options OptionsRemoveBusinessAccountProfilePhoto,
) (result APIResponse[bool], err error) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["business_connection_id"] = businessConnectionID

	return requestGeneric[bool](ctx, b, "removeBusinessAccountProfilePhoto", options)
}

// SetBusinessAccountGiftSettings changes the privacy settings pertaining to incoming gifts in a managed business account.
//
// https://core.telegram.org/bots/api#setbusinessaccountgiftsettings
func (b *Bot) SetBusinessAccountGiftSettings(
	ctx context.Context,
	businessConnectionID string,
	showGiftButton bool,
	acceptedGiftTypes AcceptedGiftTypes,
) (result APIResponse[bool], err error) {
	return requestGeneric[bool](ctx, b, "setBusinessAccountGiftSettings", map[string]any{
		"business_connection_id": businessConnectionID,
		"show_gift_button":       showGiftButton,
		"accepted_gift_types":    acceptedGiftTypes,
	})
}

// GetBusinessAccountStarBalance returns the amount of Telegram Stars owned by a managed business account.
//
// https://core.telegram.org/bots/api#getbusinessaccountstarbalance
func (b *Bot) GetBusinessAccountStarBalance(
	ctx context.Context,
	businessConnectionID string,
) (result APIResponse[StarAmount], err error) {
	return requestGeneric[StarAmount](ctx, b, "getBusinessAccountStarBalance", map[string]any{
		"business_connection_id": businessConnectionID,
	})
}

// TransferBusinessAccountStars transfers Telegram Stars from the business account balance to the bot's balance.
//
// https://core.telegram.org/bots/api#transferbusinessaccountstars
func (b *Bot) TransferBusinessAccountStars(
	ctx context.Context,
	businessConnectionID string,
	starCount int,
) (result APIResponse[bool], err error) {
	return requestGeneric[bool](ctx, b, "transferBusinessAccountStars", map[string]any{
		"business_connection_id": businessConnectionID,
		"star_count":             starCount,
	})
}

// GetBusinessAccountGifts returns the gifts received and owned by a managed business account.
//
// https://core.telegram.org/bots/api#getbusinessaccountgifts
func (b *Bot) GetBusinessAccountGifts(
	ctx context.Context,
	businessConnectionID string,
	options OptionsGetBusinessAccountGifts,
) (result APIResponse[OwnedGifts], err error) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["business_connection_id"] = businessConnectionID

	return requestGeneric[OwnedGifts](ctx, b, "getBusinessAccountGifts", options)
}

// GetUserGifts returns the gifts owned and hosted by a user.
//
// https://core.telegram.org/bots/api#getusergifts
func (b *Bot) GetUserGifts(
	ctx context.Context,
	userID int64,
	options OptionsGetUserGifts,
) (result APIResponse[OwnedGifts], err error) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["user_id"] = userID

	return requestGeneric[OwnedGifts](ctx, b, "getUserGifts", options)
}

// GetChatGifts returns the gifts owned by a chat.
//
// https://core.telegram.org/bots/api#getchatgifts
func (b *Bot) GetChatGifts(
	ctx context.Context,
	chatID ChatID,
	options OptionsGetChatGifts,
) (result APIResponse[OwnedGifts], err error) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["chat_id"] = chatID

	return requestGeneric[OwnedGifts](ctx, b, "getChatGifts", options)
}

// ConvertGiftToStars converts a given regular gift to Telegram Stars.
//
// https://core.telegram.org/bots/api#convertgifttostars
func (b *Bot) ConvertGiftToStars(
	ctx context.Context,
	businessConnectionID, ownedGiftID string,
) (result APIResponse[bool], err error) {
	return requestGeneric[bool](ctx, b, "convertGiftToStars", map[string]any{
		"business_connection_id": businessConnectionID,
		"owned_gift_id":          ownedGiftID,
	})
}

// UpgradeGift upgrades a given regular gift to a unique gift.
//
// https://core.telegram.org/bots/api#upgradegift
func (b *Bot) UpgradeGift(
	ctx context.Context,
	businessConnectionID, ownedGiftID string,
	options OptionsUpgradeGift,
) (result APIResponse[bool], err error) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["business_connection_id"] = businessConnectionID
	options["owned_gift_id"] = ownedGiftID

	return requestGeneric[bool](ctx, b, "upgradeGift", options)
}

// TransferGift transfers an owned unique gift to another user.
//
// https://core.telegram.org/bots/api#transfergift
func (b *Bot) TransferGift(
	ctx context.Context,
	businessConnectionID, ownedGiftID string,
	newOwnerChatID int64,
	options OptionsTransferGift,
) (result APIResponse[bool], err error) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["business_connection_id"] = businessConnectionID
	options["owned_gift_id"] = ownedGiftID
	options["new_owner_chat_id"] = newOwnerChatID

	return requestGeneric[bool](ctx, b, "transferGift", options)
}

// PostStory posts a story on behalf of a managed business account.
//
// https://core.telegram.org/bots/api#poststory
func (b *Bot) PostStory(
	ctx context.Context,
	businessConnectionID string,
	content InputStoryContent,
	activePeriod int,
	options OptionsPostStory,
) (result APIResponse[Story], err error) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["business_connection_id"] = businessConnectionID
	options["content"] = content
	options["active_period"] = activePeriod

	return requestGeneric[Story](ctx, b, "postStory", options)
}

// RepostStory reposts a story on behalf of a business account from another business account.
//
// https://core.telegram.org/bots/api#repoststory
func (b *Bot) RepostStory(
	ctx context.Context,
	businessConnectionID string,
	fromChatID int64,
	fromStoryID int64,
	activePeriod int,
	options OptionsRepostStory,
) (result APIResponse[Story], err error) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["business_connection_id"] = businessConnectionID
	options["from_chat_id"] = fromChatID
	options["from_story_id"] = fromStoryID
	options["active_period"] = activePeriod

	return requestGeneric[Story](ctx, b, "repostStory", options)
}

// EditStory edits a story previously posted by the bot on behalf of a managed business account.
//
// https://core.telegram.org/bots/api#editstory
func (b *Bot) EditStory(
	ctx context.Context,
	businessConnectionID string,
	storyID int64,
	content InputStoryContent,
	options OptionsEditStory,
) (result APIResponse[Story], err error) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["business_connection_id"] = businessConnectionID
	options["story_id"] = storyID
	options["content"] = content

	return requestGeneric[Story](ctx, b, "editStory", options)
}

// DeleteStory deletes a story previously posted by the bot on behalf of a managed business account.
//
// https://core.telegram.org/bots/api#deletestory
func (b *Bot) DeleteStory(
	ctx context.Context,
	businessConnectionID string,
	storyID int64,
) (result APIResponse[bool], err error) {
	return requestGeneric[bool](ctx, b, "deleteStory", map[string]any{
		"business_connection_id": businessConnectionID,
		"story_id":               storyID,
	})
}

// SendVideo sends a video file.
//
// https://core.telegram.org/bots/api#sendvideo
func (b *Bot) SendVideo(
	ctx context.Context,
	chatID ChatID,
	video InputFile,
	options OptionsSendVideo,
) (result APIResponse[Message], err error) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["chat_id"] = chatID
	options["video"] = video

	return requestGeneric[Message](ctx, b, "sendVideo", options)
}

// SendAnimation sends an animation.
//
// https://core.telegram.org/bots/api#sendanimation
func (b *Bot) SendAnimation(
	ctx context.Context,
	chatID ChatID,
	animation InputFile,
	options OptionsSendAnimation,
) (result APIResponse[Message], err error) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["chat_id"] = chatID
	options["animation"] = animation

	return requestGeneric[Message](ctx, b, "sendAnimation", options)
}

// SendVoice sends a voice file. (.ogg format only, will be played with Telegram itself))
//
// https://core.telegram.org/bots/api#sendvoice
func (b *Bot) SendVoice(
	ctx context.Context,
	chatID ChatID,
	voice InputFile,
	options OptionsSendVoice,
) (result APIResponse[Message], err error) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["chat_id"] = chatID
	options["voice"] = voice

	return requestGeneric[Message](ctx, b, "sendVoice", options)
}

// SendVideoNote sends a video note.
//
// videoNote cannot be a remote http url (not supported yet)
//
// https://core.telegram.org/bots/api#sendvideonote
func (b *Bot) SendVideoNote(
	ctx context.Context,
	chatID ChatID,
	videoNote InputFile,
	options OptionsSendVideoNote,
) (result APIResponse[Message], err error) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["chat_id"] = chatID
	options["video_note"] = videoNote

	return requestGeneric[Message](ctx, b, "sendVideoNote", options)
}

// SendPaidMedia sends paid media.
//
// https://core.telegram.org/bots/api#sendpaidmedia
func (b *Bot) SendPaidMedia(
	ctx context.Context,
	chatID ChatID,
	starCount int,
	media []InputPaidMedia,
	options OptionsSendPaidMedia,
) (result APIResponse[Message], err error) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["chat_id"] = chatID
	options["star_count"] = starCount
	options["media"] = media

	return requestGeneric[Message](ctx, b, "sendPaidMedia", options)
}

// SendMediaGroup sends a group of photos or videos as an album.
//
// https://core.telegram.org/bots/api#sendmediagroup
func (b *Bot) SendMediaGroup(
	ctx context.Context,
	chatID ChatID,
	media []InputMedia,
	options OptionsSendMediaGroup,
) (result APIResponse[[]Message], err error) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["chat_id"] = chatID
	options["media"] = media

	return requestGeneric[[]Message](ctx, b, "sendMediaGroup", options)
}

// SendLocation sends locations.
//
// https://core.telegram.org/bots/api#sendlocation
func (b *Bot) SendLocation(
	ctx context.Context,
	chatID ChatID,
	latitude, longitude float32,
	options OptionsSendLocation,
) (result APIResponse[Message], err error) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["chat_id"] = chatID
	options["latitude"] = latitude
	options["longitude"] = longitude

	return requestGeneric[Message](ctx, b, "sendLocation", options)
}

// SendVenue sends venues.
//
// https://core.telegram.org/bots/api#sendvenue
func (b *Bot) SendVenue(
	ctx context.Context,
	chatID ChatID,
	latitude, longitude float32,
	title, address string,
	options OptionsSendVenue,
) (result APIResponse[Message], err error) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["chat_id"] = chatID
	options["latitude"] = latitude
	options["longitude"] = longitude
	options["title"] = title
	options["address"] = address

	return requestGeneric[Message](ctx, b, "sendVenue", options)
}

// SendContact sends contacts.
//
// https://core.telegram.org/bots/api#sendcontact
func (b *Bot) SendContact(
	ctx context.Context,
	chatID ChatID,
	phoneNumber, firstName string,
	options OptionsSendContact,
) (result APIResponse[Message], err error) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["chat_id"] = chatID
	options["phone_number"] = phoneNumber
	options["first_name"] = firstName

	return requestGeneric[Message](ctx, b, "sendContact", options)
}

// SendPoll sends a poll.
//
// https://core.telegram.org/bots/api#sendpoll
func (b *Bot) SendPoll(
	ctx context.Context,
	chatID ChatID,
	question string,
	pollOptions []InputPollOption,
	options OptionsSendPoll,
) (result APIResponse[Message], err error) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["chat_id"] = chatID
	options["question"] = question
	options["options"] = pollOptions

	return requestGeneric[Message](ctx, b, "sendPoll", options)
}

// StopPoll stops a poll.
//
// https://core.telegram.org/bots/api#stoppoll
func (b *Bot) StopPoll(
	ctx context.Context,
	chatID ChatID,
	messageID int64,
	options OptionsStopPoll,
) (result APIResponse[Poll], err error) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["chat_id"] = chatID
	options["message_id"] = messageID

	return requestGeneric[Poll](ctx, b, "stopPoll", options)
}

// ApproveSuggestedPost approves a suggested post.
//
// https://core.telegram.org/bots/api#approvesuggestedpost
func (b *Bot) ApproveSuggestedPost(
	ctx context.Context,
	chatID int64,
	messageID int64,
	options OptionsApproveSuggestedPost,
) (result APIResponse[bool], err error) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["chat_id"] = chatID
	options["message_id"] = messageID

	return requestGeneric[bool](ctx, b, "approveSuggestedPost", options)
}

// DeclineSuggestedPost declines a suggested post.
//
// https://core.telegram.org/bots/api#declinesuggestedpost
func (b *Bot) DeclineSuggestedPost(
	ctx context.Context,
	chatID int64,
	messageID int64,
	options OptionsDeclineSuggestedPost,
) (result APIResponse[bool], err error) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["chat_id"] = chatID
	options["message_id"] = messageID

	return requestGeneric[bool](ctx, b, "declineSuggestedPost", options)
}

// SendChecklist sends a checklist.
//
// https://core.telegram.org/bots/api#sendchecklist
func (b *Bot) SendChecklist(
	ctx context.Context,
	businessConnectionID string,
	chatID ChatID,
	checklist InputChecklist,
	options OptionsSendChecklist,
) (result APIResponse[Message], err error) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["business_connection_id"] = businessConnectionID
	options["chat_id"] = chatID
	options["checklist"] = checklist

	return requestGeneric[Message](ctx, b, "sendChecklist", options)
}

// SendDice sends a random dice.
//
// https://core.telegram.org/bots/api#senddice
func (b *Bot) SendDice(
	ctx context.Context,
	chatID ChatID,
	options OptionsSendDice,
) (result APIResponse[Message], err error) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["chat_id"] = chatID

	return requestGeneric[Message](ctx, b, "sendDice", options)
}

// SendMessageDraft sends a message draft.
//
// NOTE: supported only for bots with forum topic mode enabled
//
// https://core.telegram.org/bots/api#sendmessagedraft
func (b *Bot) SendMessageDraft(
	ctx context.Context,
	chatID ChatID,
	draftID int64,
	text string,
	options OptionsSendMessageDraft,
) (result APIResponse[bool], err error) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["chat_id"] = chatID
	options["draft_id"] = draftID
	options["text"] = text

	return requestGeneric[bool](ctx, b, "sendMessageDraft", options)
}

// SendChatAction sends chat actions.
//
// https://core.telegram.org/bots/api#sendchataction
func (b *Bot) SendChatAction(
	ctx context.Context,
	chatID ChatID,
	action ChatAction,
	options OptionsSendChatAction,
) (result APIResponse[bool], err error) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["chat_id"] = chatID
	options["action"] = action

	return requestGeneric[bool](ctx, b, "sendChatAction", options)
}

// SetMessageReaction sets message reaction.
//
// https://core.telegram.org/bots/api#setmessagereaction
func (b *Bot) SetMessageReaction(
	ctx context.Context,
	chatID ChatID,
	messageID int64,
	options OptionsSetMessageReaction,
) (result APIResponse[bool], err error) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["chat_id"] = chatID
	options["message_id"] = messageID

	return requestGeneric[bool](ctx, b, "setMessageReaction", options)
}

// GetUserProfilePhotos gets user profile photos.
//
// https://core.telegram.org/bots/api#getuserprofilephotos
func (b *Bot) GetUserProfilePhotos(
	ctx context.Context,
	userID int64,
	options OptionsGetUserProfilePhotos,
) (result APIResponse[UserProfilePhotos], err error) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["user_id"] = userID

	return requestGeneric[UserProfilePhotos](ctx, b, "getUserProfilePhotos", options)
}

// https://core.telegram.org/bots/api#getuserprofileaudios
func (b *Bot) GetUserProfileAudios(
	ctx context.Context,
	userID int64,
	options OptionsGetUserProfileAudios,
) (result APIResponse[UserProfileAudios], err error) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["user_id"] = userID

	return requestGeneric[UserProfileAudios](ctx, b, "getUserProfileAudios", options)
}

// SetUserEmojiStatus changes the emoji status for a given user that previously allowed the bot to manage their emoji status via the Mini App.
//
// https://core.telegram.org/bots/api#setuseremojistatus
func (b *Bot) SetUserEmojiStatus(
	ctx context.Context,
	userID int64,
	options OptionsSetUserEmojiStatus,
) (result APIResponse[bool], err error) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["user_id"] = userID

	return requestGeneric[bool](ctx, b, "setUserEmojiStatus", options)
}

// GetFile gets file info and prepare for download.
//
// https://core.telegram.org/bots/api#getfile
func (b *Bot) GetFile(
	ctx context.Context,
	fileID string,
) (result APIResponse[File], err error) {
	// essential options
	options := map[string]any{
		"file_id": fileID,
	}

	return requestGeneric[File](ctx, b, "getFile", options)
}

// GetFileURL gets download link from a given File.
func (b *Bot) GetFileURL(file File) string {
	return fmt.Sprintf("%s%s/%s", fileBaseURL, b.token, *file.FilePath)
}

// BanChatMember bans a chat member.
//
// https://core.telegram.org/bots/api#banchatmember
func (b *Bot) BanChatMember(
	ctx context.Context,
	chatID ChatID,
	userID int64,
	options OptionsBanChatMember,
) (result APIResponse[bool], err error) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["chat_id"] = chatID
	options["user_id"] = userID

	return requestGeneric[bool](ctx, b, "banChatMember", options)
}

// LeaveChat leaves a chat.
//
// https://core.telegram.org/bots/api#leavechat
func (b *Bot) LeaveChat(
	ctx context.Context,
	chatID ChatID,
) (result APIResponse[bool], err error) {
	// essential options
	options := map[string]any{
		"chat_id": chatID,
	}

	return requestGeneric[bool](ctx, b, "leaveChat", options)
}

// UnbanChatMember unbans a chat member.
//
// https://core.telegram.org/bots/api#unbanchatmember
func (b *Bot) UnbanChatMember(
	ctx context.Context,
	chatID ChatID,
	userID int64,
	onlyIfBanned bool,
) (result APIResponse[bool], err error) {
	// essential options
	options := map[string]any{
		"chat_id":        chatID,
		"user_id":        userID,
		"only_if_banned": onlyIfBanned,
	}

	return requestGeneric[bool](ctx, b, "unbanChatMember", options)
}

// RestrictChatMember restricts a chat member.
//
// https://core.telegram.org/bots/api#restrictchatmember
func (b *Bot) RestrictChatMember(
	ctx context.Context,
	chatID ChatID,
	userID int64,
	permissions ChatPermissions,
	options OptionsRestrictChatMember,
) (result APIResponse[bool], err error) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["chat_id"] = chatID
	options["user_id"] = userID
	options["permissions"] = permissions

	return requestGeneric[bool](ctx, b, "restrictChatMember", options)
}

// PromoteChatMember promotes a chat member.
//
// https://core.telegram.org/bots/api#promotechatmember
func (b *Bot) PromoteChatMember(
	ctx context.Context,
	chatID ChatID,
	userID int64,
	options OptionsPromoteChatMember,
) (result APIResponse[bool], err error) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["chat_id"] = chatID
	options["user_id"] = userID

	return requestGeneric[bool](ctx, b, "promoteChatMember", options)
}

// SetChatAdministratorCustomTitle sets chat administrator's custom title.
//
// https://core.telegram.org/bots/api#setchatadministratorcustomtitle
func (b *Bot) SetChatAdministratorCustomTitle(
	ctx context.Context,
	chatID ChatID,
	userID int64,
	customTitle string,
) (result APIResponse[bool], err error) {
	return requestGeneric[bool](ctx, b, "setChatAdministratorCustomTitle", map[string]any{
		"chat_id":      chatID,
		"user_id":      userID,
		"custom_title": customTitle,
	})
}

// BanChatSenderChat bans a channel chat in a supergroup or a channel.
//
// https://core.telegram.org/bots/api#banchatsenderchat
func (b *Bot) BanChatSenderChat(
	ctx context.Context,
	chatID ChatID,
	senderChatID int64,
) (result APIResponse[bool], err error) {
	return requestGeneric[bool](ctx, b, "banChatSenderChat", map[string]any{
		"chat_id":        chatID,
		"sender_chat_id": senderChatID,
	})
}

// UnbanChatSenderChat unbans a previously banned channel chat in a supergroup or a channel.
//
// https://core.telegram.org/bots/api#unbanchatsenderchat
func (b *Bot) UnbanChatSenderChat(
	ctx context.Context,
	chatID ChatID,
	senderChatID int64,
) (result APIResponse[bool], err error) {
	return requestGeneric[bool](ctx, b, "unbanChatSenderChat", map[string]any{
		"chat_id":        chatID,
		"sender_chat_id": senderChatID,
	})
}

// SetChatPermissions sets permissions of a chat.
//
// https://core.telegram.org/bots/api#setchatpermissions
func (b *Bot) SetChatPermissions(
	ctx context.Context,
	chatID ChatID,
	permissions ChatPermissions,
	options OptionsSetChatPermissions,
) (result APIResponse[bool], err error) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["chat_id"] = chatID
	options["permissions"] = permissions

	return requestGeneric[bool](ctx, b, "setChatPermissions", options)
}

// ExportChatInviteLink exports a chat invite link.
//
// https://core.telegram.org/bots/api#exportchatinvitelink
func (b *Bot) ExportChatInviteLink(
	ctx context.Context,
	chatID ChatID,
) (result APIResponse[string], err error) {
	// essential options
	options := map[string]any{
		"chat_id": chatID,
	}

	return requestGeneric[string](ctx, b, "exportChatInviteLink", options)
}

// CreateChatInviteLink creates a chat invite link.
//
// https://core.telegram.org/bots/api#createchatinvitelink
func (b *Bot) CreateChatInviteLink(
	ctx context.Context,
	chatID ChatID,
	options OptionsCreateChatInviteLink,
) (result APIResponse[ChatInviteLink], err error) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["chat_id"] = chatID

	return requestGeneric[ChatInviteLink](ctx, b, "createChatInviteLink", options)
}

// EditChatInviteLink edits a chat invite link.
//
// https://core.telegram.org/bots/api#editchatinvitelink
func (b *Bot) EditChatInviteLink(
	ctx context.Context,
	chatID ChatID,
	inviteLink string,
	options OptionsCreateChatInviteLink,
) (result APIResponse[ChatInviteLink], err error) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["chat_id"] = chatID
	options["invite_link"] = inviteLink

	return requestGeneric[ChatInviteLink](ctx, b, "editChatInviteLink", options)
}

// CreateChatSubscriptionInviteLink creates a subscription invite link for a channel chat.
//
// https://core.telegram.org/bots/api#createchatsubscriptioninvitelink
func (b *Bot) CreateChatSubscriptionInviteLink(
	ctx context.Context,
	chatID ChatID,
	subscriptionPeriod, subscriptionPrice int,
	options OptionsCreateChatSubscriptionInviteLink,
) (result APIResponse[ChatInviteLink], err error) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["chat_id"] = chatID
	options["subscription_period"] = subscriptionPeriod
	options["subscription_price"] = subscriptionPrice

	return requestGeneric[ChatInviteLink](ctx, b, "createChatSubscriptionInviteLink", options)
}

// EditChatSubscriptionInviteLink edits a subscription invite link created by the bot.
//
// https://core.telegram.org/bots/api#editchatsubscriptioninvitelink
func (b *Bot) EditChatSubscriptionInviteLink(
	ctx context.Context,
	chatID ChatID,
	inviteLink string,
	options OptionsEditChatSubscriptionInviteLink,
) (result APIResponse[ChatInviteLink], err error) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["chat_id"] = chatID
	options["invite_link"] = inviteLink

	return requestGeneric[ChatInviteLink](ctx, b, "editChatSubscriptionInviteLink", options)
}

// RevokeChatInviteLink revoks a chat invite link.
//
// https://core.telegram.org/bots/api#revokechatinvitelink
func (b *Bot) RevokeChatInviteLink(
	ctx context.Context,
	chatID ChatID,
	inviteLink string,
) (result APIResponse[ChatInviteLink], err error) {
	return requestGeneric[ChatInviteLink](ctx, b, "revokeChatInviteLink", map[string]any{
		"chat_id":     chatID,
		"invite_link": inviteLink,
	})
}

// ApproveChatJoinRequest approves chat join request.
//
// https://core.telegram.org/bots/api#approvechatjoinrequest
func (b *Bot) ApproveChatJoinRequest(
	ctx context.Context,
	chatID ChatID,
	userID int64,
) (result APIResponse[bool], err error) {
	// essential options
	options := map[string]any{
		"chat_id": chatID,
		"user_id": userID,
	}

	return requestGeneric[bool](ctx, b, "approveChatJoinRequest", options)
}

// DeclineChatJoinRequest declines chat join request.
//
// https://core.telegram.org/bots/api#declinechatjoinrequest
func (b *Bot) DeclineChatJoinRequest(
	ctx context.Context,
	chatID ChatID,
	userID int64,
) (result APIResponse[bool], err error) {
	// essential options
	options := map[string]any{
		"chat_id": chatID,
		"user_id": userID,
	}

	return requestGeneric[bool](ctx, b, "declineChatJoinRequest", options)
}

// SetChatPhoto sets a chat photo.
//
// https://core.telegram.org/bots/api#setchatphoto
func (b *Bot) SetChatPhoto(
	ctx context.Context,
	chatID ChatID,
	photo InputFile,
) (result APIResponse[bool], err error) {
	// essential options
	options := map[string]any{
		"chat_id": chatID,
		"photo":   photo,
	}

	return requestGeneric[bool](ctx, b, "setChatPhoto", options)
}

// DeleteChatPhoto deletes a chat photo.
//
// https://core.telegram.org/bots/api#deletechatphoto
func (b *Bot) DeleteChatPhoto(
	ctx context.Context,
	chatID ChatID,
) (result APIResponse[bool], err error) {
	// essential options
	options := map[string]any{
		"chat_id": chatID,
	}

	return requestGeneric[bool](ctx, b, "deleteChatPhoto", options)
}

// SetChatTitle sets a chat title.
//
// https://core.telegram.org/bots/api#setchattitle
func (b *Bot) SetChatTitle(
	ctx context.Context,
	chatID ChatID,
	title string,
) (result APIResponse[bool], err error) {
	// essential options
	options := map[string]any{
		"chat_id": chatID,
		"title":   title,
	}

	return requestGeneric[bool](ctx, b, "setChatTitle", options)
}

// SetChatDescription sets a chat description.
//
// https://core.telegram.org/bots/api#setchatdescription
func (b *Bot) SetChatDescription(
	ctx context.Context,
	chatID ChatID,
	description string,
) (result APIResponse[bool], err error) {
	// essential options
	options := map[string]any{
		"chat_id":     chatID,
		"description": description,
	}

	return requestGeneric[bool](ctx, b, "setChatDescription", options)
}

// PinChatMessage pins a chat message.
//
// https://core.telegram.org/bots/api#pinchatmessage
func (b *Bot) PinChatMessage(
	ctx context.Context,
	chatID ChatID,
	messageID int64,
	options OptionsPinChatMessage,
) (result APIResponse[bool], err error) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["chat_id"] = chatID
	options["message_id"] = messageID

	return requestGeneric[bool](ctx, b, "pinChatMessage", options)
}

// UnpinChatMessage unpins a chat message.
//
// https://core.telegram.org/bots/api#unpinchatmessage
func (b *Bot) UnpinChatMessage(
	ctx context.Context,
	chatID ChatID,
	options OptionsUnpinChatMessage,
) (result APIResponse[bool], err error) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["chat_id"] = chatID

	return requestGeneric[bool](ctx, b, "unpinChatMessage", options)
}

// UnpinAllChatMessages unpins all chat messages.
//
// https://core.telegram.org/bots/api#unpinallchatmessages
func (b *Bot) UnpinAllChatMessages(
	ctx context.Context,
	chatID ChatID,
) (result APIResponse[bool], err error) {
	// essential options
	options := map[string]any{
		"chat_id": chatID,
	}

	return requestGeneric[bool](ctx, b, "unpinAllChatMessages", options)
}

// GetChat gets a chat.
//
// https://core.telegram.org/bots/api#getchat
func (b *Bot) GetChat(
	ctx context.Context,
	chatID ChatID,
) (result APIResponse[ChatFullInfo], err error) {
	// essential options
	options := map[string]any{
		"chat_id": chatID,
	}

	return requestGeneric[ChatFullInfo](ctx, b, "getChat", options)
}

// GetChatAdministrators gets chat administrators.
//
// https://core.telegram.org/bots/api#getchatadministrators
func (b *Bot) GetChatAdministrators(
	ctx context.Context,
	chatID ChatID,
) (result APIResponse[[]ChatMember], err error) {
	// essential options
	options := map[string]any{
		"chat_id": chatID,
	}

	return requestGeneric[[]ChatMember](ctx, b, "getChatAdministrators", options)
}

// GetChatMemberCount gets chat members' count.
//
// https://core.telegram.org/bots/api#getchatmembercount
func (b *Bot) GetChatMemberCount(
	ctx context.Context,
	chatID ChatID,
) (result APIResponse[int], err error) {
	// essential options
	options := map[string]any{
		"chat_id": chatID,
	}

	return requestGeneric[int](ctx, b, "getChatMemberCount", options)
}

// GetChatMember gets a chat member.
//
// https://core.telegram.org/bots/api#getchatmember
func (b *Bot) GetChatMember(
	ctx context.Context,
	chatID ChatID,
	userID int64,
) (result APIResponse[ChatMember], err error) {
	// essential options
	options := map[string]any{
		"chat_id": chatID,
		"user_id": userID,
	}

	return requestGeneric[ChatMember](ctx, b, "getChatMember", options)
}

// SetChatStickerSet sets a chat sticker set.
//
// https://core.telegram.org/bots/api#setchatstickerset
func (b *Bot) SetChatStickerSet(
	ctx context.Context,
	chatID ChatID,
	stickerSetName string,
) (result APIResponse[bool], err error) {
	// essential options
	options := map[string]any{
		"chat_id":          chatID,
		"sticker_set_name": stickerSetName,
	}

	return requestGeneric[bool](ctx, b, "setChatStickerSet", options)
}

// DeleteChatStickerSet deletes a chat sticker set.
//
// https://core.telegram.org/bots/api#deletechatstickerset
func (b *Bot) DeleteChatStickerSet(
	ctx context.Context,
	chatID ChatID,
) (result APIResponse[bool], err error) {
	// essential options
	options := map[string]any{
		"chat_id": chatID,
	}

	return requestGeneric[bool](ctx, b, "deleteChatStickerSet", options)
}

// AnswerCallbackQuery answers a callback query.
//
// https://core.telegram.org/bots/api#answercallbackquery
func (b *Bot) AnswerCallbackQuery(
	ctx context.Context,
	callbackQueryID string,
	options OptionsAnswerCallbackQuery,
) (result APIResponse[bool], err error) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["callback_query_id"] = callbackQueryID

	return requestGeneric[bool](ctx, b, "answerCallbackQuery", options)
}

// GetMyCommands fetches commands of this bot.
//
// https://core.telegram.org/bots/api#getmycommands
func (b *Bot) GetMyCommands(
	ctx context.Context,
	options OptionsGetMyCommands,
) (result APIResponse[[]BotCommand], err error) {
	return requestGeneric[[]BotCommand](ctx, b, "getMyCommands", options)
}

// SetMyName changes the bot's name.
//
// https://core.telegram.org/bots/api#setmyname
func (b *Bot) SetMyName(
	ctx context.Context,
	name string,
	options OptionsSetMyName,
) (result APIResponse[bool], err error) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["name"] = name

	return requestGeneric[bool](ctx, b, "setMyName", options)
}

// GetMyName fetches the bot's name.
//
// https://core.telegram.org/bots/api#getmyname
func (b *Bot) GetMyName(
	ctx context.Context,
	options OptionsGetMyName,
) (result APIResponse[BotName], err error) {
	return requestGeneric[BotName](ctx, b, "getMyName", options)
}

// SetMyDescription sets the bot's description.
//
// https://core.telegram.org/bots/api#setmydescription
func (b *Bot) SetMyDescription(
	ctx context.Context,
	options OptionsSetMyDescription,
) (result APIResponse[bool], err error) {
	return requestGeneric[bool](ctx, b, "setMyDescription", options)
}

// GetMyDescription gets the bot's description.
//
// https://core.telegram.org/bots/api#setmydescription
func (b *Bot) GetMyDescription(
	ctx context.Context,
	options OptionsGetMyDescription,
) (result APIResponse[BotDescription], err error) {
	return requestGeneric[BotDescription](ctx, b, "getMyDescription", options)
}

// SetMyShortDescription sets the bot's short description.
//
// https://core.telegram.org/bots/api#setmyshortdescription
func (b *Bot) SetMyShortDescription(
	ctx context.Context,
	options OptionsSetMyShortDescription,
) (result APIResponse[bool], err error) {
	return requestGeneric[bool](ctx, b, "setMyShortDescription", options)
}

// GetMyShortDescription gets the bot's short description.
//
// https://core.telegram.org/bots/api#getmyshortdescription
func (b *Bot) GetMyShortDescription(
	ctx context.Context,
	options OptionsGetMyShortDescription,
) (result APIResponse[BotShortDescription], err error) {
	return requestGeneric[BotShortDescription](ctx, b, "getMyShortDescription", options)
}

// GetUserChatBoosts gets boosts of a user.
//
// https://core.telegram.org/bots/api#getuserchatboosts
func (b *Bot) GetUserChatBoosts(
	ctx context.Context,
	chatID ChatID,
	userID int64,
) (result APIResponse[UserChatBoosts], err error) {
	// essential params
	options := map[string]any{
		"chat_id": chatID,
		"user_id": userID,
	}

	return requestGeneric[UserChatBoosts](ctx, b, "getUserChatBoosts", options)
}

// GetBusinessConnection gets a business connection.
//
// https://core.telegram.org/bots/api#getbusinessconnection
func (b *Bot) GetBusinessConnection(
	ctx context.Context,
	businessConnectionID string,
) (result APIResponse[BusinessConnection], err error) {
	// essential params
	options := map[string]any{
		"business_connection_id": businessConnectionID,
	}

	return requestGeneric[BusinessConnection](ctx, b, "getBusinessConnection", options)
}

// SetMyCommands sets commands of this bot.
//
// https://core.telegram.org/bots/api#setmycommands
func (b *Bot) SetMyCommands(
	ctx context.Context,
	commands []BotCommand,
	options OptionsSetMyCommands,
) (result APIResponse[bool], err error) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["commands"] = commands

	return requestGeneric[bool](ctx, b, "setMyCommands", options)
}

// DeleteMyCommands deletes commands of this bot.
//
// https://core.telegram.org/bots/api#deletemycommands
func (b *Bot) DeleteMyCommands(
	ctx context.Context,
	options OptionsDeleteMyCommands,
) (result APIResponse[bool], err error) {
	return requestGeneric[bool](ctx, b, "deleteMyCommands", options)
}

// SetMyProfilePhoto sets the bot's profile photo.
//
// https://core.telegram.org/bots/api#setmyprofilephoto
func (b *Bot) SetMyProfilePhoto(
	ctx context.Context,
	photo InputProfilePhoto,
) (result APIResponse[bool], err error) {
	options := map[string]any{
		"photo": photo,
	}

	return requestGeneric[bool](ctx, b, "setMyProfilePhoto", options)
}

// RemoveMyProfilePhoto deletes the bot's profile photo.
//
// https://core.telegram.org/bots/api#removemyprofilephoto
func (b *Bot) RemoveMyProfilePhoto(
	ctx context.Context,
) (result APIResponse[bool], err error) {
	return requestGeneric[bool](ctx, b, "removeMyProfilePhoto", nil)
}

// SetChatMenuButton sets chat menu button.
//
// https://core.telegram.org/bots/api#setchatmenubutton
func (b *Bot) SetChatMenuButton(
	ctx context.Context,
	options OptionsSetChatMenuButton,
) (result APIResponse[bool], err error) {
	return requestGeneric[bool](ctx, b, "setChatMenuButton", options)
}

// GetChatMenuButton fetches current chat menu button.
//
// https://core.telegram.org/bots/api#getchatmenubutton
func (b *Bot) GetChatMenuButton(
	ctx context.Context,
	options OptionsGetChatMenuButton,
) (result APIResponse[MenuButton], err error) {
	return requestGeneric[MenuButton](ctx, b, "getChatMenuButton", options)
}

// SetMyDefaultAdministratorRights sets my default administrator rights.
//
// https://core.telegram.org/bots/api#setmydefaultadministratorrights
func (b *Bot) SetMyDefaultAdministratorRights(
	ctx context.Context,
	options OptionsSetMyDefaultAdministratorRights,
) (result APIResponse[bool], err error) {
	return requestGeneric[bool](ctx, b, "setMyDefaultAdministratorRights", options)
}

// GetMyDefaultAdministratorRights gets my default administrator rights.
//
// https://core.telegram.org/bots/api#getmydefaultadministratorrights
func (b *Bot) GetMyDefaultAdministratorRights(
	ctx context.Context,
	options OptionsGetMyDefaultAdministratorRights,
) (result APIResponse[bool], err error) {
	return requestGeneric[bool](ctx, b, "getMyDefaultAdministratorRights", options)
}

// Updating messages
//
// https://core.telegram.org/bots/api#updating-messages

// EditMessageText edits text of a message.
//
// https://core.telegram.org/bots/api#editmessagetext
func (b *Bot) EditMessageText(
	ctx context.Context,
	text string,
	options OptionsEditMessageText,
) (result APIResponseMessageOrBool, err error) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["text"] = text

	return b.requestMessageOrBool(ctx, "editMessageText", options)
}

// EditMessageCaption edits caption of a message.
//
// https://core.telegram.org/bots/api#editmessagecaption
func (b *Bot) EditMessageCaption(
	ctx context.Context,
	options OptionsEditMessageCaption,
) (result APIResponseMessageOrBool, err error) {
	if options == nil {
		options = map[string]any{}
	}

	return b.requestMessageOrBool(ctx, "editMessageCaption", options)
}

// EditMessageMedia edites a media message.
//
// https://core.telegram.org/bots/api#editmessagemedia
func (b *Bot) EditMessageMedia(
	ctx context.Context,
	media InputMedia,
	options OptionsEditMessageMedia,
) (result APIResponseMessageOrBool, err error) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["media"] = media

	return b.requestMessageOrBool(ctx, "editMessageMedia", options)
}

// EditMessageLiveLocation edits live location of a message.
//
// https://core.telegram.org/bots/api#editmessagelivelocation
func (b *Bot) EditMessageLiveLocation(
	ctx context.Context,
	latitude, longitude float32,
	options OptionsEditMessageLiveLocation,
) (result APIResponseMessageOrBool, err error) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["latitude"] = latitude
	options["longitude"] = longitude

	return b.requestMessageOrBool(ctx, "editMessageLiveLocation", options)
}

// StopMessageLiveLocation stops live location of a message.
//
// https://core.telegram.org/bots/api#stopmessagelivelocation
func (b *Bot) StopMessageLiveLocation(
	ctx context.Context,
	options OptionsStopMessageLiveLocation,
) (result APIResponseMessageOrBool, err error) {
	return b.requestMessageOrBool(ctx, "stopMessageLiveLocation", options)
}

// EditMessageChecklist edits check list of a message.
//
// https://core.telegram.org/bots/api#editmessagechecklist
func (b *Bot) EditMessageChecklist(
	ctx context.Context,
	businessConnectionID string,
	chatID int64,
	messageID int64,
	checklist InputChecklist,
	options OptionsEditMessageChecklist,
) (result APIResponse[Message], err error) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["business_connection_id"] = businessConnectionID
	options["chat_id"] = chatID
	options["message_id"] = messageID
	options["checklist"] = checklist

	return requestGeneric[Message](ctx, b, "editMessageChecklist", options)
}

// EditMessageReplyMarkup edits reply markup of a message.
//
// https://core.telegram.org/bots/api#editmessagereplymarkup
func (b *Bot) EditMessageReplyMarkup(
	ctx context.Context,
	options OptionsEditMessageReplyMarkup,
) (result APIResponseMessageOrBool, err error) {
	return b.requestMessageOrBool(ctx, "editMessageReplyMarkup", options)
}

// DeleteMessage deletes a message.
//
// https://core.telegram.org/bots/api#deletemessage
func (b *Bot) DeleteMessage(
	ctx context.Context,
	chatID ChatID,
	messageID int64,
) (result APIResponse[bool], err error) {
	return requestGeneric[bool](ctx, b, "deleteMessage", map[string]any{
		"chat_id":    chatID,
		"message_id": messageID,
	})
}

// DeleteMessages deletes messages.
//
// https://core.telegram.org/bots/api#deletemessages
func (b *Bot) DeleteMessages(
	ctx context.Context,
	chatID ChatID,
	messageIDs []int64,
) (result APIResponse[bool], err error) {
	return requestGeneric[bool](ctx, b, "deleteMessages", map[string]any{
		"chat_id":     chatID,
		"message_ids": messageIDs,
	})
}

// AnswerInlineQuery sends answers to an inline query.
//
// results = array of InlineQueryResultArticle, InlineQueryResultPhoto, InlineQueryResultGif, InlineQueryResultMpeg4Gif, or InlineQueryResultVideo.
//
// https://core.telegram.org/bots/api#answerinlinequery
func (b *Bot) AnswerInlineQuery(
	ctx context.Context,
	inlineQueryID string,
	results []any,
	options OptionsAnswerInlineQuery,
) (result APIResponse[bool], err error) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["inline_query_id"] = inlineQueryID
	options["results"] = results

	return requestGeneric[bool](ctx, b, "answerInlineQuery", options)
}

// SendInvoice sends an invoice.
//
// NOTE:
// - `providerToken`: Pass "" for payments in Telegram Stars.
// - `currency`: Pass "XTR" for payments in Telegram Stars.
//
// https://core.telegram.org/bots/api#sendinvoice
func (b *Bot) SendInvoice(
	ctx context.Context,
	chatID int64,
	title, description, payload, providerToken, currency string,
	prices []LabeledPrice,
	options OptionsSendInvoice,
) (result APIResponse[Message], err error) {
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

	return requestGeneric[Message](ctx, b, "sendInvoice", options)
}

// CreateInvoiceLink creates a link for an invoice.
//
// NOTE:
// - `providerToken`: Pass "" for payments in Telegram Stars.
// - `currency`: Pass "XTR" for payments in Telegram Stars.
//
// https://core.telegram.org/bots/api#createinvoicelink
func (b *Bot) CreateInvoiceLink(
	ctx context.Context,
	title, description, payload, currency string,
	prices []LabeledPrice,
	options OptionsCreateInvoiceLink,
) (result APIResponse[string], err error) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["title"] = title
	options["description"] = description
	options["payload"] = payload
	options["currency"] = currency
	options["prices"] = prices

	return requestGeneric[string](ctx, b, "createInvoiceLink", options)
}

// AnswerShippingQuery answers a shipping query.
//
// if ok is true, shippingOptions should be provided.
// otherwise, errorMessage should be provided.
//
// https://core.telegram.org/bots/api#answershippingquery
func (b *Bot) AnswerShippingQuery(
	ctx context.Context,
	shippingQueryID string,
	ok bool,
	shippingOptions []ShippingOption,
	errorMessage *string,
) (result APIResponse[bool], err error) {
	// essential options
	options := map[string]any{
		"shipping_query_id": shippingQueryID,
		"ok":                ok,
	}

	// optional params
	if ok {
		if len(shippingOptions) > 0 {
			options["shipping_options"] = shippingOptions
		}
	} else {
		if errorMessage != nil {
			options["error_message"] = *errorMessage
		}
	}

	return requestGeneric[bool](ctx, b, "answerShippingQuery", options)
}

// AnswerPreCheckoutQuery answers a pre-checkout query.
//
// https://core.telegram.org/bots/api#answerprecheckoutquery
func (b *Bot) AnswerPreCheckoutQuery(
	ctx context.Context,
	preCheckoutQueryID string,
	ok bool,
	errorMessage *string,
) (result APIResponse[bool], err error) {
	// essential options
	options := map[string]any{
		"pre_checkout_query_id": preCheckoutQueryID,
		"ok":                    ok,
	}

	// optional params
	if !ok {
		if errorMessage != nil {
			options["error_message"] = *errorMessage
		}
	}

	return requestGeneric[bool](ctx, b, "answerPreCheckoutQuery", options)
}

// GetMyStarBalance fetches the current balance of Telegram Stars.
//
// https://core.telegram.org/bots/api#getmystarbalance
func (b *Bot) GetMyStarBalance(
	ctx context.Context,
) (result APIResponse[StarAmount], err error) {
	return requestGeneric[StarAmount](ctx, b, "getMyStarBalance", nil)
}

// GetStarTransactions gets star transactions.
//
// https://core.telegram.org/bots/api#getstartransactions
func (b *Bot) GetStarTransactions(
	ctx context.Context,
	options OptionsGetStarTransactions,
) (result APIResponse[StarTransactions], err error) {
	return requestGeneric[StarTransactions](ctx, b, "getStarTransactions", options)
}

// RefundStarPayment refunds a successful payment in Telegram Stars.
//
// https://core.telegram.org/bots/api#refundstarpayment
func (b *Bot) RefundStarPayment(
	ctx context.Context,
	userID int64,
	telegramPaymentChargeID string,
) (result APIResponse[bool], err error) {
	// essential options
	options := map[string]any{
		"user_id":                    userID,
		"telegram_payment_charge_id": telegramPaymentChargeID,
	}

	return requestGeneric[bool](ctx, b, "refundStarPayment", options)
}

// EditUserStarSubscription allows the bot to cancel or re-enable extension of a subscription.
//
// https://core.telegram.org/bots/api#edituserstarsubscription
func (b *Bot) EditUserStarSubscription(
	ctx context.Context,
	userID int64,
	telegramPaymentChargeID string,
	isCanceled bool,
) (result APIResponse[bool], err error) {
	// essential options
	options := map[string]any{
		"user_id":                    userID,
		"telegram_payment_charge_id": telegramPaymentChargeID,
		"is_canceled":                isCanceled,
	}

	return requestGeneric[bool](ctx, b, "editUserStarSubscription", options)
}

// SendGame sends a game.
//
// https://core.telegram.org/bots/api#sendgame
func (b *Bot) SendGame(
	ctx context.Context,
	chatID ChatID,
	gameShortName string,
	options OptionsSendGame,
) (result APIResponse[Message], err error) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["chat_id"] = chatID
	options["game_short_name"] = gameShortName

	return requestGeneric[Message](ctx, b, "sendGame", options)
}

// SetGameScore sets score of a game.
//
// https://core.telegram.org/bots/api#setgamescore
func (b *Bot) SetGameScore(
	ctx context.Context,
	userID int64,
	score int,
	options OptionsSetGameScore,
) (result APIResponseMessageOrBool, err error) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["user_id"] = userID
	options["score"] = score

	return b.requestMessageOrBool(ctx, "setGameScore", options)
}

// GetGameHighScores gets high scores of a game.
//
// https://core.telegram.org/bots/api#getgamehighscores
func (b *Bot) GetGameHighScores(
	ctx context.Context,
	userID int64,
	options OptionsGetGameHighScores,
) (result APIResponse[[]GameHighScore], err error) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["user_id"] = userID

	return requestGeneric[[]GameHighScore](ctx, b, "getGameHighScores", options)
}

// AnswerWebAppQuery answers a web app's query.
//
// https://core.telegram.org/bots/api#answerwebappquery
func (b *Bot) AnswerWebAppQuery(
	ctx context.Context,
	webAppQueryID string,
	res InlineQueryResult,
) (result APIResponse[SentWebAppMessage], err error) {
	options := map[string]any{
		"web_app_query_id": webAppQueryID,
		"result":           res,
	}

	return requestGeneric[SentWebAppMessage](ctx, b, "answerWebAppQuery", options)
}

// SavePreparedInlineMessage stores a message that can be sent by a user of a Mini App.
//
// https://core.telegram.org/bots/api#savepreparedinlinemessage
func (b *Bot) SavePreparedInlineMessage(
	ctx context.Context,
	userID int64,
	result InlineQueryResult,
	options OptionsSavePreparedInlineMessage,
) (res APIResponse[PreparedInlineMessage], err error) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["user_id"] = userID
	options["result"] = result

	return requestGeneric[PreparedInlineMessage](ctx, b, "savePreparedInlineMessage", options)
}

// CreateForumTopic creates a topic in a forum supergroup chat or a private chat with a user.
//
// https://core.telegram.org/bots/api#createforumtopic
func (b *Bot) CreateForumTopic(
	ctx context.Context,
	chatID ChatID,
	name string,
	options OptionsCreateForumTopic,
) (result APIResponse[ForumTopic], err error) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["chat_id"] = chatID
	options["name"] = name

	return requestGeneric[ForumTopic](ctx, b, "createForumTopic", options)
}

// EditForumTopic edits a topic in a forum supergroup chat or a private chat with a user.
//
// https://core.telegram.org/bots/api#editforumtopic
func (b *Bot) EditForumTopic(
	ctx context.Context,
	chatID ChatID,
	messageThreadID int64,
	options OptionsEditForumTopic,
) (result APIResponse[bool], err error) {
	if options == nil {
		options = map[string]any{}
	}

	// essential params
	options["chat_id"] = chatID
	options["message_thread_id"] = messageThreadID

	return requestGeneric[bool](ctx, b, "editForumTopic", options)
}

// CloseForumTopic closes an open topic in a forum supergroup chat.
//
// https://core.telegram.org/bots/api#closeforumtopic
func (b *Bot) CloseForumTopic(
	ctx context.Context,
	chatID ChatID,
	messageThreadID int64,
) (result APIResponse[bool], err error) {
	options := map[string]any{
		"chat_id":           chatID,
		"message_thread_id": messageThreadID,
	}

	return requestGeneric[bool](ctx, b, "closeForumTopic", options)
}

// ReopenForumTopic reopens a closed topic in a forum supergroup chat.
//
// https://core.telegram.org/bots/api#reopenforumtopic
func (b *Bot) ReopenForumTopic(
	ctx context.Context,
	chatID ChatID,
	messageThreadID int64,
) (result APIResponse[bool], err error) {
	options := map[string]any{
		"chat_id":           chatID,
		"message_thread_id": messageThreadID,
	}

	return requestGeneric[bool](ctx, b, "reopenForumTopic", options)
}

// DeleteForumTopic deletes a forum topic along with all its messages in a forum supergroup chat or a private chat with a user.
//
// https://core.telegram.org/bots/api#deleteforumtopic
func (b *Bot) DeleteForumTopic(
	ctx context.Context,
	chatID ChatID,
	messageThreadID int64,
) (result APIResponse[bool], err error) {
	options := map[string]any{
		"chat_id":           chatID,
		"message_thread_id": messageThreadID,
	}

	return requestGeneric[bool](ctx, b, "deleteForumTopic", options)
}

// UnpinAllForumTopicMessages clears the list of pinned messages in a forum topic in a forum supergroup chat or a private chat with a user.
//
// https://core.telegram.org/bots/api#unpinallforumtopicmessages
func (b *Bot) UnpinAllForumTopicMessages(
	ctx context.Context,
	chatID ChatID,
	messageThreadID int64,
) (result APIResponse[bool], err error) {
	options := map[string]any{
		"chat_id":           chatID,
		"message_thread_id": messageThreadID,
	}

	return requestGeneric[bool](ctx, b, "unpinAllForumTopicMessages", options)
}

// EditGeneralForumTopic edits the name of the 'General' topic in a forum supergroup chat.
//
// https://core.telegram.org/bots/api#editgeneralforumtopic
func (b *Bot) EditGeneralForumTopic(
	ctx context.Context,
	chatID ChatID,
	name string,
) (result APIResponse[bool], err error) {
	options := map[string]any{
		"chat_id": chatID,
		"name":    name,
	}

	return requestGeneric[bool](ctx, b, "editGeneralForumTopic", options)
}

// CloseGeneralForumTopic closes an open 'General' topic in a forum supergroup chat.
//
// https://core.telegram.org/bots/api#closegeneralforumtopic
func (b *Bot) CloseGeneralForumTopic(
	ctx context.Context,
	chatID ChatID,
) (result APIResponse[bool], err error) {
	options := map[string]any{
		"chat_id": chatID,
	}

	return requestGeneric[bool](ctx, b, "closeGeneralForumTopic", options)
}

// ReopenGeneralForumTopic reopens a closed 'General' topic in a forum supergroup chat.
//
// https://core.telegram.org/bots/api#reopengeneralforumtopic
func (b *Bot) ReopenGeneralForumTopic(
	ctx context.Context,
	chatID ChatID,
) (result APIResponse[bool], err error) {
	options := map[string]any{
		"chat_id": chatID,
	}

	return requestGeneric[bool](ctx, b, "reopenGeneralForumTopic", options)
}

// HideGeneralForumTopic hides the 'General' topic in a forum supergroup chat.
//
// https://core.telegram.org/bots/api#hidegeneralforumtopic
func (b *Bot) HideGeneralForumTopic(
	ctx context.Context,
	chatID ChatID,
) (result APIResponse[bool], err error) {
	options := map[string]any{
		"chat_id": chatID,
	}

	return requestGeneric[bool](ctx, b, "hideGeneralForumTopic", options)
}

// UnhideGeneralForumTopic unhides the 'General' topic in a forum supergroup chat.
//
// https://core.telegram.org/bots/api#unhidegeneralforumtopic
func (b *Bot) UnhideGeneralForumTopic(
	ctx context.Context,
	chatID ChatID,
) (result APIResponse[bool], err error) {
	options := map[string]any{
		"chat_id": chatID,
	}

	return requestGeneric[bool](ctx, b, "unhideGeneralForumTopic", options)
}

// UnpinAllGeneralForumTopicMessages clears the list of pinned messages in a General forum topic.
//
// https://core.telegram.org/bots/api#unpinallgeneralforumtopicmessages
func (b *Bot) UnpinAllGeneralForumTopicMessages(
	ctx context.Context,
	chatID ChatID,
) (result APIResponse[bool], err error) {
	options := map[string]any{
		"chat_id": chatID,
	}

	return requestGeneric[bool](ctx, b, "unpinAllGeneralForumTopicMessages", options)
}

// GetForumTopicIconStickers fetches custom emoji stickers which can be used as a forum topic icon by any user.
//
// https://core.telegram.org/bots/api#getforumtopiciconstickers
func (b *Bot) GetForumTopicIconStickers(
	ctx context.Context,
) (result APIResponse[[]Sticker], err error) {
	return requestGeneric[[]Sticker](ctx, b, "getForumTopicIconStickers", nil)
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
		case InputProfilePhoto:
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
		b.error("parameter %[1]T (%+[1]v) could not be cast to string value", param)
	case InputProfilePhoto:
		if val.Photo != nil || val.Animation != nil {
			if marshalled, err := json.Marshal(val); err == nil {
				return string(marshalled), true
			}
		}
		b.error("parameter %[1]T (%+[1]v) could not be cast to string value", param)
	default: // fallback: encode to JSON string
		json, err := json.Marshal(param)
		if err == nil {
			return string(json), true
		}
		b.error("parameter %[1]T (%+[1]v) could not be encoded as json: %s", param, err)
	}

	return "", false
}

// Send request to API server and return the response as bytes(synchronously).
//
// NOTE: If *os.File is included in the params, it will be closed automatically by this function.
func (b *Bot) request(
	ctx context.Context,
	method string,
	params map[string]any,
) (resp []byte, err error) {
	apiURL := fmt.Sprintf("%s%s/%s", apiBaseURL, b.token, method)

	b.verbose("sending request to api url: %s, params: %#v", apiURL, params)

	if checkIfFileParamExists(params) {
		// multipart form data
		resp, err = b.requestMultipartFormData(ctx, apiURL, params)
	} else {
		// www-form urlencoded
		resp, err = b.requestURLEncodedFormData(ctx, apiURL, params)
	}

	if err == nil {
		return resp, nil
	}

	return []byte{}, fmt.Errorf("%s", b.redact(err.Error()))
}

// request multipart form data
func (b *Bot) requestMultipartFormData(
	ctx context.Context,
	apiURL string,
	params map[string]any,
) (resp []byte, err error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	for k, v := range params {
		switch val := v.(type) {
		case *os.File:
			defer func() { _ = val.Close() }()

			var part io.Writer
			part, err = writer.CreateFormFile(k, val.Name())
			if err == nil {
				if _, err = io.Copy(part, val); err != nil {
					b.error("could not write to multipart: '%[1]s'", k)
				}
			} else {
				b.error("could not create form file for parameter '%[1]s' (%[2]T)", k, val)
			}
		case []byte:
			filename := fmt.Sprintf("%s.%s", k, getExtension(val))

			var part io.Writer
			part, err = writer.CreateFormFile(k, filename)
			if err == nil {
				if _, err = io.Copy(part, bytes.NewReader(val)); err != nil {
					b.error("could not write to multipart: '%[1]s'", k)
				}
			} else {
				b.error("could not create form file for parameter '%[1]s' (%[2]T)", k, val)
			}
		case InputFile:
			if val.Filepath != nil {
				var file *os.File
				if file, err = os.Open(*val.Filepath); err == nil {
					defer func() { _ = file.Close() }()

					var part io.Writer
					part, err = writer.CreateFormFile(k, file.Name())
					if err == nil {
						if _, err = io.Copy(part, file); err != nil {
							b.error("could not write to multipart: '%[1]s'", k)
						}
					} else {
						b.error("could not create form file for parameter '%[1]s' (%[2]T)", k, val)
					}
				} else {
					b.error("parameter '%[1]s' (%[2]T) could not be read from file: %[3]s", k, val, err.Error())
				}
			} else if len(val.Bytes) > 0 {
				filename := fmt.Sprintf("%s.%s", k, getExtension(val.Bytes))
				var part io.Writer
				part, err = writer.CreateFormFile(k, filename)
				if err == nil {
					if _, err = io.Copy(part, bytes.NewReader(val.Bytes)); err != nil {
						b.error("could not write %[2]T to multipart: '%[1]s'", k, val)
					}
				} else {
					b.error("could not create form file for parameter '%[1]s' (%[2]T)", k, val)
				}
			} else {
				b.error("ignoring invalid %[2]T parameter '%[1]s' without filepath nor bytes", k, val)
			}
		case InputProfilePhoto: // https://core.telegram.org/bots/api#inputprofilephoto
			if val.Filepath != nil {
				var file *os.File
				if file, err = os.Open(*val.Filepath); err == nil {
					defer func() { _ = file.Close() }()

					filename := filepath.Base(file.Name())

					var part io.Writer
					part, err = writer.CreateFormFile(filename, filename)
					if err == nil {
						if _, err = io.Copy(part, file); err == nil {
							attachedFilename := "attach://" + filename

							switch val.Type {
							case InputProfilePhotoStatic:
								val.Photo = &attachedFilename
							case InputProfilePhotoAnimated:
								val.Animation = &attachedFilename
							default:
								b.error("invalid %[1]T type: %[2]s", val, val.Type)
							}

							// write InputProfilePhoto 'value' to body
							if strValue, ok := b.paramToString(val); ok {
								if err := writer.WriteField(k, strValue); err != nil {
									b.error("failed to write field with key: '%[1]s', value: %[2]s (%[3]s)", k, strValue, err)
								}
							} else {
								b.error("invalid %[2]T parameter '%[1]s'", k, val)
							}
						} else {
							b.error("could not write to multipart: '%[1]s'", k)
						}
					} else {
						b.error("could not create form file for parameter '%[1]s' (%[2]T)", k, val)
					}
				} else {
					b.error("parameter '%[1]s' (%[2]T) could not be read from file: %[3]s", k, val, err.Error())
				}
			} else if len(val.Bytes) > 0 {
				filename := fmt.Sprintf("%s.%s", k, getExtension(val.Bytes))

				var part io.Writer
				part, err = writer.CreateFormFile(filename, filename)
				if err == nil {
					if _, err = io.Copy(part, bytes.NewReader(val.Bytes)); err == nil {
						// set attached filename
						attachedFilename := "attach://" + filename
						switch val.Type {
						case InputProfilePhotoStatic:
							val.Photo = &attachedFilename
						case InputProfilePhotoAnimated:
							val.Animation = &attachedFilename
						default:
							b.error("invalid %[1]T type: %[2]s", val, val.Type)
						}

						// write InputProfilePhoto 'value' to body
						if strValue, ok := b.paramToString(val); ok {
							if err := writer.WriteField(k, strValue); err != nil {
								b.error("failed to write field with key: '%[1]s', value: %[2]s (%[3]s)", k, strValue, err)
							}
						} else {
							b.error("invalid %[2]T parameter '%[1]s'", k, val)
						}
					} else {
						b.error("could not write %[2]T to multipart: '%[1]s'", k, val)
					}
				} else {
					b.error("could not create form file for parameter '%[1]s' (%[2]T)", k, val)
				}
			} else {
				b.error("ignoring invalid %[2]T parameter '%[1]s' without filepath nor bytes", k, val)
			}
		default:
			if strValue, ok := b.paramToString(val); ok {
				if err := writer.WriteField(k, strValue); err != nil {
					b.error("failed to write filed with key: '%s', value: %s (%s)", k, strValue, err)
				}
			}
		}
	}

	if err = writer.Close(); err != nil {
		b.error("error while closing writer (%s)", err)
	}

	var req *http.Request
	req, err = http.NewRequestWithContext(ctx, "POST", apiURL, body)
	if err == nil {
		req.Header.Add("Content-Type", writer.FormDataContentType()) // due to file parameter
		req.Close = true

		// dump request
		if b.DumpHTTP {
			// NOTE: include `body` only when verbose mode
			if dumped, err := httputil.DumpRequest(req, b.Verbose); err == nil {
				slog.Debug(">>> dumping HTTP request",
					"with_body", b.Verbose,
					"dumped", b.redact(string(dumped)),
				)
			}
		}

		var resp *http.Response
		resp, err = b.httpClient.Do(req)

		if resp != nil { // XXX - in case of http redirect
			defer func() { _ = resp.Body.Close() }()
		}

		if err == nil {
			// dump response
			if b.DumpHTTP {
				// NOTE: include `body` only when verbose mode
				if dumped, err := httputil.DumpResponse(resp, b.Verbose); err == nil {
					slog.Debug(">>> dumping HTTP response",
						"with_body", b.Verbose,
						"dumped", b.redact(string(dumped)),
					)
				}
			}

			// FIXXX: check http status code here
			var bytes []byte
			bytes, err = io.ReadAll(resp.Body)
			if err == nil {
				return bytes, nil
			}

			err = fmt.Errorf("response read error: %w", err)
		} else {
			err = fmt.Errorf("request error: %w", err)
		}
	} else {
		err = fmt.Errorf("building request error: %w", err)
	}

	return []byte{}, err
}

// request urlencoded form data
func (b *Bot) requestURLEncodedFormData(
	ctx context.Context,
	apiURL string,
	params map[string]any,
) (resp []byte, err error) {
	paramValues := url.Values{}
	for key, value := range params {
		if strValue, ok := b.paramToString(value); ok {
			paramValues[key] = []string{strValue}
		}
	}
	encoded := paramValues.Encode()

	var req *http.Request
	req, err = http.NewRequestWithContext(ctx, "POST", apiURL, bytes.NewBufferString(encoded))
	if err == nil {
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		req.Header.Add("Content-Length", strconv.Itoa(len(encoded)))
		req.Close = true

		var resp *http.Response
		resp, err = b.httpClient.Do(req)

		if resp != nil { // XXX - in case of redirect
			defer func() { _ = resp.Body.Close() }()
		}

		if err == nil {
			// FIXXX: check http status code here
			var bytes []byte
			bytes, err = io.ReadAll(resp.Body)
			if err == nil {
				return bytes, nil
			}

			err = fmt.Errorf("response read error: %w", err)
		} else {
			err = fmt.Errorf("request error: %w", err)
		}
	} else {
		err = fmt.Errorf("building request error: %w", err)
	}

	return []byte{}, err
}

// Send request for APIResponseMessageOrBool and fetch its result.
func (b *Bot) requestMessageOrBool(
	ctx context.Context,
	method string,
	params map[string]any,
) (result APIResponseMessageOrBool, err error) {
	var bytes []byte
	if bytes, err = b.request(ctx, method, params); err == nil {
		// try APIResponseMessage type,
		var resMessage APIResponse[Message]
		err = json.Unmarshal(bytes, &resMessage)
		if err == nil {
			res := APIResponseMessageOrBool{
				OK:            resMessage.OK,
				Description:   resMessage.Description,
				ResultMessage: resMessage.Result,
			}
			if !res.OK && res.Description != nil {
				err = strToErr(*res.Description)
			}
			return res, err
		}

		// then try APIResponseBool type,
		var resBool APIResponse[bool]
		err = json.Unmarshal(bytes, &resBool)
		if err == nil {
			res := APIResponseMessageOrBool{
				OK:          resBool.OK,
				Description: resBool.Description,
				ResultBool:  resBool.Result,
			}
			if !res.OK && res.Description != nil {
				err = strToErr(*res.Description)
			}
			return res, err
		}

		err = fmt.Errorf("%s failed to parse json: not in Message nor bool type (%s)", method, string(bytes))
	} else {
		err = fmt.Errorf("%s failed with error: %w", method, err)
	}

	errStr := err.Error()

	return APIResponseMessageOrBool{
		OK:          false,
		Description: &errStr,
	}, err
}

// Send request for APIResponse[T] and fetch its result.
func requestGeneric[T any](
	ctx context.Context,
	b *Bot,
	method string,
	params map[string]any,
) (result APIResponse[T], err error) {
	var bytes []byte
	if bytes, err = b.request(ctx, method, params); err == nil {
		var res APIResponse[T]
		err = json.Unmarshal(bytes, &res)
		if err == nil {
			if !res.OK && res.Description != nil {
				err = strToErr(*res.Description)
			}
			return res, err
		}
		err = fmt.Errorf("%s failed to parse json: %w (%s)", method, err, string(bytes))
	} else {
		err = fmt.Errorf("%s failed with error: %w", method, strToErr(err.Error()))
	}

	errStr := err.Error()

	return APIResponse[T]{
		OK:          false,
		Description: &errStr,
	}, err
}

// Handle Webhook request.
func (b *Bot) handleWebhook(writer http.ResponseWriter, req *http.Request) {
	defer func() { _ = req.Body.Close() }()

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
