package telegrambot

// https://core.telegram.org/bots/api#available-types

// ChatID can be `Message.Chat.Id`,
// or target channel name (in string, eg. "@channelusername")
type ChatID interface{}

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
	// (legacy) https://core.telegram.org/bots/api#markdown-style
	ParseModeMarkdown ParseMode = "Markdown"

	// https://core.telegram.org/bots/api#markdownv2-style
	ParseModeMarkdownV2 ParseMode = "MarkdownV2"

	// https://core.telegram.org/bots/api#html-style
	ParseModeHTML ParseMode = "HTML"
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
	MessageEntityTypeMention       MessageEntityType = "mention"
	MessageEntityTypeHashTag       MessageEntityType = "hashtag"
	MessageEntityTypeCashTag       MessageEntityType = "cashtag"
	MessageEntityTypeBotCommand    MessageEntityType = "bot_command"
	MessageEntityTypeURL           MessageEntityType = "url"
	MessageEntityTypeEmail         MessageEntityType = "email"
	MessageEntityTypePhoneNumber   MessageEntityType = "phone_number"
	MessageEntityTypeBold          MessageEntityType = "bold"
	MessageEntityTypeItalic        MessageEntityType = "italic"
	MessageEntityTypeUnderline     MessageEntityType = "underline"
	MessageEntityTypeStrikethrough MessageEntityType = "strikethrough"
	MessageEntityTypeCode          MessageEntityType = "code"
	MessageEntityTypePre           MessageEntityType = "pre"
	MessageEntityTypeTextLink      MessageEntityType = "text_link"
	MessageEntityTypeTextMention   MessageEntityType = "text_mention"
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
	ChatMemberStatusRestricted    ChatMemberStatus = "restricted"
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

// APIResponseBase is a base of API responses
type APIResponseBase struct {
	Ok          bool                   `json:"ok"`
	Description *string                `json:"description,omitempty"`
	Parameters  *APIResponseParameters `json:"parameters,omitempty"`
}

// APIResponseParameters is parameters in API responses
//
// https://core.telegram.org/bots/api#responseparameters
type APIResponseParameters struct {
	MigrateToChatID int64 `json:"migrate_to_chat_id,omitempty"`
	RetryAfter      int   `json:"retry_after,omitempty"`
}

// APIResponseWebhookInfo is an API response with result type: WebhookInfo
type APIResponseWebhookInfo struct {
	APIResponseBase
	Result *WebhookInfo `json:"result,omitempty"`
}

// APIResponseUser is an API response with result type: User
type APIResponseUser struct {
	APIResponseBase
	Result *User `json:"result,omitempty"`
}

// APIResponseMessage is an API response with result type: Message
type APIResponseMessage struct {
	APIResponseBase
	Result *Message `json:"result,omitempty"`
}

// APIResponseMessages is an API response with result type: []Message
type APIResponseMessages struct {
	APIResponseBase
	Result []Message `json:"result,omitempty"`
}

// APIResponseUserProfilePhotos is an API response with result type: UserProfilePhotos
type APIResponseUserProfilePhotos struct {
	APIResponseBase
	Result *UserProfilePhotos `json:"result,omitempty"`
}

// APIResponseFile is an API response with result type: File
type APIResponseFile struct {
	APIResponseBase
	Result *File `json:"result,omitempty"`
}

// APIResponseUpdates is an API response with result type: Update
type APIResponseUpdates struct {
	APIResponseBase
	Result []Update `json:"result,omitempty"`
}

// APIResponseChat is an API response with result type: Chat
type APIResponseChat struct {
	APIResponseBase
	Result *Chat `json:"result,omitempty"`
}

// APIResponseChatAdministrators is an API response with result type: ChatAdministrators
type APIResponseChatAdministrators struct {
	APIResponseBase
	Result []ChatMember `json:"result,omitempty"`
}

// APIResponseChatMember is an API response with result type: ChatMember
type APIResponseChatMember struct {
	APIResponseBase
	Result *ChatMember `json:"result,omitempty"`
}

// APIResponseInt is an API response with result type: int
type APIResponseInt struct {
	APIResponseBase
	Result int `json:"result,omitempty"`
}

// APIResponseBool is an API response with result type: bool
type APIResponseBool struct {
	APIResponseBase
	Result bool `json:"result,omitempty"`
}

// APIResponseString is an API response with result type: string
type APIResponseString struct {
	APIResponseBase
	Result *string `json:"result,omitempty"`
}

// APIResponseGameHighScores is an API response with result type: GameHighScores
type APIResponseGameHighScores struct {
	APIResponseBase
	Result []GameHighScore `json:"result,omitempty"`
}

