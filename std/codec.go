package std

import (
	"github.com/mycodeku/transtionhelper/codec"
	"github.com/mycodeku/transtionhelper/codec/types"
	cryptocodec "github.com/mycodeku/transtionhelper/crypto/codec"
	sdk "github.com/mycodeku/transtionhelper/types"
	txtypes "github.com/mycodeku/transtionhelper/types/tx"
)

// RegisterLegacyAminoCodec registers types with the Amino codec.
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	sdk.RegisterLegacyAminoCodec(cdc)
	cryptocodec.RegisterCrypto(cdc)
	codec.RegisterEvidences(cdc)
}

// RegisterInterfaces registers Interfaces from sdk/types, vesting, crypto, tx.
func RegisterInterfaces(interfaceRegistry types.InterfaceRegistry) {
	sdk.RegisterInterfaces(interfaceRegistry)
	txtypes.RegisterInterfaces(interfaceRegistry)
	cryptocodec.RegisterInterfaces(interfaceRegistry)
}
