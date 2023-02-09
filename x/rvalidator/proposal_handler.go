package rvalidator

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/mezonhub/mezonhub/x/rvalidator/keeper"
	"github.com/mezonhub/mezonhub/x/rvalidator/types"
	rvotetypes "github.com/mezonhub/mezonhub/x/rvote/types"
)

func NewProposalHandler(k keeper.Keeper) rvotetypes.Handler {
	return func(ctx sdk.Context, content rvotetypes.Content) error {
		switch c := content.(type) {
		case *types.UpdateRValidatorProposal:
			return k.ProcessUpdateRValidatorProposal(ctx, c)
		case *types.UpdateRValidatorReportProposal:
			return k.ProcessUpdateRValidatorReportProposal(ctx, c)
		default:
			return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized param proposal content type: %T", c)
		}
	}
}
