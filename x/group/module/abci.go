package module

import (
	sdk "github.com/mycodeku/transtionhelper/types"
	"github.com/mycodeku/transtionhelper/x/group/keeper"
)

func EndBlocker(ctx sdk.Context, k keeper.Keeper) {
	if err := k.UpdateTallyOfVPEndProposals(ctx); err != nil {
		panic(err)
	}
	pruneProposals(ctx, k)
}

func pruneProposals(ctx sdk.Context, k keeper.Keeper) {
	err := k.PruneProposals(ctx)
	if err != nil {
		panic(err)
	}
}
