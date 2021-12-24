package sudo

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stafiprotocol/stafihub/x/sudo/keeper"
	"github.com/stafiprotocol/stafihub/x/sudo/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
    // this line is used by starport scaffolding # genesis/module/init
	admin, err := sdk.AccAddressFromBech32(genState.Admin)
	if err != nil {
		panic(err)
	}
	k.SetAdmin(ctx, admin)

	for _, dnm := range genState.Denoms {
		k.AddDenom(ctx, dnm)
	}
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Admin = k.GetAdmin(ctx).String()
	genesis.Denoms = k.GetAllDenoms(ctx)

    // this line is used by starport scaffolding # genesis/module/export

    return genesis
}
