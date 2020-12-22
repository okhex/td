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

// Folder represents TL type `folder#ff544e65`.
// Folder
//
// See https://core.telegram.org/constructor/folder for reference.
type Folder struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Automatically add new channels to this folder
	AutofillNewBroadcasts bool
	// Automatically add joined new public supergroups to this folder
	AutofillPublicGroups bool
	// Automatically add new private chats to this folder
	AutofillNewCorrespondents bool
	// Folder ID
	ID int
	// Folder title
	Title string
	// Folder picture
	//
	// Use SetPhoto and GetPhoto helpers.
	Photo ChatPhotoClass
}

// FolderTypeID is TL type id of Folder.
const FolderTypeID = 0xff544e65

// String implements fmt.Stringer.
func (f *Folder) String() string {
	if f == nil {
		return "Folder(nil)"
	}
	var sb strings.Builder
	sb.WriteString("Folder")
	sb.WriteString("{\n")
	sb.WriteString("\tFlags: ")
	sb.WriteString(f.Flags.String())
	sb.WriteString(",\n")
	sb.WriteString("\tID: ")
	sb.WriteString(fmt.Sprint(f.ID))
	sb.WriteString(",\n")
	sb.WriteString("\tTitle: ")
	sb.WriteString(fmt.Sprint(f.Title))
	sb.WriteString(",\n")
	if f.Flags.Has(3) {
		sb.WriteString("\tPhoto: ")
		sb.WriteString(f.Photo.String())
		sb.WriteString(",\n")
	}
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (f *Folder) Encode(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't encode folder#ff544e65 as nil")
	}
	b.PutID(FolderTypeID)
	if err := f.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode folder#ff544e65: field flags: %w", err)
	}
	b.PutInt(f.ID)
	b.PutString(f.Title)
	if f.Flags.Has(3) {
		if f.Photo == nil {
			return fmt.Errorf("unable to encode folder#ff544e65: field photo is nil")
		}
		if err := f.Photo.Encode(b); err != nil {
			return fmt.Errorf("unable to encode folder#ff544e65: field photo: %w", err)
		}
	}
	return nil
}

// SetAutofillNewBroadcasts sets value of AutofillNewBroadcasts conditional field.
func (f *Folder) SetAutofillNewBroadcasts(value bool) {
	if value {
		f.Flags.Set(0)
	} else {
		f.Flags.Unset(0)
	}
}

// SetAutofillPublicGroups sets value of AutofillPublicGroups conditional field.
func (f *Folder) SetAutofillPublicGroups(value bool) {
	if value {
		f.Flags.Set(1)
	} else {
		f.Flags.Unset(1)
	}
}

// SetAutofillNewCorrespondents sets value of AutofillNewCorrespondents conditional field.
func (f *Folder) SetAutofillNewCorrespondents(value bool) {
	if value {
		f.Flags.Set(2)
	} else {
		f.Flags.Unset(2)
	}
}

// SetPhoto sets value of Photo conditional field.
func (f *Folder) SetPhoto(value ChatPhotoClass) {
	f.Flags.Set(3)
	f.Photo = value
}

// GetPhoto returns value of Photo conditional field and
// boolean which is true if field was set.
func (f *Folder) GetPhoto() (value ChatPhotoClass, ok bool) {
	if !f.Flags.Has(3) {
		return value, false
	}
	return f.Photo, true
}

// Decode implements bin.Decoder.
func (f *Folder) Decode(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't decode folder#ff544e65 to nil")
	}
	if err := b.ConsumeID(FolderTypeID); err != nil {
		return fmt.Errorf("unable to decode folder#ff544e65: %w", err)
	}
	{
		if err := f.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode folder#ff544e65: field flags: %w", err)
		}
	}
	f.AutofillNewBroadcasts = f.Flags.Has(0)
	f.AutofillPublicGroups = f.Flags.Has(1)
	f.AutofillNewCorrespondents = f.Flags.Has(2)
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode folder#ff544e65: field id: %w", err)
		}
		f.ID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode folder#ff544e65: field title: %w", err)
		}
		f.Title = value
	}
	if f.Flags.Has(3) {
		value, err := DecodeChatPhoto(b)
		if err != nil {
			return fmt.Errorf("unable to decode folder#ff544e65: field photo: %w", err)
		}
		f.Photo = value
	}
	return nil
}

// Ensuring interfaces in compile-time for Folder.
var (
	_ bin.Encoder = &Folder{}
	_ bin.Decoder = &Folder{}
)
