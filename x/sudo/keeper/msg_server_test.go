package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/mezonhub/mezonhub/testutil/keeper"
	"github.com/mezonhub/mezonhub/x/sudo/keeper"
	"github.com/mezonhub/mezonhub/x/sudo/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.SudoKeeper(t)
	return keeper.NewMsgServerImpl(k), sdk.WrapSDKContext(ctx)
}
