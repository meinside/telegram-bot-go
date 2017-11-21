package telegrambot

// https://core.telegram.org/bots/api#available-types

// ChatId can be `Message.Chat.Id`,
// or target channel name (in string, eg. "@channelusername")
type ChatId interface{}

// ChatType is a type of Chat
type ChatType string

// ChatType strings
const (
	ChatTypePrivate ChatType = "private"
	ChatTypeGroup   ChatType = "group"
	ChatTypeChannel ChatType = "channel"
)

// ParseMode is a mode of parse
type ParseMode string // parse_mode

// ParseMode strings
const (
	ParseModeMarkdown ParseMode = "Markdown"
	ParseModeHtml     ParseMode = "HTML"
)

// ChatAction is a type of action in chats
type ChatAction string

// ChatAction strings
const (
	ChatActionTyping          ChatAction = "typing"
	ChatActionUploadPhoto     ChatAction = "upload_photo"
	ChatActionRecordVideo     ChatAction = "record_video"
	ChatActionUploadVideo     ChatAction = "upload_video"
	ChatActionRecordAudio     ChatAction = "record_audio"
	ChatActionUploadAudio     ChatAction = "upload_audio"
	ChatActionUploadDocument  ChatAction = "upload_document"
	ChatActionFindLocation    ChatAction = "find_location"
	ChatActionRecordVideoNote ChatAction = "record_video_note"
	ChatActionUploadVideoNote ChatAction = "upload_video_note"
)

// InlineQueryResultType is a type of inline query result
type InlineQueryResultType string

// InlineQueryResultType strings
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
	InlineQueryResultTypeGame     InlineQueryResultType = "game"
)

// MessageEntityType is a type of MessageEntity
//
// https://core.telegram.org/bots/api#messageentity
type MessageEntityType string

// MessageEntityType strings
const (
	MessageEntityTypeMention     = "mention"
	MessageEntityTypeHashTag     = "hashtag"
	MessageEntityTypeBotCommand  = "bot_command"
	MessageEntityTypeUrl         = "url"
	MessageEntityTypeEmail       = "email"
	MessageEntityTypeBold        = "bold"
	MessageEntityTypeItalic      = "italic"
	MessageEntityTypeCode        = "code"
	MessageEntityTypePre         = "pre"
	MessageEntityTypeTextLink    = "text_link"
	MessageEntityTypeTextMention = "text_mention"
)

// ChatMemberStatus is a status of chat member
//
// https://core.telegram.org/bots/api#chatmember
type ChatMemberStatus string

// ChatMemberStatus strings
const (
	ChatMemberStatusCreator       ChatMemberStatus = "creator"
	ChatMemberStatusAdministrator ChatMemberStatus = "administrator"
	ChatMemberStatusMember        ChatMemberStatus = "member"
	ChatMemberStatusLeft          ChatMemberStatus = "left"
	ChatMemberStatusKicked        ChatMemberStatus = "kicked"
)

// MaskPositionPoint is a point in MaskPosition
//
// https://core.telegram.org/bots/api#maskposition
type MaskPositionPoint string

// MaskPosition points
const (
	MaskPositionForehead MaskPositionPoint = "forehead"
	MaskPositionEyes     MaskPositionPoint = "eyes"
	MaskPositionMouth    MaskPositionPoint = "mouth"
	MaskPositionChin     MaskPositionPoint = "chin"
)

// ApiResponseBase is a base of API responses
type ApiResponseBase struct {
	Ok          bool                   `json:"ok"`
	Description *string                `json:"description,omitempty"`
	Parameters  *ApiResponseParameters `json:"parameters,omitempty"`
}

// ApiResponseParameters is parameters in API responses
//
// https://core.telegram.org/bots/api#responseparameters
type ApiResponseParameters struct {
	MigrateToChatId int64 `json:"migrate_to_chat_id,omitempty"`
	RetryAfter      int   `json:"retry_after,omitempty"`
}

// ApiResponse is an API response
type ApiResponse struct {
	ApiResponseBase
	Result interface{} `json:"result,omitempty"`
}

// ApiResponseWebhookInfo is an API response with result type: WebhookInfo
type ApiResponseWebhookInfo struct {
	ApiResponseBase
	Result *WebhookInfo `json:"result,omitempty"`
}

// ApiResponseUser is an API response with result type: User
type ApiResponseUser struct {
	ApiResponseBase
	Result *User `json:"result,omitempty"`
}

// ApiResponseMessage is an API response with result type: Message
type ApiResponseMessage struct {
	ApiResponseBase
	Result *Message `json:"result,omitempty"`
}

// ApiResponseUserProfilePhotos is an API response with result type: UserProfilePhotos
type ApiResponseUserProfilePhotos struct {
	ApiResponseBase
	Result *UserProfilePhotos `json:"result,omitempty"`
}

// ApiResponseFile is an API response with result type: File
type ApiResponseFile struct {
	ApiResponseBase
	Result *File `json:"result,omitempty"`
}

