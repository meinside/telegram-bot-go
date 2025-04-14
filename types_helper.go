package telegrambot

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

////////////////////////////////
// Helper functions for errors
//

// converts given errStr to custom error
func strToErr(errStr string) error {
	switch {
	case strings.Contains(errStr, "Unauthorized"):
		return ErrUnauthorized{
			baseError: baseError{
				Message: errStr,
			},
		}
	case strings.Contains(errStr, "Bad Request: chat not found"):
		return ErrChatNotFound{
			baseError: baseError{
				Message: errStr,
			},
		}
	case strings.Contains(errStr, "Bad Request: user not found"):
		return ErrUserNotFound{
			baseError: baseError{
				Message: errStr,
			},
		}
	case strings.Contains(errStr, "Forbidden: user is deactivated"):
		return ErrUserDeactivated{
			baseError: baseError{
				Message: errStr,
			},
		}
	case strings.Contains(errStr, "Forbidden: bot was kicked"):
		return ErrBotKicked{
			baseError: baseError{
				Message: errStr,
			},
		}
	case strings.Contains(errStr, "Forbidden: bot blocked by user"):
		return ErrBotBlockedByUser{
			baseError: baseError{
				Message: errStr,
			},
		}
	case strings.Contains(errStr, "Forbidden: bot can't send messages to bots"):
		return ErrBotCantSendToBots{
			baseError: baseError{
				Message: errStr,
			},
		}
	case strings.Contains(errStr, "Bad request: Message not modified"):
		return ErrMessageNotModified{
			baseError: baseError{
				Message: errStr,
			},
		}
	case strings.Contains(errStr, "Bad request: Group migrated to supergroup"):
		return ErrGroupMigratedToSupergroup{
			baseError: baseError{
				Message: errStr,
			},
		}
	case strings.Contains(errStr, "Bad request: Invalid file id"):
		return ErrInvalidFileID{
			baseError: baseError{
				Message: errStr,
			},
		}
	case strings.Contains(errStr, "Conflict: Terminated by other long poll"):
		return ErrConflictedLongPoll{
			baseError: baseError{
				Message: errStr,
			},
		}
	case strings.Contains(errStr, "Conflict: can't use getUpdates method while webhook is active; use deleteWebhook to delete the webhook firs"):
		return ErrConflictedWebHook{
			baseError: baseError{
				Message: errStr,
			},
		}
	case strings.Contains(errStr, "Bad request: Wrong parameter action in request"):
		return ErrWrongParameterAction{
			baseError: baseError{
				Message: errStr,
			},
		}
	case strings.Contains(errStr, "Bad Request: message text is empty"):
		return ErrMessageEmpty{
			baseError: baseError{
				Message: errStr,
			},
		}
	case strings.Contains(errStr, "Bad Request: message is too long"):
		return ErrMessageTooLong{
			baseError: baseError{
				Message: errStr,
			},
		}
	case strings.Contains(errStr, "Bad Request: message can't be edited"):
		return ErrMessageCantBeEdited{
			baseError: baseError{
				Message: errStr,
			},
		}
	case strings.Contains(errStr, "Too many requests"):
		return ErrTooManyRequests{
			baseError: baseError{
				Message: errStr,
			},
		}
	case strings.Contains(errStr, "failed to parse json"):
		return ErrJSONParseFailed{
			baseError: baseError{
				Message: errStr,
			},
		}
	}

	// TODO: handle more errors here

	// fallback
	return ErrUnclassified{
		baseError: baseError{
			Message: errStr,
		},
	}
}

////////////////////////////////
// Helper functions for Update
//

// HasMessage checks if Update has Message.
func (u *Update) HasMessage() bool {
	return u.Message != nil
}

// HasEditedMessage checks if Update has EditedMessage.
func (u *Update) HasEditedMessage() bool {
	return u.EditedMessage != nil
}

// HasMediaGroup checks if Update is a part of a grouped media
func (u *Update) HasMediaGroup() (exists bool) {
	var message *Message
	if u.HasMessage() {
		message = u.Message
	} else if u.HasEditedMessage() {
		message = u.EditedMessage
	}

	return message != nil && message.MediaGroupID != nil
}

// MediaGroupID returns the media group id (nil if none)
func (u *Update) MediaGroupID() *string {
	var message *Message
	if u.HasMessage() {
		message = u.Message
	} else if u.HasEditedMessage() {
		message = u.EditedMessage
	}

	return message.MediaGroupID
}

// HasChannelPost checks if Update has ChannelPost.
func (u *Update) HasChannelPost() bool {
	return u.ChannelPost != nil
}

// HasEditedChannelPost checks if Update has EditedChannelPost.
func (u *Update) HasEditedChannelPost() bool {
	return u.EditedChannelPost != nil
}

// HasInlineQuery checks if Update has InlineQuery
func (u *Update) HasInlineQuery() bool {
	return u.InlineQuery != nil
}

// HasChosenInlineResult checks if Update has ChosenInlineResult
func (u *Update) HasChosenInlineResult() bool {
	return u.ChosenInlineResult != nil
}

// HasCallbackQuery checks if Update has CallbackQuery
func (u *Update) HasCallbackQuery() bool {
	return u.CallbackQuery != nil
}

// HasShippingQuery checks if Update has ShippingQuery
func (u *Update) HasShippingQuery() bool {
	return u.ShippingQuery != nil
}

// HasPreCheckoutQuery checks if Update has PreCheckoutQuery
func (u *Update) HasPreCheckoutQuery() bool {
	return u.PreCheckoutQuery != nil
}

// HasPoll checks if Update has Poll
func (u *Update) HasPoll() bool {
	return u.Poll != nil
}

// HasPollAnswer checks if Update has PollAnswer.
func (u *Update) HasPollAnswer() bool {
	return u.PollAnswer != nil
}

// HasMyChatMember checks if Update has MyChatMember.
func (u *Update) HasMyChatMember() bool {
	return u.MyChatMember != nil
}

// HasChatMember checks if Update has ChatMember.
func (u *Update) HasChatMember() bool {
	return u.ChatMember != nil
}

// HasChatJoinRequest checks if Update has ChatJoinRequest.
func (u *Update) HasChatJoinRequest() bool {
	return u.ChatJoinRequest != nil
}

// GetFrom returns the `from` value from Update.
//
// NOTE: `Poll` type doesn't have `from` property.
func (u *Update) GetFrom() *User {
	if u.HasMessage() {
		return u.Message.From
	} else if u.HasEditedMessage() {
		return u.EditedMessage.From
	} else if u.HasChannelPost() {
		return u.ChannelPost.From
	} else if u.HasEditedChannelPost() {
		return u.EditedChannelPost.From
	} else if u.HasInlineQuery() {
		return &u.InlineQuery.From
	} else if u.HasChosenInlineResult() {
		return &u.ChosenInlineResult.From
	} else if u.HasCallbackQuery() {
		return &u.CallbackQuery.From
	} else if u.HasShippingQuery() {
		return &u.ShippingQuery.From
	} else if u.HasPreCheckoutQuery() {
		return &u.PreCheckoutQuery.From
	} else if u.HasPoll() {
		return nil
	} else if u.HasPollAnswer() {
		return u.PollAnswer.User
	} else if u.HasMyChatMember() {
		return &u.MyChatMember.From
	} else if u.HasChatMember() {
		return &u.ChatMember.From
	} else if u.HasChatJoinRequest() {
		return &u.ChatJoinRequest.From
	}

	return nil
}

// GetMessage returns usable message property from Update.
func (u *Update) GetMessage() (message *Message, edited bool) {
	if u.HasMessage() {
		return u.Message, false
	} else if u.HasEditedMessage() {
		return u.EditedMessage, true
	}

	return nil, false
}

// GetChannelPost returns usable channel post property from Update.
func (u *Update) GetChannelPost() (post *Message, edited bool) {
	if u.HasChannelPost() {
		return u.ChannelPost, false
	} else if u.HasEditedChannelPost() {
		return u.EditedChannelPost, true
	}

	return nil, false
}

////////////////////////////////
// Helper functions for User
//

// InlineLink generates an inline link for User
func (u User) InlineLink() string {
	return fmt.Sprintf("tg://user?id=%d", u.ID)
}

////////////////////////////////
// Helper functions for Chat
//

////////////////////////////////
// Helper functions for Message
//

// HasForwardFrom checks if Message has Forward.
func (m *Message) HasForwardFrom() bool {
	return m.ForwardOrigin != nil &&
		(m.ForwardOrigin.SenderUser != nil || m.ForwardOrigin.SenderUserName != nil) &&
		m.ForwardOrigin.Date > 0
}

// HasForwardFromChat checks if Message has Forward from chat.
func (m *Message) HasForwardFromChat() bool {
	return m.ForwardOrigin != nil &&
		m.ForwardOrigin.SenderChat != nil &&
		m.ForwardOrigin.Date > 0
}

// HasReplyToMessage checks if Message has ReplyToMessage.
func (m *Message) HasReplyToMessage() bool {
	return m.ReplyToMessage != nil
}

// HasQuote checks if Message has Quote.
func (m *Message) HasQuote() bool {
	return m.Quote != nil
}

// HasReplyToStory checks if Message has ReplyToStory.
func (m *Message) HasReplyToStory() bool {
	return m.ReplyToStory != nil
}

// IsBot checks if Message is from bot.
func (m *Message) IsBot() bool {
	return m.ViaBot != nil && m.ViaBot.IsBot
}

// HasEditDate checks if Message has EditDate.
func (m *Message) HasEditDate() bool {
	return m.EditDate != nil
}

// HasText checks if Message has Text.
func (m *Message) HasText() bool {
	return m.Text != nil
}

// HasMessageEntities checks if Message has MessageEntities
func (m *Message) HasMessageEntities() bool {
	return len(m.Entities) > 0
}

// HasAnimation checks if Message has Animation.
func (m *Message) HasAnimation() bool {
	return m.Animation != nil
}

// HasAudio checks if Message has Audio.
func (m *Message) HasAudio() bool {
	return m.Audio != nil
}

// HasDocument checks if Message has Document.
func (m *Message) HasDocument() bool {
	return m.Document != nil
}

// HasPhoto checks if Message has Photo.
func (m *Message) HasPhoto() bool {
	return len(m.Photo) > 0
}

// LargestPhoto returns a photo with the largest file size.
func (m *Message) LargestPhoto() PhotoSize {
	if !m.HasPhoto() {
		return PhotoSize{}
	}

	maxIndex := 0
	for i, photo := range m.Photo {
		if photo.FileSize != nil && m.Photo[maxIndex].FileSize != nil && *photo.FileSize > *m.Photo[maxIndex].FileSize {
			maxIndex = i
		}
	}

	return m.Photo[maxIndex]
}

// HasSticker checks if Message has Sticker.
func (m *Message) HasSticker() bool {
	return m.Sticker != nil
}

// HasStory checks if Message has Story.
func (m *Message) HasStory() bool {
	return m.Story != nil
}

// HasVideo checks if Message has Video.
func (m *Message) HasVideo() bool {
	return m.Video != nil
}

// HasVideoNote checks if Message has VideoNote.
func (m *Message) HasVideoNote() bool {
	return m.VideoNote != nil
}

// HasVoice checks if Message has Voice.
func (m *Message) HasVoice() bool {
	return m.Voice != nil
}

// HasCaption checks if Message has Caption.
func (m *Message) HasCaption() bool {
	return m.Caption != nil
}

// HasCaptionEntities checks if Message has CaptionEntities
func (m *Message) HasCaptionEntities() bool {
	return len(m.CaptionEntities) > 0
}

// HasContact checks if Message has Contact.
func (m *Message) HasContact() bool {
	return m.Contact != nil
}

// HasDice checks if Message has Dice.
func (m *Message) HasDice() bool {
	return m.Dice != nil
}

// HasGame checks if Message has Game.
func (m *Message) HasGame() bool {
	return m.Game != nil
}

// HasPoll checks if Message has Poll.
func (m *Message) HasPoll() bool {
	return m.Poll != nil
}

// HasVenue checks if Message has Venue.
func (m *Message) HasVenue() bool {
	return m.Venue != nil
}

// HasLocation checks if Message has Location.
func (m *Message) HasLocation() bool {
	return m.Location != nil
}

// HasNewChatMembers checks if Message has NewChatParticipant.
func (m *Message) HasNewChatMembers() bool {
	return len(m.NewChatMembers) > 0
}

// HasLeftChatMember checks if Message has LeftChatParticipant.
func (m *Message) HasLeftChatMember() bool {
	return m.LeftChatMember != nil
}

// HasNewChatTitle checks if Message has NewChatTitle.
func (m *Message) HasNewChatTitle() bool {
	return m.NewChatTitle != nil
}

// HasNewChatPhoto checks if Message has NewChatPhoto.
func (m *Message) HasNewChatPhoto() bool {
	return len(m.NewChatPhoto) > 0
}

// HasMessageAutoDeleteTimerChanged checks if Message has MessageAutoDeleteTimerChanged.
func (m *Message) HasMessageAutoDeleteTimerChanged() bool {
	return m.MessageAutoDeleteTimerChanged != nil
}

// HasPinnedMessage checks if Message has PinnedMessage.
func (m *Message) HasPinnedMessage() bool {
	return m.PinnedMessage != nil
}

// HasInvoice checks if Message has Invoice.
func (m *Message) HasInvoice() bool {
	return m.Invoice != nil
}

