package keeper

import (
	sdk "github.com/mycodeku/transtionhelper/types"
	"github.com/mycodeku/transtionhelper/x/crisis/types"
)

// new crisis genesis
func (k Keeper) InitGenesis(ctx sdk.Context, data *types.GenesisState) {
	k.SetConstantFee(ctx, data.ConstantFee)
}

// ExportGenesis returns a GenesisState for a given context and keeper.
func (k Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	constantFee := k.GetConstantFee(ctx)
	return types.NewGenesisState(constantFee)
}