// API response with result type: Update
type ApiResponseUpdates struct {
	ApiResponseBase
	Result []Update `json:"result,omitempty"`
}

// ApiResponseChat is an API response with result type: Chat
type ApiResponseChat struct {
	ApiResponseBase
	Result *Chat `json:"result,omitempty"`
}

// ApiResponseChatAdministrators is an API response with result type: ChatAdministrators
type ApiResponseChatAdministrators struct {
	ApiResponseBase
	Result []ChatMember `json:"result,omitempty"`
}

// ApiResponseChatMember is an API response with result type: ChatMember
type ApiResponseChatMember struct {
	ApiResponseBase
	Result *ChatMember `json:"result,omitempty"`
}

// ApiResponseInt is an API response with result type: int
type ApiResponseInt struct {
	ApiResponseBase
	Result int `json:"result,omitempty"`
}

// ApiResponseString is an API response with result type: string
type ApiResponseString struct {
	ApiResponseBase
	Result *string `json:"result,omitempty"`
}

// API response with result type: GameHighScores
type ApiResponseGameHighScores struct {
	ApiResponseBase
	Result []GameHighScore `json:"result,omitempty"`
}

// ApiResponseStickerSet is an API response with result type: StickerSet
type ApiResponseStickerSet struct {
	ApiResponseBase
	Result *StickerSet `json:"result,omitempty"`
}

// UpdateType is a type of updates (for allowed_updates)
//
// https://core.telegram.org/bots/api#setwebhook
// https://core.telegram.org/bots/api#update
type UpdateType string

// UpdateType strings
const (
	UpdateTypeMessage            UpdateType = "message"
	UpdateTypeEditedMessage      UpdateType = "edited_message"
	UpdateTypeChannelPost        UpdateType = "channel_post"
	UpdateTypeEditedChannelPost  UpdateType = "edited_channel_post"
	UpdateTypeInlineQuery        UpdateType = "inline_query"
	UpdateTypeChosenInlineResult UpdateType = "chosen_inline_result"
	UpdateTypeCallbackQuery      UpdateType = "callback_query"
)

// WebhookInfo is a struct of webhook info
//
// https://core.telegram.org/bots/api#webhookinfo
type WebhookInfo struct {
	Url                  *string      `json:"url"`
	HasCustomCertificate bool         `json:"has_custom_certificate"`
	PendingUpdateCount   int          `json:"pending_update_count"`
	LastErrorDate        int          `json:"last_error_date,omitempty"`
	LastErrorMessage     *string      `json:"last_error_message,omitempty"`
	MaxConnections       int          `json:"max_connections,omitempty"`
	AllowedUpdates       []UpdateType `json:"allowed_updates,omitempty"`
}

// Update is a struct of an update
//
// https://core.telegram.org/bots/api#update
type Update struct {
	UpdateId           int                 `json:"update_id"`
	Message            *Message            `json:"message,omitempty"`
	EditedMessage      *Message            `json:"edited_message,omitempty"`
	ChannelPost        *Message            `json:"channel_post,omitempty"`
	EditedChannelPost  *Message            `json:"edited_channel_post,omitempty"`
	InlineQuery        *InlineQuery        `json:"inline_query,omitempty"`
	ChosenInlineResult *ChosenInlineResult `json:"chosen_inline_result,omitempty"`
	CallbackQuery      *CallbackQuery      `json:"callback_query,omitempty"`
	ShippingQuery      *ShippingQuery      `json:"shipping_query,omitempty"`
	PreCheckoutQuery   *PreCheckoutQuery   `json:"pre_checkout_query,omitempty"`
}

// User is a struct of a user
//
// https://core.telegram.org/bots/api#user
type User struct {
	Id           int     `json:"id"`
	IsBot        bool    `json:"is_bot"`
	FirstName    string  `json:"first_name"`
	LastName     *string `json:"last_name,omitempty"`
	Username     *string `json:"username,omitempty"`
	LanguageCode *string `json:"language_code,omitempty"` // https://en.wikipedia.org/wiki/IETF_language_tag
}

// Chat is a struct of a chat
//
// https://core.telegram.org/bots/api#chat
type Chat struct {
	Id                          int64      `json:"id"`
	Type                        ChatType   `json:"type"`
	Title                       *string    `json:"title,omitempty"`
	Username                    *string    `json:"username,omitempty"`
	FirstName                   *string    `json:"first_name,omitempty"`
	LastName                    *string    `json:"last_name,omitempty"`
	AllMembersAreAdministrators bool       `json:"all_members_are_administrators,omitempty"`
	Photo                       *ChatPhoto `json:"photo,omitempty"`
	Description                 *string    `json:"description,omitempty"`
	InviteLink                  *string    `json:"invite_link,omitempty"`
	PinnedMessage               *Message   `json:"pinned_message,omitempty"`
	StickerSetName              *string    `json:"sticker_set_name,omitempty"`
	CanSetStickerSet            bool       `json:"can_set_sticker_set,omitempty"`
}

