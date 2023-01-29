package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
	"lotterychain/x/lottery/types"
)

var _ = strconv.Itoa(0)

func CmdEnterLottery() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "enter-lottery [fee] [bet]",
		Short: "Broadcast message enterLottery",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argFee, err := sdk.ParseCoinNormalized(args[0])
			if err != nil {
				return err
			}
			argBet, err := sdk.ParseCoinNormalized(args[1])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgEnterLottery(
				clientCtx.GetFromAddress().String(),
				argFee,
				argBet,
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
