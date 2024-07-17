package cli

import (
    "github.com/spf13/cobra"
    "github.com/cosmos/cosmos-sdk/client"
    "github.com/cosmos/cosmos-sdk/client/flags"
    "github.com/cosmos/cosmos-sdk/client/tx"
    "github.com/cosmos/cosmos-sdk/x/erc20/types"
)

func NewTxCmd() *cobra.Command {
    txCmd := &cobra.Command{
        Use:                        types.ModuleName,
        Short:                      "ERC20 transactions subcommands",
        DisableFlagParsing:         true,
        SuggestionsMinimumDistance: 2,
        RunE:                       client.ValidateCmd,
    }

    txCmd.AddCommand(
        NewRegisterTokenPairCmd(),
        NewConvertCoinToERC20Cmd(),
        NewConvertERC20ToCoinCmd(),
    )

    return txCmd
}

func NewRegisterTokenPairCmd() *cobra.Command {
    cmd := &cobra.Command{
        Use:   "register-pair [erc20_address] [denom]",
        Short: "Register a new token pair",
        Args:  cobra.ExactArgs(2),
        RunE: func(cmd *cobra.Command, args []string) error {
            clientCtx, err := client.GetClientTxContext(cmd)
            if err != nil {
                return err
            }

            msg := types.NewMsgRegisterTokenPair(clientCtx.GetFromAddress(), args[0], args[1])
            return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
        },
    }

    flags.AddTxFlagsToCmd(cmd)
    return cmd
}

func NewConvertCoinToERC20Cmd() *cobra.Command {
    cmd := &cobra.Command{
        Use:   "convert-coin [amount]",
        Short: "Convert Cosmos Coin to ERC20",
        Args:  cobra.ExactArgs(1),
        RunE: func(cmd *cobra.Command, args []string) error {
            clientCtx, err := client.GetClientTxContext(cmd)
            if err != nil {
                return err
            }

            amount, err := sdk.ParseCoinNormalized(args[0])
            if err != nil {
                return err
            }

            msg := types.NewMsgConvertCoinToERC20(clientCtx.GetFromAddress(), amount)
            return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
        },
    }

    flags.AddTxFlagsToCmd(cmd)
    return cmd
}

func NewConvertERC20ToCoinCmd() *cobra.Command {
    cmd := &cobra.Command{
        Use:   "convert-erc20 [amount] [denom]",
        Short: "Convert ERC20 to Cosmos Coin",
        Args:  cobra.ExactArgs(2),
        RunE: func(cmd *cobra.Command, args []string) error {
            clientCtx, err := client.GetClientTxContext(cmd)
            if err != nil {
                return err
            }

            amount, ok := sdk.NewIntFromString(args[0])
            if !ok {
                return fmt.Errorf("invalid amount")
            }

            msg := types.NewMsgConvertERC20ToCoin(clientCtx.GetFromAddress(), amount, args[1])
            return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
        },
    }

    flags.AddTxFlagsToCmd(cmd)
    return cmd
}
