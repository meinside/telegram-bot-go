package telegrambot

// https://core.telegram.org/bots/api#available-methods

// MethodOptions is a type for methods' options parameter.
type MethodOptions map[string]any

// OptionsSetWebhook struct for SetWebhook().
//
// options include: `certificate`, `ip_address`, `max_connections`, `allowed_updates`, `drop_pending_updates`, and `secret_token`.
//
// https://core.telegram.org/bots/api#setwebhook
type OptionsSetWebhook MethodOptions

// SetCertificate sets the `certificate` value of OptionsSetWebhook.
func (o OptionsSetWebhook) SetCertificate(filepath string) OptionsSetWebhook {
	o["certificate"] = filepath
	return o
}

// SetIPAddress sets the `ip_address` value of OptionsSetWebhook.
func (o OptionsSetWebhook) SetIPAddress(address string) OptionsSetWebhook {
	o["ip_address"] = address
	return o
}

// SetMaxConnections sets the `max_connections` value of OptionsSetWebhook.
//
// maxConnections: 1 ~ 100 (default: 40)
func (o OptionsSetWebhook) SetMaxConnections(maxConnections int) OptionsSetWebhook {
	o["max_connections"] = maxConnections
	return o
}

// SetAllowedUpdates sets the `allowed_updates` value of OptionsSetWebhook.
func (o OptionsSetWebhook) SetAllowedUpdates(allowedUpdates []UpdateType) OptionsSetWebhook {
	o["allowed_updates"] = allowedUpdates
	return o
}

// SetDropPendingUpdates sets the `drop_pending_updates` value of OptionsSetWebhook.
func (o OptionsSetWebhook) SetDropPendingUpdates(drop bool) OptionsSetWebhook {
	o["drop_pending_updates"] = drop
	return o
}

// SetSecretToken sets the `secret_token` value of OptionsSetWebhook.
func (o OptionsSetWebhook) SetSecretToken(token string) OptionsSetWebhook {
	o["secret_token"] = token
	return o
}

// OptionsGetUpdates struct for GetUpdates().
//
// options include: `offset`, `limit`, `timeout`, and `allowed_updates`.
//
// https://core.telegram.org/bots/api#getupdates
type OptionsGetUpdates MethodOptions

// SetOffset sets the `offset` value of OptionsGetUpdates.
func (o OptionsGetUpdates) SetOffset(offset int64) OptionsGetUpdates {
	o["offset"] = offset
	return o
}

// SetLimit sets the `limit` value of OptionsGetUpdates.
func (o OptionsGetUpdates) SetLimit(limit int) OptionsGetUpdates {
	o["limit"] = limit
	return o
}

// SetTimeout sets the `timeout` value of OptionsGetUpdates.
func (o OptionsGetUpdates) SetTimeout(timeout int) OptionsGetUpdates {
	o["timeout"] = timeout
	return o
}

// SetAllowedUpdates sets the `allowed_updates` value of OptionsGetUpdates.
func (o OptionsGetUpdates) SetAllowedUpdates(allowedUpdates []AllowedUpdate) OptionsGetUpdates {
	o["allowed_updates"] = allowedUpdates
	return o
}

// OptionsSendMessage struct for SendMessage().
//
// options include: `business_connection_id`, `message_thread_id`, `direct_messages_topic_id`, `parse_mode`, `entities`, `link_preview_options`, `disable_notification`, `protect_content`, `allow_paid_broadcast`, `message_effect_id`, `suggested_post_parameters`, `reply_parameters`, and `reply_markup`.
//
// https://core.telegram.org/bots/api#sendmessage
type OptionsSendMessage MethodOptions

// SetBusinessConnectionID sets the `business_connection_id` value of OptionsSendMessage.
func (o OptionsSendMessage) SetBusinessConnectionID(businessConnectionID string) OptionsSendMessage {
	o["business_connection_id"] = businessConnectionID
	return o
}

// SetMessageThreadID sets the `message_thread_id` value of OptionsSendMessage.
func (o OptionsSendMessage) SetMessageThreadID(messageThreadID int64) OptionsSendMessage {
	o["message_thread_id"] = messageThreadID
	return o
}

// SetDirectMessagesTopicID sets the `direct_messages_topic_id` value of OptionsSendMessage.
func (o OptionsSendMessage) SetDirectMessagesTopicID(directMessagesTopicID int64) OptionsSendMessage {
	o["direct_messages_topic_id"] = directMessagesTopicID
	return o
}

// SetParseMode sets the `parse_mode` value of OptionsSendMessage.
func (o OptionsSendMessage) SetParseMode(parseMode ParseMode) OptionsSendMessage {
	o["parse_mode"] = parseMode
	return o
}

// SetEntities sets the `entities` value of OptionsSendMessage.
func (o OptionsSendMessage) SetEntities(entities []MessageEntity) OptionsSendMessage {
	o["entities"] = entities
	return o
}

// SetLinkPreviewOptions sets the `link_preview_options` value of OptionsSendMessage.
func (o OptionsSendMessage) SetLinkPreviewOptions(linkPreviewOptions LinkPreviewOptions) OptionsSendMessage {
	o["link_preview_options"] = linkPreviewOptions
	return o
}

// SetDisableNotification sets the `disable_notification` value of OptionsSendMessage.
func (o OptionsSendMessage) SetDisableNotification(disable bool) OptionsSendMessage {
	o["disable_notification"] = disable
	return o
}

// SetProtectContent sets the `protect_content` value of OptionsSendMessage.
func (o OptionsSendMessage) SetProtectContent(protect bool) OptionsSendMessage {
	o["protect_content"] = protect
	return o
}

// SetAllowPaidBroadcast sets the `allow_paid_broadcast` value of OptionsSendMessage.
func (o OptionsSendMessage) SetAllowPaidBroadcast(allow bool) OptionsSendMessage {
	o["allow_paid_broadcast"] = allow
	return o
}

// SetMessageEffectID sets the `message_effect_id` value of OptionsSendMessage.
func (o OptionsSendMessage) SetMessageEffectID(messageEffectID string) OptionsSendMessage {
	o["message_effect_id"] = messageEffectID
	return o
}

// SetSuggestedPostParameters sets the `suggested_post_parameters` value of OptionsSendMessage.
func (o OptionsSendMessage) SetSuggestedPostParameters(suggestedPostParameters SuggestedPostParameters) OptionsSendMessage {
	o["suggested_post_parameters"] = suggestedPostParameters
	return o
}

// SetReplyParameters sets the `reply_parameters` value of OptionsSendMessage.
func (o OptionsSendMessage) SetReplyParameters(replyParameters ReplyParameters) OptionsSendMessage {
	o["reply_parameters"] = replyParameters
	return o
}

// SetReplyMarkup sets the `reply_markup` value of OptionsSendMessage.
//
// `replyMarkup` can be one of InlineKeyboardMarkup, ReplyKeyboardMarkup, ReplyKeyboardRemove, or ForceReply.
func (o OptionsSendMessage) SetReplyMarkup(replyMarkup any) OptionsSendMessage {
	o["reply_markup"] = replyMarkup
	return o
}

// OptionsForwardMessage struct for ForwardMessage().
//
// options include: `message_thread_id`, `direct_messages_topic_id`, `video_start_timestamp`, `disable_notification`, `protect_content`, `message_effect_id`, and `suggested_post_parameters`.
//
// https://core.telegram.org/bots/api#forwardmessage
type OptionsForwardMessage MethodOptions

// SetMessageThreadID sets the `message_thread_id` value of OptionsForwardMessage.
func (o OptionsForwardMessage) SetMessageThreadID(messageThreadID int64) OptionsForwardMessage {
	o["message_thread_id"] = messageThreadID
	return o
}

// SetDirectMessagesTopicID sets the `direct_messages_topic_id` value of OptionsForwardMessage.
func (o OptionsForwardMessage) SetDirectMessagesTopicID(directMessagesTopicID int64) OptionsForwardMessage {
	o["direct_messages_topic_id"] = directMessagesTopicID
	return o
}

// SetVideoStartTimestamp sets the `video_start_timestamp` value of OptionsForwardMessage.
func (o OptionsForwardMessage) SetVideoStartTimestamp(videoStartTimestamp int) OptionsForwardMessage {
	o["video_start_timestamp"] = videoStartTimestamp
	return o
}

// SetDisableNotification sets the `disable_notification` value of OptionsForwardMessage.
func (o OptionsForwardMessage) SetDisableNotification(disable bool) OptionsForwardMessage {
	o["disable_notification"] = disable
	return o
}

// SetProtectContent sets the `protect_content` value of OptionsForwardMessage.
func (o OptionsForwardMessage) SetProtectContent(protect bool) OptionsForwardMessage {
	o["protect_content"] = protect
	return o
}

// SetMessageEffectID sets the `message_effect_id` value of OptionsForwardMessage.
func (o OptionsForwardMessage) SetMessageEffectID(messageEffectID string) OptionsForwardMessage {
	o["message_effect_id"] = messageEffectID
	return o
}

// SetSuggestedPostParameters sets the `suggested_post_parameters` value of OptionsForwardMessage.
func (o OptionsForwardMessage) SetSuggestedPostParameters(suggestedPostParameters SuggestedPostParameters) OptionsForwardMessage {
	o["suggested_post_parameters"] = suggestedPostParameters
	return o
}

// OptionsForwardMessages struct for ForwardMessages().
//
// options include: `message_thread_id`, `direct_messages_topic_id`, `disable_notification` and `protect_content`.
//
// https://core.telegram.org/bots/api#forwardmessages
type OptionsForwardMessages MethodOptions

// SetMessageThreadID sets the `message_thread_id` value of OptionsForwardMessages.
func (o OptionsForwardMessages) SetMessageThreadID(messageThreadID int64) OptionsForwardMessages {
	o["message_thread_id"] = messageThreadID
	return o
}

// SetDirectMessagesTopicID sets the `direct_messages_topic_id` value of OptionsForwardMessages.
func (o OptionsForwardMessages) SetDirectMessagesTopicID(directMessagesTopicID int64) OptionsForwardMessages {
	o["direct_messages_topic_id"] = directMessagesTopicID
	return o
}

// SetDisableNotification sets the `disable_notification` value of OptionsForwardMessages.
func (o OptionsForwardMessages) SetDisableNotification(disable bool) OptionsForwardMessages {
	o["disable_notification"] = disable
	return o
}

// SetProtectContent sets the `protect_content` value of OptionsForwardMessages.
func (o OptionsForwardMessages) SetProtectContent(protect bool) OptionsForwardMessages {
	o["protect_content"] = protect
	return o
}

// OptionsCopyMessage struct for CopyMessage().
//
// options include: `message_thread_id`, `direct_messages_topic_id`, `video_start_timestamp`, `caption`, `parse_mode`, `caption_entities`, `show_caption_above_media`, `disable_notification`, `protect_content`, `allow_paid_broadcast`, `message_effect_id`, `suggested_post_parameters`, `reply_parameters`, and `reply_markup`
//
// https://core.telegram.org/bots/api#copymessage
type OptionsCopyMessage MethodOptions

// SetMessageThreadID sets the `message_thread_id` value of OptionsCopyMessage.
func (o OptionsCopyMessage) SetMessageThreadID(messageThreadID int64) OptionsCopyMessage {
	o["message_thread_id"] = messageThreadID
	return o
}

// SetDirectMessagesTopicID sets the `direct_messages_topic_id` value of OptionsCopyMessage.
func (o OptionsCopyMessage) SetDirectMessagesTopicID(directMessagesTopicID int64) OptionsCopyMessage {
	o["direct_messages_topic_id"] = directMessagesTopicID
	return o
}

// SetVideoStartTimestamp sets the `video_start_timestamp` value of OptionsCopyMessage.
func (o OptionsCopyMessage) SetVideoStartTimestamp(videoStartTimestamp int) OptionsCopyMessage {
	o["video_start_timestamp"] = videoStartTimestamp
	return o
}

// SetCaption sets the `caption` value of OptionsCopyMessage.
func (o OptionsCopyMessage) SetCaption(caption string) OptionsCopyMessage {
	o["caption"] = caption
	return o
}

// SetParseMode sets the `parse_mode` value of OptionsCopyMessage.
func (o OptionsCopyMessage) SetParseMode(parseMode ParseMode) OptionsCopyMessage {
	o["parse_mode"] = parseMode
	return o
}

// SetCaptionEntities sets the `caption_entities` value of OptionsCopyMessage.
func (o OptionsCopyMessage) SetCaptionEntities(entities []MessageEntity) OptionsCopyMessage {
	o["caption_entities"] = entities
	return o
}

// SetShowCaptionAboveMedia sets the `show_caption_above_media` value of OptionsCopyMessage.
func (o OptionsCopyMessage) SetShowCaptionAboveMedia(showCaptionAboveMedia bool) OptionsCopyMessage {
	o["show_caption_above_media"] = showCaptionAboveMedia
	return o
}

// SetDisableNotification sets the `disable_notification` value of OptionsCopyMessage.
func (o OptionsCopyMessage) SetDisableNotification(disable bool) OptionsCopyMessage {
	o["disable_notification"] = disable
	return o
}

// SetProtectContent sets the `protect_content` value of OptionsCopyMessage.
func (o OptionsCopyMessage) SetProtectContent(protect bool) OptionsCopyMessage {
	o["protect_content"] = protect
	return o
}

// SetAllowPaidBroadcast sets the `allow_paid_broadcast` value of OptionsCopyMessage.
func (o OptionsCopyMessage) SetAllowPaidBroadcast(allow bool) OptionsCopyMessage {
	o["allow_paid_broadcast"] = allow
	return o
}

// SetMessageEffectID sets the `message_effect_id` value of OptionsCopyMessage.
func (o OptionsCopyMessage) SetMessageEffectID(messageEffectID string) OptionsCopyMessage {
	o["message_effect_id"] = messageEffectID
	return o
}

// SetSuggestedPostParameters sets the `suggested_post_parameters` value of OptionsCopyMessage.
func (o OptionsCopyMessage) SetSuggestedPostParameters(suggestedPostParameters SuggestedPostParameters) OptionsCopyMessage {
	o["suggested_post_parameters"] = suggestedPostParameters
	return o
}

// SetReplyParameters sets the `reply_parameters` value of OptionsCopyMessage.
func (o OptionsCopyMessage) SetReplyParameters(replyParameters ReplyParameters) OptionsCopyMessage {
	o["reply_parameters"] = replyParameters
	return o
}

// SetReplyMarkup sets the reply_markup value of OptionsCopyMessage.
//
// `replyMarkup` can be one of InlineKeyboardMarkup, ReplyKeyboardMarkup, ReplyKeyboardRemove, or ForceReply.
func (o OptionsCopyMessage) SetReplyMarkup(replyMarkup any) OptionsCopyMessage {
	o["reply_markup"] = replyMarkup
	return o
}

// OptionsCopyMessages struct for CopyMessages().
//
// options include: `message_thread_id`, `direct_messages_topic_id`, `disable_notification`, `protect_content`, and `remove_caption`
//
// https://core.telegram.org/bots/api#copymessages
type OptionsCopyMessages MethodOptions

// SetMessageThreadID sets the `message_thread_id` value of OptionsCopyMessages.
func (o OptionsCopyMessages) SetMessageThreadID(messageThreadID int64) OptionsCopyMessages {
	o["message_thread_id"] = messageThreadID
	return o
}

// SetDirectMessagesTopicID sets the `direct_messages_topic_id` value of OptionsCopyMessages.
func (o OptionsCopyMessages) SetDirectMessagesTopicID(directMessagesTopicID int64) OptionsCopyMessages {
	o["direct_messages_topic_id"] = directMessagesTopicID
	return o
}

// SetDisableNotification sets the `disable_notification` value of OptionsCopyMessages.
func (o OptionsCopyMessages) SetDisableNotification(disable bool) OptionsCopyMessages {
	o["disable_notification"] = disable
	return o
}

// SetProtectContent sets the `protect_content` value of OptionsCopyMessages.
func (o OptionsCopyMessages) SetProtectContent(protect bool) OptionsCopyMessages {
	o["protect_content"] = protect
	return o
}

// SetRemoveCaption sets the `remove_caption` value of OptionsCopyMessages.
func (o OptionsCopyMessages) SetRemoveCaption(removeCaption bool) OptionsCopyMessages {
	o["remove_caption"] = removeCaption
	return o
}

// OptionsSendPhoto struct for SendPhoto().
//
// options include: `business_connection_id`, `message_thread_id`, `direct_messages_topic_id`, `caption`, `parse_mode`, `caption_entities`, `show_caption_above_media`, `has_spoiler`, `disable_notification`, `protect_content`, `allow_paid_broadcast`,  `message_effect_id`, `suggested_post_parameters`, `reply_parameters`, and `reply_markup`.
//
// https://core.telegram.org/bots/api#sendphoto
type OptionsSendPhoto MethodOptions

// SetBusinessConnectionID sets the `business_connection_id` value of OptionsSendPhoto.
func (o OptionsSendPhoto) SetBusinessConnectionID(businessConnectionID string) OptionsSendPhoto {
	o["business_connection_id"] = businessConnectionID
	return o
}

// SetMessageThreadID sets the `message_thread_id`value of OptionsSendPhoto.
func (o OptionsSendPhoto) SetMessageThreadID(messageThreadID int64) OptionsSendPhoto {
	o["message_thread_id"] = messageThreadID
	return o
}

// SetDirectMessagesTopicID sets the `direct_messages_topic_id` value of OptionsSendPhoto.
func (o OptionsSendPhoto) SetDirectMessagesTopicID(directMessagesTopicID int64) OptionsSendPhoto {
	o["direct_messages_topic_id"] = directMessagesTopicID
	return o
}

// SetCaption sets the `caption` value of OptionsSendPhoto.
func (o OptionsSendPhoto) SetCaption(caption string) OptionsSendPhoto {
	o["caption"] = caption
	return o
}

// SetParseMode sets the `parse_mode` value of OptionsSendPhoto.
func (o OptionsSendPhoto) SetParseMode(parseMode ParseMode) OptionsSendPhoto {
	o["parse_mode"] = parseMode
	return o
}

// SetCaptionEntities sets the `caption_entities` value of OptionsSendPhoto.
func (o OptionsSendPhoto) SetCaptionEntities(entities []MessageEntity) OptionsSendPhoto {
	o["caption_entities"] = entities
	return o
}

// SetShowCaptionAboveMedia sets the `show_caption_above_media` value of OptionsSendPhoto.
func (o OptionsSendPhoto) SetShowCaptionAboveMedia(showCaptionAboveMedia bool) OptionsSendPhoto {
	o["show_caption_above_media"] = showCaptionAboveMedia
	return o
}

// SetHasSpoiler sets the `has_spoiler` value of OptionsSendPhoto.
func (o OptionsSendPhoto) SetHasSpoiler(hasSpoiler bool) OptionsSendPhoto {
	o["has_spoiler"] = hasSpoiler
	return o
}

// SetDisableNotification sets the `disable_notification` value of OptionsSendPhoto.
func (o OptionsSendPhoto) SetDisableNotification(disable bool) OptionsSendPhoto {
	o["disable_notification"] = disable
	return o
}

// SetProtectContent sets the `protect_content` value of OptionsSendPhoto.
func (o OptionsSendPhoto) SetProtectContent(protect bool) OptionsSendPhoto {
	o["protect_content"] = protect
	return o
}

// SetAllowPaidBroadcast sets the `allow_paid_broadcast` value of OptionsSendPhoto.
func (o OptionsSendPhoto) SetAllowPaidBroadcast(allow bool) OptionsSendPhoto {
	o["allow_paid_broadcast"] = allow
	return o
}

// SetMessageEffectID sets the `message_effect_id` value of OptionsSendPhoto.
func (o OptionsSendPhoto) SetMessageEffectID(messageEffectID string) OptionsSendPhoto {
	o["message_effect_id"] = messageEffectID
	return o
}

// SetSuggestedPostParameters sets the `suggested_post_parameters` value of OptionsSendPhoto.
func (o OptionsSendPhoto) SetSuggestedPostParameters(suggestedPostParameters SuggestedPostParameters) OptionsSendPhoto {
	o["suggested_post_parameters"] = suggestedPostParameters
	return o
}

// SetReplyParameters sets the `reply_parameters` value of OptionsSendPhoto.
func (o OptionsSendPhoto) SetReplyParameters(replyParameters ReplyParameters) OptionsSendPhoto {
	o["reply_parameters"] = replyParameters
	return o
}

// SetReplyMarkup sets the `reply_markup` value of OptionsSendPhoto.
//
// `replyMarkup` can be one of InlineKeyboardMarkup, ReplyKeyboardMarkup, ReplyKeyboardRemove, or ForceReply.
func (o OptionsSendPhoto) SetReplyMarkup(replyMarkup any) OptionsSendPhoto {
	o["reply_markup"] = replyMarkup
	return o
}

