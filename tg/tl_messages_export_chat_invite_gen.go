// Code generated by gotdgen, DO NOT EDIT.

package tg

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

// MessagesExportChatInviteRequest represents TL type `messages.exportChatInvite#a455de90`.
// Export an invite link for a chat
//
// See https://core.telegram.org/method/messages.exportChatInvite for reference.
type MessagesExportChatInviteRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Legacy flag, reproducing legacy behavior of this method: if set, revokes all previous
	// links before creating a new one. Kept for bot API BC, should not be used by modern
	// clients.
	LegacyRevokePermanent bool
	// Whether admin confirmation is required before admitting each separate user into the
	// chat
	RequestNeeded bool
	// Chat
	Peer InputPeerClass
	// Expiration date
	//
	// Use SetExpireDate and GetExpireDate helpers.
	ExpireDate int
	// Maximum number of users that can join using this link
	//
	// Use SetUsageLimit and GetUsageLimit helpers.
	UsageLimit int
	// Description of the invite link, visible only to administrators
	//
	// Use SetTitle and GetTitle helpers.
	Title string
	// SubscriptionPricing field of MessagesExportChatInviteRequest.
	//
	// Use SetSubscriptionPricing and GetSubscriptionPricing helpers.
	SubscriptionPricing StarsSubscriptionPricing
}

// MessagesExportChatInviteRequestTypeID is TL type id of MessagesExportChatInviteRequest.
const MessagesExportChatInviteRequestTypeID = 0xa455de90

// Ensuring interfaces in compile-time for MessagesExportChatInviteRequest.
var (
	_ bin.Encoder     = &MessagesExportChatInviteRequest{}
	_ bin.Decoder     = &MessagesExportChatInviteRequest{}
	_ bin.BareEncoder = &MessagesExportChatInviteRequest{}
	_ bin.BareDecoder = &MessagesExportChatInviteRequest{}
)

