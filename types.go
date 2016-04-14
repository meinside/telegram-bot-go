// https://core.telegram.org/bots/api#available-types
package telegrambot

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
	ParseModeHtml     ParseMode = "HTML"
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
	InlineQueryResultTypeAudio    InlineQueryResultType = "audio"
	InlineQueryResultTypeVoice    InlineQueryResultType = "voice"
	InlineQueryResultTypeDocument InlineQueryResultType = "document"
	InlineQueryResultTypeLocation InlineQueryResultType = "location"
	InlineQueryResultTypeVenue    InlineQueryResultType = "venue"
	InlineQueryResultTypeContact  InlineQueryResultType = "contact"
	InlineQueryResultTypeSticker  InlineQueryResultType = "sticker"
)

// Message Entity Types
//
// https://core.telegram.org/bots/api#messageentity
type MessageEntityType string

const (
	MessageEntityTypeMention    = "mention"
	MessageEntityTypeHashTag    = "hashtag"
	MessageEntityTypeBotCommand = "bot_command"
	MessageEntityTypeUrl        = "url"
	MessageEntityTypeEmail      = "email"
	MessageEntityTypeBold       = "bold"
	MessageEntityTypeItalic     = "italic"
	MessageEntityTypeCode       = "code"
	MessageEntityTypePre        = "pre"
	MessageEntityTypeTextLink   = "text_link"
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
	CallbackQuery      *CallbackQuery      `json:"callback_query,omitempty"`
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

// MessageEntity
//
// https://core.telegram.org/bots/api#messageentity
type MessageEntity struct {
	Type   MessageEntityType `json:"type"`
	Offset int               `json:"offset"`
	Length int               `json:"length"`
	Url    *string           `json:"url,omitempty"`
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

// Venue
//
// https://core.telegram.org/bots/api#venue
type Venue struct {
	Location     *Location `json:"location"`
	Title        *string   `json:"title"`
	Address      *string   `json:"address"`
	FoursquareId *string   `json:"foursquare_id,omitempty"`
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
	Keyboard        [][]KeyboardButton `json:"keyboard"`
	ResizeKeyboard  bool               `json:"resize_keyboard,omitempty"`
	OneTimeKeyboard bool               `json:"one_time_keyboard,omitempty"`
	Selective       bool               `json:"selective,omitempty"`
}

// KeyboardButton
//
// https://core.telegram.org/bots/api#keyboardbutton
type KeyboardButton struct {
	Text            *string `json:"text"`
	RequestContact  bool    `json:"request_contact,omitempty"`
	RequestLocation bool    `json:"request_location,omitempty"`
}

// ReplyKeyboardHide
//
// https://core.telegram.org/bots/api#replykeyboardhide
type ReplyKeyboardHide struct {
	HideKeyboard bool `json:"hide_keyboard"`
	Selective    bool `json:"selective,omitempty"`
}

// InlineKeyboardMarkup
//
// https://core.telegram.org/bots/api#inlinekeyboardmarkup
type InlineKeyboardMarkup struct {
	InlineKeyboard [][]InlineKeyboardButton `json:"inline_keyboard"`
}

// InlineKeyboardButton
//
// https://core.telegram.org/bots/api#inlinekeyboardbutton
type InlineKeyboardButton struct {
	Text              *string `json:"text"`
	Url               *string `json:"url,omitempty"`
	CallbackData      *string `json:"callback_data,omitempty"`
	SwitchInlineQuery *string `json:"switch_inline_query,omitempty"`
}

// CallbackQuery
//
// https://core.telegram.org/bots/api#callbackquery
type CallbackQuery struct {
	Id              *string  `json:"id"`
	From            *User    `json:"from"`
	Message         *Message `json:"message,omitempty"`
	InlineMessageId *string  `json:"inline_message_id,omitempty"`
	Data            *string  `json:"data"`
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
	MessageId             int             `json:"message_id"`
	From                  *User           `json:"from,omitempty"`
	Date                  int             `json:"date"`
	Chat                  *Chat           `json:"chat"`
	ForwardFrom           *User           `json:"forward_from,omitempty"`
	ForwardDate           int             `json:"forward_date,omitempty"`
	ReplyToMessage        *Message        `json:"reply_to_message,omitempty"`
	Text                  *string         `json:"text,omitempty"`
	Entities              []MessageEntity `json:"entities,omitempty"`
	Audio                 *Audio          `json:"audio,omitempty"`
	Document              *Document       `json:"document,omitempty"`
	Photo                 []PhotoSize     `json:"photo,omitempty"`
	Sticker               *Sticker        `json:"sticker,omitempty"`
	Video                 *Video          `json:"video,omitempty"`
	Voice                 *Voice          `json:"voice,omitempty"`
	Caption               *string         `json:"caption,omitempty"`
	Contact               *Contact        `json:"contact,omitempty"`
	Location              *Location       `json:"location,omitempty"`
	Venue                 *Venue          `json:"venue,omitempty"`
	NewChatMember         *User           `json:"new_chat_member,omitempty"`
	LeftChatMember        *User           `json:"left_chat_member,omitempty"`
	NewChatTitle          *string         `json:"new_chat_title,omitempty"`
	NewChatPhoto          []PhotoSize     `json:"new_chat_photo,omitempty"`
	DeleteChatPhoto       bool            `json:"delete_chat_photo,omitempty"`
	GroupChatCreated      bool            `json:"group_chat_created,omitempty"`
	SupergroupChatCreated bool            `json:"supergroup_chat_created,omitempty"`
	ChannelChatCreated    bool            `json:"channel_chat_created,omitempty"`
	MigrateToChatId       int             `json:"migrate_to_chat_id,omitempty"`
	MigrateFromChatId     int             `json:"migrate_from_chat_id,omitempty"`
	PinnedMessage         *Message        `json:"pinned_message,omitempty"`
}

// Inline query
//
// https://core.telegram.org/bots/api#inlinequery
type InlineQuery struct {
	Id       *string   `json:"id"`
	From     *User     `json:"from"`
	Location *Location `json:"location,omitempty"`
	Query    *string   `json:"query"`
	Offset   *string   `json:"offset"`
}

// Chosen inline result
//
// https://core.telegram.org/bots/api#choseninlineresult
type ChosenInlineResult struct {
	ResultId        *string   `json:"result_id"`
	From            *User     `json:"from"`
	Location        *Location `json:"location,omitempty"`
	InlineMessageId *string   `json:"inline_message_id,omitempty"`
	Query           *string   `json:"query"`
}

// Video mime types for inline query
type VideoMimeType string

const (
	VideoMimeTypeHtml VideoMimeType = "text/html"
	VideoMimeTypeMp4  VideoMimeType = "video/mp4"
)

// Document mime types for inline query
type DocumentMimeType string

const (
	DocumentMimeTypePdf DocumentMimeType = "application/pdf"
	DocumentMimeTypeZip DocumentMimeType = "application/zip"
)

// Inline query results
//
// https://core.telegram.org/bots/api#inlinequeryresult
type InlineQueryResult struct {
	Type InlineQueryResultType `json:"type"`
	Id   *string               `json:"id"`
}
type InlineQueryResultArticle struct { // https://core.telegram.org/bots/api#inlinequeryresultarticle
	InlineQueryResult
	Title               *string               `json:"title"`
	InputMessageContent InputMessageContent   `json:"input_message_content"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	Url                 *string               `json:"url,omitempty"`
	HideUrl             bool                  `json:"hide_url,omitempty"`
	Description         *string               `json:"description,omitempty"`
	ThumbUrl            *string               `json:"thumb_url,omitempty"`
	ThumbWidth          int                   `json:"thumb_width,omitempty"`
	ThumbHeight         int                   `json:"thumb_height,omitempty"`
}
type InlineQueryResultPhoto struct { // https://core.telegram.org/bots/api#inlinequeryresultphoto
	InlineQueryResult
	PhotoUrl            *string               `json:"photo_url"`
	PhotoWidth          int                   `json:"photo_width,omitempty"`
	PhotoHeight         int                   `json:"photo_height,omitempty"`
	ThumbUrl            *string               `json:"thumb_url"`
	Title               *string               `json:"title,omitempty"`
	Description         *string               `json:"description,omitempty"`
	Caption             *string               `json:"caption,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent   `json:"input_message_content,omitempty"`
}
type InlineQueryResultGif struct { // https://core.telegram.org/bots/api#inlinequeryresultgif
	InlineQueryResult
	GifUrl              *string               `json:"gif_url"`
	GifWidth            int                   `json:"gif_width,omitempty"`
	GifHeight           int                   `json:"gif_height,omitempty"`
	ThumbUrl            *string               `json:"thumb_url"`
	Title               *string               `json:"title,omitempty"`
	Caption             *string               `json:"caption,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent   `json:"input_message_content,omitempty"`
}
type InlineQueryResultMpeg4Gif struct { // https://core.telegram.org/bots/api#inlinequeryresultmpeg4gif
	InlineQueryResult
	Mpeg4Url            *string               `json:"mpeg4_url"`
	Mpeg4Width          int                   `json:"mpeg4_width,omitempty"`
	Mpeg4Height         int                   `json:"mpeg4_height,omitempty"`
	ThumbUrl            *string               `json:"thumb_url"`
	Title               *string               `json:"title,omitempty"`
	Caption             *string               `json:"caption,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent   `json:"input_message_content,omitempty"`
}
type InlineQueryResultVideo struct { // https://core.telegram.org/bots/api#inlinequeryresultvideo
	InlineQueryResult
	VideoUrl            *string               `json:"video_url"`
	MimeType            *VideoMimeType        `json:"mime_type"`
	ThumbUrl            *string               `json:"thumb_url"`
	Title               *string               `json:"title"`
	Caption             *string               `json:"caption,omitempty"`
	VideoWidth          int                   `json:"video_width,omitempty"`
	VideoHeight         int                   `json:"video_height,omitempty"`
	VideoDuration       int                   `json:"video_duration,omitempty"`
	Description         *string               `json:"description,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent   `json:"input_message_content,omitempty"`
}
type InlineQueryResultAudio struct { // https://core.telegram.org/bots/api#inlinequeryresultaudio
	InlineQueryResult
	AudioUrl            *string               `json:"audio_url"`
	Title               *string               `json:"title"`
	Performer           *string               `json:"performer,omitempty"`
	AudioDuration       int                   `json:"audio_duration,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent   `json:"input_message_content,omitempty"`
}
type InlineQueryResultVoice struct { // https://core.telegram.org/bots/api#inlinequeryresultvoice
	InlineQueryResult
	VoiceUrl            *string               `json:"voice_url"`
	Title               *string               `json:"title"`
	VoiceDuration       int                   `json:"voice_duration,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent   `json:"input_message_content,omitempty"`
}
type InlineQueryResultDocument struct { // https://core.telegram.org/bots/api#inlinequeryresultdocument
	InlineQueryResult
	Title               *string               `json:"title"`
	Caption             *string               `json:"caption,omitempty"`
	DocumentUrl         *string               `json:"document_url,omitempty"`
	MimeType            *DocumentMimeType     `json:"mime_type"`
	Description         *string               `json:"description,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent   `json:"input_message_content,omitempty"`
	ThumbUrl            *string               `json:"thumb_url,omitempty"`
	ThumbWidth          int                   `json:"thumb_width,omitempty"`
	ThumbHeight         int                   `json:"thumb_height,omitempty"`
}
type InlineQueryResultLocation struct { // https://core.telegram.org/bots/api#inlinequeryresultlocation
	InlineQueryResult
	Latitude            float32               `json:"latitude"`
	Longitude           float32               `json:"longitude"`
	Title               *string               `json:"title"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent   `json:"input_message_content,omitempty"`
	ThumbUrl            *string               `json:"thumb_url,omitempty"`
	ThumbWidth          int                   `json:"thumb_width,omitempty"`
	ThumbHeight         int                   `json:"thumb_height,omitempty"`
}
type InlineQueryResultVenue struct { // https://core.telegram.org/bots/api#inlinequeryresultvenue
	InlineQueryResult
	Latitude            float32               `json:"latitude"`
	Longitude           float32               `json:"longitude"`
	Title               *string               `json:"title"`
	Address             *string               `json:"address"`
	FoursquareId        *string               `json:"foursquare_id,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent   `json:"input_message_content,omitempty"`
	ThumbUrl            *string               `json:"thumb_url,omitempty"`
	ThumbWidth          int                   `json:"thumb_width,omitempty"`
	ThumbHeight         int                   `json:"thumb_height,omitempty"`
}
type InlineQueryResultContact struct { // https://core.telegram.org/bots/api#inlinequeryresultcontact
	InlineQueryResult
	PhoneNumber         *string               `json:"phone_number"`
	FirstName           *string               `json:"first_name"`
	LastName            *string               `json:"last_name,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent   `json:"input_message_content,omitempty"`
	ThumbUrl            *string               `json:"thumb_url,omitempty"`
	ThumbWidth          int                   `json:"thumb_width,omitempty"`
	ThumbHeight         int                   `json:"thumb_height,omitempty"`
}
type InlineQueryResultCachedPhoto struct { // https://core.telegram.org/bots/api#inlinequeryresultcachedphoto
	InlineQueryResult
	PhotoFileId         *string               `json:"photo_file_id"`
	Title               *string               `json:"title,omitempty"`
	Description         *string               `json:"description,omitempty"`
	Caption             *string               `json:"caption,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent   `json:"input_message_content,omitempty"`
}
type InlineQueryResultCachedGif struct { // https://core.telegram.org/bots/api#inlinequeryresultcachedgif
	InlineQueryResult
	GifFileId           *string               `json:"gif_file_id"`
	Title               *string               `json:"title,omitempty"`
	Caption             *string               `json:"caption,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent   `json:"input_message_content,omitempty"`
}
type InlineQueryResultCachedMpeg4Gif struct { // https://core.telegram.org/bots/api#inlinequeryresultcachedmpeg4gif
	InlineQueryResult
	Mpeg4FileId         *string               `json:"mpeg4_file_id"`
	Title               *string               `json:"title,omitempty"`
	Caption             *string               `json:"caption,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent   `json:"input_message_content,omitempty"`
}
type InlineQueryResultCachedSticker struct { // https://core.telegram.org/bots/api#inlinequeryresultcachedsticker
	InlineQueryResult
	StickerFileId       *string               `json:"sticker_file_id"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent   `json:"input_message_content,omitempty"`
}
type InlineQueryResultCachedDocument struct { // https://core.telegram.org/bots/api#inlinequeryresultcacheddocument
	InlineQueryResult
	Title               *string               `json:"title"`
	DocumentFileId      *string               `json:"document_file_id"`
	Description         *string               `json:"description,omitempty"`
	Caption             *string               `json:"caption,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent   `json:"input_message_content,omitempty"`
}
type InlineQueryResultCachedVideo struct { // https://core.telegram.org/bots/api#inlinequeryresultcachedvideo
	InlineQueryResult
	VideoFileId         *string               `json:"video_file_id"`
	Title               *string               `json:"title"`
	Description         *string               `json:"description,omitempty"`
	Caption             *string               `json:"caption,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent   `json:"input_message_content,omitempty"`
}
type InlineQueryResultCachedVoice struct { // https://core.telegram.org/bots/api#inlinequeryresultcachedvoice
	InlineQueryResult
	VoiceFileId         *string               `json:"voice_file_id"`
	Title               *string               `json:"title"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent   `json:"input_message_content,omitempty"`
}
type InlineQueryResultCachedAudio struct { // https://core.telegram.org/bots/api#inlinequeryresultcachedaudio
	InlineQueryResult
	AudioFileId         *string               `json:"audio_file_id"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent   `json:"input_message_content,omitempty"`
}

// InputMessageContent
//
// https://core.telegram.org/bots/api#inputmessagecontent
type InputMessageContent interface{}
type InputTextMessageContent struct { // https://core.telegram.org/bots/api#inputtextmessagecontent
	MessageText           *string    `json:"message_text"`
	ParseMode             *ParseMode `json:"parse_mode,omitempty"`
	DisableWebPagePreview bool       `json:"disable_web_page_preview,omitempty"`
}
type InputLocationMessageContent struct { // https://core.telegram.org/bots/api#inputlocationmessagecontent
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
}
type InputVenueMessageContent struct { // https://core.telegram.org/bots/api#inputvenuemessagecontent
	Latitude     float32 `json:"latitude"`
	Longitude    float32 `json:"longitude"`
	Title        *string `json:"title"`
	Address      *string `json:"address"`
	FoursquareId *string `json:"foursquare_id,omitempty"`
}
type InputContactMessageContent struct { // https://core.telegram.org/bots/api#inputcontactmessagecontent
	PhoneNumber *string `json:"phone_number"`
	FirstName   *string `json:"first_name"`
	LastName    *string `json:"last_name,omitempty"`
}
