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

// EmojiLanguageVector is a box for Vector<EmojiLanguage>
type EmojiLanguageVector struct {
	// Elements of Vector<EmojiLanguage>
	Elems []EmojiLanguage
}

// String implements fmt.Stringer.
func (vec *EmojiLanguageVector) String() string {
	if vec == nil {
		return "EmojiLanguageVector(nil)"
	}
	var sb strings.Builder
	sb.WriteString("EmojiLanguageVector")
	sb.WriteByte('[')
	for _, e := range vec.Elems {
		sb.WriteString(fmt.Sprint(e) + ",\n")
	}
	sb.WriteByte(']')
	return sb.String()
}

// Encode implements bin.Encoder.
func (vec *EmojiLanguageVector) Encode(b *bin.Buffer) error {
	if vec == nil {
		return fmt.Errorf("can't encode Vector<EmojiLanguage> as nil")
	}
	b.PutVectorHeader(len(vec.Elems))
	for idx, v := range vec.Elems {
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode Vector<EmojiLanguage>: field Elems element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (vec *EmojiLanguageVector) Decode(b *bin.Buffer) error {
	if vec == nil {
		return fmt.Errorf("can't decode Vector<EmojiLanguage> to nil")
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode Vector<EmojiLanguage>: field Elems: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value EmojiLanguage
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode Vector<EmojiLanguage>: field Elems: %w", err)
			}
			vec.Elems = append(vec.Elems, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for EmojiLanguageVector.
var (
	_ bin.Encoder = &EmojiLanguageVector{}
	_ bin.Decoder = &EmojiLanguageVector{}
)
