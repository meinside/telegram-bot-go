// https://core.telegram.org/bots/api#available-types
//
// following changes on 2015.10.08.

package telegrambot

const (
	ParseModeMarkdown = "Markdown"
)

type Webhook struct {
	UpdateId int     `json:"update_id"`
	Message  Message `json:"message"`
}

type ApiResult struct {
	Ok          bool        `json:"ok"`
	Description string      `json:"description,omitempty"`
	Result      interface{} `json:"result,omitempty"`
}

type ApiResultUser struct {
	Ok          bool   `json:"ok"`
	Description string `json:"description,omitempty"`
	Result      User   `json:"result,omitempty"`
}

type ApiResultMessage struct {
	Ok          bool    `json:"ok"`
	Description string  `json:"description,omitempty"`
	Result      Message `json:"result,omitempty"`
}

// https://core.telegram.org/bots/api#user
type User struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Username  string `json:"username,omitempty"`
}

// https://core.telegram.org/bots/api#chat
type Chat struct {
	Id        int    `json:"id"`
	Type      string `json:"type"` // 'private', 'group', or 'channel'
	Title     string `json:"title,omitempty"`
	Username  string `json:"username,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
}

// https://core.telegram.org/bots/api#audio
type Audio struct {
	FileId    string `json:"file_id"`
	Duration  int    `json:"duration"`
	Performer string `json:"performer,omitempty"`
	Title     string `json:"title,omitempty"`
	MimeType  string `json:"mime_type,omitempty"`
	FileSize  int    `json:"file_size,omitempty"`
}

// https://core.telegram.org/bots/api#photosize
type PhotoSize struct {
	FileId   string `json:"file_id"`
	Width    int    `json:"width"`
	Height   int    `json:"height"`
	FileSize int    `json:"file_size,omitempty"`
}

// https://core.telegram.org/bots/api#document
type Document struct {
	FileId   string    `json:"file_id"`
	Thumb    PhotoSize `json:"thumb,omitempty"`
	FileName string    `json:"file_name,omitempty"`
	MimeType string    `json:"mime_type,omitempty"`
	FileSize int       `json:"file_size,omitempty"`
}

// https://core.telegram.org/bots/api#sticker
type Sticker struct {
	FileId   string    `json:"file_id"`
	Width    int       `json:"width"`
	Height   int       `json:"height"`
	Thumb    PhotoSize `json:"thumb,omitempty"`
	FileSize int       `json:"file_size,omitempty"`
}

// https://core.telegram.org/bots/api#video
type Video struct {
	FileId   string    `json:"file_id"`
	Width    int       `json:"width"`
	Height   int       `json:"height"`
	Duration int       `json:"duration"`
	Thumb    PhotoSize `json:"thumb,omitempty"`
	MimeType string    `json:"mime_type,omitempty"`
	FileSize int       `json:"file_size,omitempty"`
}

// https://core.telegram.org/bots/api#voice
type Voice struct {
	FileId   string `json:"file_id"`
	Duration int    `json:"duration"`
	MimeType string `json:"mime_type,omitempty"`
	FileSize int    `json:"file_size,omitempty"`
}

// https://core.telegram.org/bots/api#contact
type Contact struct {
	PhoneNumber string `json:"phone_number"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name,omitempty"`
	UserId      int    `json:"user_id,omitempty"`
}

// https://core.telegram.org/bots/api#location
type Location struct {
	Longitude float32 `json:"longitude"`
	Latitude  float32 `json:"latitude"`
}

// https://core.telegram.org/bots/api#userprofilephotos
type UserProfilePhotos struct {
	TotalCount int           `json:"total_count"`
	Photos     [][]PhotoSize `json:"photos"`
}

// https://core.telegram.org/bots/api#file
type File struct {
	FileId   string `json:"file_id"`
	FileSize int    `json:"file_size,omitempty"`
	FilePath string `json:"file_path,omitempty"`
}

