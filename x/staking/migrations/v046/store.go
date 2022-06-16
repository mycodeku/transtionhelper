package v046

import (
	"github.com/mycodeku/transtionhelper/codec"
	storetypes "github.com/mycodeku/transtionhelper/store/types"
	sdk "github.com/mycodeku/transtionhelper/types"
	paramtypes "github.com/mycodeku/transtionhelper/x/params/types"
	"github.com/mycodeku/transtionhelper/x/staking/types"
)

// MigrateStore performs in-place store migrations from v0.43/v0.44 to v0.45.
// The migration includes:
//
// - Setting the MinCommissionRate param in the paramstore
func MigrateStore(ctx sdk.Context, storeKey storetypes.StoreKey, cdc codec.BinaryCodec, paramstore paramtypes.Subspace) error {
	migrateParamsStore(ctx, paramstore)

	return nil
}

func migrateParamsStore(ctx sdk.Context, paramstore paramtypes.Subspace) {
	paramstore.WithKeyTable(types.ParamKeyTable())
	paramstore.Set(ctx, types.KeyMinCommissionRate, types.DefaultMinCommissionRate)
}
