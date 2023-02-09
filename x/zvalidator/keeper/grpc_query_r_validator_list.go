package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/mezonhub/mezonhub/x/zvalidator/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ZValidatorList(goCtx context.Context, req *types.QueryZValidatorListRequest) (*types.QueryZValidatorListResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	list := make([]string, 0)
	for _, val := range k.GetSelectedZValidatorListByDenomPoolAddress(ctx, req.Denom, req.PoolAddress) {
		list = append(list, val.ValAddress)
	}

	return &types.QueryZValidatorListResponse{
		ZValidatorList: list,
	}, nil
}
