package distribution

import (
	sdk "github.com/mycodeku/transtionhelper/types"
	sdkerrors "github.com/mycodeku/transtionhelper/types/errors"
	"github.com/mycodeku/transtionhelper/x/distribution/keeper"
	"github.com/mycodeku/transtionhelper/x/distribution/types"
	govtypes "github.com/mycodeku/transtionhelper/x/gov/types/v1beta1"
)

func NewCommunityPoolSpendProposalHandler(k keeper.Keeper) govtypes.Handler {
	return func(ctx sdk.Context, content govtypes.Content) error {
		switch c := content.(type) {
		case *types.CommunityPoolSpendProposal:
			return keeper.HandleCommunityPoolSpendProposal(ctx, k, c)

		default:
			return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized distr proposal content type: %T", c)
		}
	}
}
