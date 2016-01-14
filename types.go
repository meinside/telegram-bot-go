// https://core.telegram.org/bots/api#available-types
package telegrambot

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io"
)

// Chat types
type ChatType string

const (
	ChatTypePrivate ChatType = "private"
	ChatTypeGroup   ChatType = "group"
	ChatTypeChannel ChatType = "channel"
)

// Parse modes
type ParseMode string // parse_mode

const (
	ParseModeMarkdown ParseMode = "Markdown"
)

// Chat actions
type ChatAction string

const (
	ChatActionTyping         ChatAction = "typing"
	ChatActionUploadPhoto    ChatAction = "upload_photo"
	ChatActionRecordVideo    ChatAction = "record_video"
	ChatActionUploadVideo    ChatAction = "upload_video"
	ChatActionRecordAudio    ChatAction = "record_audio"
	ChatActionUploadAudio    ChatAction = "upload_audio"
	ChatActionUploadDocument ChatAction = "upload_document"
	ChatActionFindLocation   ChatAction = "find_location"
)

// Inline query result types
type InlineQueryResultType string

const (
	InlineQueryResultTypeArticle  InlineQueryResultType = "article"
	InlineQueryResultTypePhoto    InlineQueryResultType = "photo"
	InlineQueryResultTypeGif      InlineQueryResultType = "gif"
	InlineQueryResultTypeMpeg4Gif InlineQueryResultType = "mpeg4_gif"
	InlineQueryResultTypeVideo    InlineQueryResultType = "video"
)

// Video mime types
type VideoMimeType string

const (
	VideoMimeTypeHtml VideoMimeType = "text/html"
	VideoMimeTypeMp4  VideoMimeType = "video/mp4"
)

// API result
type ApiResult struct {
	Ok          bool        `json:"ok"`
	Description *string     `json:"description,omitempty"`
	Result      interface{} `json:"result,omitempty"`
}

// API result for User
type ApiResultUser struct {
	Ok          bool    `json:"ok"`
	Description *string `json:"description,omitempty"`
	Result      *User   `json:"result,omitempty"`
}

// API result for Message
type ApiResultMessage struct {
	Ok          bool     `json:"ok"`
	Description *string  `json:"description,omitempty"`
	Result      *Message `json:"result,omitempty"`
}

// API result for UserProfilePhotos
type ApiResultUserProfilePhotos struct {
	Ok          bool               `json:"ok"`
	Description *string            `json:"description,omitempty"`
	Result      *UserProfilePhotos `json:"result,omitempty"`
}

// API result for File
type ApiResultFile struct {
	Ok          bool    `json:"ok"`
	Description *string `json:"description,omitempty"`
	Result      *File   `json:"result,omitempty"`
}

// API result for Update
type ApiResultUpdates struct {
	Ok          bool     `json:"ok"`
	Description *string  `json:"description,omitempty"`
	Result      []Update `json:"result,omitempty"`
}

// Update
//
// https://core.telegram.org/bots/api#update
type Update struct {
	UpdateId           int                 `json:"update_id"`
	Message            *Message            `json:"message,omitempty"`
	InlineQuery        *InlineQuery        `json:"inline_query,omitempty"`
	ChosenInlineResult *ChosenInlineResult `json:"chosen_inline_result,omitempty"`
}

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

// User
//
// https://core.telegram.org/bots/api#user
type User struct {
	Id        int     `json:"id"`
	FirstName *string `json:"first_name"`
	LastName  *string `json:"last_name,omitempty"`
	Username  *string `json:"username,omitempty"`
}

func (u User) String() string {
	if json, err := json.Marshal(u); err == nil {
		return fmt.Sprintf("%T%s", u, string(json))
	}
	return fmt.Sprintf("%+v", u)
}

// Chat
//
// https://core.telegram.org/bots/api#chat
type Chat struct {
	Id        int       `json:"id"`
	Type      *ChatType `json:"type"`
	Title     *string   `json:"title,omitempty"`
	Username  *string   `json:"username,omitempty"`
	FirstName *string   `json:"first_name,omitempty"`
	LastName  *string   `json:"last_name,omitempty"`
}

