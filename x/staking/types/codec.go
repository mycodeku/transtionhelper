package types

import (
	"github.com/mycodeku/transtionhelper/codec"
	"github.com/mycodeku/transtionhelper/codec/legacy"
	"github.com/mycodeku/transtionhelper/codec/types"
	sdk "github.com/mycodeku/transtionhelper/types"
	"github.com/mycodeku/transtionhelper/types/msgservice"
	"github.com/mycodeku/transtionhelper/x/authz"
)

// RegisterLegacyAminoCodec registers the necessary x/staking interfaces and concrete types
// on the provided LegacyAmino codec. These types are used for Amino JSON serialization.
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	legacy.RegisterAminoMsg(cdc, &MsgCreateValidator{}, "cosmos-sdk/MsgCreateValidator")
	legacy.RegisterAminoMsg(cdc, &MsgEditValidator{}, "cosmos-sdk/MsgEditValidator")
	legacy.RegisterAminoMsg(cdc, &MsgDelegate{}, "cosmos-sdk/MsgDelegate")
	legacy.RegisterAminoMsg(cdc, &MsgUndelegate{}, "cosmos-sdk/MsgUndelegate")
	legacy.RegisterAminoMsg(cdc, &MsgBeginRedelegate{}, "cosmos-sdk/MsgBeginRedelegate")
	legacy.RegisterAminoMsg(cdc, &MsgCancelUnbondingDelegation{}, "cosmos-sdk/MsgCancelUnbondingDelegation")

	cdc.RegisterInterface((*isStakeAuthorization_Validators)(nil), nil)
	cdc.RegisterConcrete(&StakeAuthorization_AllowList{}, "cosmos-sdk/StakeAuthorization/AllowList", nil)
	cdc.RegisterConcrete(&StakeAuthorization_DenyList{}, "cosmos-sdk/StakeAuthorization/DenyList", nil)
	cdc.RegisterConcrete(&StakeAuthorization{}, "cosmos-sdk/StakeAuthorization", nil)
}

// RegisterInterfaces registers the x/staking interfaces types with the interface registry
func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateValidator{},
		&MsgEditValidator{},
		&MsgDelegate{},
		&MsgUndelegate{},
		&MsgBeginRedelegate{},
		&MsgCancelUnbondingDelegation{},
	)
	registry.RegisterImplementations(
		(*authz.Authorization)(nil),
		&StakeAuthorization{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

func init() {
	RegisterLegacyAminoCodec(legacy.Cdc)
}
