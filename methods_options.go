package telegrambot

// https://core.telegram.org/bots/api#available-methods

// MethodOptions is a type for methods' options parameter.
type MethodOptions map[string]interface{}

// OptionsGetUpdates struct for GetUpdates().
//
// options include: offset, limit, timeout, and allowed_updates.
//
// https://core.telegram.org/bots/api#getupdates
type OptionsGetUpdates MethodOptions

// SetOffset sets the offset value of OptionsGetUpdates.
func (o OptionsGetUpdates) SetOffset(offset int) OptionsGetUpdates {
	o["offset"] = offset
	return o
}

// SetLimit sets the limit value of OptionsGetUpdates.
func (o OptionsGetUpdates) SetLimit(limit int) OptionsGetUpdates {
	o["limit"] = limit
	return o
}

// SetTimeout sets the timeout value of OptionsGetUpdates.
func (o OptionsGetUpdates) SetTimeout(timeout int) OptionsGetUpdates {
	o["timeout"] = timeout
	return o
}

// SetAllowedUpdates sets the allowed_updates value of OptionsGetUpdates.
func (o OptionsGetUpdates) SetAllowedUpdates(allowedUpdates []AllowedUpdate) OptionsGetUpdates {
	o["allowed_updates"] = allowedUpdates
	return o
}

// OptionsSendMessage struct for SendMessage().
//
// options include: parse_mode, disable_web_page_preview, disable_notification, reply_to_message_id, and reply_markup.
//
// https://core.telegram.org/bots/api#sendmessage
type OptionsSendMessage MethodOptions

// SetParseMode sets the parse_mode value of OptionsSendMessage.
func (o OptionsSendMessage) SetParseMode(parseMode ParseMode) OptionsSendMessage {
	o["parse_mode"] = parseMode
	return o
}

// SetDisableWebPagePreview sets the disable_web_page_preview value of OptionsSendMessage.
func (o OptionsSendMessage) SetDisableWebPagePreview(disable bool) OptionsSendMessage {
	o["disable_web_page_preview"] = disable
	return o
}

// SetDisableNotification sets the disable_notification value of OptionsSendMessage.
func (o OptionsSendMessage) SetDisableNotification(disable bool) OptionsSendMessage {
	o["disable_notification"] = disable
	return o
}

// SetReplyToMessageID sets the reply_to_message_id value of OptionsSendMessage.
func (o OptionsSendMessage) SetReplyToMessageID(replyToMessageID int) OptionsSendMessage {
	o["reply_to_message_id"] = replyToMessageID
	return o
}

// SetReplyMarkup sets the reply_markup value of OptionsSendMessage.
//
// replyMarkup can be one of InlineKeyboardMarkup, ReplyKeyboardMarkup, ReplyKeyboardRemove, or ForceReply.
func (o OptionsSendMessage) SetReplyMarkup(replyMarkup interface{}) OptionsSendMessage {
	o["reply_markup"] = replyMarkup
	return o
}

// OptionsForwardMessage struct for ForwardMessage().
//
// options include: disable_notification
//
// https://core.telegram.org/bots/api#forwardmessage
type OptionsForwardMessage MethodOptions

// SetDisableNotification sets the disable_notification value of OptionsForwardMessage.
func (o OptionsForwardMessage) SetDisableNotification(disable bool) OptionsForwardMessage {
	o["disable_notification"] = disable
	return o
}

// OptionsSendPhoto struct for SendPhoto().
//
// options include: caption, parse_mode, disable_notification, reply_to_message_id, and reply_markup.
//
// https://core.telegram.org/bots/api#sendphoto
type OptionsSendPhoto MethodOptions

// SetCaption sets the caption value of OptionsSendPhoto.
func (o OptionsSendPhoto) SetCaption(caption string) OptionsSendPhoto {
	o["caption"] = caption
	return o
}

// SetParseMode sets the parse_mode value of OptionsSendPhoto.
func (o OptionsSendPhoto) SetParseMode(parseMode ParseMode) OptionsSendPhoto {
	o["parse_mode"] = parseMode
	return o
}

// SetDisableNotification sets the disable_notification value of OptionsSendPhoto.
func (o OptionsSendPhoto) SetDisableNotification(disable bool) OptionsSendPhoto {
	o["disable_notification"] = disable
	return o
}

// SetReplyToMessageID sets the reply_to_message_id value of OptionsSendPhoto.
func (o OptionsSendPhoto) SetReplyToMessageID(replyToMessageID int) OptionsSendPhoto {
	o["reply_to_message_id"] = replyToMessageID
	return o
}

// SetReplyMarkup sets the reply_markup value of OptionsSendPhoto.
//
// replyMarkup can be one of InlineKeyboardMarkup, ReplyKeyboardMarkup, ReplyKeyboardRemove, or ForceReply.
func (o OptionsSendPhoto) SetReplyMarkup(replyMarkup interface{}) OptionsSendPhoto {
	o["reply_markup"] = replyMarkup
	return o
}

// OptionsSendAudio struct for SendAudio().
//
// options include: caption, parse_mode, duration, performer, title, disable_notification, reply_to_message_id, and reply_markup.
//
// https://core.telegram.org/bots/api#sendaudio
type OptionsSendAudio MethodOptions

// SetCaption sets the caption value of OptionsSendAudio.
func (o OptionsSendAudio) SetCaption(caption string) OptionsSendAudio {
	o["caption"] = caption
	return o
}

// SetParseMode sets the parse_mode value of OptionsSendAudio.
func (o OptionsSendAudio) SetParseMode(parseMode ParseMode) OptionsSendAudio {
	o["parse_mode"] = parseMode
	return o
}

// SetDuration sets the duration value of OptionsSendAudio.
func (o OptionsSendAudio) SetDuration(duration int) OptionsSendAudio {
	o["duration"] = duration
	return o
}

// SetPerformer sets the performer value of OptionsSendAudio.
func (o OptionsSendAudio) SetPerformer(performer string) OptionsSendAudio {
	o["performer"] = performer
	return o
}

// SetTitle sets the title value of OptionsSendAudio.
func (o OptionsSendAudio) SetTitle(title string) OptionsSendAudio {
	o["title"] = title
	return o
}

// SetThumb sets the thumb value of OptionsSendAudio.
//
// thumb can be one of InputFile or string.
func (o OptionsSendAudio) SetThumb(thumb interface{}) OptionsSendAudio {
	o["thumb"] = thumb
	return o
}

// SetDisableNotification sets the disable_notification value of OptionsSendAudio.
func (o OptionsSendAudio) SetDisableNotification(disable bool) OptionsSendAudio {
	o["disable_notification"] = disable
	return o
}

