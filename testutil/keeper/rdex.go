package keeper

import (
	"sync"
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	typesparams "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/mezonhub/mezonhub/x/zdex/keeper"
	"github.com/mezonhub/mezonhub/x/zdex/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
)

var (
	zdexStoreKey    = sdk.NewKVStoreKey(types.StoreKey)
	zdexMemStoreKey = storetypes.NewMemoryStoreKey(types.MemStoreKey)
	zdexOnce        sync.Once
)

func ZdexKeeper(t testing.TB) (*keeper.Keeper, sdk.Context) {
	zdexOnce.Do(func() {
		stateStore.MountStoreWithDB(zdexStoreKey, sdk.StoreTypeIAVL, db)
		stateStore.MountStoreWithDB(zdexMemStoreKey, sdk.StoreTypeMemory, nil)
	})
	require.NoError(t, stateStore.LoadLatestVersion())

	registry := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(registry)
	paramsSubspace := typesparams.NewSubspace(cdc,
		types.Amino,
		zdexStoreKey,
		zdexMemStoreKey,
		"ZdexParams",
	)
	sudoKeeper, _ := SudoKeeper(t)
	k := keeper.NewKeeper(
		cdc,
		zdexStoreKey,
		zdexMemStoreKey,
		paramsSubspace,
		BankKeeper,
		sudoKeeper,
	)

	ctx := sdk.NewContext(stateStore, tmproto.Header{}, false, log.NewNopLogger())

	// Initialize params
	k.SetParams(ctx, types.DefaultParams())

	return k, ctx
}
