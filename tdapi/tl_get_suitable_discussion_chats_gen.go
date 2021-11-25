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

// GetSuitableDiscussionChatsRequest represents TL type `getSuitableDiscussionChats#2ec5df6`.
type GetSuitableDiscussionChatsRequest struct {
}

// GetSuitableDiscussionChatsRequestTypeID is TL type id of GetSuitableDiscussionChatsRequest.
const GetSuitableDiscussionChatsRequestTypeID = 0x2ec5df6

// Ensuring interfaces in compile-time for GetSuitableDiscussionChatsRequest.
var (
	_ bin.Encoder     = &GetSuitableDiscussionChatsRequest{}
	_ bin.Decoder     = &GetSuitableDiscussionChatsRequest{}
	_ bin.BareEncoder = &GetSuitableDiscussionChatsRequest{}
	_ bin.BareDecoder = &GetSuitableDiscussionChatsRequest{}
)

func (g *GetSuitableDiscussionChatsRequest) Zero() bool {
	if g == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetSuitableDiscussionChatsRequest) String() string {
	if g == nil {
		return "GetSuitableDiscussionChatsRequest(nil)"
	}
	type Alias GetSuitableDiscussionChatsRequest
	return fmt.Sprintf("GetSuitableDiscussionChatsRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetSuitableDiscussionChatsRequest) TypeID() uint32 {
	return GetSuitableDiscussionChatsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetSuitableDiscussionChatsRequest) TypeName() string {
	return "getSuitableDiscussionChats"
}

// TypeInfo returns info about TL type.
func (g *GetSuitableDiscussionChatsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getSuitableDiscussionChats",
		ID:   GetSuitableDiscussionChatsRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetSuitableDiscussionChatsRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getSuitableDiscussionChats#2ec5df6 as nil")
	}
	b.PutID(GetSuitableDiscussionChatsRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetSuitableDiscussionChatsRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getSuitableDiscussionChats#2ec5df6 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *GetSuitableDiscussionChatsRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getSuitableDiscussionChats#2ec5df6 to nil")
	}
	if err := b.ConsumeID(GetSuitableDiscussionChatsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getSuitableDiscussionChats#2ec5df6: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetSuitableDiscussionChatsRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getSuitableDiscussionChats#2ec5df6 to nil")
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (g *GetSuitableDiscussionChatsRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getSuitableDiscussionChats#2ec5df6 as nil")
	}
	b.ObjStart()
	b.PutID("getSuitableDiscussionChats")
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (g *GetSuitableDiscussionChatsRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode getSuitableDiscussionChats#2ec5df6 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("getSuitableDiscussionChats"); err != nil {
				return fmt.Errorf("unable to decode getSuitableDiscussionChats#2ec5df6: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetSuitableDiscussionChats invokes method getSuitableDiscussionChats#2ec5df6 returning error if any.
func (c *Client) GetSuitableDiscussionChats(ctx context.Context) (*Chats, error) {
	var result Chats

	request := &GetSuitableDiscussionChatsRequest{}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}