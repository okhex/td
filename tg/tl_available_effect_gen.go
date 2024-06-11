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

// AvailableEffect represents TL type `availableEffect#93c3e27e`.
//
// See https://core.telegram.org/constructor/availableEffect for reference.
type AvailableEffect struct {
	// Flags field of AvailableEffect.
	Flags bin.Fields
	// PremiumRequired field of AvailableEffect.
	PremiumRequired bool
	// ID field of AvailableEffect.
	ID int64
	// Emoticon field of AvailableEffect.
	Emoticon string
	// StaticIconID field of AvailableEffect.
	//
	// Use SetStaticIconID and GetStaticIconID helpers.
	StaticIconID int64
	// EffectStickerID field of AvailableEffect.
	EffectStickerID int64
	// EffectAnimationID field of AvailableEffect.
	//
	// Use SetEffectAnimationID and GetEffectAnimationID helpers.
	EffectAnimationID int64
}

// AvailableEffectTypeID is TL type id of AvailableEffect.
const AvailableEffectTypeID = 0x93c3e27e

// Ensuring interfaces in compile-time for AvailableEffect.
var (
	_ bin.Encoder     = &AvailableEffect{}
	_ bin.Decoder     = &AvailableEffect{}
	_ bin.BareEncoder = &AvailableEffect{}
	_ bin.BareDecoder = &AvailableEffect{}
)

func (a *AvailableEffect) Zero() bool {
	if a == nil {
		return true
	}
	if !(a.Flags.Zero()) {
		return false
	}
	if !(a.PremiumRequired == false) {
		return false
	}
	if !(a.ID == 0) {
		return false
	}
	if !(a.Emoticon == "") {
		return false
	}
	if !(a.StaticIconID == 0) {
		return false
	}
	if !(a.EffectStickerID == 0) {
		return false
	}
	if !(a.EffectAnimationID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (a *AvailableEffect) String() string {
	if a == nil {
		return "AvailableEffect(nil)"
	}
	type Alias AvailableEffect
	return fmt.Sprintf("AvailableEffect%+v", Alias(*a))
}

// FillFrom fills AvailableEffect from given interface.
func (a *AvailableEffect) FillFrom(from interface {
	GetPremiumRequired() (value bool)
	GetID() (value int64)
	GetEmoticon() (value string)
	GetStaticIconID() (value int64, ok bool)
	GetEffectStickerID() (value int64)
	GetEffectAnimationID() (value int64, ok bool)
}) {
	a.PremiumRequired = from.GetPremiumRequired()
	a.ID = from.GetID()
	a.Emoticon = from.GetEmoticon()
	if val, ok := from.GetStaticIconID(); ok {
		a.StaticIconID = val
	}

	a.EffectStickerID = from.GetEffectStickerID()
	if val, ok := from.GetEffectAnimationID(); ok {
		a.EffectAnimationID = val
	}

}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AvailableEffect) TypeID() uint32 {
	return AvailableEffectTypeID
}

// TypeName returns name of type in TL schema.
func (*AvailableEffect) TypeName() string {
	return "availableEffect"
}

// TypeInfo returns info about TL type.
func (a *AvailableEffect) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "availableEffect",
		ID:   AvailableEffectTypeID,
	}
	if a == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "PremiumRequired",
			SchemaName: "premium_required",
			Null:       !a.Flags.Has(2),
		},
		{
			Name:       "ID",
			SchemaName: "id",
		},
		{
			Name:       "Emoticon",
			SchemaName: "emoticon",
		},
		{
			Name:       "StaticIconID",
			SchemaName: "static_icon_id",
			Null:       !a.Flags.Has(0),
		},
		{
			Name:       "EffectStickerID",
			SchemaName: "effect_sticker_id",
		},
		{
			Name:       "EffectAnimationID",
			SchemaName: "effect_animation_id",
			Null:       !a.Flags.Has(1),
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (a *AvailableEffect) SetFlags() {
	if !(a.PremiumRequired == false) {
		a.Flags.Set(2)
	}
	if !(a.StaticIconID == 0) {
		a.Flags.Set(0)
	}
	if !(a.EffectAnimationID == 0) {
		a.Flags.Set(1)
	}
}

// Encode implements bin.Encoder.
func (a *AvailableEffect) Encode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode availableEffect#93c3e27e as nil")
	}
	b.PutID(AvailableEffectTypeID)
	return a.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (a *AvailableEffect) EncodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode availableEffect#93c3e27e as nil")
	}
	a.SetFlags()
	if err := a.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode availableEffect#93c3e27e: field flags: %w", err)
	}
	b.PutLong(a.ID)
	b.PutString(a.Emoticon)
	if a.Flags.Has(0) {
		b.PutLong(a.StaticIconID)
	}
	b.PutLong(a.EffectStickerID)
	if a.Flags.Has(1) {
		b.PutLong(a.EffectAnimationID)
	}
	return nil
}