// HasSuccessfulPayment checks if Message has SuccessfulPayment.
func (m *Message) HasSuccessfulPayment() bool {
	return m.SuccessfulPayment != nil
}

// HasUsersShared checks if Message has UsersShared.
func (m *Message) HasUsersShared() bool {
	return m.UsersShared != nil
}

// HasChatShared checks if Message has ChatShared.
func (m *Message) HasChatShared() bool {
	return m.ChatShared != nil
}

// HasConnectedWebsite checks if Message has ConnectedWebsite.
func (m *Message) HasConnectedWebsite() bool {
	return m.ConnectedWebsite != nil
}

// HasWriteAccessAllowed checks if Message has WriteAccessAllowed.
func (m *Message) HasWriteAccessAllowed() bool {
	return m.WriteAccessAllowed != nil
}

// HasProximityAlertTriggered checks if Message has ProximityAlertTriggered.
func (m *Message) HasProximityAlertTriggered() bool {
	return m.ProximityAlertTriggered != nil
}

// HasBoostAdded checks if Message has BoostAdded.
func (m *Message) HasBoostAdded() bool {
	return m.BoostAdded != nil
}

// HasChatBackgroundSet checks if Message has ChatBackgroundSet.
func (m *Message) HasChatBackgroundSet() bool {
	return m.ChatBackgroundSet != nil
}

// HasForumTopicCreated checks if Message has ForumTopicCreated.
func (m *Message) HasForumTopicCreated() bool {
	return m.ForumTopicCreated != nil
}

// HasForumTopicEdited checks if Message has ForumTopicEdited.
func (m *Message) HasForumTopicEdited() bool {
	return m.ForumTopicEdited != nil
}

// HasForumTopicClosed checks if Message has ForumTopicClosed.
func (m *Message) HasForumTopicClosed() bool {
	return m.ForumTopicClosed != nil
}

// HasForumTopicReopened checks if Message has ForumTopicReopened.
func (m *Message) HasForumTopicReopened() bool {
	return m.ForumTopicReopened != nil
}

// HasGeneralForumTopicHidden checks if Message has GeneralForumTopicHidden.
func (m *Message) HasGeneralForumTopicHidden() bool {
	return m.GeneralForumTopicHidden != nil
}

// HasGeneralForumTopicUnhidden checks if Message has GeneralForumTopicUnhidden.
func (m *Message) HasGeneralForumTopicUnhidden() bool {
	return m.GeneralForumTopicUnhidden != nil
}

// HasGiveawayCreated checks if Message has GiveawayCreated.
func (m *Message) HasGiveawayCreated() bool {
	return m.GiveawayCreated != nil
}

// HasGiveaway checks if Message has Giveaway.
func (m *Message) HasGiveaway() bool {
	return m.Giveaway != nil
}

// HasGiveawayWinners checks if Message has GiveawayWinners.
func (m *Message) HasGiveawayWinners() bool {
	return m.GiveawayWinners != nil
}

// HasGiveawayCompleted checks if Message has GiveawayCompleted.
func (m *Message) HasGiveawayCompleted() bool {
	return m.GiveawayCompleted != nil
}

// HasVideoChatScheduled checks if Message has VideoChatScheduled.
func (m *Message) HasVideoChatScheduled() bool {
	return m.VideoChatScheduled != nil
}

// HasVideoChatStarted checks if Message has VideoChatStarted.
func (m *Message) HasVideoChatStarted() bool {
	return m.VideoChatStarted != nil
}

// HasVideoChatEnded checks if Message has VideoChatEnded.
func (m *Message) HasVideoChatEnded() bool {
	return m.VideoChatEnded != nil
}

// HasVideoChatParticipantsInvited checks if Message has VideoChatParticipantsInvited.
func (m *Message) HasVideoChatParticipantsInvited() bool {
	return m.VideoChatParticipantsInvited != nil
}

// HasWebAppData checks if Message has WebAppData.
func (m *Message) HasWebAppData() bool {
	return m.WebAppData != nil
}

// HasReplyMarkup checks if Message has ReplyMarkup.
func (m *Message) HasReplyMarkup() bool {
	return m.ReplyMarkup != nil
}

////////////////////////////////
// Helper functions for MaybeInaccessibleMessage
//

// IsInaccessible returns true if it is inaccessible
func (m *MaybeInaccessibleMessage) IsInaccessible() bool {
	return m.Date == 0
}

// AsMessage returns its value as `Message` or `InaccessibleMessage`
func (m *MaybeInaccessibleMessage) AsMessage() (*Message, *InaccessibleMessage) {
	if !m.IsInaccessible() {
		return (*Message)(m), nil
	} else {
		return nil, &InaccessibleMessage{
			Chat:      m.Chat,
			MessageID: m.MessageID,
			Date:      0,
		}
	}
}

////////////////////////////////
// Helper functions for InlineQuery
//

// nothing yet

////////////////////////////////
// Helper functions for ChosenInlineResult
//

// nothing yet

////////////////////////////////
// Helper functions for KeyboardButton and InlineKeyboardButton
//

// NewKeyboardButtons is a helper function for generating an array of KeyboardButtons
func NewKeyboardButtons(texts ...string) []KeyboardButton {
	keyboards := []KeyboardButton{}

	for _, text := range texts {
		keyboards = append(keyboards, KeyboardButton{
			Text: text,
		})
	}

	return keyboards
}

// NewInlineKeyboardButtonsWithURL is a helper function
// for generating an array of InlineKeyboardButtons with urls
func NewInlineKeyboardButtonsWithURL(values map[string]string) []InlineKeyboardButton {
	keyboards := []InlineKeyboardButton{}

	for text, url := range values {
		u := url
		keyboards = append(keyboards, InlineKeyboardButton{
			Text: text,
			URL:  &u,
		})
	}

	return keyboards
}

// NewInlineKeyboardButtonsWithCallbackData is a helper function
// for generating an array of InlineKeyboardButtons with callback data
func NewInlineKeyboardButtonsWithCallbackData(values map[string]string) []InlineKeyboardButton {
	return NewInlineKeyboardButtonsAsColumnsWithCallbackData(values)
}

// NewInlineKeyboardButtonsAsColumnsWithCallbackData is a helper function
// for generating an array of InlineKeyboardButtons (as columns) with callback data
func NewInlineKeyboardButtonsAsColumnsWithCallbackData(values map[string]string) []InlineKeyboardButton {
	keyboards := []InlineKeyboardButton{}

	for text, data := range values {
		callbackData := data
		keyboards = append(keyboards, InlineKeyboardButton{
			Text:         text,
			CallbackData: &callbackData,
		})
	}

	return keyboards
}

// NewInlineKeyboardButtonsAsRowsWithCallbackData is a helper function
// for generating an array of InlineKeyboardButtons (as rows) with callback data
func NewInlineKeyboardButtonsAsRowsWithCallbackData(values map[string]string) [][]InlineKeyboardButton {
	keyboards := [][]InlineKeyboardButton{}

	for text, data := range values {
		callbackData := data
		keyboards = append(keyboards, []InlineKeyboardButton{
			{
				Text:         text,
				CallbackData: &callbackData,
			},
		})
	}

	return keyboards
}

// NewInlineKeyboardButtonsWithSwitchInlineQuery is a helper function
// for generating an array of InlineKeyboardButtons with switch inline query
func NewInlineKeyboardButtonsWithSwitchInlineQuery(values map[string]string) []InlineKeyboardButton {
	keyboards := []InlineKeyboardButton{}

	for text, switchInlineQuery := range values {
		query := switchInlineQuery
		keyboards = append(keyboards, InlineKeyboardButton{
			Text:              text,
			SwitchInlineQuery: &query,
		})
	}

	return keyboards
}

////////////////////////////////
// Helper functions for LoginURL

// NewLoginURL returns a new LoginURL.
func NewLoginURL(url string) LoginURL {
	return LoginURL{
		URL: url,
	}
}

// SetForwardText sets the `forward_text` value of LoginURL.
func (u LoginURL) SetForwardText(text string) LoginURL {
	u.ForwardText = &text
	return u
}

// SetBotUsername sets the `bot_username` value of LoginURL.
func (u LoginURL) SetBotUsername(username string) LoginURL {
	u.BotUsername = &username
	return u
}

// SetRequestWriteAccess sets the `request_write_access` value of LoginURL.
func (u LoginURL) SetRequestWriteAccess(request bool) LoginURL {
	u.RequestWriteAccess = &request
	return u
}

////////////////////////////////
// Helper functions for SwitchInlineQueryChosenChat

// NewSwitchInlineQueryChosenChat returns a new SwitchInlineQueryChosenChat.
func NewSwitchInlineQueryChosenChat() SwitchInlineQueryChosenChat {
	return SwitchInlineQueryChosenChat{}
}

// SetQuery sets the `query` value of SwitchInlineQueryChosenChat.
func (c SwitchInlineQueryChosenChat) SetQuery(query string) SwitchInlineQueryChosenChat {
	c.Query = &query
	return c
}

// SetAllowUserChats sets the `allow_user_chats` value of SwitchInlineQueryChosenChat.
func (c SwitchInlineQueryChosenChat) SetAllowUserChats(allow bool) SwitchInlineQueryChosenChat {
	c.AllowUserChats = &allow
	return c
}

// SetAllowBotChats sets the `allow_bot_chats` value of SwitchInlineQueryChosenChat.
func (c SwitchInlineQueryChosenChat) SetAllowBotChats(allow bool) SwitchInlineQueryChosenChat {
	c.AllowBotChats = &allow
	return c
}

// SetAllowGroupChats sets the `allow_group_chats` value of SwitchInlineQueryChosenChat.
func (c SwitchInlineQueryChosenChat) SetAllowGroupChats(allow bool) SwitchInlineQueryChosenChat {
	c.AllowGroupChats = &allow
	return c
}

// SetAllowChannelChats sets the `allow_channel_chats` value of SwitchInlineQueryChosenChat.
func (c SwitchInlineQueryChosenChat) SetAllowChannelChats(allow bool) SwitchInlineQueryChosenChat {
	c.AllowChannelChats = &allow
	return c
}

////////////////////////////////
// Helper functions for ChatAdministratorRights

// NewChatAdministratorRights returns a new ChatAdministratorRights.
func NewChatAdministratorRights(isAnonymous, canManageChat, canDeleteMessages, canManageVideoChats, canRestrictMembers, canPromoteMembers, canChangeInfo, canInviteUsers, canPostStories, canEditStories, canDeleteStories bool) ChatAdministratorRights {
	return ChatAdministratorRights{
		IsAnonymous:         isAnonymous,
		CanManageChat:       canManageChat,
		CanDeleteMessages:   canDeleteMessages,
		CanManageVideoChats: canManageVideoChats,
		CanRestrictMembers:  canRestrictMembers,
		CanPromoteMembers:   canPromoteMembers,
		CanChangeInfo:       canChangeInfo,
		CanInviteUsers:      canInviteUsers,
		CanPostStories:      canPostStories,
		CanEditStories:      canEditStories,
		CanDeleteStories:    canDeleteStories,
	}
}

// SetCanPostMessages sets the `can_post_messages` value of ChatAdministratorRights.
func (r ChatAdministratorRights) SetCanPostMessages(canPostMessages bool) ChatAdministratorRights {
	r.CanPostMessages = &canPostMessages
	return r
}

// SetCanEditMessages sets the `can_edit_messages` value of ChatAdministratorRights.
func (r ChatAdministratorRights) SetCanEditMessages(canEditMessages bool) ChatAdministratorRights {
	r.CanEditMessages = &canEditMessages
	return r
}

// SetCanPinMessages sets the `can_pin_messages` value of ChatAdministratorRights.
func (r ChatAdministratorRights) SetCanPinMessages(canPinMessages bool) ChatAdministratorRights {
	r.CanPinMessages = &canPinMessages
	return r
}

// SetCanManageTopics sets the `can_manage_topics` value of ChatAdministratorRights.
func (r ChatAdministratorRights) SetCanManageTopics(canManageTopics bool) ChatAdministratorRights {
	r.CanManageTopics = &canManageTopics
	return r
}

////////////////////////////////
// Helper functions for ChatPermissions

// NewChatPermissions returns a new ChatPermissions. (all permissions are false)
func NewChatPermissions() ChatPermissions {
	return ChatPermissions{}
}

// SetCanSendMessages sets the `can_send_messages` value of ChatPermissions.
func (p ChatPermissions) SetCanSendMessages(can bool) ChatPermissions {
	p.CanSendMessages = &can
	return p
}

// SetCanSendAudios sets the `can_send_audios` value of ChatPermissions.
func (p ChatPermissions) SetCanSendAudios(can bool) ChatPermissions {
	p.CanSendAudios = &can
	return p
}

// SetCanSendDocuments sets the `can_send_documents` value of ChatPermissions.
func (p ChatPermissions) SetCanSendDocuments(can bool) ChatPermissions {
	p.CanSendDocuments = &can
	return p
}

// SetCanSendPhotos sets the `can_send_photos` value of ChatPermissions.
func (p ChatPermissions) SetCanSendPhotos(can bool) ChatPermissions {
	p.CanSendPhotos = &can
	return p
}

// SetCanSendVideos sets the `can_send_videos` value of ChatPermissions.
func (p ChatPermissions) SetCanSendVideos(can bool) ChatPermissions {
	p.CanSendVideos = &can
	return p
}

// SetCanSendVideoNotes sets the `can_send_video_notes` value of ChatPermissions.
func (p ChatPermissions) SetCanSendVideoNotes(can bool) ChatPermissions {
	p.CanSendVideoNotes = &can
	return p
}

