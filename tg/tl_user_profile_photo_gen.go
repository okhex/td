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

// UserProfilePhotoEmpty represents TL type `userProfilePhotoEmpty#4f11bae1`.
// Profile photo has not been set, or was hidden.
//
// See https://core.telegram.org/constructor/userProfilePhotoEmpty for reference.
type UserProfilePhotoEmpty struct {
}

// UserProfilePhotoEmptyTypeID is TL type id of UserProfilePhotoEmpty.
const UserProfilePhotoEmptyTypeID = 0x4f11bae1

// String implements fmt.Stringer.
func (u *UserProfilePhotoEmpty) String() string {
	if u == nil {
		return "UserProfilePhotoEmpty(nil)"
	}
	var sb strings.Builder
	sb.WriteString("UserProfilePhotoEmpty")
	sb.WriteString("{\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (u *UserProfilePhotoEmpty) Encode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode userProfilePhotoEmpty#4f11bae1 as nil")
	}
	b.PutID(UserProfilePhotoEmptyTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (u *UserProfilePhotoEmpty) Decode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode userProfilePhotoEmpty#4f11bae1 to nil")
	}
	if err := b.ConsumeID(UserProfilePhotoEmptyTypeID); err != nil {
		return fmt.Errorf("unable to decode userProfilePhotoEmpty#4f11bae1: %w", err)
	}
	return nil
}

// construct implements constructor of UserProfilePhotoClass.
func (u UserProfilePhotoEmpty) construct() UserProfilePhotoClass { return &u }

// Ensuring interfaces in compile-time for UserProfilePhotoEmpty.
var (
	_ bin.Encoder = &UserProfilePhotoEmpty{}
	_ bin.Decoder = &UserProfilePhotoEmpty{}

	_ UserProfilePhotoClass = &UserProfilePhotoEmpty{}
)

// UserProfilePhoto represents TL type `userProfilePhoto#69d3ab26`.
// User profile photo.
//
// See https://core.telegram.org/constructor/userProfilePhoto for reference.
type UserProfilePhoto struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Whether an animated profile picture¹ is available for this user
	//
	// Links:
	//  1) https://core.telegram.org/api/files#animated-profile-pictures
	HasVideo bool
	// Identifier of the respective photoParameter added in Layer 2¹
	//
	// Links:
	//  1) https://core.telegram.org/api/layers#layer-2
	PhotoID int64
	// Location of the file, corresponding to the small profile photo thumbnail
	PhotoSmall FileLocationToBeDeprecated
	// Location of the file, corresponding to the big profile photo thumbnail
	PhotoBig FileLocationToBeDeprecated
	// DC ID where the photo is stored
	DCID int
}

// UserProfilePhotoTypeID is TL type id of UserProfilePhoto.
const UserProfilePhotoTypeID = 0x69d3ab26

// String implements fmt.Stringer.
func (u *UserProfilePhoto) String() string {
	if u == nil {
		return "UserProfilePhoto(nil)"
	}
	var sb strings.Builder
	sb.WriteString("UserProfilePhoto")
	sb.WriteString("{\n")
	sb.WriteString("\tFlags: ")
	sb.WriteString(u.Flags.String())
	sb.WriteString(",\n")
	sb.WriteString("\tPhotoID: ")
	sb.WriteString(fmt.Sprint(u.PhotoID))
	sb.WriteString(",\n")
	sb.WriteString("\tPhotoSmall: ")
	sb.WriteString(u.PhotoSmall.String())
	sb.WriteString(",\n")
	sb.WriteString("\tPhotoBig: ")
	sb.WriteString(u.PhotoBig.String())
	sb.WriteString(",\n")
	sb.WriteString("\tDCID: ")
	sb.WriteString(fmt.Sprint(u.DCID))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (u *UserProfilePhoto) Encode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode userProfilePhoto#69d3ab26 as nil")
	}
	b.PutID(UserProfilePhotoTypeID)
	if err := u.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode userProfilePhoto#69d3ab26: field flags: %w", err)
	}
	b.PutLong(u.PhotoID)
	if err := u.PhotoSmall.Encode(b); err != nil {
		return fmt.Errorf("unable to encode userProfilePhoto#69d3ab26: field photo_small: %w", err)
	}
	if err := u.PhotoBig.Encode(b); err != nil {
		return fmt.Errorf("unable to encode userProfilePhoto#69d3ab26: field photo_big: %w", err)
	}
	b.PutInt(u.DCID)
	return nil
}

