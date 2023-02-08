package cli

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	// "github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/mezonhub/mezonhub/x/rstaking/types"
)

var (
	DefaultRelativePacketTimeoutTimestamp = uint64((time.Duration(10) * time.Minute).Nanoseconds())
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdSetInflationBase())
	cmd.AddCommand(CmdAddValToWhitelist())
	cmd.AddCommand(CmdToggleValidatorWhitelistSwitch())
	cmd.AddCommand(CmdWithdraw())
	cmd.AddCommand(CmdAddDelegatorToWhitelist())
	cmd.AddCommand(CmdToggleDelegatorWhitelistSwitch())
	cmd.AddCommand(CmdProvideToken())
	cmd.AddCommand(CmdRmDelegatorFromWhitelist())
	// this line is used by starport scaffolding # 1

	return cmd
}
