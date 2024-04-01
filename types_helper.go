package telegrambot

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io"
)

////////////////////////////////
// Helper functions for InlineQueryResult

// Generate a random UUID according to RFC-4122
//
// http://play.golang.org/p/4FkNSiUDMg
func newUUID() (string, error) {
	uuid := make([]byte, 16)
	n, err := io.ReadFull(rand.Reader, uuid)
	if n != len(uuid) || err != nil {
		return "", err
	}

	// variant bits; see section 4.1.1
	uuid[8] = uuid[8]&^0xc0 | 0x80
	// version 4 (pseudo-random); see section 4.1.3
	uuid[6] = uuid[6]&^0xf0 | 0x40

	return fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:]), nil
}

// NewInlineQueryResultArticle is a helper function for generating a new InlineQueryResultArticle
//
// https://core.telegram.org/bots/api#inlinequeryresultarticle
func NewInlineQueryResultArticle(title, messageText, description string) (newArticle *InlineQueryResultArticle, generatedID *string) {
	if id, err := newUUID(); err == nil {
		return &InlineQueryResultArticle{
			InlineQueryResult: InlineQueryResult{
				Type: InlineQueryResultTypeArticle,
				ID:   id,
			},
			Title: title,
			InputMessageContent: InputTextMessageContent{
				MessageText: messageText,
			},
			Description: &description,
		}, &id
	}

	return &InlineQueryResultArticle{}, nil
}

// NewInlineQueryResultPhoto is a helper function for generating a new InlineQueryResultPhoto
//
// Photo must be in jpeg format, < 5MB.
//
// https://core.telegram.org/bots/api#inlinequeryresultphoto
func NewInlineQueryResultPhoto(photoURL, thumbnailURL string) (newPhoto *InlineQueryResultPhoto, generatedID *string) {
	if id, err := newUUID(); err == nil {
		return &InlineQueryResultPhoto{
			InlineQueryResult: InlineQueryResult{
				Type: InlineQueryResultTypePhoto,
				ID:   id,
			},
			PhotoURL:     photoURL,
			ThumbnailURL: thumbnailURL,
		}, &id
	}

	return &InlineQueryResultPhoto{}, nil
}

// NewInlineQueryResultGif is a helper function for generating a new InlineQueryResultGif
//
// Gif must be in gif format, < 1MB.
//
// https://core.telegram.org/bots/api#inlinequeryresultgif
func NewInlineQueryResultGif(gifURL, thumbnailURL string) (newGif *InlineQueryResultGif, generatedID *string) {
	if id, err := newUUID(); err == nil {
		return &InlineQueryResultGif{
			InlineQueryResult: InlineQueryResult{
				Type: InlineQueryResultTypeGif,
				ID:   id,
			},
			GifURL:       gifURL,
			ThumbnailURL: thumbnailURL,
		}, &id
	}

	return &InlineQueryResultGif{}, nil
}

// NewInlineQueryResultMpeg4Gif is a helper function for generating a new InlineQueryResultMpeg4Gif
//
// Mpeg4 must be in H.264/MPEG-4 AVC video(wihout sound) format, < 1MB.
//
// https://core.telegram.org/bots/api#inlinequeryresultmpeg4gif
func NewInlineQueryResultMpeg4Gif(mpeg4URL, thumbnailURL string) (newMpeg4Gif *InlineQueryResultMpeg4Gif, generatedID *string) {
	if id, err := newUUID(); err == nil {
		return &InlineQueryResultMpeg4Gif{
			InlineQueryResult: InlineQueryResult{
				Type: InlineQueryResultTypeMpeg4Gif,
				ID:   id,
			},
			Mpeg4URL:     mpeg4URL,
			ThumbnailURL: thumbnailURL,
		}, &id
	}

	return &InlineQueryResultMpeg4Gif{}, nil
}

