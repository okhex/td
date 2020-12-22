// Code generated by gotdgen, DO NOT EDIT.

package tg

import (
	"context"
	"fmt"
	"strings"

	"github.com/gotd/td/bin"
)

// No-op definition for keeping imports.
var _ = bin.Buffer{}
var _ = context.Background()
var _ = fmt.Stringer(nil)
var _ = strings.Builder{}

// MessagesAddChatUserRequest represents TL type `messages.addChatUser#f9a0aa09`.
// Adds a user to a chat and sends a service message on it.
//
// See https://core.telegram.org/method/messages.addChatUser for reference.
type MessagesAddChatUserRequest struct {
	// Chat ID
	ChatID int
	// User ID to be added
	UserID InputUserClass
	// Number of last messages to be forwarded
	FwdLimit int
}

// MessagesAddChatUserRequestTypeID is TL type id of MessagesAddChatUserRequest.
const MessagesAddChatUserRequestTypeID = 0xf9a0aa09

// String implements fmt.Stringer.
func (a *MessagesAddChatUserRequest) String() string {
	if a == nil {
		return "MessagesAddChatUserRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("MessagesAddChatUserRequest")
	sb.WriteString("{\n")
	sb.WriteString("\tChatID: ")
	sb.WriteString(fmt.Sprint(a.ChatID))
	sb.WriteString(",\n")
	sb.WriteString("\tUserID: ")
	sb.WriteString(a.UserID.String())
	sb.WriteString(",\n")
	sb.WriteString("\tFwdLimit: ")
	sb.WriteString(fmt.Sprint(a.FwdLimit))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (a *MessagesAddChatUserRequest) Encode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode messages.addChatUser#f9a0aa09 as nil")
	}
	b.PutID(MessagesAddChatUserRequestTypeID)
	b.PutInt(a.ChatID)
	if a.UserID == nil {
		return fmt.Errorf("unable to encode messages.addChatUser#f9a0aa09: field user_id is nil")
	}
	if err := a.UserID.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.addChatUser#f9a0aa09: field user_id: %w", err)
	}
	b.PutInt(a.FwdLimit)
	return nil
}

// Decode implements bin.Decoder.
func (a *MessagesAddChatUserRequest) Decode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode messages.addChatUser#f9a0aa09 to nil")
	}
	if err := b.ConsumeID(MessagesAddChatUserRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.addChatUser#f9a0aa09: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.addChatUser#f9a0aa09: field chat_id: %w", err)
		}
		a.ChatID = value
	}
	{
		value, err := DecodeInputUser(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.addChatUser#f9a0aa09: field user_id: %w", err)
		}
		a.UserID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.addChatUser#f9a0aa09: field fwd_limit: %w", err)
		}
		a.FwdLimit = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesAddChatUserRequest.
var (
	_ bin.Encoder = &MessagesAddChatUserRequest{}
	_ bin.Decoder = &MessagesAddChatUserRequest{}
)

// MessagesAddChatUser invokes method messages.addChatUser#f9a0aa09 returning error if any.
// Adds a user to a chat and sends a service message on it.
//
// See https://core.telegram.org/method/messages.addChatUser for reference.
func (c *Client) MessagesAddChatUser(ctx context.Context, request *MessagesAddChatUserRequest) (UpdatesClass, error) {
	var result UpdatesBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}
