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

// Helper function for generating a new InlineQueryResultArticle
//
// https://core.telegram.org/bots/api#inlinequeryresultarticle
func NewInlineQueryResultArticle(title, messageText, description string) (newArticle *InlineQueryResultArticle, generatedId *string) {
	if id, err := newUUID(); err == nil {
		return &InlineQueryResultArticle{
			InlineQueryResult: InlineQueryResult{
				Type: InlineQueryResultTypeArticle,
				Id:   &id,
			},
			Title: &title,
			InputMessageContent: InputTextMessageContent{
				MessageText: &messageText,
			},
			Description: &description,
		}, &id
	}

	return &InlineQueryResultArticle{}, nil
}

// Helper function for generating a new InlineQueryResultPhoto
//
// Photo must be in jpeg format, < 5MB.
//
// https://core.telegram.org/bots/api#inlinequeryresultphoto
func NewInlineQueryResultPhoto(photoUrl, thumbUrl string) (newPhoto *InlineQueryResultPhoto, generatedId *string) {
	if id, err := newUUID(); err == nil {
		return &InlineQueryResultPhoto{
			InlineQueryResult: InlineQueryResult{
				Type: InlineQueryResultTypePhoto,
				Id:   &id,
			},
			PhotoUrl: &photoUrl,
			ThumbUrl: &thumbUrl,
		}, &id
	}

	return &InlineQueryResultPhoto{}, nil
}

// Helper function for generating a new InlineQueryResultGif
//
// Gif must be in gif format, < 1MB.
//
// https://core.telegram.org/bots/api#inlinequeryresultgif
func NewInlineQueryResultGif(gifUrl, thumbUrl string) (newGif *InlineQueryResultGif, generatedId *string) {
	if id, err := newUUID(); err == nil {
		return &InlineQueryResultGif{
			InlineQueryResult: InlineQueryResult{
				Type: InlineQueryResultTypeGif,
				Id:   &id,
			},
			GifUrl:   &gifUrl,
			ThumbUrl: &thumbUrl,
		}, &id
	}

	return &InlineQueryResultGif{}, nil
}

// Helper function for generating a new InlineQueryResultMpeg4Gif
//
// Mpeg4 must be in H.264/MPEG-4 AVC video(wihout sound) format, < 1MB.
//
// https://core.telegram.org/bots/api#inlinequeryresultmpeg4gif
func NewInlineQueryResultMpeg4Gif(mpeg4Url, thumbUrl string) (newMpeg4Gif *InlineQueryResultMpeg4Gif, generatedId *string) {
	if id, err := newUUID(); err == nil {
		return &InlineQueryResultMpeg4Gif{
			InlineQueryResult: InlineQueryResult{
				Type: InlineQueryResultTypeMpeg4Gif,
				Id:   &id,
			},
			Mpeg4Url: &mpeg4Url,
			ThumbUrl: &thumbUrl,
		}, &id
	}

	return &InlineQueryResultMpeg4Gif{}, nil
}

// Helper function for generating a new InlineQueryResultVideo
//
// https://core.telegram.org/bots/api#inlinequeryresultvideo
func NewInlineQueryResultVideo(videoUrl, thumbUrl, title string, mimeType VideoMimeType) (newVideo *InlineQueryResultVideo, generatedId *string) {
	if id, err := newUUID(); err == nil {
		return &InlineQueryResultVideo{
			InlineQueryResult: InlineQueryResult{
				Type: InlineQueryResultTypeVideo,
				Id:   &id,
			},
			VideoUrl: &videoUrl,
			MimeType: &mimeType,
			ThumbUrl: &thumbUrl,
			Title:    &title,
		}, &id
	}

	return &InlineQueryResultVideo{}, nil
}

// Helper function for generating a new InlineQueryResultAudio
//
// https://core.telegram.org/bots/api#inlinequeryresultaudio
func NewInlineQueryResultAudio(audioUrl, title string) (newAudio *InlineQueryResultAudio, generatedId *string) {
	if id, err := newUUID(); err == nil {
		return &InlineQueryResultAudio{
			InlineQueryResult: InlineQueryResult{
				Type: InlineQueryResultTypeAudio,
				Id:   &id,
			},
			AudioUrl: &audioUrl,
			Title:    &title,
		}, &id
	}

	return &InlineQueryResultAudio{}, nil
}

