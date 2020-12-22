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

// AccountGetThemeRequest represents TL type `account.getTheme#8d9d742b`.
// Get theme information
//
// See https://core.telegram.org/method/account.getTheme for reference.
type AccountGetThemeRequest struct {
	// Theme format, a string that identifies the theming engines supported by the client
	Format string
	// Theme
	Theme InputThemeClass
	// Document ID
	DocumentID int64
}

// AccountGetThemeRequestTypeID is TL type id of AccountGetThemeRequest.
const AccountGetThemeRequestTypeID = 0x8d9d742b

// String implements fmt.Stringer.
func (g *AccountGetThemeRequest) String() string {
	if g == nil {
		return "AccountGetThemeRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("AccountGetThemeRequest")
	sb.WriteString("{\n")
	sb.WriteString("\tFormat: ")
	sb.WriteString(fmt.Sprint(g.Format))
	sb.WriteString(",\n")
	sb.WriteString("\tTheme: ")
	sb.WriteString(g.Theme.String())
	sb.WriteString(",\n")
	sb.WriteString("\tDocumentID: ")
	sb.WriteString(fmt.Sprint(g.DocumentID))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (g *AccountGetThemeRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode account.getTheme#8d9d742b as nil")
	}
	b.PutID(AccountGetThemeRequestTypeID)
	b.PutString(g.Format)
	if g.Theme == nil {
		return fmt.Errorf("unable to encode account.getTheme#8d9d742b: field theme is nil")
	}
	if err := g.Theme.Encode(b); err != nil {
		return fmt.Errorf("unable to encode account.getTheme#8d9d742b: field theme: %w", err)
	}
	b.PutLong(g.DocumentID)
	return nil
}

// Decode implements bin.Decoder.
func (g *AccountGetThemeRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode account.getTheme#8d9d742b to nil")
	}
	if err := b.ConsumeID(AccountGetThemeRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.getTheme#8d9d742b: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode account.getTheme#8d9d742b: field format: %w", err)
		}
		g.Format = value
	}
	{
		value, err := DecodeInputTheme(b)
		if err != nil {
			return fmt.Errorf("unable to decode account.getTheme#8d9d742b: field theme: %w", err)
		}
		g.Theme = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode account.getTheme#8d9d742b: field document_id: %w", err)
		}
		g.DocumentID = value
	}
	return nil
}

// Ensuring interfaces in compile-time for AccountGetThemeRequest.
var (
	_ bin.Encoder = &AccountGetThemeRequest{}
	_ bin.Decoder = &AccountGetThemeRequest{}
)

// AccountGetTheme invokes method account.getTheme#8d9d742b returning error if any.
// Get theme information
//
// See https://core.telegram.org/method/account.getTheme for reference.
func (c *Client) AccountGetTheme(ctx context.Context, request *AccountGetThemeRequest) (*Theme, error) {
	var result Theme

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
