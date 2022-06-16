package authz

import (
	sdk "github.com/mycodeku/transtionhelper/types"
	"github.com/mycodeku/transtionhelper/x/authz/keeper"
)

// BeginBlocker is called at the begining of every block
func BeginBlocker(ctx sdk.Context, keeper keeper.Keeper) {

	// delete all the mature grants
	if err := keeper.DequeueAndDeleteExpiredGrants(ctx); err != nil {
		panic(err)
	}
}
