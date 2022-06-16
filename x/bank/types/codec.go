package types

import (
	"github.com/mycodeku/transtionhelper/codec"
	"github.com/mycodeku/transtionhelper/codec/legacy"
	"github.com/mycodeku/transtionhelper/codec/types"
	sdk "github.com/mycodeku/transtionhelper/types"
	"github.com/mycodeku/transtionhelper/types/msgservice"
	"github.com/mycodeku/transtionhelper/x/authz"
)

// RegisterLegacyAminoCodec registers the necessary x/bank interfaces and concrete types
// on the provided LegacyAmino codec. These types are used for Amino JSON serialization.
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	legacy.RegisterAminoMsg(cdc, &MsgSend{}, "cosmos-sdk/MsgSend")
	legacy.RegisterAminoMsg(cdc, &MsgMultiSend{}, "cosmos-sdk/MsgMultiSend")
	cdc.RegisterConcrete(&SendAuthorization{}, "cosmos-sdk/SendAuthorization", nil)
}

func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSend{},
		&MsgMultiSend{},
	)
	registry.RegisterImplementations(
		(*authz.Authorization)(nil),
		&SendAuthorization{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

func init() {
	RegisterLegacyAminoCodec(legacy.Cdc)
}