// InputFile represents contents of a file to be uploaded.
// Can be generated with InputFileFromXXX() functions in types_helper.go
//
// https://core.telegram.org/bots/api#inputfile
type InputFile struct {
	Filepath *string
	Url      *string
	Bytes    []byte
	FileId   *string
}

// Audio is a struct for an audio file
//
// https://core.telegram.org/bots/api#audio
type Audio struct {
	FileId    string  `json:"file_id"`
	Duration  int     `json:"duration"`
	Performer *string `json:"performer,omitempty"`
	Title     *string `json:"title,omitempty"`
	MimeType  *string `json:"mime_type,omitempty"`
	FileSize  int     `json:"file_size,omitempty"`
}

// MessageEntity is a struct of a message entity
//
// https://core.telegram.org/bots/api#messageentity
type MessageEntity struct {
	Type   MessageEntityType `json:"type"`
	Offset int               `json:"offset"`
	Length int               `json:"length"`
	Url    *string           `json:"url,omitempty"`  // for Type == "text_link" only,
	User   *User             `json:"user,omitempty"` // for Type == "text_mention" only,
}

// PhotoSize is a struct of a photo's size
//
// https://core.telegram.org/bots/api#photosize
type PhotoSize struct {
	FileId   string `json:"file_id"`
	Width    int    `json:"width"`
	Height   int    `json:"height"`
	FileSize int    `json:"file_size,omitempty"`
}

// Document is a struct for an ordinary file
//
// https://core.telegram.org/bots/api#document
type Document struct {
	FileId   string     `json:"file_id"`
	Thumb    *PhotoSize `json:"thumb,omitempty"`
	FileName *string    `json:"file_name,omitempty"`
	MimeType *string    `json:"mime_type,omitempty"`
	FileSize int        `json:"file_size,omitempty"`
}

// Sticker is a struct of a sticker
//
// https://core.telegram.org/bots/api#sticker
type Sticker struct {
	FileId       string        `json:"file_id"`
	Width        int           `json:"width"`
	Height       int           `json:"height"`
	Thumb        *PhotoSize    `json:"thumb,omitempty"`
	Emoji        *string       `json:"emoji,omitempty"`
	SetName      *string       `json:"set_name,omitempty"`
	MaskPosition *MaskPosition `json:"mask_position,omitempty"`
	FileSize     int           `json:"file_size,omitempty"`
}

// StickerSet is a struct of a sticker set
//
// https://core.telegram.org/bots/api#stickerset
type StickerSet struct {
	Name          *string   `json:"name"`
	Title         *string   `json:"title"`
	ContainsMasks bool      `json:"contains_masks"`
	Stickers      []Sticker `json:"stickers"`
}

// MaskPosition is a struct for a mask position
//
// https://core.telegram.org/bots/api#maskposition
type MaskPosition struct {
	Point  MaskPositionPoint `json:"point"`
	XShift float32           `json:"x_shift"`
	YShift float32           `json:"y_shift"`
	Scale  float32           `json:"scale"`
}

// Video is a struct for a video file
//
// https://core.telegram.org/bots/api#video
type Video struct {
	FileId   string     `json:"file_id"`
	Width    int        `json:"width"`
	Height   int        `json:"height"`
	Duration int        `json:"duration"`
	Thumb    *PhotoSize `json:"thumb,omitempty"`
	MimeType *string    `json:"mime_type,omitempty"`
	FileSize int        `json:"file_size,omitempty"`
}

// Voice is a struct for a voice file
//
// https://core.telegram.org/bots/api#voice
type Voice struct {
	FileId   string  `json:"file_id"`
	Duration int     `json:"duration"`
	MimeType *string `json:"mime_type,omitempty"`
	FileSize int     `json:"file_size,omitempty"`
}

// VideoNote is a struct for a video note
//
// https://core.telegram.org/bots/api#videonote
type VideoNote struct {
	FileId   string     `json:"file_id"`
	Length   int        `json:"length"`
	Duration int        `json:"duration"`
	Thumb    *PhotoSize `json:"thumb,omitempty"`
	FileSize int        `json:"file_size,omitempty"`
}

// Contact is a struct for a contact info
//
// https://core.telegram.org/bots/api#contact
type Contact struct {
	PhoneNumber string  `json:"phone_number"`
	FirstName   string  `json:"first_name"`
	LastName    *string `json:"last_name,omitempty"`
	UserId      int     `json:"user_id,omitempty"`
}

// Location is a struct for a location
//
// https://core.telegram.org/bots/api#location
type Location struct {
	Longitude float32 `json:"longitude"`
	Latitude  float32 `json:"latitude"`
}

// Venue is a struct of a venue
//
// https://core.telegram.org/bots/api#venue
type Venue struct {
	Location     Location `json:"location"`
	Title        string   `json:"title"`
	Address      string   `json:"address"`
	FoursquareId *string  `json:"foursquare_id,omitempty"`
}

// UserProfilePhotos is a struct for user profile photos
//
// https://core.telegram.org/bots/api#userprofilephotos
type UserProfilePhotos struct {
	TotalCount int           `json:"total_count"`
	Photos     [][]PhotoSize `json:"photos"`
}

