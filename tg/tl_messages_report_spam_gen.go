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

// MessagesReportSpamRequest represents TL type `messages.reportSpam#cf1592db`.
// Report a new incoming chat for spam, if the peer settings¹ of the chat allow us to do that
//
// Links:
//  1) https://core.telegram.org/constructor/peerSettings
//
// See https://core.telegram.org/method/messages.reportSpam for reference.
type MessagesReportSpamRequest struct {
	// Peer to report
	Peer InputPeerClass
}

// MessagesReportSpamRequestTypeID is TL type id of MessagesReportSpamRequest.
const MessagesReportSpamRequestTypeID = 0xcf1592db

// String implements fmt.Stringer.
func (r *MessagesReportSpamRequest) String() string {
	if r == nil {
		return "MessagesReportSpamRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("MessagesReportSpamRequest")
	sb.WriteString("{\n")
	sb.WriteString("\tPeer: ")
	sb.WriteString(r.Peer.String())
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (r *MessagesReportSpamRequest) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode messages.reportSpam#cf1592db as nil")
	}
	b.PutID(MessagesReportSpamRequestTypeID)
	if r.Peer == nil {
		return fmt.Errorf("unable to encode messages.reportSpam#cf1592db: field peer is nil")
	}
	if err := r.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.reportSpam#cf1592db: field peer: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (r *MessagesReportSpamRequest) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode messages.reportSpam#cf1592db to nil")
	}
	if err := b.ConsumeID(MessagesReportSpamRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.reportSpam#cf1592db: %w", err)
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.reportSpam#cf1592db: field peer: %w", err)
		}
		r.Peer = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesReportSpamRequest.
var (
	_ bin.Encoder = &MessagesReportSpamRequest{}
	_ bin.Decoder = &MessagesReportSpamRequest{}
)

// MessagesReportSpam invokes method messages.reportSpam#cf1592db returning error if any.
// Report a new incoming chat for spam, if the peer settings¹ of the chat allow us to do that
//
// Links:
//  1) https://core.telegram.org/constructor/peerSettings
//
// See https://core.telegram.org/method/messages.reportSpam for reference.
func (c *Client) MessagesReportSpam(ctx context.Context, peer InputPeerClass) (bool, error) {
	var result BoolBox

	request := &MessagesReportSpamRequest{
		Peer: peer,
	}
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