// Helper function for generating a new InlineQueryResultVoice
//
// https://core.telegram.org/bots/api#inlinequeryresultvoice
func NewInlineQueryResultVoice(voiceUrl, title string) (newVoice *InlineQueryResultVoice, generatedId *string) {
	if id, err := newUUID(); err == nil {
		return &InlineQueryResultVoice{
			InlineQueryResult: InlineQueryResult{
				Type: InlineQueryResultTypeVoice,
				Id:   &id,
			},
			VoiceUrl: &voiceUrl,
			Title:    &title,
		}, &id
	}

	return &InlineQueryResultVoice{}, nil
}

// Helper function for generating a new InlineQueryResultDocument
//
// https://core.telegram.org/bots/api#inlinequeryresultdocument
func NewInlineQueryResultDocument(documentUrl, title string, mimeType DocumentMimeType) (newDocument *InlineQueryResultDocument, generatedId *string) {
	if id, err := newUUID(); err == nil {
		return &InlineQueryResultDocument{
			InlineQueryResult: InlineQueryResult{
				Type: InlineQueryResultTypeDocument,
				Id:   &id,
			},
			Title:       &title,
			DocumentUrl: &documentUrl,
			MimeType:    &mimeType,
		}, &id
	}

	return &InlineQueryResultDocument{}, nil
}

// Helper function for generating a new InlineQueryResultLocation
//
// https://core.telegram.org/bots/api#inlinequeryresultlocation
func NewInlineQueryResultLocation(latitude, longitude float32, title string) (newLocation *InlineQueryResultLocation, generatedId *string) {
	if id, err := newUUID(); err == nil {
		return &InlineQueryResultLocation{
			InlineQueryResult: InlineQueryResult{
				Type: InlineQueryResultTypeLocation,
				Id:   &id,
			},
			Latitude:  latitude,
			Longitude: longitude,
			Title:     &title,
		}, &id
	}

	return &InlineQueryResultLocation{}, nil
}

// Helper function for generating a new InlineQueryResultVenue
//
// https://core.telegram.org/bots/api#inlinequeryresultvenue
func NewInlineQueryResultVenue(latitude, longitude float32, title, address string) (newVenue *InlineQueryResultVenue, generatedId *string) {
	if id, err := newUUID(); err == nil {
		return &InlineQueryResultVenue{
			InlineQueryResult: InlineQueryResult{
				Type: InlineQueryResultTypeVenue,
				Id:   &id,
			},
			Latitude:  latitude,
			Longitude: longitude,
			Title:     &title,
			Address:   &address,
		}, &id
	}

	return &InlineQueryResultVenue{}, nil
}

// Helper function for generating a new InlineQueryResultContact
//
// https://core.telegram.org/bots/api#inlinequeryresultcontact
func NewInlineQueryResultContact(phoneNumber, firstName, lastName string) (newContact *InlineQueryResultContact, generatedId *string) {
	if id, err := newUUID(); err == nil {
		return &InlineQueryResultContact{
			InlineQueryResult: InlineQueryResult{
				Type: InlineQueryResultTypeContact,
				Id:   &id,
			},
			PhoneNumber: &phoneNumber,
			FirstName:   &firstName,
			LastName:    &lastName,
		}, &id
	}

	return &InlineQueryResultContact{}, nil
}

// Helper function for generating a new InlineQueryResultCachedPhoto
//
// https://core.telegram.org/bots/api#inlinequeryresultcachedphoto
func NewInlineQueryResultCachedPhoto(photoFileId string) (newPhoto *InlineQueryResultCachedPhoto, generatedId *string) {
	if id, err := newUUID(); err == nil {
		return &InlineQueryResultCachedPhoto{
			InlineQueryResult: InlineQueryResult{
				Type: InlineQueryResultTypePhoto,
				Id:   &id,
			},
			PhotoFileId: &photoFileId,
		}, &id
	}

	return &InlineQueryResultCachedPhoto{}, nil
}