// File is a struct for a file
//
// https://core.telegram.org/bots/api#file
type File struct {
	FileId   string  `json:"file_id"`
	FileSize int     `json:"file_size,omitempty"`
	FilePath *string `json:"file_path,omitempty"`
}

// ReplyKeyboardMarkup is a struct for reply keyboard markups
//
// https://core.telegram.org/bots/api#replykeyboardmarkup
type ReplyKeyboardMarkup struct {
	Keyboard        [][]KeyboardButton `json:"keyboard"`
	ResizeKeyboard  bool               `json:"resize_keyboard,omitempty"`
	OneTimeKeyboard bool               `json:"one_time_keyboard,omitempty"`
	Selective       bool               `json:"selective,omitempty"`
}

// KeyboardButton is a struct of a keyboard button
//
// https://core.telegram.org/bots/api#keyboardbutton
type KeyboardButton struct {
	Text            string `json:"text"`
	RequestContact  bool   `json:"request_contact,omitempty"`
	RequestLocation bool   `json:"request_location,omitempty"`
}

// ReplyKeyboardRemove is a struct for ReplyKeyboardRemove
//
// https://core.telegram.org/bots/api#replykeyboardremove
type ReplyKeyboardRemove struct {
	RemoveKeyboard bool `json:"remove_keyboard"`
	Selective      bool `json:"selective,omitempty"`
}

// InlineKeyboardMarkup is a struct for InlineKeyboardMarkup
//
// https://core.telegram.org/bots/api#inlinekeyboardmarkup
type InlineKeyboardMarkup struct {
	InlineKeyboard [][]InlineKeyboardButton `json:"inline_keyboard"`
}

// InlineKeyboardButton is a struct for InlineKeyboardButtons
//
// https://core.telegram.org/bots/api#inlinekeyboardbutton
type InlineKeyboardButton struct {
	Text                         string        `json:"text"`
	Url                          *string       `json:"url,omitempty"`
	CallbackData                 *string       `json:"callback_data,omitempty"`
	SwitchInlineQuery            *string       `json:"switch_inline_query,omitempty"`
	SwitchInlineQueryCurrentChat *string       `json:"switch_inline_query_current_chat,omitempty"`
	CallbackGame                 *CallbackGame `json:"callback_game,omitempty"`
	Pay                          bool          `json:"pay,omitempty"`
}

// CallbackQuery is a struct for a callback query
//
// https://core.telegram.org/bots/api#callbackquery
type CallbackQuery struct {
	Id              string   `json:"id"`
	From            User     `json:"from"`
	Message         *Message `json:"message,omitempty"`
	InlineMessageId *string  `json:"inline_message_id,omitempty"`
	ChatInstance    string   `json:"chat_instance"`
	Data            *string  `json:"data,omitempty"`
	GameShortName   *string  `json:"game_short_name,omitempty"`
}

// ShippingQuery is a struct for a shipping query
//
// https://core.telegram.org/bots/api#shippingquery
type ShippingQuery struct {
	Id              string          `json:"id"`
	From            User            `json:"from"`
	InvoicePayload  string          `json:"invoice_payload"`
	ShippingAddress ShippingAddress `json:"shipping_address"`
}

// PreCheckoutQuery is a struct for a precheckout query
//
// https://core.telegram.org/bots/api#precheckoutquery
type PreCheckoutQuery struct {
	Id                string     `json:"id"`
	From              User       `json:"from"`
	Currency          string     `json:"currency"`
	TotalAmount       int        `json:"total_amount"`
	InvoicePayload    string     `json:"invoice_payload"`
	ShippingOptioniId *string    `json:"shipping_option_id,omitempty"`
	OrderInfo         *OrderInfo `json:"order_info,omitempty"`
}

// ForceReply is a struct for force-reply
//
// https://core.telegram.org/bots/api#forcereply
type ForceReply struct {
	ForceReply bool `json:"force_reply"`
	Selective  bool `json:"selective,omitempty"`
}

// ChatPhoto is a struct for a chat photo
//
// https://core.telegram.org/bots/api#chatphoto
type ChatPhoto struct {
	SmallFileId string `json:"small_file_id"`
	BigFileId   string `json:"big_file_id"`
}

// ChatMember is a struct of a chat member
//
// https://core.telegram.org/bots/api#chatmember
type ChatMember struct {
	User                  User             `json:"user"`
	Status                ChatMemberStatus `json:"status"`
	UntilDate             int              `json:"until_date,omitempty"`
	CanBeEdited           bool             `json:"can_be_edited,omitempty"`
	CanChangeInfo         bool             `json:"can_change_info,omitempty"`
	CanPostMessages       bool             `json:"can_post_messages,omitempty"`
	CanEditMessages       bool             `json:"can_edit_messages,omitempty"`
	CanDeleteMessages     bool             `json:"can_delete_messages,omitempty"`
	CanInviteUsers        bool             `json:"can_invite_users,omitempty"`
	CanRestrictMembers    bool             `json:"can_restrict_members,omitempty"`
	CanPinMessages        bool             `json:"can_pin_messages,omitempty"`
	CanPromoteMembers     bool             `json:"can_promote_members,omitempty"`
	CanSendMessages       bool             `json:"can_send_messages,omitempty"`
	CanSendMediaMessages  bool             `json:"can_send_media_messages,omitempty"`
	CanSendOtherMessages  bool             `json:"can_send_other_messages,omitempty"`
	CanAddWebPagePreviews bool             `json:"can_add_web_page_previews,omitempty"`
}

