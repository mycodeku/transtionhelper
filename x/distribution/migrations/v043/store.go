package v043

import (
	storetypes "github.com/mycodeku/transtionhelper/store/types"
	sdk "github.com/mycodeku/transtionhelper/types"
	v040distribution "github.com/mycodeku/transtionhelper/x/distribution/migrations/v040"
)

// MigrateStore performs in-place store migrations from v0.40 to v0.43. The
// migration includes:
//
// - Change addresses to be length-prefixed.
func MigrateStore(ctx sdk.Context, storeKey storetypes.StoreKey) error {
	store := ctx.KVStore(storeKey)
	MigratePrefixAddress(store, v040distribution.ValidatorOutstandingRewardsPrefix)
	MigratePrefixAddress(store, v040distribution.DelegatorWithdrawAddrPrefix)
	MigratePrefixAddressAddress(store, v040distribution.DelegatorStartingInfoPrefix)
	MigratePrefixAddressBytes(store, v040distribution.ValidatorHistoricalRewardsPrefix)
	MigratePrefixAddress(store, v040distribution.ValidatorCurrentRewardsPrefix)
	MigratePrefixAddress(store, v040distribution.ValidatorAccumulatedCommissionPrefix)
	MigratePrefixAddressBytes(store, v040distribution.ValidatorSlashEventPrefix)

	return nil
}
