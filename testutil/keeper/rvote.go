package keeper

import (
	"sync"
	"testing"

	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	ledgermodule "github.com/mezonhub/mezonhub/x/ledger"
	ledgertypes "github.com/mezonhub/mezonhub/x/ledger/types"
	"github.com/mezonhub/mezonhub/x/zvote/keeper"
	"github.com/mezonhub/mezonhub/x/zvote/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
)

var (
	zvoteStoreKey    = sdk.NewKVStoreKey(types.StoreKey)
	zvoteMemStoreKey = storetypes.NewMemoryStoreKey(types.MemStoreKey)
	zvoteOnce        sync.Once
)

func RvoteKeeper(t testing.TB) (*keeper.Keeper, sdk.Context) {
	zvoteOnce.Do(func() {
		stateStore.MountStoreWithDB(zvoteStoreKey, sdk.StoreTypeIAVL, db)
		stateStore.MountStoreWithDB(zvoteMemStoreKey, sdk.StoreTypeMemory, nil)
	})

	sudoKeeper, _ := SudoKeeper(t)
	relayersKeeper, _ := RelayersKeeper(t)
	ledgerKeeper, _ := LedgerKeeper(t)
	require.NoError(t, stateStore.LoadLatestVersion())

	zvoteRouter := types.NewRouter()
	zvoteRouter.AddRoute(ledgertypes.RouterKey, ledgermodule.NewProposalHandler(ledgerKeeper))
	zvoteKeeper := keeper.NewKeeper(
		cdc,
		zvoteStoreKey,
		zvoteMemStoreKey,

		sudoKeeper,
		relayersKeeper,
		zvoteRouter,
	)

	ctx := sdk.NewContext(stateStore, tmproto.Header{}, false, log.NewNopLogger())
	return zvoteKeeper, ctx
}
