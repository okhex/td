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

// CollectibleItemTypeUsername represents TL type `collectibleItemTypeUsername#1b56e7d1`.
type CollectibleItemTypeUsername struct {
	// The username
	Username string
}

// CollectibleItemTypeUsernameTypeID is TL type id of CollectibleItemTypeUsername.
const CollectibleItemTypeUsernameTypeID = 0x1b56e7d1

// construct implements constructor of CollectibleItemTypeClass.
func (c CollectibleItemTypeUsername) construct() CollectibleItemTypeClass { return &c }

// Ensuring interfaces in compile-time for CollectibleItemTypeUsername.
var (
	_ bin.Encoder     = &CollectibleItemTypeUsername{}
	_ bin.Decoder     = &CollectibleItemTypeUsername{}
	_ bin.BareEncoder = &CollectibleItemTypeUsername{}
	_ bin.BareDecoder = &CollectibleItemTypeUsername{}

	_ CollectibleItemTypeClass = &CollectibleItemTypeUsername{}
)

func (c *CollectibleItemTypeUsername) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.Username == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *CollectibleItemTypeUsername) String() string {
	if c == nil {
		return "CollectibleItemTypeUsername(nil)"
	}
	type Alias CollectibleItemTypeUsername
	return fmt.Sprintf("CollectibleItemTypeUsername%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*CollectibleItemTypeUsername) TypeID() uint32 {
	return CollectibleItemTypeUsernameTypeID
}

// TypeName returns name of type in TL schema.
func (*CollectibleItemTypeUsername) TypeName() string {
	return "collectibleItemTypeUsername"
}

// TypeInfo returns info about TL type.
func (c *CollectibleItemTypeUsername) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "collectibleItemTypeUsername",
		ID:   CollectibleItemTypeUsernameTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Username",
			SchemaName: "username",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *CollectibleItemTypeUsername) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode collectibleItemTypeUsername#1b56e7d1 as nil")
	}
	b.PutID(CollectibleItemTypeUsernameTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *CollectibleItemTypeUsername) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode collectibleItemTypeUsername#1b56e7d1 as nil")
	}
	b.PutString(c.Username)
	return nil
}

