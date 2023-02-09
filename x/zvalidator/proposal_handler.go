package zvalidator

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/mezonhub/mezonhub/x/zvalidator/keeper"
	"github.com/mezonhub/mezonhub/x/zvalidator/types"
	zvotetypes "github.com/mezonhub/mezonhub/x/zvote/types"
)

func NewProposalHandler(k keeper.Keeper) zvotetypes.Handler {
	return func(ctx sdk.Context, content zvotetypes.Content) error {
		switch c := content.(type) {
		case *types.UpdateZValidatorProposal:
			return k.ProcessUpdateZValidatorProposal(ctx, c)
		case *types.UpdateZValidatorReportProposal:
			return k.ProcessUpdateZValidatorReportProposal(ctx, c)
		default:
			return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized param proposal content type: %T", c)
		}
	}
}