// Helper function for generating a new InlineQueryResultCachedGif
//
// https://core.telegram.org/bots/api#inlinequeryresultcachedgif
func NewInlineQueryResultCachedGif(gifFileId string) (newGif *InlineQueryResultCachedGif, generatedId *string) {
	if id, err := newUUID(); err == nil {
		return &InlineQueryResultCachedGif{
			InlineQueryResult: InlineQueryResult{
				Type: InlineQueryResultTypeGif,
				Id:   &id,
			},
			GifFileId: &gifFileId,
		}, &id
	}

	return &InlineQueryResultCachedGif{}, nil
}

// Helper function for generating a new InlineQueryResultCachedMpeg4Gif
//
// https://core.telegram.org/bots/api#inlinequeryresultcachedmpeg4gif
func NewInlineQueryResultCachedMpeg4Gif(mpeg4FileId string) (newMpeg4Gif *InlineQueryResultCachedMpeg4Gif, generatedId *string) {
	if id, err := newUUID(); err == nil {
		return &InlineQueryResultCachedMpeg4Gif{
			InlineQueryResult: InlineQueryResult{
				Type: InlineQueryResultTypeMpeg4Gif,
				Id:   &id,
			},
			Mpeg4FileId: &mpeg4FileId,
		}, &id
	}

	return &InlineQueryResultCachedMpeg4Gif{}, nil
}

// Helper function for generating a new InlineQueryResultCachedSticker
//
// https://core.telegram.org/bots/api#inlinequeryresultcachedsticker
func NewInlineQueryResultCachedSticker(stickerFileId string) (newSticker *InlineQueryResultCachedSticker, generatedId *string) {
	if id, err := newUUID(); err == nil {
		return &InlineQueryResultCachedSticker{
			InlineQueryResult: InlineQueryResult{
				Type: InlineQueryResultTypeSticker,
				Id:   &id,
			},
			StickerFileId: &stickerFileId,
		}, &id
	}

	return &InlineQueryResultCachedSticker{}, nil
}

// Helper function for generating a new InlineQueryResultCachedDocument
//
// https://core.telegram.org/bots/api#inlinequeryresultcacheddocument
func NewInlineQueryResultCachedDocument(title, documentFileId string) (newDocument *InlineQueryResultCachedDocument, generatedId *string) {
	if id, err := newUUID(); err == nil {
		return &InlineQueryResultCachedDocument{
			InlineQueryResult: InlineQueryResult{
				Type: InlineQueryResultTypeDocument,
				Id:   &id,
			},
			Title:          &title,
			DocumentFileId: &documentFileId,
		}, &id
	}

	return &InlineQueryResultCachedDocument{}, nil
}

// Helper function for generating a new InlineQueryResultCachedVideo
//
// https://core.telegram.org/bots/api#inlinequeryresultcachedvideo
func NewInlineQueryResultCachedVideo(title, videoFileId string) (newVideo *InlineQueryResultCachedVideo, generatedId *string) {
	if id, err := newUUID(); err == nil {
		return &InlineQueryResultCachedVideo{
			InlineQueryResult: InlineQueryResult{
				Type: InlineQueryResultTypeVideo,
				Id:   &id,
			},
			Title:       &title,
			VideoFileId: &videoFileId,
		}, &id
	}

	return &InlineQueryResultCachedVideo{}, nil
}

// Helper function for generating a new InlineQueryResultCachedVoice
//
// https://core.telegram.org/bots/api#inlinequeryresultcachedvoice
func NewInlineQueryResultCachedVoice(title, voiceFileId string) (newVoice *InlineQueryResultCachedVoice, generatedId *string) {
	if id, err := newUUID(); err == nil {
		return &InlineQueryResultCachedVoice{
			InlineQueryResult: InlineQueryResult{
				Type: InlineQueryResultTypeVoice,
				Id:   &id,
			},
			Title:       &title,
			VoiceFileId: &voiceFileId,
		}, &id
	}

	return &InlineQueryResultCachedVoice{}, nil
}

