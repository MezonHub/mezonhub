package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/mezonhub/mezonhub/x/zvalidator/types"
	sudoTypes "github.com/mezonhub/mezonhub/x/sudo/types"
)

func (k msgServer) AddZValidator(goCtx context.Context, msg *types.MsgAddZValidator) (*types.MsgAddZValidatorResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	isAdmin := k.sudoKeeper.IsAdmin(ctx, msg.Creator)
	if !isAdmin {
		return nil, sudoTypes.ErrCreatorNotAdmin
	}

	if err := k.ZBankKeeper.CheckValAddress(ctx, msg.Denom, msg.ValAddress); err != nil {
		return nil, err
	}

	if err := k.ZBankKeeper.CheckAccAddress(ctx, msg.Denom, msg.PoolAddress); err != nil {
		return nil, err
	}

	zValidator := types.ZValidator{
		Denom:       msg.Denom,
		PoolAddress: msg.PoolAddress,
		ValAddress:  msg.ValAddress,
	}

	if k.Keeper.HasSelectedZValidator(ctx, &zValidator) {
		return nil, types.ErrZValidatorAlreadyExist
	}

	snapShots := k.ledgerKeeper.CurrentEraSnapshots(ctx, msg.Denom)
	if len(snapShots.ShotIds) > 0 {
		return nil, types.ErrLedgerIsBusyWithEra
	}

	chainEra, found := k.ledgerKeeper.GetChainEra(ctx, msg.Denom)
	if !found {
		return nil, types.ErrLedgerChainEraNotExist
	}

	k.Keeper.AddSelectedZValidator(ctx, &zValidator)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeAddZValidator,
			sdk.NewAttribute(types.AttributeKeyDenom, msg.Denom),
			sdk.NewAttribute(types.AttributeKeyPoolAddress, msg.PoolAddress),
			sdk.NewAttribute(types.AttributeKeyChainEra, fmt.Sprintf("%d", chainEra.Era)),
			sdk.NewAttribute(types.AttributeKeyAddedAddress, msg.ValAddress),
		),
	)
	return &types.MsgAddZValidatorResponse{}, nil
}
