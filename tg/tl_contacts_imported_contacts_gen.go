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

// ContactsImportedContacts represents TL type `contacts.importedContacts#77d01c3b`.
// Info on succesfully imported contacts.
//
// See https://core.telegram.org/constructor/contacts.importedContacts for reference.
type ContactsImportedContacts struct {
	// List of succesfully imported contacts
	Imported []ImportedContact
	// Popular contacts
	PopularInvites []PopularContact
	// List of contact ids that could not be imported due to system limitation and will need to be imported at a later date.Parameter added in Layer 13¹
	//
	// Links:
	//  1) https://core.telegram.org/api/layers#layer-13
	RetryContacts []int64
	// List of users
	Users []UserClass
}

// ContactsImportedContactsTypeID is TL type id of ContactsImportedContacts.
const ContactsImportedContactsTypeID = 0x77d01c3b

// String implements fmt.Stringer.
func (i *ContactsImportedContacts) String() string {
	if i == nil {
		return "ContactsImportedContacts(nil)"
	}
	var sb strings.Builder
	sb.WriteString("ContactsImportedContacts")
	sb.WriteString("{\n")
	sb.WriteByte('[')
	for _, v := range i.Imported {
		sb.WriteString(fmt.Sprint(v))
	}
	sb.WriteByte(']')
	sb.WriteByte('[')
	for _, v := range i.PopularInvites {
		sb.WriteString(fmt.Sprint(v))
	}
	sb.WriteByte(']')
	sb.WriteByte('[')
	for _, v := range i.RetryContacts {
		sb.WriteString(fmt.Sprint(v))
	}
	sb.WriteByte(']')
	sb.WriteByte('[')
	for _, v := range i.Users {
		sb.WriteString(fmt.Sprint(v))
	}
	sb.WriteByte(']')
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (i *ContactsImportedContacts) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode contacts.importedContacts#77d01c3b as nil")
	}
	b.PutID(ContactsImportedContactsTypeID)
	b.PutVectorHeader(len(i.Imported))
	for idx, v := range i.Imported {
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode contacts.importedContacts#77d01c3b: field imported element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(i.PopularInvites))
	for idx, v := range i.PopularInvites {
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode contacts.importedContacts#77d01c3b: field popular_invites element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(i.RetryContacts))
	for _, v := range i.RetryContacts {
		b.PutLong(v)
	}
	b.PutVectorHeader(len(i.Users))
	for idx, v := range i.Users {
		if v == nil {
			return fmt.Errorf("unable to encode contacts.importedContacts#77d01c3b: field users element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode contacts.importedContacts#77d01c3b: field users element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (i *ContactsImportedContacts) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode contacts.importedContacts#77d01c3b to nil")
	}
	if err := b.ConsumeID(ContactsImportedContactsTypeID); err != nil {
		return fmt.Errorf("unable to decode contacts.importedContacts#77d01c3b: %w", err)
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode contacts.importedContacts#77d01c3b: field imported: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value ImportedContact
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode contacts.importedContacts#77d01c3b: field imported: %w", err)
			}
			i.Imported = append(i.Imported, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode contacts.importedContacts#77d01c3b: field popular_invites: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value PopularContact
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode contacts.importedContacts#77d01c3b: field popular_invites: %w", err)
			}
			i.PopularInvites = append(i.PopularInvites, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode contacts.importedContacts#77d01c3b: field retry_contacts: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Long()
			if err != nil {
				return fmt.Errorf("unable to decode contacts.importedContacts#77d01c3b: field retry_contacts: %w", err)
			}
			i.RetryContacts = append(i.RetryContacts, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode contacts.importedContacts#77d01c3b: field users: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeUser(b)
			if err != nil {
				return fmt.Errorf("unable to decode contacts.importedContacts#77d01c3b: field users: %w", err)
			}
			i.Users = append(i.Users, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for ContactsImportedContacts.
var (
	_ bin.Encoder = &ContactsImportedContacts{}
	_ bin.Decoder = &ContactsImportedContacts{}
)
