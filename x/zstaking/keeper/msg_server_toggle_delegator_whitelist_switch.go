package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/mezonhub/mezonhub/x/zstaking/types"
	sudotypes "github.com/mezonhub/mezonhub/x/sudo/types"
)

func (k msgServer) ToggleDelegatorWhitelistSwitch(goCtx context.Context, msg *types.MsgToggleDelegatorWhitelistSwitch) (*types.MsgToggleDelegatorWhitelistSwitchResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	isAdmin := k.sudoKeeper.IsAdmin(ctx, msg.Creator)
	if !isAdmin {
		return nil, sudotypes.ErrCreatorNotAdmin
	}
	k.Keeper.ToggleDelegatorWhitelistSwitch(ctx)

	return &types.MsgToggleDelegatorWhitelistSwitchResponse{}, nil
}