func (c Chat) String() string {
	if json, err := json.Marshal(c); err == nil {
		return fmt.Sprintf("%T%s", c, string(json))
	}
	return fmt.Sprintf("%+v", c)
}

// Audio
//
// https://core.telegram.org/bots/api#audio
type Audio struct {
	FileId    *string `json:"file_id"`
	Duration  int     `json:"duration"`
	Performer *string `json:"performer,omitempty"`
	Title     *string `json:"title,omitempty"`
	MimeType  *string `json:"mime_type,omitempty"`
	FileSize  int     `json:"file_size,omitempty"`
}

// PhotoSize
//
// https://core.telegram.org/bots/api#photosize
type PhotoSize struct {
	FileId   *string `json:"file_id"`
	Width    int     `json:"width"`
	Height   int     `json:"height"`
	FileSize int     `json:"file_size,omitempty"`
}

// Document
//
// https://core.telegram.org/bots/api#document
type Document struct {
	FileId   *string    `json:"file_id"`
	Thumb    *PhotoSize `json:"thumb,omitempty"`
	FileName *string    `json:"file_name,omitempty"`
	MimeType *string    `json:"mime_type,omitempty"`
	FileSize int        `json:"file_size,omitempty"`
}

// Sticker
//
// https://core.telegram.org/bots/api#sticker
type Sticker struct {
	FileId   *string    `json:"file_id"`
	Width    int        `json:"width"`
	Height   int        `json:"height"`
	Thumb    *PhotoSize `json:"thumb,omitempty"`
	FileSize int        `json:"file_size,omitempty"`
}

// Video
//
// https://core.telegram.org/bots/api#video
type Video struct {
	FileId   *string    `json:"file_id"`
	Width    int        `json:"width"`
	Height   int        `json:"height"`
	Duration int        `json:"duration"`
	Thumb    *PhotoSize `json:"thumb,omitempty"`
	MimeType *string    `json:"mime_type,omitempty"`
	FileSize int        `json:"file_size,omitempty"`
}

// Voice
//
// https://core.telegram.org/bots/api#voice
type Voice struct {
	FileId   *string `json:"file_id"`
	Duration int     `json:"duration"`
	MimeType *string `json:"mime_type,omitempty"`
	FileSize int     `json:"file_size,omitempty"`
}

// Contact
//
// https://core.telegram.org/bots/api#contact
type Contact struct {
	PhoneNumber *string `json:"phone_number"`
	FirstName   *string `json:"first_name"`
	LastName    *string `json:"last_name,omitempty"`
	UserId      int     `json:"user_id,omitempty"`
}

// Location
//
// https://core.telegram.org/bots/api#location
type Location struct {
	Longitude float32 `json:"longitude"`
	Latitude  float32 `json:"latitude"`
}

// UserProfilePhotos
//
// https://core.telegram.org/bots/api#userprofilephotos
type UserProfilePhotos struct {
	TotalCount int           `json:"total_count"`
	Photos     [][]PhotoSize `json:"photos"`
}

// File
//
// https://core.telegram.org/bots/api#file
type File struct {
	FileId   *string `json:"file_id"`
	FileSize int     `json:"file_size,omitempty"`
	FilePath *string `json:"file_path,omitempty"`
}

// ReplyKeyboardMarkup
//
// https://core.telegram.org/bots/api#replykeyboardmarkup
type ReplyKeyboardMarkup struct {
	Keyboard        [][]string `json:"keyboard"`
	ResizeKeyboard  bool       `json:"resize_keyboard,omitempty"`
	OneTimeKeyboard bool       `json:"one_time_keyboard,omitempty"`
	Selective       bool       `json:"selective,omitempty"`
}

