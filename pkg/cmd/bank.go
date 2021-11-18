package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	BankCmd.PersistentFlags().StringVar(&clientURL, "client-url", "http://127.0.0.1:7545", "etheum client url, default to be ")
	BankCmd.PersistentFlags().StringVar(&privateKey, "private-key", "", "private key in hex format")
	BankCmd.MarkFlagRequired("private-key")
	BankCmd.PersistentFlags().StringVar(&contractAddress, "contract-address", "", "contract address in hex format")
	BankCmd.MarkFlagRequired("contract-address")
	// BankCmd.PersistentFlags().Int64Var(&value, "value", 0, "value to transfer along with the transaction")
	BankCmd.PersistentFlags().Int64Var(&gasPrice, "gas-price", 1000000, "gas price for the transaction")
	BankCmd.PersistentFlags().Uint64Var(&gasLimit, "gas-limit", 3000000, "gas limit for the transaction")
	BankCmd.AddCommand(FundCmd)
	BankCmd.AddCommand(TransferCmd)
	BankCmd.AddCommand(BalanceCmd)
}

var BankCmd = &cobra.Command{
	Use:   "bank",
	Short: "bank --client-url [CLIENT_URL] --private-key [PRIVATE_KEY] --contract-address [CONTRACT_ADDRESS]",
}
