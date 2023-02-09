package bank

import (
	"encoding/json"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/mezonhub/mezonhub/utils"
)

var (
	_ module.AppModuleBasic = AppModuleBasic{}
)

// AppModuleBasic defines the basic application module used by the staking module.
type AppModuleBasic struct {
	bank.AppModuleBasic
}

// DefaultGenesis returns default genesis state as raw bytes for the gov
// module.
func (am AppModuleBasic) DefaultGenesis(cdc codec.JSONCodec) json.RawMessage {
	defaultGenesisState := types.DefaultGenesisState()
	// {
	//     "description": "The native staking token of the MeZon Hub",
	//     "denom_units": [
	//         {
	//             "denom": "umez",
	//             "exponent": 0,
	//             "aliases": [
	//                 "micromez"
	//             ]
	//         },
	//         {
	//             "denom": "mmez",
	//             "exponent": 3,
	//             "aliases": [
	//               "millimez"
	//             ]
	//         },
	//         {
	//             "denom": "mez",
	//             "exponent": 6
	//         }
	//     ],
	//     "base": "umez",
	//     "display": "mez",
	//     "name": "MEZ",
	//     "symbol": "MEZ"
	// }
	defaultGenesisState.DenomMetadata = append(defaultGenesisState.DenomMetadata,
		types.Metadata{
			Description: "The native staking token of the MeZon Hub",
			DenomUnits: []*types.DenomUnit{
				{
					Denom:    utils.FisDenom,
					Exponent: 0,
					Aliases:  []string{"micromez"},
				},
				{
					Denom:    "mmez",
					Exponent: 3,
					Aliases:  []string{"millimez"},
				},
				{
					Denom:    "mez",
					Exponent: 6,
				},
			},
			Base:    utils.FisDenom,
			Display: "mez",
			Name:    "MEZ",
			Symbol:  "MEZ",
		},
	)
	return cdc.MustMarshalJSON(defaultGenesisState)
}