func (e *MessagesExportChatInviteRequest) Zero() bool {
	if e == nil {
		return true
	}
	if !(e.Flags.Zero()) {
		return false
	}
	if !(e.LegacyRevokePermanent == false) {
		return false
	}
	if !(e.RequestNeeded == false) {
		return false
	}
	if !(e.Peer == nil) {
		return false
	}
	if !(e.ExpireDate == 0) {
		return false
	}
	if !(e.UsageLimit == 0) {
		return false
	}
	if !(e.Title == "") {
		return false
	}
	if !(e.SubscriptionPricing.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (e *MessagesExportChatInviteRequest) String() string {
	if e == nil {
		return "MessagesExportChatInviteRequest(nil)"
	}
	type Alias MessagesExportChatInviteRequest
	return fmt.Sprintf("MessagesExportChatInviteRequest%+v", Alias(*e))
}

// FillFrom fills MessagesExportChatInviteRequest from given interface.
func (e *MessagesExportChatInviteRequest) FillFrom(from interface {
	GetLegacyRevokePermanent() (value bool)
	GetRequestNeeded() (value bool)
	GetPeer() (value InputPeerClass)
	GetExpireDate() (value int, ok bool)
	GetUsageLimit() (value int, ok bool)
	GetTitle() (value string, ok bool)
	GetSubscriptionPricing() (value StarsSubscriptionPricing, ok bool)
}) {
	e.LegacyRevokePermanent = from.GetLegacyRevokePermanent()
	e.RequestNeeded = from.GetRequestNeeded()
	e.Peer = from.GetPeer()
	if val, ok := from.GetExpireDate(); ok {
		e.ExpireDate = val
	}

	if val, ok := from.GetUsageLimit(); ok {
		e.UsageLimit = val
	}

	if val, ok := from.GetTitle(); ok {
		e.Title = val
	}

	if val, ok := from.GetSubscriptionPricing(); ok {
		e.SubscriptionPricing = val
	}

}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesExportChatInviteRequest) TypeID() uint32 {
	return MessagesExportChatInviteRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesExportChatInviteRequest) TypeName() string {
	return "messages.exportChatInvite"
}

// TypeInfo returns info about TL type.
func (e *MessagesExportChatInviteRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.exportChatInvite",
		ID:   MessagesExportChatInviteRequestTypeID,
	}
	if e == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "LegacyRevokePermanent",
			SchemaName: "legacy_revoke_permanent",
			Null:       !e.Flags.Has(2),
		},
		{
			Name:       "RequestNeeded",
			SchemaName: "request_needed",
			Null:       !e.Flags.Has(3),
		},
		{
			Name:       "Peer",
			SchemaName: "peer",
		},
		{
			Name:       "ExpireDate",
			SchemaName: "expire_date",
			Null:       !e.Flags.Has(0),
		},
		{
			Name:       "UsageLimit",
			SchemaName: "usage_limit",
			Null:       !e.Flags.Has(1),
		},
		{
			Name:       "Title",
			SchemaName: "title",
			Null:       !e.Flags.Has(4),
		},
		{
			Name:       "SubscriptionPricing",
			SchemaName: "subscription_pricing",
			Null:       !e.Flags.Has(5),
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (e *MessagesExportChatInviteRequest) SetFlags() {
	if !(e.LegacyRevokePermanent == false) {
		e.Flags.Set(2)
	}
	if !(e.RequestNeeded == false) {
		e.Flags.Set(3)
	}
	if !(e.ExpireDate == 0) {
		e.Flags.Set(0)
	}
	if !(e.UsageLimit == 0) {
		e.Flags.Set(1)
	}
	if !(e.Title == "") {
		e.Flags.Set(4)
	}
	if !(e.SubscriptionPricing.Zero()) {
		e.Flags.Set(5)
	}
}

// Encode implements bin.Encoder.
func (e *MessagesExportChatInviteRequest) Encode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode messages.exportChatInvite#a455de90 as nil")
	}
	b.PutID(MessagesExportChatInviteRequestTypeID)
	return e.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (e *MessagesExportChatInviteRequest) EncodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode messages.exportChatInvite#a455de90 as nil")
	}
	e.SetFlags()
	if err := e.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.exportChatInvite#a455de90: field flags: %w", err)
	}
	if e.Peer == nil {
		return fmt.Errorf("unable to encode messages.exportChatInvite#a455de90: field peer is nil")
	}
	if err := e.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.exportChatInvite#a455de90: field peer: %w", err)
	}
	if e.Flags.Has(0) {
		b.PutInt(e.ExpireDate)
	}
	if e.Flags.Has(1) {
		b.PutInt(e.UsageLimit)
	}
	if e.Flags.Has(4) {
		b.PutString(e.Title)
	}
	if e.Flags.Has(5) {
		if err := e.SubscriptionPricing.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.exportChatInvite#a455de90: field subscription_pricing: %w", err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (e *MessagesExportChatInviteRequest) Decode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode messages.exportChatInvite#a455de90 to nil")
	}
	if err := b.ConsumeID(MessagesExportChatInviteRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.exportChatInvite#a455de90: %w", err)
	}
	return e.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (e *MessagesExportChatInviteRequest) DecodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode messages.exportChatInvite#a455de90 to nil")
	}
	{
		if err := e.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.exportChatInvite#a455de90: field flags: %w", err)
		}
	}
	e.LegacyRevokePermanent = e.Flags.Has(2)
	e.RequestNeeded = e.Flags.Has(3)
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.exportChatInvite#a455de90: field peer: %w", err)
		}
		e.Peer = value
	}
	if e.Flags.Has(0) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.exportChatInvite#a455de90: field expire_date: %w", err)
		}
		e.ExpireDate = value
	}
	if e.Flags.Has(1) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.exportChatInvite#a455de90: field usage_limit: %w", err)
		}
		e.UsageLimit = value
	}
	if e.Flags.Has(4) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode messages.exportChatInvite#a455de90: field title: %w", err)
		}
		e.Title = value
	}
	if e.Flags.Has(5) {
		if err := e.SubscriptionPricing.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.exportChatInvite#a455de90: field subscription_pricing: %w", err)
		}
	}
	return nil
}

// SetLegacyRevokePermanent sets value of LegacyRevokePermanent conditional field.
func (e *MessagesExportChatInviteRequest) SetLegacyRevokePermanent(value bool) {
	if value {
		e.Flags.Set(2)
		e.LegacyRevokePermanent = true
	} else {
		e.Flags.Unset(2)
		e.LegacyRevokePermanent = false
	}
}