// Decode implements bin.Decoder.
func (c *CollectibleItemTypeUsername) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode collectibleItemTypeUsername#1b56e7d1 to nil")
	}
	if err := b.ConsumeID(CollectibleItemTypeUsernameTypeID); err != nil {
		return fmt.Errorf("unable to decode collectibleItemTypeUsername#1b56e7d1: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *CollectibleItemTypeUsername) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode collectibleItemTypeUsername#1b56e7d1 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode collectibleItemTypeUsername#1b56e7d1: field username: %w", err)
		}
		c.Username = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (c *CollectibleItemTypeUsername) EncodeTDLibJSON(b tdjson.Encoder) error {
	if c == nil {
		return fmt.Errorf("can't encode collectibleItemTypeUsername#1b56e7d1 as nil")
	}
	b.ObjStart()
	b.PutID("collectibleItemTypeUsername")
	b.Comma()
	b.FieldStart("username")
	b.PutString(c.Username)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (c *CollectibleItemTypeUsername) DecodeTDLibJSON(b tdjson.Decoder) error {
	if c == nil {
		return fmt.Errorf("can't decode collectibleItemTypeUsername#1b56e7d1 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("collectibleItemTypeUsername"); err != nil {
				return fmt.Errorf("unable to decode collectibleItemTypeUsername#1b56e7d1: %w", err)
			}
		case "username":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode collectibleItemTypeUsername#1b56e7d1: field username: %w", err)
			}
			c.Username = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetUsername returns value of Username field.
func (c *CollectibleItemTypeUsername) GetUsername() (value string) {
	if c == nil {
		return
	}
	return c.Username
}

// CollectibleItemTypePhoneNumber represents TL type `collectibleItemTypePhoneNumber#4ae0e142`.
type CollectibleItemTypePhoneNumber struct {
	// The phone number
	PhoneNumber string
}

// CollectibleItemTypePhoneNumberTypeID is TL type id of CollectibleItemTypePhoneNumber.
const CollectibleItemTypePhoneNumberTypeID = 0x4ae0e142

// construct implements constructor of CollectibleItemTypeClass.
func (c CollectibleItemTypePhoneNumber) construct() CollectibleItemTypeClass { return &c }

// Ensuring interfaces in compile-time for CollectibleItemTypePhoneNumber.
var (
	_ bin.Encoder     = &CollectibleItemTypePhoneNumber{}
	_ bin.Decoder     = &CollectibleItemTypePhoneNumber{}
	_ bin.BareEncoder = &CollectibleItemTypePhoneNumber{}
	_ bin.BareDecoder = &CollectibleItemTypePhoneNumber{}

	_ CollectibleItemTypeClass = &CollectibleItemTypePhoneNumber{}
)

func (c *CollectibleItemTypePhoneNumber) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.PhoneNumber == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *CollectibleItemTypePhoneNumber) String() string {
	if c == nil {
		return "CollectibleItemTypePhoneNumber(nil)"
	}
	type Alias CollectibleItemTypePhoneNumber
	return fmt.Sprintf("CollectibleItemTypePhoneNumber%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*CollectibleItemTypePhoneNumber) TypeID() uint32 {
	return CollectibleItemTypePhoneNumberTypeID
}

// TypeName returns name of type in TL schema.
func (*CollectibleItemTypePhoneNumber) TypeName() string {
	return "collectibleItemTypePhoneNumber"
}

// TypeInfo returns info about TL type.
func (c *CollectibleItemTypePhoneNumber) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "collectibleItemTypePhoneNumber",
		ID:   CollectibleItemTypePhoneNumberTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "PhoneNumber",
			SchemaName: "phone_number",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *CollectibleItemTypePhoneNumber) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode collectibleItemTypePhoneNumber#4ae0e142 as nil")
	}
	b.PutID(CollectibleItemTypePhoneNumberTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *CollectibleItemTypePhoneNumber) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode collectibleItemTypePhoneNumber#4ae0e142 as nil")
	}
	b.PutString(c.PhoneNumber)
	return nil
}

