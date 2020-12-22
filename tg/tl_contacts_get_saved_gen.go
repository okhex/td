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

// ContactsGetSavedRequest represents TL type `contacts.getSaved#82f1e39f`.
// Get all contacts
//
// See https://core.telegram.org/method/contacts.getSaved for reference.
type ContactsGetSavedRequest struct {
}

// ContactsGetSavedRequestTypeID is TL type id of ContactsGetSavedRequest.
const ContactsGetSavedRequestTypeID = 0x82f1e39f

// String implements fmt.Stringer.
func (g *ContactsGetSavedRequest) String() string {
	if g == nil {
		return "ContactsGetSavedRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("ContactsGetSavedRequest")
	sb.WriteString("{\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (g *ContactsGetSavedRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode contacts.getSaved#82f1e39f as nil")
	}
	b.PutID(ContactsGetSavedRequestTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (g *ContactsGetSavedRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode contacts.getSaved#82f1e39f to nil")
	}
	if err := b.ConsumeID(ContactsGetSavedRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode contacts.getSaved#82f1e39f: %w", err)
	}
	return nil
}

// Ensuring interfaces in compile-time for ContactsGetSavedRequest.
var (
	_ bin.Encoder = &ContactsGetSavedRequest{}
	_ bin.Decoder = &ContactsGetSavedRequest{}
)

// ContactsGetSaved invokes method contacts.getSaved#82f1e39f returning error if any.
// Get all contacts
//
// See https://core.telegram.org/method/contacts.getSaved for reference.
func (c *Client) ContactsGetSaved(ctx context.Context) ([]SavedPhoneContact, error) {
	var result SavedPhoneContactVector

	request := &ContactsGetSavedRequest{}
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Elems, nil
}
