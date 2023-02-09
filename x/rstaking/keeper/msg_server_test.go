package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/mezonhub/mezonhub/testutil/keeper"
	"github.com/mezonhub/mezonhub/x/rstaking/keeper"
	"github.com/mezonhub/mezonhub/x/rstaking/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.RStakingKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