// SetReplyToMessageID sets the reply_to_message_id value of OptionsSendAudio.
func (o OptionsSendAudio) SetReplyToMessageID(replyToMessageID int) OptionsSendAudio {
	o["reply_to_message_id"] = replyToMessageID
	return o
}

// SetReplyMarkup sets the reply_markup value of OptionsSendAudio.
//
// replyMarkup can be one of InlineKeyboardMarkup, ReplyKeyboardMarkup, ReplyKeyboardRemove, or ForceReply.
func (o OptionsSendAudio) SetReplyMarkup(replyMarkup interface{}) OptionsSendAudio {
	o["reply_markup"] = replyMarkup
	return o
}

// OptionsSendDocument struct for SendDocument().
//
// options include: caption, parse_mode, disable_notification, reply_to_message_id, and reply_markup.
//
// https://core.telegram.org/bots/api#senddocument
type OptionsSendDocument MethodOptions

// SetThumb sets the thumb value of OptionsSendDocument.
//
// thumb can be one of InputFile or string.
func (o OptionsSendDocument) SetThumb(thumb interface{}) OptionsSendDocument {
	o["thumb"] = thumb
	return o
}

// SetCaption sets the caption value of OptionsSendDocument.
func (o OptionsSendDocument) SetCaption(caption string) OptionsSendDocument {
	o["caption"] = caption
	return o
}

// SetParseMode sets the parse_mode value of OptionsSendDocument.
func (o OptionsSendDocument) SetParseMode(parseMode ParseMode) OptionsSendDocument {
	o["parse_mode"] = parseMode
	return o
}

// SetDisableNotification sets the disable_notification value of OptionsSendDocument.
func (o OptionsSendDocument) SetDisableNotification(disable bool) OptionsSendDocument {
	o["disable_notification"] = disable
	return o
}

// SetReplyToMessageID sets the reply_to_message_id value of OptionsSendDocument.
func (o OptionsSendDocument) SetReplyToMessageID(replyToMessageID int) OptionsSendDocument {
	o["reply_to_message_id"] = replyToMessageID
	return o
}

// SetReplyMarkup sets the reply_markup value of OptionsSendDocument.
//
// replyMarkup can be one of InlineKeyboardMarkup, ReplyKeyboardMarkup, ReplyKeyboardRemove, or ForceReply.
func (o OptionsSendDocument) SetReplyMarkup(replyMarkup interface{}) OptionsSendDocument {
	o["reply_markup"] = replyMarkup
	return o
}

// OptionsSendSticker struct for SendSticker().
//
// options include: disable_notification, reply_to_message_id, and reply_markup.
//
// https://core.telegram.org/bots/api#sendsticker
type OptionsSendSticker MethodOptions

// SetDisableNotification sets the disable_notification value of OptionsSendSticker.
func (o OptionsSendSticker) SetDisableNotification(disable bool) OptionsSendSticker {
	o["disable_notification"] = disable
	return o
}

// SetReplyToMessageID sets the reply_to_message_id value of OptionsSendSticker.
func (o OptionsSendSticker) SetReplyToMessageID(replyToMessageID int) OptionsSendSticker {
	o["reply_to_message_id"] = replyToMessageID
	return o
}

// SetReplyMarkup sets the reply_markup value of OptionsSendSticker.
//
// replyMarkup can be one of InlineKeyboardMarkup, ReplyKeyboardMarkup, ReplyKeyboardRemove, or ForceReply.
func (o OptionsSendSticker) SetReplyMarkup(replyMarkup interface{}) OptionsSendSticker {
	o["reply_markup"] = replyMarkup
	return o
}

// OptionsCreateNewStickerSet struct for CreateNewStickerSet().
//
// options include: `png_sticker`, `tgs_sticker`, `contains_masks`, and `mask_position`
//
// https://core.telegram.org/bots/api#createnewstickerset
type OptionsCreateNewStickerSet MethodOptions

// SetPNGSticker sets the `png_sticker` value of OptionsCreateNewStickerSet.
func (o OptionsCreateNewStickerSet) SetPNGSticker(pngSticker InputFile) OptionsCreateNewStickerSet {
	o["png_sticker"] = pngSticker
	return o
}

// SetPNGStickerString sets the `png_sticker` value of OptionsCreateNewStickerSet.
//
// `thumb` can be a file_id or a http url to a file
func (o OptionsCreateNewStickerSet) SetPNGStickerString(pngSticker string) OptionsCreateNewStickerSet {
	o["png_sticker"] = pngSticker
	return o
}

// SetTGSSticker sets the `tgs_sticker` value of OptionsCreateNewStickerSet.
func (o OptionsCreateNewStickerSet) SetTGSSticker(tgsSticker InputFile) OptionsCreateNewStickerSet {
	o["tgs_sticker"] = tgsSticker
	return o
}

// SetContainsMasks sets the contains_masks value of OptionsCreateNewStickerSet.
func (o OptionsCreateNewStickerSet) SetContainsMasks(containsMasks bool) OptionsCreateNewStickerSet {
	o["contains_masks"] = containsMasks
	return o
}

// SetMaskPosition sets the mask_position value of OptionsCreateNewStickerSet.
func (o OptionsCreateNewStickerSet) SetMaskPosition(maskPosition MaskPosition) OptionsCreateNewStickerSet {
	o["mask_position"] = maskPosition
	return o
}

// OptionsAddStickerToSet struct for AddStickerToSet()
//
// options include: `png_sticker`, `tgs_sticker`, and `mask_position`
//
// https://core.telegram.org/bots/api#addstickertoset
type OptionsAddStickerToSet MethodOptions

// SetPNGSticker sets the `png_sticker` value of OptionsCreateNew.
func (o OptionsAddStickerToSet) SetPNGSticker(pngSticker InputFile) OptionsAddStickerToSet {
	o["png_sticker"] = pngSticker
	return o
}

// SetPNGStickerString sets the `png_sticker` value of OptionsCreateNew.
//
// `thumb` can be a file_id or a http url to a file
func (o OptionsAddStickerToSet) SetPNGStickerString(pngSticker string) OptionsAddStickerToSet {
	o["png_sticker"] = pngSticker
	return o
}

// SetTGSSticker sets the `tgs_sticker` value of OptionsCreateNewStickerSet.
func (o OptionsAddStickerToSet) SetTGSSticker(tgsSticker InputFile) OptionsAddStickerToSet {
	o["tgs_sticker"] = tgsSticker
	return o
}

// SetMaskPosition sets the mask_position value of OptionsAddStickerToSet.
func (o OptionsAddStickerToSet) SetMaskPosition(maskPosition MaskPosition) OptionsAddStickerToSet {
	o["mask_position"] = maskPosition
	return o
}

// OptionsSetStickerSetThumb struct for SetStickerSetThumb()
//
// options include: `thumb`
//
// https://core.telegram.org/bots/api#setstickersetthumb
type OptionsSetStickerSetThumb MethodOptions