// NewInlineQueryResultVideo is a helper function for generating a new InlineQueryResultVideo
//
// https://core.telegram.org/bots/api#inlinequeryresultvideo
func NewInlineQueryResultVideo(videoURL, thumbnailURL, title string, mimeType VideoMimeType) (newVideo *InlineQueryResultVideo, generatedID *string) {
	if id, err := newUUID(); err == nil {
		return &InlineQueryResultVideo{
			InlineQueryResult: InlineQueryResult{
				Type: InlineQueryResultTypeVideo,
				ID:   id,
			},
			VideoURL:     videoURL,
			MimeType:     mimeType,
			ThumbnailURL: thumbnailURL,
			Title:        title,
		}, &id
	}

	return &InlineQueryResultVideo{}, nil
}

// NewInlineQueryResultAudio is a helper function for generating a new InlineQueryResultAudio
//
// https://core.telegram.org/bots/api#inlinequeryresultaudio
func NewInlineQueryResultAudio(audioURL, title string) (newAudio *InlineQueryResultAudio, generatedID *string) {
	if id, err := newUUID(); err == nil {
		return &InlineQueryResultAudio{
			InlineQueryResult: InlineQueryResult{
				Type: InlineQueryResultTypeAudio,
				ID:   id,
			},
			AudioURL: audioURL,
			Title:    title,
		}, &id
	}

	return &InlineQueryResultAudio{}, nil
}

// NewInlineQueryResultVoice is a helper function for generating a new InlineQueryResultVoice
//
// https://core.telegram.org/bots/api#inlinequeryresultvoice
func NewInlineQueryResultVoice(voiceURL, title string) (newVoice *InlineQueryResultVoice, generatedID *string) {
	if id, err := newUUID(); err == nil {
		return &InlineQueryResultVoice{
			InlineQueryResult: InlineQueryResult{
				Type: InlineQueryResultTypeVoice,
				ID:   id,
			},
			VoiceURL: voiceURL,
			Title:    title,
		}, &id
	}

	return &InlineQueryResultVoice{}, nil
}

// NewInlineQueryResultDocument is a helper function for generating a new InlineQueryResultDocument
//
// https://core.telegram.org/bots/api#inlinequeryresultdocument
func NewInlineQueryResultDocument(documentURL, title string, mimeType DocumentMimeType) (newDocument *InlineQueryResultDocument, generatedID *string) {
	if id, err := newUUID(); err == nil {
		return &InlineQueryResultDocument{
			InlineQueryResult: InlineQueryResult{
				Type: InlineQueryResultTypeDocument,
				ID:   id,
			},
			Title:       title,
			DocumentURL: documentURL,
			MimeType:    mimeType,
		}, &id
	}

	return &InlineQueryResultDocument{}, nil
}

// NewInlineQueryResultLocation is a helper function for generating a new InlineQueryResultLocation
//
// https://core.telegram.org/bots/api#inlinequeryresultlocation
func NewInlineQueryResultLocation(latitude, longitude float32, title string) (newLocation *InlineQueryResultLocation, generatedID *string) {
	if id, err := newUUID(); err == nil {
		return &InlineQueryResultLocation{
			InlineQueryResult: InlineQueryResult{
				Type: InlineQueryResultTypeLocation,
				ID:   id,
			},
			Latitude:  latitude,
			Longitude: longitude,
			Title:     title,
		}, &id
	}

	return &InlineQueryResultLocation{}, nil
}

// NewInlineQueryResultVenue is a helper function for generating a new InlineQueryResultVenue
//
// https://core.telegram.org/bots/api#inlinequeryresultvenue
func NewInlineQueryResultVenue(latitude, longitude float32, title, address string) (newVenue *InlineQueryResultVenue, generatedID *string) {
	if id, err := newUUID(); err == nil {
		return &InlineQueryResultVenue{
			InlineQueryResult: InlineQueryResult{
				Type: InlineQueryResultTypeVenue,
				ID:   id,
			},
			Latitude:  latitude,
			Longitude: longitude,
			Title:     title,
			Address:   address,
		}, &id
	}

	return &InlineQueryResultVenue{}, nil
}

