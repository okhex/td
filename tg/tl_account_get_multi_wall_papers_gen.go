// Code generated by gotdgen, DO NOT EDIT.

package tg

import (
	"context"
	"fmt"

	"github.com/ernado/td/bin"
)

// No-op definition for keeping imports.
var _ = bin.Buffer{}
var _ = context.Background()
var _ = fmt.Stringer(nil)

// AccountGetMultiWallPapersRequest represents TL type `account.getMultiWallPapers#65ad71dc`.
type AccountGetMultiWallPapersRequest struct {
	// Wallpapers field of AccountGetMultiWallPapersRequest.
	Wallpapers []InputWallPaperClass
}

// AccountGetMultiWallPapersRequestTypeID is TL type id of AccountGetMultiWallPapersRequest.
const AccountGetMultiWallPapersRequestTypeID = 0x65ad71dc

// Encode implements bin.Encoder.
func (g *AccountGetMultiWallPapersRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode account.getMultiWallPapers#65ad71dc as nil")
	}
	b.PutID(AccountGetMultiWallPapersRequestTypeID)
	b.PutVectorHeader(len(g.Wallpapers))
	for idx, v := range g.Wallpapers {
		if v == nil {
			return fmt.Errorf("unable to encode account.getMultiWallPapers#65ad71dc: field wallpapers element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode account.getMultiWallPapers#65ad71dc: field wallpapers element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *AccountGetMultiWallPapersRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode account.getMultiWallPapers#65ad71dc to nil")
	}
	if err := b.ConsumeID(AccountGetMultiWallPapersRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.getMultiWallPapers#65ad71dc: %w", err)
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode account.getMultiWallPapers#65ad71dc: field wallpapers: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeInputWallPaper(b)
			if err != nil {
				return fmt.Errorf("unable to decode account.getMultiWallPapers#65ad71dc: field wallpapers: %w", err)
			}
			g.Wallpapers = append(g.Wallpapers, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for AccountGetMultiWallPapersRequest.
var (
	_ bin.Encoder = &AccountGetMultiWallPapersRequest{}
	_ bin.Decoder = &AccountGetMultiWallPapersRequest{}
)