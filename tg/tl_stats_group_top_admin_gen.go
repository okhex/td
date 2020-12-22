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

// StatsGroupTopAdmin represents TL type `statsGroupTopAdmin#6014f412`.
// Information about an active admin in a supergroup
//
// See https://core.telegram.org/constructor/statsGroupTopAdmin for reference.
type StatsGroupTopAdmin struct {
	// User ID
	UserID int
	// Number of deleted messages for statistics¹ period in consideration
	//
	// Links:
	//  1) https://core.telegram.org/api/stats
	Deleted int
	// Number of kicked users for statistics¹ period in consideration
	//
	// Links:
	//  1) https://core.telegram.org/api/stats
	Kicked int
	// Number of banned users for statistics¹ period in consideration
	//
	// Links:
	//  1) https://core.telegram.org/api/stats
	Banned int
}

// StatsGroupTopAdminTypeID is TL type id of StatsGroupTopAdmin.
const StatsGroupTopAdminTypeID = 0x6014f412

// String implements fmt.Stringer.
func (s *StatsGroupTopAdmin) String() string {
	if s == nil {
		return "StatsGroupTopAdmin(nil)"
	}
	var sb strings.Builder
	sb.WriteString("StatsGroupTopAdmin")
	sb.WriteString("{\n")
	sb.WriteString("\tUserID: ")
	sb.WriteString(fmt.Sprint(s.UserID))
	sb.WriteString(",\n")
	sb.WriteString("\tDeleted: ")
	sb.WriteString(fmt.Sprint(s.Deleted))
	sb.WriteString(",\n")
	sb.WriteString("\tKicked: ")
	sb.WriteString(fmt.Sprint(s.Kicked))
	sb.WriteString(",\n")
	sb.WriteString("\tBanned: ")
	sb.WriteString(fmt.Sprint(s.Banned))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (s *StatsGroupTopAdmin) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode statsGroupTopAdmin#6014f412 as nil")
	}
	b.PutID(StatsGroupTopAdminTypeID)
	b.PutInt(s.UserID)
	b.PutInt(s.Deleted)
	b.PutInt(s.Kicked)
	b.PutInt(s.Banned)
	return nil
}

// Decode implements bin.Decoder.
func (s *StatsGroupTopAdmin) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode statsGroupTopAdmin#6014f412 to nil")
	}
	if err := b.ConsumeID(StatsGroupTopAdminTypeID); err != nil {
		return fmt.Errorf("unable to decode statsGroupTopAdmin#6014f412: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode statsGroupTopAdmin#6014f412: field user_id: %w", err)
		}
		s.UserID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode statsGroupTopAdmin#6014f412: field deleted: %w", err)
		}
		s.Deleted = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode statsGroupTopAdmin#6014f412: field kicked: %w", err)
		}
		s.Kicked = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode statsGroupTopAdmin#6014f412: field banned: %w", err)
		}
		s.Banned = value
	}
	return nil
}

// Ensuring interfaces in compile-time for StatsGroupTopAdmin.
var (
	_ bin.Encoder = &StatsGroupTopAdmin{}
	_ bin.Decoder = &StatsGroupTopAdmin{}
)