// NewInlineQueryResultContact is a helper function for generating a new InlineQueryResultContact
//
// https://core.telegram.org/bots/api#inlinequeryresultcontact
func NewInlineQueryResultContact(phoneNumber, firstName string) (newContact *InlineQueryResultContact, generatedID *string) {
	if id, err := newUUID(); err == nil {
		return &InlineQueryResultContact{
			InlineQueryResult: InlineQueryResult{
				Type: InlineQueryResultTypeContact,
				ID:   id,
			},
			PhoneNumber: phoneNumber,
			FirstName:   firstName,
		}, &id
	}

	return &InlineQueryResultContact{}, nil
}

// NewInlineQueryResultCachedPhoto is a helper function for generating a new InlineQueryResultCachedPhoto
//
// https://core.telegram.org/bots/api#inlinequeryresultcachedphoto
func NewInlineQueryResultCachedPhoto(photoFileID string) (newPhoto *InlineQueryResultCachedPhoto, generatedID *string) {
	if id, err := newUUID(); err == nil {
		return &InlineQueryResultCachedPhoto{
			InlineQueryResult: InlineQueryResult{
				Type: InlineQueryResultTypePhoto,
				ID:   id,
			},
			PhotoFileID: photoFileID,
		}, &id
	}

	return &InlineQueryResultCachedPhoto{}, nil
}

// NewInlineQueryResultCachedGif is a helper function for generating a new InlineQueryResultCachedGif
//
// https://core.telegram.org/bots/api#inlinequeryresultcachedgif
func NewInlineQueryResultCachedGif(gifFileID string) (newGif *InlineQueryResultCachedGif, generatedID *string) {
	if id, err := newUUID(); err == nil {
		return &InlineQueryResultCachedGif{
			InlineQueryResult: InlineQueryResult{
				Type: InlineQueryResultTypeGif,
				ID:   id,
			},
			GifFileID: gifFileID,
		}, &id
	}

	return &InlineQueryResultCachedGif{}, nil
}

// NewInlineQueryResultCachedMpeg4Gif is a helper function for generating a new InlineQueryResultCachedMpeg4Gif
//
// https://core.telegram.org/bots/api#inlinequeryresultcachedmpeg4gif
func NewInlineQueryResultCachedMpeg4Gif(mpeg4FileID string) (newMpeg4Gif *InlineQueryResultCachedMpeg4Gif, generatedID *string) {
	if id, err := newUUID(); err == nil {
		return &InlineQueryResultCachedMpeg4Gif{
			InlineQueryResult: InlineQueryResult{
				Type: InlineQueryResultTypeMpeg4Gif,
				ID:   id,
			},
			Mpeg4FileID: mpeg4FileID,
		}, &id
	}

	return &InlineQueryResultCachedMpeg4Gif{}, nil
}

// NewInlineQueryResultCachedSticker is a helper function for generating a new InlineQueryResultCachedSticker
//
// https://core.telegram.org/bots/api#inlinequeryresultcachedsticker
func NewInlineQueryResultCachedSticker(stickerFileID string) (newSticker *InlineQueryResultCachedSticker, generatedID *string) {
	if id, err := newUUID(); err == nil {
		return &InlineQueryResultCachedSticker{
			InlineQueryResult: InlineQueryResult{
				Type: InlineQueryResultTypeSticker,
				ID:   id,
			},
			StickerFileID: stickerFileID,
		}, &id
	}

	return &InlineQueryResultCachedSticker{}, nil
}

// NewInlineQueryResultCachedDocument is a helper function for generating a new InlineQueryResultCachedDocument
//
// https://core.telegram.org/bots/api#inlinequeryresultcacheddocument
func NewInlineQueryResultCachedDocument(title, documentFileID string) (newDocument *InlineQueryResultCachedDocument, generatedID *string) {
	if id, err := newUUID(); err == nil {
		return &InlineQueryResultCachedDocument{
			InlineQueryResult: InlineQueryResult{
				Type: InlineQueryResultTypeDocument,
				ID:   id,
			},
			Title:          title,
			DocumentFileID: documentFileID,
		}, &id
	}

	return &InlineQueryResultCachedDocument{}, nil
}

