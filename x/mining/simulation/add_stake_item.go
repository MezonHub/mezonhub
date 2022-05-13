package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/stafihub/stafihub/x/mining/keeper"
	"github.com/stafihub/stafihub/x/mining/types"
)

func SimulateMsgAddStakeItem(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgAddStakeItem{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the AddStakeItem simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "AddStakeItem simulation not implemented"), nil, nil
	}
}