// OptionsSendAudio struct for SendAudio().
//
// options include: `business_connection_id`, `message_thread_id`, `direct_messages_topic_id`, `caption`, `parse_mode`, `caption_entities`, `duration`, `performer`, `title`, `thumbnail`, `disable_notification`, `protect_content`, `allow_paid_broadcast`, `message_effect_id`, `suggested_post_parameters`, `reply_parameters`, and `reply_markup`.
//
// https://core.telegram.org/bots/api#sendaudio
type OptionsSendAudio MethodOptions

// SetBusinessConnectionID sets the `business_connection_id` value of OptionsSendAudio.
func (o OptionsSendAudio) SetBusinessConnectionID(businessConnectionID string) OptionsSendAudio {
	o["business_connection_id"] = businessConnectionID
	return o
}

// SetMessageThreadID sets the `message_thread_id` value of OptionsSendAudio.
func (o OptionsSendAudio) SetMessageThreadID(messageThreadID int64) OptionsSendAudio {
	o["message_thread_id"] = messageThreadID
	return o
}

// SetDirectMessagesTopicID sets the `direct_messages_topic_id` value of OptionsSendAudio.
func (o OptionsSendAudio) SetDirectMessagesTopicID(directMessagesTopicID int64) OptionsSendAudio {
	o["direct_messages_topic_id"] = directMessagesTopicID
	return o
}

// SetCaption sets the `caption` value of OptionsSendAudio.
func (o OptionsSendAudio) SetCaption(caption string) OptionsSendAudio {
	o["caption"] = caption
	return o
}

// SetParseMode sets the `parse_mode` value of OptionsSendAudio.
func (o OptionsSendAudio) SetParseMode(parseMode ParseMode) OptionsSendAudio {
	o["parse_mode"] = parseMode
	return o
}

// SetCaptionEntities sets the `caption_entities` value of OptionsSendAudio.
func (o OptionsSendAudio) SetCaptionEntities(entities []MessageEntity) OptionsSendAudio {
	o["caption_entities"] = entities
	return o
}

// SetDuration sets the `duration` value of OptionsSendAudio.
func (o OptionsSendAudio) SetDuration(duration int) OptionsSendAudio {
	o["duration"] = duration
	return o
}

// SetPerformer sets the `performer` value of OptionsSendAudio.
func (o OptionsSendAudio) SetPerformer(performer string) OptionsSendAudio {
	o["performer"] = performer
	return o
}

// SetTitle sets the `title` value of OptionsSendAudio.
func (o OptionsSendAudio) SetTitle(title string) OptionsSendAudio {
	o["title"] = title
	return o
}

// SetThumbnail sets the `thumbnail` value of OptionsSendAudio.
//
// `thumbnail` can be one of InputFile or string.
func (o OptionsSendAudio) SetThumbnail(thumbnail any) OptionsSendAudio {
	o["thumbnail"] = thumbnail
	return o
}

// SetDisableNotification sets the `disable_notification` value of OptionsSendAudio.
func (o OptionsSendAudio) SetDisableNotification(disable bool) OptionsSendAudio {
	o["disable_notification"] = disable
	return o
}

// SetProtectContent sets the `protect_content` value of OptionsSendAudio.
func (o OptionsSendAudio) SetProtectContent(protect bool) OptionsSendAudio {
	o["protect_content"] = protect
	return o
}

// SetAllowPaidBroadcast sets the `allow_paid_broadcast` value of OptionsSendAudio.
func (o OptionsSendAudio) SetAllowPaidBroadcast(allow bool) OptionsSendAudio {
	o["allow_paid_broadcast"] = allow
	return o
}

// SetMessageEffectID sets the `message_effect_id` value of OptionsSendAudio.
func (o OptionsSendAudio) SetMessageEffectID(messageEffectID string) OptionsSendAudio {
	o["message_effect_id"] = messageEffectID
	return o
}

// SetSuggestedPostParameters sets the `suggested_post_parameters` value of OptionsSendAudio.
func (o OptionsSendAudio) SetSuggestedPostParameters(suggestedPostParameters SuggestedPostParameters) OptionsSendAudio {
	o["suggested_post_parameters"] = suggestedPostParameters
	return o
}

// SetReplyParameters sets the `reply_parameters` value of OptionsSendAudio.
func (o OptionsSendAudio) SetReplyParameters(replyParameters ReplyParameters) OptionsSendAudio {
	o["reply_parameters"] = replyParameters
	return o
}

// SetReplyMarkup sets the `reply_markup` value of OptionsSendAudio.
//
// `replyMarkup` can be one of InlineKeyboardMarkup, ReplyKeyboardMarkup, ReplyKeyboardRemove, or ForceReply.
func (o OptionsSendAudio) SetReplyMarkup(replyMarkup any) OptionsSendAudio {
	o["reply_markup"] = replyMarkup
	return o
}

// OptionsSendDocument struct for SendDocument().
//
// options include: `business_connection_id`, `message_thread_id`, `direct_messages_topic_id`, `thumbnail`, `caption`, `parse_mode`, `caption_entities`, `disable_content_type_detection`, `disable_notification`, `protect_content`, `allow_paid_broadcast`, `message_effect_id`, `suggested_post_parameters`, `reply_parameters`, and `reply_markup`.
//
// https://core.telegram.org/bots/api#senddocument
type OptionsSendDocument MethodOptions

// SetBusinessConnectionID sets the `business_connection_id` value of OptionsSendDocument.
func (o OptionsSendDocument) SetBusinessConnectionID(businessConnectionID string) OptionsSendDocument {
	o["business_connection_id"] = businessConnectionID
	return o
}

// SetMessageThreadID sets the `message_thread_id` value of OptionsSendDocument.
func (o OptionsSendDocument) SetMessageThreadID(messageThreadID int64) OptionsSendDocument {
	o["message_thread_id"] = messageThreadID
	return o
}

// SetDirectMessagesTopicID sets the `direct_messages_topic_id` value of OptionsSendDocument.
func (o OptionsSendDocument) SetDirectMessagesTopicID(directMessagesTopicID int64) OptionsSendDocument {
	o["direct_messages_topic_id"] = directMessagesTopicID
	return o
}

// SetThumbnail sets the thumbnail value of OptionsSendDocument.
//
// `thumbnail` can be one of InputFile or string.
func (o OptionsSendDocument) SetThumbnail(thumbnail any) OptionsSendDocument {
	o["thumbnail"] = thumbnail
	return o
}

// SetCaption sets the `caption` value of OptionsSendDocument.
func (o OptionsSendDocument) SetCaption(caption string) OptionsSendDocument {
	o["caption"] = caption
	return o
}

// SetParseMode sets the `parse_mode` value of OptionsSendDocument.
func (o OptionsSendDocument) SetParseMode(parseMode ParseMode) OptionsSendDocument {
	o["parse_mode"] = parseMode
	return o
}

// SetCaptionEntities sets the `caption_entities` value of OptionsSendDocument.
func (o OptionsSendDocument) SetCaptionEntities(entities []MessageEntity) OptionsSendDocument {
	o["caption_entities"] = entities
	return o
}

// SetDisableContentTypeDetection sets the `disable_content_type_detection` value of OptionsSendDocument.
func (o OptionsSendDocument) SetDisableContentTypeDetection(disable bool) OptionsSendDocument {
	o["disable_content_type_detection"] = disable
	return o
}

// SetDisableNotification sets the `disable_notification` value of OptionsSendDocument.
func (o OptionsSendDocument) SetDisableNotification(disable bool) OptionsSendDocument {
	o["disable_notification"] = disable
	return o
}

// SetProtectContent sets the `protect_content` value of OptionsSendDocument.
func (o OptionsSendDocument) SetProtectContent(protect bool) OptionsSendDocument {
	o["protect_content"] = protect
	return o
}

// SetAllowPaidBroadcast sets the `allow_paid_broadcast` value of OptionsSendDocument.
func (o OptionsSendDocument) SetAllowPaidBroadcast(allow bool) OptionsSendDocument {
	o["allow_paid_broadcast"] = allow
	return o
}

// SetMessageEffectID sets the `message_effect_id` value of OptionsSendDocument.
func (o OptionsSendDocument) SetMessageEffectID(messageEffectID string) OptionsSendDocument {
	o["message_effect_id"] = messageEffectID
	return o
}

// SetSuggestedPostParameters sets the `suggested_post_parameters` value of OptionsSendDocument.
func (o OptionsSendDocument) SetSuggestedPostParameters(suggestedPostParameters SuggestedPostParameters) OptionsSendDocument {
	o["suggested_post_parameters"] = suggestedPostParameters
	return o
}

// SetReplyParameters sets the `reply_parameters` value of OptionsSendDocument.
func (o OptionsSendDocument) SetReplyParameters(replyParameters ReplyParameters) OptionsSendDocument {
	o["reply_parameters"] = replyParameters
	return o
}

// SetReplyMarkup sets the reply_markup value of OptionsSendDocument.
//
// `replyMarkup` can be one of InlineKeyboardMarkup, ReplyKeyboardMarkup, ReplyKeyboardRemove, or ForceReply.
func (o OptionsSendDocument) SetReplyMarkup(replyMarkup any) OptionsSendDocument {
	o["reply_markup"] = replyMarkup
	return o
}

// OptionsSendSticker struct for SendSticker().
//
// options include: `business_connection_id`, `message_thread_id`, `direct_messages_topic_id`, `emoji`, `disable_notification`, `protect_content`, `allow_paid_broadcast`, `message_effect_id`, `suggested_post_parameters`, `reply_parameters`, and `reply_markup`.
//
// https://core.telegram.org/bots/api#sendsticker
type OptionsSendSticker MethodOptions

// SetBusinessConnectionID sets the `business_connection_id` value of OptionsSendSticker.
func (o OptionsSendSticker) SetBusinessConnectionID(businessConnectionID string) OptionsSendSticker {
	o["business_connection_id"] = businessConnectionID
	return o
}

// SetMessageThreadID sets the `message_thread_id` value of OptionsSendSticker.
func (o OptionsSendSticker) SetMessageThreadID(messageThreadID int64) OptionsSendSticker {
	o["message_thread_id"] = messageThreadID
	return o
}

// SetDirectMessagesTopicID sets the `direct_messages_topic_id` value of OptionsSendSticker.
func (o OptionsSendSticker) SetDirectMessagesTopicID(directMessagesTopicID int64) OptionsSendSticker {
	o["direct_messages_topic_id"] = directMessagesTopicID
	return o
}

// SetEmoji sets the `emoji` value of OptionsSendSticker.
func (o OptionsSendSticker) SetEmoji(emoji string) OptionsSendSticker {
	o["emoji"] = emoji
	return o
}

// SetDisableNotification sets the `disable_notification` value of OptionsSendSticker.
func (o OptionsSendSticker) SetDisableNotification(disable bool) OptionsSendSticker {
	o["disable_notification"] = disable
	return o
}

// SetProtectContent sets the `protect_content` value of OptionsSendSticker.
func (o OptionsSendSticker) SetProtectContent(protect bool) OptionsSendSticker {
	o["protect_content"] = protect
	return o
}

// SetAllowPaidBroadcast sets the `allow_paid_broadcast` value of OptionsSendSticker.
func (o OptionsSendSticker) SetAllowPaidBroadcast(allow bool) OptionsSendSticker {
	o["allow_paid_broadcast"] = allow
	return o
}

// SetMessageEffectID sets the `message_effect_id` value of OptionsSendSticker.
func (o OptionsSendSticker) SetMessageEffectID(messageEffectID string) OptionsSendSticker {
	o["message_effect_id"] = messageEffectID
	return o
}

// SetSuggestedPostParameters sets the `suggested_post_parameters` value of OptionsSendSticker.
func (o OptionsSendSticker) SetSuggestedPostParameters(suggestedPostParameters SuggestedPostParameters) OptionsSendSticker {
	o["suggested_post_parameters"] = suggestedPostParameters
	return o
}

// SetReplyParameters sets the `reply_parameters` value of OptionsSendSticker.
func (o OptionsSendSticker) SetReplyParameters(replyParameters ReplyParameters) OptionsSendSticker {
	o["reply_parameters"] = replyParameters
	return o
}

// SetReplyMarkup sets the `reply_markup` value of OptionsSendSticker.
//
// `replyMarkup` can be one of InlineKeyboardMarkup, ReplyKeyboardMarkup, ReplyKeyboardRemove, or ForceReply.
func (o OptionsSendSticker) SetReplyMarkup(replyMarkup any) OptionsSendSticker {
	o["reply_markup"] = replyMarkup
	return o
}

// OptionsCreateNewStickerSet struct for CreateNewStickerSet().
//
// options include: `sticker_type`, and `needs_repainting`.
//
// https://core.telegram.org/bots/api#createnewstickerset
type OptionsCreateNewStickerSet MethodOptions

// SetStickerType sets the `sticker_type` value of OptionsCreateNewStickerSet.
func (o OptionsCreateNewStickerSet) SetStickerType(stickerType StickerType) OptionsCreateNewStickerSet {
	o["sticker_type"] = stickerType
	return o
}

// SetNeedsRepainting sets the `needs_repainting` value of OptionsCreateNewStickerSet.
func (o OptionsCreateNewStickerSet) SetNeedsRepainting(needsRepainting bool) OptionsCreateNewStickerSet {
	o["needs_repainting"] = needsRepainting
	return o
}

// OptionsAddStickerToSet struct for AddStickerToSet()
//
// options include: nothing for now
//
// https://core.telegram.org/bots/api#addstickertoset
type OptionsAddStickerToSet MethodOptions

// OptionsSetStickerSetThumbnail struct for SetStickerSetThumbnail()
//
// options include: `thumbnail`.
//
// https://core.telegram.org/bots/api#setstickersetthumbnail
type OptionsSetStickerSetThumbnail MethodOptions

// SetThumbnail sets the `thumbnail` value of OptionsSetStickerSetThumbnail.
func (o OptionsSetStickerSetThumbnail) SetThumbnail(thumbnail InputFile) OptionsSetStickerSetThumbnail {
	o["thumbnail"] = thumbnail
	return o
}

// SetThumbnailString sets the `thumbnail` value of OptionsSetStickerSetThumbnail.
//
// `thumbnail` can be a file_id or a http url to a file
func (o OptionsSetStickerSetThumbnail) SetThumbnailString(thumbnail string) OptionsSetStickerSetThumbnail {
	o["thumbnail"] = thumbnail
	return o
}

// OptionsSetCustomEmojiStickerSetThumbnail struct for SetCustomEmojiStickerSet()
//
// options include: `custom_emoji_id`.
//
// https://core.telegram.org/bots/api#setcustomemojistickersetthumbnail
type OptionsSetCustomEmojiStickerSetThumbnail MethodOptions

// SetCustomEmojiID sets the `custom_emoji_id` value of OptionsSetCustomEmojiStickerSetThumbnail.
func (o OptionsSetCustomEmojiStickerSetThumbnail) SetCustomEmojiID(customEmojiID string) OptionsSetCustomEmojiStickerSetThumbnail {
	o["custom_emoji_id"] = customEmojiID
	return o
}

// OptionsSetStickerMaskPosition struct for SetStickerMaskPosition()
//
// options include: `mask_position`.
//
// https://core.telegram.org/bots/api#setstickermaskposition
type OptionsSetStickerMaskPosition MethodOptions

// SetMaskPosition sets the `mask_position` value of OptionsSetStickerMaskPosition.
func (o OptionsSetStickerMaskPosition) SetMaskPosition(maskPosition MaskPosition) OptionsSetStickerMaskPosition {
	o["mask_position"] = maskPosition
	return o
}

// OptionsSendGift struct for SendGift().
//
// options include: `user_id`, `chat_id`, `pay_for_upgrade`, `text`, `text_parse_mode`, and `text_entities`.
//
// https://core.telegram.org/bots/api#sendgift
type OptionsSendGift MethodOptions

// SetUserID sets the `user_id` value of OptionsSendGift.
func (o OptionsSendGift) SetUserID(userID int64) OptionsSendGift {
	o["user_id"] = userID
	return o
}

// SetChatID sets the `chat_id` value of OptionsSendGift.
func (o OptionsSendGift) SetChatID(chatID ChatID) OptionsSendGift {
	o["chat_id"] = chatID
	return o
}

// SetPayForUpgrade sets the `pay_for_upgrade` value of OptionsSendGift.
func (o OptionsSendGift) SetPayForUpgrade(payForUpgrade bool) OptionsSendGift {
	o["pay_for_upgrade"] = payForUpgrade
	return o
}

// SetText sets the `text` value of OptionsSendGift.
func (o OptionsSendGift) SetText(text string) OptionsSendGift {
	o["text"] = text
	return o
}

// SetTextParseMode sets the `text_parse_mode` value of OptionsSendGift.
func (o OptionsSendGift) SetTextParseMode(textParseMode string) OptionsSendGift {
	o["text_parse_mode"] = textParseMode
	return o
}

// SetTextEntities sets the `text_entities` value of OptionsSendGift.
func (o OptionsSendGift) SetTextEntities(textEntities []MessageEntity) OptionsSendGift {
	o["text_entities"] = textEntities
	return o
}

// OptionsVerifyUser struct for VerifyUser().
//
// options include: `custom_description`.
//
// https://core.telegram.org/bots/api#verifyuser
type OptionsVerifyUser MethodOptions

// SetCustomDescription sets the `custom_description` value of OptionsVerifyUser.
func (o OptionsVerifyUser) SetCustomDescription(customDescription string) OptionsVerifyUser {
	o["custom_description"] = customDescription
	return o
}

// OptionsVerifyChat struct for VerifyChat().
//
// options include: `custom_description`.
//
// https://core.telegram.org/bots/api#verifychat
type OptionsVerifyChat MethodOptions

// SetCustomDescription sets the `custom_description` value of OptionsVerifyChat.
func (o OptionsVerifyChat) SetCustomDescription(customDescription string) OptionsVerifyChat {
	o["custom_description"] = customDescription
	return o
}

// OptionsSendVideo struct for SendVideo().
//
// options include: `business_connection_id`, `message_thread_id`, `direct_messages_topic_id`, `duration`, `width`, `height`, `thumbnail`, `cover`, `start_timestamp`, `caption`, `parse_mode`, `caption_entities`, `show_caption_above_media`, `has_spoiler`, `supports_streaming`, `disable_notification`, `protect_content`, `allow_paid_broadcast`, `message_effect_id`, `suggested_post_parameters`, `reply_parameters`, and `reply_markup`.
//
// https://core.telegram.org/bots/api#sendvideo
type OptionsSendVideo MethodOptions

// SetBusinessConnectionID sets the `business_connection_id` value of OptionsSendVideo.
func (o OptionsSendVideo) SetBusinessConnectionID(businessConnectionID string) OptionsSendVideo {
	o["business_connection_id"] = businessConnectionID
	return o
}

// SetMessageThreadID sets the `message_thread_id` value of OptionsSendVideo.
func (o OptionsSendVideo) SetMessageThreadID(messageThreadID int64) OptionsSendVideo {
	o["message_thread_id"] = messageThreadID
	return o
}

// SetDirectMessagesTopicID sets the `direct_messages_topic_id` value of OptionsSendVideo.
func (o OptionsSendVideo) SetDirectMessagesTopicID(directMessagesTopicID int64) OptionsSendVideo {
	o["direct_messages_topic_id"] = directMessagesTopicID
	return o
}

// SetDuration sets the `duration` value of OptionsSendVideo.
func (o OptionsSendVideo) SetDuration(duration int) OptionsSendVideo {
	o["duration"] = duration
	return o
}

// SetWidth sets the `width` value of OptionsSendVideo.
func (o OptionsSendVideo) SetWidth(width int) OptionsSendVideo {
	o["width"] = width
	return o
}

// SetHeight sets the `height` value of OptionsSendVideo.
func (o OptionsSendVideo) SetHeight(height int) OptionsSendVideo {
	o["height"] = height
	return o
}

// SetThumbnail sets the `thumbnail` value of OptionsSendVideo.
//
// `thumbnail` can be one of InputFile or string.
func (o OptionsSendVideo) SetThumbnail(thumbnail any) OptionsSendVideo {
	o["thumbnail"] = thumbnail
	return o
}

// SetCover sets the `cover` value of OptionsSendVideo.
//
// `cover` can be one of InputFile or string.
func (o OptionsSendVideo) SetCover(cover any) OptionsSendVideo {
	o["cover"] = cover
	return o
}

// SetStartTimestamp sets the `start_timestamp` value of OptionsSendVideo.
func (o OptionsSendVideo) SetStartTimestamp(startTimestamp int) OptionsSendVideo {
	o["start_timestamp"] = startTimestamp
	return o
}

// SetCaption sets the `caption` value of OptionsSendVideo.
func (o OptionsSendVideo) SetCaption(caption string) OptionsSendVideo {
	o["caption"] = caption
	return o
}

// SetParseMode sets the `parse_mode` value of OptionsSendVideo.
func (o OptionsSendVideo) SetParseMode(parseMode ParseMode) OptionsSendVideo {
	o["parse_mode"] = parseMode
	return o
}

