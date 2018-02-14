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
func NewInlineQueryResultArticle(title, messageText, description string) (newArticle *InlineQueryResultArticle, generatedId *string) {
	if id, err := newUUID(); err == nil {
		return &InlineQueryResultArticle{
			InlineQueryResult: InlineQueryResult{
				Type: InlineQueryResultTypeArticle,
				Id:   id,
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
func NewInlineQueryResultPhoto(photoUrl, thumbUrl string) (newPhoto *InlineQueryResultPhoto, generatedId *string) {
	if id, err := newUUID(); err == nil {
		return &InlineQueryResultPhoto{
			InlineQueryResult: InlineQueryResult{
				Type: InlineQueryResultTypePhoto,
				Id:   id,
			},
			PhotoUrl: photoUrl,
			ThumbUrl: thumbUrl,
		}, &id
	}

	return &InlineQueryResultPhoto{}, nil
}

// NewInlineQueryResultGif is a helper function for generating a new InlineQueryResultGif
//
// Gif must be in gif format, < 1MB.
//
// https://core.telegram.org/bots/api#inlinequeryresultgif
func NewInlineQueryResultGif(gifUrl, thumbUrl string) (newGif *InlineQueryResultGif, generatedId *string) {
	if id, err := newUUID(); err == nil {
		return &InlineQueryResultGif{
			InlineQueryResult: InlineQueryResult{
				Type: InlineQueryResultTypeGif,
				Id:   id,
			},
			GifUrl:   gifUrl,
			ThumbUrl: thumbUrl,
		}, &id
	}

	return &InlineQueryResultGif{}, nil
}

// NewInlineQueryResultMpeg4Gif is a helper function for generating a new InlineQueryResultMpeg4Gif
//
// Mpeg4 must be in H.264/MPEG-4 AVC video(wihout sound) format, < 1MB.
//
// https://core.telegram.org/bots/api#inlinequeryresultmpeg4gif
func NewInlineQueryResultMpeg4Gif(mpeg4Url, thumbUrl string) (newMpeg4Gif *InlineQueryResultMpeg4Gif, generatedId *string) {
	if id, err := newUUID(); err == nil {
		return &InlineQueryResultMpeg4Gif{
			InlineQueryResult: InlineQueryResult{
				Type: InlineQueryResultTypeMpeg4Gif,
				Id:   id,
			},
			Mpeg4Url: mpeg4Url,
			ThumbUrl: thumbUrl,
		}, &id
	}

	return &InlineQueryResultMpeg4Gif{}, nil
}

// NewInlineQueryResultVideo is a helper function for generating a new InlineQueryResultVideo
//
// https://core.telegram.org/bots/api#inlinequeryresultvideo
func NewInlineQueryResultVideo(videoUrl, thumbUrl, title string, mimeType VideoMimeType) (newVideo *InlineQueryResultVideo, generatedId *string) {
	if id, err := newUUID(); err == nil {
		return &InlineQueryResultVideo{
			InlineQueryResult: InlineQueryResult{
				Type: InlineQueryResultTypeVideo,
				Id:   id,
			},
			VideoUrl: videoUrl,
			MimeType: mimeType,
			ThumbUrl: thumbUrl,
			Title:    title,
		}, &id
	}

	return &InlineQueryResultVideo{}, nil
}

// NewInlineQueryResultAudio is a helper function for generating a new InlineQueryResultAudio
//
// https://core.telegram.org/bots/api#inlinequeryresultaudio
func NewInlineQueryResultAudio(audioUrl, title string) (newAudio *InlineQueryResultAudio, generatedId *string) {
	if id, err := newUUID(); err == nil {
		return &InlineQueryResultAudio{
			InlineQueryResult: InlineQueryResult{
				Type: InlineQueryResultTypeAudio,
				Id:   id,
			},
			AudioUrl: audioUrl,
			Title:    title,
		}, &id
	}

	return &InlineQueryResultAudio{}, nil
}

// NewInlineQueryResultVoice is a helper function for generating a new InlineQueryResultVoice
//
// https://core.telegram.org/bots/api#inlinequeryresultvoice
func NewInlineQueryResultVoice(voiceUrl, title string) (newVoice *InlineQueryResultVoice, generatedId *string) {
	if id, err := newUUID(); err == nil {
		return &InlineQueryResultVoice{
			InlineQueryResult: InlineQueryResult{
				Type: InlineQueryResultTypeVoice,
				Id:   id,
			},
			VoiceUrl: voiceUrl,
			Title:    title,
		}, &id
	}

	return &InlineQueryResultVoice{}, nil
}

// NewInlineQueryResultDocument is a helper function for generating a new InlineQueryResultDocument
//
// https://core.telegram.org/bots/api#inlinequeryresultdocument
func NewInlineQueryResultDocument(documentUrl, title string, mimeType DocumentMimeType) (newDocument *InlineQueryResultDocument, generatedId *string) {
	if id, err := newUUID(); err == nil {
		return &InlineQueryResultDocument{
			InlineQueryResult: InlineQueryResult{
				Type: InlineQueryResultTypeDocument,
				Id:   id,
			},
			Title:       title,
			DocumentUrl: documentUrl,
			MimeType:    mimeType,
		}, &id
	}

	return &InlineQueryResultDocument{}, nil
}

// NewInlineQueryResultLocation is a helper function for generating a new InlineQueryResultLocation
//
// https://core.telegram.org/bots/api#inlinequeryresultlocation
func NewInlineQueryResultLocation(latitude, longitude float32, title string) (newLocation *InlineQueryResultLocation, generatedId *string) {
	if id, err := newUUID(); err == nil {
		return &InlineQueryResultLocation{
			InlineQueryResult: InlineQueryResult{
				Type: InlineQueryResultTypeLocation,
				Id:   id,
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
func NewInlineQueryResultVenue(latitude, longitude float32, title, address string) (newVenue *InlineQueryResultVenue, generatedId *string) {
	if id, err := newUUID(); err == nil {
		return &InlineQueryResultVenue{
			InlineQueryResult: InlineQueryResult{
				Type: InlineQueryResultTypeVenue,
				Id:   id,
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
func NewInlineQueryResultContact(phoneNumber, firstName string) (newContact *InlineQueryResultContact, generatedId *string) {
	if id, err := newUUID(); err == nil {
		return &InlineQueryResultContact{
			InlineQueryResult: InlineQueryResult{
				Type: InlineQueryResultTypeContact,
				Id:   id,
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
func NewInlineQueryResultCachedPhoto(photoFileId string) (newPhoto *InlineQueryResultCachedPhoto, generatedId *string) {
	if id, err := newUUID(); err == nil {
		return &InlineQueryResultCachedPhoto{
			InlineQueryResult: InlineQueryResult{
				Type: InlineQueryResultTypePhoto,
				Id:   id,
			},
			PhotoFileId: photoFileId,
		}, &id
	}

	return &InlineQueryResultCachedPhoto{}, nil
}

// NewInlineQueryResultCachedGif is a helper function for generating a new InlineQueryResultCachedGif
//
// https://core.telegram.org/bots/api#inlinequeryresultcachedgif
func NewInlineQueryResultCachedGif(gifFileId string) (newGif *InlineQueryResultCachedGif, generatedId *string) {
	if id, err := newUUID(); err == nil {
		return &InlineQueryResultCachedGif{
			InlineQueryResult: InlineQueryResult{
				Type: InlineQueryResultTypeGif,
				Id:   id,
			},
			GifFileId: gifFileId,
		}, &id
	}

	return &InlineQueryResultCachedGif{}, nil
}

// NewInlineQueryResultCachedMpeg4Gif is a helper function for generating a new InlineQueryResultCachedMpeg4Gif
//
// https://core.telegram.org/bots/api#inlinequeryresultcachedmpeg4gif
func NewInlineQueryResultCachedMpeg4Gif(mpeg4FileId string) (newMpeg4Gif *InlineQueryResultCachedMpeg4Gif, generatedId *string) {
	if id, err := newUUID(); err == nil {
		return &InlineQueryResultCachedMpeg4Gif{
			InlineQueryResult: InlineQueryResult{
				Type: InlineQueryResultTypeMpeg4Gif,
				Id:   id,
			},
			Mpeg4FileId: mpeg4FileId,
		}, &id
	}

	return &InlineQueryResultCachedMpeg4Gif{}, nil
}

// NewInlineQueryResultCachedSticker is a helper function for generating a new InlineQueryResultCachedSticker
//
// https://core.telegram.org/bots/api#inlinequeryresultcachedsticker
func NewInlineQueryResultCachedSticker(stickerFileId string) (newSticker *InlineQueryResultCachedSticker, generatedId *string) {
	if id, err := newUUID(); err == nil {
		return &InlineQueryResultCachedSticker{
			InlineQueryResult: InlineQueryResult{
				Type: InlineQueryResultTypeSticker,
				Id:   id,
			},
			StickerFileId: stickerFileId,
		}, &id
	}

	return &InlineQueryResultCachedSticker{}, nil
}

// NewInlineQueryResultCachedDocument is a helper function for generating a new InlineQueryResultCachedDocument
//
// https://core.telegram.org/bots/api#inlinequeryresultcacheddocument
func NewInlineQueryResultCachedDocument(title, documentFileId string) (newDocument *InlineQueryResultCachedDocument, generatedId *string) {
	if id, err := newUUID(); err == nil {
		return &InlineQueryResultCachedDocument{
			InlineQueryResult: InlineQueryResult{
				Type: InlineQueryResultTypeDocument,
				Id:   id,
			},
			Title:          title,
			DocumentFileId: documentFileId,
		}, &id
	}

	return &InlineQueryResultCachedDocument{}, nil
}

// NewInlineQueryResultCachedVideo is a helper function for generating a new InlineQueryResultCachedVideo
//
// https://core.telegram.org/bots/api#inlinequeryresultcachedvideo
func NewInlineQueryResultCachedVideo(title, videoFileId string) (newVideo *InlineQueryResultCachedVideo, generatedId *string) {
	if id, err := newUUID(); err == nil {
		return &InlineQueryResultCachedVideo{
			InlineQueryResult: InlineQueryResult{
				Type: InlineQueryResultTypeVideo,
				Id:   id,
			},
			Title:       title,
			VideoFileId: videoFileId,
		}, &id
	}

	return &InlineQueryResultCachedVideo{}, nil
}

// NewInlineQueryResultCachedVoice is a helper function for generating a new InlineQueryResultCachedVoice
//
// https://core.telegram.org/bots/api#inlinequeryresultcachedvoice
func NewInlineQueryResultCachedVoice(title, voiceFileId string) (newVoice *InlineQueryResultCachedVoice, generatedId *string) {
	if id, err := newUUID(); err == nil {
		return &InlineQueryResultCachedVoice{
			InlineQueryResult: InlineQueryResult{
				Type: InlineQueryResultTypeVoice,
				Id:   id,
			},
			Title:       title,
			VoiceFileId: voiceFileId,
		}, &id
	}

	return &InlineQueryResultCachedVoice{}, nil
}

// NewInlineQueryResultCachedAudio is a helper function for generating a new InlineQueryResultCachedAudio
//
// https://core.telegram.org/bots/api#inlinequeryresultcachedaudio
func NewInlineQueryResultCachedAudio(audioFileId string) (newAudio *InlineQueryResultCachedAudio, generatedId *string) {
	if id, err := newUUID(); err == nil {
		return &InlineQueryResultCachedAudio{
			InlineQueryResult: InlineQueryResult{
				Type: InlineQueryResultTypeAudio,
				Id:   id,
			},
			AudioFileId: audioFileId,
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

////////////////////////////////
// Helper functions for User
//

// String function for User
func (u User) String() string {
	return structToString(u)
}

// InlineLink generates an inline link for User
func (u User) InlineLink() string {
	return fmt.Sprintf("tg://user?id=%d", u.Id)
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
	return m.ForwardFrom != nil && m.ForwardDate > 0
}

// HasForwardFromChat checks if Message has Forward from chat.
func (m *Message) HasForwardFromChat() bool {
	return m.ForwardFromChat != nil && m.ForwardDate > 0
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

// HasPhoto checks if Message has Photo.
func (m *Message) HasPhoto() bool {
	return len(m.Photo) > 0
}

// LargestPhoto returns a photo with the largest file size.
func (m *Message) LargestPhoto() PhotoSize {
	var maxIndex int = 0
	for i, photo := range m.Photo {
		if photo.FileSize > m.Photo[maxIndex].FileSize {
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
	return m.DeleteChatPhoto
}

// HasGroupChatCreated checks if Message has GroupChatCreated.
func (m *Message) HasGroupChatCreated() bool {
	return m.GroupChatCreated
}

// HasSupergroupChatCreated checks if Message has SupergroupChatCreated.
func (m *Message) HasSupergroupChatCreated() bool {
	return m.SupergroupChatCreated
}

// HasChannelChatCreated checks if Message has ChannelChatCreated.
func (m *Message) HasChannelChatCreated() bool {
	return m.ChannelChatCreated
}

// HasMigrateToChatId checks if Message has MigrateToChatId.
func (m *Message) HasMigrateToChatId() bool {
	return m.MigrateToChatId > 0
}

// HasMigrateFromChatId checks if Message has MigrateFromChatId.
func (m *Message) HasMigrateFromChatId() bool {
	return m.MigrateFromChatId > 0
}

// HasPinnedMessage checks if Message has PinnedMessage.
func (m *Message) HasPinnedMessage() bool {
	return m.PinnedMessage != nil
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

// NewInlineKeyboardButtonsWithUrl is a helper function
// for generating an array of InlineKeyboardButtons with urls
func NewInlineKeyboardButtonsWithUrl(values map[string]string) []InlineKeyboardButton {
	keyboards := []InlineKeyboardButton{}

	for text, url := range values {
		u := url
		keyboards = append(keyboards, InlineKeyboardButton{
			Text: text,
			Url:  &u,
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
			InlineKeyboardButton{
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
// Other helper functions

// interface to string (in JSON format)
func structToString(v interface{}) string {
	if json, err := json.Marshal(v); err == nil {
		return fmt.Sprintf("%T%s", v, string(json))
	} else {
		return err.Error()
	}
}

// InputFileFromFilepath generates an InputFile from given filepath
func InputFileFromFilepath(filepath string) InputFile {
	return InputFile{
		Filepath: &filepath,
	}
}

// InputFileFromUrl generates an InputFile from given url
func InputFileFromUrl(url string) InputFile {
	return InputFile{
		Url: &url,
	}
}

// InputFileFromBytes generates an InputFile from given bytes array
func InputFileFromBytes(bytes []byte) InputFile {
	return InputFile{
		Bytes: bytes,
	}
}

// InputFileFromFileId generates an InputFile from given file id
func InputFileFromFileId(fileId string) InputFile {
	return InputFile{
		FileId: &fileId,
	}
}
