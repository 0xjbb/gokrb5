package gssapi

import "github.com/0xjbb/gofork/encoding/asn1"

// GSS-API context flags assigned numbers.
const (
	ContextFlagDeleg    = 1
	ContextFlagMutual   = 2
	ContextFlagReplay   = 4
	ContextFlagSequence = 8
	ContextFlagConf     = 16
	ContextFlagInteg    = 32
	ContextFlagAnon     = 64
)

// ContextFlags flags for GSSAPI
type ContextFlags asn1.BitString

// NewContextFlags creates a new ContextFlags instance.
func NewContextFlags() ContextFlags {
	var c ContextFlags
	c.BitLength = 32
	c.Bytes = make([]byte, 4)
	return c
}