// SetHasVideo sets value of HasVideo conditional field.
func (u *UserProfilePhoto) SetHasVideo(value bool) {
	if value {
		u.Flags.Set(0)
	} else {
		u.Flags.Unset(0)
	}
}

// Decode implements bin.Decoder.
func (u *UserProfilePhoto) Decode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode userProfilePhoto#69d3ab26 to nil")
	}
	if err := b.ConsumeID(UserProfilePhotoTypeID); err != nil {
		return fmt.Errorf("unable to decode userProfilePhoto#69d3ab26: %w", err)
	}
	{
		if err := u.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode userProfilePhoto#69d3ab26: field flags: %w", err)
		}
	}
	u.HasVideo = u.Flags.Has(0)
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode userProfilePhoto#69d3ab26: field photo_id: %w", err)
		}
		u.PhotoID = value
	}
	{
		if err := u.PhotoSmall.Decode(b); err != nil {
			return fmt.Errorf("unable to decode userProfilePhoto#69d3ab26: field photo_small: %w", err)
		}
	}
	{
		if err := u.PhotoBig.Decode(b); err != nil {
			return fmt.Errorf("unable to decode userProfilePhoto#69d3ab26: field photo_big: %w", err)
		}
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode userProfilePhoto#69d3ab26: field dc_id: %w", err)
		}
		u.DCID = value
	}
	return nil
}

// construct implements constructor of UserProfilePhotoClass.
func (u UserProfilePhoto) construct() UserProfilePhotoClass { return &u }

// Ensuring interfaces in compile-time for UserProfilePhoto.
var (
	_ bin.Encoder = &UserProfilePhoto{}
	_ bin.Decoder = &UserProfilePhoto{}

	_ UserProfilePhotoClass = &UserProfilePhoto{}
)

// UserProfilePhotoClass represents UserProfilePhoto generic type.
//
// See https://core.telegram.org/type/UserProfilePhoto for reference.
//
// Example:
//  g, err := DecodeUserProfilePhoto(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *UserProfilePhotoEmpty: // userProfilePhotoEmpty#4f11bae1
//  case *UserProfilePhoto: // userProfilePhoto#69d3ab26
//  default: panic(v)
//  }
type UserProfilePhotoClass interface {
	bin.Encoder
	bin.Decoder
	construct() UserProfilePhotoClass
	fmt.Stringer
}

// DecodeUserProfilePhoto implements binary de-serialization for UserProfilePhotoClass.
func DecodeUserProfilePhoto(buf *bin.Buffer) (UserProfilePhotoClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case UserProfilePhotoEmptyTypeID:
		// Decoding userProfilePhotoEmpty#4f11bae1.
		v := UserProfilePhotoEmpty{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode UserProfilePhotoClass: %w", err)
		}
		return &v, nil
	case UserProfilePhotoTypeID:
		// Decoding userProfilePhoto#69d3ab26.
		v := UserProfilePhoto{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode UserProfilePhotoClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode UserProfilePhotoClass: %w", bin.NewUnexpectedID(id))
	}
}

// UserProfilePhoto boxes the UserProfilePhotoClass providing a helper.
type UserProfilePhotoBox struct {
	UserProfilePhoto UserProfilePhotoClass
}

// Decode implements bin.Decoder for UserProfilePhotoBox.
func (b *UserProfilePhotoBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode UserProfilePhotoBox to nil")
	}
	v, err := DecodeUserProfilePhoto(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.UserProfilePhoto = v
	return nil
}

// Encode implements bin.Encode for UserProfilePhotoBox.
func (b *UserProfilePhotoBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.UserProfilePhoto == nil {
		return fmt.Errorf("unable to encode UserProfilePhotoClass as nil")
	}
	return b.UserProfilePhoto.Encode(buf)
}
