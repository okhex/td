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

// PaymentsPaymentForm represents TL type `payments.paymentForm#3f56aea3`.
// Payment form
//
// See https://core.telegram.org/constructor/payments.paymentForm for reference.
type PaymentsPaymentForm struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Whether the user can choose to save credentials.
	CanSaveCredentials bool
	// Indicates that the user can save payment credentials, but only after setting up a 2FA password¹ (currently the account doesn't have a 2FA password²)
	//
	// Links:
	//  1) https://core.telegram.org/api/srp
	//  2) https://core.telegram.org/api/srp
	PasswordMissing bool
	// Bot ID
	BotID int
	// Invoice
	Invoice Invoice
	// Payment provider ID.
	ProviderID int
	// Payment form URL
	URL string
	// Payment provider name.One of the following:- stripe
	//
	// Use SetNativeProvider and GetNativeProvider helpers.
	NativeProvider string
	// Contains information about the payment provider, if available, to support it natively without the need for opening the URL.A JSON object that can contain the following fields:- publishable_key: Stripe API publishable key- apple_pay_merchant_id: Apple Pay merchant ID- android_pay_public_key: Android Pay public key- android_pay_bgcolor: Android Pay form background color- android_pay_inverse: Whether to use the dark theme in the Android Pay form- need_country: True, if the user country must be provided,- need_zip: True, if the user ZIP/postal code must be provided,- need_cardholder_name: True, if the cardholder name must be provided
	//
	// Use SetNativeParams and GetNativeParams helpers.
	NativeParams DataJSON
	// Saved server-side order information
	//
	// Use SetSavedInfo and GetSavedInfo helpers.
	SavedInfo PaymentRequestedInfo
	// Contains information about saved card credentials
	//
	// Use SetSavedCredentials and GetSavedCredentials helpers.
	SavedCredentials PaymentSavedCredentialsCard
	// Users
	Users []UserClass
}

// PaymentsPaymentFormTypeID is TL type id of PaymentsPaymentForm.
const PaymentsPaymentFormTypeID = 0x3f56aea3

// String implements fmt.Stringer.
func (p *PaymentsPaymentForm) String() string {
	if p == nil {
		return "PaymentsPaymentForm(nil)"
	}
	var sb strings.Builder
	sb.WriteString("PaymentsPaymentForm")
	sb.WriteString("{\n")
	sb.WriteString("\tFlags: ")
	sb.WriteString(p.Flags.String())
	sb.WriteString(",\n")
	sb.WriteString("\tBotID: ")
	sb.WriteString(fmt.Sprint(p.BotID))
	sb.WriteString(",\n")
	sb.WriteString("\tInvoice: ")
	sb.WriteString(p.Invoice.String())
	sb.WriteString(",\n")
	sb.WriteString("\tProviderID: ")
	sb.WriteString(fmt.Sprint(p.ProviderID))
	sb.WriteString(",\n")
	sb.WriteString("\tURL: ")
	sb.WriteString(fmt.Sprint(p.URL))
	sb.WriteString(",\n")
	if p.Flags.Has(4) {
		sb.WriteString("\tNativeProvider: ")
		sb.WriteString(fmt.Sprint(p.NativeProvider))
		sb.WriteString(",\n")
	}
	if p.Flags.Has(4) {
		sb.WriteString("\tNativeParams: ")
		sb.WriteString(p.NativeParams.String())
		sb.WriteString(",\n")
	}
	if p.Flags.Has(0) {
		sb.WriteString("\tSavedInfo: ")
		sb.WriteString(p.SavedInfo.String())
		sb.WriteString(",\n")
	}
	if p.Flags.Has(1) {
		sb.WriteString("\tSavedCredentials: ")
		sb.WriteString(p.SavedCredentials.String())
		sb.WriteString(",\n")
	}
	sb.WriteByte('[')
	for _, v := range p.Users {
		sb.WriteString(fmt.Sprint(v))
	}
	sb.WriteByte(']')
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (p *PaymentsPaymentForm) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode payments.paymentForm#3f56aea3 as nil")
	}
	b.PutID(PaymentsPaymentFormTypeID)
	if err := p.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode payments.paymentForm#3f56aea3: field flags: %w", err)
	}
	b.PutInt(p.BotID)
	if err := p.Invoice.Encode(b); err != nil {
		return fmt.Errorf("unable to encode payments.paymentForm#3f56aea3: field invoice: %w", err)
	}
	b.PutInt(p.ProviderID)
	b.PutString(p.URL)
	if p.Flags.Has(4) {
		b.PutString(p.NativeProvider)
	}
	if p.Flags.Has(4) {
		if err := p.NativeParams.Encode(b); err != nil {
			return fmt.Errorf("unable to encode payments.paymentForm#3f56aea3: field native_params: %w", err)
		}
	}
	if p.Flags.Has(0) {
		if err := p.SavedInfo.Encode(b); err != nil {
			return fmt.Errorf("unable to encode payments.paymentForm#3f56aea3: field saved_info: %w", err)
		}
	}
	if p.Flags.Has(1) {
		if err := p.SavedCredentials.Encode(b); err != nil {
			return fmt.Errorf("unable to encode payments.paymentForm#3f56aea3: field saved_credentials: %w", err)
		}
	}
	b.PutVectorHeader(len(p.Users))
	for idx, v := range p.Users {
		if v == nil {
			return fmt.Errorf("unable to encode payments.paymentForm#3f56aea3: field users element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode payments.paymentForm#3f56aea3: field users element with index %d: %w", idx, err)
		}
	}
	return nil
}

