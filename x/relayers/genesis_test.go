package relayers_test

import (
	"testing"

	keepertest "github.com/mezonhub/mezonhub/testutil/keeper"
	"github.com/mezonhub/mezonhub/x/relayers"
	"github.com/mezonhub/mezonhub/x/relayers/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.RelayersKeeper(t)
	relayers.InitGenesis(ctx, *k, genesisState)
	got := relayers.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	require.Len(t, got.Relayers, len(genesisState.Relayers))
	require.Subset(t, genesisState.Relayers, got.Relayers)
	require.Len(t, got.Thresholds, len(genesisState.Thresholds))
	require.Subset(t, genesisState.Thresholds, got.Thresholds)
	// this line is used by starport scaffolding # genesis/test/assert
}