// Message is a struct of a message
//
// https://core.telegram.org/bots/api#message
type Message struct {
	MessageId             int                `json:"message_id"`
	From                  *User              `json:"from,omitempty"`
	Date                  int                `json:"date"`
	Chat                  Chat               `json:"chat"`
	ForwardFrom           *User              `json:"forward_from,omitempty"`
	ForwardFromChat       *Chat              `json:"forward_from_chat,omitempty"`
	ForwardFromMessageId  int                `json:"forward_from_message_id,omitempty"`
	ForwardSignature      *string            `json:"forward_signature,omitempty"`
	ForwardDate           int                `json:"forward_date,omitempty"`
	ReplyToMessage        *Message           `json:"reply_to_message,omitempty"`
	EditDate              int                `json:"edit_date,omitempty"`
	AuthorSignature       *string            `json:"author_signature,omitempty"`
	Text                  *string            `json:"text,omitempty"`
	Entities              []MessageEntity    `json:"entities,omitempty"`
	CaptionEntities       []MessageEntity    `json:"caption_entities,omitempty"`
	Audio                 *Audio             `json:"audio,omitempty"`
	Document              *Document          `json:"document,omitempty"`
	Game                  *Game              `json:"game,omitempty"`
	Photo                 []PhotoSize        `json:"photo,omitempty"`
	Sticker               *Sticker           `json:"sticker,omitempty"`
	Video                 *Video             `json:"video,omitempty"`
	Voice                 *Voice             `json:"voice,omitempty"`
	VideoNote             *VideoNote         `json:"video_note,omitempty"`
	Caption               *string            `json:"caption,omitempty"`
	Contact               *Contact           `json:"contact,omitempty"`
	Location              *Location          `json:"location,omitempty"`
	Venue                 *Venue             `json:"venue,omitempty"`
	NewChatMembers        []User             `json:"new_chat_members,omitempty"`
	LeftChatMember        *User              `json:"left_chat_member,omitempty"`
	NewChatTitle          *string            `json:"new_chat_title,omitempty"`
	NewChatPhoto          []PhotoSize        `json:"new_chat_photo,omitempty"`
	DeleteChatPhoto       bool               `json:"delete_chat_photo,omitempty"`
	GroupChatCreated      bool               `json:"group_chat_created,omitempty"`
	SupergroupChatCreated bool               `json:"supergroup_chat_created,omitempty"`
	ChannelChatCreated    bool               `json:"channel_chat_created,omitempty"`
	MigrateToChatId       int64              `json:"migrate_to_chat_id,omitempty"`
	MigrateFromChatId     int64              `json:"migrate_from_chat_id,omitempty"`
	PinnedMessage         *Message           `json:"pinned_message,omitempty"`
	Invoice               *Invoice           `json:"invoice,omitempty"`
	SuccessfulPayment     *SuccessfulPayment `json:"successful_payment,omitempty"`
}

// InlineQuery is a struct of an inline query
//
// https://core.telegram.org/bots/api#inlinequery
type InlineQuery struct {
	Id       string    `json:"id"`
	From     User      `json:"from"`
	Location *Location `json:"location,omitempty"`
	Query    string    `json:"query"`
	Offset   string    `json:"offset"`
}

// ChosenInlineResult is a struct for a chosen inline result
//
// https://core.telegram.org/bots/api#choseninlineresult
type ChosenInlineResult struct {
	ResultId        string    `json:"result_id"`
	From            User      `json:"from"`
	Location        *Location `json:"location,omitempty"`
	InlineMessageId *string   `json:"inline_message_id,omitempty"`
	Query           string    `json:"query"`
}

// VideoMimeType is a video mime type for an inline query
type VideoMimeType string

// VideoMimeType strings
const (
	VideoMimeTypeHtml VideoMimeType = "text/html"
	VideoMimeTypeMp4  VideoMimeType = "video/mp4"
)

// DocumentMimeType is a document mime type for an inline query
type DocumentMimeType string

// DocumentMimeType strings
const (
	DocumentMimeTypePdf DocumentMimeType = "application/pdf"
	DocumentMimeTypeZip DocumentMimeType = "application/zip"
)

// InlineQueryResult is a struct for inline query results
//
// https://core.telegram.org/bots/api#inlinequeryresult
type InlineQueryResult struct {
	Type InlineQueryResultType `json:"type"`
	Id   string                `json:"id"`
}