// SetCanSaveCredentials sets value of CanSaveCredentials conditional field.
func (p *PaymentsPaymentForm) SetCanSaveCredentials(value bool) {
	if value {
		p.Flags.Set(2)
	} else {
		p.Flags.Unset(2)
	}
}

// SetPasswordMissing sets value of PasswordMissing conditional field.
func (p *PaymentsPaymentForm) SetPasswordMissing(value bool) {
	if value {
		p.Flags.Set(3)
	} else {
		p.Flags.Unset(3)
	}
}

// SetNativeProvider sets value of NativeProvider conditional field.
func (p *PaymentsPaymentForm) SetNativeProvider(value string) {
	p.Flags.Set(4)
	p.NativeProvider = value
}

// GetNativeProvider returns value of NativeProvider conditional field and
// boolean which is true if field was set.
func (p *PaymentsPaymentForm) GetNativeProvider() (value string, ok bool) {
	if !p.Flags.Has(4) {
		return value, false
	}
	return p.NativeProvider, true
}

// SetNativeParams sets value of NativeParams conditional field.
func (p *PaymentsPaymentForm) SetNativeParams(value DataJSON) {
	p.Flags.Set(4)
	p.NativeParams = value
}

// GetNativeParams returns value of NativeParams conditional field and
// boolean which is true if field was set.
func (p *PaymentsPaymentForm) GetNativeParams() (value DataJSON, ok bool) {
	if !p.Flags.Has(4) {
		return value, false
	}
	return p.NativeParams, true
}

// SetSavedInfo sets value of SavedInfo conditional field.
func (p *PaymentsPaymentForm) SetSavedInfo(value PaymentRequestedInfo) {
	p.Flags.Set(0)
	p.SavedInfo = value
}

// GetSavedInfo returns value of SavedInfo conditional field and
// boolean which is true if field was set.
func (p *PaymentsPaymentForm) GetSavedInfo() (value PaymentRequestedInfo, ok bool) {
	if !p.Flags.Has(0) {
		return value, false
	}
	return p.SavedInfo, true
}

// SetSavedCredentials sets value of SavedCredentials conditional field.
func (p *PaymentsPaymentForm) SetSavedCredentials(value PaymentSavedCredentialsCard) {
	p.Flags.Set(1)
	p.SavedCredentials = value
}

// GetSavedCredentials returns value of SavedCredentials conditional field and
// boolean which is true if field was set.
func (p *PaymentsPaymentForm) GetSavedCredentials() (value PaymentSavedCredentialsCard, ok bool) {
	if !p.Flags.Has(1) {
		return value, false
	}
	return p.SavedCredentials, true
}

// Decode implements bin.Decoder.
func (p *PaymentsPaymentForm) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode payments.paymentForm#3f56aea3 to nil")
	}
	if err := b.ConsumeID(PaymentsPaymentFormTypeID); err != nil {
		return fmt.Errorf("unable to decode payments.paymentForm#3f56aea3: %w", err)
	}
	{
		if err := p.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode payments.paymentForm#3f56aea3: field flags: %w", err)
		}
	}
	p.CanSaveCredentials = p.Flags.Has(2)
	p.PasswordMissing = p.Flags.Has(3)
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode payments.paymentForm#3f56aea3: field bot_id: %w", err)
		}
		p.BotID = value
	}
	{
		if err := p.Invoice.Decode(b); err != nil {
			return fmt.Errorf("unable to decode payments.paymentForm#3f56aea3: field invoice: %w", err)
		}
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode payments.paymentForm#3f56aea3: field provider_id: %w", err)
		}
		p.ProviderID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode payments.paymentForm#3f56aea3: field url: %w", err)
		}
		p.URL = value
	}
	if p.Flags.Has(4) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode payments.paymentForm#3f56aea3: field native_provider: %w", err)
		}
		p.NativeProvider = value
	}
	if p.Flags.Has(4) {
		if err := p.NativeParams.Decode(b); err != nil {
			return fmt.Errorf("unable to decode payments.paymentForm#3f56aea3: field native_params: %w", err)
		}
	}
	if p.Flags.Has(0) {
		if err := p.SavedInfo.Decode(b); err != nil {
			return fmt.Errorf("unable to decode payments.paymentForm#3f56aea3: field saved_info: %w", err)
		}
	}
	if p.Flags.Has(1) {
		if err := p.SavedCredentials.Decode(b); err != nil {
			return fmt.Errorf("unable to decode payments.paymentForm#3f56aea3: field saved_credentials: %w", err)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode payments.paymentForm#3f56aea3: field users: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeUser(b)
			if err != nil {
				return fmt.Errorf("unable to decode payments.paymentForm#3f56aea3: field users: %w", err)
			}
			p.Users = append(p.Users, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for PaymentsPaymentForm.
var (
	_ bin.Encoder = &PaymentsPaymentForm{}
	_ bin.Decoder = &PaymentsPaymentForm{}
)
