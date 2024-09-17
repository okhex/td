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

// GetStarGiveawayPaymentOptionsRequest represents TL type `getStarGiveawayPaymentOptions#cb5bdb1e`.
type GetStarGiveawayPaymentOptionsRequest struct {
}

// GetStarGiveawayPaymentOptionsRequestTypeID is TL type id of GetStarGiveawayPaymentOptionsRequest.
const GetStarGiveawayPaymentOptionsRequestTypeID = 0xcb5bdb1e

// Ensuring interfaces in compile-time for GetStarGiveawayPaymentOptionsRequest.
var (
	_ bin.Encoder     = &GetStarGiveawayPaymentOptionsRequest{}
	_ bin.Decoder     = &GetStarGiveawayPaymentOptionsRequest{}
	_ bin.BareEncoder = &GetStarGiveawayPaymentOptionsRequest{}
	_ bin.BareDecoder = &GetStarGiveawayPaymentOptionsRequest{}
)

func (g *GetStarGiveawayPaymentOptionsRequest) Zero() bool {
	if g == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetStarGiveawayPaymentOptionsRequest) String() string {
	if g == nil {
		return "GetStarGiveawayPaymentOptionsRequest(nil)"
	}
	type Alias GetStarGiveawayPaymentOptionsRequest
	return fmt.Sprintf("GetStarGiveawayPaymentOptionsRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetStarGiveawayPaymentOptionsRequest) TypeID() uint32 {
	return GetStarGiveawayPaymentOptionsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetStarGiveawayPaymentOptionsRequest) TypeName() string {
	return "getStarGiveawayPaymentOptions"
}

// TypeInfo returns info about TL type.
func (g *GetStarGiveawayPaymentOptionsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getStarGiveawayPaymentOptions",
		ID:   GetStarGiveawayPaymentOptionsRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetStarGiveawayPaymentOptionsRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getStarGiveawayPaymentOptions#cb5bdb1e as nil")
	}
	b.PutID(GetStarGiveawayPaymentOptionsRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetStarGiveawayPaymentOptionsRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getStarGiveawayPaymentOptions#cb5bdb1e as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *GetStarGiveawayPaymentOptionsRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getStarGiveawayPaymentOptions#cb5bdb1e to nil")
	}
	if err := b.ConsumeID(GetStarGiveawayPaymentOptionsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getStarGiveawayPaymentOptions#cb5bdb1e: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetStarGiveawayPaymentOptionsRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getStarGiveawayPaymentOptions#cb5bdb1e to nil")
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (g *GetStarGiveawayPaymentOptionsRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getStarGiveawayPaymentOptions#cb5bdb1e as nil")
	}
	b.ObjStart()
	b.PutID("getStarGiveawayPaymentOptions")
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (g *GetStarGiveawayPaymentOptionsRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode getStarGiveawayPaymentOptions#cb5bdb1e to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("getStarGiveawayPaymentOptions"); err != nil {
				return fmt.Errorf("unable to decode getStarGiveawayPaymentOptions#cb5bdb1e: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetStarGiveawayPaymentOptions invokes method getStarGiveawayPaymentOptions#cb5bdb1e returning error if any.
func (c *Client) GetStarGiveawayPaymentOptions(ctx context.Context) (*StarGiveawayPaymentOptions, error) {
	var result StarGiveawayPaymentOptions

	request := &GetStarGiveawayPaymentOptionsRequest{}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}