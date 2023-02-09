package zdex_test

import (
	"testing"

	keepertest "github.com/mezonhub/mezonhub/testutil/keeper"
	"github.com/mezonhub/mezonhub/testutil/nullify"
	"github.com/mezonhub/mezonhub/x/zdex"
	"github.com/mezonhub/mezonhub/x/zdex/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.ZdexKeeper(t)
	zdex.InitGenesis(ctx, *k, genesisState)
	got := zdex.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