// ReplyKeyboardHide
//
// https://core.telegram.org/bots/api#replykeyboardhide
type ReplyKeyboardHide struct {
	HideKeyboard bool `json:"hide_keyboard"`
	Selective    bool `json:"selective,omitempty"`
}

// ForceReply
//
// https://core.telegram.org/bots/api#forcereply
type ForceReply struct {
	ForceReply bool `json:"force_reply"`
	Selective  bool `json:"selective,omitempty"`
}

// Message
//
// https://core.telegram.org/bots/api#message
type Message struct {
	MessageId             int         `json:"message_id"`
	From                  *User       `json:"from,omitempty"`
	Date                  int         `json:"date"`
	Chat                  *Chat       `json:"chat"`
	ForwardFrom           *User       `json:"forward_from,omitempty"`
	ForwardDate           int         `json:"forward_date,omitempty"`
	ReplyToMessage        *Message    `json:"reply_to_message,omitempty"`
	Text                  *string     `json:"text,omitempty"`
	Audio                 *Audio      `json:"audio,omitempty"`
	Document              *Document   `json:"document,omitempty"`
	Photo                 []PhotoSize `json:"photo,omitempty"`
	Sticker               *Sticker    `json:"sticker,omitempty"`
	Video                 *Video      `json:"video,omitempty"`
	Voice                 *Voice      `json:"voice,omitempty"`
	Caption               *string     `json:"caption,omitempty"`
	Contact               *Contact    `json:"contact,omitempty"`
	Location              *Location   `json:"location,omitempty"`
	NewChatParticipant    *User       `json:"new_chat_participant,omitempty"`
	LeftChatParticipant   *User       `json:"left_chat_participant,omitempty"`
	NewChatTitle          *string     `json:"new_chat_title,omitempty"`
	NewChatPhoto          []PhotoSize `json:"new_chat_photo,omitempty"`
	DeleteChatPhoto       bool        `json:"delete_chat_photo,omitempty"`
	GroupChatCreated      bool        `json:"group_chat_created,omitempty"`
	SupergroupChatCreated bool        `json:"supergroup_chat_created,omitempty"`
	ChannelChatCreated    bool        `json:"channel_chat_created,omitempty"`
	MigrateToChatId       int         `json:"migrate_to_chat_id,omitempty"`
	MigrateFromChatId     int         `json:"migrate_from_chat_id,omitempty"`
}

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

// Check if Message has NewChatParticipant.
func (m *Message) HasNewChatParticipant() bool {
	return m.NewChatParticipant != nil
}