// NewInlineQueryResultCachedVideo is a helper function for generating a new InlineQueryResultCachedVideo
//
// https://core.telegram.org/bots/api#inlinequeryresultcachedvideo
func NewInlineQueryResultCachedVideo(title, videoFileID string) (newVideo *InlineQueryResultCachedVideo, generatedID *string) {
	if id, err := newUUID(); err == nil {
		return &InlineQueryResultCachedVideo{
			InlineQueryResult: InlineQueryResult{
				Type: InlineQueryResultTypeVideo,
				ID:   id,
			},
			Title:       title,
			VideoFileID: videoFileID,
		}, &id
	}

	return &InlineQueryResultCachedVideo{}, nil
}

// NewInlineQueryResultCachedVoice is a helper function for generating a new InlineQueryResultCachedVoice
//
// https://core.telegram.org/bots/api#inlinequeryresultcachedvoice
func NewInlineQueryResultCachedVoice(title, voiceFileID string) (newVoice *InlineQueryResultCachedVoice, generatedID *string) {
	if id, err := newUUID(); err == nil {
		return &InlineQueryResultCachedVoice{
			InlineQueryResult: InlineQueryResult{
				Type: InlineQueryResultTypeVoice,
				ID:   id,
			},
			Title:       title,
			VoiceFileID: voiceFileID,
		}, &id
	}

	return &InlineQueryResultCachedVoice{}, nil
}

// NewInlineQueryResultCachedAudio is a helper function for generating a new InlineQueryResultCachedAudio
//
// https://core.telegram.org/bots/api#inlinequeryresultcachedaudio
func NewInlineQueryResultCachedAudio(audioFileID string) (newAudio *InlineQueryResultCachedAudio, generatedID *string) {
	if id, err := newUUID(); err == nil {
		return &InlineQueryResultCachedAudio{
			InlineQueryResult: InlineQueryResult{
				Type: InlineQueryResultTypeAudio,
				ID:   id,
			},
			AudioFileID: audioFileID,
		}, &id
	}

	return &InlineQueryResultCachedAudio{}, nil
}

////////////////////////////////
// Helper functions for Update
//

// String function for Update
func (u Update) String() string {
	return structToString(u)
}

// HasMessage checks if Update has Message.
func (u *Update) HasMessage() bool {
	return u.Message != nil
}

// HasEditedMessage checks if Update has EditedMessage.
func (u *Update) HasEditedMessage() bool {
	return u.EditedMessage != nil
}

// HasChannelPost checks if Update has ChannelPost.
func (u *Update) HasChannelPost() bool {
	return u.ChannelPost != nil
}

// HasEditedChannelPost checks if Update has EditedChannelPost.
func (u *Update) HasEditedChannelPost() bool {
	return u.EditedChannelPost != nil
}

// HasInlineQuery checks if Update has InlineQuery
func (u *Update) HasInlineQuery() bool {
	return u.InlineQuery != nil
}

// HasChosenInlineResult checks if Update has ChosenInlineResult
func (u *Update) HasChosenInlineResult() bool {
	return u.ChosenInlineResult != nil
}

// HasCallbackQuery checks if Update has CallbackQuery
func (u *Update) HasCallbackQuery() bool {
	return u.CallbackQuery != nil
}

// HasShippingQuery checks if Update has ShippingQuery
func (u *Update) HasShippingQuery() bool {
	return u.ShippingQuery != nil
}

// HasPreCheckoutQuery checks if Update has PreCheckoutQuery
func (u *Update) HasPreCheckoutQuery() bool {
	return u.PreCheckoutQuery != nil
}

// HasPoll checks if Update has Poll
func (u *Update) HasPoll() bool {
	return u.Poll != nil
}