// SetCanSendVoiceNotes sets the `can_send_voice_notes` value of ChatPermissions.
func (p ChatPermissions) SetCanSendVoiceNotes(can bool) ChatPermissions {
	p.CanSendVoiceNotes = &can
	return p
}

// SetCanSendPolls sets the `can_send_polls` value of ChatPermissions.
func (p ChatPermissions) SetCanSendPolls(can bool) ChatPermissions {
	p.CanSendPolls = &can
	return p
}

// SetCanSendOtherMessages sets the `can_send_other_messages` value of ChatPermissions.
func (p ChatPermissions) SetCanSendOtherMessages(can bool) ChatPermissions {
	p.CanSendOtherMessages = &can
	return p
}

// SetCanAddWebPagePreviews sets the `can_add_web_page_previews` value of ChatPermissions.
func (p ChatPermissions) SetCanAddWebPagePreviews(can bool) ChatPermissions {
	p.CanAddWebPagePreviews = &can
	return p
}

// SetCanChangeInfo sets the `can_change_info` value of ChatPermissions.
func (p ChatPermissions) SetCanChangeInfo(can bool) ChatPermissions {
	p.CanChangeInfo = &can
	return p
}

// SetCanInviteUsers sets the `can_invite_users` value of ChatPermissions.
func (p ChatPermissions) SetCanInviteUsers(can bool) ChatPermissions {
	p.CanInviteUsers = &can
	return p
}

// SetCanPinMessages sets the `can_pin_messages` value of ChatPermissions.
func (p ChatPermissions) SetCanPinMessages(can bool) ChatPermissions {
	p.CanPinMessages = &can
	return p
}

// SetCanManageTopics sets the `can_manage_topics` value of ChatPermissions.
func (p ChatPermissions) SetCanManageTopics(can bool) ChatPermissions {
	p.CanManageTopics = &can
	return p
}

////////////////////////////////
// Helper functions for CallbackQuery

// nothing yet

////////////////////////////////
// Helper functions for ReactionType

// NewEmojiReaction returns a ReactionType with emoji.
//
// https://core.telegram.org/bots/api#reactiontypeemoji
func NewEmojiReaction(emoji string) ReactionType {
	return ReactionType{
		Type:  "emoji",
		Emoji: &emoji,
	}
}

// NewCustomEmojiReaction returns a ReactionType with custom emoji.
//
// https://core.telegram.org/bots/api#reactiontypecustomemoji
func NewCustomEmojiReaction(customEmojiID string) ReactionType {
	return ReactionType{
		Type:          "custom_emoji",
		CustomEmojiID: &customEmojiID,
	}
}

// NewMessageReactionWithEmoji returns a new OptionsSetMessageReaction with an emoji string for function `SetMessageReaction`.
func NewMessageReactionWithEmoji(emoji string) OptionsSetMessageReaction {
	return OptionsSetMessageReaction{}.
		SetReaction([]ReactionType{
			NewEmojiReaction(emoji),
		})
}

// NewPaidReaction returns a new ReactionType with type 'paid'.
//
// https://core.telegram.org/bots/api#reactiontypepaid
func NewPaidReaction() ReactionType {
	return ReactionType{
		Type: "paid",
	}
}

////////////////////////////////
// Helper functions for InputFile

// NewInputFileFromFilepath generates an InputFile from given filepath
func NewInputFileFromFilepath(filepath string) InputFile {
	return InputFile{
		Filepath: &filepath,
	}
}

// NewInputFileFromURL generates an InputFile from given url
func NewInputFileFromURL(url string) InputFile {
	return InputFile{
		URL: &url,
	}
}

// NewInputFileFromBytes generates an InputFile from given bytes array
func NewInputFileFromBytes(bytes []byte) InputFile {
	return InputFile{
		Bytes: bytes,
	}
}

// NewInputFileFromFileID generates an InputFile from given file id
func NewInputFileFromFileID(fileID string) InputFile {
	return InputFile{
		FileID: &fileID,
	}
}

////////////////////////////////
// Helper functions for MessageEntity

// NewMessageEntity returns a new MessageEntity.
func NewMessageEntity(typ MessageEntityType, offset, length int) MessageEntity {
	return MessageEntity{
		Type:   typ,
		Offset: offset,
		Length: length,
	}
}

// SetURL sets the `url` value of MessageEntity.
func (m MessageEntity) SetURL(url string) MessageEntity {
	if m.Type == MessageEntityTypeTextLink {
		m.URL = &url
	}
	return m
}

// SetUser sets the `user` value of MessageEntity.
func (m MessageEntity) SetUser(user User) MessageEntity {
	if m.Type == MessageEntityTypeTextMention {
		m.User = &user
	}
	return m
}

// SetLanguage sets the `language` value of MessageEntity.
func (m MessageEntity) SetLanguage(language string) MessageEntity {
	if m.Type == MessageEntityTypePre {
		m.Language = &language
	}
	return m
}

// SetCustomEmojiID sets the `custom_emoji_id` value of MessageEntity.
func (m MessageEntity) SetCustomEmojiID(customEmojiID string) MessageEntity {
	if m.Type == MessageEntityTypeCustomEmoji {
		m.CustomEmojiID = &customEmojiID
	}
	return m
}

////////////////////////////////
// Helper functions for ReplyParameters

// NewReplyParameters returns a new ReplyParameters.
func NewReplyParameters(messageID int64) ReplyParameters {
	return ReplyParameters{
		MessageID: messageID,
	}
}

// SetChatID sets the `chat_id` value of ReplyParameters.
func (p ReplyParameters) SetChatID(chatID ChatID) ReplyParameters {
	p.ChatID = &chatID
	return p
}

// SetAllowSendingWithoutReply sets the `allow_sending_without_reply` value of ReplyParameters.
func (p ReplyParameters) SetAllowSendingWithoutReply(allowSendingWithoutReply bool) ReplyParameters {
	p.AllowSendingWithoutReply = &allowSendingWithoutReply
	return p
}

// SetQuote sets the `quote` value of ReplyParameters.
func (p ReplyParameters) SetQuote(quote string) ReplyParameters {
	p.Quote = &quote
	return p
}

// SetQuoteParseMode sets the `quote_parse_mode` value of ReplyParameters.
func (p ReplyParameters) SetQuoteParseMode(parseMode ParseMode) ReplyParameters {
	p.QuoteParseMode = &parseMode
	return p
}

// SetQuoteEntities sets the `quote_entities` value of ReplyParameters.
func (p ReplyParameters) SetQuoteEntities(entities []MessageEntity) ReplyParameters {
	p.QuoteEntities = entities
	return p
}

// SetQuotePosition sets the `quote_position` value of ReplyParameters.
func (p ReplyParameters) SetQuotePosition(position int) ReplyParameters {
	p.QuotePosition = &position
	return p
}

////////////////////////////////
// Helper functions for MaskPosition

// NewMaskPosition returns a new MaskPosition.
func NewMaskPosition(point MaskPositionPoint, xShift, yShift, scale float32) MaskPosition {
	return MaskPosition{
		Point:  point,
		XShift: xShift,
		YShift: yShift,
		Scale:  scale,
	}
}

////////////////////////////////
// Helper functions for InputSticker

// NewInputSticker returns a new InputSticker.
func NewInputSticker(sticker any, format StickerFormat, emojiList []string) InputSticker {
	return InputSticker{
		Sticker:   sticker,
		Format:    format,
		EmojiList: emojiList,
	}
}

// SetMaskPosition sets the `mask_position` value of InputSticker.
func (s InputSticker) SetMaskPosition(maskPosition MaskPosition) InputSticker {
	s.MaskPosition = &maskPosition
	return s
}

// SetKeywords sets the `keywords` value of InputSticker.
func (s InputSticker) SetKeywords(keywords []string) InputSticker {
	s.Keywords = keywords
	return s
}

////////////////////////////////
// Helper functions for LinkPreviewOptions

// NewLinkPreviewOptions returns a new LinkPreviewOptions.
func NewLinkPreviewOptions() LinkPreviewOptions {
	return LinkPreviewOptions{}
}

// SetIsDisabled sets the `is_disabled` value of LinkPreviewOptions.
func (o LinkPreviewOptions) SetIsDisabled(disabled bool) LinkPreviewOptions {
	o.IsDisabled = &disabled
	return o
}

// SetURL sets the `url` value of LinkPreviewOptions.
func (o LinkPreviewOptions) SetURL(url string) LinkPreviewOptions {
	o.URL = &url
	return o
}

// SetPreferSmallMedia sets the `prefer_small_media` value of LinkPreviewOptions.
func (o LinkPreviewOptions) SetPreferSmallMedia(preferSmallMedia bool) LinkPreviewOptions {
	o.PreferSmallMedia = &preferSmallMedia
	return o
}

// SetPreferLargeMedia sets the `prefer_large_media` value of LinkPreviewOptions.
func (o LinkPreviewOptions) SetPreferLargeMedia(preferLargeMedia bool) LinkPreviewOptions {
	o.PreferLargeMedia = &preferLargeMedia
	return o
}

// SetShowAboveText sets the `show_above_text` value of LinkPreviewOptions.
func (o LinkPreviewOptions) SetShowAboveText(showAboveText bool) LinkPreviewOptions {
	o.ShowAboveText = &showAboveText
	return o
}

////////////////////////////////
// Helper functions for ReplyKeyboardMarkup

// NewReplyKeyboardMarkup returns a new ReplyKeyboardMarkup.
func NewReplyKeyboardMarkup(keyboard [][]KeyboardButton) ReplyKeyboardMarkup {
	return ReplyKeyboardMarkup{
		Keyboard: keyboard,
	}
}

// SetIsPersistent sets the `is_persistent` value of ReplyKeyboardMarkup.
func (m ReplyKeyboardMarkup) SetIsPersistent(persistent bool) ReplyKeyboardMarkup {
	m.IsPersistent = &persistent
	return m
}

// SetResizeKeyboard sets the `resize_keyboard` value of ReplyKeyboardMarkup.
func (m ReplyKeyboardMarkup) SetResizeKeyboard(resizeKeyboard bool) ReplyKeyboardMarkup {
	m.ResizeKeyboard = &resizeKeyboard
	return m
}

// SetOneTimeKeyboard sets the `one_time_keyboard` value of ReplyKeyboardMarkup.
func (m ReplyKeyboardMarkup) SetOneTimeKeyboard(oneTimeKeyboard bool) ReplyKeyboardMarkup {
	m.OneTimeKeyboard = &oneTimeKeyboard
	return m
}

// SetInputFieldPlaceholder sets the `input_field_placeholder` value of ReplyKeyboardMarkup.
func (m ReplyKeyboardMarkup) SetInputFieldPlaceholder(placeholder string) ReplyKeyboardMarkup {
	m.InputFieldPlaceholder = &placeholder
	return m
}

// SetSelective sets the `selective` value of ReplyKeyboardMarkup.
func (m ReplyKeyboardMarkup) SetSelective(selective bool) ReplyKeyboardMarkup {
	m.Selective = &selective
	return m
}

////////////////////////////////
// Helper functions for KeyboardButton

// NewKeyboardButton returns a new KeyboardButton.
func NewKeyboardButton(text string) KeyboardButton {
	return KeyboardButton{
		Text: text,
	}
}

// SetRequestUsers sets the `request_users` value of KeyboardButton.
func (b KeyboardButton) SetRequestUsers(requestUsers KeyboardButtonRequestUsers) KeyboardButton {
	b.RequestUsers = &requestUsers
	return b
}

// SetRequestChat sets the `request_chat` value of KeyboardButton.
func (b KeyboardButton) SetRequestChat(requestChat KeyboardButtonRequestChat) KeyboardButton {
	b.RequestChat = &requestChat
	return b
}

// SetRequestContact sets the `request_contact` value of KeyboardButton.
func (b KeyboardButton) SetRequestContact(requestContact bool) KeyboardButton {
	b.RequestContact = &requestContact
	return b
}

// SetRequestLocation sets the `request_location` value of KeyboardButton.
func (b KeyboardButton) SetRequestLocation(requestLocation bool) KeyboardButton {
	b.RequestLocation = &requestLocation
	return b
}

// SetRequestPoll sets the `request_poll` value of KeyboardButton.
func (b KeyboardButton) SetRequestPoll(requestPoll KeyboardButtonPollType) KeyboardButton {
	b.RequestPoll = &requestPoll
	return b
}

// SetWebApp sets the `web_app` value of KeyboardButton.
func (b KeyboardButton) SetWebApp(webApp WebAppInfo) KeyboardButton {
	b.WebApp = &webApp
	return b
}

////////////////////////////////
// Helper functions for KeyboardButtonRequestUsers

// NewKeyboardButtonRequestUsers returns a new KeyboardButtonRequestUsers.
func NewKeyboardButtonRequestUsers(requestID int64) KeyboardButtonRequestUsers {
	return KeyboardButtonRequestUsers{
		RequestID: requestID,
	}
}

// SetUserIsBot sets the `user_is_bot` value of KeyboardButtonRequestUsers.
func (u KeyboardButtonRequestUsers) SetUserIsBot(userIsBot bool) KeyboardButtonRequestUsers {
	u.UserIsBot = &userIsBot
	return u
}

// SetUserIsPremium sets the `user_is_premium` value of KeyboardButtonRequestUsers.
func (u KeyboardButtonRequestUsers) SetUserIsPremium(userIsPremium bool) KeyboardButtonRequestUsers {
	u.UserIsPremium = &userIsPremium
	return u
}