// APIResponseStickerSet is an API response with result type: StickerSet
type APIResponseStickerSet struct {
	APIResponseBase
	Result *StickerSet `json:"result,omitempty"`
}

// APIResponseMessageOrBool is an API response with result type: Message or bool
type APIResponseMessageOrBool struct {
	APIResponseBase
	ResultMessage *Message
	ResultBool    *bool
}

// APIResponsePoll is an API response with result type: Poll
type APIResponsePoll struct {
	APIResponseBase
	Result *Poll `json:"result,omitempty"`
}

// APIResponseBotCommands is an API response with result type: []BotCommand
type APIResponseBotCommands struct {
	APIResponseBase
	Result []BotCommand `json:"result,omitempty"`
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
	UpdateTypeShippingQuery      UpdateType = "shipping_query"
	UpdateTypePreCheckoutQuery   UpdateType = "pre_checkout_query"
	UpdateTypePoll               UpdateType = "poll"
)

// WebhookInfo is a struct of webhook info
//
// https://core.telegram.org/bots/api#webhookinfo
type WebhookInfo struct {
	URL                  *string      `json:"url"`
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
	UpdateID           int                 `json:"update_id"`
	Message            *Message            `json:"message,omitempty"`
	EditedMessage      *Message            `json:"edited_message,omitempty"`
	ChannelPost        *Message            `json:"channel_post,omitempty"`
	EditedChannelPost  *Message            `json:"edited_channel_post,omitempty"`
	InlineQuery        *InlineQuery        `json:"inline_query,omitempty"`
	ChosenInlineResult *ChosenInlineResult `json:"chosen_inline_result,omitempty"`
	CallbackQuery      *CallbackQuery      `json:"callback_query,omitempty"`
	ShippingQuery      *ShippingQuery      `json:"shipping_query,omitempty"`
	PreCheckoutQuery   *PreCheckoutQuery   `json:"pre_checkout_query,omitempty"`
	Poll               *Poll               `json:"poll,omitempty"`
	PollAnswer         *PollAnswer         `json:"poll_answer,omitempty"`
}

// AllowedUpdate is a type for 'allowed_updates'
type AllowedUpdate string

// AllowedUpdate type constants constants constants constants
const (
	AllowMessage            AllowedUpdate = "message"
	AllowEditedMessage      AllowedUpdate = "edited_message"
	AllowChannelPost        AllowedUpdate = "channel_post"
	AllowEditedChannelPost  AllowedUpdate = "edited_channel_post"
	AllowInlineQuery        AllowedUpdate = "inline_query"
	AllowChosenInlineResult AllowedUpdate = "chosen_inline_result"
	AllowCallbackQuery      AllowedUpdate = "callback_query"
	AllowShippingQuery      AllowedUpdate = "shipping_query"
	AllowPreCheckoutQuery   AllowedUpdate = "pre_checkout_query"
)

// User is a struct of a user
//
// https://core.telegram.org/bots/api#user
type User struct {
	ID                      int     `json:"id"`
	IsBot                   bool    `json:"is_bot"`
	FirstName               string  `json:"first_name"`
	LastName                *string `json:"last_name,omitempty"`
	Username                *string `json:"username,omitempty"`
	LanguageCode            *string `json:"language_code,omitempty"`               // https://en.wikipedia.org/wiki/IETF_language_tag
	CanJoinGroups           bool    `json:"can_join_groups,omitempty"`             // returned only in GetMe()
	CanReadAllGroupMessages bool    `json:"can_read_all_group_messages,omitempty"` // returned only in GetMe()
	SupportsInlineQueries   bool    `json:"supports_inline_queries,omitempty"`     // returned only in GetMe()
}

// Chat is a struct of a chat
//
// https://core.telegram.org/bots/api#chat
type Chat struct {
	ID               int64            `json:"id"`
	Type             ChatType         `json:"type"`
	Title            *string          `json:"title,omitempty"`
	Username         *string          `json:"username,omitempty"`
	FirstName        *string          `json:"first_name,omitempty"`
	LastName         *string          `json:"last_name,omitempty"`
	Photo            *ChatPhoto       `json:"photo,omitempty"`
	Description      *string          `json:"description,omitempty"`
	InviteLink       *string          `json:"invite_link,omitempty"`
	PinnedMessage    *Message         `json:"pinned_message,omitempty"`
	Permissions      *ChatPermissions `json:"permissions,omitempty"`
	SlowModeDelay    int              `json:"slow_mode_delay,omitempty"`
	StickerSetName   *string          `json:"sticker_set_name,omitempty"`
	CanSetStickerSet bool             `json:"can_set_sticker_set,omitempty"`
}

// InputMediaType is a type of InputMedia
type InputMediaType string