// HasPollAnswer checks if Update has PollAnswer.
func (u *Update) HasPollAnswer() bool {
	return u.PollAnswer != nil
}

// HasMyChatMember checks if Update has MyChatMember.
func (u *Update) HasMyChatMember() bool {
	return u.MyChatMember != nil
}

// HasChatMember checks if Update has ChatMember.
func (u *Update) HasChatMember() bool {
	return u.ChatMember != nil
}

// HasChatJoinRequest checks if Update has ChatJoinRequest.
func (u *Update) HasChatJoinRequest() bool {
	return u.ChatJoinRequest != nil
}

// GetFrom returns the `from` value from Update.
//
// NOTE: `Poll` type doesn't have `from` property.
func (u *Update) GetFrom() *User {
	if u.HasMessage() {
		return u.Message.From
	} else if u.HasEditedMessage() {
		return u.EditedMessage.From
	} else if u.HasChannelPost() {
		return u.ChannelPost.From
	} else if u.HasEditedChannelPost() {
		return u.EditedChannelPost.From
	} else if u.HasInlineQuery() {
		return &u.InlineQuery.From
	} else if u.HasChosenInlineResult() {
		return &u.ChosenInlineResult.From
	} else if u.HasCallbackQuery() {
		return &u.CallbackQuery.From
	} else if u.HasShippingQuery() {
		return &u.ShippingQuery.From
	} else if u.HasPreCheckoutQuery() {
		return &u.PreCheckoutQuery.From
	} else if u.HasPoll() {
		return nil
	} else if u.HasPollAnswer() {
		return u.PollAnswer.User
	} else if u.HasMyChatMember() {
		return &u.MyChatMember.From
	} else if u.HasChatMember() {
		return &u.ChatMember.From
	} else if u.HasChatJoinRequest() {
		return &u.ChatJoinRequest.From
	}

	return nil
}

// GetMessage returns usable message property from Update.
func (u *Update) GetMessage() (message *Message, edited bool) {
	if u.HasMessage() {
		return u.Message, false
	} else if u.HasEditedMessage() {
		return u.EditedMessage, true
	}

	return nil, false
}

// GetChannelPost returns usable channel post property from Update.
func (u *Update) GetChannelPost() (post *Message, edited bool) {
	if u.HasChannelPost() {
		return u.ChannelPost, false
	} else if u.HasEditedChannelPost() {
		return u.EditedChannelPost, true
	}

	return nil, false
}

////////////////////////////////
// Helper functions for User
//

// String function for User
func (u User) String() string {
	return structToString(u)
}

// InlineLink generates an inline link for User
func (u User) InlineLink() string {
	return fmt.Sprintf("tg://user?id=%d", u.ID)
}

////////////////////////////////
// Helper functions for Chat
//

// String function for Chat
func (c Chat) String() string {
	return structToString(c)
}

////////////////////////////////
// Helper functions for Message
//

// String function for Message
func (m Message) String() string {
	return structToString(m)
}

// HasForwardFrom checks if Message has Forward.
func (m *Message) HasForwardFrom() bool {
	return m.ForwardOrigin != nil &&
		(m.ForwardOrigin.SenderUser != nil || m.ForwardOrigin.SenderUserName != nil) &&
		m.ForwardOrigin.Date > 0
}

// HasForwardFromChat checks if Message has Forward from chat.
func (m *Message) HasForwardFromChat() bool {
	return m.ForwardOrigin != nil &&
		m.ForwardOrigin.SenderChat != nil &&
		m.ForwardOrigin.Date > 0
}

// HasReplyTo checks if Message has ReplyTo.
func (m *Message) HasReplyTo() bool {
	return m.ReplyToMessage != nil
}

// HasText checks if Message has Text.
func (m *Message) HasText() bool {
	return m.Text != nil
}

// HasMessageEntities checks if Message has MessageEntities
func (m *Message) HasMessageEntities() bool {
	return len(m.Entities) > 0
}

