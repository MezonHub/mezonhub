package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/mezonhub/mezonhub/x/zvalidator/types"
)

type (
	Keeper struct {
		cdc        codec.BinaryCodec
		storeKey   sdk.StoreKey
		memKey     sdk.StoreKey
		paramstore paramtypes.Subspace

		sudoKeeper   types.SudoKeeper
		ZBankKeeper  types.ZBankKeeper
		ledgerKeeper types.LedgerKeeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey sdk.StoreKey,
	ps paramtypes.Subspace,
	sudoKeeper types.SudoKeeper,
	zBankKeeper types.ZBankKeeper,
	ledgerKeeper types.LedgerKeeper,
) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{

		cdc:          cdc,
		storeKey:     storeKey,
		memKey:       memKey,
		paramstore:   ps,
		sudoKeeper:   sudoKeeper,
		ZBankKeeper:  zBankKeeper,
		ledgerKeeper: ledgerKeeper,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) AddSelectedZValidator(ctx sdk.Context, zValidator *types.ZValidator) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.SelectedRValdidatorStoreKey(zValidator.Denom, zValidator.PoolAddress, zValidator.ValAddress), []byte{})
}

func (k Keeper) RemoveSelectedZValidator(ctx sdk.Context, zValidator *types.ZValidator) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.SelectedRValdidatorStoreKey(zValidator.Denom, zValidator.PoolAddress, zValidator.ValAddress))
}

func (k Keeper) HasSelectedZValidator(ctx sdk.Context, zValidator *types.ZValidator) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.SelectedRValdidatorStoreKey(zValidator.Denom, zValidator.PoolAddress, zValidator.ValAddress))
}

func (k Keeper) GetSelectedZValidatorListByDenomPoolAddress(ctx sdk.Context, denom, poolAddress string) []*types.ZValidator {
	store := ctx.KVStore(k.storeKey)
	denomLen := len([]byte(denom))
	poolAddressLen := len([]byte(poolAddress))

	key := make([]byte, 1+1+denomLen+1+poolAddressLen)
	copy(key[0:], types.SelectedZValidatorStoreKeyPrefix)
	key[1] = byte(denomLen)
	copy(key[2:], []byte(denom))
	key[2+denomLen] = byte(poolAddressLen)
	copy(key[2+denomLen+1:], []byte(poolAddress))

	iterator := sdk.KVStorePrefixIterator(store, key)
	defer iterator.Close()

	list := make([]*types.ZValidator, 0)
	for ; iterator.Valid(); iterator.Next() {
		key := iterator.Key()

		address := string(key[2+denomLen+1+poolAddressLen+1:])

		zValidator := types.ZValidator{
			Denom:       denom,
			PoolAddress: poolAddress,
			ValAddress:  address,
		}

		list = append(list, &zValidator)
	}
	return list
}

func (k Keeper) GetSelectedZValidatorList(ctx sdk.Context) []*types.ZValidator {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.SelectedZValidatorStoreKeyPrefix)
	defer iterator.Close()

	list := make([]*types.ZValidator, 0)
	for ; iterator.Valid(); iterator.Next() {
		key := iterator.Key()

		denomLen := int(key[1])
		denom := string(key[2 : 2+denomLen])
		poolAddressLen := int(key[2+denomLen])
		poolAddress := string(key[2+denomLen+1 : 2+denomLen+1+poolAddressLen])
		valAddress := string(key[2+denomLen+1+poolAddressLen+1:])

		zValidator := types.ZValidator{
			Denom:       denom,
			PoolAddress: poolAddress,
			ValAddress:  valAddress,
		}

		list = append(list, &zValidator)
	}
	return list
}

func (k Keeper) SetLatestVotedCycle(ctx sdk.Context, cycle *types.Cycle) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.LatestVotedCycleStoreKey(cycle.Denom, cycle.PoolAddress), k.cdc.MustMarshal(cycle))
}

func (k Keeper) GetLatestVotedCycle(ctx sdk.Context, denom, poolAddress string) *types.Cycle {
	store := ctx.KVStore(k.storeKey)
	bts := store.Get(types.LatestVotedCycleStoreKey(denom, poolAddress))
	if bts == nil {
		return &types.Cycle{
			Denom:       denom,
			PoolAddress: poolAddress,
			Version:     0,
			Number:      0,
		}
	}
	cycle := types.Cycle{}
	k.cdc.MustUnmarshal(bts, &cycle)

	return &cycle
}

func (k Keeper) GetAllLatestVotedCycle(ctx sdk.Context) []*types.Cycle {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.LatestVotedCycleStoreKeyPrefix)
	defer iterator.Close()

	cycleList := make([]*types.Cycle, 0)
	for ; iterator.Valid(); iterator.Next() {
		cycle := types.Cycle{}
		k.cdc.MustUnmarshal(iterator.Value(), &cycle)
		cycleList = append(cycleList, &cycle)
	}

	return cycleList
}

func (k Keeper) SetLatestDealedCycle(ctx sdk.Context, cycle *types.Cycle) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.LatestDealedCycleStoreKey(cycle.Denom, cycle.PoolAddress), k.cdc.MustMarshal(cycle))
}

