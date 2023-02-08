package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/mezonhub/mezonhub/x/rdex/types"
	sudotypes "github.com/mezonhub/mezonhub/x/sudo/types"
)

func (k msgServer) ToggleProviderSwitch(goCtx context.Context, msg *types.MsgToggleProviderSwitch) (*types.MsgToggleProviderSwitchResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.sudoKeeper.IsAdmin(ctx, msg.Creator) {
		return nil, sudotypes.ErrCreatorNotAdmin
	}

	k.Keeper.ToggleProviderSwitch(ctx)
	return &types.MsgToggleProviderSwitchResponse{}, nil
}
