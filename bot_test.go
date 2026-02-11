// bot_test.go
//
// created on: 2023.11.10.

package telegrambot

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"strings"
	"testing"
	"time"
)

// test timeouts and their handling
func TestTimeout(t *testing.T) {
	_token := os.Getenv("TOKEN")
	_verbose := os.Getenv("VERBOSE")

	client := NewClient(_token)
	client.Verbose = _verbose == "true"

	if len(_token) <= 0 {
		t.Errorf("environment variable `TOKEN` is needed")
	}

	slog.Info("testing timeouts...")

	ctx, cancel := context.WithTimeout(context.TODO(), 1*time.Nanosecond)
	defer cancel()

	// intentional timeout
	if _, err := client.GetMe(ctx); err != nil {
		var e ErrContextTimeout
		if !errors.As(err, &e) {
			t.Errorf("expected `ErrContextTimeout` but got: %[1]s (%[1]T)", err)
		}
	}
}

// test polling updates
func TestPollingUpdates(t *testing.T) {
	_token := os.Getenv("TOKEN")
	_verbose := os.Getenv("VERBOSE")

	client := NewClient(_token)
	client.Verbose = _verbose == "true"

	if len(_token) <= 0 {
		t.Errorf("environment variable `TOKEN` is needed")
	}

	slog.Info("testing polling updates...")

	if deleted, err := client.DeleteWebhook(context.TODO(), true); err != nil {
		t.Errorf("failed to delete webhook before testing polling updates: %s", *deleted.Description)
	} else {
		go func() {
			time.Sleep(10 * time.Second) // sleep for a while,

			client.StopPollingUpdates() // stop polling
		}()

		// polling is synchronous
		client.StartPollingUpdates(0, 1, func(b *Bot, update Update, err error) {
			if err != nil {
				t.Errorf("error while polling updates: %s", err)
			}
		}, []AllowedUpdate{AllowMessage})

		slog.Info("stopped polling updates")
	}
}