// InlineQueryResultArticle is a struct for InlineQueryResultArticle
type InlineQueryResultArticle struct { // https://core.telegram.org/bots/api#inlinequeryresultarticle
	InlineQueryResult
	Title               string                `json:"title"`
	InputMessageContent InputMessageContent   `json:"input_message_content"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	Url                 *string               `json:"url,omitempty"`
	HideUrl             bool                  `json:"hide_url,omitempty"`
	Description         *string               `json:"description,omitempty"`
	ThumbUrl            *string               `json:"thumb_url,omitempty"`
	ThumbWidth          int                   `json:"thumb_width,omitempty"`
	ThumbHeight         int                   `json:"thumb_height,omitempty"`
}

// InlineQueryResultPhoto is a struct for InlineQueryResultPhoto
type InlineQueryResultPhoto struct { // https://core.telegram.org/bots/api#inlinequeryresultphoto
	InlineQueryResult
	PhotoUrl            string                `json:"photo_url"`
	PhotoWidth          int                   `json:"photo_width,omitempty"`
	PhotoHeight         int                   `json:"photo_height,omitempty"`
	ThumbUrl            string                `json:"thumb_url"`
	Title               *string               `json:"title,omitempty"`
	Description         *string               `json:"description,omitempty"`
	Caption             *string               `json:"caption,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
}

// InlineQueryResultGif is a struct for InlineQueryResultGif
type InlineQueryResultGif struct { // https://core.telegram.org/bots/api#inlinequeryresultgif
	InlineQueryResult
	GifUrl              string                `json:"gif_url"`
	GifWidth            int                   `json:"gif_width,omitempty"`
	GifHeight           int                   `json:"gif_height,omitempty"`
	GifDuration         int                   `json:"gif_duration,omitempty"`
	ThumbUrl            string                `json:"thumb_url"`
	Title               *string               `json:"title,omitempty"`
	Caption             *string               `json:"caption,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
}

// InlineQueryResultMpeg4Gif is a struct for InlineQueryResultMpeg4Gif
type InlineQueryResultMpeg4Gif struct { // https://core.telegram.org/bots/api#inlinequeryresultmpeg4gif
	InlineQueryResult
	Mpeg4Url            string                `json:"mpeg4_url"`
	Mpeg4Width          int                   `json:"mpeg4_width,omitempty"`
	Mpeg4Height         int                   `json:"mpeg4_height,omitempty"`
	Mpeg4Duration       int                   `json:"mpeg4_duration,omitempty"`
	ThumbUrl            string                `json:"thumb_url"`
	Title               *string               `json:"title,omitempty"`
	Caption             *string               `json:"caption,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
}

// InlineQueryResultVideo is a struct of InlineQueryResultVideo
type InlineQueryResultVideo struct { // https://core.telegram.org/bots/api#inlinequeryresultvideo
	InlineQueryResult
	VideoUrl            string                `json:"video_url"`
	MimeType            VideoMimeType         `json:"mime_type"`
	ThumbUrl            string                `json:"thumb_url"`
	Title               string                `json:"title"`
	Caption             *string               `json:"caption,omitempty"`
	VideoWidth          int                   `json:"video_width,omitempty"`
	VideoHeight         int                   `json:"video_height,omitempty"`
	VideoDuration       int                   `json:"video_duration,omitempty"`
	Description         *string               `json:"description,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
}

// InlineQueryResultAudio is a struct of InlineQueryResultAudio
type InlineQueryResultAudio struct { // https://core.telegram.org/bots/api#inlinequeryresultaudio
	InlineQueryResult
	AudioUrl            string                `json:"audio_url"`
	Title               string                `json:"title"`
	Caption             *string               `json:"caption,omitempty"`
	Performer           *string               `json:"performer,omitempty"`
	AudioDuration       int                   `json:"audio_duration,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent InputMessageContent   `json:"input_message_content,omitempty"`
}

// InlineQueryResultVoice is a struct of InlineQueryResultVoice
type InlineQueryResultVoice struct { // https://core.telegram.org/bots/api#inlinequeryresultvoice
	InlineQueryResult
	VoiceUrl            string                `json:"voice_url"`
	Title               string                `json:"title"`
	Caption             *string               `json:"caption,omitempty"`
	VoiceDuration       int                   `json:"voice_duration,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
}

// InlineQueryResultDocument is a struct of InlineQueryResultDocument
type InlineQueryResultDocument struct { // https://core.telegram.org/bots/api#inlinequeryresultdocument
	InlineQueryResult
	Title               string                `json:"title"`
	Caption             *string               `json:"caption,omitempty"`
	DocumentUrl         string                `json:"document_url"`
	MimeType            DocumentMimeType      `json:"mime_type"`
	Description         *string               `json:"description,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
	ThumbUrl            *string               `json:"thumb_url,omitempty"`
	ThumbWidth          int                   `json:"thumb_width,omitempty"`
	ThumbHeight         int                   `json:"thumb_height,omitempty"`
}