// SetMaxQuantity sets the `max_quantity` value of KeyboardButtonRequestUsers.
func (u KeyboardButtonRequestUsers) SetMaxQuantity(maxQuantity int) KeyboardButtonRequestUsers {
	u.MaxQuantity = &maxQuantity
	return u
}

// SetRequestName sets the `request_name` value of KeyboardButtonRequestUsers.
func (u KeyboardButtonRequestUsers) SetRequestName(requestName bool) KeyboardButtonRequestUsers {
	u.RequestName = &requestName
	return u
}

// SetRequestUsername sets the `request_username` value of KeyboardButtonRequestUsers.
func (u KeyboardButtonRequestUsers) SetRequestUsername(requestUsername bool) KeyboardButtonRequestUsers {
	u.RequestUsername = &requestUsername
	return u
}

// SetRequestPhoto sets the `request_photo` value of KeyboardButtonRequestUsers.
func (u KeyboardButtonRequestUsers) SetRequestPhoto(requestPhoto bool) KeyboardButtonRequestUsers {
	u.RequestPhoto = &requestPhoto
	return u
}

////////////////////////////////
// Helper functions for KeyboardButtonRequestChat

// NewKeyboardButtonRequestChat returns a new KeyboardButtonRequestChat.
func NewKeyboardButtonRequestChat(requestID int64, isChannel bool) KeyboardButtonRequestChat {
	return KeyboardButtonRequestChat{
		RequestID:     requestID,
		ChatIsChannel: isChannel,
	}
}

// SetChatIsForum sets the `chat_is_forum` value of KeyboardButtonRequestChat.
func (c KeyboardButtonRequestChat) SetChatIsForum(isForum bool) KeyboardButtonRequestChat {
	c.ChatIsForum = &isForum
	return c
}

// SetChatHasUsername sets the `chat_has_username` value of KeyboardButtonRequestChat.
func (c KeyboardButtonRequestChat) SetChatHasUsername(hasUsername bool) KeyboardButtonRequestChat {
	c.ChatHasUsername = &hasUsername
	return c
}

// SetChatIsCreated sets the `chat_is_created` value of KeyboardButtonRequestChat.
func (c KeyboardButtonRequestChat) SetChatIsCreated(isCreated bool) KeyboardButtonRequestChat {
	c.ChatIsCreated = &isCreated
	return c
}

// SetUserAdministratorRights sets the `user_administrator_rights` value of KeyboardButtonRequestChat.
func (c KeyboardButtonRequestChat) SetUserAdministratorRights(userAdminRights ChatAdministratorRights) KeyboardButtonRequestChat {
	c.UserAdministratorRights = &userAdminRights
	return c
}

// SetBotAdministratorRights sets the `bot_administrator_rights` value of KeyboardButtonRequestChat.
func (c KeyboardButtonRequestChat) SetBotAdministratorRights(botAdminRights ChatAdministratorRights) KeyboardButtonRequestChat {
	c.BotAdministratorRights = &botAdminRights
	return c
}

// SetBotIsMember sets the `bot_is_member` value of KeyboardButtonRequestChat.
func (c KeyboardButtonRequestChat) SetBotIsMember(isMember bool) KeyboardButtonRequestChat {
	c.BotIsMember = &isMember
	return c
}

// SetRequestTitle sets the `request_title` value of KeyboardButtonRequestChat.
func (c KeyboardButtonRequestChat) SetRequestTitle(requestTitle bool) KeyboardButtonRequestChat {
	c.RequestTitle = &requestTitle
	return c
}

// SetRequestUsername sets the `request_username` value of KeyboardButtonRequestChat.
func (c KeyboardButtonRequestChat) SetRequestUsername(requestUsername bool) KeyboardButtonRequestChat {
	c.RequestUsername = &requestUsername
	return c
}

// SetRequestPhoto sets the `request_photo` value of KeyboardButtonRequestChat.
func (c KeyboardButtonRequestChat) SetRequestPhoto(requestPhoto bool) KeyboardButtonRequestChat {
	c.RequestPhoto = &requestPhoto
	return c
}

////////////////////////////////
// Helper functions for KeyboardButtonPollType

// NewKeyboardButtonPollType returns a new KeyboardButtonPollType.
func NewKeyboardButtonPollType(typ string) KeyboardButtonPollType {
	return KeyboardButtonPollType{
		Type: &typ,
	}
}

////////////////////////////////
// Helper functions for ReplyKeyboardRemove

// NewReplyKeyboardRemove returns a new ReplyKeyboardRemove.
func NewReplyKeyboardRemove(remove bool) ReplyKeyboardRemove {
	return ReplyKeyboardRemove{
		RemoveKeyboard: remove,
	}
}

// SetSelective sets the `selective` value of ReplyKeyboardRemove.
func (r ReplyKeyboardRemove) SetSelective(selective bool) ReplyKeyboardRemove {
	r.Selective = &selective
	return r
}

////////////////////////////////
// Helper functions for InlineKeyboardMarkup

// NewInlineKeyboardMarkup returns a new InlineKeyboardMarkup with given rows of keyboard buttons.
func NewInlineKeyboardMarkup(keyboard [][]InlineKeyboardButton) InlineKeyboardMarkup {
	return InlineKeyboardMarkup{
		InlineKeyboard: keyboard,
	}
}

////////////////////////////////
// Helper functions for InlineKeyboardButton

// NewInlineKeyboardButton returns a new InlineKeyboardButton with given text.
func NewInlineKeyboardButton(text string) InlineKeyboardButton {
	return InlineKeyboardButton{
		Text: text,
	}
}

// SetURL sets the `url` value of InlineKeyboardButton.
func (b InlineKeyboardButton) SetURL(url string) InlineKeyboardButton {
	b.URL = &url
	return b
}

// SetLoginURL sets the `login_url` value of InlineKeyboardButton.
func (b InlineKeyboardButton) SetLoginURL(loginURL LoginURL) InlineKeyboardButton {
	b.LoginURL = &loginURL
	return b
}

// SetCallbackData sets the `callback_data` value of InlineKeyboardButton.
func (b InlineKeyboardButton) SetCallbackData(data string) InlineKeyboardButton {
	b.CallbackData = &data
	return b
}

// SetWebApp sets the `web_app` value of InlineKeyboardButton.
func (b InlineKeyboardButton) SetWebApp(webApp WebAppInfo) InlineKeyboardButton {
	b.WebApp = &webApp
	return b
}

// SetSwichInlineQuery sets the `switch_inline_query` value of InlineKeyboardButton.
func (b InlineKeyboardButton) SetSwichInlineQuery(query string) InlineKeyboardButton {
	b.SwitchInlineQuery = &query
	return b
}

// SetSwichInlineQueryCurrentChat sets the `switch_inline_query_current_chat` value of InlineKeyboardButton.
func (b InlineKeyboardButton) SetSwichInlineQueryCurrentChat(query string) InlineKeyboardButton {
	b.SwitchInlineQueryCurrentChat = &query
	return b
}

// SetSwichInlineQueryChosenChat sets the `switch_inline_query_chosen_chat` value of InlineKeyboardButton.
func (b InlineKeyboardButton) SetSwichInlineQueryChosenChat(chosenChat SwitchInlineQueryChosenChat) InlineKeyboardButton {
	b.SwitchInlineQueryChosenChat = &chosenChat
	return b
}

// SetCallbackGame sets the `callback_game` value of InlineKeyboardButton.
func (b InlineKeyboardButton) SetCallbackGame(callbackGame CallbackGame) InlineKeyboardButton {
	b.CallbackGame = &callbackGame
	return b
}

// SetPay sets the `pay` value of InlineKeyboardButton.
func (b InlineKeyboardButton) SetPay(pay bool) InlineKeyboardButton {
	b.Pay = &pay
	return b
}

////////////////////////////////
// Helper functions for InlineQueryResultsButton

// NewInlineQueryResultsButton returns a new InlineQueryResultsButton with given text.
//
// https://core.telegram.org/bots/api#inlinequeryresultsbutton
func NewInlineQueryResultsButton(text string) InlineQueryResultsButton {
	return InlineQueryResultsButton{
		Text: text,
	}
}

// SetWebApp sets the `web_app` value of InlineQueryResultsButton.
func (b InlineQueryResultsButton) SetWebApp(webApp WebAppInfo) InlineQueryResultsButton {
	b.WebApp = &webApp
	return b
}

// SetStartParameter sets the `start_parameter` value of InlineQueryResultsButton.
func (b InlineQueryResultsButton) SetStartParameter(startParameter string) InlineQueryResultsButton {
	b.StartParameter = &startParameter
	return b
}

////////////////////////////////
// Helper functions for InlineQueryResult

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

// NewInlineQueryResultArticle is a helper function for generating a new InlineQueryResultArticle.
//
// https://core.telegram.org/bots/api#inlinequeryresultarticle
func NewInlineQueryResultArticle(title, messageText, description string) (newArticle InlineQueryResultArticle, generatedID *string) {
	if id, err := newUUID(); err == nil {
		return InlineQueryResultArticle{
			InlineQueryResult: InlineQueryResult{
				Type: InlineQueryResultTypeArticle,
				ID:   id,
			},
			Title: title,
			InputMessageContent: InputTextMessageContent{
				MessageText: messageText,
			},
			Description: &description,
		}, &id
	}

	return InlineQueryResultArticle{}, nil
}

// SetReplyMarkup sets the `reply_markup` value of InlineQueryResultArticle.
func (r InlineQueryResultArticle) SetReplyMarkup(markup InlineKeyboardMarkup) InlineQueryResultArticle {
	r.ReplyMarkup = &markup
	return r
}

// SetURL sets the `url` value of InlineQueryResultArticle.
//
// NOTE: Set an empty string for hiding the URL.
func (r InlineQueryResultArticle) SetURL(url string) InlineQueryResultArticle {
	r.URL = &url
	return r
}

// SetDescription sets the `description` value of InlineQueryResultArticle.
func (r InlineQueryResultArticle) SetDescription(description string) InlineQueryResultArticle {
	r.Description = &description
	return r
}

// SetThumbnailURL sets the `thumbnail_url` value of InlineQueryResultArticle.
func (r InlineQueryResultArticle) SetThumbnailURL(thumbnailURL string) InlineQueryResultArticle {
	r.ThumbnailURL = &thumbnailURL
	return r
}

// SetThumbnailWidth sets the `thumbnail_width` value of InlineQueryResultArticle.
func (r InlineQueryResultArticle) SetThumbnailWidth(width int) InlineQueryResultArticle {
	r.ThumbnailWidth = &width
	return r
}

// SetThumbnailHeight sets the `thumbnail_height` value of InlineQueryResultArticle.
func (r InlineQueryResultArticle) SetThumbnailHeight(height int) InlineQueryResultArticle {
	r.ThumbnailHeight = &height
	return r
}

// NewInlineQueryResultPhoto is a helper function for generating a new InlineQueryResultPhoto.
//
// Photo must be in jpeg format, < 5MB.
//
// https://core.telegram.org/bots/api#inlinequeryresultphoto
func NewInlineQueryResultPhoto(photoURL, thumbnailURL string) (newPhoto InlineQueryResultPhoto, generatedID *string) {
	if id, err := newUUID(); err == nil {
		return InlineQueryResultPhoto{
			InlineQueryResult: InlineQueryResult{
				Type: InlineQueryResultTypePhoto,
				ID:   id,
			},
			PhotoURL:     photoURL,
			ThumbnailURL: thumbnailURL,
		}, &id
	}

	return InlineQueryResultPhoto{}, nil
}

// SetPhotoWidth sets the `photo_width` value of InlineQueryResultPhoto.
func (r InlineQueryResultPhoto) SetPhotoWidth(width int) InlineQueryResultPhoto {
	r.PhotoWidth = &width
	return r
}

// SetPhotoHeight sets the `photo_height` value of InlineQueryResultPhoto.
func (r InlineQueryResultPhoto) SetPhotoHeight(height int) InlineQueryResultPhoto {
	r.PhotoHeight = &height
	return r
}

// SetTitle sets the `title` value of InlineQueryResultPhoto.
func (r InlineQueryResultPhoto) SetTitle(title string) InlineQueryResultPhoto {
	r.Title = &title
	return r
}

// SetDescription sets the `description` value of InlineQueryResultPhoto.
func (r InlineQueryResultPhoto) SetDescription(description string) InlineQueryResultPhoto {
	r.Description = &description
	return r
}

// SetCaption sets the `caption` value of InlineQueryResultPhoto.
func (r InlineQueryResultPhoto) SetCaption(caption string) InlineQueryResultPhoto {
	r.Caption = &caption
	return r
}

// SetParseMode sets the `parse_mode` value of InlineQueryResultPhoto.
func (r InlineQueryResultPhoto) SetParseMode(parseMode ParseMode) InlineQueryResultPhoto {
	r.ParseMode = &parseMode
	return r
}

// SetCaptionEntities sets the `caption_entities` value of InlineQueryResultPhoto.
func (r InlineQueryResultPhoto) SetCaptionEntities(entities []MessageEntity) InlineQueryResultPhoto {
	r.CaptionEntities = entities
	return r
}

// SetReplyMarkup sets the `reply_markup` value of InlineQueryResultPhoto.
func (r InlineQueryResultPhoto) SetReplyMarkup(markup InlineKeyboardMarkup) InlineQueryResultPhoto {
	r.ReplyMarkup = &markup
	return r
}

