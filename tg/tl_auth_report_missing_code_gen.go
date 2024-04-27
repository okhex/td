// Code generated by gotdgen, DO NOT EDIT.

package tg

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

// AuthReportMissingCodeRequest represents TL type `auth.reportMissingCode#cb9deff6`.
//
// See https://core.telegram.org/method/auth.reportMissingCode for reference.
type AuthReportMissingCodeRequest struct {
	// PhoneNumber field of AuthReportMissingCodeRequest.
	PhoneNumber string
	// PhoneCodeHash field of AuthReportMissingCodeRequest.
	PhoneCodeHash string
	// Mnc field of AuthReportMissingCodeRequest.
	Mnc string
}

// AuthReportMissingCodeRequestTypeID is TL type id of AuthReportMissingCodeRequest.
const AuthReportMissingCodeRequestTypeID = 0xcb9deff6

// Ensuring interfaces in compile-time for AuthReportMissingCodeRequest.
var (
	_ bin.Encoder     = &AuthReportMissingCodeRequest{}
	_ bin.Decoder     = &AuthReportMissingCodeRequest{}
	_ bin.BareEncoder = &AuthReportMissingCodeRequest{}
	_ bin.BareDecoder = &AuthReportMissingCodeRequest{}
)

func (r *AuthReportMissingCodeRequest) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.PhoneNumber == "") {
		return false
	}
	if !(r.PhoneCodeHash == "") {
		return false
	}
	if !(r.Mnc == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *AuthReportMissingCodeRequest) String() string {
	if r == nil {
		return "AuthReportMissingCodeRequest(nil)"
	}
	type Alias AuthReportMissingCodeRequest
	return fmt.Sprintf("AuthReportMissingCodeRequest%+v", Alias(*r))
}

// FillFrom fills AuthReportMissingCodeRequest from given interface.
func (r *AuthReportMissingCodeRequest) FillFrom(from interface {
	GetPhoneNumber() (value string)
	GetPhoneCodeHash() (value string)
	GetMnc() (value string)
}) {
	r.PhoneNumber = from.GetPhoneNumber()
	r.PhoneCodeHash = from.GetPhoneCodeHash()
	r.Mnc = from.GetMnc()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AuthReportMissingCodeRequest) TypeID() uint32 {
	return AuthReportMissingCodeRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*AuthReportMissingCodeRequest) TypeName() string {
	return "auth.reportMissingCode"
}

// TypeInfo returns info about TL type.
func (r *AuthReportMissingCodeRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "auth.reportMissingCode",
		ID:   AuthReportMissingCodeRequestTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "PhoneNumber",
			SchemaName: "phone_number",
		},
		{
			Name:       "PhoneCodeHash",
			SchemaName: "phone_code_hash",
		},
		{
			Name:       "Mnc",
			SchemaName: "mnc",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (r *AuthReportMissingCodeRequest) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode auth.reportMissingCode#cb9deff6 as nil")
	}
	b.PutID(AuthReportMissingCodeRequestTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *AuthReportMissingCodeRequest) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode auth.reportMissingCode#cb9deff6 as nil")
	}
	b.PutString(r.PhoneNumber)
	b.PutString(r.PhoneCodeHash)
	b.PutString(r.Mnc)
	return nil
}

// Decode implements bin.Decoder.
func (r *AuthReportMissingCodeRequest) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode auth.reportMissingCode#cb9deff6 to nil")
	}
	if err := b.ConsumeID(AuthReportMissingCodeRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode auth.reportMissingCode#cb9deff6: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *AuthReportMissingCodeRequest) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode auth.reportMissingCode#cb9deff6 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode auth.reportMissingCode#cb9deff6: field phone_number: %w", err)
		}
		r.PhoneNumber = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode auth.reportMissingCode#cb9deff6: field phone_code_hash: %w", err)
		}
		r.PhoneCodeHash = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode auth.reportMissingCode#cb9deff6: field mnc: %w", err)
		}
		r.Mnc = value
	}
	return nil
}

// GetPhoneNumber returns value of PhoneNumber field.
func (r *AuthReportMissingCodeRequest) GetPhoneNumber() (value string) {
	if r == nil {
		return
	}
	return r.PhoneNumber
}

// GetPhoneCodeHash returns value of PhoneCodeHash field.
func (r *AuthReportMissingCodeRequest) GetPhoneCodeHash() (value string) {
	if r == nil {
		return
	}
	return r.PhoneCodeHash
}

// GetMnc returns value of Mnc field.
func (r *AuthReportMissingCodeRequest) GetMnc() (value string) {
	if r == nil {
		return
	}
	return r.Mnc
}

// AuthReportMissingCode invokes method auth.reportMissingCode#cb9deff6 returning error if any.
//
// See https://core.telegram.org/method/auth.reportMissingCode for reference.
func (c *Client) AuthReportMissingCode(ctx context.Context, request *AuthReportMissingCodeRequest) (bool, error) {
	var result BoolBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
