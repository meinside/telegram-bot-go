// https://core.telegram.org/bots/api#available-types

package telegrambot

const (
	ParseModeMarkdown = "Markdown" // parse_mode: Markdown
)

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

// Webhook
type Webhook struct {
	UpdateId int      `json:"update_id"`
	Message  *Message `json:"message"`
}

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
	UpdateId int      `json:"update_id"`
	Message  *Message `json:"message,omitempty"`
}

// User
//
// https://core.telegram.org/bots/api#user
type User struct {
	Id        int     `json:"id"`
	FirstName *string `json:"first_name,omitempty"`
	LastName  *string `json:"last_name,omitempty"`
	Username  *string `json:"username,omitempty"`
}

// Chat
//
// https://core.telegram.org/bots/api#chat
type Chat struct {
	Id        int     `json:"id"`
	Type      *string `json:"type"` // 'private', 'group', or 'channel'
	Title     *string `json:"title,omitempty"`
	Username  *string `json:"username,omitempty"`
	FirstName *string `json:"first_name,omitempty"`
	LastName  *string `json:"last_name,omitempty"`
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
	MessageId           int         `json:"message_id"`
	From                *User       `json:"from"`
	Date                int         `json:"date"`
	Chat                *Chat       `json:"chat"`
	ForwardFrom         *User       `json:"forward_from,omitempty"`
	ForwardDate         int         `json:"forward_date,omitempty"`
	ReplyToMessage      *Message    `json:"reply_to_message,omitempty"`
	Text                *string     `json:"text,omitempty"`
	Audio               *Audio      `json:"audio,omitempty"`
	Document            *Document   `json:"document,omitempty"`
	Photo               []PhotoSize `json:"photo,omitempty"`
	Sticker             *Sticker    `json:"sticker,omitempty"`
	Video               *Video      `json:"video,omitempty"`
	Voice               *Voice      `json:"voice,omitempty"`
	Caption             *string     `json:"caption,omitempty"`
	Contact             *Contact    `json:"contact,omitempty"`
	Location            *Location   `json:"location,omitempty"`
	NewChatParticipant  *User       `json:"new_chat_participant,omitempty"`
	LeftChatParticipant *User       `json:"left_chat_participant,omitempty"`
	NewChatTitle        *string     `json:"new_chat_title,omitempty"`
	NewChatPhoto        []PhotoSize `json:"new_chat_photo,omitempty"`
	DeleteChatPhoto     bool        `json:"delete_chat_photo,omitempty"`
	GroupChatCreated    bool        `json:"group_chat_created,omitempty"`
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
