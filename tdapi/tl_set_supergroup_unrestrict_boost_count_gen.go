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

// SetSupergroupUnrestrictBoostCountRequest represents TL type `setSupergroupUnrestrictBoostCount#39ce30a3`.
type SetSupergroupUnrestrictBoostCountRequest struct {
	// Identifier of the supergroup
	SupergroupID int64
	// New value of the unrestrict_boost_count supergroup setting; 0-8. Use 0 to remove the
	// setting
	UnrestrictBoostCount int32
}

// SetSupergroupUnrestrictBoostCountRequestTypeID is TL type id of SetSupergroupUnrestrictBoostCountRequest.
const SetSupergroupUnrestrictBoostCountRequestTypeID = 0x39ce30a3

// Ensuring interfaces in compile-time for SetSupergroupUnrestrictBoostCountRequest.
var (
	_ bin.Encoder     = &SetSupergroupUnrestrictBoostCountRequest{}
	_ bin.Decoder     = &SetSupergroupUnrestrictBoostCountRequest{}
	_ bin.BareEncoder = &SetSupergroupUnrestrictBoostCountRequest{}
	_ bin.BareDecoder = &SetSupergroupUnrestrictBoostCountRequest{}
)

func (s *SetSupergroupUnrestrictBoostCountRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.SupergroupID == 0) {
		return false
	}
	if !(s.UnrestrictBoostCount == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *SetSupergroupUnrestrictBoostCountRequest) String() string {
	if s == nil {
		return "SetSupergroupUnrestrictBoostCountRequest(nil)"
	}
	type Alias SetSupergroupUnrestrictBoostCountRequest
	return fmt.Sprintf("SetSupergroupUnrestrictBoostCountRequest%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*SetSupergroupUnrestrictBoostCountRequest) TypeID() uint32 {
	return SetSupergroupUnrestrictBoostCountRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*SetSupergroupUnrestrictBoostCountRequest) TypeName() string {
	return "setSupergroupUnrestrictBoostCount"
}

// TypeInfo returns info about TL type.
func (s *SetSupergroupUnrestrictBoostCountRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "setSupergroupUnrestrictBoostCount",
		ID:   SetSupergroupUnrestrictBoostCountRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "SupergroupID",
			SchemaName: "supergroup_id",
		},
		{
			Name:       "UnrestrictBoostCount",
			SchemaName: "unrestrict_boost_count",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *SetSupergroupUnrestrictBoostCountRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode setSupergroupUnrestrictBoostCount#39ce30a3 as nil")
	}
	b.PutID(SetSupergroupUnrestrictBoostCountRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *SetSupergroupUnrestrictBoostCountRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode setSupergroupUnrestrictBoostCount#39ce30a3 as nil")
	}
	b.PutInt53(s.SupergroupID)
	b.PutInt32(s.UnrestrictBoostCount)
	return nil
}

// Decode implements bin.Decoder.
func (s *SetSupergroupUnrestrictBoostCountRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode setSupergroupUnrestrictBoostCount#39ce30a3 to nil")
	}
	if err := b.ConsumeID(SetSupergroupUnrestrictBoostCountRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode setSupergroupUnrestrictBoostCount#39ce30a3: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *SetSupergroupUnrestrictBoostCountRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode setSupergroupUnrestrictBoostCount#39ce30a3 to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode setSupergroupUnrestrictBoostCount#39ce30a3: field supergroup_id: %w", err)
		}
		s.SupergroupID = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode setSupergroupUnrestrictBoostCount#39ce30a3: field unrestrict_boost_count: %w", err)
		}
		s.UnrestrictBoostCount = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (s *SetSupergroupUnrestrictBoostCountRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode setSupergroupUnrestrictBoostCount#39ce30a3 as nil")
	}
	b.ObjStart()
	b.PutID("setSupergroupUnrestrictBoostCount")
	b.Comma()
	b.FieldStart("supergroup_id")
	b.PutInt53(s.SupergroupID)
	b.Comma()
	b.FieldStart("unrestrict_boost_count")
	b.PutInt32(s.UnrestrictBoostCount)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (s *SetSupergroupUnrestrictBoostCountRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode setSupergroupUnrestrictBoostCount#39ce30a3 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("setSupergroupUnrestrictBoostCount"); err != nil {
				return fmt.Errorf("unable to decode setSupergroupUnrestrictBoostCount#39ce30a3: %w", err)
			}
		case "supergroup_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode setSupergroupUnrestrictBoostCount#39ce30a3: field supergroup_id: %w", err)
			}
			s.SupergroupID = value
		case "unrestrict_boost_count":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode setSupergroupUnrestrictBoostCount#39ce30a3: field unrestrict_boost_count: %w", err)
			}
			s.UnrestrictBoostCount = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetSupergroupID returns value of SupergroupID field.
func (s *SetSupergroupUnrestrictBoostCountRequest) GetSupergroupID() (value int64) {
	if s == nil {
		return
	}
	return s.SupergroupID
}

// GetUnrestrictBoostCount returns value of UnrestrictBoostCount field.
func (s *SetSupergroupUnrestrictBoostCountRequest) GetUnrestrictBoostCount() (value int32) {
	if s == nil {
		return
	}
	return s.UnrestrictBoostCount
}

// SetSupergroupUnrestrictBoostCount invokes method setSupergroupUnrestrictBoostCount#39ce30a3 returning error if any.
func (c *Client) SetSupergroupUnrestrictBoostCount(ctx context.Context, request *SetSupergroupUnrestrictBoostCountRequest) error {
	var ok Ok

	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}