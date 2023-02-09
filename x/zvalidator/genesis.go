package zvalidator

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/mezonhub/mezonhub/x/zvalidator/keeper"
	"github.com/mezonhub/mezonhub/x/zvalidator/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
	for _, cycleSeconds := range genState.CycleSecondsList {
		k.SetCycleSeconds(ctx, cycleSeconds)
	}
	for _, latestDealedCycle := range genState.LatestDealedCycleList {
		k.SetLatestDealedCycle(ctx, latestDealedCycle)
	}

	for _, latestVotedCycle := range genState.LatestVotedCycleList {
		k.SetLatestVotedCycle(ctx, latestVotedCycle)
	}
	for _, selectedZValidator := range genState.SelectedZValidatorList {
		if err := k.ZBankKeeper.CheckValAddress(ctx, selectedZValidator.Denom, selectedZValidator.ValAddress); err != nil {
			panic(err)
		}
		if err := k.ZBankKeeper.CheckAccAddress(ctx, selectedZValidator.Denom, selectedZValidator.PoolAddress); err != nil {
			panic(err)
		}
		k.AddSelectedZValidator(ctx, selectedZValidator)
	}

	for _, shuffleSeconds := range genState.ShuffleSecondsList {
		k.SetShuffleSeconds(ctx, shuffleSeconds)
	}

	for _, dealingZValidator := range genState.DealingZValidatorList {
		if err := k.ZBankKeeper.CheckValAddress(ctx, dealingZValidator.Denom, dealingZValidator.NewValAddress); err != nil {
			panic(err)
		}
		if err := k.ZBankKeeper.CheckValAddress(ctx, dealingZValidator.Denom, dealingZValidator.OldValAddress); err != nil {
			panic(err)
		}
		if err := k.ZBankKeeper.CheckAccAddress(ctx, dealingZValidator.Denom, dealingZValidator.PoolAddress); err != nil {
			panic(err)
		}
		k.SetDealingZValidator(ctx, dealingZValidator)
	}
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)
	genesis.CycleSecondsList = k.GetAllCycleSeconds(ctx)
	genesis.LatestDealedCycleList = k.GetAllLatestDealedCycle(ctx)
	genesis.LatestVotedCycleList = k.GetAllLatestVotedCycle(ctx)
	genesis.SelectedZValidatorList = k.GetSelectedZValidatorList(ctx)
	genesis.ShuffleSecondsList = k.GetAllShuffleSeconds(ctx)
	genesis.DealingZValidatorList = k.GetAllDealingZvalidators(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
