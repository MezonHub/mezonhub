package cli

import (
	"strconv"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"github.com/mezonhub/mezonhub/x/zvalidator/types"
)

var _ = strconv.Itoa(0)

func CmdInitZValidator() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "init-r-validator [denom] [pool-address] [address-list]",
		Short: "Init zvalidator",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argDenom := args[0]
			argPoolAddress := args[1]
			argAddressList := strings.Split(args[2], ":")

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgInitZValidator(
				clientCtx.GetFromAddress().String(),
				argDenom,
				argPoolAddress,
				argAddressList,
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