// InputMediaType strings
const (
	InputMediaAnimation InputMediaType = "animation" // https://core.telegram.org/bots/api#inputmediaanimation
	InputMediaDocument  InputMediaType = "document"  // https://core.telegram.org/bots/api#inputmediadocument
	InputMediaAudio     InputMediaType = "audio"     // https://core.telegram.org/bots/api#inputmediaaudio
	InputMediaPhoto     InputMediaType = "photo"     // https://core.telegram.org/bots/api#inputmediaphoto
	InputMediaVideo     InputMediaType = "video"     // https://core.telegram.org/bots/api#inputmediavideo
)

// InputMedia represents the content of a media message to be sent.
//
// https://core.telegram.org/bots/api#inputmedia
type InputMedia struct {
	Type              InputMediaType `json:"type"`
	Media             string         `json:"media"`
	Thumb             *InputFile     `json:"thumb,omitempty"` // video, animation, audio, document
	Caption           *string        `json:"caption,omitempty"`
	ParseMode         *ParseMode     `json:"parse_mode,omitempty"`
	Width             int            `json:"width,omitempty"`              // video, animation
	Height            int            `json:"height,omitempty"`             // video, animation
	Duration          int            `json:"duration,omitempty"`           // video, animation
	Performer         *string        `json:"performer,omitempty"`          // audio only
	Title             *string        `json:"title,omitempty"`              // audio only
	SupportsStreaming bool           `json:"supports_streaming,omitempty"` // video only
}

// InputFile represents contents of a file to be uploaded.
// Can be generated with InputFileFromXXX() functions in types_helper.go
//
// https://core.telegram.org/bots/api#inputfile
type InputFile struct {
	Filepath *string
	URL      *string
	Bytes    []byte
	FileID   *string
}

// Audio is a struct for an audio file
//
// https://core.telegram.org/bots/api#audio
type Audio struct {
	FileID       string     `json:"file_id"`
	FileUniqueID string     `json:"file_unique_id"`
	Duration     int        `json:"duration"`
	Performer    *string    `json:"performer,omitempty"`
	Title        *string    `json:"title,omitempty"`
	MimeType     *string    `json:"mime_type,omitempty"`
	FileSize     int        `json:"file_size,omitempty"`
	Thumb        *PhotoSize `json:"thumb,omitempty"`
}

// MessageEntity is a struct of a message entity
//
// https://core.telegram.org/bots/api#messageentity
type MessageEntity struct {
	Type     MessageEntityType `json:"type"`
	Offset   int               `json:"offset"`
	Length   int               `json:"length"`
	URL      *string           `json:"url,omitempty"`      // when Type == MessageEntityTypeTextLink
	User     *User             `json:"user,omitempty"`     // when Type == MessageEntityTypeTextMention
	Language *string           `json:"language,omitempty"` // when Type == MessageEntityTypePre
}

// PhotoSize is a struct of a photo's size
//
// https://core.telegram.org/bots/api#photosize
type PhotoSize struct {
	FileID       string `json:"file_id"`
	FileUniqueID string `json:"file_unique_id"`
	Width        int    `json:"width"`
	Height       int    `json:"height"`
	FileSize     int    `json:"file_size,omitempty"`
}

// Document is a struct for an ordinary file
//
// https://core.telegram.org/bots/api#document
type Document struct {
	FileID       string     `json:"file_id"`
	FileUniqueID string     `json:"file_unique_id"`
	Thumb        *PhotoSize `json:"thumb,omitempty"`
	FileName     *string    `json:"file_name,omitempty"`
	MimeType     *string    `json:"mime_type,omitempty"`
	FileSize     int        `json:"file_size,omitempty"`
}

// Sticker is a struct of a sticker
//
// https://core.telegram.org/bots/api#sticker
type Sticker struct {
	FileID       string        `json:"file_id"`
	FileUniqueID string        `json:"file_unique_id"`
	Width        int           `json:"width"`
	Height       int           `json:"height"`
	IsAnimated   bool          `json:"is_animated"`
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
	Name          string     `json:"name"`
	Title         string     `json:"title"`
	IsAnimated    bool       `json:"is_animated"`
	ContainsMasks bool       `json:"contains_masks"`
	Stickers      []Sticker  `json:"stickers"`
	Thumb         *PhotoSize `json:"thumb,omitempty"`
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
	FileID       string     `json:"file_id"`
	FileUniqueID string     `json:"file_unique_id"`
	Width        int        `json:"width"`
	Height       int        `json:"height"`
	Duration     int        `json:"duration"`
	Thumb        *PhotoSize `json:"thumb,omitempty"`
	MimeType     *string    `json:"mime_type,omitempty"`
	FileSize     int        `json:"file_size,omitempty"`
}

