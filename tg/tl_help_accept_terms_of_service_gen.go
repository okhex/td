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

// HelpAcceptTermsOfServiceRequest represents TL type `help.acceptTermsOfService#ee72f79a`.
// Accept the new terms of service
//
// See https://core.telegram.org/method/help.acceptTermsOfService for reference.
type HelpAcceptTermsOfServiceRequest struct {
	// ID of terms of service
	ID DataJSON
}

// HelpAcceptTermsOfServiceRequestTypeID is TL type id of HelpAcceptTermsOfServiceRequest.
const HelpAcceptTermsOfServiceRequestTypeID = 0xee72f79a

// String implements fmt.Stringer.
func (a *HelpAcceptTermsOfServiceRequest) String() string {
	if a == nil {
		return "HelpAcceptTermsOfServiceRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("HelpAcceptTermsOfServiceRequest")
	sb.WriteString("{\n")
	sb.WriteString("\tID: ")
	sb.WriteString(a.ID.String())
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (a *HelpAcceptTermsOfServiceRequest) Encode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode help.acceptTermsOfService#ee72f79a as nil")
	}
	b.PutID(HelpAcceptTermsOfServiceRequestTypeID)
	if err := a.ID.Encode(b); err != nil {
		return fmt.Errorf("unable to encode help.acceptTermsOfService#ee72f79a: field id: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (a *HelpAcceptTermsOfServiceRequest) Decode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode help.acceptTermsOfService#ee72f79a to nil")
	}
	if err := b.ConsumeID(HelpAcceptTermsOfServiceRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode help.acceptTermsOfService#ee72f79a: %w", err)
	}
	{
		if err := a.ID.Decode(b); err != nil {
			return fmt.Errorf("unable to decode help.acceptTermsOfService#ee72f79a: field id: %w", err)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for HelpAcceptTermsOfServiceRequest.
var (
	_ bin.Encoder = &HelpAcceptTermsOfServiceRequest{}
	_ bin.Decoder = &HelpAcceptTermsOfServiceRequest{}
)

// HelpAcceptTermsOfService invokes method help.acceptTermsOfService#ee72f79a returning error if any.
// Accept the new terms of service
//
// See https://core.telegram.org/method/help.acceptTermsOfService for reference.
func (c *Client) HelpAcceptTermsOfService(ctx context.Context, id DataJSON) (bool, error) {
	var result BoolBox

	request := &HelpAcceptTermsOfServiceRequest{
		ID: id,
	}
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
