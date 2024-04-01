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

// SharedUser represents TL type `sharedUser#117724f7`.
type SharedUser struct {
	// User identifier
	UserID int64
	// First name of the user; for bots only
	FirstName string
	// Last name of the user; for bots only
	LastName string
	// Username of the user; for bots only
	Username string
	// Profile photo of the user; for bots only; may be null
	Photo Photo
}

// SharedUserTypeID is TL type id of SharedUser.
const SharedUserTypeID = 0x117724f7

// Ensuring interfaces in compile-time for SharedUser.
var (
	_ bin.Encoder     = &SharedUser{}
	_ bin.Decoder     = &SharedUser{}
	_ bin.BareEncoder = &SharedUser{}
	_ bin.BareDecoder = &SharedUser{}
)

func (s *SharedUser) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.UserID == 0) {
		return false
	}
	if !(s.FirstName == "") {
		return false
	}
	if !(s.LastName == "") {
		return false
	}
	if !(s.Username == "") {
		return false
	}
	if !(s.Photo.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *SharedUser) String() string {
	if s == nil {
		return "SharedUser(nil)"
	}
	type Alias SharedUser
	return fmt.Sprintf("SharedUser%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*SharedUser) TypeID() uint32 {
	return SharedUserTypeID
}

// TypeName returns name of type in TL schema.
func (*SharedUser) TypeName() string {
	return "sharedUser"
}

// TypeInfo returns info about TL type.
func (s *SharedUser) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "sharedUser",
		ID:   SharedUserTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "UserID",
			SchemaName: "user_id",
		},
		{
			Name:       "FirstName",
			SchemaName: "first_name",
		},
		{
			Name:       "LastName",
			SchemaName: "last_name",
		},
		{
			Name:       "Username",
			SchemaName: "username",
		},
		{
			Name:       "Photo",
			SchemaName: "photo",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *SharedUser) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode sharedUser#117724f7 as nil")
	}
	b.PutID(SharedUserTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *SharedUser) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode sharedUser#117724f7 as nil")
	}
	b.PutInt53(s.UserID)
	b.PutString(s.FirstName)
	b.PutString(s.LastName)
	b.PutString(s.Username)
	if err := s.Photo.Encode(b); err != nil {
		return fmt.Errorf("unable to encode sharedUser#117724f7: field photo: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *SharedUser) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode sharedUser#117724f7 to nil")
	}
	if err := b.ConsumeID(SharedUserTypeID); err != nil {
		return fmt.Errorf("unable to decode sharedUser#117724f7: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *SharedUser) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode sharedUser#117724f7 to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode sharedUser#117724f7: field user_id: %w", err)
		}
		s.UserID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode sharedUser#117724f7: field first_name: %w", err)
		}
		s.FirstName = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode sharedUser#117724f7: field last_name: %w", err)
		}
		s.LastName = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode sharedUser#117724f7: field username: %w", err)
		}
		s.Username = value
	}
	{
		if err := s.Photo.Decode(b); err != nil {
			return fmt.Errorf("unable to decode sharedUser#117724f7: field photo: %w", err)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (s *SharedUser) EncodeTDLibJSON(b tdjson.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode sharedUser#117724f7 as nil")
	}
	b.ObjStart()
	b.PutID("sharedUser")
	b.Comma()
	b.FieldStart("user_id")
	b.PutInt53(s.UserID)
	b.Comma()
	b.FieldStart("first_name")
	b.PutString(s.FirstName)
	b.Comma()
	b.FieldStart("last_name")
	b.PutString(s.LastName)
	b.Comma()
	b.FieldStart("username")
	b.PutString(s.Username)
	b.Comma()
	b.FieldStart("photo")
	if err := s.Photo.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode sharedUser#117724f7: field photo: %w", err)
	}
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (s *SharedUser) DecodeTDLibJSON(b tdjson.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode sharedUser#117724f7 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("sharedUser"); err != nil {
				return fmt.Errorf("unable to decode sharedUser#117724f7: %w", err)
			}
		case "user_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode sharedUser#117724f7: field user_id: %w", err)
			}
			s.UserID = value
		case "first_name":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode sharedUser#117724f7: field first_name: %w", err)
			}
			s.FirstName = value
		case "last_name":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode sharedUser#117724f7: field last_name: %w", err)
			}
			s.LastName = value
		case "username":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode sharedUser#117724f7: field username: %w", err)
			}
			s.Username = value
		case "photo":
			if err := s.Photo.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode sharedUser#117724f7: field photo: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetUserID returns value of UserID field.
func (s *SharedUser) GetUserID() (value int64) {
	if s == nil {
		return
	}
	return s.UserID
}

// GetFirstName returns value of FirstName field.
func (s *SharedUser) GetFirstName() (value string) {
	if s == nil {
		return
	}
	return s.FirstName
}

// GetLastName returns value of LastName field.
func (s *SharedUser) GetLastName() (value string) {
	if s == nil {
		return
	}
	return s.LastName
}

// GetUsername returns value of Username field.
func (s *SharedUser) GetUsername() (value string) {
	if s == nil {
		return
	}
	return s.Username
}

// GetPhoto returns value of Photo field.
func (s *SharedUser) GetPhoto() (value Photo) {
	if s == nil {
		return
	}
	return s.Photo
}