// Voice is a struct for a voice file
//
// https://core.telegram.org/bots/api#voice
type Voice struct {
	FileID       string  `json:"file_id"`
	FileUniqueID string  `json:"file_unique_id"`
	Duration     int     `json:"duration"`
	MimeType     *string `json:"mime_type,omitempty"`
	FileSize     int     `json:"file_size,omitempty"`
}

// VideoNote is a struct for a video note
//
// https://core.telegram.org/bots/api#videonote
type VideoNote struct {
	FileID       string     `json:"file_id"`
	FileUniqueID string     `json:"file_unique_id"`
	Length       int        `json:"length"`
	Duration     int        `json:"duration"`
	Thumb        *PhotoSize `json:"thumb,omitempty"`
	FileSize     int        `json:"file_size,omitempty"`
}

// Contact is a struct for a contact info
//
// https://core.telegram.org/bots/api#contact
type Contact struct {
	PhoneNumber string  `json:"phone_number"`
	FirstName   string  `json:"first_name"`
	LastName    *string `json:"last_name,omitempty"`
	UserID      int     `json:"user_id,omitempty"`
	VCard       *string `json:"vcard,omitempty"` // https://en.wikipedia.org/wiki/VCard
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
	Location       Location `json:"location"`
	Title          string   `json:"title"`
	Address        string   `json:"address"`
	FoursquareID   *string  `json:"foursquare_id,omitempty"`
	FoursquareType *string  `json:"foursquare_type,omitempty"`
}

// Poll is a struct of a poll
//
// https://core.telegram.org/bots/api#poll
type Poll struct {
	ID                    string       `json:"id"`
	Question              string       `json:"question"` // 1~255 chars
	Options               []PollOption `json:"options"`
	TotalVoterCount       int          `json:"total_voter_count"`
	IsClosed              bool         `json:"is_closed"`
	IsAnonymous           bool         `json:"is_anonymous"`
	Type                  string       `json:"type"` // "quiz" or "regular"
	AllowsMultipleAnswers bool         `json:"allows_multiple_answers"`
}

// PollOption is a struct of a poll option
//
// https://core.telegram.org/bots/api#polloption
type PollOption struct {
	Text       string `json:"text"` // 1~100 chars
	VoterCount int    `json:"voter_count"`
}

// PollAnswer is a struct of a poll answer
//
// https://core.telegram.org/bots/api#pollanswer
type PollAnswer struct {
	PollID    string `json:"poll_id"`
	User      User   `json:"user"`
	OptionIDs []int  `json:"option_ids"`
}

// Dice is a struct for dice in message
//
// https://core.telegram.org/bots/api#senddice
type Dice struct {
	Value int `json:"value"` // 1-6
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
	FileID       string  `json:"file_id"`
	FileUniqueID string  `json:"file_unique_id"`
	FileSize     int     `json:"file_size,omitempty"`
	FilePath     *string `json:"file_path,omitempty"`
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
	Text            string                  `json:"text"`
	RequestContact  bool                    `json:"request_contact,omitempty"`
	RequestLocation bool                    `json:"request_location,omitempty"`
	RequestPoll     *KeyboardButtonPollType `json:"request_poll,omitempty"`
}

// KeyboardButtonPollType is a struct for KeyboardButtonPollType
//
// https://core.telegram.org/bots/api#keyboardbuttonpolltype
type KeyboardButtonPollType struct {
	Type *string `json:"type,omitempty"` // "quiz", "regular", or anything
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
	URL                          *string       `json:"url,omitempty"`
	LoginURL                     *LoginURL     `json:"login_url,omitempty"`
	CallbackData                 *string       `json:"callback_data,omitempty"`
	SwitchInlineQuery            *string       `json:"switch_inline_query,omitempty"`
	SwitchInlineQueryCurrentChat *string       `json:"switch_inline_query_current_chat,omitempty"`
	CallbackGame                 *CallbackGame `json:"callback_game,omitempty"`
	Pay                          bool          `json:"pay,omitempty"`
}

// LoginURL is a struct for LoginURL
//
// https://core.telegram.org/bots/api#loginurl
type LoginURL struct {
	URL                string  `json:"url"`
	ForwardText        *string `json:"forward_text,omitempty"`
	BotUsername        *string `json:"bot_username,omitempty"`
	RequestWriteAccess bool    `json:"request_write_access,omitempty"`
}

// CallbackQuery is a struct for a callback query
//
// https://core.telegram.org/bots/api#callbackquery
type CallbackQuery struct {
	ID              string   `json:"id"`
	From            User     `json:"from"`
	Message         *Message `json:"message,omitempty"`
	InlineMessageID *string  `json:"inline_message_id,omitempty"`
	ChatInstance    string   `json:"chat_instance"`
	Data            *string  `json:"data,omitempty"`
	GameShortName   *string  `json:"game_short_name,omitempty"`
}

