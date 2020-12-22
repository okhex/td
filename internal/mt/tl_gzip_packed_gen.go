// Code generated by gotdgen, DO NOT EDIT.

package mt

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

// GzipPacked represents TL type `gzip_packed#3072cfa1`.
type GzipPacked struct {
	// PackedData field of GzipPacked.
	PackedData []byte
}

// GzipPackedTypeID is TL type id of GzipPacked.
const GzipPackedTypeID = 0x3072cfa1

// String implements fmt.Stringer.
func (g *GzipPacked) String() string {
	if g == nil {
		return "GzipPacked(nil)"
	}
	var sb strings.Builder
	sb.WriteString("GzipPacked")
	sb.WriteString("{\n")
	sb.WriteString("\tPackedData: ")
	sb.WriteString(fmt.Sprint(g.PackedData))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (g *GzipPacked) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode gzip_packed#3072cfa1 as nil")
	}
	b.PutID(GzipPackedTypeID)
	b.PutBytes(g.PackedData)
	return nil
}

// Decode implements bin.Decoder.
func (g *GzipPacked) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode gzip_packed#3072cfa1 to nil")
	}
	if err := b.ConsumeID(GzipPackedTypeID); err != nil {
		return fmt.Errorf("unable to decode gzip_packed#3072cfa1: %w", err)
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode gzip_packed#3072cfa1: field packed_data: %w", err)
		}
		g.PackedData = value
	}
	return nil
}

// Ensuring interfaces in compile-time for GzipPacked.
var (
	_ bin.Encoder = &GzipPacked{}
	_ bin.Decoder = &GzipPacked{}
)