// SetInputMessageContent sets the `input_message_content` value of InlineQueryResultPhoto.
func (r InlineQueryResultPhoto) SetInputMessageContent(content InputMessageContent) InlineQueryResultPhoto {
	r.InputMessageContent = &content
	return r
}

// NewInlineQueryResultGif is a helper function for generating a new InlineQueryResultGif.
//
// Gif must be in gif format, < 1MB.
//
// https://core.telegram.org/bots/api#inlinequeryresultgif
func NewInlineQueryResultGif(gifURL, thumbnailURL string) (newGif InlineQueryResultGif, generatedID *string) {
	if id, err := newUUID(); err == nil {
		return InlineQueryResultGif{
			InlineQueryResult: InlineQueryResult{
				Type: InlineQueryResultTypeGif,
				ID:   id,
			},
			GifURL:       gifURL,
			ThumbnailURL: thumbnailURL,
		}, &id
	}

	return InlineQueryResultGif{}, nil
}

// SetGifWidth sets the `gif_width` value of InlineQueryResultGif.
func (r InlineQueryResultGif) SetGifWidth(width int) InlineQueryResultGif {
	r.GifWidth = &width
	return r
}

// SetGifHeight sets the `gif_height` value of InlineQueryResultGif.
func (r InlineQueryResultGif) SetGifHeight(height int) InlineQueryResultGif {
	r.GifHeight = &height
	return r
}

// SetGifDuration sets the `gif_duration` value of InlineQueryResultGif.
func (r InlineQueryResultGif) SetGifDuration(duration int) InlineQueryResultGif {
	r.GifDuration = &duration
	return r
}

// SetThumbnailMimeType sets the `thumbnail_mime_type` value of InlineQueryResultGif.
func (r InlineQueryResultGif) SetThumbnailMimeType(mimeType ThumbnailMimeType) InlineQueryResultGif {
	r.ThumbnailMimeType = &mimeType
	return r
}

// SetTitle sets the `title` value of InlineQueryResultGif.
func (r InlineQueryResultGif) SetTitle(title string) InlineQueryResultGif {
	r.Title = &title
	return r
}

// SetCaption sets the `caption` value of InlineQueryResultGif.
func (r InlineQueryResultGif) SetCaption(caption string) InlineQueryResultGif {
	r.Caption = &caption
	return r
}

// SetParseMode sets the `parse_mode` value of InlineQueryResultGif.
func (r InlineQueryResultGif) SetParseMode(parseMode ParseMode) InlineQueryResultGif {
	r.ParseMode = &parseMode
	return r
}

// SetCaptionEntities sets the `caption_entities` value of InlineQueryResultGif.
func (r InlineQueryResultGif) SetCaptionEntities(entities []MessageEntity) InlineQueryResultGif {
	r.CaptionEntities = entities
	return r
}

// SetReplyMarkup sets the `reply_markup` value of InlineQueryResultGif.
func (r InlineQueryResultGif) SetReplyMarkup(markup InlineKeyboardMarkup) InlineQueryResultGif {
	r.ReplyMarkup = &markup
	return r
}

// SetInputMessageContent sets the `input_message_content` value of InlineQueryResultGif.
func (r InlineQueryResultGif) SetInputMessageContent(content InputMessageContent) InlineQueryResultGif {
	r.InputMessageContent = &content
	return r
}

// NewInlineQueryResultMpeg4Gif is a helper function for generating a new InlineQueryResultMpeg4Gif.
//
// Mpeg4 must be in H.264/MPEG-4 AVC video(wihout sound) format, < 1MB.
//
// https://core.telegram.org/bots/api#inlinequeryresultmpeg4gif
func NewInlineQueryResultMpeg4Gif(mpeg4URL, thumbnailURL string) (newMpeg4Gif InlineQueryResultMpeg4Gif, generatedID *string) {
	if id, err := newUUID(); err == nil {
		return InlineQueryResultMpeg4Gif{
			InlineQueryResult: InlineQueryResult{
				Type: InlineQueryResultTypeMpeg4Gif,
				ID:   id,
			},
			Mpeg4URL:     mpeg4URL,
			ThumbnailURL: thumbnailURL,
		}, &id
	}

	return InlineQueryResultMpeg4Gif{}, nil
}

// SetMpeg4Width sets the `mpeg4_width` value of InlineQueryResultMpeg4Gif.
func (r InlineQueryResultMpeg4Gif) SetMpeg4Width(width int) InlineQueryResultMpeg4Gif {
	r.Mpeg4Width = &width
	return r
}

// SetMpeg4Height sets the `mpeg4_height` value of InlineQueryResultMpeg4Gif.
func (r InlineQueryResultMpeg4Gif) SetMpeg4Height(height int) InlineQueryResultMpeg4Gif {
	r.Mpeg4Height = &height
	return r
}

// SetMpeg4Duration sets the `mpeg4_duration` value of InlineQueryResultMpeg4Gif.
func (r InlineQueryResultMpeg4Gif) SetMpeg4Duration(duration int) InlineQueryResultMpeg4Gif {
	r.Mpeg4Duration = &duration
	return r
}

// SetThumbnailMimeType sets the `thumbnail_mime_type` value of InlineQueryResultMpeg4Gif.
func (r InlineQueryResultMpeg4Gif) SetThumbnailMimeType(mimeType ThumbnailMimeType) InlineQueryResultMpeg4Gif {
	r.ThumbnailMimeType = &mimeType
	return r
}

// SetTitle sets the `title` value of InlineQueryResultMpeg4Gif.
func (r InlineQueryResultMpeg4Gif) SetTitle(title string) InlineQueryResultMpeg4Gif {
	r.Title = &title
	return r
}

// SetCaption sets the `caption` value of InlineQueryResultMpeg4Gif.
func (r InlineQueryResultMpeg4Gif) SetCaption(caption string) InlineQueryResultMpeg4Gif {
	r.Caption = &caption
	return r
}

// SetParseMode sets the `parse_mode` value of InlineQueryResultMpeg4Gif.
func (r InlineQueryResultMpeg4Gif) SetParseMode(parseMode ParseMode) InlineQueryResultMpeg4Gif {
	r.ParseMode = &parseMode
	return r
}

// SetCaptionEntities sets the `caption_entities` value of InlineQueryResultMpeg4Gif.
func (r InlineQueryResultMpeg4Gif) SetCaptionEntities(entities []MessageEntity) InlineQueryResultMpeg4Gif {
	r.CaptionEntities = entities
	return r
}

// SetReplyMarkup sets the `reply_markup` value of InlineQueryResultMpeg4Gif.
func (r InlineQueryResultMpeg4Gif) SetReplyMarkup(markup InlineKeyboardMarkup) InlineQueryResultMpeg4Gif {
	r.ReplyMarkup = &markup
	return r
}

// SetInputMessageContent sets the `input_message_content` value of InlineQueryResultMpeg4Gif.
func (r InlineQueryResultMpeg4Gif) SetInputMessageContent(content InputMessageContent) InlineQueryResultMpeg4Gif {
	r.InputMessageContent = &content
	return r
}

// NewInlineQueryResultVideo is a helper function for generating a new InlineQueryResultVideo.
//
// https://core.telegram.org/bots/api#inlinequeryresultvideo
func NewInlineQueryResultVideo(videoURL, thumbnailURL, title string, mimeType VideoMimeType) (newVideo InlineQueryResultVideo, generatedID *string) {
	if id, err := newUUID(); err == nil {
		return InlineQueryResultVideo{
			InlineQueryResult: InlineQueryResult{
				Type: InlineQueryResultTypeVideo,
				ID:   id,
			},
			VideoURL:     videoURL,
			MimeType:     mimeType,
			ThumbnailURL: thumbnailURL,
			Title:        title,
		}, &id
	}

	return InlineQueryResultVideo{}, nil
}

// SetCaption sets the `caption` value of InlineQueryResultVideo.
func (r InlineQueryResultVideo) SetCaption(caption string) InlineQueryResultVideo {
	r.Caption = &caption
	return r
}

// SetParseMode sets the `parse_mode` value of InlineQueryResultVideo.
func (r InlineQueryResultVideo) SetParseMode(parseMode ParseMode) InlineQueryResultVideo {
	r.ParseMode = &parseMode
	return r
}

// SetCaptionEntities sets the `caption_entities` value of InlineQueryResultVideo.
func (r InlineQueryResultVideo) SetCaptionEntities(entities []MessageEntity) InlineQueryResultVideo {
	r.CaptionEntities = entities
	return r
}

// SetVideoWidth sets the `video_width` value of InlineQueryResultVideo.
func (r InlineQueryResultVideo) SetVideoWidth(width int) InlineQueryResultVideo {
	r.VideoWidth = &width
	return r
}

// SetVideoHeight sets the `video_height` value of InlineQueryResultVideo.
func (r InlineQueryResultVideo) SetVideoHeight(height int) InlineQueryResultVideo {
	r.VideoHeight = &height
	return r
}

// SetVideoDuration sets the `video_duration` value of InlineQueryResultVideo.
func (r InlineQueryResultVideo) SetVideoDuration(duration int) InlineQueryResultVideo {
	r.VideoDuration = &duration
	return r
}

// SetDescription sets the `description` value of InlineQueryResultVideo.
func (r InlineQueryResultVideo) SetDescription(description string) InlineQueryResultVideo {
	r.Description = &description
	return r
}

// SetReplyMarkup sets the `reply_markup` value of InlineQueryResultVideo.
func (r InlineQueryResultVideo) SetReplyMarkup(markup InlineKeyboardMarkup) InlineQueryResultVideo {
	r.ReplyMarkup = &markup
	return r
}

// SetInputMessageContent sets the `input_message_content` value of InlineQueryResultVideo.
func (r InlineQueryResultVideo) SetInputMessageContent(content InputMessageContent) InlineQueryResultVideo {
	r.InputMessageContent = &content
	return r
}

// NewInlineQueryResultAudio is a helper function for generating a new InlineQueryResultAudio.
//
// https://core.telegram.org/bots/api#inlinequeryresultaudio
func NewInlineQueryResultAudio(audioURL, title string) (newAudio InlineQueryResultAudio, generatedID *string) {
	if id, err := newUUID(); err == nil {
		return InlineQueryResultAudio{
			InlineQueryResult: InlineQueryResult{
				Type: InlineQueryResultTypeAudio,
				ID:   id,
			},
			AudioURL: audioURL,
			Title:    title,
		}, &id
	}

	return InlineQueryResultAudio{}, nil
}

// SetCaption sets the `caption` value of InlineQueryResultAudio.
func (r InlineQueryResultAudio) SetCaption(caption string) InlineQueryResultAudio {
	r.Caption = &caption
	return r
}

// SetParseMode sets the `parse_mode` value of InlineQueryResultAudio.
func (r InlineQueryResultAudio) SetParseMode(parseMode ParseMode) InlineQueryResultAudio {
	r.ParseMode = &parseMode
	return r
}

// SetCaptionEntities sets the `caption_entities` value of InlineQueryResultAudio.
func (r InlineQueryResultAudio) SetCaptionEntities(entities []MessageEntity) InlineQueryResultAudio {
	r.CaptionEntities = entities
	return r
}

// SetPerformer sets the `performer` value of InlineQueryResultAudio.
func (r InlineQueryResultAudio) SetPerformer(performer string) InlineQueryResultAudio {
	r.Performer = &performer
	return r
}

// SetAudioDuration sets the `audio_duration` value of InlineQueryResultAudio.
func (r InlineQueryResultAudio) SetAudioDuration(duration int) InlineQueryResultAudio {
	r.AudioDuration = &duration
	return r
}

// SetReplyMarkup sets the `reply_markup` value of InlineQueryResultAudio.
func (r InlineQueryResultAudio) SetReplyMarkup(markup InlineKeyboardMarkup) InlineQueryResultAudio {
	r.ReplyMarkup = &markup
	return r
}

// SetInputMessageContent sets the `input_message_content` value of InlineQueryResultAudio.
func (r InlineQueryResultAudio) SetInputMessageContent(content InputMessageContent) InlineQueryResultAudio {
	r.InputMessageContent = &content
	return r
}

// NewInlineQueryResultVoice is a helper function for generating a new InlineQueryResultVoice.
//
// https://core.telegram.org/bots/api#inlinequeryresultvoice
func NewInlineQueryResultVoice(voiceURL, title string) (newVoice InlineQueryResultVoice, generatedID *string) {
	if id, err := newUUID(); err == nil {
		return InlineQueryResultVoice{
			InlineQueryResult: InlineQueryResult{
				Type: InlineQueryResultTypeVoice,
				ID:   id,
			},
			VoiceURL: voiceURL,
			Title:    title,
		}, &id
	}

	return InlineQueryResultVoice{}, nil
}

// SetCaption sets the `caption` value of InlineQueryResultVoice.
func (r InlineQueryResultVoice) SetCaption(caption string) InlineQueryResultVoice {
	r.Caption = &caption
	return r
}

// SetParseMode sets the `parse_mode` value of InlineQueryResultVoice.
func (r InlineQueryResultVoice) SetParseMode(parseMode ParseMode) InlineQueryResultVoice {
	r.ParseMode = &parseMode
	return r
}

// SetCaptionEntities sets the `caption_entities` value of InlineQueryResultVoice.
func (r InlineQueryResultVoice) SetCaptionEntities(entities []MessageEntity) InlineQueryResultVoice {
	r.CaptionEntities = entities
	return r
}