// InlineQueryResultLocation is a struct of InlineQueryResultLocation
type InlineQueryResultLocation struct { // https://core.telegram.org/bots/api#inlinequeryresultlocation
	InlineQueryResult
	Latitude            float32               `json:"latitude"`
	Longitude           float32               `json:"longitude"`
	Title               string                `json:"title"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
	ThumbUrl            *string               `json:"thumb_url,omitempty"`
	ThumbWidth          int                   `json:"thumb_width,omitempty"`
	ThumbHeight         int                   `json:"thumb_height,omitempty"`
}

// InlineQueryResultVenue is a struct of InlineQueryResultVenue
type InlineQueryResultVenue struct { // https://core.telegram.org/bots/api#inlinequeryresultvenue
	InlineQueryResult
	Latitude            float32               `json:"latitude"`
	Longitude           float32               `json:"longitude"`
	Title               string                `json:"title"`
	Address             string                `json:"address"`
	FoursquareId        *string               `json:"foursquare_id,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
	ThumbUrl            *string               `json:"thumb_url,omitempty"`
	ThumbWidth          int                   `json:"thumb_width,omitempty"`
	ThumbHeight         int                   `json:"thumb_height,omitempty"`
}

// InlineQueryResultContact is a struct of InlineQueryResultContact
type InlineQueryResultContact struct { // https://core.telegram.org/bots/api#inlinequeryresultcontact
	InlineQueryResult
	PhoneNumber         string                `json:"phone_number"`
	FirstName           string                `json:"first_name"`
	LastName            *string               `json:"last_name,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
	ThumbUrl            *string               `json:"thumb_url,omitempty"`
	ThumbWidth          int                   `json:"thumb_width,omitempty"`
	ThumbHeight         int                   `json:"thumb_height,omitempty"`
}

// InlineQueryResultGame is a struct of InlineQueryResultGame
type InlineQueryResultGame struct { // https://core.telegram.org/bots/api#inlinequeryresultgame
	InlineQueryResult
	GameShortName string                `json:"game_short_name"`
	ReplyMarkup   *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// InlineQueryResultCachedPhoto is a struct of InlineQueryResultCachedPhoto
type InlineQueryResultCachedPhoto struct { // https://core.telegram.org/bots/api#inlinequeryresultcachedphoto
	InlineQueryResult
	PhotoFileId         string                `json:"photo_file_id"`
	Title               *string               `json:"title,omitempty"`
	Description         *string               `json:"description,omitempty"`
	Caption             *string               `json:"caption,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
}

// InlineQueryResultCachedGif is a struct of InlineQueryResultCachedGif
type InlineQueryResultCachedGif struct { // https://core.telegram.org/bots/api#inlinequeryresultcachedgif
	InlineQueryResult
	GifFileId           string                `json:"gif_file_id"`
	Title               *string               `json:"title,omitempty"`
	Caption             *string               `json:"caption,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
}

// InlineQueryResultCachedMpeg4Gif is a struct of InlineQueryResultCachedMpeg4Gif
type InlineQueryResultCachedMpeg4Gif struct { // https://core.telegram.org/bots/api#inlinequeryresultcachedmpeg4gif
	InlineQueryResult
	Mpeg4FileId         string                `json:"mpeg4_file_id"`
	Title               *string               `json:"title,omitempty"`
	Caption             *string               `json:"caption,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
}

// InlineQueryResultCachedSticker is a struct of InlineQueryResultCachedSticker
type InlineQueryResultCachedSticker struct { // https://core.telegram.org/bots/api#inlinequeryresultcachedsticker
	InlineQueryResult
	StickerFileId       string                `json:"sticker_file_id"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
}

// InlineQueryResultCachedDocument is a struct of InlineQueryResultCachedDocument
type InlineQueryResultCachedDocument struct { // https://core.telegram.org/bots/api#inlinequeryresultcacheddocument
	InlineQueryResult
	Title               string                `json:"title"`
	DocumentFileId      string                `json:"document_file_id"`
	Description         *string               `json:"description,omitempty"`
	Caption             *string               `json:"caption,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
}

// InlineQueryResultCachedVideo is a struct of InlineQueryResultCachedVideo
type InlineQueryResultCachedVideo struct { // https://core.telegram.org/bots/api#inlinequeryresultcachedvideo
	InlineQueryResult
	VideoFileId         string                `json:"video_file_id"`
	Title               string                `json:"title"`
	Description         *string               `json:"description,omitempty"`
	Caption             *string               `json:"caption,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
}

// InlineQueryResultCachedVoice is a struct of InlineQueryResultCachedVoice
type InlineQueryResultCachedVoice struct { // https://core.telegram.org/bots/api#inlinequeryresultcachedvoice
	InlineQueryResult
	VoiceFileId         string                `json:"voice_file_id"`
	Title               string                `json:"title"`
	Caption             *string               `json:"caption,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
}