func (k Keeper) GetLatestDealedCycle(ctx sdk.Context, denom, poolAddress string) *types.Cycle {
	store := ctx.KVStore(k.storeKey)
	bts := store.Get(types.LatestDealedCycleStoreKey(denom, poolAddress))
	if bts == nil {
		return &types.Cycle{
			Denom:       denom,
			PoolAddress: poolAddress,
			Version:     0,
			Number:      0,
		}
	}
	cycle := types.Cycle{}
	k.cdc.MustUnmarshal(bts, &cycle)

	return &cycle
}

func (k Keeper) GetAllLatestDealedCycle(ctx sdk.Context) []*types.Cycle {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.LatestDealedCycleStoreKeyPrefix)
	defer iterator.Close()

	cycleList := make([]*types.Cycle, 0)
	for ; iterator.Valid(); iterator.Next() {
		cycle := types.Cycle{}
		k.cdc.MustUnmarshal(iterator.Value(), &cycle)
		cycleList = append(cycleList, &cycle)
	}

	return cycleList
}

func (k Keeper) SetCycleSeconds(ctx sdk.Context, cycleSeconds *types.CycleSeconds) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.CycleSecondsStoreKey(cycleSeconds.Denom), k.cdc.MustMarshal(cycleSeconds))
}

func (k Keeper) GetCycleSeconds(ctx sdk.Context, denom string) *types.CycleSeconds {
	store := ctx.KVStore(k.storeKey)
	bts := store.Get(types.CycleSecondsStoreKey(denom))
	if bts == nil {
		return &types.CycleSeconds{
			Denom:   denom,
			Version: 0,
			Seconds: types.DefaultCycleSeconds,
		}
	}

	cycleSeconds := types.CycleSeconds{}
	k.cdc.MustUnmarshal(bts, &cycleSeconds)
	return &cycleSeconds
}

func (k Keeper) GetAllCycleSeconds(ctx sdk.Context) []*types.CycleSeconds {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.CycleSecondsStoreKeyPrefix)
	defer iterator.Close()

	cycleSecondsList := make([]*types.CycleSeconds, 0)
	for ; iterator.Valid(); iterator.Next() {
		cycleSeconds := types.CycleSeconds{}
		k.cdc.MustUnmarshal(iterator.Value(), &cycleSeconds)
		cycleSecondsList = append(cycleSecondsList, &cycleSeconds)
	}
	return cycleSecondsList
}

func (k Keeper) SetShuffleSeconds(ctx sdk.Context, shuffleSeconds *types.ShuffleSeconds) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.ShuffleSecondsStoreKey(shuffleSeconds.Denom), k.cdc.MustMarshal(shuffleSeconds))
}

func (k Keeper) GetShuffleSeconds(ctx sdk.Context, denom string) *types.ShuffleSeconds {
	store := ctx.KVStore(k.storeKey)
	bts := store.Get(types.ShuffleSecondsStoreKey(denom))
	if bts == nil {
		return &types.ShuffleSeconds{
			Denom:   denom,
			Version: 0,
			Seconds: types.DefaultShuffleSeconds,
		}
	}
	shuffleSeconds := types.ShuffleSeconds{}

	k.cdc.MustUnmarshal(bts, &shuffleSeconds)

	return &shuffleSeconds
}

func (k Keeper) GetAllShuffleSeconds(ctx sdk.Context) []*types.ShuffleSeconds {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.ShuffleSecondsStoreKeyPrefix)
	defer iterator.Close()

	shuffleSecondsList := make([]*types.ShuffleSeconds, 0)
	for ; iterator.Valid(); iterator.Next() {
		shuffleSeconds := types.ShuffleSeconds{}
		k.cdc.MustUnmarshal(iterator.Value(), &shuffleSeconds)
		shuffleSecondsList = append(shuffleSecondsList, &shuffleSeconds)
	}

	return shuffleSecondsList
}

func (k Keeper) SetDealingZValidator(ctx sdk.Context, zValidator *types.DealingZValidator) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.DealingZValidatorStoreKey(zValidator.Denom, zValidator.PoolAddress), k.cdc.MustMarshal(zValidator))
}

func (k Keeper) RemoveDealingZValidator(ctx sdk.Context, denom, poolAddress string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.DealingZValidatorStoreKey(denom, poolAddress))
}

func (k Keeper) GetDealingZValidator(ctx sdk.Context, denom, poolAddress string) (*types.DealingZValidator, bool) {
	store := ctx.KVStore(k.storeKey)
	bts := store.Get(types.DealingZValidatorStoreKey(denom, poolAddress))
	if bts == nil {
		return nil, false
	}
	zvalidator := types.DealingZValidator{}
	k.cdc.MustUnmarshal(bts, &zvalidator)

	return &zvalidator, true
}

func (k Keeper) GetAllDealingZvalidators(ctx sdk.Context) []*types.DealingZValidator {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.DealingZValidatorStoreKeyPrefix)
	defer iterator.Close()

	list := make([]*types.DealingZValidator, 0)
	for ; iterator.Valid(); iterator.Next() {
		dealingZvalidator := types.DealingZValidator{}
		k.cdc.MustUnmarshal(iterator.Value(), &dealingZvalidator)
		list = append(list, &dealingZvalidator)
	}
	return list
}