// SetVoiceDuration sets the `voice_duration` value of InlineQueryResultVoice.
func (r InlineQueryResultVoice) SetVoiceDuration(duration int) InlineQueryResultVoice {
	r.VoiceDuration = &duration
	return r
}

// SetReplyMarkup sets the `reply_markup` value of InlineQueryResultVoice.
func (r InlineQueryResultVoice) SetReplyMarkup(markup InlineKeyboardMarkup) InlineQueryResultVoice {
	r.ReplyMarkup = &markup
	return r
}

// SetInputMessageContent sets the `input_message_content` value of InlineQueryResultVoice.
func (r InlineQueryResultVoice) SetInputMessageContent(content InputMessageContent) InlineQueryResultVoice {
	r.InputMessageContent = &content
	return r
}

// NewInlineQueryResultDocument is a helper function for generating a new InlineQueryResultDocument.
//
// https://core.telegram.org/bots/api#inlinequeryresultdocument
func NewInlineQueryResultDocument(documentURL, title string, mimeType DocumentMimeType) (newDocument InlineQueryResultDocument, generatedID *string) {
	if id, err := newUUID(); err == nil {
		return InlineQueryResultDocument{
			InlineQueryResult: InlineQueryResult{
				Type: InlineQueryResultTypeDocument,
				ID:   id,
			},
			Title:       title,
			DocumentURL: documentURL,
			MimeType:    mimeType,
		}, &id
	}

	return InlineQueryResultDocument{}, nil
}

// SetCaption sets the `caption` value of InlineQueryResultDocument.
func (r InlineQueryResultDocument) SetCaption(caption string) InlineQueryResultDocument {
	r.Caption = &caption
	return r
}

// SetParseMode sets the `parse_mode` value of InlineQueryResultDocument.
func (r InlineQueryResultDocument) SetParseMode(parseMode ParseMode) InlineQueryResultDocument {
	r.ParseMode = &parseMode
	return r
}

// SetCaptionEntities sets the `caption_entities` value of InlineQueryResultDocument.
func (r InlineQueryResultDocument) SetCaptionEntities(entities []MessageEntity) InlineQueryResultDocument {
	r.CaptionEntities = entities
	return r
}

// SetDescription sets the `description` value of InlineQueryResultDocument.
func (r InlineQueryResultDocument) SetDescription(description string) InlineQueryResultDocument {
	r.Description = &description
	return r
}

// SetReplyMarkup sets the `reply_markup` value of InlineQueryResultDocument.
func (r InlineQueryResultDocument) SetReplyMarkup(markup InlineKeyboardMarkup) InlineQueryResultDocument {
	r.ReplyMarkup = &markup
	return r
}

// SetInputMessageContent sets the `input_message_content` value of InlineQueryResultDocument.
func (r InlineQueryResultDocument) SetInputMessageContent(content InputMessageContent) InlineQueryResultDocument {
	r.InputMessageContent = &content
	return r
}

// SetThumbnailURL sets the `thumbnail_url` value of InlineQueryResultDocument.
func (r InlineQueryResultDocument) SetThumbnailURL(thumbnailURL string) InlineQueryResultDocument {
	r.ThumbnailURL = &thumbnailURL
	return r
}

// SetThumbnailWidth sets the `thumbnail_width` value of InlineQueryResultDocument.
func (r InlineQueryResultDocument) SetThumbnailWidth(thumbnailWidth int) InlineQueryResultDocument {
	r.ThumbnailWidth = &thumbnailWidth
	return r
}

// SetThumbnailHeight sets the `thumbnail_height` value of InlineQueryResultDocument.
func (r InlineQueryResultDocument) SetThumbnailHeight(thumbnailHeight int) InlineQueryResultDocument {
	r.ThumbnailHeight = &thumbnailHeight
	return r
}

// NewInlineQueryResultLocation is a helper function for generating a new InlineQueryResultLocation.
//
// https://core.telegram.org/bots/api#inlinequeryresultlocation
func NewInlineQueryResultLocation(latitude, longitude float32, title string) (newLocation InlineQueryResultLocation, generatedID *string) {
	if id, err := newUUID(); err == nil {
		return InlineQueryResultLocation{
			InlineQueryResult: InlineQueryResult{
				Type: InlineQueryResultTypeLocation,
				ID:   id,
			},
			Latitude:  latitude,
			Longitude: longitude,
			Title:     title,
		}, &id
	}

	return InlineQueryResultLocation{}, nil
}

// SetHorizontalAccuracy sets the `horizontal_accuracy` value of InlineQueryResultLocation.
func (r InlineQueryResultLocation) SetHorizontalAccuracy(horizontalAccuracy float32) InlineQueryResultLocation {
	r.HorizontalAccuracy = &horizontalAccuracy
	return r
}

// SetLivePeriod sets the `live_period` value of InlineQueryResultLocation.
func (r InlineQueryResultLocation) SetLivePeriod(livePeriod int) InlineQueryResultLocation {
	r.LivePeriod = &livePeriod
	return r
}

// SetHeading sets the `heading` value of InlineQueryResultLocation.
func (r InlineQueryResultLocation) SetHeading(heading int) InlineQueryResultLocation {
	r.Heading = &heading
	return r
}

// SetProximityAlertRadius sets the `proximity_alert_radius` value of InlineQueryResultLocation.
func (r InlineQueryResultLocation) SetProximityAlertRadius(radius int) InlineQueryResultLocation {
	r.ProximityAlertRadius = &radius
	return r
}

// SetReplyMarkup sets the `reply_markup` value of InlineQueryResultLocation.
func (r InlineQueryResultLocation) SetReplyMarkup(markup InlineKeyboardMarkup) InlineQueryResultLocation {
	r.ReplyMarkup = &markup
	return r
}

// SetInputMessageContent sets the `input_message_content` value of InlineQueryResultLocation.
func (r InlineQueryResultLocation) SetInputMessageContent(content InputMessageContent) InlineQueryResultLocation {
	r.InputMessageContent = &content
	return r
}

// SetThumbnailURL sets the `thumbnail_url` value of InlineQueryResultLocation.
func (r InlineQueryResultLocation) SetThumbnailURL(thumbnailURL string) InlineQueryResultLocation {
	r.ThumbnailURL = &thumbnailURL
	return r
}

// SetThumbnailWidth sets the `thumbnail_width` value of InlineQueryResultLocation.
func (r InlineQueryResultLocation) SetThumbnailWidth(thumbnailWidth int) InlineQueryResultLocation {
	r.ThumbnailWidth = &thumbnailWidth
	return r
}

// SetThumbnailHeight sets the `thumbnail_height` value of InlineQueryResultLocation.
func (r InlineQueryResultLocation) SetThumbnailHeight(thumbnailHeight int) InlineQueryResultLocation {
	r.ThumbnailHeight = &thumbnailHeight
	return r
}

// NewInlineQueryResultVenue is a helper function for generating a new InlineQueryResultVenue.
//
// https://core.telegram.org/bots/api#inlinequeryresultvenue
func NewInlineQueryResultVenue(latitude, longitude float32, title, address string) (newVenue InlineQueryResultVenue, generatedID *string) {
	if id, err := newUUID(); err == nil {
		return InlineQueryResultVenue{
			InlineQueryResult: InlineQueryResult{
				Type: InlineQueryResultTypeVenue,
				ID:   id,
			},
			Latitude:  latitude,
			Longitude: longitude,
			Title:     title,
			Address:   address,
		}, &id
	}

	return InlineQueryResultVenue{}, nil
}

// SetFoursquareID sets the `foursquare_id` value of InlineQueryResultVenue.
func (r InlineQueryResultVenue) SetFoursquareID(foursquareID string) InlineQueryResultVenue {
	r.FoursquareID = &foursquareID
	return r
}

// SetFoursquareType sets the `foursquare_type` value of InlineQueryResultVenue.
func (r InlineQueryResultVenue) SetFoursquareType(foursquareType string) InlineQueryResultVenue {
	r.FoursquareType = &foursquareType
	return r
}

// SetGooglePlaceID sets the `google_place_id` value of InlineQueryResultVenue.
func (r InlineQueryResultVenue) SetGooglePlaceID(googlePlaceID string) InlineQueryResultVenue {
	r.GooglePlaceID = &googlePlaceID
	return r
}

// SetGooglePlaceType sets the `google_place_type` value of InlineQueryResultVenue.
func (r InlineQueryResultVenue) SetGooglePlaceType(googlePlaceType string) InlineQueryResultVenue {
	r.GooglePlaceType = &googlePlaceType
	return r
}

// SetReplyMarkup sets the `reply_markup` value of InlineQueryResultVenue.
func (r InlineQueryResultVenue) SetReplyMarkup(markup InlineKeyboardMarkup) InlineQueryResultVenue {
	r.ReplyMarkup = &markup
	return r
}

// SetInputMessageContent sets the `input_message_content` value of InlineQueryResultVenue.
func (r InlineQueryResultVenue) SetInputMessageContent(content InputMessageContent) InlineQueryResultVenue {
	r.InputMessageContent = &content
	return r
}

// SetThumbnailURL sets the `thumbnail_url` value of InlineQueryResultVenue.
func (r InlineQueryResultVenue) SetThumbnailURL(thumbnailURL string) InlineQueryResultVenue {
	r.ThumbnailURL = &thumbnailURL
	return r
}

// SetThumbnailWidth sets the `thumbnail_width` value of InlineQueryResultVenue.
func (r InlineQueryResultVenue) SetThumbnailWidth(thumbnailWidth int) InlineQueryResultVenue {
	r.ThumbnailWidth = &thumbnailWidth
	return r
}

// SetThumbnailHeight sets the `thumbnail_height` value of InlineQueryResultVenue.
func (r InlineQueryResultVenue) SetThumbnailHeight(thumbnailHeight int) InlineQueryResultVenue {
	r.ThumbnailHeight = &thumbnailHeight
	return r
}

// NewInlineQueryResultContact is a helper function for generating a new InlineQueryResultContact.
//
// https://core.telegram.org/bots/api#inlinequeryresultcontact
func NewInlineQueryResultContact(phoneNumber, firstName string) (newContact InlineQueryResultContact, generatedID *string) {
	if id, err := newUUID(); err == nil {
		return InlineQueryResultContact{
			InlineQueryResult: InlineQueryResult{
				Type: InlineQueryResultTypeContact,
				ID:   id,
			},
			PhoneNumber: phoneNumber,
			FirstName:   firstName,
		}, &id
	}

	return InlineQueryResultContact{}, nil
}

// SetLastName sets the `last_name` value of InlineQueryResultContact.
func (r InlineQueryResultContact) SetLastName(lastName string) InlineQueryResultContact {
	r.LastName = &lastName
	return r
}

// SetVCard sets the `vcard` value of InlineQueryResultContact.
func (r InlineQueryResultContact) SetVCard(vCard string) InlineQueryResultContact {
	r.VCard = &vCard
	return r
}

// SetReplyMarkup sets the `reply_markup` value of InlineQueryResultContact.
func (r InlineQueryResultContact) SetReplyMarkup(markup InlineKeyboardMarkup) InlineQueryResultContact {
	r.ReplyMarkup = &markup
	return r
}

// SetInputMessageContent sets the `input_message_content` value of InlineQueryResultContact.
func (r InlineQueryResultContact) SetInputMessageContent(content InputMessageContent) InlineQueryResultContact {
	r.InputMessageContent = &content
	return r
}

// SetThumbnailURL sets the `thumbnail_url` value of InlineQueryResultContact.
func (r InlineQueryResultContact) SetThumbnailURL(thumbnailURL string) InlineQueryResultContact {
	r.ThumbnailURL = &thumbnailURL
	return r
}

// SetThumbnailWidth sets the `thumbnail_width` value of InlineQueryResultContact.
func (r InlineQueryResultContact) SetThumbnailWidth(thumbnailWidth int) InlineQueryResultContact {
	r.ThumbnailWidth = &thumbnailWidth
	return r
}

// SetThumbnailHeight sets the `thumbnail_height` value of InlineQueryResultContact.
func (r InlineQueryResultContact) SetThumbnailHeight(thumbnailHeight int) InlineQueryResultContact {
	r.ThumbnailHeight = &thumbnailHeight
	return r
}

// NewInlineQueryResultGame is a helper function for generating a new InlineQueryResultGame.
//
// https://core.telegram.org/bots/api#inlinequeryresultgame
func NewInlineQueryResultGame(shortName string) (newGame InlineQueryResultGame, generatedID *string) {
	if id, err := newUUID(); err == nil {
		return InlineQueryResultGame{
			InlineQueryResult: InlineQueryResult{
				Type: InlineQueryResultTypeGame,
				ID:   id,
			},
			GameShortName: shortName,
		}, &id
	}

	return InlineQueryResultGame{}, nil
}

// SetReplyMarkup sets the `reply_markup` value of InlineQueryResultGame.
func (r InlineQueryResultGame) SetReplyMarkup(markup InlineKeyboardMarkup) InlineQueryResultGame {
	r.ReplyMarkup = &markup
	return r
}

// NewInlineQueryResultCachedPhoto is a helper function for generating a new InlineQueryResultCachedPhoto.
//
// https://core.telegram.org/bots/api#inlinequeryresultcachedphoto
func NewInlineQueryResultCachedPhoto(photoFileID string) (newPhoto InlineQueryResultCachedPhoto, generatedID *string) {
	if id, err := newUUID(); err == nil {
		return InlineQueryResultCachedPhoto{
			InlineQueryResult: InlineQueryResult{
				Type: InlineQueryResultTypePhoto,
				ID:   id,
			},
			PhotoFileID: photoFileID,
		}, &id
	}

	return InlineQueryResultCachedPhoto{}, nil
}

