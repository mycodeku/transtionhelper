package types

import (
	"github.com/mycodeku/transtionhelper/codec"
	cryptocodec "github.com/mycodeku/transtionhelper/crypto/codec"
)

var (
	amino = codec.NewLegacyAmino()
)

func init() {
	cryptocodec.RegisterCrypto(amino)
	amino.Seal()
}
