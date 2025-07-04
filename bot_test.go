// bot_test.go
//
// created on: 2023.11.10.

package telegrambot

import (
	"errors"
	"fmt"
	"log/slog"
	"os"
	"strings"
	"testing"
	"time"
)

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

	if deleted := client.DeleteWebhook(true); !deleted.Ok {
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
	_chatID := os.Getenv("CHAT_ID") // NOTE: `chat_id` of a group chat
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
	if me := client.GetMe(); !me.Ok {
		t.Errorf("failed to get me: %s", *me.Description)
	} else {
		////////////////////////////////
		// (webhook)
		//
		// SetWebhook
		if webhook := client.SetWebhook("testdomain.com", 8443, OptionsSetWebhook{}); !webhook.Ok {
			t.Errorf("failed to set webhook: %s", *webhook.Description)
		} else {
			// GetWebhookInfo
			if webhook := client.GetWebhookInfo(); !webhook.Ok {
				t.Errorf("failed to get webhook info: %s", *webhook.Description)
			}

			// DeleteWebhook
			if deleted := client.DeleteWebhook(false); !deleted.Ok {
				t.Errorf("failed to delete webhook: %s", *webhook.Description)
			}

			////////////////////////////////
			// (general methods)
			//
			// GetUpdates
			if updates := client.GetUpdates(OptionsGetUpdates{}); !updates.Ok {
				t.Errorf("failed to get updates: %s", *updates.Description)
			}
			// TODO: LogOut
			// TODO: Close
			// SendMessage
			if sent := client.SendMessage(_chatID, "test message", OptionsSendMessage{}); !sent.Ok {
				t.Errorf("failed to send message: %s", *sent.Description)
			} else {
				// EditMessageText
				if edited := client.EditMessageText("edited message", OptionsEditMessageText{}.
					SetIDs(_chatID, sent.Result.MessageID)); !edited.Ok {
					t.Errorf("failed to edit message text: %s", *edited.Description)
				}
				// CopyMessage
				if copied := client.CopyMessage(_chatID, _chatID, sent.Result.MessageID, OptionsCopyMessage{}); !copied.Ok {
					t.Errorf("failed to copy message: %s", *copied.Description)
				}
				// ForwardMessage
				if forwarded := client.ForwardMessage(_chatID, _chatID, sent.Result.MessageID, OptionsForwardMessage{}); !forwarded.Ok {
					t.Errorf("failed to forward message: %s", *forwarded.Description)
				}
			}
			// TODO: CopyMessages
			// TODO: ForwardMessages
			// SendPhoto
			if photo := client.SendPhoto(_chatID, NewInputFileFromFilepath("./samples/_files/gopher.png"), OptionsSendPhoto{}); !photo.Ok {
				t.Errorf("failed to send photo: %s", *photo.Description)
			} else {
				// EditMessageCaption
				if caption := client.EditMessageCaption(OptionsEditMessageCaption{}.
					SetIDs(_chatID, photo.Result.MessageID).
					SetCaption("edited caption")); !caption.Ok {
					t.Errorf("failed to edit message caption: %s", *caption.Description)
				}
			}
			// TODO: SendAudio
			// TODO: SendDocument
			if doc := client.SendDocument(_chatID, NewInputFileFromFilepath("./samples/_files/gopher.png"), OptionsSendDocument{}); !doc.Ok {
				t.Errorf("failed to send document: %s", *doc.Description)
			} else {
				// GetFile
				if file := client.GetFile(doc.Result.Document.FileID); !file.Ok {
					t.Errorf("failed to get file: %s", *file.Description)
				}
				// DeleteMessage
				if deleted := client.DeleteMessage(_chatID, doc.Result.MessageID); !deleted.Ok {
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
			if location := client.SendLocation(_chatID, 37.5665, 126.9780, OptionsSendLocation{}); !location.Ok {
				t.Errorf("failed to send location: %s", *location.Description)
			}
			// TODO: SendVenue
			// SendContact
			if contact := client.SendContact(_chatID, "911", "Nine-One-One", OptionsSendContact{}); !contact.Ok {
				t.Errorf("failed to send contact: %s", *contact.Description)
			}
			// SendPoll
			if poll := client.SendPoll(_chatID, "The earth is...?", []InputPollOption{
				{Text: "flat"},
				{Text: "round"},
				{Text: "nothing"},
			}, OptionsSendPoll{}); !poll.Ok {
				t.Errorf("failed to send poll: %s", *poll.Description)
			} else {
				// StopPoll
				if stopped := client.StopPoll(_chatID, poll.Result.MessageID, OptionsStopPoll{}); !stopped.Ok {
					t.Errorf("failed to stop poll: %s", *stopped.Description)
				}
			}
			// TODO: SendChecklist
			// SendDice
			if dice := client.SendDice(_chatID, OptionsSendDice{}); !dice.Ok {
				t.Errorf("failed to send dice: %s", *dice.Description)
			}
			// SendChatAction
			if action := client.SendChatAction(_chatID, ChatActionTyping, OptionsSendChatAction{}); !action.Ok {
				t.Errorf("failed to send chat action: %s", *action.Description)
			}
			// TODO: GetUserProfilePhotos
			// TODO: ApproveChatJoinRequest
			// TODO: DeclineChatJoinRequest
			// TODO: GetMyCommands
			// GetMyName
			if name := client.GetMyName(OptionsGetMyName{}); !name.Ok {
				t.Errorf("failed to get my name: %s", *name.Description)
			} else {
				newName := "telegram api test bot"

				if name.Result.Name != newName {
					// SetMyName
					if name := client.SetMyName(newName, OptionsSetMyName{}); !name.Ok {
						t.Errorf("failed to set my name: %s", *name.Description)
					}
				}
			}
			// SetMyDescription
			if desc := client.SetMyDescription(OptionsSetMyDescription{}.
				SetDescription("A bot for testing library: telegram-bot-go")); !desc.Ok {
				t.Errorf("failed to set my description: %s", *desc.Description)
			}
			// GetMyDescription
			if desc := client.GetMyDescription(OptionsGetMyDescription{}); !desc.Ok {
				t.Errorf("failed to get my description: %s", *desc.Description)
			}
			// SetMyShortDescription
			if desc := client.SetMyShortDescription(OptionsSetMyShortDescription{}.
				SetShortDescription("telegram-bot-go")); !desc.Ok {
				t.Errorf("failed to set my short description: %s", *desc.Description)
			}
			// GetMyShortDescription
			if desc := client.GetMyShortDescription(OptionsGetMyShortDescription{}); !desc.Ok {
				t.Errorf("failed to get my short description: %s", *desc.Description)
			}
			// TODO: GetUserChatBoosts
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
			// TODO: ConvertGiftToStars
			// TODO: UpgradeGift
			// TODO: TransferGift
			// TODO: PostStory
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
			if chat := client.GetChat(_chatID); !chat.Ok {
				t.Errorf("failed to get chat: %s", *chat.Description)
			}
			// GetChatAdministrators
			if admins := client.GetChatAdministrators(_chatID); !admins.Ok {
				t.Errorf("failed to get chat administrators: %s", *admins.Description)
			}
			// GetChatMemberCount
			if count := client.GetChatMemberCount(_chatID); !count.Ok {
				t.Errorf("failed to get chat member count: %s", *count.Description)
			}
			// TODO: GetChatMember
			// TODO: CreateChat
			// TODO: SetChatTitle
			// SetChatDescription
			if desc := client.SetChatDescription(_chatID, fmt.Sprintf("[telegram-bot-go] chat_id: %s (last update: %d)", _chatID, time.Now().Unix())); !desc.Ok {
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
			if balance := client.GetMyStarBalance(); !balance.Ok {
				t.Errorf("failed to get my star balance: %s", *balance.Description)
			}
			// TODO: GetStarTransactions
			// TODO: RefundStarPayment
			// TODO: EditUserStarSubscription

			////////////////////////////////
			// (forum)
			//
			// TODO: CreateForumTopic
			// TODO: EditForumTopic
			// TODO: CloseForumTopic
			// TODO: ReopenForumTopic
			// TODO: DeleteForumTopic
			// TODO: UnpinAllForumTopicMessages
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
	if res := unauthClient.SendMessage(_chatID, "unauthorized", OptionsSendMessage{}); !res.Ok {
		var e ErrUnauthorized
		if !errors.As(res.Error, &e) {
			t.Errorf("should have failed with ErrUnauthorized, but got: %s", res.Error)
		}
	} else {
		t.Errorf("should have failed to send unauthorized request")
	}

	// ErrChatNotFound
	if res := client.SendMessage(0, "no-such-chat", OptionsSendMessage{}); !res.Ok {
		var e ErrChatNotFound
		if !errors.As(res.Error, &e) {
			t.Errorf("should have failed with ErrChatNotFound, but got: %s", res.Error)
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
	if res := client.SendMessage(_chatID, "", OptionsSendMessage{}); !res.Ok {
		var e ErrMessageEmpty
		if !errors.As(res.Error, &e) {
			t.Errorf("should have failed with ErrMessageEmpty but got: %s", res.Error)
		}
	} else {
		t.Errorf("should have failed to send an empty message")
	}

	// ErrMessageTooLong
	longLongMessage := strings.Repeat("a", 4097)
	if res := client.SendMessage(_chatID, longLongMessage, OptionsSendMessage{}); !res.Ok {
		var e ErrMessageTooLong
		if !errors.As(res.Error, &e) {
			t.Errorf("should have failed with ErrMessageTooLong but got: %s", res.Error)
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
