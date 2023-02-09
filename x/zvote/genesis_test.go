package zvote_test

import (
	"testing"

	keepertest "github.com/mezonhub/mezonhub/testutil/keeper"
	"github.com/mezonhub/mezonhub/x/zvote"
	"github.com/mezonhub/mezonhub/x/zvote/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.RvoteKeeper(t)
	zvote.InitGenesis(ctx, *k, genesisState)
	got := zvote.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	// this line is used by starport scaffolding # genesis/test/assert
}
