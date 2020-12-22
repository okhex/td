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

// PhoneConfirmCallRequest represents TL type `phone.confirmCall#2efe1722`.
// Complete phone call E2E encryption key exchange »¹
//
// Links:
//  1) https://core.telegram.org/api/end-to-end/voice-calls
//
// See https://core.telegram.org/method/phone.confirmCall for reference.
type PhoneConfirmCallRequest struct {
	// The phone call
	Peer InputPhoneCall
	// Parameter for E2E encryption key exchange »¹
	//
	// Links:
	//  1) https://core.telegram.org/api/end-to-end/voice-calls
	GA []byte
	// Key fingerprint
	KeyFingerprint int64
	// Phone call settings
	Protocol PhoneCallProtocol
}

// PhoneConfirmCallRequestTypeID is TL type id of PhoneConfirmCallRequest.
const PhoneConfirmCallRequestTypeID = 0x2efe1722

// String implements fmt.Stringer.
func (c *PhoneConfirmCallRequest) String() string {
	if c == nil {
		return "PhoneConfirmCallRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("PhoneConfirmCallRequest")
	sb.WriteString("{\n")
	sb.WriteString("\tPeer: ")
	sb.WriteString(c.Peer.String())
	sb.WriteString(",\n")
	sb.WriteString("\tGA: ")
	sb.WriteString(fmt.Sprint(c.GA))
	sb.WriteString(",\n")
	sb.WriteString("\tKeyFingerprint: ")
	sb.WriteString(fmt.Sprint(c.KeyFingerprint))
	sb.WriteString(",\n")
	sb.WriteString("\tProtocol: ")
	sb.WriteString(c.Protocol.String())
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (c *PhoneConfirmCallRequest) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode phone.confirmCall#2efe1722 as nil")
	}
	b.PutID(PhoneConfirmCallRequestTypeID)
	if err := c.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode phone.confirmCall#2efe1722: field peer: %w", err)
	}
	b.PutBytes(c.GA)
	b.PutLong(c.KeyFingerprint)
	if err := c.Protocol.Encode(b); err != nil {
		return fmt.Errorf("unable to encode phone.confirmCall#2efe1722: field protocol: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (c *PhoneConfirmCallRequest) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode phone.confirmCall#2efe1722 to nil")
	}
	if err := b.ConsumeID(PhoneConfirmCallRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode phone.confirmCall#2efe1722: %w", err)
	}
	{
		if err := c.Peer.Decode(b); err != nil {
			return fmt.Errorf("unable to decode phone.confirmCall#2efe1722: field peer: %w", err)
		}
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode phone.confirmCall#2efe1722: field g_a: %w", err)
		}
		c.GA = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode phone.confirmCall#2efe1722: field key_fingerprint: %w", err)
		}
		c.KeyFingerprint = value
	}
	{
		if err := c.Protocol.Decode(b); err != nil {
			return fmt.Errorf("unable to decode phone.confirmCall#2efe1722: field protocol: %w", err)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for PhoneConfirmCallRequest.
var (
	_ bin.Encoder = &PhoneConfirmCallRequest{}
	_ bin.Decoder = &PhoneConfirmCallRequest{}
)

// PhoneConfirmCall invokes method phone.confirmCall#2efe1722 returning error if any.
// Complete phone call E2E encryption key exchange »¹
//
// Links:
//  1) https://core.telegram.org/api/end-to-end/voice-calls
//
// See https://core.telegram.org/method/phone.confirmCall for reference.
func (c *Client) PhoneConfirmCall(ctx context.Context, request *PhoneConfirmCallRequest) (*PhonePhoneCall, error) {
	var result PhonePhoneCall

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