// HasAudio checks if Message has Audio.
func (m *Message) HasAudio() bool {
	return m.Audio != nil
}

// HasDocument checks if Message has Document.
func (m *Message) HasDocument() bool {
	return m.Document != nil
}

// HasAnimation checks if Message has Animation.
func (m *Message) HasAnimation() bool {
	return m.Animation != nil
}

// HasGame checks if Message has Game.
func (m *Message) HasGame() bool {
	return m.Game != nil
}

// HasPhoto checks if Message has Photo.
func (m *Message) HasPhoto() bool {
	return len(m.Photo) > 0
}

// LargestPhoto returns a photo with the largest file size.
func (m *Message) LargestPhoto() PhotoSize {
	if !m.HasPhoto() {
		return PhotoSize{}
	}

	maxIndex := 0
	for i, photo := range m.Photo {
		if photo.FileSize != nil && m.Photo[maxIndex].FileSize != nil && *photo.FileSize > *m.Photo[maxIndex].FileSize {
			maxIndex = i
		}
	}

	return m.Photo[maxIndex]
}

// HasSticker checks if Message has Sticker.
func (m *Message) HasSticker() bool {
	return m.Sticker != nil
}

// HasVideo checks if Message has Video.
func (m *Message) HasVideo() bool {
	return m.Video != nil
}

// HasVoice checks if Message has Voice.
func (m *Message) HasVoice() bool {
	return m.Voice != nil
}

// HasCaption checks if Message has Caption.
func (m *Message) HasCaption() bool {
	return m.Caption != nil
}

// HasContact checks if Message has Contact.
func (m *Message) HasContact() bool {
	return m.Contact != nil
}

// HasLocation checks if Message has Location.
func (m *Message) HasLocation() bool {
	return m.Location != nil
}

// HasVenue checks if Message has Venue.
func (m *Message) HasVenue() bool {
	return m.Venue != nil
}

// HasPoll checks if Message has Poll.
func (m *Message) HasPoll() bool {
	return m.Poll != nil
}

// HasNewChatMembers checks if Message has NewChatParticipant.
func (m *Message) HasNewChatMembers() bool {
	return len(m.NewChatMembers) > 0
}

// HasLeftChatMember checks if Message has LeftChatParticipant.
func (m *Message) HasLeftChatMember() bool {
	return m.LeftChatMember != nil
}

// HasNewChatTitle checks if Message has NewChatTitle.
func (m *Message) HasNewChatTitle() bool {
	return m.NewChatTitle != nil
}

// HasNewChatPhoto checks if Message has NewChatPhoto.
func (m *Message) HasNewChatPhoto() bool {
	return len(m.NewChatPhoto) > 0
}

// HasDeleteChatPhoto checks if Message has DeleteChatPhoto.
func (m *Message) HasDeleteChatPhoto() bool {
	if b := m.DeleteChatPhoto; b != nil {
		return *b
	}
	return false
}

// HasGroupChatCreated checks if Message has GroupChatCreated.
func (m *Message) HasGroupChatCreated() bool {
	if b := m.GroupChatCreated; b != nil {
		return *b
	}
	return false
}

// HasSupergroupChatCreated checks if Message has SupergroupChatCreated.
func (m *Message) HasSupergroupChatCreated() bool {
	if b := m.SupergroupChatCreated; b != nil {
		return *b
	}
	return false
}

// HasChannelChatCreated checks if Message has ChannelChatCreated.
func (m *Message) HasChannelChatCreated() bool {
	if b := m.ChannelChatCreated; b != nil {
		return *b
	}
	return false
}

// HasPinnedMessage checks if Message has PinnedMessage.
func (m *Message) HasPinnedMessage() bool {
	return m.PinnedMessage != nil
}

////////////////////////////////
// Helper functions for MaybeInaccessibleMessage
//

// IsInaccessible returns true if it is inaccessible
func (m *MaybeInaccessibleMessage) IsInaccessible() bool {
	return m.Date == 0
}