// GetLegacyRevokePermanent returns value of LegacyRevokePermanent conditional field.
func (e *MessagesExportChatInviteRequest) GetLegacyRevokePermanent() (value bool) {
	if e == nil {
		return
	}
	return e.Flags.Has(2)
}

// SetRequestNeeded sets value of RequestNeeded conditional field.
func (e *MessagesExportChatInviteRequest) SetRequestNeeded(value bool) {
	if value {
		e.Flags.Set(3)
		e.RequestNeeded = true
	} else {
		e.Flags.Unset(3)
		e.RequestNeeded = false
	}
}

// GetRequestNeeded returns value of RequestNeeded conditional field.
func (e *MessagesExportChatInviteRequest) GetRequestNeeded() (value bool) {
	if e == nil {
		return
	}
	return e.Flags.Has(3)
}

// GetPeer returns value of Peer field.
func (e *MessagesExportChatInviteRequest) GetPeer() (value InputPeerClass) {
	if e == nil {
		return
	}
	return e.Peer
}

// SetExpireDate sets value of ExpireDate conditional field.
func (e *MessagesExportChatInviteRequest) SetExpireDate(value int) {
	e.Flags.Set(0)
	e.ExpireDate = value
}

// GetExpireDate returns value of ExpireDate conditional field and
// boolean which is true if field was set.
func (e *MessagesExportChatInviteRequest) GetExpireDate() (value int, ok bool) {
	if e == nil {
		return
	}
	if !e.Flags.Has(0) {
		return value, false
	}
	return e.ExpireDate, true
}

// SetUsageLimit sets value of UsageLimit conditional field.
func (e *MessagesExportChatInviteRequest) SetUsageLimit(value int) {
	e.Flags.Set(1)
	e.UsageLimit = value
}

// GetUsageLimit returns value of UsageLimit conditional field and
// boolean which is true if field was set.
func (e *MessagesExportChatInviteRequest) GetUsageLimit() (value int, ok bool) {
	if e == nil {
		return
	}
	if !e.Flags.Has(1) {
		return value, false
	}
	return e.UsageLimit, true
}

// SetTitle sets value of Title conditional field.
func (e *MessagesExportChatInviteRequest) SetTitle(value string) {
	e.Flags.Set(4)
	e.Title = value
}

// GetTitle returns value of Title conditional field and
// boolean which is true if field was set.
func (e *MessagesExportChatInviteRequest) GetTitle() (value string, ok bool) {
	if e == nil {
		return
	}
	if !e.Flags.Has(4) {
		return value, false
	}
	return e.Title, true
}

// SetSubscriptionPricing sets value of SubscriptionPricing conditional field.
func (e *MessagesExportChatInviteRequest) SetSubscriptionPricing(value StarsSubscriptionPricing) {
	e.Flags.Set(5)
	e.SubscriptionPricing = value
}

// GetSubscriptionPricing returns value of SubscriptionPricing conditional field and
// boolean which is true if field was set.
func (e *MessagesExportChatInviteRequest) GetSubscriptionPricing() (value StarsSubscriptionPricing, ok bool) {
	if e == nil {
		return
	}
	if !e.Flags.Has(5) {
		return value, false
	}
	return e.SubscriptionPricing, true
}

// MessagesExportChatInvite invokes method messages.exportChatInvite#a455de90 returning error if any.
// Export an invite link for a chat
//
// Possible errors:
//
//	400 CHANNEL_INVALID: The provided channel is invalid.
//	400 CHANNEL_PRIVATE: You haven't joined this channel/supergroup.
//	400 CHAT_ADMIN_REQUIRED: You must be an admin in this chat to do this.
//	400 CHAT_ID_INVALID: The provided chat id is invalid.
//	403 CHAT_WRITE_FORBIDDEN: You can't write in this chat.
//	400 EXPIRE_DATE_INVALID: The specified expiration date is invalid.
//	400 MSG_ID_INVALID: Invalid message ID provided.
//	400 PEER_ID_INVALID: The provided peer id is invalid.
//	400 USAGE_LIMIT_INVALID: The specified usage limit is invalid.
//
// See https://core.telegram.org/method/messages.exportChatInvite for reference.
// Can be used by bots.
func (c *Client) MessagesExportChatInvite(ctx context.Context, request *MessagesExportChatInviteRequest) (ExportedChatInviteClass, error) {
	var result ExportedChatInviteBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.ExportedChatInvite, nil
}