// SetTitle sets the `title` value of InlineQueryResultCachedPhoto.
func (r InlineQueryResultCachedPhoto) SetTitle(title string) InlineQueryResultCachedPhoto {
	r.Title = &title
	return r
}

// SetDescription sets the `description` value of InlineQueryResultCachedPhoto.
func (r InlineQueryResultCachedPhoto) SetDescription(description string) InlineQueryResultCachedPhoto {
	r.Description = &description
	return r
}

// SetCaption sets the `caption` value of InlineQueryResultCachedPhoto.
func (r InlineQueryResultCachedPhoto) SetCaption(caption string) InlineQueryResultCachedPhoto {
	r.Caption = &caption
	return r
}

// SetParseMode sets the `parse_mode` value of InlineQueryResultCachedPhoto.
func (r InlineQueryResultCachedPhoto) SetParseMode(parseMode ParseMode) InlineQueryResultCachedPhoto {
	r.ParseMode = &parseMode
	return r
}

// SetCaptionEntities sets the `caption_entities` value of InlineQueryResultCachedPhoto.
func (r InlineQueryResultCachedPhoto) SetCaptionEntities(entities []MessageEntity) InlineQueryResultCachedPhoto {
	r.CaptionEntities = entities
	return r
}

// SetReplyMarkup sets the `reply_markup` value of InlineQueryResultCachedPhoto.
func (r InlineQueryResultCachedPhoto) SetReplyMarkup(markup InlineKeyboardMarkup) InlineQueryResultCachedPhoto {
	r.ReplyMarkup = &markup
	return r
}

// SetInputMessageContent sets the `input_message_content` value of InlineQueryResultCachedPhoto.
func (r InlineQueryResultCachedPhoto) SetInputMessageContent(content InputMessageContent) InlineQueryResultCachedPhoto {
	r.InputMessageContent = &content
	return r
}

// NewInlineQueryResultCachedGif is a helper function for generating a new InlineQueryResultCachedGif.
//
// https://core.telegram.org/bots/api#inlinequeryresultcachedgif
func NewInlineQueryResultCachedGif(gifFileID string) (newGif InlineQueryResultCachedGif, generatedID *string) {
	if id, err := newUUID(); err == nil {
		return InlineQueryResultCachedGif{
			InlineQueryResult: InlineQueryResult{
				Type: InlineQueryResultTypeGif,
				ID:   id,
			},
			GifFileID: gifFileID,
		}, &id
	}

	return InlineQueryResultCachedGif{}, nil
}

// SetTitle sets the `title` value of InlineQueryResultCachedGif.
func (r InlineQueryResultCachedGif) SetTitle(title string) InlineQueryResultCachedGif {
	r.Title = &title
	return r
}

// SetCaption sets the `caption` value of InlineQueryResultCachedGif.
func (r InlineQueryResultCachedGif) SetCaption(caption string) InlineQueryResultCachedGif {
	r.Caption = &caption
	return r
}

// SetParseMode sets the `parse_mode` value of InlineQueryResultCachedGif.
func (r InlineQueryResultCachedGif) SetParseMode(parseMode ParseMode) InlineQueryResultCachedGif {
	r.ParseMode = &parseMode
	return r
}

// SetCaptionEntities sets the `caption_entities` value of InlineQueryResultCachedGif.
func (r InlineQueryResultCachedGif) SetCaptionEntities(entities []MessageEntity) InlineQueryResultCachedGif {
	r.CaptionEntities = entities
	return r
}

// SetReplyMarkup sets the `reply_markup` value of InlineQueryResultCachedGif.
func (r InlineQueryResultCachedGif) SetReplyMarkup(markup InlineKeyboardMarkup) InlineQueryResultCachedGif {
	r.ReplyMarkup = &markup
	return r
}

// SetInputMessageContent sets the `input_message_content` value of InlineQueryResultCachedGif.
func (r InlineQueryResultCachedGif) SetInputMessageContent(content InputMessageContent) InlineQueryResultCachedGif {
	r.InputMessageContent = &content
	return r
}

// NewInlineQueryResultCachedMpeg4Gif is a helper function for generating a new InlineQueryResultCachedMpeg4Gif.
//
// https://core.telegram.org/bots/api#inlinequeryresultcachedmpeg4gif
func NewInlineQueryResultCachedMpeg4Gif(mpeg4FileID string) (newMpeg4Gif InlineQueryResultCachedMpeg4Gif, generatedID *string) {
	if id, err := newUUID(); err == nil {
		return InlineQueryResultCachedMpeg4Gif{
			InlineQueryResult: InlineQueryResult{
				Type: InlineQueryResultTypeMpeg4Gif,
				ID:   id,
			},
			Mpeg4FileID: mpeg4FileID,
		}, &id
	}

	return InlineQueryResultCachedMpeg4Gif{}, nil
}

// SetTitle sets the `title` value of InlineQueryResultCachedMpeg4Gif.
func (r InlineQueryResultCachedMpeg4Gif) SetTitle(title string) InlineQueryResultCachedMpeg4Gif {
	r.Title = &title
	return r
}

// SetCaption sets the `caption` value of InlineQueryResultCachedMpeg4Gif.
func (r InlineQueryResultCachedMpeg4Gif) SetCaption(caption string) InlineQueryResultCachedMpeg4Gif {
	r.Caption = &caption
	return r
}

// SetParseMode sets the `parse_mode` value of InlineQueryResultCachedMpeg4Gif.
func (r InlineQueryResultCachedMpeg4Gif) SetParseMode(parseMode ParseMode) InlineQueryResultCachedMpeg4Gif {
	r.ParseMode = &parseMode
	return r
}

// SetCaptionEntities sets the `caption_entities` value of InlineQueryResultCachedMpeg4Gif.
func (r InlineQueryResultCachedMpeg4Gif) SetCaptionEntities(entities []MessageEntity) InlineQueryResultCachedMpeg4Gif {
	r.CaptionEntities = entities
	return r
}

// SetReplyMarkup sets the `reply_markup` value of InlineQueryResultCachedMpeg4Gif.
func (r InlineQueryResultCachedMpeg4Gif) SetReplyMarkup(markup InlineKeyboardMarkup) InlineQueryResultCachedMpeg4Gif {
	r.ReplyMarkup = &markup
	return r
}

// SetInputMessageContent sets the `input_message_content` value of InlineQueryResultCachedMpeg4Gif.
func (r InlineQueryResultCachedMpeg4Gif) SetInputMessageContent(content InputMessageContent) InlineQueryResultCachedMpeg4Gif {
	r.InputMessageContent = &content
	return r
}

// NewInlineQueryResultCachedSticker is a helper function for generating a new InlineQueryResultCachedSticker.
//
// https://core.telegram.org/bots/api#inlinequeryresultcachedsticker
func NewInlineQueryResultCachedSticker(stickerFileID string) (newSticker InlineQueryResultCachedSticker, generatedID *string) {
	if id, err := newUUID(); err == nil {
		return InlineQueryResultCachedSticker{
			InlineQueryResult: InlineQueryResult{
				Type: InlineQueryResultTypeSticker,
				ID:   id,
			},
			StickerFileID: stickerFileID,
		}, &id
	}

	return InlineQueryResultCachedSticker{}, nil
}

// SetReplyMarkup sets the `reply_markup` value of InlineQueryResultCachedSticker.
func (r InlineQueryResultCachedSticker) SetReplyMarkup(markup InlineKeyboardMarkup) InlineQueryResultCachedSticker {
	r.ReplyMarkup = &markup
	return r
}

// SetInputMessageContent sets the `input_message_content` value of InlineQueryResultCachedSticker.
func (r InlineQueryResultCachedSticker) SetInputMessageContent(content InputMessageContent) InlineQueryResultCachedSticker {
	r.InputMessageContent = &content
	return r
}

// NewInlineQueryResultCachedDocument is a helper function for generating a new InlineQueryResultCachedDocument.
//
// https://core.telegram.org/bots/api#inlinequeryresultcacheddocument
func NewInlineQueryResultCachedDocument(title, documentFileID string) (newDocument InlineQueryResultCachedDocument, generatedID *string) {
	if id, err := newUUID(); err == nil {
		return InlineQueryResultCachedDocument{
			InlineQueryResult: InlineQueryResult{
				Type: InlineQueryResultTypeDocument,
				ID:   id,
			},
			Title:          title,
			DocumentFileID: documentFileID,
		}, &id
	}

	return InlineQueryResultCachedDocument{}, nil
}

// SetCaption sets the `caption` value of InlineQueryResultCachedDocument.
func (r InlineQueryResultCachedDocument) SetCaption(caption string) InlineQueryResultCachedDocument {
	r.Caption = &caption
	return r
}

// SetParseMode sets the `parse_mode` value of InlineQueryResultCachedDocument.
func (r InlineQueryResultCachedDocument) SetParseMode(parseMode ParseMode) InlineQueryResultCachedDocument {
	r.ParseMode = &parseMode
	return r
}

// SetCaptionEntities sets the `caption_entities` value of InlineQueryResultCachedDocument.
func (r InlineQueryResultCachedDocument) SetCaptionEntities(entities []MessageEntity) InlineQueryResultCachedDocument {
	r.CaptionEntities = entities
	return r
}

// SetDescription sets the `description` value of InlineQueryResultCachedDocument.
func (r InlineQueryResultCachedDocument) SetDescription(description string) InlineQueryResultCachedDocument {
	r.Description = &description
	return r
}

// SetReplyMarkup sets the `reply_markup` value of InlineQueryResultCachedDocument.
func (r InlineQueryResultCachedDocument) SetReplyMarkup(markup InlineKeyboardMarkup) InlineQueryResultCachedDocument {
	r.ReplyMarkup = &markup
	return r
}

// SetInputMessageContent sets the `input_message_content` value of InlineQueryResultCachedDocument.
func (r InlineQueryResultCachedDocument) SetInputMessageContent(content InputMessageContent) InlineQueryResultCachedDocument {
	r.InputMessageContent = &content
	return r
}

// NewInlineQueryResultCachedVideo is a helper function for generating a new InlineQueryResultCachedVideo.
//
// https://core.telegram.org/bots/api#inlinequeryresultcachedvideo
func NewInlineQueryResultCachedVideo(title, videoFileID string) (newVideo InlineQueryResultCachedVideo, generatedID *string) {
	if id, err := newUUID(); err == nil {
		return InlineQueryResultCachedVideo{
			InlineQueryResult: InlineQueryResult{
				Type: InlineQueryResultTypeVideo,
				ID:   id,
			},
			Title:       title,
			VideoFileID: videoFileID,
		}, &id
	}

	return InlineQueryResultCachedVideo{}, nil
}

// SetCaption sets the `caption` value of InlineQueryResultCachedVideo.
func (r InlineQueryResultCachedVideo) SetCaption(caption string) InlineQueryResultCachedVideo {
	r.Caption = &caption
	return r
}

// SetParseMode sets the `parse_mode` value of InlineQueryResultCachedVideo.
func (r InlineQueryResultCachedVideo) SetParseMode(parseMode ParseMode) InlineQueryResultCachedVideo {
	r.ParseMode = &parseMode
	return r
}

// SetCaptionEntities sets the `caption_entities` value of InlineQueryResultCachedVideo.
func (r InlineQueryResultCachedVideo) SetCaptionEntities(entities []MessageEntity) InlineQueryResultCachedVideo {
	r.CaptionEntities = entities
	return r
}

// SetDescription sets the `description` value of InlineQueryResultCachedVideo.
func (r InlineQueryResultCachedVideo) SetDescription(description string) InlineQueryResultCachedVideo {
	r.Description = &description
	return r
}

// SetReplyMarkup sets the `reply_markup` value of InlineQueryResultCachedVideo.
func (r InlineQueryResultCachedVideo) SetReplyMarkup(markup InlineKeyboardMarkup) InlineQueryResultCachedVideo {
	r.ReplyMarkup = &markup
	return r
}

// SetInputMessageContent sets the `input_message_content` value of InlineQueryResultCachedVideo.
func (r InlineQueryResultCachedVideo) SetInputMessageContent(content InputMessageContent) InlineQueryResultCachedVideo {
	r.InputMessageContent = &content
	return r
}

// NewInlineQueryResultCachedVoice is a helper function for generating a new InlineQueryResultCachedVoice.
//
// https://core.telegram.org/bots/api#inlinequeryresultcachedvoice
func NewInlineQueryResultCachedVoice(title, voiceFileID string) (newVoice InlineQueryResultCachedVoice, generatedID *string) {
	if id, err := newUUID(); err == nil {
		return InlineQueryResultCachedVoice{
			InlineQueryResult: InlineQueryResult{
				Type: InlineQueryResultTypeVoice,
				ID:   id,
			},
			Title:       title,
			VoiceFileID: voiceFileID,
		}, &id
	}

	return InlineQueryResultCachedVoice{}, nil
}

// SetCaption sets the `caption` value of InlineQueryResultCachedVoice.
func (r InlineQueryResultCachedVoice) SetCaption(caption string) InlineQueryResultCachedVoice {
	r.Caption = &caption
	return r
}

