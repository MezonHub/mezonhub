package cli

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/mezonhub/mezonhub/x/zvalidator/types"
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

	cmd.AddCommand(CmdInitZValidator())
	cmd.AddCommand(CmdUpdateZValidator())
	cmd.AddCommand(CmdSetCycleSeconds())
	cmd.AddCommand(CmdSetShuffleSeconds())
	cmd.AddCommand(CmdUpdateZValidatorReport())
	cmd.AddCommand(CmdAddZValidator())
	cmd.AddCommand(CmdRmZValidator())
	// this line is used by starport scaffolding # 1

	return cmd
}