// Helper function for generating a new InlineQueryResultCachedAudio
//
// https://core.telegram.org/bots/api#inlinequeryresultcachedaudio
func NewInlineQueryResultCachedAudio(audioFileId string) (newAudio *InlineQueryResultCachedAudio, generatedId *string) {
	if id, err := newUUID(); err == nil {
		return &InlineQueryResultCachedAudio{
			InlineQueryResult: InlineQueryResult{
				Type: InlineQueryResultTypeAudio,
				Id:   &id,
			},
			AudioFileId: &audioFileId,
		}, &id
	}

	return &InlineQueryResultCachedAudio{}, nil
}

////////////////////////////////
// Helper functions for Update
//

// String function for Update
func (u Update) String() string {
	if json, err := json.Marshal(u); err == nil {
		return fmt.Sprintf("%T%s", u, string(json))
	}
	return fmt.Sprintf("%+v", u)
}

// Check if Update has Message.
func (u *Update) HasMessage() bool {
	return u.Message != nil
}

// Check if Update has InlineQuery
func (u *Update) HasInlineQuery() bool {
	return u.InlineQuery != nil
}

// Check if Update has ChosenInlineResult
func (u *Update) HasChosenInlineResult() bool {
	return u.ChosenInlineResult != nil
}

// Check if Update has CallbackQuery
func (u *Update) HasCallbackQuery() bool {
	return u.CallbackQuery != nil
}

////////////////////////////////
// Helper functions for User
//

// String function for User
func (u User) String() string {
	if json, err := json.Marshal(u); err == nil {
		return fmt.Sprintf("%T%s", u, string(json))
	}
	return fmt.Sprintf("%+v", u)
}

////////////////////////////////
// Helper functions for Chat
//

// String function for Chat
func (c Chat) String() string {
	if json, err := json.Marshal(c); err == nil {
		return fmt.Sprintf("%T%s", c, string(json))
	}
	return fmt.Sprintf("%+v", c)
}

////////////////////////////////
// Helper functions for Message
//

// String function for Message
func (m Message) String() string {
	if json, err := json.Marshal(m); err == nil {
		return fmt.Sprintf("%T%s", m, string(json))
	}
	return fmt.Sprintf("%+v", m)
}

// Check if Message has Forward.
func (m *Message) HasForward() bool {
	return m.ForwardFrom != nil && m.ForwardDate > 0
}

// Check if Message has ReplyTo.
func (m *Message) HasReplyTo() bool {
	return m.ReplyToMessage != nil
}

// Check if Message has Text.
func (m *Message) HasText() bool {
	return m.Text != nil
}

// Check if Message has MessageEntities
func (m *Message) HasMessageEntities() bool {
	return len(m.Entities) > 0
}

// Check if Message has Audio.
func (m *Message) HasAudio() bool {
	return m.Audio != nil
}

// Check if Message has Document.
func (m *Message) HasDocument() bool {
	return m.Document != nil
}

// Check if Message has Photo.
func (m *Message) HasPhoto() bool {
	return len(m.Photo) > 0
}

// Check if Message has Sticker.
func (m *Message) HasSticker() bool {
	return m.Sticker != nil
}

// Check if Message has Video.
func (m *Message) HasVideo() bool {
	return m.Video != nil
}

// Check if Message has Voice.
func (m *Message) HasVoice() bool {
	return m.Voice != nil
}

// Check if Message has Caption.
func (m *Message) HasCaption() bool {
	return m.Caption != nil
}

// Check if Message has Contact.
func (m *Message) HasContact() bool {
	return m.Contact != nil
}

// Check if Message has Location.
func (m *Message) HasLocation() bool {
	return m.Location != nil
}

// Check if Message has Venue.
func (m *Message) HasVenue() bool {
	return m.Venue != nil
}

// Check if Message has NewChatParticipant.
func (m *Message) HasNewChatMember() bool {
	return m.NewChatMember != nil
}

