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

// MsgResendReq represents TL type `msg_resend_req#7d861a08`.
type MsgResendReq struct {
	// MsgIds field of MsgResendReq.
	MsgIds []int64
}

// MsgResendReqTypeID is TL type id of MsgResendReq.
const MsgResendReqTypeID = 0x7d861a08

// String implements fmt.Stringer.
func (m *MsgResendReq) String() string {
	if m == nil {
		return "MsgResendReq(nil)"
	}
	var sb strings.Builder
	sb.WriteString("MsgResendReq")
	sb.WriteString("{\n")
	sb.WriteByte('[')
	for _, v := range m.MsgIds {
		sb.WriteString(fmt.Sprint(v))
	}
	sb.WriteByte(']')
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (m *MsgResendReq) Encode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode msg_resend_req#7d861a08 as nil")
	}
	b.PutID(MsgResendReqTypeID)
	b.PutVectorHeader(len(m.MsgIds))
	for _, v := range m.MsgIds {
		b.PutLong(v)
	}
	return nil
}

// Decode implements bin.Decoder.
func (m *MsgResendReq) Decode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode msg_resend_req#7d861a08 to nil")
	}
	if err := b.ConsumeID(MsgResendReqTypeID); err != nil {
		return fmt.Errorf("unable to decode msg_resend_req#7d861a08: %w", err)
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode msg_resend_req#7d861a08: field msg_ids: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Long()
			if err != nil {
				return fmt.Errorf("unable to decode msg_resend_req#7d861a08: field msg_ids: %w", err)
			}
			m.MsgIds = append(m.MsgIds, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for MsgResendReq.
var (
	_ bin.Encoder = &MsgResendReq{}
	_ bin.Decoder = &MsgResendReq{}
)