// SetCaptionEntities sets the `caption_entities` value of OptionsSendVideo.
func (o OptionsSendVideo) SetCaptionEntities(entities []MessageEntity) OptionsSendVideo {
	o["caption_entities"] = entities
	return o
}

// SetShowCaptionAboveMedia sets the `show_caption_above_media` value of OptionsSendVideo.
func (o OptionsSendVideo) SetShowCaptionAboveMedia(showCaptionAboveMedia bool) OptionsSendVideo {
	o["show_caption_above_media"] = showCaptionAboveMedia
	return o
}

// SetHasSpoiler sets the `has_spoiler` value of OptionsSendVideo.
func (o OptionsSendVideo) SetHasSpoiler(hasSpoiler bool) OptionsSendVideo {
	o["has_spoiler"] = hasSpoiler
	return o
}

// SetSupportsStreaming sets the `supports_streaming` value of OptionsSendVideo.
func (o OptionsSendVideo) SetSupportsStreaming(supportsStreaming bool) OptionsSendVideo {
	o["supports_streaming"] = supportsStreaming
	return o
}

// SetDisableNotification sets the `disable_notification` value of OptionsSendVideo.
func (o OptionsSendVideo) SetDisableNotification(disable bool) OptionsSendVideo {
	o["disable_notification"] = disable
	return o
}

// SetProtectContent sets the `protect_content` value of OptionsSendVideo.
func (o OptionsSendVideo) SetProtectContent(protect bool) OptionsSendVideo {
	o["protect_content"] = protect
	return o
}

// SetAllowPaidBroadcast sets the `allow_paid_broadcast` value of OptionsSendVideo.
func (o OptionsSendVideo) SetAllowPaidBroadcast(allow bool) OptionsSendVideo {
	o["allow_paid_broadcast"] = allow
	return o
}

// SetMessageEffectID sets the `message_effect_id` value of OptionsSendVideo.
func (o OptionsSendVideo) SetMessageEffectID(messageEffectID string) OptionsSendVideo {
	o["message_effect_id"] = messageEffectID
	return o
}

// SetSuggestedPostParameters sets the `suggested_post_parameters` value of OptionsSendVideo.
func (o OptionsSendVideo) SetSuggestedPostParameters(suggestedPostParameters SuggestedPostParameters) OptionsSendVideo {
	o["suggested_post_parameters"] = suggestedPostParameters
	return o
}

// SetReplyParameters sets the `reply_parameters` value of OptionsSendVideo.
func (o OptionsSendVideo) SetReplyParameters(replyParameters ReplyParameters) OptionsSendVideo {
	o["reply_parameters"] = replyParameters
	return o
}

// SetReplyMarkup sets the `reply_markup` value of OptionsSendVideo.
//
// `replyMarkup` can be one of InlineKeyboardMarkup, ReplyKeyboardMarkup, ReplyKeyboardRemove, or ForceReply.
func (o OptionsSendVideo) SetReplyMarkup(replyMarkup any) OptionsSendVideo {
	o["reply_markup"] = replyMarkup
	return o
}

// OptionsSendAnimation struct for SendAnimation().
//
// options include: `business_connection_id`, `message_thread_id`, `direct_messages_topic_id`, `duration`, `width`, `height`, `thumbnail`, `caption`, `parse_mode`, `caption_entities`, `show_caption_above_media`, `has_spoiler`, `disable_notification`, `protect_content`, `allow_paid_broadcast`, `message_effect_id`, `suggested_post_parameters`, `reply_parameters`, and `reply_markup`.
//
// https://core.telegram.org/bots/api#sendanimation
type OptionsSendAnimation MethodOptions

// SetBusinessConnectionID sets the `business_connection_id` value of OptionsSendAnimation.
func (o OptionsSendAnimation) SetBusinessConnectionID(businessConnectionID string) OptionsSendAnimation {
	o["business_connection_id"] = businessConnectionID
	return o
}

// SetMessageThreadID sets the `message_thread_id` value of OptionsSendAnimation.
func (o OptionsSendAnimation) SetMessageThreadID(messageThreadID int64) OptionsSendAnimation {
	o["message_thread_id"] = messageThreadID
	return o
}

// SetDirectMessagesTopicID sets the `direct_messages_topic_id` value of OptionsSendAnimation.
func (o OptionsSendAnimation) SetDirectMessagesTopicID(directMessagesTopicID int64) OptionsSendAnimation {
	o["direct_messages_topic_id"] = directMessagesTopicID
	return o
}

// SetDuration sets the `duration` value of OptionsSendAnimation.
func (o OptionsSendAnimation) SetDuration(duration int) OptionsSendAnimation {
	o["duration"] = duration
	return o
}

// SetWidth sets the `width` value of OptionsSendAnimation.
func (o OptionsSendAnimation) SetWidth(width int) OptionsSendAnimation {
	o["width"] = width
	return o
}

// SetHeight sets the `height` value of OptionsSendAnimation.
func (o OptionsSendAnimation) SetHeight(height int) OptionsSendAnimation {
	o["height"] = height
	return o
}

// SetThumbnail sets the `thumbnail` value of OptionsSendAnimation.
//
// `thumbnail` can be one of InputFile or string.
func (o OptionsSendAnimation) SetThumbnail(thumbnail any) OptionsSendAnimation {
	o["thumbnail"] = thumbnail
	return o
}

// SetCaption sets the `caption` value of OptionsSendAnimation.
func (o OptionsSendAnimation) SetCaption(caption string) OptionsSendAnimation {
	o["caption"] = caption
	return o
}

// SetParseMode sets the `parse_mode` value of OptionsSendAnimation.
func (o OptionsSendAnimation) SetParseMode(parseMode ParseMode) OptionsSendAnimation {
	o["parse_mode"] = parseMode
	return o
}

// SetCaptionEntities sets the `caption_entities` value of OptionsSendAnimation.
func (o OptionsSendAnimation) SetCaptionEntities(entities []MessageEntity) OptionsSendAnimation {
	o["caption_entities"] = entities
	return o
}

// SetShowCaptionAboveMedia sets the `show_caption_above_media` value of OptionsSendAnimation.
func (o OptionsSendAnimation) SetShowCaptionAboveMedia(showCaptionAboveMedia bool) OptionsSendAnimation {
	o["show_caption_above_media"] = showCaptionAboveMedia
	return o
}

// SetHasSpoiler sets the `has_spoiler` value of OptionsSendAnimation.
func (o OptionsSendAnimation) SetHasSpoiler(hasSpoiler bool) OptionsSendAnimation {
	o["has_spoiler"] = hasSpoiler
	return o
}

// SetDisableNotification sets the `disable_notification` value of OptionsSendAnimation.
func (o OptionsSendAnimation) SetDisableNotification(disable bool) OptionsSendAnimation {
	o["disable_notification"] = disable
	return o
}

// SetProtectContent sets the `protect_content` value of OptionsSendAnimation.
func (o OptionsSendAnimation) SetProtectContent(protect bool) OptionsSendAnimation {
	o["protect_content"] = protect
	return o
}

// SetAllowPaidBroadcast sets the `allow_paid_broadcast` value of OptionsSendAnimation.
func (o OptionsSendAnimation) SetAllowPaidBroadcast(allow bool) OptionsSendAnimation {
	o["allow_paid_broadcast"] = allow
	return o
}

// SetMessageEffectID sets the `message_effect_id` value of OptionsSendAnimation.
func (o OptionsSendAnimation) SetMessageEffectID(messageEffectID string) OptionsSendAnimation {
	o["message_effect_id"] = messageEffectID
	return o
}

// SetSuggestedPostParameters sets the `suggested_post_parameters` value of OptionsSendAnimation.
func (o OptionsSendAnimation) SetSuggestedPostParameters(suggestedPostParameters SuggestedPostParameters) OptionsSendAnimation {
	o["suggested_post_parameters"] = suggestedPostParameters
	return o
}

// SetReplyParameters sets the `reply_parameters` value of OptionsSendAnimation.
func (o OptionsSendAnimation) SetReplyParameters(replyParameters ReplyParameters) OptionsSendAnimation {
	o["reply_parameters"] = replyParameters
	return o
}

// SetReplyMarkup sets the `reply_markup` value of OptionsSendAnimation.
//
// `replyMarkup` can be one of InlineKeyboardMarkup, ReplyKeyboardMarkup, ReplyKeyboardRemove, or ForceReply.
func (o OptionsSendAnimation) SetReplyMarkup(replyMarkup any) OptionsSendAnimation {
	o["reply_markup"] = replyMarkup
	return o
}

// OptionsSendVoice struct for SendVoice().
//
// options include: `business_connection_id`, `message_thread_id`, `direct_messages_topic_id`, `caption`, `parse_mode`, `caption_entities`, `duration`, `disable_notification`, `protect_content`, `allow_paid_broadcast`, `message_effect_id`, `suggested_post_parameters`, `reply_parameters`, and `reply_markup`.
//
// https://core.telegram.org/bots/api#sendvoice
type OptionsSendVoice MethodOptions

// SetBusinessConnectionID sets the `business_connection_id` value of OptionsSendVoice.
func (o OptionsSendVoice) SetBusinessConnectionID(businessConnectionID string) OptionsSendVoice {
	o["business_connection_id"] = businessConnectionID
	return o
}

// SetMessageThreadID sets the `message_thread_id` value of OptionsSendVoice.
func (o OptionsSendVoice) SetMessageThreadID(messageThreadID int64) OptionsSendVoice {
	o["message_thread_id"] = messageThreadID
	return o
}

// SetDirectMessagesTopicID sets the `direct_messages_topic_id` value of OptionsSendVoice.
func (o OptionsSendVoice) SetDirectMessagesTopicID(directMessagesTopicID int64) OptionsSendVoice {
	o["direct_messages_topic_id"] = directMessagesTopicID
	return o
}

// SetCaption sets the `caption` value of OptionsSendVoice.
func (o OptionsSendVoice) SetCaption(caption string) OptionsSendVoice {
	o["caption"] = caption
	return o
}

// SetParseMode sets the `parse_mode` value of OptionsSendVoice.
func (o OptionsSendVoice) SetParseMode(parseMode ParseMode) OptionsSendVoice {
	o["parse_mode"] = parseMode
	return o
}

// SetCaptionEntities sets the `caption_entities` value of OptionsSendVoice.
func (o OptionsSendVoice) SetCaptionEntities(entities []MessageEntity) OptionsSendVoice {
	o["caption_entities"] = entities
	return o
}

// SetDuration sets the `duration` value of OptionsSendVoice.
func (o OptionsSendVoice) SetDuration(duration int) OptionsSendVoice {
	o["duration"] = duration
	return o
}

// SetDisableNotification sets the `disable_notification` value of OptionsSendVoice.
func (o OptionsSendVoice) SetDisableNotification(disable bool) OptionsSendVoice {
	o["disable_notification"] = disable
	return o
}

// SetProtectContent sets the `protect_content` value of OptionsSendVoice.
func (o OptionsSendVoice) SetProtectContent(protect bool) OptionsSendVoice {
	o["protect_content"] = protect
	return o
}

// SetAllowPaidBroadcast sets the `allow_paid_broadcast` value of OptionsSendVoice.
func (o OptionsSendVoice) SetAllowPaidBroadcast(allow bool) OptionsSendVoice {
	o["allow_paid_broadcast"] = allow
	return o
}

// SetMessageEffectID sets the `message_effect_id` value of OptionsSendVoice.
func (o OptionsSendVoice) SetMessageEffectID(messageEffectID string) OptionsSendVoice {
	o["message_effect_id"] = messageEffectID
	return o
}

// SetSuggestedPostParameters sets the `suggested_post_parameters` value of OptionsSendVoice.
func (o OptionsSendVoice) SetSuggestedPostParameters(suggestedPostParameters SuggestedPostParameters) OptionsSendVoice {
	o["suggested_post_parameters"] = suggestedPostParameters
	return o
}

// SetReplyParameters sets the `reply_parameters` value of OptionsSendVoice.
func (o OptionsSendVoice) SetReplyParameters(replyParameters ReplyParameters) OptionsSendVoice {
	o["reply_parameters"] = replyParameters
	return o
}

// SetReplyMarkup sets the `reply_markup` value of OptionsSendVoice.
//
// `replyMarkup` can be one of InlineKeyboardMarkup, ReplyKeyboardMarkup, ReplyKeyboardRemove, or ForceReply.
func (o OptionsSendVoice) SetReplyMarkup(replyMarkup any) OptionsSendVoice {
	o["reply_markup"] = replyMarkup
	return o
}

// OptionsSendVideoNote struct for SendVideoNote().
//
// options include: `business_connection_id`, `message_thread_id`, `direct_messages_topic_id`, `duration`, `length`, `thumbnail`, `disable_notification`, `protect_content`, `allow_paid_broadcast`, `message_effect_id`, `suggested_post_parameters`, `reply_parameters`, and `reply_markup`.
// (XXX: API returns 'Bad Request: wrong video note length' when length is not given / 2017.05.19.)
//
// https://core.telegram.org/bots/api#sendvideonote
type OptionsSendVideoNote MethodOptions

// SetBusinessConnectionID sets the `business_connection_id` value of OptionsSendVideoNote.
func (o OptionsSendVideoNote) SetBusinessConnectionID(businessConnectionID string) OptionsSendVideoNote {
	o["business_connection_id"] = businessConnectionID
	return o
}

// SetMessageThreadID sets the `message_thread_id` value of OptionsSendVideoNote.
func (o OptionsSendVideoNote) SetMessageThreadID(messageThreadID int64) OptionsSendVideoNote {
	o["message_thread_id"] = messageThreadID
	return o
}

// SetDirectMessagesTopicID sets the `direct_messages_topic_id` value of OptionsSendVideoNote.
func (o OptionsSendVideoNote) SetDirectMessagesTopicID(directMessagesTopicID int64) OptionsSendVideoNote {
	o["direct_messages_topic_id"] = directMessagesTopicID
	return o
}

// SetDuration sets the `duration` value of OptionsSendVideoNote.
func (o OptionsSendVideoNote) SetDuration(duration int) OptionsSendVideoNote {
	o["duration"] = duration
	return o
}

// SetLength sets the `length` value of OptionsSendVideoNote.
func (o OptionsSendVideoNote) SetLength(length int) OptionsSendVideoNote {
	o["length"] = length
	return o
}

// SetThumbnail sets the `thumbnail` value of OptionsSendVideoNote.
//
// `thumbnail` can be one of InputFile or string.
func (o OptionsSendVideoNote) SetThumbnail(thumbnail any) OptionsSendVideoNote {
	o["thumbnail"] = thumbnail
	return o
}

// SetDisableNotification sets the `disable_notification` value of OptionsSendVideoNote.
func (o OptionsSendVideoNote) SetDisableNotification(disable bool) OptionsSendVideoNote {
	o["disable_notification"] = disable
	return o
}

// SetProtectContent sets the `protect_content` value of OptionsSendVideoNote.
func (o OptionsSendVideoNote) SetProtectContent(protect bool) OptionsSendVideoNote {
	o["protect_content"] = protect
	return o
}

// SetAllowPaidBroadcast sets the `allow_paid_broadcast` value of OptionsSendVideoNote.
func (o OptionsSendVideoNote) SetAllowPaidBroadcast(allow bool) OptionsSendVideoNote {
	o["allow_paid_broadcast"] = allow
	return o
}

// SetMessageEffectID sets the `message_effect_id` value of OptionsSendVideoNote.
func (o OptionsSendVideoNote) SetMessageEffectID(messageEffectID string) OptionsSendVideoNote {
	o["message_effect_id"] = messageEffectID
	return o
}

// SetSuggestedPostParameters sets the `suggested_post_parameters` value of OptionsSendVideoNote.
func (o OptionsSendVideoNote) SetSuggestedPostParameters(suggestedPostParameters SuggestedPostParameters) OptionsSendVideoNote {
	o["suggested_post_parameters"] = suggestedPostParameters
	return o
}

// SetReplyParameters sets the `reply_parameters` value of OptionsSendVideoNote.
func (o OptionsSendVideoNote) SetReplyParameters(replyParameters ReplyParameters) OptionsSendVideoNote {
	o["reply_parameters"] = replyParameters
	return o
}

// SetReplyMarkup sets the `reply_markup` value of OptionsSendVideoNote.
//
// `replyMarkup` can be one of InlineKeyboardMarkup, ReplyKeyboardMarkup, ReplyKeyboardRemove, or ForceReply.
func (o OptionsSendVideoNote) SetReplyMarkup(replyMarkup any) OptionsSendVideoNote {
	o["reply_markup"] = replyMarkup
	return o
}

// OptionsSendPaidMedia struct for SendPaideMedia().
//
// options include: `business_connection_id`, `message_thread_id`, `direct_messages_topic_id`, `payload`, `caption`, `parse_mode`, `caption_entities`, `show_caption_above_media`, `disable_notification`, `protect_content`, `allow_paid_broadcast`, `suggested_post_parameters`, `reply_parameters`, and `reply_markup`.
//
// https://core.telegram.org/bots/api#sendpaidmedia
type OptionsSendPaidMedia MethodOptions

// SetBusinessConnectionID sets the `business_connection_id` value of OptionsSendPaidMedia.
func (o OptionsSendPaidMedia) SetBusinessConnectionID(businessConnectionID string) OptionsSendPaidMedia {
	o["business_connection_id"] = businessConnectionID
	return o
}

// SetMessageThreadID sets the `message_thread_id` value of OptionsSendPaidMedia.
func (o OptionsSendPaidMedia) SetMessageThreadID(messageThreadID int64) OptionsSendPaidMedia {
	o["message_thread_id"] = messageThreadID
	return o
}

// SetDirectMessagesTopicID sets the `direct_messages_topic_id` value of OptionsSendPaidMedia.
func (o OptionsSendPaidMedia) SetDirectMessagesTopicID(directMessagesTopicID int64) OptionsSendPaidMedia {
	o["direct_messages_topic_id"] = directMessagesTopicID
	return o
}

// SetPayload sets the `payload` value of OptionsSendPaidMedia.
func (o OptionsSendPaidMedia) SetPayload(payload string) OptionsSendPaidMedia {
	o["payload"] = payload
	return o
}

// SetCaption sets the `caption` value of OptionsSendPaidMedia.
func (o OptionsSendPaidMedia) SetCaption(caption string) OptionsSendPaidMedia {
	o["caption"] = caption
	return o
}

// SetParseMode sets the `parse_mode` value of OptionsSendPaidMedia.
func (o OptionsSendPaidMedia) SetParseMode(parseMode ParseMode) OptionsSendPaidMedia {
	o["parse_mode"] = parseMode
	return o
}

// SetCaptionEntities sets the `caption_entities` value of OptionsSendPaidMedia.
func (o OptionsSendPaidMedia) SetCaptionEntities(captionEntities []MessageEntity) OptionsSendPaidMedia {
	o["caption_entities"] = captionEntities
	return o
}

// SetShowCaptionAboveMedia sets the `show_caption_above_media` value of OptionsSendPaidMedia.
func (o OptionsSendPaidMedia) SetShowCaptionAboveMedia(showCaptionAboveMedia bool) OptionsSendPaidMedia {
	o["show_caption_above_media"] = showCaptionAboveMedia
	return o
}

// SetDisableNotification sets the `disable_notification` value of OptionsSendPaidMedia.
func (o OptionsSendPaidMedia) SetDisableNotification(disableNotification bool) OptionsSendPaidMedia {
	o["disable_notification"] = disableNotification
	return o
}

// SetProtectContent sets the `protect_content` value of OptionsSendPaidMedia.
func (o OptionsSendPaidMedia) SetProtectContent(protectContent bool) OptionsSendPaidMedia {
	o["protect_content"] = protectContent
	return o
}

// SetAllowPaidBroadcast sets the `allow_paid_broadcast` value of OptionsSendPaidMedia.
func (o OptionsSendPaidMedia) SetAllowPaidBroadcast(allow bool) OptionsSendPaidMedia {
	o["allow_paid_broadcast"] = allow
	return o
}

// SetSuggestedPostParameters sets the `suggested_post_parameters` value of OptionsSendPaidMedia.
func (o OptionsSendPaidMedia) SetSuggestedPostParameters(suggestedPostParameters SuggestedPostParameters) OptionsSendPaidMedia {
	o["suggested_post_parameters"] = suggestedPostParameters
	return o
}

// SetReplyParameters sets the `reply_parameters` value of OptionsSendPaidMedia.
func (o OptionsSendPaidMedia) SetReplyParameters(replyParameters ReplyParameters) OptionsSendPaidMedia {
	o["reply_parameters"] = replyParameters
	return o
}

// SetReplyMarkup sets the `reply_markup` value of OptionsSendPaidMedia.
func (o OptionsSendPaidMedia) SetReplyMarkup(replyMarkup any) OptionsSendPaidMedia {
	o["reply_markup"] = replyMarkup
	return o
}

