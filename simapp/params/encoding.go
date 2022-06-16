package params

import (
	"github.com/mycodeku/transtionhelper/client"
	"github.com/mycodeku/transtionhelper/codec"
	"github.com/mycodeku/transtionhelper/codec/types"
)

// EncodingConfig specifies the concrete encoding types to use for a given app.
// This is provided for compatibility between protobuf and amino implementations.
type EncodingConfig struct {
	InterfaceRegistry types.InterfaceRegistry
	Codec             codec.Codec
	TxConfig          client.TxConfig
	Amino             *codec.LegacyAmino
}
