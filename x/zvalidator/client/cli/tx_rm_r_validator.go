package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"github.com/mezonhub/mezonhub/x/zvalidator/types"
)

var _ = strconv.Itoa(0)

func CmdRmZValidator() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "rm-r-validator [denom] [pool-address] [old-address] [new-address]",
		Short: "Remove zvalidator",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argDenom := args[0]
			argPoolAddress := args[1]
			argOldAddress := args[2]
			argNewAddress := args[3]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgRmZValidator(
				clientCtx.GetFromAddress().String(),
				argDenom,
				argPoolAddress,
				argOldAddress,
				argNewAddress,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