// SetThumb sets the `thumb` value of OptionsSetStickerSetThumb.
func (o OptionsSetStickerSetThumb) SetThumb(thumb InputFile) OptionsSetStickerSetThumb {
	o["thumb"] = thumb
	return o
}

// SetThumbString sets the `thumb` value of OptionsSetStickerSetThumb.
//
// `thumb` can be a file_id or a http url to a file
func (o OptionsSetStickerSetThumb) SetThumbString(thumb string) OptionsSetStickerSetThumb {
	o["thumb"] = thumb
	return o
}

// OptionsSendVideo struct for SendVideo().
//
// options include: duration, caption, parse_mode, supports_streaming, disable_notification, reply_to_message_id, and reply_markup.
//
// https://core.telegram.org/bots/api#sendvideo
type OptionsSendVideo MethodOptions

// SetDuration sets the duration value of OptionsSendVideo.
func (o OptionsSendVideo) SetDuration(duration int) OptionsSendVideo {
	o["duration"] = duration
	return o
}

// SetWidth sets the width value of OptionsSendVideo.
func (o OptionsSendVideo) SetWidth(width int) OptionsSendVideo {
	o["width"] = width
	return o
}

// SetHeight sets the height value of OptionsSendVideo.
func (o OptionsSendVideo) SetHeight(height int) OptionsSendVideo {
	o["height"] = height
	return o
}

// SetThumb sets the thumb value of OptionsSendVideo.
//
// thumb can be one of InputFile or string.
func (o OptionsSendVideo) SetThumb(thumb interface{}) OptionsSendVideo {
	o["thumb"] = thumb
	return o
}

// SetCaption sets the caption value of OptionsSendVideo.
func (o OptionsSendVideo) SetCaption(caption string) OptionsSendVideo {
	o["caption"] = caption
	return o
}

// SetParseMode sets the parse_mode value of OptionsSendVideo.
func (o OptionsSendVideo) SetParseMode(parseMode ParseMode) OptionsSendVideo {
	o["parse_mode"] = parseMode
	return o
}

// SetSupportsStreaming sets the supports_streaming value of OptionsSendVideo.
func (o OptionsSendVideo) SetSupportsStreaming(supportsStreaming bool) OptionsSendVideo {
	o["supports_streaming"] = supportsStreaming
	return o
}

// SetDisableNotification sets the disable_notification value of OptionsSendVideo.
func (o OptionsSendVideo) SetDisableNotification(disable bool) OptionsSendVideo {
	o["disable_notification"] = disable
	return o
}

// SetReplyToMessageID sets the reply_to_message_id value of OptionsSendVideo.
func (o OptionsSendVideo) SetReplyToMessageID(replyToMessageID int) OptionsSendVideo {
	o["reply_to_message_id"] = replyToMessageID
	return o
}

// SetReplyMarkup sets the reply_markup value of OptionsSendVideo.
//
// replyMarkup can be one of InlineKeyboardMarkup, ReplyKeyboardMarkup, ReplyKeyboardRemove, or ForceReply.
func (o OptionsSendVideo) SetReplyMarkup(replyMarkup interface{}) OptionsSendVideo {
	o["reply_markup"] = replyMarkup
	return o
}

// OptionsSendAnimation struct for SendAnimation().
//
// options include: duration, width, height, thumb, caption, parse_mode, disable_notification, reply_to_message_id, and reply_markup.
//
// https://core.telegram.org/bots/api#sendanimation
type OptionsSendAnimation MethodOptions

// SetDuration sets the duration value of OptionsSendAnimation.
func (o OptionsSendAnimation) SetDuration(duration int) OptionsSendAnimation {
	o["duration"] = duration
	return o
}

// SetWidth sets the width value of OptionsSendAnimation.
func (o OptionsSendAnimation) SetWidth(width int) OptionsSendAnimation {
	o["width"] = width
	return o
}

// SetHeight sets the height value of OptionsSendAnimation.
func (o OptionsSendAnimation) SetHeight(height int) OptionsSendAnimation {
	o["height"] = height
	return o
}

// SetThumb sets the thumb value of OptionsSendAnimation.
//
// thumb can be one of InputFile or string.
func (o OptionsSendAnimation) SetThumb(thumb interface{}) OptionsSendAnimation {
	o["thumb"] = thumb
	return o
}

// SetCaption sets the caption value of OptionsSendAnimation.
func (o OptionsSendAnimation) SetCaption(caption string) OptionsSendAnimation {
	o["caption"] = caption
	return o
}

// SetParseMode sets the parse_mode value of OptionsSendAnimation.
func (o OptionsSendAnimation) SetParseMode(parseMode ParseMode) OptionsSendAnimation {
	o["parse_mode"] = parseMode
	return o
}

// SetDisableNotification sets the disable_notification value of OptionsSendAnimation.
func (o OptionsSendAnimation) SetDisableNotification(disable bool) OptionsSendAnimation {
	o["disable_notification"] = disable
	return o
}

// SetReplyToMessageID sets the reply_to_message_id value of OptionsSendAnimation.
func (o OptionsSendAnimation) SetReplyToMessageID(replyToMessageID int) OptionsSendAnimation {
	o["reply_to_message_id"] = replyToMessageID
	return o
}

// SetReplyMarkup sets the reply_markup value of OptionsSendAnimation.
//
// replyMarkup can be one of InlineKeyboardMarkup, ReplyKeyboardMarkup, ReplyKeyboardRemove, or ForceReply.
func (o OptionsSendAnimation) SetReplyMarkup(replyMarkup interface{}) OptionsSendAnimation {
	o["reply_markup"] = replyMarkup
	return o
}

// OptionsSendVoice struct for SendVoice().
//
// options include: caption, parse_mode, duration, disable_notification, reply_to_message_id, and reply_markup.
//
// https://core.telegram.org/bots/api#sendvoice
type OptionsSendVoice MethodOptions

// SetCaption sets the caption value of OptionsSendVoice.
func (o OptionsSendVoice) SetCaption(caption string) OptionsSendVoice {
	o["caption"] = caption
	return o
}

// SetParseMode sets the parse_mode value of OptionsSendVoice.
func (o OptionsSendVoice) SetParseMode(parseMode ParseMode) OptionsSendVoice {
	o["parse_mode"] = parseMode
	return o
}

// SetDuration sets the duration value of OptionsSendVoice.
func (o OptionsSendVoice) SetDuration(duration int) OptionsSendVoice {
	o["duration"] = duration
	return o
}

// SetDisableNotification sets the disable_notification value of OptionsSendVoice.
func (o OptionsSendVoice) SetDisableNotification(disable bool) OptionsSendVoice {
	o["disable_notification"] = disable
	return o
}