// OptionsSendMediaGroup struct for SendMediaGroup().
//
// options include: `business_connection_id`, `message_thread_id`, `direct_messages_topic_id`, `disable_notification`, `protect_content`, `allow_paid_broadcast`, `message_effect_id`, and `reply_parameters`.
//
// https://core.telegram.org/bots/api#sendmediagroup
type OptionsSendMediaGroup MethodOptions

// SetBusinessConnectionID sets the `business_connection_id` value of OptionsSendMediaGroup.
func (o OptionsSendMediaGroup) SetBusinessConnectionID(businessConnectionID string) OptionsSendMediaGroup {
	o["business_connection_id"] = businessConnectionID
	return o
}

// SetMessageThreadID sets the `message_thread_id` value of OptionsSendMediaGroup.
func (o OptionsSendMediaGroup) SetMessageThreadID(messageThreadID int64) OptionsSendMediaGroup {
	o["message_thread_id"] = messageThreadID
	return o
}

// SetDirectMessagesTopicID sets the `direct_messages_topic_id` value of OptionsSendMediaGroup.
func (o OptionsSendMediaGroup) SetDirectMessagesTopicID(directMessagesTopicID int64) OptionsSendMediaGroup {
	o["direct_messages_topic_id"] = directMessagesTopicID
	return o
}

// SetDisableNotification sets the `disable_notification` value of OptionsSendMediaGroup.
func (o OptionsSendMediaGroup) SetDisableNotification(disable bool) OptionsSendMediaGroup {
	o["disable_notification"] = disable
	return o
}

// SetProtectContent sets the `protect_content` value of OptionsSendMediaGroup.
func (o OptionsSendMediaGroup) SetProtectContent(protect bool) OptionsSendMediaGroup {
	o["protect_content"] = protect
	return o
}

// SetAllowPaidBroadcast sets the `allow_paid_broadcast` value of OptionsSendMediaGroup.
func (o OptionsSendMediaGroup) SetAllowPaidBroadcast(allow bool) OptionsSendMediaGroup {
	o["allow_paid_broadcast"] = allow
	return o
}

// SetMessageEffectID sets the `message_effect_id` value of OptionsSendMediaGroup.
func (o OptionsSendMediaGroup) SetMessageEffectID(messageEffectID string) OptionsSendMediaGroup {
	o["message_effect_id"] = messageEffectID
	return o
}

// SetReplyParameters sets the `reply_parameters` value of OptionsSendMediaGroup.
func (o OptionsSendMediaGroup) SetReplyParameters(replyParameters ReplyParameters) OptionsSendMediaGroup {
	o["reply_parameters"] = replyParameters
	return o
}

// OptionsSendLocation struct for SendLocation()
//
// options include: `business_connection_id`, `message_thread_id`, `direct_messages_topic_id`, `horizontal_accuracy`, `live_period`, `heading`, `proximity_alert_radius`, `disable_notification`, `protect_content`, `allow_paid_broadcast`, `message_effect_id`, `suggested_post_parameters`, `reply_parameters`, and `reply_markup`.
//
// https://core.telegram.org/bots/api#sendlocation
type OptionsSendLocation MethodOptions

// SetBusinessConnectionID sets the `business_connection_id` value of OptionsSendLocation.
func (o OptionsSendLocation) SetBusinessConnectionID(businessConnectionID string) OptionsSendLocation {
	o["business_connection_id"] = businessConnectionID
	return o
}

// SetMessageThreadID sets the `message_thread_id` value of OptionsSendLocation.
func (o OptionsSendLocation) SetMessageThreadID(messageThreadID int64) OptionsSendLocation {
	o["message_thread_id"] = messageThreadID
	return o
}

// SetDirectMessagesTopicID sets the `direct_messages_topic_id` value of OptionsSendLocation.
func (o OptionsSendLocation) SetDirectMessagesTopicID(directMessagesTopicID int64) OptionsSendLocation {
	o["direct_messages_topic_id"] = directMessagesTopicID
	return o
}

// SetHorizontalAccuracy sets the `horizontal_accuracy` value of OptionsSendLocation.
func (o OptionsSendLocation) SetHorizontalAccuracy(horizontalAccuracy float32) OptionsSendLocation {
	o["horizontal_accuracy"] = horizontalAccuracy
	return o
}

// SetLivePeriod sets the `live_period` value of OptionsSendLocation.
func (o OptionsSendLocation) SetLivePeriod(livePeriod int) OptionsSendLocation {
	o["live_period"] = livePeriod
	return o
}

// SetHeading sets the `heading` value of OptionsSendLocation.
func (o OptionsSendLocation) SetHeading(heading int) OptionsSendLocation {
	o["heading"] = heading
	return o
}

// SetProximityAlertRadius sets the `proximity_alert_radius` value of OptionsSendLocation.
func (o OptionsSendLocation) SetProximityAlertRadius(proximityAlertRadius int) OptionsSendLocation {
	o["proximity_alert_radius"] = proximityAlertRadius
	return o
}

// SetDisableNotification sets the `disable_notification` value of OptionsSendLocation.
func (o OptionsSendLocation) SetDisableNotification(disable bool) OptionsSendLocation {
	o["disable_notification"] = disable
	return o
}

// SetProtectContent sets the `protect_content` value of OptionsSendLocation.
func (o OptionsSendLocation) SetProtectContent(protect bool) OptionsSendLocation {
	o["protect_content"] = protect
	return o
}

// SetAllowPaidBroadcast sets the `allow_paid_broadcast` value of OptionsSendLocation.
func (o OptionsSendLocation) SetAllowPaidBroadcast(allow bool) OptionsSendLocation {
	o["allow_paid_broadcast"] = allow
	return o
}

// SetMessageEffectID sets the `message_effect_id` value of OptionsSendLocation.
func (o OptionsSendLocation) SetMessageEffectID(messageEffectID string) OptionsSendLocation {
	o["message_effect_id"] = messageEffectID
	return o
}

// SetSuggestedPostParameters sets the `suggested_post_parameters` value of OptionsSendLocation.
func (o OptionsSendLocation) SetSuggestedPostParameters(suggestedPostParameters SuggestedPostParameters) OptionsSendLocation {
	o["suggested_post_parameters"] = suggestedPostParameters
	return o
}

// SetReplyParameters sets the `reply_parameters` value of OptionsSendLocation.
func (o OptionsSendLocation) SetReplyParameters(replyParameters ReplyParameters) OptionsSendLocation {
	o["reply_parameters"] = replyParameters
	return o
}

// SetReplyMarkup sets the `reply_markup` value of OptionsSendLocation.
//
// `replyMarkup` can be one of InlineKeyboardMarkup, ReplyKeyboardMarkup, ReplyKeyboardRemove, or ForceReply.
func (o OptionsSendLocation) SetReplyMarkup(replyMarkup any) OptionsSendLocation {
	o["reply_markup"] = replyMarkup
	return o
}

// OptionsSendVenue struct for SendVenue().
//
// options include: `business_connection_id`, `message_thread_id`, `direct_messages_topic_id`, `foursquare_id`, `foursquare_type`, `google_place_id`, `google_place_type`, `disable_notification`, `protect_content`, `allow_paid_broadcast`, `message_effect_id`, `suggested_post_parameters`, `reply_parameters`, and `reply_markup`.
//
// https://core.telegram.org/bots/api#sendvenue
type OptionsSendVenue MethodOptions

// SetBusinessConnectionID sets the `business_connection_id` value of OptionsSendVenue.
func (o OptionsSendVenue) SetBusinessConnectionID(businessConnectionID string) OptionsSendVenue {
	o["business_connection_id"] = businessConnectionID
	return o
}

// SetMessageThreadID sets the `message_thread_id` value of OptionsSendVenue.
func (o OptionsSendVenue) SetMessageThreadID(messageThreadID int64) OptionsSendVenue {
	o["message_thread_id"] = messageThreadID
	return o
}

// SetDirectMessagesTopicID sets the `direct_messages_topic_id` value of OptionsSendVenue.
func (o OptionsSendVenue) SetDirectMessagesTopicID(directMessagesTopicID int64) OptionsSendVenue {
	o["direct_messages_topic_id"] = directMessagesTopicID
	return o
}

// SetFoursquareID sets the `foursquare_id` value of OptionsSendVenue.
func (o OptionsSendVenue) SetFoursquareID(foursquareID string) OptionsSendVenue {
	o["foursquare_id"] = foursquareID
	return o
}

// SetFoursquareType sets the `foursquare_type` value of OptionsSendVenue.
func (o OptionsSendVenue) SetFoursquareType(foursquareType string) OptionsSendVenue {
	o["foursquare_type"] = foursquareType
	return o
}

// SetGooglePlaceID sets the `google_place_id` value of OptionsSendVenue.
func (o OptionsSendVenue) SetGooglePlaceID(googlePlaceID string) OptionsSendVenue {
	o["google_place_id"] = googlePlaceID
	return o
}

// SetGooglePlaceType sets the `google_place_type` value of OptionsSendVenue.
func (o OptionsSendVenue) SetGooglePlaceType(googlePlaceType string) OptionsSendVenue {
	o["google_place_type"] = googlePlaceType
	return o
}

// SetDisableNotification sets the `disable_notification` value of OptionsSendVenue.
func (o OptionsSendVenue) SetDisableNotification(disable bool) OptionsSendVenue {
	o["disable_notification"] = disable
	return o
}

// SetProtectContent sets the `protect_content` value of OptionsSendVenue.
func (o OptionsSendVenue) SetProtectContent(protect bool) OptionsSendVenue {
	o["protect_content"] = protect
	return o
}

// SetAllowPaidBroadcast sets the `allow_paid_broadcast` value of OptionsSendVenue.
func (o OptionsSendVenue) SetAllowPaidBroadcast(allow bool) OptionsSendVenue {
	o["allow_paid_broadcast"] = allow
	return o
}

// SetMessageEffectID sets the `message_effect_id` value of OptionsSendVenue.
func (o OptionsSendVenue) SetMessageEffectID(messageEffectID string) OptionsSendVenue {
	o["message_effect_id"] = messageEffectID
	return o
}

// SetSuggestedPostParameters sets the `suggested_post_parameters` value of OptionsSendVenue.
func (o OptionsSendVenue) SetSuggestedPostParameters(suggestedPostParameters SuggestedPostParameters) OptionsSendVenue {
	o["suggested_post_parameters"] = suggestedPostParameters
	return o
}

// SetReplyParameters sets the `reply_parameters` value of OptionsSendVenue.
func (o OptionsSendVenue) SetReplyParameters(replyParameters ReplyParameters) OptionsSendVenue {
	o["reply_parameters"] = replyParameters
	return o
}

// SetReplyMarkup sets the `reply_markup` value of OptionsSendVenue.
//
// `replyMarkup` can be one of InlineKeyboardMarkup, ReplyKeyboardMarkup, ReplyKeyboardRemove, or ForceReply.
func (o OptionsSendVenue) SetReplyMarkup(replyMarkup any) OptionsSendVenue {
	o["reply_markup"] = replyMarkup
	return o
}

// OptionsSendPoll struct for SendPoll().
//
// options include: `business_connection_id`, `message_thread_id`, `question_parse_mode`, `question_entities`, `is_anonymous`, `type`, `allows_multiple_answers`, `correct_option_id`, `explanation`, `explanation_parse_mode`, `explanation_entities`, `open_period`, `close_date`, `is_closed`, `disable_notification`, `protect_content`, `allow_paid_broadcast`, `message_effect_id`, `reply_parameters`, and `reply_markup`.
//
// https://core.telegram.org/bots/api#sendpoll
type OptionsSendPoll MethodOptions

// SetBusinessConnectionID sets the `business_connection_id` value of OptionsSendPoll.
func (o OptionsSendPoll) SetBusinessConnectionID(businessConnectionID string) OptionsSendPoll {
	o["business_connection_id"] = businessConnectionID
	return o
}

// SetMessageThreadID sets the `message_thread_id` value of OptionsSendPoll.
func (o OptionsSendPoll) SetMessageThreadID(messageThreadID int64) OptionsSendPoll {
	o["message_thread_id"] = messageThreadID
	return o
}

// SetQuestionParseMode sets the `question_parse_mode` value of OptionsSendPoll.
func (o OptionsSendPoll) SetQuestionParseMode(questionParseMode ParseMode) OptionsSendPoll {
	o["question_parse_mode"] = questionParseMode
	return o
}

// SetQuestionEntities sets the `question_entities` value of OptionsSendPoll.
func (o OptionsSendPoll) SetQuestionEntities(questionEntities []MessageEntity) OptionsSendPoll {
	o["question_entities"] = questionEntities
	return o
}

// SetIsAnonymous sets the `is_anonymous` value of OptionsSendPoll.
func (o OptionsSendPoll) SetIsAnonymous(isAnonymous bool) OptionsSendPoll {
	o["is_anonymous"] = isAnonymous
	return o
}

// SetType sets the `type` value of OptionsSendPoll.
func (o OptionsSendPoll) SetType(newType string) OptionsSendPoll {
	o["type"] = newType
	return o
}

// SetAllowsMultipleAnswers sets the `allows_multiple_answers` value of OptionsSendPoll.
func (o OptionsSendPoll) SetAllowsMultipleAnswers(allowsMultipleAnswers bool) OptionsSendPoll {
	o["allows_multiple_answers"] = allowsMultipleAnswers
	return o
}

// SetCorrectOptionID sets the `correct_option_id` value of OptionsSendPoll.
func (o OptionsSendPoll) SetCorrectOptionID(correctOptionID int) OptionsSendPoll {
	o["correct_option_id"] = correctOptionID
	return o
}

// SetExplanation sets the `explanation` value of OptionsSendPoll.
func (o OptionsSendPoll) SetExplanation(explanation string) OptionsSendPoll {
	o["explanation"] = explanation
	return o
}

// SetExplanationParseMode sets the `explanation_parse_mode` value of OptionsSendPoll.
func (o OptionsSendPoll) SetExplanationParseMode(explanationParseMode ParseMode) OptionsSendPoll {
	o["explanation_parse_mode"] = explanationParseMode
	return o
}

// SetExplanationEntities sets the `explanation_entities` value of OptionsSendPoll.
func (o OptionsSendPoll) SetExplanationEntities(entities []MessageEntity) OptionsSendPoll {
	o["explanation_entities"] = entities
	return o
}

// SetOpenPeriod sets the `open_period` value of OptionsSendPoll.
func (o OptionsSendPoll) SetOpenPeriod(openPeriod int) OptionsSendPoll {
	o["open_period"] = openPeriod
	return o
}

// SetCloseDate sets the `close_date` value of OptionsSendPoll.
func (o OptionsSendPoll) SetCloseDate(closeDate int) OptionsSendPoll {
	o["close_date"] = closeDate
	return o
}

// SetIsClosed sets the `is_closed` value of OptionsSendPoll.
func (o OptionsSendPoll) SetIsClosed(isClosed bool) OptionsSendPoll {
	o["is_closed"] = isClosed
	return o
}

// SetDisableNotification sets the `disable_notification` value of OptionsSendPoll.
func (o OptionsSendPoll) SetDisableNotification(disable bool) OptionsSendPoll {
	o["disable_notification"] = disable
	return o
}

// SetProtectContent sets the `protect_content` value of OptionsSendPoll.
func (o OptionsSendPoll) SetProtectContent(protect bool) OptionsSendPoll {
	o["protect_content"] = protect
	return o
}

// SetAllowPaidBroadcast sets the `allow_paid_broadcast` value of OptionsSendPoll.
func (o OptionsSendPoll) SetAllowPaidBroadcast(allow bool) OptionsSendPoll {
	o["allow_paid_broadcast"] = allow
	return o
}

// SetMessageEffectID sets the `message_effect_id` value of OptionsSendPoll.
func (o OptionsSendPoll) SetMessageEffectID(messageEffectID string) OptionsSendPoll {
	o["message_effect_id"] = messageEffectID
	return o
}

// SetReplyParameters sets the `reply_parameters` value of OptionsSendPoll.
func (o OptionsSendPoll) SetReplyParameters(replyParameters ReplyParameters) OptionsSendPoll {
	o["reply_parameters"] = replyParameters
	return o
}

// SetReplyMarkup sets the `reply_markup` value of OptionsSendPoll.
//
// `replyMarkup` can be one of InlineKeyboardMarkup, ReplyKeyboardMarkup, ReplyKeyboardRemove, or ForceReply.
func (o OptionsSendPoll) SetReplyMarkup(replyMarkup any) OptionsSendPoll {
	o["reply_markup"] = replyMarkup
	return o
}

// OptionsStopPoll struct for StopPoll().
//
// options include: `business_connection_id`, and `reply_markup`.
//
// https://core.telegram.org/bots/api#stoppoll
type OptionsStopPoll MethodOptions

// SetBusinessConnectionID sets the `business_connection_id` value of OptionsStopPoll.
func (o OptionsStopPoll) SetBusinessConnectionID(businessConnectionID string) OptionsStopPoll {
	o["business_connection_id"] = businessConnectionID
	return o
}

// SetReplyMarkup sets the `reply_markup` value of OptionsStopPoll.
func (o OptionsStopPoll) SetReplyMarkup(replyMarkup InlineKeyboardMarkup) OptionsStopPoll {
	o["reply_markup"] = replyMarkup
	return o
}

// OptionsApproveSuggestedPost struct for ApproveSuggestedPost().
//
// options include: `send_date`.
//
// https://core.telegram.org/bots/api#approvesuggestedpost
type OptionsApproveSuggestedPost MethodOptions

// SetSendDate sets the `send_date` value of OptionsApproveSuggestedPost.
func (o OptionsApproveSuggestedPost) SetSendDate(sendDate int) OptionsApproveSuggestedPost {
	o["send_date"] = sendDate
	return o
}

// OptionsDeclineSuggestedPost struct for DeclineSuggestedPost().
//
// options include: `send_date`.
//
// https://core.telegram.org/bots/api#declinesuggestedpost
type OptionsDeclineSuggestedPost MethodOptions

// SetSendDate sets the `send_date` value of OptionsApproveSuggestedPost.
func (o OptionsDeclineSuggestedPost) SetSendDate(sendDate int) OptionsDeclineSuggestedPost {
	o["send_date"] = sendDate
	return o
}

// OptionsSendChecklist struct for SendChecklist().
//
// options include: `disable_notification`, `protect_content`, `allow_paid_broadcast`, `message_effect_id`, `reply_parameters`, and `reply_markup`.
//
// https://core.telegram.org/bots/api#sendchecklist
type OptionsSendChecklist MethodOptions

// SetDisableNotification sets the `disable_notification` value of OptionsSendChecklist.
func (o OptionsSendChecklist) SetDisableNotification(disable bool) OptionsSendChecklist {
	o["disable_notification"] = disable
	return o
}

// SetProtectContent sets the `protect_content` value of OptionsSendChecklist.
func (o OptionsSendChecklist) SetProtectContent(protect bool) OptionsSendChecklist {
	o["protect_content"] = protect
	return o
}

// SetMessageEffectID sets the `message_effect_id` value of OptionsSendChecklist.
func (o OptionsSendChecklist) SetMessageEffectID(messageEffectID string) OptionsSendChecklist {
	o["message_effect_id"] = messageEffectID
	return o
}

// SetReplyParameters sets the `reply_parameters` value of OptionsSendChecklist.
func (o OptionsSendChecklist) SetReplyParameters(replyParameters ReplyParameters) OptionsSendChecklist {
	o["reply_parameters"] = replyParameters
	return o
}

// SetReplyMarkup sets the `reply_markup` value of OptionsSendChecklist.
func (o OptionsSendChecklist) SetReplyMarkup(replyMarkup InlineKeyboardMarkup) OptionsSendChecklist {
	o["reply_markup"] = replyMarkup
	return o
}

// OptionsSendDice struct for SendDice().
//
// options include: `business_connection_id`, `message_thread_id`, `direct_messages_topic_id`, `emoji`, `disable_notification`, `protect_content`, `allow_paid_broadcast`, `message_effect_id`, `suggested_post_parameters`, `reply_parameters`, and `reply_markup`.
//
// https://core.telegram.org/bots/api#senddice
type OptionsSendDice MethodOptions

// SetBusinessConnectionID sets the `business_connection_id` value of OptionsSendDice.
func (o OptionsSendDice) SetBusinessConnectionID(businessConnectionID string) OptionsSendDice {
	o["business_connection_id"] = businessConnectionID
	return o
}

// SetMessageThreadID sets the `message_thread_id` value of OptionsSendDice.
func (o OptionsSendDice) SetMessageThreadID(messageThreadID int64) OptionsSendDice {
	o["message_thread_id"] = messageThreadID
	return o
}

// SetDirectMessagesTopicID sets the `direct_messages_topic_id` value of OptionsSendDice.
func (o OptionsSendDice) SetDirectMessagesTopicID(directMessagesTopicID int64) OptionsSendDice {
	o["direct_messages_topic_id"] = directMessagesTopicID
	return o
}