// InlineQueryResultCachedAudio is a struct of InlineQueryResultCachedAudio
type InlineQueryResultCachedAudio struct { // https://core.telegram.org/bots/api#inlinequeryresultcachedaudio
	InlineQueryResult
	AudioFileId         string                `json:"audio_file_id"`
	Caption             *string               `json:"caption,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
}

// InputMessageContent is a generic type of input message content types
//
// https://core.telegram.org/bots/api#inputmessagecontent
type InputMessageContent interface{}

// InputTextMessageContent is a struct of InputTextMessageContent
type InputTextMessageContent struct { // https://core.telegram.org/bots/api#inputtextmessagecontent
	MessageText           string     `json:"message_text"`
	ParseMode             *ParseMode `json:"parse_mode,omitempty"`
	DisableWebPagePreview bool       `json:"disable_web_page_preview,omitempty"`
}

// InputLocationMessageContent is a struct of InputLocationMessageContent
type InputLocationMessageContent struct { // https://core.telegram.org/bots/api#inputlocationmessagecontent
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
}

// InputVenueMessageContent is a struct of InputVenueMessageContent
type InputVenueMessageContent struct { // https://core.telegram.org/bots/api#inputvenuemessagecontent
	Latitude     float32 `json:"latitude"`
	Longitude    float32 `json:"longitude"`
	Title        string  `json:"title"`
	Address      string  `json:"address"`
	FoursquareId *string `json:"foursquare_id,omitempty"`
}

// InputContactMessageContent is a struct of InputContactMessageContent
type InputContactMessageContent struct { // https://core.telegram.org/bots/api#inputcontactmessagecontent
	PhoneNumber string  `json:"phone_number"`
	FirstName   string  `json:"first_name"`
	LastName    *string `json:"last_name,omitempty"`
}

// CallbackGame is for callback of games
//
// https://core.telegram.org/bots/api#callbackgame
type CallbackGame struct {
	// has nothing yet
}

// Game is a struct of Game
//
// https://core.telegram.org/bots/api#game
type Game struct {
	Title        string          `json:"title"`
	Description  string          `json:"description"`
	Photo        []PhotoSize     `json:"photo"`
	Text         *string         `json:"text,omitempty"`
	TextEntities []MessageEntity `json:"text_entities,omitempty"`
	Animation    *Animation      `json:"animation,omitempty"`
}

// Animation is a struct of Animation
//
// https://core.telegram.org/bots/api#animation
type Animation struct {
	FileId   string     `json:"file_id"`
	Thumb    *PhotoSize `json:"thumb,omitempty"`
	FileName *string    `json:"file_name,omitempty"`
	MimeType *string    `json:"mime_type,omitempty"`
	FileSize int        `json:"file_size,omitempty"`
}

// GameHighScore is a struct of GameHighScore
//
// https://core.telegram.org/bots/api#gamehighscore
type GameHighScore struct {
	Position int  `json:"position"`
	User     User `json:"user"`
	Score    int  `json:"score"`
}

// Invoice is a struct of Invoice
//
// https://core.telegram.org/bots/api#invoice
type Invoice struct {
	Title          string `json:"title"`
	Description    string `json:"description"`
	StartParameter string `json:"start_parameter"`
	Currency       string `json:"currency"`     // https://core.telegram.org/bots/payments#supported-currencies
	TotalAmount    int    `json:"total_amount"` // https://core.telegram.org/bots/payments/currencies.json
}

// SuccessfulPayment is a struct of successful payments
//
// https://core.telegram.org/bots/api#successfulpayment
type SuccessfulPayment struct {
	Currency                string     `json:"currency"`
	TotalAmount             int        `json:"total_amount"`
	InvoicePayload          string     `json:"invoice_payload"`
	ShippingOptionId        *string    `json:"shipping_option_id,omitempty"`
	OrderInfo               *OrderInfo `json:"order_info,omitempty"`
	TelegramPaymentChargeId string     `json:"telegram_payment_charge_id"`
	ProviderPaymentChargeId string     `json:"provider_payment_charge_id"`
}

// OrderInfo is a struct of order info
//
// https://core.telegram.org/bots/api#orderinfo
type OrderInfo struct {
	Name            *string          `json:"name,omitempty"`
	PhoneNumber     *string          `json:"phone_number,omitempty"`
	Email           *string          `json:"email,omitempty"`
	ShippingAddress *ShippingAddress `json:"shipping_address,omitempty"`
}

// ShippingOption is a struct of an option of the shipping
//
// https://core.telegram.org/bots/api#shippingoption
type ShippingOption struct {
	Id     string         `json:"id"`
	Title  string         `json:"title"`
	Prices []LabeledPrice `json:"prices"`
}

// LabeledPrice is a struct of labeled prices
//
// https://core.telegram.org/bots/api#labeledprice
type LabeledPrice struct {
	Label  string `json:"label"`
	Amount int    `json:"amount"`
}

// ShippingAddress is a struct of shipping address
//
// https://core.telegram.org/bots/api#shippingaddress
type ShippingAddress struct {
	CountryCode string `json:"country_code"`
	State       string `json:"state"`
	City        string `json:"city"`
	StreetLine1 string `json:"street_line1"`
	StreetLine2 string `json:"street_line2"`
	PostCode    string `json:"post_code"`
}
