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

// CloseBirthdayUser represents TL type `closeBirthdayUser#800659ee`.
type CloseBirthdayUser struct {
	// User identifier
	UserID int64
	// Birthdate of the user
	Birthdate Birthdate
}

// CloseBirthdayUserTypeID is TL type id of CloseBirthdayUser.
const CloseBirthdayUserTypeID = 0x800659ee

// Ensuring interfaces in compile-time for CloseBirthdayUser.
var (
	_ bin.Encoder     = &CloseBirthdayUser{}
	_ bin.Decoder     = &CloseBirthdayUser{}
	_ bin.BareEncoder = &CloseBirthdayUser{}
	_ bin.BareDecoder = &CloseBirthdayUser{}
)

func (c *CloseBirthdayUser) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.UserID == 0) {
		return false
	}
	if !(c.Birthdate.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *CloseBirthdayUser) String() string {
	if c == nil {
		return "CloseBirthdayUser(nil)"
	}
	type Alias CloseBirthdayUser
	return fmt.Sprintf("CloseBirthdayUser%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*CloseBirthdayUser) TypeID() uint32 {
	return CloseBirthdayUserTypeID
}

// TypeName returns name of type in TL schema.
func (*CloseBirthdayUser) TypeName() string {
	return "closeBirthdayUser"
}

// TypeInfo returns info about TL type.
func (c *CloseBirthdayUser) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "closeBirthdayUser",
		ID:   CloseBirthdayUserTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "UserID",
			SchemaName: "user_id",
		},
		{
			Name:       "Birthdate",
			SchemaName: "birthdate",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *CloseBirthdayUser) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode closeBirthdayUser#800659ee as nil")
	}
	b.PutID(CloseBirthdayUserTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *CloseBirthdayUser) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode closeBirthdayUser#800659ee as nil")
	}
	b.PutInt53(c.UserID)
	if err := c.Birthdate.Encode(b); err != nil {
		return fmt.Errorf("unable to encode closeBirthdayUser#800659ee: field birthdate: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (c *CloseBirthdayUser) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode closeBirthdayUser#800659ee to nil")
	}
	if err := b.ConsumeID(CloseBirthdayUserTypeID); err != nil {
		return fmt.Errorf("unable to decode closeBirthdayUser#800659ee: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *CloseBirthdayUser) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode closeBirthdayUser#800659ee to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode closeBirthdayUser#800659ee: field user_id: %w", err)
		}
		c.UserID = value
	}
	{
		if err := c.Birthdate.Decode(b); err != nil {
			return fmt.Errorf("unable to decode closeBirthdayUser#800659ee: field birthdate: %w", err)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (c *CloseBirthdayUser) EncodeTDLibJSON(b tdjson.Encoder) error {
	if c == nil {
		return fmt.Errorf("can't encode closeBirthdayUser#800659ee as nil")
	}
	b.ObjStart()
	b.PutID("closeBirthdayUser")
	b.Comma()
	b.FieldStart("user_id")
	b.PutInt53(c.UserID)
	b.Comma()
	b.FieldStart("birthdate")
	if err := c.Birthdate.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode closeBirthdayUser#800659ee: field birthdate: %w", err)
	}
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (c *CloseBirthdayUser) DecodeTDLibJSON(b tdjson.Decoder) error {
	if c == nil {
		return fmt.Errorf("can't decode closeBirthdayUser#800659ee to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("closeBirthdayUser"); err != nil {
				return fmt.Errorf("unable to decode closeBirthdayUser#800659ee: %w", err)
			}
		case "user_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode closeBirthdayUser#800659ee: field user_id: %w", err)
			}
			c.UserID = value
		case "birthdate":
			if err := c.Birthdate.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode closeBirthdayUser#800659ee: field birthdate: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetUserID returns value of UserID field.
func (c *CloseBirthdayUser) GetUserID() (value int64) {
	if c == nil {
		return
	}
	return c.UserID
}

// GetBirthdate returns value of Birthdate field.
func (c *CloseBirthdayUser) GetBirthdate() (value Birthdate) {
	if c == nil {
		return
	}
	return c.Birthdate
}