// SetEmoji sets the `emoji` value of OptionsSendDice.
//
// `emoji` can be one of: 🎲 (1~6), 🎯 (1~6), 🎳 (1~6), 🏀 (1~5), ⚽ (1~5), or 🎰 (1~64); default: 🎲
func (o OptionsSendDice) SetEmoji(emoji string) OptionsSendDice {
	o["emoji"] = emoji
	return o
}

// SetDisableNotification sets the `disable_notification` value of OptionsSendDice.
func (o OptionsSendDice) SetDisableNotification(disable bool) OptionsSendDice {
	o["disable_notification"] = disable
	return o
}

// SetProtectContent sets the `protect_content` value of OptionsSendDice.
func (o OptionsSendDice) SetProtectContent(protect bool) OptionsSendDice {
	o["protect_content"] = protect
	return o
}

// SetAllowPaidBroadcast sets the `allow_paid_broadcast` value of OptionsSendDice.
func (o OptionsSendDice) SetAllowPaidBroadcast(allow bool) OptionsSendDice {
	o["allow_paid_broadcast"] = allow
	return o
}

// SetMessageEffectID sets the `message_effect_id` value of OptionsSendDice.
func (o OptionsSendDice) SetMessageEffectID(messageEffectID string) OptionsSendDice {
	o["message_effect_id"] = messageEffectID
	return o
}

// SetSuggestedPostParameters sets the `suggested_post_parameters` value of OptionsSendDice.
func (o OptionsSendDice) SetSuggestedPostParameters(suggestedPostParameters SuggestedPostParameters) OptionsSendDice {
	o["suggested_post_parameters"] = suggestedPostParameters
	return o
}

// SetReplyParameters sets the `reply_parameters` value of OptionsSendDice.
func (o OptionsSendDice) SetReplyParameters(replyParameters ReplyParameters) OptionsSendDice {
	o["reply_parameters"] = replyParameters
	return o
}

// SetReplyMarkup sets the `reply_markup` value of OptionsSendDice.
//
// `replyMarkup` can be one of InlineKeyboardMarkup, ReplyKeyboardMarkup, ReplyKeyboardRemove, or ForceReply.
func (o OptionsSendDice) SetReplyMarkup(replyMarkup any) OptionsSendDice {
	o["reply_markup"] = replyMarkup
	return o
}

// OptionsSendMessageDraft struct for SendMessageDraft().
//
// options include: `message_thread_id`, `parse_mode`, and `entities`.
//
// https://core.telegram.org/bots/api#sendmessagedraft
type OptionsSendMessageDraft MethodOptions

// SetMessageThreadID sets the `message_thread_id` value of OptionsSendMessageDraft.
func (o OptionsSendMessageDraft) SetMessageThreadID(messageThreadID int64) OptionsSendMessageDraft {
	o["message_thread_id"] = messageThreadID
	return o
}

// SetParseMode sets the `parse_mode` value of OptionsSendMessageDraft.
func (o OptionsSendMessageDraft) SetParseMode(parseMode ParseMode) OptionsSendMessageDraft {
	o["parse_mode"] = parseMode
	return o
}

// SetEntities sets the `entities` value of OptionsSendMessageDraft.
func (o OptionsSendMessageDraft) SetEntities(entities []MessageEntity) OptionsSendMessageDraft {
	o["entities"] = entities
	return o
}

// OptionsSendChatAction struct for SendChatAction().
//
// options include: `business_connection_id`, and `message_thread_id`.
//
// https://core.telegram.org/bots/api#sendchataction
type OptionsSendChatAction MethodOptions

// SetBusinessConnectionID sets the `business_connection_id` value of OptionsSendChatAction.
func (o OptionsSendChatAction) SetBusinessConnectionID(businessConnectionID string) OptionsSendChatAction {
	o["business_connection_id"] = businessConnectionID
	return o
}

// SetMessageThreadID sets the `message_thread_id` value of OptionsSendChatAction.
func (o OptionsSendChatAction) SetMessageThreadID(messageThreadID int64) OptionsSendChatAction {
	o["message_thread_id"] = messageThreadID
	return o
}

// OptionsSetMessageReaction struct for SetMessageReaction().
//
// options include: `reaction`, and `is_big`.
//
// https://core.telegram.org/bots/api#setmessagereaction
type OptionsSetMessageReaction MethodOptions

// SetReaction sets the `reaction` value of OptionsSetMessageReaction.
func (o OptionsSetMessageReaction) SetReaction(reactions []ReactionType) OptionsSetMessageReaction {
	o["reaction"] = reactions
	return o
}

// SetIsBig sets the `is_big` value of OptionsSetMessageReaction.
func (o OptionsSetMessageReaction) SetIsBig(isBig bool) OptionsSetMessageReaction {
	o["is_big"] = isBig
	return o
}

// OptionsSendContact struct for SendContact().
//
// options include: `business_connection_id`, `message_thread_id`, `direct_messages_topic_id`, `last_name`, `vcard`, `disable_notification`, `protect_content`, `allow_paid_broadcast`, `message_effect_id`, `suggested_post_parameters`, `reply_parameters`, and `reply_markup`.
//
// https://core.telegram.org/bots/api#sendcontact
type OptionsSendContact MethodOptions

// SetBusinessConnectionID sets the `business_connection_id` value of OptionsSendContact.
func (o OptionsSendContact) SetBusinessConnectionID(businessConnectionID string) OptionsSendContact {
	o["business_connection_id"] = businessConnectionID
	return o
}

// SetMessageThreadID sets the `message_thread_id` value of OptionsSendContact.
func (o OptionsSendContact) SetMessageThreadID(messageThreadID int64) OptionsSendContact {
	o["message_thread_id"] = messageThreadID
	return o
}

// SetDirectMessagesTopicID sets the `direct_messages_topic_id` value of OptionsSendContact.
func (o OptionsSendContact) SetDirectMessagesTopicID(directMessagesTopicID int64) OptionsSendContact {
	o["direct_messages_topic_id"] = directMessagesTopicID
	return o
}

// SetLastName sets the `last_name` value of OptionsSendContact.
func (o OptionsSendContact) SetLastName(lastName string) OptionsSendContact {
	o["last_name"] = lastName
	return o
}

// SetVCard sets the `vcard` value of OptionsSendContact.
func (o OptionsSendContact) SetVCard(vCard string) OptionsSendContact {
	o["vcard"] = vCard
	return o
}

// SetDisableNotification sets the `disable_notification` value of OptionsSendContact.
func (o OptionsSendContact) SetDisableNotification(disable bool) OptionsSendContact {
	o["disable_notification"] = disable
	return o
}

// SetProtectContent sets the `protect_content` value of OptionsSendContact.
func (o OptionsSendContact) SetProtectContent(protect bool) OptionsSendContact {
	o["protect_content"] = protect
	return o
}

// SetAllowPaidBroadcast sets the `allow_paid_broadcast` value of OptionsSendContact.
func (o OptionsSendContact) SetAllowPaidBroadcast(allow bool) OptionsSendContact {
	o["allow_paid_broadcast"] = allow
	return o
}

// SetMessageEffectID sets the `message_effect_id` value of OptionsSendContact.
func (o OptionsSendContact) SetMessageEffectID(messageEffectID string) OptionsSendContact {
	o["message_effect_id"] = messageEffectID
	return o
}

// SetSuggestedPostParameters sets the `suggested_post_parameters` value of OptionsSendContact.
func (o OptionsSendContact) SetSuggestedPostParameters(suggestedPostParameters SuggestedPostParameters) OptionsSendContact {
	o["suggested_post_parameters"] = suggestedPostParameters
	return o
}

// SetReplyParameters sets the `reply_parameters` value of OptionsSendContact.
func (o OptionsSendContact) SetReplyParameters(replyParameters ReplyParameters) OptionsSendContact {
	o["reply_parameters"] = replyParameters
	return o
}

// SetReplyMarkup sets the `reply_markup` value of OptionsSendContact.
//
// `replyMarkup` can be one of InlineKeyboardMarkup, ReplyKeyboardMarkup, ReplyKeyboardRemove, or ForceReply.
func (o OptionsSendContact) SetReplyMarkup(replyMarkup any) OptionsSendContact {
	o["reply_markup"] = replyMarkup
	return o
}

// OptionsGetUserProfilePhotos struct for GetUserProfilePhotos().
//
// options include: `offset` and `limit`.
//
// https://core.telegram.org/bots/api#getuserprofilephotos
type OptionsGetUserProfilePhotos MethodOptions

// SetOffset sets the `offset` value of OptionsGetUserProfilePhotos.
func (o OptionsGetUserProfilePhotos) SetOffset(offset int) OptionsGetUserProfilePhotos {
	o["offset"] = offset
	return o
}

// SetLimit sets the `limit` value of OptionsGetUserProfilePhotos.
func (o OptionsGetUserProfilePhotos) SetLimit(limit int) OptionsGetUserProfilePhotos {
	o["limit"] = limit
	return o
}

// OptionsGetUserProfileAudios struct for GetUserProfileAudios().
//
// options include: `offset` and `limit`.
//
// https://core.telegram.org/bots/api#getuserprofileaudios
type OptionsGetUserProfileAudios MethodOptions

// SetOffset sets the `offset` value of OptionsGetUserProfileAudios.
func (o OptionsGetUserProfileAudios) SetOffset(offset int) OptionsGetUserProfileAudios {
	o["offset"] = offset
	return o
}

// SetLimit sets the `limit` value of OptionsGetUserProfileAudios.
func (o OptionsGetUserProfileAudios) SetLimit(limit int) OptionsGetUserProfileAudios {
	o["limit"] = limit
	return o
}

// OptionsSetUserEmojiStatus struct for SetUserEmojiStatus().
//
// options include: `emoji_status_custom_emoji_id`, and `emoji_status_expiration_date`.
type OptionsSetUserEmojiStatus MethodOptions

// SetEmojiStatusCustomEmojiID sets the `emoji_status_custom_emoji_id` value of OptionsSetUserEmojiStatus.
func (o OptionsSetUserEmojiStatus) SetEmojiStatusCustomEmojiID(customEmojiID string) OptionsSetUserEmojiStatus {
	o["emoji_status_custom_emoji_id"] = customEmojiID
	return o
}

// SetEmojiStatusExpirationDate sets the `emoji_status_expiration_date` value of OptionsSetUserEmojiStatus.
func (o OptionsSetUserEmojiStatus) SetEmojiStatusExpirationDate(expirationDate int) OptionsSetUserEmojiStatus {
	o["emoji_status_expiration_date"] = expirationDate
	return o
}

// OptionsBanChatMember struct for BanChatMember().
//
// options include: `until_date` and `revoke_messages`.
//
// https://core.telegram.org/bots/api#banchatmember
type OptionsBanChatMember MethodOptions

// SetUntilDate sets the `until_date` value of OptionsBanChatMember.
func (o OptionsBanChatMember) SetUntilDate(untilDate int) OptionsBanChatMember {
	o["until_date"] = untilDate
	return o
}

// SetRevokeMessages sets the `revoke_messages` value of OptionsBanChatMember.
func (o OptionsBanChatMember) SetRevokeMessages(revokeMessages bool) OptionsBanChatMember {
	o["revoke_messages"] = revokeMessages
	return o
}

// OptionsRestrictChatMember struct for RestrictChatMember().
//
// options include: `use_independent_chat_permissions`, and `until_date`
//
// https://core.telegram.org/bots/api#restrictchatmember
type OptionsRestrictChatMember MethodOptions

// SetUserIndependentChatPermissions sets the `use_independent_chat_permissions` value of OptionsRestrictChatMember.
func (o OptionsRestrictChatMember) SetUserIndependentChatPermissions(val bool) OptionsRestrictChatMember {
	o["use_independent_chat_permissions"] = val
	return o
}

// SetUntilDate sets the `until_date` value of OptionsRestrictChatMember.
func (o OptionsRestrictChatMember) SetUntilDate(until int) OptionsRestrictChatMember {
	o["until_date"] = until
	return o
}

// OptionsPromoteChatMember struct for PromoteChatMember().
//
// options include: `is_anonymous`, `can_manage_chat`, `can_post_messages`, `can_edit_messages`, `can_delete_messages`, `can_manage_video_chats`, `can_restrict_members`, `can_promote_members`, `can_change_info`, `can_invite_users`, `can_pin_messages`, `can_manage_topics`, and `can_manage_direct_messages`.
//
// https://core.telegram.org/bots/api#promotechatmember
type OptionsPromoteChatMember MethodOptions

// SetIsAnonymous sets the `is_anonymous` value of OptionsPromoteChatMember.
func (o OptionsPromoteChatMember) SetIsAnonymous(anonymous bool) OptionsPromoteChatMember {
	o["is_anonymous"] = anonymous
	return o
}

// SetCanChangeInfo sets the `can_change_info` value of OptionsPromoteChatMember.
func (o OptionsPromoteChatMember) SetCanChangeInfo(can bool) OptionsPromoteChatMember {
	o["can_change_info"] = can
	return o
}

// SetCanManageChat sets the `can_manage_chat` value of OptionsPromoteChatMember.
func (o OptionsPromoteChatMember) SetCanManageChat(can bool) OptionsPromoteChatMember {
	o["can_manage_chat"] = can
	return o
}

// SetCanPostMessages sets the `can_post_messages` value of OptionsPromoteChatMember.
func (o OptionsPromoteChatMember) SetCanPostMessages(can bool) OptionsPromoteChatMember {
	o["can_post_messages"] = can
	return o
}

// SetCanEditMessages sets the `can_edit_messages` value of OptionsPromoteChatMember.
func (o OptionsPromoteChatMember) SetCanEditMessages(can bool) OptionsPromoteChatMember {
	o["can_edit_messages"] = can
	return o
}

// SetCanDeleteMessages sets the `can_delete_messages` value of OptionsPromoteChatMember.
func (o OptionsPromoteChatMember) SetCanDeleteMessages(can bool) OptionsPromoteChatMember {
	o["can_delete_messages"] = can
	return o
}

// SetCanPostStories sets the `can_post_stories` value of OptionsPromoteChatMember.
func (o OptionsPromoteChatMember) SetCanPostStories(can bool) OptionsPromoteChatMember {
	o["can_post_stories"] = can
	return o
}

// SetCanEditStories sets the `can_edit_stories` value of OptionsPromoteChatMember.
func (o OptionsPromoteChatMember) SetCanEditStories(can bool) OptionsPromoteChatMember {
	o["can_edit_stories"] = can
	return o
}

// SetCanDeleteStories sets the `can_delete_stories` value of OptionsPromoteChatMember.
func (o OptionsPromoteChatMember) SetCanDeleteStories(can bool) OptionsPromoteChatMember {
	o["can_delete_stories"] = can
	return o
}

// SetCanManageVideoChats sets the `can_manage_video_chats` value of OptionsPromoteChatMember.
func (o OptionsPromoteChatMember) SetCanManageVideoChats(can bool) OptionsPromoteChatMember {
	o["can_manage_video_chats"] = can
	return o
}

// SetCanInviteUsers sets the `can_invite_users` value of OptionsPromoteChatMember.
func (o OptionsPromoteChatMember) SetCanInviteUsers(can bool) OptionsPromoteChatMember {
	o["can_invite_users"] = can
	return o
}

// SetCanRestrictMembers sets the `can_restrict_members` value of OptionsPromoteChatMember.
func (o OptionsPromoteChatMember) SetCanRestrictMembers(can bool) OptionsPromoteChatMember {
	o["can_restrict_members"] = can
	return o
}

// SetCanPinMessages sets the `can_pin_messages` value of OptionsPromoteChatMember.
func (o OptionsPromoteChatMember) SetCanPinMessages(can bool) OptionsPromoteChatMember {
	o["can_pin_messages"] = can
	return o
}

// SetCanPromoteMembers sets the `can_promote_members` value of OptionsPromoteChatMember.
func (o OptionsPromoteChatMember) SetCanPromoteMembers(can bool) OptionsPromoteChatMember {
	o["can_promote_members"] = can
	return o
}

// SetCanManageTopics sets the `can_manage_topics` value of OptionsPromoteChatMember.
func (o OptionsPromoteChatMember) SetCanManageTopics(can bool) OptionsPromoteChatMember {
	o["can_manage_topics"] = can
	return o
}

// SetCanManageDirectMessages sets the `can_manage_direct_messages` value of OptionsPromoteChatMember.
func (o OptionsPromoteChatMember) SetCanManageDirectMessages(can bool) OptionsPromoteChatMember {
	o["can_manage_direct_messages"] = can
	return o
}

// OptionsSetChatPermissions struct for SetChatPermissions
//
// options include: `use_independent_chat_permissions`.
//
// https://core.telegram.org/bots/api#setchatpermissions
type OptionsSetChatPermissions MethodOptions

// SetUserIndependentChatPermissions sets the `use_independent_chat_permissions` value of OptionsRestrictChatMember.
func (o OptionsSetChatPermissions) SetUserIndependentChatPermissions(val bool) OptionsSetChatPermissions {
	o["use_independent_chat_permissions"] = val
	return o
}

// OptionsCreateChatInviteLink struct for CreateChatInviteLink
//
// options include: `name`, `expire_date`, `member_limit`, and `creates_join_request`.
//
// https://core.telegram.org/bots/api#createchatinvitelink
type OptionsCreateChatInviteLink MethodOptions

// SetName sets the `name` value of OptionsCreateChatInviteLink.
func (o OptionsCreateChatInviteLink) SetName(name string) OptionsCreateChatInviteLink {
	o["name"] = name
	return o
}

// SetExpireDate sets the `expire_date` value of OptionsCreateChatInviteLink.
func (o OptionsCreateChatInviteLink) SetExpireDate(expireDate int) OptionsCreateChatInviteLink {
	o["expire_date"] = expireDate
	return o
}

// SetMemberLimit sets the `member_limit` value of OptionsCreateChatInviteLink.
func (o OptionsCreateChatInviteLink) SetMemberLimit(memberLimit int) OptionsCreateChatInviteLink {
	o["member_limit"] = memberLimit
	return o
}

// SetCreatesJoinRequest sets the `creates_join_request` value of OptionsCreateChatInviteLink.
func (o OptionsCreateChatInviteLink) SetCreatesJoinRequest(createsJoinRequest bool) OptionsCreateChatInviteLink {
	o["creates_join_request"] = createsJoinRequest
	return o
}

// OptionsCreateChatSubscriptionInviteLink struct for CreateChatSubscriptionInviteLink
//
// options include: `name`.
//
// https://core.telegram.org/bots/api#createchatsubscriptioninvitelink
type OptionsCreateChatSubscriptionInviteLink MethodOptions

// SetName sets the `name` value of OptionsCreateChatSubscriptionInviteLink.
func (o OptionsCreateChatSubscriptionInviteLink) SetName(name string) OptionsCreateChatSubscriptionInviteLink {
	o["name"] = name
	return o
}

// OptionsEditChatSubscriptionInviteLink struct for EditChatSubscriptionInviteLink
//
// https://core.telegram.org/bots/api#editchatsubscriptioninvitelink
type OptionsEditChatSubscriptionInviteLink MethodOptions

// SetName sets the `name` value of OptionsEditChatSubscriptionInviteLink.
func (o OptionsEditChatSubscriptionInviteLink) SetName(name string) OptionsEditChatSubscriptionInviteLink {
	o["name"] = name
	return o
}

// OptionsPinChatMessage struct for PinChatMessage
//
// options include: `business_connection_id`, and `disable_notification`.
//
// https://core.telegram.org/bots/api#pinchatmessage
type OptionsPinChatMessage MethodOptions

// SetBusinessConnectionID sets the `business_connection_id` value of OptionsPinChatMessage.
func (o OptionsPinChatMessage) SetBusinessConnectionID(businessConnectionID string) OptionsPinChatMessage {
	o["business_connection_id"] = businessConnectionID
	return o
}

// SetDisableNotification sets the `disable_notification` value of OptionsPinChatMessage.
func (o OptionsPinChatMessage) SetDisableNotification(disable bool) OptionsPinChatMessage {
	o["disable_notification"] = disable
	return o
}

// OptionsUnpinChatMessage struct for UnpinChatMessage
//
// options include: `business_connection_id`, and `message_id`.
//
// https://core.telegram.org/bots/api#unpinchatmessage
type OptionsUnpinChatMessage MethodOptions

// SetBusinessConnectionID sets the `business_connection_id` value of OptionsUnpinChatMessage.
func (o OptionsUnpinChatMessage) SetBusinessConnectionID(businessConnectionID string) OptionsUnpinChatMessage {
	o["business_connection_id"] = businessConnectionID
	return o
}

// SetMessageID set the `message_id` value of OptionsUnpinChatMessage.
func (o OptionsUnpinChatMessage) SetMessageID(messageID int64) OptionsUnpinChatMessage {
	o["message_id"] = messageID
	return o
}