// Check if Message has LeftChatParticipant.
func (m *Message) HasLeftChatParticipant() bool {
	return m.LeftChatParticipant != nil
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

// Inline query
//
// https://core.telegram.org/bots/api#inlinequery
type InlineQuery struct {
	Id     *string `json:"id"`
	From   *User   `json:"from"`
	Query  *string `json:"query"`
	Offset *string `json:"offset"`
}

func (i InlineQuery) String() string {
	if json, err := json.Marshal(i); err == nil {
		return fmt.Sprintf("%T%s", i, string(json))
	}
	return fmt.Sprintf("%+v", i)
}

// Chosen inline result
//
// https://core.telegram.org/bots/api#choseninlineresult
type ChosenInlineResult struct {
	ResultId *string `json:"result_id"`
	From     *User   `json:"from"`
	Query    *string `json:"query"`
}

func (c ChosenInlineResult) String() string {
	if json, err := json.Marshal(c); err == nil {
		return fmt.Sprintf("%T%s", c, string(json))
	}
	return fmt.Sprintf("%+v", c)
}

// Inline query results
//
// https://core.telegram.org/bots/api#inlinequeryresult
type InlineQueryResult struct {
	Type InlineQueryResultType `json:"type"`
	Id   *string               `json:"id"`
}
type InlineQueryResultArticle struct {
	InlineQueryResult
	Title                 *string    `json:"title"`
	MessageText           *string    `json:"message_text"`
	ParseMode             *ParseMode `json:"parse_mode,omitempty"`
	DisableWebPagePreview bool       `json:"disable_web_page_preview,omitempty"`
	Url                   *string    `json:"url,omitempty"`
	HideUrl               bool       `json:"hide_url,omitempty"`
	Description           *string    `json:"description,omitempty"`
	ThumbUrl              *string    `json:"thumb_url,omitempty"`
	ThumbWidth            int        `json:"thumb_width,omitempty"`
	ThumbHeight           int        `json:"thumb_height,omitempty"`
}
type InlineQueryResultPhoto struct {
	InlineQueryResult
	PhotoUrl              *string    `json:"photo_url"`
	PhotoWidth            int        `json:"photo_width,omitempty"`
	PhotoHeight           int        `json:"photo_height,omitempty"`
	ThumbUrl              *string    `json:"thumb_url"`
	Title                 *string    `json:"title,omitempty"`
	Description           *string    `json:"description,omitempty"`
	Caption               *string    `json:"caption,omitempty"`
	MessageText           *string    `json:"message_text,omitempty"`
	ParseMode             *ParseMode `json:"parse_mode,omitempty"`
	DisableWebPagePreview bool       `json:"disable_web_page_preview,omitempty"`
}
type InlineQueryResultGif struct {
	InlineQueryResult
	GifUrl                *string    `json:"gif_url"`
	GifWidth              int        `json:"gif_width,omitempty"`
	GifHeight             int        `json:"gif_height,omitempty"`
	ThumbUrl              *string    `json:"thumb_url"`
	Title                 *string    `json:"title,omitempty"`
	Caption               *string    `json:"caption,omitempty"`
	MessageText           *string    `json:"message_text,omitempty"`
	ParseMode             *ParseMode `json:"parse_mode,omitempty"`
	DisableWebPagePreview bool       `json:"disable_web_page_preview,omitempty"`
}
type InlineQueryResultMpeg4Gif struct {
	InlineQueryResult
	Mpeg4Url              *string    `json:"mpeg4_url"`
	Mpeg4Width            int        `json:"mpeg4_width,omitempty"`
	Mpeg4Height           int        `json:"mpeg4_height,omitempty"`
	ThumbUrl              *string    `json:"thumb_url"`
	Title                 *string    `json:"title,omitempty"`
	Caption               *string    `json:"caption,omitempty"`
	MessageText           *string    `json:"message_text,omitempty"`
	ParseMode             *ParseMode `json:"parse_mode,omitempty"`
	DisableWebPagePreview bool       `json:"disable_web_page_preview,omitempty"`
}
type InlineQueryResultVideo struct {
	InlineQueryResult
	VideoUrl              *string        `json:"video_url"`
	MimeType              *VideoMimeType `json:"mime_type"`
	MessageText           *string        `json:"message_text"`
	ParseMode             *ParseMode     `json:"parse_mode,omitempty"`
	DisableWebPagePreview bool           `json:"disable_web_page_preview,omitempty"`
	VideoWidth            int            `json:"video_width,omitempty"`
	VideoHeight           int            `json:"video_height,omitempty"`
	VideoDuration         int            `json:"video_duration,omitempty"`
	ThumbUrl              *string        `json:"thumb_url"`
	Title                 *string        `json:"title"`
	Description           *string        `json:"description,omitempty"`
}

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
			Title:       &title,
			MessageText: &messageText,
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
func NewInlineQueryResultVideo(videoUrl, thumbUrl, title, messageText string, mimeType VideoMimeType) (newVideo *InlineQueryResultVideo, generatedId *string) {
	if id, err := newUUID(); err == nil {
		return &InlineQueryResultVideo{
			InlineQueryResult: InlineQueryResult{
				Type: InlineQueryResultTypeVideo,
				Id:   &id,
			},
			VideoUrl:    &videoUrl,
			MimeType:    &mimeType,
			MessageText: &messageText,
			ThumbUrl:    &thumbUrl,
			Title:       &title,
		}, &id
	}

	return &InlineQueryResultVideo{}, nil
}