// test API method functions
func TestMethods(t *testing.T) {
	_token := os.Getenv("TOKEN")
	_chatID := os.Getenv("CHAT_ID") // NOTE: `chat_id` of a group chat with topics
	_verbose := os.Getenv("VERBOSE")

	client := NewClient(_token)
	client.Verbose = _verbose == "true"

	if len(_token) <= 0 || len(_chatID) <= 0 {
		t.Errorf("environment variables `TOKEN` and `CHAT_ID` are needed")
	}

	slog.Info("testing API method functions...")

	////////////////////////////////
	// (bot info)
	//
	// GetMe
	if me, _ := client.GetMe(context.TODO()); !me.OK {
		t.Errorf("failed to get me: %s", *me.Description)
	} else {
		if client.Verbose {
			marshalled, _ := json.Marshal(me.Result)
			slog.Info("fetched bot info", "user", marshalled)
		}

		// show warnings on missing permissions
		if me.Result.SupportsInlineQueries != nil && !*me.Result.SupportsInlineQueries {
			slog.Warn("bot does not support inline queries")
		}
		if me.Result.HasTopicsEnabled != nil && !*me.Result.HasTopicsEnabled {
			slog.Warn("bot does not have topics enabled")
		}
		if me.Result.AllowsUsersToCreateTopics != nil && !*me.Result.AllowsUsersToCreateTopics {
			slog.Warn("bot does not allow users to create topics")
		}

		////////////////////////////////
		// (webhook)
		//
		// SetWebhook
		if webhook, _ := client.SetWebhook(
			context.TODO(),
			"testdomain.com",
			8443,
			OptionsSetWebhook{},
		); !webhook.OK {
			t.Errorf("failed to set webhook: %s", *webhook.Description)
		} else {
			// GetWebhookInfo
			if webhook, _ := client.GetWebhookInfo(context.TODO()); !webhook.OK {
				t.Errorf("failed to get webhook info: %s", *webhook.Description)
			}

			// DeleteWebhook
			if deleted, _ := client.DeleteWebhook(
				context.TODO(),
				false,
			); !deleted.OK {
				t.Errorf("failed to delete webhook: %s", *webhook.Description)
			}

			////////////////////////////////
			// (general methods)
			//
			// GetUpdates
			if updates, _ := client.GetUpdates(
				context.TODO(),
				OptionsGetUpdates{},
			); !updates.OK {
				t.Errorf("failed to get updates: %s", *updates.Description)
			}
			// TODO: LogOut
			// TODO: Close
			// SendMessage
			if sent, _ := client.SendMessage(
				context.TODO(),
				_chatID,
				"test message",
				OptionsSendMessage{},
			); !sent.OK {
				t.Errorf("failed to send message: %s", *sent.Description)
			} else {
				// EditMessageText
				if edited, _ := client.EditMessageText(
					context.TODO(),
					"edited message",
					OptionsEditMessageText{}.
						SetIDs(_chatID, sent.Result.MessageID),
				); !edited.OK {
					t.Errorf("failed to edit message text: %s", *edited.Description)
				}
				// CopyMessage
				if copied, _ := client.CopyMessage(
					context.TODO(),
					_chatID,
					_chatID,
					sent.Result.MessageID,
					OptionsCopyMessage{},
				); !copied.OK {
					t.Errorf("failed to copy message: %s", *copied.Description)
				}
				// ForwardMessage
				if forwarded, _ := client.ForwardMessage(
					context.TODO(),
					_chatID,
					_chatID,
					sent.Result.MessageID,
					OptionsForwardMessage{},
				); !forwarded.OK {
					t.Errorf("failed to forward message: %s", *forwarded.Description)
				}
			}
			// TODO: SendMessageDraft
			// TODO: CopyMessages
			// TODO: ForwardMessages
			// SendPhoto
			if photo, _ := client.SendPhoto(
				context.TODO(),
				_chatID,
				NewInputFileFromFilepath("./samples/_files/gopher.png"),
				OptionsSendPhoto{},
			); !photo.OK {
				t.Errorf("failed to send photo: %s", *photo.Description)
			} else {
				// EditMessageCaption
				if caption, _ := client.EditMessageCaption(
					context.TODO(),
					OptionsEditMessageCaption{}.
						SetIDs(_chatID, photo.Result.MessageID).
						SetCaption("edited caption"),
				); !caption.OK {
					t.Errorf("failed to edit message caption: %s", *caption.Description)
				}
			}
			// TODO: SendAudio
			// TODO: SendDocument
			if doc, _ := client.SendDocument(
				context.TODO(),
				_chatID,
				NewInputFileFromFilepath("./samples/_files/gopher.png"),
				OptionsSendDocument{},
			); !doc.OK {
				t.Errorf("failed to send document: %s", *doc.Description)
			} else {
				// GetFile
				if file, _ := client.GetFile(
					context.TODO(),
					doc.Result.Document.FileID,
				); !file.OK {
					t.Errorf("failed to get file: %s", *file.Description)
				}
				// DeleteMessage
				if deleted, _ := client.DeleteMessage(
					context.TODO(),
					_chatID,
					doc.Result.MessageID,
				); !deleted.OK {
					t.Errorf("failed to delete message: %s", *deleted.Description)
				}
			}
			// TODO: DeleteMessages
			// TODO: SendSticker
			// TODO: SendVideo
			// TODO: SendAnimation
			// TODO: SendVoice
			// TODO: SendVideoNote
			// TODO: SendPaidMedia
			// TODO: SendMediaGroup
			// SendLocation
			if location, _ := client.SendLocation(
				context.TODO(),
				_chatID,
				37.5665,
				126.9780,
				OptionsSendLocation{},
			); !location.OK {
				t.Errorf("failed to send location: %s", *location.Description)
			}
			// TODO: SendVenue
			// SendContact
			if contact, _ := client.SendContact(
				context.TODO(),
				_chatID,
				"911",
				"Nine-One-One",
				OptionsSendContact{},
			); !contact.OK {
				t.Errorf("failed to send contact: %s", *contact.Description)
			}
			// SendPoll
			if poll, _ := client.SendPoll(
				context.TODO(),
				_chatID,
				"The earth is...?",
				[]InputPollOption{
					{Text: "flat"},
					{Text: "round"},
					{Text: "nothing"},
				},
				OptionsSendPoll{},
			); !poll.OK {
				t.Errorf("failed to send poll: %s", *poll.Description)
			} else {
				// StopPoll
				if stopped, _ := client.StopPoll(
					context.TODO(),
					_chatID,
					poll.Result.MessageID,
					OptionsStopPoll{},
				); !stopped.OK {
					t.Errorf("failed to stop poll: %s", *stopped.Description)
				}
			}
			// TODO: ApproveSuggestedPost
			// TODO: DeclineSuggestedPost
			// TODO: SendChecklist
			// SendDice
			if dice, _ := client.SendDice(
				context.TODO(),
				_chatID,
				OptionsSendDice{},
			); !dice.OK {
				t.Errorf("failed to send dice: %s", *dice.Description)
			}
			// SendChatAction
			if action, _ := client.SendChatAction(
				context.TODO(),
				_chatID,
				ChatActionTyping,
				OptionsSendChatAction{},
			); !action.OK {
				t.Errorf("failed to send chat action: %s", *action.Description)
			}
			// TODO: GetUserProfilePhotos
			// TODO: ApproveChatJoinRequest
			// TODO: DeclineChatJoinRequest
			// TODO: GetMyCommands
			// GetMyName
			if name, _ := client.GetMyName(
				context.TODO(),
				OptionsGetMyName{},
			); !name.OK {
				t.Errorf("failed to get my name: %s", *name.Description)
			} else {
				newName := "telegram api test bot"

				if name.Result.Name != newName {
					// SetMyName
					if name, _ := client.SetMyName(
						context.TODO(),
						newName,
						OptionsSetMyName{},
					); !name.OK {
						t.Errorf("failed to set my name: %s", *name.Description)
					}
				}
			}
			// SetMyDescription
			if desc, _ := client.SetMyDescription(
				context.TODO(),
				OptionsSetMyDescription{}.
					SetDescription("A bot for testing library: telegram-bot-go"),
			); !desc.OK {
				t.Errorf("failed to set my description: %s", *desc.Description)
			}
			// GetMyDescription
			if desc, _ := client.GetMyDescription(
				context.TODO(),
				OptionsGetMyDescription{},
			); !desc.OK {
				t.Errorf("failed to get my description: %s", *desc.Description)
			}
			// SetMyShortDescription
			if desc, _ := client.SetMyShortDescription(
				context.TODO(),
				OptionsSetMyShortDescription{}.
					SetShortDescription("telegram-bot-go"),
			); !desc.OK {
				t.Errorf("failed to set my short description: %s", *desc.Description)
			}
			// GetMyShortDescription
			if desc, _ := client.GetMyShortDescription(
				context.TODO(),
				OptionsGetMyShortDescription{},
			); !desc.OK {
				t.Errorf("failed to get my short description: %s", *desc.Description)
			}
			// TODO: GetUserProfileAudios
			// TODO: GetUserChatBoosts
			// RemoveMyProfilePhoto - FIXME: Bad Request: BOT_FALLBACK_UNSUPPORTED
			if removed, _ := client.RemoveMyProfilePhoto(
				context.TODO(),
			); !removed.OK {
				t.Errorf("failed to remove my profile photo: %s", *removed.Description)
			}
			// SetMyProfilePhoto
			if photo, _ := client.SetMyProfilePhoto(
				context.TODO(),
				NewInputProfilePhotoFromFilepath(
					InputProfilePhotoStatic,
					"./samples/_files/gopher.jpg",
				),
			); !photo.OK {
				t.Errorf("failed to set my profile photo: %s", *photo.Description)
			}
			// TODO: SetMyCommands
			// TODO: DeleteMyCommands
			// TODO: SetChatMenuButton
			// TODO: GetChatMenuButton
			// TODO: SetMyDefaultAdministratorRights
			// TODO: GetMyDefaultAdministratorRights
			// TODO: EditMessageMedia
			// TODO: EditMessageReplyMarkup
			// TODO: EditMessageLiveLocation
			// TODO: StopMessageLiveLocation
			// TODO: EditMessageChecklist

			////////////////////////////////
			// (business connection)
			//
			// TODO: GetBusinessConnection
			// TODO: ReadBusinessMessage
			// TODO: DeleteBusinessMessages
			// TODO: SetBusinessAccountName
			// TODO: SetBusinessAccountUsername
			// TODO: SetBusinessAccountBio
			// TODO: SetBusinessAccountProfilePhoto
			// TODO: RemoveBusinessAccountProfilePhoto
			// TODO: SetBusinessAccountGiftSettings
			// TODO: GetBusinessAccountStarBalance
			// TODO: TransferBusinessAccountStars
			// TODO: GetBusinessAccountGifts
			// TODO: GetUserGifts
			// TODO: GetChatGifts
			// TODO: ConvertGiftToStars
			// TODO: UpgradeGift
			// TODO: TransferGift
			// TODO: PostStory
			// TODO: RepostStory
			// TODO: EditStory
			// TODO: DeleteStory

			////////////////////////////////
			// (callback query)
			//
			// TODO: AnswerCallbackQuery

			////////////////////////////////
			// (inline query)
			//
			// TODO: AnswerInlineQuery

			////////////////////////////////
			// (sticker)
			//
			// TODO: SendSticker
			// TODO: GetStickerSet
			// TODO: GetCustomEmojiStickers
			// TODO: UploadStickerFile
			// TODO: CreateNewsStickerSet
			// TODO: AddStickerToSet
			// TODO: SetStickerPositionInSet
			// TODO: DeleteStickerFromSet
			// TODO: ReplaceStickerInSet
			// TODO: SetStickerSetThumbnail
			// TODO: SetCustomEmojiStickerSetThumbnail
			// TODO: SetStickerSetTitle
			// TODO: DeleteStickerSet
			// TODO: SetStickerEmojiList
			// TODO: SetStickerKeywords
			// TODO: SetStickerMaskPosition
			// TODO: SetChatStickerSet
			// TODO: DeleteChatStickerSet

			////////////////////////////////
			// (chat administration)
			//
			// GetChat
			if chat, _ := client.GetChat(
				context.TODO(),
				_chatID,
			); !chat.OK {
				t.Errorf("failed to get chat: %s", *chat.Description)
			}
			// GetChatAdministrators
			if admins, _ := client.GetChatAdministrators(
				context.TODO(),
				_chatID,
			); !admins.OK {
				t.Errorf("failed to get chat administrators: %s", *admins.Description)
			}
			// GetChatMemberCount
			if count, _ := client.GetChatMemberCount(
				context.TODO(),
				_chatID,
			); !count.OK {
				t.Errorf("failed to get chat member count: %s", *count.Description)
			}
			// TODO: GetChatMember
			// TODO: CreateChat
			// TODO: SetChatTitle
			// SetChatDescription
			if desc, _ := client.SetChatDescription(
				context.TODO(),
				_chatID,
				fmt.Sprintf("[telegram-bot-go] chat_id: %s (last update: %d)",
					_chatID,
					time.Now().Unix(),
				),
			); !desc.OK {
				t.Errorf("failed to set chat description: %s", *desc.Description)
			}
			// TODO: BanChatMember
			// TODO: LeaveChat
			// TODO: UnbanChatMember
			// TODO: RestrictChatMember
			// TODO: PromoteChatMember
			// TODO: SetChatAdministratorCustomTitle
			// TODO: BanChatSenderChat
			// TODO: UnbanChatSenderChat
			// TODO: SetChatPermissions
			// TODO: SetChatPhoto
			// TODO: DeleteChatPhoto
			// TODO: PinChatMessage
			// TODO: UnpinChatMessage
			// TODO: UnpinAllChatMessages
			// TODO: ExportChatInviteLink
			// TODO: CreateChatInviteLink
			// TODO: EditChatInviteLink
			// TODO: CreateChatSubscriptionInviteLink
			// TODO: EditChatSubscriptionInviteLink
			// TODO: RevokeChatInviteLink

			////////////////////////////////
			// (shopping)
			//
			// TODO: SendInvoice
			// TODO: CreateInvoiceLink
			// TODO: AnswerShippingQuery
			// TODO: AnswerPreCheckoutQuery
			// GetMyStarBalance
			if balance, _ := client.GetMyStarBalance(context.TODO()); !balance.OK {
				t.Errorf("failed to get my star balance: %s", *balance.Description)
			}
			// TODO: GetStarTransactions
			// TODO: RefundStarPayment
			// TODO: EditUserStarSubscription

			////////////////////////////////
			// (forum)
			//
			// CreateForumTopic
			if created, _ := client.CreateForumTopic(
				context.TODO(),
				_chatID,
				fmt.Sprintf("forum topic with chat_id: %s", _chatID),
				OptionsCreateForumTopic{},
			); created.OK {
				_messageThreadID := created.Result.MessageThreadID

				// EditForumTopic
				if edited, _ := client.EditForumTopic(
					context.TODO(),
					_chatID,
					_messageThreadID,
					OptionsEditForumTopic{}.
						SetName(
							fmt.Sprintf(
								"updated forum topic with chat_id: %s, message_thread_id: %d",
								_chatID,
								_messageThreadID,
							),
						),
				); !edited.OK {
					t.Errorf("failed to edit forum topic: %s", *edited.Description)
				}

				// UnpinAllForumTopicMessages
				if unpinned, _ := client.UnpinAllForumTopicMessages(
					context.TODO(),
					_chatID,
					_messageThreadID,
				); !unpinned.OK {
					t.Errorf("failed to unpin all forum topic messages: %s", *unpinned.Description)
				}

				// DeleteForumTopic
				if deleted, _ := client.DeleteForumTopic(
					context.TODO(),
					_chatID,
					_messageThreadID,
				); !deleted.OK {
					t.Errorf("failed to delete forum topic: %s", *deleted.Description)
				}
			} else {
				t.Errorf("failed to create forum topic: %s", *created.Description)
			}
			// TODO: CloseForumTopic
			// TODO: ReopenForumTopic
			// TODO: EditGeneralForumTopic
			// TODO: CloseGeneralForumTopic
			// TODO: ReopenGeneralForumTopic
			// TODO: HideGeneralForumTopic
			// TODO: UnhideGeneralForumTopic
			// TODO: UnpinAllGeneralForumTopicMessages
			// TODO: GetForumTopicIconStickers

			////////////////////////////////
			// (game)
			//
			// TODO: SendGame
			// TODO: SetGameScore
			// TODO: GetGameHighScores

			////////////////////////////////
			// (reaction)
			//
			// TODO: SetMessageReaction

			////////////////////////////////
			// (gift)
			//
			// TODO: GetAvailableGifts
			// TODO: SendGift
			// TODO: GiftPremiumSubscription

			////////////////////////////////
			// (verification)
			//
			// TODO: VerifyUser
			// TODO: RemoveUserVerification
			// TODO: VerifyChat
			// TODO: RemoveChatVerification

			////////////////////////////////
			// (webapp)
			//
			// TODO: AnswerWebAppQuery
			// TODO: SetUserEmojiStatus
			// TODO: SavePreparedInlineMessage
		}
	}
}