// OptionsAnswerCallbackQuery struct for AnswerCallbackQuery().
//
// options include: `text`, `show_alert`, `url`, and `cache_time`.
//
// https://core.telegram.org/bots/api#answercallbackquery
type OptionsAnswerCallbackQuery MethodOptions

// SetText sets the `text` value of OptionsAnswerCallbackQuery.
func (o OptionsAnswerCallbackQuery) SetText(text string) OptionsAnswerCallbackQuery {
	o["text"] = text
	return o
}

func (o OptionsAnswerCallbackQuery) SetShowAlert(showAlert bool) OptionsAnswerCallbackQuery {
	o["show_alert"] = showAlert
	return o
}

// SetURL sets the `url` value of OptionsAnswerCallbackQuery.
func (o OptionsAnswerCallbackQuery) SetURL(url string) OptionsAnswerCallbackQuery {
	o["url"] = url
	return o
}

// SetCacheTime sets the `cache_time` value of OptionsAnswerCallbackQuery.
func (o OptionsAnswerCallbackQuery) SetCacheTime(cacheTime int) OptionsAnswerCallbackQuery {
	o["cache_time"] = cacheTime
	return o
}

// OptionsGetMyCommands struct for GetMyCommands().
//
// options include: `scope`, and `language_code`.
//
// https://core.telegram.org/bots/api#getmycommands
type OptionsGetMyCommands MethodOptions

// SetScope sets the `scope` value of OptionsGetMyCommands.
//
// `scope` can be one of: BotCommandScopeDefault, BotCommandScopeAllPrivateChats, BotCommandScopeAllGroupChats, BotCommandScopeAllChatAdministrators, BotCommandScopeChat, BotCommandScopeChatAdministrators, or BotCommandScopeChatMember.
func (o OptionsGetMyCommands) SetScope(scope any) OptionsGetMyCommands {
	o["scope"] = scope
	return o
}

// SetLanguageCode sets the `language_code` value of OptionsGetMyCommands.
//
// `language_code` is a two-letter ISO 639-1 language code and can be empty.
func (o OptionsGetMyCommands) SetLanguageCode(languageCode string) OptionsGetMyCommands {
	o["language_code"] = languageCode
	return o
}

// OptionsSetMyCommands struct for SetMyCommands().
//
// options include: `scope`, and `language_code`
//
// https://core.telegram.org/bots/api#setmycommands
type OptionsSetMyCommands MethodOptions

// SetScope sets the `scope` value of OptionsSetMyCommands.
//
// `scope` can be one of: BotCommandScopeDefault, BotCommandScopeAllPrivateChats, BotCommandScopeAllGroupChats, BotCommandScopeAllChatAdministrators, BotCommandScopeChat, BotCommandScopeChatAdministrators, or BotCommandScopeChatMember.
func (o OptionsSetMyCommands) SetScope(scope any) OptionsSetMyCommands {
	o["scope"] = scope
	return o
}

// SetLanguageCode sets the `language_code` value of OptionsSetMyCommands.
//
// `language_code` is a two-letter ISO 639-1 language code and can be empty.
func (o OptionsSetMyCommands) SetLanguageCode(languageCode string) OptionsSetMyCommands {
	o["language_code"] = languageCode
	return o
}

// OptionsDeleteMyCommands struct for DeleteMyCommands().
//
// options include: `scope`, and `language_code`.
//
// https://core.telegram.org/bots/api#deletemycommands
type OptionsDeleteMyCommands MethodOptions

// SetScope sets the `scope` value of OptionsDeleteMyCommands.
//
// `scope` can be one of: BotCommandScopeDefault, BotCommandScopeAllPrivateChats, BotCommandScopeAllGroupChats, BotCommandScopeAllChatAdministrators, BotCommandScopeChat, BotCommandScopeChatAdministrators, or BotCommandScopeChatMember.
func (o OptionsDeleteMyCommands) SetScope(scope any) OptionsDeleteMyCommands {
	o["scope"] = scope
	return o
}

// SetLanguageCode sets the `language_code` value of OptionsDeleteMyCommands.
//
// `language_code` is a two-letter ISO 639-1 language code and can be empty.
func (o OptionsDeleteMyCommands) SetLanguageCode(languageCode string) OptionsDeleteMyCommands {
	o["language_code"] = languageCode
	return o
}

// OptionsSetMyName struct for SetMyName().
//
// options include: `language_code`.
//
// https://core.telegram.org/bots/api#setmyname
type OptionsSetMyName MethodOptions

// SetLanguageCode sets the `language_code` value of OptionsSetMyName.
//
// `language_code` is a two-letter ISO 639-1 language code and can be empty.
func (o OptionsSetMyName) SetLanguageCode(languageCode string) OptionsSetMyName {
	o["language_code"] = languageCode
	return o
}

// OptionsGetMyName struct for GetMyName().
//
// options include: `language_code`.
//
// https://core.telegram.org/bots/api#getmyname
type OptionsGetMyName MethodOptions

// SetLanguageCode sets the `language_code` value of OptionsGetMyName.
//
// `language_code` is a two-letter ISO 639-1 language code and can be empty.
func (o OptionsGetMyName) SetLanguageCode(languageCode string) OptionsGetMyName {
	o["language_code"] = languageCode
	return o
}

// OptionsSetMyDescription struct for SetMyDescription().
//
// options include: `description`, and `language_code`.
//
// https://core.telegram.org/bots/api#setmydescription
type OptionsSetMyDescription MethodOptions

// SetDescription sets the `description` value of OptionsSetMyDescription.
func (o OptionsSetMyDescription) SetDescription(description string) OptionsSetMyDescription {
	o["description"] = description
	return o
}

// SetLanguageCode sets the `language_code` value of OptionsSetMyDescription.
//
// `language_code` is a two-letter ISO 639-1 language code and can be empty.
func (o OptionsSetMyDescription) SetLanguageCode(languageCode string) OptionsSetMyDescription {
	o["language_code"] = languageCode
	return o
}

// OptionsGetMyDescription struct for GetMyDescription().
//
// options include: `language_code`.
//
// https://core.telegram.org/bots/api#getmydescription
type OptionsGetMyDescription MethodOptions

// SetLanguageCode sets the `language_code` value of OptionsGetMyDescription.
//
// `language_code` is a two-letter ISO 639-1 language code and can be empty.
func (o OptionsGetMyDescription) SetLanguageCode(languageCode string) OptionsGetMyDescription {
	o["language_code"] = languageCode
	return o
}

// OptionsSetMyShortDescription struct for SetMyShortDescription().
//
// options include: `short_description`, and `language_code`.
//
// https://core.telegram.org/bots/api#setmyshortdescription
type OptionsSetMyShortDescription MethodOptions

// SetShortDescription sets the `short_description` value of OptionsSetMyShortDescription.
func (o OptionsSetMyShortDescription) SetShortDescription(shortDescription string) OptionsSetMyShortDescription {
	o["short_description"] = shortDescription
	return o
}

// SetLanguageCode sets the `language_code` value of OptionsSetMyShortDescription.
//
// `language_code` is a two-letter ISO 639-1 language code and can be empty.
func (o OptionsSetMyShortDescription) SetLanguageCode(languageCode string) OptionsSetMyShortDescription {
	o["language_code"] = languageCode
	return o
}

// OptionsGetMyShortDescription struct for GetMyShortDescription().
//
// options include: `language_code`.
//
// https://core.telegram.org/bots/api#getmyshortdescription
type OptionsGetMyShortDescription MethodOptions

// SetLanguageCode sets the `language_code` value of OptionsGetMyShortDescription.
//
// `language_code` is a two-letter ISO 639-1 language code and can be empty.
func (o OptionsGetMyShortDescription) SetLanguageCode(languageCode string) OptionsGetMyShortDescription {
	o["language_code"] = languageCode
	return o
}

// OptionsSetChatMenuButton struct for SetChatMenuButton().
//
// options include: `chat_id`, and `menu_button`.
//
// https://core.telegram.org/bots/api#setchatmenubutton
type OptionsSetChatMenuButton MethodOptions

// SetChatID sets the `chat_id` value of OptionsSetChatMenuButton.
func (o OptionsSetChatMenuButton) SetChatID(chatID ChatID) OptionsSetChatMenuButton {
	o["chat_id"] = chatID
	return o
}

// SetMenuButton sets the `menu_button` value of OptionsSetChatMenuButton.
func (o OptionsSetChatMenuButton) SetMenuButton(menuButton MenuButton) OptionsSetChatMenuButton {
	o["menu_button"] = menuButton
	return o
}

// OptionsGetChatMenuButton struct for GetChatMenuButton().
//
// options include: `chat_id`.
//
// https://core.telegram.org/bots/api#getchatmenubutton
type OptionsGetChatMenuButton MethodOptions

// SetChatID sets the `chat_id` value of OptionsGetChatMenuButton.
func (o OptionsGetChatMenuButton) SetChatID(chatID ChatID) OptionsGetChatMenuButton {
	o["chat_id"] = chatID
	return o
}

// OptionsSetMyDefaultAdministratorRights struct for SetMyDefaultAdministratorRights().
//
// options include: `rights`, and `for_channels`.
//
// https://core.telegram.org/bots/api#setmydefaultadministratorrights
type OptionsSetMyDefaultAdministratorRights MethodOptions

// SetRights sets the `rights` value of OptionsSetMyDefaultAdministratorRights.
func (o OptionsSetMyDefaultAdministratorRights) SetRights(rights ChatAdministratorRights) OptionsSetMyDefaultAdministratorRights {
	o["rights"] = rights
	return o
}

// SetForChannels sets the `for_channels` value of OptionsSetMyDefaultAdministratorRights.
func (o OptionsSetMyDefaultAdministratorRights) SetForChannels(forChannels bool) OptionsSetMyDefaultAdministratorRights {
	o["for_channels"] = forChannels
	return o
}

// OptionsGetMyDefaultAdministratorRights struct for GetMyDefaultAdministratorRights().
//
// options include: `for_channels`.
//
// https://core.telegram.org/bots/api#getmydefaultadministratorrights
type OptionsGetMyDefaultAdministratorRights MethodOptions

// SetForChannels sets the `for_channels` value of OptionsGetMyDefaultAdministratorRights.
func (o OptionsGetMyDefaultAdministratorRights) SetForChannels(forChannels bool) OptionsGetMyDefaultAdministratorRights {
	o["for_channels"] = forChannels
	return o
}

// OptionsEditMessageText struct for EditMessageText().
//
// required options: `chat_id` + `message_id` (when `inline_message_id` is not given)
//
//	or `inline_message_id` (when `chat_id` & `message_id` is not given)
//
// other options: `business_connection_id`, `parse_mode`, `entities`, `link_preview_options`, and `reply_markup`
//
// https://core.telegram.org/bots/api#editmessagetext
type OptionsEditMessageText MethodOptions

// SetBusinessConnectionID sets the `business_connection_id` value of OptionsEditMessageText.
func (o OptionsEditMessageText) SetBusinessConnectionID(businessConnectionID string) OptionsEditMessageText {
	o["business_connection_id"] = businessConnectionID
	return o
}

// SetIDs sets the `chat_id` and `message_id` values of OptionsEditMessageText.
func (o OptionsEditMessageText) SetIDs(chatID ChatID, messageID int64) OptionsEditMessageText {
	o["chat_id"] = chatID
	o["message_id"] = messageID
	return o
}

// SetInlineMessageID sets the `inline_message_id` value of OptionsEditMessageText.
func (o OptionsEditMessageText) SetInlineMessageID(inlineMessageID string) OptionsEditMessageText {
	o["inline_message_id"] = inlineMessageID
	return o
}

// SetParseMode sets the `parse_mode` value of OptionsEditMessageText.
func (o OptionsEditMessageText) SetParseMode(parseMode ParseMode) OptionsEditMessageText {
	o["parse_mode"] = parseMode
	return o
}

// SetEntities sets the `entities` value of OptionsEditMessageText.
func (o OptionsEditMessageText) SetEntities(entities []MessageEntity) OptionsEditMessageText {
	o["entities"] = entities
	return o
}

// SetLinkPreviewOptions sets the `link_preview_options` value of OptionsEditMessageText.
func (o OptionsEditMessageText) SetLinkPreviewOptions(linkPreviewOptions LinkPreviewOptions) OptionsEditMessageText {
	o["link_preview_options"] = linkPreviewOptions
	return o
}

// SetReplyMarkup sets the `reply_markup` value of OptionsEditMessageText.
func (o OptionsEditMessageText) SetReplyMarkup(replyMarkup InlineKeyboardMarkup) OptionsEditMessageText {
	o["reply_markup"] = replyMarkup
	return o
}

// OptionsEditMessageCaption struct for EditMessageCaption().
//
// required options: `chat_id` + `message_id` (when `inline_message_id` is not given)
//
//	or `inline_message_id` (when `chat_id` & `message_id` is not given)
//
// other options: `business_connection_id`, `caption`, `parse_mode`, `caption_entities`, `show_caption_above_media`, or `reply_markup`
//
// https://core.telegram.org/bots/api#editmessagecaption
type OptionsEditMessageCaption MethodOptions

// SetBusinessConnectionID sets the `business_connection_id` value of OptionsEditMessageCaption.
func (o OptionsEditMessageCaption) SetBusinessConnectionID(businessConnectionID string) OptionsEditMessageCaption {
	o["business_connection_id"] = businessConnectionID
	return o
}

// SetIDs sets the `chat_id` and `message_id` values of OptionsEditMessageCaption.
func (o OptionsEditMessageCaption) SetIDs(chatID ChatID, messageID int64) OptionsEditMessageCaption {
	o["chat_id"] = chatID
	o["message_id"] = messageID
	return o
}

// SetInlineMessageID sets the `inline_message_id` value of OptionsEditMessageCaption.
func (o OptionsEditMessageCaption) SetInlineMessageID(inlineMessageID string) OptionsEditMessageCaption {
	o["inline_message_id"] = inlineMessageID
	return o
}

// SetCaption sets the `caption` value of OptionsEditMessageCaption.
func (o OptionsEditMessageCaption) SetCaption(caption string) OptionsEditMessageCaption {
	o["caption"] = caption
	return o
}

// SetParseMode sets the `parse_mode` value of OptionsEditMessageCaption.
func (o OptionsEditMessageCaption) SetParseMode(parseMode ParseMode) OptionsEditMessageCaption {
	o["parse_mode"] = parseMode
	return o
}

// SetCaptionEntities sets the `caption_entities` value of OptionsEditMessageCaption.
func (o OptionsEditMessageCaption) SetCaptionEntities(entities []MessageEntity) OptionsEditMessageCaption {
	o["caption_entities"] = entities
	return o
}

// SetShowCaptionAboveMedia sets the `show_caption_above_media` value of OptionsEditMessageCaption.
func (o OptionsEditMessageCaption) SetShowCaptionAboveMedia(showCaptionAboveMedia bool) OptionsEditMessageCaption {
	o["show_caption_above_media"] = showCaptionAboveMedia
	return o
}

// SetReplyMarkup sets the `reply_markup` value of OptionsEditMessageCaption.
func (o OptionsEditMessageCaption) SetReplyMarkup(replyMarkup InlineKeyboardMarkup) OptionsEditMessageCaption {
	o["reply_markup"] = replyMarkup
	return o
}

// OptionsEditMessageMedia struct for EditMessageMedia()
//
// required options: `chat_id` + `message_id` (when `inline_message_id` is not given)
//
//	or `inline_message_id` (when `chat_id` & `message_id` is not given)
//
// other options: `business_connection_id`, and `reply_markup`
//
// https://core.telegram.org/bots/api#editmessagemedia
type OptionsEditMessageMedia MethodOptions

// SetBusinessConnectionID sets the `business_connection_id` value of OptionsEditMessageMedia.
func (o OptionsEditMessageMedia) SetBusinessConnectionID(businessConnectionID string) OptionsEditMessageMedia {
	o["business_connection_id"] = businessConnectionID
	return o
}

// SetIDs sets the `chat_id` and `message_id` values of OptionsEditMessageMedia.
func (o OptionsEditMessageMedia) SetIDs(chatID ChatID, messageID int64) OptionsEditMessageMedia {
	o["chat_id"] = chatID
	o["message_id"] = messageID
	return o
}

// SetInlineMessageID sets the `inline_message_id` value of OptionsEditMessageMedia.
func (o OptionsEditMessageMedia) SetInlineMessageID(inlineMessageID string) OptionsEditMessageMedia {
	o["inline_message_id"] = inlineMessageID
	return o
}

// SetReplyMarkup sets the `reply_markup` value of OptionsEditMessageMedia.
func (o OptionsEditMessageMedia) SetReplyMarkup(replyMarkup InlineKeyboardMarkup) OptionsEditMessageMedia {
	o["reply_markup"] = replyMarkup
	return o
}

// OptionsEditMessageReplyMarkup struct for EditMessageReplyMarkup()
//
// required options: `chat_id` + `message_id` (when `inline_message_id` is not given)
//
//	or `inline_message_id` (when `chat_id` & `message_id` is not given)
//
// other options: `business_connection_id`, `reply_markup`
//
// https://core.telegram.org/bots/api#editmessagereplymarkup
type OptionsEditMessageReplyMarkup MethodOptions

// SetBusinessConnectionID sets the `business_connection_id` value of OptionsEditMessageReplyMarkup.
func (o OptionsEditMessageReplyMarkup) SetBusinessConnectionID(businessConnectionID string) OptionsEditMessageReplyMarkup {
	o["business_connection_id"] = businessConnectionID
	return o
}

// SetIDs sets the `chat_id` and `message_id` values of OptionsEditMessageReplyMarkup.
func (o OptionsEditMessageReplyMarkup) SetIDs(chatID ChatID, messageID int64) OptionsEditMessageReplyMarkup {
	o["chat_id"] = chatID
	o["message_id"] = messageID
	return o
}

// SetInlineMessageID sets the `inline_message_id` value of OptionsEditMessageReplyMarkup.
func (o OptionsEditMessageReplyMarkup) SetInlineMessageID(inlineMessageID string) OptionsEditMessageReplyMarkup {
	o["inline_message_id"] = inlineMessageID
	return o
}

// SetReplyMarkup sets the `reply_markup` value of OptionsEditMessageReplyMarkup.
func (o OptionsEditMessageReplyMarkup) SetReplyMarkup(replyMarkup InlineKeyboardMarkup) OptionsEditMessageReplyMarkup {
	o["reply_markup"] = replyMarkup
	return o
}

// OptionsEditMessageLiveLocation struct for EditMessageLiveLocation()
//
// required options: `chat_id` + `message_id` (when `inline_message_id` is not given)
//
//	or `inline_message_id` (when `chat_id` & `message_id` is not given)
//
// other options: `business_connection_id`, `live_period`, `horizontal_accuracy`, `heading`, `proximity_alert_radius`, `reply_markup`
//
// https://core.telegram.org/bots/api#editmessagelivelocation
type OptionsEditMessageLiveLocation MethodOptions

// SetBusinessConnectionID sets the `business_connection_id` value of OptionsEditMessageLiveLocation.
func (o OptionsEditMessageLiveLocation) SetBusinessConnectionID(businessConnectionID string) OptionsEditMessageLiveLocation {
	o["business_connection_id"] = businessConnectionID
	return o
}

// SetIDs sets the `chat_id` and `message_id` values of OptionsEditMessageLiveLocation.
func (o OptionsEditMessageLiveLocation) SetIDs(chatID ChatID, messageID int64) OptionsEditMessageLiveLocation {
	o["chat_id"] = chatID
	o["message_id"] = messageID
	return o
}

// SetInlineMessageID sets the `inline_message_id` value of OptionsEditMessageLiveLocation.
func (o OptionsEditMessageLiveLocation) SetInlineMessageID(inlineMessageID string) OptionsEditMessageLiveLocation {
	o["inline_message_id"] = inlineMessageID
	return o
}

// SetLivePeriod sets the `live_period` value of OptionsEditMessageLiveLocation.
func (o OptionsEditMessageLiveLocation) SetLivePeriod(livePeriod int) OptionsEditMessageLiveLocation {
	o["live_period"] = livePeriod
	return o
}

// SetHorizontalAccuracy sets the `horizontal_accuracy` value of OptionsEditMessageLiveLocation.
func (o OptionsEditMessageLiveLocation) SetHorizontalAccuracy(horizontalAccuracy float32) OptionsEditMessageLiveLocation {
	o["horizontal_accuracy"] = horizontalAccuracy
	return o
}

// SetHeading sets the `heading` value of OptionsEditMessageLiveLocation.
func (o OptionsEditMessageLiveLocation) SetHeading(heading int) OptionsEditMessageLiveLocation {
	o["heading"] = heading
	return o
}

// SetProximityAlertRadius sets the `proximity_alert_radius` value of OptionsEditMessageLiveLocation.
func (o OptionsEditMessageLiveLocation) SetProximityAlertRadius(proximityAlertRadius int) OptionsEditMessageLiveLocation {
	o["proximity_alert_radius"] = proximityAlertRadius
	return o
}