// AsMessage returns its value as `Message` or `InaccessibleMessage`
func (m *MaybeInaccessibleMessage) AsMessage() (*Message, *InaccessibleMessage) {
	if !m.IsInaccessible() {
		return (*Message)(m), nil
	} else {
		return nil, &InaccessibleMessage{
			Chat:      m.Chat,
			MessageID: m.MessageID,
			Date:      0,
		}
	}
}

////////////////////////////////
// Helper functions for InlineQuery
//

// String function for InlineQuery
func (i InlineQuery) String() string {
	return structToString(i)
}

////////////////////////////////
// Helper functions for ChosenInlineResult
//

// String function for ChosenInlineResult
func (c ChosenInlineResult) String() string {
	return structToString(c)
}

////////////////////////////////
// Helper functions for KeyboardButton and InlineKeyboardButton
//

// NewKeyboardButtons is a helper function for generating an array of KeyboardButtons
func NewKeyboardButtons(texts ...string) []KeyboardButton {
	keyboards := []KeyboardButton{}

	for _, text := range texts {
		keyboards = append(keyboards, KeyboardButton{
			Text: text,
		})
	}

	return keyboards
}

// NewInlineKeyboardButtonsWithURL is a helper function
// for generating an array of InlineKeyboardButtons with urls
func NewInlineKeyboardButtonsWithURL(values map[string]string) []InlineKeyboardButton {
	keyboards := []InlineKeyboardButton{}

	for text, url := range values {
		u := url
		keyboards = append(keyboards, InlineKeyboardButton{
			Text: text,
			URL:  &u,
		})
	}

	return keyboards
}

// NewInlineKeyboardButtonsWithCallbackData is a helper function
// for generating an array of InlineKeyboardButtons with callback data
func NewInlineKeyboardButtonsWithCallbackData(values map[string]string) []InlineKeyboardButton {
	return NewInlineKeyboardButtonsAsColumnsWithCallbackData(values)
}

// NewInlineKeyboardButtonsAsColumnsWithCallbackData is a helper function
// for generating an array of InlineKeyboardButtons (as columns) with callback data
func NewInlineKeyboardButtonsAsColumnsWithCallbackData(values map[string]string) []InlineKeyboardButton {
	keyboards := []InlineKeyboardButton{}

	for text, data := range values {
		callbackData := data
		keyboards = append(keyboards, InlineKeyboardButton{
			Text:         text,
			CallbackData: &callbackData,
		})
	}

	return keyboards
}

// NewInlineKeyboardButtonsAsRowsWithCallbackData is a helper function
// for generating an array of InlineKeyboardButtons (as rows) with callback data
func NewInlineKeyboardButtonsAsRowsWithCallbackData(values map[string]string) [][]InlineKeyboardButton {
	keyboards := [][]InlineKeyboardButton{}

	for text, data := range values {
		callbackData := data
		keyboards = append(keyboards, []InlineKeyboardButton{
			{
				Text:         text,
				CallbackData: &callbackData,
			},
		})
	}

	return keyboards
}

// NewInlineKeyboardButtonsWithSwitchInlineQuery is a helper function
// for generating an array of InlineKeyboardButtons with switch inline query
func NewInlineKeyboardButtonsWithSwitchInlineQuery(values map[string]string) []InlineKeyboardButton {
	keyboards := []InlineKeyboardButton{}

	for text, switchInlineQuery := range values {
		query := switchInlineQuery
		keyboards = append(keyboards, InlineKeyboardButton{
			Text:              text,
			SwitchInlineQuery: &query,
		})
	}

	return keyboards
}

////////////////////////////////
// Helper functions for CallbackQuery

// String function for CallbackQuery
func (q CallbackQuery) String() string {
	return structToString(q)
}

////////////////////////////////
// Helper functions for ReactionType

// NewEmojiReaction returns a ReactionType with emoji.
func NewEmojiReaction(emoji string) ReactionType {
	return ReactionType{
		Type:  "emoji",
		Emoji: &emoji,
	}
}