// https://core.telegram.org/bots/api#replykeyboardmarkup
type ReplyKeyboardMarkup struct {
	Keyboard        [][]string `json:"keyboard"`
	ResizeKeyboard  bool       `json:"resize_keyboard,omitempty"`
	OneTimeKeyboard bool       `json:"one_time_keyboard,omitempty"`
	Selective       bool       `json:"selective,omitempty"`
}

// https://core.telegram.org/bots/api#replykeyboardhide
type ReplyKeyboardHide struct {
	HideKeyboard bool `json:"hide_keyboard"`
	Selective    bool `json:"selective,omitempty"`
}

// https://core.telegram.org/bots/api#forcereply
type ForceReply struct {
	ForceReply bool `json:"force_reply"`
	Selective  bool `json:"selective,omitempty"`
}

// https://core.telegram.org/bots/api#message
type Message struct {
	MessageId           int         `json:"message_id"`
	From                User        `json:"from"`
	Date                int         `json:"date"`
	Chat                Chat        `json:"chat"`
	ForwardFrom         User        `json:"forward_from,omitempty"`
	ForwardDate         int         `json:"forward_date,omitempty"`
	ReplyToMessage      *Message    `json:"reply_to_message,omitempty"`
	Text                string      `json:"text,omitempty"`
	Audio               Audio       `json:"audio,omitempty"`
	Document            Document    `json:"document,omitempty"`
	Photo               []PhotoSize `json:"photo,omitempty"`
	Sticker             Sticker     `json:"sticker,omitempty"`
	Video               Video       `json:"video,omitempty"`
	Voice               Voice       `json:"voice,omitempty"`
	Caption             string      `json:"caption,omitempty"`
	Contact             Contact     `json:"contact,omitempty"`
	Location            Location    `json:"location,omitempty"`
	NewChatParticipant  User        `json:"new_chat_participant,omitempty"`
	LeftChatParticipant User        `json:"left_chat_participant,omitempty"`
	NewChatTitle        string      `json:"new_chat_title,omitempty"`
	NewChatPhoto        []PhotoSize `json:"new_chat_photo,omitempty"`
	DeleteChatPhoto     bool        `json:"delete_chat_photo,omitempty"`
	GroupChatCreated    bool        `json:"group_chat_created,omitempty"`
}

func (m *Message) HasForward() bool {
	return m.ForwardDate > 0
}

func (m *Message) HasReplyTo() bool {
	return m.ReplyToMessage != nil
}

func (m *Message) HasText() bool {
	return m.Text != ""
}

func (m *Message) HasAudio() bool {
	return m.Audio.FileId != ""
}

func (m *Message) HasDocument() bool {
	return m.Document.FileId != ""
}

func (m *Message) HasPhoto() bool {
	return len(m.Photo) > 0
}

func (m *Message) HasSticker() bool {
	return m.Sticker.FileId != ""
}

func (m *Message) HasVideo() bool {
	return m.Video.FileId != ""
}

func (m *Message) HasCaption() bool {
	return m.Caption != ""
}

func (m *Message) HasContact() bool {
	return m.Contact.FirstName != ""
}

func (m *Message) HasLocation() bool {
	return m.Location.Longitude > 0 && m.Location.Latitude > 0
}

func (m *Message) HasNewChatParticipant() bool {
	return m.NewChatParticipant.Id > 0
}

func (m *Message) HasLeftChatParticipant() bool {
	return m.LeftChatParticipant.Id > 0
}

func (m *Message) HasNewChatTitle() bool {
	return m.NewChatTitle != ""
}

func (m *Message) HasNewChatPhoto() bool {
	return len(m.NewChatPhoto) > 0
}

func (m *Message) HasDeleteChatPhoto() bool {
	return m.DeleteChatPhoto
}

func (m *Message) HasGroupChatCreated() bool {
	return m.GroupChatCreated
}