// SetReplyToMessageID sets the reply_to_message_id value of OptionsSendVoice.
func (o OptionsSendVoice) SetReplyToMessageID(replyToMessageID int) OptionsSendVoice {
	o["reply_to_message_id"] = replyToMessageID
	return o
}

// SetReplyMarkup sets the reply_markup value of OptionsSendVoice.
//
// replyMarkup can be one of InlineKeyboardMarkup, ReplyKeyboardMarkup, ReplyKeyboardRemove, or ForceReply.
func (o OptionsSendVoice) SetReplyMarkup(replyMarkup interface{}) OptionsSendVoice {
	o["reply_markup"] = replyMarkup
	return o
}

// OptionsSendVideoNote struct for SendVideoNote().
//
// options include: duration, length, disable_notification, reply_to_message_id, and reply_markup.
// (XXX: API returns 'Bad Request: wrong video note length' when length is not given / 2017.05.19.)
//
// https://core.telegram.org/bots/api#sendvideonote
type OptionsSendVideoNote MethodOptions

// SetDuration sets the duration value of OptionsSendVideoNote.
func (o OptionsSendVideoNote) SetDuration(duration int) OptionsSendVideoNote {
	o["duration"] = duration
	return o
}

// SetLength sets the duration value of OptionsSendVideoNote.
func (o OptionsSendVideoNote) SetLength(length int) OptionsSendVideoNote {
	o["length"] = length
	return o
}

// SetThumb sets the thumb value of OptionsSendVideoNote.
//
// thumb can be one of InputFile or string.
func (o OptionsSendVideoNote) SetThumb(thumb interface{}) OptionsSendVideoNote {
	o["thumb"] = thumb
	return o
}

// SetDisableNotification sets the disable_notification value of OptionsSendVideoNote.
func (o OptionsSendVideoNote) SetDisableNotification(disable bool) OptionsSendVideoNote {
	o["disable_notification"] = disable
	return o
}

// SetReplyToMessageID sets the reply_to_message_id value of OptionsSendVideoNote.
func (o OptionsSendVideoNote) SetReplyToMessageID(replyToMessageID int) OptionsSendVideoNote {
	o["reply_to_message_id"] = replyToMessageID
	return o
}

// SetReplyMarkup sets the reply_markup value of OptionsSendVideoNote.
//
// replyMarkup can be one of InlineKeyboardMarkup, ReplyKeyboardMarkup, ReplyKeyboardRemove, or ForceReply.
func (o OptionsSendVideoNote) SetReplyMarkup(replyMarkup interface{}) OptionsSendVideoNote {
	o["reply_markup"] = replyMarkup
	return o
}

// OptionsSendMediaGroup struct for SendMediaGroup().
//
// options include: disable_notification, and reply_to_message_id
//
// https://core.telegram.org/bots/api#sendmediagroup
type OptionsSendMediaGroup MethodOptions

// SetDisableNotification sets the disable_notification value of OptionsSendMediaGroup.
func (o OptionsSendMediaGroup) SetDisableNotification(disable bool) OptionsSendMediaGroup {
	o["disable_notification"] = disable
	return o
}

// SetReplyToMessageID sets the reply_to_message_id value of OptionsSendMediaGroup.
func (o OptionsSendMediaGroup) SetReplyToMessageID(replyToMessageID int) OptionsSendMediaGroup {
	o["reply_to_message_id"] = replyToMessageID
	return o
}

// OptionsSendLocation struct for SendLocation()
//
// options include: disable_notification, reply_to_message_id, and reply_markup.
//
// https://core.telegram.org/bots/api#sendlocation
type OptionsSendLocation MethodOptions

// SetLivePeriod sets the live_period value of OptionsSendLocation.
func (o OptionsSendLocation) SetLivePeriod(livePeriod int) OptionsSendLocation {
	o["live_period"] = livePeriod
	return o
}

// SetDisableNotification sets the disable_notification value of OptionsSendLocation.
func (o OptionsSendLocation) SetDisableNotification(disable bool) OptionsSendLocation {
	o["disable_notification"] = disable
	return o
}

// SetReplyToMessageID sets the reply_to_message_id value of OptionsSendLocation.
func (o OptionsSendLocation) SetReplyToMessageID(replyToMessageID int) OptionsSendLocation {
	o["reply_to_message_id"] = replyToMessageID
	return o
}

// SetReplyMarkup sets the reply_markup value of OptionsSendLocation.
//
// replyMarkup can be one of InlineKeyboardMarkup, ReplyKeyboardMarkup, ReplyKeyboardRemove, or ForceReply.
func (o OptionsSendLocation) SetReplyMarkup(replyMarkup interface{}) OptionsSendLocation {
	o["reply_markup"] = replyMarkup
	return o
}

// OptionsSendVenue struct for SendVenue().
//
// options include: foursquare_id, foursquare_type, disable_notification, reply_to_message_id, and reply_markup.
//
// https://core.telegram.org/bots/api#sendvenue
type OptionsSendVenue MethodOptions

// SetFoursquareID sets the foursquare_id value of OptionsSendVenue.
func (o OptionsSendVenue) SetFoursquareID(foursquareID string) OptionsSendVenue {
	o["foursquare_id"] = foursquareID
	return o
}

// SetFoursquareType sets the foursquare_type value of OptionsSendVenue.
func (o OptionsSendVenue) SetFoursquareType(foursquareType string) OptionsSendVenue {
	o["foursquare_type"] = foursquareType
	return o
}

// SetDisableNotification sets the disable_notification value of OptionsSendVenue.
func (o OptionsSendVenue) SetDisableNotification(disable bool) OptionsSendVenue {
	o["disable_notification"] = disable
	return o
}

// SetReplyToMessageID sets the reply_to_message_id value of OptionsSendVenue.
func (o OptionsSendVenue) SetReplyToMessageID(replyToMessageID int) OptionsSendVenue {
	o["reply_to_message_id"] = replyToMessageID
	return o
}

// SetReplyMarkup sets the reply_markup value of OptionsSendVenue.
//
// replyMarkup can be one of InlineKeyboardMarkup, ReplyKeyboardMarkup, ReplyKeyboardRemove, or ForceReply.
func (o OptionsSendVenue) SetReplyMarkup(replyMarkup interface{}) OptionsSendVenue {
	o["reply_markup"] = replyMarkup
	return o
}

// OptionsSendPoll struct for SendPoll().
//
// options include: disable_notification, reply_to_message_id, reply_markup, is_anonymous, type, allows_multiple_answers, correct_option_id, and is_closed.
//
// https://core.telegram.org/bots/api#sendpoll
type OptionsSendPoll MethodOptions

// SetDisableNotification sets the disable_notification value of OptionsSendPoll.
func (o OptionsSendPoll) SetDisableNotification(disable bool) OptionsSendPoll {
	o["disable_notification"] = disable
	return o
}

