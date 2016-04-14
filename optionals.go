package gotelebot

import (
	"net/url"
	"strconv"

	"github.com/likoo/gotelebot/types"
)

type Optional interface {
	AppendPayload(payload *url.Values)
}

// Optional parameters for SendMessage method
type SendMessageOptional struct {
	// Send Markdown, if you want Telegram apps to show bold, italic and inline URLs in your bot's message.
	ParseMode *string
	// Disables link previews for links in this message
	DisableWebPagePreview *bool
	DisableNotification   *bool
	ReplyToMessageId      *int
	ReplyMarkup           *types.ReplyMarkup
}

func (opt *SendMessageOptional) AppendPayload(payload *url.Values) {
	if opt.ParseMode != nil {
		payload.Add("parse_mode", *opt.ParseMode)
	}
	if opt.DisableWebPagePreview != nil {
		payload.Add("disable_web_page_preview", strconv.FormatBool(*opt.DisableWebPagePreview))
	}
	if opt.DisableNotification != nil {
		payload.Add("disable_notification", strconv.FormatBool(*opt.DisableNotification))
	}
	if opt.ReplyToMessageId != nil {
		payload.Add("reply_to_message_id", strconv.Itoa(*opt.ReplyToMessageId))
	}
	if opt.ReplyMarkup != nil {
		jsonStr, _ := (*opt.ReplyMarkup).ToJson()
		payload.Add("reply_markup", jsonStr)
	}
}

// Optional parameters for SendPhoto method
type SendPhotoOptional struct {
	// Photo caption
	Caption          *string
	ReplyToMessageId *int
	ReplyMarkup      *types.ReplyMarkup
}

func (opt *SendPhotoOptional) AppendPayload(payload *url.Values) {
	if opt.Caption != nil {
		payload.Add("caption", *opt.Caption)
	}
	if opt.ReplyToMessageId != nil {
		payload.Add("reply_to_message_id", strconv.Itoa(*opt.ReplyToMessageId))
	}
	if opt.ReplyMarkup != nil {
		jsonStr, _ := (*opt.ReplyMarkup).ToJson()
		payload.Add("reply_markup", jsonStr)
	}
}

// Optional parameters for SendAudio method
type SendAudioOptional struct {
	// Duration of the audio in seconds
	Duration         *int
	Performer        *string
	Title            *string
	ReplyToMessageId *int
	ReplyMarkup      *types.ReplyMarkup
}

func (opt *SendAudioOptional) AppendPayload(payload *url.Values) {
	if opt.Duration != nil {
		payload.Add("duration", strconv.Itoa(*opt.Duration))
	}
	if opt.Performer != nil {
		payload.Add("performer", *opt.Performer)
	}
	if opt.Title != nil {
		payload.Add("title", *opt.Title)
	}
	if opt.ReplyToMessageId != nil {
		payload.Add("reply_to_message_id", strconv.Itoa(*opt.ReplyToMessageId))
	}
	if opt.ReplyMarkup != nil {
		jsonStr, _ := (*opt.ReplyMarkup).ToJson()
		payload.Add("reply_markup", jsonStr)
	}
}

// Optional parameters for SendDocument method
type SendDocumentOptional struct {
	ReplyToMessageId *int
	ReplyMarkup      *types.ReplyMarkup
}

func (opt *SendDocumentOptional) AppendPayload(payload *url.Values) {
	if opt.ReplyToMessageId != nil {
		payload.Add("reply_to_message_id", strconv.Itoa(*opt.ReplyToMessageId))
	}
	if opt.ReplyMarkup != nil {
		jsonStr, _ := (*opt.ReplyMarkup).ToJson()
		payload.Add("reply_markup", jsonStr)
	}
}

// Optional parameters for SendSticker method
type SendStickerOptional struct {
	ReplyToMessageId *int
	ReplyMarkup      *types.ReplyMarkup
}

func (opt *SendStickerOptional) AppendPayload(payload *url.Values) {
	if opt.ReplyToMessageId != nil {
		payload.Add("reply_to_message_id", strconv.Itoa(*opt.ReplyToMessageId))
	}
	if opt.ReplyMarkup != nil {
		jsonStr, _ := (*opt.ReplyMarkup).ToJson()
		payload.Add("reply_markup", jsonStr)
	}
}

// Optional parameters for SendVideo method
type SendVideoOptional struct {
	Duration         *int
	Caption          *string
	ReplyToMessageId *int
	ReplyMarkup      *types.ReplyMarkup
}

func (opt *SendVideoOptional) AppendPayload(payload *url.Values) {
	if opt.Duration != nil {
		payload.Add("duration", strconv.Itoa(*opt.Duration))
	}
	if opt.Caption != nil {
		payload.Add("caption", *opt.Caption)
	}
	if opt.ReplyToMessageId != nil {
		payload.Add("reply_to_message_id", strconv.Itoa(*opt.ReplyToMessageId))
	}
	if opt.ReplyMarkup != nil {
		jsonStr, _ := (*opt.ReplyMarkup).ToJson()
		payload.Add("reply_markup", jsonStr)
	}
}

// Optional parameters for SendVoice method
type SendVoiceOptional struct {
	Duration         *int
	ReplyToMessageId *int
	ReplyMarkup      *types.ReplyMarkup
}

func (opt *SendVoiceOptional) AppendPayload(payload *url.Values) {
	if opt.Duration != nil {
		payload.Add("duration", strconv.Itoa(*opt.Duration))
	}
	if opt.ReplyToMessageId != nil {
		payload.Add("reply_to_message_id", strconv.Itoa(*opt.ReplyToMessageId))
	}
	if opt.ReplyMarkup != nil {
		jsonStr, _ := (*opt.ReplyMarkup).ToJson()
		payload.Add("reply_markup", jsonStr)
	}
}

// Optional parameters for SendLocation method
type SendLocationOptional struct {
	ReplyToMessageId *int
	ReplyMarkup      *types.ReplyMarkup
}

func (opt *SendLocationOptional) AppendPayload(payload *url.Values) {
	if opt.ReplyToMessageId != nil {
		payload.Add("reply_to_message_id", strconv.Itoa(*opt.ReplyToMessageId))
	}
	if opt.ReplyMarkup != nil {
		jsonStr, _ := (*opt.ReplyMarkup).ToJson()
		payload.Add("reply_markup", jsonStr)
	}
}

type GetUserProfilePhotosOptional struct {
	Offset *int
	Limit  *int
}

func (opt *GetUserProfilePhotosOptional) AppendPayload(payload *url.Values) {
	if opt.Offset != nil {
		payload.Add("offset", strconv.Itoa(*opt.Offset))
	}
	if opt.Limit != nil {
		payload.Add("limit", strconv.Itoa(*opt.Limit))
	}
}

type AnswerInlineQueryOptional struct {
	CacheTime  *int
	IsPersonal *bool
	NextOffset *string
}

func (opt *AnswerInlineQueryOptional) AppendPayload(payload *url.Values) {
	if opt.CacheTime != nil {
		payload.Add("cache_time", strconv.Itoa(*opt.CacheTime))
	}
	if opt.IsPersonal != nil {
		if *opt.IsPersonal {
			payload.Add("is_personal", "True")
		} else {
			payload.Add("is_personal", "False")
		}
	}
	if opt.NextOffset != nil {
		payload.Add("next_offset", *opt.NextOffset)
	}
}