// SetReplyMarkup sets the `reply_markup` value of OptionsEditMessageLiveLocation.
func (o OptionsEditMessageLiveLocation) SetReplyMarkup(replyMarkup InlineKeyboardMarkup) OptionsEditMessageLiveLocation {
	o["reply_markup"] = replyMarkup
	return o
}

// OptionsStopMessageLiveLocation struct for StopMessageLiveLocation()
//
// required options: `chat_id` + `message_id` (when `inline_message_id` is not given)
//
//	or `inline_message_id` (when `chat_id` & `message_id` is not given)
//
// other options: `business_connection_id`, `reply_markup`
//
// https://core.telegram.org/bots/api#stopmessagelivelocation
type OptionsStopMessageLiveLocation MethodOptions

// SetBusinessConnectionID sets the `business_connection_id` value of OptionsStopMessageLiveLocation.
func (o OptionsStopMessageLiveLocation) SetBusinessConnectionID(businessConnectionID string) OptionsStopMessageLiveLocation {
	o["business_connection_id"] = businessConnectionID
	return o
}

// SetIDs sets the `chat_id` and `message_id` values of OptionsStopMessageLiveLocation.
func (o OptionsStopMessageLiveLocation) SetIDs(chatID ChatID, messageID int64) OptionsStopMessageLiveLocation {
	o["chat_id"] = chatID
	o["message_id"] = messageID
	return o
}

// SetInlineMessageID sets the `inline_message_id` value of OptionsStopMessageLiveLocation.
func (o OptionsStopMessageLiveLocation) SetInlineMessageID(inlineMessageID string) OptionsStopMessageLiveLocation {
	o["inline_message_id"] = inlineMessageID
	return o
}

// SetReplyMarkup sets the `reply_markup` value of OptionsStopMessageLiveLocation.
func (o OptionsStopMessageLiveLocation) SetReplyMarkup(replyMarkup InlineKeyboardMarkup) OptionsStopMessageLiveLocation {
	o["reply_markup"] = replyMarkup
	return o
}

// OptionsEditMessageChecklist struct for EditMessageChecklist()
//
// options include: `reply_markup`.
//
// https://core.telegram.org/bots/api#editmessagechecklist
type OptionsEditMessageChecklist MethodOptions

// SetReplyMarkup sets the `reply_markup` value of OptionsEditMessageChecklist.
func (o OptionsEditMessageChecklist) SetReplyMarkup(replyMarkup InlineKeyboardMarkup) OptionsEditMessageChecklist {
	o["reply_markup"] = replyMarkup
	return o
}

// OptionsAnswerInlineQuery struct for AnswerInlineQuery().
//
// options include: `cache_time`, `is_personal`, `next_offset`, `switch_pm_text`, and `switch_pm_parameter`.
//
// https://core.telegram.org/bots/api#answerinlinequery
type OptionsAnswerInlineQuery MethodOptions

// SetCacheTime sets the `cache_time` value of OptionsAnswerInlineQuery.
func (o OptionsAnswerInlineQuery) SetCacheTime(cacheTime int) OptionsAnswerInlineQuery {
	o["cache_time"] = cacheTime
	return o
}

// SetIsPersonal sets the `is_personal` value of OptionsAnswerInlineQuery.
func (o OptionsAnswerInlineQuery) SetIsPersonal(isPersonal bool) OptionsAnswerInlineQuery {
	o["is_personal"] = isPersonal
	return o
}

// SetNextOffset sets the `next_offset` value of OptionsAnswerInlineQuery.
func (o OptionsAnswerInlineQuery) SetNextOffset(nextOffset string) OptionsAnswerInlineQuery {
	o["next_offset"] = nextOffset
	return o
}

// SetButton sets the `button` value of OptionsAnswerInlineQuery.
func (o OptionsAnswerInlineQuery) SetButton(button InlineQueryResultsButton) OptionsAnswerInlineQuery {
	o["button"] = button
	return o
}

// OptionsSendInvoice struct for SendInvoice().
//
// options include: `message_thread_id`, `direct_messages_topic_id`, `max_tip_amount`, `suggested_tip_amounts`, `start_parameter`, `provider_data`, `photo_url`, `photo_size`, `photo_width`, `photo_height`, `need_name`, `need_phone_number`, `need_email`, `need_shipping_address`, `send_phone_number_to_provider`, `send_email_to_provider`, `is_flexible`, `disable_notification`, `protect_content`, `allow_paid_broadcast`, `message_effect_id`, `suggested_post_parameters`, `reply_parameters`, and `reply_markup`.
//
// https://core.telegram.org/bots/api#sendinvoice
type OptionsSendInvoice MethodOptions

// SetMessageThreadID sets the `message_thread_id` value of OptionsSendInvoice.
func (o OptionsSendInvoice) SetMessageThreadID(messageThreadID int64) OptionsSendInvoice {
	o["message_thread_id"] = messageThreadID
	return o
}

// SetDirectMessagesTopicID sets the `direct_messages_topic_id` value of OptionsSendInvoice.
func (o OptionsSendInvoice) SetDirectMessagesTopicID(directMessagesTopicID int64) OptionsSendInvoice {
	o["direct_messages_topic_id"] = directMessagesTopicID
	return o
}

// SetMaxTipAmount sets the `max_tip_amount` value of OptionsSendInvoice.
func (o OptionsSendInvoice) SetMaxTipAmount(maxTipAmount int) OptionsSendInvoice {
	o["max_tip_amount"] = maxTipAmount
	return o
}

// SetSuggestedTipAmounts sets the `suggested_tip_amounts` value of OptionsSendInvoice.
func (o OptionsSendInvoice) SetSuggestedTipAmounts(suggestedTipAmounts []int) OptionsSendInvoice {
	o["suggested_tip_amounts"] = suggestedTipAmounts
	return o
}

// SetStartParameter sets the `start_parameter` value of OptionsSendInvoice.
func (o OptionsSendInvoice) SetStartParameter(startParameter string) OptionsSendInvoice {
	o["start_parameter"] = startParameter
	return o
}

// SetProviderData sets the `provider_data` value of OptionsSendInvoice.
func (o OptionsSendInvoice) SetProviderData(providerData string) OptionsSendInvoice {
	o["provider_data"] = providerData
	return o
}

// SetPhotoURL sets the `photo_url` value of OptionsSendInvoice.
func (o OptionsSendInvoice) SetPhotoURL(photoURL string) OptionsSendInvoice {
	o["photo_url"] = photoURL
	return o
}

// SetPhotoSize sets the `photo_size` value of OptionsSendInvoice.
func (o OptionsSendInvoice) SetPhotoSize(photoSize int) OptionsSendInvoice {
	o["photo_size"] = photoSize
	return o
}

// SetPhotoWidth sets the `photoWidth` value of OptionsSendInvoice.
func (o OptionsSendInvoice) SetPhotoWidth(photoWidth int) OptionsSendInvoice {
	o["photo_width"] = photoWidth
	return o
}

// SetPhotoHeight sets the `photo_height` value of OptionsSendInvoice.
func (o OptionsSendInvoice) SetPhotoHeight(photoHeight int) OptionsSendInvoice {
	o["photo_height"] = photoHeight
	return o
}

// SetNeedName sets the `need_name` value of OptionsSendInvoice.
func (o OptionsSendInvoice) SetNeedName(needName bool) OptionsSendInvoice {
	o["need_name"] = needName
	return o
}

// SetNeedPhoneNumber sets the `need_phone_number` value of OptionsSendInvoice.
func (o OptionsSendInvoice) SetNeedPhoneNumber(needPhoneNumber bool) OptionsSendInvoice {
	o["need_phone_number"] = needPhoneNumber
	return o
}

// SetNeedEmail sets the `need_email` value of OptionsSendInvoice.
func (o OptionsSendInvoice) SetNeedEmail(needEmail bool) OptionsSendInvoice {
	o["need_email"] = needEmail
	return o
}

// SetNeedShippingAddress sets the `need_shipping_address` value of OptionsSendInvoice.
func (o OptionsSendInvoice) SetNeedShippingAddress(needShippingAddr bool) OptionsSendInvoice {
	o["need_shipping_address"] = needShippingAddr
	return o
}

// SetSendPhoneNumberToProvider sets the `send_phone_number_to_provider` value of OptionsSendInvoice.
func (o OptionsSendInvoice) SetSendPhoneNumberToProvider(sendPhoneNumberToProvider bool) OptionsSendInvoice {
	o["send_phone_number_to_provider"] = sendPhoneNumberToProvider
	return o
}

// SetSendEmailToProvider sets the `send_email_to_provider` value of OptionsSendInvoice.
func (o OptionsSendInvoice) SetSendEmailToProvider(sendEmailToProvider bool) OptionsSendInvoice {
	o["send_email_to_provider"] = sendEmailToProvider
	return o
}

// SetIsFlexible sets the `is_flexible` value of OptionsSendInvoice.
func (o OptionsSendInvoice) SetIsFlexible(isFlexible bool) OptionsSendInvoice {
	o["is_flexible"] = isFlexible
	return o
}

// SetDisableNotification sets the `disable_notification` value of OptionsSendInvoice.
func (o OptionsSendInvoice) SetDisableNotification(disable bool) OptionsSendInvoice {
	o["disable_notification"] = disable
	return o
}

// SetProtectContent sets the `protect_content` value of OptionsSendInvoice.
func (o OptionsSendInvoice) SetProtectContent(protect bool) OptionsSendInvoice {
	o["protect_content"] = protect
	return o
}

// SetAllowPaidBroadcast sets the `allow_paid_broadcast` value of OptionsSendInvoice.
func (o OptionsSendInvoice) SetAllowPaidBroadcast(allow bool) OptionsSendInvoice {
	o["allow_paid_broadcast"] = allow
	return o
}

// SetMessageEffectID sets the `message_effect_id` value of OptionsSendInvoice.
func (o OptionsSendInvoice) SetMessageEffectID(messageEffectID string) OptionsSendInvoice {
	o["message_effect_id"] = messageEffectID
	return o
}

// SetSuggestedPostParameters sets the `suggested_post_parameters` value of OptionsSendInvoice.
func (o OptionsSendInvoice) SetSuggestedPostParameters(suggestedPostParameters SuggestedPostParameters) OptionsSendInvoice {
	o["suggested_post_parameters"] = suggestedPostParameters
	return o
}

// SetReplyParameters sets the `reply_parameters` value of OptionsSendInvoice.
func (o OptionsSendInvoice) SetReplyParameters(replyParameters ReplyParameters) OptionsSendInvoice {
	o["reply_parameters"] = replyParameters
	return o
}

// SetReplyMarkup sets the `reply_markup` value of OptionsSendInvoice.
func (o OptionsSendInvoice) SetReplyMarkup(replyMarkup InlineKeyboardMarkup) OptionsSendInvoice {
	o["reply_markup"] = replyMarkup
	return o
}

// OptionsCreateInvoiceLink struct for CreateInvoiceLink().
//
// options include: `business_connection_id`, `provider_token`, `subscription_period`, `max_tip_amount`, `suggested_tip_amounts`, `provider_data`, `photo_url`, `photo_size`, `photo_width`, `photo_height`, `need_name`, `need_phone_number`, `need_email`, `need_shipping_address`, `send_phone_number_to_provider`, `send_email_to_provider`, and `is_flexible`.
//
// https://core.telegram.org/bots/api#createinvoicelink
type OptionsCreateInvoiceLink MethodOptions

// SetBusinessConnectionID sets the `business_connection_id` value of OptionsCreateInvoiceLink.
func (o OptionsCreateInvoiceLink) SetBusinessConnectionID(businessConnectionID string) OptionsCreateInvoiceLink {
	o["business_connection_id"] = businessConnectionID
	return o
}

// SetProviderToken sets the `provider_token` value of OptionsCreateInvoiceLink.
func (o OptionsCreateInvoiceLink) SetProviderToken(providerToken string) OptionsCreateInvoiceLink {
	o["provider_token"] = providerToken
	return o
}

// SetSubscriptionPeriod sets the `subscription_period` value of OptionsCreateInvoiceLink.
func (o OptionsCreateInvoiceLink) SetSubscriptionPeriod(subscriptionPeriod int) OptionsCreateInvoiceLink {
	o["subscription_period"] = subscriptionPeriod
	return o
}

// SetMaxTipAmount sets the `max_tip_amount` value of OptionsCreateInvoiceLink.
func (o OptionsCreateInvoiceLink) SetMaxTipAmount(maxTipAmount int) OptionsCreateInvoiceLink {
	o["max_tip_amount"] = maxTipAmount
	return o
}

// SetSuggestedTipAmounts sets the `suggested_tip_amounts` value of OptionsCreateInvoiceLink.
func (o OptionsCreateInvoiceLink) SetSuggestedTipAmounts(suggestedTipAmounts []int) OptionsCreateInvoiceLink {
	o["suggested_tip_amounts"] = suggestedTipAmounts
	return o
}

// SetProviderData sets the `provider_data` value of OptionsCreateInvoiceLink.
func (o OptionsCreateInvoiceLink) SetProviderData(providerData string) OptionsCreateInvoiceLink {
	o["provider_data"] = providerData
	return o
}

// SetPhotoURL sets the `photo_url` value of OptionsCreateInvoiceLink.
func (o OptionsCreateInvoiceLink) SetPhotoURL(photoURL string) OptionsCreateInvoiceLink {
	o["photo_url"] = photoURL
	return o
}

// SetPhotoSize sets the `photo_size` value of OptionsCreateInvoiceLink.
func (o OptionsCreateInvoiceLink) SetPhotoSize(photoSize int) OptionsCreateInvoiceLink {
	o["photo_size"] = photoSize
	return o
}

// SetPhotoWidth sets the `photoWidth` value of OptionsCreateInvoiceLink.
func (o OptionsCreateInvoiceLink) SetPhotoWidth(photoWidth int) OptionsCreateInvoiceLink {
	o["photo_width"] = photoWidth
	return o
}

// SetPhotoHeight sets the `photo_height` value of OptionsCreateInvoiceLink.
func (o OptionsCreateInvoiceLink) SetPhotoHeight(photoHeight int) OptionsCreateInvoiceLink {
	o["photo_height"] = photoHeight
	return o
}

// SetNeedName sets the `need_name` value of OptionsCreateInvoiceLink.
func (o OptionsCreateInvoiceLink) SetNeedName(needName bool) OptionsCreateInvoiceLink {
	o["need_name"] = needName
	return o
}

// SetNeedPhoneNumber sets the `need_phone_number` value of OptionsCreateInvoiceLink.
func (o OptionsCreateInvoiceLink) SetNeedPhoneNumber(needPhoneNumber bool) OptionsCreateInvoiceLink {
	o["need_phone_number"] = needPhoneNumber
	return o
}

// SetNeedEmail sets the `need_email` value of OptionsCreateInvoiceLink.
func (o OptionsCreateInvoiceLink) SetNeedEmail(needEmail bool) OptionsCreateInvoiceLink {
	o["need_email"] = needEmail
	return o
}

// SetNeedShippingAddress sets the `need_shipping_address` value of OptionsCreateInvoiceLink.
func (o OptionsCreateInvoiceLink) SetNeedShippingAddress(needShippingAddr bool) OptionsCreateInvoiceLink {
	o["need_shipping_address"] = needShippingAddr
	return o
}

// SetSendPhoneNumberToProvider sets the `send_phone_number_to_provider` value of OptionsCreateInvoiceLink.
func (o OptionsCreateInvoiceLink) SetSendPhoneNumberToProvider(sendPhoneNumberToProvider bool) OptionsCreateInvoiceLink {
	o["send_phone_number_to_provider"] = sendPhoneNumberToProvider
	return o
}

// SetSendEmailToProvider sets the `send_email_to_provider` value of OptionsCreateInvoiceLink.
func (o OptionsCreateInvoiceLink) SetSendEmailToProvider(sendEmailToProvider bool) OptionsCreateInvoiceLink {
	o["send_email_to_provider"] = sendEmailToProvider
	return o
}

// SetIsFlexible sets the `is_flexible` value of OptionsCreateInvoiceLink.
func (o OptionsCreateInvoiceLink) SetIsFlexible(isFlexible bool) OptionsCreateInvoiceLink {
	o["is_flexible"] = isFlexible
	return o
}

// OptionsGetStarTransactions struct for GetStarTransactions().
//
// options include: `offset`, and `limit`.
//
// https://core.telegram.org/bots/api#getstartransactions
type OptionsGetStarTransactions MethodOptions

// SetOffset sets the `offset` value of OptionsGetStarTransactions.
func (o OptionsGetStarTransactions) SetOffset(offset int) OptionsGetStarTransactions {
	o["offset"] = offset
	return o
}

// SetLimit sets the `limit` value of OptionsGetStarTransactions.
//
// Accepted `limit` value is 1-100. (default: 100)
func (o OptionsGetStarTransactions) SetLimit(limit int) OptionsGetStarTransactions {
	o["limit"] = limit
	return o
}

// OptionsSendGame struct for SendGame()
//
// options include: `business_connection_id`, `message_thread_id`, `disable_notification`, `protect_content`, `allow_paid_broadcast`, `message_effect_id`, `reply_parameters`, and `reply_markup`.
//
// https://core.telegram.org/bots/api#sendgame
type OptionsSendGame MethodOptions

// SetBusinessConnectionID sets the `business_connection_id` value of OptionsSendGame.
func (o OptionsSendGame) SetBusinessConnectionID(businessConnectionID string) OptionsSendGame {
	o["business_connection_id"] = businessConnectionID
	return o
}

// SetMessageThreadID sets the `message_thread_id` value of OptionsSendGame.
func (o OptionsSendGame) SetMessageThreadID(messageThreadID int64) OptionsSendGame {
	o["message_thread_id"] = messageThreadID
	return o
}

// SetDisableNotification sets the `disable_notification` value of OptionsSendGame.
func (o OptionsSendGame) SetDisableNotification(disable bool) OptionsSendGame {
	o["disable_notification"] = disable
	return o
}

// SetProtectContent sets the `protect_content` value of OptionsSendGame.
func (o OptionsSendGame) SetProtectContent(protect bool) OptionsSendGame {
	o["protect_content"] = protect
	return o
}

// SetAllowPaidBroadcast sets the `allow_paid_broadcast` value of OptionsSendGame.
func (o OptionsSendGame) SetAllowPaidBroadcast(allow bool) OptionsSendGame {
	o["allow_paid_broadcast"] = allow
	return o
}

// SetMessageEffectID sets the `message_effect_id` value of OptionsSendGame.
func (o OptionsSendGame) SetMessageEffectID(messageEffectID string) OptionsSendGame {
	o["message_effect_id"] = messageEffectID
	return o
}

// SetReplyParameters sets the `reply_parameters` value of OptionsSendGame.
func (o OptionsSendGame) SetReplyParameters(replyParameters ReplyParameters) OptionsSendGame {
	o["reply_parameters"] = replyParameters
	return o
}

// SetReplyMarkup sets the `reply_markup` value of OptionsSendGame.
func (o OptionsSendGame) SetReplyMarkup(replyMarkup InlineKeyboardMarkup) OptionsSendGame {
	o["reply_markup"] = replyMarkup
	return o
}

// OptionsSetGameScore struct for SetGameScore().
//
// required options: `chat_id` + `message_id` (when `inline_message_id` is not given)
//
//	or `inline_message_id` (when `chat_id` & `message_id` is not given)
//
// other options: `force`, and `disable_edit_message`
//
// https://core.telegram.org/bots/api#setgamescore
type OptionsSetGameScore MethodOptions

// SetForce sets the `force` value of OptionsSetGameScore.
func (o OptionsSetGameScore) SetForce(force bool) OptionsSetGameScore {
	o["force"] = force
	return o
}

// SetDisableEditMessage sets the `disable_edit_message` value of OptionsSetGameScore.
func (o OptionsSetGameScore) SetDisableEditMessage(disableEditMessage bool) OptionsSetGameScore {
	o["disable_edit_message"] = disableEditMessage
	return o
}

// SetIDs sets the `chat_id` and `message_id` values of OptionsSetGameScore.
func (o OptionsSetGameScore) SetIDs(chatID ChatID, messageID int64) OptionsSetGameScore {
	o["chat_id"] = chatID
	o["message_id"] = messageID
	return o
}

// SetInlineMessageID sets the `inline_message_id` value of OptionsSetGameScore.
func (o OptionsSetGameScore) SetInlineMessageID(inlineMessageID string) OptionsSetGameScore {
	o["inline_message_id"] = inlineMessageID
	return o
}

// OptionsGetGameHighScores struct for GetGameHighScores().
//
// required options: `chat_id` + `message_id` (when `inline_message_id` is not given)
//
//	or `inline_message_id` (when `chat_id` & `message_id` is not given)
//
// https://core.telegram.org/bots/api#getgamehighscores
type OptionsGetGameHighScores MethodOptions

