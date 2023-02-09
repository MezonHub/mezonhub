package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/mezonhub/mezonhub/x/zvalidator/types"
	sudoTypes "github.com/mezonhub/mezonhub/x/sudo/types"
)

func (k msgServer) RmZValidator(goCtx context.Context, msg *types.MsgRmZValidator) (*types.MsgRmZValidatorResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	isAdmin := k.sudoKeeper.IsAdmin(ctx, msg.Creator)
	if !isAdmin {
		return nil, sudoTypes.ErrCreatorNotAdmin
	}
	newVal := types.ZValidator{
		Denom:       msg.Denom,
		PoolAddress: msg.PoolAddress,
		ValAddress:  msg.NewAddress,
	}
	if !k.Keeper.HasSelectedZValidator(ctx, &newVal) {
		return nil, types.ErrZValidatorNotExist
	}

	latestVotedCycle := k.GetLatestVotedCycle(ctx, msg.Denom, msg.PoolAddress)
	willUseCycle := types.Cycle{
		Denom:       msg.Denom,
		PoolAddress: msg.PoolAddress,
		Version:     latestVotedCycle.Version,
		Number:      latestVotedCycle.Number + 1,
	}

	proposal := types.NewUpdateZValidatorProposal(msg.Creator, msg.Denom, msg.PoolAddress, msg.OldAddress, msg.NewAddress, &willUseCycle)

	err := k.ProcessUpdateZValidatorProposal(ctx, proposal)
	if err != nil {
		return nil, err
	}
	return &types.MsgRmZValidatorResponse{}, nil
}