// ShippingQuery is a struct for a shipping query
//
// https://core.telegram.org/bots/api#shippingquery
type ShippingQuery struct {
	ID              string          `json:"id"`
	From            User            `json:"from"`
	InvoicePayload  string          `json:"invoice_payload"`
	ShippingAddress ShippingAddress `json:"shipping_address"`
}

// PreCheckoutQuery is a struct for a precheckout query
//
// https://core.telegram.org/bots/api#precheckoutquery
type PreCheckoutQuery struct {
	ID               string     `json:"id"`
	From             User       `json:"from"`
	Currency         string     `json:"currency"`
	TotalAmount      int        `json:"total_amount"`
	InvoicePayload   string     `json:"invoice_payload"`
	ShippingOptionID *string    `json:"shipping_option_id,omitempty"`
	OrderInfo        *OrderInfo `json:"order_info,omitempty"`
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
	SmallFileID       string `json:"small_file_id"`
	SmallFileUniqueID string `json:"small_file_unique_id"`
	BigFileID         string `json:"big_file_id"`
	BigFileUniqueID   string `json:"big_file_unique_id"`
}

// ChatMember is a struct of a chat member
//
// https://core.telegram.org/bots/api#chatmember
type ChatMember struct {
	User                  User             `json:"user"`
	Status                ChatMemberStatus `json:"status"`
	CustomTitle           *string          `json:"custom_title,omitempty"`              // owner and administrators only
	UntilDate             int              `json:"until_date,omitempty"`                // restricted and kicked only
	CanBeEdited           bool             `json:"can_be_edited,omitempty"`             // administrators only
	CanPostMessages       bool             `json:"can_post_messages,omitempty"`         // administrators only
	CanEditMessages       bool             `json:"can_edit_messages,omitempty"`         // administrators only
	CanDeleteMessages     bool             `json:"can_delete_messages,omitempty"`       // administrators only
	CanRestrictMembers    bool             `json:"can_restrict_members,omitempty"`      // administrators only
	CanPromoteMembers     bool             `json:"can_promote_members,omitempty"`       // administrators only
	CanChangeInfo         bool             `json:"can_change_info,omitempty"`           // administrators and restricted only
	CanInviteUsers        bool             `json:"can_invite_users,omitempty"`          // administrators and restricted only
	CanPinMessages        bool             `json:"can_pin_messages,omitempty"`          // administrators and restricted only
	IsMember              bool             `json:"is_member,omitempty"`                 // restricted only
	CanSendMessages       bool             `json:"can_send_messages,omitempty"`         // restricted only
	CanSendMediaMessages  bool             `json:"can_send_media_messages,omitempty"`   // restricted only
	CanSendPolls          bool             `json:"can_send_polls,omitempty"`            // restricted only
	CanSendOtherMessages  bool             `json:"can_send_other_messages,omitempty"`   // restricted only
	CanAddWebPagePreviews bool             `json:"can_add_web_page_previews,omitempty"` // restricted only
}

// ChatPermissions is a struct of chat permissions
//
// https://core.telegram.org/bots/api#chatpermissions
type ChatPermissions struct {
	CanSendMessages       bool `json:"can_send_messages,omitempty"`
	CanSendMediaMessages  bool `json:"can_send_media_messages,omitempty"`
	CanSendPolls          bool `json:"can_send_polls,omitempty"`
	CanSendOtherMessages  bool `json:"can_send_other_messages,omitempty"`
	CanAddWebPagePreviews bool `json:"can_add_web_page_previews,omitempty"`
	CanChangeInfo         bool `json:"can_change_info,omitempty"`
	CanInviteUsers        bool `json:"can_invite_users,omitempty"`
	CanPinMessages        bool `json:"can_pin_messages,omitempty"`
}

// BotCommand is a struct of a bot command
//
// https://core.telegram.org/bots/api#botcommand
type BotCommand struct {
	Command     string `json:"command"`
	Description string `json:"description"`
}