// SetIDs sets the `chat_id` and `message_id` values of OptionsGetGameHighScores.
func (o OptionsGetGameHighScores) SetIDs(chatID ChatID, messageID int64) OptionsGetGameHighScores {
	o["chat_id"] = chatID
	o["message_id"] = messageID
	return o
}

// SetInlineMessageID sets the `inline_message_id` value of OptionsGetGameHighScores.
func (o OptionsGetGameHighScores) SetInlineMessageID(inlineMessageID string) OptionsGetGameHighScores {
	o["inline_message_id"] = inlineMessageID
	return o
}

// OptionsSavePreparedInlineMessage struct for SavePreparedInlineMessage().
//
// https://core.telegram.org/bots/api#savepreparedinlinemessage
type OptionsSavePreparedInlineMessage MethodOptions

// SetAllowUserChats sets the `allow_user_chats` value of OptionsSavePreparedInlineMessage.
func (o OptionsSavePreparedInlineMessage) SetAllowUserChats(allowUserChats bool) OptionsSavePreparedInlineMessage {
	o["allow_user_chats"] = allowUserChats
	return o
}

// SetAllowBotChats sets the `allow_bot_chats` value of OptionsSavePreparedInlineMessage.
func (o OptionsSavePreparedInlineMessage) SetAllowBotChats(allowBotChats bool) OptionsSavePreparedInlineMessage {
	o["allow_bot_chats"] = allowBotChats
	return o
}

// SetAllowGroupChats sets the `allow_group_chats` value of OptionsSavePreparedInlineMessage.
func (o OptionsSavePreparedInlineMessage) SetAllowGroupChats(allowGroupChats bool) OptionsSavePreparedInlineMessage {
	o["allow_group_chats"] = allowGroupChats
	return o
}

// SetAllowChannelChats sets the `allow_channel_chats` value of OptionsSavePreparedInlineMessage.
func (o OptionsSavePreparedInlineMessage) SetAllowChannelChats(allowChannelChats bool) OptionsSavePreparedInlineMessage {
	o["allow_channel_chats"] = allowChannelChats
	return o
}

// OptionsCreateForumTopic struct for CreateForumTopic().
//
// https://core.telegram.org/bots/api#createforumtopic
type OptionsCreateForumTopic MethodOptions

// SetIconColor sets the `icon_color` value of OptionsCreateForumTopic.
func (o OptionsCreateForumTopic) SetIconColor(iconColor int) OptionsCreateForumTopic {
	o["icon_color"] = iconColor
	return o
}

// SetIconCustomEmojiID sets the `icon_custom_emoji_id` value of OptionsCreateForumTopic.
func (o OptionsCreateForumTopic) SetIconCustomEmojiID(iconCustomEmojiID string) OptionsCreateForumTopic {
	o["icon_custom_emoji_id"] = iconCustomEmojiID
	return o
}

// OptionsEditForumTopic struct for EditForumTopic().
//
// https://core.telegram.org/bots/api#editforumtopic
type OptionsEditForumTopic MethodOptions

// SetName sets the `name` value of OptionsEditForumTopic.
func (o OptionsEditForumTopic) SetName(name string) OptionsEditForumTopic {
	o["name"] = name
	return o
}

// SetIconCustomEmojiID sets the `icon_custom_emoji_id` value of OptionsEditForumTopic.
func (o OptionsEditForumTopic) SetIconCustomEmojiID(iconCustomEmojiID string) OptionsEditForumTopic {
	o["icon_custom_emoji_id"] = iconCustomEmojiID
	return o
}

// OptionsSetBusinessAccountName struct for SetBusinessAccountName().
//
// options include: `last_name`.
//
// https://core.telegram.org/bots/api#setbusinessaccountname
type OptionsSetBusinessAccountName MethodOptions

// SetLastName sets the `last_name` value of OptionsSetBusinessAccountName.
func (o OptionsSetBusinessAccountName) SetLastName(lastName string) OptionsSetBusinessAccountName {
	o["last_name"] = lastName
	return o
}

// OptionsSetBusinessAccountUsername struct for SetBusinessAccountUsername().
//
// options include: `username`.
//
// https://core.telegram.org/bots/api#setbusinessaccountusername
type OptionsSetBusinessAccountUsername MethodOptions

// SetUsername sets the `username` value of OptionsSetBusinessAccountUsername.
func (o OptionsSetBusinessAccountUsername) SetUsername(username string) OptionsSetBusinessAccountUsername {
	o["username"] = username
	return o
}

// OptionsSetBusinessAccountBio struct for SetBusinessAccountBio().
//
// options include: `bio`.
//
// https://core.telegram.org/bots/api#setbusinessaccountbio
type OptionsSetBusinessAccountBio MethodOptions

// SetBio sets the `bio` value of OptionsSetBusinessAccountBio.
func (o OptionsSetBusinessAccountBio) SetBio(bio string) OptionsSetBusinessAccountBio {
	o["bio"] = bio
	return o
}

// OptionsSetBusinessAccountProfilePhoto struct for SetBusinessAccountProfilePhoto().
//
// options include: `is_public`.
//
// https://core.telegram.org/bots/api#setbusinessaccountprofilephoto
type OptionsSetBusinessAccountProfilePhoto MethodOptions

// SetIsPublic sets the `is_public` value of OptionsSetBusinessAccountProfilePhoto.
func (o OptionsSetBusinessAccountProfilePhoto) SetIsPublic(isPublic bool) OptionsSetBusinessAccountProfilePhoto {
	o["is_public"] = isPublic
	return o
}

// OptionsRemoveBusinessAccountProfilePhoto struct for RemoveBusinessAccountProfilePhoto().
//
// options include: `is_public`.
//
// https://core.telegram.org/bots/api#removebusinessaccountprofilephoto
type OptionsRemoveBusinessAccountProfilePhoto MethodOptions

// SetIsPublic sets the `is_public` value of OptionsRemoveBusinessAccountProfilePhoto.
func (o OptionsRemoveBusinessAccountProfilePhoto) SetIsPublic(isPublic bool) OptionsRemoveBusinessAccountProfilePhoto {
	o["is_public"] = isPublic
	return o
}

// OptionsGetBusinessAccountGifts struct for GetBusinessAccountGifts().
//
// options include: `exclude_unsaved`, `exclude_saved`, `exclude_unlimited`, `exclude_limited_upgradable`,
// `exclude_limited_non_upgradable`, `exclude_unique`, `exclude_from_blockchain`, `sort_by_price`, `offset`, and `limit`.
//
// https://core.telegram.org/bots/api#getbusinessaccountgifts
type OptionsGetBusinessAccountGifts MethodOptions

// SetExcludeUnsaved sets the `exclude_unsaved` value of OptionsGetBusinessAccountGifts.
func (o OptionsGetBusinessAccountGifts) SetExcludeUnsaved(excludeUnsaved bool) OptionsGetBusinessAccountGifts {
	o["exclude_unsaved"] = excludeUnsaved
	return o
}

// SetExcludeSaved sets the `exclude_saved` value of OptionsGetBusinessAccountGifts.
func (o OptionsGetBusinessAccountGifts) SetExcludeSaved(excludeSaved bool) OptionsGetBusinessAccountGifts {
	o["exclude_saved"] = excludeSaved
	return o
}

// SetExcludeUnlimited sets the `exclude_unlimited` value of OptionsGetBusinessAccountGifts.
func (o OptionsGetBusinessAccountGifts) SetExcludeUnlimited(excludeUnlimited bool) OptionsGetBusinessAccountGifts {
	o["exclude_unlimited"] = excludeUnlimited
	return o
}

// SetExcludeLimitedUpgradable sets the `exclude_limited_upgradable` value of OptionsGetBusinessAccountGifts.
func (o OptionsGetBusinessAccountGifts) SetExcludeLimitedUpgradable(excludeLimitedUpgradable bool) OptionsGetBusinessAccountGifts {
	o["exclude_limited_upgradable"] = excludeLimitedUpgradable
	return o
}

// SetExcludeLimitedNonUpgradable sets the `exclude_limited_non_upgradable` value of OptionsGetBusinessAccountGifts.
func (o OptionsGetBusinessAccountGifts) SetExcludeLimitedNonUpgradable(excludeLimitedNonUpgradable bool) OptionsGetBusinessAccountGifts {
	o["exclude_limited_non_upgradable"] = excludeLimitedNonUpgradable
	return o
}

// SetExcludeUnique sets the `exclude_unique` value of OptionsGetBusinessAccountGifts.
func (o OptionsGetBusinessAccountGifts) SetExcludeUnique(excludeUnique bool) OptionsGetBusinessAccountGifts {
	o["exclude_unique"] = excludeUnique
	return o
}

// SetExcludeFromBlockchain sets the `exclude_from_blockchain` value of OptionsGetBusinessAccountGifts.
func (o OptionsGetBusinessAccountGifts) SetExcludeFromBlockchain(excludeFromBlockchain bool) OptionsGetBusinessAccountGifts {
	o["exclude_from_blockchain"] = excludeFromBlockchain
	return o
}

// SetSortByPrice sets the `sort_by_price` value of OptionsGetBusinessAccountGifts.
func (o OptionsGetBusinessAccountGifts) SetSortByPrice(sortByPrice bool) OptionsGetBusinessAccountGifts {
	o["sort_by_price"] = sortByPrice
	return o
}

// SetOffset sets the `offset` value of OptionsGetBusinessAccountGifts.
func (o OptionsGetBusinessAccountGifts) SetOffset(offset string) OptionsGetBusinessAccountGifts {
	o["offset"] = offset
	return o
}

// SetLimit sets the `limit` value of OptionsGetBusinessAccountGifts.
func (o OptionsGetBusinessAccountGifts) SetLimit(limit int) OptionsGetBusinessAccountGifts {
	o["limit"] = limit
	return o
}

// OptionsGetUserGifts struct for GetUserGifts().
//
// options include: `exclude_unlimited`, `exclude_limited_upgradable`, `exclude_limited_non_upgradable`, `exclude_from_blockchain`, `exclude_unique`, `sort_by_price`, `offset`, and `limit`.
//
// https://core.telegram.org/bots/api#getusergifts
type OptionsGetUserGifts MethodOptions

// SetExcludeUnlimited sets the `exclude_unlimited` value of OptionsGetUserGifts.
func (o OptionsGetUserGifts) SetExcludeUnlimited(excludeUnlimited bool) OptionsGetUserGifts {
	o["exclude_unlimited"] = excludeUnlimited
	return o
}

// SetExcludeLimitedUpgradable sets the `exclude_limited_upgradable` value of OptionsGetUserGifts.
func (o OptionsGetUserGifts) SetExcludeLimitedUpgradable(excludeLimitedUpgradable bool) OptionsGetUserGifts {
	o["exclude_limited_upgradable"] = excludeLimitedUpgradable
	return o
}

// SetExcludeLimitedNonUpgradable sets the `exclude_limited_non_upgradable` value of OptionsGetUserGifts.
func (o OptionsGetUserGifts) SetExcludeLimitedNonUpgradable(excludeLimitedNonUpgradable bool) OptionsGetUserGifts {
	o["exclude_limited_non_upgradable"] = excludeLimitedNonUpgradable
	return o
}

// SetExcludeFromBlockchain sets the `exclude_from_blockchain` value of OptionsGetUserGifts.
func (o OptionsGetUserGifts) SetExcludeFromBlockchain(excludeFromBlockchain bool) OptionsGetUserGifts {
	o["exclude_from_blockchain"] = excludeFromBlockchain
	return o
}

// SetExcludeUnique sets the `exclude_unique` value of OptionsGetUserGifts.
func (o OptionsGetUserGifts) SetExcludeUnique(excludeUnique bool) OptionsGetUserGifts {
	o["exclude_unique"] = excludeUnique
	return o
}

// SetSortByPrice sets the `sort_by_price` value of OptionsGetUserGifts.
func (o OptionsGetUserGifts) SetSortByPrice(sortByPrice bool) OptionsGetUserGifts {
	o["sort_by_price"] = sortByPrice
	return o
}

// SetOffset sets the `offset` value of OptionsGetUserGifts.
func (o OptionsGetUserGifts) SetOffset(offset string) OptionsGetUserGifts {
	o["offset"] = offset
	return o
}

// SetLimit sets the `limit` value of OptionsGetUserGifts.
func (o OptionsGetUserGifts) SetLimit(limit int) OptionsGetUserGifts {
	o["limit"] = limit
	return o
}

// OptionsGetChatGifts struct for GetChatGifts().
//
// options include: `exclude_unsaved`, `exclude_saved`, `exclude_unlimited`, `exclude_limited_upgradable`, `exclude_limited_non_upgradable`, `exclude_from_blockchain`, `exclude_unique`, `sort_by_price`, `offset`, and `limit`.
//
// https://core.telegram.org/bots/api#getchatgifts

type OptionsGetChatGifts MethodOptions

// SetExcludeUnsaved sets the `exclude_unsaved` value of OptionsGetChatGifts.
func (o OptionsGetChatGifts) SetExcludeUnsaved(excludeUnsaved bool) OptionsGetChatGifts {
	o["exclude_unsaved"] = excludeUnsaved
	return o
}

// SetExcludeSaved sets the `exclude_saved` value of OptionsGetChatGifts.
func (o OptionsGetChatGifts) SetExcludeSaved(excludeSaved bool) OptionsGetChatGifts {
	o["exclude_saved"] = excludeSaved
	return o
}

// SetExcludeUnlimited sets the `exclude_unlimited` value of OptionsGetChatGifts.
func (o OptionsGetChatGifts) SetExcludeUnlimited(excludeUnlimited bool) OptionsGetChatGifts {
	o["exclude_unlimited"] = excludeUnlimited
	return o
}

// SetExcludeLimitedUpgradable sets the `exclude_limited_upgradable` value of OptionsGetChatGifts.
func (o OptionsGetChatGifts) SetExcludeLimitedUpgradable(excludeLimitedUpgradable bool) OptionsGetChatGifts {
	o["exclude_limited_upgradable"] = excludeLimitedUpgradable
	return o
}

// SetExcludeLimitedNonUpgradable sets the `exclude_limited_non_upgradable` value of OptionsGetChatGifts.
func (o OptionsGetChatGifts) SetExcludeLimitedNonUpgradable(excludeLimitedNonUpgradable bool) OptionsGetChatGifts {
	o["exclude_limited_non_upgradable"] = excludeLimitedNonUpgradable
	return o
}

// SetExcludeFromBlockchain sets the `exclude_from_blockchain` value of OptionsGetChatGifts.
func (o OptionsGetChatGifts) SetExcludeFromBlockchain(excludeFromBlockchain bool) OptionsGetChatGifts {
	o["exclude_from_blockchain"] = excludeFromBlockchain
	return o
}

// SetExcludeUnique sets the `exclude_unique` value of OptionsGetChatGifts.
func (o OptionsGetChatGifts) SetExcludeUnique(excludeUnique bool) OptionsGetChatGifts {
	o["exclude_unique"] = excludeUnique
	return o
}

// SetSortByPrice sets the `sort_by_price` value of OptionsGetChatGifts.
func (o OptionsGetChatGifts) SetSortByPrice(sortByPrice bool) OptionsGetChatGifts {
	o["sort_by_price"] = sortByPrice
	return o
}

// SetOffset sets the `offset` value of OptionsGetChatGifts.
func (o OptionsGetChatGifts) SetOffset(offset string) OptionsGetChatGifts {
	o["offset"] = offset
	return o
}

// SetLimit sets the `limit` value of OptionsGetChatGifts.
func (o OptionsGetChatGifts) SetLimit(limit int) OptionsGetChatGifts {
	o["limit"] = limit
	return o
}

// OptionsUpgradeGift struct for UpgradeGift().
//
// options include: `keep_original_details`, and `star_count`.
//
// https://core.telegram.org/bots/api#upgradegift
type OptionsUpgradeGift MethodOptions

// SetKeepOriginalDetails sets the `keep_original_details` value of OptionsUpgradeGift.
func (o OptionsUpgradeGift) SetKeepOriginalDetails(keep bool) OptionsUpgradeGift {
	o["keep_original_details"] = keep
	return o
}

// SetStarCount sets the `star_count` value of OptionsUpgradeGift.
func (o OptionsUpgradeGift) SetStarCount(starCount int) OptionsUpgradeGift {
	o["star_count"] = starCount
	return o
}

// OptionsTransferGift struct for TransferGift().
//
// options include: `star_count`.
//
// https://core.telegram.org/bots/api#transfergift
type OptionsTransferGift MethodOptions

// SetStarCount sets the `star_count` value of OptionsTrasnferGift.
func (o OptionsTransferGift) SetStarCount(starCount int) OptionsTransferGift {
	o["star_count"] = starCount
	return o
}

// OptionsPostStory struct for PostStory().
//
// options include: `caption`, `parse_mode`, `caption_entities`, `areas`, `post_to_chat_page`, and `protect_content`.
//
// https://core.telegram.org/bots/api#poststory
type OptionsPostStory MethodOptions

// SetCaption sets the `caption` value of OptionsPostStory.
func (o OptionsPostStory) SetCaption(caption string) OptionsPostStory {
	o["caption"] = caption
	return o
}

// SetParseMode sets the `parse_mode` value of OptionsPostStory.
func (o OptionsPostStory) SetParseMode(parseMode ParseMode) OptionsPostStory {
	o["parse_mode"] = parseMode
	return o
}

// SetCaptionEntities sets the `caption_entities` value of OptionsPostStory.
func (o OptionsPostStory) SetCaptionEntities(captionEntities []MessageEntity) OptionsPostStory {
	o["caption_entities"] = captionEntities
	return o
}

// SetAreas sets the `areas` value of OptionsPostStory.
func (o OptionsPostStory) SetAreas(areas []StoryArea) OptionsPostStory {
	o["areas"] = areas
	return o
}

// SetPostToChatPage sets the `post_to_chat_page` value of OptionsPostStory.
func (o OptionsPostStory) SetPostToChatPage(postToChatPage bool) OptionsPostStory {
	o["post_to_chat_page"] = postToChatPage
	return o
}

// SetProtectContent sets the `protect_content` value of OptionsPostStory.
func (o OptionsPostStory) SetProtectContent(protectContent bool) OptionsPostStory {
	o["protect_content"] = protectContent
	return o
}

// OptionsRepostStory struct for RepostStory().
//
// options include: `post_to_chat_page`, and `protect_content`.
//
// https://core.telegram.org/bots/api#repoststory
type OptionsRepostStory MethodOptions

// SetPostToChatPage sets the `post_to_chat_page` value of OptionsRepostStory.
func (o OptionsRepostStory) SetPostToChatPage(postToChatPage bool) OptionsRepostStory {
	o["post_to_chat_page"] = postToChatPage
	return o
}

// SetProtectContent sets the `protect_content` value of OptionsRepostStory.
func (o OptionsRepostStory) SetProtectContent(protectContent bool) OptionsRepostStory {
	o["protect_content"] = protectContent
	return o
}

// OptionsEditStory struct for EditStory().
//
// options include: `caption`, `parse_mode`, `caption_entities`, and `areas`.
//
// https://core.telegram.org/bots/api#editstory
type OptionsEditStory MethodOptions

// SetCaption sets the `caption` value of OptionsEditStory.
func (o OptionsEditStory) SetCaption(caption string) OptionsEditStory {
	o["caption"] = caption
	return o
}

// SetParseMode sets the `parse_mode` value of OptionsEditStory.
func (o OptionsEditStory) SetParseMode(parseMode ParseMode) OptionsEditStory {
	o["parse_mode"] = parseMode
	return o
}

// SetCaptionEntities sets the `caption_entities` value of OptionsEditStory.
func (o OptionsEditStory) SetCaptionEntities(captionEntities []MessageEntity) OptionsEditStory {
	o["caption_entities"] = captionEntities
	return o
}

// SetAreas sets the `areas` value of OptionsEditStory.
func (o OptionsEditStory) SetAreas(areas []StoryArea) OptionsEditStory {
	o["areas"] = areas
	return o
}

// OptionsGiftPremiumSubscription struct for GiftPremiumSubscription().
//
// options include: `text`, `text_parse_mode`, and `text_entities`.
//
// https://core.telegram.org/bots/api#giftpremiumsubscription
type OptionsGiftPremiumSubscription MethodOptions

// SetText sets the `text` value of OptionsGiftPremiumSubscription.
func (o OptionsGiftPremiumSubscription) SetText(text string) OptionsGiftPremiumSubscription {
	o["text"] = text
	return o
}

// SetTextParseMode sets the `text_parse_mode` value of OptionsGiftPremiumSubscription.
func (o OptionsGiftPremiumSubscription) SetTextParseMode(parseMode ParseMode) OptionsGiftPremiumSubscription {
	o["text_parse_mode"] = parseMode
	return o
}

// SetTextEntities sets the `text_entities` value of OptionsGiftPremiumSubscription.
func (o OptionsGiftPremiumSubscription) SetTextEntities(textEntities []MessageEntity) OptionsGiftPremiumSubscription {
	o["text_entities"] = textEntities
	return o
}
