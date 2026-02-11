package telegrambot

import "encoding/json"

// https://core.telegram.org/bots/api#available-types

// ChatID can be `Message.Chat.Id`,
// or target channel name (in string, eg. "@channelusername")
type ChatID any

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
	ChatActionRecordVoice     ChatAction = "record_voice"
	ChatActionUploadVoice     ChatAction = "upload_voice"
	ChatActionUploadDocument  ChatAction = "upload_document"
	ChatActionChooseSticker   ChatAction = "choose_sticker"
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

// ThumbnailMimeType is a type of inline query result's thumbnail mime type
type ThumbnailMimeType string

// ThumbnailMimeType strings
const (
	ThumbnailMimeTypeImageJpeg ThumbnailMimeType = "image/jpeg"
	ThumbnailMimeTypeImageGif  ThumbnailMimeType = "image/gif"
	ThumbnailMimeTypeVideoMp4  ThumbnailMimeType = "video/mp4"
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
	MessageEntityTypeSpoiler       MessageEntityType = "spoiler"
	MessageEntityTypeBlockquote    MessageEntityType = "blockquote"
	MessageEntityTypeCode          MessageEntityType = "code"
	MessageEntityTypePre           MessageEntityType = "pre"
	MessageEntityTypeTextLink      MessageEntityType = "text_link"
	MessageEntityTypeTextMention   MessageEntityType = "text_mention"
	MessageEntityTypeCustomEmoji   MessageEntityType = "custom_emoji"
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
	ChatMemberStatusBanned        ChatMemberStatus = "kicked"
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

// base error type
type baseError struct {
	Message string
}

// Error returns error message.
func (e baseError) Error() string {
	return e.Message
}

// custom error types
//
// referenced: https://github.com/TelegramBotAPI/errors
type (
	ErrUnauthorized              struct{ baseError } // for error: 'Unauthorized'
	ErrChatNotFound              struct{ baseError } // for error: 'Bad Request: chat not found'
	ErrUserNotFound              struct{ baseError } // for error: 'Bad Request: user not found'
	ErrUserDeactivated           struct{ baseError } // for error: 'Forbidden: user is deactivated'
	ErrBotKicked                 struct{ baseError } // for error: 'Forbidden: bot was kicked'
	ErrBotBlockedByUser          struct{ baseError } // for error: 'Forbidden: bot blocked by user'
	ErrBotCantSendToBots         struct{ baseError } // for error: 'Forbidden: bot can't send messages to bots'
	ErrMessageNotModified        struct{ baseError } // for error: 'Bad request: Message not modified'
	ErrGroupMigratedToSupergroup struct{ baseError } // for error: 'Bad request: Group migrated to supergroup'
	ErrInvalidFileID             struct{ baseError } // for error: 'Bad request: Invalid file id'
	ErrConflictedLongPoll        struct{ baseError } // for error: 'Conflict: Terminated by other long poll'
	ErrConflictedWebHook         struct{ baseError } // for error: 'Conflict: can't use getUpdates method while webhook is active; use deleteWebhook to delete the webhook first'
	ErrWrongParameterAction      struct{ baseError } // for error: 'Bad request: Wrong parameter action in request'
	ErrMessageEmpty              struct{ baseError } // for error: 'Bad Request: message text is empty'
	ErrMessageTooLong            struct{ baseError } // for error: 'Bad Request: message is too long'
	ErrMessageCantBeEdited       struct{ baseError } // for error: 'Bad Request: message can't be edited'
	ErrTooManyRequests           struct{ baseError } // for error: 'Too many requests'
	ErrJSONParseFailed           struct{ baseError } // for error: 'failed to parse json'
	ErrContextTimeout            struct{ baseError } // for error: 'context deadline exceeded'
	ErrUnclassified              struct{ baseError } // for unclassified errors
)

// APIResponse is a base of API responses
type APIResponse[T any] struct {
	OK          bool    `json:"ok"`
	Description *string `json:"description,omitempty"`

	Parameters *APIResponseParameters `json:"parameters,omitempty"`

	Result *T `json:"result,omitempty"`
}

// APIResponseMessageOrBool type for ambiguous type of `result`
type APIResponseMessageOrBool struct {
	OK          bool    `json:"ok"`
	Description *string `json:"description,omitempty"`

	Parameters *APIResponseParameters `json:"parameters,omitempty"`

	ResultMessage *Message `json:"result_message,omitempty"`
	ResultBool    *bool    `json:"result_bool,omitempty"`
}

// APIResponseParameters is parameters in API responses
//
// https://core.telegram.org/bots/api#responseparameters
type APIResponseParameters struct {
	MigrateToChatID *int64 `json:"migrate_to_chat_id,omitempty"`
	RetryAfter      *int   `json:"retry_after,omitempty"`
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
	URL                          *string      `json:"url"`
	HasCustomCertificate         bool         `json:"has_custom_certificate"`
	PendingUpdateCount           int          `json:"pending_update_count"`
	IPAddress                    *string      `json:"ip_address,omitempty"`
	LastErrorDate                *int         `json:"last_error_date,omitempty"`
	LastErrorMessage             *string      `json:"last_error_message,omitempty"`
	LastSynchronizationErrorDate *int         `json:"last_synchronization_error_date,omitempty"`
	MaxConnections               *int         `json:"max_connections,omitempty"`
	AllowedUpdates               []UpdateType `json:"allowed_updates,omitempty"`
}

// Update is a struct of an update
//
// https://core.telegram.org/bots/api#update
type Update struct {
	UpdateID                int64                        `json:"update_id"`
	Message                 *Message                     `json:"message,omitempty"`
	EditedMessage           *Message                     `json:"edited_message,omitempty"`
	ChannelPost             *Message                     `json:"channel_post,omitempty"`
	EditedChannelPost       *Message                     `json:"edited_channel_post,omitempty"`
	BusinessConnection      *BusinessConnection          `json:"business_connection,omitempty"`
	BusinessMessage         *Message                     `json:"business_message,omitempty"`
	EditedBusinessMessage   *Message                     `json:"edited_business_message,omitempty"`
	DeletedBusinessMessages *BusinessMessagesDeleted     `json:"deleted_business_messages,omitempty"`
	MessageReaction         *MessageReactionUpdated      `json:"message_reaction,omitempty"`
	MessageReactionCount    *MessageReactionCountUpdated `json:"message_reaction_count,omitempty"`
	InlineQuery             *InlineQuery                 `json:"inline_query,omitempty"`
	ChosenInlineResult      *ChosenInlineResult          `json:"chosen_inline_result,omitempty"`
	CallbackQuery           *CallbackQuery               `json:"callback_query,omitempty"`
	ShippingQuery           *ShippingQuery               `json:"shipping_query,omitempty"`
	PreCheckoutQuery        *PreCheckoutQuery            `json:"pre_checkout_query,omitempty"`
	PurchasedPaidMedia      *PaidMediaPurchased          `json:"purchased_paid_media,omitempty"`
	Poll                    *Poll                        `json:"poll,omitempty"`
	PollAnswer              *PollAnswer                  `json:"poll_answer,omitempty"`
	MyChatMember            *ChatMemberUpdated           `json:"my_chat_member,omitempty"`
	ChatMember              *ChatMemberUpdated           `json:"chat_member,omitempty"`
	ChatJoinRequest         *ChatJoinRequest             `json:"chat_join_request,omitempty"`
	ChatBoost               *ChatBoostUpdated            `json:"chat_boost,omitempty"`
	RemovedChatBoost        *ChatBoostRemoved            `json:"removed_chat_boost,omitempty"`
}

// AllowedUpdate is a type for 'allowed_updates'
type AllowedUpdate string

// AllowedUpdate type constants
//
// https://core.telegram.org/bots/api#update
const (
	AllowMessage                 AllowedUpdate = "message"
	AllowEditedMessage           AllowedUpdate = "edited_message"
	AllowChannelPost             AllowedUpdate = "channel_post"
	AllowEditedChannelPost       AllowedUpdate = "edited_channel_post"
	AllowBusinessConnection      AllowedUpdate = "business_connection"
	AllowBusinessMessage         AllowedUpdate = "business_message"
	AllowEditedBusinessMessage   AllowedUpdate = "edited_business_message"
	AllowDeletedBusinessMessages AllowedUpdate = "deleted_business_messages"
	AllowMessageReaction         AllowedUpdate = "message_reaction"       // NOTE: must be an admin, and need to be explicitly specified
	AllowMessageReactionCount    AllowedUpdate = "message_reaction_count" // NOTE: must be an admin, and need to be explicitly specified
	AllowInlineQuery             AllowedUpdate = "inline_query"
	AllowChosenInlineResult      AllowedUpdate = "chosen_inline_result"
	AllowCallbackQuery           AllowedUpdate = "callback_query"
	AllowShippingQuery           AllowedUpdate = "shipping_query"
	AllowPreCheckoutQuery        AllowedUpdate = "pre_checkout_query"
	AllowPurchasedPaidMedia      AllowedUpdate = "purchased_paid_media"
	AllowPoll                    AllowedUpdate = "poll"
	AllowPollAnswer              AllowedUpdate = "poll_answer"
	AllowMyChatMember            AllowedUpdate = "my_chat_member"
	AllowChatMember              AllowedUpdate = "chat_member"        // NOTE: must be an admin, and need to be explicitly specified
	AllowChatJoinRequest         AllowedUpdate = "chat_join_request"  // NOTE: must have `can_invite_users` admin right
	AllowChatBoost               AllowedUpdate = "chat_boost"         // NOTE: must be an admin
	AllowRemovedChatBoost        AllowedUpdate = "removed_chat_boost" // NOTE: must be an admin
)

// User is a struct of a user
//
// https://core.telegram.org/bots/api#user
type User struct {
	ID                        int64   `json:"id"`
	IsBot                     bool    `json:"is_bot"`
	FirstName                 string  `json:"first_name"`
	LastName                  *string `json:"last_name,omitempty"`
	Username                  *string `json:"username,omitempty"`
	LanguageCode              *string `json:"language_code,omitempty"` // https://en.wikipedia.org/wiki/IETF_language_tag
	IsPremium                 *bool   `json:"is_premium,omitempty"`
	AddedToAttachmentMenu     *bool   `json:"added_to_attachment_menu,omitempty"`
	CanJoinGroups             *bool   `json:"can_join_groups,omitempty"`               // returned only in GetMe()
	CanReadAllGroupMessages   *bool   `json:"can_read_all_group_messages,omitempty"`   // returned only in GetMe()
	SupportsInlineQueries     *bool   `json:"supports_inline_queries,omitempty"`       // returned only in GetMe()
	CanConnectToBusiness      *bool   `json:"can_connect_to_business,omitempty"`       // returned only in GetMe()
	HasMainWebApp             *bool   `json:"has_main_web_app,omitempty"`              // returned only in GetMe()
	HasTopicsEnabled          *bool   `json:"has_topics_enabled,omitempty"`            // returned only in GetMe()
	AllowsUsersToCreateTopics *bool   `json:"allows_users_to_create_topics,omitempty"` // returned only in GetMe()
}

// Chat is a struct of a chat
//
// https://core.telegram.org/bots/api#chat
type Chat struct {
	ID               int64    `json:"id"`
	Type             ChatType `json:"type"`
	Title            *string  `json:"title,omitempty"`
	Username         *string  `json:"username,omitempty"`
	FirstName        *string  `json:"first_name,omitempty"`
	LastName         *string  `json:"last_name,omitempty"`
	IsForum          *bool    `json:"is_forum,omitempty"`
	IsDirectMessages *bool    `json:"is_direct_messages,omitempty"`
}

// ChatFullInfo is a struct for a full info of chat
//
// https://core.telegram.org/bots/api#chatfullinfo
type ChatFullInfo struct {
	ID                                 int64                 `json:"id"`
	Type                               ChatType              `json:"type"`
	Title                              *string               `json:"title,omitempty"`
	Username                           *string               `json:"username,omitempty"`
	FirstName                          *string               `json:"first_name,omitempty"`
	LastName                           *string               `json:"last_name,omitempty"`
	IsForum                            *bool                 `json:"is_forum,omitempty"`
	IsDirectMessages                   *bool                 `json:"is_direct_messages,omitempty"`
	AccentColorID                      int                   `json:"accent_color_id"`
	MaxReactionCount                   int                   `json:"max_reaction_count"`
	Photo                              *ChatPhoto            `json:"photo,omitempty"`
	ActiveUsernames                    []string              `json:"active_usernames,omitempty"`
	Birthdate                          *Birthdate            `json:"birthdate,omitempty"`
	BusinessIntro                      *BusinessIntro        `json:"business_intro,omitempty"`
	BusinessLocation                   *BusinessLocation     `json:"business_location,omitempty"`
	BusinessOpeningHours               *BusinessOpeningHours `json:"business_opening_hours,omitempty"`
	PersonalChat                       *Chat                 `json:"personal_chat,omitempty"`
	ParentChat                         *Chat                 `json:"parent_chat,omitempty"`
	AvailableReactions                 []ReactionType        `json:"available_reactions,omitempty"`
	BackgroundCustomEmojiID            *string               `json:"background_custom_emoji_id,omitempty"`
	ProfileAccentColorID               *int                  `json:"profile_accent_color_id,omitempty"`
	ProfileBackgroundCustomEmojiID     *string               `json:"profile_background_custom_emoji_id,omitempty"`
	EmojiStatusCustomEmojiID           *string               `json:"emoji_status_custom_emoji_id,omitempty"`
	EmojiStatusExpirationDate          *int                  `json:"emoji_status_expiration_date,omitempty"`
	Bio                                *string               `json:"bio,omitempty"`
	HasPrivateForwards                 *bool                 `json:"has_private_forwards,omitempty"`
	HasRestrictedVoiceAndVideoMessages *bool                 `json:"has_restricted_voice_and_video_messages,omitempty"`
	JoinToSendMessages                 *bool                 `json:"join_to_send_messages,omitempty"`
	JoinByRequest                      *bool                 `json:"join_by_request,omitempty"`
	Description                        *string               `json:"description,omitempty"`
	InviteLink                         *string               `json:"invite_link,omitempty"`
	PinnedMessage                      *Message              `json:"pinned_message,omitempty"`
	Permissions                        *ChatPermissions      `json:"permissions,omitempty"`
	AcceptedGiftTypes                  *AcceptedGiftTypes    `json:"accepted_gift_types,omitempty"`
	CanSendPaidMedia                   *bool                 `json:"can_send_paid_media,omitempty"`
	SlowModeDelay                      *int                  `json:"slow_mode_delay,omitempty"`
	UnrestrictBoostCount               *int                  `json:"unrestrict_boost_count,omitempty"`
	MessageAutoDeleteTime              *int                  `json:"message_auto_delete_time,omitempty"`
	HasAggressiveAntiSpamEnabled       *bool                 `json:"has_aggressive_anti_spam_enabled,omitempty"`
	HasHiddenMembers                   *bool                 `json:"has_hidden_members,omitempty"`
	HasProtectedContent                *bool                 `json:"has_protected_content,omitempty"`
	HasVisibleHistory                  *bool                 `json:"has_visible_history,omitempty"`
	StickerSetName                     *string               `json:"sticker_set_name,omitempty"`
	CanSetStickerSet                   *bool                 `json:"can_set_sticker_set,omitempty"`
	CustomEmojiStickerSetName          *string               `json:"custom_emoji_sticker_set_name,omitempty"`
	LinkedChatID                       *int64                `json:"linked_chat_id,omitempty"`
	Location                           *ChatLocation         `json:"location,omitempty"`
	Rating                             *UserRating           `json:"rating,omitempty"`
	FirstProfileAudio                  *Audio                `json:"first_profile_audio,omitempty"`
	UniqueGiftColors                   *UniqueGiftColors     `json:"unique_gift_colors,omitempty"`
	PaidMessageStarCount               *int                  `json:"paid_message_star_count,omitempty"`
}

// InputMediaType is a type of InputMedia
type InputMediaType string

// InputMediaType strings
const (
	InputMediaTypeAnimation InputMediaType = "animation"
	InputMediaTypeDocument  InputMediaType = "document"
	InputMediaTypeAudio     InputMediaType = "audio"
	InputMediaTypePhoto     InputMediaType = "photo"
	InputMediaTypeVideo     InputMediaType = "video"
)

// InputMedia represents the content of a media message to be sent.
//
// NOTE: Can be one of InputMediaAnimation, InputMediaDocument, InputMediaAudio, InputMediaPhoto, or InputMediaVideo.
//
// https://core.telegram.org/bots/api#inputmedia
type InputMedia any

// InputMediaAnimation is a struct of an animation
//
// https://core.telegram.org/bots/api#inputmediaanimation
type InputMediaAnimation struct {
	Type                  InputMediaType  `json:"type"` // == InputMediaTypeAnimation
	Media                 string          `json:"media"`
	Thumbnail             *InputFile      `json:"thumbnail,omitempty"`
	Caption               *string         `json:"caption,omitempty"`
	ParseMode             *ParseMode      `json:"parse_mode,omitempty"`
	CaptionEntities       []MessageEntity `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia *bool           `json:"show_caption_above_media,omitempty"`
	Width                 *int            `json:"width,omitempty"`
	Height                *int            `json:"height,omitempty"`
	Duration              *int            `json:"duration,omitempty"`
	HasSpoiler            *bool           `json:"has_spoiler,omitempty"`
}

// InputMediaDocument is a struct of a document
//
// https://core.telegram.org/bots/api#inputmediadocument
type InputMediaDocument struct {
	Type                        InputMediaType  `json:"type"` // == InputMediaTypeDocument
	Media                       string          `json:"media"`
	Thumbnail                   *InputFile      `json:"thumbnail,omitempty"`
	Caption                     *string         `json:"caption,omitempty"`
	ParseMode                   *ParseMode      `json:"parse_mode,omitempty"`
	CaptionEntities             []MessageEntity `json:"caption_entities,omitempty"`
	DisableContentTypeDetection *bool           `json:"disable_content_type_detection,omitempty"`
}

// InputMediaAudio is a struct of an audio
//
// https://core.telegram.org/bots/api#inputmediaaudio
type InputMediaAudio struct {
	Type            InputMediaType  `json:"type"` // == InputMediaTypeAudio
	Media           string          `json:"media"`
	Thumbnail       *InputFile      `json:"thumbnail,omitempty"`
	Caption         *string         `json:"caption,omitempty"`
	ParseMode       *ParseMode      `json:"parse_mode,omitempty"`
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
	Duration        *int            `json:"duration,omitempty"`
	Performer       *string         `json:"performer,omitempty"`
	Title           *string         `json:"title,omitempty"`
}

// InputMediaPhoto is a struct of a photo
//
// https://core.telegram.org/bots/api#inputmediaphoto
type InputMediaPhoto struct {
	Type                  InputMediaType  `json:"type"` // == InputMediaTypePhoto
	Media                 string          `json:"media"`
	Caption               *string         `json:"caption,omitempty"`
	ParseMode             *ParseMode      `json:"parse_mode,omitempty"`
	CaptionEntities       []MessageEntity `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia *bool           `json:"show_caption_above_media,omitempty"`
	HasSpoiler            *bool           `json:"has_spoiler,omitempty"`
}

// InputMediaVideo is a struct of a video
//
// https://core.telegram.org/bots/api#inputmediavideo
type InputMediaVideo struct {
	Type                  InputMediaType  `json:"type"` // == InputMediaTypeVideo
	Media                 string          `json:"media"`
	Thumbnail             *InputFile      `json:"thumbnail,omitempty"`
	Cover                 *string         `json:"cover,omitempty"`
	StartTimestamp        *int            `json:"start_timestamp,omitempty"`
	Caption               *string         `json:"caption,omitempty"`
	ParseMode             *ParseMode      `json:"parse_mode,omitempty"`
	CaptionEntities       []MessageEntity `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia *bool           `json:"show_caption_above_media,omitempty"`
	Width                 *int            `json:"width,omitempty"`
	Height                *int            `json:"height,omitempty"`
	Duration              *int            `json:"duration,omitempty"`
	SupportsStreaming     *bool           `json:"supports_streaming,omitempty"`
	HasSpoiler            *bool           `json:"has_spoiler,omitempty"`
}

// InputFile represents contents of a file to be uploaded.
//
// NOTE: Can be generated with NewInputFileFromXXX() functions in types_helper.go
//
// https://core.telegram.org/bots/api#inputfile
type InputFile struct {
	Filepath *string
	URL      *string
	Bytes    []byte
	FileID   *string
}

// InputPaidMedia can be one of `InputPaidMediaPhoto` or `InputPaidMediaVideo`
//
// https://core.telegram.org/bots/api#inputpaidmedia
type InputPaidMedia any

// InputPaidMediaPhoto struct
//
// https://core.telegram.org/bots/api#inputpaidmediaphoto
type InputPaidMediaPhoto struct {
	Type  string `json:"type"` // == "photo"
	Media string `json:"media"`
}

// InputPaidMediaVideo struct
//
// https://core.telegram.org/bots/api#inputpaidmediavideo
type InputPaidMediaVideo struct {
	Type              string  `json:"type"` // == "video"
	Media             string  `json:"media"`
	Thumbnail         any     `json:"thumbnail,omitempty"` // `InputFile` or string
	Cover             *string `json:"cover,omitempty"`
	StartTimestamp    *int    `json:"start_timestamp,omitempty"`
	Width             *int    `json:"width,omitempty"`
	Height            *int    `json:"height,omitempty"`
	Duration          *int    `json:"duration,omitempty"`
	SupportsStreaming *bool   `json:"supports_streaming,omitempty"`
}

// StickerFormat is a format of sticker
type StickerFormat string

// StickerFormat strings
const (
	StickerFormatStatic   StickerFormat = "static"
	StickerFormatAnimated StickerFormat = "animated"
	StickerFormatVideo    StickerFormat = "video"
)

// StickerType is a type of sticker
type StickerType string

// StickerType strings
const (
	StickerTypeRegular     StickerType = "regular"
	StickerTypeMask        StickerType = "mask"
	StickerTypeCustomEmoji StickerType = "custom_emoji"
)

// Audio is a struct for an audio file
//
// https://core.telegram.org/bots/api#audio
type Audio struct {
	FileID       string     `json:"file_id"`
	FileUniqueID string     `json:"file_unique_id"`
	Duration     int        `json:"duration"`
	Performer    *string    `json:"performer,omitempty"`
	Title        *string    `json:"title,omitempty"`
	FileName     *string    `json:"file_name,omitempty"`
	MimeType     *string    `json:"mime_type,omitempty"`
	FileSize     *int       `json:"file_size,omitempty"`
	Thumbnail    *PhotoSize `json:"thumbnail,omitempty"`
}

// MessageEntity is a struct of a message entity
//
// NOTE: Can be generated with NewMessageEntity() function in types_helper.go
//
// https://core.telegram.org/bots/api#messageentity
type MessageEntity struct {
	Type          MessageEntityType `json:"type"`
	Offset        int               `json:"offset"`
	Length        int               `json:"length"`
	URL           *string           `json:"url,omitempty"`             // when Type == MessageEntityTypeTextLink
	User          *User             `json:"user,omitempty"`            // when Type == MessageEntityTypeTextMention
	Language      *string           `json:"language,omitempty"`        // when Type == MessageEntityTypePre
	CustomEmojiID *string           `json:"custom_emoji_id,omitempty"` // when Type == MessageEntityTypeCustomEmoji
}

// TextQuote is a struct of a text quote
//
// https://core.telegram.org/bots/api#textquote
type TextQuote struct {
	Text     string          `json:"text"`
	Entities []MessageEntity `json:"entities,omitempty"`
	Position int             `json:"position"`
	IsManual *bool           `json:"is_manual,omitempty"`
}

// ExternalReplyInfo is a struct of an external reply info of a message
//
// https://core.telegram.org/bots/api#externalreplyinfo
type ExternalReplyInfo struct {
	Origin             MessageOrigin       `json:"origin"`
	Chat               *Chat               `json:"chat,omitempty"`
	MessageID          *int64              `json:"message_id,omitempty"`
	LinkPreviewOptions *LinkPreviewOptions `json:"link_preview_options,omitempty"`
	Animation          *Animation          `json:"animation,omitempty"`
	Audio              *Audio              `json:"audio,omitempty"`
	Document           *Document           `json:"document,omitempty"`
	PaidMedia          *PaidMediaInfo      `json:"paid_media,omitempty"`
	Photo              []PhotoSize         `json:"photo,omitempty"`
	Sticker            *Sticker            `json:"sticker,omitempty"`
	Story              *Story              `json:"story,omitempty"`
	Video              *Video              `json:"video,omitempty"`
	VideoNote          *VideoNote          `json:"video_note,omitempty"`
	Voice              *Voice              `json:"voice,omitempty"`
	HasMediaSpoiler    bool                `json:"has_media_spoiler,omitempty"`
	Checklist          *Checklist          `json:"checklist,omitempty"`
	Contact            *Contact            `json:"contact,omitempty"`
	Dice               *Dice               `json:"dice,omitempty"`
	Game               *Game               `json:"game,omitempty"`
	Giveaway           *Giveaway           `json:"giveaway,omitempty"`
	GiveawayWinners    *GiveawayWinners    `json:"giveaway_winners,omitempty"`
	Invoice            *Invoice            `json:"invoice,omitempty"`
	Location           *Location           `json:"location,omitempty"`
	Poll               *Poll               `json:"poll,omitempty"`
	Venue              *Venue              `json:"venue,omitempty"`
}

// ReplyParameters is a struct for replying messages
//
// NOTE: Can be generated with NewReplyParameters() function in types_helper.go
//
// https://core.telegram.org/bots/api#replyparameters
type ReplyParameters struct {
	MessageID                int64           `json:"message_id"`
	ChatID                   *ChatID         `json:"chat_id,omitempty"`
	AllowSendingWithoutReply *bool           `json:"allow_sending_without_reply,omitempty"`
	Quote                    *string         `json:"quote,omitempty"`
	QuoteParseMode           *ParseMode      `json:"quote_parse_mode,omitempty"`
	QuoteEntities            []MessageEntity `json:"quote_entities,omitempty"`
	QuotePosition            *int            `json:"quote_position,omitempty"`
	ChecklistTaskID          *int64          `json:"checklist_task_id,omitempty"`
}

// MessageOrigin struct for describing the origin of a message
//
// https://core.telegram.org/bots/api#messageorigin
type MessageOrigin struct {
	Type string `json:"type"`
	Date int    `json:"date"`

	// https://core.telegram.org/bots/api#messageoriginuser
	SenderUser *User `json:"sender_user,omitempty"`

	// https://core.telegram.org/bots/api#messageoriginhiddenuser
	SenderUserName *string `json:"sender_user_name,omitempty"`

	// https://core.telegram.org/bots/api#messageoriginchat
	SenderChat      *Chat   `json:"sender_chat,omitempty"`
	AuthorSignature *string `json:"author_signature,omitempty"`

	// https://core.telegram.org/bots/api#messageoriginchannel
	Chat      *Chat  `json:"chat,omitempty"`
	MessageID *int64 `json:"message_id,omitempty"`
	// AuthorSignature *string `json:"author_signature,omitempty"`
}

// PhotoSize is a struct of a photo's size
//
// https://core.telegram.org/bots/api#photosize
type PhotoSize struct {
	FileID       string `json:"file_id"`
	FileUniqueID string `json:"file_unique_id"`
	Width        int    `json:"width"`
	Height       int    `json:"height"`
	FileSize     *int   `json:"file_size,omitempty"`
}

// Document is a struct for an ordinary file
//
// https://core.telegram.org/bots/api#document
type Document struct {
	FileID       string     `json:"file_id"`
	FileUniqueID string     `json:"file_unique_id"`
	Thumbnail    *PhotoSize `json:"thumbnail,omitempty"`
	FileName     *string    `json:"file_name,omitempty"`
	MimeType     *string    `json:"mime_type,omitempty"`
	FileSize     int        `json:"file_size,omitempty"`
}

// Sticker is a struct of a sticker
//
// https://core.telegram.org/bots/api#sticker
type Sticker struct {
	FileID           string        `json:"file_id"`
	FileUniqueID     string        `json:"file_unique_id"`
	Type             StickerType   `json:"type"`
	Width            int           `json:"width"`
	Height           int           `json:"height"`
	IsAnimated       bool          `json:"is_animated"`
	IsVideo          bool          `json:"is_video"`
	Thumbnail        *PhotoSize    `json:"thumbnail,omitempty"`
	Emoji            *string       `json:"emoji,omitempty"`
	SetName          *string       `json:"set_name,omitempty"`
	PremiumAnimation *File         `json:"premium_animation,omitempty"`
	MaskPosition     *MaskPosition `json:"mask_position,omitempty"`
	CustomEmojiID    *string       `json:"custom_emoji_id,omitempty"`
	NeedsRepainting  *bool         `json:"needs_repainting,omitempty"`
	FileSize         *int          `json:"file_size,omitempty"`
}

// StickerSet is a struct of a sticker set
//
// https://core.telegram.org/bots/api#stickerset
type StickerSet struct {
	Name        string      `json:"name"`
	Title       string      `json:"title"`
	StickerType StickerType `json:"sticker_type"`
	Stickers    []Sticker   `json:"stickers"`
	Thumbnail   *PhotoSize  `json:"thumbnail,omitempty"`
}

// MaskPosition is a struct for a mask position
//
// NOTE: Can be generated with NewMaskPosition() function in types_helper.go
//
// https://core.telegram.org/bots/api#maskposition
type MaskPosition struct {
	Point  MaskPositionPoint `json:"point"`
	XShift float32           `json:"x_shift"`
	YShift float32           `json:"y_shift"`
	Scale  float32           `json:"scale"`
}

// InputSticker is a struct for a sticker
//
// NOTE: Can be generated with NewInputSticker() function in types_helper.go
//
// https://core.telegram.org/bots/api#inputsticker
type InputSticker struct {
	Sticker      any           `json:"sticker"` // InputFile or `file_id`
	Format       StickerFormat `json:"format"`  // "static" for .webp or .png, "animated" for .tgs, "video" for .webm
	EmojiList    []string      `json:"emoji_list"`
	MaskPosition *MaskPosition `json:"mask_position,omitempty"`
	Keywords     []string      `json:"keywords,omitempty"`
}

// Story is a struct for a forwarded story of a message
//
// https://core.telegram.org/bots/api#story
type Story struct {
	Chat Chat  `json:"chat"`
	ID   int64 `json:"id"`
}

// VideoQuality is a struct for a video quality
//
// https://core.telegram.org/bots/api#videoquality
type VideoQuality struct {
	FileID       string `json:"file_id"`
	FileUniqueID string `json:"file_unique_id"`
	Width        int    `json:"width"`
	Height       int    `json:"height"`
	Codec        string `json:"codec"` // eg. "h264", "h265", or "av01"
	FileSize     *int64 `json:"file_size,omitempty"`
}

// Video is a struct for a video file
//
// https://core.telegram.org/bots/api#video
type Video struct {
	FileID         string         `json:"file_id"`
	FileUniqueID   string         `json:"file_unique_id"`
	Width          int            `json:"width"`
	Height         int            `json:"height"`
	Duration       int            `json:"duration"`
	Thumbnail      *PhotoSize     `json:"thumbnail,omitempty"`
	Cover          []PhotoSize    `json:"cover,omitempty"`
	StartTimestamp *int           `json:"start_timestamp,omitempty"`
	Qualities      []VideoQuality `json:"qualities,omitempty"`
	FileName       *string        `json:"file_name,omitempty"`
	MimeType       *string        `json:"mime_type,omitempty"`
	FileSize       *int           `json:"file_size,omitempty"`
}

// PaidMediaInfo struct
//
// https://core.telegram.org/bots/api#paidmediainfo
type PaidMediaInfo struct {
	StarCount int         `json:"star_count"`
	PaidMedia []PaidMedia `json:"paid_media"`
}

// PaidMedia can be one of `PaidMediaPreview`, `PaidMediaPhoto`, or `PaidMediaVideo`
//
// https://core.telegram.org/bots/api#paidmedia
type PaidMedia any

// PaidMediaPreview struct
//
// https://core.telegram.org/bots/api#paidmediapreview
type PaidMediaPreview struct {
	Type     string `json:"type"` // == "preview"
	Width    int    `json:"width"`
	Height   int    `json:"height"`
	Duration int    `json:"duration"`
}

// PaidMediaPhoto struct
//
// https://core.telegram.org/bots/api#paidmediaphoto
type PaidMediaPhoto struct {
	Type  string      `json:"type"` // == "photo"
	Photo []PhotoSize `json:"photo"`
}

// PaidMediaVideo struct
//
// https://core.telegram.org/bots/api#paidmediavideo
type PaidMediaVideo struct {
	Type  string `json:"type"` // == "video"
	Video Video  `json:"video"`
}

// Voice is a struct for a voice file
//
// https://core.telegram.org/bots/api#voice
type Voice struct {
	FileID       string  `json:"file_id"`
	FileUniqueID string  `json:"file_unique_id"`
	Duration     int     `json:"duration"`
	MimeType     *string `json:"mime_type,omitempty"`
	FileSize     *int    `json:"file_size,omitempty"`
}

// VideoNote is a struct for a video note
//
// https://core.telegram.org/bots/api#videonote
type VideoNote struct {
	FileID       string     `json:"file_id"`
	FileUniqueID string     `json:"file_unique_id"`
	Length       int        `json:"length"`
	Duration     int        `json:"duration"`
	Thumbnail    *PhotoSize `json:"thumbnail,omitempty"`
	FileSize     *int       `json:"file_size,omitempty"`
}

// Contact is a struct for a contact info
//
// https://core.telegram.org/bots/api#contact
type Contact struct {
	PhoneNumber string  `json:"phone_number"`
	FirstName   string  `json:"first_name"`
	LastName    *string `json:"last_name,omitempty"`
	UserID      *int64  `json:"user_id,omitempty"`
	VCard       *string `json:"vcard,omitempty"` // https://en.wikipedia.org/wiki/VCard
}

// Location is a struct for a location
//
// https://core.telegram.org/bots/api#location
type Location struct {
	Longitude            float32 `json:"longitude"`
	Latitude             float32 `json:"latitude"`
	HorizontalAccuracy   float32 `json:"horizontal_accuracy,omitempty"`
	LivePeriod           *int    `json:"live_period,omitempty"`
	Heading              *int    `json:"heading,omitempty"`
	ProximityAlertRadius *int    `json:"proximity_alert_radius,omitempty"`
}

// Venue is a struct of a venue
//
// https://core.telegram.org/bots/api#venue
type Venue struct {
	Location        Location `json:"location"`
	Title           string   `json:"title"`
	Address         string   `json:"address"`
	FoursquareID    *string  `json:"foursquare_id,omitempty"`
	FoursquareType  *string  `json:"foursquare_type,omitempty"`
	GooglePlaceID   *string  `json:"google_place_id,omitempty"`
	GooglePlaceType *string  `json:"google_place_type,omitempty"`
}

// WebAppData is a struct of a web app data
//
// https://core.telegram.org/bots/api#webappdata
type WebAppData struct {
	Data       string `json:"data"`
	ButtonText string `json:"button_text"`
}

// ProximityAlertTriggered is a struct of priximity alert triggered object
//
// https://core.telegram.org/bots/api#proximityalerttriggered
type ProximityAlertTriggered struct {
	Traveler User `json:"traveler"`
	Watcher  User `json:"watcher"`
	Distance int  `json:"distance"`
}

// ChatBoostAdded is a struct of an added boost to a chat
//
// https://core.telegram.org/bots/api#chatboostadded
type ChatBoostAdded struct {
	BoostCount int `json:"boost_count"`
}

// Poll is a struct of a poll
//
// https://core.telegram.org/bots/api#poll
type Poll struct {
	ID                    string          `json:"id"`
	Question              string          `json:"question"` // 1~255 chars
	QuestionEntities      []MessageEntity `json:"question_entities,omitempty"`
	Options               []PollOption    `json:"options"`
	TotalVoterCount       int             `json:"total_voter_count"`
	IsClosed              bool            `json:"is_closed"`
	IsAnonymous           bool            `json:"is_anonymous"`
	Type                  string          `json:"type"` // "quiz" or "regular"
	AllowsMultipleAnswers bool            `json:"allows_multiple_answers"`
	CorrectOptionID       *int            `json:"correct_option_id,omitempty"`
	Explanation           *string         `json:"explanation,omitempty"`
	ExplanationEntities   []MessageEntity `json:"explanation_entities,omitempty"`
	OpenPeriod            *int            `json:"open_period,omitempty"`
	CloseDate             *int            `json:"close_date,omitempty"`
}

// PollOption is a struct of a poll option
//
// https://core.telegram.org/bots/api#polloption
type PollOption struct {
	Text         string          `json:"text"` // 1~100 chars
	TextEntities []MessageEntity `json:"text_entities,omitempty"`
	VoterCount   int             `json:"voter_count"`
}

// InputPollOption is a struct of an input poll option
//
// https://core.telegram.org/bots/api#inputpolloption
type InputPollOption struct {
	Text          string          `json:"text"` // 1~100 chars
	TextParseMode ParseMode       `json:"text_parse_mode,omitempty"`
	TextEntities  []MessageEntity `json:"text_entities,omitempty"`
}

// PollAnswer is a struct of a poll answer
//
// https://core.telegram.org/bots/api#pollanswer
type PollAnswer struct {
	PollID    string `json:"poll_id"`
	VoterChat *Chat  `json:"voter_chat,omitempty"`
	User      *User  `json:"user,omitempty"`
	OptionIDs []int  `json:"option_ids"`
}

// ChecklistTask is a struct of a checklist task
//
// https://core.telegram.org/bots/api#checklisttask
type ChecklistTask struct {
	ID              int64           `json:"id"`
	Text            string          `json:"text"`
	TextEntities    []MessageEntity `json:"text_entities,omitempty"`
	CompletedByUser *User           `json:"completed_by_user,omitempty"`
	CompletedByChat *Chat           `json:"completed_by_chat,omitempty"`
	CompletionDate  *int64          `json:"completion_date,omitempty"`
}

// Checklist is a struct of a checklist
//
// https://core.telegram.org/bots/api#checklist
type Checklist struct {
	Title                    string          `json:"title"`
	TitleEntities            []MessageEntity `json:"title_entities,omitempty"`
	Tasks                    []ChecklistTask `json:"tasks"`
	OthersCanAddTasks        *bool           `json:"others_can_add_tasks,omitempty"`
	OthersCanMarkTasksAsDone *bool           `json:"others_can_mark_tasks_as_done,omitempty"`
}

// InputChecklistTask is a struct of an input checklist task
//
// https://core.telegram.org/bots/api#inputchecklisttask
type InputChecklistTask struct {
	ID           int64           `json:"id"`
	Text         string          `json:"text"`
	ParseMode    *ParseMode      `json:"parse_mode,omitempty"`
	TextEntities []MessageEntity `json:"text_entities,omitempty"`
}

// InputChecklist is a struct of an input checklist
//
// https://core.telegram.org/bots/api#inputchecklist
type InputChecklist struct {
	Title                    string               `json:"title"`
	ParseMode                *ParseMode           `json:"parse_mode,omitempty"`
	TitleEntities            []MessageEntity      `json:"title_entities,omitempty"`
	Tasks                    []InputChecklistTask `json:"tasks"`
	OthersCanAddTasks        *bool                `json:"others_can_add_tasks,omitempty"`
	OthersCanMarkTasksAsDone *bool                `json:"others_can_mark_tasks_as_done,omitempty"`
}

// ChecklistTasksDone is a struct for service message: checklist tasks done
//
// https://core.telegram.org/bots/api#checklisttasksdone
type ChecklistTasksDone struct {
	ChecklistMessage       *Message `json:"checklist_message,omitempty"`
	MarkedAsDoneTaskIDs    []int64  `json:"marked_as_done_task_ids,omitempty"`
	MarkedAsNotDoneTaskIDs []int64  `json:"marked_as_not_done_task_ids,omitempty"`
}

// ChecklistTasksAdded is a struct for service message: checklist tasks added
//
// https://core.telegram.org/bots/api#checklisttasksadded
type ChecklistTasksAdded struct {
	ChecklistMessage *Message        `json:"checklist_message,omitempty"`
	Tasks            []ChecklistTask `json:"tasks"`
}

// Dice is a struct for dice in message
//
// https://core.telegram.org/bots/api#dice
type Dice struct {
	Emoji string `json:"emoji"`
	Value int    `json:"value"` // 1-6 for dice, dart, and bowling; 1-5 for basketball and football; 1-64 for slotmachine;
}

// MessageAutoDeleteTimerChanged is service message: message auto delete timer changed
//
// https://core.telegram.org/bots/api#messageautodeletetimerchanged
type MessageAutoDeleteTimerChanged struct {
	MessageAutoDeleteTime int `json:"message_auto_delete_time"`
}

// ChatBackground is a struct for chat background
//
// https://core.telegram.org/bots/api#chatbackground
type ChatBackground struct {
	Type BackgroundType `json:"type"`
}

// BackgroundTypeType for types of BackgroundType
type BackgroundTypeType string

// BackgroundTypeType constants
const (
	BackgroundTypeFill      BackgroundTypeType = "fill"       // https://core.telegram.org/bots/api#backgroundtypefill
	BackgroundTypeWallpaper BackgroundTypeType = "wallpaper"  // https://core.telegram.org/bots/api#backgroundtypewallpaper
	BackgroundTypePattern   BackgroundTypeType = "pattern"    // https://core.telegram.org/bots/api#backgroundtypepattern
	BackgroundTypeChatTheme BackgroundTypeType = "chat_theme" // https://core.telegram.org/bots/api#backgroundtypechattheme
)

// BackgroundType is a struct for a type of background
//
// https://core.telegram.org/bots/api#backgroundtype
type BackgroundType struct {
	Type BackgroundTypeType `json:"type"`

	// type == "fill"
	Fill             *BackgroundFill `json:"fill,omitempty"`
	DarkThemeDimming *int            `json:"dark_theme_dimming,omitempty"`

	// type == "wallpaper"
	Document *Document `json:"document,omitempty"`
	// DarkThemeDimming *int `json:"dark_theme_dimming,omitempty"`
	IsBlurred *bool `json:"is_blurred,omitempty"`
	IsMoving  *bool `json:"is_moving,omitempty"`

	// type == "pattern"
	// Document *Document `json:"document,omitempty"`
	// Fill *BackgroundFill `json:"fill,omitempty"`
	Intensity  *int  `json:"intensity,omitempty"`
	IsInverted *bool `json:"is_inverted,omitempty"`
	// IsMoving *bool `json:"is_moving,omitempty"`

	// type == "chat_theme"
	ThemeName *string `json:"theme_name,omitempty"`
}

// BackgroundFillType for types of BackgroundFill
type BackgroundFillType string

// BackgroundFillType constants
const (
	BackgroundFillTypeSolid            BackgroundFillType = "solid"
	BackgroundFillTypeGradient         BackgroundFillType = "gradient"
	BackgroundFillTypeFreeformGradient BackgroundFillType = "freeform_gradient"
)

// BackgroundFill is a struct for a type of background fill
//
// https://core.telegram.org/bots/api#backgroundfill
type BackgroundFill struct {
	Type BackgroundFillType `json:"type"`

	// type == "solid"
	Color *int `json:"color,omitempty"`

	// type == "gradient"
	TopColor      *int `json:"top_color,omitempty"`      // RGB24
	BottomColor   *int `json:"bottom_color,omitempty"`   // RGB24
	RotationAngle *int `json:"rotation_angle,omitempty"` // 0-359

	// type == "freeform_gradient"
	Colors []int `json:"colors,omitempty"`
}

// ForumTopicCreated is a struct for a new forum topic created in the chat.
//
// https://core.telegram.org/bots/api#forumtopiccreated
type ForumTopicCreated struct {
	Name              string  `json:"name"`
	IconColor         int     `json:"icon_color"`
	IconCustomEmojiID *string `json:"icon_custom_emoji_id,omitempty"`
	IsNameImplicit    *bool   `json:"is_name_implicit,omitempty"`
}

// ForumTopicClosed is a struct for a closed forum topic in the chat.
//
// https://core.telegram.org/bots/api#forumtopicclosed
type ForumTopicClosed struct{}

// ForumTopicEdited is a struct for a edited forum topic in the chat.
//
// https://core.telegram.org/bots/api#forumtopicedited
type ForumTopicEdited struct {
	Name              *string `json:"name,omitempty"`
	IconCustomEmojiID *string `json:"icon_custom_emoji_id,omitempty"`
}

// ForumTopicReopened is a struct for a reopened forum topic in the chat.
//
// https://core.telegram.org/bots/api#forumtopicreopened
type ForumTopicReopened struct{}

// GeneralForumTopicHidden is a struct for a hidden general forum topic in the chat.
//
// https://core.telegram.org/bots/api#generalforumtopichidden
type GeneralForumTopicHidden struct{}

// GeneralForumTopicUnhidden is a struct for an unhidden general forum topic in the chat.
//
// https://core.telegram.org/bots/api#generalforumtopicunhidden
type GeneralForumTopicUnhidden struct{}

// SharedUser is a struct for a user which was shared with the bot using KeyboardButtonRequestUser button.
//
// https://core.telegram.org/bots/api#shareduser
type SharedUser struct {
	UserID    int64       `json:"user_id"`
	FirstName *string     `json:"first_name,omitempty"`
	LastName  *string     `json:"last_name,omitempty"`
	Username  *string     `json:"username,omitempty"`
	Photo     []PhotoSize `json:"photo,omitempty"`
}

// UsersShared is a struct for users who shared the message.
//
// https://core.telegram.org/bots/api#usersshared
type UsersShared struct {
	RequestID int64        `json:"request_id"`
	Users     []SharedUser `json:"users"`
}

// ChatShared is a struct for a chat which shared the message.
//
// https://core.telegram.org/bots/api#chatshared
type ChatShared struct {
	RequestID int64       `json:"request_id"`
	ChatID    int64       `json:"chat_id"`
	Title     *string     `json:"title,omitempty"`
	Username  *string     `json:"username,omitempty"`
	Photo     []PhotoSize `json:"photo,omitempty"`
}

// WriteAccessAllowed is a struct for an allowed write access in the chat.
//
// https://core.telegram.org/bots/api#writeaccessallowed
type WriteAccessAllowed struct {
	FromRequest        *bool   `json:"from_request,omitempty"`
	WebAppName         *string `json:"web_app_name,omitempty"`
	FromAttachmentMenu *bool   `json:"from_attachment_menu,omitempty"`
}

// VideoChatStarted is a struct for service message: video chat started.
//
// https://core.telegram.org/bots/api#videochatstarted
type VideoChatStarted struct{}

// VideoChatEnded is a struct for service message: video chat ended
//
// https://core.telegram.org/bots/api#videochatended
type VideoChatEnded struct {
	Duration int `json:"duration"`
}

// VideoChatScheduled is a struct for servoice message: video chat scheduled
//
// https://core.telegram.org/bots/api#videochatscheduled
type VideoChatScheduled struct {
	StartDate int `json:"start_date"`
}

// VideoChatParticipantsInvited is a struct for service message: new members invited to video chat
//
// https://core.telegram.org/bots/api#videochatparticipantsinvited
type VideoChatParticipantsInvited struct {
	Users []User `json:"users"`
}

// PaidMessagePriceChanged describes a service message about a change in the price of paid messages within a chat.
//
// https://core.telegram.org/bots/api#paidmessagepricechanged
type PaidMessagePriceChanged struct {
	PaidMessageStarCount int `json:"paid_message_star_count"`
}

// DirectMessagePriceChanged is a struct for a service message about a change in the price of direct messages.
//
// https://core.telegram.org/bots/api#directmessagepricechanged
type DirectMessagePriceChanged struct {
	AreDirectMessagesEnabled bool `json:"are_direct_messages_enabled"`
	DirectMessageStarCount   *int `json:"direct_message_star_count,omitempty"`
}

// SuggestedPostApproved is a struct for a service message about the approval of a suggested post.
//
// https://core.telegram.org/bots/api#suggestedpostapproved
type SuggestedPostApproved struct {
	SuggestedPostMessage *Message            `json:"suggested_post_message,omitempty"`
	Price                *SuggestedPostPrice `json:"price,omitempty"`
	SendDate             int                 `json:"send_date"`
}

// SuggestedPostApprovalFailed is a struct for a service message about the failed approval of a suggested post.
//
// https://core.telegram.org/bots/api#suggestedpostapprovalfailed
type SuggestedPostApprovalFailed struct {
	SuggestedPostMessage *Message           `json:"suggested_post_message,omitempty"`
	Price                SuggestedPostPrice `json:"price"`
}

// SuggestedPostDeclined is a struct for a service message about the rejection of a suggested post.
//
// https://core.telegram.org/bots/api#suggestedpostdeclined
type SuggestedPostDeclined struct {
	SuggestedPostMessage *Message `json:"suggested_post_message,omitempty"`
	Comment              *string  `json:"comment,omitempty"`
}

// SuggestedPostPaid is a struct for a service message about a successful payment for a suggested post.
//
// https://core.telegram.org/bots/api#suggestedpostpaid
type SuggestedPostPaid struct {
	SuggestedPostMessage *Message    `json:"suggested_post_message,omitempty"`
	Currency             string      `jso:"currency"`
	Amount               int         `json:"amount,omitempty"`
	StarAmount           *StarAmount `json:"star_amount,omitempty"`
}

// SuggestedPostRefunded is a struct for a service message about a payment refund for a suggested post.
//
// https://core.telegram.org/bots/api#suggestedpostrefunded
type SuggestedPostRefunded struct {
	SuggestedPostMessage *Message `json:"suggested_post_message,omitempty"`
	Reason               string   `json:"reason"`
}

// Giveaway struct for giveaways
//
// https://core.telegram.org/bots/api#giveaway
type Giveaway struct {
	Chats                         []Chat   `json:"chats"`
	WinnersSelectionDate          int      `json:"winners_selection_date"`
	WinnerCount                   int      `json:"winner_count"`
	OnlyNewMembers                *bool    `json:"only_new_members,omitempty"`
	HasPublicWinners              *bool    `json:"has_public_winners,omitempty"`
	PrizeDescription              *string  `json:"prize_description,omitempty"`
	CountryCodes                  []string `json:"country_codes,omitempty"`
	PrizeStarCount                *int     `json:"prize_star_count,omitempty"`
	PremiumSubscriptionMonthCount *int     `json:"premium_subscription_month_count,omitempty"`
}

// GiveawayCreated struct for service message about the creation of giveaway
//
// https://core.telegram.org/bots/api#giveawaycreated
type GiveawayCreated struct {
	PrizeStarCount *int `json:"prize_star_count,omitempty"`
}

// GiveawayWinners struct for representing the completion of a giveaway with public winners
//
// https://core.telegram.org/bots/api#giveawaywinners
type GiveawayWinners struct {
	Chat                          Chat    `json:"chat"`
	GiveawayMessageID             int64   `json:"giveaway_message_id"`
	WinnersSelectionDate          int     `json:"winners_selection_date"`
	WinnerCount                   int     `json:"winner_count"`
	Winners                       []User  `json:"winners"`
	AdditionalChatCount           *int    `json:"additional_chat_count,omitempty"`
	PrizeStarCount                *int    `json:"prize_star_count,omitempty"`
	PremiumSubscriptionMonthCount *int    `json:"premium_subscription_month_count,omitempty"`
	UnclaimedPrizeCount           *int    `json:"unclaimed_prize_count,omitempty"`
	OnlyNewMembers                *bool   `json:"only_new_members,omitempty"`
	WasRefunded                   *bool   `json:"was_refunded,omitempty"`
	PrizeDescription              *string `json:"prize_description,omitempty"`
}

// GiveawayCompleted struct for service message about the completion of a giveaway without public winners
//
// https://core.telegram.org/bots/api#giveawaycompleted
type GiveawayCompleted struct {
	WinnerCount         int      `json:"winner_count"`
	UnclaimedPrizeCount *int     `json:"unclaimed_prize_count,omitempty"`
	GiveawayMessage     *Message `json:"giveaway_message,omitempty"`
	IsStarGiveaway      *bool    `json:"is_star_giveaway,omitempty"`
}

// LinkPreviewOptions is a struct for link preview
//
// NOTE: Can be generated with NewLinkPreviewOptions() function in types_helper.go
//
// https://core.telegram.org/bots/api#linkpreviewoptions
type LinkPreviewOptions struct {
	IsDisabled       *bool   `json:"is_disabled,omitempty"`
	URL              *string `json:"url,omitempty"`
	PreferSmallMedia *bool   `json:"prefer_small_media,omitempty"`
	PreferLargeMedia *bool   `json:"prefer_large_media,omitempty"`
	ShowAboveText    *bool   `json:"show_above_text,omitempty"`
}

// SuggestedPostPrice is a struct for suggested post price
//
// https://core.telegram.org/bots/api#suggestedpostprice
type SuggestedPostPrice struct {
	Currency string `json:"currency"`
	Amount   int64  `json:"amount"`
}

// SuggestedPostStateType is a type of suggested post info's state
type SuggestedPostStateType string

const (
	SuggestedPostStatePending  SuggestedPostStateType = "pending"
	SuggestedPostStateApproved SuggestedPostStateType = "approved"
	SuggestedPostStateDeclined SuggestedPostStateType = "declined"
)

// SuggestedPostInfo is a struct for suggested post info
//
// https://core.telegram.org/bots/api#suggestedpostinfo
type SuggestedPostInfo struct {
	State    SuggestedPostStateType `json:"state"`
	Price    *SuggestedPostPrice    `json:"price,omitempty"`
	SendDate *int                   `json:"send_date,omitempty"`
}

// SuggestedPostParameters is a struct for suggested post parameters
//
// https://core.telegram.org/bots/api#suggestedpostparameters
type SuggestedPostParameters struct {
	Price    *SuggestedPostPrice `json:"price,omitempty"`
	SendDate *int                `json:"send_date,omitempty"`
}

// DirectMessagesTopic is a struct for direct messages topic
//
// https://core.telegram.org/bots/api#directmessagestopic
type DirectMessagesTopic struct {
	TopicID int64 `json:"topic_id"`
	User    *User `json:"user,omitempty"`
}

// UserProfilePhotos is a struct for user profile photos
//
// https://core.telegram.org/bots/api#userprofilephotos
type UserProfilePhotos struct {
	TotalCount int           `json:"total_count"`
	Photos     [][]PhotoSize `json:"photos"`
}

// UserProfileAudios is a struct for user profile audios
//
// https://core.telegram.org/bots/api#userprofileaudios
type UserProfileAudios struct {
	TotalCount int     `json:"total_count"`
	Audios     []Audio `json:"audios"`
}

// File is a struct for a file
//
// https://core.telegram.org/bots/api#file
type File struct {
	FileID       string  `json:"file_id"`
	FileUniqueID string  `json:"file_unique_id"`
	FileSize     *int    `json:"file_size,omitempty"`
	FilePath     *string `json:"file_path,omitempty"`
}

// ReplyKeyboardMarkup is a struct for reply keyboard markups
//
// NOTE: Can be generated with NewReplyKeyboardMarkup() function in types_helper.go
//
// https://core.telegram.org/bots/api#replykeyboardmarkup
type ReplyKeyboardMarkup struct {
	Keyboard              [][]KeyboardButton `json:"keyboard"`
	IsPersistent          *bool              `json:"is_persistent,omitempty"`
	ResizeKeyboard        *bool              `json:"resize_keyboard,omitempty"`
	OneTimeKeyboard       *bool              `json:"one_time_keyboard,omitempty"`
	InputFieldPlaceholder *string            `json:"input_field_placeholder,omitempty"` // 1-64 characters
	Selective             *bool              `json:"selective,omitempty"`
}

// KeyboardButton is a struct of a keyboard button
//
// NOTE: Can be generated with NewKeyboardButton() function in types_helper.go
//
// https://core.telegram.org/bots/api#keyboardbutton
type KeyboardButton struct {
	Text              string                      `json:"text"`
	IconCustomEmojiID *string                     `json:"icon_custom_emoji_id,omitempty"`
	Style             *KeyboardStyle              `json:"style,omitempty"`
	RequestUsers      *KeyboardButtonRequestUsers `json:"request_users,omitempty"`
	RequestChat       *KeyboardButtonRequestChat  `json:"request_chat,omitempty"`
	RequestContact    *bool                       `json:"request_contact,omitempty"`
	RequestLocation   *bool                       `json:"request_location,omitempty"`
	RequestPoll       *KeyboardButtonPollType     `json:"request_poll,omitempty"`
	WebApp            *WebAppInfo                 `json:"web_app,omitempty"`
}

type KeyboardStyle string

const (
	KeyboardStyleDanger  KeyboardStyle = "danger"  // red
	KeyboardStyleSuccess KeyboardStyle = "success" // green
	KeyboardStylePrimary KeyboardStyle = "primary" // blue
)

// KeyboardButtonRequestUsers is a struct for `request_users` in KeyboardButton
//
// NOTE: Can be generated with NewKeyboardButtonRequestUsers() function in types_helper.go
//
// https://core.telegram.org/bots/api#keyboardbuttonrequestusers
type KeyboardButtonRequestUsers struct {
	RequestID       int64 `json:"request_id"`
	UserIsBot       *bool `json:"user_is_bot,omitempty"`
	UserIsPremium   *bool `json:"user_is_premium,omitempty"`
	MaxQuantity     *int  `json:"max_quantity,omitempty"`
	RequestName     *bool `json:"request_name,omitempty"`
	RequestUsername *bool `json:"request_username,omitempty"`
	RequestPhoto    *bool `json:"request_photo,omitempty"`
}

// KeyboardButtonRequestChat is a struct for `request_chat` in KeyboardButton
//
// NOTE: Can be generated with NewKeyboardButtonRequestChat() function in types_helper.go
//
// https://core.telegram.org/bots/api#keyboardbuttonrequestchat
type KeyboardButtonRequestChat struct {
	RequestID               int64                    `json:"request_id"`
	ChatIsChannel           bool                     `json:"chat_is_channel"`
	ChatIsForum             *bool                    `json:"chat_is_forum,omitempty"`
	ChatHasUsername         *bool                    `json:"chat_has_username,omitempty"`
	ChatIsCreated           *bool                    `json:"chat_is_created,omitempty"`
	UserAdministratorRights *ChatAdministratorRights `json:"user_administrator_rights,omitempty"`
	BotAdministratorRights  *ChatAdministratorRights `json:"bot_administrator_rights,omitempty"`
	BotIsMember             *bool                    `json:"bot_is_member,omitempty"`
	RequestTitle            *bool                    `json:"request_title,omitempty"`
	RequestUsername         *bool                    `json:"request_username,omitempty"`
	RequestPhoto            *bool                    `json:"request_photo,omitempty"`
}

// KeyboardButtonPollType is a struct for KeyboardButtonPollType
//
// NOTE: Can be generated with NewKeyboardButtonPollType() function in types_helper.go
//
// https://core.telegram.org/bots/api#keyboardbuttonpolltype
type KeyboardButtonPollType struct {
	Type *string `json:"type,omitempty"` // "quiz", "regular", or anything
}

// ReplyKeyboardRemove is a struct for ReplyKeyboardRemove
//
// NOTE: Can be generated with NewReplyKeyboardRemove() function in types_helper.go
//
// https://core.telegram.org/bots/api#replykeyboardremove
type ReplyKeyboardRemove struct {
	RemoveKeyboard bool  `json:"remove_keyboard"`
	Selective      *bool `json:"selective,omitempty"`
}

// InlineKeyboardMarkup is a struct for InlineKeyboardMarkup
//
// NOTE: Can be generated with NewInlineKeyboardMarkup() function in types_helper.go
//
// https://core.telegram.org/bots/api#inlinekeyboardmarkup
type InlineKeyboardMarkup struct {
	InlineKeyboard [][]InlineKeyboardButton `json:"inline_keyboard"`
}

// InlineKeyboardButton is a struct for InlineKeyboardButtons
//
// NOTE: Can be generated with NewInlineKeyboardButton() function in types_helper.go
//
// https://core.telegram.org/bots/api#inlinekeyboardbutton
type InlineKeyboardButton struct {
	Text                         string                       `json:"text"`
	IconCustomEmojiID            *string                      `json:"icon_custom_emoji_id,omitempty"`
	Style                        *KeyboardStyle               `json:"style,omitempty"`
	URL                          *string                      `json:"url,omitempty"`
	CallbackData                 *string                      `json:"callback_data,omitempty"`
	WebApp                       *WebAppInfo                  `json:"web_app,omitempty"`
	LoginURL                     *LoginURL                    `json:"login_url,omitempty"`
	SwitchInlineQuery            *string                      `json:"switch_inline_query,omitempty"`
	SwitchInlineQueryCurrentChat *string                      `json:"switch_inline_query_current_chat,omitempty"`
	SwitchInlineQueryChosenChat  *SwitchInlineQueryChosenChat `json:"switch_inline_query_chosen_chat,omitempty"`
	CopyText                     *CopyTextButton              `json:"copy_text,omitempty"`
	CallbackGame                 *CallbackGame                `json:"callback_game,omitempty"`
	Pay                          *bool                        `json:"pay,omitempty"`
}

// LoginURL is a struct for LoginURL
//
// NOTE: Can be generated with NewLoginURL() function in types_helper.go
//
// https://core.telegram.org/bots/api#loginurl
type LoginURL struct {
	URL                string  `json:"url"`
	ForwardText        *string `json:"forward_text,omitempty"`
	BotUsername        *string `json:"bot_username,omitempty"`
	RequestWriteAccess *bool   `json:"request_write_access,omitempty"`
}

// SwitchInlineQueryChosenChat is a struct for SwitchInlineQueryChosenChat
//
// NOTE: Can be generated with NewSwitchInlineQueryChosenChat() function in types_helper.go
//
// https://core.telegram.org/bots/api#switchinlinequerychosenchat
type SwitchInlineQueryChosenChat struct {
	Query             *string `json:"query,omitempty"`
	AllowUserChats    *bool   `json:"allow_user_chats,omitempty"`
	AllowBotChats     *bool   `json:"allow_bot_chats,omitempty"`
	AllowGroupChats   *bool   `json:"allow_group_chats,omitempty"`
	AllowChannelChats *bool   `json:"allow_channel_chats,omitempty"`
}

// CopyTextButton is a struct for CopyTextButton
//
// https://core.telegram.org/bots/api#copytextbutton
type CopyTextButton struct {
	Text string `json:"text"` // 1~256 characters
}

// CallbackQuery is a struct for a callback query
//
// https://core.telegram.org/bots/api#callbackquery
type CallbackQuery struct {
	ID              string                    `json:"id"`
	From            User                      `json:"from"`
	Message         *MaybeInaccessibleMessage `json:"message,omitempty"`
	InlineMessageID *string                   `json:"inline_message_id,omitempty"`
	ChatInstance    string                    `json:"chat_instance"`
	Data            *string                   `json:"data,omitempty"`
	GameShortName   *string                   `json:"game_short_name,omitempty"`
}

// RefundedPayment is a struct for a refunded payment
//
// https://core.telegram.org/bots/api#refundedpayment
type RefundedPayment struct {
	Currency                string  `json:"currency"`
	TotalAmount             int     `json:"total_amount"`
	InvoicePayload          string  `json:"invoice_payload"`
	TelegramPaymentChargeID string  `json:"telegram_payment_charge_id"`
	ProviderPaymentChargeID *string `json:"provider_payment_charge_id,omitempty"`
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

// PaidMediaPurchased is a struct for a paid media purchase.
//
// https://core.telegram.org/bots/api#paidmediapurchased
type PaidMediaPurchased struct {
	From             User   `json:"from"`
	PaidMediaPayload string `json:"paid_media_payload"`
}

// RevenueWithdrawalStateType is a type of revenue withdrawl state
//
// https://core.telegram.org/bots/api#revenuewithdrawalstate
type RevenueWithdrawalStateType string

const (
	RevenueWithdrawalStatePending   RevenueWithdrawalStateType = "pending"
	RevenueWithdrawalStateSucceeded RevenueWithdrawalStateType = "succeeded"
	RevenueWithdrawalStateFailed    RevenueWithdrawalStateType = "failed"
)

// RevenueWithdrawalState is a struct for a state of revenue withdrawl
//
// https://core.telegram.org/bots/api#revenuewithdrawalstate
type RevenueWithdrawalState struct {
	Type RevenueWithdrawalStateType `json:"type"`

	// when Type == RevenueWithdrawalStateSucceeded
	Date *int    `json:"date,omitempty"`
	URL  *string `json:"url,omitempty"`
}

// TransactionPartnerType is a type of transaction partner
//
// https://core.telegram.org/bots/api#transactionpartner
type TransactionPartnerType string

const (
	TransactionPartnerUser             TransactionPartnerType = "user"
	TransactionPartnerChat             TransactionPartnerType = "chat"
	TransactionPartnerAffiliateProgram TransactionPartnerType = "affiliate_program"
	TransactionPartnerFragment         TransactionPartnerType = "fragment"
	TransactionPartnerTelegramAds      TransactionPartnerType = "telegram_ads"
	TransactionPartnerTelegramAPI      TransactionPartnerType = "telegram_api"
	TransactionPartnerOther            TransactionPartnerType = "other"
)

// TransactionPartnerUserTransactionType is a transaction type of TransactionPartnerUser
//
// https://core.telegram.org/bots/api#transactionpartneruser
type TransactionPartnerUserTransactionType string

const (
	TransactionPartnerUserTransactionInvoicePayment          TransactionPartnerUserTransactionType = "invoice_payment"
	TransactionPartnerUserTransactionPaidMediaPayment        TransactionPartnerUserTransactionType = "paid_media_payment"
	TransactionPartnerUserTransactionGiftPurchase            TransactionPartnerUserTransactionType = "gift_purchase"
	TransactionPartnerUserTransactionPremiumPurchase         TransactionPartnerUserTransactionType = "premium_purchase"
	TransactionPartnerUserTransactionBusinessAccountTransfer TransactionPartnerUserTransactionType = "business_account_transfer"
)

// TransactionPartner is a struct for a transaction partner
//
// https://core.telegram.org/bots/api#transactionpartner
type TransactionPartner struct {
	Type TransactionPartnerType `json:"type"`

	// when Type == TransactionPartnerUser
	TransactionType             *TransactionPartnerUserTransactionType `json:"transaction_type,omitempty"`
	User                        *User                                  `json:"user,omitempty"`
	Affiliate                   *AffiliateInfo                         `json:"affiliate,omitempty"`
	InvoicePayload              *string                                `json:"invoice_payload,omitempty"`
	SubscriptionPeriod          *int                                   `json:"subscription_period,omitempty"`
	PaidMedia                   []PaidMedia                            `json:"paid_media,omitempty"`
	PaidMediaPayload            *string                                `json:"paid_media_payload,omitempty"`
	PremiumSubscriptionDuration *int                                   `json:"premium_subscription_duration,omitempty"`

	// when Type == TransactionPartnerChat
	Chat *Chat `json:"chat,omitempty"`

	// when Type == TransactionPartnerUser, or Type == TransactionPartnerChat
	Gift *string `json:"gift,omitempty"`

	// when Type == TransactionPartnerAffiliateProgram
	SponsorUser       *User `json:"sponsor_user,omitempty"`
	CommissionPerMile *int  `json:"commission_per_mille,omitempty"`

	// when Type == TransactionPartnerFragment
	WithdrawlState *RevenueWithdrawalState `json:"withdrawal_state,omitempty"`

	// when Type == TransactionParterTelegramAPI
	RequestCount *int `json:"request_count,omitempty"`
}

// AffiliateInfo is a struct for a affiliate of transaction
//
// https://core.telegram.org/bots/api#affiliateinfo
type AffiliateInfo struct {
	AffiliateUser     *User `json:"affiliate_user,omitempty"`
	AffiliateChat     *Chat `json:"affiliate_chat,omitempty"`
	CommissionPerMile int   `json:"commission_per_mille"`
	Amount            int   `json:"amount"`
	NanostarAmount    *int  `json:"nanostar_amount,omitempty"`
}

// StarTransaction is a struct for a star transaction
//
// https://core.telegram.org/bots/api#startransaction
type StarTransaction struct {
	ID             string              `json:"id"`
	Amount         int                 `json:"amount"`
	NanostarAmount *int                `json:"nanostar_amount,omitempty"`
	Date           int                 `json:"date"`
	Source         *TransactionPartner `json:"source,omitempty"`   // only for incoming transactions
	Receiver       *TransactionPartner `json:"receiver,omitempty"` // only for outgoing transactions
}

// StarTransactions is a struct for star transactions
//
// https://core.telegram.org/bots/api#startransactions
type StarTransactions struct {
	Transactions []StarTransaction `json:"transactions"`
}

// ForceReply is a struct for force-reply
//
// https://core.telegram.org/bots/api#forcereply
type ForceReply struct {
	ForceReply            bool    `json:"force_reply"`
	InputFieldPlaceholder *string `json:"input_field_placeholder,omitempty"` // 1-64 characters
	Selective             *bool   `json:"selective,omitempty"`
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

// ChatInviteLink is a struct of an invite link for a chat
//
// https://core.telegram.org/bots/api#chatinvitelink
type ChatInviteLink struct {
	InviteLink              string  `json:"invite_link"`
	Creator                 User    `json:"creator"`
	CreatesJoinRequest      bool    `json:"creates_join_request"`
	IsPrimary               bool    `json:"is_primary"`
	IsRevoked               bool    `json:"is_revoked"`
	Name                    *string `json:"name,omitempty"`
	ExpireDate              *int    `json:"expire_date,omitempty"`
	MemberLimit             *int    `json:"member_limit,omitempty"`
	PendingJoinRequestCount *int    `json:"pending_join_request_count,omitempty"`
}

// ChatAdministratorRights is a struct of chat administrator's rights
//
// NOTE: Can be generated with NewChatAdministratorRights() function in types_helper.go
//
// https://core.telegram.org/bots/api#chatadministratorrights
type ChatAdministratorRights struct {
	IsAnonymous             bool  `json:"is_anonymous"`
	CanManageChat           bool  `json:"can_manage_chat"`
	CanDeleteMessages       bool  `json:"can_delete_messages"`
	CanManageVideoChats     bool  `json:"can_manage_video_chats"`
	CanRestrictMembers      bool  `json:"can_restrict_members"`
	CanPromoteMembers       bool  `json:"can_promote_members"`
	CanChangeInfo           bool  `json:"can_change_info"`
	CanInviteUsers          bool  `json:"can_invite_users"`
	CanPostStories          bool  `json:"can_post_stories"`
	CanEditStories          bool  `json:"can_edit_stories"`
	CanDeleteStories        bool  `json:"can_delete_stories"`
	CanPostMessages         *bool `json:"can_post_messages,omitempty"`
	CanEditMessages         *bool `json:"can_edit_messages,omitempty"`
	CanPinMessages          *bool `json:"can_pin_messages,omitempty"`
	CanManageTopics         *bool `json:"can_manage_topics,omitempty"`
	CanManageDirectMessages *bool `json:"can_manage_direct_messages,omitempty"`
}

// ChatMember is a struct of a chat member
//
// https://core.telegram.org/bots/api#chatmember
type ChatMember struct {
	Status                ChatMemberStatus `json:"status"`
	User                  User             `json:"user"`
	IsAnonymous           *bool            `json:"is_anonymous,omitempty"`              // owner and administrators only
	CustomTitle           *string          `json:"custom_title,omitempty"`              // owner and administrators only
	CanBeEdited           *bool            `json:"can_be_edited,omitempty"`             // administrators only
	CanManageChat         *bool            `json:"can_manage_chat,omitempty"`           // administrators only
	CanPostMessages       *bool            `json:"can_post_messages,omitempty"`         // administrators only
	CanEditMessages       *bool            `json:"can_edit_messages,omitempty"`         // administrators only
	CanDeleteMessages     *bool            `json:"can_delete_messages,omitempty"`       // administrators only
	CanManageVideoChats   *bool            `json:"can_manage_video_chats,omitempty"`    // administrators only
	CanRestrictMembers    *bool            `json:"can_restrict_members,omitempty"`      // administrators only
	CanPromoteMembers     *bool            `json:"can_promote_members,omitempty"`       // administrators only
	CanChangeInfo         *bool            `json:"can_change_info,omitempty"`           // administrators and restricted only
	CanInviteUsers        *bool            `json:"can_invite_users,omitempty"`          // administrators and restricted only
	CanPinMessages        *bool            `json:"can_pin_messages,omitempty"`          // administrators and restricted only
	CanManageTopics       *bool            `json:"can_manage_topics,omitempty"`         // administrators and restricted only
	IsMember              *bool            `json:"is_member,omitempty"`                 // restricted only
	CanSendMessages       *bool            `json:"can_send_messages,omitempty"`         // restricted only
	CanSendMediaMessages  *bool            `json:"can_send_media_messages,omitempty"`   // restricted only
	CanSendPolls          *bool            `json:"can_send_polls,omitempty"`            // restricted only
	CanSendOtherMessages  *bool            `json:"can_send_other_messages,omitempty"`   // restricted only
	CanAddWebPagePreviews *bool            `json:"can_add_web_page_previews,omitempty"` // restricted only
	UntilDate             *int             `json:"until_date,omitempty"`                // restricted and kicked only
}

// ChatMemberUpdated is a struct of an updated chat member
//
// https://core.telegram.org/bots/api#chatmemberupdated
type ChatMemberUpdated struct {
	Chat                    Chat            `json:"chat"`
	From                    User            `json:"from"`
	Date                    int             `json:"date"`
	OldChatMember           ChatMember      `json:"old_chat_member"`
	NewChatMember           ChatMember      `json:"new_chat_member"`
	InviteLink              *ChatInviteLink `json:"invite_link,omitempty"`
	ViaJoinRequest          *bool           `json:"via_join_request,omitempty"`
	ViaChatFolderInviteLink *bool           `json:"via_chat_folder_invite_link,omitempty"`
}

// ChatPermissions is a struct of chat permissions
//
// NOTE: Can be generated with NewChatPermissions() function in types_helper.go
//
// https://core.telegram.org/bots/api#chatpermissions
type ChatPermissions struct {
	CanSendMessages       *bool `json:"can_send_messages,omitempty"`
	CanSendAudios         *bool `json:"can_send_audios,omitempty"`
	CanSendDocuments      *bool `json:"can_send_documents,omitempty"`
	CanSendPhotos         *bool `json:"can_send_photos,omitempty"`
	CanSendVideos         *bool `json:"can_send_videos,omitempty"`
	CanSendVideoNotes     *bool `json:"can_send_video_notes,omitempty"`
	CanSendVoiceNotes     *bool `json:"can_send_voice_notes,omitempty"`
	CanSendPolls          *bool `json:"can_send_polls,omitempty"`
	CanSendOtherMessages  *bool `json:"can_send_other_messages,omitempty"`
	CanAddWebPagePreviews *bool `json:"can_add_web_page_previews,omitempty"`
	CanChangeInfo         *bool `json:"can_change_info,omitempty"`
	CanInviteUsers        *bool `json:"can_invite_users,omitempty"`
	CanPinMessages        *bool `json:"can_pin_messages,omitempty"`
	CanManageTopics       *bool `json:"can_manage_topics,omitempty"`
}

// Birthdate is a struct of birth date
//
// https://core.telegram.org/bots/api#birthdate
type Birthdate struct {
	Day   int  `json:"day"`
	Month int  `json:"month"`
	Year  *int `json:"year,omitempty"`
}

// BusinessIntro is a struct of introduction of a business
//
// https://core.telegram.org/bots/api#businessintro
type BusinessIntro struct {
	Title   *string  `json:"title,omitempty"`
	Message *string  `json:"message,omitempty"`
	Sticker *Sticker `json:"sticker,omitempty"`
}

// BusinessLocation is a struct of a business location
//
// https://core.telegram.org/bots/api#businesslocation
type BusinessLocation struct {
	Address  string    `json:"address"`
	Location *Location `json:"location,omitempty"`
}

// BusinessOpeningHoursInterval is a struct of an opening hours interval of business
//
// https://core.telegram.org/bots/api#businessopeninghoursinterval
type BusinessOpeningHoursInterval struct {
	OpeningMinute int `json:"opening_minute"`
	ClosingMinute int `json:"closing_minute"`
}

// BusinessOpeningHours is a struct of an opening hours of business
//
// https://core.telegram.org/bots/api#businessopeninghours
type BusinessOpeningHours struct {
	TimeZoneName string                         `json:"time_zone_name"`
	OpeningHours []BusinessOpeningHoursInterval `json:"opening_hours"`
}

// UserRating is a struct which describes the rating of a user based on their Telegram Star spendings.
//
// https://core.telegram.org/bots/api#userrating
type UserRating struct {
	Level              int  `json:"level"`
	Rating             int  `json:"rating"`
	CurrentLevelRating int  `json:"current_level_rating"`
	NextLevelRating    *int `json:"next_level_rating,omitempty"`
}

// StoryAreaPosition describes the position of a clickable area within a story.
//
// https://core.telegram.org/bots/api#storyareaposition
type StoryAreaPosition struct {
	XPercentage            float32 `json:"x_percentage"`
	YPercentage            float32 `json:"y_percentage"`
	WidthPercentage        float32 `json:"width_percentage"`
	HeightPercentage       float32 `json:"height_percentage"`
	RotationAngle          float32 `json:"rotation_angle"`
	CornerRadiusPercentage float32 `json:"corner_radius_percentage"`
}

// LocationAddress describes the physical address of a location.
//
// https://core.telegram.org/bots/api#locationaddress
type LocationAddress struct {
	CountryCode string  `json:"country_code"`
	State       *string `json:"state,omitempty"`
	City        *string `json:"city,omitempty"`
	Street      *string `json:"street,omitempty"`
}

// StoryAreaTypeType is a type of StoryAreaType
type StoryAreaTypeType string

const (
	StoryAreaTypeLocation          StoryAreaTypeType = "location"
	StoryAreaTypeSuggestedReaction StoryAreaTypeType = "suggested_reaction"
	StoryAreaTypeLink              StoryAreaTypeType = "link"
	StoryAreaTypeWeather           StoryAreaTypeType = "weather"
	StoryAreaTypeUniqueGift        StoryAreaTypeType = "unique_gift"
)

// StoryAreaType describes the type of a clickable area on a story.
//
// https://core.telegram.org/bots/api#storyareatype
type StoryAreaType struct {
	Type StoryAreaTypeType `json:"type"`

	// Type == TypeStoryAreaLocation
	// https://core.telegram.org/bots/api#storyareatypelocation
	Latitude  *float32         `json:"latitude,omitempty"`
	Longitude *float32         `json:"longitude,omitempty"`
	Address   *LocationAddress `json:"address,omitempty"`

	// Type == TypeStoryAreaSuggestedReaction
	// https://core.telegram.org/bots/api#storyareatypesuggestedreaction
	ReactionType *ReactionType `json:"reaction_type,omitempty"`
	IsDark       *bool         `json:"is_dark,omitempty"`
	IsFlipped    *bool         `json:"is_flipped,omitempty"`

	// Type == TypeStoryAreaLink
	// https://core.telegram.org/bots/api#storyareatypelink
	URL *string `json:"url,omitempty"`

	// Type == TypeStoryAreaWeather
	// https://core.telegram.org/bots/api#storyareatypeweather
	Temperature     *float32 `json:"temperature,omitempty"`
	Emoji           *string  `json:"emoji,omitempty"`
	BackgroundColor *int     `json:"background_color,omitempty"`

	// Type == TypeStoryAreaUniqueGift
	// https://core.telegram.org/bots/api#storyareatypeuniquegift
	Name *string `json:"name,omitempty"`
}

// InputStoryContent describes the content of a story to post.
//
// https://core.telegram.org/bots/api#inputstorycontent
type InputStoryContent struct {
	Type string `json:"type"`

	// Type == "photo"
	// https://core.telegram.org/bots/api#inputstorycontentphoto
	Photo *string `json:"photo,omitempty"`

	// Type == "video"
	// https://core.telegram.org/bots/api#inputstorycontentvideo
	Video               *string  `json:"video,omitempty"`
	Duration            *float32 `json:"duration,omitempty"`
	CoverFrameTimestamp *float32 `json:"cover_frame_timestamp,omitempty"`
	IsAnimation         *bool    `json:"is_animation,omitempty"`
}

// StoryArea describes a clickable area on a story media.
//
// https://core.telegram.org/bots/api#storyarea
type StoryArea struct {
	Position StoryAreaPosition `json:"position"`
	Type     StoryAreaType     `json:"type"`
}

// ChatMemberOwner is a struct of a chat member who is an owner.
//
// https://core.telegram.org/bots/api#chatmemberowner
type ChatMemberOwner struct {
	Status      string  `json:"status"` // = "creator"
	User        User    `json:"user"`
	IsAnonymous bool    `json:"is_anonymous"`
	CustomTitle *string `json:"custom_title,omitempty"`
}

// ChatMemberAdministrator is a struct of a chat member who is an administrator.
//
// https://core.telegram.org/bots/api#chatmemberadministrator
type ChatMemberAdministrator struct {
	Status                  string  `json:"status"` // = "administrator"
	User                    User    `json:"user"`
	CanBeEdited             bool    `json:"can_be_edited"`
	IsAnonymous             bool    `json:"is_anonymous"`
	CanManageChat           bool    `json:"can_manage_chat"`
	CanDeleteMessages       bool    `json:"can_delete_messages"`
	CanManageVideoChats     bool    `json:"can_manage_video_chats"`
	CanRestrictMembers      bool    `json:"can_restrict_members"`
	CanPromoteMembers       bool    `json:"can_promote_members"`
	CanChangeInfo           bool    `json:"can_change_info"`
	CanInviteUsers          bool    `json:"can_invite_users"`
	CanPostStories          bool    `json:"can_post_stories"`
	CanEditStories          bool    `json:"can_edit_stories"`
	CanDeleteStories        bool    `json:"can_delete_stories"`
	CanPostMessages         *bool   `json:"can_post_messages,omitempty"`
	CanEditMessages         *bool   `json:"can_edit_messages,omitempty"`
	CanPinMessages          *bool   `json:"can_pin_messages,omitempty"`
	CanManageTopics         *bool   `json:"can_manage_topics,omitempty"`
	CanManageDirectMessages *bool   `json:"can_manage_direct_messages,omitempty"`
	CustomTitle             *string `json:"custom_title,omitempty"`
}

// ChatMemberMember is a struct of a chat member.
//
// https://core.telegram.org/bots/api#chatmembermember
type ChatMemberMember struct {
	Status    string `json:"status"` // = "member"
	User      User   `json:"user"`
	UntilDate int    `json:"until_date"`
}

// ChatMemberRestricted is a struct of chat member who is restricted
//
// https://core.telegram.org/bots/api#chatmemberrestricted
type ChatMemberRestricted struct {
	Status                 string `json:"status"` // = "restricted"
	User                   User   `json:"user"`
	IsMember               bool   `json:"is_member"`
	CanChangeInfo          bool   `json:"can_change_info"`
	CanInviteUsers         bool   `json:"can_invite_users"`
	CanPinMessages         bool   `json:"can_pin_messages"`
	CanManageTopics        bool   `json:"can_manage_topics"`
	CanSendMessages        bool   `json:"can_send_messages"`
	CanSendAudios          bool   `json:"can_send_audios"`
	CanSendDocuments       bool   `json:"can_send_documents"`
	CanSendPhotos          bool   `json:"can_send_photos"`
	CanSendVideos          bool   `json:"can_send_videos"`
	CanSendVideoNotes      bool   `json:"can_send_video_notes"`
	CanSendVoiceNotes      bool   `json:"can_send_voice_notes"`
	CanSendPolls           bool   `json:"can_send_polls"`
	CanSendOtherMessages   bool   `json:"can_send_other_messages"`
	CanSendWebPagePreviews bool   `json:"can_add_web_page_previews"`
	UntilDate              int    `json:"until_date"`
}

// ChatMemberLeft is a struct of a chat member who left.
//
// https://core.telegram.org/bots/api#chatmemberleft
type ChatMemberLeft struct {
	Status string `json:"status"` // = "left"
	User   User   `json:"user"`
}

// ChatMemberBanned is a struct of a chat member who is banned.
//
// https://core.telegram.org/bots/api#chatmemberbanned
type ChatMemberBanned struct {
	Status    string `json:"status"` // = "kicked"
	User      User   `json:"user"`
	UntilDate int    `json:"until_date"`
}

// ChatLocation is a struct of chat location
//
// https://core.telegram.org/bots/api#chatlocation
type ChatLocation struct {
	Location Location `json:"location"`
	Address  string   `json:"address"`
}

// ChatJoinRequest is a struct of chat join request
//
// https://core.telegram.org/bots/api#chatjoinrequest
type ChatJoinRequest struct {
	Chat       Chat            `json:"chat"`
	From       User            `json:"from"`
	UserChatID int64           `json:"user_chat_id"`
	Date       int             `json:"date"`
	Bio        *string         `json:"bio,omitempty"`
	InviteLink *ChatInviteLink `json:"invite_link,omitempty"`
}

// BotCommand is a struct of a bot command
//
// https://core.telegram.org/bots/api#botcommand
type BotCommand struct {
	Command     string `json:"command"`
	Description string `json:"description"`
}

// BotName is a struct of a bot's name
//
// https://core.telegram.org/bots/api#botname
type BotName struct {
	Name string `json:"name"`
}

// BotDescription is a struct of a bot's description
//
// https://core.telegram.org/bots/api#botdescription
type BotDescription struct {
	Description string `json:"description"`
}

// BotShortDescription is a struct of a bot's short description
//
// https://core.telegram.org/bots/api#botshortdescription
type BotShortDescription struct {
	ShortDescription string `json:"short_description"`
}

// BotCommandScopeType type
//
// https://core.telegram.org/bots/api#botcommandscope
type BotCommandScopeType string

// BotCommandScopeType constants
const (
	BotCommandScopeTypeDefault               BotCommandScopeType = "default"
	BotCommandScopeTypeAllPrivateChats       BotCommandScopeType = "all_private_chats"
	BotCommandScopeTypeAllGroupChats         BotCommandScopeType = "all_group_chats"
	BotCommandScopeTypeAllChatAdministrators BotCommandScopeType = "all_chat_administrators"
	BotCommandScopeTypeChat                  BotCommandScopeType = "chat"
	BotCommandScopeTypeChatAdministrators    BotCommandScopeType = "chat_administrators"
	BotCommandScopeTypeChatMember            BotCommandScopeType = "chat_member"
)

// BotCommandScopeDefault represents the bot command scopes
//
// https://core.telegram.org/bots/api#botcommandscopedefault
type BotCommandScopeDefault struct {
	Type BotCommandScopeType `json:"type"` // = "default"
}

// BotCommandScopeAllPrivateChats represents the bot command scopes
//
// https://core.telegram.org/bots/api#botcommandscopeallprivatechats
type BotCommandScopeAllPrivateChats BotCommandScopeDefault // = "all_private_chats"

// BotCommandScopeAllGroupChats represents the bot command scopes
//
// https://core.telegram.org/bots/api#botcommandscopeallgroupchats
type BotCommandScopeAllGroupChats BotCommandScopeDefault // = "all_group_chats"

// BotCommandScopeAllChatAdministrators represents the bot command scopes
//
// https://core.telegram.org/bots/api#botcommandscopeallchatadministrators
type BotCommandScopeAllChatAdministrators BotCommandScopeDefault // = "all_chat_administrators"

// BotCommandScopeChat represents the bot command scopes
//
// https://core.telegram.org/bots/api#botcommandscopechat
type BotCommandScopeChat struct {
	BotCommandScopeDefault        // = "chat"
	ChatID                 ChatID `json:"chat_id"`
}

// BotCommandScopeChatAdministrators represents the bot command scopes
//
// https://core.telegram.org/bots/api#botcommandscopechatadministrators
type BotCommandScopeChatAdministrators struct {
	BotCommandScopeDefault        // = "chat_administrators"
	ChatID                 ChatID `json:"chat_id"`
}

// BotCommandScopeChatMember represents the bot command scopes
//
// https://core.telegram.org/bots/api#botcommandscopechatmember
type BotCommandScopeChatMember struct {
	BotCommandScopeDefault        // = "chat_member"
	ChatID                 ChatID `json:"chat_id"`
	UserID                 int64  `json:"user_id"`
}

// Message is a struct of a message
//
// https://core.telegram.org/bots/api#message
type Message struct {
	MessageID                     int64                          `json:"message_id"`
	MessageThreadID               *int64                         `json:"message_thread_id,omitempty"`
	DirectMessagesTopic           *DirectMessagesTopic           `json:"direct_messages_topic,omitempty"`
	From                          *User                          `json:"from,omitempty"`
	SenderChat                    *Chat                          `json:"sender_chat,omitempty"`
	SenderBoostCount              *int                           `json:"sender_boost_count,omitempty"`
	SenderBusinessBot             *User                          `json:"sender_business_bot,omitempty"`
	Date                          int                            `json:"date"`
	BusinessConnectionID          *string                        `json:"business_connection_id,omitempty"`
	Chat                          Chat                           `json:"chat"`
	ForwardOrigin                 *MessageOrigin                 `json:"forward_origin,omitempty"`
	IsTopicMessage                *bool                          `json:"is_topic_message,omitempty"`
	IsAutomaticForward            *bool                          `json:"is_automatic_forward,omitempty"`
	ReplyToMessage                *Message                       `json:"reply_to_message,omitempty"`
	ExternalReply                 *ExternalReplyInfo             `json:"external_reply,omitempty"`
	Quote                         *TextQuote                     `json:"quote,omitempty"`
	ReplyToStory                  *Story                         `json:"reply_to_story,omitempty"`
	ReplyToChecklistTaskID        *int64                         `json:"reply_to_checklist_task_id,omitempty"`
	ViaBot                        *User                          `json:"via_bot,omitempty"`
	EditDate                      *int                           `json:"edit_date,omitempty"`
	HasProtectedContent           *bool                          `json:"has_protected_content,omitempty"`
	IsFromOffline                 *bool                          `json:"is_from_offline,omitempty"`
	IsPaidPost                    *bool                          `json:"is_paid_post,omitempty"`
	MediaGroupID                  *string                        `json:"media_group_id,omitempty"`
	AuthorSignature               *string                        `json:"author_signature,omitempty"`
	PaidStarCount                 *int                           `json:"paid_star_count,omitempty"`
	Text                          *string                        `json:"text,omitempty"`
	Entities                      []MessageEntity                `json:"entities,omitempty"`
	LinkPreviewOptions            *LinkPreviewOptions            `json:"link_preview_options,omitempty"`
	SuggestedPostInfo             *SuggestedPostInfo             `json:"suggested_post_info,omitempty"`
	EffectID                      *string                        `json:"effect_id,omitempty"`
	Animation                     *Animation                     `json:"animation,omitempty"`
	Audio                         *Audio                         `json:"audio,omitempty"`
	Document                      *Document                      `json:"document,omitempty"`
	PaidMedia                     *PaidMediaInfo                 `json:"paid_media,omitempty"`
	Photo                         []PhotoSize                    `json:"photo,omitempty"`
	Sticker                       *Sticker                       `json:"sticker,omitempty"`
	Story                         *Story                         `json:"story,omitempty"`
	Video                         *Video                         `json:"video,omitempty"`
	VideoNote                     *VideoNote                     `json:"video_note,omitempty"`
	Voice                         *Voice                         `json:"voice,omitempty"`
	Caption                       *string                        `json:"caption,omitempty"`
	CaptionEntities               []MessageEntity                `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia         *bool                          `json:"show_caption_above_media,omitempty"`
	HasMediaSpoiler               *bool                          `json:"has_media_spoiler,omitempty"`
	Checklist                     *Checklist                     `json:"checklist,omitempty"`
	Contact                       *Contact                       `json:"contact,omitempty"`
	Dice                          *Dice                          `json:"dice,omitempty"`
	Game                          *Game                          `json:"game,omitempty"`
	Poll                          *Poll                          `json:"poll,omitempty"`
	Venue                         *Venue                         `json:"venue,omitempty"`
	Location                      *Location                      `json:"location,omitempty"`
	NewChatMembers                []User                         `json:"new_chat_members,omitempty"`
	LeftChatMember                *User                          `json:"left_chat_member,omitempty"`
	ChatOwnerLeft                 *ChatOwnerLeft                 `json:"chat_owner_left,omitempty"`
	ChatOwnerChanged              *ChatOwnerChanged              `json:"chat_owner_changed,omitempty"`
	NewChatTitle                  *string                        `json:"new_chat_title,omitempty"`
	NewChatPhoto                  []PhotoSize                    `json:"new_chat_photo,omitempty"`
	DeleteChatPhoto               *bool                          `json:"delete_chat_photo,omitempty"`
	GroupChatCreated              *bool                          `json:"group_chat_created,omitempty"`
	SupergroupChatCreated         *bool                          `json:"supergroup_chat_created,omitempty"`
	ChannelChatCreated            *bool                          `json:"channel_chat_created,omitempty"`
	MessageAutoDeleteTimerChanged *MessageAutoDeleteTimerChanged `json:"message_auto_delete_timer_changed,omitempty"`
	MigrateToChatID               *int64                         `json:"migrate_to_chat_id,omitempty"`
	MigrateFromChatID             *int64                         `json:"migrate_from_chat_id,omitempty"`
	PinnedMessage                 *MaybeInaccessibleMessage      `json:"pinned_message,omitempty"`
	Invoice                       *Invoice                       `json:"invoice,omitempty"`
	SuccessfulPayment             *SuccessfulPayment             `json:"successful_payment,omitempty"`
	RefundedPayment               *RefundedPayment               `json:"refunded_payment,omitempty"`
	UsersShared                   *UsersShared                   `json:"users_shared,omitempty"`
	ChatShared                    *ChatShared                    `json:"chat_shared,omitempty"`
	Gift                          *GiftInfo                      `json:"gift,omitempty"`
	UniqueGift                    *UniqueGiftInfo                `json:"unique_gift,omitempty"`
	GiftUpgradeSent               *GiftInfo                      `json:"gift_upgrade_sent,omitempty"`
	ConnectedWebsite              *string                        `json:"connected_website,omitempty"`
	WriteAccessAllowed            *WriteAccessAllowed            `json:"write_access_allowed,omitempty"`
	// PassportData          *PassportData         `json:"passport_data,omitempty"` // NOT IMPLEMENTED: https://core.telegram.org/bots/api#passportdata
	ProximityAlertTriggered      *ProximityAlertTriggered      `json:"proximity_alert_triggered,omitempty"`
	BoostAdded                   *ChatBoostAdded               `json:"boost_added,omitempty"`
	ChatBackgroundSet            *ChatBackground               `json:"chat_background_set,omitempty"`
	ChecklistTasksDone           *ChecklistTasksDone           `json:"checklist_tasks_done,omitempty"`
	ChecklistTasksAdded          *ChecklistTasksAdded          `json:"checklist_tasks_added,omitempty"`
	DirectMessagePriceChanged    *DirectMessagePriceChanged    `json:"direct_message_price_changed,omitempty"`
	ForumTopicCreated            *ForumTopicCreated            `json:"forum_topic_created,omitempty"`
	ForumTopicEdited             *ForumTopicEdited             `json:"forum_topic_edited,omitempty"`
	ForumTopicClosed             *ForumTopicClosed             `json:"forum_topic_closed,omitempty"`
	ForumTopicReopened           *ForumTopicReopened           `json:"forum_topic_reopened,omitempty"`
	GeneralForumTopicHidden      *GeneralForumTopicHidden      `json:"general_forum_topic_hidden,omitempty"`
	GeneralForumTopicUnhidden    *GeneralForumTopicUnhidden    `json:"general_forum_topic_unhidden,omitempty"`
	GiveawayCreated              *GiveawayCreated              `json:"giveaway_created,omitempty"`
	Giveaway                     *Giveaway                     `json:"giveaway,omitempty"`
	GiveawayWinners              *GiveawayWinners              `json:"giveaway_winners,omitempty"`
	GiveawayCompleted            *GiveawayCompleted            `json:"giveaway_completed,omitempty"`
	PaidMessagePriceChanged      *PaidMessagePriceChanged      `json:"paid_message_price_changed,omitempty"`
	SuggestedPostApproved        *SuggestedPostApproved        `json:"suggested_post_approved,omitempty"`
	SuggestedPostApprovalFailed  *SuggestedPostApprovalFailed  `json:"suggested_post_approval_failed,omitempty"`
	SuggestedPostDeclined        *SuggestedPostDeclined        `json:"suggested_post_declined,omitempty"`
	SuggestedPostPaid            *SuggestedPostPaid            `json:"suggested_post_paid,omitempty"`
	SuggestedPostRefunded        *SuggestedPostRefunded        `json:"suggested_post_refunded,omitempty"`
	VideoChatScheduled           *VideoChatScheduled           `json:"video_chat_scheduled,omitempty"`
	VideoChatStarted             *VideoChatStarted             `json:"video_chat_started,omitempty"`
	VideoChatEnded               *VideoChatEnded               `json:"video_chat_ended,omitempty"`
	VideoChatParticipantsInvited *VideoChatParticipantsInvited `json:"video_chat_participants_invited,omitempty"`
	WebAppData                   *WebAppData                   `json:"web_app_data,omitempty"`
	ReplyMarkup                  *InlineKeyboardMarkup         `json:"reply_markup,omitempty"`
}

// MaybeInaccessibleMessage is a struct of a message that can be one of `Message` or `InaccessibleMessage`
//
// https://core.telegram.org/bots/api#maybeinaccessiblemessage
type MaybeInaccessibleMessage Message

// InaccessibleMessage is a struct of an inaccessible message
//
// https://core.telegram.org/bots/api#inaccessiblemessage
type InaccessibleMessage struct {
	Chat      Chat  `json:"chat"`
	MessageID int64 `json:"message_id"`
	Date      int   `json:"date"` // NOTE: always 0
}

// MessageID is a struct of message id
//
// https://core.telegram.org/bots/api#messageid
type MessageID struct {
	MessageID int64 `json:"message_id"`
}

// InlineQuery is a struct of an inline query
//
// https://core.telegram.org/bots/api#inlinequery
type InlineQuery struct {
	ID       string    `json:"id"`
	From     User      `json:"from"`
	Query    string    `json:"query"`
	Offset   string    `json:"offset"`
	ChatType *string   `json:"chat_type,omitempty"`
	Location *Location `json:"location,omitempty"`
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

// InlineQueryResultsButton is a struct for inline query results button
//
// https://core.telegram.org/bots/api#inlinequeryresultsbutton
type InlineQueryResultsButton struct {
	Text           string      `json:"text"`
	WebApp         *WebAppInfo `json:"web_app,omitempty"`
	StartParameter *string     `json:"start_parameter,omitempty"`
}

// InlineQueryResult is a struct for inline query results
//
// NOTE: Can be generated with NewInlineQueryResult*() functions in types_helper.go
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
	URL                 *string               `json:"url,omitempty"` // NOTE: pass an empty string for hiding the URL
	Description         *string               `json:"description,omitempty"`
	ThumbnailURL        *string               `json:"thumbnail_url,omitempty"`
	ThumbnailWidth      *int                  `json:"thumbnail_width,omitempty"`
	ThumbnailHeight     *int                  `json:"thumbnail_height,omitempty"`
}

// InlineQueryResultPhoto is a struct for InlineQueryResultPhoto
type InlineQueryResultPhoto struct { // https://core.telegram.org/bots/api#inlinequeryresultphoto
	InlineQueryResult
	PhotoURL              string                `json:"photo_url"`
	PhotoWidth            *int                  `json:"photo_width,omitempty"`
	PhotoHeight           *int                  `json:"photo_height,omitempty"`
	ThumbnailURL          string                `json:"thumbnail_url"`
	Title                 *string               `json:"title,omitempty"`
	Description           *string               `json:"description,omitempty"`
	Caption               *string               `json:"caption,omitempty"`
	ParseMode             *ParseMode            `json:"parse_mode,omitempty"`
	CaptionEntities       []MessageEntity       `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia *bool                 `json:"show_caption_above_media,omitempty"`
	ReplyMarkup           *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent   *InputMessageContent  `json:"input_message_content,omitempty"`
}

// InlineQueryResultGif is a struct for InlineQueryResultGif
type InlineQueryResultGif struct { // https://core.telegram.org/bots/api#inlinequeryresultgif
	InlineQueryResult
	GifURL                string                `json:"gif_url"`
	GifWidth              *int                  `json:"gif_width,omitempty"`
	GifHeight             *int                  `json:"gif_height,omitempty"`
	GifDuration           *int                  `json:"gif_duration,omitempty"`
	ThumbnailURL          string                `json:"thumbnail_url"`
	ThumbnailMimeType     *ThumbnailMimeType    `json:"thumbnail_mime_type,omitempty"`
	Title                 *string               `json:"title,omitempty"`
	Caption               *string               `json:"caption,omitempty"`
	ParseMode             *ParseMode            `json:"parse_mode,omitempty"`
	CaptionEntities       []MessageEntity       `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia *bool                 `json:"show_caption_above_media,omitempty"`
	ReplyMarkup           *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent   *InputMessageContent  `json:"input_message_content,omitempty"`
}

// InlineQueryResultMpeg4Gif is a struct for InlineQueryResultMpeg4Gif
type InlineQueryResultMpeg4Gif struct { // https://core.telegram.org/bots/api#inlinequeryresultmpeg4gif
	InlineQueryResult
	Mpeg4URL              string                `json:"mpeg4_url"`
	Mpeg4Width            *int                  `json:"mpeg4_width,omitempty"`
	Mpeg4Height           *int                  `json:"mpeg4_height,omitempty"`
	Mpeg4Duration         *int                  `json:"mpeg4_duration,omitempty"`
	ThumbnailURL          string                `json:"thumbnail_url"`
	ThumbnailMimeType     *ThumbnailMimeType    `json:"thumbnail_mime_type,omitempty"`
	Title                 *string               `json:"title,omitempty"`
	Caption               *string               `json:"caption,omitempty"`
	ParseMode             *ParseMode            `json:"parse_mode,omitempty"`
	CaptionEntities       []MessageEntity       `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia *bool                 `json:"show_caption_above_media,omitempty"`
	ReplyMarkup           *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent   *InputMessageContent  `json:"input_message_content,omitempty"`
}

// InlineQueryResultVideo is a struct of InlineQueryResultVideo
type InlineQueryResultVideo struct { // https://core.telegram.org/bots/api#inlinequeryresultvideo
	InlineQueryResult
	VideoURL              string                `json:"video_url"`
	MimeType              VideoMimeType         `json:"mime_type"`
	ThumbnailURL          string                `json:"thumbnail_url"`
	Title                 string                `json:"title"`
	Caption               *string               `json:"caption,omitempty"`
	ParseMode             *ParseMode            `json:"parse_mode,omitempty"`
	CaptionEntities       []MessageEntity       `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia *bool                 `json:"show_caption_above_media,omitempty"`
	VideoWidth            *int                  `json:"video_width,omitempty"`
	VideoHeight           *int                  `json:"video_height,omitempty"`
	VideoDuration         *int                  `json:"video_duration,omitempty"`
	Description           *string               `json:"description,omitempty"`
	ReplyMarkup           *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent   *InputMessageContent  `json:"input_message_content,omitempty"`
}

// InlineQueryResultAudio is a struct of InlineQueryResultAudio
type InlineQueryResultAudio struct { // https://core.telegram.org/bots/api#inlinequeryresultaudio
	InlineQueryResult
	AudioURL            string                `json:"audio_url"`
	Title               string                `json:"title"`
	Caption             *string               `json:"caption,omitempty"`
	ParseMode           *ParseMode            `json:"parse_mode,omitempty"`
	CaptionEntities     []MessageEntity       `json:"caption_entities,omitempty"`
	Performer           *string               `json:"performer,omitempty"`
	AudioDuration       *int                  `json:"audio_duration,omitempty"`
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
	CaptionEntities     []MessageEntity       `json:"caption_entities,omitempty"`
	VoiceDuration       *int                  `json:"voice_duration,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
}

// InlineQueryResultDocument is a struct of InlineQueryResultDocument
type InlineQueryResultDocument struct { // https://core.telegram.org/bots/api#inlinequeryresultdocument
	InlineQueryResult
	Title               string                `json:"title"`
	Caption             *string               `json:"caption,omitempty"`
	ParseMode           *ParseMode            `json:"parse_mode,omitempty"`
	CaptionEntities     []MessageEntity       `json:"caption_entities,omitempty"`
	DocumentURL         string                `json:"document_url"`
	MimeType            DocumentMimeType      `json:"mime_type"`
	Description         *string               `json:"description,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
	ThumbnailURL        *string               `json:"thumbnail_url,omitempty"`
	ThumbnailWidth      *int                  `json:"thumbnail_width,omitempty"`
	ThumbnailHeight     *int                  `json:"thumbnail_height,omitempty"`
}

// InlineQueryResultLocation is a struct of InlineQueryResultLocation
type InlineQueryResultLocation struct { // https://core.telegram.org/bots/api#inlinequeryresultlocation
	InlineQueryResult
	Latitude             float32               `json:"latitude"`
	Longitude            float32               `json:"longitude"`
	Title                string                `json:"title"`
	HorizontalAccuracy   *float32              `json:"horizontal_accuracy,omitempty"`
	LivePeriod           *int                  `json:"live_period,omitempty"`
	Heading              *int                  `json:"heading,omitempty"`
	ProximityAlertRadius *int                  `json:"proximity_alert_radius,omitempty"`
	ReplyMarkup          *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent  *InputMessageContent  `json:"input_message_content,omitempty"`
	ThumbnailURL         *string               `json:"thumbnail_url,omitempty"`
	ThumbnailWidth       *int                  `json:"thumbnail_width,omitempty"`
	ThumbnailHeight      *int                  `json:"thumbnail_height,omitempty"`
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
	GooglePlaceID       *string               `json:"google_place_id,omitempty"`
	GooglePlaceType     *string               `json:"google_place_type,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
	ThumbnailURL        *string               `json:"thumbnail_url,omitempty"`
	ThumbnailWidth      *int                  `json:"thumbnail_width,omitempty"`
	ThumbnailHeight     *int                  `json:"thumbnail_height,omitempty"`
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
	ThumbnailURL        *string               `json:"thumbnail_url,omitempty"`
	ThumbnailWidth      *int                  `json:"thumbnail_width,omitempty"`
	ThumbnailHeight     *int                  `json:"thumbnail_height,omitempty"`
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
	PhotoFileID           string                `json:"photo_file_id"`
	Title                 *string               `json:"title,omitempty"`
	Description           *string               `json:"description,omitempty"`
	Caption               *string               `json:"caption,omitempty"`
	ParseMode             *ParseMode            `json:"parse_mode,omitempty"`
	CaptionEntities       []MessageEntity       `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia *bool                 `json:"show_caption_above_media,omitempty"`
	ReplyMarkup           *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent   *InputMessageContent  `json:"input_message_content,omitempty"`
}

// InlineQueryResultCachedGif is a struct of InlineQueryResultCachedGif
type InlineQueryResultCachedGif struct { // https://core.telegram.org/bots/api#inlinequeryresultcachedgif
	InlineQueryResult
	GifFileID             string                `json:"gif_file_id"`
	Title                 *string               `json:"title,omitempty"`
	Caption               *string               `json:"caption,omitempty"`
	ParseMode             *ParseMode            `json:"parse_mode,omitempty"`
	CaptionEntities       []MessageEntity       `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia *bool                 `json:"show_caption_above_media,omitempty"`
	ReplyMarkup           *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent   *InputMessageContent  `json:"input_message_content,omitempty"`
}

// InlineQueryResultCachedMpeg4Gif is a struct of InlineQueryResultCachedMpeg4Gif
type InlineQueryResultCachedMpeg4Gif struct { // https://core.telegram.org/bots/api#inlinequeryresultcachedmpeg4gif
	InlineQueryResult
	Mpeg4FileID           string                `json:"mpeg4_file_id"`
	Title                 *string               `json:"title,omitempty"`
	Caption               *string               `json:"caption,omitempty"`
	ParseMode             *ParseMode            `json:"parse_mode,omitempty"`
	CaptionEntities       []MessageEntity       `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia *bool                 `json:"show_caption_above_media,omitempty"`
	ReplyMarkup           *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent   *InputMessageContent  `json:"input_message_content,omitempty"`
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
	CaptionEntities     []MessageEntity       `json:"caption_entities,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
}

// InlineQueryResultCachedVideo is a struct of InlineQueryResultCachedVideo
type InlineQueryResultCachedVideo struct { // https://core.telegram.org/bots/api#inlinequeryresultcachedvideo
	InlineQueryResult
	VideoFileID           string                `json:"video_file_id"`
	Title                 string                `json:"title"`
	Description           *string               `json:"description,omitempty"`
	Caption               *string               `json:"caption,omitempty"`
	ParseMode             *ParseMode            `json:"parse_mode,omitempty"`
	CaptionEntities       []MessageEntity       `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia *bool                 `json:"show_caption_above_media,omitempty"`
	ReplyMarkup           *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent   *InputMessageContent  `json:"input_message_content,omitempty"`
}

// InlineQueryResultCachedVoice is a struct of InlineQueryResultCachedVoice
type InlineQueryResultCachedVoice struct { // https://core.telegram.org/bots/api#inlinequeryresultcachedvoice
	InlineQueryResult
	VoiceFileID         string                `json:"voice_file_id"`
	Title               string                `json:"title"`
	Caption             *string               `json:"caption,omitempty"`
	ParseMode           *ParseMode            `json:"parse_mode,omitempty"`
	CaptionEntities     []MessageEntity       `json:"caption_entities,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
}

// InlineQueryResultCachedAudio is a struct of InlineQueryResultCachedAudio
type InlineQueryResultCachedAudio struct { // https://core.telegram.org/bots/api#inlinequeryresultcachedaudio
	InlineQueryResult
	AudioFileID         string                `json:"audio_file_id"`
	Caption             *string               `json:"caption,omitempty"`
	ParseMode           *ParseMode            `json:"parse_mode,omitempty"`
	CaptionEntities     []MessageEntity       `json:"caption_entities,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent  `json:"input_message_content,omitempty"`
}

// InputMessageContent is a generic type of input message content types
//
// (can be one of `InputTextMessageContent`, `InputLocationMessageContent`, `InputVenueMessageContent`, `InputContactMessageContent`, or `InputInvoiceMessageContent`)
//
// NOTE: Can be generated with NewInput*MessageContent() function in types_helper.go
//
// https://core.telegram.org/bots/api#inputmessagecontent
type InputMessageContent any

// InputTextMessageContent is a struct of InputTextMessageContent
type InputTextMessageContent struct { // https://core.telegram.org/bots/api#inputtextmessagecontent
	InputMessageContent

	MessageText        string              `json:"message_text"`
	ParseMode          *ParseMode          `json:"parse_mode,omitempty"`
	CaptionEntities    []MessageEntity     `json:"caption_entities,omitempty"`
	LinkPreviewOptions *LinkPreviewOptions `json:"link_preview_options,omitempty"`
}

// InputLocationMessageContent is a struct of InputLocationMessageContent
type InputLocationMessageContent struct { // https://core.telegram.org/bots/api#inputlocationmessagecontent
	InputMessageContent

	Latitude             float32  `json:"latitude"`
	Longitude            float32  `json:"longitude"`
	HorizontalAccuracy   *float32 `json:"horizontal_accuracy,omitempty"`
	LivePeriod           *int     `json:"live_period,omitempty"`
	Heading              *int     `json:"heading,omitempty"`
	ProximityAlertRadius *int     `json:"proximity_alert_radius,omitempty"`
}

// InputVenueMessageContent is a struct of InputVenueMessageContent
type InputVenueMessageContent struct { // https://core.telegram.org/bots/api#inputvenuemessagecontent
	InputMessageContent

	Latitude        float32 `json:"latitude"`
	Longitude       float32 `json:"longitude"`
	Title           string  `json:"title"`
	Address         string  `json:"address"`
	FoursquareID    *string `json:"foursquare_id,omitempty"`
	FoursquareType  *string `json:"foursquare_type,omitempty"`
	GooglePlaceID   *string `json:"google_place_id,omitempty"`
	GooglePlaceType *string `json:"google_place_type,omitempty"`
}

// InputContactMessageContent is a struct of InputContactMessageContent
type InputContactMessageContent struct { // https://core.telegram.org/bots/api#inputcontactmessagecontent
	InputMessageContent

	PhoneNumber string  `json:"phone_number"`
	FirstName   string  `json:"first_name"`
	LastName    *string `json:"last_name,omitempty"`
	VCard       *string `json:"vcard,omitempty"` // https://en.wikipedia.org/wiki/VCard
}

// InputInvoiceMessageContent is a struct of InputInvoiceMessageContent
//
// NOTE:
// - `ProviderToken`: Set "" for payments in Telegram Stars.
// - `Currency`: Set "XTR" for payments in Telegram Stars.
type InputInvoiceMessageContent struct { // https://core.telegram.org/bots/api#inputinvoicemessagecontent
	InputMessageContent

	Title                     string         `json:"title"`
	Description               string         `json:"description"`
	Payload                   string         `json:"payload"`
	ProviderToken             string         `json:"provider_token"`
	Currency                  string         `json:"currency"`
	Prices                    []LabeledPrice `json:"prices"`
	MaxTipAmount              *int           `json:"max_tip_amount,omitempty"`
	SuggestedTipAmounts       []int          `json:"suggested_tip_amounts,omitempty"`
	ProviderData              *string        `json:"provider_data,omitempty"`
	PhotoURL                  *string        `json:"photo_url,omitempty"`
	PhotoSize                 *int           `json:"photo_size,omitempty"`
	PhotoWidth                *int           `json:"photo_width,omitempty"`
	PhotoHeight               *int           `json:"photo_height,omitempty"`
	NeedName                  *bool          `json:"need_name,omitempty"`
	NeedPhoneNumber           *bool          `json:"need_phone_number,omitempty"`
	NeedEmail                 *bool          `json:"need_email,omitempty"`
	NeedShippingAddress       *bool          `json:"need_shipping_address,omitempty"`
	SendPhoneNumberToProvider *bool          `json:"send_phone_number_to_provider,omitempty"`
	SendEmailToProvider       *bool          `json:"send_email_to_provider,omitempty"`
	IsFlexible                *bool          `json:"is_flexible,omitempty"`
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
	Thumbnail    *PhotoSize `json:"thumbnail,omitempty"`
	FileName     *string    `json:"file_name,omitempty"`
	MimeType     *string    `json:"mime_type,omitempty"`
	FileSize     *int       `json:"file_size,omitempty"`
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
	Currency                   string     `json:"currency"`
	TotalAmount                int        `json:"total_amount"`
	InvoicePayload             string     `json:"invoice_payload"`
	SubscriptionExpirationDate *int       `json:"subscription_expiration_date,omitempty"`
	IsRecurring                *bool      `json:"is_recurring,omitempty"`
	IsFirstRecurring           *bool      `json:"is_first_recurring,omitempty"`
	ShippingOptionID           *string    `json:"shipping_option_id,omitempty"`
	OrderInfo                  *OrderInfo `json:"order_info,omitempty"`
	TelegramPaymentChargeID    string     `json:"telegram_payment_charge_id"`
	ProviderPaymentChargeID    string     `json:"provider_payment_charge_id"`
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

// WebAppInfo is a struct of web app's information
//
// https://core.telegram.org/bots/api#webappinfo
type WebAppInfo struct {
	URL string `json:"url"`
}

// SentWebAppMessage is a struct for an inline message sent by web app
//
// https://core.telegram.org/bots/api#sentwebappmessage
type SentWebAppMessage struct {
	InlineMessageID *string `json:"inline_message_id,omitempty"`
}

// PreparedInlineMessage is a struct for a prepared inline message
//
// https://core.telegram.org/bots/api#preparedinlinemessage
type PreparedInlineMessage struct {
	ID             string `json:"id"`
	ExpirationDate int    `json:"expiration_date"`
}

// MenuButton is a generic type of the bot's menu buttons
//
// https://core.telegram.org/bots/api#menubutton
type MenuButton any

// MenuButtonCommands is a struct for a menu button which opens the bot's commands list
//
// https://core.telegram.org/bots/api#menubuttoncommands
type MenuButtonCommands struct {
	Type string `json:"type"` // = "commands"
}

// MenuButtonWebApp is a struct for a menu button which launches a web app
//
// https://core.telegram.org/bots/api#menubuttonwebapp
type MenuButtonWebApp struct {
	Type   string     `json:"type"` // = "web_app"
	Text   string     `json:"text"`
	WebApp WebAppInfo `json:"web_app"`
}

// MenuButtonDefault is a struct for a menu button with no specific value
//
// https://core.telegram.org/bots/api#menubuttondefault
type MenuButtonDefault struct {
	Type string `json:"type"` // = "default"
}

// ForumTopic is a struct for a forum topic.
//
// https://core.telegram.org/bots/api#forumtopic
type ForumTopic struct {
	MessageThreadID   int64   `json:"message_thread_id"`
	Name              string  `json:"name"`
	IconColor         int     `json:"icon_color"`
	IconCustomEmojiID *string `json:"icon_custom_emoji_id,omitempty"`
	IsNameImplicit    *bool   `json:"is_name_implicit,omitempty"`
}

// ReactionType is a struct for a reaction
//
// NOTE: Can be generated with New*Reaction*() functions in types_helper.go
//
// https://core.telegram.org/bots/api#reactiontype
type ReactionType struct {
	Type          string  `json:"type"`
	Emoji         *string `json:"emoji,omitempty"`
	CustomEmojiID *string `json:"custom_emoji_id,omitempty"`
}

// ReactionCount is a struct for a count of reactions
//
// https://core.telegram.org/bots/api#reactioncount
type ReactionCount struct {
	Type       ReactionType `json:"type"`
	TotalCount int          `json:"total_count"`
}

// MessageReactionUpdated is a struct for a change of a reaction on a message
//
// https://core.telegram.org/bots/api#messagereactionupdated
type MessageReactionUpdated struct {
	Chat        Chat           `json:"chat"`
	MessageID   int64          `json:"message_id"`
	User        *User          `json:"user,omitempty"`
	ActorChat   *Chat          `json:"actor_chat,omitempty"`
	Date        int            `json:"date"`
	OldReaction []ReactionType `json:"old_reaction"`
	NewReaction []ReactionType `json:"new_reaction"`
}

// MessageReactionCountUpdated is a struct for changes with anonymous reactions on a message
//
// https://core.telegram.org/bots/api#messagereactioncountupdated
type MessageReactionCountUpdated struct {
	Chat      Chat            `json:"chat"`
	MessageID int64           `json:"message_id"`
	Date      int             `json:"date"`
	Reactions []ReactionCount `json:"reactions"`
}

// ChatBoostSource is a struct for sources of chat boosts
//
// https://core.telegram.org/bots/api#chatboostsource
type ChatBoostSource struct {
	Source string `json:"source"`

	User *User `json:"user,omitempty"`

	// source == "giveaway"
	GiveawayMessageID *int64 `json:"giveaway_message_id,omitempty"`
	PrizeStarCount    *int   `json:"prize_star_count,omitempty"`
	IsUnclaimed       *bool  `json:"is_unclaimed,omitempty"`
}

// ChatBoost is a struct for a chat boost
//
// https://core.telegram.org/bots/api#chatboost
type ChatBoost struct {
	BoostID        string          `json:"boost_id"`
	AddDate        int             `json:"add_date"`
	ExpirationDate int             `json:"expiration_date"`
	Source         ChatBoostSource `json:"source"`
}

// ChatBoostUpdated is a struct for an added or changed boost
//
// https://core.telegram.org/bots/api#chatboostupdated
type ChatBoostUpdated struct {
	Chat  Chat      `json:"chat"`
	Boost ChatBoost `json:"boost"`
}

// ChatBoostRemoved is a struct for a removed boost
//
// https://core.telegram.org/bots/api#chatboostremoved
type ChatBoostRemoved struct {
	Chat       Chat            `json:"chat"`
	BoostID    string          `json:"boost_id"`
	RemoveDate int             `json:"remove_date"`
	Source     ChatBoostSource `json:"source"`
}

// ChatOwnerLeft is a struct for a left chat owner.
//
// https://core.telegram.org/bots/api#chatownerleft
type ChatOwnerLeft struct {
	NewOwner *User `json:"new_owner,omitempty"`
}

// ChatOwnerChanged is a struct for a changed chat owner.
//
// https://core.telegram.org/bots/api#chatownerchanged
type ChatOwnerChanged struct {
	NewOwner User `json:"new_owner"`
}

// UserChatBoosts is a struct for a user's chat boosts
//
// https://core.telegram.org/bots/api#userchatboosts
type UserChatBoosts struct {
	Boosts []ChatBoost `json:"boosts"`
}

// BusinessConnection is a struct for a business connection
//
// https://core.telegram.org/bots/api#businessconnection
type BusinessConnection struct {
	ID         string             `json:"id"`
	User       User               `json:"user"`
	UserChatID int64              `json:"user_chat_id"`
	Date       int                `json:"date"`
	CanReply   bool               `json:"can_reply"`
	Rights     *BusinessBotRights `json:"rights,omitempty"`
	IsEnabled  bool               `json:"is_enabled"`
}

// BusinessBotRights is a struct for business bot rights
//
// https://core.telegram.org/bots/api#businessbotrights
type BusinessBotRights struct {
	CanReply                   *bool `json:"can_reply,omitempty"`
	CanReadMessages            *bool `json:"can_read_messages,omitempty"`
	CanDeleteOutgoingMessages  *bool `json:"can_delete_outgoing_messages,omitempty"`
	CanDeleteAllMessages       *bool `json:"can_delete_all_messages,omitempty"`
	CanEditName                *bool `json:"can_edit_name,omitempty"`
	CanEditBio                 *bool `json:"can_edit_bio,omitempty"`
	CanEditProfilePhoto        *bool `json:"can_edit_profile_photo,omitempty"`
	CanEditUsername            *bool `json:"can_edit_username,omitempty"`
	CanChangeGiftSettings      *bool `json:"can_change_gift_settings,omitempty"`
	CanViewGiftsAndStars       *bool `json:"can_view_gifts_and_stars,omitempty"`
	CanConvertGiftsToStars     *bool `json:"can_convert_gifts_to_stars,omitempty"`
	CanTransferAndUpgradeGifts *bool `json:"can_transfer_and_upgrade_gifts,omitempty"`
	CanTransferStars           *bool `json:"can_transfer_stars,omitempty"`
	CanManageStories           *bool `json:"can_manage_stories,omitempty"`
}

// BusinessMessagesDeleted is a struct sent when messages are deleted from connected business accounts
//
// https://core.telegram.org/bots/api#businessmessagesdeleted
type BusinessMessagesDeleted struct {
	BusinessConnectionID string  `json:"business_connection_id"`
	Chat                 Chat    `json:"chat"`
	MessageIDs           []int64 `json:"message_ids"`
}

// GiftBackground is a struct for a gift background.
//
// https://core.telegram.org/bots/api#giftbackground
type GiftBackground struct {
	CenterColor int `json:"center_color"`
	EdgeColor   int `json:"edge_color"`
	TextColor   int `json:"text_color"`
}

// Gift represents a gift that can be sent by the bot.
//
// https://core.telegram.org/bots/api#gift
type Gift struct {
	ID                     string          `json:"id"`
	Sticker                Sticker         `json:"sticker"`
	StarCount              int             `json:"star_count"`
	UpgradeStarCount       *int            `json:"upgrade_star_count,omitempty"`
	IsPremium              *bool           `json:"is_premium,omitempty"`
	HasColors              *bool           `json:"has_colors,omitempty"`
	TotalCount             *int            `json:"total_count,omitempty"`
	RemainingCount         *int            `json:"remaining_count,omitempty"`
	PersonalTotalCount     *int            `json:"personal_total_count,omitempty"`
	PersonalRemainingCount *int            `json:"personal_remaining_count,omitempty"`
	Background             *GiftBackground `json:"background,omitempty"`
	UniqueGiftVariantCount *int            `json:"unique_gift_variant_count,omitempty"`
	PublisherChat          *Chat           `json:"publisher_chat,omitempty"`
}

// Gifts represents a list of gifts.
//
// https://core.telegram.org/bots/api#gifts
type Gifts struct {
	Gifts []Gift `json:"gifts"`
}

type Rarity string

const (
	RarityUncommon  Rarity = "uncommon"
	RarityRare      Rarity = "rare"
	RarityEpic      Rarity = "epic"
	RarityLegendary Rarity = "legendary"
)

// UniqueGiftModel describes the model of a unique gift.
//
// https://core.telegram.org/bots/api#uniquegiftmodel
type UniqueGiftModel struct {
	Name           string  `json:"name"`
	Sticker        Sticker `json:"sticker"`
	RarityPerMille int     `json:"rarity_per_mille"`
	Rarity         *Rarity `json:"rarity,omitempty"`
}

// UniqueGiftSymbol describes the symbol shown on the pattern of a unique gift.
//
// https://core.telegram.org/bots/api#uniquegiftsymbol
type UniqueGiftSymbol struct {
	Name           string  `json:"name"`
	Sticker        Sticker `json:"sticker"`
	RarityPerMille int     `json:"rarity_per_mille"`
}

// UniqueGiftBackdropColors describes the colors of the backdrop of a unique gift.
//
// https://core.telegram.org/bots/api#uniquegiftbackdropcolors
type UniqueGiftBackdropColors struct {
	CenterColor int `json:"center_color"`
	EdgeColor   int `json:"edge_color"`
	SymbolColor int `json:"symbol_color"`
	TextColor   int `json:"text_color"`
}

// UniqueGiftBackdrop describes the backdrop of a unique gift.
//
// https://core.telegram.org/bots/api#uniquegiftbackdrop
type UniqueGiftBackdrop struct {
	Name           string                   `json:"name"`
	Colors         UniqueGiftBackdropColors `json:"colors"`
	RarityPerMille int                      `json:"rarity_per_mille"`
}

// UniqueGiftColors contains information about the color scheme for a user's name, message replies and link previews based on a unique gift.
//
// https://core.telegram.org/bots/api#uniquegiftcolors
type UniqueGiftColors struct {
	ModelCustomEmojiID    string `json:"model_custom_emoji_id"`
	SymbolCustomEmojiID   string `json:"symbol_custom_emoji_id"`
	LightThemeMainColor   int    `json:"light_theme_main_color"`
	LightThemeOtherColors []int  `json:"light_theme_other_colors"`
	DarkThemeMainColor    int    `json:"dark_theme_main_color"`
	DarkThemeOtherColors  []int  `json:"dark_theme_other_colors"`
}

// UniqueGift describes a unique gift that was upgraded from a regular gift.
//
// https://core.telegram.org/bots/api#uniquegift
type UniqueGift struct {
	GiftID           string             `json:"gift_id"`
	BaseName         string             `json:"base_name"`
	Name             string             `json:"name"`
	Number           int                `json:"number"`
	Model            UniqueGiftModel    `json:"model"`
	Symbol           UniqueGiftSymbol   `json:"symbol"`
	Backdrop         UniqueGiftBackdrop `json:"backdrop"`
	IsPremium        *bool              `json:"is_premium,omitempty"`
	IsBurned         *bool              `json:"is_burned,omitempty"`
	IsFromBlockchain *bool              `json:"is_from_blockchain,omitempty"`
	Colors           *UniqueGiftColors  `json:"colors,omitempty"`
	PublisherChat    *Chat              `json:"publisher_chat,omitempty"`
}

// GiftInfo describes a service message about a regular gift that was sent or received.
//
// https://core.telegram.org/bots/api#giftinfo
type GiftInfo struct {
	Gift                    Gift            `json:"gift"`
	OwnedGiftID             *string         `json:"owned_gift_id,omitempty"`
	ConvertStarCount        *int            `json:"convert_star_count,omitempty"`
	PrepaidUpgradeStarCount *int            `json:"prepaid_upgrade_star_count,omitempty"`
	IsUpgradeSeparate       *bool           `json:"is_upgrade_separate,omitempty"`
	CanBeUpgraded           *bool           `json:"can_be_upgraded,omitempty"`
	Text                    *string         `json:"text,omitempty"`
	Entities                []MessageEntity `json:"entities,omitempty"`
	IsPrivate               *bool           `json:"is_private,omitempty"`
	UniqueGiftNumber        *int            `json:"unique_gift_number,omitempty"`
}

// UniqueGiftInfo describes a service message about a unique gift that was sent or received.
//
// https://core.telegram.org/bots/api#uniquegiftinfo
type UniqueGiftInfo struct {
	Gift               UniqueGift           `json:"gift"`
	Origin             UniqueGiftInfoOrigin `json:"origin"`
	LastResaleCurrency *string              `json:"last_resale_currency,omitempty"`
	LastResaleAmount   *int                 `json:"last_resale_amount,omitempty"`
	OwnedGiftID        *string              `json:"owned_gift_id,omitempty"`
	TransferStarCount  *int                 `json:"transfer_star_count,omitempty"`
	NextTransferDate   *int                 `json:"next_transfer_date,omitempty"`
}

// UniqueGiftInfoOrigin is the origin of a unique gift info.
type UniqueGiftInfoOrigin string

// UniqueGiftInfoOrigin constants
const (
	UniqueGiftInfoOriginUpgrade       UniqueGiftInfoOrigin = "upgrade"
	UniqueGiftInfoOriginTransfer      UniqueGiftInfoOrigin = "transfer"
	UniqueGiftInfoOriginResale        UniqueGiftInfoOrigin = "resale"
	UniqueGiftInfoOriginGiftedUpgrade UniqueGiftInfoOrigin = "gifted_upgrade"
	UniqueGiftInfoOriginOffer         UniqueGiftInfoOrigin = "offer"
)

type InputProfilePhotoType string

const (
	InputProfilePhotoStatic   InputProfilePhotoType = "static"
	InputProfilePhotoAnimated InputProfilePhotoType = "animated"
)

// InputProfilePhoto describes a profile photo to set.
//
// https://core.telegram.org/bots/api#inputprofilephoto
type InputProfilePhoto struct {
	Type InputProfilePhotoType `json:"type"`

	// when Type == InputProfilePhotoStatic
	//
	// https://core.telegram.org/bots/api#inputprofilephotostatic
	Photo *string `json:"photo,omitempty"`

	// when Type == InputProfilePhotoAnimated
	//
	// https://core.telegram.org/bots/api#inputprofilephotoanimated
	Animation          *string  `json:"animation,omitempty"`
	MainFrameTimestamp *float32 `json:"main_frame_timestamp,omitempty"`

	// actual data for file upload
	Filepath *string `json:"-"`
	Bytes    []byte  `json:"-"`
}

// AcceptedGiftTypes describes the types of gifts that can be gifted to a user or a chat.
//
// https://core.telegram.org/bots/api#acceptedgifttypes
type AcceptedGiftTypes struct {
	UnlimitedGifts      bool `json:"unlimited_gifts"`
	LimitedGifts        bool `json:"limited_gifts"`
	UniqueGifts         bool `json:"unique_gifts"`
	PremiumSubscription bool `json:"premium_subscription"`
	GiftsFromChannels   bool `json:"gifts_from_channels"`
}

// StarAmount describes an amount of Telegram Stars.
//
// https://core.telegram.org/bots/api#staramount
type StarAmount struct {
	Amount         int  `json:"amount"`
	NanostarAmount *int `json:"nanostar_amount,omitempty"`
}

// OwnedGifts contains the list of gifts received and owned by a user or a chat.
//
// https://core.telegram.org/bots/api#ownedgifts
type OwnedGifts struct {
	TotalCount int         `json:"total_count"`
	Gifts      []OwnedGift `json:"gifts"`
	NextOffset *string     `json:"next_offset,omitempty"`
}

// OwnedGift describes a gift received and owned by a user or a chat.
//
// https://core.telegram.org/bots/api#ownedgift
type OwnedGift struct {
	Type        string          `json:"type"`
	Gift        json.RawMessage `json:"gift"` // one of `Gift` or `UniqueGift`
	OwnedGiftID *string         `json:"owned_gift_id,omitempty"`
	SenderUser  *User           `json:"sender_user,omitempty"`
	SendDate    int             `json:"send_date"`
	IsSaved     *bool           `json:"is_saved,omitempty"`

	// Type == "regular"
	Text                    *string         `json:"text,omitempty"`
	Entities                []MessageEntity `json:"entities,omitempty"`
	IsPrivate               *bool           `json:"is_private,omitempty"`
	CanBeUpgraded           *bool           `json:"can_be_upgraded,omitempty"`
	WasRefunded             *bool           `json:"was_refunded,omitempty"`
	ConvertStarCount        *int            `json:"convert_star_count,omitempty"`
	PrepaidUpgradeStarCount *int            `json:"prepaid_upgrade_star_count,omitempty"`
	IsUpgradeSeparate       *bool           `json:"is_upgrade_separate,omitempty"`
	UniqueGiftNumber        *int            `json:"unique_gift_number,omitempty"`

	// Type == "unique"
	CanBeTransferred  *bool `json:"can_be_transferred,omitempty"`
	TransferStarCount *int  `json:"transfer_star_count,omitempty"`
	NextTransferDate  *int  `json:"next_transfer_date,omitempty"`
}
