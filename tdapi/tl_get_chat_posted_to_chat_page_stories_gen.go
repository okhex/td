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

// GetChatPostedToChatPageStoriesRequest represents TL type `getChatPostedToChatPageStories#fd3bc72b`.
type GetChatPostedToChatPageStoriesRequest struct {
	// Chat identifier
	ChatID int64
	// Identifier of the story starting from which stories must be returned; use 0 to get
	// results from pinned and the newest story
	FromStoryID int32
	// The maximum number of stories to be returned
	Limit int32
}

// GetChatPostedToChatPageStoriesRequestTypeID is TL type id of GetChatPostedToChatPageStoriesRequest.
const GetChatPostedToChatPageStoriesRequestTypeID = 0xfd3bc72b

// Ensuring interfaces in compile-time for GetChatPostedToChatPageStoriesRequest.
var (
	_ bin.Encoder     = &GetChatPostedToChatPageStoriesRequest{}
	_ bin.Decoder     = &GetChatPostedToChatPageStoriesRequest{}
	_ bin.BareEncoder = &GetChatPostedToChatPageStoriesRequest{}
	_ bin.BareDecoder = &GetChatPostedToChatPageStoriesRequest{}
)

func (g *GetChatPostedToChatPageStoriesRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.ChatID == 0) {
		return false
	}
	if !(g.FromStoryID == 0) {
		return false
	}
	if !(g.Limit == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetChatPostedToChatPageStoriesRequest) String() string {
	if g == nil {
		return "GetChatPostedToChatPageStoriesRequest(nil)"
	}
	type Alias GetChatPostedToChatPageStoriesRequest
	return fmt.Sprintf("GetChatPostedToChatPageStoriesRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetChatPostedToChatPageStoriesRequest) TypeID() uint32 {
	return GetChatPostedToChatPageStoriesRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetChatPostedToChatPageStoriesRequest) TypeName() string {
	return "getChatPostedToChatPageStories"
}

// TypeInfo returns info about TL type.
func (g *GetChatPostedToChatPageStoriesRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getChatPostedToChatPageStories",
		ID:   GetChatPostedToChatPageStoriesRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ChatID",
			SchemaName: "chat_id",
		},
		{
			Name:       "FromStoryID",
			SchemaName: "from_story_id",
		},
		{
			Name:       "Limit",
			SchemaName: "limit",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetChatPostedToChatPageStoriesRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getChatPostedToChatPageStories#fd3bc72b as nil")
	}
	b.PutID(GetChatPostedToChatPageStoriesRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetChatPostedToChatPageStoriesRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getChatPostedToChatPageStories#fd3bc72b as nil")
	}
	b.PutInt53(g.ChatID)
	b.PutInt32(g.FromStoryID)
	b.PutInt32(g.Limit)
	return nil
}

// Decode implements bin.Decoder.
func (g *GetChatPostedToChatPageStoriesRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getChatPostedToChatPageStories#fd3bc72b to nil")
	}
	if err := b.ConsumeID(GetChatPostedToChatPageStoriesRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getChatPostedToChatPageStories#fd3bc72b: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetChatPostedToChatPageStoriesRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getChatPostedToChatPageStories#fd3bc72b to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode getChatPostedToChatPageStories#fd3bc72b: field chat_id: %w", err)
		}
		g.ChatID = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode getChatPostedToChatPageStories#fd3bc72b: field from_story_id: %w", err)
		}
		g.FromStoryID = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode getChatPostedToChatPageStories#fd3bc72b: field limit: %w", err)
		}
		g.Limit = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (g *GetChatPostedToChatPageStoriesRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getChatPostedToChatPageStories#fd3bc72b as nil")
	}
	b.ObjStart()
	b.PutID("getChatPostedToChatPageStories")
	b.Comma()
	b.FieldStart("chat_id")
	b.PutInt53(g.ChatID)
	b.Comma()
	b.FieldStart("from_story_id")
	b.PutInt32(g.FromStoryID)
	b.Comma()
	b.FieldStart("limit")
	b.PutInt32(g.Limit)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (g *GetChatPostedToChatPageStoriesRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode getChatPostedToChatPageStories#fd3bc72b to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("getChatPostedToChatPageStories"); err != nil {
				return fmt.Errorf("unable to decode getChatPostedToChatPageStories#fd3bc72b: %w", err)
			}
		case "chat_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode getChatPostedToChatPageStories#fd3bc72b: field chat_id: %w", err)
			}
			g.ChatID = value
		case "from_story_id":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode getChatPostedToChatPageStories#fd3bc72b: field from_story_id: %w", err)
			}
			g.FromStoryID = value
		case "limit":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode getChatPostedToChatPageStories#fd3bc72b: field limit: %w", err)
			}
			g.Limit = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetChatID returns value of ChatID field.
func (g *GetChatPostedToChatPageStoriesRequest) GetChatID() (value int64) {
	if g == nil {
		return
	}
	return g.ChatID
}

// GetFromStoryID returns value of FromStoryID field.
func (g *GetChatPostedToChatPageStoriesRequest) GetFromStoryID() (value int32) {
	if g == nil {
		return
	}
	return g.FromStoryID
}

// GetLimit returns value of Limit field.
func (g *GetChatPostedToChatPageStoriesRequest) GetLimit() (value int32) {
	if g == nil {
		return
	}
	return g.Limit
}

// GetChatPostedToChatPageStories invokes method getChatPostedToChatPageStories#fd3bc72b returning error if any.
func (c *Client) GetChatPostedToChatPageStories(ctx context.Context, request *GetChatPostedToChatPageStoriesRequest) (*Stories, error) {
	var result Stories

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
