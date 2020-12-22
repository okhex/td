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

// AccountSaveWallPaperRequest represents TL type `account.saveWallPaper#6c5a5b37`.
// Install/uninstall wallpaper
//
// See https://core.telegram.org/method/account.saveWallPaper for reference.
type AccountSaveWallPaperRequest struct {
	// Wallpaper to save
	Wallpaper InputWallPaperClass
	// Uninstall wallpaper?
	Unsave bool
	// Wallpaper settings
	Settings WallPaperSettings
}

// AccountSaveWallPaperRequestTypeID is TL type id of AccountSaveWallPaperRequest.
const AccountSaveWallPaperRequestTypeID = 0x6c5a5b37

// String implements fmt.Stringer.
func (s *AccountSaveWallPaperRequest) String() string {
	if s == nil {
		return "AccountSaveWallPaperRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("AccountSaveWallPaperRequest")
	sb.WriteString("{\n")
	sb.WriteString("\tWallpaper: ")
	sb.WriteString(s.Wallpaper.String())
	sb.WriteString(",\n")
	sb.WriteString("\tUnsave: ")
	sb.WriteString(fmt.Sprint(s.Unsave))
	sb.WriteString(",\n")
	sb.WriteString("\tSettings: ")
	sb.WriteString(s.Settings.String())
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (s *AccountSaveWallPaperRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode account.saveWallPaper#6c5a5b37 as nil")
	}
	b.PutID(AccountSaveWallPaperRequestTypeID)
	if s.Wallpaper == nil {
		return fmt.Errorf("unable to encode account.saveWallPaper#6c5a5b37: field wallpaper is nil")
	}
	if err := s.Wallpaper.Encode(b); err != nil {
		return fmt.Errorf("unable to encode account.saveWallPaper#6c5a5b37: field wallpaper: %w", err)
	}
	b.PutBool(s.Unsave)
	if err := s.Settings.Encode(b); err != nil {
		return fmt.Errorf("unable to encode account.saveWallPaper#6c5a5b37: field settings: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *AccountSaveWallPaperRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode account.saveWallPaper#6c5a5b37 to nil")
	}
	if err := b.ConsumeID(AccountSaveWallPaperRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.saveWallPaper#6c5a5b37: %w", err)
	}
	{
		value, err := DecodeInputWallPaper(b)
		if err != nil {
			return fmt.Errorf("unable to decode account.saveWallPaper#6c5a5b37: field wallpaper: %w", err)
		}
		s.Wallpaper = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode account.saveWallPaper#6c5a5b37: field unsave: %w", err)
		}
		s.Unsave = value
	}
	{
		if err := s.Settings.Decode(b); err != nil {
			return fmt.Errorf("unable to decode account.saveWallPaper#6c5a5b37: field settings: %w", err)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for AccountSaveWallPaperRequest.
var (
	_ bin.Encoder = &AccountSaveWallPaperRequest{}
	_ bin.Decoder = &AccountSaveWallPaperRequest{}
)

// AccountSaveWallPaper invokes method account.saveWallPaper#6c5a5b37 returning error if any.
// Install/uninstall wallpaper
//
// See https://core.telegram.org/method/account.saveWallPaper for reference.
func (c *Client) AccountSaveWallPaper(ctx context.Context, request *AccountSaveWallPaperRequest) (bool, error) {
	var result BoolBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
