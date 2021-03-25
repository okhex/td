package inline

import (
	"github.com/gotd/td/telegram/message/entity"
	"github.com/gotd/td/telegram/message/markup"
	"github.com/gotd/td/telegram/message/styling"
	"github.com/gotd/td/tg"
)

// MessageTextBuilder is a builder of inline result text message.
type MessageTextBuilder struct {
	message *tg.InputBotInlineMessageText
	option  styling.StyledTextOption
	options []styling.StyledTextOption
}

func (b *MessageTextBuilder) apply() (tg.InputBotInlineMessageClass, error) {
	tb := entity.Builder{}
	if err := styling.Perform(&tb, b.option, b.options...); err != nil {
		return nil, err
	}
	msg, entities := tb.Complete()
	r := *b.message

	r.Message = msg
	r.Entities = entities
	return &r, nil
}

// MessageText creates new message text option builder.
func MessageText(msg string) *MessageTextBuilder {
	return MessageStyledText(styling.Plain(msg))
}

// MessageStyledText creates new message text option builder.
func MessageStyledText(text styling.StyledTextOption, texts ...styling.StyledTextOption) *MessageTextBuilder {
	return &MessageTextBuilder{
		message: &tg.InputBotInlineMessageText{},
		option:  text,
		options: texts,
	}
}

// NoWebpage sets flag to disable generation of the webpage preview.
func (b *MessageTextBuilder) NoWebpage() *MessageTextBuilder {
	b.message.NoWebpage = true
	return b
}

// Markup sets reply markup for sending bot buttons.
// NB: markup will not be used, if you send multiple media attachments.
func (b *MessageTextBuilder) Markup(m tg.ReplyMarkupClass) *MessageTextBuilder {
	b.message.ReplyMarkup = m
	return b
}

// Row sets single row keyboard markup  for sending bot buttons.
// NB: markup will not be used, if you send multiple media attachments.
func (b *MessageTextBuilder) Row(button tg.KeyboardButtonClass, buttons ...tg.KeyboardButtonClass) *MessageTextBuilder {
	return b.Markup(markup.InlineRow(button, buttons...))
}