// Decode implements bin.Decoder.
func (c *CollectibleItemTypePhoneNumber) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode collectibleItemTypePhoneNumber#4ae0e142 to nil")
	}
	if err := b.ConsumeID(CollectibleItemTypePhoneNumberTypeID); err != nil {
		return fmt.Errorf("unable to decode collectibleItemTypePhoneNumber#4ae0e142: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *CollectibleItemTypePhoneNumber) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode collectibleItemTypePhoneNumber#4ae0e142 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode collectibleItemTypePhoneNumber#4ae0e142: field phone_number: %w", err)
		}
		c.PhoneNumber = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (c *CollectibleItemTypePhoneNumber) EncodeTDLibJSON(b tdjson.Encoder) error {
	if c == nil {
		return fmt.Errorf("can't encode collectibleItemTypePhoneNumber#4ae0e142 as nil")
	}
	b.ObjStart()
	b.PutID("collectibleItemTypePhoneNumber")
	b.Comma()
	b.FieldStart("phone_number")
	b.PutString(c.PhoneNumber)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (c *CollectibleItemTypePhoneNumber) DecodeTDLibJSON(b tdjson.Decoder) error {
	if c == nil {
		return fmt.Errorf("can't decode collectibleItemTypePhoneNumber#4ae0e142 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("collectibleItemTypePhoneNumber"); err != nil {
				return fmt.Errorf("unable to decode collectibleItemTypePhoneNumber#4ae0e142: %w", err)
			}
		case "phone_number":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode collectibleItemTypePhoneNumber#4ae0e142: field phone_number: %w", err)
			}
			c.PhoneNumber = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetPhoneNumber returns value of PhoneNumber field.
func (c *CollectibleItemTypePhoneNumber) GetPhoneNumber() (value string) {
	if c == nil {
		return
	}
	return c.PhoneNumber
}

// CollectibleItemTypeClassName is schema name of CollectibleItemTypeClass.
const CollectibleItemTypeClassName = "CollectibleItemType"

// CollectibleItemTypeClass represents CollectibleItemType generic type.
//
// Example:
//
//	g, err := tdapi.DecodeCollectibleItemType(buf)
//	if err != nil {
//	    panic(err)
//	}
//	switch v := g.(type) {
//	case *tdapi.CollectibleItemTypeUsername: // collectibleItemTypeUsername#1b56e7d1
//	case *tdapi.CollectibleItemTypePhoneNumber: // collectibleItemTypePhoneNumber#4ae0e142
//	default: panic(v)
//	}
type CollectibleItemTypeClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() CollectibleItemTypeClass

	// TypeID returns type id in TL schema.
	//
	// See https://core.telegram.org/mtproto/TL-tl#remarks.
	TypeID() uint32
	// TypeName returns name of type in TL schema.
	TypeName() string
	// String implements fmt.Stringer.
	String() string
	// Zero returns true if current object has a zero value.
	Zero() bool

	EncodeTDLibJSON(b tdjson.Encoder) error
	DecodeTDLibJSON(b tdjson.Decoder) error
}

// DecodeCollectibleItemType implements binary de-serialization for CollectibleItemTypeClass.
func DecodeCollectibleItemType(buf *bin.Buffer) (CollectibleItemTypeClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case CollectibleItemTypeUsernameTypeID:
		// Decoding collectibleItemTypeUsername#1b56e7d1.
		v := CollectibleItemTypeUsername{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode CollectibleItemTypeClass: %w", err)
		}
		return &v, nil
	case CollectibleItemTypePhoneNumberTypeID:
		// Decoding collectibleItemTypePhoneNumber#4ae0e142.
		v := CollectibleItemTypePhoneNumber{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode CollectibleItemTypeClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode CollectibleItemTypeClass: %w", bin.NewUnexpectedID(id))
	}
}

// DecodeTDLibJSONCollectibleItemType implements binary de-serialization for CollectibleItemTypeClass.
func DecodeTDLibJSONCollectibleItemType(buf tdjson.Decoder) (CollectibleItemTypeClass, error) {
	id, err := buf.FindTypeID()
	if err != nil {
		return nil, err
	}
	switch id {
	case "collectibleItemTypeUsername":
		// Decoding collectibleItemTypeUsername#1b56e7d1.
		v := CollectibleItemTypeUsername{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode CollectibleItemTypeClass: %w", err)
		}
		return &v, nil
	case "collectibleItemTypePhoneNumber":
		// Decoding collectibleItemTypePhoneNumber#4ae0e142.
		v := CollectibleItemTypePhoneNumber{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode CollectibleItemTypeClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode CollectibleItemTypeClass: %w", tdjson.NewUnexpectedID(id))
	}
}

// CollectibleItemType boxes the CollectibleItemTypeClass providing a helper.
type CollectibleItemTypeBox struct {
	CollectibleItemType CollectibleItemTypeClass
}

// Decode implements bin.Decoder for CollectibleItemTypeBox.
func (b *CollectibleItemTypeBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode CollectibleItemTypeBox to nil")
	}
	v, err := DecodeCollectibleItemType(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.CollectibleItemType = v
	return nil
}

// Encode implements bin.Encode for CollectibleItemTypeBox.
func (b *CollectibleItemTypeBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.CollectibleItemType == nil {
		return fmt.Errorf("unable to encode CollectibleItemTypeClass as nil")
	}
	return b.CollectibleItemType.Encode(buf)
}

// DecodeTDLibJSON implements bin.Decoder for CollectibleItemTypeBox.
func (b *CollectibleItemTypeBox) DecodeTDLibJSON(buf tdjson.Decoder) error {
	if b == nil {
		return fmt.Errorf("unable to decode CollectibleItemTypeBox to nil")
	}
	v, err := DecodeTDLibJSONCollectibleItemType(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.CollectibleItemType = v
	return nil
}

// EncodeTDLibJSON implements bin.Encode for CollectibleItemTypeBox.
func (b *CollectibleItemTypeBox) EncodeTDLibJSON(buf tdjson.Encoder) error {
	if b == nil || b.CollectibleItemType == nil {
		return fmt.Errorf("unable to encode CollectibleItemTypeClass as nil")
	}
	return b.CollectibleItemType.EncodeTDLibJSON(buf)
}
