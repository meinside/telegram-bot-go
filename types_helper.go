package telegrambot

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io"
)

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

// Helper functions for User
//

// String function for User
func (u User) String() string {
	if json, err := json.Marshal(u); err == nil {
		return fmt.Sprintf("%T%s", u, string(json))
	}
	return fmt.Sprintf("%+v", u)
}

// Helper functions for Chat
//

// String function for Chat
func (c Chat) String() string {
	if json, err := json.Marshal(c); err == nil {
		return fmt.Sprintf("%T%s", c, string(json))
	}
	return fmt.Sprintf("%+v", c)
}

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
	return m.ForwardDate > 0
}

// Check if Message has ReplyTo.
func (m *Message) HasReplyTo() bool {
	return m.ReplyToMessage != nil
}

// Check if Message has Text.
func (m *Message) HasText() bool {
	return m.Text != nil
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

// Helper functions for InlineQuery
//

// String function for InlineQuery
func (i InlineQuery) String() string {
	if json, err := json.Marshal(i); err == nil {
		return fmt.Sprintf("%T%s", i, string(json))
	}
	return fmt.Sprintf("%+v", i)
}

// Helper functions for ChosenInlineResult
//

// String function for ChosenInlineResult
func (c ChosenInlineResult) String() string {
	if json, err := json.Marshal(c); err == nil {
		return fmt.Sprintf("%T%s", c, string(json))
	}
	return fmt.Sprintf("%+v", c)
}