// NewCustomEmojiReaction returns a ReactionType with custom emoji.
func NewCustomEmojiReaction(customEmojiID string) ReactionType {
	return ReactionType{
		Type:          "custom_emoji",
		CustomEmojiID: &customEmojiID,
	}
}

// NewMessageReactionWithEmoji returns a new OptionsSetMessageReaction with an emoji string for function `SetMessageReaction`.
func NewMessageReactionWithEmoji(emoji string) OptionsSetMessageReaction {
	return OptionsSetMessageReaction{}.
		SetReaction([]ReactionType{
			NewEmojiReaction(emoji),
		})
}

////////////////////////////////
// Helper functions for MessageOrigin

// NewMessageOriginUser returns a new MessageOriginUser.
func NewMessageOriginUser(user User, date int) MessageOrigin {
	return MessageOrigin{
		Type:       "user",
		Date:       date,
		SenderUser: &user,
	}
}

// NewMessageOriginHiddenUser returns a new MessageOriginHiddenUser.
func NewMessageOriginHiddenUser(senderUserName string, date int) MessageOrigin {
	return MessageOrigin{
		Type:           "hidden_user",
		Date:           date,
		SenderUserName: &senderUserName,
	}
}

// NewMessageOriginChat returns a new MessageOriginChat.
//
// `authorSignature` can be nil.
func NewMessageOriginChat(senderChat Chat, date int, authorSignature *string) MessageOrigin {
	return MessageOrigin{
		Type:            "chat",
		Date:            date,
		SenderChat:      &senderChat,
		AuthorSignature: authorSignature,
	}
}

// NewMessageOriginChannel returns a new MessageOriginChannel.
//
// `authorSignature` can be nil.
func NewMessageOriginChannel(chat Chat, messageID int64, date int, authorSignature *string) MessageOrigin {
	return MessageOrigin{
		Type:            "channel",
		Date:            date,
		Chat:            &chat,
		MessageID:       &messageID,
		AuthorSignature: authorSignature,
	}
}

////////////////////////////////
// Helper functions for ChatBoostSource

// NewChatBoostSourcePremium returns a new ChatBoostSourcePremium.
func NewChatBoostSourcePremium(user User) ChatBoostSource {
	return ChatBoostSource{
		Source: "premium",
		User:   &user,
	}
}

// NewChatBoostSourceGiftCode returns a new ChatBoostSourceGiftCode.
func NewChatBoostSourceGiftCode(user User) ChatBoostSource {
	return ChatBoostSource{
		Source: "gift_code",
		User:   &user,
	}
}

// NewChatBoostSourceGiveaway returns a new ChatBoostSourceGiveaway.
func NewChatBoostSourceGiveaway(giveawayMessageID int64, user *User, isUnclaimed bool) ChatBoostSource {
	return ChatBoostSource{
		Source:            "giveaway",
		GiveawayMessageID: &giveawayMessageID,
		User:              user,
		IsUnclaimed:       &isUnclaimed,
	}
}

////////////////////////////////
// Other helper functions

// interface to string (in JSON format)
func structToString(v any) string {
	json, err := json.Marshal(v)

	if err == nil {
		return fmt.Sprintf("%T%s", v, string(json))
	}

	return err.Error()
}

// InputFileFromFilepath generates an InputFile from given filepath
func InputFileFromFilepath(filepath string) InputFile {
	return InputFile{
		Filepath: &filepath,
	}
}

// InputFileFromURL generates an InputFile from given url
func InputFileFromURL(url string) InputFile {
	return InputFile{
		URL: &url,
	}
}

// InputFileFromBytes generates an InputFile from given bytes array
func InputFileFromBytes(bytes []byte) InputFile {
	return InputFile{
		Bytes: bytes,
	}
}

// InputFileFromFileID generates an InputFile from given file id
func InputFileFromFileID(fileID string) InputFile {
	return InputFile{
		FileID: &fileID,
	}
}
