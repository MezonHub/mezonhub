package ledger_test

import (
	"testing"

	keepertest "github.com/mezonhub/mezonhub/testutil/keeper"
	"github.com/mezonhub/mezonhub/x/ledger"
	"github.com/mezonhub/mezonhub/x/ledger/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.LedgerKeeper(t)
	ledger.InitGenesis(ctx, k, genesisState)
	got := ledger.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	// this line is used by starport scaffolding # genesis/test/assert
}
