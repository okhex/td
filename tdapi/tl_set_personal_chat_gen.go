// Code generated by gotdgen, DO NOT EDIT.

package tdapi

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"

	"go.uber.org/multierr"

	"github.com/gotd/td/bin"
	"github.com/gotd/td/tdjson"
	"github.com/gotd/td/tdp"
	"github.com/gotd/td/tgerr"
)

// No-op definition for keeping imports.
var (
	_ = bin.Buffer{}
	_ = context.Background()
	_ = fmt.Stringer(nil)
	_ = strings.Builder{}
	_ = errors.Is
	_ = multierr.AppendInto
	_ = sort.Ints
	_ = tdp.Format
	_ = tgerr.Error{}
	_ = tdjson.Encoder{}
)

// SetPersonalChatRequest represents TL type `setPersonalChat#c04babb4`.
type SetPersonalChatRequest struct {
	// Identifier of the new personal chat; pass 0 to remove the chat. Use
	// getSuitablePersonalChats to get suitable chats
	ChatID int64
}

// SetPersonalChatRequestTypeID is TL type id of SetPersonalChatRequest.
const SetPersonalChatRequestTypeID = 0xc04babb4

// Ensuring interfaces in compile-time for SetPersonalChatRequest.
var (
	_ bin.Encoder     = &SetPersonalChatRequest{}
	_ bin.Decoder     = &SetPersonalChatRequest{}
	_ bin.BareEncoder = &SetPersonalChatRequest{}
	_ bin.BareDecoder = &SetPersonalChatRequest{}
)

func (s *SetPersonalChatRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.ChatID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *SetPersonalChatRequest) String() string {
	if s == nil {
		return "SetPersonalChatRequest(nil)"
	}
	type Alias SetPersonalChatRequest
	return fmt.Sprintf("SetPersonalChatRequest%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*SetPersonalChatRequest) TypeID() uint32 {
	return SetPersonalChatRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*SetPersonalChatRequest) TypeName() string {
	return "setPersonalChat"
}

// TypeInfo returns info about TL type.
func (s *SetPersonalChatRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "setPersonalChat",
		ID:   SetPersonalChatRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ChatID",
			SchemaName: "chat_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *SetPersonalChatRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode setPersonalChat#c04babb4 as nil")
	}
	b.PutID(SetPersonalChatRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *SetPersonalChatRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode setPersonalChat#c04babb4 as nil")
	}
	b.PutInt53(s.ChatID)
	return nil
}

// Decode implements bin.Decoder.
func (s *SetPersonalChatRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode setPersonalChat#c04babb4 to nil")
	}
	if err := b.ConsumeID(SetPersonalChatRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode setPersonalChat#c04babb4: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *SetPersonalChatRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode setPersonalChat#c04babb4 to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode setPersonalChat#c04babb4: field chat_id: %w", err)
		}
		s.ChatID = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (s *SetPersonalChatRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode setPersonalChat#c04babb4 as nil")
	}
	b.ObjStart()
	b.PutID("setPersonalChat")
	b.Comma()
	b.FieldStart("chat_id")
	b.PutInt53(s.ChatID)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (s *SetPersonalChatRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode setPersonalChat#c04babb4 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("setPersonalChat"); err != nil {
				return fmt.Errorf("unable to decode setPersonalChat#c04babb4: %w", err)
			}
		case "chat_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode setPersonalChat#c04babb4: field chat_id: %w", err)
			}
			s.ChatID = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetChatID returns value of ChatID field.
func (s *SetPersonalChatRequest) GetChatID() (value int64) {
	if s == nil {
		return
	}
	return s.ChatID
}

// SetPersonalChat invokes method setPersonalChat#c04babb4 returning error if any.
func (c *Client) SetPersonalChat(ctx context.Context, chatid int64) error {
	var ok Ok

	request := &SetPersonalChatRequest{
		ChatID: chatid,
	}
	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