// SetReplyToMessageID sets the reply_to_message_id value of OptionsSendPoll.
func (o OptionsSendPoll) SetReplyToMessageID(replyToMessageID int) OptionsSendPoll {
	o["reply_to_message_id"] = replyToMessageID
	return o
}

// SetReplyMarkup sets the reply_markup value of OptionsSendPoll.
//
// replyMarkup can be one of InlineKeyboardMarkup, ReplyKeyboardMarkup, ReplyKeyboardRemove, or ForceReply.
func (o OptionsSendPoll) SetReplyMarkup(replyMarkup interface{}) OptionsSendPoll {
	o["reply_markup"] = replyMarkup
	return o
}

// SetIsAnonymous sets the is_anonymous value of OptionsSendPoll.
func (o OptionsSendPoll) SetIsAnonymous(isAnonymous bool) OptionsSendPoll {
	o["is_anonymous"] = isAnonymous
	return o
}

// SetType sets the type value of OptionsSendPoll.
func (o OptionsSendPoll) SetType(newType string) OptionsSendPoll {
	o["type"] = newType
	return o
}

// SetAllowsMultipleAnswers sets the allows_multiple_answers value of OptionsSendPoll.
func (o OptionsSendPoll) SetAllowsMultipleAnswers(allowsMultipleAnswers bool) OptionsSendPoll {
	o["allows_multiple_answers"] = allowsMultipleAnswers
	return o
}

// SetCorrectOptionID sets the correct_option_id value of OptionsSendPoll.
func (o OptionsSendPoll) SetCorrectOptionID(correctOptionID int) OptionsSendPoll {
	o["correct_option_id"] = correctOptionID
	return o
}

// SetIsClosed sets the is_closed value of OptionsSendPoll.
func (o OptionsSendPoll) SetIsClosed(isClosed bool) OptionsSendPoll {
	o["is_closed"] = isClosed
	return o
}

// OptionsStopPoll struct for StopPoll().
//
// options include: reply_markup.
//
// https://core.telegram.org/bots/api#stoppoll
type OptionsStopPoll MethodOptions

// SetReplyMarkup sets the reply_markup value of OptionsStopPoll.
func (o OptionsStopPoll) SetReplyMarkup(replyMarkup InlineKeyboardMarkup) OptionsStopPoll {
	o["reply_markup"] = replyMarkup
	return o
}

// OptionsSendDice struct for SendDice().
//
// options include: disable_notification, reply_to_message_id, and reply_markup.
//
// https://core.telegram.org/bots/api#senddice
type OptionsSendDice MethodOptions

// SetDisableNotification sets the disable_notification value of OptionsSendDice.
func (o OptionsSendDice) SetDisableNotification(disable bool) OptionsSendDice {
	o["disable_notification"] = disable
	return o
}

// SetReplyToMessageID sets the reply_to_message_id value of OptionsSendDice.
func (o OptionsSendDice) SetReplyToMessageID(replyToMessageID int) OptionsSendDice {
	o["reply_to_message_id"] = replyToMessageID
	return o
}

// SetReplyMarkup sets the reply_markup value of OptionsSendDice.
//
// replyMarkup can be one of InlineKeyboardMarkup, ReplyKeyboardMarkup, ReplyKeyboardRemove, or ForceReply.
func (o OptionsSendDice) SetReplyMarkup(replyMarkup interface{}) OptionsSendDice {
	o["reply_markup"] = replyMarkup
	return o
}

// OptionsSendContact struct for SendContact()
//
// options include: last_name, vcard, disable_notification, reply_to_message_id, and reply_markup.
//
// https://core.telegram.org/bots/api#sendcontact
type OptionsSendContact MethodOptions

// SetLastName sets the last_name value of OptionsSendContact.
func (o OptionsSendContact) SetLastName(lastName string) OptionsSendContact {
	o["last_name"] = lastName
	return o
}

// SetVCard sets the vcard value of OptionsSendContact.
func (o OptionsSendContact) SetVCard(vCard string) OptionsSendContact {
	o["vcard"] = vCard
	return o
}

// SetDisableNotification sets the disable_notification value of OptionsSendContact.
func (o OptionsSendContact) SetDisableNotification(disable bool) OptionsSendContact {
	o["disable_notification"] = disable
	return o
}

// SetReplyToMessageID sets the reply_to_message_id value of OptionsSendContact.
func (o OptionsSendContact) SetReplyToMessageID(replyToMessageID int) OptionsSendContact {
	o["reply_to_message_id"] = replyToMessageID
	return o
}

// SetReplyMarkup sets the reply_markup value of OptionsSendContact.
//
// replyMarkup can be one of InlineKeyboardMarkup, ReplyKeyboardMarkup, ReplyKeyboardRemove, or ForceReply.
func (o OptionsSendContact) SetReplyMarkup(replyMarkup interface{}) OptionsSendContact {
	o["reply_markup"] = replyMarkup
	return o
}

// OptionsGetUserProfilePhotos struct for GetUserProfilePhotos().
//
// options include: offset and limit.
//
// https://core.telegram.org/bots/api#getuserprofilephotos
type OptionsGetUserProfilePhotos MethodOptions

// SetOffset sets the offset value of OptionsGetUserProfilePhotos.
func (o OptionsGetUserProfilePhotos) SetOffset(offset int) OptionsGetUserProfilePhotos {
	o["offset"] = offset
	return o
}

// SetLimit sets the limit value of OptionsGetUserProfilePhotos.
func (o OptionsGetUserProfilePhotos) SetLimit(limit int) OptionsGetUserProfilePhotos {
	o["limit"] = limit
	return o
}

// OptionsRestrictChatMember struct for RestrictChatMember().
//
// options include: until_date
//
// https://core.telegram.org/bots/api#restrictchatmember
type OptionsRestrictChatMember MethodOptions

// SetUntilDate sets the until_date value of OptionsRestrictChatMember.
func (o OptionsRestrictChatMember) SetUntilDate(until int) OptionsRestrictChatMember {
	o["until_date"] = until
	return o
}

// OptionsPromoteChatMember struct for PromoteChatMember().
//
// options include: can_change_info, can_post_messages, can_edit_messages, can_delete_messages, can_invite_users, can_restrict_members, can_pin_messages, and can_promote_members
//
// https://core.telegram.org/bots/api#promotechatmember
type OptionsPromoteChatMember MethodOptions

// SetCanChangeInfo sets the can_change_info value of OptionsPromoteChatMember.
func (o OptionsPromoteChatMember) SetCanChangeInfo(can bool) OptionsPromoteChatMember {
	o["can_change_info"] = can
	return o
}