// Decode implements bin.Decoder.
func (a *AvailableEffect) Decode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode availableEffect#93c3e27e to nil")
	}
	if err := b.ConsumeID(AvailableEffectTypeID); err != nil {
		return fmt.Errorf("unable to decode availableEffect#93c3e27e: %w", err)
	}
	return a.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (a *AvailableEffect) DecodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode availableEffect#93c3e27e to nil")
	}
	{
		if err := a.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode availableEffect#93c3e27e: field flags: %w", err)
		}
	}
	a.PremiumRequired = a.Flags.Has(2)
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode availableEffect#93c3e27e: field id: %w", err)
		}
		a.ID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode availableEffect#93c3e27e: field emoticon: %w", err)
		}
		a.Emoticon = value
	}
	if a.Flags.Has(0) {
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode availableEffect#93c3e27e: field static_icon_id: %w", err)
		}
		a.StaticIconID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode availableEffect#93c3e27e: field effect_sticker_id: %w", err)
		}
		a.EffectStickerID = value
	}
	if a.Flags.Has(1) {
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode availableEffect#93c3e27e: field effect_animation_id: %w", err)
		}
		a.EffectAnimationID = value
	}
	return nil
}

// SetPremiumRequired sets value of PremiumRequired conditional field.
func (a *AvailableEffect) SetPremiumRequired(value bool) {
	if value {
		a.Flags.Set(2)
		a.PremiumRequired = true
	} else {
		a.Flags.Unset(2)
		a.PremiumRequired = false
	}
}

// GetPremiumRequired returns value of PremiumRequired conditional field.
func (a *AvailableEffect) GetPremiumRequired() (value bool) {
	if a == nil {
		return
	}
	return a.Flags.Has(2)
}

// GetID returns value of ID field.
func (a *AvailableEffect) GetID() (value int64) {
	if a == nil {
		return
	}
	return a.ID
}

// GetEmoticon returns value of Emoticon field.
func (a *AvailableEffect) GetEmoticon() (value string) {
	if a == nil {
		return
	}
	return a.Emoticon
}

// SetStaticIconID sets value of StaticIconID conditional field.
func (a *AvailableEffect) SetStaticIconID(value int64) {
	a.Flags.Set(0)
	a.StaticIconID = value
}

// GetStaticIconID returns value of StaticIconID conditional field and
// boolean which is true if field was set.
func (a *AvailableEffect) GetStaticIconID() (value int64, ok bool) {
	if a == nil {
		return
	}
	if !a.Flags.Has(0) {
		return value, false
	}
	return a.StaticIconID, true
}

// GetEffectStickerID returns value of EffectStickerID field.
func (a *AvailableEffect) GetEffectStickerID() (value int64) {
	if a == nil {
		return
	}
	return a.EffectStickerID
}

// SetEffectAnimationID sets value of EffectAnimationID conditional field.
func (a *AvailableEffect) SetEffectAnimationID(value int64) {
	a.Flags.Set(1)
	a.EffectAnimationID = value
}

// GetEffectAnimationID returns value of EffectAnimationID conditional field and
// boolean which is true if field was set.
func (a *AvailableEffect) GetEffectAnimationID() (value int64, ok bool) {
	if a == nil {
		return
	}
	if !a.Flags.Has(1) {
		return value, false
	}
	return a.EffectAnimationID, true
}
