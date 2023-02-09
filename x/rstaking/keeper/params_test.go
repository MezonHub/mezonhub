package keeper_test

import (
	"testing"

	testkeeper "github.com/mezonhub/mezonhub/testutil/keeper"
	"github.com/mezonhub/mezonhub/x/rstaking/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.RStakingKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