// SetCanPostMessages sets the can_post_messages value of OptionsPromoteChatMember.
func (o OptionsPromoteChatMember) SetCanPostMessages(can bool) OptionsPromoteChatMember {
	o["can_post_messages"] = can
	return o
}

// SetCanEditMessages sets the can_edit_messages value of OptionsPromoteChatMember.
func (o OptionsPromoteChatMember) SetCanEditMessages(can bool) OptionsPromoteChatMember {
	o["can_edit_messages"] = can
	return o
}

// SetCanDeleteMessages sets the can_delete_messages value of OptionsPromoteChatMember.
func (o OptionsPromoteChatMember) SetCanDeleteMessages(can bool) OptionsPromoteChatMember {
	o["can_delete_messages"] = can
	return o
}

// SetCanInviteUsers sets the can_invite_users value of OptionsPromoteChatMember.
func (o OptionsPromoteChatMember) SetCanInviteUsers(can bool) OptionsPromoteChatMember {
	o["can_invite_users"] = can
	return o
}

// SetCanRestrictMembers sets the can_restrict_members value of OptionsPromoteChatMember.
func (o OptionsPromoteChatMember) SetCanRestrictMembers(can bool) OptionsPromoteChatMember {
	o["can_restrict_members"] = can
	return o
}

// SetCanPinMessages sets the can_pin_messages value of OptionsPromoteChatMember.
func (o OptionsPromoteChatMember) SetCanPinMessages(can bool) OptionsPromoteChatMember {
	o["can_pin_messages"] = can
	return o
}

// SetCanPromoteMembers sets the can_promote_members value of OptionsPromoteChatMember.
func (o OptionsPromoteChatMember) SetCanPromoteMembers(can bool) OptionsPromoteChatMember {
	o["can_promote_members"] = can
	return o
}

// OptionsPinChatMessage struct for PinChatMessage
//
// options include: disable_notification
//
// https://core.telegram.org/bots/api#pinchatmessage
type OptionsPinChatMessage MethodOptions

// SetDisableNotification sets the disable_notification value of OptionsPinChatMessage.
func (o OptionsPinChatMessage) SetDisableNotification(disable bool) OptionsPinChatMessage {
	o["disable_notification"] = disable
	return o
}

// OptionsAnswerCallbackQuery struct for AnswerCallbackQuery().
//
// options include: text, show_alert, url, and cache_time
//
// https://core.telegram.org/bots/api#answercallbackquery
type OptionsAnswerCallbackQuery MethodOptions

// SetURL sets the url value of OptionsAnswerCallbackQuery.
func (o OptionsAnswerCallbackQuery) SetURL(url string) OptionsAnswerCallbackQuery {
	o["url"] = url
	return o
}

// SetCacheTime sets the cache_time value of OptionsAnswerCallbackQuery.
func (o OptionsAnswerCallbackQuery) SetCacheTime(cacheTime int) OptionsAnswerCallbackQuery {
	o["cache_time"] = cacheTime
	return o
}

// OptionsEditMessageText struct for EditMessageText()
//
// required options: chat_id + message_id (when inline_message_id is not given)
//                or inline_message_id (when chat_id & message_id is not given)
//
// other options: parse_mode, disable_web_page_preview, and reply_markup
//
// https://core.telegram.org/bots/api#editmessagetext
type OptionsEditMessageText MethodOptions

// SetIDs sets the chat_id and message_id values of OptionsEditMessageText.
func (o OptionsEditMessageText) SetIDs(chatID ChatID, messageID int) OptionsEditMessageText {
	o["chat_id"] = chatID
	o["message_id"] = messageID
	return o
}

// SetInlineMessageID sets the inline_message_id value of OptionsEditMessageText.
func (o OptionsEditMessageText) SetInlineMessageID(inlineMessageID string) OptionsEditMessageText {
	o["inline_message_id"] = inlineMessageID
	return o
}

// SetParseMode sets the parse_mode value of OptionsEditMessageText.
func (o OptionsEditMessageText) SetParseMode(parseMode ParseMode) OptionsEditMessageText {
	o["parse_mode"] = parseMode
	return o
}

// SetDisableWebPagePreview sets the disable_web_page_preview value of OptionsEditMessageText.
func (o OptionsEditMessageText) SetDisableWebPagePreview(disable bool) OptionsEditMessageText {
	o["disable_web_page_preview"] = disable
	return o
}

// SetReplyMarkup sets the reply_markup value of OptionsEditMessageText.
func (o OptionsEditMessageText) SetReplyMarkup(replyMarkup InlineKeyboardMarkup) OptionsEditMessageText {
	o["reply_markup"] = replyMarkup
	return o
}

// OptionsEditMessageCaption struct for EditMessageCaption().
//
// required options: chat_id + message_id (when inline_message_id is not given)
//                or inline_message_id (when chat_id & message_id is not given)
//
// other options: parse_mode, or reply_markup
//
// https://core.telegram.org/bots/api#editmessagecaption
type OptionsEditMessageCaption MethodOptions

// SetIDs sets the chat_id and message_id values of OptionsEditMessageCaption.
func (o OptionsEditMessageCaption) SetIDs(chatID ChatID, messageID int) OptionsEditMessageCaption {
	o["chat_id"] = chatID
	o["message_id"] = messageID
	return o
}

// SetInlineMessageID sets the inline_message_id value of OptionsEditMessageCaption.
func (o OptionsEditMessageCaption) SetInlineMessageID(inlineMessageID string) OptionsEditMessageCaption {
	o["inline_message_id"] = inlineMessageID
	return o
}

// SetCaption sets the caption value of OptionsEditMessageCaption.
func (o OptionsEditMessageCaption) SetCaption(caption string) OptionsEditMessageCaption {
	o["caption"] = caption
	return o
}

// SetParseMode sets the parse_mode value of OptionsEditMessageCaption.
func (o OptionsEditMessageCaption) SetParseMode(parseMode ParseMode) OptionsEditMessageCaption {
	o["parse_mode"] = parseMode
	return o
}

// SetReplyMarkup sets the reply_markup value of OptionsEditMessageCaption.
func (o OptionsEditMessageCaption) SetReplyMarkup(replyMarkup InlineKeyboardMarkup) OptionsEditMessageCaption {
	o["reply_markup"] = replyMarkup
	return o
}

// OptionsEditMessageMedia struct for EditMessageMedia()
//
// required options: chat_id + message_id (when inline_message_id is not given)
//                or inline_message_id (when chat_id & message_id is not given)
//
// other options: reply_markup
//
// https://core.telegram.org/bots/api#editmessagemedia
type OptionsEditMessageMedia MethodOptions

// SetIDs sets the chat_id and message_id values of OptionsEditMessageMedia.
func (o OptionsEditMessageMedia) SetIDs(chatID ChatID, messageID int) OptionsEditMessageMedia {
	o["chat_id"] = chatID
	o["message_id"] = messageID
	return o
}

