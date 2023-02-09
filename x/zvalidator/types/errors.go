package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/zvalidator module sentinel errors
var (
	ErrZValidatorAlreadyExist              = sdkerrors.Register(ModuleName, 1101, "zValidator already exist")
	ErrZValidatorNotExist                  = sdkerrors.Register(ModuleName, 1102, "zValidator not exist")
	ErrCycleBehindLatestCycle              = sdkerrors.Register(ModuleName, 1103, "cycle behind latest voted cycle")
	ErrCycleVersionNotMatch                = sdkerrors.Register(ModuleName, 1104, "cycle version not match")
	ErrLatestVotedCycleNotDealed           = sdkerrors.Register(ModuleName, 1105, "latest voted cycle not dealed")
	ErrLedgerIsBusyWithEra                 = sdkerrors.Register(ModuleName, 1106, "ledger is busy with era")
	ErrReportCycleNotMatchLatestVotedCycle = sdkerrors.Register(ModuleName, 1107, "report cycle not match latest voted cycle")
	ErrLedgerChainEraNotExist              = sdkerrors.Register(ModuleName, 1108, "ledger chain era not exist")
	ErrDealingZvalidatorNotFound           = sdkerrors.Register(ModuleName, 1109, "dealing zvalidator not found")
	ErrOldEqualNewZValidator               = sdkerrors.Register(ModuleName, 1110, "old euqal new zValidator")
	ErrZValidatorAlreadyInit               = sdkerrors.Register(ModuleName, 1111, "zValidator already init")
)
