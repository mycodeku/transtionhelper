package v043

import (
	storetypes "github.com/mycodeku/transtionhelper/store/types"
	sdk "github.com/mycodeku/transtionhelper/types"
	v043distribution "github.com/mycodeku/transtionhelper/x/distribution/migrations/v043"
	v040slashing "github.com/mycodeku/transtionhelper/x/slashing/migrations/v040"
)

// MigrateStore performs in-place store migrations from v0.40 to v0.43. The
// migration includes:
//
// - Change addresses to be length-prefixed.
func MigrateStore(ctx sdk.Context, storeKey storetypes.StoreKey) error {
	store := ctx.KVStore(storeKey)
	v043distribution.MigratePrefixAddress(store, v040slashing.ValidatorSigningInfoKeyPrefix)
	v043distribution.MigratePrefixAddressBytes(store, v040slashing.ValidatorMissedBlockBitArrayKeyPrefix)
	v043distribution.MigratePrefixAddress(store, v040slashing.AddrPubkeyRelationKeyPrefix)

	return nil
}