// SetInlineMessageID sets the inline_message_id value of OptionsEditMessageMedia.
func (o OptionsEditMessageMedia) SetInlineMessageID(inlineMessageID string) OptionsEditMessageMedia {
	o["inline_message_id"] = inlineMessageID
	return o
}

// SetReplyMarkup sets the reply_markup value of OptionsEditMessageMedia.
func (o OptionsEditMessageMedia) SetReplyMarkup(replyMarkup InlineKeyboardMarkup) OptionsEditMessageMedia {
	o["reply_markup"] = replyMarkup
	return o
}

// OptionsEditMessageReplyMarkup struct for EditMessageReplyMarkup()
//
// required options: chat_id + message_id (when inline_message_id is not given)
//                or inline_message_id (when chat_id & message_id is not given)
//
// other options: reply_markup
//
// https://core.telegram.org/bots/api#editmessagereplymarkup
type OptionsEditMessageReplyMarkup MethodOptions

// SetIDs sets the chat_id and message_id values of OptionsEditMessageReplyMarkup.
func (o OptionsEditMessageReplyMarkup) SetIDs(chatID ChatID, messageID int) OptionsEditMessageReplyMarkup {
	o["chat_id"] = chatID
	o["message_id"] = messageID
	return o
}

// SetInlineMessageID sets the inline_message_id value of OptionsEditMessageReplyMarkup.
func (o OptionsEditMessageReplyMarkup) SetInlineMessageID(inlineMessageID string) OptionsEditMessageReplyMarkup {
	o["inline_message_id"] = inlineMessageID
	return o
}

// SetReplyMarkup sets the reply_markup value of OptionsEditMessageReplyMarkup.
func (o OptionsEditMessageReplyMarkup) SetReplyMarkup(replyMarkup InlineKeyboardMarkup) OptionsEditMessageReplyMarkup {
	o["reply_markup"] = replyMarkup
	return o
}

// OptionsEditMessageLiveLocation struct for EditMessageLiveLocation()
//
// required options: chat_id + message_id (when inline_message_id is not given)
//                or inline_message_id (when chat_id & message_id is not given)
//
// other options: reply_markup
//
// https://core.telegram.org/bots/api#editmessagelivelocation
type OptionsEditMessageLiveLocation MethodOptions

// SetIDs sets the chat_id and message_id values of OptionsEditMessageLiveLocation.
func (o OptionsEditMessageLiveLocation) SetIDs(chatID ChatID, messageID int) OptionsEditMessageLiveLocation {
	o["chat_id"] = chatID
	o["message_id"] = messageID
	return o
}

// SetInlineMessageID sets the inline_message_id value of OptionsEditMessageLiveLocation.
func (o OptionsEditMessageLiveLocation) SetInlineMessageID(inlineMessageID string) OptionsEditMessageLiveLocation {
	o["inline_message_id"] = inlineMessageID
	return o
}

// SetReplyMarkup sets the reply_markup value of OptionsEditMessageLiveLocation.
func (o OptionsEditMessageLiveLocation) SetReplyMarkup(replyMarkup InlineKeyboardMarkup) OptionsEditMessageLiveLocation {
	o["reply_markup"] = replyMarkup
	return o
}

// OptionsStopMessageLiveLocation struct for StopMessageLiveLocation()
//
// required options: chat_id + message_id (when inline_message_id is not given)
//                or inline_message_id (when chat_id & message_id is not given)
//
// other options: reply_markup
//
// https://core.telegram.org/bots/api#stopmessagelivelocation
type OptionsStopMessageLiveLocation MethodOptions

// SetIDs sets the chat_id and message_id values of OptionsStopMessageLiveLocation.
func (o OptionsStopMessageLiveLocation) SetIDs(chatID ChatID, messageID int) OptionsStopMessageLiveLocation {
	o["chat_id"] = chatID
	o["message_id"] = messageID
	return o
}

// SetInlineMessageID sets the inline_message_id value of OptionsStopMessageLiveLocation.
func (o OptionsStopMessageLiveLocation) SetInlineMessageID(inlineMessageID string) OptionsStopMessageLiveLocation {
	o["inline_message_id"] = inlineMessageID
	return o
}

// SetReplyMarkup sets the reply_markup value of OptionsStopMessageLiveLocation.
func (o OptionsStopMessageLiveLocation) SetReplyMarkup(replyMarkup InlineKeyboardMarkup) OptionsStopMessageLiveLocation {
	o["reply_markup"] = replyMarkup
	return o
}

// OptionsAnswerInlineQuery struct for AnswerInlineQuery().
//
// options include: cache_time, is_personal, next_offset, switch_pm_text, and switch_pm_parameter.
//
// https://core.telegram.org/bots/api#answerinlinequery
type OptionsAnswerInlineQuery MethodOptions

// SetCacheTime sets the cache_time value of OptionsAnswerInlineQuery.
func (o OptionsAnswerInlineQuery) SetCacheTime(cacheTime int) OptionsAnswerInlineQuery {
	o["cache_time"] = cacheTime
	return o
}

// SetIsPersonal sets the is_personal value of OptionsAnswerInlineQuery.
func (o OptionsAnswerInlineQuery) SetIsPersonal(isPersonal bool) OptionsAnswerInlineQuery {
	o["is_personal"] = isPersonal
	return o
}

// SetNextOffset sets the next_offset value of OptionsAnswerInlineQuery.
func (o OptionsAnswerInlineQuery) SetNextOffset(nextOffset string) OptionsAnswerInlineQuery {
	o["next_offset"] = nextOffset
	return o
}

// SetSwitchPmText sets the switch_pm_text value of OptionsAnswerInlineQuery.
func (o OptionsAnswerInlineQuery) SetSwitchPmText(switchPmText string) OptionsAnswerInlineQuery {
	o["switch_pm_text"] = switchPmText
	return o
}

// SetSwitchPmParameter sets the switch_pm_parameter value of OptionsAnswerInlineQuery.
func (o OptionsAnswerInlineQuery) SetSwitchPmParameter(switchPmParam string) OptionsAnswerInlineQuery {
	o["switch_pm_parameter"] = switchPmParam
	return o
}

// OptionsSendInvoice struct for SendInvoice().
//
// options include: provider_data, photo_url, photo_size, photo_width, photo_height, need_name, need_phone_number, need_email, need_shipping_address, send_phone_number_to_provider, send_email_to_provider, is_flexible, disable_notification, reply_to_message_id, and reply_markup
//
// https://core.telegram.org/bots/api#sendinvoice
type OptionsSendInvoice MethodOptions

// SetProviderData sets the provider_data value of OptionsSendInvoice.
func (o OptionsSendInvoice) SetProviderData(providerData string) OptionsSendInvoice {
	o["provider_data"] = providerData
	return o
}