// SetParseMode sets the `parse_mode` value of InlineQueryResultCachedVoice.
func (r InlineQueryResultCachedVoice) SetParseMode(parseMode ParseMode) InlineQueryResultCachedVoice {
	r.ParseMode = &parseMode
	return r
}

// SetCaptionEntities sets the `caption_entities` value of InlineQueryResultCachedVoice.
func (r InlineQueryResultCachedVoice) SetCaptionEntities(entities []MessageEntity) InlineQueryResultCachedVoice {
	r.CaptionEntities = entities
	return r
}

// SetReplyMarkup sets the `reply_markup` value of InlineQueryResultCachedVoice.
func (r InlineQueryResultCachedVoice) SetReplyMarkup(markup InlineKeyboardMarkup) InlineQueryResultCachedVoice {
	r.ReplyMarkup = &markup
	return r
}

// SetInputMessageContent sets the `input_message_content` value of InlineQueryResultCachedVoice.
func (r InlineQueryResultCachedVoice) SetInputMessageContent(content InputMessageContent) InlineQueryResultCachedVoice {
	r.InputMessageContent = &content
	return r
}

// NewInlineQueryResultCachedAudio is a helper function for generating a new InlineQueryResultCachedAudio.
//
// https://core.telegram.org/bots/api#inlinequeryresultcachedaudio
func NewInlineQueryResultCachedAudio(audioFileID string) (newAudio InlineQueryResultCachedAudio, generatedID *string) {
	if id, err := newUUID(); err == nil {
		return InlineQueryResultCachedAudio{
			InlineQueryResult: InlineQueryResult{
				Type: InlineQueryResultTypeAudio,
				ID:   id,
			},
			AudioFileID: audioFileID,
		}, &id
	}

	return InlineQueryResultCachedAudio{}, nil
}

// SetCaption sets the `caption` value of InlineQueryResultCachedAudio.
func (r InlineQueryResultCachedAudio) SetCaption(caption string) InlineQueryResultCachedAudio {
	r.Caption = &caption
	return r
}

// SetParseMode sets the `parse_mode` value of InlineQueryResultCachedAudio.
func (r InlineQueryResultCachedAudio) SetParseMode(parseMode ParseMode) InlineQueryResultCachedAudio {
	r.ParseMode = &parseMode
	return r
}

// SetCaptionEntities sets the `caption_entities` value of InlineQueryResultCachedAudio.
func (r InlineQueryResultCachedAudio) SetCaptionEntities(entities []MessageEntity) InlineQueryResultCachedAudio {
	r.CaptionEntities = entities
	return r
}

// SetReplyMarkup sets the `reply_markup` value of InlineQueryResultCachedAudio.
func (r InlineQueryResultCachedAudio) SetReplyMarkup(markup InlineKeyboardMarkup) InlineQueryResultCachedAudio {
	r.ReplyMarkup = &markup
	return r
}

// SetInputMessageContent sets the `input_message_content` value of InlineQueryResultCachedAudio.
func (r InlineQueryResultCachedAudio) SetInputMessageContent(content InputMessageContent) InlineQueryResultCachedAudio {
	r.InputMessageContent = &content
	return r
}

////////////////////////////////
// Helper functions for InputMessageContent

// NewInputTextMessageContent returns a new InputTextMessageContent with given text.
func NewInputTextMessageContent(text string) InputTextMessageContent {
	return InputTextMessageContent{
		MessageText: text,
	}
}

// SetParseMode sets the `parse_mode` value of InputTextMessageContent.
func (c InputTextMessageContent) SetParseMode(parseMode ParseMode) InputTextMessageContent {
	c.ParseMode = &parseMode
	return c
}

// SetCaptionEntities sets the `caption_entities` value of InputTextMessageContent.
func (c InputTextMessageContent) SetCaptionEntities(entities []MessageEntity) InputTextMessageContent {
	c.CaptionEntities = entities
	return c
}

// SetLinkPreviewOptions sets the `link_preview_options` value of InputTextMessageContent.
func (c InputTextMessageContent) SetLinkPreviewOptions(options LinkPreviewOptions) InputTextMessageContent {
	c.LinkPreviewOptions = &options
	return c
}

// NewInputLocationMessageContent returns a new InputLocationMessageContent.
func NewInputLocationMessageContent(latitude, longitude float32) InputLocationMessageContent {
	return InputLocationMessageContent{
		Latitude:  latitude,
		Longitude: longitude,
	}
}

// SetHorizontalAccuracy sets the `horizontal_accuracy` value of InputLocationMessageContent.
func (c InputLocationMessageContent) SetHorizontalAccuracy(horizontalAccuracy float32) InputLocationMessageContent {
	c.HorizontalAccuracy = &horizontalAccuracy
	return c
}

// SetLivePeriod sets the `live_period` value of InputLocationMessageContent.
func (c InputLocationMessageContent) SetLivePeriod(livePeriod int) InputLocationMessageContent {
	c.LivePeriod = &livePeriod
	return c
}

// SetHeading sets the `heading` value of InputLocationMessageContent.
func (c InputLocationMessageContent) SetHeading(heading int) InputLocationMessageContent {
	c.Heading = &heading
	return c
}

// SetProximityAlertRadius sets the `proximity_alert_radius` value of InputLocationMessageContent.
func (c InputLocationMessageContent) SetProximityAlertRadius(radius int) InputLocationMessageContent {
	c.ProximityAlertRadius = &radius
	return c
}

// NewInputVenueMessageContent returns a new InputVenueMessageContent.
func NewInputVenueMessageContent(latitude, longitude float32, title, address string) InputVenueMessageContent {
	return InputVenueMessageContent{
		Latitude:  latitude,
		Longitude: longitude,
		Title:     title,
		Address:   address,
	}
}

// SetFoursquareID sets the `foursquare_id` value of InputVenueMessageContent.
func (c InputVenueMessageContent) SetFoursquareID(foursquareID string) InputVenueMessageContent {
	c.FoursquareID = &foursquareID
	return c
}

// SetFoursquareType sets the `foursquare_type` value of InputVenueMessageContent.
func (c InputVenueMessageContent) SetFoursquareType(foursquareType string) InputVenueMessageContent {
	c.FoursquareType = &foursquareType
	return c
}

// SetGooglePlaceID sets the `google_place_id` value of InputVenueMessageContent.
func (c InputVenueMessageContent) SetGooglePlaceID(googlePlaceID string) InputVenueMessageContent {
	c.GooglePlaceID = &googlePlaceID
	return c
}

// SetGooglePlaceType sets the `google_place_type` value of InputVenueMessageContent.
func (c InputVenueMessageContent) SetGooglePlaceType(googlePlaceType string) InputVenueMessageContent {
	c.GooglePlaceType = &googlePlaceType
	return c
}

// NewInputContactMessageContent returns a new InputContactMessageContent.
func NewInputContactMessageContent(phoneNumber, firstName string) InputContactMessageContent {
	return InputContactMessageContent{
		PhoneNumber: phoneNumber,
		FirstName:   firstName,
	}
}

// SetLastName sets the `last_name` value of InputContactMessageContent.
func (c InputContactMessageContent) SetLastName(lastName string) InputContactMessageContent {
	c.LastName = &lastName
	return c
}

// SetVCard sets the `vcard` value of InputContactMessageContent.
func (c InputContactMessageContent) SetVCard(vCard string) InputContactMessageContent {
	c.VCard = &vCard
	return c
}

// NewInputInvoiceMessageContent returns a new InputInvoiceMessageContent.
func NewInputInvoiceMessageContent(title, description, payload, providerToken, currency string, prices []LabeledPrice) InputInvoiceMessageContent {
	return InputInvoiceMessageContent{
		Title:         title,
		Description:   description,
		Payload:       payload,
		ProviderToken: providerToken,
		Currency:      currency,
		Prices:        prices,
	}
}

// SetMaxTipAmount sets the `max_tip_amount` value of InputInvoiceMessageContent.
func (c InputInvoiceMessageContent) SetMaxTipAmount(amount int) InputInvoiceMessageContent {
	c.MaxTipAmount = &amount
	return c
}

// SetSuggestedTipAmounts sets the `suggested_tip_amounts` value of InputInvoiceMessageContent.
func (c InputInvoiceMessageContent) SetSuggestedTipAmounts(amounts []int) InputInvoiceMessageContent {
	c.SuggestedTipAmounts = amounts
	return c
}

// SetProviderData sets the `provider_data` value of InputInvoiceMessageContent.
func (c InputInvoiceMessageContent) SetProviderData(providerData string) InputInvoiceMessageContent {
	c.ProviderData = &providerData
	return c
}

// SetPhotoURL sets the `photo_url` value of InputInvoiceMessageContent.
func (c InputInvoiceMessageContent) SetPhotoURL(photoURL string) InputInvoiceMessageContent {
	c.PhotoURL = &photoURL
	return c
}

// SetPhotoSize sets the `photo_size` value of InputInvoiceMessageContent.
func (c InputInvoiceMessageContent) SetPhotoSize(photoSize int) InputInvoiceMessageContent {
	c.PhotoSize = &photoSize
	return c
}

// SetPhotoWidth sets the `photo_width` value of InputInvoiceMessageContent.
func (c InputInvoiceMessageContent) SetPhotoWidth(photoWidth int) InputInvoiceMessageContent {
	c.PhotoWidth = &photoWidth
	return c
}

// SetPhotoHeight sets the `photo_height` value of InputInvoiceMessageContent.
func (c InputInvoiceMessageContent) SetPhotoHeight(photoHeight int) InputInvoiceMessageContent {
	c.PhotoHeight = &photoHeight
	return c
}

// SetNeedName sets the `need_name` value of InputInvoiceMessageContent.
func (c InputInvoiceMessageContent) SetNeedName(needName bool) InputInvoiceMessageContent {
	c.NeedName = &needName
	return c
}

// SetNeedPhoneNumber sets the `need_phone_number` value of InputInvoiceMessageContent.
func (c InputInvoiceMessageContent) SetNeedPhoneNumber(needPhoneNumber bool) InputInvoiceMessageContent {
	c.NeedPhoneNumber = &needPhoneNumber
	return c
}

// SetNeedEmail sets the `need_email` value of InputInvoiceMessageContent.
func (c InputInvoiceMessageContent) SetNeedEmail(needEmail bool) InputInvoiceMessageContent {
	c.NeedEmail = &needEmail
	return c
}

// SetNeedShippingAddress sets the `need_shipping_address` value of InputInvoiceMessageContent.
func (c InputInvoiceMessageContent) SetNeedShippingAddress(needShippingAddress bool) InputInvoiceMessageContent {
	c.NeedShippingAddress = &needShippingAddress
	return c
}

// SetSendPhoneNumberToProvider sets the `send_phone_number_to_provider` value of InputInvoiceMessageContent.
func (c InputInvoiceMessageContent) SetSendPhoneNumberToProvider(sendPhoneNumberToProvider bool) InputInvoiceMessageContent {
	c.SendPhoneNumberToProvider = &sendPhoneNumberToProvider
	return c
}

// SetSendEmailToProvider sets the `send_email_to_provider` value of InputInvoiceMessageContent.
func (c InputInvoiceMessageContent) SetSendEmailToProvider(sendEmailToProvider bool) InputInvoiceMessageContent {
	c.SendEmailToProvider = &sendEmailToProvider
	return c
}

// SetIsFlexible sets the `is_flexible` value of InputInvoiceMessageContent.
func (c InputInvoiceMessageContent) SetIsFlexible(isFlexible bool) InputInvoiceMessageContent {
	c.IsFlexible = &isFlexible
	return c
}

////////////////////////////////
// Helper functions for WebAppInfo

// NOTE: NOT IMPLEMENTED YET

////////////////////////////////
// Helper functions for ChatBoostSource

// NewChatBoostSourcePremium returns a new ChatBoostSourcePremium.
func NewChatBoostSourcePremium(user User) ChatBoostSource {
	return ChatBoostSource{
		Source: "premium",
		User:   &user,
	}
}

// NewChatBoostSourceGiftCode returns a new ChatBoostSourceGiftCode.
func NewChatBoostSourceGiftCode(user User) ChatBoostSource {
	return ChatBoostSource{
		Source: "gift_code",
		User:   &user,
	}
}

// NewChatBoostSourceGiveaway returns a new ChatBoostSourceGiveaway.
func NewChatBoostSourceGiveaway(giveawayMessageID int64, user *User, isUnclaimed bool) ChatBoostSource {
	return ChatBoostSource{
		Source:            "giveaway",
		GiveawayMessageID: &giveawayMessageID,
		User:              user,
		IsUnclaimed:       &isUnclaimed,
	}
}

////////////////////////////////
// Helper functions for OwnedGift

// GiftAsGift returns `gift` as Gift.
func (g OwnedGift) GiftAsGift() (result *Gift, err error) {
	if g.Type == "regular" {
		if err = json.Unmarshal(g.Gift, &result); err == nil {
			return result, nil
		} else {
			return nil, fmt.Errorf("failed to unmarshal gift: %w", err)
		}
	}
	return nil, fmt.Errorf("gift is not in regular type")
}

// GiftAsUniqueGift returns `gift` as UniqueGift.
func (g OwnedGift) GiftAsUniqueGift() (result *UniqueGift, err error) {
	if g.Type == "unique" {
		if err = json.Unmarshal(g.Gift, &result); err == nil {
			return result, nil
		} else {
			return nil, fmt.Errorf("failed to unmarshal unique gift: %w", err)
		}
	}
	return nil, fmt.Errorf("gift is not in unique type")
}