// Check if Message has LeftChatParticipant.
func (m *Message) HasLeftChatMember() bool {
	return m.LeftChatMember != nil
}

// Check if Message has NewChatTitle.
func (m *Message) HasNewChatTitle() bool {
	return m.NewChatTitle != nil
}

// Check if Message has NewChatPhoto.
func (m *Message) HasNewChatPhoto() bool {
	return len(m.NewChatPhoto) > 0
}

// Check if Message has DeleteChatPhoto.
func (m *Message) HasDeleteChatPhoto() bool {
	return m.DeleteChatPhoto
}

// Check if Message has GroupChatCreated.
func (m *Message) HasGroupChatCreated() bool {
	return m.GroupChatCreated
}

// Check if Message has SupergroupChatCreated.
func (m *Message) HasSupergroupChatCreated() bool {
	return m.SupergroupChatCreated
}

// Check if Message has ChannelChatCreated.
func (m *Message) HasChannelChatCreated() bool {
	return m.ChannelChatCreated
}

// Check if Message has MigrateToChatId.
func (m *Message) HasMigrateToChatId() bool {
	return m.MigrateToChatId > 0
}

// Check if Message has MigrateFromChatId.
func (m *Message) HasMigrateFromChatId() bool {
	return m.MigrateFromChatId > 0
}

// Check if Message has PinnedMessage.
func (m *Message) HasPinnedMessage() bool {
	return m.PinnedMessage != nil
}

////////////////////////////////
// Helper functions for InlineQuery
//

// String function for InlineQuery
func (i InlineQuery) String() string {
	if json, err := json.Marshal(i); err == nil {
		return fmt.Sprintf("%T%s", i, string(json))
	}
	return fmt.Sprintf("%+v", i)
}

////////////////////////////////
// Helper functions for ChosenInlineResult
//

// String function for ChosenInlineResult
func (c ChosenInlineResult) String() string {
	if json, err := json.Marshal(c); err == nil {
		return fmt.Sprintf("%T%s", c, string(json))
	}
	return fmt.Sprintf("%+v", c)
}

////////////////////////////////
// Helper functions for KeyboardButton and InlineKeyboardButton
//

// Helper function for generating an array of KeyboardButtons
func NewKeyboardButtons(texts ...string) []KeyboardButton {
	keyboards := []KeyboardButton{}

	for _, text := range texts {
		keyboards = append(keyboards, KeyboardButton{
			Text: text,
		})
	}

	return keyboards
}

// Helper function for generating an array of InlineKeyboardButtons with url
func NewInlineKeyboardButtonsWithUrl(values map[string]string) []InlineKeyboardButton {
	keyboards := []InlineKeyboardButton{}

	for text, url := range values {
		keyboards = append(keyboards, InlineKeyboardButton{
			Text: text,
			Url:  url,
		})
	}

	return keyboards
}

// Helper function for generating an array of InlineKeyboardButtons with callback data
func NewInlineKeyboardButtonsWithCallbackData(values map[string]string) []InlineKeyboardButton {
	keyboards := []InlineKeyboardButton{}

	for text, callbackData := range values {
		keyboards = append(keyboards, InlineKeyboardButton{
			Text:         text,
			CallbackData: callbackData,
		})
	}

	return keyboards
}

// Helper function for generating an array of InlineKeyboardButtons with switch inline query
func NewInlineKeyboardButtonsWithSwitchInlineQuery(values map[string]string) []InlineKeyboardButton {
	keyboards := []InlineKeyboardButton{}

	for text, switchInlineQuery := range values {
		keyboards = append(keyboards, InlineKeyboardButton{
			Text:              text,
			SwitchInlineQuery: switchInlineQuery,
		})
	}

	return keyboards
}

////////////////////////////////
// Helper functions for CallbackQuery

// String function for CallbackQuery
func (q CallbackQuery) String() string {
	if json, err := json.Marshal(q); err == nil {
		return fmt.Sprintf("%T%s", q, string(json))
	}
	return fmt.Sprintf("%+v", q)
}