// Message is a struct of a message
//
// https://core.telegram.org/bots/api#message
type Message struct {
	MessageID             int                `json:"message_id"`
	From                  *User              `json:"from,omitempty"`
	Date                  int                `json:"date"`
	Chat                  Chat               `json:"chat"`
	ForwardFrom           *User              `json:"forward_from,omitempty"`
	ForwardFromChat       *Chat              `json:"forward_from_chat,omitempty"`
	ForwardFromMessageID  int                `json:"forward_from_message_id,omitempty"`
	ForwardSignature      *string            `json:"forward_signature,omitempty"`
	ForwardSenderName     *string            `json:"forward_sender_name,omitempty"`
	ForwardDate           int                `json:"forward_date,omitempty"`
	ReplyToMessage        *Message           `json:"reply_to_message,omitempty"`
	EditDate              int                `json:"edit_date,omitempty"`
	AuthorSignature       *string            `json:"author_signature,omitempty"`
	Text                  *string            `json:"text,omitempty"`
	Entities              []MessageEntity    `json:"entities,omitempty"`
	CaptionEntities       []MessageEntity    `json:"caption_entities,omitempty"`
	Audio                 *Audio             `json:"audio,omitempty"`
	Document              *Document          `json:"document,omitempty"`
	Animation             *Animation         `json:"animation,omitempty"`
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
	Poll                  *Poll              `json:"poll,omitempty"`
	Dice                  *Dice              `json:"dice,omitempty"`
	NewChatMembers        []User             `json:"new_chat_members,omitempty"`
	LeftChatMember        *User              `json:"left_chat_member,omitempty"`
	NewChatTitle          *string            `json:"new_chat_title,omitempty"`
	NewChatPhoto          []PhotoSize        `json:"new_chat_photo,omitempty"`
	DeleteChatPhoto       bool               `json:"delete_chat_photo,omitempty"`
	GroupChatCreated      bool               `json:"group_chat_created,omitempty"`
	SupergroupChatCreated bool               `json:"supergroup_chat_created,omitempty"`
	ChannelChatCreated    bool               `json:"channel_chat_created,omitempty"`
	MigrateToChatID       int64              `json:"migrate_to_chat_id,omitempty"`
	MigrateFromChatID     int64              `json:"migrate_from_chat_id,omitempty"`
	PinnedMessage         *Message           `json:"pinned_message,omitempty"`
	Invoice               *Invoice           `json:"invoice,omitempty"`
	SuccessfulPayment     *SuccessfulPayment `json:"successful_payment,omitempty"`
	ConnectedWebsite      *string            `json:"connected_website,omitempty"`
	//PassportData          *PassportData         `json:"passport_data,omitempty"` // NOT IMPLEMENTED: https://core.telegram.org/bots/api#passportdata
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// InlineQuery is a struct of an inline query
//
// https://core.telegram.org/bots/api#inlinequery
type InlineQuery struct {
	ID       string    `json:"id"`
	From     User      `json:"from"`
	Location *Location `json:"location,omitempty"`
	Query    string    `json:"query"`
	Offset   string    `json:"offset"`
}

// ChosenInlineResult is a struct for a chosen inline result
//
// https://core.telegram.org/bots/api#choseninlineresult
type ChosenInlineResult struct {
	ResultID        string    `json:"result_id"`
	From            User      `json:"from"`
	Location        *Location `json:"location,omitempty"`
	InlineMessageID *string   `json:"inline_message_id,omitempty"`
	Query           string    `json:"query"`
}

// VideoMimeType is a video mime type for an inline query
type VideoMimeType string

// VideoMimeType strings
const (
	VideoMimeTypeHTML VideoMimeType = "text/html"
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
	ID   string                `json:"id"`
}

// InlineQueryResultArticle is a struct for InlineQueryResultArticle
type InlineQueryResultArticle struct { // https://core.telegram.org/bots/api#inlinequeryresultarticle
	InlineQueryResult
	Title               string                `json:"title"`
	InputMessageContent InputMessageContent   `json:"input_message_content"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	URL                 *string               `json:"url,omitempty"`
	HideURL             bool                  `json:"hide_url,omitempty"`
	Description         *string               `json:"description,omitempty"`
	ThumbURL            *string               `json:"thumb_url,omitempty"`
	ThumbWidth          int                   `json:"thumb_width,omitempty"`
	ThumbHeight         int                   `json:"thumb_height,omitempty"`
}

// InlineQueryResultPhoto is a struct for InlineQueryResultPhoto
type InlineQueryResultPhoto struct { // https://core.telegram.org/bots/api#inlinequeryresultphoto
	InlineQueryResult
	PhotoURL            string                `json:"photo_url"`
	PhotoWidth          int                   `json:"photo_width,omitempty"`
	PhotoHeight         int                   `json:"photo_height,omitempty"`
	ThumbURL            string                `json:"thumb_url"`
	Title               *string               `json:"title,omitempty"`
	Description         *string               `json:"description,omitempty"`
	Caption             *string               `json:"caption,omitempty"`
	ParseMode           *ParseMode            `json:"parse_mode,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
}

// InlineQueryResultGif is a struct for InlineQueryResultGif
type InlineQueryResultGif struct { // https://core.telegram.org/bots/api#inlinequeryresultgif
	InlineQueryResult
	GifURL              string                `json:"gif_url"`
	GifWidth            int                   `json:"gif_width,omitempty"`
	GifHeight           int                   `json:"gif_height,omitempty"`
	GifDuration         int                   `json:"gif_duration,omitempty"`
	ThumbURL            string                `json:"thumb_url"`
	Title               *string               `json:"title,omitempty"`
	Caption             *string               `json:"caption,omitempty"`
	ParseMode           *ParseMode            `json:"parse_mode,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
}

// InlineQueryResultMpeg4Gif is a struct for InlineQueryResultMpeg4Gif
type InlineQueryResultMpeg4Gif struct { // https://core.telegram.org/bots/api#inlinequeryresultmpeg4gif
	InlineQueryResult
	Mpeg4URL            string                `json:"mpeg4_url"`
	Mpeg4Width          int                   `json:"mpeg4_width,omitempty"`
	Mpeg4Height         int                   `json:"mpeg4_height,omitempty"`
	Mpeg4Duration       int                   `json:"mpeg4_duration,omitempty"`
	ThumbURL            string                `json:"thumb_url"`
	Title               *string               `json:"title,omitempty"`
	Caption             *string               `json:"caption,omitempty"`
	ParseMode           *ParseMode            `json:"parse_mode,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
}

// InlineQueryResultVideo is a struct of InlineQueryResultVideo
type InlineQueryResultVideo struct { // https://core.telegram.org/bots/api#inlinequeryresultvideo
	InlineQueryResult
	VideoURL            string                `json:"video_url"`
	MimeType            VideoMimeType         `json:"mime_type"`
	ThumbURL            string                `json:"thumb_url"`
	Title               string                `json:"title"`
	Caption             *string               `json:"caption,omitempty"`
	ParseMode           *ParseMode            `json:"parse_mode,omitempty"`
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
	AudioURL            string                `json:"audio_url"`
	Title               string                `json:"title"`
	Caption             *string               `json:"caption,omitempty"`
	ParseMode           *ParseMode            `json:"parse_mode,omitempty"`
	Performer           *string               `json:"performer,omitempty"`
	AudioDuration       int                   `json:"audio_duration,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
}

// InlineQueryResultVoice is a struct of InlineQueryResultVoice
type InlineQueryResultVoice struct { // https://core.telegram.org/bots/api#inlinequeryresultvoice
	InlineQueryResult
	VoiceURL            string                `json:"voice_url"`
	Title               string                `json:"title"`
	Caption             *string               `json:"caption,omitempty"`
	ParseMode           *ParseMode            `json:"parse_mode,omitempty"`
	VoiceDuration       int                   `json:"voice_duration,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
}

// InlineQueryResultDocument is a struct of InlineQueryResultDocument
type InlineQueryResultDocument struct { // https://core.telegram.org/bots/api#inlinequeryresultdocument
	InlineQueryResult
	Title               string                `json:"title"`
	Caption             *string               `json:"caption,omitempty"`
	ParseMode           *ParseMode            `json:"parse_mode,omitempty"`
	DocumentURL         string                `json:"document_url"`
	MimeType            DocumentMimeType      `json:"mime_type"`
	Description         *string               `json:"description,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
	ThumbURL            *string               `json:"thumb_url,omitempty"`
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
	ThumbURL            *string               `json:"thumb_url,omitempty"`
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
	FoursquareID        *string               `json:"foursquare_id,omitempty"`
	FoursquareType      *string               `json:"foursquare_type,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
	ThumbURL            *string               `json:"thumb_url,omitempty"`
	ThumbWidth          int                   `json:"thumb_width,omitempty"`
	ThumbHeight         int                   `json:"thumb_height,omitempty"`
}

// InlineQueryResultContact is a struct of InlineQueryResultContact
type InlineQueryResultContact struct { // https://core.telegram.org/bots/api#inlinequeryresultcontact
	InlineQueryResult
	PhoneNumber         string                `json:"phone_number"`
	FirstName           string                `json:"first_name"`
	LastName            *string               `json:"last_name,omitempty"`
	VCard               *string               `json:"vcard,omitempty"` // https://en.wikipedia.org/wiki/VCard
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
	ThumbURL            *string               `json:"thumb_url,omitempty"`
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
	PhotoFileID         string                `json:"photo_file_id"`
	Title               *string               `json:"title,omitempty"`
	Description         *string               `json:"description,omitempty"`
	Caption             *string               `json:"caption,omitempty"`
	ParseMode           *ParseMode            `json:"parse_mode,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
}

// InlineQueryResultCachedGif is a struct of InlineQueryResultCachedGif
type InlineQueryResultCachedGif struct { // https://core.telegram.org/bots/api#inlinequeryresultcachedgif
	InlineQueryResult
	GifFileID           string                `json:"gif_file_id"`
	Title               *string               `json:"title,omitempty"`
	Caption             *string               `json:"caption,omitempty"`
	ParseMode           *ParseMode            `json:"parse_mode,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
}

// InlineQueryResultCachedMpeg4Gif is a struct of InlineQueryResultCachedMpeg4Gif
type InlineQueryResultCachedMpeg4Gif struct { // https://core.telegram.org/bots/api#inlinequeryresultcachedmpeg4gif
	InlineQueryResult
	Mpeg4FileID         string                `json:"mpeg4_file_id"`
	Title               *string               `json:"title,omitempty"`
	Caption             *string               `json:"caption,omitempty"`
	ParseMode           *ParseMode            `json:"parse_mode,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
}

// InlineQueryResultCachedSticker is a struct of InlineQueryResultCachedSticker
type InlineQueryResultCachedSticker struct { // https://core.telegram.org/bots/api#inlinequeryresultcachedsticker
	InlineQueryResult
	StickerFileID       string                `json:"sticker_file_id"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
}

// InlineQueryResultCachedDocument is a struct of InlineQueryResultCachedDocument
type InlineQueryResultCachedDocument struct { // https://core.telegram.org/bots/api#inlinequeryresultcacheddocument
	InlineQueryResult
	Title               string                `json:"title"`
	DocumentFileID      string                `json:"document_file_id"`
	Description         *string               `json:"description,omitempty"`
	Caption             *string               `json:"caption,omitempty"`
	ParseMode           *ParseMode            `json:"parse_mode,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
}

// InlineQueryResultCachedVideo is a struct of InlineQueryResultCachedVideo
type InlineQueryResultCachedVideo struct { // https://core.telegram.org/bots/api#inlinequeryresultcachedvideo
	InlineQueryResult
	VideoFileID         string                `json:"video_file_id"`
	Title               string                `json:"title"`
	Description         *string               `json:"description,omitempty"`
	Caption             *string               `json:"caption,omitempty"`
	ParseMode           *ParseMode            `json:"parse_mode,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
}

// InlineQueryResultCachedVoice is a struct of InlineQueryResultCachedVoice
type InlineQueryResultCachedVoice struct { // https://core.telegram.org/bots/api#inlinequeryresultcachedvoice
	InlineQueryResult
	VoiceFileID         string                `json:"voice_file_id"`
	Title               string                `json:"title"`
	Caption             *string               `json:"caption,omitempty"`
	ParseMode           *ParseMode            `json:"parse_mode,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
}

// InlineQueryResultCachedAudio is a struct of InlineQueryResultCachedAudio
type InlineQueryResultCachedAudio struct { // https://core.telegram.org/bots/api#inlinequeryresultcachedaudio
	InlineQueryResult
	AudioFileID         string                `json:"audio_file_id"`
	Caption             *string               `json:"caption,omitempty"`
	ParseMode           *ParseMode            `json:"parse_mode,omitempty"`
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
	Latitude       float32 `json:"latitude"`
	Longitude      float32 `json:"longitude"`
	Title          string  `json:"title"`
	Address        string  `json:"address"`
	FoursquareID   *string `json:"foursquare_id,omitempty"`
	FoursquareType *string `json:"foursquare_type,omitempty"`
}

// InputContactMessageContent is a struct of InputContactMessageContent
type InputContactMessageContent struct { // https://core.telegram.org/bots/api#inputcontactmessagecontent
	PhoneNumber string  `json:"phone_number"`
	FirstName   string  `json:"first_name"`
	LastName    *string `json:"last_name,omitempty"`
	VCard       *string `json:"vcard,omitempty"` // https://en.wikipedia.org/wiki/VCard
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
	FileID       string     `json:"file_id"`
	FileUniqueID string     `json:"file_unique_id"`
	Width        int        `json:"width"`
	Height       int        `json:"height"`
	Duration     int        `json:"duration"`
	Thumb        *PhotoSize `json:"thumb,omitempty"`
	FileName     *string    `json:"file_name,omitempty"`
	MimeType     *string    `json:"mime_type,omitempty"`
	FileSize     int        `json:"file_size,omitempty"`
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
	ShippingOptionID        *string    `json:"shipping_option_id,omitempty"`
	OrderInfo               *OrderInfo `json:"order_info,omitempty"`
	TelegramPaymentChargeID string     `json:"telegram_payment_charge_id"`
	ProviderPaymentChargeID string     `json:"provider_payment_charge_id"`
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
	ID     string         `json:"id"`
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