// SetPhotoURL sets the photo_url value of OptionsSendInvoice.
func (o OptionsSendInvoice) SetPhotoURL(photoURL string) OptionsSendInvoice {
	o["photo_url"] = photoURL
	return o
}

// SetPhotoSize sets the photo_size value of OptionsSendInvoice.
func (o OptionsSendInvoice) SetPhotoSize(photoSize int) OptionsSendInvoice {
	o["photo_size"] = photoSize
	return o
}

// SetPhotoWidth sets the photoWidth value of OptionsSendInvoice.
func (o OptionsSendInvoice) SetPhotoWidth(photoWidth int) OptionsSendInvoice {
	o["photo_width"] = photoWidth
	return o
}

// SetPhotoHeight sets the photo_height value of OptionsSendInvoice.
func (o OptionsSendInvoice) SetPhotoHeight(photoHeight int) OptionsSendInvoice {
	o["photo_height"] = photoHeight
	return o
}

// SetNeedName sets the need_name value of OptionsSendInvoice.
func (o OptionsSendInvoice) SetNeedName(needName bool) OptionsSendInvoice {
	o["need_name"] = needName
	return o
}

// SetNeedPhoneNumber sets the need_phone_number value of OptionsSendInvoice.
func (o OptionsSendInvoice) SetNeedPhoneNumber(needPhoneNumber bool) OptionsSendInvoice {
	o["need_phone_number"] = needPhoneNumber
	return o
}

// SetNeedEmail sets the need_email value of OptionsSendInvoice.
func (o OptionsSendInvoice) SetNeedEmail(needEmail bool) OptionsSendInvoice {
	o["need_email"] = needEmail
	return o
}

// SetNeedShippingAddress sets the need_shipping_address value of OptionsSendInvoice.
func (o OptionsSendInvoice) SetNeedShippingAddress(needShippingAddr bool) OptionsSendInvoice {
	o["need_shipping_address"] = needShippingAddr
	return o
}

// SetSendPhoneNumberToProvider sets the send_phone_number_to_provider value of OptionsSendInvoice.
func (o OptionsSendInvoice) SetSendPhoneNumberToProvider(sendPhoneNumberToProvider bool) OptionsSendInvoice {
	o["send_phone_number_to_provider"] = sendPhoneNumberToProvider
	return o
}

// SetSendEmailToProvider sets the send_email_to_provider value of OptionsSendInvoice.
func (o OptionsSendInvoice) SetSendEmailToProvider(sendEmailToProvider bool) OptionsSendInvoice {
	o["send_email_to_provider"] = sendEmailToProvider
	return o
}

// SetIsFlexible sets the is_flexible value of OptionsSendInvoice.
func (o OptionsSendInvoice) SetIsFlexible(isFlexible bool) OptionsSendInvoice {
	o["is_flexible"] = isFlexible
	return o
}

// SetDisableNotification sets the disable_notification value of OptionsSendInvoice.
func (o OptionsSendInvoice) SetDisableNotification(disable bool) OptionsSendInvoice {
	o["disable_notification"] = disable
	return o
}

// SetReplyToMessageID sets the reply_to_message_id value of OptionsSendInvoice.
func (o OptionsSendInvoice) SetReplyToMessageID(replyToMessageID int) OptionsSendInvoice {
	o["reply_to_message_id"] = replyToMessageID
	return o
}

// SetReplyMarkup sets the reply_markup value of OptionsSendInvoice.
func (o OptionsSendInvoice) SetReplyMarkup(replyMarkup InlineKeyboardMarkup) OptionsSendInvoice {
	o["reply_markup"] = replyMarkup
	return o
}

// OptionsSendGame struct for SendGame()
//
// options include: disable_notification, reply_to_message_id, and reply_markup.
//
// https://core.telegram.org/bots/api#sendgame
type OptionsSendGame MethodOptions

// SetDisableNotification sets the disable_notification value of OptionsSendGame.
func (o OptionsSendGame) SetDisableNotification(disable bool) OptionsSendGame {
	o["disable_notification"] = disable
	return o
}

// SetReplyToMessageID sets the reply_to_message_id value of OptionsSendGame.
func (o OptionsSendGame) SetReplyToMessageID(replyToMessageID int) OptionsSendGame {
	o["reply_to_message_id"] = replyToMessageID
	return o
}

// SetReplyMarkup sets the reply_markup value of OptionsSendGame.
func (o OptionsSendGame) SetReplyMarkup(replyMarkup InlineKeyboardMarkup) OptionsSendGame {
	o["reply_markup"] = replyMarkup
	return o
}

// OptionsSetGameScore struct for SetGameScore().
//
// required options: chat_id + message_id (when inline_message_id is not given)
//                or inline_message_id (when chat_id & message_id is not given)
//
// other options: force, and disable_edit_message
//
// https://core.telegram.org/bots/api#setgamescore
type OptionsSetGameScore MethodOptions

// SetForce sets the force value of OptionsSetGameScore.
func (o OptionsSetGameScore) SetForce(force bool) OptionsSetGameScore {
	o["force"] = force
	return o
}

// SetDisableEditMessage sets the disable_edit_message value of OptionsSetGameScore.
func (o OptionsSetGameScore) SetDisableEditMessage(disableEditMessage bool) OptionsSetGameScore {
	o["disable_edit_message"] = disableEditMessage
	return o
}

// SetIDs sets the chat_id and message_id values of OptionsSetGameScore.
func (o OptionsSetGameScore) SetIDs(chatID ChatID, messageID int) OptionsSetGameScore {
	o["chat_id"] = chatID
	o["message_id"] = messageID
	return o
}

// SetInlineMessageID sets the inline_message_id value of OptionsSetGameScore.
func (o OptionsSetGameScore) SetInlineMessageID(inlineMessageID string) OptionsSetGameScore {
	o["inline_message_id"] = inlineMessageID
	return o
}

// OptionsGetGameHighScores struct for GetGameHighScores().
//
// required options: chat_id + message_id (when inline_message_id is not given)
//                or inline_message_id (when chat_id & message_id is not given)
//
// https://core.telegram.org/bots/api#getgamehighscores
type OptionsGetGameHighScores MethodOptions

// SetIDs sets the chat_id and message_id values of OptionsGetGameHighScores.
func (o OptionsGetGameHighScores) SetIDs(chatID ChatID, messageID int) OptionsGetGameHighScores {
	o["chat_id"] = chatID
	o["message_id"] = messageID
	return o
}

// SetInlineMessageID sets the inline_message_id value of OptionsGetGameHighScores.
func (o OptionsGetGameHighScores) SetInlineMessageID(inlineMessageID string) OptionsGetGameHighScores {
	o["inline_message_id"] = inlineMessageID
	return o
}