// test (un)classified errors
func TestErrors(t *testing.T) {
	_token := os.Getenv("TOKEN")
	_chatID := os.Getenv("CHAT_ID") // NOTE: `chat_id` of a group chat
	_verbose := os.Getenv("VERBOSE")

	client := NewClient(_token)
	client.Verbose = _verbose == "true"

	if len(_token) <= 0 || len(_chatID) <= 0 {
		t.Errorf("environment variables `TOKEN` and `CHAT_ID` are needed")
	}

	slog.Info("testing classification of errors...")

	// ErrUnauthorized
	unauthClient := NewClient("000000000:UNAUTHORIZEDabcdefghijklmnopqrs-0_Z")
	unauthClient.Verbose = _verbose == "true"
	if sent, err := unauthClient.SendMessage(
		context.TODO(),
		_chatID,
		"unauthorized",
		OptionsSendMessage{},
	); !sent.OK {
		var e ErrUnauthorized
		if !errors.As(err, &e) {
			t.Errorf("should have failed with ErrUnauthorized, but got: %s", err)
		}
	} else {
		t.Errorf("should have failed to send unauthorized request")
	}

	// ErrChatNotFound
	if sent, err := client.SendMessage(
		context.TODO(),
		0,
		"no-such-chat",
		OptionsSendMessage{},
	); !sent.OK {
		var e ErrChatNotFound
		if !errors.As(err, &e) {
			t.Errorf("should have failed with ErrChatNotFound, but got: %s", err)
		}
	} else {
		t.Errorf("should have failed to send message to a non-existent chat")
	}

	// TODO: ErrUserNotFound

	// TODO: ErrUserDeactivated

	// TODO: ErrBotKicked

	// TODO: ErrBotBlockedByUser

	// TODO: ErrBotCantSendToBots

	// TODO: ErrMessageNotModified

	// TODO: ErrGroupMigratedToSupergroup

	// TODO: ErrInvalidFileID

	// TODO: ErrConflictedLongPoll

	// TODO: ErrConflictedWebHook

	// TODO: ErrWrongParameterAction

	// ErrMessageEmpty
	if sent, err := client.SendMessage(
		context.TODO(),
		_chatID,
		"",
		OptionsSendMessage{},
	); !sent.OK {
		var e ErrMessageEmpty
		if !errors.As(err, &e) {
			t.Errorf("should have failed with ErrMessageEmpty but got: %s", err)
		}
	} else {
		t.Errorf("should have failed to send an empty message")
	}

	// ErrMessageTooLong
	longLongMessage := strings.Repeat("a", 4097)
	if sent, err := client.SendMessage(
		context.TODO(),
		_chatID,
		longLongMessage,
		OptionsSendMessage{},
	); !sent.OK {
		var e ErrMessageTooLong
		if !errors.As(err, &e) {
			t.Errorf("should have failed with ErrMessageTooLong but got: %s", err)
		}
	} else {
		t.Errorf("should have failed to send a long message")
	}

	// ErrMessageCantBeEdited
	// TODO: ErrMessageCantBeEdited

	// ErrTooManyRequests
	// TODO: ErrTooManyRequests

	// ErrJSONParseFailed
	// TODO: ErrJSONParseFailed

	// TODO: add more errors here

	// ErrUnclassified
	// TODO: ErrUnclassified
}
