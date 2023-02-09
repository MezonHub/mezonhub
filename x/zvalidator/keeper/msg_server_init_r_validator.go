package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/mezonhub/mezonhub/x/zvalidator/types"
	sudoTypes "github.com/mezonhub/mezonhub/x/sudo/types"
)

// init zvalidator and can only init once
func (k msgServer) InitZValidator(goCtx context.Context, msg *types.MsgInitZValidator) (*types.MsgInitZValidatorResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	isAdmin := k.sudoKeeper.IsAdmin(ctx, msg.Creator)
	if !isAdmin {
		return nil, sudoTypes.ErrCreatorNotAdmin
	}

	zvalidatorList := k.Keeper.GetSelectedZValidatorListByDenomPoolAddress(ctx, msg.Denom, msg.PoolAddress)
	if len(zvalidatorList) > 0 {
		return nil, types.ErrZValidatorAlreadyInit
	}

	if err := k.ZBankKeeper.CheckAccAddress(ctx, msg.Denom, msg.PoolAddress); err != nil {
		return nil, err
	}

	addresses := ""
	for _, address := range msg.ValAddressList {
		if err := k.ZBankKeeper.CheckValAddress(ctx, msg.Denom, address); err != nil {
			return nil, err
		}
		zValidator := types.ZValidator{
			Denom:       msg.Denom,
			PoolAddress: msg.PoolAddress,
			ValAddress:  address,
		}

		if k.Keeper.HasSelectedZValidator(ctx, &zValidator) {
			return nil, types.ErrZValidatorAlreadyExist
		}

		k.Keeper.AddSelectedZValidator(ctx, &zValidator)

		addresses = addresses + ":" + address
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeInitZValidator,
			sdk.NewAttribute(types.AttributeKeyDenom, msg.Denom),
			sdk.NewAttribute(types.AttributeKeyPoolAddress, msg.PoolAddress),
			sdk.NewAttribute(types.AttributeKeyAddresses, addresses[1:]),
		),
	)
	return &types.MsgInitZValidatorResponse{}, nil
}
