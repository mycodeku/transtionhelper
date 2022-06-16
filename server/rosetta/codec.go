package rosetta

import (
	"github.com/mycodeku/transtionhelper/codec"
	codectypes "github.com/mycodeku/transtionhelper/codec/types"
	cryptocodec "github.com/mycodeku/transtionhelper/crypto/codec"
	authcodec "github.com/mycodeku/transtionhelper/x/auth/types"
	bankcodec "github.com/mycodeku/transtionhelper/x/bank/types"
)

// MakeCodec generates the codec required to interact
// with the cosmos APIs used by the rosetta gateway
func MakeCodec() (*codec.ProtoCodec, codectypes.InterfaceRegistry) {
	ir := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(ir)

	authcodec.RegisterInterfaces(ir)
	bankcodec.RegisterInterfaces(ir)
	cryptocodec.RegisterInterfaces(ir)

	return cdc, ir
}
